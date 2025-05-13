package handlers

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"zencloud-backend/pkg/utils"
	"zencloud-backend/pkg/utils/environment"
)

func DownloadHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		utils.HandleError(w, fmt.Errorf("method not allowed"), http.StatusMethodNotAllowed)
		return
	}

	// Get the filename from the query parameters
	filename := r.URL.Query().Get("filename")
	if filename == "" {
		utils.HandleError(w, fmt.Errorf("filename is required"), http.StatusBadRequest)
		return
	}
	// Open the file for reading
	file, err := os.Open(environment.GetStorageLocation() + filename)
	if err != nil {
		utils.HandleError(w, err, http.StatusNotFound)
		return
	}
	defer file.Close()
	// Set the headers for the response
	info, err := file.Stat()
	if err != nil {
		utils.HandleError(w, err, http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Disposition", "attachment; filename="+filename)
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Transfer-Encoding", "binary")
	w.Header().Set("Content-Length", fmt.Sprintf("%d", info.Size()))

	// Write the file content to the response
	_, err = io.Copy(w, file)
	if err != nil {
		utils.HandleError(w, err, http.StatusInternalServerError)
		return
	}
	// Respond with success
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("{\"message\": \"File downloaded successfully\"}"))
}
