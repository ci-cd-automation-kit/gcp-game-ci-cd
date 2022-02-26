# Follow the tutorial in the cloud shell
You can follow the below tutorial by clicking the **Open in Cloud Shell**, the tutorial will guide you through the whole process.

[![Open in Cloud Shell](https://gstatic.com/cloudssh/images/open-btn.png)](https://ssh.cloud.google.com/cloudshell/open?cloudshell_git_repo=https://github.com/grogu-gaming/tutorial.git&cloudshell_tutorial=tutorial.md)

For more information, you can also refer to the following steps.

## Set up GitLab SSH key

Reference: [how to generate a SSH key](https://gitlab.endpoints.cn-gaming-cicd.cloud.goog/help/ssh/README#generating-a-new-ssh-key-pair)

Open a terminal. You may want to use an email address for the comment "&lt;comment>".


```
ssh-keygen -t rsa -b 2048 -C "<comment>"
```


This command generate a SSH key in **~/.ssh**, which include** id_rsa** and **id_rsa.pub**


```
cat ~/.ssh/id_rsa.pub ​​| clip
```


Copy the content in **id_rsa.pub**, paste it on GitLab SSH Key.


<img src="https://lh4.googleusercontent.com/F_kptPNQnTFf5NTjfzXmaxJXQiucZ2hNCpmxwwK9JS1A1-ledHV1BL5ATMQyy6eIPQ512ENYGvOgDd1oU2OWFn6hr1dMjGjO5eAq6NXfb_H0ZK7OTD4jXb29Kv1nWxwBn2WlOe1rK2AAvoujprsqxT9abYaw80azflzp1suWOc0TQmxG=s0" alt="image text" title="imae Title" />



Click **Add key.**

Test the SSH connection,  replacing gitlab.example.com with your GitLab instance URL.


```
ssh -T git@gitlab.example.com
```


If this is the first time you connect, you should verify the authenticity of the GitLab host. If you see a message like:


```
The authenticity of host 'gitlab.example.com (35.231.145.151)' can't be established.
ECDSA key fingerprint is SHA256:HbW3g8zUjNSksFbqTiUWPWg2Bq1x8xdGUrliXFzSnUw.
Are you sure you want to continue connecting (yes/no)? yes
Warning: Permanently added 'gitlab.example.com' (ECDSA) to the list of known hosts.
```


Type yes and press Enter.



![alt_text](https://lh6.googleusercontent.com/ZnlcGM-oRd7In2e2Im59_6KDplh0X2GFxmRZ5IWGMs3qdndYh72bft9PwixjuaQ_ZjmVYN2CfcArHX6PveT0dUxJXz1gUNyVRnvghr7WJ_TqGDy-23wCY5iqLw7Mi8YETVs23Cu-8iStUubFYoYfYXZTAVCkDASxQo4y6mQReR_A2m2J=s0 "image_tooltip")


You should receive a **Welcome to GitLab, @username! **Message.


## Save SSH key to Secret Manager

On the **Secret Manager** page, click **Create Secret**.

On the **Create secret** page, under **Name**, enter a name for your secret.

In the **Secret value** field, enter a name for your secret or upload a file. \
To upload your SSH key file, click **Upload** to include your **~/.ssh/id_rsa** file.

Or paste the content in **~/.ssh/id_rsa** file to the **Secret value** field.

Leave the **Regions** section unchanged.

Click the **Create secret** button to create your secret.


## Set up GKE cluster and Agones


We will use [terraform scripts](https://gitlab.endpoints.cn-gaming-cicd.cloud.goog/gaming-ci-cd-automation/core/-/blob/main/main.tf) to create a VPC and GKE cluster, use Helm to install Agones.
We will use a [cloud build config file](https://gitlab.endpoints.cn-gaming-cicd.cloud.goog/gaming-ci-cd-automation/core/-/blob/main/cloud-build/cloud_build_terraform.yaml) to apply terraform scripts. It includes the following steps:
1. SSH authentication and clone the Gitlab repository
2. Replace the project id field
3. Replace the cluster name field
4. Terraform init, plan and apply

Go to the cloud build, create a webhook trigger:
1. Click **Create trigger**.
2. Enter the following trigger settings:
    1. **Name**: A name for your trigger.
    2. **Event**: Select **Webhook event** to set up your trigger to start builds in response to incoming webhook events.
    3. **Webhook URL**: Use the webhook URL to authenticate incoming webhook events.
    4. use an existing secret:
        1. Select **Use existing**.
        2. In the **Secret** field, select the name of the secret you want to use from the drop-down menu or follow the instructions to add a secret by resource ID.
        3. In the **Secret version** field, select your secret version from the drop-down menu.
3. **Configuration**: Copy and paste the [cloud build config file](https://gitlab.endpoints.cn-gaming-cicd.cloud.goog/gaming-ci-cd-automation/core/-/blob/main/cloud-build/cloud_build_terraform.yaml).
4. **Substitutions**: define 2 substitution variables using this field.

    
![alt_text](https://lh4.googleusercontent.com/yr0xBiBJC3xn_FxX9OfyCcMilj6bNYHwKmZC-fsMGIeXdcYVQB4VDAwAArolk0mrwM0r9MOVGrq_s_2jJ7y97zcbnpE4XG0fbQjhZIFzJtZ0cUc5EG8XlYSG_n_NJsDNpYGIBImxSnrkAdM1vISA5DROKGiWfzunrdFdbTt5MKybKwBH=s0 "image_tooltip")

These substitution variables will be replaced with the parameters passed in the webhook POST request.\
**_CLUSTER_NAME='$(body.cluster_name)'** specifies the cluster name of the cluster you are going to create.\
**_REGION='$(body.region)** specifies the region of the cluster.

Use the “curl” command to trigger cloud build, specify the cluster name and the region  using the JSON format, the parameters will be replaced dynamically in the cloud build config file.

```
curl -X POST -H "application/json" "https://cloudbuild.googleapis.com/v1/projects/cn-gaming-cicd/triggers/core-trigger:webhook?key=AIzaSyBe45IJgmgE3Ssoekh77cl3Z1riBmG7N_I&secret=9LvtgKaNTnKA1ATGBA_167JJPIWynkbB" -d '{"cluster_name":"cluster-xyz", "region":"us-central1"}'
```



## Apply the configuration of the Fleet by push events


We will use a [fleet config file](https://gitlab.endpoints.cn-gaming-cicd.cloud.goog/gaming-ci-cd-automation/core/-/blob/main/modules/agones/fleet_configs_simple.yaml).In this demo, we use the simple game server to test the Quilkin proxy. 
1. Create a fleet with 2 replica simple game servers. we’ll take the example container  that Agones provides for the simple game server.
2. Create a ConfigMap to store the yaml for a static configuration for Quilkin that will accept connections on port 26002 and route then to the simple game server on port 7654.
        
    The CaptureBytes filter will find the first 3 bytes within a packet, and capture it into Filter Dynamic Metadata, so that it can be utilised by filters further down the chain, which is the TokenRouter filter. This TokenRouter Filter compares the token found in the Filter Dynamic Metadata from the  CaptureBytes Filter, and compares it to Endpoint's tokens, and sends packets to those Endpoints only if there is a match. 
        
    In this example, the base64 encoded token is “YWJj”, if the token is found in the first 3 bytes within the packet, it will be removed and sended the rest of the message to the “127.0.0.1:7654”, which is listened by the simple game server. 
    
    
```yaml
apiVersion: v1
kind: ConfigMap
metadata:
  name: quilkin-config-simple
data:
  quilkin.yaml: |
    version: v1alpha1
    proxy:
      id: proxy-demo-simple
      port: 26002
    static:
      filters:
        - name: quilkin.extensions.filters.debug.v1alpha1.Debug # Debug filter: every packet will be logged
          config:
            id: debug-1
    
        - name: quilkin.extensions.filters.capture_bytes.v1alpha1.CaptureBytes # Capture and remove the authentication token
          config:
              strategy: PREFIX
              size: 3
              remove: true
              metadataKey: quilkin.dev
        - name: quilkin.extensions.filters.token_router.v1alpha1.TokenRouter
          config:
              metadataKey: quilkin.dev
      endpoints: 
          address: 127.0.0.1:7654
          metadata:
              quilkin.dev:
                tokens:
                    - YWJj # abc
                    - MXg3aWp5Ng== # Authentication is provided by these ids, and matched against 
                    - OGdqM3YyaQ== # the value stored in Filter dynamic metadata
```    
3. Run Quilkin alongside each dedicated game server as a sidecar.

We will use the [cloud build config file](https://gitlab.endpoints.cn-gaming-cicd.cloud.goog/gaming-ci-cd-automation/core/-/blob/main/cloud-build/cloud_build_fleet_configs.yaml) to apply the fleet configurations. It includes the following steps:
1. SSH authentication and clone the Gitlab repository
2. Connect the GKE cluster
3. Apply the fleet_config.yaml file

Create the following substitution variables.

    

![alt_text](https://lh3.googleusercontent.com/Qv6vlAcIdwEUO23cT8iJJIaM9k6BPVHqHsPyJoj2jXfyA9PAUoQojmq9SdQEhlzQrKVOOehrDrScJlGbpdSPFsspk14NwKmHH-9K6BT-n687iphEXDNTFoxPAmziwspcMVdokZQ901ARvl14wekgpbJIxvlKplgoh-LI8l7STFSqPVNb=s0 "image_tooltip")

These substitution variables will be replaced with the parameters passed in the webhook POST request.\
**_CLUSTER_NAME='$(body.cluster_name)'** specifies the cluster name of the cluster you are going to create.\
**_REGION='$(body.region)** specifies the region of the cluster.\
**_CONFIG_FILE='$(body.config_file)'** specifies the config file name of the fleet config file in the GitLab.



Use “curl” command to trigger cloud build, it will apply the fleet configurations with specified yaml file name. In this demo, we use the simple game server to test the Quilkin proxy, so “config_file” is “fleet_configs_simple.yaml”. 

```
curl -X POST -H "application/json" "https://cloudbuild.googleapis.com/v1/projects/cn-gaming-cicd/triggers/fleet-config-trigger:webhook?key=AIzaSyBe45IJgmgE3Ssoekh77cl3Z1riBmG7N_I&secret=xh8twKFPVoA4SG2J9MxXY20sxD50EMVd" -d '{"cluster_name":"cluster-xyz", "region":"us-central1", "config_file":"fleet_configs_simple.yaml"}'
```


Test the fleet creation.
1. In Cloud Build history, it was successfully built.
2. In the cloud shell, verify that gameservers were created, and the state is ready.

```
Kubectl get gs
```


The output should look like this:

```bash
NAME                            STATE  ADDRESS     PORT         NODE                                                 AGE
simple-game-server-nndvr-5gvch  Ready  IP_ADDRESS  PORT_NUMBER  gke-cluster-xyz-cluster-xyz-node-pool-db3d2ee8-dlg7  11s
simple-game-server-nndvr-vqwpl  Ready  IP_ADDRESS  PORT_NUMBER  gke-cluster-xyz-cluster-xyz-node-pool-db3d2ee8-dlg7  11s
```


## Test the Quilkin Proxy

Run [SimpleGameServerQuilkinTest.py](https://gitlab.endpoints.cn-gaming-cicd.cloud.goog/gaming-ci-cd-automation/core/-/blob/main/QuilkinTest/SimpleGameServerQuilkinTest.py) to send a UDP package of message “abcEXIT”, “abc” is the first 3 bytes of the message, and it matches the token, so it will be removed by the filter, and pass “EXIT” to the game server.

The simple game server will echo the message back, and the message should be “EXIT”.




![alt_text](https://lh6.googleusercontent.com/nRr5fcxSAqIxuu3NFXZUM7SOu_3prci91C2WUTseGl9O0D5bux8vs7OOmefyym9BnqRFsfRfpOTWVrIfzdh6wT6Uof3xOMWz8hJDqEAqcey-WAxYoCGTsQUHcxrgSm3Nwj4tCnf5O4q2AVfIKxRQOkj9DIJifoA6BqsMQYtYmZ8SYQBr=s0 "image_tooltip")


After the game server “simple-game-server-7qsr8-tlmvh” receives an “EXIT” message, it will shut down the game server, and create a new one. You should see that a new game server “simple-game-server-7qsr8-gszqf” is created, the port is changed from “7040” to “7537”.





![alt_text](https://lh6.googleusercontent.com/4j-ItSw0pi-V-7loosXBChW4NMy6TjDOnIDYnSaiyYCVSYBZhnDKu-03ptKH7Mf4GcHiBI9gniTJXcOt-5kW82Z1AYWxric8AVmWezBF3opn99SrC9EXG7qYztiYk9HGG14bm8G9c_DJkL60-wVmr-o1w43vVQDi7O36QbR4LJkJzXN5=s0 "image_tooltip")



