# Microservice Using Go

This project containing simple example about how microservice work. The idea is we have a `config.yaml` that containing `server` which the value of service that will be made. and then make a route `v0,v1,...,vn` using `createRouter()` from `controller.go`.

endpoint :

```
127.0.0.1:5000/v0/hello-world
127.0.0.1:5000/v1/hello-world
...
127.0.0.1:5000/v{n}/hello-world
```

example : 

![](./img/1.png)
![](./img/2.png)