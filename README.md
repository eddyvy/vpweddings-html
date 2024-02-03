Deploy GCLOUD:

```
gcloud builds submit --tag gcr.io/PROJECT-ID/vp-eu --project PROJECT-ID
gcloud run deploy --image gcr.io/PROJECT-ID/vp-eu --project PROJECT-ID --platform managed --allow-unauthenticated
```
