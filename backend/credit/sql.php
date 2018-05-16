<?php
 $conn = mysqli_connect("localhost","root","","credit_card") or die("Connect DB Error");
 $code = $_POST["code"];
 $month = $_POST["month"];
 $day = $_POST["day"];
 $name = $_POST["nameCard"];
 $cvv = $_POST["cvv"];
 $cvv_check = $conn->query("SELECT cvv FROM member WHERE id_card='".$code."'")->fetch_assoc();
 $order_detail = $conn->query("SELECT * FROM payment WHERE order_id='".$_POST["order"]."'")->fetch_assoc();
 if ($cvv_check["cvv"] == $cvv) {
     $query = $conn->query("SELECT * FROM member WHERE id_card='".$code."'")->fetch_assoc();
     if ($order_detail["amount"] < $query["money"]) {
        echo "DONE";
         $conn->query("UPDATE member SET money = money-".$order_detail["amount"]." WHERE id_card =".$code);
         $conn->query("UPDATE payment SET status=1 WHERE order_id='".$_POST["order"]."'");
     }
     else
     {
        echo "NEMO";
     }
 }
 else
 {
    echo "INPS";
 }
?>