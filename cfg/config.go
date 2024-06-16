package cfg

type Config struct {
	Participents []struct {
		Name  string    `yaml:"name"`
		Teams [3]string `yaml:"teams"`
	} `yaml:"participents"`
}
