package app

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

//go:generate minimock -i homework-2/internal/app.Repository -o ./repository_mock_test.go -n RepositoryMock

import (
	"context"
	pb "homework-2/pkg/api"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
)

// RepositoryMock implements Repository
type RepositoryMock struct {
	t minimock.Tester

	funcAddShare          func(ctx context.Context, shareReq *pb.Share) (err error)
	inspectFuncAddShare   func(ctx context.Context, shareReq *pb.Share)
	afterAddShareCounter  uint64
	beforeAddShareCounter uint64
	AddShareMock          mRepositoryMockAddShare

	funcGetShare          func(ctx context.Context, shareReq *pb.Share) (shareRes *pb.Share, err error)
	inspectFuncGetShare   func(ctx context.Context, shareReq *pb.Share)
	afterGetShareCounter  uint64
	beforeGetShareCounter uint64
	GetShareMock          mRepositoryMockGetShare
}

// NewRepositoryMock returns a mock for Repository
func NewRepositoryMock(t minimock.Tester) *RepositoryMock {
	m := &RepositoryMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.AddShareMock = mRepositoryMockAddShare{mock: m}
	m.AddShareMock.callArgs = []*RepositoryMockAddShareParams{}

	m.GetShareMock = mRepositoryMockGetShare{mock: m}
	m.GetShareMock.callArgs = []*RepositoryMockGetShareParams{}

	return m
}

type mRepositoryMockAddShare struct {
	mock               *RepositoryMock
	defaultExpectation *RepositoryMockAddShareExpectation
	expectations       []*RepositoryMockAddShareExpectation

	callArgs []*RepositoryMockAddShareParams
	mutex    sync.RWMutex
}

// RepositoryMockAddShareExpectation specifies expectation struct of the Repository.AddShare
type RepositoryMockAddShareExpectation struct {
	mock    *RepositoryMock
	params  *RepositoryMockAddShareParams
	results *RepositoryMockAddShareResults
	Counter uint64
}

// RepositoryMockAddShareParams contains parameters of the Repository.AddShare
type RepositoryMockAddShareParams struct {
	ctx      context.Context
	shareReq *pb.Share
}

// RepositoryMockAddShareResults contains results of the Repository.AddShare
type RepositoryMockAddShareResults struct {
	err error
}

// Expect sets up expected params for Repository.AddShare
func (mmAddShare *mRepositoryMockAddShare) Expect(ctx context.Context, shareReq *pb.Share) *mRepositoryMockAddShare {
	if mmAddShare.mock.funcAddShare != nil {
		mmAddShare.mock.t.Fatalf("RepositoryMock.AddShare mock is already set by Set")
	}

	if mmAddShare.defaultExpectation == nil {
		mmAddShare.defaultExpectation = &RepositoryMockAddShareExpectation{}
	}

	mmAddShare.defaultExpectation.params = &RepositoryMockAddShareParams{ctx, shareReq}
	for _, e := range mmAddShare.expectations {
		if minimock.Equal(e.params, mmAddShare.defaultExpectation.params) {
			mmAddShare.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmAddShare.defaultExpectation.params)
		}
	}

	return mmAddShare
}

// Inspect accepts an inspector function that has same arguments as the Repository.AddShare
func (mmAddShare *mRepositoryMockAddShare) Inspect(f func(ctx context.Context, shareReq *pb.Share)) *mRepositoryMockAddShare {
	if mmAddShare.mock.inspectFuncAddShare != nil {
		mmAddShare.mock.t.Fatalf("Inspect function is already set for RepositoryMock.AddShare")
	}

	mmAddShare.mock.inspectFuncAddShare = f

	return mmAddShare
}

// Return sets up results that will be returned by Repository.AddShare
func (mmAddShare *mRepositoryMockAddShare) Return(err error) *RepositoryMock {
	if mmAddShare.mock.funcAddShare != nil {
		mmAddShare.mock.t.Fatalf("RepositoryMock.AddShare mock is already set by Set")
	}

	if mmAddShare.defaultExpectation == nil {
		mmAddShare.defaultExpectation = &RepositoryMockAddShareExpectation{mock: mmAddShare.mock}
	}
	mmAddShare.defaultExpectation.results = &RepositoryMockAddShareResults{err}
	return mmAddShare.mock
}

