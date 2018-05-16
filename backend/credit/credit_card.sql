-- phpMyAdmin SQL Dump
-- version 4.7.7
-- https://www.phpmyadmin.net/
--
-- Host: 127.0.0.1
-- Generation Time: Apr 23, 2018 at 11:50 AM
-- Server version: 10.1.30-MariaDB
-- PHP Version: 7.2.2

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET AUTOCOMMIT = 0;
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `credit_card`
--
CREATE DATABASE IF NOT EXISTS `credit_card` DEFAULT CHARACTER SET utf8 COLLATE utf8_general_ci;
USE `credit_card`;

-- --------------------------------------------------------

--
-- Table structure for table `member`
--

CREATE TABLE `member` (
  `id_card` varchar(16) NOT NULL,
  `name` varchar(30) NOT NULL,
  `exp_date` varchar(10) NOT NULL,
  `cvv` varchar(3) NOT NULL,
  `money` int(10) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data for table `member`
--

INSERT INTO `member` (`id_card`, `name`, `exp_date`, `cvv`, `money`) VALUES
('1345887945621235', 'Jaturapong khamngoen', '06/22', '456', 100000),
('4568135445687895', 'Sirapop sitthi', '03/19', '759', 45000),
('5468467815464897', 'Supakit chumjanjira', '06/21', '753', 45000),
('5645231256485465', 'Thunyawan sithi', '05/20', '125', 50000),
('7456123456478956', 'Sahaphap panya', '05/12', '654', 60000),
('9562123545781254', 'Suthipong Kaowdang', '04/20', '546', 45000);

-- --------------------------------------------------------

--
-- Table structure for table `payment`
--

CREATE TABLE `payment` (
  `order_id` varchar(10) NOT NULL,
  `amount` int(10) NOT NULL,
  `status` tinyint(1) NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

--
-- Dumping data for table `payment`
--

INSERT INTO `payment` (`order_id`, `amount`, `status`) VALUES
('A551CD1447', 2000, 0),
('A551CD1449', 5000, 1);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `member`
--
ALTER TABLE `member`
  ADD PRIMARY KEY (`id_card`);

--
-- Indexes for table `payment`
--
ALTER TABLE `payment`
  ADD PRIMARY KEY (`order_id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
