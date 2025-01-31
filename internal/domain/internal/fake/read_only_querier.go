// Code generated by counterfeiter. DO NOT EDIT.
package fake

import (
	"context"
	"sync"

	"github.com/crhntr/muxt-example-htmx-sortable/internal/database"
)

type ReadOnlyQuerier struct {
	ListByIDStub        func(context.Context, int32) (database.List, error)
	listByIDMutex       sync.RWMutex
	listByIDArgsForCall []struct {
		arg1 context.Context
		arg2 int32
	}
	listByIDReturns struct {
		result1 database.List
		result2 error
	}
	listByIDReturnsOnCall map[int]struct {
		result1 database.List
		result2 error
	}
	ListsStub        func(context.Context) ([]database.List, error)
	listsMutex       sync.RWMutex
	listsArgsForCall []struct {
		arg1 context.Context
	}
	listsReturns struct {
		result1 []database.List
		result2 error
	}
	listsReturnsOnCall map[int]struct {
		result1 []database.List
		result2 error
	}
	TasksByListIDStub        func(context.Context, int32) ([]database.TasksByListIDRow, error)
	tasksByListIDMutex       sync.RWMutex
	tasksByListIDArgsForCall []struct {
		arg1 context.Context
		arg2 int32
	}
	tasksByListIDReturns struct {
		result1 []database.TasksByListIDRow
		result2 error
	}
	tasksByListIDReturnsOnCall map[int]struct {
		result1 []database.TasksByListIDRow
		result2 error
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *ReadOnlyQuerier) ListByID(arg1 context.Context, arg2 int32) (database.List, error) {
	fake.listByIDMutex.Lock()
	ret, specificReturn := fake.listByIDReturnsOnCall[len(fake.listByIDArgsForCall)]
	fake.listByIDArgsForCall = append(fake.listByIDArgsForCall, struct {
		arg1 context.Context
		arg2 int32
	}{arg1, arg2})
	stub := fake.ListByIDStub
	fakeReturns := fake.listByIDReturns
	fake.recordInvocation("ListByID", []interface{}{arg1, arg2})
	fake.listByIDMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *ReadOnlyQuerier) ListByIDCallCount() int {
	fake.listByIDMutex.RLock()
	defer fake.listByIDMutex.RUnlock()
	return len(fake.listByIDArgsForCall)
}

func (fake *ReadOnlyQuerier) ListByIDCalls(stub func(context.Context, int32) (database.List, error)) {
	fake.listByIDMutex.Lock()
	defer fake.listByIDMutex.Unlock()
	fake.ListByIDStub = stub
}

func (fake *ReadOnlyQuerier) ListByIDArgsForCall(i int) (context.Context, int32) {
	fake.listByIDMutex.RLock()
	defer fake.listByIDMutex.RUnlock()
	argsForCall := fake.listByIDArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *ReadOnlyQuerier) ListByIDReturns(result1 database.List, result2 error) {
	fake.listByIDMutex.Lock()
	defer fake.listByIDMutex.Unlock()
	fake.ListByIDStub = nil
	fake.listByIDReturns = struct {
		result1 database.List
		result2 error
	}{result1, result2}
}

func (fake *ReadOnlyQuerier) ListByIDReturnsOnCall(i int, result1 database.List, result2 error) {
	fake.listByIDMutex.Lock()
	defer fake.listByIDMutex.Unlock()
	fake.ListByIDStub = nil
	if fake.listByIDReturnsOnCall == nil {
		fake.listByIDReturnsOnCall = make(map[int]struct {
			result1 database.List
			result2 error
		})
	}
	fake.listByIDReturnsOnCall[i] = struct {
		result1 database.List
		result2 error
	}{result1, result2}
}

func (fake *ReadOnlyQuerier) Lists(arg1 context.Context) ([]database.List, error) {
	fake.listsMutex.Lock()
	ret, specificReturn := fake.listsReturnsOnCall[len(fake.listsArgsForCall)]
	fake.listsArgsForCall = append(fake.listsArgsForCall, struct {
		arg1 context.Context
	}{arg1})
	stub := fake.ListsStub
	fakeReturns := fake.listsReturns
	fake.recordInvocation("Lists", []interface{}{arg1})
	fake.listsMutex.Unlock()
	if stub != nil {
		return stub(arg1)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *ReadOnlyQuerier) ListsCallCount() int {
	fake.listsMutex.RLock()
	defer fake.listsMutex.RUnlock()
	return len(fake.listsArgsForCall)
}

func (fake *ReadOnlyQuerier) ListsCalls(stub func(context.Context) ([]database.List, error)) {
	fake.listsMutex.Lock()
	defer fake.listsMutex.Unlock()
	fake.ListsStub = stub
}

func (fake *ReadOnlyQuerier) ListsArgsForCall(i int) context.Context {
	fake.listsMutex.RLock()
	defer fake.listsMutex.RUnlock()
	argsForCall := fake.listsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *ReadOnlyQuerier) ListsReturns(result1 []database.List, result2 error) {
	fake.listsMutex.Lock()
	defer fake.listsMutex.Unlock()
	fake.ListsStub = nil
	fake.listsReturns = struct {
		result1 []database.List
		result2 error
	}{result1, result2}
}

func (fake *ReadOnlyQuerier) ListsReturnsOnCall(i int, result1 []database.List, result2 error) {
	fake.listsMutex.Lock()
	defer fake.listsMutex.Unlock()
	fake.ListsStub = nil
	if fake.listsReturnsOnCall == nil {
		fake.listsReturnsOnCall = make(map[int]struct {
			result1 []database.List
			result2 error
		})
	}
	fake.listsReturnsOnCall[i] = struct {
		result1 []database.List
		result2 error
	}{result1, result2}
}

func (fake *ReadOnlyQuerier) TasksByListID(arg1 context.Context, arg2 int32) ([]database.TasksByListIDRow, error) {
	fake.tasksByListIDMutex.Lock()
	ret, specificReturn := fake.tasksByListIDReturnsOnCall[len(fake.tasksByListIDArgsForCall)]
	fake.tasksByListIDArgsForCall = append(fake.tasksByListIDArgsForCall, struct {
		arg1 context.Context
		arg2 int32
	}{arg1, arg2})
	stub := fake.TasksByListIDStub
	fakeReturns := fake.tasksByListIDReturns
	fake.recordInvocation("TasksByListID", []interface{}{arg1, arg2})
	fake.tasksByListIDMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fakeReturns.result1, fakeReturns.result2
}

func (fake *ReadOnlyQuerier) TasksByListIDCallCount() int {
	fake.tasksByListIDMutex.RLock()
	defer fake.tasksByListIDMutex.RUnlock()
	return len(fake.tasksByListIDArgsForCall)
}

func (fake *ReadOnlyQuerier) TasksByListIDCalls(stub func(context.Context, int32) ([]database.TasksByListIDRow, error)) {
	fake.tasksByListIDMutex.Lock()
	defer fake.tasksByListIDMutex.Unlock()
	fake.TasksByListIDStub = stub
}

func (fake *ReadOnlyQuerier) TasksByListIDArgsForCall(i int) (context.Context, int32) {
	fake.tasksByListIDMutex.RLock()
	defer fake.tasksByListIDMutex.RUnlock()
	argsForCall := fake.tasksByListIDArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *ReadOnlyQuerier) TasksByListIDReturns(result1 []database.TasksByListIDRow, result2 error) {
	fake.tasksByListIDMutex.Lock()
	defer fake.tasksByListIDMutex.Unlock()
	fake.TasksByListIDStub = nil
	fake.tasksByListIDReturns = struct {
		result1 []database.TasksByListIDRow
		result2 error
	}{result1, result2}
}

func (fake *ReadOnlyQuerier) TasksByListIDReturnsOnCall(i int, result1 []database.TasksByListIDRow, result2 error) {
	fake.tasksByListIDMutex.Lock()
	defer fake.tasksByListIDMutex.Unlock()
	fake.TasksByListIDStub = nil
	if fake.tasksByListIDReturnsOnCall == nil {
		fake.tasksByListIDReturnsOnCall = make(map[int]struct {
			result1 []database.TasksByListIDRow
			result2 error
		})
	}
	fake.tasksByListIDReturnsOnCall[i] = struct {
		result1 []database.TasksByListIDRow
		result2 error
	}{result1, result2}
}

func (fake *ReadOnlyQuerier) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.listByIDMutex.RLock()
	defer fake.listByIDMutex.RUnlock()
	fake.listsMutex.RLock()
	defer fake.listsMutex.RUnlock()
	fake.tasksByListIDMutex.RLock()
	defer fake.tasksByListIDMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *ReadOnlyQuerier) recordInvocation(key string, args []interface{}) {
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

var _ database.ReadOnlyQuerier = new(ReadOnlyQuerier)
