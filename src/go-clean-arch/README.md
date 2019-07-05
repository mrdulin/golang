# serverless-functions-go

## Deploy

### Deploy on Cloud Run

1. Build docker image using `Dockerfile.cloudrun` file

```bash
docker build -f Dockerfile.cloudrun -t sls-fns-go .
```

2. Push the docker image to docker registry

For example, using `GCR`

```bash
docker tag sls-fns-go gcr.io/${projectId}/sls-fns-go
```

```bash
docker push gcr.io/${projectId}/sls-fns-go
```

3. Deploy using `gcloud` sdk

Make sure the version of `gcloud` sdk is `251.0.0`, it seems there is an issue about latest version(`253.0.0`) when deploy resource on cloud run

```bash
gcloud components update --version 251.0.0
```

deploy shell script [here](./scripts/cloud-run/deploy.sh)

### Deploy as Cloud Functions

`cd` to the root directory of the cloud function, for example:

```bash
cd ./infrastructure/functions/getAdPerformanceReport
```

Run `./deploy.sh` script:

```bash
./deploy.sh
```
