package main

import (
    "bytes"
    "encoding/json"
    "fmt"
    "io/ioutil"
    "net/http"
    // "github.com/jmoiron/jsonq"
    "strings"
    "strconv"
    "github.com/gorilla/mux"
    "log"
    "crypto/rand"
    "net/url"
    "github.com/gorilla/handlers"
    "github.com/Tomasen/realip"
)

 var use_url = "http://10.80.145.232:8080"

func main() {
    fmt.Println("Starting the agent api...")
    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/getAirline", getAllAirline).Methods("GET")
    router.HandleFunc("/getAirline/{from}", getAirline).Methods("GET")
    router.HandleFunc("/getAirline/{from}/{goat}", getAirlineDest).Methods("GET")
    router.HandleFunc("/getAirline/{from}/{goat}/{seat}", takeASeat).Methods("POST")
    router.HandleFunc("/getPrice/{from}/{goat}", getPriceReturn).Methods("GET")
    router.HandleFunc("/getHotel/{source}/{quantity}/{datein}/{dateout}", find_hotel).Methods("GET")
    router.HandleFunc("/getHotel/{id}/{quantity}/{datein}/{dateout}", reservHotel).Methods("POST")
    // router.HandleFunc("/godFunction1/{from}/{goat}/{seat}/", getPriceReturn).Methods("GET")
    // log.Fatal(http.ListenAndServe(":3000", router))
    router.HandleFunc("/checkReserveAirline", checkReserveAirlineStatus).Methods("GET")
    router.HandleFunc("/checkReserveHotel", checkReserveHotelStatus).Methods("GET")
    router.HandleFunc("/cancelOrder", cancelOrder).Methods("GET")
    router.HandleFunc("/createOrder/{cost}", create_order).Methods("GET")
    router.HandleFunc("/payment/{card_num}/{month}/{day}/{nameCard}/{cvv}/{order}", paypaypay).Methods("POST")
    router.HandleFunc("/getIpAddress", getIpAddress).Methods("GET")
    log.Fatal(http.ListenAndServe(":3000", handlers.CORS(handlers.AllowedHeaders([]string{"application/json", "Content-Type", "Authorization"}), handlers.AllowedMethods([]string{"GET", "POST", "PUT", "HEAD", "OPTIONS"}), handlers.AllowedOrigins([]string{"*"}))(router)))

}

func getAllAirline(w http.ResponseWriter, r *http.Request) {
    response, err := http.Get(use_url + "/getAirline/")
    if err != nil {
        fmt.Printf("The HTTP request failed with error %s\n", err)
    } else {
        data, _ := ioutil.ReadAll(response.Body)
        fmt.Fprintln(w,string(data))

    }
}

func getAirline(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    province := vars["from"]
    response, err := http.Get(use_url + "/getAirline/" + province)
    if err != nil {
        fmt.Printf("The HTTP request failed with error %s\n", err)
    } else {
        data, _ := ioutil.ReadAll(response.Body)
        fmt.Fprintln(w,string(data))
    }
}

func getAirlineDest(w http.ResponseWriter, r *http.Request) {
    var check string
    var check2 bool
    clientIP := realip.FromRequest(r)
    fmt.Println("GET / from", clientIP)
    check = checkConnected(clientIP) // add person (string,string,string)
    fmt.Println("check:" + check)
    if check != "false" && check == "" {
      check2 = addNewClient(clientIP)
      if(check2 == true){
            fmt.Println("add new client")
        }else{
            fmt.Println("no need add new client")
        }   
    } 
    vars := mux.Vars(r)
    source := vars["from"]
    destination := vars["goat"]
    response, err := http.Get(use_url + "/getAirline/" + source + "/" + destination)
    if err != nil {
        fmt.Printf("The HTTP request failed with error %s\n", err)
    } else {
       data, _ := ioutil.ReadAll(response.Body)
       fmt.Fprintln(w,string(data)) 
    }

}

func takeASeat(w http.ResponseWriter, r *http.Request){
    clientIP := realip.FromRequest(r)
    fmt.Println("GET / from", clientIP)
    vars := mux.Vars(r)
    source := vars["from"]
    destination  := vars["goat"]
    seat := vars["seat"]
    jsonData := map[string]string{}
    jsonValue, _ := json.Marshal(jsonData)
    fmt.Printf(use_url + "/getAirline/" + source + "/" + destination + "/" +  seat)
    response, err := http.Post(use_url+ "/getAirline/" + source + "/" + destination + "/" +  seat, "application/json", bytes.NewBuffer(jsonValue))
    if err != nil {
        fmt.Printf("The HTTP request failed with error %s\n", err)
    } else {
        data, _ := ioutil.ReadAll(response.Body)
        fmt.Fprintln(w,string(data))
        reserveAirline(source,destination,seat,clientIP)
    }
}

func getPriceReturn(w http.ResponseWriter, r *http.Request){
    vars := mux.Vars(r)
    source := vars["from"]
    destination  := vars["goat"]
    fmt.Printf("http://10.80.146.227/getAirline/" + source + "/" + destination)
    response, err := http.Get(use_url + "/getAirline/" + source + "/" + destination)
    if err != nil {
        fmt.Printf("The HTTP request failed with error %s\n", err)
    } else {
        data, _ := ioutil.ReadAll(response.Body)
        price := strconv.Itoa(getPrice(string(data)))
        fmt.Fprintln(w,price)
        
    }
}

func getPrice(data string) int{
    i := strings.Index(string(data), "price")
    //fmt.Println("Index: ", i)
    i = i+8
    priceString := ""
    for string(data[i]) != "\"" {
        priceString += string(data[i])
        i++
    }
    price, err := strconv.Atoi(priceString)
    if err != nil {
        fmt.Println(err)
    }
    //fmt.Println(price)
    return price
}

