
variable subnet_id {
  description = "This has to be the public id of the subnet that will be created in the vpc"

}
variable "nat_name" {
  description = "This is the name of the NAT"
}

variable "region" {
  description = "Region to create this resource in"
  default = "us-west-1"
}