# Pod Controller

Pods are the minimum snap-in of Kubernetes, which can be divided into two classes according to the way they are created :

- Autonomous POD: A pod created directly by Kubernetes, which will not be rebuilt after it is deleted

- Controller created Pods: Pods created by Kubernetes through the controller, which are automatically rebuilt after being deleted

Some common controller:

- ReplicationController：Has been abandoned.
- ReplicaSet：Ensure that the number of copies is always maintained as expected, and support pod number expansion and scaling, image version upgrade
- Deployment：Control POD by controlling replica set and support rolling upgrade and version rollback
- Horizontal Pod Autoscaler：The number of PODS can be automatically adjusted horizontally according to the cluster load to achieve peak load cutting and valley filling
- DaemonSet：Run only one copy on a specified node in the cluster, typically for daemon class tasks
- Job：It creates pods that exit as soon as they complete their tasks, requiring no reboot or rebuild, and are used to perform one-off tasks
- Cronjob：It creates pods that are responsible for periodic task control and do not need to run continuously in the background
- StatefulSet：Manage stateful applications

---

### ReplicaSet(RS)

The main role of replica set is to ensure the normal operation of a certain number of PODS. It continuously monitors the running status of these pods and will restart or rebuild the pods once they fail. At the same time, it also supports the number of POD expansion and mirror version upgrade or downgrade.

Replica Set Resource List file:

```yaml
apiVersion: apps/v1 
kind: ReplicaSet       
metadata: 
  name: # rs name
  namespace: 
  labels: 
    controller: rs
spec: 
  replicas: 3 # Copy number
  selector: # Selector, which specifies which pods the controller manages
    matchLabels:      # Labels matching rule
      app: nginx-pod
    matchExpressions: 
      - {key: app, operator: In, values: [nginx-pod]}
  template: # Template. When the number of copies is insufficient, a POD copy is created based on the following template
    metadata:
      labels:
        app: nginx-pod
    spec:
      containers:
      - name: nginx
        image: nginx:1.17.1
        ports:
        - containerPort: 80
```

```yaml
# example
apiVersion: apps/v1
kind: ReplicaSet   
metadata:
  name: pc-replicaset
  namespace: dev
spec:
  replicas: 3
  selector: 
    matchLabels:
      app: nginx-pod
  template:
    metadata:
      labels:
        app: nginx-pod
    spec:
      containers:
      - name: nginx
        image: nginx:1.17.1
```

```bash
# create 
[root@master ~]# kubectl create -f pc-replicaset.yaml
replicaset.apps/pc-replicaset created

#check 
[root@master ~]# kubectl get rs pc-replicaset -n dev -o wide
NAME          DESIRED   CURRENT READY AGE   CONTAINERS   IMAGES             SELECTOR
pc-replicaset 3         3       3     22s   nginx        nginx:1.17.1       app=nginx-pod

# edit
# edit relicate in file and rebuild
[root@master ~]# kubectl edit rs pc-replicaset -n dev
replicaset.apps/pc-replicaset edited

# use command to change replicate number
[root@master ~]# kubectl scale rs pc-replicaset --replicas=2 -n dev
replicaset.apps/pc-replicaset scaled

# update
# edit image in yaml file and rebuild
[root@master ~]# kubectl edit rs pc-replicaset -n dev
replicaset.apps/pc-replicaset edited

# use command
[root@master ~]# kubectl set image rs pc-replicaset nginx=nginx:1.17.1  -n dev
replicaset.apps/pc-replicaset image updated
```

### Deployment(Deploy)

```yaml
apiVersion: apps/v1 
kind: Deployment       
metadata: 
  name: 
  namespace: 
  labels:
    controller: deploy
spec: 
  replicas: 3
  revisionHistoryLimit: 3 # Preserve historical versions
  paused: false # Suspension of deployment
  progressDeadlineSeconds: 600 # default 600
  strategy: 
    type: RollingUpdate # Rolling update strategy
    rollingUpdate: 
      maxSurge: 30% # The maximum number of extra copies that can exist, which can be a percentage or an integer
      maxUnavailable: 30% # Maximum The maximum number of pods in the unavailable state, which can be a percentage or an integer
  selector: d
    matchLabels:      
      app: nginx-pod
    matchExpressions: 
      - {key: app, operator: In, values: [nginx-pod]}
  template: 
    metadata:
      labels:
        app: nginx-pod
    spec:
      containers:
      - name: nginx
        image: nginx:1.17.1
        ports:
        - containerPort: 80
```

Change replicate number and image are same as rs before.

Deployment supports two update strategies: 'rebuild update' and 'rolling update'. The policy type can be specified by 'strategy'. Two attributes are supported:

```bash
strategy：
  type：Specifies the policy type. Two policies are supported
    Recreate：All existing pods are killed before new ones are created
    RollingUpdate：There are two versions of pod in the update process
  rollingUpdate：This parameter is valid when type is rolling Update and is used to set parameters for Rolling Update. Two properties are supported:
    maxUnavailable：Used to specify the maximum number of pods that cannot be used during the upgrade. The default is 25%.
    maxSurge： Used to specify the maximum number of pods that can be exceeded during an upgrade. The default is 25%.
```

**version recover**

```
- status       The upgrade status is displayed
- history     The upgrade history is displayed

- pause       Pause the version upgrade process
- resume      Continue the stalled version upgrade process
- restart      Restart the version upgrade process
- undo       Rollback to a previous version (you can use --to-revision to rollback to a specified version)
```