//Set uses given function f to mock the Repository.AddShare method
func (mmAddShare *mRepositoryMockAddShare) Set(f func(ctx context.Context, shareReq *pb.Share) (err error)) *RepositoryMock {
	if mmAddShare.defaultExpectation != nil {
		mmAddShare.mock.t.Fatalf("Default expectation is already set for the Repository.AddShare method")
	}

	if len(mmAddShare.expectations) > 0 {
		mmAddShare.mock.t.Fatalf("Some expectations are already set for the Repository.AddShare method")
	}

	mmAddShare.mock.funcAddShare = f
	return mmAddShare.mock
}

// When sets expectation for the Repository.AddShare which will trigger the result defined by the following
// Then helper
func (mmAddShare *mRepositoryMockAddShare) When(ctx context.Context, shareReq *pb.Share) *RepositoryMockAddShareExpectation {
	if mmAddShare.mock.funcAddShare != nil {
		mmAddShare.mock.t.Fatalf("RepositoryMock.AddShare mock is already set by Set")
	}

	expectation := &RepositoryMockAddShareExpectation{
		mock:   mmAddShare.mock,
		params: &RepositoryMockAddShareParams{ctx, shareReq},
	}
	mmAddShare.expectations = append(mmAddShare.expectations, expectation)
	return expectation
}

// Then sets up Repository.AddShare return parameters for the expectation previously defined by the When method
func (e *RepositoryMockAddShareExpectation) Then(err error) *RepositoryMock {
	e.results = &RepositoryMockAddShareResults{err}
	return e.mock
}

// AddShare implements Repository
func (mmAddShare *RepositoryMock) AddShare(ctx context.Context, shareReq *pb.Share) (err error) {
	mm_atomic.AddUint64(&mmAddShare.beforeAddShareCounter, 1)
	defer mm_atomic.AddUint64(&mmAddShare.afterAddShareCounter, 1)

	if mmAddShare.inspectFuncAddShare != nil {
		mmAddShare.inspectFuncAddShare(ctx, shareReq)
	}

	mm_params := &RepositoryMockAddShareParams{ctx, shareReq}

	// Record call args
	mmAddShare.AddShareMock.mutex.Lock()
	mmAddShare.AddShareMock.callArgs = append(mmAddShare.AddShareMock.callArgs, mm_params)
	mmAddShare.AddShareMock.mutex.Unlock()

	for _, e := range mmAddShare.AddShareMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmAddShare.AddShareMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmAddShare.AddShareMock.defaultExpectation.Counter, 1)
		mm_want := mmAddShare.AddShareMock.defaultExpectation.params
		mm_got := RepositoryMockAddShareParams{ctx, shareReq}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmAddShare.t.Errorf("RepositoryMock.AddShare got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmAddShare.AddShareMock.defaultExpectation.results
		if mm_results == nil {
			mmAddShare.t.Fatal("No results are set for the RepositoryMock.AddShare")
		}
		return (*mm_results).err
	}
	if mmAddShare.funcAddShare != nil {
		return mmAddShare.funcAddShare(ctx, shareReq)
	}
	mmAddShare.t.Fatalf("Unexpected call to RepositoryMock.AddShare. %v %v", ctx, shareReq)
	return
}

// AddShareAfterCounter returns a count of finished RepositoryMock.AddShare invocations
func (mmAddShare *RepositoryMock) AddShareAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmAddShare.afterAddShareCounter)
}

// AddShareBeforeCounter returns a count of RepositoryMock.AddShare invocations
func (mmAddShare *RepositoryMock) AddShareBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmAddShare.beforeAddShareCounter)
}

// Calls returns a list of arguments used in each call to RepositoryMock.AddShare.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmAddShare *mRepositoryMockAddShare) Calls() []*RepositoryMockAddShareParams {
	mmAddShare.mutex.RLock()

	argCopy := make([]*RepositoryMockAddShareParams, len(mmAddShare.callArgs))
	copy(argCopy, mmAddShare.callArgs)

	mmAddShare.mutex.RUnlock()

	return argCopy
}

// MinimockAddShareDone returns true if the count of the AddShare invocations corresponds
// the number of defined expectations
func (m *RepositoryMock) MinimockAddShareDone() bool {
	for _, e := range m.AddShareMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.AddShareMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterAddShareCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcAddShare != nil && mm_atomic.LoadUint64(&m.afterAddShareCounter) < 1 {
		return false
	}
	return true
}

