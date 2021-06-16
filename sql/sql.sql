CREATE DATABASE IF NOT EXISTS Authentication;

USE Authentication;

DROP TABLE IF EXISTS usuarios;

CREATE TABLE usuarios(
    id int auto_increment primary key,
    nome varchar(50) not null,
    nick varchar(50) not null unique,
    email varchar(50) not null unique,
    senha varchar(100) not null unique,
    criadoem timestamp default current_timestamp()
) ENGINE=INNODB;