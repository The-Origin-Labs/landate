output "instance_1_ip_addr" {
  value = aws_instance.instance_landate.public_ip
}

output "instance_2_ip_addr" {
  value = aws_instance.instance_landate.private_ip
}
