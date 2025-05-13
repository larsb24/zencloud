package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"zencloud-backend/pkg/utils"
	"zencloud-backend/pkg/utils/environment"
)

func UploadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		utils.HandleError(w, fmt.Errorf("method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	err := r.ParseMultipartForm(10 << 20) // limit upload size to 10 MB
	if err != nil {
		utils.HandleError(w, err, http.StatusBadRequest)
		return
	}

	file, fileHeader, err := r.FormFile("file")
	if err != nil {
		utils.HandleError(w, err, http.StatusBadRequest)
		return
	}
	defer file.Close()

	// Create a file on the server to save the uploaded content
	dst, err := os.Create(environment.GetStorageLocation() + fileHeader.Filename)
	if err != nil {
		utils.HandleError(w, err, http.StatusInternalServerError)
		return
	}
	//TODO check if file already exists
	defer dst.Close()

	// Copy the uploaded file's content to the destination file
	_, err = io.Copy(dst, file)
	if err != nil {
		utils.HandleError(w, err, http.StatusInternalServerError)
		return
	}

	// Respond with success
	w.Header().Set("Content-Type", "application/json")
	response := fmt.Sprintf("{\"filename\": \"%s\", \"size\": %d}", fileHeader.Filename, fileHeader.Size)
	w.Write([]byte(response))
}
