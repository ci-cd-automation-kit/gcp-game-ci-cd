#!/bin/bash
NAMESPACE=default # replace with any namespace
EXTERNAL_IP=$(kubectl get services agones-allocator -n agones-system -o jsonpath='{.status.loadBalancer.ingress[0].ip}')
KEY_FILE=client.key
CERT_FILE=client.crt
TLS_CA_FILE=ca.crt

# allocator-client.default secret is created only when using helm installation. Otherwise generate the client certificate and replace the following.
# In case of MacOS replace "base64 -d" with "base64 -D"
kubectl get secret allocator-client.default -n "${NAMESPACE}" -ojsonpath="{.data.tls\.crt}" | base64 -d > "${CERT_FILE}"
kubectl get secret allocator-client.default -n "${NAMESPACE}" -ojsonpath="{.data.tls\.key}" | base64 -d > "${KEY_FILE}"
kubectl get secret allocator-tls-ca -n agones-system -ojsonpath="{.data.tls-ca\.crt}" | base64 -d > "${TLS_CA_FILE}"

call_gs()
{
    # 30 players keep calling the gs
    for i in {1..30}
    do
        printf "player $i calls the game server."
        curl -k --key ${KEY_FILE} --cert ${CERT_FILE} --cacert ${TLS_CA_FILE} -H "Content-Type: application/json" --data '{"namespace":"'${NAMESPACE}'"}' https://${EXTERNAL_IP}/gameserverallocation -XPOST | awk -F "," '{print $1}' >> log.txt
        echo $'\n' >> log.txt
        sleep 1
    done
}

update()
{
    echo "check the gs: \n"
    # check the gs
    kubectl get gs

    sleep 2
    # change replica of fleet v2 to 10
    echo "create fleet v2, 30 game servers"
    echo "create fleet v2, 30 game servers" >> log.txt
    echo $'\n' >> log.txt
    kubectl apply -f ../modules/agones/fleet_configs_simple_v2.yaml

    echo "---------------\n"
    echo "sleep 20 sec \n"
    sleep 20
    echo "---------------\n"

    echo "check the gs: \n"
    # check the gs
    kubectl get gs

    # change replica of fleet v1 to 0
    echo "change replica of fleet v1 to 0"
    echo "change replica of fleet v1 to 0" >> log.txt
    echo $'\n' >> log.txt
    kubectl scale fleet simple-game-server --replicas=0

    echo "---------------"
    echo "sleep 5 sec \n"
    sleep 5
    echo "---------------"
    echo "check the gs: "
    # check the gs
    kubectl get gs
}

call_gs & update