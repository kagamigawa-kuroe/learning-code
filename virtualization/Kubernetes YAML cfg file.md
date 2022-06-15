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



