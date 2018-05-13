package main

import (
	"Go4Chan/JsonStructs"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
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
	defer response.Body.Close()

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil || contents == nil {
		fmt.Println(err)
	}

	var jsonRoot = JsonStructs.GetAllBoards()
	var allBoards = jsonRoot.Boards

	for index, board := range allBoards {
		fmt.Printf("%d - %s \n", index, board.Title)
	}

	var pagesRoot = JsonStructs.GetAllPages("w")
	for index, page := range pagesRoot {
		fmt.Printf("Index: %d Page: %d \n", index, page.Page)
		for index, thread := range page.Threads {
			fmt.Printf("Index: %d Thread: %d\n", index, thread.No)
		}
	}

	fmt.Println("Jaga")

}
