#export PROJECT_ID=flius-demo
export CLUSTER_NAME=gaming-demo-1

gcloud config set project flius-demo
gcloud config set compute/region us-central1
gcloud config set compute/zone us-central1-a

PROJECT_ID=$(gcloud config get-value project)
GCP_REGION=$(gcloud config get-value compute/region)

gcloud services enable container.googleapis.com
gcloud services enable dns.googleapis.com

gcloud builds submit --substitutions=_PROJECT_ID=${PROJECT_ID},_CLUSTER_NAME=${CLUSTER_NAME},_GCP_REGION=${GCP_REGION}
