provider "google" {
  version = "~> 2.10"
}

provider "google-beta" {
  version = "~> 2.10"
}

variable "project_id" {
  default = ""
}

variable "region" {
  default = ""
}

variable "cluster_name" {
  default = "agones-terraform-example"
}

// Install latest version of agones
variable "agones_version" {
  default = ""
}

variable "machine_type" {
  default = "e2-standard-4"
}

// Note: This is the number of gameserver nodes. The Agones module will automatically create an additional
// two node pools with 1 node each for "agones-system" and "agones-metrics".
variable "node_count" {
  default = "4"
}
variable "zone" {
  default     = "us-west1-c"
  description = "The GCP zone to create the cluster in"
}

variable "network" {
  default     = "default"
  description = "The name of the VPC network to attach the cluster and firewall rule to"
}

variable "subnetwork" {
  default     = ""
  description = "The subnetwork to host the cluster in. Required field if network value isn't 'default'."
}

variable "log_level" {
  default = "info"
}

variable "feature_gates" {
  default = ""
}

variable "windows_node_count" {
  default = "0"
}

variable "windows_machine_type" {
  default = "e2-standard-4"
}

module "gke_cluster" {
  // ***************************************************************************************************
  // Update ?ref= to the agones release you are installing. For example, ?ref=release-1.8.0 corresponds
  // to Agones version 1.8.0
  // ***************************************************************************************************
  source = "./modules/gke"
  cluster_name = var.cluster_name
  project_id = var.project_id
  region  = var.region
}

module "helm_agones" {
  #source = "git::https://github.com/googleforgames/agones.git//install/terraform/modules/helm3/?ref=main"
  source = "./modules/agones"

  agones_version         = var.agones_version
  values_file            = ""
  feature_gates          = var.feature_gates
  host                   = module.gke_cluster.host
  token                  = module.gke_cluster.token
  cluster_ca_certificate = module.gke_cluster.cluster_ca_certificate
  log_level              = var.log_level
}

output "host" {
  value = module.gke_cluster.host
}
#output "token" {
#  value = module.gke_cluster.token
#}
output "cluster_ca_certificate" {
  value = module.gke_cluster.cluster_ca_certificate
}