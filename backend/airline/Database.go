package main

import (
    "database/sql"
    _"github.com/go-sql-driver/mysql" // _ mean i will use it don't worry. i will. plz don't fu*king error.
    // "fmt"
    "strconv"
)
// check ------------------------------------------------------------------------------------------------------
func getAirline(startat string,stopat string) (listairline []Airline) {
    db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/morrasoom_airline") // ID mySQL plz use your own
    if err != nil {
        panic(err.Error())
    }
    defer db.Close() // close db after compile is finished
    // fu*king golang are not allow to break a fu*king long line mySQL command fu*kyou golang.
    var combine string
    if (startat != "" || stopat != ""){
        combine = "WHERE "
    }
    if (startat != ""){
        combine = combine+"p1.pnameeng = '"+startat+"'"
    }
    if (stopat != ""){
        if (startat != ""){
            combine = combine + " and "
        }
        combine = combine+"p2.pnameeng = '"+stopat+"'"
    }
    read, err := db.Query("SELECT p1.pnameeng as Startat, p2.pnameeng as Doneat, airLineID, gate, timetogo, reach_des, max_seat,reserved_seat, price FROM airline a INNER JOIN province p1 ON a.src = p1.pid INNER JOIN province p2 ON a.des = p2.pid "+combine+"")
    if err != nil {
        panic(err.Error())
    }
    defer read.Close()

    for read.Next() {
        var airlineData Airline
        err = read.Scan(&airlineData.From, &airlineData.Goto, 
        &airlineData.AirlineID, &airlineData.Gate, &airlineData.Timetogo, 
        &airlineData.Reach, &airlineData.Max_seat, &airlineData.Reserved_seat, &airlineData.Price )

        if err != nil {
            panic(err.Error())
        }

        listairline = append(listairline,airlineData)
    }
    return
}
// ---------------------------------------------------------------------------------------------------------------------
// update ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++
func addsentairline(startat string,stopat string,seat string) (bool){ // update
    airlinedata := (getAirline(startat,stopat)[0])

    temp, _ := strconv.Atoi(airlinedata.Reserved_seat)
    temp2, _ := strconv.Atoi(seat)
    temp3, _ := strconv.Atoi(airlinedata.Max_seat)
    if (temp+temp2 > temp3 || temp+temp2 < 0){
        return false
    }


    db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/morrasoom_airline") // ID mySQL plz use your own
    if err != nil {
        panic(err.Error())
    }
    defer db.Close() // close db after compile is finished

    var combine,key string
    temp4 := strconv.Itoa(temp + temp2)
    key = "a.reserved_seat = "+ temp4
    combine = "p1.pnameeng = '"+ startat +"' AND p2.pnameeng = '"+ stopat +"'"
    stmt, err := db.Query("UPDATE airline a INNER JOIN province p1 ON a.src = p1.pid INNER JOIN province p2 ON a.des = p2.pid SET "+ key +" WHERE "+ combine +"")
    if err != nil {
        panic(err.Error())
    }
    defer stmt.Close()
// UPDATE airline a
// INNER JOIN province p1 ON a.src = p1.pid INNER JOIN province p2 ON a.des = p2.pid
// SET a.reserved_seat = 0
// WHERE p1.pnameeng = "chiang mai" AND p2.pnameeng = "bangkok"
    return true
}
// ++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++++