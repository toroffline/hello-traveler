<?php

// Input = booking_id
// Output = true/false

if ($_SERVER['REQUEST_METHOD'] == 'GET') {

    $bid = $_GET['id'];
    
    $sql = "SELECT booking_id 
            FROM booking 
            WHERE booking_id = '$bid'";

    $con = mysqli_connect("localhost","root","","hotel") or die("Connect DB Error");
    $con->query("SET NAMES UTF8");
    $query = $con->query($sql);

    $res = "false";

    while ($result = $query->fetch_array())  {
        $sql = "DELETE FROM booking WHERE booking_id = '$bid'";
        $con->query($sql);
        $res = "true";
        break;
    }

    echo $res;
}

?>