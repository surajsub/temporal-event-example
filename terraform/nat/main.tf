# Resource: aws_eip
# https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/eip
provider "aws" {
  region = var.region
}

resource "aws_eip" "temporal_nat" {
  vpc = true

  tags = {
    Name = var.nat_name
  }
}

# Resource: aws_nat_gateway
# https://registry.terraform.io/providers/hashicorp/aws/latest/docs/resources/nat_gateway
resource "aws_nat_gateway" "temporal" {
  allocation_id = aws_eip.temporal_nat.id
  subnet_id = var.subnet_id


  tags = {
    Name = var.nat_name
    created_by = "terraform"
  }


}


output "nat_id" {
  value = aws_eip.temporal_nat.id
}

output nat_gateway_id {
  value= aws_nat_gateway.temporal.id
}

output "nat_gateway_allocation_id" {
  value = aws_nat_gateway.temporal.allocation_id
}