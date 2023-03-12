terraform {
  required_providers {
    executor = {
      version = "~> 1.0.0"
      source  = "girdharshubham.com/execprovider/exec"
    }
  }
}

resource "executor" "hello_world" {
  command    = "Hello World"
  entrypoint = "echo"
}