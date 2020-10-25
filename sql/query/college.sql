
-- name: InsertDepartment :one

insert into departments (
  name  , hod_id
) values  (
  $1 , $2
) returning *;


-- name: InsertInstructor :one
insert into instructors (
  name  
) values (
  $1
) returning *;


-- name: InsertStudent :one
insert into students (
  name
) values (
  $1
) returning *;

-- name: InsertCourse :one
insert into courses (
  name , department_id , instructor_id
) values (
  $1 , $2 , $3
) returning *;


-- name: InsertEnroll :one
insert into enroll (
  student_id , course_id
)  values (
  $1 , $2
) returning *;