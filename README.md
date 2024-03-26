# go-pubsub-flamework

Go package implementing the `sfomuseum/go-pubsub.Publisher` interface for Flamework log events.

## Documentation

Documentation is incomplete

## Example

### apilogsql://

```
$> go run -mod vendor cmd/subscribe/main.go \
	-subscriber-uri redis://api_log \
	-publisher-uri 'apilogsql://mysql?dsn={DB_USER}:{DB_PSWD}@tcp({DB_HOST}:{DB_PORT})/{DB_TABLE}'
```

## See also:

* https://github.com/sfomuseum/go-pubsub
