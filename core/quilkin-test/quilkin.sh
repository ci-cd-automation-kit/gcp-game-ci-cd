#!/bin/bash

# ./quilkin.sh --project-id cn-gaming-cicd --cluster-name opq-zyj --region us-central1

# get the ip address and the port of one of the game server
while [ ! -z "$1" ]; do
  case "$1" in
     --project-id)
         shift
         PROJECT_ID=$1
         ;;
     --cluster-name)
         shift
         CLUSTER_NAME=$1
         ;;
     --region)
         shift
         REGION=$1
         ;;
     *)
        echo "unknown parameters"
        ;;
  esac
shift
done

gcloud container clusters get-credentials $CLUSTER_NAME --region $REGION --project $PROJECT_ID

IP=`kubectl get gs | grep blue | awk 'NR==1 {print $3}'`
PORT=`kubectl get gs | grep blue | awk 'NR==1 {print $4}'`

python quilkin.py $IP $PORT
