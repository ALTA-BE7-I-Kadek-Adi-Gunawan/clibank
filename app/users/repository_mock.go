package users

import (
	"github.com/stretchr/testify/mock"
	"gorm.io/gorm"
)

type MockUserRepository struct {
	mock.Mock
}

func (m *MockUserRepository) FindUsers() []User {
	args := m.Called()
	return args.Get(0).([]User)
}

func (m *MockUserRepository) Create(data CreateUserDto) (User, error) {
	args := m.Called(data)
	return args.Get(0).(User), args.Error(1)
}

func (m *MockUserRepository) Init(db *gorm.DB) {
	m.Called(db)
}
func (m *MockUserRepository) FindByID(id int) (*User, error) {
	args := m.Called(id)
	return args.Get(0).(*User), args.Error(1)
}
func (m *MockUserRepository) FindByPhone(phone string) (*User, error) {
	args := m.Called(phone)
	return args.Get(0).(*User), args.Error(1)
}

func (m *MockUserRepository) CheckPin(phone string, pin string) (bool, error) {
	args := m.Called(phone, pin)
	return args.Bool(0), args.Error(1)
}

func (m *MockUserRepository) Update(phone string, data UpdateUserDto) (User, error) {
	args := m.Called(phone, data)
	return args.Get(0).(User), args.Error(1)
}

func (m *MockUserRepository) Delete(phone string) error {
	args := m.Called(phone)
	return args.Error(0)
}
