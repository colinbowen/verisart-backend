// get cert
// delete cert
// transfer cert
// accept cert
// create cert

package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var certsDB []certificate

type app struct {
	Router *mux.Router
	DB     *sql.DB
}

type user struct {
	ID    string `json:"id"`
	Email string `json:"email"`
	Name  string `json:"name"`
}

type certificate struct {
	ID        string `json:"id"`
	Title     string `json:"title"`
	CreatedAt int    `json:"createdAt"`
	OwnerID   string `json:"ownerID"`
	Year      int    `json:"year"`
	Note      string `json:"note"`
}

type transfer struct {
	To     string `json:"id"`
	Status string `json:"status"`
}

func main() {
	// Init router
	r := mux.NewRouter()

	// Mock Data - @todo - implement DB
	certsDB = append(certsDB, certificate{ID: "1", Title: "certificate 1", CreatedAt: 000000, OwnerID: "Owner1", Year: 2019, Note: "Note Here"})
	certsDB = append(certsDB, certificate{ID: "6", Title: "certificate 6", CreatedAt: 000000, OwnerID: "Owner1", Year: 2019, Note: "Note Here"})
	certsDB = append(certsDB, certificate{ID: "2", Title: "certificate 2", CreatedAt: 000000, OwnerID: "Owner2", Year: 2019, Note: "Note Here"})
	certsDB = append(certsDB, certificate{ID: "3", Title: "certificate 3", CreatedAt: 000000, OwnerID: "Owner3", Year: 2019, Note: "Note Here"})
	certsDB = append(certsDB, certificate{ID: "4", Title: "certificate 4", CreatedAt: 000000, OwnerID: "Owner4", Year: 2019, Note: "Note Here"})
	certsDB = append(certsDB, certificate{ID: "5", Title: "certificate 5", CreatedAt: 000000, OwnerID: "Owner5", Year: 2019, Note: "Note Here"})

	// RouteHandlers / Endpoints
	// Get All certificates
	r.HandleFunc("/api/certificates", certificateHandler).Methods("GET")
	// Create / Update / Delete Certificates
	r.HandleFunc("/api/certificates/{id}", certificateHandler).Methods("POST", "DELETE", "PUT")
	// Get all user owned certificates
	r.HandleFunc("/users/{userid}/certificates", userCertificatesHandler).Methods("GET")
	// Transfer a certificate (Create / Accept)
	r.HandleFunc("/api/transfer", transferCertificateHandler).Methods("PATCH")

	log.Fatal(http.ListenAndServe(":8000", r))
}

func certificateHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	if r.Method == "GET" {
		// get certificate
		getCertificates(w, r)
	}
	if r.Method == "POST" {
		// create certificate
		createCertificate(w, r)
	}

	if r.Method == "DELETE" {
		// delete certificate
		deleteCertificate(w, r)
	}

	if r.Method == "PUT" {
		// update certificate
		updateCertificate(w, r)
	}
}

func userCertificatesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	var userCertificates []certificate
	for _, item := range certsDB {
		if item.ID == vars["userid"] {
			userCertificates = append(userCertificates, item)
		}
	}
	json.NewEncoder(w).Encode(userCertificates)
	return
}

func transferCertificateHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	fmt.Println(vars)

}

// Function Handlers

func getCertificates(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(certsDB)
}

func getUserCertificates(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r) // Get vars
	// Loop though certificats and find one with id
	for _, item := range certsDB {
		if item.ID == vars["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	json.NewEncoder(w).Encode(&certificate{})
}

func createCertificate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(certsDB)

	var certificate certificate
	_ = json.NewDecoder(r.Body).Decode(&certificate)
	certsDB = append(certsDB, certificate)
	json.NewEncoder(w).Encode(certificate)
}

func deleteCertificate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	for index, item := range certsDB {
		if item.ID == vars["id"] {
			certsDB = append(certsDB[:index], certsDB[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(certsDB)
	return
}

func updateCertificate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	for index, item := range certsDB {
		if item.ID == vars["id"] {
			certsDB = append(certsDB[:index], certsDB[index+1:]...)

			var certificate certificate
			_ = json.NewDecoder(r.Body).Decode(&certificate)
			certsDB = append(certsDB, certificate)
			json.NewEncoder(w).Encode(certificate)
			return
		}
	}
	json.NewEncoder(w).Encode(certsDB)
	return
}

func transferCertificate(w http.ResponseWriter, r *http.Request) {

}

func acceptCertificate(w http.ResponseWriter, r *http.Request) {

}
