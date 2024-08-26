variable "cidr_block" {
  description="CDIR Block for the VPC"
}



variable "name" {
  description = "Name of the VPC"
}

variable "region" {
  description = "Region to create this resource in"
  default = "us-west-1"
}