# EdgeX Rules Possibilities

A brief exploration of options for the [EdgeX](edgex.discourse.group) Rules Engine

# Summary

As of this writing, the [EdgeX Rules Engine](https://docs.edgexfoundry.org/Ch-RulesEngine.html) is based on [Drools](https://www.drools.org/). The challenge is the footprint for this container is somewhat large for the edge. In a perfect world the replacement would be written in Go or something with a similar small footprint. It also needs to be maintained, not a one off someone threw into the wild.

# Options

See PDF in this repo

# Running the Examples

As part of the exercise, I played a little. Really more to prove to myself that I could express the simple EdgeX IFTT type rules.

## Durable Rules

* [Install and start Redis](https://redis.io/topics/quickstart)
* [Install Durable Rules](https://github.com/jruizgit/rules/blob/master/docs/py/reference.md)
* Run my little example

```shell
python edge-event.py
```

* Test
```shell
curl http://localhost:5000/motortoofastsignal/events -H Content-type:application/json -X POST -d '{"device": "562114e9e4b0385849b96cd8", "parameter": "RPM", "value": 1200}'

curl http://localhost:5000/motortoofastsignal/events -H Content-type:application/json -X POST -d '{"device": "562114e9e4b0385849b96cd8", "parameter": "RPM", "value": 1201}'
  ```

## JsonLogic

* Build

```go
go build
```

* Test

```go
go run edge_event.go
```
