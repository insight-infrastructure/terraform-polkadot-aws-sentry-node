# terraform-polkadot-aws-sentry-node

WIP - Do not use

## Features

This module sets up

## Terraform versions

For Terraform v0.12.0+

## Usage

```
module "this" {
    source = "github.com/robc-io/terraform-polkadot-aws-sentry-node"

}
```

## Examples

- [simple](https://github.com/robc-io/terraform-polkadot-aws-sentry-node/tree/master/examples/simple)

## Known issues
No issue is creating limit on this module.

<!-- BEGINNING OF PRE-COMMIT-TERRAFORM DOCS HOOK -->
## Providers

| Name | Version |
|------|---------|
| aws | n/a |
| random | n/a |

## Inputs

| Name | Description | Type | Default | Required |
|------|-------------|------|---------|:-----:|
| create\_eip | Boolean to create elastic IP | `bool` | `false` | no |
| ebs\_volume\_size | EBS volume size | `string` | `0` | no |
| environment | The environment | `string` | `""` | no |
| instance\_type | Instance type | `string` | `"t2.micro"` | no |
| monitoring | Boolean for cloudwatch | `bool` | `false` | no |
| namespace | The namespace to deploy into | `string` | `""` | no |
| network\_name | The network name, ie kusama / mainnet | `string` | `""` | no |
| node\_name | Name of the node | `string` | `""` | no |
| owner | Owner of the infrastructure | `string` | `""` | no |
| public\_key\_path | The path to the public ssh key | `string` | `""` | no |
| root\_volume\_size | Root volume size | `string` | `0` | no |
| security\_group\_filter\_map | A map of of tags to query for a security group - ie {Name = "main"} | `map(string)` | n/a | yes |
| security\_group\_id | The id of the security group to run in | `string` | `""` | no |
| security\_group\_name | The name of the security group to query on | `string` | `""` | no |
| stage | The stage of the deployment | `string` | `""` | no |
| subnet\_id | The id of the subnet. | `string` | `""` | no |
| volume\_path | Volume path | `string` | `"/dev/xvdf"` | no |
| vpc\_filter\_map | A map of of tags to query for a vpc - ie {Name = "main"} | `map(string)` | n/a | yes |
| vpc\_id | The vpc id of the subnet | `string` | `""` | no |

## Outputs

No output.

<!-- END OF PRE-COMMIT-TERRAFORM DOCS HOOK -->

## Authors

Module managed by [robc-io](github.com/robc-io)

## Credits

- [Anton Babenko](https://github.com/antonbabenko)

## License

Apache 2 Licensed. See LICENSE for full details.