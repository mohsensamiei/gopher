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
`
	GitKeep = ``
	GitCI   = `
stages:
  - test
  - build

test:
  stage: test
  image: golang:latest
  before_script:
    - go install golang.org/x/tools/go/analysis/passes/shadow/cmd/shadow@latest
    - go mod download
  script:
    - go vet -vettool="$(which shadow)" ./...
    - go test -v ./...
  only:
    - merge_requests
    - tags
  tags:
    - test

build:
  stage: build
  before_script:
    - docker login -u ${CI_REGISTRY_USER} -p ${CI_REGISTRY_PASSWORD} ${CI_REGISTRY}
    - SERVICE=$(echo ${CI_COMMIT_TAG} | cut -d- -f1)
    - VERSION=$(echo ${CI_COMMIT_TAG} | cut -d- -f2)
  script:
    - docker build .
      -f services/${SERVICE}/Dockerfile
      -t ${CI_REGISTRY_IMAGE}/${SERVICE}:${VERSION}
      --build-arg GOPROXY=${GOPROXY}
      --build-arg VERSION=${VERSION}
    - docker push ${CI_REGISTRY_IMAGE}/${SERVICE}:${VERSION}
  only:
    - tags
  tags:
    - build
`
)
