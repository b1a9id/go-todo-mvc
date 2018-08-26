// Package model contains the types for schema 'gtm'.
package model

// Code generated by xo. DO NOT EDIT.

import (
	"database/sql"
	"database/sql/driver"
	"encoding/csv"
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"
)

// Brand represents a row from 'gtm.brand'.
type Brand struct {
	ID uint `json:"id"` // id
	Name string `json:"name"` // name
	Designer sql.NullString `json:"designer"` // designer
	CreatedAt time.Time `json:"created_at"` // created_at
	UpdatedAt time.Time `json:"updated_at"` // updated_at

	// xo fields
	_exists, _deleted bool

}


// Exists determines if the Brand exists in the database.
func (b *Brand) Exists() bool {
	return b._exists
}

// Deleted provides information if the Brand has been deleted from the database.
func (b *Brand) Deleted() bool {
	return b._deleted
}

// Insert inserts the Brand to the database.
func (b *Brand) Insert(db XODB) error {
	var err error

	// if already exist, bail
	if b._exists {
		return errors.New("insert failed: already exists")
	}



	// sql insert query, primary key must be provided
	const sqlstr = `INSERT INTO gtm.brand (` +
		`id, name, designer, created_at, updated_at` +
		`) VALUES (` +
		`?, ?, ?, ?, ?` +
		`)`

	// run query
	XOLog(sqlstr, b.ID, b.Name, b.Designer, b.CreatedAt, b.UpdatedAt)
	_, err = db.Exec(sqlstr, b.ID, b.Name, b.Designer, b.CreatedAt, b.UpdatedAt)
	if err != nil {
		return err
	}

	// set existence
	b._exists = true


	return nil
}


	// Update updates the Brand in the database.
	func (b *Brand) Update(db XODB) error {
		var err error

		// if doesn't exist, bail
		if !b._exists {
			return errors.New("update failed: does not exist")
		}

		// if deleted, bail
		if b._deleted {
			return errors.New("update failed: marked for deletion")
		}

		
			// sql query
			const sqlstr = `UPDATE gtm.brand SET ` +
				`name = ?, designer = ?, created_at = ?, updated_at = ?` +
				` WHERE id = ?`

			// run query
			XOLog(sqlstr, b.Name, b.Designer, b.CreatedAt, b.UpdatedAt, b.ID)
			_, err = db.Exec(sqlstr, b.Name, b.Designer, b.CreatedAt, b.UpdatedAt, b.ID)
			return err
	}

	// Save saves the Brand to the database.
	func (b *Brand) Save(db XODB) error {
		if b.Exists() {
			return b.Update(db)
		}

		return b.Insert(db)
	}


// Delete deletes the Brand from the database.
func (b *Brand) Delete(db XODB) error {
	var err error

	// if doesn't exist, bail
	if !b._exists {
		return nil
	}

	// if deleted, bail
	if b._deleted {
		return nil
	}

	
		// sql query
		const sqlstr = `DELETE FROM gtm.brand WHERE id = ?`

		// run query
		XOLog(sqlstr, b.ID)
		_, err = db.Exec(sqlstr, b.ID)
		if err != nil {
			return err
		}

	// set deleted
	b._deleted = true

	return nil
}

// BrandByID retrieves a row from 'gtm.brand' as a Brand.
//
// Generated from index 'brand_id_pkey'.
func BrandByID(db XODB, id uint) (*Brand, error) {
	var err error

	// sql query
	const sqlstr = `SELECT ` +
		`id, name, designer, created_at, updated_at ` +
		`FROM gtm.brand ` +
		`WHERE id = ?`

	// run query
	XOLog(sqlstr, id)
	b := Brand{
		_exists: true,
	}

	err = db.QueryRow(sqlstr, id).Scan(&b.ID, &b.Name, &b.Designer, &b.CreatedAt, &b.UpdatedAt)
	if err != nil {
		return nil, err
	}

	return &b, nil
}

