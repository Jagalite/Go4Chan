package JsonStructs

import (
	"net/http"
        "fmt"
	"encoding/json"
)

type JsonPages []struct {
	Page    int `json:"page"`
	Threads []struct {
		No           int `json:"no"`
		LastModified int `json:"last_modified"`
	} `json:"threads"`
}

func GetAllPages(board string) JsonPages {
	response, err := http.Get("https://a.4cdn.org/" + board + "/threads.json")
	if err != nil {
		fmt.Println(err)
        }
	defer response.Body.Close();

        var pages JsonPages
	json.NewDecoder(response.Body).Decode(&pages)

	return pages
}
