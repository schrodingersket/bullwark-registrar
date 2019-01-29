package cli

type Config interface {
    Configure(map[ConfigType]Config)
}

type Configs []Config

var ConfigList = Configs{
    CoreConfig{},
    ConsulConfig{},
    TraefikConfig{},
    ClientConfig{},
}
