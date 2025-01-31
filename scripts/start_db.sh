source .envrc

echo "Pulling PostgreSQL Docker image..."
docker pull "$POSTGRES_IMAGE"

if [ "$(docker ps -aq -f name="${CONTAINER_NAME}")" ]; then
  echo "A container named '$CONTAINER_NAME' already exists."

  # If the container is running, stop it (optional).
  if [ "$(docker ps -aq -f status=running -f name="${CONTAINER_NAME}")" ]; then
    echo "Stopping the running container '${CONTAINER_NAME}'..."
    docker stop "${CONTAINER_NAME}"
  fi

  # Remove the old container (optional).
  echo "Removing the existing container '${CONTAINER_NAME}'..."
  docker rm "${CONTAINER_NAME}"
fi

docker run -d \
  --name "${CONTAINER_NAME}" \
  -p "${HOST_PORT}":"${CONTAINER_PORT}" \
  -e POSTGRES_USER="${POSTGRES_USER}" \
  -e POSTGRES_PASSWORD="${POSTGRES_PASSWORD}" \
  -e POSTGRES_DB="${POSTGRES_DB}" \
  "${POSTGRES_IMAGE}"
