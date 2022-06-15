// Code generated by mockery v2.12.3. DO NOT EDIT.

package mocks

import (
	context "context"

	domain "github.com/bangumi/server/internal/domain"
	mock "github.com/stretchr/testify/mock"

	model "github.com/bangumi/server/internal/model"
)

// PersonRepo is an autogenerated mock type for the PersonRepo type
type PersonRepo struct {
	mock.Mock
}

type PersonRepo_Expecter struct {
	mock *mock.Mock
}

func (_m *PersonRepo) EXPECT() *PersonRepo_Expecter {
	return &PersonRepo_Expecter{mock: &_m.Mock}
}

// Get provides a mock function with given fields: ctx, id
func (_m *PersonRepo) Get(ctx context.Context, id model.PersonID) (model.Person, error) {
	ret := _m.Called(ctx, id)

	var r0 model.Person
	if rf, ok := ret.Get(0).(func(context.Context, model.PersonID) model.Person); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(model.Person)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, model.PersonID) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PersonRepo_Get_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Get'
type PersonRepo_Get_Call struct {
	*mock.Call
}

// Get is a helper method to define mock.On call
//  - ctx context.Context
//  - id model.PersonID
func (_e *PersonRepo_Expecter) Get(ctx interface{}, id interface{}) *PersonRepo_Get_Call {
	return &PersonRepo_Get_Call{Call: _e.mock.On("Get", ctx, id)}
}

func (_c *PersonRepo_Get_Call) Run(run func(ctx context.Context, id model.PersonID)) *PersonRepo_Get_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(model.PersonID))
	})
	return _c
}

func (_c *PersonRepo_Get_Call) Return(_a0 model.Person, _a1 error) *PersonRepo_Get_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetByIDs provides a mock function with given fields: ctx, ids
func (_m *PersonRepo) GetByIDs(ctx context.Context, ids ...model.PersonID) (map[model.PersonID]model.Person, error) {
	_va := make([]interface{}, len(ids))
	for _i := range ids {
		_va[_i] = ids[_i]
	}
	var _ca []interface{}
	_ca = append(_ca, ctx)
	_ca = append(_ca, _va...)
	ret := _m.Called(_ca...)

	var r0 map[model.PersonID]model.Person
	if rf, ok := ret.Get(0).(func(context.Context, ...model.PersonID) map[model.PersonID]model.Person); ok {
		r0 = rf(ctx, ids...)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[model.PersonID]model.Person)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, ...model.PersonID) error); ok {
		r1 = rf(ctx, ids...)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PersonRepo_GetByIDs_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetByIDs'
type PersonRepo_GetByIDs_Call struct {
	*mock.Call
}

// GetByIDs is a helper method to define mock.On call
//  - ctx context.Context
//  - ids ...model.PersonID
func (_e *PersonRepo_Expecter) GetByIDs(ctx interface{}, ids ...interface{}) *PersonRepo_GetByIDs_Call {
	return &PersonRepo_GetByIDs_Call{Call: _e.mock.On("GetByIDs",
		append([]interface{}{ctx}, ids...)...)}
}

func (_c *PersonRepo_GetByIDs_Call) Run(run func(ctx context.Context, ids ...model.PersonID)) *PersonRepo_GetByIDs_Call {
	_c.Call.Run(func(args mock.Arguments) {
		variadicArgs := make([]model.PersonID, len(args)-1)
		for i, a := range args[1:] {
			if a != nil {
				variadicArgs[i] = a.(model.PersonID)
			}
		}
		run(args[0].(context.Context), variadicArgs...)
	})
	return _c
}

func (_c *PersonRepo_GetByIDs_Call) Return(_a0 map[model.PersonID]model.Person, _a1 error) *PersonRepo_GetByIDs_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetCharacterRelated provides a mock function with given fields: ctx, subjectID
func (_m *PersonRepo) GetCharacterRelated(ctx context.Context, subjectID model.CharacterID) ([]domain.PersonCharacterRelation, error) {
	ret := _m.Called(ctx, subjectID)

	var r0 []domain.PersonCharacterRelation
	if rf, ok := ret.Get(0).(func(context.Context, model.CharacterID) []domain.PersonCharacterRelation); ok {
		r0 = rf(ctx, subjectID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.PersonCharacterRelation)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, model.CharacterID) error); ok {
		r1 = rf(ctx, subjectID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PersonRepo_GetCharacterRelated_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetCharacterRelated'
type PersonRepo_GetCharacterRelated_Call struct {
	*mock.Call
}

// GetCharacterRelated is a helper method to define mock.On call
//  - ctx context.Context
//  - subjectID model.CharacterID
func (_e *PersonRepo_Expecter) GetCharacterRelated(ctx interface{}, subjectID interface{}) *PersonRepo_GetCharacterRelated_Call {
	return &PersonRepo_GetCharacterRelated_Call{Call: _e.mock.On("GetCharacterRelated", ctx, subjectID)}
}

func (_c *PersonRepo_GetCharacterRelated_Call) Run(run func(ctx context.Context, subjectID model.CharacterID)) *PersonRepo_GetCharacterRelated_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(model.CharacterID))
	})
	return _c
}

func (_c *PersonRepo_GetCharacterRelated_Call) Return(_a0 []domain.PersonCharacterRelation, _a1 error) *PersonRepo_GetCharacterRelated_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

// GetSubjectRelated provides a mock function with given fields: ctx, subjectID
func (_m *PersonRepo) GetSubjectRelated(ctx context.Context, subjectID model.SubjectID) ([]domain.SubjectPersonRelation, error) {
	ret := _m.Called(ctx, subjectID)

	var r0 []domain.SubjectPersonRelation
	if rf, ok := ret.Get(0).(func(context.Context, model.SubjectID) []domain.SubjectPersonRelation); ok {
		r0 = rf(ctx, subjectID)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]domain.SubjectPersonRelation)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, model.SubjectID) error); ok {
		r1 = rf(ctx, subjectID)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// PersonRepo_GetSubjectRelated_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSubjectRelated'
type PersonRepo_GetSubjectRelated_Call struct {
	*mock.Call
}

// GetSubjectRelated is a helper method to define mock.On call
//  - ctx context.Context
//  - subjectID model.SubjectID
func (_e *PersonRepo_Expecter) GetSubjectRelated(ctx interface{}, subjectID interface{}) *PersonRepo_GetSubjectRelated_Call {
	return &PersonRepo_GetSubjectRelated_Call{Call: _e.mock.On("GetSubjectRelated", ctx, subjectID)}
}

func (_c *PersonRepo_GetSubjectRelated_Call) Run(run func(ctx context.Context, subjectID model.SubjectID)) *PersonRepo_GetSubjectRelated_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(model.SubjectID))
	})
	return _c
}

func (_c *PersonRepo_GetSubjectRelated_Call) Return(_a0 []domain.SubjectPersonRelation, _a1 error) *PersonRepo_GetSubjectRelated_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

type NewPersonRepoT interface {
	mock.TestingT
	Cleanup(func())
}

// NewPersonRepo creates a new instance of PersonRepo. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewPersonRepo(t NewPersonRepoT) *PersonRepo {
	mock := &PersonRepo{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
