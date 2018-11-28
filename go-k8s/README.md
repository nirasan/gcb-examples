
```
# Create k8s cluster
gcloud beta container --project "go-slack-bot-169307" clusters create "gcb-cluster-1" --zone "us-central1-a" --username "admin" --cluster-version "1.9.7-gke.11" --machine-type "f1-micro" --image-type "COS" --disk-type "pd-standard" --disk-size "100" --scopes "https://www.googleapis.com/auth/compute","https://www.googleapis.com/auth/devstorage.read_only","https://www.googleapis.com/auth/logging.write","https://www.googleapis.com/auth/monitoring","https://www.googleapis.com/auth/servicecontrol","https://www.googleapis.com/auth/service.management.readonly","https://www.googleapis.com/auth/trace.append" --num-nodes "3" --enable-cloud-logging --enable-cloud-monitoring --no-enable-ip-alias --network "projects/go-slack-bot-169307/global/networks/default" --subnetwork "projects/go-slack-bot-169307/regions/us-central1/subnetworks/default" --addons HorizontalPodAutoscaling,HttpLoadBalancing,KubernetesDashboard --enable-autoupgrade --enable-autorepair

# Connect k8s cluster
gcloud container clusters get-credentials gcb-cluster-1 --zone us-central1-a --project go-slack-bot-169307

# Apply deployment
kubectl apply -f k8s.yaml
```
