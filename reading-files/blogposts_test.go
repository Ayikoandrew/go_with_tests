package blogposts_test

import (
	"errors"
	"io/fs"
	"testing"
	"testing/fstest"

	blogposts "github.com/Ayikoandrew/gwt/reading-files"
)

func TestBlogposts(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {

		fs := fstest.MapFS{
			"hello-world.md": {Data: []byte("Title: Hello, TDD world")},
			//"hello-twitch.md": {Data: []byte("Title: Hello, twitchy world")},
		}
		posts, err := blogposts.PostsFromFs(fs)

		if err != nil {
			t.Fatal(err)
		}

		if len(posts) != len(fs) {
			t.Errorf("expected %d posts but got %d posts", len(posts), len(fs))
		}

		expectedBlogpost := blogposts.Post{Title: "Hello, TDD world"}

		if posts[0] != expectedBlogpost {
			t.Errorf("got %s, want %s", posts[0], expectedBlogpost)
		}
	})

	t.Run("failing fs", func(t *testing.T) {
		_, err := blogposts.PostsFromFs(FailingFs{})

		if err == nil {
			t.Errorf("expected an error did get one")
		}
	})
}

type FailingFs struct{}

func (FailingFs) Open(string) (fs.File, error) {
	return nil, errors.New("oh no i always fail")
}
