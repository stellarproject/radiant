![Blackbird](docs/blackbird.png)

# Blackbird
Blackbird is a GRPC proxy service using [Caddy](https://caddyserver.com).

Blackbird uses a datastore to store server info.  By default, there is a simple in-memory datastore.  You can implement whatever you want to integrate with external systems.

# Usage
To start the server, run:

```
$> blackbird
```

This will start both the proxy and GRPC servers.

There is a Go client available to assist in usage:

```go
    client, _ := blackbird.NewClient(addr)

    timeout := time.Second * 30
    upstreams := []string{
        "http://1.2.3.4",
        "http://5.6.7.8",
    }
    opts := []blackbird.AddOpts{
    	blackbird.WithUpstreams(upstreams...),
    	blackbird.WithTimeouts(timeout),
        blackbird.WithTLS,
    }

    // add the server
    _ = client.AddServer(host, opts...)

    // reload the proxy to take effect
    _ = client.Reload()

    // remove the server
    _ = client.RemoveServer(host)

    // reload to take effect
    _ = client.Reload()
```
It is safe to reload as often as you wish.  There is zero downtime for the reload operation.

There is also a CLI that can be used directly or as reference.
