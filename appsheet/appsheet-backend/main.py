from flask import escape
from kubernetes import client, config
# from google.auth import compute_engine
# from google.cloud.container_v1 import ClusterManagerClient
from os import path
import requests
import yaml
import socket
import time
import base64
import json
import os

def get_access_token():
    METADATA_URL = 'http://metadata.google.internal/computeMetadata/v1/'
    METADATA_HEADERS = {'Metadata-Flavor': 'Google'}
    SERVICE_ACCOUNT = 'default'

    url = '{}instance/service-accounts/{}/token'.format(METADATA_URL, SERVICE_ACCOUNT)

    # Request an access token from the metadata server.
    r = requests.get(url, headers=METADATA_HEADERS)
    r.raise_for_status()

    # Extract the access token from the response.
    access_token = r.json()['access_token']

    print(access_token)

    return access_token

# def Init_K8s_Client(access_token):
# configuration = client.Configuration()
# configuration.host = f"https://34.82.213.24:443"
# configuration.verify_ssl = False
# configuration.api_key = {"authorization": "Bearer " + access_token}
# client.Configuration.set_default(configuration)

def AddGameServer(access_token):
    configuration = client.Configuration()
    k8s_host = os.environ.get('k8s_endpoint', 'Specified environment variable is not set.')
    configuration.host = "https://" + k8s_host + ":443"
    configuration.verify_ssl = False
    configuration.api_key = {"authorization": "Bearer " + access_token}
    client.Configuration.set_default(configuration)
    
    # load gameserverallocate template("./localpackage/gameserverallocate.yaml")
    with open(path.join(path.dirname(__file__), "./localpackage/gameserverallocate.yaml")) as f:
        gameserverallocate_obj = yaml.safe_load(f)

    crdclient = client.CustomObjectsApi()

    group = 'allocation.agones.dev' # str | the custom resource's group
    version = 'v1' # str | the custom resource's version
    namespace = 'default' # str | The custom resource's namespace
    plural = 'gameserverallocations' # str | the custom resource's plural name. For TPRs this would be lowercase plural kind.
    # name = 'demo-allocate' # str | the custom object's name
    # body = None # object | The JSON schema of the Resource to patch.

    response = crdclient.create_namespaced_custom_object(group, version, namespace, plural, body=gameserverallocate_obj)

    print(response)
    
    if response["status"]["state"] == "Allocated":
        gameserver_address = response["status"]["address"]
        gameserver_port = response["status"]["ports"][0]["port"]
        gameserver_name = response["metadata"]["name"]
        if gameserver_name.split("-")[1] == "blue":
            gameserver_name = "Blue"
        else:
            gameserver_name = "Green"
    else :
        gameserver_address = "0.0.0.0"
        gameserver_port = 0
        gameserver_name = ""
    
    return gameserver_address, gameserver_port, gameserver_name

def UpdateFleetReplica(access_token, fleet_name, replica):
    configuration = client.Configuration()
    k8s_host = os.environ.get('k8s_endpoint', 'Specified environment variable is not set.')
    configuration.host = "https://" + k8s_host + ":443"
    configuration.verify_ssl = False
    configuration.api_key = {"authorization": "Bearer " + access_token}
    client.Configuration.set_default(configuration)

    with open(path.join(path.dirname(__file__), "./localpackage/fleet.yaml")) as f:
        fleetobj = yaml.safe_load(f)
        fleetobj["spec"]["replicas"] = replica
        fleetobj["metadata"]["name"] = fleet_name
        print(fleetobj)
    
    group = 'agones.dev' # str | the custom resource's group
    version = 'v1' # str | the custom resource's version
    namespace = 'default' # str | The custom resource's namespace
    plural = 'fleets' # str | the custom resource's plural name. For TPRs this would be lowercase plural kind.
    name = fleet_name # str | the custom object's name
    body = None # object | The JSON schema of the Resource to patch.

    crdclient = client.CustomObjectsApi()

    response = crdclient.patch_namespaced_custom_object(group, version, namespace, plural, name, fleetobj)

    print(response)

def ConnGameServer(gameserver_address, gameserver_port, updatemode):
    if updatemode == "Add":
        message = b"Hello"
        s = socket.socket(socket.AF_INET, socket.SOCK_DGRAM, 0)
        s.sendto(message, (gameserver_address, gameserver_port))
        time.sleep(3)
        response_data, response_address = s.recvfrom(4096)
        s.close()
        return response_data.decode('utf-8')
    else :
        message = b"EXIT"
        s = socket.socket(socket.AF_INET, socket.SOCK_DGRAM, 0)
        s.sendto(message, (gameserver_address, gameserver_port))
        time.sleep(3)
        response_data, response_address = s.recvfrom(4096)
        s.close()
        print(response_data.decode('utf-8'))
        return response_data.decode('utf-8')

