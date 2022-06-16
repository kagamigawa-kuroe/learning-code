# Kubernetes

As we known, one pod contains several containers. 

There exist a base (root) container named PAUSE in every Pod, who provides IP for his Pod, and show the status of the whole Pod.

K8s provides us with PAUSE container so it always exists, we needn't to do extra configuration by ourselves 

---

### Pod YAML configuration

some basic fields:

```YAML
# * represent mandatory
apiVersion: v1     # * version
kind: Pod       　 # * kind of resource, like Pod
metadata:       　 # * Standard object's metadata
  name: string     # * Pod's name
  namespace: string  # Pod's namespace, default is "default"
  labels:       　　  # Labels
    - name: string      　          
spec:  # * particular definition of pod 
  containers:  # * list of containers
  - name: string   # * container's name
    image: string  # * image's name
    imagePullPolicy: [ Always|Never|IfNotPresent ]  # Strategy for getting image 
    command: [string]   # command will execute when container start
    args: [string]      # parameters of command
    workingDir: string  # working dir of container
    volumeMounts:       # configuration about volume
    - name: string      # name of volume
      mountPath: string # absolute path of it 
      readOnly: boolean # if read only or not
    ports: # port configuration
      containerPort: int  # container's port
      hostPort: int       # port exposed (host machine's port)
      protocol: string    # protocol of port (tcp/udp)
    env:   # environment parameters' list
    - name: string  
      value: string 
    resources: # resources setting
      limits:  # limit of resources
        cpu: string     # number of CPU cores
        memory: string  # memory size
      requests: # setting of resources' request 
        cpu: string    # CPU requests
        memory: string # memory request
    lifecycle: # hock of lifecycle
		postStart: # execute after container stop
		preStop: # execute before container stop
    livenessProbe:  # configuration about container status check
      exec:       　 # exec way
        command: [string]  
      httpGet:       # httpget way
        path: string
        port: number
        host: string
        scheme: string
        HttpHeaders:
        - name: string
          value: string
      tcpSocket:     # tcpSocket way
         port: number
       initialDelaySeconds: 0       # the time of first detection
       timeoutSeconds: 0    　　    # maximum waiting time of detection
       periodSeconds: 0     　　    # regular detection interval
       successThreshold: 0
       failureThreshold: 0
       securityContext:
         privileged: false
  restartPolicy: [Always | Never | OnFailure]  # Strategy of pod restarting
  nodeName: <string> # to refer which server (node) the pod will run on
  nodeSelector: obeject # 
  imagePullSecrets: # 
  - name: string
  hostNetwork: false   # use host machine's network or not
  volumes:   # public volume( shared by all containers in pod)
  - name: string    # name
    emptyDir: {}       # type emptyDir
    hostPath: string   # type hostPath
      path: string      　　        
    secret:       　　　# type secret
      scretname: string  
      items:     
      - key: string
        path: string
    configMap:         # type configmap
      name: string
      items:
      - key: string
        path: string
```

##### Five level 1 fields

```bash
╭─hongzhe@hongzhe-OMEN--HP ~ 
╰─$ kubectl explain pod                                                                     

KIND:     Pod
VERSION:  v1

DESCRIPTION:
     Pod is a collection of containers that can run on a host. This resource is
     created by clients and scheduled onto hosts.

FIELDS:
   apiVersion   <string>
     APIVersion defines the versioned schema of this representation of an
     object. Servers should convert recognized schemas to the latest internal
     value, and may reject unrecognized values. More info:
     https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources

   kind <string>
     Kind is a string value representing the REST resource this object
     represents. Servers may infer this from the endpoint the client submits
     requests to. Cannot be updated. In CamelCase. More info:
     https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds

   metadata     <Object>
     Standard object's metadata. More info:
     https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#metadata

   spec <Object>
     Specification of the desired behavior of the pod. More info:
     https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status

   status       <Object>
     Most recently observed status of the pod. This data may not be up to date.
     Populated by the system. Read-only. More info:
     https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#spec-and-status
```

#### Basic configuration

