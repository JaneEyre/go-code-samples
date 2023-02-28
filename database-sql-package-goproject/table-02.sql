DROP TABLE IF EXISTS album_trx;
CREATE TABLE album_trx (
                             trx_id    INT AUTO_INCREMENT NOT NULL,
                             trx_check     INT UNSIGNED,
                             PRIMARY KEY (`trx_id`)
);
