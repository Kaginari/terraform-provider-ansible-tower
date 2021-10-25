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

resource "ansible-tower_inventory_script" "script" {
  name = "tf scriptssdsdddsds"
  description = "dsdsd"
  organization_id = ansible-tower_organisation.organisation.id
  script = <<EOT
#!/usr/bin/env python
echo "hey"
EOT

}
resource "ansible-tower_inventory_source" "source_custom_script" {
  name = "cxcdsfdsffffx"
  inventory_id = ansible-tower_inventory.inventory.id
  source = "custom"
  source_script = ansible-tower_inventory_script.script.id
}

resource "ansible-tower_project" "vault" {
  name                 = "test playbook"
  scm_type             = "git"
  scm_url              = "https://github.com/Kaginari/ansible-playbook-tower-test"
  scm_branch           = "main"
  scm_update_on_launch = true
  organisation_id      = ansible-tower_organisation.organisation.id
}
resource "ansible-tower_inventory_source" "source" {
  name = "cfdfdxcx"
  inventory_id = ansible-tower_inventory.inventory.id
  source_project_id = ansible-tower_project.vault.id
  source_path= ""
  source = "scm"

}
resource "ansible-tower_job_template" "template" {

  name           = "test-job-template"
  inventory_id   = ansible-tower_inventory.inventory.id
  project_id     = ansible-tower_project.vault.id
  playbook       = "main.yml"
  job_type       = "run"
  become_enabled = true
}
