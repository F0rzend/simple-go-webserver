package domain

import (
	"errors"
	"net/mail"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	bitcoinDomain "github.com/F0rzend/simple-go-webserver/app/bitcoin/domain"
)

func TestNewUser(t *testing.T) {
	t.Parallel()

	now := time.Now()

	testCases := []struct {
		testName string

		id            uint64
		name          string
		username      string
		email         string
		bitcoinAmount float64
		usdBalance    float64
		createdAt     time.Time
		updatedAt     time.Time

		expected *User
		err      error
	}{
		{
			testName: "success",

			id:            1,
			name:          "John Doe",
			username:      "johndoe",
			email:         "johndoe@gmail.com",
			bitcoinAmount: 0.0,
			usdBalance:    0.0,
			createdAt:     now,
			updatedAt:     now,

			expected: &User{
				ID:        1,
				Name:      "John Doe",
				Username:  "johndoe",
				Email:     &mail.Address{Name: "", Address: "johndoe@gmail.com"},
				Balance:   Balance{USD: bitcoinDomain.MustNewUSD(0), BTC: bitcoinDomain.MustNewBTC(0)},
				CreatedAt: now,
				UpdatedAt: now,
			},
			err: nil,
		},
		{
			testName: "wrong email",

			id:            1,
			name:          "John Doe",
			username:      "johndoe",
			email:         "johndoe",
			bitcoinAmount: 0.0,
			usdBalance:    0.0,
			createdAt:     now,
			updatedAt:     now,

			expected: nil,
			err:      errors.New("mail: missing '@' or angle-addr"),
		},
		{
			testName: "empty name",

			id:            1,
			name:          "",
			username:      "johndoe",
			email:         "johndoe@gmail.com",
			bitcoinAmount: 0.0,
			usdBalance:    0.0,
			createdAt:     now,
			updatedAt:     now,

			expected: nil,
			err:      ErrNameEmpty,
		},
		{
			testName: "empty username",

			id:            1,
			name:          "John Doe",
			username:      "",
			email:         "johndoe@gmail.com",
			bitcoinAmount: 0.0,
			usdBalance:    0.0,
			createdAt:     now,
			updatedAt:     now,

			expected: nil,
			err:      ErrUsernameEmpty,
		},
		{
			testName: "small btc amount",

			id:            1,
			name:          "John Doe",
			username:      "johndoe",
			email:         "johndoe@gmail.com",
			bitcoinAmount: -1,
			usdBalance:    0.0,
			createdAt:     now,
			updatedAt:     now,

			expected: nil,
			err:      bitcoinDomain.ErrNegativeCurrency(-1),
		},
		{
			testName: "small usd amount",

			id:            1,
			name:          "John Doe",
			username:      "johndoe",
			email:         "johndoe@gmail.com",
			bitcoinAmount: 0.0,
			usdBalance:    -1,
			createdAt:     now,
			updatedAt:     now,

			expected: nil,
			err:      bitcoinDomain.ErrNegativeCurrency(-1),
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.testName, func(t *testing.T) {
			t.Parallel()

			user, err := NewUser(
				tc.id,
				tc.name,
				tc.username,
				tc.email,
				tc.bitcoinAmount,
				tc.usdBalance,
				tc.createdAt,
				tc.updatedAt,
			)

			assert.Equal(t, err, tc.err)
			assert.Equal(t, tc.expected, user)
		})
	}
}

