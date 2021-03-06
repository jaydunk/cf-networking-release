// Code generated by counterfeiter. DO NOT EDIT.
package fakes

import (
	"context"
	"sync"
	"time"
)

type ContextAdapter struct {
	WithTimeoutStub        func(context.Context, time.Duration) (context.Context, context.CancelFunc)
	withTimeoutMutex       sync.RWMutex
	withTimeoutArgsForCall []struct {
		arg1 context.Context
		arg2 time.Duration
	}
	withTimeoutReturns struct {
		result1 context.Context
		result2 context.CancelFunc
	}
	withTimeoutReturnsOnCall map[int]struct {
		result1 context.Context
		result2 context.CancelFunc
	}
	invocations      map[string][][]interface{}
	invocationsMutex sync.RWMutex
}

func (fake *ContextAdapter) WithTimeout(arg1 context.Context, arg2 time.Duration) (context.Context, context.CancelFunc) {
	fake.withTimeoutMutex.Lock()
	ret, specificReturn := fake.withTimeoutReturnsOnCall[len(fake.withTimeoutArgsForCall)]
	fake.withTimeoutArgsForCall = append(fake.withTimeoutArgsForCall, struct {
		arg1 context.Context
		arg2 time.Duration
	}{arg1, arg2})
	fake.recordInvocation("WithTimeout", []interface{}{arg1, arg2})
	fake.withTimeoutMutex.Unlock()
	if fake.WithTimeoutStub != nil {
		return fake.WithTimeoutStub(arg1, arg2)
	}
	if specificReturn {
		return ret.result1, ret.result2
	}
	return fake.withTimeoutReturns.result1, fake.withTimeoutReturns.result2
}

func (fake *ContextAdapter) WithTimeoutCallCount() int {
	fake.withTimeoutMutex.RLock()
	defer fake.withTimeoutMutex.RUnlock()
	return len(fake.withTimeoutArgsForCall)
}

func (fake *ContextAdapter) WithTimeoutArgsForCall(i int) (context.Context, time.Duration) {
	fake.withTimeoutMutex.RLock()
	defer fake.withTimeoutMutex.RUnlock()
	return fake.withTimeoutArgsForCall[i].arg1, fake.withTimeoutArgsForCall[i].arg2
}

func (fake *ContextAdapter) WithTimeoutReturns(result1 context.Context, result2 context.CancelFunc) {
	fake.WithTimeoutStub = nil
	fake.withTimeoutReturns = struct {
		result1 context.Context
		result2 context.CancelFunc
	}{result1, result2}
}

func (fake *ContextAdapter) WithTimeoutReturnsOnCall(i int, result1 context.Context, result2 context.CancelFunc) {
	fake.WithTimeoutStub = nil
	if fake.withTimeoutReturnsOnCall == nil {
		fake.withTimeoutReturnsOnCall = make(map[int]struct {
			result1 context.Context
			result2 context.CancelFunc
		})
	}
	fake.withTimeoutReturnsOnCall[i] = struct {
		result1 context.Context
		result2 context.CancelFunc
	}{result1, result2}
}

func (fake *ContextAdapter) Invocations() map[string][][]interface{} {
	fake.invocationsMutex.RLock()
	defer fake.invocationsMutex.RUnlock()
	fake.withTimeoutMutex.RLock()
	defer fake.withTimeoutMutex.RUnlock()
	copiedInvocations := map[string][][]interface{}{}
	for key, value := range fake.invocations {
		copiedInvocations[key] = value
	}
	return copiedInvocations
}

func (fake *ContextAdapter) recordInvocation(key string, args []interface{}) {
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
