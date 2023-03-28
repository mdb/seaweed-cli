package config

type Config struct {
	Spots []Spot `yaml:"spots"`
}

type Spot struct {
	ID   string `yaml:"id,omitempty"`
	Name string `yaml:"name,omitempty"`
}

type ConfigParser struct{}

func (parser ConfigParser) getDefaultConfig() Config {
	return Config{
		Spots: []Spot{{
			ID:   "391",
			Name: "Ocean City, NJ",
		}},
	}
}

func ParseConfig(path string) (Config, error) {
	parser := ConfigParser{}

	return parser.getDefaultConfig(), nil
}
