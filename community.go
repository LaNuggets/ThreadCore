package main

import (
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
	"strings"
)

func Community(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/community/" {
		http.Redirect(w, r, "/search", http.StatusSeeOther)
		return
	}

	tmpl, err := template.ParseFiles("./templates/community.html") // Read the home page
	if err != nil {
		log.Printf("\033[31mError parsing template: %v\033[0m", err)
		http.Error(w, "Internal error, template not found.", http.StatusInternalServerError)
		return
	}

	communityName := strings.ReplaceAll(r.URL.Path, "/community/", "")
	if strings.Contains(communityName, "/") {
		http.Redirect(w, r, "/search", http.StatusSeeOther)
	}

	communityPage := struct {
		Name string
	}{
		Name: communityName,
	}

	//Upload Image
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
	//

	err = tmpl.Execute(w, communityPage)
	if err != nil {
		log.Printf("\033[31mError executing template: %v\033[0m", err)
		http.Error(w, "Internal error", http.StatusInternalServerError)
		return
	}
}
