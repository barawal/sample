package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func main() {
	http.HandleFunc("/", handler)
	f, err := os.Create("pid.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	l, err := f.WriteString(strconv.Itoa(os.Getpid()))
	if err != nil {
		fmt.Println(err)
		f.Close()
		return
	}
	fmt.Println(l, "bytes written successfully")
	err = f.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	log.Fatal(http.ListenAndServe(":8080", nil))
}
