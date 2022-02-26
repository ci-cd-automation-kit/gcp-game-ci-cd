# agones-prometheus
Agones native support Prometheus for metrics monitor, users can leverage this repo to setup the Prometheus and Grafana on demo environment.
![Grafana Sample Dashboard](/image/dashboard-sample.png)

## Prometheus Setup
Cloud Shell
```
helm repo add prometheus-community https://prometheus-community.github.io/helm-charts
helm repo update

helm upgrade --install --wait prom prometheus-community/prometheus --version 11.16.2 --namespace metrics \
    --set server.global.scrape_interval=30s \
    --set server.persistentVolume.enabled=true \
    --set server.persistentVolume.size=64Gi \
    -f ./prometheus.yaml
```

## Grafana Setup
Cloud Shell, replace the your-admin-password to grafana initial admin_password.
```
kubectl apply -f .grafana/

helm repo add grafana https://grafana.github.io/helm-charts
helm repo update

helm upgrade --install --wait grafana grafana/grafana --version=5.7.10 --namespace metrics \
  --set adminPassword=<your-admin-password> -f ./grafana.yaml
```

## Access Grafana
Local Shell
```
kubectl port-forward deployments/grafana 3000 -n metrics
```

Login using **username**:root, **password**:your-admin-password. Select the **Agones GameServers** Dashboard.
