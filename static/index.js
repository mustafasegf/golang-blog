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
      const target = document.getElementById("content")
      target.innerHTML += `
      <h1>${data.title}</h1>
      <h3>${data.name}</h3>
      <p>${data.content}</p>`
      loadComment(id)
    });
}

const loadComment = (id) => {
  fetch(`http://localhost:3000/api/blogs/${id}/comments`)
    .then(response => response.json())
    .then(data => {
      console.log(data)
      const comments = document.getElementById("comments")
      if (data === null) {
        comments.innerHTML += `<p> None </p>`
      } else {
        data.forEach(e => {
          comments.innerHTML += `
          <p id="comment-${e.id}">
            ${e.name}:<span id="span-comment-${e.id}">${e.comment}</span>
            <span class="delete" onclick="deleteComment(${e.id})">delete</span>
            <span class="edit" onclick="editComment(${e.id})">edit</span>
          </p>`
        });
      }
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
      'Authorization': `bearer ${window.sessionStorage.accessToken}`
    },
    body: JSON.stringify(data)
  })
    .then(response => response.json())
    .then(resdata => {
      console.log(resdata)
      const comments = document.getElementById("comments")
      if ('id' in resdata) {
        comments.innerHTML += `
          <p id="comment-${resdata.id}">
            ${resdata.name}:<span id="span-comment-${resdata.id}">${resdata.comment} </span>
            <span class="delete" onclick="deleteComment(${e.id})">delete</span>
            <span class="edit" onclick="editComment(${e.id})">edit</span>
          </p>`
      }
    })
}

const deleteComment = (id) => {
  console.log(id)
  fetch(`http://localhost:3000/api/comments/${id}`, {
    method: 'DELETE',
    headers: {
      'Authorization': `bearer ${window.sessionStorage.accessToken}`
    },
  })
    .then(response => response.json())
    .then(resdata => {
      console.log('data', resdata)
      if ('Comment deleted' === resdata) {
        const comment = document.getElementById("comments").querySelector(`#comment-${id}`)
        comment.remove()
      }
    })
}

const editComment = (id) => {
  console.log(id)
  const comment = document.getElementById("comments").querySelector(`#comment-${id}`)
  const form = document.createElement("div");

  form.innerHTML += `
  <form action="" id="edit-form" method="POST" onsubmit="fetchEditComment(event, ${id});">
    <label for="comment"> comment</label>
    <input type="text" name="comment" />
    <input type="submit" name="submit" value="edit" />
  </form>`
  if (comment.querySelector('#edit-form') !== null ) {
    comment.querySelector('#edit-form').remove()
  } else {
    comment.appendChild(form)
  }
}

const fetchEditComment = (e, id) => {
  e.preventDefault()
  const form = document.getElementById("edit-form")
  const formData = new FormData(form);
  let data = {}
  for (let key of formData.keys()) {
    data[key] = formData.get(key);
  }
  console.log(data)
  fetch(`http://localhost:3000/api/comments/${id}`, {
    method: 'PATCH',
    headers: {
      'Content-Type': 'application/json',
      'Authorization': `bearer ${window.sessionStorage.accessToken}`
    },
    body: JSON.stringify(data)
  })
    .then(response => response.json())
    .then(resdata => {
      console.log(resdata)
      const comment = document.getElementById("comments").querySelector(`#span-comment-${id}`)
      comment.innerHTML = data.comment
    })
}