// Code generated by http://github.com/gojuno/minimock (v3.3.12). DO NOT EDIT.

package http

//go:generate minimock -i github.com/etilite/xlsx-builder/internal/delivery/http.Builder -o zzz_builder_mock_test.go -n BuilderMock -p http

import (
	"io"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"github.com/gojuno/minimock/v3"
)

// BuilderMock implements Builder
type BuilderMock struct {
	t          minimock.Tester
	finishOnce sync.Once

	funcBuild          func(r io.Reader, w io.Writer) (err error)
	inspectFuncBuild   func(r io.Reader, w io.Writer)
	afterBuildCounter  uint64
	beforeBuildCounter uint64
	BuildMock          mBuilderMockBuild
}

// NewBuilderMock returns a mock for Builder
func NewBuilderMock(t minimock.Tester) *BuilderMock {
	m := &BuilderMock{t: t}

	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.BuildMock = mBuilderMockBuild{mock: m}
	m.BuildMock.callArgs = []*BuilderMockBuildParams{}

	t.Cleanup(m.MinimockFinish)

	return m
}

type mBuilderMockBuild struct {
	optional           bool
	mock               *BuilderMock
	defaultExpectation *BuilderMockBuildExpectation
	expectations       []*BuilderMockBuildExpectation

	callArgs []*BuilderMockBuildParams
	mutex    sync.RWMutex

	expectedInvocations uint64
}

// BuilderMockBuildExpectation specifies expectation struct of the Builder.Build
type BuilderMockBuildExpectation struct {
	mock      *BuilderMock
	params    *BuilderMockBuildParams
	paramPtrs *BuilderMockBuildParamPtrs
	results   *BuilderMockBuildResults
	Counter   uint64
}

// BuilderMockBuildParams contains parameters of the Builder.Build
type BuilderMockBuildParams struct {
	r io.Reader
	w io.Writer
}

// BuilderMockBuildParamPtrs contains pointers to parameters of the Builder.Build
type BuilderMockBuildParamPtrs struct {
	r *io.Reader
	w *io.Writer
}

// BuilderMockBuildResults contains results of the Builder.Build
type BuilderMockBuildResults struct {
	err error
}

// Marks this method to be optional. The default behavior of any method with Return() is '1 or more', meaning
// the test will fail minimock's automatic final call check if the mocked method was not called at least once.
// Optional() makes method check to work in '0 or more' mode.
// It is NOT RECOMMENDED to use this option unless you really need it, as default behaviour helps to
// catch the problems when the expected method call is totally skipped during test run.
func (mmBuild *mBuilderMockBuild) Optional() *mBuilderMockBuild {
	mmBuild.optional = true
	return mmBuild
}

// Expect sets up expected params for Builder.Build
func (mmBuild *mBuilderMockBuild) Expect(r io.Reader, w io.Writer) *mBuilderMockBuild {
	if mmBuild.mock.funcBuild != nil {
		mmBuild.mock.t.Fatalf("BuilderMock.Build mock is already set by Set")
	}

	if mmBuild.defaultExpectation == nil {
		mmBuild.defaultExpectation = &BuilderMockBuildExpectation{}
	}

	if mmBuild.defaultExpectation.paramPtrs != nil {
		mmBuild.mock.t.Fatalf("BuilderMock.Build mock is already set by ExpectParams functions")
	}

	mmBuild.defaultExpectation.params = &BuilderMockBuildParams{r, w}
	for _, e := range mmBuild.expectations {
		if minimock.Equal(e.params, mmBuild.defaultExpectation.params) {
			mmBuild.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmBuild.defaultExpectation.params)
		}
	}

	return mmBuild
}

// ExpectRParam1 sets up expected param r for Builder.Build
func (mmBuild *mBuilderMockBuild) ExpectRParam1(r io.Reader) *mBuilderMockBuild {
	if mmBuild.mock.funcBuild != nil {
		mmBuild.mock.t.Fatalf("BuilderMock.Build mock is already set by Set")
	}

	if mmBuild.defaultExpectation == nil {
		mmBuild.defaultExpectation = &BuilderMockBuildExpectation{}
	}

	if mmBuild.defaultExpectation.params != nil {
		mmBuild.mock.t.Fatalf("BuilderMock.Build mock is already set by Expect")
	}

	if mmBuild.defaultExpectation.paramPtrs == nil {
		mmBuild.defaultExpectation.paramPtrs = &BuilderMockBuildParamPtrs{}
	}
	mmBuild.defaultExpectation.paramPtrs.r = &r

	return mmBuild
}

