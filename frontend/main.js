
const hostUrl = "http://localhost:8080";

window.onload = () => {
  console.log("startedd......")
  document.getElementById("students").onclick = showStudents;
  document.getElementById("instructors").onclick = showInstructors;
  document.getElementById("departments").onclick = showDepatments;
  document.getElementById("courses").onclick = showCourses;
}

function showStudents() {

  console.log("showing students table");

  fetch(hostUrl+"/students")
  .then(res =>  {
    console.log(res);
  })
}

function showDepatments() {
  console.log("showing departments table");
}

function showCourses() {
  console.log("showing courses table");
}

function showInstructors() {
  console.log("showing instructor table");
}