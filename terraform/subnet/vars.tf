
variable "vpc_id" {
  description = "The VPC to create this subnet in"
}

variable "subnet_size" {
  description = "Subnet_size"

}

variable "subnet_name" {
  description = "Private Subnet Zone"

}


variable "region" {
  default = "us-west-1"
}

