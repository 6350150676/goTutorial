package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

type Todo struct {
	UserID    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func getreq() {
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/4")
	if err != nil {
		fmt.Println("Error getting:", err)
		return
	}
	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println("Error in getting Response:", res.Status)
		return
	}

	var todo Todo
	err = json.NewDecoder(res.Body).Decode(&todo)
	if err != nil {
		fmt.Println("Error decoding:", err)
		return
	}

	fmt.Println("Todo:", todo)
}
func Postreq() {
	todo := Todo{
		UserID:    23,
		Title:     "Prince Kumar",
		Completed: true,
	}
	jsonData, _ := json.Marshal(todo)
	jsonString := string(jsonData)

	jsonReader := strings.NewReader(jsonString)

	myURL := "https://jsonplaceholder.typicode.com/todos"
	res, _ := http.Post(myURL, "application/json", jsonReader)

	defer res.Body.Close()
	data, _ := ioutil.ReadAll(res.Body)
	fmt.Println("Response : ", string(data))
}
func performUpdateRequest() {
	todo := Todo{
		UserID:    232,
		Title:     "Prince Kuwwdadfmar",
		Completed: true,
	}
	jsonData, _ := json.Marshal(todo)

	jsonString := string(jsonData)

	jsonReader := strings.NewReader(jsonString)

	const myurl = "https://jsonplaceholder.typicode.com/todos/1"

	req, _ := http.NewRequest(http.MethodPut, myurl, jsonReader)
	req.Header.Set("Content-type", "application/json") //always put this for better telling that we are sending json data
	client := http.Client{}
	res, _ := client.Do(req)
	defer res.Body.Close()
	data, _ := ioutil.ReadAll(res.Body)
	fmt.Println("Response : ", string(data))
}

func main() {
	fmt.Println("Learning CRUD...")
	Postreq()
	performUpdateRequest()
}
