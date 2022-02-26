curl -X POST -H "application/json" \
"https://cloudbuild.googleapis.com/v1/projects/cn-gaming-cicd/triggers/fan-trigger:webhook?key=AIzaSyBe45IJgmgE3Ssoekh77cl3Z1riBmG7N_I&secret=Ii_D0vWNiflXrzb4lGddK4u6ELBoyd1g" \
 -d '{"cluster_name":"fan-cluster", "region":"us-central1"}'

https://gitlab.endpoints.cn-gaming-cicd.cloud.goog/gaming-ci-cd-automation/core/-/blob/main/shell-scripts/demo.sh


1. Notice the public key & private key for ssh to gitlab
2. secret for webhook is a diffrent secret for cloud build trigger

./demo.sh --project-id cn-gaming-cicd --secret-name fan-demo-script-secret --trigger-name fan-demo-trigger \
--config ../cloud-build/cloud_build_fleet_configs.yaml --cluster-name fan-cluster --region us-central1 \
--region us-central1 --config-gs fleet_configs_simple.yaml --create-or-build gameserver


1. set up environment for gke & install angones

2. demo.sh ->
case1. set up environment
case2. cloud build trigger and create a fleet
case3. update fleet

3. quilkin demo
quilkin.py <ip> <port>

4. cd demo for upgrade fleet from v1 to v2

#expose weavescope
kubectl port-forward -n weave "$(kubectl get -n weave pod --selector=weave-scope-component=app -o jsonpath='{.items..metadata.name}')" 4040

#expose prometheus
export POD_NAME=$(kubectl get pods --namespace metrics -l "app=prometheus,component=server" -o jsonpath="{.items[0].metadata.name}")
  kubectl --namespace metrics port-forward $POD_NAME 9090

#expose Grafana
export POD_NAME=$(kubectl get pods --namespace metrics -l "app.kubernetes.io/name=grafana,app.kubernetes.io/instance=grafana" -o jsonpath="{.items[0].metadata.name}")
     kubectl --namespace metrics port-forward $POD_NAME 3000

5. ci demo
see notes.txt in ci folder