```yaml
# here is a basic yaml file
apiVersion: v1
kind: Pod
metadata:
  name: pod-base
  labels:
    name: myapp
spec:
  containers:
  - name: nginx
    image: nginx:1.17.1
    # 3 options: always/IfNotPresent/Never
    # always means it always pull image from remote repository
    # IfNotPresent means if image exist in local, we use it directly
    # Never means we never use remote image
    imagePullPolicy: Always 
    resources:
      limits:
        memory: "128Mi"
        cpu: "500m"
    ports:
      - containerPort: 80
        name: nginx_port
        protocol: TCP
      # hostport 8080 (exposed IP in host machine)
      # hostIP (to refer run on which server node)
  - name: busybox
    image: busybox:1.30
    # the command that will be executed when container created
    command: ["/bin/sh","-c","touch /tmp/hello.txt; while true;do /bin/echo $(data +%T) >> /tmp/hello/txt; sleep3; done;"]
    resources:
      # to limit resource usage of one container in the pod
      # to avoid one container use too much resource of one pod
      # and other pod can't work
      limits:
        memory: "128Mi"
        cpu: "500m"
    # request:
    # to minimum resource needed by this container
```

```bash
# create pod
╭─hongzhe@hongzhe-OMEN--HP ~/code/K8s 
╰─$ kubectl apply -f pod-base.yaml 
pod/pod-base created

# enter container
╭─hongzhe@hongzhe-OMEN--HP ~/code/K8s 
╰─$ kubectl exec pod-base -it -c busybox /bin/sh                                             
kubectl exec [POD] [COMMAND] is DEPRECATED and will be removed in a future version. Use kubectl exec [POD] -- [COMMAND] instead.
/ # ls
bin   dev   etc   home  proc  root  sys   tmp   usr   var
```

#### pod lifecycle

The time range from creation to completion of a POD object is commonly referred to as the pod life cycle, which consists of the following processes:

- creation of pod
- run init container
- run main container

  - lifecycle hock when pod start (post start)、lifecycle hock before pod stop (pre stop)

  - liveness probe、readiness probe
- pod stop

<img src="./image/image-20200412111402706.png" alt="image-20200412111402706" style="border:solid 1px" />

Throughout its life cycle, a POD appears in five different status, as follows:

- Pending：Apiserver has created the POD resource object, but it has not yet been scheduled or is in the process of downloading the image
- Running：The POD has been scheduled to a node and all containers have been created by Kubelet
- Succeeded：All containers in the POD have successfully terminated and will not be restarted
- Failed：All containers have terminated, but at least one container failed to terminate, that is, the container returned an exit status with a non-zero value
- Unknown：Apiserver fails to obtain the status information of the POD object, which is usually caused by a network communication failure

---

#### Pod creation process

1. The user submits the POD information to the API Server through Kubectl or other API clients

2. The API Server starts generating information about the POD object, stores the information into etCD, and then returns confirmation information to the client

3. API Server starts to reflect changes to pod objects in ETCD, and other components use the Watch mechanism to track changes on API Server

4. The Scheduler finds a new POD object to create, starts assigning hosts to pods and updates the resulting information to API Server

5. Kubelet on node finds pod scheduling, tries to call docker to start the container, and sends the result back to API server

6. The API Server stores the received POD status information into the ETCD

#### Termination process for pod 

1. The user sends the command to the API Server to delete the POD object
2. The POD object information in the API Servcer is updated over time, and during the grace period (default 30s), the POD is considered dead
3. Mark the POD as terminating
4. Kubelet starts the POD shutdown process when it monitors the pod object's terminating state
5. The endpoint controller monitors the shutdown behavior of a POD object and removes it from the list of endpoints for all service resources matching this endpoint
6. If the current POD object defines a Pre Stop hook handler, execution will start synchronously when it is marked terminating
7. The container process in the POD object receives a stop signal
8. After the grace period, if there are still running processes in the POD, the POD object receives a signal to terminate immediately
9. Kubelet requests the API Server to set the grace period for this POD resource to 0 to complete the deletion operation, at which point the POD is no longer visible to the user

---

### Initialize the container

The initialization container is the container that runs before the pod's main container is started. It does some pre-loading of the main container. It has two characteristics:

