package cli

import (
    "flag"
)

type TraefikConfig struct {
    Url           *string
    AdminUrl      *string
}

func (c TraefikConfig) Configure(configMap map[ConfigType]Config) {

    configMap[TraefikConfigType] = TraefikConfig{
        Url:      flag.String("traefik-url", "http://127.0.0.1:80", "Traefik base URL"),
        AdminUrl:      flag.String("traefik-admin-url", "http://127.0.0.1:8080", "Traefik base admin URL"),
    }

}
