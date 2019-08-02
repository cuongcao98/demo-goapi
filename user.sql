-- phpMyAdmin SQL Dump
-- version 4.9.0.1
-- https://www.phpmyadmin.net/
--
-- Host: localhost:8889
-- Generation Time: Aug 02, 2019 at 11:02 AM
-- Server version: 5.7.26
-- PHP Version: 7.3.7

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
SET time_zone = "+00:00";

--
-- Database: `user`
--

-- --------------------------------------------------------

--
-- Table structure for table `thongtins`
--

CREATE TABLE `thongtins` (
  `id` int(11) NOT NULL,
  `username` varchar(100) COLLATE utf8_unicode_ci NOT NULL,
  `password` varchar(100) COLLATE utf8_unicode_ci NOT NULL,
  `email` varchar(100) COLLATE utf8_unicode_ci NOT NULL,
  `birth` date NOT NULL,
  `sex` int(1) NOT NULL,
  `phone` int(20) NOT NULL,
  `national_id` int(11) NOT NULL,
  `height` float NOT NULL
) ENGINE=InnoDB DEFAULT CHARSET=utf8 COLLATE=utf8_unicode_ci;

--
-- Dumping data for table `thongtins`
--

INSERT INTO `thongtins` (`id`, `username`, `password`, `email`, `birth`, `sex`, `phone`, `national_id`, `height`) VALUES
(1, 'cuong', '123', 'cuong@gmail.com', '2018-10-10', 1, 9090909, 84, 1.7),
(3, 'bang', '1234', 'bang@gmail.com', '2002-10-02', 1, 987654, 34, 1.4),
(4, 'cuong', 'ewqq', 'haha@', '2019-08-13', 1, 909092, 22, 1.7),
(5, 'bang', '81dc9bdb52d04dc20036dbd8313ed055', 'bang@gmail.com', '2002-10-02', 1, 987654, 34, 1.4);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `thongtins`
--
ALTER TABLE `thongtins`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `thongtins`
--
ALTER TABLE `thongtins`
  MODIFY `id` int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=6;
