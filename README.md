# Bullwark Service Registrar

This project is a simple service which provides REST endpoints for client services to use
to register with the Bullwark service discovery implementation. It is intended to be small, fast,
and light.

## Prerequisites

- Go 1.10+ OR Docker


## Build

With native Go:

```bash
make build
```

With Docker:

```bash
make docker-build
```

## Run (Server Mode)

```bash
./bin/bullward-registrar
```

You can then access the registrar at http://localhost:8000.

## Run (CLI Mode)

This binary can also be used as a client to sideload a service.

To do so, (assuming you're running in the Vagrant configuration) run: 

```bash
./bin/bullwark-registrar --client=true --generate-service-ids --config=src/examples/client.yml --registrar-host=192.168.50.2

```

The above command in the standard Vagrant setup illustrates how to service on various paths; the end result of the above
command is that traefik is configured to proxy both itself and Consul out to port 80.

Consul: 192.168.50.2/ui/
Traefik: 192.168.50.2/dashboard/
