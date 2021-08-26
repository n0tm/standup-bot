package telegram

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

//go:generate minimock -i standup-bot/src/telegram.Middleware -o ./middleware_mock.go -n MiddlewareMock

import (
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
)

// MiddlewareMock implements Middleware
type MiddlewareMock struct {
	t minimock.Tester

	funcMiddleware          func(u1 Update) (err error)
	inspectFuncMiddleware   func(u1 Update)
	afterMiddlewareCounter  uint64
	beforeMiddlewareCounter uint64
	MiddlewareMock          mMiddlewareMockMiddleware
}

// NewMiddlewareMock returns a mock for Middleware
func NewMiddlewareMock(t minimock.Tester) *MiddlewareMock {
	m := &MiddlewareMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.MiddlewareMock = mMiddlewareMockMiddleware{mock: m}
	m.MiddlewareMock.callArgs = []*MiddlewareMockMiddlewareParams{}

	return m
}

type mMiddlewareMockMiddleware struct {
	mock               *MiddlewareMock
	defaultExpectation *MiddlewareMockMiddlewareExpectation
	expectations       []*MiddlewareMockMiddlewareExpectation

	callArgs []*MiddlewareMockMiddlewareParams
	mutex    sync.RWMutex
}

// MiddlewareMockMiddlewareExpectation specifies expectation struct of the Middleware.Middleware
type MiddlewareMockMiddlewareExpectation struct {
	mock    *MiddlewareMock
	params  *MiddlewareMockMiddlewareParams
	results *MiddlewareMockMiddlewareResults
	Counter uint64
}

// MiddlewareMockMiddlewareParams contains parameters of the Middleware.Middleware
type MiddlewareMockMiddlewareParams struct {
	u1 Update
}

// MiddlewareMockMiddlewareResults contains results of the Middleware.Middleware
type MiddlewareMockMiddlewareResults struct {
	err error
}

// Expect sets up expected params for Middleware.Middleware
func (mmMiddleware *mMiddlewareMockMiddleware) Expect(u1 Update) *mMiddlewareMockMiddleware {
	if mmMiddleware.mock.funcMiddleware != nil {
		mmMiddleware.mock.t.Fatalf("MiddlewareMock.Middleware mock is already set by Set")
	}

	if mmMiddleware.defaultExpectation == nil {
		mmMiddleware.defaultExpectation = &MiddlewareMockMiddlewareExpectation{}
	}

	mmMiddleware.defaultExpectation.params = &MiddlewareMockMiddlewareParams{u1}
	for _, e := range mmMiddleware.expectations {
		if minimock.Equal(e.params, mmMiddleware.defaultExpectation.params) {
			mmMiddleware.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmMiddleware.defaultExpectation.params)
		}
	}

	return mmMiddleware
}

// Inspect accepts an inspector function that has same arguments as the Middleware.Middleware
func (mmMiddleware *mMiddlewareMockMiddleware) Inspect(f func(u1 Update)) *mMiddlewareMockMiddleware {
	if mmMiddleware.mock.inspectFuncMiddleware != nil {
		mmMiddleware.mock.t.Fatalf("Inspect function is already set for MiddlewareMock.Middleware")
	}

	mmMiddleware.mock.inspectFuncMiddleware = f

	return mmMiddleware
}

// Return sets up results that will be returned by Middleware.Middleware
func (mmMiddleware *mMiddlewareMockMiddleware) Return(err error) *MiddlewareMock {
	if mmMiddleware.mock.funcMiddleware != nil {
		mmMiddleware.mock.t.Fatalf("MiddlewareMock.Middleware mock is already set by Set")
	}

	if mmMiddleware.defaultExpectation == nil {
		mmMiddleware.defaultExpectation = &MiddlewareMockMiddlewareExpectation{mock: mmMiddleware.mock}
	}
	mmMiddleware.defaultExpectation.results = &MiddlewareMockMiddlewareResults{err}
	return mmMiddleware.mock
}

//Set uses given function f to mock the Middleware.Middleware method
func (mmMiddleware *mMiddlewareMockMiddleware) Set(f func(u1 Update) (err error)) *MiddlewareMock {
	if mmMiddleware.defaultExpectation != nil {
		mmMiddleware.mock.t.Fatalf("Default expectation is already set for the Middleware.Middleware method")
	}

	if len(mmMiddleware.expectations) > 0 {
		mmMiddleware.mock.t.Fatalf("Some expectations are already set for the Middleware.Middleware method")
	}

	mmMiddleware.mock.funcMiddleware = f
	return mmMiddleware.mock
}

