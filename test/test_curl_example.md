## Add Worker

`curl -X POST http://192.168.0.100:8090/addworker -H 'Content-Type: application/json' -d '{"ip":"192.168.0.101","name":"oak", "tags" : ["123", "get"]}'`

# Get Workers

`curl -X GET http://192.168.0.100:8090/getworkers`