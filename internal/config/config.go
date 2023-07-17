package config

type Config struct {
	// api-server
	Server string
	// database
	Database        string
	MaxConn         int
	MaxIdleConn     int
	MaxLifetimeConn int
}
