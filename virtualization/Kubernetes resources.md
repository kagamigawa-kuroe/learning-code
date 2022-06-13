## Kubernetes

5 basic resources in K8s :

- Namespace
- Pod
- Label
- Deployment
- Service

---

### Namespace

Namespace is a method used to group pods in K8s.

Pods in differnet group can't access each other.

##### How to check your namespace

```
# check all namespaces
╭─hongzhe@hongzhe-OMEN--HP ~ 
╰─$ kubectl get ns
NAME              STATUS   AGE
default           Active   23h  default namespace
kube-node-lease   Active   23h
kube-public       Active   23h  public ns, everyone can access
kube-system       Active   23h  all resource created by system

# check certain ns
╭─hongzhe@hongzhe-OMEN--HP ~ 
╰─$ kubectl get ns default                                                     
NAME      STATUS   AGE
default   Active   23h

# check pods in a ns
╭─hongzhe@hongzhe-OMEN--HP ~ 
╰─$ kubectl get pods -n kube-system
NAME                               READY   STATUS    RESTARTS        AGE
coredns-64897985d-7qz5v            1/1     Running   2 (35m ago)     23h
etcd-minikube                      1/1     Running   2 (35m ago)     23h
kube-apiserver-minikube            1/1     Running   2 (35m ago)     23h
kube-controller-manager-minikube   1/1     Running   2 (3h26m ago)   23h
kube-proxy-qqkd4                   1/1     Running   2 (35m ago)     23h
kube-scheduler-minikube            1/1     Running   2 (35m ago)     23h
storage-provisioner                1/1     Running   5 (33m ago)     23h

# show infomation of a ns
╭─hongzhe@hongzhe-OMEN--HP ~ 
╰─$ kubectl describe ns default    
Name:         default
Labels:       kubernetes.io/metadata.name=default
Annotations:  <none>
Status:       Active

No resource quota.
No LimitRange resource.
```

##### crud namespace

```
╭─hongzhe@hongzhe-OMEN--HP ~ 
╰─$ kubectl create ns dev                                    
namespace/dev created

╭─hongzhe@hongzhe-OMEN--HP ~ 
╰─$ kubectl get ns dev   
NAME   STATUS   AGE
dev    Active   7s

╭─hongzhe@hongzhe-OMEN--HP ~ 
╰─$ kubectl delete ns dev
namespace "dev" deleted
```

##### Use yaml file

```yaml
apiVersion: v1
kind: Namespacec
metadata:
	name: dev
```

---

### Pod

#### create pods

```
╭─hongzhe@hongzhe-OMEN--HP ~ 
╰─$ kubectl run nginx2 --image=nginx:1.17.1 --port=80 --namespace dev
pod/nginx2 created
```

- --image : to refer image
- --port : to choose port
- --namespace : to refer ns

#### Check pods

```
╭─hongzhe@hongzhe-OMEN--HP ~ 
╰─$ kubectl get pod -n dev -o wide
NAME     READY   STATUS    RESTARTS   AGE    IP           NODE       NOMINATED NODE   READINESS GATES
nginx2   1/1     Running   0          8m4s   172.17.0.5   minikube   <none>           <none>
```

### ！！！Pod IP , cluster IP, and real IP

Cluster IP , it's ip of service：

```
╭─hongzhe@hongzhe-OMEN--HP ~ 
╰─$ kubectl get svc                                                                        
NAME             TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)          AGE
hello-minikube   NodePort    10.107.31.169   <none>        8080:31588/TCP   24h
kubernetes       ClusterIP   10.96.0.1       <none>        443/TCP          24h
nginx            NodePort    10.105.52.18    <none>        80:31764/TCP     24h
```

Pod IP, it't IP of Pod :

```
╭─hongzhe@hongzhe-OMEN--HP ~ 
╰─$ kubectl get pods -o wide                                                                 
NAME                              READY   STATUS    RESTARTS      AGE   IP           NODE       NOMINATED NODE   READINESS GATES
hello-minikube-7bc9d7884c-22cgk   1/1     Running   2 (90m ago)   24h   172.17.0.3   minikube   <none>           <none>
nginx-7cbb8cd5d8-n699c            1/1     Running   2 (90m ago)   23h   172.17.0.4   minikube   <none>           <none>
```

Service includes a series of pod. 

Normally, if we want access a service, we mapping the port of service IP (cluster IP) to real (outside) IP, so that we can access it from outside.

Obviously if we enter into a pod, we can access to other pods in the same namespace by pod IP.

---

 ### Label

