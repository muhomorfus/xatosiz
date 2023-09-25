package mocks

// Code generated by http://github.com/gojuno/minimock (dev). DO NOT EDIT.

//go:generate minimock -i git.iu7.bmstu.ru/kav20u129/ppo/backend/internal/ports.GroupRepository -o ./internal/mocks/group_repository_mock.go -n GroupRepositoryMock

import (
	"context"
	"sync"
	mm_atomic "sync/atomic"
	mm_time "time"

	"git.iu7.bmstu.ru/kav20u129/ppo/backend/internal/domain/models"
	"github.com/gojuno/minimock/v3"
	"github.com/google/uuid"
)

// GroupRepositoryMock implements ports.GroupRepository
type GroupRepositoryMock struct {
	t minimock.Tester

	funcCreate          func(ctx context.Context) (gp1 *models.Group, err error)
	inspectFuncCreate   func(ctx context.Context)
	afterCreateCounter  uint64
	beforeCreateCounter uint64
	CreateMock          mGroupRepositoryMockCreate

	funcGet          func(ctx context.Context, id uuid.UUID) (gp1 *models.Group, err error)
	inspectFuncGet   func(ctx context.Context, id uuid.UUID)
	afterGetCounter  uint64
	beforeGetCounter uint64
	GetMock          mGroupRepositoryMockGet

	funcGetList          func(ctx context.Context, filters models.Filters) (gpa1 []*models.GroupPreview, err error)
	inspectFuncGetList   func(ctx context.Context, filters models.Filters)
	afterGetListCounter  uint64
	beforeGetListCounter uint64
	GetListMock          mGroupRepositoryMockGetList
}

// NewGroupRepositoryMock returns a mock for ports.GroupRepository
func NewGroupRepositoryMock(t minimock.Tester) *GroupRepositoryMock {
	m := &GroupRepositoryMock{t: t}
	if controller, ok := t.(minimock.MockController); ok {
		controller.RegisterMocker(m)
	}

	m.CreateMock = mGroupRepositoryMockCreate{mock: m}
	m.CreateMock.callArgs = []*GroupRepositoryMockCreateParams{}

	m.GetMock = mGroupRepositoryMockGet{mock: m}
	m.GetMock.callArgs = []*GroupRepositoryMockGetParams{}

	m.GetListMock = mGroupRepositoryMockGetList{mock: m}
	m.GetListMock.callArgs = []*GroupRepositoryMockGetListParams{}

	return m
}

type mGroupRepositoryMockCreate struct {
	mock               *GroupRepositoryMock
	defaultExpectation *GroupRepositoryMockCreateExpectation
	expectations       []*GroupRepositoryMockCreateExpectation

	callArgs []*GroupRepositoryMockCreateParams
	mutex    sync.RWMutex
}

// GroupRepositoryMockCreateExpectation specifies expectation struct of the GroupRepository.Create
type GroupRepositoryMockCreateExpectation struct {
	mock    *GroupRepositoryMock
	params  *GroupRepositoryMockCreateParams
	results *GroupRepositoryMockCreateResults
	Counter uint64
}

// GroupRepositoryMockCreateParams contains parameters of the GroupRepository.Create
type GroupRepositoryMockCreateParams struct {
	ctx context.Context
}

// GroupRepositoryMockCreateResults contains results of the GroupRepository.Create
type GroupRepositoryMockCreateResults struct {
	gp1 *models.Group
	err error
}

// Expect sets up expected params for GroupRepository.Create
func (mmCreate *mGroupRepositoryMockCreate) Expect(ctx context.Context) *mGroupRepositoryMockCreate {
	if mmCreate.mock.funcCreate != nil {
		mmCreate.mock.t.Fatalf("GroupRepositoryMock.Create mock is already set by Set")
	}

	if mmCreate.defaultExpectation == nil {
		mmCreate.defaultExpectation = &GroupRepositoryMockCreateExpectation{}
	}

	mmCreate.defaultExpectation.params = &GroupRepositoryMockCreateParams{ctx}
	for _, e := range mmCreate.expectations {
		if minimock.Equal(e.params, mmCreate.defaultExpectation.params) {
			mmCreate.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmCreate.defaultExpectation.params)
		}
	}

	return mmCreate
}

