// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package bitcoinhandlers

import (
	"github.com/F0rzend/simple-go-webserver/app/aggregate/bitcoin/entity"
	"sync"
)

// Ensure, that MockBitcoinService does implement BitcoinService.
// If this is not the case, regenerate this file with moq.
var _ BitcoinService = &MockBitcoinService{}

// MockBitcoinService is a mock implementation of BitcoinService.
//
// 	func TestSomethingThatUsesBitcoinService(t *testing.T) {
//
// 		// make and configure a mocked BitcoinService
// 		mockedBitcoinService := &MockBitcoinService{
// 			GetBTCPriceFunc: func() bitcoinentity.BTCPrice {
// 				panic("mock out the GetBTCPrice method")
// 			},
// 			SetBTCPriceFunc: func(newPrice float64) error {
// 				panic("mock out the SetBTCPrice method")
// 			},
// 		}
//
// 		// use mockedBitcoinService in code that requires BitcoinService
// 		// and then make assertions.
//
// 	}
type MockBitcoinService struct {
	// GetBTCPriceFunc mocks the GetBTCPrice method.
	GetBTCPriceFunc func() bitcoinentity.BTCPrice

	// SetBTCPriceFunc mocks the SetBTCPrice method.
	SetBTCPriceFunc func(newPrice float64) error

	// calls tracks calls to the methods.
	calls struct {
		// GetBTCPrice holds details about calls to the GetBTCPrice method.
		GetBTCPrice []struct {
		}
		// SetBTCPrice holds details about calls to the SetBTCPrice method.
		SetBTCPrice []struct {
			// NewPrice is the newPrice argument value.
			NewPrice float64
		}
	}
	lockGetBTCPrice sync.RWMutex
	lockSetBTCPrice sync.RWMutex
}

// GetBTCPrice calls GetBTCPriceFunc.
func (mock *MockBitcoinService) GetBTCPrice() bitcoinentity.BTCPrice {
	if mock.GetBTCPriceFunc == nil {
		panic("MockBitcoinService.GetBTCPriceFunc: method is nil but BitcoinService.GetBTCPrice was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGetBTCPrice.Lock()
	mock.calls.GetBTCPrice = append(mock.calls.GetBTCPrice, callInfo)
	mock.lockGetBTCPrice.Unlock()
	return mock.GetBTCPriceFunc()
}

// GetBTCPriceCalls gets all the calls that were made to GetBTCPrice.
// Check the length with:
//     len(mockedBitcoinService.GetBTCPriceCalls())
func (mock *MockBitcoinService) GetBTCPriceCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGetBTCPrice.RLock()
	calls = mock.calls.GetBTCPrice
	mock.lockGetBTCPrice.RUnlock()
	return calls
}

// SetBTCPrice calls SetBTCPriceFunc.
func (mock *MockBitcoinService) SetBTCPrice(newPrice float64) error {
	if mock.SetBTCPriceFunc == nil {
		panic("MockBitcoinService.SetBTCPriceFunc: method is nil but BitcoinService.SetBTCPrice was just called")
	}
	callInfo := struct {
		NewPrice float64
	}{
		NewPrice: newPrice,
	}
	mock.lockSetBTCPrice.Lock()
	mock.calls.SetBTCPrice = append(mock.calls.SetBTCPrice, callInfo)
	mock.lockSetBTCPrice.Unlock()
	return mock.SetBTCPriceFunc(newPrice)
}

// SetBTCPriceCalls gets all the calls that were made to SetBTCPrice.
// Check the length with:
//     len(mockedBitcoinService.SetBTCPriceCalls())
func (mock *MockBitcoinService) SetBTCPriceCalls() []struct {
	NewPrice float64
} {
	var calls []struct {
		NewPrice float64
	}
	mock.lockSetBTCPrice.RLock()
	calls = mock.calls.SetBTCPrice
	mock.lockSetBTCPrice.RUnlock()
	return calls
}
