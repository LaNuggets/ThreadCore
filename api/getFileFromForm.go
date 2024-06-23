package api

import (
	"errors"
	"fmt"
	"io"
	"log"
	"mime/multipart"
	"os"
)

/*
! GetFileFromForm is called on when Creating or modifing a post, user or community
! takes the file, filehandler and path
! checks if the file dosent already exists to prevent errors
! saves the file in the server with the corrseponding path wanted using the "os" standard go package
! check for errors
*/
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

/*
! DeleteFile is called on when Creating or modifing a post, user or community
! takes the path of the file we want to remove from the server to avoid taking to much space
! removes the file using the "os" standard go package
! Checks for potential errors
*/
func DeleteFile(path string) {
	if path == "/static/images/bannerTemplate.png" || path == "/static/images/profileTemplate.png" || path == "/static/images/medidaTemplate.png" {
	} else {
		if _, err := os.Stat("." + path); errors.Is(err, os.ErrNotExist) {
		} else {
			e := os.Remove("." + path)
			if e != nil {
				log.Fatal(e)
			}
		}
	}

}
