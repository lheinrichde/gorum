package handlers

import (
	"database/sql"
	"errors"

	"github.com/lheinrichde/gorum/pkg/db"
)

// Board handler
func Board(request map[string]interface{}, username string, auth bool) interface{} {
	var err error

	// check if board id provided
	boardID := GetInt(request, "boardID")
	if boardID == 0 {
		// not provided
		return errors.New("400")
	}

	// define variables
	var sort, category int
	var name, description, icon string

	// query user by id
	err = db.DB.QueryRow("SELECT boardname, boarddescription, boardicon, sort, category FROM boards WHERE id = $1;", boardID).Scan(&name, &description, &icon, &sort, &category)

	// check not found
	if err == sql.ErrNoRows {
		// return not found
		return errors.New("404")
	} else if err != nil {
		// return error
		return err
	}

	// board map to write
	thread := map[string]interface{}{}
	thread["id"] = boardID
	thread["name"] = name
	thread["description"] = description
	thread["icon"] = icon
	thread["sort"] = sort
	thread["category"] = category

	// write map
	return thread
}