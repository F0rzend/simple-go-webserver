package userentity

import (
	"net/mail"
	"testing"
	"time"

	"github.com/F0rzend/simple-go-webserver/app/common"

	"github.com/F0rzend/simple-go-webserver/app/tests"

	"github.com/stretchr/testify/require"

	"github.com/stretchr/testify/assert"

	"github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/entity"
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
		checkErr tests.ErrorChecker
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
				Balance:   Balance{USD: bitcoinentity.NewUSD(0), BTC: bitcoinentity.NewBTC(0)},
				CreatedAt: now,
				UpdatedAt: now,
			},
			checkErr: assert.NoError,
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
			checkErr: tests.AssertErrorFlag(common.FlagInvalidArgument),
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
			checkErr: tests.AssertErrorFlag(common.FlagInvalidArgument),
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
			checkErr: tests.AssertErrorFlag(common.FlagInvalidArgument),
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

			expected: &User{
				ID:        1,
				Name:      "John Doe",
				Username:  "johndoe",
				Email:     &mail.Address{Name: "", Address: "johndoe@gmail.com"},
				Balance:   Balance{USD: bitcoinentity.NewUSD(0), BTC: bitcoinentity.NewBTC(-1)},
				CreatedAt: now,
				UpdatedAt: now,
			},
			checkErr: assert.NoError,
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

			expected: &User{
				ID:        1,
				Name:      "John Doe",
				Username:  "johndoe",
				Email:     &mail.Address{Name: "", Address: "johndoe@gmail.com"},
				Balance:   Balance{USD: bitcoinentity.NewUSD(-1), BTC: bitcoinentity.NewBTC(0)},
				CreatedAt: now,
				UpdatedAt: now,
			},
			checkErr: assert.NoError,
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

			tc.checkErr(t, err)
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
		amount   bitcoinentity.USD
		expected Balance
		checkErr tests.ErrorChecker
	}{
		{
			name: "success deposit",
			user: User{
				Balance: Balance{
					USD: bitcoinentity.NewUSD(0),
					BTC: bitcoinentity.NewBTC(0),
				},
			},
			action: DepositUSDAction,
			amount: bitcoinentity.NewUSD(1),
			expected: Balance{
				USD: bitcoinentity.NewUSD(1),
				BTC: bitcoinentity.NewBTC(0),
			},
			checkErr: assert.NoError,
		},
		{
			name: "success withdraw",
			user: User{
				Balance: Balance{
					USD: bitcoinentity.NewUSD(1),
					BTC: bitcoinentity.NewBTC(0),
				},
			},
			action: WithdrawUSDAction,
			amount: bitcoinentity.NewUSD(1),
			expected: Balance{
				USD: bitcoinentity.NewUSD(0),
				BTC: bitcoinentity.NewBTC(0),
			},
			checkErr: assert.NoError,
		},
		{
			name: "insufficient funds",
			user: User{
				Balance: Balance{
					USD: bitcoinentity.NewUSD(0),
					BTC: bitcoinentity.NewBTC(0),
				},
			},
			action: WithdrawUSDAction,
			amount: bitcoinentity.NewUSD(1),
			expected: Balance{
				USD: bitcoinentity.NewUSD(0),
				BTC: bitcoinentity.NewBTC(0),
			},
			checkErr: tests.AssertErrorFlag(common.FlagInvalidArgument),
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			user := &tc.user
			err := user.ChangeUSDBalance(tc.action, tc.amount)

			tc.checkErr(t, err)
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
		amount   bitcoinentity.BTC
		btcPrice bitcoinentity.USD
		expected Balance
		checkErr tests.ErrorChecker
	}{
		{
			name: "success buy",
			user: User{
				Balance: Balance{
					USD: bitcoinentity.NewUSD(1),
					BTC: bitcoinentity.NewBTC(0),
				},
			},
			action:   BuyBTCAction,
			amount:   bitcoinentity.NewBTC(1),
			btcPrice: bitcoinentity.NewUSD(1),
			expected: Balance{
				USD: bitcoinentity.NewUSD(0),
				BTC: bitcoinentity.NewBTC(1),
			},
			checkErr: assert.NoError,
		},
		{
			name: "success sale",
			user: User{
				Balance: Balance{
					USD: bitcoinentity.NewUSD(0),
					BTC: bitcoinentity.NewBTC(1),
				},
			},
			action:   SellBTCAction,
			amount:   bitcoinentity.NewBTC(1),
			btcPrice: bitcoinentity.NewUSD(1),
			expected: Balance{
				USD: bitcoinentity.NewUSD(1),
				BTC: bitcoinentity.NewBTC(0),
			},
			checkErr: assert.NoError,
		},
		{
			name: "insufficient funds on buy",
			user: User{
				Balance: Balance{
					USD: bitcoinentity.NewUSD(0),
					BTC: bitcoinentity.NewBTC(0),
				},
			},
			action:   BuyBTCAction,
			amount:   bitcoinentity.NewBTC(1),
			btcPrice: bitcoinentity.NewUSD(1),
			expected: Balance{
				USD: bitcoinentity.NewUSD(0),
				BTC: bitcoinentity.NewBTC(0),
			},
			checkErr: tests.AssertErrorFlag(common.FlagInvalidArgument),
		},
		{
			name: "insufficient funds on sell",
			user: User{
				Balance: Balance{
					USD: bitcoinentity.NewUSD(0),
					BTC: bitcoinentity.NewBTC(0),
				},
			},
			action:   SellBTCAction,
			amount:   bitcoinentity.NewBTC(1),
			btcPrice: bitcoinentity.NewUSD(1),
			expected: Balance{
				USD: bitcoinentity.NewUSD(0),
				BTC: bitcoinentity.NewBTC(0),
			},
			checkErr: tests.AssertErrorFlag(common.FlagInvalidArgument),
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			user := &tc.user
			price, err := bitcoinentity.NewBTCPrice(tc.btcPrice, now)
			require.NoError(t, err)

			err = user.ChangeBTCBalance(tc.action, tc.amount, price)

			tc.checkErr(t, err)
			assert.True(t, tc.expected.BTC.Equal(user.Balance.BTC))
			assert.True(t, tc.expected.USD.Equal(user.Balance.USD))
		})
	}
}

func TestUser_ParseEmail(t *testing.T) {
	t.Parallel()

	testCases := []struct {
		name     string
		email    string
		checkErr tests.ErrorChecker
	}{
		{
			name:     "success",
			email:    "test@mail.com",
			checkErr: assert.NoError,
		},
		{
			name:     "invalid mail",
			email:    "test",
			checkErr: tests.AssertErrorFlag(common.FlagInvalidArgument),
		},
	}

	for _, tc := range testCases {
		tc := tc
		t.Run(tc.name, func(t *testing.T) {
			t.Parallel()

			_, err := ParseEmail(tc.email)

			tc.checkErr(t, err)
		})
	}
}
