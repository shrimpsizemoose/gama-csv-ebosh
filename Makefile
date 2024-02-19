REGISTRY=ghcr.io
NAMESPACE=shrimpsizemoose
APP=gama-csv-ebosh-5678
VERSION=$(shell git describe --tags)

SERVICE_TAG=${REGISTRY}/${NAMESPACE}/${APP}:${VERSION}

docker-build:
	docker build -t ${SERVICE_TAG} .

docker-run:
	docker run -p 5678:5678 ${SERVICE_TAG}

docker-push:
	docker push ${SERVICE_TAG}
