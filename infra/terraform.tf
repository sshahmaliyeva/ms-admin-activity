variable "DOCKER_IMG_TAG" {}
variable "CIRCLE_BUILD_NUM" {}
variable "environment" {}

terraform {
  backend "s3" {
    bucket = "digital.pashabank.eks.terraform.configs"
    key    = "ms-admin-activity.tfstate"
    region = "us-east-1"
  }
}

module "microservice" {
  source              = ""
  service_name        = "ms-admin-activity"
  service_version     = "1"
  service_image       = "ictcontact/ms-admin-activity:${var.DOCKER_IMG_TAG}"
  service_replicas    = "2"
  cluster_environment = "${var.environment}"
  pod_version         = "${var.CIRCLE_BUILD_NUM}"
}
