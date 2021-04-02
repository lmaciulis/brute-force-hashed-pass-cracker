package config

type Config struct {
	AvailableChars string         `json:"available_chars"`
	MaxPassLength  int            `json:"max_pass_length"`
	Prefixes       PrefixesConfig `json:"prefixes"`
	Suffixes       SuffixesConfig `json:"suffixes"`
}

type PrefixesConfig struct {
	Enabled bool     `json:"enabled"`
	List    []string `json:"list"`
}

type SuffixesConfig struct {
	Enabled bool     `json:"enabled"`
	List    []string `json:"list"`
}
