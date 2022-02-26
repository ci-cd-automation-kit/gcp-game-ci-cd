# appsheet-backend

## Setup Quickstart
  By using this guide, users can easily setup the appsheet demo application for GameCICD demo.

### Step 1 -- Setup Appsheet
Access [Appsheet](https://www.appsheet.com/Template/Apps), Copy the GameCICD-Template application.

**Change to naviage path**

Get Appsheet AppID. Select the app you create in previous step, go to **Manage->Integration Page**, Select **IN: from cloud services to your app**, record down the **App ID** (Sample like **4b5d9262-92bc-4386-9f78-213cbae64d68**). Click the **Enable** button, then click **Create Application Access Key**, get the **Access Key** by clicking **Show Access Key**, record down the **Access Key**.

### Step 2 -- Get Demo K8s Environment Endpoint
This app just using appsheet as frontend, all the action taken in the UI will trigger an Cloud Function that manipulating the K8s CRD on the GKE envrionment. Each demo environment has its own GKE environment with unique endpoint address. This step will collect the K8s Environment Endpoint that hosted on the GKE, which will be used to initiate the Cloud Function.
**Cloud Shell**, Replace the **Cluster_Name** and **Cluster_Region** with your environment information.
```
gcloud container clusters describe **Cluster_Name** --region **Cluster_Region** --format="value(endpoint)"
```
Record down the ouput as k8s_endpoint.

### Step 3 -- Deployment the Appsheet-Backend Code to Cloud Function.
For users's Project the host outside the **google.com**, we suggest deploy the appsheet-backend cloud function with unauthenticated invoke option, it will be simple for your appsheet setup. For users's Project the host under the **google.com**, cause the Org security policy will deny the unauthenticated invoke, users will need to execute Add-On step to using the token authentication method.
**Cloud Shell**, Replace the **K8s_Endpoint**, **appsheet_app_id** and **appsheet_accesskey** with your environment information.
```
gcloud functions deploy demo_appsheet_backend --runtime python39 --trigger-http --allow-unauthenticated --set-env-vars k8s_endpoint=**K8s_Endpoint collect from step 2**,appsheet_app_id=**app id collect from step 1**,appsheet_accesskey=**app access key collect from step 1**
```

**Add-On Step (Optional: Only apply to demo project hosted under google.com)**
Generate Google Cloud Function invoke Token.
**Cloud Shell**
```
gcloud auth print-identity-token
```
Record Down the token, Go to [Appsheet](https://www.appsheet.com/Template/Apps), Select the app you created in previous step, go to Automation->FleetBot, select ChangeFleet->UpdateFleet, on the right side edit tab change the HTTP Headers in following format. Replace the **Token** with the one generated before. 
```
Authorization: "bearer **Token Generate in Previous Step**"
```
Follow the same step to update GameServerBot->ChangeGameServer->UpdateGameServer.

### Step 4 -- Use Appsheet Frontend Demo.
Current release Appsheet Frontend support:
* Change Agones Fleet Replica
* Apply New GameServer
* Delete GameServer
