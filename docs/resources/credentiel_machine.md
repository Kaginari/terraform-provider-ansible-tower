
# ansible-tower_credential_machine

`ansible-tower_credential_machine` Credentials machine Credentials are utilized by Tower for authentication when launching Jobs against machines.

## Example Usage with password

```hcl


resource "ansible-tower_credential_machine" "credential" {
  organisation_id     = ansible-tower_organisation.organisation.id
  name                = "acc-machine-credential"
  username            = "test"
  password            = "pwd"
}
```

## Example Usage with ssh key

```hcl
resource "ansible-tower_credential_machine" "credential" {
  organisation_id     = ansible-tower_organisation.organisation.id
  name                = "acc-machine-credential"
  username            = "test"
  ssh_key_data        = file("${path.module}/files/id_rsa")
  ssh_public_key_data = file("${path.module}/files/id_rsa.pub")
  ssh_key_unlock      = "test"  # if private key is encrypted
}

```

## Argument Reference

The following arguments are supported:

* `name` - Name of this credential. (string, required)
* `organisation_id` - Organization containing this credential. (id, required)
* `description` - Optional description of this credential. (string, default="")
* `username` - (Optional) credential machine USERNAME (string,  default="")
* `password` - (Optional)  credential machine PASSWORD (string, default="")
* `ssh_key_data` - (Optional)  credential machine SSH_KEY  (string, default="")
* `ssh_public_key_data` - (Optional)  credential machine SIGNED SSH CERTIFICATE  (string, default="")
* `ssh_key_unlock` - (Optional)  credential machine SSH_KEY_PASSPHRASE  (string, default="")



## Import

Ansible Tower Credential machine can be imported using the id, e.g. for a Credential with id : 125

```sh
$ terraform import ansible-tower_credential_machine.example 125
```