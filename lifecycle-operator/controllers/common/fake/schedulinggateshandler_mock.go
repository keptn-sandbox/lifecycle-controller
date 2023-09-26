// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package fake

import (
	"context"
	lfcv1alpha3 "github.com/keptn/lifecycle-toolkit/lifecycle-operator/apis/lifecycle/v1alpha3"
	"sync"
)

// ISchedulingGatesHandlerMock is a mock implementation of common.ISchedulingGatesHandler.
//
//	func TestSomethingThatUsesISchedulingGatesHandler(t *testing.T) {
//
//		// make and configure a mocked common.ISchedulingGatesHandler
//		mockedISchedulingGatesHandler := &ISchedulingGatesHandlerMock{
//			EnabledFunc: func() bool {
//				panic("mock out the Enabled method")
//			},
//			RemoveGatesFunc: func(ctx context.Context, workloadInstance *lfcv1alpha3.KeptnWorkloadInstance) error {
//				panic("mock out the RemoveGates method")
//			},
//		}
//
//		// use mockedISchedulingGatesHandler in code that requires common.ISchedulingGatesHandler
//		// and then make assertions.
//
//	}
type ISchedulingGatesHandlerMock struct {
	// EnabledFunc mocks the Enabled method.
	EnabledFunc func() bool

	// RemoveGatesFunc mocks the RemoveGates method.
	RemoveGatesFunc func(ctx context.Context, workloadInstance *lfcv1alpha3.KeptnWorkloadInstance) error

	// calls tracks calls to the methods.
	calls struct {
		// Enabled holds details about calls to the Enabled method.
		Enabled []struct {
		}
		// RemoveGates holds details about calls to the RemoveGates method.
		RemoveGates []struct {
			// Ctx is the ctx argument value.
			Ctx context.Context
			// WorkloadInstance is the workloadInstance argument value.
			WorkloadInstance *lfcv1alpha3.KeptnWorkloadInstance
		}
	}
	lockEnabled sync.RWMutex
	lockRemoveGates         sync.RWMutex
}

// Enabled calls EnabledFunc.
func (mock *ISchedulingGatesHandlerMock) Enabled() bool {
	if mock.EnabledFunc == nil {
		panic("ISchedulingGatesHandlerMock.EnabledFunc: method is nil but ISchedulingGatesHandler.Enabled was just called")
	}
	callInfo := struct {
	}{}
	mock.lockEnabled.Lock()
	mock.calls.Enabled = append(mock.calls.Enabled, callInfo)
	mock.lockEnabled.Unlock()
	return mock.EnabledFunc()
}

// EnabledCalls gets all the calls that were made to Enabled.
// Check the length with:
//
//	len(mockedISchedulingGatesHandler.EnabledCalls())
func (mock *ISchedulingGatesHandlerMock) EnabledCalls() []struct {
} {
	var calls []struct {
	}
	mock.lockEnabled.RLock()
	calls = mock.calls.Enabled
	mock.lockEnabled.RUnlock()
	return calls
}

// RemoveGates calls RemoveGatesFunc.
func (mock *ISchedulingGatesHandlerMock) RemoveGates(ctx context.Context, workloadInstance *lfcv1alpha3.KeptnWorkloadInstance) error {
	if mock.RemoveGatesFunc == nil {
		panic("ISchedulingGatesHandlerMock.RemoveGatesFunc: method is nil but ISchedulingGatesHandler.RemoveGates was just called")
	}
	callInfo := struct {
		Ctx              context.Context
		WorkloadInstance *lfcv1alpha3.KeptnWorkloadInstance
	}{
		Ctx:              ctx,
		WorkloadInstance: workloadInstance,
	}
	mock.lockRemoveGates.Lock()
	mock.calls.RemoveGates = append(mock.calls.RemoveGates, callInfo)
	mock.lockRemoveGates.Unlock()
	return mock.RemoveGatesFunc(ctx, workloadInstance)
}

// RemoveGatesCalls gets all the calls that were made to RemoveGates.
// Check the length with:
//
//	len(mockedISchedulingGatesHandler.RemoveGatesCalls())
func (mock *ISchedulingGatesHandlerMock) RemoveGatesCalls() []struct {
	Ctx              context.Context
	WorkloadInstance *lfcv1alpha3.KeptnWorkloadInstance
} {
	var calls []struct {
		Ctx              context.Context
		WorkloadInstance *lfcv1alpha3.KeptnWorkloadInstance
	}
	mock.lockRemoveGates.RLock()
	calls = mock.calls.RemoveGates
	mock.lockRemoveGates.RUnlock()
	return calls
}
