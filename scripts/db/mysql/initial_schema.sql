CREATE DATABASE IF NOT EXISTS `silverlabs` DEFAULT CHARACTER SET latin1;

DROP TABLE IF EXISTS `Account`;
CREATE TABLE `Account` (
    `AccountID` int NOT NULL AUTO_INCREMENT,
    `DocumentNumber` varchar(20),
    PRIMARY KEY (`AccountID`)
);

DROP TABLE IF EXISTS `Transaction`;
CREATE TABLE `Transaction` (
    `TransactionID` int NOT NULL AUTO_INCREMENT,
    `AccountID` int,
    `OperationTypeID` int,
    `Amount` decimal(15,2),
    `EventDate` datetime(5),
    PRIMARY KEY (`TransactionID`)
);