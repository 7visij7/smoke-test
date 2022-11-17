# Post-deploy-smoke-test
> Smoke testing allows to verify that REST API perform correclty.
> 
# Configuration
> There is [config.yaml](https://github.com/7visij7/smoke-test/blob/master/smoke-test/post-deployment-smoke-tests/config/http.yml) where set up of checks with different parametrs.
> Available next params:
+ *base_url* - base url for checks
+ *path* - path of location
+ *method* - HTTP Request Methods
+ *response_code* - expected response code
+ *response_body* - expected response body
+ *playload*
+ *auth* - authorization parameters 
  + *login*
  +  *password*
+ *headers* - HTTP headers. You can add any headers. Below examples:
  + *Content-Type: application/json*
  + *Accept: application/json*
  + *Authorization: SuperToken*


## Build and run application
```Bash
go build -o notification-center
./notification-center
```
---
## Docker

> Build Docker image from a [Dockerfile](https://github.com/7visij7/smoke-test/blob/master/smoke-test/helm/templates/test/smoke-test.yaml)
```
docker build -t IMAGENAME
```
> Start application
```
docker run -it --rm IMAGENAME
```

## Helm test
> Can use post-deploy-smoke-test to execute test after delpoy with Helm. For this add [manifest](https://github.com/7visij7/smoke-test/blob/master/Dockerfile) to your helm chart.
