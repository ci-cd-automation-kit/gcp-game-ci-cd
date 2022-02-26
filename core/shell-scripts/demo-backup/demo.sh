#!/bin/bash
# chmod +x demo.sh
usage()
{
    printf "
case 1: create a gke trigger and trigger it\n
./demo.sh --project-id cn-gaming-cicd --secret-name opq-zyj --trigger-name opq-zyj --config terraform/tf-cloud-build.yaml --cluster-name opq-zyj --region us-central1 --create-or-build cluster
\n
case 2: create a gs config trigger and trigger it\n
./demo.sh --project-id cn-gaming-cicd --secret-name opq-zyj-gs --trigger-name opq-zyj-gs --config terraform/fleet-cloud-build.yaml --cluster-name opq-zyj --region us-central1 --create-or-build gameserver --config-gs fleet_configs_simple.yaml
\n
case 3: after you created the trigger, you want to trigger it to change the game server configuration\n
./demo.sh --project-id cn-gaming-cicd --secret-name opq-zyj-gs --trigger-name opq-zyj-gs --cluster-name opq-zyj --region us-central1 --create-or-build gameserver --config-gs fleet_configs_simple.yaml
\n
"
}

CREATE="NULL"
while [ ! -z "$1" ]; do
  case "$1" in
     --project-id)
         shift
         PROJECT_ID=$1 
         ;;
     --secret-name)
         shift
         SECRET_NAME=$1
         ;;
     --trigger-name)
         shift
         TRIGGER_NAME=$1
         ;;
     --config)
         shift
         INLINE_CONFIG=$1 # the local cloud build inline config file path. 
         ;;
     --cluster-name)
         shift
         CLUSTER_NAME=$1
         ;;
     --region)
         shift
         REGION=$1
         ;;
     --config-gs)
         shift
         CONFIG_GS=$1 # the fleet config file path in the GitLab repo, only defined in case 2 and case 3.
         ;;
     --create-or-build)
         shift
         CREATE=$1
         ;;
     --help|--h)
         usage
         ;;     
     *)
        echo "error: unknown parameters"
        echo "please use --help or -h to see the usage"
        ;;
  esac
shift
done

# create a secret for generating the webhook URL
create_secret() 
{
    # method 1: create a random string as the secret value
    SECRET=$(cat /dev/urandom | tr -dc 'a-zA-Z0-9' | fold -w 32 | head -n 1)
    gcloud secrets create $SECRET_NAME \
        --replication-policy="automatic"
    echo -n $SECRET | \
        gcloud secrets versions add $SECRET_NAME --data-file=-
    # method 2: add self-defined secret value from a text file
    # gcloud secrets versions add secret-id --data-file="/path/to/file.txt"

    # get the project number
    PROJECT_NUMBER=`gcloud projects describe $PROJECT_ID | grep 'projectNumber: ' | awk -F "'" '{print $2}'`
    # grant the service account the secretAccessor role to the secret
    gcloud projects add-iam-policy-binding $PROJECT_ID \
    --member serviceAccount:service-$PROJECT_NUMBER@gcp-sa-cloudbuild.iam.gserviceaccount.com \
    --role roles/secretmanager.secretAccessor
}

# case 1: create a gke trigger and trigger it
create_cluster() 
{
    # create a webhook trigger, specify the yaml config file, create 2 substitutions variables
    gcloud alpha builds triggers create webhook \
        --name=$TRIGGER_NAME\
        --secret=projects/$PROJECT_ID/secrets/$SECRET_NAME/versions/1 \
        --substitutions=_CLUSTER_NAME='$(body.cluster_name)',_REGION='$(body.region)' \
        --inline-config=$INLINE_CONFIG
    # get the API_KEY
    API_KEY=`curl -v --silent https://apikeys.googleapis.com/v1/projects/$PROJECT_ID/apiKeys?access_token=$(gcloud auth print-access-token) --stderr - | grep 'currentKey' | awk -F '"' '{print $4}'`
    # curl command to trigger cloud build 
    curl -X POST -H "application/json" "https://cloudbuild.googleapis.com/v1/projects/$PROJECT_ID/triggers/${TRIGGER_NAME}:webhook?key=${API_KEY}&secret=${SECRET}" -d '{"cluster_name":"'"$CLUSTER_NAME"'", "region":"'"$REGION"'"}'
}

# case 2: create a gs config trigger and trigger it
create_gameserver()
{
    # create a webhook trigger, specify the yaml config file, create 2 substitutions variables
    gcloud alpha builds triggers create webhook \
    --name=$TRIGGER_NAME\
    --secret=projects/$PROJECT_ID/secrets/$SECRET_NAME/versions/1 \
    --substitutions=_CLUSTER_NAME='$(body.cluster_name)',_REGION='$(body.region)',_CONFIG_FILE='$(body.config_file)' \
    --inline-config=$INLINE_CONFIG

    # get the API_KEY
    API_KEY=`curl -v --silent https://apikeys.googleapis.com/v1/projects/$PROJECT_ID/apiKeys?access_token=$(gcloud auth print-access-token) --stderr - | grep 'currentKey' | awk -F '"' '{print $4}'`
    # curl command to trigger cloud build 
    curl -X POST -H "application/json" "https://cloudbuild.googleapis.com/v1/projects/$PROJECT_ID/triggers/${TRIGGER_NAME}:webhook?key=${API_KEY}&secret=${SECRET}" -d '{"cluster_name":"'"$CLUSTER_NAME"'", "region":"'"$REGION"'", "config_file":"'"$CONFIG_GS"'"}'
}


# case 3: after you created the trigger, you want to trigger it to change the game server configuration.
config_gameserver() 
{
    # giving its trigger name and some parameters
    # get the secret value by giving its secret name
    SECRET=`gcloud secrets versions access 1 --secret=$SECRET_NAME`
    # get the API_KEY
    API_KEY=`curl -v --silent https://apikeys.googleapis.com/v1/projects/$PROJECT_ID/apiKeys?access_token=$(gcloud auth print-access-token) --stderr - | grep 'currentKey' | awk -F '"' '{print $4}'`
    # curl command to trigger cloud build 
    curl -X POST -H "application/json" "https://cloudbuild.googleapis.com/v1/projects/$PROJECT_ID/triggers/${TRIGGER_NAME}:webhook?key=${API_KEY}&secret=${SECRET}" -d '{"cluster_name":"'"$CLUSTER_NAME"'", "region":"'"$REGION"'", "config_file":"'"$CONFIG_GS"'"}'
}

case $CREATE in
    "cluster")       
        create_secret
        create_cluster
        ;;
    "gameserver")
        create_secret
        create_gameserver
        ;;            
    "build")       
        config_gameserver
        ;;
    "NULL")
        echo "please define --create-or-build"
        ;;
    *)
        echo "unknown trigger type"
        ;;
esac   
