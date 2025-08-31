package blogposts

import (
	"io"
	"io/fs"
	"strings"
)

type Post struct {
	Title string
}

func PostsFromFs(fileSystem fs.FS) ([]Post, error) {
	dir, err := fs.ReadDir(fileSystem, ".")
	if err != nil {
		return nil, err
	}
	var posts []Post

	for _, f := range dir {
		posts = append(posts, makePostFromFile(fileSystem, f))
	}
	return posts, nil
}

func makePostFromFile(fileSystem fs.FS, f fs.DirEntry) Post {
	blogfile, _ := fileSystem.Open(f.Name())
	return newPost(blogfile)
}

func newPost(blogfile io.Reader) Post {
	fileContent, _ := io.ReadAll(blogfile)
	title := strings.TrimPrefix(string(fileContent), "Title: ")

	return Post{title}
}
