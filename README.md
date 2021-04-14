# testt

This is a simple url shortener service.

You can provide a long url through the graph query and you'd get 
a short url on a specified host that redirects to the inital long url

<img width="1620" alt="Screenshot 2021-04-14 at 17 37 07" src="https://user-images.githubusercontent.com/46886694/114749966-635bca00-9d4b-11eb-90d3-ea85915bcac6.png">


```shell script
# To generate graphql schemas
# Note: see server/graph/schemas and also ./gqlgen.yml

$ make schema 
```

```shell script
# to start the service locally

$ make local 
```
```shell script
# For tests

$ make test 
```

```shell script
# to build binary

$ make build 
```
