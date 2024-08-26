provider "aws" {
  region = var.region
}

resource "aws_subnet" "private_temporal" {
  vpc_id            = var.vpc_id
  cidr_block        = var.subnet_size
  availability_zone = "us-west-1a"

  tags = {
    "Name"                            = var.subnet_name
  }
}


output "subnet_id" {
  value = aws_subnet.private_temporal.id
}

output "subnet_arn" {
  value = aws_subnet.private_temporal.arn
}