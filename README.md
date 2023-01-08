## Description
Examples of my fast coding for job interviews.
(Production ready code looks some different).

## Starting
Start application in docker container.

To start application without docker
-   in first start use command *make imports*
-   use command *make run args="--config ./config/cfg.yaml"*

## Config
Config example
```yaml
http_listen: ":8080"
log_file : "./log.txt"
log_level : "debug"
```
## Usage libs
-	github.com/spf13/cobra v1.2.1: CLI
-	github.com/spf13/viper v1.10.1: unmarshal yaml
-	go.uber.org/zap v1.17.0: logging
-   github.com/gin-gonic/gin: REST API
