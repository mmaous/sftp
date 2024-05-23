.PHONY: all build run clean


# Name of the Docker image
DOCKER_REPO := mmaous
IMAGE_NAME := sftp
TAG := v1.0.0

# envs used to run transfer
X_HOST=$(HOST)
X_PORT=$(PORT)
X_VM_USERNAME=$(VM_USERNAME)
X_VM_PASSWORD=$(VM_PASSWORD)
X_LOCAL_FOLDER_PATH=$(LOCAL_FOLDER_PATH)
X_REMOTE_FOLDER_PATH=$(REMOTE_FOLDER_PATH)

# Build the Docker image
docker-build:
		docker build -t $(DOCKER_REPO)/$(IMAGE_NAME):$(TAG) .

# Run the Docker container with environment variables
docker-run:
		docker run --rm \
		-e HOST=$(X_HOST) \
		-e PORT=$(X_PORT) \
		-e VM_USERNAME=$(X_VM_USERNAME) \
		-e VM_PASSWORD=$(X_VM_PASSWORD) \
		-e LOCAL_FOLDER_PATH=/sftp-data \
		-e REMOTE_FOLDER_PATH=$(X_REMOTE_FOLDER_PATH) \
		-v /$(X_LOCAL_FOLDER_PATH):/sftp-data \
		$(DOCKER_REPO)/$(IMAGE_NAME):$(TAG)

# Clean up Docker images and containers
docker-push:
		docker push $(DOCKER_REPO)/$(IMAGE_NAME):$(TAG)
