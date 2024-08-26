variable vpc_id {
  description = "VPC ID"
}

variable "nat_id" {
  description = "NAT Id to be used for the route"
}

variable "igw_id" {
  description = "Internet Gateway"
}

variable "subnet_id" {
  description = "THe public subnet id created for this resource"
}



variable "region" {
  description = "The region of the SG"
  default ="us-west-1"
}