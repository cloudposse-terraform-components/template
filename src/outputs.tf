output "mock" {
  description = "Mock output example for the Cloud Posse Terraform component template test"
  value       = local.enabled ? "hello ${basename(abspath(path.module))}" : ""
}
