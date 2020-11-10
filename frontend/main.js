
const hostUrl = window.location.href;
const box  = document.getElementById("box");
const input = document.getElementById("input");
window.onload = () => {
  console.log("startedd......")
  document.getElementById("students").onclick = showStudents;
  document.getElementById("instructors").onclick = showInstructors;
  document.getElementById("departments").onclick = showDepatments;
  document.getElementById("courses").onclick = showCourses;
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
  console.log("showing departments table");
}

function showCourses() {
  console.log("showing courses table");
}

function showInstructors() {
  console.log("showing instructor table");
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