```
[root@master ~]# kubectl rollout status deploy pc-deployment -n dev
deployment "pc-deployment" successfully rolled out

[root@master ~]# kubectl rollout history deploy pc-deployment -n dev
deployment.apps/pc-deployment
REVISION  CHANGE-CAUSE
1         kubectl create --filename=pc-deployment.yaml --record=true
2         kubectl create --filename=pc-deployment.yaml --record=true
3         kubectl create --filename=pc-deployment.yaml --record=true

[root@master ~]# kubectl rollout undo deployment pc-deployment --to-revision=1 -n dev
deployment.apps/pc-deployment rolled back
```

###  Service

In Kubernetes, POD is the carrier of the application program. We can access the application program through the IP of POD, but the IP address of POD is not fixed, which means that it is not convenient to directly use the IP of POD to access the service.

To solve this problem, Kubernetes provides the resource Service, which aggregates pods that provide the same service and provides a unified entry address. You can access the pod service by accessing the entry address of the Service.

In many cases, service is just a concept. The kube-proxy service process is actually used. Each node runs a Kube-proxy service process. When a service is created, information about the service is written to etcd via apI-server. Kube-proxy detects the service changes based on the listening mechanism, and then converts the latest service information into the corresponding access rules .

```yaml
kind: Service  
apiVersion: v1  
metadata: 
  name: service 
  namespace: dev 
spec: 
  selector: 
    app: nginx
  type: # Service type: specifies the access mode of service
  clusterIP:  # IP address of the virtual service
  sessionAffinity: # Session affinity: Supports client IP and None
  ports: 
    - protocol: TCP 
      port: 3017  # service port
      targetPort: 5003 # pod port
      nodePort: 31122 # host machine port
```

---

### Data storage

Volume is a shared directory in POD that can be accessed by multiple containers. It is defined on POD and then mounted to a specific file directory by multiple containers in a POD. Kubernetes uses volume to realize data sharing and persistent storage between different containers in the same POD. The life container of a volume is independent of the life cycle of the individual container in the POD, and data in the volume is not lost when the container terminates or restarts.

##### 1. EmptyDir

Empty dir is the most basic volume type. An empty dir is an empty directory on host.

Empty dir is created when pod is assigned to Node, its initial content is empty, and there is no need to specify the corresponding directory file on the host, because Kubernetes will automatically assign a directory, and the data in empty dir will be permanently deleted when POD is destroyed.

It be used as temporary directory or a way to change data between two containers.

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: volume-emptydir
  namespace: dev
spec:
  containers:
  - name: nginx
    image: nginx:1.14-alpine
    ports:
    - containerPort: 80
    volumeMounts:  # add volume in nginx
    - name: logs-volume
      mountPath: /var/log/nginx
  - name: busybox # add volume in busybox
    image: busybox:1.30
    command: ["/bin/sh","-c","tail -f /logs/access.log"] 
    volumeMounts:  
    - name: logs-volume
      mountPath: /logs
  volumes: # declaration
  - name: logs-volume
    emptyDir: {}
```

##### 2.HostPath

The host path is to attach an actual directory from the Node host to the pod for the container to use. This design ensures that the POD is destroyed, but the data basis can be stored on the Node host.

```yaml
apiVersion: v1
kind: Pod
metadata:
  name: volume-hostpath
  namespace: dev
spec:
  containers:
  - name: nginx
    image: nginx:1.17.1
    ports:
    - containerPort: 80
    volumeMounts:
    - name: logs-volume
      mountPath: /var/log/nginx
  - name: busybox
    image: busybox:1.30
    command: ["/bin/sh","-c","tail -f /logs/access.log"]
    volumeMounts:
    - name: logs-volume
      mountPath: /logs
  volumes:
  - name: logs-volume
    hostPath: 
      path: /root/logs
      type: DirectoryOrCreate  # if directory not exists, k8s'll create it
```

---

### Advanced storage

Host path can solve the problem of data persistence, but once the node node fails, if pod is moved to another node, the problem will occur again. In this case, you need to prepare a separate network storage system, more commonly used NFS, CIFS.

##### NFS

NFS is a network file storage system, you can set up a NFS server, and then connect the POD storage directly to the NFS system, so that no matter how the POD is transferred on the node, as long as the node and NFS connection is ok, the data can be successfully accessed.

First step is to install NFS in a server and register this server into k8s pod, and then it will work.

##### PV/PVC

```yaml
# create PV (warehouse of data)
apiVersion: v1
kind: PersistentVolume
metadata:
  name:  pv3
spec:
  capacity: 
    storage: 3Gi
  accessModes:
  - ReadWriteMany
  persistentVolumeReclaimPolicy: Retain
  nfs:
    path: /root/data/pv3
    server: 192.168.109.100
    
# create PVC (request to PV)
apiVersion: v1
kind: PersistentVolumeClaim
metadata:
  name: pvc3
  namespace: dev
spec:
  accessModes: 
  - ReadWriteMany
  resources:
    requests:
      storage: 1Gi
      
      
# add volume
apiVersion: v1
kind: Pod
metadata:
  name: pod2
  namespace: dev
spec:
  containers:
  - name: busybox
    image: busybox:1.30
    command: ["/bin/sh","-c","while true;do echo pod2 >> /root/out.txt; sleep 10; done;"]
    volumeMounts:
    - name: volume
      mountPath: /root/
  volumes:
    - name: volume
      persistentVolumeClaim:
        claimName: pvc3
        readOnly: false        
```



