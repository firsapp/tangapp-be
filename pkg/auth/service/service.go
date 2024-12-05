package service

// AuthRepository defines the datastore contract to be implemented.
type AuthRepository interface {
	//Auth
}

// Auth defines the application service with the required dependencies.
type Auth struct {
	AuthRepo AuthRepository
}

// NewAuth returns an instance of Units service.
func NewAuth(AuthRepo AuthRepository) *Auth {
	return &Auth{
		AuthRepo: AuthRepo,
	}
}
