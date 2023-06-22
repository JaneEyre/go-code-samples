package fetchuser

import (
	"testing"

	"github.com/stretchr/testify/mock"
)

type MockTestifyFetcher struct {
	mock.Mock
}

func (m *MockTestifyFetcher) FetchData(id int) (User, error) {
	args := m.Called(id)
	return args.Get(0).(User), args.Error(1)
}

func TestProcessUser_TestifyMock(t *testing.T) {
	user := User{ID: 1, Name: "Alice"}
	mockFetcher := new(MockTestifyFetcher)
	mockFetcher.On("FetchData", 1).Return(user, nil)

	result, err := ProcessUser(mockFetcher, 1)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if result != user {
		t.Errorf("expected user: %v, got: %v", user, result)
	}
}
