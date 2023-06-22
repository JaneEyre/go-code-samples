package fetchuser

import (
	"testing"
)

type MockInterfaceFetcher2 struct {
	FetchDataFunc func(id int) (User, error)
}

func (m *MockInterfaceFetcher2) FetchData(id int) (User, error) {
	return m.FetchDataFunc(id)
}

func TestProcessUser_InterfaceMock2(t *testing.T) {
	user := User{ID: 1, Name: "Alice"}
	mockFetcher := &MockInterfaceFetcher2{
		FetchDataFunc: func(id int) (User, error) {
			return user, nil
		}}

	result, err := ProcessUser(mockFetcher, 1)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if result != user {
		t.Errorf("Expected user: %v, got: %v", user, result)
	}
}
