package utils

import (
	"fmt"

	scribble "github.com/nanobox-io/golang-scribble"
)

// GetDatabase returns a scribble database object to store the command examples in
func GetDatabase() *scribble.Driver {
	dir := "./data"
	db, err := scribble.New(dir, nil)
	if err != nil {
		fmt.Println("Error", err)
	}
	return db
}
