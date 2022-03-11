terraform {
    required_providers {
        google = {
            source = "hashicorp/google"
            version = "~> 4.13.0" 
        }
    }
}

provider "google" {
    project = "toflysim"
    region = "us-central1"
    zone = "us-central1-a"
}

resource "google_cloud_run_service" "publicApi" {
    name = "iothub-public-api"
    location = "us-central1"

    template {
        spec {
            containers {
                image = "us-docker.pkg.dev/cloudrun/container/hello"
            }
        }
    }

    traffic {
        percent = 100
        latest_revision = true
    }
}

data "google_iam_policy" "noauth" {
  binding {
    role = "roles/run.invoker"
    members = [
      "allUsers",
    ]
  }
}

resource "google_cloud_run_service_iam_policy" "noauth" {
  location    = google_cloud_run_service.publicApi.location
  project     = google_cloud_run_service.publicApi.project
  service     = google_cloud_run_service.publicApi.name

  policy_data = data.google_iam_policy.noauth.policy_data
}
