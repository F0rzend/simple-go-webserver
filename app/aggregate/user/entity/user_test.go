package entity

import (
	"net/mail"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"

	bitcoinEntity "github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/entity"
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
				Balance:   Balance{USD: bitcoinEntity.MustNewUSD(0), BTC: bitcoinEntity.MustNewBTC(0)},
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
			err:      ErrInvalidEmail,
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
			err:      bitcoinEntity.ErrNegativeCurrency,
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
			err:      bitcoinEntity.ErrNegativeCurrency,
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
		action   Action
		amount   bitcoinEntity.USD
		expected Balance
		err      error
	}{
		{
			name: "success deposit",
			user: User{
				Balance: Balance{
					USD: bitcoinEntity.MustNewUSD(0),
					BTC: bitcoinEntity.MustNewBTC(0),
				},
			},
			action: DepositUSDAction,
			amount: bitcoinEntity.MustNewUSD(1),
			expected: Balance{
				USD: bitcoinEntity.MustNewUSD(1),
				BTC: bitcoinEntity.MustNewBTC(0),
			},
			err: nil,
		},
		{
			name: "success withdraw",
			user: User{
				Balance: Balance{
					USD: bitcoinEntity.MustNewUSD(1),
					BTC: bitcoinEntity.MustNewBTC(0),
				},
			},
			action: WithdrawUSDAction,
			amount: bitcoinEntity.MustNewUSD(1),
			expected: Balance{
				USD: bitcoinEntity.MustNewUSD(0),
				BTC: bitcoinEntity.MustNewBTC(0),
			},
			err: nil,
		},
		{
			name: "insufficient funds",
			user: User{
				Balance: Balance{
					USD: bitcoinEntity.MustNewUSD(0),
					BTC: bitcoinEntity.MustNewBTC(0),
				},
			},
			action: WithdrawUSDAction,
			amount: bitcoinEntity.MustNewUSD(1),
			expected: Balance{
				USD: bitcoinEntity.MustNewUSD(0),
				BTC: bitcoinEntity.MustNewBTC(0),
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

	now := time.Now()

	testCases := []struct {
		name     string
		user     User
		action   Action
		amount   bitcoinEntity.BTC
		btcPrice bitcoinEntity.BTCPrice
		expected Balance
		err      error
	}{
		{
			name: "success buy",
			user: User{
				Balance: Balance{
					USD: bitcoinEntity.MustNewUSD(1),
					BTC: bitcoinEntity.MustNewBTC(0),
				},
			},
			action:   BuyBTCAction,
			amount:   bitcoinEntity.MustNewBTC(1),
			btcPrice: bitcoinEntity.NewBTCPrice(bitcoinEntity.MustNewUSD(1), now),
			expected: Balance{
				USD: bitcoinEntity.MustNewUSD(0),
				BTC: bitcoinEntity.MustNewBTC(1),
			},
			err: nil,
		},
		{
			name: "success sale",
			user: User{
				Balance: Balance{
					USD: bitcoinEntity.MustNewUSD(0),
					BTC: bitcoinEntity.MustNewBTC(1),
				},
			},
			action:   SellBTCAction,
			amount:   bitcoinEntity.MustNewBTC(1),
			btcPrice: bitcoinEntity.NewBTCPrice(bitcoinEntity.MustNewUSD(1), now),
			expected: Balance{
				USD: bitcoinEntity.MustNewUSD(1),
				BTC: bitcoinEntity.MustNewBTC(0),
			},
			err: nil,
		},
		{
			name: "insufficient funds on buy",
			user: User{
				Balance: Balance{
					USD: bitcoinEntity.MustNewUSD(0),
					BTC: bitcoinEntity.MustNewBTC(0),
				},
			},
			action:   BuyBTCAction,
			amount:   bitcoinEntity.MustNewBTC(1),
			btcPrice: bitcoinEntity.NewBTCPrice(bitcoinEntity.MustNewUSD(1), now),
			expected: Balance{
				USD: bitcoinEntity.MustNewUSD(0),
				BTC: bitcoinEntity.MustNewBTC(0),
			},
			err: ErrInsufficientFunds,
		},
		{
			name: "insufficient funds on sell",
			user: User{
				Balance: Balance{
					USD: bitcoinEntity.MustNewUSD(0),
					BTC: bitcoinEntity.MustNewBTC(0),
				},
			},
			action:   SellBTCAction,
			amount:   bitcoinEntity.MustNewBTC(1),
			btcPrice: bitcoinEntity.NewBTCPrice(bitcoinEntity.MustNewUSD(1), now),
			expected: Balance{
				USD: bitcoinEntity.MustNewUSD(0),
				BTC: bitcoinEntity.MustNewBTC(0),
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

func TestUser_ParseEmail(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name  string
		email string
		err   error
	}{
		{
			name:  "success",
			email: "test@mail.com",
			err:   nil,
		},
		{
			name:  "invalid mail",
			email: "test",
			err:   ErrInvalidEmail,
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			_, err := ParseEmail(tc.email)

			assert.ErrorIs(t, err, tc.err)
		})
	}
}
