package client

import (
    "fmt"
    "io/ioutil"
    "os"

    "gopkg.in/yaml.v2"

    "github.com/bullwark-registrar/common"
    "github.com/bullwark-registrar/common/cli"
    "github.com/bullwark-registrar/registration"
)

type Config struct {
    Services  []registration.Request `json:"services"`
}

func ParseClientConfig() Config {

    var appConfig = common.Configs[cli.ClientConfigType].(cli.ClientConfig)
    var cliConfigPath = *appConfig.ConfigPath
    var cliConfig Config

    // Check if config file exists
    //
    _, err := os.Stat(cliConfigPath)

    if err != nil {
        fmt.Printf("err: %v\n", err)
        os.Exit(1)
    }

    // Read and unmarshal config
    //
    ymlFile, err := ioutil.ReadFile(cliConfigPath)

    err = yaml.Unmarshal(ymlFile, &cliConfig)
    if err != nil {
        fmt.Printf("err: %v\n", err)
        os.Exit(1)
    }

    return cliConfig
}
