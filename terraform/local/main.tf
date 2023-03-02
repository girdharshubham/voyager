resource "null_resource" "unit" {
  provisioner "local-exec" {
    command = "echo kitty"
  }
}
