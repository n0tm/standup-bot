package telegram

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

//go:generate minimock -i standup-bot/src/telegram.User -o ./user_mock.go -n UserMock

import (
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
)

// UserMock implements User
type UserMock struct {
	t minimock.Tester

	funcGetId          func() (u1 uint)
	inspectFuncGetId   func()
	afterGetIdCounter  uint64
	beforeGetIdCounter uint64
	GetIdMock          mUserMockGetId

	funcSetId          func(u1 uint)
	inspectFuncSetId   func(u1 uint)
	afterSetIdCounter  uint64
	beforeSetIdCounter uint64
	SetIdMock          mUserMockSetId
}

// NewUserMock returns a mock for User
func NewUserMock(t minimock.Tester) *UserMock {
	m := &UserMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.GetIdMock = mUserMockGetId{mock: m}

	m.SetIdMock = mUserMockSetId{mock: m}
	m.SetIdMock.callArgs = []*UserMockSetIdParams{}

	return m
}

type mUserMockGetId struct {
	mock               *UserMock
	defaultExpectation *UserMockGetIdExpectation
	expectations       []*UserMockGetIdExpectation
}

// UserMockGetIdExpectation specifies expectation struct of the User.GetId
type UserMockGetIdExpectation struct {
	mock *UserMock

	results *UserMockGetIdResults
	Counter uint64
}

// UserMockGetIdResults contains results of the User.GetId
type UserMockGetIdResults struct {
	u1 uint
}

// Expect sets up expected params for User.GetId
func (mmGetId *mUserMockGetId) Expect() *mUserMockGetId {
	if mmGetId.mock.funcGetId != nil {
		mmGetId.mock.t.Fatalf("UserMock.GetId mock is already set by Set")
	}

	if mmGetId.defaultExpectation == nil {
		mmGetId.defaultExpectation = &UserMockGetIdExpectation{}
	}

	return mmGetId
}

// Inspect accepts an inspector function that has same arguments as the User.GetId
func (mmGetId *mUserMockGetId) Inspect(f func()) *mUserMockGetId {
	if mmGetId.mock.inspectFuncGetId != nil {
		mmGetId.mock.t.Fatalf("Inspect function is already set for UserMock.GetId")
	}

	mmGetId.mock.inspectFuncGetId = f

	return mmGetId
}

// Return sets up results that will be returned by User.GetId
func (mmGetId *mUserMockGetId) Return(u1 uint) *UserMock {
	if mmGetId.mock.funcGetId != nil {
		mmGetId.mock.t.Fatalf("UserMock.GetId mock is already set by Set")
	}

	if mmGetId.defaultExpectation == nil {
		mmGetId.defaultExpectation = &UserMockGetIdExpectation{mock: mmGetId.mock}
	}
	mmGetId.defaultExpectation.results = &UserMockGetIdResults{u1}
	return mmGetId.mock
}

//Set uses given function f to mock the User.GetId method
func (mmGetId *mUserMockGetId) Set(f func() (u1 uint)) *UserMock {
	if mmGetId.defaultExpectation != nil {
		mmGetId.mock.t.Fatalf("Default expectation is already set for the User.GetId method")
	}

	if len(mmGetId.expectations) > 0 {
		mmGetId.mock.t.Fatalf("Some expectations are already set for the User.GetId method")
	}

	mmGetId.mock.funcGetId = f
	return mmGetId.mock
}

// GetId implements User
func (mmGetId *UserMock) GetId() (u1 uint) {
	mm_atomic.AddUint64(&mmGetId.beforeGetIdCounter, 1)
	defer mm_atomic.AddUint64(&mmGetId.afterGetIdCounter, 1)

	if mmGetId.inspectFuncGetId != nil {
		mmGetId.inspectFuncGetId()
	}

	if mmGetId.GetIdMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmGetId.GetIdMock.defaultExpectation.Counter, 1)

		mm_results := mmGetId.GetIdMock.defaultExpectation.results
		if mm_results == nil {
			mmGetId.t.Fatal("No results are set for the UserMock.GetId")
		}
		return (*mm_results).u1
	}
	if mmGetId.funcGetId != nil {
		return mmGetId.funcGetId()
	}
	mmGetId.t.Fatalf("Unexpected call to UserMock.GetId.")
	return
}

// GetIdAfterCounter returns a count of finished UserMock.GetId invocations
func (mmGetId *UserMock) GetIdAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetId.afterGetIdCounter)
}

// GetIdBeforeCounter returns a count of UserMock.GetId invocations
func (mmGetId *UserMock) GetIdBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetId.beforeGetIdCounter)
}

// MinimockGetIdDone returns true if the count of the GetId invocations corresponds
// the number of defined expectations
func (m *UserMock) MinimockGetIdDone() bool {
	for _, e := range m.GetIdMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetIdMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetIdCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetId != nil && mm_atomic.LoadUint64(&m.afterGetIdCounter) < 1 {
		return false
	}
	return true
}

// MinimockGetIdInspect logs each unmet expectation
func (m *UserMock) MinimockGetIdInspect() {
	for _, e := range m.GetIdMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Error("Expected call to UserMock.GetId")
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetIdMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetIdCounter) < 1 {
		m.t.Error("Expected call to UserMock.GetId")
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetId != nil && mm_atomic.LoadUint64(&m.afterGetIdCounter) < 1 {
		m.t.Error("Expected call to UserMock.GetId")
	}
}

type mUserMockSetId struct {
	mock               *UserMock
	defaultExpectation *UserMockSetIdExpectation
	expectations       []*UserMockSetIdExpectation

	callArgs []*UserMockSetIdParams
	mutex    sync.RWMutex
}

