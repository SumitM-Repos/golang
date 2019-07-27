package utility

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"sort"
)

type User struct {
	ID        int32  `json:"id"`
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	Email     string `json:"email"`
	Gender    string `json:"gender"`
	IPAddress string `json:"ip_address"`
	Timestamp string `json:"timestamp"`
}

func ReadJsonFile() []User {

	file, err := ioutil.ReadFile("./test.json")
	if err != nil {
		errors.New("Unable to read File")
	}

	data := []User{}

	err = json.Unmarshal([]byte(file), &data)

	if err != nil {
		errors.New("Unable to read File")
	}
	return data

}

func SortJsonValuesByTimeStamp(file []User) []User {

	sort.SliceStable(file, func(i, j int) bool {
		return file[i].Timestamp < file[j].Timestamp
	})
	return file
}
