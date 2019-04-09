package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var certsDB []certificate

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
	r.HandleFunc("/api/certificates", certificateHandler).Methods("GET")
	r.HandleFunc("/api/certificates/{id}", certificateHandler).Methods("POST", "DELETE", "PUT")
	r.HandleFunc("/users/{userid}/certificates", userCertificatesHandler).Methods("GET")
	r.HandleFunc("/api/{certificatesID}/transfer", transferCertificateHandler).Methods("POST")
	r.HandleFunc("/api/{certificatesID}/transfer", transferCertificateHandler).Methods("PUT")

	log.Fatal(http.ListenAndServe(":8000", r))
}

func certificateHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		getCertificates(w, r)
	}
	if r.Method == "POST" {
		createCertificate(w, r)
	}

	if r.Method == "DELETE" {
		deleteCertificate(w, r)
	}

	if r.Method == "PUT" {
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
	err := json.NewEncoder(w).Encode(userCertificates)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func transferCertificateHandler(w http.ResponseWriter, r *http.Request) {
	// @TODO - IMPLEMENT TRANSFER CERTIFICATE
}

// Function Handlers

func getCertificates(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(certsDB)
	if err != nil {
		log.Fatal(err)
	}
}

func getUserCertificates(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(r) // Get vars
	// Loop though certificats and find one with id
	for _, item := range certsDB {
		if item.ID == vars["id"] {
			err := json.NewEncoder(w).Encode(item)
			if err != nil {
				log.Fatal(err)
			}
			return
		}
	}
	err := json.NewEncoder(w).Encode(&certificate{})
	if err != nil {
		log.Fatal(err)
	}
}

func createCertificate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	err := json.NewEncoder(w).Encode(certsDB)
	if err != nil {
		log.Fatal(err)
	}
	var certificate certificate
	err = json.NewDecoder(r.Body).Decode(&certificate)
	if err != nil {
		log.Fatal(err)
	}
	certsDB = append(certsDB, certificate)
	err = json.NewEncoder(w).Encode(certificate)
	if err != nil {
		log.Fatal(err)
	}
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
	err := json.NewEncoder(w).Encode(certsDB)
	if err != nil {
		log.Fatal(err)
	}
	return
}

func updateCertificate(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	for index, item := range certsDB {
		if item.ID == vars["id"] {
			certsDB = append(certsDB[:index], certsDB[index+1:]...)

			var certificate certificate
			err := json.NewDecoder(r.Body).Decode(&certificate)
			if err != nil {
				log.Fatal(err)
			}
			certsDB = append(certsDB, certificate)
			err = json.NewEncoder(w).Encode(certificate)
			if err != nil {
				log.Fatal(err)
			}
			return
		}
	}
	err := json.NewEncoder(w).Encode(certsDB)
	if err != nil {
		log.Fatal(err)
	}

	return
}

func transferCertificate(w http.ResponseWriter, r *http.Request) {

}

func acceptCertificate(w http.ResponseWriter, r *http.Request) {

}
