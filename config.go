package go_poro

type Config struct {
	BaseURL string
	ChampionPath string
	tagURL string
}

func NewConfig(opts ...ConfigOption) Config {
	const (
		base         = "https://www.leagueofgraphs.com"
		championPath = "/summoner/champions"
	)

	conf := Config{
		BaseURL: base,
		ChampionPath: championPath,
	}

	for _, opt := range opts {
		opt(&conf)
	}

	conf.tagURL = conf.BaseURL + conf.ChampionPath

	return conf
}

type ConfigOption = func(*Config)

func ConfigBaseUrlOpt(url string) ConfigOption {
	return func(c *Config) {
		c.BaseURL = url
	}
}