package authn

// Method defines an authentication method
type Method interface {
	Authenticate() ([]byte, error)
	IsAvailable() bool
}
