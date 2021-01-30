CREATE TABLE user
(
    id            INT PRIMARY KEY AUTO_INCREMENT,
    username      VARCHAR(50) UNIQUE,
    password_hash VARCHAR(255),
    created_at    DATETIME DEFAULT NOW()
);

CREATE TABLE profile
(
    user_id   INT PRIMARY KEY,
    name      VARCHAR(50),
    surname   VARCHAR(50),
    age       INT,
    gender    VARCHAR(50),
    interests TEXT,
    city      VARCHAR(100),
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
