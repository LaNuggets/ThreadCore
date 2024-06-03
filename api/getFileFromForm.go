package api

import (
	"errors"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"os"
)

func GetFileFromForm(file multipart.File, handler *multipart.FileHeader, err error, path string) {
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
		return
	}
	defer file.Close()
	// fmt.Printf("Uploaded File: %+v\n", handler.Filename)
	// fmt.Printf("File Size: %+v\n", handler.Size)
	// fmt.Printf("MIME Header: %+v\n", handler.Header)

	if _, err := os.Stat("." + path); errors.Is(err, os.ErrNotExist) {
		// file does not exist
	} else {
		e := os.Remove("." + path)
		if e != nil {
			log.Fatal(e)
		}
	}

	f, err := os.OpenFile("."+path, os.O_WRONLY|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	io.Copy(f, file)
}
