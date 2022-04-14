# README #

To test it's required to have Docker and Go installed. If you are using another tools for performance, go is not required to be installed in your machine.

### How to run? ###

* ./buildImages.sh && docker-compose up


### How to test perf? ###

[Bombardier](https://github.com/codesenberg/bombardier)

* go get -u github.com/codesenberg/bombardier
* ~/go/bin/bombardier -c 200 -n 10000 http://localhost:10000/
* ~/go/bin/bombardier -c 200 -n 10000 http://localhost:10001/
