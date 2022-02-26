provider "google" {
  project = var.project_id
  region  = var.region
}


# VPC
resource "google_compute_network" "vpc" {
  name                    = "${var.cluster_name}-vpc"
  auto_create_subnetworks = "false"
}

# Subnet
resource "google_compute_subnetwork" "subnet" {
  name          = "${var.cluster_name}-subnet"
  region        = var.region
  network       = google_compute_network.vpc.name
  ip_cidr_range = "10.10.0.0/24"

}

# firewall rules
resource "google_compute_firewall" "rules" {
  project     = var.project_id
  name        = "${var.cluster_name}-firewall"
  network     = google_compute_network.vpc.name
  description = "Creates firewall rule targeting tagged instances"

  allow {
    protocol  = "udp"
    ports     = ["7000-8000"]
  }
  target_tags = [var.tags]
}

output "region" {
  value       = var.region
  description = "region"
}