// Inspect accepts an inspector function that has same arguments as the GroupRepository.Create
func (mmCreate *mGroupRepositoryMockCreate) Inspect(f func(ctx context.Context)) *mGroupRepositoryMockCreate {
	if mmCreate.mock.inspectFuncCreate != nil {
		mmCreate.mock.t.Fatalf("Inspect function is already set for GroupRepositoryMock.Create")
	}

	mmCreate.mock.inspectFuncCreate = f

	return mmCreate
}

// Return sets up results that will be returned by GroupRepository.Create
func (mmCreate *mGroupRepositoryMockCreate) Return(gp1 *models.Group, err error) *GroupRepositoryMock {
	if mmCreate.mock.funcCreate != nil {
		mmCreate.mock.t.Fatalf("GroupRepositoryMock.Create mock is already set by Set")
	}

	if mmCreate.defaultExpectation == nil {
		mmCreate.defaultExpectation = &GroupRepositoryMockCreateExpectation{mock: mmCreate.mock}
	}
	mmCreate.defaultExpectation.results = &GroupRepositoryMockCreateResults{gp1, err}
	return mmCreate.mock
}

// Set uses given function f to mock the GroupRepository.Create method
func (mmCreate *mGroupRepositoryMockCreate) Set(f func(ctx context.Context) (gp1 *models.Group, err error)) *GroupRepositoryMock {
	if mmCreate.defaultExpectation != nil {
		mmCreate.mock.t.Fatalf("Default expectation is already set for the GroupRepository.Create method")
	}

	if len(mmCreate.expectations) > 0 {
		mmCreate.mock.t.Fatalf("Some expectations are already set for the GroupRepository.Create method")
	}

	mmCreate.mock.funcCreate = f
	return mmCreate.mock
}

// When sets expectation for the GroupRepository.Create which will trigger the result defined by the following
// Then helper
func (mmCreate *mGroupRepositoryMockCreate) When(ctx context.Context) *GroupRepositoryMockCreateExpectation {
	if mmCreate.mock.funcCreate != nil {
		mmCreate.mock.t.Fatalf("GroupRepositoryMock.Create mock is already set by Set")
	}

	expectation := &GroupRepositoryMockCreateExpectation{
		mock:   mmCreate.mock,
		params: &GroupRepositoryMockCreateParams{ctx},
	}
	mmCreate.expectations = append(mmCreate.expectations, expectation)
	return expectation
}

// Then sets up GroupRepository.Create return parameters for the expectation previously defined by the When method
func (e *GroupRepositoryMockCreateExpectation) Then(gp1 *models.Group, err error) *GroupRepositoryMock {
	e.results = &GroupRepositoryMockCreateResults{gp1, err}
	return e.mock
}

// Create implements ports.GroupRepository
func (mmCreate *GroupRepositoryMock) Create(ctx context.Context) (gp1 *models.Group, err error) {
	mm_atomic.AddUint64(&mmCreate.beforeCreateCounter, 1)
	defer mm_atomic.AddUint64(&mmCreate.afterCreateCounter, 1)

	if mmCreate.inspectFuncCreate != nil {
		mmCreate.inspectFuncCreate(ctx)
	}

	mm_params := &GroupRepositoryMockCreateParams{ctx}

	// Record call args
	mmCreate.CreateMock.mutex.Lock()
	mmCreate.CreateMock.callArgs = append(mmCreate.CreateMock.callArgs, mm_params)
	mmCreate.CreateMock.mutex.Unlock()

	for _, e := range mmCreate.CreateMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.gp1, e.results.err
		}
	}

	if mmCreate.CreateMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmCreate.CreateMock.defaultExpectation.Counter, 1)
		mm_want := mmCreate.CreateMock.defaultExpectation.params
		mm_got := GroupRepositoryMockCreateParams{ctx}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmCreate.t.Errorf("GroupRepositoryMock.Create got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmCreate.CreateMock.defaultExpectation.results
		if mm_results == nil {
			mmCreate.t.Fatal("No results are set for the GroupRepositoryMock.Create")
		}
		return (*mm_results).gp1, (*mm_results).err
	}
	if mmCreate.funcCreate != nil {
		return mmCreate.funcCreate(ctx)
	}
	mmCreate.t.Fatalf("Unexpected call to GroupRepositoryMock.Create. %v", ctx)
	return
}

