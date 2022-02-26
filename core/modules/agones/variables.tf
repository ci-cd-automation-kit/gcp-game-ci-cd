#Helm variables

variable "chart" {
  default = "agones"
}

variable "force_update" {
  default = "true"
}

variable "agones_version" {
  default = ""
}

variable "udp_expose" {
  default = "true"
}

variable "log_level" {
  default = "info"
}

variable "feature_gates" {
  default = ""
}

variable "host" {}

variable "token" {}

variable "cluster_ca_certificate" {}

variable "crd_cleanup" {
  default = "true"
}

variable "image_registry" {
  default = "gcr.io/agones-images"
}

variable "pull_policy" {
  default = "IfNotPresent"
}

variable "always_pull_sidecar" {
  default = "false"
}

variable "image_pull_secret" {
  default = ""
}

variable "ping_service_type" {
  default = "LoadBalancer"
}

variable "values_file" {
  default = ""
}

variable "gameserver_minPort" {
  default = "7000"
}

variable "gameserver_maxPort" {
  default = "8000"
}

variable "gameserver_namespaces" {
  default = ["default"]
  type    = list(string)
}