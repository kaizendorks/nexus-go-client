// Copyright (c) 2020 Kaizen Dorks.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

/*
Package nexus can be used for interacting with the Nexus Repository REST API.

As the name suggests this is a minimal Nexus API client written in go for interacting with the APIs defined here:
	https://help.sonatype.com/repomanager3/rest-and-integration-api

It does not implement any extra logic other than the one supported by the official Nexus API
(its more like a proxy that just passes things back and forward), leaving it up to the users to handle advanced logic.
This reduces the amount of code that is not under the user's control, in addition to making this library simpler and easier to keep it in sync with the Nexus APIs.


The most basic usage would be to check if the Nexus server is running and ready to receive requests.
Fot this we would start of by passing in a basic ClientConfig to the NewClient constructor, followed by a call to one of the status API endpoints.

	client := nexus.NewClient(nexus.ClientConfig{
		Host:		  nexusHost,
		Username: nexusUsername,
		Password: nexusPassword,
	})

	err := client.Status.StatusWritable()
	if err != nil {
		// Report Nexus server is not readable/writable.
	}

For control over the HTTP client, add extra options to the ClientConfig.
Example 1: allowing insecure https.
// TODO:
*/
package nexus
