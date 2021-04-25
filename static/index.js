const loadBlog = () => {
  fetch("http://localhost:3000/api/blogs")
    .then(response => response.json())
    .then(data => {
      console.log(data)
      const target = document.getElementById("blog")
      target.innerHTML += `<tr>
      <th>Title</th>
      <th>Content</th>
      <th>Author</th>
      <th>url</th>
      </tr>`
      data.forEach(e => {
        let tr = document.createElement('tr');
        tr.innerHTML += `
        <td>${e.title}</td>
        <td>${e.content}</td>
        <td>${e.name}</td>
        <td><a href=http://localhost:3000/blogs/${e.id}>click here</a></td>`;
        target.appendChild(tr);
      });
    });
}

const loadBlogById = (id) => {
  fetch(`http://localhost:3000/api/blogs/${id}`)
    .then(response => response.json())
    .then(data => {
      console.log(data)
      const target = document.getElementById("blog")
      let div = document.createElement('div');
      div.innerHTML += `
      <h1>${data.title}</h1>
      <h3>${data.name}</h3>
      <p>${data.content}</p>`
      target.appendChild(div);

      loadComment(id)
    });
}
const loadComment = (id) => {
  fetch(`http://localhost:3000/api/blogs/${id}/comments`)
    .then(response => response.json())
    .then(data => {
      console.log(data)
      const target = document.getElementById("blog")
      let comment = document.createElement('h3');
      comment.innerHTML = "Comments"
      target.appendChild(comment);
      data.forEach(e => {
        let p = document.createElement('p');
        p.innerHTML = e.comment;
        target.appendChild(p);
      });
    });
}

const register = () => {
  const form = document.getElementById("form")
  form.addEventListener("submit", (e) => {
    e.preventDefault()
    // let Name = document.getElementsByName("name")[0]
    // let Username = document.getElementsByName("username")
    // let Password = document.getElementsByName("password")
    // const form = document.querySelector('form')
    const formData  = new FormData(form);
    let data = {}
    for (let key of formData.keys()) {
      data[key] = formData.get(key);
    }
    console.log(data)

    fetch('http://localhost:3000/api/users', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(data)
    })
      .then(response => response.json())
      .then(data => console.log(data))
  })
}