terraform {
  required_providers {
    docker = {
      source  = "kreuzwerker/docker"
      version = "~> 3.0.1"
    }
  }
}

provider "docker" {}

# POSGRESQL CONFIGURATION
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

# MONGO CONFIGURATION
resource "docker_image" "mongo" {
  name         = "mongo"
  keep_locally = false
}

resource "docker_volume" "mongo-data" {
  name   = "mongo-data"
  driver = "local"
  driver_opts {
    type   = "none"
    o      = "bind"
    device = "/path/on/host"
  }
}

resource "docker_container" "mongo" {
  name  = "mongo"
  image = docker_image.mongo.latest

  ports {
    internal = 27017
    external = 27017
  }

  environment = [
    "MONGO_INITDB_ROOT_USERNAME=root",
    "MONGO_INITDB_ROOT_PASSWORD=root",
  ]

  volumes {
    container_path = "/data/db"
    volume_name    = docker_volume.mongo-data.name
    read_only      = false
    consistency    = "consistent"
  }
}

# MONGO-EXPRESS CONFIGURATION
resource "docker_image" "mongo-express" {
  name         = "mongo-express"
  depends_on   = [docker_image.mongo]
  keep_locally = false
}

resource "docker-container" "mongo-express" {
  name  = "mongo-express"
  image = docker_image.mongo-express.latest

  restart = "always"

  ports {
    internal = 8081
    external = 8081
  }

  environment = [
    "ME_CONFIG_MONGODB_ADMINUSERNAME=root",
    "ME_CONFIG_MONGODB_ADMINPASSWORD=root",
    "ME_CONFIG_MONGODB_URL=mongodb://root:root@mongo:27017/",
  ]
}

# CONSUL CONFIGURATION
resource "docker_image" "consul" {
  name         = "hashicorp/consul:latest"
  keep_locally = false
}

resource "docker_container" "consul" {
  name  = "consul-badger"
  image = docker_image.consul.latest

  restart = "always"

  ports {
    internal = 8500
    external = 8500
    protocol = "tcp"
  }

  ports {
    internal = 8600
    external = 8600
    protocol = "tcp"
  }

  ports {
    internal = 8600
    external = 8600
    protocol = "udp"
  }

  network_advanced {
    name   = "consul"
    driver = "bridge"
  }

  command = "agent -server -ui -node=server-1 -bootstrap-expect=1 -client=0.0.0.0"
}

resource "docker_network" "consul" {
  name   = "consul"
  driver = "bridge"
}

# Rabbit Message Queue Configuration
resource "docker_image" "rabbitmq" {
  name         = "rabbitmq:3-management"
  keep_locally = false
}

resource "docker_container" "rabbitmq" {
  name     = "rabbitmq"
  image    = docker_image.rabbitmq.name
  hostname = "my-rabbit"

  ports {
    internal = 15672
    external = 15672
  }

  ports {
    internal = 5672
    external = 5672
  }
}


