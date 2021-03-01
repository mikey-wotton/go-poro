package go_poro

const (
	base         = "https://www.leagueofgraphs.com"
	championPath = "/summoner/champions"
)

type Config struct {
	BaseURL string
	ChampionPath string
	tagURL string
}

func DefaultConfig() Config {
	return Config{
		BaseURL: base,
		ChampionPath: championPath,
		tagURL: base + championPath,
	}
}