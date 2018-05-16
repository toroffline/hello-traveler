package main

import (
    "database/sql"
    _"github.com/go-sql-driver/mysql" 
    "fmt"
)
// check ------------------------------------------------------------------------------------------------------
func checkConnected(clientIP string) (string) {
    db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/agent") // ID mySQL plz use your own
    if err != nil {
        panic(err.Error())
    }
    defer db.Close() // close db after compile is finished
    // fu*king golang are not allow to break a fu*king long line mySQL command fu*kyou golang.
    read, err := db.Query("SELECT clientIP FROM currentUser WHERE clientIP = " + "'" + clientIP + "'")
    if err != nil {
        panic(err.Error())
        return "false";
    }
    defer read.Close()
    var temp string
    for read.Next(){
        err =  read.Scan(&temp);
    }    
    if err != nil {
        panic(err.Error())
    }
    return temp

}
// ---------------------------------------------------------------------------------------------------------------------
func addNewClient(clientIP string) (bool){ // update
    db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/agent") 
    if err != nil {
        panic(err.Error())
    }
    defer db.Close() // close db after compile is finished
    fmt.Println("Add " + clientIP)
    stmt, err := db.Query("INSERT INTO `currentUser`(`clientIP`) VALUES ('" + clientIP +"')" )
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

func checkReserveAirline(clientIP string) (bool){ // update
    db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/agent") 
    if err != nil {
        panic(err.Error())
    }
    defer db.Close() // close db after compile is finished
    read, err := db.Query("SELECT `airline_reserve_status` FROM `currentuser` WHERE `clientIP` = '" + clientIP + "'" )
    if err != nil {
        panic(err.Error())
    }
    defer read.Close()
    var temp string
    for read.Next(){
        err =  read.Scan(&temp);
    }    
    if err != nil {
        panic(err.Error())
    }
    if( temp == "0" ){
        return false;
    }else if( temp == "1" ){
        return true;
    }
    return false;
}

func checkReserveHotel(clientIP string) (bool){ // update
    db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/agent") 
    if err != nil {
        panic(err.Error())
    }
    defer db.Close() // close db after compile is finished
    read, err := db.Query("SELECT `hotel_reserve_status` FROM `currentuser` WHERE `clientIP` = '" + clientIP + "'" )
    if err != nil {
        panic(err.Error())
    }
    defer read.Close()
    var temp string
    for read.Next(){
        err =  read.Scan(&temp);
    }    
    if err != nil {
        panic(err.Error())
    }
    if( temp == "0" ){
        return false;
    }else if(temp == "1"){
        return true;
    }
    return false;
}

func reserveAirline(startat string,stopat string,seat string, clientIP string) (bool){ // update
    db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/agent") // ID mySQL plz use your own
    if err != nil {
        panic(err.Error())
    }
    defer db.Close() // close db after compile is finished
    var id = getId(clientIP);
    stmt, err := db.Query("INSERT INTO `reserveAirline`(`id`, `source`, `dest`, `seat`) VALUES ('"+id+"','" + startat+ "','" + stopat + "', " + seat + ")")
    if err != nil {
        panic(err.Error())
    }
    defer stmt.Close()
    stmt2, err2 := db.Query("UPDATE `currentuser` SET `airline_reserve_status`= 1  WHERE `id` ='" + id + "'")
    if err2 != nil {
        panic(err.Error())
    }
    defer stmt2.Close()

// UPDATE airline a
// INNER JOIN province p1 ON a.src = p1.pid INNER JOIN province p2 ON a.des = p2.pid
// SET a.reserved_seat = 0
// WHERE p1.pnameeng = "chiang mai" AND p2.pnameeng = "bangkok"
    return true
}

func reserveHotel(h_id string,datein string,dateout string, quantity string, clientIP string, book_id string) (bool){ // update
    db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/agent") // ID mySQL plz use your own
    if err != nil {
        panic(err.Error())
    }
    defer db.Close() // close db after compile is finished
    var id = getId(clientIP);
    stmt, err := db.Query("INSERT INTO `reservehotel`(`id`, `h_id`, `datein`, `dateout`, `quantity`, `book_id`) VALUES ('" + id +"','" + h_id+ "','" + datein + "', '" + dateout + "'," + quantity +", " + book_id +")")
    if err != nil {
        panic(err.Error())
    }
    defer stmt.Close()
    stmt2, err2 := db.Query("UPDATE `currentuser` SET `hotel_reserve_status`= 1  WHERE `id` ='" + id + "'")
    if err2 != nil {
        panic(err.Error())
    }
    defer stmt2.Close()

// UPDATE airline a
// INNER JOIN province p1 ON a.src = p1.pid INNER JOIN province p2 ON a.des = p2.pid
// SET a.reserved_seat = 0
// WHERE p1.pnameeng = "chiang mai" AND p2.pnameeng = "bangkok"
    return true
}

func getId(clientIP string)(string){
    db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/agent") 
    if err != nil {
        panic(err.Error())
    }
    defer db.Close() // close db after compile is finished
    read, err := db.Query("SELECT `id` FROM `currentuser` WHERE `clientIP` = '" + clientIP + "'" )
    if err != nil {
        panic(err.Error())
    }
    defer read.Close()
    var temp string
    for read.Next(){
        err =  read.Scan(&temp);
    }
    return temp    
}

func delAirOrder(clientIP string)(bool){
    db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/agent") // ID mySQL plz use your own
    if err != nil {
        panic(err.Error())
    }
    defer db.Close() // close db after compile is finished
    var id = getId(clientIP);
    stmt2, err2 := db.Query("UPDATE `currentuser` SET `hotel_reserve_status`= 0  WHERE `id` ='" + id + "'")
    if err2 != nil {
        panic(err.Error())
    }
    defer stmt2.Close()
    stmt, err := db.Query("UPDATE `currentuser` SET `airline_reserve_status`= 0  WHERE `id` ='" + id + "'")
    if err2 != nil {
        panic(err.Error())
    }
    defer stmt.Close()
    stmt3, err3 := db.Query("DELETE FROM `reserveairline` WHERE `id` ='" + id + "'")
    if err3 != nil {
        panic(err.Error())
    }
    defer stmt3.Close()
    stmt4, err4 := db.Query("DELETE FROM `reservehotel` WHERE `id` ='" + id + "'")
    if err4 != nil {
        panic(err.Error())
    }
    defer stmt4.Close()
    stmt5, err5 := db.Query("DELETE FROM `currentuser` WHERE `clientIP` ='" + clientIP + "'")
    if err5 != nil {
        panic(err.Error())
    }
    defer stmt5.Close()
    return true;
}

func getBookId(clientIP string)(string){
     db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/agent") // ID mySQL plz use your own
    if err != nil {
        panic(err.Error())
    }
    defer db.Close() // close db after compile is finished
    var id = getId(clientIP);
     read, err := db.Query("SELECT `book_id` FROM `reservehotel` WHERE `id` = '" + id + "'" )
    if err != nil {
        panic(err.Error())
    }
    defer read.Close()
    var temp string
    for read.Next(){
        err =  read.Scan(&temp);
    }
    return temp    
}