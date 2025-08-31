package blogposts_test

import (
	"errors"
	"io/fs"
	"reflect"
	"testing"
	"testing/fstest"

	blogposts "github.com/Ayikoandrew/gwt/reading-files"
)

func TestPostsFromFs(t *testing.T) {
	t.Run("happy path", func(t *testing.T) {

		fs := fstest.MapFS{
			"hello-world.md": {Data: []byte(`Title: Hello, TDD world!
Description: lol
Tags: tag,lol`)},
		}
		posts, err := blogposts.PostsFromFs(fs)

		if err != nil {
			t.Fatal(err)
		}

		if len(posts) != len(fs) {
			t.Errorf("expected %d posts but got %d posts", len(posts), len(fs))
		}

		expectedBlogpost := blogposts.Post{
			Title:       "Hello, TDD world!",
			Description: "lol",
			Tags:        []string{"tag", "lol"},
		}

		assertResponse(t, posts[0], expectedBlogpost)
	})

	t.Run("failing fs", func(t *testing.T) {
		_, err := blogposts.PostsFromFs(FailingFs{})

		if err == nil {
			t.Errorf("expected an error did get one")
		}
	})
}

func assertResponse(t *testing.T, got, want blogposts.Post) {
	t.Helper()

	if !reflect.DeepEqual(got, want) {
		t.Errorf("got %s, want %s", got, want)
	}
}

type FailingFs struct{}

func (FailingFs) Open(string) (fs.File, error) {
	return nil, errors.New("oh no i always fail")
}
