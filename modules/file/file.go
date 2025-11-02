package file

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

const storagePath = "./storage/"

func ensureStorage() {
	if _, err := os.Stat(storagePath); os.IsNotExist(err) {
		os.Mkdir(storagePath, 0755)
	}
}

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Use POST", http.StatusMethodNotAllowed)
		return
	}

	file, header, err := r.FormFile("file")
	if err != nil {
		http.Error(w, "Invalid file", http.StatusBadRequest)
		return
	}
	defer file.Close()

	ensureStorage()
	out, err := os.Create(storagePath + header.Filename)
	if err != nil {
		http.Error(w, "Cannot save file", http.StatusInternalServerError)
		return
	}
	defer out.Close()

	io.Copy(out, file)
	fmt.Fprintf(w, "Uploaded: %s\n", header.Filename)
}

func DownloadHandler(w http.ResponseWriter, r *http.Request) {
	filename := r.URL.Query().Get("file")
	if filename == "" {
		http.Error(w, "Specify file param", http.StatusBadRequest)
		return
	}

	http.ServeFile(w, r, storagePath+filename)
}

func ListHandler(w http.ResponseWriter, r *http.Request) {
	ensureStorage()
	files, err := os.ReadDir(storagePath)
	if err != nil {
		http.Error(w, "Cannot read storage", http.StatusInternalServerError)
		return
	}

	for _, f := range files {
		fmt.Fprintln(w, f.Name())
	}
}
