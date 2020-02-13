package staff

import (
	"api-oaxaca-com/packages/db"
	"api-oaxaca-com/packages/auth"
)

// type dbWrapper wraps a database connection and provides higher-level
// operations/abstractions on the database..
type dbWrapper struct {
	db *db.DB
}

// CreateStaff takes a pointer to an instance of a createInput struct and
// then creates a new staff model - which is then inserted into the
// database.
func (w *dbWrapper) CreateStaff(i *createInput) error {
	passwordHash, err := auth.HashPassword(i.Password)
	if err != nil {
		return err
	}

	// construct record
	r := staff{
		RestaurantID: i.RestaurantID,
		FullName: i.FullName,
		Score: i.Score,
		PasswordHash: passwordHash,
		HasPassedTraining: i.HasPassedTraining,
		Role: i.Role,
	}

	// insert record and handle error
	if _, err := w.db.Conn.
		Model(&r).
		Insert(); err != nil {
		return err
	} else {
		return nil
	}
}

// GetStaff takes an id and fetches the staff record in the database
// with that id.
func (w *dbWrapper) GetStaff(id int64) (*staff, error) {
	// construct empty record
	r := new(staff)

	// run query and handle error
	if err := w.db.Conn.
		Model(r).
		Where("staff_id = ?", id).
		Select(); err != nil {
		return nil, err
	} else {
		return r, nil
	}
}

// GetAllStaff returns a slice containing all staff records in the
// database.
func (w *dbWrapper) GetAllStaff() ([]staff, error) {
	// define space for result set
	var staffs []staff

	// run query and handle error
	if err := w.db.Conn.
		Model(&staffs).
		Select(); err != nil {
		return nil, err
	} else {
		return staffs, nil
	}
}

// UpdateStaff takes an id and a pointer to an instance of an updateInput
// struct and then fetches the staff record in the database with the
// corresponding id - and updates its fields.
func (w *dbWrapper) UpdateStaff(id int64, i *updateInput) error {
	// construct empty record
	r := staff{
		StaffID: id,
	}

	// assign changed fields
	if i.RestaurantID != nil {
		r.RestaurantID = *i.RestaurantID
	}
	if i.FullName != nil {
		r.FullName = *i.FullName
	}
	if i.Score != nil {
		r.Score = *i.Score
	}
	if i.Password != nil {
		if passwordHash, err := auth.HashPassword(*i.Password); err != nil {
			return err
		} else {
			r.PasswordHash = passwordHash
		}
	}
	if i.HasPassedTraining != nil {
		r.HasPassedTraining = *i.HasPassedTraining
	}
	if i.Role != nil {
		r.Role = *i.Role
	}

	// run query and handle error
	if _, err := w.db.Conn.
		Model(&r).
		WherePK().
		UpdateNotZero(); err != nil {
		return err
	} else {
		return nil
	}
}

// DeleteStaff takes an id and deletes the record in the database with
// that id.
func (w *dbWrapper) DeleteStaff(id int64) error {
	// construct record with ID only
	r := staff{
		StaffID: id,
	}

	// run query and handle error
	if _, err := w.db.Conn.
		Model(&r).
		WherePK().
		Delete(); err != nil {
		return err
	} else {
		return nil
	}
}
