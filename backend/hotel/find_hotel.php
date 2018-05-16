<?php

$resArray = array();

// Input = province, booking_qty, check_in, check_out
// Output = hotel_id, hotel_name, tel, price, available_room

if ($_SERVER['REQUEST_METHOD'] == 'GET') {

    $province   = $_GET['provi'];
    $quantity   = $_GET['qty']; // Qty
    $date_in    = $_GET['in'];
    $date_out   = $_GET['out'];

    // $province   = "เชียงใหม่";
    // $quantity   = "1";
    // $date_in    = "2018-04-20";
    // $date_out   = "2018-04-25";

    $sql = "SELECT  a.hotel_id,
                    a.name,
                    a.telephone,
                    a.price,
                    (a.room_qty - (IFNULL((
                        SELECT  SUM(b.qty)
                        FROM    booking AS b
                        WHERE   a.hotel_id = b.hotel_id
                        AND     (b.check_in  BETWEEN '$date_in' AND '$date_out'
                        OR      b.check_out BETWEEN '$date_in' AND '$date_out')
                        GROUP BY b.hotel_id
                    ),0))) AS avl
            FROM    hotel AS a
            WHERE   a.province = '$province'";

    $con = mysqli_connect("localhost","root","","hotel") or die("Connect DB Error");
    $con->query("SET NAMES UTF8");
    $query = $con->query($sql);

    while ($result = $query->fetch_array()) {
        $id     = $result["hotel_id"];
        $name   = $result["name"];
        $tel    = $result["telephone"];
        $price  = $result["price"];
        $avl    = $result["avl"];
        
        if ($avl >= $quantity) {
            array_push( $resArray, 
                array(
                    "Id"    => $id,
                    "Name"  => $name,
                    "Tel"   => $tel,
                    "Price" => $price,
                    "Avl"   => $avl
                )
            );
        }
    }

    $json = json_encode($resArray);
    echo $json;
}

?>