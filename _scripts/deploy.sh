docker login -e="$QUAY_EMAIL" -u="$QUAY_USERNAME" -p="$QUAY_PASSWORD" quay.io
DOCKER_REPO=quay.io/ DOCKER_TAG=0.0.1 make -C .. build docker-build docker-push
