package api

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

func CreateCommunity(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseMultipartForm(10 << 20)

		file, handler, err := r.FormFile("myFile")
		if err == http.ErrMissingFile {
			fmt.Println("no file uploaded")
		} else {
			if err != nil {
				fmt.Println("Error Retrieving the File")
				fmt.Println(err)
				return
			}
			defer file.Close()
			fmt.Printf("Uploaded File: %+v\n", handler.Filename)
			fmt.Printf("File Size: %+v\n", handler.Size)
			fmt.Printf("MIME Header: %+v\n", handler.Header)

			f, err := os.OpenFile("./imagesTest/"+handler.Filename, os.O_WRONLY|os.O_CREATE, 0666)
			if err != nil {
				fmt.Println(err)
				return
			}
			defer f.Close()
			io.Copy(f, file)
		}

		name := r.FormValue("name")

		http.Redirect(w, r, "/community/"+name, http.StatusSeeOther)
	}
}
