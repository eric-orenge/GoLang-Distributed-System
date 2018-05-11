# Distributed System Working with Flexible Data.


Includes: Module to collect votes from twitter statuses, another module to count the votes, api and web  modules.
##  Background info.
Open and maintain a connection to Twitter's streaming APIs looking for any mention of the options, if  any, add mentioned option into NSQ que awaiting to 
be consumed ie counted and saved on a persistent disk. An api  then exposes the results accessible  via web service. Abstraction of these modules makes it possible to run
one module without necessarily the other module running.

## A little setup

Install NSQ

> brew install nsq

Install mongodb

> sudo dnf install mongodb (centos,fedora)

## Get dependencies

> go get gopkg.in/tylerb/graceful.v1


> go get gopkg.in/mgo.v2


> go get github.com/bitly/go-nsq


> go get github.com/garyburd/go-oauth/oauth


> go get github.com/joeshaw/envdecode

## How to run

In one terminal

> nsqlookupd

In another terminal

> nsqd --lookupd-tcp-address=localhost:4160

Start mongod(daemon) to do heavy lifting on behalf of mongo - in another terminal

> mongod

Run counter to get number of votes cast per option.

```
cd counter
go build –o counter
./counter

```
Stream  tweets while checking options are available, if available,  cast vote. Be sure that you have the appropriate environment variables set(consumer key,consumer secret, token key and token secret), otherwise you will see errors when you run the program

```
cd ../twittervotes
go build –o twittervotes
./twittervotes

```

Api to  make the results available for viewing

```
cd ../api
go build –o api
./api

```

Web service to  consume the data

```

cd ../web
go build –o web
./web

```

## Results

View polls at (http://localhost:8081) you can also add and delete a polls

#### merciArsene


