
job "health" {

    datacenters = ["*"]

    type = "service"

    group "trevatk" {
        count = 1

        network {

            mode = "bridge"

            port "http" {
                static = -1
                to = 8090
            }
        }

        service {
            name = "health-structx-io"
            tags = [
                "traefik.enable=true"
            ]
            port = "http"
            provider = "consul"
        }

        task "server" {

            driver = "docker"

            config {
                image = "trevatk/healthcheck:v0.0.1"
                ports = ["http"]
            }

            env {
                HTTP_SERVER_PORT = 8090
            }

            resources {
                cpu    = 500
                memory = 256
            }
        }
    }
}