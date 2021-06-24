// Code generated by mockery v2.8.0. DO NOT EDIT.

package mock

import (
	context "context"

	model "github.com/MuZaZaVr/account-service/internal/model"
	mock "github.com/stretchr/testify/mock"
)

// Account is an autogenerated mock type for the Account type
type Account struct {
	mock.Mock
}

// Create provides a mock function with given fields: ctx, account
func (_m *Account) Create(ctx context.Context, account model.AccountDTO) (string, error) {
	ret := _m.Called(ctx, account)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, model.AccountDTO) string); ok {
		r0 = rf(ctx, account)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, model.AccountDTO) error); ok {
		r1 = rf(ctx, account)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Delete provides a mock function with given fields: ctx, id
func (_m *Account) Delete(ctx context.Context, id string) (string, error) {
	ret := _m.Called(ctx, id)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, string) string); ok {
		r0 = rf(ctx, id)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindAccountsByCredentialsAddress provides a mock function with given fields: ctx, credentialsAddress
func (_m *Account) FindAccountsByCredentialsAddress(ctx context.Context, credentialsAddress string) ([]model.AccountDTO, error) {
	ret := _m.Called(ctx, credentialsAddress)

	var r0 []model.AccountDTO
	if rf, ok := ret.Get(0).(func(context.Context, string) []model.AccountDTO); ok {
		r0 = rf(ctx, credentialsAddress)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.AccountDTO)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, credentialsAddress)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindAccountsByCredentialsAge provides a mock function with given fields: ctx, credentialsAge
func (_m *Account) FindAccountsByCredentialsAge(ctx context.Context, credentialsAge int) ([]model.AccountDTO, error) {
	ret := _m.Called(ctx, credentialsAge)

	var r0 []model.AccountDTO
	if rf, ok := ret.Get(0).(func(context.Context, int) []model.AccountDTO); ok {
		r0 = rf(ctx, credentialsAge)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.AccountDTO)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, credentialsAge)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindAccountsByCredentialsCity provides a mock function with given fields: ctx, credentialsCity
func (_m *Account) FindAccountsByCredentialsCity(ctx context.Context, credentialsCity string) ([]model.AccountDTO, error) {
	ret := _m.Called(ctx, credentialsCity)

	var r0 []model.AccountDTO
	if rf, ok := ret.Get(0).(func(context.Context, string) []model.AccountDTO); ok {
		r0 = rf(ctx, credentialsCity)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.AccountDTO)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, credentialsCity)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindAccountsByCredentialsEmail provides a mock function with given fields: ctx, credentialsEmail
func (_m *Account) FindAccountsByCredentialsEmail(ctx context.Context, credentialsEmail string) ([]model.AccountDTO, error) {
	ret := _m.Called(ctx, credentialsEmail)

	var r0 []model.AccountDTO
	if rf, ok := ret.Get(0).(func(context.Context, string) []model.AccountDTO); ok {
		r0 = rf(ctx, credentialsEmail)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.AccountDTO)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, credentialsEmail)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindAccountsByCredentialsLogin provides a mock function with given fields: ctx, credentialsLogin
func (_m *Account) FindAccountsByCredentialsLogin(ctx context.Context, credentialsLogin string) ([]model.AccountDTO, error) {
	ret := _m.Called(ctx, credentialsLogin)

	var r0 []model.AccountDTO
	if rf, ok := ret.Get(0).(func(context.Context, string) []model.AccountDTO); ok {
		r0 = rf(ctx, credentialsLogin)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.AccountDTO)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, credentialsLogin)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindAccountsByCredentialsMiddlename provides a mock function with given fields: ctx, credentialsMiddlename
func (_m *Account) FindAccountsByCredentialsMiddlename(ctx context.Context, credentialsMiddlename string) ([]model.AccountDTO, error) {
	ret := _m.Called(ctx, credentialsMiddlename)

	var r0 []model.AccountDTO
	if rf, ok := ret.Get(0).(func(context.Context, string) []model.AccountDTO); ok {
		r0 = rf(ctx, credentialsMiddlename)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.AccountDTO)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, credentialsMiddlename)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindAccountsByCredentialsName provides a mock function with given fields: ctx, credentialsName
func (_m *Account) FindAccountsByCredentialsName(ctx context.Context, credentialsName string) ([]model.AccountDTO, error) {
	ret := _m.Called(ctx, credentialsName)

	var r0 []model.AccountDTO
	if rf, ok := ret.Get(0).(func(context.Context, string) []model.AccountDTO); ok {
		r0 = rf(ctx, credentialsName)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.AccountDTO)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, credentialsName)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindAccountsByCredentialsPhone provides a mock function with given fields: ctx, credentialsPhone
func (_m *Account) FindAccountsByCredentialsPhone(ctx context.Context, credentialsPhone string) ([]model.AccountDTO, error) {
	ret := _m.Called(ctx, credentialsPhone)

	var r0 []model.AccountDTO
	if rf, ok := ret.Get(0).(func(context.Context, string) []model.AccountDTO); ok {
		r0 = rf(ctx, credentialsPhone)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.AccountDTO)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, credentialsPhone)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindAccountsByCredentialsSurname provides a mock function with given fields: ctx, credentialsSurname
func (_m *Account) FindAccountsByCredentialsSurname(ctx context.Context, credentialsSurname string) ([]model.AccountDTO, error) {
	ret := _m.Called(ctx, credentialsSurname)

	var r0 []model.AccountDTO
	if rf, ok := ret.Get(0).(func(context.Context, string) []model.AccountDTO); ok {
		r0 = rf(ctx, credentialsSurname)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.AccountDTO)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, credentialsSurname)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindAccountsByUserID provides a mock function with given fields: ctx, id
func (_m *Account) FindAccountsByUserID(ctx context.Context, id int) ([]model.AccountDTO, error) {
	ret := _m.Called(ctx, id)

	var r0 []model.AccountDTO
	if rf, ok := ret.Get(0).(func(context.Context, int) []model.AccountDTO); ok {
		r0 = rf(ctx, id)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]model.AccountDTO)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, int) error); ok {
		r1 = rf(ctx, id)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// FindByName provides a mock function with given fields: ctx, name
func (_m *Account) FindByName(ctx context.Context, name string) (*model.AccountDTO, error) {
	ret := _m.Called(ctx, name)

	var r0 *model.AccountDTO
	if rf, ok := ret.Get(0).(func(context.Context, string) *model.AccountDTO); ok {
		r0 = rf(ctx, name)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*model.AccountDTO)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// IsExist provides a mock function with given fields: ctx, name
func (_m *Account) IsExist(ctx context.Context, name string) (bool, error) {
	ret := _m.Called(ctx, name)

	var r0 bool
	if rf, ok := ret.Get(0).(func(context.Context, string) bool); ok {
		r0 = rf(ctx, name)
	} else {
		r0 = ret.Get(0).(bool)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string) error); ok {
		r1 = rf(ctx, name)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Update provides a mock function with given fields: ctx, id, newAccount
func (_m *Account) Update(ctx context.Context, id string, newAccount model.AccountDTO) (string, error) {
	ret := _m.Called(ctx, id, newAccount)

	var r0 string
	if rf, ok := ret.Get(0).(func(context.Context, string, model.AccountDTO) string); ok {
		r0 = rf(ctx, id, newAccount)
	} else {
		r0 = ret.Get(0).(string)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, string, model.AccountDTO) error); ok {
		r1 = rf(ctx, id, newAccount)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}
