// Code generated by counterfeiter. DO NOT EDIT.
package fake

import (
	"context"
	"sync"

	"github.com/crhntr/muxt-example-htmx-sortable/internal/database/internal/transaction"
	pgx "github.com/jackc/pgx/v5"
)

type TxBeginner struct {
	BeginTxStub        func(context.Context, pgx.TxOptions) (pgx.Tx, error)
	beginTxMutex       sync.RWMutex
	beginTxArgsForCall []struct {
		arg1 context.Context
		arg2 pgx.TxOptions
	}
	beginTxReturns struct {
		result1 pgx.Tx
		result2 error
	}
	beginTxReturnsOnCall map[int]struct {
		result1 pgx.Tx
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *TxBeginner) BeginTx(arg1 context.Context, arg2 pgx.TxOptions) (pgx.Tx, error) {
	fake.beginTxMutex.Lock()
	ret, specificReturn := fake.beginTxReturnsOnCall[len(fake.beginTxArgsForCall)]
	fake.beginTxArgsForCall = append(fake.beginTxArgsForCall, struct {
		arg1 context.Context
		arg2 pgx.TxOptions
	}{arg1, arg2})
	stub := fake.BeginTxStub
	fakeReturns := fake.beginTxReturns
	fake.recordInvocation("BeginTx", []interface{}{arg1, arg2})
	fake.beginTxMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *TxBeginner) BeginTxCallCount() int {
	fake.beginTxMutex.RLock()
	defer fake.beginTxMutex.RUnlock()
	return len(fake.beginTxArgsForCall)
}

func (fake *TxBeginner) BeginTxCalls(stub func(context.Context, pgx.TxOptions) (pgx.Tx, error)) {
	fake.beginTxMutex.Lock()
	defer fake.beginTxMutex.Unlock()
	fake.BeginTxStub = stub
}

func (fake *TxBeginner) BeginTxArgsForCall(i int) (context.Context, pgx.TxOptions) {
	fake.beginTxMutex.RLock()
	defer fake.beginTxMutex.RUnlock()
	argsForCall := fake.beginTxArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *TxBeginner) BeginTxReturns(result1 pgx.Tx, result2 error) {
	fake.beginTxMutex.Lock()
	defer fake.beginTxMutex.Unlock()
	fake.BeginTxStub = nil
	fake.beginTxReturns = struct {
		result1 pgx.Tx
		result2 error
	}{result1, result2}
}

func (fake *TxBeginner) BeginTxReturnsOnCall(i int, result1 pgx.Tx, result2 error) {
	fake.beginTxMutex.Lock()
	defer fake.beginTxMutex.Unlock()
	fake.BeginTxStub = nil
	if fake.beginTxReturnsOnCall == nil {
		fake.beginTxReturnsOnCall = make(map[int]struct {
			result1 pgx.Tx
			result2 error
		})
	}
	fake.beginTxReturnsOnCall[i] = struct {
		result1 pgx.Tx
		result2 error
	}{result1, result2}
}

func (fake *TxBeginner) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.beginTxMutex.RLock()
	defer fake.beginTxMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *TxBeginner) recordInvocation(key string, args []interface{}) {
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

var _ transaction.TxBeginner = new(TxBeginner)
