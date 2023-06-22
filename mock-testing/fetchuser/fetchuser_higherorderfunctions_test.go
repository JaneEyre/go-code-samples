package fetchuser

import (
	"testing"
)

type FetchDataFunc func(id int) (User, error)

func (f FetchDataFunc) FetchData(id int) (User, error) {
	return f(id)
}

func TestProcessUser_HigherOrderFunctions(t *testing.T) {
	user := User{ID: 1, Name: "Alice"}
	var mockFetcher FetchDataFunc = func(id int) (User, error) {
		return user, nil
	}

	result, err := ProcessUser(mockFetcher, 1)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if result != user {
		t.Errorf("Expected user: %v, got: %v", user, result)
	}
}
