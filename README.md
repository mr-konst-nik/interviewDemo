## Description
Starting web server with setting from yaml with flag *-- config file.name*, otherwise server will starting by defaults.

Config example
```yaml
http_listen: "127.0.0.1:8080"
log_file : "./log/log.txt"
log_level : "debug"
```
## Usage libs
-	github.com/spf13/cobra v1.2.1: CLI
-	github.com/spf13/viper v1.10.1: unmarshal yaml
-	go.uber.org/zap v1.17.0: logging
