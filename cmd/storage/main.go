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

	restoredFile, err := st.GetByID(file.ID)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("map", st.Files)
	fmt.Println("it uploaded", file)
	fmt.Println("it restored", restoredFile)
}