// CreateAfterCounter returns a count of finished GroupRepositoryMock.Create invocations
func (mmCreate *GroupRepositoryMock) CreateAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCreate.afterCreateCounter)
}

// CreateBeforeCounter returns a count of GroupRepositoryMock.Create invocations
func (mmCreate *GroupRepositoryMock) CreateBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmCreate.beforeCreateCounter)
}

// Calls returns a list of arguments used in each call to GroupRepositoryMock.Create.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmCreate *mGroupRepositoryMockCreate) Calls() []*GroupRepositoryMockCreateParams {
	mmCreate.mutex.RLock()

	argCopy := make([]*GroupRepositoryMockCreateParams, len(mmCreate.callArgs))
	copy(argCopy, mmCreate.callArgs)

	mmCreate.mutex.RUnlock()

	return argCopy
}

// MinimockCreateDone returns true if the count of the Create invocations corresponds
// the number of defined expectations
func (m *GroupRepositoryMock) MinimockCreateDone() bool {
	for _, e := range m.CreateMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.CreateMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterCreateCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCreate != nil && mm_atomic.LoadUint64(&m.afterCreateCounter) < 1 {
		return false
	}
	return true
}

// MinimockCreateInspect logs each unmet expectation
func (m *GroupRepositoryMock) MinimockCreateInspect() {
	for _, e := range m.CreateMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to GroupRepositoryMock.Create with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.CreateMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterCreateCounter) < 1 {
		if m.CreateMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to GroupRepositoryMock.Create")
		} else {
			m.t.Errorf("Expected call to GroupRepositoryMock.Create with params: %#v", *m.CreateMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcCreate != nil && mm_atomic.LoadUint64(&m.afterCreateCounter) < 1 {
		m.t.Error("Expected call to GroupRepositoryMock.Create")
	}
}

type mGroupRepositoryMockGet struct {
	mock               *GroupRepositoryMock
	defaultExpectation *GroupRepositoryMockGetExpectation
	expectations       []*GroupRepositoryMockGetExpectation

	callArgs []*GroupRepositoryMockGetParams
	mutex    sync.RWMutex
}

// GroupRepositoryMockGetExpectation specifies expectation struct of the GroupRepository.Get
type GroupRepositoryMockGetExpectation struct {
	mock    *GroupRepositoryMock
	params  *GroupRepositoryMockGetParams
	results *GroupRepositoryMockGetResults
	Counter uint64
}

// GroupRepositoryMockGetParams contains parameters of the GroupRepository.Get
type GroupRepositoryMockGetParams struct {
	ctx context.Context
	id  uuid.UUID
}

// GroupRepositoryMockGetResults contains results of the GroupRepository.Get
type GroupRepositoryMockGetResults struct {
	gp1 *models.Group
	err error
}

// Expect sets up expected params for GroupRepository.Get
func (mmGet *mGroupRepositoryMockGet) Expect(ctx context.Context, id uuid.UUID) *mGroupRepositoryMockGet {
	if mmGet.mock.funcGet != nil {
		mmGet.mock.t.Fatalf("GroupRepositoryMock.Get mock is already set by Set")
	}

	if mmGet.defaultExpectation == nil {
		mmGet.defaultExpectation = &GroupRepositoryMockGetExpectation{}
	}

	mmGet.defaultExpectation.params = &GroupRepositoryMockGetParams{ctx, id}
	for _, e := range mmGet.expectations {
		if minimock.Equal(e.params, mmGet.defaultExpectation.params) {
			mmGet.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmGet.defaultExpectation.params)
		}
	}

	return mmGet
}

