package server

// Does not matter what server (mux, chi, echo) we use
// our code must implement the Run method
type Server interface {
	Run()
}
