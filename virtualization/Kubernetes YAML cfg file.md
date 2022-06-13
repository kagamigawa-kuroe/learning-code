## Kubernetes

As we known, one pod contains several containers. 

There exist a base (root) container named PAUSE in every Pod, who provides IP for his Pod, and show the status of the whole Pod.

K8s provides us with PAUSE container so it always exists, we needn't to do extra configuration by ourselves 

---

### Pod YAML configuration

some basic contributions:

```YAML
apiVersion: v1     # version, necessary

```

