package blogposts

import (
	"bufio"
	"io"
	"io/fs"
	"strings"
)

type Post struct {
	Title       string
	Description string
	Tags        []string
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
	// fileContent, _ := io.ReadAll(blogfile)
	// title := strings.TrimPrefix(string(fileContent), "Title: ")
	scanner := bufio.NewScanner(blogfile)
	readLine := func(prefix string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), prefix)
	}

	title := readLine("Title: ")
	description := readLine("Description: ")
	tags := strings.Split(readLine("Tags: "), ",")
	return Post{
		Title:       title,
		Description: description,
		Tags:        tags,
	}
}
