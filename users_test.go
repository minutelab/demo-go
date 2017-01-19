package demo

import (
	"reflect"
	"testing"

	"github.com/minutelab/go-mlabtest/sqitchdb"
)

func TestUsers(t *testing.T) {
	db, _ := sqitchdb.New(t, "schema", "", nil)
	defer db.Close()

	Init(db.Conn())

	// Store first user
	u1 := User{
		ID:    "user1",
		Name:  "John Doe",
		Email: "jdoe@gmail.com",
	}

	if err := Store(u1); err != nil {
		t.Fatal("Failed storing user1:", err)
	}

	// Try to get it
	u2, err := Get("user1")
	if err != nil {
		t.Fatal("Failed getting user1:", err)
	}
	t.Log("Got user1:", *u2)
	if !reflect.DeepEqual(u1, *u2) {
		t.Error("Got different users")
	}

	// Storing the same user twice should fail
	if err := Store(u1); err == nil {
		t.Error("Error: stored the same user twice")
	} else {
		t.Log("As expected cannot store the same user twice:", err)
	}

	// try to get no user
	if u, err := Get("none"); err != nil {
		t.Log("As expected cannot get none user:", err)
	} else {
		t.Error("Error: got a non-existing user: ", *u)
	}
}
