

**# ansible-tower_credential_type

`ansible-tower_credential_type` custom credential type


## Example Usage

```hcl
resource "ansible-tower_credential_type" "example" {
  name           = "credential_type"
  input {
    id = "username"
    type = "string"
    label = "USERNAME"
  }
  input {
    id = "password"
    type = "string"
    label = "PASSWORD"
    secret = true
  }
  input {
    id = "url"
    type = "string"
    label = "URI"
    format = "url"
    multiline = false
  }
}
```



## Argument Reference

The following arguments are supported:

* `name` - Name of this custom inventory script. (string, required)
* `kind` - (Optional)  Kind of inventory being represented. (choice)
    * **cloud** :  (default)
    * **net** 
* `input` - (Optional) Credential type fields accept up to 5 

## Import

Ansible Tower credential type can be imported using the id, e.g. for a credential type  with id : 120

```sh  
$ terraform import ansible-tower_credential_type 120  
```**