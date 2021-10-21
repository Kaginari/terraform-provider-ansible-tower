module github.com/Kaginari/terraform-provider-ansible-tower

go 1.15

require (
	github.com/hashicorp/terraform-plugin-sdk/v2 v2.8.0
	github.com/mitchellh/mapstructure v1.1.2
	github.com/mrcrilly/goawx v0.1.4
)

replace github.com/mrcrilly/goawx => github.com/ITMonta/goawx v0.1.6
