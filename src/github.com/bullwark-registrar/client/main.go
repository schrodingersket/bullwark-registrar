package client

import (
    "errors"
    "fmt"
    "github.com/dghubble/sling"
    "net/url"
    "os"

    "github.com/bullwark-registrar/common"
    "github.com/bullwark-registrar/common/cli"
    "github.com/bullwark-registrar/handlers/app"
    "github.com/bullwark-registrar/registration"
)

func RunRegistrarClient() {

    var cliConfig = ParseClientConfig()
    clientConfig, ok := common.Configs[cli.ClientConfigType].(cli.ClientConfig)

    if !ok {
        terminate("Unable to parse configuration.")
    }

    fmt.Println("Setting configuration for Traefik services...")

    // Parse all CLI options and create Traefik configuration
    //
    for _, service := range cliConfig.Services {
        PopulateServiceId(*clientConfig.GenerateServiceId, service)
    }

    // Register or deregister services
    //
    if *clientConfig.Deregister {
        if err := DeregisterServices(clientConfig, cliConfig.Services); err != nil {
            terminate(err)
        }
    } else {
        if err := RegisterServices(clientConfig, cliConfig.Services); err != nil {
            terminate(err)
        }
    }
}

func PopulateServiceId(generateServiceId bool, request registration.Request) {

    if generateServiceId {

        if len(request.ServiceId) == 0  {

            // Set service id if not already defined
            //
            request.ServiceId = common.RandString(app.ServiceIdLength)
        } else {
            fmt.Printf("A service id was provided for the '%s' service" +
                ", but the --generate-service-ids flag was set. Defaulting to " +
                "provided service id '%s' instead of generating a new one.\n",
                request.ServiceKey, request.ServiceId)
        }

    } else {
        terminate(fmt.Sprintf("No service id provided for service " +
            "'%s'. If this was intentional, re-run this client with the " +
            "--generate-service-ids flag.",
            request.ServiceKey))
    }

    fmt.Printf("Registering service '%s' with id '%s'.\n",
        request.ServiceKey, request.ServiceId)
}

func RegisterServices(config cli.ClientConfig, requests []registration.Request) error {

    // Generate registration URL
    //
    var registrarUrl = url.URL{
        Scheme: *config.RegistrarScheme,
        Host: fmt.Sprintf("%s:%d", *config.RegistrarHost, *config.RegistrarPort),
        Path: fmt.Sprintf("%s/register/bulk", *config.RegistrarBaseURL),
    }

    // Send request
    //
    var responseBody interface{}
    req, err := sling.New().
        Post(registrarUrl.String()).
        BodyJSON(requests).
        ReceiveSuccess(responseBody)

    if err != nil {
        return err
    }

    if req.StatusCode / 100 != 2 {
        return errors.New(fmt.Sprintf("Request failed with status code" +
            " %d.", req.StatusCode))
    } else {
        fmt.Println("Success.")
    }

    return nil
}

func DeregisterServices(config cli.ClientConfig, requests []registration.Request) error {

    // Generate registration URL
    //
    var registrarUrl = url.URL{
        Scheme: *config.RegistrarScheme,
        Host: fmt.Sprintf("%s:%d", *config.RegistrarHost, *config.RegistrarPort),
        Path: fmt.Sprintf("%s/deregister/bulk", *config.RegistrarBaseURL),
    }

    // Send request
    //
    var responseBody interface{}
    req, err := sling.New().
        Post(registrarUrl.String()).
        BodyJSON(requests).
        ReceiveSuccess(responseBody)

    if err != nil {
        return err
    }

    if req.StatusCode / 100 != 2 {
        return errors.New(fmt.Sprintf("Request failed with status code" +
            " %d.", req.StatusCode))
    } else {
        fmt.Println("Success.")
    }

    return nil
}

func terminate(message interface{}) {
    fmt.Println(message)
    fmt.Println("Exiting.")
    os.Exit(1)
}
