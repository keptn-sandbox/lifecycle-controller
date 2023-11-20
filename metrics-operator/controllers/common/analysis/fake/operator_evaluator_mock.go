// Code generated by moq; DO NOT EDIT.
// github.com/matryer/moq

package fake

import (
	metricsapi "github.com/keptn/lifecycle-toolkit/metrics-operator/api/v1beta1"
	"github.com/keptn/lifecycle-toolkit/metrics-operator/controllers/common/analysis/types"
	"sync"
)

// IOperatorEvaluatorMock is a mock implementation of analysis.IOperatorEvaluator.
//
//	func TestSomethingThatUsesIOperatorEvaluator(t *testing.T) {
//
//		// make and configure a mocked analysis.IOperatorEvaluator
//		mockedIOperatorEvaluator := &IOperatorEvaluatorMock{
//			EvaluateFunc: func(val float64, criteria *metricsapi.Operator) types.OperatorResult {
//				panic("mock out the Evaluate method")
//			},
//		}
//
//		// use mockedIOperatorEvaluator in code that requires analysis.IOperatorEvaluator
//		// and then make assertions.
//
//	}
type IOperatorEvaluatorMock struct {
	// EvaluateFunc mocks the Evaluate method.
	EvaluateFunc func(val float64, criteria *metricsapi.Operator) types.OperatorResult

	// calls tracks calls to the methods.
	calls struct {
		// Evaluate holds details about calls to the Evaluate method.
		Evaluate []struct {
			// Val is the val argument value.
			Val float64
			// Criteria is the criteria argument value.
			Criteria *metricsapi.Operator
		}
	}
	lockEvaluate sync.RWMutex
}

// Evaluate calls EvaluateFunc.
func (mock *IOperatorEvaluatorMock) Evaluate(val float64, criteria *metricsapi.Operator) types.OperatorResult {
	if mock.EvaluateFunc == nil {
		panic("IOperatorEvaluatorMock.EvaluateFunc: method is nil but IOperatorEvaluator.Evaluate was just called")
	}
	callInfo := struct {
		Val      float64
		Criteria *metricsapi.Operator
	}{
		Val:      val,
		Criteria: criteria,
	}
	mock.lockEvaluate.Lock()
	mock.calls.Evaluate = append(mock.calls.Evaluate, callInfo)
	mock.lockEvaluate.Unlock()
	return mock.EvaluateFunc(val, criteria)
}

// EvaluateCalls gets all the calls that were made to Evaluate.
// Check the length with:
//
//	len(mockedIOperatorEvaluator.EvaluateCalls())
func (mock *IOperatorEvaluatorMock) EvaluateCalls() []struct {
	Val      float64
	Criteria *metricsapi.Operator
} {
	var calls []struct {
		Val      float64
		Criteria *metricsapi.Operator
	}
	mock.lockEvaluate.RLock()
	calls = mock.calls.Evaluate
	mock.lockEvaluate.RUnlock()
	return calls
}
