package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
        "strconv"
	"github.com/gorilla/mux"
)

type operands struct {
	Operand1 float64 `json:"operand1,omitempty"`
	Operand2 float64 `json:"operand2,omitempty"`
}

func add(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	//decoder := json.NewDecoder(r.Body)
	//var request operands
	//err := decoder.Decode(&request)
	//if err != nil {
	//	http.Error(w, "{}", http.StatusInternalServerError)
	//}
        vars := mux.Vars(r)
        //Operand1 := vars["Operand1"]
        Operand1, _ := strconv.Atoi(vars["Operand1"])
        //Operand2 := vars["Operand2"]
        Operand2, _ := strconv.Atoi(vars["Operand2"])
	response := Operand1 + Operand2
	json.NewEncoder(w).Encode(response)
}

func subtract(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
//	decoder := json.NewDecoder(r.Body)
//	var request operands
//	err := decoder.Decode(&request)
//	if err != nil {
//		http.Error(w, "{}", http.StatusInternalServerError)
//	}
//	response := request.Operand1 - request.Operand2
        vars := mux.Vars(r)
        //Operand1 := vars["Operand1"]
        Operand1, _ := strconv.Atoi(vars["Operand1"])
        //Operand2 := vars["Operand2"]
        Operand2, _ := strconv.Atoi(vars["Operand2"])
        response := Operand1 - Operand2

	json.NewEncoder(w).Encode(response)
}

func multiply(w http.ResponseWriter, r *http.Request) {
//	w.Header().Set("Content-Type", "application/json")
//	decoder := json.NewDecoder(r.Body)
//	var request operands
//	err := decoder.Decode(&request)
//	if err != nil {
//		http.Error(w, "{}", http.StatusInternalServerError)
//	}
//	response := request.Operand1 * request.Operand2
        vars := mux.Vars(r)
        //Operand1 := vars["Operand1"]
        Operand1, _ := strconv.Atoi(vars["Operand1"])
        //Operand2 := vars["Operand2"]
        Operand2, _ := strconv.Atoi(vars["Operand2"])
        response := Operand1 * Operand2

	json.NewEncoder(w).Encode(response)
}

func status(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "ok")
}

func main() {

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/status", status).Methods("GET")
	router.HandleFunc("/add/{Operand1}/{Operand2}", add).Methods("GET")
	router.HandleFunc("/subtract/{Operand1}/{Operand2}", subtract).Methods("GET")
	router.HandleFunc("/multiply/{Operand1}/{Operand2}", multiply).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}
