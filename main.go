// get cert
// delete cert
// transfer cert
// accept cert
// create cert

package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var certs []Certificate

type User struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

type Certificate struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	CreatedAt int    `json:"createdAt"`
	OwnerID   string `json:"ownerID"`
	Year      int    `json:"year"`
	Note      string `json:"note"`
}

type Transfer struct {
	To     string `json:"id"`
	Status string `json:"status"`
}

func main() {
	// Init router
	r := mux.NewRouter()

	// Mock Data - @todo - implement DB
	certs = append(certs, Certificate{ID: "1", Title: "Certificate 1", CreatedAt: 000000, OwnerID: "Owner1", Year: 2019, Note: "Note Here"})
	certs = append(certs, Certificate{ID: "2", Title: "Certificate 2", CreatedAt: 000000, OwnerID: "Owner2", Year: 2019, Note: "Note Here"})
	certs = append(certs, Certificate{ID: "3", Title: "Certificate 3", CreatedAt: 000000, OwnerID: "Owner3", Year: 2019, Note: "Note Here"})
	certs = append(certs, Certificate{ID: "4", Title: "Certificate 4", CreatedAt: 000000, OwnerID: "Owner4", Year: 2019, Note: "Note Here"})
	certs = append(certs, Certificate{ID: "5", Title: "Certificate 5", CreatedAt: 000000, OwnerID: "Owner5", Year: 2019, Note: "Note Here"})

	// RouteHandlers / Endpoints
	r.HandleFunc("/api/certs", getCertificates).Methods("GET", "POST", "DELETE", "UPDATE")
	r.HandleFunc("/api/certs/{id}", getUserCertificate).Methods("GET")
	r.HandleFunc("/api/transfer", transferCertificate).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8000", r))
}

func getCertificates(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(certs)
}

func getUserCertificate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	params := mux.Vars(r) // Get Params
	// Loop though certificats and find one with id
	for _, item := range certs {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&Certificate{})
}

func createCertificate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(certs)

	var certificate Certificate
	_ = json.NewDecoder(r.Body).Decode(&certificate)
	certificate.ID = strconv.Itoa(rand.Intn(10000000)) // Mock ID
	//certificates = append(certificates, certificate)
	json.NewEncoder(w).Encode(certificate)
}

func deleteCertificate(w http.ResponseWriter, r *http.Request) {

}

func updateCertificate(w http.ResponseWriter, r *http.Request) {

}

func transferCertificate(w http.ResponseWriter, r *http.Request) {

}

func acceptCertificate(w http.ResponseWriter, r *http.Request) {

}
