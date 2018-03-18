package configs

// Environment specifies environment for instance development, staging, production, etc..
type Environment int

const (
	// Development environment
	Development Environment = iota
	// Testing environment
	Testing
	// Staging environment
	Staging
	// Production environment
	Production
)

var environments = [...]string{
	"Development",
	"Testing",
	"Staging",
	"Production",
}

// String returns string representation of environment.
func (e Environment) String() string { return environments[e] }
