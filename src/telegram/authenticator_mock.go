package telegram

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

//go:generate minimock -i standup-bot/src/telegram.Authenticator -o ./authenticator_mock.go -n AuthenticatorMock

import (
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
)

// AuthenticatorMock implements Authenticator
type AuthenticatorMock struct {
	t minimock.Tester

	funcAuthenticate          func(u1 Update) (u2 User, err error)
	inspectFuncAuthenticate   func(u1 Update)
	afterAuthenticateCounter  uint64
	beforeAuthenticateCounter uint64
	AuthenticateMock          mAuthenticatorMockAuthenticate
}

// NewAuthenticatorMock returns a mock for Authenticator
func NewAuthenticatorMock(t minimock.Tester) *AuthenticatorMock {
	m := &AuthenticatorMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.AuthenticateMock = mAuthenticatorMockAuthenticate{mock: m}
	m.AuthenticateMock.callArgs = []*AuthenticatorMockAuthenticateParams{}

	return m
}

type mAuthenticatorMockAuthenticate struct {
	mock               *AuthenticatorMock
	defaultExpectation *AuthenticatorMockAuthenticateExpectation
	expectations       []*AuthenticatorMockAuthenticateExpectation

	callArgs []*AuthenticatorMockAuthenticateParams
	mutex    sync.RWMutex
}

// AuthenticatorMockAuthenticateExpectation specifies expectation struct of the Authenticator.Authenticate
type AuthenticatorMockAuthenticateExpectation struct {
	mock    *AuthenticatorMock
	params  *AuthenticatorMockAuthenticateParams
	results *AuthenticatorMockAuthenticateResults
	Counter uint64
}

// AuthenticatorMockAuthenticateParams contains parameters of the Authenticator.Authenticate
type AuthenticatorMockAuthenticateParams struct {
	u1 Update
}

// AuthenticatorMockAuthenticateResults contains results of the Authenticator.Authenticate
type AuthenticatorMockAuthenticateResults struct {
	u2  User
	err error
}

// Expect sets up expected params for Authenticator.Authenticate
func (mmAuthenticate *mAuthenticatorMockAuthenticate) Expect(u1 Update) *mAuthenticatorMockAuthenticate {
	if mmAuthenticate.mock.funcAuthenticate != nil {
		mmAuthenticate.mock.t.Fatalf("AuthenticatorMock.Authenticate mock is already set by Set")
	}

	if mmAuthenticate.defaultExpectation == nil {
		mmAuthenticate.defaultExpectation = &AuthenticatorMockAuthenticateExpectation{}
	}

	mmAuthenticate.defaultExpectation.params = &AuthenticatorMockAuthenticateParams{u1}
	for _, e := range mmAuthenticate.expectations {
		if minimock.Equal(e.params, mmAuthenticate.defaultExpectation.params) {
			mmAuthenticate.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmAuthenticate.defaultExpectation.params)
		}
	}

	return mmAuthenticate
}

// Inspect accepts an inspector function that has same arguments as the Authenticator.Authenticate
func (mmAuthenticate *mAuthenticatorMockAuthenticate) Inspect(f func(u1 Update)) *mAuthenticatorMockAuthenticate {
	if mmAuthenticate.mock.inspectFuncAuthenticate != nil {
		mmAuthenticate.mock.t.Fatalf("Inspect function is already set for AuthenticatorMock.Authenticate")
	}

	mmAuthenticate.mock.inspectFuncAuthenticate = f

	return mmAuthenticate
}

// Return sets up results that will be returned by Authenticator.Authenticate
func (mmAuthenticate *mAuthenticatorMockAuthenticate) Return(u2 User, err error) *AuthenticatorMock {
	if mmAuthenticate.mock.funcAuthenticate != nil {
		mmAuthenticate.mock.t.Fatalf("AuthenticatorMock.Authenticate mock is already set by Set")
	}

	if mmAuthenticate.defaultExpectation == nil {
		mmAuthenticate.defaultExpectation = &AuthenticatorMockAuthenticateExpectation{mock: mmAuthenticate.mock}
	}
	mmAuthenticate.defaultExpectation.results = &AuthenticatorMockAuthenticateResults{u2, err}
	return mmAuthenticate.mock
}

//Set uses given function f to mock the Authenticator.Authenticate method
func (mmAuthenticate *mAuthenticatorMockAuthenticate) Set(f func(u1 Update) (u2 User, err error)) *AuthenticatorMock {
	if mmAuthenticate.defaultExpectation != nil {
		mmAuthenticate.mock.t.Fatalf("Default expectation is already set for the Authenticator.Authenticate method")
	}

	if len(mmAuthenticate.expectations) > 0 {
		mmAuthenticate.mock.t.Fatalf("Some expectations are already set for the Authenticator.Authenticate method")
	}

	mmAuthenticate.mock.funcAuthenticate = f
	return mmAuthenticate.mock
}

