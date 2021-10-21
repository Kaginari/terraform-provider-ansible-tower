# Terraform provider Ansible-tower

![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/Kaginari/terraform-provider-ansible-tower?logo=go&style=flat-square)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/Kaginari/terraform-provider-ansible-tower?logo=git&style=flat-square)
![GitHub](https://img.shields.io/github/license/Kaginari/terraform-provider-ansible-tower?color=yellow&style=flat-square)
![GitHub Workflow Status](https://img.shields.io/github/workflow/status/Kaginari/terraform-provider-ansible-tower/golangci?logo=github&style=flat-square)
![GitHub issues](https://img.shields.io/github/issues/Kaginari/terraform-provider-ansible-tower?logo=github&style=flat-square)


This repository is a [Terraform](https://www.terraform.io) Provider for Ansible tower (awx)  
 
### Requirements

- [Terraform](https://www.terraform.io/downloads.html) >= 0.13
- [Go](https://golang.org/doc/install) >= 1.15

### Installation

1. Clone the repository
1. Enter the repository directory
1. Build the provider using the `make install` command:

````bash
git clone https://github.com/Kaginari/terraform-provider-mongodb
cd terraform-provider-ansible-tower
make install
````

### To test locally

**1.1: launch awx tower**


````bash
cd example
docker-compose up -d
````

*follow the instruction in this link*

https://debugthis.dev/posts/2020/04/setting-up-ansible-awx-using-a-docker-environment-part-2-the-docker-compose-approach/


**1.4 :  user in tower**

* default user : admin
* default password : password

**2: Build the provider**

follow the [Installation](#Installation)

**3: Use the provider**

````bash
cd example
make apply
````
