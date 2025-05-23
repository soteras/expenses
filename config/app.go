package config

type app struct {
	ENV string
}

func NewApp() *app {
	return &app{
		ENV: GetEnv("GO_ENV"),
	}
}
