CREATE TABLE `book`
(
    `book` varchar(255) NOT NULL COMMENT 'book_name',
    `price` int NOT NULL COMMENT 'book_price',
    PRIMARY KEY (`book`)
)ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;