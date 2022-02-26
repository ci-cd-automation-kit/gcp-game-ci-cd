blue_green_launch() {
  kubectl scale fleet supertuxkart-green --replicas=0
  kubectl scale fleet supertuxkart-blue --replicas=2
  sleep 20
  NAMESPACE=default # replace with any namespace
  EXTERNAL_IP=$(kubectl get services agones-allocator -n agones-system -o jsonpath='{.status.loadBalancer.ingress[0].ip}')
  KEY_FILE=client.key
  CERT_FILE=client.crt
  TLS_CA_FILE=ca.crt
  FLEET="supertuxkart-blue"
  #echo $EXTERNAL_IP
  allocated=$(curl -k -s --key ${KEY_FILE} --cert ${CERT_FILE} --cacert ${TLS_CA_FILE} -H "Content-Type: application/json" --data '{"namespace":"'${NAMESPACE}'", "requiredGameServerSelector": {"matchLabels":{"agones.dev/fleet":"'${FLEET}'"}}}' https://${EXTERNAL_IP}/gameserverallocation -XPOST)
  #echo $allocated
  ip=$(echo ${allocated} | jq -r '.address')
  echo "ip=$ip"
  port=$(echo ${allocated} | jq '.ports[0].port')
  echo "port=$port"
  open demo://${ip}:${port} -n
  open demo://${ip}:${port} -n
  sleep 60

  allocated=$(curl -k -s --key ${KEY_FILE} --cert ${CERT_FILE} --cacert ${TLS_CA_FILE} -H "Content-Type: application/json" --data '{"namespace":"'${NAMESPACE}'", "requiredGameServerSelector": {"matchLabels":{"agones.dev/fleet":"'${FLEET}'"}}}' https://${EXTERNAL_IP}/gameserverallocation -XPOST)
  #echo $allocated
  ip=$(echo ${allocated} | jq -r '.address')
  echo "ip=$ip"
  port=$(echo ${allocated} | jq '.ports[0].port')
  eecho "port=$port"
  open demo://${ip}:${port} -n
  open demo://${ip}:${port} -n
}
blue_green_launch
