
# ansible-tower_credential_scm

`ansible-tower_credential_scm` Credentials scm are utilized by Tower for  synchronizing with inventory sources, and importing project content from a version control system.

## Example Usage with password

```hcl


resource "ansible-tower_credential_scm" "example" {
  organisation_id = ansible-tower_organisation.organisation.id
  name            = "acc-scm-credential"
  username        = "test"
  password        = "password"
}
```

## Example Usage with ssh key

```hcl
resource "ansible-tower_credential_scm" "credential" {
  organisation_id = ansible-tower_organisation.organisation.id
  name            = "acc-scm-credential"
  username        = "test"
  ssh_key_data    = file("${path.module}/files/id_rsa")
  ssh_key_unlock  = "passphrase"
}
```

## Argument Reference

The following arguments are supported:

* `name` - Name of this credential. (string, required)
* `organisation_id` - Organization containing this credential. (id, required)
* `description` - Optional description of this credential. (string, default="")
* `username` - (Optional) credential scm USERNAME (string,  default="")
* `password` - (Optional)  credential scm PASSWORD (string, default="")
* `ssh_key_data` - (Optional)  credential scm SSH_KEY  (string, default="")
* `ssh_key_unlock` - (Optional)  credential scm SSH_KEY_PASSPHRASE  (string, default="")



## Import

Ansible Tower Credential SCM can be imported using the id, e.g. for a Credential with id : 125

```sh
$ terraform import ansible-tower_credential_scm.example 125
```