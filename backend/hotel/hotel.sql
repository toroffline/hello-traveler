-- phpMyAdmin SQL Dump
-- version 4.6.5.2
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Apr 29, 2018 at 12:05 PM
-- Server version: 10.1.21-MariaDB
-- PHP Version: 5.6.30

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `hotel`
--

-- --------------------------------------------------------

--
-- Table structure for table `booking`
--

CREATE TABLE `booking` (
  `booking_id` int(11) NOT NULL,
  `hotel_id` int(11) NOT NULL,
  `check_in` date NOT NULL,
  `check_out` date NOT NULL,
  `qty` int(11) NOT NULL,
  `total_price` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data for table `booking`
--

INSERT INTO `booking` (`booking_id`, `hotel_id`, `check_in`, `check_out`, `qty`, `total_price`) VALUES
(1, 2, '2018-04-22', '2018-04-24', 2, 2),
(2, 2, '2018-04-24', '2018-04-27', 1, 3),
(3, 2, '2018-04-22', '2018-04-25', 3, 4),
(4, 3, '2018-04-22', '2018-04-26', 4, 5),
(5, 3, '2018-04-22', '2018-04-26', 2, 6),
(6, 4, '2018-04-25', '2018-04-27', 5, 7),
(7, 5, '2018-04-23', '2018-04-28', 2, 8),
(8, 6, '2018-04-21', '2018-04-22', 5, 9),
(9, 6, '2018-04-23', '2018-04-25', 1, 10),
(10, 7, '2018-04-21', '2018-04-24', 4, 11),
(11, 8, '2018-04-24', '2018-04-26', 2, 12),
(12, 8, '2018-04-25', '2018-04-26', 2, 13),
(13, 9, '2018-04-25', '2018-04-28', 4, 14),
(14, 9, '2018-04-26', '2018-04-27', 3, 15),
(15, 10, '2018-04-22', '2018-04-23', 2, 16),
(16, 10, '2018-04-24', '2018-04-26', 3, 17),
(17, 13, '2018-04-25', '2018-04-26', 4, 18),
(18, 13, '2018-04-26', '2018-04-29', 2, 19),
(19, 13, '2018-04-26', '2018-04-30', 2, 20);

-- --------------------------------------------------------

--
-- Table structure for table `hotel`
--

CREATE TABLE `hotel` (
  `hotel_id` int(11) NOT NULL,
  `name` varchar(40) NOT NULL,
  `address` varchar(100) NOT NULL,
  `province` varchar(40) NOT NULL,
  `telephone` varchar(11) NOT NULL,
  `room_qty` int(11) NOT NULL,
  `price` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data for table `hotel`
--

INSERT INTO `hotel` (`hotel_id`, `name`, `address`, `province`, `telephone`, `room_qty`, `price`) VALUES
(1, 'โรงแรมดวงตะวันเชียงใหม่', '132 ถนน ลอยเคราะห์ ตำบล ช้างคลาน Chang Mai เชียงใหม่ 50100', 'เชียงใหม่', '053-905000', 5, 1464),
(2, 'โมนา เชียงใหม่ บูติค', 'Chang Klan, Mueang, เชียงใหม่ 50100', 'เชียงใหม่', '095-7895196', 6, 1280),
(3, 'โรงแรมเชียงใหม่เกท', '11/10 ถนนสุริยวงศ์ ตำบล หายยา อำเภอเมืองเชียงใหม่ เชียงใหม่ 50100', 'เชียงใหม่', '053-203895', 6, 1040),
(4, 'โรงแรมอิมพีเรียลแม่ปิง', '153 Sridonchai Road, Chang Klan, Chiang Mai 50100', 'เชียงใหม่', '053-283900', 6, 1420),
(5, 'เลอ เมอริเดียน เชียงใหม่', '108 Chang Klan Road Tambol Chang Klan, Amphur Muang Tambol Chang Klan Chiang Mai TH-50 50100', 'เชียงใหม่', '053-253666', 5, 3506),
(6, 'มณีนาราคร', '99 ถนน ศรีดอนไชย ตำบลช้างคลาน อำเภอเมืองเชียงใหม่ เชียงใหม่ 50100', 'เชียงใหม่', '053-999555', 5, 865),
(7, 'Holiday Inn ChiangMai', '318/1 ถนนเชียงใหม่-ลำพูน, อำเภอเมืองเชียงใหม่ เชียงใหม่ 50000', 'เชียงใหม่', '053-275300', 7, 1489),
(8, 'The Empress Chiang Mai Hotel', '199/42 Chang Klan Rd, ตำบล ช้างคลาน อำเภอเมืองเชียงใหม่ เชียงใหม่ 50100', 'เชียงใหม่', '053-253199', 8, 2350),
(9, 'โรงแรมโลตัสปางสวนแก้ว', '21 ถนน ห้วยแก้ว ต.สุเทพ อำเภอเมืองเชียงใหม่ เชียงใหม่ 50200', 'เชียงใหม่', '053-224333', 7, 1686),
(10, 'Ibis Phuket Patong', '10 Chalermphrakiat Road, Kathu District, Phuket, 83150', 'ภูเก็ต', '02-6591066', 7, 1500),
(11, 'คาลิมา รีสอร์ท แอนด์ สปา', '338/1 ถนน พระบารมี ตำบลป่าตอง อำเภอ กะทู้ ภูเก็ต 83150', 'ภูเก็ต', '076-358999', 7, 3985),
(12, 'Angsana Laguna Phuket', '10 Moo 4, Srisoonthorn Road, Cherngtalay, อำเภอ ถลาง Phuket 83110', 'ภูเก็ต', '076-358500', 7, 4022),
(13, 'Sleep With Me Hotel design hotel @ paton', 'Sleep With Me Hotel design hotel @ patong', 'ภูเก็ต', '076-363333', 7, 1765),
(14, 'Burasari Phuket', '18/110 Ruamjai Road ตำบลป่าตอง Amphur Kathu Phuket ภูเก็ต 83150', 'ภูเก็ต', '076-292929', 7, 1350),
(15, 'ปริภัส ป่าตอง รีสอร์ท', '230 ถนน ราษฏร์อุทิศ 200 ปี ตำบลป่าตอง อำเภอ กะทู้ ภูเก็ต 83150', 'ภูเก็ต', '076-349840', 7, 1350),
(16, 'โรงแรมอันดาคีรา', '222,222/1,222/2 Phangmuang Sai Ko Road, ตำบลป่าตอง อำเภอ กะทู้ ภูเก็ต 83150', 'ภูเก็ต', '076-366188', 7, 2300),
(17, 'เดอะชาร์ม รีสอร์ท ภูเก็ต', '212 ถนนทวีวงศ์ ตำบลป่าตอง อำเภอ กะทู้ ภูเก็ต 83150', 'ภูเก็ต', '076-318300', 6, 2560),
(18, 'Bauman Residence Hotel', '201 ถ.ผังเมือง สาย ก กะทู้ ภูเก็ต 83150', 'ภูเก็ต', '076-310999', 9, 1200),
(19, 'Deevana Patong Resort & Spa', '43/2 Raj-U-Thid, ถนน 200 ปี ตำบลป่าตอง อำเภอ กะทู้ ภูเก็ต 83150', 'ภูเก็ต', '076-317179', 4, 2185),
(20, 'โรงแรม แรมแบรนดท์ กรุงเทพฯ', '19 ซอย สุขุมวิท 18 แขวง คลองเตย เขต คลองเตย กรุงเทพมหานคร 10110', 'กรุงเทพ', '02-2617100', 9, 1985),
(21, 'Ibis Hotel Bangkok Siam', '927 ถนน พระรามที่ 1 แขวง วังใหม่ เขต ปทุมวัน กรุงเทพมหานคร 10330', 'กรุงเทพ', '02-6592888', 4, 1700),
(22, 'Rambuttri Village Inn & Plaza', '95 Soi Ram Buttri, Chakkra Phong Road, Phra Nakorn, กรุงเทพมหานคร 10200', 'กรุงเทพ', '02-2829162', 1, 950),
(23, 'Lemontea Hotel', '55 เพชรบุรี 15 แขวง ราชเทวี เขต ราชเทวี กรุงเทพมหานคร 10400', 'กรุงเทพ', '02-6935499', 0, 1450),
(24, 'ดีแอนด์ดีอินน์', '68-70 ถนน ข้าวสาร แขวง ตลาดยอด เขต พระนคร กรุงเทพมหานคร 10200', 'กรุงเทพ', '02-6290526', 4, 965),
(25, 'Bangkok Palace Hotel', '1091/336 Soi New Phetchaburi, Makkasan, Ratchathewi, Bangkok 10400', 'กรุงเทพ', '02-8909000', 0, 1625),
(26, 'Khaosan Palace Hotel', '139 ถนนข้าวสาร แขวง ตลาดยอด เขต พระนคร กรุงเทพมหานคร 10200', 'กรุงเทพ', '02-2820578', 5, 3950),
(27, 'ริกก้า อินน์', 'ถนนข้าวสาร แขวง ตลาดยอด เขต พระนคร กรุงเทพมหานคร 10200', 'กรุงเทพ', '02-2827511', 7, 1050),
(28, 'โรงแรมเอเชีย กรุงเทพ', '296 ถนน พญาไท แขวง ถนนเพชรบุรี เขต ราชเทวี กรุงเทพมหานคร 10400', 'กรุงเทพ', '02-2170808', 5, 2500),
(29, 'โรงแรม นูโวซิตี้', '889 Rama I Rd, แขวง ปทุมวัน เขต ปทุมวัน กรุงเทพมหานคร 10330', 'กรุงเทพ', '02-2827500', 6, 2149),
(30, 'Pullman Khon Kaen Raja Orchid', '9-9 Prachasumran Road, Nai Muang, Muang, 40000', 'ขอนแก่น', '043-913333', 4, 2925),
(31, 'โรงแรม กลาเซียร์', 'Muang อำเภอเมืองขอนแก่น ขอนแก่น 40000', 'ขอนแก่น', '043-334999', 4, 1500),
(32, 'Chada Verandah Hotel', 'ซอย ศรีจันทร์ 3 ตำบล ในเมือง อำเภอเมืองขอนแก่น ขอนแก่น 40000', 'ขอนแก่น', '088-0356823', 4, 669),
(33, 'โรงแรม OMG Place', '53/1 Soi Pimpasut 1 ในเมือง เมือง 40000', 'ขอนแก่น', '085-6813131', 4, 699),
(34, 'โรงแรมฮ็อป อินน์ ขอนแก่น HOP INN Khon Ka', '90/609 หมู่ที่ 4 ตำบล ในเมือง 90/609 Moo 4, Nai Mueang Sub District, Mueang Khon Kaen District, Khon', 'ขอนแก่น', '02-6592899', 4, 650),
(35, 'โรมาโฮเต็ล', '50/2 กลางเมือง ซอย กลางเมือง 2 ตำบลในเมือง อำเภอเมืองขอนแก่น ขอนแก่น 40000', 'ขอนแก่น', '043-334444', 4, 999),
(36, 'โรงแรม แมมมอธ รีสอร์ท ขอนแก่น', '406/1 ซอยหนองวัด 4 ในเมือง เมือง ขอนแก่น 40000', 'ขอนแก่น', '043-226246', 5, 1250),
(37, 'โรงแรมเอสเอฟบิซ', '98 ซอย ดรุณสำราญ ตำบล ในเมือง อำเภอเมืองขอนแก่น ขอนแก่น 40000', 'ขอนแก่น', '043-226500', 8, 2569),
(38, 'โรงแรมขอนแก่นออร์คิด', '146/96 Soi Way Pachaiyawan Road, ขอนแก่น 40000', 'ขอนแก่น', '043-246068', 9, 750),
(39, 'Piman Garden Boutique Hotel', 'ตำบล ในเมือง อำเภอเมืองขอนแก่น ขอนแก่น 40000', 'ขอนแก่น', '043-334111', 9, 1250);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `booking`
--
ALTER TABLE `booking`
  ADD PRIMARY KEY (`booking_id`),
  ADD KEY `book_hotel` (`hotel_id`);

--
-- Indexes for table `hotel`
--
ALTER TABLE `hotel`
  ADD PRIMARY KEY (`hotel_id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `hotel`
--
ALTER TABLE `hotel`
  MODIFY `hotel_id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=40;
--
-- Constraints for dumped tables
--

--
-- Constraints for table `booking`
--
ALTER TABLE `booking`
  ADD CONSTRAINT `book_hotel` FOREIGN KEY (`hotel_id`) REFERENCES `hotel` (`hotel_id`) ON DELETE CASCADE ON UPDATE CASCADE;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
