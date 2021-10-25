
# ansible-tower_project



## Example Usage

```hcl
data "ansible-tower_organization" "default" {
  name = "Default"
}

resource "ansible-tower_project" "base_service_config" {
  name                 = "vault cluster playbook"
  scm_type             = "git"
  scm_url              = "https://gitlab.com/nt-factory/2021/admin/vault"
  scm_branch           = "feature/cluster-playbook"
  scm_update_on_launch = true
  organisation_id      = data.ansible-tower_organization.default.id
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of this project
* `organisation_id` - (Required) Numeric ID of the project organization
* `scm_type` - (Required) One of "" (manual), git, hg, svn
* `description` - (Optional) Optional description of this project.
* `local_path` - (Optional) Local path (relative to PROJECTS_ROOT) containing playbooks and related files for this project.
* `scm_branch` - (Optional) Specific branch, tag or commit to checkout.
* `scm_clean` - (Optional)
* `scm_credential_id` - (Optional) Numeric ID of the scm used credential
* `scm_delete_on_update` - (Optional)
* `scm_update_cache_timeout` - (Optional)
* `scm_update_on_launch` - (Optional)
* `scm_url` - (Optional) 

## Import

Ansible Tower project can be imported using the id, e.g. for an Inventory source with id : 50

```sh
$ terraform import ansible-tower_project.example 50
```**
