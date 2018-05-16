-- phpMyAdmin SQL Dump
-- version 4.5.1
-- http://www.phpmyadmin.net
--
-- Host: 127.0.0.1
-- Generation Time: Apr 22, 2018 at 01:59 AM
-- Server version: 10.1.19-MariaDB
-- PHP Version: 5.6.28

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `morrasoom_airline`
--

-- --------------------------------------------------------

--
-- Table structure for table `airline`
--

CREATE TABLE `airline` (
  `src` int(2) NOT NULL,
  `des` int(2) NOT NULL,
  `airLineID` int(5) NOT NULL,
  `gate` int(11) NOT NULL,
  `timetogo` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  `reach_des` timestamp NOT NULL DEFAULT '0000-00-00 00:00:00',
  `max_seat` int(11) NOT NULL,
  `reserved_seat` int(11) NOT NULL,
  `price` int(11) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf32 COLLATE=utf32_thai_520_w2;

--
-- Dumping data for table `airline`
--

INSERT INTO `airline` (`src`, `des`, `airLineID`, `gate`, `timetogo`, `reach_des`, `max_seat`, `reserved_seat`, `price`) VALUES
(11, 12, 1001112, 5, '2018-04-24 09:00:00', '2018-04-24 13:00:00', 50, 0, 1500),
(11, 13, 1001113, 8, '2018-04-24 09:30:00', '2018-04-24 13:30:00', 40, 0, 1300),
(11, 14, 1001114, 3, '2018-04-24 10:00:00', '2018-04-24 14:00:00', 45, 0, 2000),
(12, 11, 1001211, 4, '2018-04-24 09:00:00', '2018-04-24 13:00:00', 40, 0, 1450),
(12, 13, 1001213, 6, '2018-04-24 09:30:00', '2018-04-24 13:30:00', 30, 0, 1200),
(12, 14, 1001214, 7, '2018-04-24 10:00:00', '2018-04-24 14:00:00', 40, 0, 1300),
(13, 11, 1001311, 5, '2018-04-24 09:00:00', '2018-04-24 13:00:00', 40, 0, 1250),
(13, 12, 1001312, 1, '2018-04-24 09:30:00', '2018-04-24 13:30:00', 42, 0, 1400),
(13, 14, 1001314, 4, '2018-04-24 10:00:00', '2018-04-24 13:00:00', 38, 0, 1600),
(14, 11, 1001411, 4, '2018-04-24 09:00:00', '2018-04-24 13:00:00', 28, 0, 2000),
(14, 12, 1001412, 2, '2018-04-24 09:30:00', '2018-04-24 13:30:00', 38, 0, 1700),
(14, 13, 1001413, 6, '2018-04-24 10:00:00', '2018-04-24 14:00:00', 43, 0, 1650);

-- --------------------------------------------------------

--
-- Table structure for table `province`
--

CREATE TABLE `province` (
  `pid` int(2) NOT NULL,
  `pname` varchar(50) COLLATE utf32_thai_520_w2 NOT NULL,
  `pnameeng` varchar(50) COLLATE utf32_thai_520_w2 DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf32 COLLATE=utf32_thai_520_w2;

--
-- Dumping data for table `province`
--

INSERT INTO `province` (`pid`, `pname`, `pnameeng`) VALUES
(11, 'เชียงใหม่', 'chiang mai'),
(12, 'กรุงเทพ', 'bangkok'),
(13, 'ขอนแก่น', 'khon kaen'),
(14, 'ภูเก็ต', 'phuket');

-- --------------------------------------------------------

--
-- Table structure for table `reservation`
--

CREATE TABLE `reservation` (
  `pid` varchar(13) COLLATE utf32_thai_520_w2 NOT NULL,
  `airLineID` varchar(13) COLLATE utf32_thai_520_w2 NOT NULL,
  `name` text COLLATE utf32_thai_520_w2 NOT NULL,
  `tel` text COLLATE utf32_thai_520_w2 NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf32 COLLATE=utf32_thai_520_w2;

--
-- Indexes for dumped tables
--

--
-- Indexes for table `airline`
--
ALTER TABLE `airline`
  ADD PRIMARY KEY (`airLineID`);

--
-- Indexes for table `province`
--
ALTER TABLE `province`
  ADD PRIMARY KEY (`pid`);

--
-- Indexes for table `reservation`
--
ALTER TABLE `reservation`
  ADD PRIMARY KEY (`pid`,`airLineID`);

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