def UpdateAppsheetGameServerStatusTable(game_server_name, fleet_name, server_response, updatemode):
    url = "https://api.appsheet.com/api/v2/apps/" + os.environ.get('appsheet_app_id', 'Specified environment variable is not set.') + "/tables/gameserverstatus/Action"
    appsheet_application_accesskey = os.environ.get('appsheet_accesskey', 'Specified environment variable is not set.')
    
    headers = {
    'ApplicationAccessKey': appsheet_application_accesskey,
    'Content-Type': 'application/json'
    }

    if updatemode == "Add":
        payload = json.dumps({
            "Action": "Edit",
            "Properties": {
                "Locale": "en-US",
                "RunAsUserEmail": "user2@mydomain.com"
            },
            "Rows": [
                {
                "GameServerName": game_server_name,
                "FleetName": fleet_name,
                "ServerResponse": server_response
                }
            ]
            })
    else:
        payload = json.dumps({
        "Action": "Delete",
        "Properties": {
            "Locale": "en-US",
            "RunAsUserEmail": "user2@mydomain.com"
        },
        "Rows": [
            {
            "GameServerName": game_server_name
            }
        ]
        })

    response = requests.request("POST", url, headers=headers, data=payload)

    print(response.text)        



def demo_appsheet_backend(request):
    """HTTP Cloud Function.
    Args:
        request (flask.Request): The request object.
        <https://flask.palletsprojects.com/en/1.1.x/api/#incoming-request-data>
    Returns:
        The response text, or any set of values that can be turned into a
        Response object using `make_response`
        <https://flask.palletsprojects.com/en/1.1.x/api/#flask.make_response>.
    """
    request_json = request.get_json(silent=True)
    print(request_json)
    request_args = request.args

    # # config.load_kube_config("./localpackage/config")
    # with open(path.join(path.dirname(__file__), "./localpackage/gameserverallocate.yaml")) as f:
    #     fleetobj = yaml.safe_load(f)
    #     fleetobj["spec"]["replicas"] = 10
    #     print(fleetobj["spec"]["replicas"])
    
    # crdclient = client.CustomObjectsApi()

    # group = 'allocation.agones.dev' # str | the custom resource's group
    # version = 'v1' # str | the custom resource's version
    # namespace = 'default' # str | The custom resource's namespace
    # plural = 'gameserverallocations' # str | the custom resource's plural name. For TPRs this would be lowercase plural kind.
    # name = 'demo-allocate' # str | the custom object's name
    # body = None # object | The JSON schema of the Resource to patch.

    # response = crdclient.create_namespaced_custom_object(group, version, namespace, plural, body=fleetobj)

    # print(response)

    if request_json and 'UpdateMode' in request_json and request_json["UpdateMode"] == "Add" and request_json["TableName"] == "gameserverstatus":
        access_token = get_access_token()
        # Init_K8s_Client(access_token)
        gameserver_address, gameserver_port, gameserver_name = AddGameServer(access_token)
        response_data = ConnGameServer(gameserver_address, gameserver_port, "Add")
        server_response = gameserver_address + ":" + str(gameserver_port) + ":" + response_data
        UpdateAppsheetGameServerStatusTable(request_json["Data"]["GameServerName"], gameserver_name, server_response, "Add")
        return ("Done", 200)
    elif request_json and 'UpdateMode' in request_json and request_json["UpdateMode"] == "Delete" and request_json["TableName"] == "gameserverstatus":
        access_token = get_access_token()
        # Init_K8s_Client(access_token)
        gameserver_info = request_json["Data"]["ServerResponse"]
        print(gameserver_info)
        gameserver_address = gameserver_info.split(":")[0]
        gameserver_port = gameserver_info.split(":")[1]
        response_data = ConnGameServer(gameserver_address, int(gameserver_port), "Delete")
        # server_response = gameserver_address + ":" + str(gameserver_port) + ":" + response_data
        # UpdateAppsheetGameServerStatusTable(request_json["Data"]["GameServerName"], request_json["Data"]["FleetName"], server_response, "Delete")
        return ("Done", 200)
    elif request_json and 'UpdateMode' in request_json and request_json["UpdateMode"] == "Update" and request_json["TableName"] == "appsheetdemo":
        access_token = get_access_token()
        # Init_K8s_Client(access_token)
        fleet_name = request_json["Data"]["FleetName"]
        print(fleet_name)
        if fleet_name == "Blue":
            fleet_name="supertuxkart-blue"
        else:
            fleet_name="supertuxkart-green"
        UpdateFleetReplica(access_token, fleet_name.lower(), int(request_json["Data"]["Replica"]))
        return ("Done", 200)
    elif request_args and 'name' in request_args:
        name = request_args['name']
    else:
        name = 'World'
    return 'Hello {}!'.format(escape(name))
