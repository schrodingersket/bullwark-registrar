package cli

import (
    "flag"
)

type CoreConfig struct {
    Client         *bool
    Port           *int
    Verbose        *bool
    BaseURL        *string
}

func (c CoreConfig) Configure(configMap map[ConfigType]Config) {

  configMap[CoreConfigType] = CoreConfig{
      Client           :   flag.Bool("client", false, "Run in CLI mode."),
      Port             :   flag.Int("port", 8000, "Port to run on."),
      Verbose          :   flag.Bool("verbose", false, "Turn on verbose logging"),
      BaseURL          :   flag.String("base-url", "/", "Determines the base URL which the application serves at."),
  }

}
