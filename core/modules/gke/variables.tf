variable "project_id" {
  default = ""
}

variable "cluster_name" {
  default = ""
}

variable "region" {
  default = ""
}

variable "tags" {
  default = "cluster-tags"
}

variable "gke_num_nodes" {
  default     = 1
  description = "number of gke nodes"
}

variable "machine_type" {
  default     = "e2-standard-4"
  description = "machine type of gke nodes"
}
