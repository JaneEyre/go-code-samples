package fetchuser

import (
	"testing"
)

func TestProcessUser_HigherOrderFunctions(t *testing.T) {
	user := User{ID: 1, Name: "Alice"}

	var mockFetcher FetchDataFunc = func(url string, id int) (User, error) {
		return user, nil
	}

	result, err := ProcessUserHOF(mockFetcher, "noURL", 1)
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if result != user {
		t.Errorf("Expected user: %v, got: %v", user, result)
	}
}
