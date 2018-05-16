<?php
// ขั้นตอนการใช้งาน
// - ส่งค่าแบบ GET/POST โดยมี method ให้เลือกคือ 
//  1.get คือแสดงข้อมูล order ที่ร้องขอว่าชำระเงินสำเร็จหรือไม่ 
//  2.set คือ สร้างรายการชำระเงินใหม่โดยต้องส่งตัวแปรมาให้ครบทั้ง 2 ตัวแปรคือ order_id,amount
// ความหมาย API
// order-id : รหัสรายการสั่งซื้อ
// status = 0 : ยังไม่ได้ชำระ
// status = 1 : ชำระเงินสำเร็จ
// status = 2 : ไม่มีข้อมูลรายการสั่งซื้อที่ร้องขอ
header("Access-Control-Allow-Origin: *");
header("Access-Control-Allow-Headers: access");
header("Access-Control-Allow-Methods: GET,POST");
header("Access-Control-Allow-Credentials: true");
header('Content-Type: application/json;charset=utf-8');
if (isset($_GET["order_id"]))
{
    if (isset($_GET["method"]) && strcmp($_GET["method"],"get") == 0) //รับ api แบบ GET และ ส่งข้อมูลออก
    {
        $conn = mysqli_connect("localhost","root","","credit_card") or die("Connect DB Error");
        $conn->query("SET NAMES UTF8");
        $query = $conn->query("SELECT * FROM payment WHERE order_id='".$_GET["order_id"]."'");
        if ($query->num_rows > 0) {
            $query= $query->fetch_array();
            $output=array(
                "title" => "Payment status",
                "order-id" => $_GET["order_id"],
                "status" => $query["status"]
            );
                echo json_encode($output);
        }
        else //ไม่พบข้อมูลของ order
        {
            $output=array(
                "title" => "Payment status",
                "order-id" => $_GET["order_id"],
                "status" => 2
            );
                echo json_encode($output);
        }
    }
    else if (isset($_GET["method"]) && strcmp($_GET["method"],"set") == 0) //รับ api แบบ GET และ สร้างข้อมูลใหม่
    { 
        $conn = mysqli_connect("localhost","root","","credit_card") or die("Connect DB Error");
        $conn->query("SET NAMES UTF8");
        $conn->query("INSERT INTO payment(order_id,amount,status) VALUES ('".$_GET["order_id"]."','".$_GET["amount"]."',0)");
        echo "DONE";
    }
}
elseif (isset($_POST["order_id"])) {
    if (isset($_POST["method"]) && strcmp($_POST["method"],"get") == 0) //รับ api แบบ POST และ ส่งข้อมูลออก
    {
        $conn = mysqli_connect("localhost","root","","credit_card") or die("Connect DB Error");
        $conn->query("SET NAMES UTF8");
        $query = $conn->query("SELECT * FROM payment WHERE order_id='".$_POST["order_id"]."'");
        if ($query->num_rows > 0) {
            $query= $query->fetch_array();
            $output=array(
                "title" => "Payment status",
                "order-id" => $_POST["order_id"],
                "status" => $query["status"]
            );
                echo json_encode($output);
        }
        else //ไม่พบข้อมูลของ order
        {
            $output=array(
                "title" => "Payment status",
                "order-id" => $_POST["order_id"],
                "status" => 2
            );
                echo json_encode($output);
        }
    }
    else if (isset($_POST["method"]) && strcmp($_POST["method"],"set") == 0) //รับ api แบบ POST และ สร้างข้อมูลใหม่
    {
        $conn = mysqli_connect("localhost","root","","credit_card") or die("Connect DB Error");
        $conn->query("SET NAMES UTF8");
        $conn->query("INSERT INTO payment(order_id,amount,status) VALUES ('".$_POST["order_id"]."','".$_POST["amount"]."',0)");
    }
}
?>