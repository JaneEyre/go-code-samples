package fetchuser

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type fetchFunc func(url string, id int) (User, error)

func realFetchData(url string, id int) (User, error) {
	resp, err := http.Get(fmt.Sprintf("%s/users/%d", url, id))
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

func ProcessUserHOF(fetchData fetchFunc, url string, id int) (User, error) {
	user, err := fetchData(url, id)
	if err != nil {
		return User{}, err
	}
	// Process the user data.
	return user, nil
}
