package main

import (
    "fmt"
    "github.com/gorilla/handlers"
    "net/http"

    "github.com/bullwark-registrar/client"
    "github.com/bullwark-registrar/common"
    "github.com/bullwark-registrar/routes"
)

func main() {

    // Parse command-line flags
    //
    common.ConfigureFromFlags()


    if common.IsClientMode() {

        // Run client registration
        //
        client.RunRegistrarClient()

    } else {

        allowedOrigins := handlers.AllowedOrigins([]string{"*"})
        allowedMethods := handlers.AllowedMethods([]string{"GET", "POST", "DELETE", "PUT"})
        allowedHeaders := handlers.AllowedHeaders([]string{"X-Requested-With", "Content-Type"})

        // Define API routes
        //
        router := handlers.CORS(allowedOrigins, allowedMethods,
            allowedHeaders)(routes.NewRouter())

        // Start server
        //
        fmt.Printf("Registrar server started on :%d\n", common.GetListenPort())
        if err := http.ListenAndServe(fmt.Sprintf(":%d", common.GetListenPort()), router); err != nil {
            panic(err)
        }
    }
}
