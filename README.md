# go-example-tuto04
Program used by my [tutorial](https://github.com/richardpct/aws-terraform-tuto04)  
It creates a web server and displays two stuffs:

* the environment
* the count of requests using a redis server

# usage
    $ go get github.com/richardpct/go-example-tuto04
    $ go/bin/go-example-tuto04 -redishost ${database_host} -redispass ${database_pass} -env ${environment}
