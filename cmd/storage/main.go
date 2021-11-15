package main

import (
	"fmt"
	"log"
	"storage/internal/storage"
	//"github.com/Postoev-Alexander/storage/internal/storage"
)

func main() {
	st := storage.NewStorage()

	file, err := st.Upload("test.txt", []byte("hello"))
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("it uploaded", file)
}
