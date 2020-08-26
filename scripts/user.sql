drop table users;

create table users (
    id varchar(50) primary key ,
    username varchar(100) unique,
    password varchar(100)
)

select *
from users;

delete from users where username = "thuocnv";
