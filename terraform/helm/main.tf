resource "helm_release" "busybox" {
  name       = "busybox"
  repository = "oci://localhost:5001"
  chart      = "busybox"

  values = [
    file("${path.cwd}/chart/busybox/values.yaml")
  ]
}
