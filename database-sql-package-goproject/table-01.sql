DROP TABLE IF EXISTS album;
CREATE TABLE album (
                       id         INT AUTO_INCREMENT NOT NULL,
                       title      VARCHAR(128) NOT NULL,
                       artist     VARCHAR(255) NOT NULL,
                       price      DECIMAL(5,2) NOT NULL,
                       quantity   INT UNSIGNED,
                       PRIMARY KEY (`id`)
);

INSERT INTO album
(title, artist, price, quantity)
VALUES
    ('Blue Train', 'John Coltrane', 56.99, 5),
    ('Giant Steps', 'John Coltrane', 63.99, 62),
    ('Jeru', 'Gerry Mulligan', 17.99, 0),
    ('Sarah Vaughan', 'Sarah Vaughan', 34.98, 127);
