package main

import (
    "log"
    "net/http"
    "github.com/gorilla/mux"
    "fmt"
    "encoding/json"
   "github.com/gorilla/handlers"
)

func main() {

    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/getAirline", getall).Methods("GET")
    router.HandleFunc("/getAirline/{from}", getfrom).Methods("GET")
    router.HandleFunc("/getAirline/{from}/{goat}", getfromgo).Methods("GET")
    router.HandleFunc("/getAirline/{from}/{goat}/{seat}", addseat).Methods("POST")

 log.Fatal(http.ListenAndServe(":8080", handlers.CORS(handlers.AllowedHeaders([]string{"application/json", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router)))
}

func getall(w http.ResponseWriter, r *http.Request) {
    airlinedata, err := json.Marshal(getAirline("","")) // all data in db
    if err != nil {
        panic(err)
    }
    fmt.Fprintln(w,string(airlinedata))
}

func getfrom(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    from := vars["from"]
    airlinedata, err := json.Marshal(getAirline(from,"")) // from (string,"")
    if err != nil {
        panic(err)
    }
    fmt.Fprintln(w,string(airlinedata))
}

func getfromgo(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    from := vars["from"]
    goat := vars["goat"]
    airlinedata, err := json.Marshal(getAirline(from,goat)) // from to (string,string)
    if err != nil {
        panic(err)
    }
    fmt.Fprintln(w,string(airlinedata))
}

func addseat(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    from := vars["from"]
    goat := vars["goat"]
    seat := vars["seat"]
    var check bool
    check = addsentairline(from,goat,seat) // add person (string,string,string)
    if check == false {
        fmt.Fprintln(w,"Fail ,Can't Do it ,Run out of seat")
    } else {
        fmt.Fprintln(w,"Done")
    }
}