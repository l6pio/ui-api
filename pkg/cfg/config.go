package cfg

type DbConfig struct {
	DbAddr, DbUser, DbPass, DbName string
}

type Config struct {
	*DbConfig
}

func New(dbAddr string, dbUser string, dbPass string) *Config {
	return &Config{
		DbConfig: &DbConfig{
			DbAddr: dbAddr,
			DbUser: dbUser,
			DbPass: dbPass,
			DbName: "l6p-report",
		},
	}
}