func tokenGenerator() string {
    b := make([]byte, 5)
    rand.Read(b)
    return fmt.Sprintf("%x", b)
}

func create_order(w http.ResponseWriter, r *http.Request) {
    dest := "http://127.0.0.1/concurrent/credit/api-payment.php?"
    method := "method=set&order_id="
    vars := mux.Vars(r)
    cost := vars["cost"]
    order_id := tokenGenerator()
    fmt.Println("order_id: " + order_id)
    URL := dest + method + order_id + "&amount=" + cost
    // fmt.Println(URL)
    res, err := http.Get(URL)
    if err != nil {
        // log.Fatal(err)
    }
    body, err := ioutil.ReadAll(res.Body)
    res.Body.Close()
    
    result := string(body[len(body)-4:len(body)])

    // fmt.Println(result)
    if result == "DONE" {
        fmt.Fprintln(w,order_id)
    }else{
       fmt.Fprintln(w,"Failed")
    }
}

func paypaypay(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    card_num := vars["card_num"]
    month := vars["month"]
    day := vars["day"]
    nameCard := vars["nameCard"]
    cvv := vars["cvv"]
    order_id := vars["order"]
    location := "http://127.0.0.1/concurrent/credit/sql.php"
    data := url.Values{
        "code": {card_num},
        "month": {month},
        "day": {day},
        "nameCard": {nameCard},
        "cvv": {cvv},
        "order": {order_id},
    }
    fmt.Println("eiie")
    resp, err := http.PostForm(
        location,
        data,
    )

    if err != nil {
        log.Println("post err")
        panic(err)
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    
    if err != nil {
        log.Println("ioutil err")
        panic(err)
    }
    result := string(body[len(body)-4:len(body)])
    fmt.Println(result)
    if result == "INPS" {
        fmt.Fprintln(w,"FAIL: Password incorrect")
    }else if result == "NEMO" {
        fmt.Fprintln(w,"FAIL: Not enough money")
    }else if result == "DONE"{
        fmt.Fprintln(w,"Success")
    }else{
        fmt.Fprintln(w,string(body))
    }
}

func find_hotel(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    provi := vars["source"]
    quan := vars["quantity"]
    d_in := vars["datein"]
    d_out := vars["dateout"]
    dest := "http://127.0.0.1/concurrent/hotel/find_hotel.php?"
    URL := dest + "provi=" + provi + "&qty=" + quan + "&in=" + d_in + "&out=" + d_out
 
    res, err := http.Get(URL)
    if err != nil {
        log.Fatal(err)
    }
    body, err := ioutil.ReadAll(res.Body)

    res.Body.Close()
    // fmt.Println(body)
    fmt.Fprintln(w,string(body))

    // var msg Message
    // err = json.Unmarshal(body, &msg)
    // if err != nil {
    //  log.Fatal(err)
    //  return
    // }

    // fmt.Println(msg) 

    // output, err := json.Marshal(msg)
    // if err != nil {
    //  log.Fatal(err)
    //  return
    // }

    // fmt.Println(output)

    // y, err := json.Marshal(body)
    // if err != nil {
    //     panic(err)
    // }
    // fmt.Println(string(y))
}

func reservHotel(w http.ResponseWriter, r *http.Request) {
    clientIP := realip.FromRequest(r)
    vars := mux.Vars(r)
    h_id := vars["id"]
    quan := vars["quantity"]
    d_in := vars["datein"]
    d_out := vars["dateout"]
    dest := "http://127.0.0.1/concurrent/hotel/book_hotel.php?"
    URL := dest + "id=" + h_id + "&qty=" + quan + "&in=" + d_in + "&out=" + d_out
    
    res, err := http.Get(URL)
    if err != nil {
        log.Fatal(err)
    }
    body, err := ioutil.ReadAll(res.Body)
    reserveHotel(h_id, d_in, d_out, quan,clientIP, string(body))
    res.Body.Close()
    // fmt.Println(body)
    fmt.Fprintln(w,string(body))

}

func checkReserveAirlineStatus(w http.ResponseWriter, r *http.Request) {
      var check string
      clientIP := realip.FromRequest(r)
    check = checkConnected(clientIP) // add person (string,string,string)
    fmt.Println("check:" + check)
    if check != "false" && check == "" {
      check2 := addNewClient(clientIP)
      if(check2 == true){
            fmt.Println("add new client")
        }else{
            fmt.Println("no need add new client")
        }   
    } 
      check3 := checkReserveAirline(clientIP)
      if( check3 == true ){
        fmt.Fprintln(w,"1") 
      }else{
        fmt.Fprintln(w,"0") 
      }
}


func checkReserveHotelStatus(w http.ResponseWriter, r *http.Request) {
      var check bool
      clientIP := realip.FromRequest(r)
      check = checkReserveHotel(clientIP)
      if( check == true ){
        fmt.Fprintln(w,"1") 
      }else{
        fmt.Fprintln(w,"0") 
      }
}



func cancelOrder(w http.ResponseWriter, r *http.Request){
    var check bool
    clientIP := realip.FromRequest(r)
    book_id := getBookId(clientIP);
    dest := "http://127.0.0.1/concurrent/hotel/cancel_hotel.php?"
    URL := dest + "id=" + book_id
    res, err := http.Get(URL)
    if err != nil {
        log.Fatal(err)
    }

    res.Body.Close()

    check = delAirOrder(clientIP)

      if( check == true ){
        fmt.Fprintln(w,"1") 
      }else{
        fmt.Fprintln(w,"0") 
      }

}


func getIpAddress(w http.ResponseWriter, r *http.Request) {
     clientIP := realip.FromRequest(r)
     fmt.Fprintln(w,string(clientIP))
}