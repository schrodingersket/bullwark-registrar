package cli

import (
    "flag"
)

type ConsulConfig struct {
    Host     *string
    Port     *int
    Scheme   *string
    Prefix   *string
    Watch    *bool
}

func (c ConsulConfig) Configure(configMap map[ConfigType]Config) {

    configMap[ConsulConfigType] = ConsulConfig{
        Host:        flag.String("consul-host", "127.0.0.1", "Consul host"),
        Port:        flag.Int("consul-port", 8500, "Consul port"),
        Scheme:      flag.String("consul-scheme", "http", "Consul scheme"),
        Prefix:      flag.String("consul-prefix", "traefik", "Consul prefix"),
    }

}
