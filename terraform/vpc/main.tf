provider "aws" {
  region = var.region
}

resource "aws_vpc" "temporal" {
  cidr_block = var.cidr_block

 tags = {
   Name = var.name
   created_by = "terraform"
 }
}

output "vpc_id" {
  value = aws_vpc.temporal.id
}

output "vpc_cidr_block" {
  value = aws_vpc.temporal.cidr_block
}