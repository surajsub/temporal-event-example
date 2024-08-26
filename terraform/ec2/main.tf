
provider "aws" {
  region = var.region
}


resource "aws_key_pair" "temporal" {
  key_name   = "temporal-key"
public_key= <SSH KEY>
}

resource "aws_instance" "temporal" {


  ami                    = var.amiId
  instance_type          = var.aws_instance_type
  subnet_id              = var.subnet_id
  key_name               = aws_key_pair.temporal.key_name
  vpc_security_group_ids = [var.sg_id]


  tags = {

    Name = var.instance_name
    created_by = "terraform"

  }

}


output "instance_id" {
  value = aws_instance.temporal.id
}

output "instance_public_ip" {
  value = aws_instance.temporal.public_ip
}
