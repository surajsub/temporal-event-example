variable "vpc_id" {
  description = "The vpc id to create this sg in"
}

variable "region" {
  description = "The region of the SG"
  default ="us-west-1"
}