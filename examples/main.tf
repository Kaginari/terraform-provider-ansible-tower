terraform {
  required_version = ">= 0.13"

  required_providers {
    ansible-tower = {
      source = "registry.terraform.io/Kaginari/ansible-tower"
    }
  }
}
variable "username" {
  description = "ansible tower  username"
  default = "admin"
}
variable "password" {
  description = "ansible tower password"
  default = "password"
}

variable "host" {
  description = "ansible tower host"
  default = "http://127.0.0.1"
}
provider "ansible-tower" {
  tower_host = var.host
  tower_username = var.username
  tower_password = var.password
}

resource "ansible-tower_organisation" "organisation" {
  name = "test organisation"
  description = "desc"
}

resource "ansible-tower_inventory" "inventory" {
  name = "test inventory"
  description = "test dsd"
  organisation_id = ansible-tower_organisation.organisation.id
  kind = ""
  host_filter = ""
  variable {
    key = "sas"
    value = "sasaa"
  }
  variable {
    key = "monta"
    value = "[ a , b ]"
  }
}
//dGVzdGludmVudG9yeS4xMA==