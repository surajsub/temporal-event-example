provider "aws" {
  region = var.region
}

resource "aws_route_table" "private" {
  vpc_id = var.vpc_id

  route {
      cidr_block                 = "0.0.0.0/0"
      nat_gateway_id             = var.nat_id
    }


  tags = {
    Name = "temporal-private"
  }
}

resource "aws_route_table" "public" {
  vpc_id = var.vpc_id

  route {
      cidr_block                 = "0.0.0.0/0"
      gateway_id                 = var.igw_id
    }


  tags = {
    Name = "temporal-public"
  }
}


// Since we are creating one public and one private subnet we will need two route associations

resource "aws_route_table_association" "public-temporal" {
  subnet_id      = var.subnet_id
  route_table_id = aws_route_table.public.id
  depends_on = [aws_route_table.public]
}



output "rt_public_id" {
  value = aws_route_table_association.public-temporal.id
}

output "rt_table_arn" {
  value = aws_route_table.public.arn
}