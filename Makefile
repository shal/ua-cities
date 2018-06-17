VERSION ?= $(shell git describe --tags --abbrev=0 2>/dev/null || echo latest)
APP     ?= ua-cities
IMAGE   := ashanaakh/$(APP):$(VERSION)

.PHONY: default build push run

default: build run

run:
	@echo '> Starting "$(APP)...'
	@docker run --name $(APP) \
							-p 8080:8080 \
							-d $(IMAGE)

push: build
	@echo '> Pushing "$(APP)" docker image with version: "$(VERSION)"'
	@docker push $(IMAGE)

build:
	@echo '> Building "$(APP)" docker image...'
	@docker build -t $(IMAGE) .