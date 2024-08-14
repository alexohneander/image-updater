# image-updater

## Title

### Development

If you are running a Minikube cluster, you can build this image directly on the Docker engine of the Minikube node without pushing it to a registry. To build the image on Minikube:

```bash
eval $(minikube docker-env)
docker build -t image-updater .
```

If you have RBAC enabled on your cluster, use the following snippet to create role binding which will grant the default service account view permissions.

```bash
kubectl create clusterrolebinding default-view --clusterrole=view --serviceaccount=default:default
```

Then, run the image in a Pod with a single instance Deployment:

```bash
kubectl run --rm -i image-updater-demo --image=image-updater

There are 4 pods in the cluster
There are 4 pods in the cluster
There are 4 pods in the cluster
...
```
