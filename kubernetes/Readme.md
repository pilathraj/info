# Kubernetes
- Open Source Container orchestration tool
- Developed by Google
- Help to Manage Containerized Applications in different environment

### Features:
 - High Availability (Zero Downtime)
 - Scalability (High Performance)
 - Disaster Recovery (backup & restore)

### Components
1. **Node**
   - Colletion of pods running.
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
8. **StateFulSet:**
9. **Deployment:**
    - blue print of pod is know as deployment.
    - No.of replicates you want to specify, while deploying pod.
    - Abstract layer of pod.
    - In Practies,we mostly working on Deployment, not in pods.
    - DB Can't replicated via Deployment.

### Install minikube
**Windows > Cmd > Run Administrator**, Then
```cmd
choco install minikube
minikube start
```
### minikube UI
```cmd
minikube dashboard
```
