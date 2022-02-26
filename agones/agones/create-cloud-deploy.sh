PROJECT_ID=flius-vpc-2
sed -i'' -e "s/PROJECT_ID/${PROJECT_ID}/" ./examples/simple-game-server/clouddeploy.yaml
gcloud beta deploy apply --file ./examples/simple-game-server/clouddeploy.yaml --region=us-central1 --project=${PROJECT_ID}
sed -i'' -e "s/${PROJECT_ID}/PROJECT_ID/" ./examples/simple-game-server/clouddeploy.yaml
