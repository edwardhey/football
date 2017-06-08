package models

import (
	"testing"
)

func TestNewId(t *testing.T) {
	idGenrator, _ := NewSnowFlake(1)
	// id2, _ := NewSnowFlake(2)

	id1, _ := idGenrator.Next()
	id2, _ := idGenrator.Next()
	if id1 == id2 {
		t.Error("id genr same")
	}
	// fmt.Println(NewActivatyID())
	// t.Error("aaaaaaaaaa")
}
