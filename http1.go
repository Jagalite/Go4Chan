package main

import (
	"io"
	"io/ioutil"
	"net/http"
	"fmt"
)

func hello(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Hello, World! \n")
}

func main() {
	//http.HandleFunc("/", hello)
	//fs := http.FileServer(http.Dir("/"))
	//http.Handle("/", fs)
	//http.ListenAndServe(":1337", nil)

	response, err := http.Get("https://a.4cdn.org/boards.json")
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close();

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(contents))

}
