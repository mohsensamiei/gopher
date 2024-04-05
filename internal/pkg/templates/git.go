package templates

const (
	GitIgnore = `
# IDE
.idea
.metals
.vscode

# OS files
.DS_Store
Thumbs.db

# Secret data
.env

# Auto generated
/logs
/build
/vendor
*.pb.go
`
	GitKeep = ``
	GitCI   = `
stages:
  - test
  - build
  - deploy

variables:
  DEPLOY_TARGET:
    value: "none"
    options:
      - "none"
      - "development"
      - "staging"
      - "production"
    description: "The deployment target. Set to 'none' by default."

test:
  stage: test
  image: golang:latest
  before_script:
    - go install golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow@latest
    - go mod download
  script:
    - go vet -vettool="$(which shadow)" ./...
    - go test -v ./...
  rules:
    - if: $CI_PIPELINE_SOURCE == "merge_request_event" || $DEPLOY_TARGET != "none"
  tags:
    - test

build:
  stage: build
  environment:
    name: ${DEPLOY_TARGET}
  before_script:
    - docker login -u ${CI_REGISTRY_USER} -p ${CI_REGISTRY_PASSWORD} ${CI_REGISTRY}
    - export VERSION=$(echo ${CI_COMMIT_TAG:-${CI_COMMIT_SHORT_SHA}}.${CI_PIPELINE_ID})
  script:
    - docker compose -f deploy/docker-compose.build.yml build --no-cache
    - docker compose -f deploy/docker-compose.build.yml push
  only:
    variables:
      - $DEPLOY_TARGET != "none"
  tags:
    - build

deploy:
  stage: deploy
  environment:
    name: ${DEPLOY_TARGET}
  before_script:
    - docker login -u ${ARTIFACTORY_USER} -p ${ARTIFACTORY_PASS} ${DOCKER_ARTIFACTORY}
    - docker login -u ${CI_REGISTRY_USER} -p ${CI_REGISTRY_PASSWORD} ${CI_REGISTRY}
    - export VERSION=$(echo ${CI_COMMIT_TAG:-${CI_COMMIT_SHORT_SHA}}.${CI_PIPELINE_ID})
  script:
    - docker compose -f deploy/docker-compose.deploy.yml up -d
  only:
    variables:
      - $DEPLOY_TARGET != "none"
  tags:
    - ${DEPLOY_TARGET}
`
)
