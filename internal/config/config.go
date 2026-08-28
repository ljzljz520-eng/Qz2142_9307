package config

import "os"

type Config struct{ Addr, DBPath string }

func Load() Config {
	a := os.Getenv("MATH_ADDR")
	if a == "" {
		a = ":8080"
	}
	p := os.Getenv("MATH_DB")
	if p == "" {
		p = "mathrush.db"
	}
	return Config{a, p}
}
func (c Config) IsValid() bool { return c.Addr != "" && c.DBPath != "" }
