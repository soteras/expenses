package config

type app struct {
	ENV string
}

func NewAPP() *app {
	return &app{
		ENV: GetEnv("GO_ENV"),
	}
}
