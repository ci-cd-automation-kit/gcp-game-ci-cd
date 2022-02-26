apiVersion: "agones.dev/v1"
kind: Fleet
metadata:
  name: xonotic
spec:
  replicas: 2
  strategy:
    type: Recreate
  template:
    spec:
      ports:
      - name: default
        containerPort: 26000
      health:
        initialDelaySeconds: 30
        periodSeconds: 60
      template:
        spec:
          containers:
          - name: xonotic
            image: gcr.io/cn-gaming-cicd/hello-cloudbuild-webhook:COMMIT_SHA

            