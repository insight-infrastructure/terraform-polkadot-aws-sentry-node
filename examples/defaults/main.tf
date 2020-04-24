variable "public_key" {}
variable "private_key_path" {}
variable "create_ansible" {}

module "network" {
  source = "github.com/insight-infrastructure/terraform-aws-polkadot-network.git?ref=master"
  sentry_enabled = true
  num_azs = 1
}

module "default" {
  source            = "../.."
  public_key        = var.public_key
  subnet_id         = module.network.public_subnets[0]
  security_group_id = module.network.sentry_security_group_id
  private_key_path  = var.private_key_path
  create_ansible    = var.create_ansible
}
