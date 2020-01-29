package authentication

// type user represents a member of staff in the database, but only contains
// StaffID and PasswordHash fields. The actual staff table is created by the
// staff resource.
type staff struct {
  StaffID int64 `pg:",pk" json:"staff_id"`
  PasswordHash []byte `pg:",notnull" json:"password_hash"`
}
