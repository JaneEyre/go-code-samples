package fetchuser

import (
	"testing"

	"github.com/golang/mock/gomock"
)

func TestProcessUser_Mockgen(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	user := User{ID: 1, Name: "Alice"}
	mockFetcher := NewMockAPIFetcher(ctrl)
	mockFetcher.EXPECT().FetchData(1).Return(user, nil)

	result, err := ProcessUser(mockFetcher, 1)
	if err != nil {
		t.Errorf("unexpected error: %v", err)
	}

	if result != user {
		t.Errorf("expected user: %v, got: %v", user, result)
	}
}