1. The initialization container must run until complete. If an initialization container fails, Kubernetes needs to restart it until it completes successfully
2. Initializing containers must be done in a defined order, with subsequent containers running if and only if the current one succeeds

There are many application scenarios for initializing containers. The following are some of the most common:

- Provides utilities or custom code that are not available in the main container image
- The initialization container starts and runs sequentially before the application container, so it can be used to delay the start of the application container until its dependent conditions are met

A example :

Suppose you want to run nginx as the main container, but you need to be able to connect to mysql and Redis servers before you can run nginx .

We can add two init container to make sure we have mysql and redis.

```
apiVersion: v1
kind: Pod
metadata:
  name: pod-initcontainer
  namespace: dev
spec:
  containers:
  - name: main-container
    image: nginx:1.17.1
    ports: 
    - name: nginx-port
      containerPort: 80
  initContainers:
  - name: test-mysql
    image: busybox:1.30
    command: ['sh', '-c', 'until ping 192.168.109.201 -c 1 ; do echo waiting for mysql...; sleep 2; done;']
  - name: test-redis
    image: busybox:1.30
    command: ['sh', '-c', 'until ping 192.168.109.202 -c 1 ; do echo waiting for reids...; sleep 2; done;']
```

### Hook function

Hook functions can sense events in their lifecycle and run user-specified program code when the appropriate time comes.

Kubernetes provides two hook functions after the main container starts and before it stops:

- post start：Execute after the container is created, or restart the container if it fails
- pre stop  ：Execute before the container terminates, after which the container terminates successfully, blocking the operation to delete the container until it completes

The hook handler supports defining actions in one of three ways:

- Exec command: Executes a command within the container

```
  lifecycle:
    postStart: 
      exec:
        command:
        - cat
        - /tmp/hello.txt
```

- TCPSocket：try to access referred socket in current container

```
  lifecycle:
    postStart:
      tcpSocket:
        port: 8080
```

- HttpGET: sends an HTTP request to a URL in the current container

```
  lifecycle:
    postStart:
      httpGet:
        path: /
        port: 80 
        host: 192.168.109.100 
        scheme: HTTP 
```

#### Container detection

Container detection is a traditional mechanism used to check whether application instances in containers are working properly to ensure service availability. If the instance is not detected as expected, Kubernetes removes the instance from service. Kubernetes provides two types of probes for container probing:

- liveness probes：This command is used to check whether the application instance is running properly

- readiness probes：This command is used to check whether the application instance can receive requests

Liveness Probe determines whether to restart the container. Readiness Probe determines whether to forward requests to the container.

Probes also has three ways to realize, really same with hook function.

```bash
# example
apiVersion: v1
kind: Pod
metadata:
  name: pod-liveness-exec
  namespace: dev
spec:
  containers:
  - name: nginx
    image: nginx:1.17.1
    ports: 
    - name: nginx-port
      containerPort: 80
    # Exec command: executes a command in the container. If the exit code of the command execution is 0, the program is considered normal; otherwise, it is abnormal
    livenessProbe:
      exec:
        command: ["/bin/cat","/tmp/hello.txt"] 
    # or
    # tcpSocket will be made to access the port of a user container, and if the connection can be established, the program is considered normal, otherwise it is not
    livenessProbe:
      tcpSocket:
        port: 8080 # Attempt to access port 8080
    # or
    # Call the URL of the Web application in the container. If the status code returned is between 200 and 399, the program is considered normal; otherwise, it is abnormal
    livenessProbe:
    tcpSocket:
      port: 8080
```

#### Restart strategy

In the previous section, kubernetes will restart the pod in which the container is located if there is a problem with the probe. This is determined by the pod restart policy.

- Always ：The container restarts automatically when it fails, which is also the default.
- OnFailure ： Restart when the container terminates and the exit code is not 0
- Never ： Do not restart the container regardless of its state

```yaml
# example
apiVersion: v1
kind: Pod
metadata:
  name: pod-restartpolicy
  namespace: dev
spec:
  containers:
  - name: nginx
    image: nginx:1.17.1
    ports:
    - name: nginx-port
      containerPort: 80
    livenessProbe:
      httpGet:
        scheme: HTTP
        port: 80
        path: /hello
  restartPolicy: Never 
```

