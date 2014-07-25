# logostash

Emit logstash JSON events from Golang types.

## Usage

First, import the package:

```go
import (
       "github.com/DavidHuie/logostash"
)
```

Create a client, specifying the type of connection to use
and the address of the logstash server:

```go
client, err := logostash.NewClient("udp", "my-logstash-server:8888")
```

Now, emit any JSON-encodeable type:

```go
message := map[string]string{
	"type":    "my-application",
	"message": "something's f'ed up",
}
err := client.SendJson(message)
```

## Copyright

Copyright (c) 2014 David Huie. See LICENSE.txt for further details.
