

**# ansible-tower_inventory_script

`ansible-tower_inventory_script` custom inventory scripts available in Tower


## Example Usage

```hcl
resource "ansible-tower_inventory_script" "script" {
  name = "test script"
  description = "description"
  organization_id = ansible-tower_organisation.organisation.id
  script = <<EOT
#!/usr/bin/env python
echo "hey"
EOT

}
```



## Argument Reference

The following arguments are supported:

* `name` - Name of this custom inventory script. (string, required)
* `description` - Optional description of this custom inventory script. (string, default="")
* `organization_id` - Organization owning this inventory script (id, required)
* `script` - (string, required)

## Import

Ansible Tower Inventory script can be imported using the id, e.g. for an Inventory script with id : 120

```sh  
$ terraform import ansible-tower_inventory_script 120  
```**