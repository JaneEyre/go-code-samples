package fetchuser

import (
	"testing"
)

type MockInterfaceFetcher1 struct{}

var user = User{ID: 1, Name: "Alice"}

func (m *MockInterfaceFetcher1) FetchData(_ int) (User, error) {
	return user, nil
}

func TestProcessUser_InterfaceMock1(t *testing.T) {

	result, err := ProcessUser(&MockInterfaceFetcher1{}, 1)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if result != user {
		t.Errorf("Expected user: %v, got: %v", user, result)
	}
}