// UserMockSetIdExpectation specifies expectation struct of the User.SetId
type UserMockSetIdExpectation struct {
	mock   *UserMock
	params *UserMockSetIdParams

	Counter uint64
}

// UserMockSetIdParams contains parameters of the User.SetId
type UserMockSetIdParams struct {
	u1 uint
}

// Expect sets up expected params for User.SetId
func (mmSetId *mUserMockSetId) Expect(u1 uint) *mUserMockSetId {
	if mmSetId.mock.funcSetId != nil {
		mmSetId.mock.t.Fatalf("UserMock.SetId mock is already set by Set")
	}

	if mmSetId.defaultExpectation == nil {
		mmSetId.defaultExpectation = &UserMockSetIdExpectation{}
	}

	mmSetId.defaultExpectation.params = &UserMockSetIdParams{u1}
	for _, e := range mmSetId.expectations {
		if minimock.Equal(e.params, mmSetId.defaultExpectation.params) {
			mmSetId.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmSetId.defaultExpectation.params)
		}
	}

	return mmSetId
}

// Inspect accepts an inspector function that has same arguments as the User.SetId
func (mmSetId *mUserMockSetId) Inspect(f func(u1 uint)) *mUserMockSetId {
	if mmSetId.mock.inspectFuncSetId != nil {
		mmSetId.mock.t.Fatalf("Inspect function is already set for UserMock.SetId")
	}

	mmSetId.mock.inspectFuncSetId = f

	return mmSetId
}

// Return sets up results that will be returned by User.SetId
func (mmSetId *mUserMockSetId) Return() *UserMock {
	if mmSetId.mock.funcSetId != nil {
		mmSetId.mock.t.Fatalf("UserMock.SetId mock is already set by Set")
	}

	if mmSetId.defaultExpectation == nil {
		mmSetId.defaultExpectation = &UserMockSetIdExpectation{mock: mmSetId.mock}
	}

	return mmSetId.mock
}

//Set uses given function f to mock the User.SetId method
func (mmSetId *mUserMockSetId) Set(f func(u1 uint)) *UserMock {
	if mmSetId.defaultExpectation != nil {
		mmSetId.mock.t.Fatalf("Default expectation is already set for the User.SetId method")
	}

	if len(mmSetId.expectations) > 0 {
		mmSetId.mock.t.Fatalf("Some expectations are already set for the User.SetId method")
	}

	mmSetId.mock.funcSetId = f
	return mmSetId.mock
}

// SetId implements User
func (mmSetId *UserMock) SetId(u1 uint) {
	mm_atomic.AddUint64(&mmSetId.beforeSetIdCounter, 1)
	defer mm_atomic.AddUint64(&mmSetId.afterSetIdCounter, 1)

	if mmSetId.inspectFuncSetId != nil {
		mmSetId.inspectFuncSetId(u1)
	}

	mm_params := &UserMockSetIdParams{u1}

	// Record call args
	mmSetId.SetIdMock.mutex.Lock()
	mmSetId.SetIdMock.callArgs = append(mmSetId.SetIdMock.callArgs, mm_params)
	mmSetId.SetIdMock.mutex.Unlock()

	for _, e := range mmSetId.SetIdMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return
		}
	}

	if mmSetId.SetIdMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmSetId.SetIdMock.defaultExpectation.Counter, 1)
		mm_want := mmSetId.SetIdMock.defaultExpectation.params
		mm_got := UserMockSetIdParams{u1}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmSetId.t.Errorf("UserMock.SetId got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		return

	}
	if mmSetId.funcSetId != nil {
		mmSetId.funcSetId(u1)
		return
	}
	mmSetId.t.Fatalf("Unexpected call to UserMock.SetId. %v", u1)

}

// SetIdAfterCounter returns a count of finished UserMock.SetId invocations
func (mmSetId *UserMock) SetIdAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmSetId.afterSetIdCounter)
}

// SetIdBeforeCounter returns a count of UserMock.SetId invocations
func (mmSetId *UserMock) SetIdBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmSetId.beforeSetIdCounter)
}

// Calls returns a list of arguments used in each call to UserMock.SetId.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmSetId *mUserMockSetId) Calls() []*UserMockSetIdParams {
	mmSetId.mutex.RLock()

	argCopy := make([]*UserMockSetIdParams, len(mmSetId.callArgs))
	copy(argCopy, mmSetId.callArgs)

	mmSetId.mutex.RUnlock()

	return argCopy
}

// MinimockSetIdDone returns true if the count of the SetId invocations corresponds
// the number of defined expectations
func (m *UserMock) MinimockSetIdDone() bool {
	for _, e := range m.SetIdMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.SetIdMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterSetIdCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcSetId != nil && mm_atomic.LoadUint64(&m.afterSetIdCounter) < 1 {
		return false
	}
	return true
}

// MinimockSetIdInspect logs each unmet expectation
func (m *UserMock) MinimockSetIdInspect() {
	for _, e := range m.SetIdMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to UserMock.SetId with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.SetIdMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterSetIdCounter) < 1 {
		if m.SetIdMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to UserMock.SetId")
		} else {
			m.t.Errorf("Expected call to UserMock.SetId with params: %#v", *m.SetIdMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcSetId != nil && mm_atomic.LoadUint64(&m.afterSetIdCounter) < 1 {
		m.t.Error("Expected call to UserMock.SetId")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *UserMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockGetIdInspect()

		m.MinimockSetIdInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *UserMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *UserMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockGetIdDone() &&
		m.MinimockSetIdDone()
}