// ExpectWParam2 sets up expected param w for Builder.Build
func (mmBuild *mBuilderMockBuild) ExpectWParam2(w io.Writer) *mBuilderMockBuild {
	if mmBuild.mock.funcBuild != nil {
		mmBuild.mock.t.Fatalf("BuilderMock.Build mock is already set by Set")
	}

	if mmBuild.defaultExpectation == nil {
		mmBuild.defaultExpectation = &BuilderMockBuildExpectation{}
	}

	if mmBuild.defaultExpectation.params != nil {
		mmBuild.mock.t.Fatalf("BuilderMock.Build mock is already set by Expect")
	}

	if mmBuild.defaultExpectation.paramPtrs == nil {
		mmBuild.defaultExpectation.paramPtrs = &BuilderMockBuildParamPtrs{}
	}
	mmBuild.defaultExpectation.paramPtrs.w = &w

	return mmBuild
}

// Inspect accepts an inspector function that has same arguments as the Builder.Build
func (mmBuild *mBuilderMockBuild) Inspect(f func(r io.Reader, w io.Writer)) *mBuilderMockBuild {
	if mmBuild.mock.inspectFuncBuild != nil {
		mmBuild.mock.t.Fatalf("Inspect function is already set for BuilderMock.Build")
	}

	mmBuild.mock.inspectFuncBuild = f

	return mmBuild
}

// Return sets up results that will be returned by Builder.Build
func (mmBuild *mBuilderMockBuild) Return(err error) *BuilderMock {
	if mmBuild.mock.funcBuild != nil {
		mmBuild.mock.t.Fatalf("BuilderMock.Build mock is already set by Set")
	}

	if mmBuild.defaultExpectation == nil {
		mmBuild.defaultExpectation = &BuilderMockBuildExpectation{mock: mmBuild.mock}
	}
	mmBuild.defaultExpectation.results = &BuilderMockBuildResults{err}
	return mmBuild.mock
}

// Set uses given function f to mock the Builder.Build method
func (mmBuild *mBuilderMockBuild) Set(f func(r io.Reader, w io.Writer) (err error)) *BuilderMock {
	if mmBuild.defaultExpectation != nil {
		mmBuild.mock.t.Fatalf("Default expectation is already set for the Builder.Build method")
	}

	if len(mmBuild.expectations) > 0 {
		mmBuild.mock.t.Fatalf("Some expectations are already set for the Builder.Build method")
	}

	mmBuild.mock.funcBuild = f
	return mmBuild.mock
}

// When sets expectation for the Builder.Build which will trigger the result defined by the following
// Then helper
func (mmBuild *mBuilderMockBuild) When(r io.Reader, w io.Writer) *BuilderMockBuildExpectation {
	if mmBuild.mock.funcBuild != nil {
		mmBuild.mock.t.Fatalf("BuilderMock.Build mock is already set by Set")
	}

	expectation := &BuilderMockBuildExpectation{
		mock:   mmBuild.mock,
		params: &BuilderMockBuildParams{r, w},
	}
	mmBuild.expectations = append(mmBuild.expectations, expectation)
	return expectation
}

// Then sets up Builder.Build return parameters for the expectation previously defined by the When method
func (e *BuilderMockBuildExpectation) Then(err error) *BuilderMock {
	e.results = &BuilderMockBuildResults{err}
	return e.mock
}

// Times sets number of times Builder.Build should be invoked
func (mmBuild *mBuilderMockBuild) Times(n uint64) *mBuilderMockBuild {
	if n == 0 {
		mmBuild.mock.t.Fatalf("Times of BuilderMock.Build mock can not be zero")
	}
	mm_atomic.StoreUint64(&mmBuild.expectedInvocations, n)
	return mmBuild
}

func (mmBuild *mBuilderMockBuild) invocationsDone() bool {
	if len(mmBuild.expectations) == 0 && mmBuild.defaultExpectation == nil && mmBuild.mock.funcBuild == nil {
		return true
	}

	totalInvocations := mm_atomic.LoadUint64(&mmBuild.mock.afterBuildCounter)
	expectedInvocations := mm_atomic.LoadUint64(&mmBuild.expectedInvocations)

	return totalInvocations > 0 && (expectedInvocations == 0 || expectedInvocations == totalInvocations)
}

