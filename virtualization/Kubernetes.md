### Kubernetes

---

Some components who compose K8s : 

##### In master server :

- Scheduler : to assign tasks
- Replication controller : to control the number of Replication
- Api server : to provide some api to control k8s
- Etcd : database to implement persistence

##### In node :

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

###### 2. Install every components by ourselves

.....

---

##### components :

1. install docker

2. Install basic components : kubelet kubeadm kubectl

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
   
   #edit some configuration in kublet
   vim /usr/lib/systemd/system/kubelet.service
   
   ```
   KUBELET_CGROUP_ARGS="--cgroup-driver=systemd"
   KUBE_PROXY_MODE="ipvs"
   ```
   ````

----

#### use kubeadm to initialize

---

1. Use command ``kubeadm config images list``  to check vertion information.

   You can use Docker to download components with different versions by yourselves

   ```
   root@debian:~# kubeadm config images list
   k8s.gcr.io/kube-apiserver:v1.24.1
   k8s.gcr.io/kube-controller-manager:v1.24.1
   k8s.gcr.io/kube-scheduler:v1.24.1
   k8s.gcr.io/kube-proxy:v1.24.1
   k8s.gcr.io/pause:3.7
   k8s.gcr.io/etcd:3.5.3-0
   k8s.gcr.io/coredns/coredns:v1.8.6
   ```

2. In master :
   
