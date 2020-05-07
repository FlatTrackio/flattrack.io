# staging.flattrack.io

> Deploying the staging site

```
kubectl create ns flattrackio-site-home-staging

kubectl -n create secret generic flattrackio-postgres-secret --from-literal=postgres-password=MYPASSWORD
```

```
kubectl apply -k .
```
