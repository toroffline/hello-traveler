<?php

// Input = hotel_id, booking_qty, check_in, check_out
// Output = booking_id

function getBookId() {
    $sql = "SELECT      booking_id
            FROM        booking
            ORDER BY    booking_id DESC
            LIMIT       1";

    $con = mysqli_connect("localhost","root","","hotel") or die("Connect DB Error");
    $con->query("SET NAMES UTF8");
    $query = $con->query($sql);

    while ($result = $query->fetch_array())  {
        return $result["booking_id"];
    }
}

if ($_SERVER['REQUEST_METHOD'] == 'GET') {

    $hid        = $_GET['id'];
    $quantity   = $_GET['qty'];
    $date_in    = $_GET['in'];
    $date_out   = $_GET['out'];

    // $hid        = "1";
    // $quantity   = "1";
    // $date_in    = "2018-04-20";
    // $date_out   = "2018-04-25";

    $sql = "SELECT  a.price,
                    (a.room_qty - (IFNULL((
                        SELECT  SUM(b.qty)
                        FROM    booking AS b
                        WHERE   a.hotel_id = b.hotel_id
                        AND     (b.check_in  BETWEEN '$date_in' AND '$date_out'
                        OR       b.check_out BETWEEN '$date_in' AND '$date_out')
                        GROUP BY b.hotel_id
                    ),0))) AS avl
            FROM    hotel AS a
            WHERE   a.hotel_id = '$hid'";

    $con = mysqli_connect("localhost","root","","hotel") or die("Connect DB Error");
    $con->query("SET NAMES UTF8");
    $query = $con->query($sql);

    $bid = "false";

    while ($result = $query->fetch_array())  {
        $price  = $result["price"];
        $avl    = $result["avl"];
        
        if ($avl >= $quantity)
        {
            $price = $price * $quantity;
            $bid =  getBookId() + 1;

            $sql = "INSERT INTO booking
                    VALUES ('$bid', '$hid', '$date_in', '$date_out', '$quantity', '$price')";

            $con->query($sql);
            break;
        }
    }

    echo $bid;
}

?>