// Inspect accepts an inspector function that has same arguments as the GroupRepository.Get
func (mmGet *mGroupRepositoryMockGet) Inspect(f func(ctx context.Context, id uuid.UUID)) *mGroupRepositoryMockGet {
	if mmGet.mock.inspectFuncGet != nil {
		mmGet.mock.t.Fatalf("Inspect function is already set for GroupRepositoryMock.Get")
	}

	mmGet.mock.inspectFuncGet = f

	return mmGet
}

// Return sets up results that will be returned by GroupRepository.Get
func (mmGet *mGroupRepositoryMockGet) Return(gp1 *models.Group, err error) *GroupRepositoryMock {
	if mmGet.mock.funcGet != nil {
		mmGet.mock.t.Fatalf("GroupRepositoryMock.Get mock is already set by Set")
	}

	if mmGet.defaultExpectation == nil {
		mmGet.defaultExpectation = &GroupRepositoryMockGetExpectation{mock: mmGet.mock}
	}
	mmGet.defaultExpectation.results = &GroupRepositoryMockGetResults{gp1, err}
	return mmGet.mock
}

// Set uses given function f to mock the GroupRepository.Get method
func (mmGet *mGroupRepositoryMockGet) Set(f func(ctx context.Context, id uuid.UUID) (gp1 *models.Group, err error)) *GroupRepositoryMock {
	if mmGet.defaultExpectation != nil {
		mmGet.mock.t.Fatalf("Default expectation is already set for the GroupRepository.Get method")
	}

	if len(mmGet.expectations) > 0 {
		mmGet.mock.t.Fatalf("Some expectations are already set for the GroupRepository.Get method")
	}

	mmGet.mock.funcGet = f
	return mmGet.mock
}

// When sets expectation for the GroupRepository.Get which will trigger the result defined by the following
// Then helper
func (mmGet *mGroupRepositoryMockGet) When(ctx context.Context, id uuid.UUID) *GroupRepositoryMockGetExpectation {
	if mmGet.mock.funcGet != nil {
		mmGet.mock.t.Fatalf("GroupRepositoryMock.Get mock is already set by Set")
	}

	expectation := &GroupRepositoryMockGetExpectation{
		mock:   mmGet.mock,
		params: &GroupRepositoryMockGetParams{ctx, id},
	}
	mmGet.expectations = append(mmGet.expectations, expectation)
	return expectation
}

// Then sets up GroupRepository.Get return parameters for the expectation previously defined by the When method
func (e *GroupRepositoryMockGetExpectation) Then(gp1 *models.Group, err error) *GroupRepositoryMock {
	e.results = &GroupRepositoryMockGetResults{gp1, err}
	return e.mock
}

// Get implements ports.GroupRepository
func (mmGet *GroupRepositoryMock) Get(ctx context.Context, id uuid.UUID) (gp1 *models.Group, err error) {
	mm_atomic.AddUint64(&mmGet.beforeGetCounter, 1)
	defer mm_atomic.AddUint64(&mmGet.afterGetCounter, 1)

	if mmGet.inspectFuncGet != nil {
		mmGet.inspectFuncGet(ctx, id)
	}

	mm_params := &GroupRepositoryMockGetParams{ctx, id}

	// Record call args
	mmGet.GetMock.mutex.Lock()
	mmGet.GetMock.callArgs = append(mmGet.GetMock.callArgs, mm_params)
	mmGet.GetMock.mutex.Unlock()

	for _, e := range mmGet.GetMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.gp1, e.results.err
		}
	}

	if mmGet.GetMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmGet.GetMock.defaultExpectation.Counter, 1)
		mm_want := mmGet.GetMock.defaultExpectation.params
		mm_got := GroupRepositoryMockGetParams{ctx, id}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmGet.t.Errorf("GroupRepositoryMock.Get got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmGet.GetMock.defaultExpectation.results
		if mm_results == nil {
			mmGet.t.Fatal("No results are set for the GroupRepositoryMock.Get")
		}
		return (*mm_results).gp1, (*mm_results).err
	}
	if mmGet.funcGet != nil {
		return mmGet.funcGet(ctx, id)
	}
	mmGet.t.Fatalf("Unexpected call to GroupRepositoryMock.Get. %v %v", ctx, id)
	return
}