// MinimockAddShareInspect logs each unmet expectation
func (m *RepositoryMock) MinimockAddShareInspect() {
	for _, e := range m.AddShareMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to RepositoryMock.AddShare with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.AddShareMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterAddShareCounter) < 1 {
		if m.AddShareMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to RepositoryMock.AddShare")
		} else {
			m.t.Errorf("Expected call to RepositoryMock.AddShare with params: %#v", *m.AddShareMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcAddShare != nil && mm_atomic.LoadUint64(&m.afterAddShareCounter) < 1 {
		m.t.Error("Expected call to RepositoryMock.AddShare")
	}
}

type mRepositoryMockGetShare struct {
	mock               *RepositoryMock
	defaultExpectation *RepositoryMockGetShareExpectation
	expectations       []*RepositoryMockGetShareExpectation

	callArgs []*RepositoryMockGetShareParams
	mutex    sync.RWMutex
}

// RepositoryMockGetShareExpectation specifies expectation struct of the Repository.GetShare
type RepositoryMockGetShareExpectation struct {
	mock    *RepositoryMock
	params  *RepositoryMockGetShareParams
	results *RepositoryMockGetShareResults
	Counter uint64
}

// RepositoryMockGetShareParams contains parameters of the Repository.GetShare
type RepositoryMockGetShareParams struct {
	ctx      context.Context
	shareReq *pb.Share
}

// RepositoryMockGetShareResults contains results of the Repository.GetShare
type RepositoryMockGetShareResults struct {
	shareRes *pb.Share
	err      error
}

// Expect sets up expected params for Repository.GetShare
func (mmGetShare *mRepositoryMockGetShare) Expect(ctx context.Context, shareReq *pb.Share) *mRepositoryMockGetShare {
	if mmGetShare.mock.funcGetShare != nil {
		mmGetShare.mock.t.Fatalf("RepositoryMock.GetShare mock is already set by Set")
	}

	if mmGetShare.defaultExpectation == nil {
		mmGetShare.defaultExpectation = &RepositoryMockGetShareExpectation{}
	}

	mmGetShare.defaultExpectation.params = &RepositoryMockGetShareParams{ctx, shareReq}
	for _, e := range mmGetShare.expectations {
		if minimock.Equal(e.params, mmGetShare.defaultExpectation.params) {
			mmGetShare.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmGetShare.defaultExpectation.params)
		}
	}

	return mmGetShare
}

// Inspect accepts an inspector function that has same arguments as the Repository.GetShare
func (mmGetShare *mRepositoryMockGetShare) Inspect(f func(ctx context.Context, shareReq *pb.Share)) *mRepositoryMockGetShare {
	if mmGetShare.mock.inspectFuncGetShare != nil {
		mmGetShare.mock.t.Fatalf("Inspect function is already set for RepositoryMock.GetShare")
	}

	mmGetShare.mock.inspectFuncGetShare = f

	return mmGetShare
}

// Return sets up results that will be returned by Repository.GetShare
func (mmGetShare *mRepositoryMockGetShare) Return(shareRes *pb.Share, err error) *RepositoryMock {
	if mmGetShare.mock.funcGetShare != nil {
		mmGetShare.mock.t.Fatalf("RepositoryMock.GetShare mock is already set by Set")
	}

	if mmGetShare.defaultExpectation == nil {
		mmGetShare.defaultExpectation = &RepositoryMockGetShareExpectation{mock: mmGetShare.mock}
	}
	mmGetShare.defaultExpectation.results = &RepositoryMockGetShareResults{shareRes, err}
	return mmGetShare.mock
}

//Set uses given function f to mock the Repository.GetShare method
func (mmGetShare *mRepositoryMockGetShare) Set(f func(ctx context.Context, shareReq *pb.Share) (shareRes *pb.Share, err error)) *RepositoryMock {
	if mmGetShare.defaultExpectation != nil {
		mmGetShare.mock.t.Fatalf("Default expectation is already set for the Repository.GetShare method")
	}

	if len(mmGetShare.expectations) > 0 {
		mmGetShare.mock.t.Fatalf("Some expectations are already set for the Repository.GetShare method")
	}

	mmGetShare.mock.funcGetShare = f
	return mmGetShare.mock
}

