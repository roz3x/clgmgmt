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
  email text not null , 
  age int not null , 
  id bigserial primary key not null 
);
create table students (
  name text not null , 
  age int not null  , 
  email text not null , 
  id bigserial primary key not null 
);

create table enroll (
  student_id int not null , 
  course_id int not null 
);

insert into students values
('Alfreds Futterkiste', 20 ,  'af@xyz.com') , 
('Blauer See Delikatessen' , 21 , 'bsd@xyz.com'),
('Blondel père et fils' , 22  , 'bpef@xyz.com'),
('Bottom-Dollar Marketse' , 21  , 'bdm@xyz.com'),
('Comércio Mineiro'  ,19  , 'cm@xyx.com'),
('Ernst Handel	' , 30 , 'eh@xyz.com' ),
('Antonio Moreno Taquería', 21 , 'amt@xyz.com') ;