// GetAfterCounter returns a count of finished GroupRepositoryMock.Get invocations
func (mmGet *GroupRepositoryMock) GetAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGet.afterGetCounter)
}

// GetBeforeCounter returns a count of GroupRepositoryMock.Get invocations
func (mmGet *GroupRepositoryMock) GetBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGet.beforeGetCounter)
}

// Calls returns a list of arguments used in each call to GroupRepositoryMock.Get.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmGet *mGroupRepositoryMockGet) Calls() []*GroupRepositoryMockGetParams {
	mmGet.mutex.RLock()

	argCopy := make([]*GroupRepositoryMockGetParams, len(mmGet.callArgs))
	copy(argCopy, mmGet.callArgs)

	mmGet.mutex.RUnlock()

	return argCopy
}

// MinimockGetDone returns true if the count of the Get invocations corresponds
// the number of defined expectations
func (m *GroupRepositoryMock) MinimockGetDone() bool {
	for _, e := range m.GetMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGet != nil && mm_atomic.LoadUint64(&m.afterGetCounter) < 1 {
		return false
	}
	return true
}

// MinimockGetInspect logs each unmet expectation
func (m *GroupRepositoryMock) MinimockGetInspect() {
	for _, e := range m.GetMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to GroupRepositoryMock.Get with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetCounter) < 1 {
		if m.GetMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to GroupRepositoryMock.Get")
		} else {
			m.t.Errorf("Expected call to GroupRepositoryMock.Get with params: %#v", *m.GetMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGet != nil && mm_atomic.LoadUint64(&m.afterGetCounter) < 1 {
		m.t.Error("Expected call to GroupRepositoryMock.Get")
	}
}

type mGroupRepositoryMockGetList struct {
	mock               *GroupRepositoryMock
	defaultExpectation *GroupRepositoryMockGetListExpectation
	expectations       []*GroupRepositoryMockGetListExpectation

	callArgs []*GroupRepositoryMockGetListParams
	mutex    sync.RWMutex
}

// GroupRepositoryMockGetListExpectation specifies expectation struct of the GroupRepository.GetList
type GroupRepositoryMockGetListExpectation struct {
	mock    *GroupRepositoryMock
	params  *GroupRepositoryMockGetListParams
	results *GroupRepositoryMockGetListResults
	Counter uint64
}

// GroupRepositoryMockGetListParams contains parameters of the GroupRepository.GetList
type GroupRepositoryMockGetListParams struct {
	ctx     context.Context
	filters models.Filters
}

// GroupRepositoryMockGetListResults contains results of the GroupRepository.GetList
type GroupRepositoryMockGetListResults struct {
	gpa1 []*models.GroupPreview
	err  error
}

// Expect sets up expected params for GroupRepository.GetList
func (mmGetList *mGroupRepositoryMockGetList) Expect(ctx context.Context, filters models.Filters) *mGroupRepositoryMockGetList {
	if mmGetList.mock.funcGetList != nil {
		mmGetList.mock.t.Fatalf("GroupRepositoryMock.GetList mock is already set by Set")
	}

	if mmGetList.defaultExpectation == nil {
		mmGetList.defaultExpectation = &GroupRepositoryMockGetListExpectation{}
	}

	mmGetList.defaultExpectation.params = &GroupRepositoryMockGetListParams{ctx, filters}
	for _, e := range mmGetList.expectations {
		if minimock.Equal(e.params, mmGetList.defaultExpectation.params) {
			mmGetList.mock.t.Fatalf("Expectation set by When has same params: %#v", *mmGetList.defaultExpectation.params)
		}
	}

	return mmGetList
}

// Inspect accepts an inspector function that has same arguments as the GroupRepository.GetList
func (mmGetList *mGroupRepositoryMockGetList) Inspect(f func(ctx context.Context, filters models.Filters)) *mGroupRepositoryMockGetList {
	if mmGetList.mock.inspectFuncGetList != nil {
		mmGetList.mock.t.Fatalf("Inspect function is already set for GroupRepositoryMock.GetList")
	}

	mmGetList.mock.inspectFuncGetList = f

	return mmGetList
}

// Return sets up results that will be returned by GroupRepository.GetList
func (mmGetList *mGroupRepositoryMockGetList) Return(gpa1 []*models.GroupPreview, err error) *GroupRepositoryMock {
	if mmGetList.mock.funcGetList != nil {
		mmGetList.mock.t.Fatalf("GroupRepositoryMock.GetList mock is already set by Set")
	}

	if mmGetList.defaultExpectation == nil {
		mmGetList.defaultExpectation = &GroupRepositoryMockGetListExpectation{mock: mmGetList.mock}
	}
	mmGetList.defaultExpectation.results = &GroupRepositoryMockGetListResults{gpa1, err}
	return mmGetList.mock
}

// Set uses given function f to mock the GroupRepository.GetList method
func (mmGetList *mGroupRepositoryMockGetList) Set(f func(ctx context.Context, filters models.Filters) (gpa1 []*models.GroupPreview, err error)) *GroupRepositoryMock {
	if mmGetList.defaultExpectation != nil {
		mmGetList.mock.t.Fatalf("Default expectation is already set for the GroupRepository.GetList method")
	}

	if len(mmGetList.expectations) > 0 {
		mmGetList.mock.t.Fatalf("Some expectations are already set for the GroupRepository.GetList method")
	}

	mmGetList.mock.funcGetList = f
	return mmGetList.mock
}

// When sets expectation for the GroupRepository.GetList which will trigger the result defined by the following
// Then helper
func (mmGetList *mGroupRepositoryMockGetList) When(ctx context.Context, filters models.Filters) *GroupRepositoryMockGetListExpectation {
	if mmGetList.mock.funcGetList != nil {
		mmGetList.mock.t.Fatalf("GroupRepositoryMock.GetList mock is already set by Set")
	}

	expectation := &GroupRepositoryMockGetListExpectation{
		mock:   mmGetList.mock,
		params: &GroupRepositoryMockGetListParams{ctx, filters},
	}
	mmGetList.expectations = append(mmGetList.expectations, expectation)
	return expectation
}

// Then sets up GroupRepository.GetList return parameters for the expectation previously defined by the When method
func (e *GroupRepositoryMockGetListExpectation) Then(gpa1 []*models.GroupPreview, err error) *GroupRepositoryMock {
	e.results = &GroupRepositoryMockGetListResults{gpa1, err}
	return e.mock
}

// GetList implements ports.GroupRepository
func (mmGetList *GroupRepositoryMock) GetList(ctx context.Context, filters models.Filters) (gpa1 []*models.GroupPreview, err error) {
	mm_atomic.AddUint64(&mmGetList.beforeGetListCounter, 1)
	defer mm_atomic.AddUint64(&mmGetList.afterGetListCounter, 1)

	if mmGetList.inspectFuncGetList != nil {
		mmGetList.inspectFuncGetList(ctx, filters)
	}

	mm_params := &GroupRepositoryMockGetListParams{ctx, filters}

	// Record call args
	mmGetList.GetListMock.mutex.Lock()
	mmGetList.GetListMock.callArgs = append(mmGetList.GetListMock.callArgs, mm_params)
	mmGetList.GetListMock.mutex.Unlock()

	for _, e := range mmGetList.GetListMock.expectations {
		if minimock.Equal(e.params, mm_params) {
			mm_atomic.AddUint64(&e.Counter, 1)
			return e.results.gpa1, e.results.err
		}
	}

	if mmGetList.GetListMock.defaultExpectation != nil {
		mm_atomic.AddUint64(&mmGetList.GetListMock.defaultExpectation.Counter, 1)
		mm_want := mmGetList.GetListMock.defaultExpectation.params
		mm_got := GroupRepositoryMockGetListParams{ctx, filters}
		if mm_want != nil && !minimock.Equal(*mm_want, mm_got) {
			mmGetList.t.Errorf("GroupRepositoryMock.GetList got unexpected parameters, want: %#v, got: %#v%s\n", *mm_want, mm_got, minimock.Diff(*mm_want, mm_got))
		}

		mm_results := mmGetList.GetListMock.defaultExpectation.results
		if mm_results == nil {
			mmGetList.t.Fatal("No results are set for the GroupRepositoryMock.GetList")
		}
		return (*mm_results).gpa1, (*mm_results).err
	}
	if mmGetList.funcGetList != nil {
		return mmGetList.funcGetList(ctx, filters)
	}
	mmGetList.t.Fatalf("Unexpected call to GroupRepositoryMock.GetList. %v %v", ctx, filters)
	return
}

// GetListAfterCounter returns a count of finished GroupRepositoryMock.GetList invocations
func (mmGetList *GroupRepositoryMock) GetListAfterCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetList.afterGetListCounter)
}

