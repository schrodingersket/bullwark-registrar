package cli

import (
    "flag"
)

type ClientConfig struct {
    GenerateServiceId  *bool
    ConfigPath         *string
    RegistrarHost      *string
    RegistrarScheme    *string
    RegistrarPort      *int
    RegistrarBaseURL   *string
    Deregister         *bool
}

func (c ClientConfig) Configure(configMap map[ConfigType]Config) {

    configMap[ClientConfigType] = ClientConfig{
        GenerateServiceId :   flag.Bool("generate-service-ids", false, "Set this to true to enable automatic generation of service ids"),
        ConfigPath        :   flag.String("config", "services.yml", "Path to CLI config, if running in CLI mode"),
        RegistrarHost     :   flag.String("registrar-host", "127.0.0.1", "Registrar server  host used for CLI"),
        RegistrarPort     :   flag.Int("registrar-port", 8000, "Registrar server port used for CLI"),
        RegistrarScheme   :   flag.String("registrar-scheme", "http", "Registrar scheme used for CLI"),
        RegistrarBaseURL  :   flag.String("registrar-base-url", "", "Registrar scheme used for CLI"),
        Deregister        :   flag.Bool("deregister", false, "Set this to true to deregister services"),
    }
}
