package routes

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func RoutePostFile(w http.ResponseWriter, r *http.Request) {

	log.Println("Received POST request to upload file")

	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// Lecture de l'ID de la requête
	id := r.FormValue("id")
	if id == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "ID is required")
		return
	}

	// Vérification de l'existence du dossier
	if _, err := os.Stat("data/" + id); os.IsNotExist(err) {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "Folder with ID %s not found", id)
		return
	} else if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error checking folder for ID %s", id)
		return
	}

	// Récupération du fichier depuis la requête
	file, header, err := r.FormFile("file")
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Error retrieving file from request")
		return
	}
	defer file.Close()

	// Création du fichier dans le dossier correspondant
	filePath := fmt.Sprintf("%s/%s", "data/"+id+"/.ssh", header.Filename)
	newFile, err := os.Create(filePath)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error creating file %s in folder %s", header.Filename, id)
		return
	}
	defer newFile.Close()

	// Copie du contenu du fichier
	_, err = io.Copy(newFile, file)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error copying file content for file %s in folder %s", header.Filename, id)
		return
	}

	// Réponse
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "File %s successfully uploaded to folder %s", header.Filename, id)
}
