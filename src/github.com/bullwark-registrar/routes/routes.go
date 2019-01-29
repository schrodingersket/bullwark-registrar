package routes

import (
    "encoding/json"
    "net/http"

    "github.com/gorilla/mux"

    "github.com/bullwark-registrar/common"
    "github.com/bullwark-registrar/common/cli"
    "github.com/bullwark-registrar/handlers/app"
    "github.com/bullwark-registrar/handlers/middlewares"
)

var routes = middlewares.Routes{
    middlewares.Route{
        Name: "Register",
        Method: "POST",
        Pattern: "/register",
        HandlerFunc: app.Register,
        Middlewares: []middlewares.Middleware{
            middlewares.Json,
        },
    },
    middlewares.Route{
        Name: "Register",
        Method: "POST",
        Pattern: "/register/bulk",
        HandlerFunc: app.RegisterBulk,
        Middlewares: []middlewares.Middleware{
            middlewares.Json,
        },
    },
}

// Defines a new router for the application based on the above routes.
//
func NewRouter() *mux.Router {

    var coreConfig = common.Configs[cli.CoreConfigType].(cli.CoreConfig)
    parentRouter := mux.NewRouter()
    router := parentRouter.PathPrefix(*coreConfig.BaseURL).Subrouter()

    // Add not found handler
    //
    router.NotFoundHandler = http.HandlerFunc(func (w http.ResponseWriter, req *http.Request) {

        w.Header()["Content-Type"] = []string{"application/json; charset=utf-8"}

        w.WriteHeader(http.StatusNotFound)
        _ = json.NewEncoder(w).Encode(map[string]interface{}{
            "message": "Not Found",
        })
    })


    for _, route := range routes {

        // Every handler gets a logger
        //
        var handler = middlewares.Logger(route.HandlerFunc, route)

        //var handler = route.HandlerFunc;

        // Apply middlewares in the order they show up in the slice
        //
        for i:= len(route.Middlewares) - 1; i>= 0; i-- {
          handler = route.Middlewares[i](handler, route)
        }

        // Add handler
        //
        router.
            Methods(route.Method).
            Path(route.Pattern).
            Name(route.Name).
            Handler(handler)
    }

    return router
}
