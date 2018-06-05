package main

import (
	"encoding/json"
	"fmt"
	"go-graphql-api/musicutil"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/graphql-go/graphql"
)

func GetDiscs(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576))
	if err != nil {
		log.Fatalln("Error GetDiscs", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := r.Body.Close(); err != nil {
		log.Fatalln("Error GetDiscs", err)
	}
	var apolloQuery map[string]interface{}
	fmt.Println("Received request GetDiscs")
	if err := json.Unmarshal(body, &apolloQuery); err != nil { // unmarshall body contents as a type query
		fmt.Println(err)
		fmt.Println("Error on Unmarshalling!!!")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Fatalln("Error GetDiscs unmarshalling data", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	query := apolloQuery["query"]
	variables := apolloQuery["variables"]
	result := graphql.Do(graphql.Params{
		Schema:         musicutil.MusicSchema,
		RequestString:  query.(string),
		VariableValues: variables.(map[string]interface{}),
	})
	json.NewEncoder(w).Encode(result)
	w.WriteHeader(http.StatusOK)
	return
}

func main() {
	router := mux.NewRouter().StrictSlash(true)
	var handler http.Handler
	handler = http.HandlerFunc(GetDiscs)
	router.Methods("POST").Path("/").Name("GetDiscs").Handler(handler)

	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	headersOk := handlers.AllowedHeaders([]string{"X-Requested-With", "Accept", "Content-Type", "Content-Length", "Accept-Encoding", "Accept-Language", "X-CSRF-Token", "Authorization"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "HEAD", "OPTIONS"})

	fmt.Println("Now server is running on port 8090")

	// launch server
	log.Fatal(http.ListenAndServe(":8090",
		handlers.CORS(allowedOrigins, headersOk, allowedMethods)(router)))
}
