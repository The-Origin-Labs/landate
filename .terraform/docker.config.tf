terraform {
  required_providers {
    docker = {
      source  = "kreuzwerker/docker"
      version = "~> 3.0.1"
    }
  }
}

provider "docker" {}

resource "docker_image" "postgres" {
  name         = "postgres:13-alpine"
  keep_locally = false
}

resource "docker_volume" "postgres_data" {
  name   = "postgres-data"
  driver = "local"
  driver_opts {
    type   = "none"
    o      = "bind"
    device = "/path/on/host"
  }
}

resource "docker_container" "postgres" {
  name  = "postgres"
  image = docker_image.postgres.latest

  restart = "always"
  ports {
    internal = 5432
    external = 5432
  }

  environment = [
    "POSTGRES_USER=postgres",
    "POSTGRES_PASSWORD=dev",
    "POSTGRES_DB=heimdb",
  ]

  volumes {
    container_path = "/var/lib/postgresql/data"
    volume_name    = docker_volume.postgres_data.name
    read_only      = false
    consistency    = "consistent"
  }
}


resource "docker_image" "mongo" {
  name         = "mongo"
  keep_locally = false
}

resource "docker_image" "mongo-express" {
  name         = "mongo-express"
  depends_on   = [docker_image.mongo]
  keep_locally = false
}

resource "docker_image" "consul" {
  name         = "hashicorp/consul:latest"
  keep_locally = false
}

resource "docker_image" "rabbitmq" {
  name         = "rabbitmq:3-management"
  keep_locally = false
}