// When sets expectation for the Middleware.Middleware which will trigger the result defined by the following
// Then helper
func (mmMiddleware *mMiddlewareMockMiddleware) When(u1 Update) *MiddlewareMockMiddlewareExpectation {
	if mmMiddleware.mock.funcMiddleware != nil {
		mmMiddleware.mock.t.Fatalf("MiddlewareMock.Middleware mock is already set by Set")
	}

	expectation := &MiddlewareMockMiddlewareExpectation{
		mock:   mmMiddleware.mock,
		params: &MiddlewareMockMiddlewareParams{u1},
	}
	mmMiddleware.expectations = append(mmMiddleware.expectations, expectation)
	return expectation
}

// Then sets up Middleware.Middleware return parameters for the expectation previously defined by the When method
func (e *MiddlewareMockMiddlewareExpectation) Then(err error) *MiddlewareMock {
	e.results = &MiddlewareMockMiddlewareResults{err}
	return e.mock
}

// Middleware implements Middleware
func (mmMiddleware *MiddlewareMock) Middleware(u1 Update) (err error) {
	mm_atomic.AddUint64(&mmMiddleware.beforeMiddlewareCounter, 1)
	defer mm_atomic.AddUint64(&mmMiddleware.afterMiddlewareCounter, 1)

	if mmMiddleware.inspectFuncMiddleware != nil {
		mmMiddleware.inspectFuncMiddleware(u1)
	}

	mm_params := &MiddlewareMockMiddlewareParams{u1}

	// Record call args
	mmMiddleware.MiddlewareMock.mutex.Lock()
	mmMiddleware.MiddlewareMock.callArgs = append(mmMiddleware.MiddlewareMock.callArgs, mm_params)
	mmMiddleware.MiddlewareMock.mutex.Unlock()

	for _, e := range mmMiddleware.MiddlewareMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmMiddleware.MiddlewareMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmMiddleware.MiddlewareMock.defaultExpectation.Counter, 1)
		mm_want := mmMiddleware.MiddlewareMock.defaultExpectation.params
		mm_got := MiddlewareMockMiddlewareParams{u1}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmMiddleware.t.Errorf("MiddlewareMock.Middleware got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmMiddleware.MiddlewareMock.defaultExpectation.results
		if mm_results == nil {
			mmMiddleware.t.Fatal("No results are set for the MiddlewareMock.Middleware")
		}
		return (*mm_results).err
	}
	if mmMiddleware.funcMiddleware != nil {
		return mmMiddleware.funcMiddleware(u1)
	}
	mmMiddleware.t.Fatalf("Unexpected call to MiddlewareMock.Middleware. %v", u1)
	return
}

// MiddlewareAfterCounter returns a count of finished MiddlewareMock.Middleware invocations
func (mmMiddleware *MiddlewareMock) MiddlewareAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmMiddleware.afterMiddlewareCounter)
}

// MiddlewareBeforeCounter returns a count of MiddlewareMock.Middleware invocations
func (mmMiddleware *MiddlewareMock) MiddlewareBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmMiddleware.beforeMiddlewareCounter)
}

// Calls returns a list of arguments used in each call to MiddlewareMock.Middleware.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmMiddleware *mMiddlewareMockMiddleware) Calls() []*MiddlewareMockMiddlewareParams {
	mmMiddleware.mutex.RLock()

	argCopy := make([]*MiddlewareMockMiddlewareParams, len(mmMiddleware.callArgs))
	copy(argCopy, mmMiddleware.callArgs)

	mmMiddleware.mutex.RUnlock()

	return argCopy
}

// MinimockMiddlewareDone returns true if the count of the Middleware invocations corresponds
// the number of defined expectations
func (m *MiddlewareMock) MinimockMiddlewareDone() bool {
	for _, e := range m.MiddlewareMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.MiddlewareMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterMiddlewareCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcMiddleware != nil && mm_atomic.LoadUint64(&m.afterMiddlewareCounter) < 1 {
		return false
	}
	return true
}

// MinimockMiddlewareInspect logs each unmet expectation
func (m *MiddlewareMock) MinimockMiddlewareInspect() {
	for _, e := range m.MiddlewareMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to MiddlewareMock.Middleware with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.MiddlewareMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterMiddlewareCounter) < 1 {
		if m.MiddlewareMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to MiddlewareMock.Middleware")
		} else {
			m.t.Errorf("Expected call to MiddlewareMock.Middleware with params: %#v", *m.MiddlewareMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcMiddleware != nil && mm_atomic.LoadUint64(&m.afterMiddlewareCounter) < 1 {
		m.t.Error("Expected call to MiddlewareMock.Middleware")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *MiddlewareMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockMiddlewareInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *MiddlewareMock) MinimockWait(timeout mm_time.Duration) {
	timeoutCh := mm_time.After(timeout)
	for {
		if m.minimockDone() {
			return
		}
		select {
		case <-timeoutCh:
			m.MinimockFinish()
			return
		case <-mm_time.After(10 * mm_time.Millisecond):
		}
	}
}

func (m *MiddlewareMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockMiddlewareDone()
}
