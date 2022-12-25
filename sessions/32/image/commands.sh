cd image/

# Build the image
docker build -f Dockerfile -t 7learn-k8s-content-gen:v3 .

# Load the image into the K8s cluster
# Replace <YOUR_CLUSTER_NAME> with your kind cluster name (kind get clusters)
kind load docker-image 7learn-k8s-content-gen:v3 --name <YOUR_CLUSTER_NAME>

cd ..
