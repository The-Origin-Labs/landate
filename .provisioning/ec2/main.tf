terraform {
    required_providers {
      aws = {
        source = "hashicorp/aws",
        version = "~> 3.0"
      }
    }
}

provider "aws" {
    region = var.region
}

locals {
  extra_tag = "landate-tag"
}

resource "aws_instance" "instance_landate" {
    ami = var.ami
    instance_type = var.instance_type

    tags = {
      Name = var.instance_name
      ExtraTag = local.extra_tag
    }
}