# [Pile2](https://github.com/arapov/pile2)

## Run in OpenShift
```
$ export DOCKER_ID_USER="username"
$ oc new-app $DOCKER_ID_USER/s2i-pile2-centos7~https://github.com/$DOCKER_ID_USER/pile2.git
```

## Manually trigger build in OpenShift
```
$ oc start-build pile2
```

## How to run it?
```
$ npm install
$ npm run init
$ npm run watch
```
