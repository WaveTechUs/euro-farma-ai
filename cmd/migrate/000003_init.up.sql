create table user (
    id int auto_increment,
    name varchar(255) not null,
    email varchar(255) not null,
    password varchar(255) not null,
    role varchar(255) not null,
    primary key (id) 
);
