package config

//	"os"

// ENV ...
type ENV struct {
	IsDev bool

	ZookeeperURI     string
	ZookeeperTestURI string

	// App port
	AppUserPort string

	// Database
	Database struct {
		URI            string
		TestName       string
		UserName       string
	}

	GRPCUri string
}

var env ENV

// InitENV ...
func InitENV() {
	env = ENV{
		IsDev: true,
	}
}

// GetEnv ...
func GetEnv() *ENV {
	return &env
}
