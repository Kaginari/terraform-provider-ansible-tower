
# ansible-tower_inventory_source

`ansible-tower_inventory_source` Inventory source 
* currently, supports [source](#source) : 
  * `scm` : Sourced from a Project
  * `custom` : Custom Script

## Example Usage with source `scm`

```hcl
resource "ansible-tower_organisation" "organisation" {
  name = "test organisation"
  description = "desc"
}
resource "ansible-tower_inventory" "inventory" {
  name = "test inventory"
  description = "test dsd"
  organisation_id = ansible-tower_organisation.organisation.id
}

resource "ansible-tower_inventory_source" "source" {
  name = "test source"
  inventory_id = ansible-tower_inventory.inventory.id
  source_project_id = 6
  source_path= ""
  source = "scm" ##(Optional : default 'scm')

}

```

## Example Usage with source `custom`

```hcl
resource "ansible-tower_organisation" "organisation" {
  name = "test organisation"
  description = "desc"
}
resource "ansible-tower_inventory" "inventory" {
  name = "test inventory"
  description = "test dsd"
  organisation_id = ansible-tower_organisation.organisation.id
}
resource "ansible-tower_inventory_script" "script" {
  name = "test inventory script script"
  description = "desc"
  organization_id = ansible-tower_organisation.organisation.id
  script = <<EOT
#!/usr/bin/env python
echo "hey"
EOT

}
resource "ansible-tower_inventory_source" "source_custom_script" {
  name = "test with custom script"
  inventory_id = ansible-tower_inventory.inventory.id
  source = "custom"
  source_script = ansible-tower_inventory_script.script.id
}
```

## Argument Reference

The following arguments are supported:

* `name` - Name of this inventory. (string, required)
* `description` - Optional description of this inventory source. (string, default="")
* `inventory_id` - (id, required)
*  <a id="source">`source`</a> -  (Optional) . (choice)
    * `scm` :  Sourced from a Project. (default)
    * `custom` : Custom Script.
* `source_script` - (Optional ,Needed only when source custom) (int , id, default=None)
* `credential_id` - Cloud credential to use for inventory updates. (integer, default=None)
* `source_path` - (string, default="")
* `source_project_id` - Project containing inventory file used as source. (id, default=None)
* `update_on_launch` - (boolean, default=False)
* `update_cache_timeout` - (integer, default=0)
* `overwrite_vars` -  Overwrite local variables from remote inventory source. (boolean, default=False)
* `overwrite` -   Overwrite local groups and hosts from remote inventory source. (boolean, default=False)
* `verbosity` - (Optional) . (choice)
  * `0` :  (WARNING)
  * `1` :  (INFO) (default)
  * `2` :  (DEBUG)
  
## Import

Ansible Tower Inventory source can be imported using the id, e.g. for an Inventory source with id : 1033

```sh
$ terraform import ansible-tower_inventory.example 1033
```**


