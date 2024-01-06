// Code generated by MockGen. DO NOT EDIT.
// Source: C:\Users\erons\TESTES-MOCKS\teste8\database\database.go
//
// Generated by this command:
//
//	mockgen -source=C:\Users\erons\TESTES-MOCKS\teste8\database\database.go -destination=C:\Users\erons\TESTES-MOCKS\database\gomock\mocks_database.go -package=mocks
//

// Package mocks is a generated GoMock package.
package mocks

import (
	sql "database/sql"
	reflect "reflect"

	models "github.com/eron97/testes-mocks.git/teste8/models"
	gomock "go.uber.org/mock/gomock"
)

// MockDatabase is a mock of Database interface.
type MockDatabase struct {
	ctrl     *gomock.Controller
	recorder *MockDatabaseMockRecorder
}

// MockDatabaseMockRecorder is the mock recorder for MockDatabase.
type MockDatabaseMockRecorder struct {
	mock *MockDatabase
}

// NewMockDatabase creates a new mock instance.
func NewMockDatabase(ctrl *gomock.Controller) *MockDatabase {
	mock := &MockDatabase{ctrl: ctrl}
	mock.recorder = &MockDatabaseMockRecorder{mock}
	return mock
}

// EXPECT returns an object that allows the caller to indicate expected use.
func (m *MockDatabase) EXPECT() *MockDatabaseMockRecorder {
	return m.recorder
}

// GetAllUsers mocks base method.
func (m *MockDatabase) GetAllUsers() ([]models.User, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "GetAllUsers")
	ret0, _ := ret[0].([]models.User)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// GetAllUsers indicates an expected call of GetAllUsers.
func (mr *MockDatabaseMockRecorder) GetAllUsers() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "GetAllUsers", reflect.TypeOf((*MockDatabase)(nil).GetAllUsers))
}

// InitDB mocks base method.
func (m *MockDatabase) InitDB() (*sql.DB, error) {
	m.ctrl.T.Helper()
	ret := m.ctrl.Call(m, "InitDB")
	ret0, _ := ret[0].(*sql.DB)
	ret1, _ := ret[1].(error)
	return ret0, ret1
}

// InitDB indicates an expected call of InitDB.
func (mr *MockDatabaseMockRecorder) InitDB() *gomock.Call {
	mr.mock.ctrl.T.Helper()
	return mr.mock.ctrl.RecordCallWithMethodType(mr.mock, "InitDB", reflect.TypeOf((*MockDatabase)(nil).InitDB))
}
