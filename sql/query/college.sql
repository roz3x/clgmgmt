
-- name: SelectStudents :many 
select * from students;

-- name: SelectEnroll :many 
select * from enroll;

-- name: SelectDepartment :many 
select * from departments;

-- name: SelectInstructor :many 
select * from instructors;

-- name: SelectCourses :many 
select * from courses;

-- name: InsertDepartment :one
insert into departments 
( 
   name , hod_id
) values(
  $1  , $2 
) returning *;

-- name: InsertStudent :one
insert into students (
  name , email , age 
) values (
  $1 , $2 , $3
) returning *;

-- name: InsertInstructor :one
insert into instructors (
 name , email , age
) values (
  $1 , $2 , $3
) returning *;


-- name: InsertCourse :one 
insert into courses (
  name ,  instructor_id , department_id 
) values (
  $1 , $2 , $3
) returning *;

-- name: InsertEnroll :one
insert into enroll (
 student_id  ,  course_id 
) values (
  $1 , $2
) returning * ;

