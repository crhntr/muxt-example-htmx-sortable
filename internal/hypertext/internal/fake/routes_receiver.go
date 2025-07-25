// Code generated by counterfeiter. DO NOT EDIT.
package fake

import (
	"context"
	"sync"

	"github.com/crhntr/muxt-example-htmx-sortable/internal/database"
	"github.com/crhntr/muxt-example-htmx-sortable/internal/hypertext"
)

type RoutesReceiver struct {
	GetListStub        func(context.Context, int32) hypertext.ListData
	getListMutex       sync.RWMutex
	getListArgsForCall []struct {
		arg1 context.Context
		arg2 int32
	}
	getListReturns struct {
		result1 hypertext.ListData
	}
	getListReturnsOnCall map[int]struct {
		result1 hypertext.ListData
	}
	ListsStub        func(context.Context) []database.List
	listsMutex       sync.RWMutex
	listsArgsForCall []struct {
		arg1 context.Context
	}
	listsReturns struct {
		result1 []database.List
	}
	listsReturnsOnCall map[int]struct {
		result1 []database.List
	}
	UpdateListStub        func(context.Context, int32, hypertext.UpdateListForm) hypertext.ListData
	updateListMutex       sync.RWMutex
	updateListArgsForCall []struct {
		arg1 context.Context
		arg2 int32
		arg3 hypertext.UpdateListForm
	}
	updateListReturns struct {
		result1 hypertext.ListData
	}
	updateListReturnsOnCall map[int]struct {
		result1 hypertext.ListData
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *RoutesReceiver) GetList(arg1 context.Context, arg2 int32) hypertext.ListData {
	fake.getListMutex.Lock()
	ret, specificReturn := fake.getListReturnsOnCall[len(fake.getListArgsForCall)]
	fake.getListArgsForCall = append(fake.getListArgsForCall, struct {
		arg1 context.Context
		arg2 int32
	}{arg1, arg2})
	stub := fake.GetListStub
	fakeReturns := fake.getListReturns
	fake.recordInvocation("GetList", []interface{}{arg1, arg2})
	fake.getListMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *RoutesReceiver) GetListCallCount() int {
	fake.getListMutex.RLock()
	defer fake.getListMutex.RUnlock()
	return len(fake.getListArgsForCall)
}

func (fake *RoutesReceiver) GetListCalls(stub func(context.Context, int32) hypertext.ListData) {
	fake.getListMutex.Lock()
	defer fake.getListMutex.Unlock()
	fake.GetListStub = stub
}

func (fake *RoutesReceiver) GetListArgsForCall(i int) (context.Context, int32) {
	fake.getListMutex.RLock()
	defer fake.getListMutex.RUnlock()
	argsForCall := fake.getListArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2
}

func (fake *RoutesReceiver) GetListReturns(result1 hypertext.ListData) {
	fake.getListMutex.Lock()
	defer fake.getListMutex.Unlock()
	fake.GetListStub = nil
	fake.getListReturns = struct {
		result1 hypertext.ListData
	}{result1}
}

func (fake *RoutesReceiver) GetListReturnsOnCall(i int, result1 hypertext.ListData) {
	fake.getListMutex.Lock()
	defer fake.getListMutex.Unlock()
	fake.GetListStub = nil
	if fake.getListReturnsOnCall == nil {
		fake.getListReturnsOnCall = make(map[int]struct {
			result1 hypertext.ListData
		})
	}
	fake.getListReturnsOnCall[i] = struct {
		result1 hypertext.ListData
	}{result1}
}

func (fake *RoutesReceiver) Lists(arg1 context.Context) []database.List {
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
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *RoutesReceiver) ListsCallCount() int {
	fake.listsMutex.RLock()
	defer fake.listsMutex.RUnlock()
	return len(fake.listsArgsForCall)
}

func (fake *RoutesReceiver) ListsCalls(stub func(context.Context) []database.List) {
	fake.listsMutex.Lock()
	defer fake.listsMutex.Unlock()
	fake.ListsStub = stub
}

func (fake *RoutesReceiver) ListsArgsForCall(i int) context.Context {
	fake.listsMutex.RLock()
	defer fake.listsMutex.RUnlock()
	argsForCall := fake.listsArgsForCall[i]
	return argsForCall.arg1
}

func (fake *RoutesReceiver) ListsReturns(result1 []database.List) {
	fake.listsMutex.Lock()
	defer fake.listsMutex.Unlock()
	fake.ListsStub = nil
	fake.listsReturns = struct {
		result1 []database.List
	}{result1}
}

func (fake *RoutesReceiver) ListsReturnsOnCall(i int, result1 []database.List) {
	fake.listsMutex.Lock()
	defer fake.listsMutex.Unlock()
	fake.ListsStub = nil
	if fake.listsReturnsOnCall == nil {
		fake.listsReturnsOnCall = make(map[int]struct {
			result1 []database.List
		})
	}
	fake.listsReturnsOnCall[i] = struct {
		result1 []database.List
	}{result1}
}

func (fake *RoutesReceiver) UpdateList(arg1 context.Context, arg2 int32, arg3 hypertext.UpdateListForm) hypertext.ListData {
	fake.updateListMutex.Lock()
	ret, specificReturn := fake.updateListReturnsOnCall[len(fake.updateListArgsForCall)]
	fake.updateListArgsForCall = append(fake.updateListArgsForCall, struct {
		arg1 context.Context
		arg2 int32
		arg3 hypertext.UpdateListForm
	}{arg1, arg2, arg3})
	stub := fake.UpdateListStub
	fakeReturns := fake.updateListReturns
	fake.recordInvocation("UpdateList", []interface{}{arg1, arg2, arg3})
	fake.updateListMutex.Unlock()
	if stub != nil {
		return stub(arg1, arg2, arg3)
	}
	if specificReturn {
		return ret.result1
	}
	return fakeReturns.result1
}

func (fake *RoutesReceiver) UpdateListCallCount() int {
	fake.updateListMutex.RLock()
	defer fake.updateListMutex.RUnlock()
	return len(fake.updateListArgsForCall)
}

func (fake *RoutesReceiver) UpdateListCalls(stub func(context.Context, int32, hypertext.UpdateListForm) hypertext.ListData) {
	fake.updateListMutex.Lock()
	defer fake.updateListMutex.Unlock()
	fake.UpdateListStub = stub
}

func (fake *RoutesReceiver) UpdateListArgsForCall(i int) (context.Context, int32, hypertext.UpdateListForm) {
	fake.updateListMutex.RLock()
	defer fake.updateListMutex.RUnlock()
	argsForCall := fake.updateListArgsForCall[i]
	return argsForCall.arg1, argsForCall.arg2, argsForCall.arg3
}

func (fake *RoutesReceiver) UpdateListReturns(result1 hypertext.ListData) {
	fake.updateListMutex.Lock()
	defer fake.updateListMutex.Unlock()
	fake.UpdateListStub = nil
	fake.updateListReturns = struct {
		result1 hypertext.ListData
	}{result1}
}

func (fake *RoutesReceiver) UpdateListReturnsOnCall(i int, result1 hypertext.ListData) {
	fake.updateListMutex.Lock()
	defer fake.updateListMutex.Unlock()
	fake.UpdateListStub = nil
	if fake.updateListReturnsOnCall == nil {
		fake.updateListReturnsOnCall = make(map[int]struct {
			result1 hypertext.ListData
		})
	}
	fake.updateListReturnsOnCall[i] = struct {
		result1 hypertext.ListData
	}{result1}
}

func (fake *RoutesReceiver) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *RoutesReceiver) recordInvocation(key string, args []interface{}) {
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

var _ hypertext.RoutesReceiver = new(RoutesReceiver)
