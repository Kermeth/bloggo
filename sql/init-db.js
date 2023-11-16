db.createCollection("posts")

for (let i = 0; i < 1000; i++) {
    let date = new Date(Date.now()).toISOString();
    let post = { id: `${i}`, title: "title", content: "content", created: date }
    db.posts.insert(post)
}
