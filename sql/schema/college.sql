create table departments (
  id bigserial primary key,
  name text ,
  hod_id int
);

create table courses (
  name text  ,
  id bigserial primary key , 
  instructor_id integer,
  department_id int
);


create table instructors (
  name text , 
  id bigserial primary key , 
);