// Build implements Builder
func (mmBuild *BuilderMock) Build(r io.Reader, w io.Writer) (err error) {
	mm_atomic.AddUint64(&mmBuild.beforeBuildCounter, 1)
	defer mm_atomic.AddUint64(&mmBuild.afterBuildCounter, 1)

	if mmBuild.inspectFuncBuild != nil {
		mmBuild.inspectFuncBuild(r, w)
	}

	mm_params := BuilderMockBuildParams{r, w}

	// Record call args
	mmBuild.BuildMock.mutex.Lock()
	mmBuild.BuildMock.callArgs = append(mmBuild.BuildMock.callArgs, &mm_params)
	mmBuild.BuildMock.mutex.Unlock()

	for _, e := range mmBuild.BuildMock.expectations {
		if minimock.Equal(*e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.err
		}
	}

	if mmBuild.BuildMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmBuild.BuildMock.defaultExpectation.Counter, 1)
		mm_want := mmBuild.BuildMock.defaultExpectation.params
		mm_want_ptrs := mmBuild.BuildMock.defaultExpectation.paramPtrs

		mm_got := BuilderMockBuildParams{r, w}

		if mm_want_ptrs != nil {

			if mm_want_ptrs.r != nil && !minimock.Equal(*mm_want_ptrs.r, mm_got.r) {
				mmBuild.t.Errorf("BuilderMock.Build got unexpected parameter r, want: %#v, got: %#v%s\n", *mm_want_ptrs.r, mm_got.r, minimock.Diff(*mm_want_ptrs.r, mm_got.r))
			}

			if mm_want_ptrs.w != nil && !minimock.Equal(*mm_want_ptrs.w, mm_got.w) {
				mmBuild.t.Errorf("BuilderMock.Build got unexpected parameter w, want: %#v, got: %#v%s\n", *mm_want_ptrs.w, mm_got.w, minimock.Diff(*mm_want_ptrs.w, mm_got.w))
			}

		} else if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmBuild.t.Errorf("BuilderMock.Build got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmBuild.BuildMock.defaultExpectation.results
		if mm_results == nil {
			mmBuild.t.Fatal("No results are set for the BuilderMock.Build")
		}
		return (*mm_results).err
	}
	if mmBuild.funcBuild != nil {
		return mmBuild.funcBuild(r, w)
	}
	mmBuild.t.Fatalf("Unexpected call to BuilderMock.Build. %v %v", r, w)
	return
}

// BuildAfterCounter returns a count of finished BuilderMock.Build invocations
func (mmBuild *BuilderMock) BuildAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmBuild.afterBuildCounter)
}

// BuildBeforeCounter returns a count of BuilderMock.Build invocations
func (mmBuild *BuilderMock) BuildBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmBuild.beforeBuildCounter)
}

// Calls returns a list of arguments used in each call to BuilderMock.Build.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmBuild *mBuilderMockBuild) Calls() []*BuilderMockBuildParams {
	mmBuild.mutex.RLock()

	argCopy := make([]*BuilderMockBuildParams, len(mmBuild.callArgs))
	copy(argCopy, mmBuild.callArgs)

	mmBuild.mutex.RUnlock()

	return argCopy
}

// MinimockBuildDone returns true if the count of the Build invocations corresponds
// the number of defined expectations
func (m *BuilderMock) MinimockBuildDone() bool {
	if m.BuildMock.optional {
		// Optional methods provide '0 or more' call count restriction.
		return true
	}

	for _, e := range m.BuildMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	return m.BuildMock.invocationsDone()
}

// MinimockBuildInspect logs each unmet expectation
func (m *BuilderMock) MinimockBuildInspect() {
	for _, e := range m.BuildMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to BuilderMock.Build with params: %#v", *e.params)
		}
	}

	afterBuildCounter := mm_atomic.LoadUint64(&m.afterBuildCounter)
	// if default expectation was set then invocations count should be greater than zero
	if m.BuildMock.defaultExpectation != nil && afterBuildCounter < 1 {
		if m.BuildMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to BuilderMock.Build")
		} else {
			m.t.Errorf("Expected call to BuilderMock.Build with params: %#v", *m.BuildMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcBuild != nil && afterBuildCounter < 1 {
		m.t.Error("Expected call to BuilderMock.Build")
	}

	if !m.BuildMock.invocationsDone() && afterBuildCounter > 0 {
		m.t.Errorf("Expected %d calls to BuilderMock.Build but found %d calls",
			mm_atomic.LoadUint64(&m.BuildMock.expectedInvocations), afterBuildCounter)
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *BuilderMock) MinimockFinish() {
	m.finishOnce.Do(func() {
		if !m.minimockDone() {
			m.MinimockBuildInspect()
		}
	})
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *BuilderMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *BuilderMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockBuildDone()
}