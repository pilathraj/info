# Kubernetes
- Open Source Container orchestration tool
- Developed by Google
- Help to Manage Containerized Applications in different environment

### Features:
 - High Availability (Zero Downtime)
 - Scalability (High Performance)
 - Disaster Recovery (backup & restore)

### Basic Components
1. **Worker Node**
   - Each node has multiple pods on it
   - 3 processes must installed on every node.
      1. Container runtime
      2. Kubelet - communication between node vs  container.
      3. Kube Proxy - forwards the requests to service
   - **Worker node** actually do the work
2. **Master node**
      - How to Schedule a pod?
      - How to re-Schedule / re-start?
      - How to monitor?
      - Join a new node?
      - All ^^^^ these managing process done by master node.
      - 4 processes is running on every master node
         1. API Server
             -  Cluster Gateway
             -  Client -> Some request -> API server -> Validate request -> Some other processes... -> pod
             -  It's handle both Query and Updates
             -  Single entry point into the Cluster, (so we'll manage scecurily) 
          2. Scheduler 
             - Schedule new Pod -> API server ->Scheduler (I know where to put the pod) -> start application pod on one of the worker nodes
          3. Controller manager
             - Detect the cluster state changes
             - Controller manager -> Scheduler -> start new pod on the worker nodes
          4. etcd
              - Cluster brain
              - Key and value pair
              - Means, the cluster changes getting stores in the key, value pair.
              - In practice, K8s cluster made multiple master node.
2. **Pod:**
   - Smallest units in K8s
   - Abstruction over the container
   - You can only interact with kubernetes layers
   - Usually one application in a Pod
   - Each pod gets their own IP address
   - When Pod die, New IP address assign on re-creation
3. **Service:**
    - Pod's IP address might change on re-creation, to resolve IP-address changes, service component comes into the Picture.
    - Service is having static/permanent IP address
    - Life cycle of Pod & service is not connected, Means even pod die, Service ip address remains un-changed.
    - Service also act as a load balancer.
    - 2 types of services
        - Internal service  (ex, Database connection)
        - External service (ex,  access application on external browser)
4. **Ingress:**
    - 192.168.2.1:8080/api ---> This url form is not good for end product, we need a https://my-app.com/api
    - Ingress redirect a domain to exaxt service IP address and port.
5. **ConfigMap:**
    - Store the external configuration data.
    - Key and value pair.
    - Store in the plan text.
6. **Secrets:**
   - Similiar to ConfigMap, but store the data secure manner.
   - base64 encoded format.   
7. **Volumns:** 
   -If pod contains my-app container and db container, when pod restart, all the data has been gone, even in the database.
   - All the presistent data(DB data) or log should available, even pod is restart. It resolve use of Volumns component.
   - Volumns allow you to store the data in your host machine or Remote server.
   - K8s doesn't manage data persistance -Administrator responsibility to manage the presistance data effectivily.
8. **Deployment:**
    - blue print of pod is know as deployment.
    - No.of replicates you want to specify, while deploying pod.
    - Abstract layer of pod.   
    - DB Can't replicated via Deployment.
    - **Do** -Replicate a pod, Scale up / Scale down.
    - In Practies,we mostly working on Deployment, not in pods.
9. **StateFulSet:**
    - If you replicate the DB, we need to ensure, the data read & write correctly in different pods. Else it might caused the data-inconsistency. 
    - StateFullSet Component address this issue.
    - Any statefull application or Databases need to create via StatefulSet Component, not ~Deployment~
    -  **Do** -Replicate a pod, Scale up / Scale down pods, in addition to that synchronize read & write operation.
    -  Deploying the DB in StateFullSet i not easy, Mostly Database deployed outside a K8s cluster in most cases.

### Example:
   - Very small cluster alteast contain 2 master node  and 3 worker nodes.
   - Master node do less work, so it having less resources in terms of CPU, Memory and Storage.

### Layers of Abstraction
 - Deployment manganged replicasets.
           |||
           vvv
 - Replicaset manganged a pod.
           |||
           vvv
 - Pod is the abtraction of container.
           |||
           vvv
 - Container.
 - **Everything below the deployment automatically managed by kubernetes.**


### Local setup:  Install minikube & kubectl
**Windows > Cmd > Run Administrator**, Then
```cmd
choco install minikube
minikube start
```
### minikube UI
```cmd
minikube dashboard
```
### kubectl
- kubectl get nodes - Return nodes list
- kubectl get pod - return pods list
- kubectl get services - return service list
- kubectl create deployment name --image=image [--dry-run] [option]  -- Deployed pods.

- **CRUD commands:**
   * kubectl create deployment [name]
   * kubectl edit deployment [name]
   * kubectl delete deployment [name]
- **Status of different K8s components**
   * kubectl get nodes | pod | services | replicaset | deployment
- **Debugging pod**
   * Kubectl logs [pod name]

```cmd
kubectl create deployment ngnix-depl --image=nginx
O/P: deployment.apps/ngnix-depl created
kubectl get deployment
O/P:
NAME         READY   UP-TO-DATE   AVAILABLE   AGE
ngnix-depl   0/1     1            0           101s
kubectl get pod
ngnix-depl-5b9c968f6b-xbrsk   1/1     Running   0          5m47s
```
- Replicaset is automatically created, you can check it
```cmd
kubectl get replicaset
o/p:
NAME                    DESIRED   CURRENT   READY   AGE
ngnix-depl-5b9c968f6b   1         1         1       5d20h
```
- Get pod
```cmd
kubectl get pod
NAME                          READY   STATUS    RESTARTS   AGE
ngnix-depl-5b9c968f6b-xbrsk   1/1     Running   1          5d20h
```
- Edit deployment
```cmd
kubectl edit deployment ngnix-depl 
```
- we get auto generated deployment file
```yaml
# Please edit the object below. Lines beginning with a '#' will be ignored,
# and an empty file will abort the edit. If an error occurs while saving this file will be
# reopened with the relevant failures.
#
apiVersion: apps/v1
kind: Deployment
metadata:
  annotations:
    deployment.kubernetes.io/revision: "1"
  creationTimestamp: "2021-04-30T12:24:26Z"
  generation: 1
  labels:
    app: ngnix-depl
  name: ngnix-depl
  namespace: default
  resourceVersion: "41671"
  uid: 9006a7bb-35f1-4365-804e-69f1f77df30b
spec:
  progressDeadlineSeconds: 600
  replicas: 1
  revisionHistoryLimit: 10
  selector:
    matchLabels:
      app: ngnix-depl
  strategy:
    rollingUpdate:
      maxSurge: 25%
      maxUnavailable: 25%
    type: RollingUpdate
  template:
    metadata:
      creationTimestamp: null
      labels:
        app: ngnix-depl
    spec:
      containers:
      - image: nginx
        imagePullPolicy: Always
        name: nginx
        resources: {}
        terminationMessagePath: /dev/termination-log
        terminationMessagePolicy: File
      dnsPolicy: ClusterFirst
      restartPolicy: Always
      schedulerName: default-scheduler
      securityContext: {}
      terminationGracePeriodSeconds: 30
status:
  availableReplicas: 1
  conditions:
  - lastTransitionTime: "2021-04-30T12:24:26Z"
    lastUpdateTime: "2021-04-30T12:28:42Z"
    message: ReplicaSet "ngnix-depl-5b9c968f6b" has successfully progressed.
    reason: NewReplicaSetAvailable
    status: "True"
    type: Progressing
  - lastTransitionTime: "2021-05-06T08:52:54Z"
    lastUpdateTime: "2021-05-06T08:52:54Z"
    message: Deployment has minimum availability.
    reason: MinimumReplicasAvailable
    status: "True"
    type: Available
  observedGeneration: 1
  readyReplicas: 1
  replicas: 1
  updatedReplicas: 1
```
-- set image: nginx:1.16 and save the file. Then run
```cmd
kubectl get pod
NAME                          READY   STATUS        RESTARTS   AGE
ngnix-depl-5b9c968f6b-xbrsk   1/1     Terminating   1          5d21h
ngnix-depl-5fff49f8bb-vvh72   1/1     Running       0          50s
```
- kubectl logs <pod-name>
```cmd
 kubectl logs ngnix-depl-5fff49f8bb-vvh72
 ```
- Describe pod
```cmd
   kubectl create deployment mongo-dpl --image=mongo
   kebectl get pod
 ```
 O/P:
 ```cmd
 NAME                          READY   STATUS              RESTARTS   AGE
mongo-dpl-7675b8cd4c-2jfnb    0/1     ContainerCreating   0          13s
ngnix-depl-5fff49f8bb-vvh72   1/1     Running             0          65m
```
```cmd
 kubectl describe pod mongo-dpl-7675b8cd4c-2jfnb
 O/P:
 Name:           mongo-dpl-7675b8cd4c-2jfnb
Namespace:      default
Priority:       0
Node:           minikube/192.168.49.2
Start Time:     Thu, 06 May 2021 16:30:41 +0530
Labels:         app=mongo-dpl
                pod-template-hash=7675b8cd4c
Annotations:    <none>
Status:         Pending
IP:
IPs:            <none>
Controlled By:  ReplicaSet/mongo-dpl-7675b8cd4c
Containers:
  mongo:
    Container ID:
    Image:          mongo
    Image ID:
    Port:           <none>
    Host Port:      <none>
    State:          Waiting
      Reason:       ContainerCreating
    Ready:          False
    Restart Count:  0
    Environment:    <none>
    Mounts:
      /var/run/secrets/kubernetes.io/serviceaccount from default-token-zgmd5 (ro)
Conditions:
  Type              Status
  Initialized       True
  Ready             False
  ContainersReady   False
  PodScheduled      True
Volumes:
  default-token-zgmd5:
    Type:        Secret (a volume populated by a Secret)
    SecretName:  default-token-zgmd5
    Optional:    false
QoS Class:       BestEffort
Node-Selectors:  <none>
Tolerations:     node.kubernetes.io/not-ready:NoExecute op=Exists for 300s
                 node.kubernetes.io/unreachable:NoExecute op=Exists for 300s
Events:
  Type    Reason     Age   From               Message
  ----    ------     ----  ----               -------
  Normal  Scheduled  46s   default-scheduler  Successfully assigned default/mongo-dpl-7675b8cd4c-2jfnb to minikube
  Normal  Pulling    45s   kubelet            Pulling image "mongo"
```
- run the same comments again, you may get output:
```cmd
 ....
 ....
 Events:
  Type    Reason     Age    From               Message
  ----    ------     ----   ----               -------
  Normal  Scheduled  6m23s  default-scheduler  Successfully assigned default/mongo-dpl-7675b8cd4c-2jfnb to minikube
  Normal  Pulling    6m23s  kubelet            Pulling image "mongo"
  Normal  Pulled     5m25s  kubelet            Successfully pulled image "mongo" in 57.8312863s
  Normal  Created    5m25s  kubelet            Created container mongo
  Normal  Started    5m25s  kubelet            Started container mongo
 ```
- Run interactive mode
```cmd
 kubectl exec -it mongo-dpl-7675b8cd4c-2jfnb -- bin/bash
 ```
 - delete the deployment
 ```cmd
 kubectl delete deployment mongo-dpl
 ```
 - **Apply** create via config file.
 ```cmd
 kubectl apply -f config-file.yaml 
 ```
