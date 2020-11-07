# Install&Run from scratch
## swagger as a docker image
``~ docker pull quay.io/goswagger/swagger``

``~ alias swagger="docker run --rm -it -e GOPATH=$HOME/go:/go -v $HOME:$HOME -w $(pwd) quay.io/goswagger/swagger"``

## Validate swagger spec and generate client code

``~ swagger validate swagger/client-spec.yaml``

``~ swagger generate client -f swagger/client-spec.yaml``

## Finally, run the application.

``go run cmd/main.go``
