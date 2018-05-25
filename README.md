# Scavenge and Survive Core

This is the core storage backend and processing server for the
[Scavenge and Survive](https://github.com/Southclaws/ScavengeSurvive) gamemode.

The gamemode communicates with this application via a HTTP API using JSON as a
data encoding.

Aside from providing a method of storing data, this application also processes
some data and makes additional metadata about players and the world available to
the gamemode and any additional applications that are authorised to access it.
