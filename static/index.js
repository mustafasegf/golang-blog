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
      <h1 id="title">${data.title}</h1>
      <h3 id="author">Author: ${data.name}</h3>
      <p id="content">${data.content}</p>
      <p>
        <span class="delete" onclick="deleteBlog(${id})">delete</span>
        <span class="edit" onclick="editBlog(${id})">edit</span>
      </p>`
      loadComment(id)
    });
}

const addBlog = (e) => {
  const form = document.getElementById("form")
  e.preventDefault()
  const formData = new FormData(form);
  let data = {}
  for (let key of formData.keys()) {
    data[key] = formData.get(key);
  }
  console.log(data)
  fetch(`http://localhost:3000/api/blogs`, {
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
      const blog = document.getElementById("blog")
      if ('title' in resdata) {
        blog.innerHTML += `
        <td>${resdata.title}</td>
        <td>${resdata.content}</td>
        <td>${resdata.name}</td>
        <td><a href=http://localhost:3000/blogs/${resdata.id}>click here</a></td>`;
      }
    })
}

const editBlog = (id) => {
  console.log(id)
  const blog = document.getElementById("content").querySelector('#content')
  const form = document.createElement("div")

  form.innerHTML += `
  <form action="" id="edit-blog-form" method="POST" onsubmit="fetchEditBlog(event, ${id});">
    <label for="title"> title</label>
    <input type="text" name="title" />
    <label for="content"> content</label>
    <input type="text" name="content" />
    <input type="submit" name="submit" value="edit" />
  </form>`
  if (blog.querySelector('#edit-blog-form') !== null) {
    blog.querySelector('#edit-blog-form').remove()
  } else {
    blog.appendChild(form)
  }
}

const fetchEditBlog = (e, id) => {
  e.preventDefault()
  const form = document.getElementById("edit-blog-form")
  const formData = new FormData(form);
  let data = {}
  for (let key of formData.keys()) {
    data[key] = formData.get(key);
  }
  console.log(data)
  fetch(`http://localhost:3000/api/blogs/${id}`, {
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
      const blog = document.getElementById("content")
      const title = blog.querySelector('#title')
      const content = blog.querySelector('#content')
      if (data.title !== '' && data.title !== undefined) {
        title.innerHTML = data.title
      }

      if (data.content !== '' && data.content !== undefined) {
        content.innerHTML = data.content
      }
    })
}

const deleteBlog = (id) => {
  console.log(id)
  fetch(`http://localhost:3000/api/blogs/${id}`, {
    method: 'DELETE',
    headers: {
      'Authorization': `bearer ${window.sessionStorage.accessToken}`
    },
    body: { id }
  })
    .then(response => response.json())
    .then(resdata => {
      console.log('data', resdata)
      if ('Blog deleted' === resdata) {
        window.location = '/'
      }
    })
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
  if (comment.querySelector('#edit-form') !== null) {
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