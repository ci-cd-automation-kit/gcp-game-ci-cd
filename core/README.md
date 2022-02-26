# 0. Setup environment & prepare cert files for calling allocation api
```
gcloud auth login
gcloud config set project <PROJECT_ID>
cd core
./setup.sh #GKE cluster and tools provisioning, including agones, fleet, prometheus, weavescope, grafana
cd shell-scripts/demo/
./prepare_cert.sh
```

# 1. CI demo: cloud deploy for pipeline to deploy, skaffold for ci, cloud build for whole workflow
```
cd ../../agones/agones
export PROJECT_ID=<PROJECT_ID>
./create-cloud-deploy.sh

gcloud builds submit --config=cloudbuild-clouddeploy.yaml --substitutions=TAG_NAME=$(date +%y%m%d-%H%M%S),_PROJECT_ID=${PROJECT_ID}
kubectl get gameserver
nc -u <ip> <port>
hello
ACK V1: hello
```
See notes.txt in ci folder, you need to update main.go in examples/simple-game-server to see the updates after deploy to corresponding environment

# 2.  Demo for blue green and launch supertuxkart
```
export POD_NAME=$(kubectl get pods --namespace metrics -l "app.kubernetes.io/name=grafana,app.kubernetes.io/instance=grafana" -o jsonpath="{.items[0].metadata.name}")
     kubectl --namespace metrics port-forward $POD_NAME 3000
cd shell-scripts/demo/
./blue_green_launch.sh
```
Demo for supertuxkart game, view the dashboard in grafana

# 3. Demo multi players in supertuxkart
```
cd shell-scripts/demo/
./demo_supertuxkart.sh
```
will allocate one game server for two players, then allocate new game server for new two players

# 4. Allocate a gameserver and quilkin test
```
cd shell-scripts/demo/
./allocate_gameserver.sh
cd quilkin-test
python quilkin.py $IP $PORT
```

# 5. Appsheet for gaming Server scaling
```
kubectl port-forward -n weave "$(kubectl get -n weave pod --selector=weave-scope-component=app -o jsonpath='{.items..metadata.name}')" 4040
```
Scale gaming server in appsheet UI