#### Pod scheduling

​    By default, the scheduler component uses a scheduler algorithm to determine which node a POD runs on, a process that is not manually controlled. But in practice, this is not enough, because in many cases, we want to control some pods to reach some nodes, so how do we do it? This requires to understand the scheduling rules of Kubernetes for POD. Kubernetes provides four major scheduling modes:

- Automatic scheduling: Which node to run on is completely calculated by scheduler through a series of algorithms
- Directional scheduling: Node name, node selector
- Affinity scheduling: Node affinity, Pod Affinity, and POD anti-affinity
- Taints and toleration scheduling

#### 1. Directional scheduling

Directional scheduling refers to scheduling a pod to a node we want by declaring a **node name** or **node selector** on the pod. Note that the scheduling is mandatory, which means that even if the target node does not exist, it will be scheduled up, but the POD will fail.

**NodeName**

   Node name Is used to enforce constraints to schedule pods to nodes with the specified name. In this way, the scheduler skips scheduler's scheduling logic and dispatches pods directly to the node with the specified name。

**NodeSelector**

​    Node selector is used to schedule pods to nodes with the specified label added. It is implemented by kubernetes' label-selector mechanism, that is, before pod is created, scheduler uses the match node selector scheduling policy to match the label, find the target node, and then schedule pod to the target node. The matching rule is a mandatory constraint.

```yaml
# example
apiVersion: v1
kind: Pod
metadata:
  name: pod-nodename
  namespace: dev
spec:
  containers:
  - name: nginx
    image: nginx:1.17.1
    
  nodeName: node1 # Specifies scheduling to node1
  # or
  nodeSelector: 
    nodeenv: pro # Specifies scheduling to nodes with the nodeenv=pro label
```

#### Affinity scheduling

In the previous section, we introduced two directional scheduling methods that are very convenient to use, but have a problem. If there are no nodes that meet the criteria, a POD will not run, even if there is a list of nodes available in the cluster, which limits its usage scenarios.

Based on the above problems, Kubernetes also provides affinity scheduling. It is extended on the basis of node selector, and can be configured to preferentially select nodes that meet the conditions for scheduling. If not, it can also be scheduled to nodes that do not meet the conditions, making scheduling more flexible.

Affinity falls into three main categories

- nodeAffinity: Target nodes to solve the problem of which nodes a POD can be scheduled to

- podAffinity :  Target PODS to solve the problem of which existing PODS can be deployed in the same topology domain

- podAntiAffinity :  Target PODS to solve the problem that pods cannot be deployed in the same topology domain as existing pods

Some practical scenarios :

**Affinity** : If two applications frequently interact, it is necessary to use affinity to make the two applications as close as possible to reduce performance loss caused by network communication.

**Anti-affinity** : If multiple copies of applications are deployed, it is necessary to use anti-affinity to disperse application instances and distribute them on each node to improve service availability.

---

##### **NodeAffinity**

```yaml
pod.spec.affinity.nodeAffinity
  requiredDuringSchedulingIgnoredDuringExecution  
  # Node nodes must meet all the specified rules, which is equivalent to a hard limit
  # Node selection list
    nodeSelectorTerms  
      matchFields   # List of node selector requirements by node field
      matchExpressions   # List of node selector requirements by node label (recommended)
        key    
        values 
        operator # Relational operator
  preferredDuringSchedulingIgnoredDuringExecution 
    # Preferentially schedule nodes that meet the specified rules, equivalent to soft limit 
    preference   # A node selector item associated with the corresponding weight
      matchFields   # List of node selector requirements by node field
      matchExpressions   # List of node selector requirements by node label (recommended)
        key    
        values 
        operator 
	weight # The weight
```

```
some operator :

- matchExpressions:
  - key: nodeenv              # exist
    operator: Exists
  - key: nodeenv              # in
    operator: In
    values: ["xxx","yyy"]
  - key: nodeenv              # great than
    operator: Gt
    values: "xxx"
```

