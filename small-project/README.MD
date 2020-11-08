# Description
The service listen on 3009 port by default.
http handler:
POST /api/login
{"login":"admin", "password":"admin"}

You can try to login with another predefined user:
login/password - user/user
# Install&Run
## run PostgreSQL in docker
``~ ./build/build.sh``

## ensure dependencies

``~ go mod tidy``

``~ go mod vendor``

## Finally, run the application.

``go run cmd/main.go``