// GetListBeforeCounter returns a count of GroupRepositoryMock.GetList invocations
func (mmGetList *GroupRepositoryMock) GetListBeforeCounter() uint64 {
	return mm_atomic.LoadUint64(&mmGetList.beforeGetListCounter)
}

// Calls returns a list of arguments used in each call to GroupRepositoryMock.GetList.
// The list is in the same order as the calls were made (i.e. recent calls have a higher index)
func (mmGetList *mGroupRepositoryMockGetList) Calls() []*GroupRepositoryMockGetListParams {
	mmGetList.mutex.RLock()

	argCopy := make([]*GroupRepositoryMockGetListParams, len(mmGetList.callArgs))
	copy(argCopy, mmGetList.callArgs)

	mmGetList.mutex.RUnlock()

	return argCopy
}

// MinimockGetListDone returns true if the count of the GetList invocations corresponds
// the number of defined expectations
func (m *GroupRepositoryMock) MinimockGetListDone() bool {
	for _, e := range m.GetListMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			return false
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetListMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetListCounter) < 1 {
		return false
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetList != nil && mm_atomic.LoadUint64(&m.afterGetListCounter) < 1 {
		return false
	}
	return true
}

// MinimockGetListInspect logs each unmet expectation
func (m *GroupRepositoryMock) MinimockGetListInspect() {
	for _, e := range m.GetListMock.expectations {
		if mm_atomic.LoadUint64(&e.Counter) < 1 {
			m.t.Errorf("Expected call to GroupRepositoryMock.GetList with params: %#v", *e.params)
		}
	}

	// if default expectation was set then invocations count should be greater than zero
	if m.GetListMock.defaultExpectation != nil && mm_atomic.LoadUint64(&m.afterGetListCounter) < 1 {
		if m.GetListMock.defaultExpectation.params == nil {
			m.t.Error("Expected call to GroupRepositoryMock.GetList")
		} else {
			m.t.Errorf("Expected call to GroupRepositoryMock.GetList with params: %#v", *m.GetListMock.defaultExpectation.params)
		}
	}
	// if func was set then invocations count should be greater than zero
	if m.funcGetList != nil && mm_atomic.LoadUint64(&m.afterGetListCounter) < 1 {
		m.t.Error("Expected call to GroupRepositoryMock.GetList")
	}
}

// MinimockFinish checks that all mocked methods have been called the expected number of times
func (m *GroupRepositoryMock) MinimockFinish() {
	if !m.minimockDone() {
		m.MinimockCreateInspect()

		m.MinimockGetInspect()

		m.MinimockGetListInspect()
		m.t.FailNow()
	}
}

// MinimockWait waits for all mocked methods to be called the expected number of times
func (m *GroupRepositoryMock) MinimockWait(timeout mm_time.Duration) {
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

func (m *GroupRepositoryMock) minimockDone() bool {
	done := true
	return done &&
		m.MinimockCreateDone() &&
		m.MinimockGetDone() &&
		m.MinimockGetListDone()
}