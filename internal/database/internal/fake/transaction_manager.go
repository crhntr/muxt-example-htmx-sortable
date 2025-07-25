// Code generated by counterfeiter. DO NOT EDIT.
package fake

import (
	"context"
	"sync"

	"github.com/crhntr/muxt-example-htmx-sortable/internal/database"
	"github.com/crhntr/transaction"
	pgx "github.com/jackc/pgx/v5"
)

type TransactionManager struct {
	CallStub        func(context.Context, pgx.TxOptions, transaction.Func) error
	callMutex       sync.RWMutex
	callArgsForCall []struct {
		arg1 context.Context
		arg2 pgx.TxOptions
		arg3 transaction.Func
	}
	callReturns struct {
		result1 error
	}
	callReturnsOnCall map[int]struct {
		result1 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *TransactionManager) Call(arg1 context.Context, arg2 pgx.TxOptions, arg3 transaction.Func) error {
	fake.callMutex.Lock()
	ret, specificReturn := fake.callReturnsOnCall[len(fake.callArgsForCall)]
	fake.callArgsForCall = append(fake.callArgsForCall, struct {
		arg1 context.Context
		arg2 pgx.TxOptions
		arg3 transaction.Func
	}{arg1, arg2, arg3})
	stub := fake.CallStub
	fakeReturns := fake.callReturns
	fake.recordInvocation("Call", []interface{}{arg1, arg2, arg3})
	fake.callMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *TransactionManager) CallCallCount() int {
	fake.callMutex.RLock()
	defer fake.callMutex.RUnlock()
	return len(fake.callArgsForCall)
}

func (fake *TransactionManager) CallCalls(stub func(context.Context, pgx.TxOptions, transaction.Func) error) {
	fake.callMutex.Lock()
	defer fake.callMutex.Unlock()
	fake.CallStub = stub
}

func (fake *TransactionManager) CallArgsForCall(i int) (context.Context, pgx.TxOptions, transaction.Func) {
	fake.callMutex.RLock()
	defer fake.callMutex.RUnlock()
	argsForCall := fake.callArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *TransactionManager) CallReturns(result1 error) {
	fake.callMutex.Lock()
	defer fake.callMutex.Unlock()
	fake.CallStub = nil
	fake.callReturns = struct {
		result1 error
	}{result1}
}

func (fake *TransactionManager) CallReturnsOnCall(i int, result1 error) {
	fake.callMutex.Lock()
	defer fake.callMutex.Unlock()
	fake.CallStub = nil
	if fake.callReturnsOnCall == nil {
		fake.callReturnsOnCall = make(map[int]struct {
			result1 error
		})
	}
	fake.callReturnsOnCall[i] = struct {
		result1 error
	}{result1}
}

func (fake *TransactionManager) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *TransactionManager) recordInvocation(key string, args []interface{}) {
	fake.invocationsMutex.Lock()
	defer fake.invocationsMutex.Unlock()
	if fake.invocations == nil {
		fake.invocations = map[string][][]interface{}{}
	}
	if fake.invocations[key] == nil {
		fake.invocations[key] = [][]interface{}{}
	}
	fake.invocations[key] = append(fake.invocations[key], args)
}

var _ database.Caller = new(TransactionManager)
