package config

import (
	"github.com/cristalhq/aconfig"
	"github.com/cristalhq/aconfig/aconfigdotenv"
)

func Load(cfg interface{}) error {
	return aconfig.LoaderFor(cfg, aconfig.Config{
		SkipEnv:            false,
		EnvPrefix:          "SOCIALIST",
		FailOnFileNotFound: false,
		Files:              []string{".env"},
		FileDecoders: map[string]aconfig.FileDecoder{
			".env": aconfigdotenv.New(),
		},
	}).Load()
}
