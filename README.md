# slash-command-random

This repository contains a [serverless](https://serverless.com/) function exposing an HTTP endpoint
to generate random integers. It was designed as a proof-of-concept for a Mattermost tech talk on
slash commands.

## Deployment

### Serverless

See [Getting Started with Serverless](https://serverless.com/framework/docs/getting-started/) for
instructions on how to deploy this function to the compute provider of your choice.

This repository is configured to use AWS Lambda, but feel free to modify for your own purposes.

### Compiling

Simply invoke `make` to generate the requisite binaries:

    make

Then deploy using serverless:

    serverless deploy

Once deployed, serverless will provide the URL to an HTTP endpoint against which the function
can be invoked.

## API

The function integrates with Mattermost as a slash command, parsing a `text` query string parameter
in respond to an HTTP GET request.

## Empty or missing text parameter

The API returns a random integer between 0 and 100, inclusive.

## Single integer parameter

The API returns a random integer between 0 and the given integer, inclusive.

Note that the associated tech talk video incorrectly suggested in the autocomplete help text that 
this was a minimum, not a maximum. 

## Single string parameter: `die` or `dice`

The API returns a random integer between 1 and 6, inclusive.

## Two integer parameters, separated by a space

The API returns a random integer in the inclusive range given by the input parameters.
