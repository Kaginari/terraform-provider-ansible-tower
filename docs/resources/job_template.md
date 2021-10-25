# ansible-tower_job_template

A job template is a definition and set of parameters for running an Ansible job. Job templates are useful to execute the same job many times. Job templates also encourage the reuse of Ansible playbook content and collaboration between teams. While the REST API allows for the execution of jobs directly, Tower requires that you first create a job template.

## Example Usage
```hcl


resource "ansible-tower_job_template" "template_example" {
    name           = "test-job-template"
    job_type       = "run"
    inventory_id   = ansible-tower_inventory.example_inventory.id
    project_id     = ansible-tower_project.example_project.id
    playbook       = "main.yml"
    become_enabled = true
}
```

## Argument Reference

The following arguments are supported:

* `name` - (Required) Name of this workflow job template. (string, required)
* `playbook` - (Required) (string, default=``)
* `description` - (Optional)  description of this job template. (string, default="")
* `organisation_id` - (Optional) The organization used to determine access to this template. (id, default=``)
* `inventory_id` - (Required) Inventory applied as a prompt, assuming job template prompts for inventory.
* `project_id` - (Required) (id, default=``)
* `job_type` - (Optional) . (choice)
    * `run` :  Run (default)
    * `check` :   Check
* `allow_simultaneous` - (Optional)
* `ask_inventory_on_launch` - (Optional)
* `ask_limit_on_launch` - (Optional)
* `ask_scm_branch_on_launch` - (Optional)
* `ask_variables_on_launch` - (Optional)
* `limit` - (Optional)
* `survey_enabled` - (Optional)
* `variables` - (Optional)
* `webhook_credential` - (Optional)
* `webhook_service` - (Optional) 
* `verbosity` - (Optional) . (choice)
    * `0` :  (Normal) (default)
    * `1` :  (Verbose)
    * `2` :  (More Verbose)
    * `3` :  (DEBUG)
    * `4` :  (Connection Debug)
    * `5` :  (DEBUG) 

## Import

Ansible Tower Job Template can be imported using the id, e.g. for a Job Template with id : 10

```sh
$ terraform import ansible-tower_job_template.example 10
```**