```yaml
# example 1
apiVersion: v1
kind: Pod
metadata:
  name: pod-nodeaffinity-required
  namespace: dev
spec:
  containers:
  - name: nginx
    image: nginx:1.17.1
  affinity:  #亲和性设置
    nodeAffinity: #设置node亲和性
      requiredDuringSchedulingIgnoredDuringExecution: # 硬限制
        nodeSelectorTerms:
        - matchExpressions: # 匹配env的值在["xxx","yyy"]中的标签
          - key: nodeenv
            operator: In
            values: ["xxx","yyy"]
        
# example 2
apiVersion: v1
kind: Pod
metadata:
  name: pod-nodeaffinity-preferred
  namespace: dev
spec:
  containers:
  - name: nginx
    image: nginx:1.17.1
  affinity:  #亲和性设置
    nodeAffinity: #设置node亲和性
      preferredDuringSchedulingIgnoredDuringExecution: # 软限制
      - weight: 1
        preference:
          matchExpressions: # 匹配env的值在["xxx","yyy"]中的标签(当前环境没有)
          - key: nodeenv
            operator: In
            values: ["xxx","yyy"]
```

**PodAffinity**

Pod Affinity mainly implements the function of placing a newly created POD in the same area as a reference POD by taking the running POD as a reference.

```
pod.spec.affinity.podAffinity
  requiredDuringSchedulingIgnoredDuringExecution  The hard limit
    namespaces       Specify a namespace that is referenced to pod
    topologyKey      Specify the scheduling scope
    labelSelector    Label selector
      matchExpressions  List of node selector requirements by node label (recommended)
        key    
        values 
        operator 
      matchLabels    
  preferredDuringSchedulingIgnoredDuringExecution The soft limit
    podAffinityTerm  options
      namespaces      
      topologyKey
      labelSelector
        matchExpressions  
          key    
          values 
          operator
        matchLabels 
    weight 
  
The Topology key is used to specify the scheduling scope, for example:
If kubernetes. IO /hostname is specified, node is used to distinguish between nodes
If beta.kubernetes. IO/OS is specified, it is distinguished by the operating system type of the node
```

```yaml
# example
apiVersion: v1
kind: Pod
metadata:
  name: pod-podaffinity-required
  namespace: dev
spec:
  containers:
  - name: nginx
    image: nginx:1.17.1
  affinity: 
    podAffinity: 
      requiredDuringSchedulingIgnoredDuringExecution: 
      - labelSelector:
          matchExpressions: 
          - key: podenv
            operator: In
            values: ["xxx","yyy"]
        topologyKey: kubernetes.io/hostname
```

---

### Taint and tolerance

**Taints**

​    The previous scheduling method is to stand on the pod point of view, by adding attributes to the pod, to determine whether to schedule to the specified node, in fact, we can stand on the node point of view, by adding **stain** attribute on the node, to decide whether to allow pod scheduling.

​    Nodes are stained with pods and have a mutually exclusive relationship with them, thus rejecting pod scheduling and even ejecting existing pods.

##### add label about taints

```bash
# set taint
kubectl taint nodes node1 key=value:effect

# remove taint
kubectl taint nodes node1 key:effect-

# remove all taint
kubectl taint nodes node1 key-
```

- PreferNoSchedule：Kubernetes will try to avoid scheduling pods on nodes with this stain unless there are no other nodes to schedule
- NoSchedule：Kubernetes will not dispatch pods to nodes with the stain, but will not affect existing pods on the current node
- NoExecute：Kubernetes will not dispatch pods to nodes with the stain and will remove existing pods from nodes

---

##### Toleration

We can add a stain to node to deny pod scheduling. However, if we want to dispatch a pod to a node with a stain, what should we do? The answer is to use tolerate.

```yaml
# example
apiVersion: v1
kind: Pod
metadata:
  name: pod-toleration
  namespace: dev
spec:
  containers:
  - name: nginx
    image: nginx:1.17.1
  tolerations:      
  - key: "tag"        
    operator: "Equal" 
    value: "heima"    
    effect: "NoExecute"   
```

