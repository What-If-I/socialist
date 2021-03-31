CREATE TABLE user
(
    id            INT PRIMARY KEY AUTO_INCREMENT,
    username      VARCHAR(50) UNIQUE NOT NULL,
    password_hash VARCHAR(255)       NOT NULL,
    created_at    datetime DEFAULT now()
);

CREATE TABLE profile
(
    user_id   INT PRIMARY KEY,
    name      VARCHAR(50)  NOT NULL,
    surname   VARCHAR(50)  NOT NULL,
    age       INT          NOT NULL,
    gender    VARCHAR(50)  NOT NULL,
    interests text         NOT NULL,
    city      VARCHAR(100) NOT NULL,
    FOREIGN KEY (user_id) REFERENCES user (id)
        ON UPDATE CASCADE ON DELETE RESTRICT
);

CREATE TABLE friends
(
    user_id   INT,
    friend_id INT,
    FOREIGN KEY (user_id) REFERENCES user (id)
        ON UPDATE CASCADE ON DELETE NO ACTION,
    FOREIGN KEY (friend_id) REFERENCES user (id)
        ON UPDATE CASCADE ON DELETE NO ACTION
);
