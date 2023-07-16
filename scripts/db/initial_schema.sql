CREATE DATABASE IF NOT EXISTS `silverlabs` DEFAULT CHARACTER SET latin1;

DROP TABLE IF EXISTS `Account`;
CREATE TABLE `Account` (
    `AccountID` int NOT NULL AUTO_INCREMENT,
    `DocumentNumber` varchar(20),
    PRIMARY KEY (`AccountID`)
)