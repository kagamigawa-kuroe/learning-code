### Kubernetes

---

Some components who compose K8s : 

##### In master server node :

- Scheduler : to assign tasks
- Replication controller : to control the number of Replication
- Api server : to provide some api to control k8s
- Etcd : database to implement persistence

##### In other node :

- Kublet : to create Pod in node ( Pod is a cluster of container )
- Kube proxy : Realize pod communication, and load balancing

##### Other :

- CoreDNS  : DNS for node communication
- DASHBORD : B/S access system
- PROMETHEUS : a monitor of k8s
- Elk : log component

---

#### Pod

---

The smallest constituent unit of K8S.

It contains multiple containers, and those containers share the same IP and the data volume.

##### Controller in pod :

ReplicationController/ReplicaSet(in new version) : control the number of replication.

Deployment : Automatic management tool

---

#### Build a cluster

---

Before we start, we need to make sure that time in every server are synchronous.

use command:

``systemtl start chronyd``

Use command ``data`` to check it.

----

##### two way

###### 1. kubeadm :

1. Create a master node, and use ``kubeadm init``
2. Add other nodes to the master, ``kubeadm join {ip of the master:port of the master}``

###### 2. Minikube

---

### Kubeadm

##### install docker

##### Install basic components : kubelet kubeadm kubectl

````shell
#Update the apt package index and install packages needed to use the Kubernetes apt repository:
sudo apt-get update
sudo apt-get install -y apt-transport-https ca-certificates curl

#Download the Google Cloud public signing key:
sudo curl -fsSLo /usr/share/keyrings/kubernetes-archive-keyring.gpg https://packages.cloud.google.com/apt/doc/apt-key.gpg

#Add the Kubernetes apt repository:
echo "deb [signed-by=/usr/share/keyrings/kubernetes-archive-keyring.gpg] https://apt.kubernetes.io/ kubernetes-xenial main" | sudo tee /etc/apt/sources.list.d/kubernetes.list

#Update apt package index, install kubelet, kubeadm and kubectl, and pin their version:
sudo apt-get update
sudo apt-get install -y kubelet kubeadm kubectl
sudo apt-mark hold kubelet kubeadm kubectl
````

----

### Minikube

##### install

```
curl -LO https://storage.googleapis.com/minikube/releases/latest/minikube-linux-amd64
sudo install minikube-linux-amd64 /usr/local/bin/minikube
```

##### start

```shell
#here we start a first k8s node
minikube start
kubectl create deployment hello-minikube --image=k8s.gcr.io/echoserver:1.4
kubectl expose deployment hello-minikube --type=NodePort --port=8080
kubectl get services hello-minikube
minikube service hello-minikube
```

##### a simple example

```
#deploy nginx
kubectl create deployment nginx --image=nginx:1.14-alpine  
kubectl 
#check
kubectl get pods,svc
```

---

### Resource management of K8s

1. use command to manage pod directly

   ``kubectl run nginx-pod --image=nginx:1.17.1 --port=80 ``

2. use configuration file to manage pod

   ``kubectl apply -f nginx-pod.yaml``

3.  use command and configuration file to manage pod

   ``kubectl create -f nginx-pod.yaml``

---

### Some command to manage resources

#### standard format 

``kubectl + command + type + name + flags``

- command : operation like create, get , delete
- type : type of resource like deployment, pod, service
- name : name of resource
- flags : extra parameters

```bash
# get command to have a check
# check inforamtion of all pod
kubectl get pod

# node information
kubectl get nodes

# check one pod
kubectl get podname

# add option -o wide to show more detail
kubectl get podname -o wide
```

```bash
# some basic command
create edit get patch delete explain 

#some command for running and debugging
run expose describe logs 
attach (enter container) 
exec (like exec in docker) 
cp rollout 
scale autoscale (change the number of pod)

#others
apply (use configuration file to manage pod)
label (comment)
```

##### namespace

```
# namespace is something to classify pods
# here is a example

kubectl create ns dev
kubectl run pod --image=nginx:1.17.1 -n dev

kubectl delete pods podname -n dev
kubectl delete ns dev
```

---

### Command + cfg file to manage resources

The core idea is to use a yaml file to manage resource

A example of yaml cfg file :

```yaml
apiVersion: v1
kind: Namespace
metadata: 
	name: dev
	
---

apiVersion: v1
kind: Pod
metadata: 
	name: nginxpod
	namespace: dev
spec:
	containers: 
	- name: nginx-containers
	  image: nginx:1.17.1
```

use command ``create`` :

```
kubectl create -f nginxpod.yaml
```

then we will get feed-back:

```
namespace/dedv created
pod/nginxpod created
```

---

### Only cfg file to manage resources

```
kubectl apply -f nginxpod.yaml
```

