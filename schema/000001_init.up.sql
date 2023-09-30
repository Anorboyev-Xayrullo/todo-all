CREATE TABLE book
 (
    id Bigserial Not Null,
    name Varchar (255) Not Null,
    author Varchar (255) Not Null,
    type bool,
    number int,
    description Varchar (255) Not Null,
    language Varchar (255) Not Null,
    year Varchar(100) Not Null
);

CREATE TABLE genres 
(
    id Bigserial Not Null,
    genre Varchar(20) Not Null
);