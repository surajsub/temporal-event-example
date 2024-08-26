/* Resource file for internet gateway creation */
resource "aws_internet_gateway" "temporal" {

 vpc_id = var.vpc_id
 tags = {
   Name = var.igw_name
   created_by = "temporal"
 }

}

output "igw_id" {
  value = aws_internet_gateway.temporal.id
}

output "igw_arn" {
  value = aws_internet_gateway.temporal.arn
}
