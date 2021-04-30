# Kubernetes
- Open Source Container orchestration tool
- Developed by Google
- Help to Manage Containerized Applications in different environment

### Features:
 - High Availability (Zero Downtime)
 - Scalability (High Performance)
 - Disaster Recovery (backup & restore)

### Components
1. **Pod:**
   - Smallest units in K8s
   - Abstruction over the container
   - You can only interact with kubernetes layers
   - Usually one application in a Pod
   - Each pod gets their own IP address
   - When Pod die, New IP address assign on re-creation
2. **Service:**
    - Pod's IP address might change on re-creation, to resolve IP-address changes, service component comes into the Picture.
    - Service is having static/permanent IP address
    - Life cycle of Pod & service is not connected, Means even pod die, Service ip address remains un-changed.
    - 2 types of services
        - Internal service  (ex, Database connection)
        - External service (ex,  access application on external browser)
3. **Ingress:**
    - 192.168.2.1:8080/api ---> This url form is not good for end product, we need a https://my-app.com/api
    - Ingress redirect a domain to exaxt service IP address and port.
4. **ConfigMap:**
    - Store the external configuration data.
    - Key and value pair.
    - Store in the plan text.
5. **Secrets:**
   - Similiar to ConfigMap, but store the data secure manner.
   - base64 encoded format.   
6. **Volumns:** 
   -If pod contains my-app container and db container, when pod restart, all the data has been gone, even in the database.
   - All the presistend data(DB data) or log should available, even pod is restart. It resolve use of Volumns component.
8. **StateFulSet:**
9. **Deloyment:**


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
