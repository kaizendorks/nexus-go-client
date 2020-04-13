# nexus-go-client
>A minimal Golang client for interacting with the Nexus Repository APIs.

## Description

As the name suggests this is a minimal NexusAPI client written in go for interacting with the APIs defined here: https://help.sonatype.com/repomanager3/rest-and-integration-api

It does not implement any extra logic other than the one supported by the official Nexus API (its more like a proxy that just passes things back and forward), leaving it up to the users to handle advanced logic. The idea is that this reduces the amount of code that is not under the user's control making this library simpler and it easier to keep it in sync with the Nexus APIs.
