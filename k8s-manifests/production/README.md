# flattrack.io

> Deploying the production site

```
kubectl create ns flattrackio-home

kubectl -n flattrackio-site create secret generic flattrackio-postgres-secret --from-literal=postgres-password=MYPASSWORD
```

```
kubectl apply -k .
```
