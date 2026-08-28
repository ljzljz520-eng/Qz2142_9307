package config

func Default() Config            { return Config{Addr: ":8080", DBPath: "mathrush.db"} }
func (c Config) Address() string { return c.Addr }
func (c Config) Path() string    { return c.DBPath }
