package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	handler "github.com/mjehanno/todo-back/handlers"
)

const port = 8000

func main() {

	http.HandleFunc("/task", handler.TaskHandler)
	sPort := fmt.Sprintf(":%d", port)
	if err := http.ListenAndServe(sPort, nil); err != nil {
		fmt.Println(err)
		log.Fatal(err)
	} else {
		fmt.Println("*_*")
		log.SetOutput(os.Stdout)
		log.Println("Server is running on port ", port)
	}
}
