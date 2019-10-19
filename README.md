[![Build Status](https://travis-ci.com/shihanng/oauth-demo.svg?branch=develop)](https://travis-ci.com/shihanng/oauth-demo)

## Enviroment Variables

See [OAuth2_Proxy's doc](https://pusher.github.io/oauth2_proxy/auth-configuration#google-auth-provider) for how to obtain the following:
```
export OAUTH2_PROXY_CLIENT_ID=
export OAUTH2_PROXY_COOKIE_SECRET=
export OAUTH2_PROXY_CLIENT_SECRET=
```


## Run
```
docker-compose -f deployments/docker-compose.yml up --build
```
