package role

/*
Role - A user defined role that be attached to a User. Provides a structure
to attach permissions to
*/
type Role struct {
}

/*
NewRole -
*/
func NewRole() *Role {
	return &Role{}
}