func TestUser_ChangeUSDBalance(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		user     User
		action   bitcoinDomain.USDAction
		amount   bitcoinDomain.USD
		expected Balance
		err      error
	}{
		{
			name: "success deposit",
			user: User{
				Balance: Balance{
					USD: bitcoinDomain.MustNewUSD(0),
					BTC: bitcoinDomain.MustNewBTC(0),
				},
			},
			action: bitcoinDomain.DepositUSDAction,
			amount: bitcoinDomain.MustNewUSD(1),
			expected: Balance{
				USD: bitcoinDomain.MustNewUSD(1),
				BTC: bitcoinDomain.MustNewBTC(0),
			},
			err: nil,
		},
		{
			name: "success withdraw",
			user: User{
				Balance: Balance{
					USD: bitcoinDomain.MustNewUSD(1),
					BTC: bitcoinDomain.MustNewBTC(0),
				},
			},
			action: bitcoinDomain.WithdrawUSDAction,
			amount: bitcoinDomain.MustNewUSD(1),
			expected: Balance{
				USD: bitcoinDomain.MustNewUSD(0),
				BTC: bitcoinDomain.MustNewBTC(0),
			},
			err: nil,
		},
		{
			name: "insufficient funds",
			user: User{
				Balance: Balance{
					USD: bitcoinDomain.MustNewUSD(0),
					BTC: bitcoinDomain.MustNewBTC(0),
				},
			},
			action: bitcoinDomain.WithdrawUSDAction,
			amount: bitcoinDomain.MustNewUSD(1),
			expected: Balance{
				USD: bitcoinDomain.MustNewUSD(0),
				BTC: bitcoinDomain.MustNewBTC(0),
			},
			err: ErrInsufficientFunds(1),
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			user := &tc.user
			err := user.ChangeUSDBalance(tc.action, tc.amount)

			assert.ErrorIs(t, err, tc.err)
			assert.True(t, tc.expected.BTC.Equal(user.Balance.BTC))
			assert.True(t, tc.expected.USD.Equal(user.Balance.USD))
		})
	}
}

func TestUser_ChangeBTCBalance(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		user     User
		action   bitcoinDomain.BTCAction
		amount   bitcoinDomain.BTC
		btcPrice bitcoinDomain.BTCPrice
		expected Balance
		err      error
	}{
		{
			name: "success buy",
			user: User{
				Balance: Balance{
					USD: bitcoinDomain.MustNewUSD(1),
					BTC: bitcoinDomain.MustNewBTC(0),
				},
			},
			action:   bitcoinDomain.BuyBTCAction,
			amount:   bitcoinDomain.MustNewBTC(1),
			btcPrice: bitcoinDomain.NewBTCPrice(bitcoinDomain.MustNewUSD(1)),
			expected: Balance{
				USD: bitcoinDomain.MustNewUSD(0),
				BTC: bitcoinDomain.MustNewBTC(1),
			},
			err: nil,
		},
		{
			name: "success sale",
			user: User{
				Balance: Balance{
					USD: bitcoinDomain.MustNewUSD(0),
					BTC: bitcoinDomain.MustNewBTC(1),
				},
			},
			action:   bitcoinDomain.SellBTCAction,
			amount:   bitcoinDomain.MustNewBTC(1),
			btcPrice: bitcoinDomain.NewBTCPrice(bitcoinDomain.MustNewUSD(1)),
			expected: Balance{
				USD: bitcoinDomain.MustNewUSD(1),
				BTC: bitcoinDomain.MustNewBTC(0),
			},
			err: nil,
		},
		{
			name: "insufficient funds on buy",
			user: User{
				Balance: Balance{
					USD: bitcoinDomain.MustNewUSD(0),
					BTC: bitcoinDomain.MustNewBTC(0),
				},
			},
			action:   bitcoinDomain.BuyBTCAction,
			amount:   bitcoinDomain.MustNewBTC(1),
			btcPrice: bitcoinDomain.NewBTCPrice(bitcoinDomain.MustNewUSD(1)),
			expected: Balance{
				USD: bitcoinDomain.MustNewUSD(0),
				BTC: bitcoinDomain.MustNewBTC(0),
			},
			err: ErrInsufficientFunds(1),
		},
		{
			name: "insufficient funds on sell",
			user: User{
				Balance: Balance{
					USD: bitcoinDomain.MustNewUSD(0),
					BTC: bitcoinDomain.MustNewBTC(0),
				},
			},
			action:   bitcoinDomain.SellBTCAction,
			amount:   bitcoinDomain.MustNewBTC(1),
			btcPrice: bitcoinDomain.NewBTCPrice(bitcoinDomain.MustNewUSD(1)),
			expected: Balance{
				USD: bitcoinDomain.MustNewUSD(0),
				BTC: bitcoinDomain.MustNewBTC(0),
			},
			err: ErrInsufficientFunds(1),
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			user := &tc.user
			err := user.ChangeBTCBalance(tc.action, tc.amount, tc.btcPrice)

			assert.ErrorIs(t, err, tc.err)
			assert.True(t, tc.expected.BTC.Equal(user.Balance.BTC))
			assert.True(t, tc.expected.USD.Equal(user.Balance.USD))
		})
	}
}
