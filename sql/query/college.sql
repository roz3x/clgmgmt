
-- name: InsertDepartment :one
insert into departments 
( 
   name , hod_id
) values(
  $1  , $2 
) returning *;


-- name: InsertCourse :one
insert into  courses (
  name , department_id 
) values (
  $1 , $2 
) returning *;


