create table survey (
    id int auto_increment,
    createdAt date not null,
    description varchar(255) not null,
    topic varchar(255) not null,
    status varchar(255) not null,
    resumo varchar(255) not null,
    conclusions varchar(255) not null,
    method varchar(255) not null,
    keywords varchar(255) not null,
    primary key (id) 
);
