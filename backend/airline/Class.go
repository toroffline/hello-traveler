/*
* @Author: OVerKillLeD
* @Date:   2018-04-21 18:32:47
* @Last Modified by:   OVerKillLeD
* @Last Modified time: 2018-04-21 23:27:45
*/

package main

type Airline struct {
    From  string    `json:"Startat"`
    Goto    string    `json:"Doneat"`
    AirlineID  string    `json:"airlineID"`
    Gate  string    `json:"gate"`
    Timetogo  string    `json:"timetogo"`
    Reach  string    `json:"reach_des"`
    Max_seat  string    `json:"max_seat"`
    Reserved_seat  string    `json:"reserved_seat"`
    Price   string `json:"price"`
}

type airlineData Airline
type listairline []Airline