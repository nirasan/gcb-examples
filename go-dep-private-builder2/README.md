# Preparing

* Enable Cloud KMS
* Add permission to Cloud Build Service Account to use Cloud KMS
* Get github.com Access Token
* Set github.com Access Token to Cloud Build Trigger Variables named 'GITHUB_ACCESS_TOKEN'

# How to run

```
gcloud builds submit --config=go-dep-private-builder2/cloudbuild.yaml \
    --substitutions=_GITHUB_ACCESS_TOKEN=""
```
