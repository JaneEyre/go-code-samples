package fetchuser

import (
	"testing"
)

type MockInterfaceFetcher struct {
	u User
}

func (m *MockInterfaceFetcher) FetchData(_ int) (User, error) {
	return m.u, nil
}

func TestProcessUser_InterfaceMock(t *testing.T) {
	user := User{ID: 1, Name: "Alice"}
	result, err := ProcessUser(&MockInterfaceFetcher{user}, 1)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if result != user {
		t.Errorf("Expected user: %v, got: %v", user, result)
	}
}
