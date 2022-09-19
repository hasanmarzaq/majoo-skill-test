-- phpMyAdmin SQL Dump
-- version 5.2.0
-- https://www.phpmyadmin.net/
--
-- Host: localhost
-- Generation Time: Sep 19, 2022 at 02:35 AM
-- Server version: 5.7.33
-- PHP Version: 7.4.19

SET SQL_MODE = "NO_AUTO_VALUE_ON_ZERO";
START TRANSACTION;
SET time_zone = "+00:00";


/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET @OLD_CHARACTER_SET_RESULTS=@@CHARACTER_SET_RESULTS */;
/*!40101 SET @OLD_COLLATION_CONNECTION=@@COLLATION_CONNECTION */;
/*!40101 SET NAMES utf8mb4 */;

--
-- Database: `test-majoo`
--

-- --------------------------------------------------------

--
-- Table structure for table `areas`
--

CREATE TABLE `areas` (
  `id` bigint(20) NOT NULL,
  `area_value` bigint(20) DEFAULT NULL,
  `type` longtext
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `areas`
--

INSERT INTO `areas` (`id`, `area_value`, `type`) VALUES
(1, 100, 'persegi');

-- --------------------------------------------------------

--
-- Table structure for table `merchants`
--

CREATE TABLE `merchants` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `user_id` bigint(20) UNSIGNED NOT NULL,
  `merchant_name` varchar(40) NOT NULL,
  `uuid` varchar(50) NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `created_by` bigint(20) UNSIGNED DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `updated_by` bigint(20) UNSIGNED DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `deleted_by` bigint(20) UNSIGNED DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `merchants`
--

INSERT INTO `merchants` (`id`, `user_id`, `merchant_name`, `uuid`, `created_at`, `created_by`, `updated_at`, `updated_by`, `deleted_at`, `deleted_by`) VALUES
(1, 1, 'merchant 1', '10e89f9d-9ab9-4f8d-b8ee-80ab59b85497', '2022-09-19 00:30:10.991', 1, '2022-09-19 00:30:10.991', 0, NULL, 0),
(2, 2, 'merchant 2', '1dbee1e4-6316-4924-94d6-416c655672ac', '2022-09-19 00:30:28.869', 2, '2022-09-19 00:30:28.869', 0, NULL, 0),
(3, 3, 'merchant 3', 'ade78ebf-f8a8-4c9c-aa84-dc36b7319d99', '2022-09-19 00:30:43.291', 3, '2022-09-19 00:30:43.291', 0, NULL, 0);

-- --------------------------------------------------------

--
-- Table structure for table `outlets`
--

CREATE TABLE `outlets` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `merchant_id` bigint(20) UNSIGNED NOT NULL,
  `outlet_name` varchar(40) NOT NULL,
  `uuid` varchar(50) NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `created_by` bigint(20) UNSIGNED DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `updated_by` bigint(20) UNSIGNED DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `deleted_by` bigint(20) UNSIGNED DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `outlets`
--

INSERT INTO `outlets` (`id`, `merchant_id`, `outlet_name`, `uuid`, `created_at`, `created_by`, `updated_at`, `updated_by`, `deleted_at`, `deleted_by`) VALUES
(1, 1, 'outlet 1', '3780b58a-6033-4d95-b2d5-ca361aec8c18', '2022-09-19 01:19:42.168', 3, '2022-09-19 01:19:42.168', 0, NULL, 0),
(2, 1, 'outlet 2', '9dcf590f-a7b2-4042-9bae-b3b07ad143d4', '2022-09-19 01:20:14.045', 3, '2022-09-19 01:20:14.045', 0, NULL, 0),
(3, 1, 'outlet 3', 'df685608-4fc8-4e8c-99f3-a915efde8a9f', '2022-09-19 01:20:22.625', 3, '2022-09-19 01:20:22.625', 0, NULL, 0),
(4, 2, 'outlet 4', '10c5b111-a5d8-4b2b-b467-2f7001adb678', '2022-09-19 01:20:32.486', 3, '2022-09-19 01:20:32.486', 0, NULL, 0),
(5, 2, 'outlet 5', '7aee2955-f33c-48b8-a441-c88273ba0651', '2022-09-19 01:20:36.537', 3, '2022-09-19 01:20:36.537', 0, NULL, 0),
(6, 2, 'outlet 6', '6964ac62-664c-43fc-b720-380ef19d55b9', '2022-09-19 01:20:39.704', 3, '2022-09-19 01:20:39.704', 0, NULL, 0),
(7, 3, 'outlet 7', '7ad32348-9329-4634-aea6-e4ceae04008b', '2022-09-19 01:20:47.320', 3, '2022-09-19 01:20:47.320', 0, NULL, 0),
(8, 3, 'outlet 8', '91a4544c-7751-4300-a200-bb1e585d1d14', '2022-09-19 01:20:53.316', 3, '2022-09-19 01:20:53.316', 0, NULL, 0),
(9, 3, 'outlet 9', 'd72f0045-7418-48c6-b60d-2ec61d41cf94', '2022-09-19 01:30:17.777', 3, '2022-09-19 01:30:17.777', 0, NULL, 0),
(10, 3, 'outlet 10', '45d9a1ce-9ae4-4e27-a3f9-33af0505c971', '2022-09-19 08:57:07.677', 3, '2022-09-19 08:57:07.677', 0, NULL, 0);

-- --------------------------------------------------------

--
-- Table structure for table `transactions`
--

CREATE TABLE `transactions` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `merchant_id` bigint(20) UNSIGNED NOT NULL,
  `outlet_id` bigint(20) UNSIGNED NOT NULL,
  `bill_total` double NOT NULL,
  `uuid` varchar(50) NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `created_by` bigint(20) UNSIGNED DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `updated_by` bigint(20) UNSIGNED DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `deleted_by` bigint(20) UNSIGNED DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `transactions`
--

INSERT INTO `transactions` (`id`, `merchant_id`, `outlet_id`, `bill_total`, `uuid`, `created_at`, `created_by`, `updated_at`, `updated_by`, `deleted_at`, `deleted_by`) VALUES
(1, 3, 7, 20, '730213a6-1eef-4581-a646-ee134db4716d', '2022-09-19 02:47:29.319', 3, '2022-09-19 02:47:29.319', 0, NULL, 0),
(2, 3, 7, 20000, '3885188c-8a0f-47ac-9c6e-e11555885068', '2022-09-19 02:48:22.774', 3, '2022-09-19 02:48:22.774', 0, NULL, 0);

-- --------------------------------------------------------

--
-- Table structure for table `users`
--

CREATE TABLE `users` (
  `id` bigint(20) UNSIGNED NOT NULL,
  `name` varchar(45) DEFAULT NULL,
  `user_name` varchar(45) DEFAULT NULL,
  `password` varchar(255) DEFAULT NULL,
  `uuid` varchar(50) NOT NULL,
  `created_at` datetime(3) DEFAULT NULL,
  `created_by` bigint(20) UNSIGNED DEFAULT NULL,
  `updated_at` datetime(3) DEFAULT NULL,
  `updated_by` bigint(20) UNSIGNED DEFAULT NULL,
  `deleted_at` datetime(3) DEFAULT NULL,
  `deleted_by` bigint(20) UNSIGNED DEFAULT NULL
) ENGINE=InnoDB DEFAULT CHARSET=latin1;

--
-- Dumping data for table `users`
--

INSERT INTO `users` (`id`, `name`, `user_name`, `password`, `uuid`, `created_at`, `created_by`, `updated_at`, `updated_by`, `deleted_at`, `deleted_by`) VALUES
(1, 'pengguna 1', 'pengguna1', '$2a$04$KvCNowLz6rdKFxo1rA/YYuhfy3V.S3E9GFhF2VcYSzg4AMNfXVU6i', '65845bd8-0b51-47de-9850-1b0e6f6a0bd9', '2022-09-19 00:08:52.204', 0, '2022-09-19 00:08:52.204', 0, NULL, 0),
(2, 'pengguna 2', 'pengguna2', '$2a$04$4k31FyVyIpV71UbOY5hE2.FBkKamPCE0vlBk8ByyU8ORiXXi9.5/i', 'e4b63c94-f1e7-40d9-8244-d19fc8a66cde', '2022-09-19 00:09:01.457', 0, '2022-09-19 00:09:01.457', 0, NULL, 0),
(3, 'pengguna 3', 'pengguna3', '$2a$04$IprrU0fMbuqfEv3i1YV3DuAPFSM6HO4M9zw4RbMqK6ZrXwEaoLzay', '64b73ab3-f029-4d5f-ac63-e9bd2198f02b', '2022-09-19 00:09:10.027', 0, '2022-09-19 00:09:10.027', 0, NULL, 0);

--
-- Indexes for dumped tables
--

--
-- Indexes for table `areas`
--
ALTER TABLE `areas`
  ADD PRIMARY KEY (`id`);

--
-- Indexes for table `merchants`
--
ALTER TABLE `merchants`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_merchants_user` (`user_id`);

--
-- Indexes for table `outlets`
--
ALTER TABLE `outlets`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_outlets_merchant` (`merchant_id`);

--
-- Indexes for table `transactions`
--
ALTER TABLE `transactions`
  ADD PRIMARY KEY (`id`),
  ADD KEY `fk_transactions_merchant` (`merchant_id`),
  ADD KEY `fk_transactions_outlet` (`outlet_id`);

--
-- Indexes for table `users`
--
ALTER TABLE `users`
  ADD PRIMARY KEY (`id`);

--
-- AUTO_INCREMENT for dumped tables
--

--
-- AUTO_INCREMENT for table `areas`
--
ALTER TABLE `areas`
  MODIFY `id` bigint(20) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;

--
-- AUTO_INCREMENT for table `merchants`
--
ALTER TABLE `merchants`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- AUTO_INCREMENT for table `outlets`
--
ALTER TABLE `outlets`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=11;

--
-- AUTO_INCREMENT for table `transactions`
--
ALTER TABLE `transactions`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=3;

--
-- AUTO_INCREMENT for table `users`
--
ALTER TABLE `users`
  MODIFY `id` bigint(20) UNSIGNED NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=4;

--
-- Constraints for dumped tables
--

--
-- Constraints for table `merchants`
--
ALTER TABLE `merchants`
  ADD CONSTRAINT `fk_merchants_user` FOREIGN KEY (`user_id`) REFERENCES `users` (`id`);

--
-- Constraints for table `outlets`
--
ALTER TABLE `outlets`
  ADD CONSTRAINT `fk_outlets_merchant` FOREIGN KEY (`merchant_id`) REFERENCES `merchants` (`id`);

--
-- Constraints for table `transactions`
--
ALTER TABLE `transactions`
  ADD CONSTRAINT `fk_transactions_merchant` FOREIGN KEY (`merchant_id`) REFERENCES `merchants` (`id`),
  ADD CONSTRAINT `fk_transactions_outlet` FOREIGN KEY (`outlet_id`) REFERENCES `outlets` (`id`);
COMMIT;

/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
/*!40101 SET CHARACTER_SET_RESULTS=@OLD_CHARACTER_SET_RESULTS */;
/*!40101 SET COLLATION_CONNECTION=@OLD_COLLATION_CONNECTION */;
