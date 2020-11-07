# Description
The service make some API requests with predefined parameters to [currencies exchange API](exchangeratesapi.io).
There is no way to change those parameters by passing as program arguments or somehow else except changing the code.
# Install&Run from scratch
## swagger as a docker image
``~ docker pull quay.io/goswagger/swagger``

``~ alias swagger="docker run --rm -it -e GOPATH=$HOME/go:/go -v $HOME:$HOME -w $(pwd) quay.io/goswagger/swagger"``

## Validate swagger spec and generate client code

``~ swagger validate swagger/client-spec.yaml``

``~ swagger generate client -f swagger/client-spec.yaml``

## Finally, run the application.

``go run cmd/main.go``
