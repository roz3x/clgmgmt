create table departments (
  id bigserial primary key,
  name text not null,
  hod_id int not null
);

create table instructors (
  id bigserial primary key,
  name text  not null
);

create table students (
  id bigserial primary key,
  name text not null 
);

create table courses (
  id bigserial primary key,
  name text not null , 
  department_id int not null ,
  instructor_id int  not null 
);

create table enroll (
  student_id int not null , 
  course_id int not null 
);