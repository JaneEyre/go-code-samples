package fetchuser

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type APIFetcher interface {
	FetchData(id int) (User, error)
}

type RealAPIFetcher struct {
	ApiURL string
}

func (ra *RealAPIFetcher) FetchData(id int) (User, error) {
	resp, err := http.Get(fmt.Sprintf("%s/users/%d", ra.ApiURL, id))
	if err != nil {
		return User{}, err
	}
	defer resp.Body.Close()

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return User{}, err
	}

	var user User
	err = json.Unmarshal(bodyBytes, &user)
	return user, err
}

func ProcessUser(fetcher APIFetcher, id int) (User, error) {
	user, err := fetcher.FetchData(id)
	if err != nil {
		return User{}, err
	}
	// Process the user data.
	return user, nil
}
