package cli

type ConfigType int

const (
    CoreConfigType    ConfigType = iota
    ConsulConfigType
    TraefikConfigType
    ClientConfigType
)
