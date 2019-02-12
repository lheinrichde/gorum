package handlers

import (
	"errors"
	"log"

	"github.com/ltheinrich/gorum/internal/pkg/db"
)

// DeleteAccount handler
func DeleteAccount(data HandlerData) interface{} {
	var err error

	// authenticate
	if !data.Authenticated {
		// not authenticated
		return errors.New("403")
	}

	// delete account
	_, err = db.DB.Exec("DELETE FROM users WHERE username = $1;", data.Username)
	if err != nil {
		// print and return error
		log.Println(err)
		return err
	}

	// write map
	return map[string]interface{}{"success": true}
}
