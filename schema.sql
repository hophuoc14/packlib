create table users (
    id varchar(50) primary key,
    username varchar(100) not null unique,
    password varchar(100) not null,
    email varchar(100) not null unique,
    employee_id varchar(50) unique,
    created_at timestamp not null default current_timestamp,
    updated_at timestamp not null default current_timestamp on update current_timestamp,
    deleted_at timestamp null,
    foreign key (employee_id) references employees(id) on update cascade on delete set null
)