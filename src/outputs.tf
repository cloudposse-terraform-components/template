output "mock" {
  description = "Mock output example for the Cloud Posse Terraform component template test 9"
  value       = local.enabled ? "hello ${basename(abspath(path.module))}" : ""
}
