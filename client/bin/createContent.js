let fs = require ('fs')
let path = require ('path')

const markdownContent = `---
title: ''
description: ''
cover: ''
---`


// 构建文件名
const fileName = `${(new Date()).getTime()}.md`

const dir = path.resolve(__dirname, '../content')
// 文件内容
const filePath = path.join(dir, fileName)

// 写入文件内容
fs.writeFile(filePath, markdownContent, (err) => {
  if (err) {
    console.error('Error writing Markdown file:', err)
  } else {
    console.log(`Markdown file ${fileName} created successfully.`)
  }
})
