use test;

CREATE TABLE users (
    id int(auto_increment) primary key,
    name varchar(100) not null,
    password varchar(100) not null,
)ENGINE=INNODB;

INSERT INTO test.users (name,password) VALUES
    ('haposan','haposan123');
