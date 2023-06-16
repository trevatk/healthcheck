
job "health" {

    datacenters = ["us-mountain-1"]

    type = "service"

    group "trevatk" {
        count = 1

        network {

            port "http" {
                static = -1
                to = 8090
            }
        }

        service {
            name = "health-structx-io"
            tags = [
                "traefik.enable=true",
                "traefik.http.routers.health.rule=Host(`health.structx.io`)",
                "traefik.http.routers.health.tls=true",
                "treafik.http.routers.tls.certresolver=myresolver"
            ]
            port = "http"
            provider = "consul"

            check {
                name = "alive"
                type = "http"
                port = "http"
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
                HTTP_SERVER_PORT = 8090
            }

            resources {
                cpu    = 500
                memory = 256
            }
        }
    }
}