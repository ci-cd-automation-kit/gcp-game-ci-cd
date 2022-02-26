allocate() {
  init
  read -t 30 -p "Please input fleet name: " FLEET
  echo "Fleet name: $FLEET"

  kubectl scale fleet $FLEET --replicas=3
  sleep 10
  NAMESPACE=default # replace with any namespace
  EXTERNAL_IP=$(kubectl get services agones-allocator -n agones-system -o jsonpath='{.status.loadBalancer.ingress[0].ip}')
  KEY_FILE=client.key
  CERT_FILE=client.crt
  TLS_CA_FILE=ca.crt
  #echo $EXTERNAL_IP
  allocated=$(curl -k -s --key ${KEY_FILE} --cert ${CERT_FILE} --cacert ${TLS_CA_FILE} -H "Content-Type: application/json" --data '{"namespace":"'${NAMESPACE}'", "requiredGameServerSelector": {"matchLabels":{"agones.dev/fleet":"'${FLEET}'"}}}' https://${EXTERNAL_IP}/gameserverallocation -XPOST)
  #echo $allocated
  ip=$(echo ${allocated} | jq -r '.address')
  echo $ip
  port=$(echo ${allocated} | jq '.ports[0].port')
  echo $port
}

init() {
  for each in $(kubectl get fleet -o jsonpath="{.items[*].metadata.name}");
  do
    kubectl scale fleet $each --replicas=0
  done
  sleep 10
  for each in $(kubectl get gameserver -o jsonpath="{.items[*].metadata.name}");
  do
    kubectl delete gameserver $each
  done
}
allocate
