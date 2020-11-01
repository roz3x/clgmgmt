create table departments (
  id bigserial primary key not null  ,
  name text not null ,
  hod_id int not null 
);

create table courses (
  name text not null  ,
  id bigserial primary key not null , 
  instructor_id integer not null ,
  department_id int not null 
);

create table instructors (
  name text not null , 
  id bigserial primary key not null 
);
create table students (
  name text not null , 
  id bigserial primary key not null 
)