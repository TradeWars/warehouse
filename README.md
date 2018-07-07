# Scavenge and Survive Core

This is the core storage backend and processing server for the
[Scavenge and Survive](https://github.com/Southclaws/ScavengeSurvive) gamemode.

The gamemode communicates with this application via a HTTP API using JSON as a
data encoding.

Aside from providing a method of storing data, this application also processes
some data and makes additional metadata about players and the world available to
the gamemode and any additional applications that are authorised to access it.

## Security

This is not designed to be a public facing API, hence the lack of proper
security. It is to be run in a private network, accessible only by the Scavenge
and Survive game server and any other applications that are authorised to read
and write data.

With that said, there is still a basic `Authorization` header key check for all
requests.

## Development

To develop, you must fill `.env` and have a running MongoDB instance. You can do
all of this locally by using the makefile.

The environment variables are in the `Config` struct in the `server` package. It
uses `github.com/kelseyhightower/envconfig` to load these from environment
variables and the `split_words` tag so `MongoHost` becomes
`WAREHOUSE_MONGO_HOST`.

Run `make mongodb` to spin up a MongoDB instance and a Mongo-Express container
which is a small web frontend for viewing the MongoDB database.

Once that's all done, you can build with `make fast` which just wraps `go build`
with a version number injection. Run the app by just executing `./scc` there are
no flags or arguments, it reads everything from `.env`.

## Production

Production doesn't differ much from development, just ensure your database has
`--auth` specified and a user specifically for the application. Make sure
`WAREHOUSE_AUTH` is set to something strong.

In Scavenge and Survive, configure the host:port and auth string in the
`settings.ini` file.

## Integration

You can integrate other applications such as web backends to the API. Everything
is documented in the application index, which can be accessed by simply sending
a `GET` request to the server's root path `/`. This documentation includes
example "accepts" and "returns" values, accepts may be either JSON or a
comma-separated list of URL query parameters. All endpoints return a "status"
object with three fields: `result` which contains some object, value, array
related to the request, `success` which is either true or false and `message`
which is some success-related message - if the request failed, this is an error
message and if it succeeded, it should be absent with exception to the `/` path.
