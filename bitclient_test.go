package bitclient

import (
	"fmt"
	"testing"
)

func TestCreateRepository(t *testing.T) {
	c := NewBitClient("http://localhost", "Foo", "foo")

	data := CreateRepositoryRequest{
		Name:     "test1",
		ScmId:    "git",
		Forkable: false,
	}
	r, e := c.CreateRepository("FOO", data)
	fmt.Println(r, e)
}
