db.createCollection("posts")

content = `
# My Awesome Blog Post

## Introduction

Welcome to my blog! In this post, I'll be sharing some thoughts on a topic that's been on my mind lately. Grab a cup of coffee and let's dive in.

## The Importance of Learning Markdown

Markdown is a lightweight markup language that allows you to write and format text easily. Whether you're a blogger, a developer, or just someone who loves to take notes, learning Markdown can be a game-changer. Here are a few reasons why:

### 1. **Simplicity**

Markdown is incredibly easy to learn. You don't need to be a coding expert to use it. With just a few simple symbols, you can create well-structured documents.

### 2. **Compatibility**

Markdown is widely supported across different platforms and applications. Whether you're writing a blog post, updating your GitHub readme, or taking notes in a Markdown editor, your content will look consistent.

### 3. **Focus on Content**

One of the best things about Markdown is that it allows you to focus on your content rather than formatting. You can quickly add headings, lists, and links without getting bogged down by complex formatting options.

## Getting Started with Markdown

If you're new to Markdown, here's a quick guide to get you started:

### Headers

\`\`\`markdown
# This is a H1 header
## This is a H2 header
### This is a H3 header
\`\`\`
`

for (let i = 0; i < 1000; i++) {
    let date = new Date(Date.now()).toISOString();
    let post = { id: `${i}`, title: "title", content: content, created: date }
    db.posts.insert(post)
}
