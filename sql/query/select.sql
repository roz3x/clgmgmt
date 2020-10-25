
-- name: SelectDepartments :many
select * from departments;


-- name: SeleteEnrollDetails :many
select students.name ,
enroll.course_id , 
courses.instructor_id 
from  
enroll left join students on enroll.student_id = students.id 
left join courses on courses.id = enroll.course_id ;
