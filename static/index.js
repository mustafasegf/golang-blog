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
      if (data === null) {
        data = [{ comment: 'none' }]
      }
      data.forEach(e => {
        let p = document.createElement('p');
        p.innerHTML = e.comment;
        target.appendChild(p);
      });
    });
}

const register = (e) => {
  const form = document.getElementById("form")
  e.preventDefault()
  const formData = new FormData(form);
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
    .then(resdata => {
      console.log(resdata)
      if (!('error' in resdata)) {
        window.location = '/login'
      }
    })
}

const login = (e) => {
  const form = document.getElementById("form")
  e.preventDefault()
  const formData = new FormData(form);
  let data = {}
  for (let key of formData.keys()) {
    data[key] = formData.get(key);
  }
  console.log(data)

  fetch('http://localhost:3000/api/users/login', {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
    },
    body: JSON.stringify(data)
  })
    .then(response => response.json())
    .then(resdata => {
      console.log(resdata)
      window.sessionStorage.accessToken = resdata.access_token
      if (!('error' in resdata)) {
        window.location = '/'
      }
    })
}

const addComment = (e, id) => {
  const form = document.getElementById("form")
  e.preventDefault()
  const formData = new FormData(form);
  let data = {}
  for (let key of formData.keys()) {
    data[key] = formData.get(key);
  }
  console.log(data)
  fetch(`http://localhost:3000/api/blogs/${id}/comments`, {
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
      'Authorization' : `bearer ${window.sessionStorage.accessToken}`
    },
    body: JSON.stringify(data)
  })
    .then(response => response.json())
    .then(resdata => {
      console.log(resdata)
    })
}