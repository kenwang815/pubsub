# Simple pubsub demo
Starting the main program automatically subscribe to the sensor topic, and the sensor publish the data to the broker. Subscribers will get sensor data through the broker.

## Quick start
```
go run main.go
Channel: sensorChannel, Topic: sensor, DataEvent: { "id": a2a2054b-1231-46a2-b1d0-f26b714892de, "name": bedroom, "temperature": 22.813981, "humidity": 68.215273, "battery": 95, "timestamp": 2020-12-20T15:53:40+08:00 }
Channel: sensorChannel, Topic: sensor, DataEvent: { "id": a2a2054b-1231-46a2-b1d0-f26b714892de, "name": bedroom, "temperature": 22.313143, "humidity": 52.739125, "battery": 95, "timestamp": 2020-12-20T15:53:41+08:00 }
...
...
```