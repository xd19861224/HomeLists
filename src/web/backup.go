package web

import (
	"log"
	"strings"
	"html"
	"net/http"
	"io"
	"os"
	"time"
)

func backup(w http.ResponseWriter, r *http.Request) {
	urlString := html.EscapeString(r.URL.Path)
	tags := strings.Split(urlString, "/")

	method := tags[2]
	currentTime := time.Now()

	switch method {
	case "create":
		backupString := "-" + currentTime.Format("2006-01-02T15-04")

		sourceFile, _ := os.Open(Data.Config.DbPath)
		defer sourceFile.Close()

		newFile, _ := os.Create(Data.Config.DbPath + backupString)
		defer newFile.Close()

		io.Copy(newFile, sourceFile)

		http.Redirect(w, r, r.Header.Get("Referer"), 302)

	case "download":
		filename := "HomeLists-backup-" + currentTime.Format("2006-01-02T15-04") + ".db"

		w.Header().Set("Content-Disposition", "attachment; filename="+filename)
		w.Header().Set("Content-Type", "application/octet-stream")
		http.ServeFile(w, r, Data.Config.DbPath)
	}
}

func upload(w http.ResponseWriter, r *http.Request) {
    uploadFile, _, err := r.FormFile("dbfile")
	if err != nil {
		log.Println("Upload error:", err)
	} else {
		defer uploadFile.Close()

		newFile, _ := os.Create(Data.Config.DbPath)
		defer newFile.Close()

		io.Copy(newFile, uploadFile)
	}
	http.Redirect(w, r, r.Header.Get("Referer"), 302)
}