// When sets expectation for the Repository.GetShare which will trigger the result defined by the following
// Then helper
func (mmGetShare *mRepositoryMockGetShare) When(ctx context.Context, shareReq *pb.Share) *RepositoryMockGetShareExpectation {
	if mmGetShare.mock.funcGetShare != nil {
		mmGetShare.mock.t.Fatalf("RepositoryMock.GetShare mock is already set by Set")
	}

	expectation := &RepositoryMockGetShareExpectation{
		mock:   mmGetShare.mock,
		params: &RepositoryMockGetShareParams{ctx, shareReq},
	}
	mmGetShare.expectations = append(mmGetShare.expectations, expectation)
	return expectation
}

// Then sets up Repository.GetShare return parameters for the expectation previously defined by the When method
func (e *RepositoryMockGetShareExpectation) Then(shareRes *pb.Share, err error) *RepositoryMock {
	e.results = &RepositoryMockGetShareResults{shareRes, err}
	return e.mock
}

// GetShare implements Repository
func (mmGetShare *RepositoryMock) GetShare(ctx context.Context, shareReq *pb.Share) (shareRes *pb.Share, err error) {
	mm_atomic.AddUint64(&mmGetShare.beforeGetShareCounter, 1)
	defer mm_atomic.AddUint64(&mmGetShare.afterGetShareCounter, 1)

	if mmGetShare.inspectFuncGetShare != nil {
		mmGetShare.inspectFuncGetShare(ctx, shareReq)
	}

	mm_params := &RepositoryMockGetShareParams{ctx, shareReq}

	// Record call args
	mmGetShare.GetShareMock.mutex.Lock()
	mmGetShare.GetShareMock.callArgs = append(mmGetShare.GetShareMock.callArgs, mm_params)
	mmGetShare.GetShareMock.mutex.Unlock()

	for _, e := range mmGetShare.GetShareMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.shareRes, e.results.err
		}
	}

	if mmGetShare.GetShareMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmGetShare.GetShareMock.defaultExpectation.Counter, 1)
		mm_want := mmGetShare.GetShareMock.defaultExpectation.params
		mm_got := RepositoryMockGetShareParams{ctx, shareReq}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmGetShare.t.Errorf("RepositoryMock.GetShare got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmGetShare.GetShareMock.defaultExpectation.results
		if mm_results == nil {
			mmGetShare.t.Fatal("No results are set for the RepositoryMock.GetShare")
		}
		return (*mm_results).shareRes, (*mm_results).err
	}
	if mmGetShare.funcGetShare != nil {
		return mmGetShare.funcGetShare(ctx, shareReq)
	}
	mmGetShare.t.Fatalf("Unexpected call to RepositoryMock.GetShare. %v %v", ctx, shareReq)
	return
}

// GetShareAfterCounter returns a count of finished RepositoryMock.GetShare invocations
func (mmGetShare *RepositoryMock) GetShareAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetShare.afterGetShareCounter)
}

// GetShareBeforeCounter returns a count of RepositoryMock.GetShare invocations
func (mmGetShare *RepositoryMock) GetShareBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetShare.beforeGetShareCounter)
}

// Calls returns a list of arguments used in each call to RepositoryMock.GetShare.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmGetShare *mRepositoryMockGetShare) Calls() []*RepositoryMockGetShareParams {
	mmGetShare.mutex.RLock()

	argCopy := make([]*RepositoryMockGetShareParams, len(mmGetShare.callArgs))
	copy(argCopy, mmGetShare.callArgs)

	mmGetShare.mutex.RUnlock()

	return argCopy
}

// MinimockGetShareDone returns true if the count of the GetShare invocations corresponds
// the number of defined expectations
func (m *RepositoryMock) MinimockGetShareDone() bool {
	for _, e := range m.GetShareMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetShareMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetShareCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetShare != nil && mm_atomic.LoadUint64(&m.afterGetShareCounter) < 1 {
		return false
	}
	return true
}

// MinimockGetShareInspect logs each unmet expectation
func (m *RepositoryMock) MinimockGetShareInspect() {
	for _, e := range m.GetShareMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to RepositoryMock.GetShare with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetShareMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetShareCounter) < 1 {
		if m.GetShareMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to RepositoryMock.GetShare")
		} else {
			m.t.Errorf("Expected call to RepositoryMock.GetShare with params: %#v", *m.GetShareMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetShare != nil && mm_atomic.LoadUint64(&m.afterGetShareCounter) < 1 {
		m.t.Error("Expected call to RepositoryMock.GetShare")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *RepositoryMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockAddShareInspect()

		m.MinimockGetShareInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *RepositoryMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *RepositoryMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockAddShareDone() &&
		m.MinimockGetShareDone()
}