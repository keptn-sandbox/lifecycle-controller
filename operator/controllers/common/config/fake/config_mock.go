// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package fake

import (
	"sync"
	"time"
)

// MockConfig is a mock implementation of config.IConfig.
//
//	func TestSomethingThatUsesIConfig(t *testing.T) {
//
//		// make and configure a mocked config.IConfig
//		mockedIConfig := &MockConfig{
//			GetCreationRequestTimeoutFunc: func() time.Duration {
//				panic("mock out the GetCreationRequestTimeout method")
//			},
//			SetCreationRequestTimeoutFunc: func(value time.Duration)  {
//				panic("mock out the SetCreationRequestTimeout method")
//			},
//		}
//
//		// use mockedIConfig in code that requires config.IConfig
//		// and then make assertions.
//
//	}
type MockConfig struct {
	// GetCreationRequestTimeoutFunc mocks the GetCreationRequestTimeout method.
	GetCreationRequestTimeoutFunc func() time.Duration

	// SetCreationRequestTimeoutFunc mocks the SetCreationRequestTimeout method.
	SetCreationRequestTimeoutFunc func(value time.Duration)

	// calls tracks calls to the methods.
	calls struct {
		// GetCreationRequestTimeout holds details about calls to the GetCreationRequestTimeout method.
		GetCreationRequestTimeout []struct {
		}
		// SetCreationRequestTimeout holds details about calls to the SetCreationRequestTimeout method.
		SetCreationRequestTimeout []struct {
			// Value is the value argument value.
			Value time.Duration
		}
	}
	lockGetCreationRequestTimeout sync.RWMutex
	lockSetCreationRequestTimeout sync.RWMutex
}

// GetCreationRequestTimeout calls GetCreationRequestTimeoutFunc.
func (mock *MockConfig) GetCreationRequestTimeout() time.Duration {
	if mock.GetCreationRequestTimeoutFunc == nil {
		panic("MockConfig.GetCreationRequestTimeoutFunc: method is nil but IConfig.GetCreationRequestTimeout was just called")
	}
	callInfo := struct {
	}{}
	mock.lockGetCreationRequestTimeout.Lock()
	mock.calls.GetCreationRequestTimeout = append(mock.calls.GetCreationRequestTimeout, callInfo)
	mock.lockGetCreationRequestTimeout.Unlock()
	return mock.GetCreationRequestTimeoutFunc()
}

// GetCreationRequestTimeoutCalls gets all the calls that were made to GetCreationRequestTimeout.
// Check the length with:
//
//	len(mockedIConfig.GetCreationRequestTimeoutCalls())
func (mock *MockConfig) GetCreationRequestTimeoutCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockGetCreationRequestTimeout.RLock()
	calls = mock.calls.GetCreationRequestTimeout
	mock.lockGetCreationRequestTimeout.RUnlock()
	return calls
}

// SetCreationRequestTimeout calls SetCreationRequestTimeoutFunc.
func (mock *MockConfig) SetCreationRequestTimeout(value time.Duration) {
	if mock.SetCreationRequestTimeoutFunc == nil {
		panic("MockConfig.SetCreationRequestTimeoutFunc: method is nil but IConfig.SetCreationRequestTimeout was just called")
	}
	callInfo := struct {
		Value time.Duration
	}{
		Value: value,
	}
	mock.lockSetCreationRequestTimeout.Lock()
	mock.calls.SetCreationRequestTimeout = append(mock.calls.SetCreationRequestTimeout, callInfo)
	mock.lockSetCreationRequestTimeout.Unlock()
	mock.SetCreationRequestTimeoutFunc(value)
}

// SetCreationRequestTimeoutCalls gets all the calls that were made to SetCreationRequestTimeout.
// Check the length with:
//
//	len(mockedIConfig.SetCreationRequestTimeoutCalls())
func (mock *MockConfig) SetCreationRequestTimeoutCalls() []struct {
	Value time.Duration
} {
	var calls []struct {
		Value time.Duration
	}
	mock.lockSetCreationRequestTimeout.RLock()
	calls = mock.calls.SetCreationRequestTimeout
	mock.lockSetCreationRequestTimeout.RUnlock()
	return calls
}