Label is also a way to group, that means to assgin a group identification to each pod, we can select and distinguish pods according to it.

A source can own more than one label.

##### Create labels

```bash
╭─hongzhe@hongzhe-OMEN--HP ~ 
╰─$ kubectl run nginx2 --image=nginx:1.17.1 --port=82 -n dev
pod/nginx2 created

# --show-labels
╭─hongzhe@hongzhe-OMEN--HP ~ 
╰─$ kubectl get pod -n dev --show-labels                                                    
NAME     READY   STATUS    RESTARTS   AGE   LABELS
nginx2   1/1     Running   0          33s   run=nginx2

# set label
╭─hongzhe@hongzhe-OMEN--HP ~ 
╰─$ kubectl label pod nginx2 -n dev version=1.0                                             pod/nginx2 labeled

# check
╭─hongzhe@hongzhe-OMEN--HP ~ 
╰─$ kubectl get pod -n dev --show-labels       
NAME     READY   STATUS    RESTARTS   AGE    LABELS
nginx2   1/1     Running   0          2m8s   run=nginx2,version=1.0

# update
╭─hongzhe@hongzhe-OMEN--HP ~ 
╰─$ kubectl label pod nginx2 -n dev version=2.0 --overwrite
pod/nginx2 labeled
```

##### select by label

```bash
╭─hongzhe@hongzhe-OMEN--HP ~ 
╰─$ kubectl get pod -l "version=1.0" -n dev --show-labels  
NAME     READY   STATUS    RESTARTS   AGE   LABELS
nginx3   1/1     Running   0          61s   run=nginx3,version=1.0

╭─hongzhe@hongzhe-OMEN--HP ~ 
╰─$ kubectl get pod -l "version!=1.0" -n dev --show-labels
NAME     READY   STATUS    RESTARTS   AGE     LABELS
nginx2   1/1     Running   0          5m17s   run=nginx2,version=2.0
```

##### delete label

```bash
╭─hongzhe@hongzhe-OMEN--HP ~ 
╰─$ kubectl label pod nginx2 -n dev version-                
pod/nginx2 unlabeled

╭─hongzhe@hongzhe-OMEN--HP ~ 
╰─$ kubectl get pod -n dev --show-labels                    
NAME     READY   STATUS    RESTARTS   AGE     LABELS
nginx2   1/1     Running   0          6m30s   run=nginx2
nginx3   1/1     Running   0          2m31s   run=nginx3,version=1.0
```

---

### deployment

A way to management pod who will help you create/delete pod automaticlly.

##### create

```
╭─hongzhe@hongzhe-OMEN--HP ~ 
╰─$ kubectl create deployment nginx --image=nginx:1.17 --port=80 --replicas=3 -n dev
deployment.apps/nginx created

╭─hongzhe@hongzhe-OMEN--HP ~ 
╰─$ kubectl get pod -n dev                                                          
NAME                     READY   STATUS              RESTARTS   AGE
nginx-5757b68bb6-5hlkh   0/1     ContainerCreating   0          7s
nginx-5757b68bb6-gmvvs   0/1     ContainerCreating   0          7s
nginx-5757b68bb6-t8xlt   0/1     ContainerCreating   0          7s

# when you delete anyone of them, deployment will create another one instead it 
```

##### delete

```
╭─hongzhe@hongzhe-OMEN--HP ~ 
╰─$ kubectl get deploy -n dev                                                              
NAME    READY   UP-TO-DATE   AVAILABLE   AGE
nginx   3/3     3            3           4m14s

╭─hongzhe@hongzhe-OMEN--HP ~ 
╰─$ kubectl delete deploy nginx  -n dev                                                    
deployment.apps "nginx" deleted

╭─hongzhe@hongzhe-OMEN--HP ~ 
╰─$ kubectl get pod -n dev                                                          
No resources found in dev namespace.
```

---

### Service

We have already mentioned it before, service is a group of pod who provide interface to outside togeteher.

##### expose service

```
╭─hongzhe@hongzhe-OMEN--HP ~ 
╰─$ kubectl expose deploy nginx --name=svc-nginx1 --type=ClusterIP --port=80 --target-port=80 -n dev
service/svc-nginx1 exposed

#check
╭─hongzhe@hongzhe-OMEN--HP ~ 
╰─$ kubectl get svc svc-nginx1 -n dev -o wide                                               
NAME         TYPE        CLUSTER-IP      EXTERNAL-IP   PORT(S)   AGE   SELECTOR
svc-nginx1   ClusterIP   10.97.192.148   <none>        80/TCP    38s   app=nginx
```





