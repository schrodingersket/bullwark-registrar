package app

import (
    "encoding/json"
    "fmt"
    "net/http"

    "github.com/bullwark-registrar/common"
    "github.com/bullwark-registrar/registration"
)

const (
    ServiceIdLength = 32
)

var registrar registration.Registrar

func Register(w http.ResponseWriter, r *http.Request) {

    var registrationRequest registration.Request
    var serviceId = common.RandString(ServiceIdLength)
    var encoder = json.NewEncoder(w)

    // TODO: Replace with DI
    //
    if registrar == nil {
        registrar = registration.NewTraefikRegistrar()
    }

    // Parse body
    //
    err := json.NewDecoder(r.Body).Decode(&registrationRequest)

    if err != nil {
        fmt.Println(err)
        _ = encoder.Encode(registration.Response{
            Status: "error",
            Reason: fmt.Sprintf("%s", err),
        })
        return
    }

    // Register with service registration impl
    //
    err = registrar.Register(registrationRequest)

    if err != nil {
        fmt.Println(err)
        _ = encoder.Encode(registration.Response{
            Status: "error",
            Reason: fmt.Sprintf("%s", err),
        })
        return
    }

    // Success
    //

    _ = encoder.Encode(registration.Response{
        ServiceId: serviceId,
        Status:    "success",
    })
    return
}

func RegisterBulk(w http.ResponseWriter, r *http.Request) {

    var registrationRequest []registration.Request
    var encoder = json.NewEncoder(w)

    // TODO: Replace with DI
    //
    if registrar == nil {
        registrar = registration.NewTraefikRegistrar()
    }

    // Parse body
    //
    var err = json.NewDecoder(r.Body).Decode(&registrationRequest)

    if err != nil {
        fmt.Println(err)
        _ = encoder.Encode(registration.Response{
            Status: "error",
            Reason: fmt.Sprintf("%s", err),
        })
        return
    }

    // Register with service registration impl
    //
    for _, service := range registrationRequest {
        var err = registrar.Register(service)

        if err != nil {
            fmt.Println(err)
            _ = encoder.Encode(registration.Response{
                Status: "error",
                Reason: fmt.Sprintf("%s", err),
            })
            return
        }
    }

    _ = encoder.Encode(registration.Response{
            Status:    "success",
        })
    return
}
