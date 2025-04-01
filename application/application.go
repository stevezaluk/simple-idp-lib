package application

/*
Application - A user defined application. Users will define these and authorize there applications
to use API's that are defined
*/
type Application struct {
}

/*
NewApplication - A constructor for the Application struct
*/
func NewApplication() *Application {
	return &Application{}
}
