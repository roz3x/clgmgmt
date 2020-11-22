
const hostUrl = window.location.href;
const box  = document.getElementById("box");
const input = document.getElementById("input");
window.onload = () => {
  console.log("startedd......")
  document.getElementById("students").onclick = showStudents;
  document.getElementById("instructors").onclick = showInstructors;
  document.getElementById("departments").onclick = showDepatments;
  document.getElementById("courses").onclick = showCourses;
  document.getElementById("enroll").onclick = showEnroll;
}

function showStudents() {
  fetch(hostUrl+"/students", {
    method: "GET",
  })
  .then(e => e.json())
  .then(e => {
    var content = `
<table class="thetable">
  <tr style="color:#aaa;font-size:1.3em;">
    <th>Name</th>
    <th>Email</th>
    <th>Age</th>
  </tr>
`;
    for ( var i of e) {
      content += `<tr>`;
      content += `<th>${i.name}</th>`;
      content += `<th>${i.email}</th>`;
      content += `<th>${i.age}</th>`;
      content += `</tr>`;
    }
    content += `</table>`;
    box.innerHTML = content;
  })
  var InputBoxContent = `
  <div id="d">
  student's name<input id="name" type="text">
  email<input id="email" type="text">
  age<input type="number" id="age"> 
  <button id="btn">submit</button>
  </div>
  `
  input.innerHTML  = InputBoxContent;
  document.getElementById("btn").onclick = createStudent;
}

function showDepatments() {
  box.innerHTML = "";
  fetch(hostUrl+"/departments", {
    method: "GET",
  })
  .then(e => e.json())
  .then(e => {
    var content = `
<table class="thetable">
  <tr style="color:#aaa;font-size:1.3em;">
    <th>Name</th>
  </tr>
`;
    for ( var i of e) {
      content += `<tr>`;
      content += `<th>${i.name}</th>`;
      content += `</tr>`;
    }
    content += `</table>`;
    box.innerHTML = content;
  })
  var InputBoxContent = `
  <div id="d">
  department's name<input id="name" type="text">
  <button id="btn">submit</button>
  </div>
  `
  input.innerHTML  = InputBoxContent;
  document.getElementById("btn").onclick = createDepartment;

}


function showCourses() {
  box.innerHTML = "";
  fetch(hostUrl+"/courses", {
    method: "GET",
  })
  .then(e => e.json())
  .then(e => {
    var content = `
<table class="thetable">
  <tr style="color:#aaa;font-size:1.3em;">
    <th>Instructor ID</th>
    <th>Department ID</th>
  </tr>
`;
    for ( var i of e) {
      content += `<tr>`;
      content += `<th>${i.instructor_id}</th>`;
      content += `<th>${i.department_id}</th>`;
      content += `</tr>`;
    }
    content += `</table>`;
    box.innerHTML = content;
  })
  var InputBoxContent = `
  <div id="d">
  department's id<input id="d_id" type="number">
  instructor's id<input id="i_id" type="number">
  <button id="btn">submit</button>
  </div>
  `
  input.innerHTML  = InputBoxContent;
  document.getElementById("btn").onclick = createCourse;

}

function showInstructors() {
  box.innerHTML = "";
  fetch(hostUrl+"/instructors", {
    method: "GET",
  })
  .then(e => e.json())
  .then(e => {
    var content = `
<table class="thetable">
  <tr style="color:#aaa;font-size:1.3em;">
    <th>name</th>
    <th>email </th>
    <th>age </th> 
  </tr>
`;
    for ( var i of e) {
      content += `<tr>`;
      content += `<th>${i.name}</th>`;
      content += `<th>${i.email}</th>`;
      content += `<th>${i.age}</th>`;
      content += `</tr>`;
    }
    content += `</table>`;
    box.innerHTML = content;
  })
  var InputBoxContent = `
  <div id="d">
  name <input id="name" type="text">
  email<input id="email" type="text">
  age<input id="age" type="number">
  <button id="btn">submit</button>
  </div>
  `
  input.innerHTML  = InputBoxContent;
  document.getElementById("btn").onclick = creteInstructor;

}

function createStudent() {
  var name = document.getElementById("name").value;
  var email = document.getElementById("email").value;
  var age  = document.getElementById("age").value;
  fetch(hostUrl+"create_student/",{
    method: 'post',
    body: JSON.stringify({
      "name":name,
      "email":email,
      "age": age,
    }),
  }).then(()=> {
    // window.location.reload()
  }).catch(e => console.log(e))
}



function createDepartment() {
  var name = document.getElementById("name").value;
  fetch(hostUrl+"create_department",{
    method: 'post',
    body: JSON.stringify({
      "name":name,
    }),
  }).then(()=> {
     // window.location.reload()
  }).catch(e => console.log(e))
}

function createCourse() {
  var i = document.getElementById("i_id").value
  var d = document.getElementById("d_id").value
  fetch(hostUrl+"create_course",{
    method: 'post',
    body: JSON.stringify({
      "department_id": parseInt(d),
      "instructor_id":parseInt(i),
    }),
  }).then(()=> {
     // window.location.reload()
  }).catch(e => console.log(e))
}
function creteInstructor() {
  var name = document.getElementById("name").value
  var email = document.getElementById("email").value
  var age = document.getElementById("age").value
  fetch(hostUrl+"create_instructor",{
    method: 'post',
    body: JSON.stringify({
      "name": name, 
      "email": email,
      "age": parseInt(age),
    }),
  }).then(()=> {
     // window.location.reload()
  }).catch(e => console.log(e))
}

function showEnroll() {
  fetch(hostUrl+"/enrolls", {
    method: "GET",
  })
  .then(e => e.json())
  .then(e => {
    var content = `
<table class="thetable">
  <tr style="color:#aaa;font-size:1.3em;">
    <th>student's id</th>
    <th>course's id </th>
  </tr>
`;
console.log(e);
    for ( var i of e) {
      content += `<tr>`;
      content += `<th>${i.student_id}</th>`;
      content += `<th>${i.course_id}</th>`;
      content += `</tr>`;
    }
    content += `</table>`;
    box.innerHTML = content;
  })
  .catch( e => console.log(e))
  var InputBoxContent = `
  <div id="d">
  student's id<input id="s_id" type="number">
  course id <input id="c_id" type="number">
  <button id="btn">submit</button>
  </div>
  `
  input.innerHTML  = InputBoxContent;
  document.getElementById("btn").onclick = createEnroll;
}

function createEnroll() {
  var s = document.getElementById("s_id").value;
  var c = document.getElementById("c_id").value;
  fetch(hostUrl+"create_enroll",{
    method: 'post',
    body: JSON.stringify({
      "course_id": parseInt(c),
      "student_id": parseInt(s),
    }),
  }).then(()=> {
    // window.location.reload()
  }).catch(e => console.log(e))
}