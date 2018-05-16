//---------------------------------------
start
    commandline go run *.go (it will open at port 8080)
stop
    open task Manager and kill start.go
-----------------------------------------//
Use database MySQL
    file name "morrasoom_airline.sql"
----------------------------------------//
command at port (start.go)
localhost:8080/ (will show all data)
localhost:8080/{from} (will show all startlocation) ex. localhost:8080/chiang mai
localhost:8080/{from}/{stop} (same at top but have stoplocation) ex. localhost:8080/chiang mai/bangkok
localhost:8080/{from}/{stop}/{seat} (add seat) ex. localhost:8080/chiang mai/bangkok/5 ***methods{POST}***
---------------------------------------//
