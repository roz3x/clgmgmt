create table departments (
  id bigserial primary key,
  _name text ,
  hod_id int
);

create table instructors (
  id bigserial primary key,
  _name text 
);

create table students (
  id bigserial primary key,
  _name text 
);

create table courses (
  id bigserial primary key,
  _name text , 
  department_id int ,
  instructor_id int
);

create table enroll (
  student_id int , 
  course_id int 
);