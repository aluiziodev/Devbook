create database DevBook;
use DevBook;

drop table if exists usuarios;

create table usuarios(
	id int auto_increment primary key,
    nome varchar(50) not null,
    nick varchar(50) not null unique,
    email varchar(50) not null unique,
    senha varchar(50) not null unique,
    data_inicio timestamp default current_timestamp()
) engine=InnoDB;



