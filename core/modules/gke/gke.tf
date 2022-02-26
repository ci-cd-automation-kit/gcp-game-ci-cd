data "google_client_config" "default" {}

# GKE cluster
resource "google_container_cluster" "primary" {
  name         = var.cluster_name
  location     = var.region
  #node_version = "1.18"
  #min_master_version = "1.18"

  remove_default_node_pool = true
  initial_node_count       = 1

  network    = google_compute_network.vpc.name
  subnetwork = google_compute_subnetwork.subnet.name

  # Enable Workload Identity for cluster
  #workload_identity_config {
  #  identity_namespace = "${var.project_id}.svc.id.goog"
  #}

}

# Separately Managed Node Pool
resource "google_container_node_pool" "primary_nodes" {
  name       = "${google_container_cluster.primary.name}-node-pool"
  location   = var.region
  cluster    = google_container_cluster.primary.name
  node_count = var.gke_num_nodes

  node_config {
    oauth_scopes = [
      "https://www.googleapis.com/auth/cloud-platform"
    ]

    labels = {
      environment = "dev",
      cluster     = var.cluster_name
    }

    machine_type = var.machine_type

    tags = [var.tags]

    # Enable Workload Identity for node pool
    # workload_metadata_config {
    #   node_metadata = "GKE_METADATA_SERVER"
    # }
    }

    #initial_node_count = 1

    autoscaling {
      min_node_count = 2
      max_node_count = 10
    }

    #management {
    #  auto_repair  = true
    #  auto_upgrade = true
    #}
}

output "kubernetes_cluster_name" {
  value       = google_container_cluster.primary.name
  description = "GKE Cluster Name"
}

output "cluster_ca_certificate" {
  value = base64decode(google_container_cluster.primary.master_auth.0.cluster_ca_certificate)
}

output "host" {
  value = "https://${google_container_cluster.primary.endpoint}"
}

output "token" {
  value = data.google_client_config.default.access_token
}
