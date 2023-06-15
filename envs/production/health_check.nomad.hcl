
job "health" {

    datacenters = ["mountain"]

    type = "service"

    group "trevatk" {
        count = 1

        network {

            mode = "bridge"

            port "http" {
                to = -1
            }
        }

        service {
            name = "health-structx-io"
            tags = [
                "traefik.enable=true",
                "traefik.http.routers.health.rule=Host(`health.structx.io`)",
            ]
            port = "http"

            connect {
                sidecar_service {}
            }

            check {
                name = "alive"
                type = "http"
                path = "/health"
                interval = "1m"
                timeout = "10s"
            }
        }

        task "server" {

            driver = "docker"

            config {
                image = "trevatk/healthcheck:v0.0.1"
                ports = ["http"]
            }

            env {
                HTTP_SERVER_PORT = "${NOMAD_PORT_http}"
            }

            resources {
                cpu    = 500
                memory = 256
            }
        }
    }
}