// When sets expectation for the Authenticator.Authenticate which will trigger the result defined by the following
// Then helper
func (mmAuthenticate *mAuthenticatorMockAuthenticate) When(u1 Update) *AuthenticatorMockAuthenticateExpectation {
	if mmAuthenticate.mock.funcAuthenticate != nil {
		mmAuthenticate.mock.t.Fatalf("AuthenticatorMock.Authenticate mock is already set by Set")
	}

	expectation := &AuthenticatorMockAuthenticateExpectation{
		mock:   mmAuthenticate.mock,
		params: &AuthenticatorMockAuthenticateParams{u1},
	}
	mmAuthenticate.expectations = append(mmAuthenticate.expectations, expectation)
	return expectation
}

// Then sets up Authenticator.Authenticate return parameters for the expectation previously defined by the When method
func (e *AuthenticatorMockAuthenticateExpectation) Then(u2 User, err error) *AuthenticatorMock {
	e.results = &AuthenticatorMockAuthenticateResults{u2, err}
	return e.mock
}

// Authenticate implements Authenticator
func (mmAuthenticate *AuthenticatorMock) Authenticate(u1 Update) (u2 User, err error) {
	mm_atomic.AddUint64(&mmAuthenticate.beforeAuthenticateCounter, 1)
	defer mm_atomic.AddUint64(&mmAuthenticate.afterAuthenticateCounter, 1)

	if mmAuthenticate.inspectFuncAuthenticate != nil {
		mmAuthenticate.inspectFuncAuthenticate(u1)
	}

	mm_params := &AuthenticatorMockAuthenticateParams{u1}

	// Record call args
	mmAuthenticate.AuthenticateMock.mutex.Lock()
	mmAuthenticate.AuthenticateMock.callArgs = append(mmAuthenticate.AuthenticateMock.callArgs, mm_params)
	mmAuthenticate.AuthenticateMock.mutex.Unlock()

	for _, e := range mmAuthenticate.AuthenticateMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.u2, e.results.err
		}
	}

	if mmAuthenticate.AuthenticateMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmAuthenticate.AuthenticateMock.defaultExpectation.Counter, 1)
		mm_want := mmAuthenticate.AuthenticateMock.defaultExpectation.params
		mm_got := AuthenticatorMockAuthenticateParams{u1}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmAuthenticate.t.Errorf("AuthenticatorMock.Authenticate got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmAuthenticate.AuthenticateMock.defaultExpectation.results
		if mm_results == nil {
			mmAuthenticate.t.Fatal("No results are set for the AuthenticatorMock.Authenticate")
		}
		return (*mm_results).u2, (*mm_results).err
	}
	if mmAuthenticate.funcAuthenticate != nil {
		return mmAuthenticate.funcAuthenticate(u1)
	}
	mmAuthenticate.t.Fatalf("Unexpected call to AuthenticatorMock.Authenticate. %v", u1)
	return
}

// AuthenticateAfterCounter returns a count of finished AuthenticatorMock.Authenticate invocations
func (mmAuthenticate *AuthenticatorMock) AuthenticateAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmAuthenticate.afterAuthenticateCounter)
}

// AuthenticateBeforeCounter returns a count of AuthenticatorMock.Authenticate invocations
func (mmAuthenticate *AuthenticatorMock) AuthenticateBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmAuthenticate.beforeAuthenticateCounter)
}

// Calls returns a list of arguments used in each call to AuthenticatorMock.Authenticate.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmAuthenticate *mAuthenticatorMockAuthenticate) Calls() []*AuthenticatorMockAuthenticateParams {
	mmAuthenticate.mutex.RLock()

	argCopy := make([]*AuthenticatorMockAuthenticateParams, len(mmAuthenticate.callArgs))
	copy(argCopy, mmAuthenticate.callArgs)

	mmAuthenticate.mutex.RUnlock()

	return argCopy
}

// MinimockAuthenticateDone returns true if the count of the Authenticate invocations corresponds
// the number of defined expectations
func (m *AuthenticatorMock) MinimockAuthenticateDone() bool {
	for _, e := range m.AuthenticateMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.AuthenticateMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterAuthenticateCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcAuthenticate != nil && mm_atomic.LoadUint64(&m.afterAuthenticateCounter) < 1 {
		return false
	}
	return true
}

// MinimockAuthenticateInspect logs each unmet expectation
func (m *AuthenticatorMock) MinimockAuthenticateInspect() {
	for _, e := range m.AuthenticateMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to AuthenticatorMock.Authenticate with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.AuthenticateMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterAuthenticateCounter) < 1 {
		if m.AuthenticateMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to AuthenticatorMock.Authenticate")
		} else {
			m.t.Errorf("Expected call to AuthenticatorMock.Authenticate with params: %#v", *m.AuthenticateMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcAuthenticate != nil && mm_atomic.LoadUint64(&m.afterAuthenticateCounter) < 1 {
		m.t.Error("Expected call to AuthenticatorMock.Authenticate")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *AuthenticatorMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockAuthenticateInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *AuthenticatorMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *AuthenticatorMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockAuthenticateDone()
}