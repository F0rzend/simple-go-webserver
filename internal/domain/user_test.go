package domain

import (
	"errors"
	"net/mail"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
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
				Balance:   Balance{USD: MustNewUSD(0), BTC: MustNewBTC(0)},
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
			err:      ErrNegativeBTC,
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
			err:      ErrNegativeUSD,
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
		action   USDAction
		amount   USD
		expected Balance
		err      error
	}{
		{
			name: "success deposit",
			user: User{
				Balance: Balance{
					USD: MustNewUSD(0),
					BTC: MustNewBTC(0),
				},
			},
			action: DepositUSDAction,
			amount: MustNewUSD(1),
			expected: Balance{
				USD: MustNewUSD(1),
				BTC: MustNewBTC(0),
			},
			err: nil,
		},
		{
			name: "success withdraw",
			user: User{
				Balance: Balance{
					USD: MustNewUSD(1),
					BTC: MustNewBTC(0),
				},
			},
			action: WithdrawUSDAction,
			amount: MustNewUSD(1),
			expected: Balance{
				USD: MustNewUSD(0),
				BTC: MustNewBTC(0),
			},
			err: nil,
		},
		{
			name: "insufficient funds",
			user: User{
				Balance: Balance{
					USD: MustNewUSD(0),
					BTC: MustNewBTC(0),
				},
			},
			action: WithdrawUSDAction,
			amount: MustNewUSD(1),
			expected: Balance{
				USD: MustNewUSD(0),
				BTC: MustNewBTC(0),
			},
			err: ErrInsufficientFunds,
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
		action   BTCAction
		amount   BTC
		btcPrice BTCPrice
		expected Balance
		err      error
	}{
		{
			name: "success buy",
			user: User{
				Balance: Balance{
					USD: MustNewUSD(1),
					BTC: MustNewBTC(0),
				},
			},
			action:   BuyBTCAction,
			amount:   MustNewBTC(1),
			btcPrice: NewBTCPrice(MustNewUSD(1)),
			expected: Balance{
				USD: MustNewUSD(0),
				BTC: MustNewBTC(1),
			},
			err: nil,
		},
		{
			name: "success sale",
			user: User{
				Balance: Balance{
					USD: MustNewUSD(0),
					BTC: MustNewBTC(1),
				},
			},
			action:   SellBTCAction,
			amount:   MustNewBTC(1),
			btcPrice: NewBTCPrice(MustNewUSD(1)),
			expected: Balance{
				USD: MustNewUSD(1),
				BTC: MustNewBTC(0),
			},
			err: nil,
		},
		{
			name: "insufficient funds on buy",
			user: User{
				Balance: Balance{
					USD: MustNewUSD(0),
					BTC: MustNewBTC(0),
				},
			},
			action:   BuyBTCAction,
			amount:   MustNewBTC(1),
			btcPrice: NewBTCPrice(MustNewUSD(1)),
			expected: Balance{
				USD: MustNewUSD(0),
				BTC: MustNewBTC(0),
			},
			err: ErrInsufficientFunds,
		},
		{
			name: "insufficient funds on sell",
			user: User{
				Balance: Balance{
					USD: MustNewUSD(0),
					BTC: MustNewBTC(0),
				},
			},
			action:   SellBTCAction,
			amount:   MustNewBTC(1),
			btcPrice: NewBTCPrice(MustNewUSD(1)),
			expected: Balance{
				USD: MustNewUSD(0),
				BTC: MustNewBTC(0),
			},
			err: ErrInsufficientFunds,
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
