[![Generic badge](https://img.shields.io/badge/development-in&nbsp;progress-orange.svg)](https://shields.io/)

<img align="left" width="225" height="200" src="https://static.wikia.nocookie.net/studio-ghibli/images/b/bb/Sheeta_3.png/revision/latest/scale-to-width-down/350?cb=20221030023634">

# Sheeta

CDK deployment infrastructure can be found here: [link](https://github.com/iancullinane/sheeta-infrastructure)

Cloud bot for discord. Has features for slash commands (auto-generated even), direct chat, cloud deployment, and more to come. Everything in various states of progress. This is currently in heavy development. Originally I utilized a websocket approach, but have switched to a lambda. Use caution.

## Usage

For now this will work best being deployed behind a AWS API Gateway HTTP API v2. The reason for this is the container in ecr uses the ecr build image to simplify development. 

This repo will automatically build to a configured ecr repository via the `drone.yaml`.

### Application Commands

Before you can use a `slash` command in discord it my be created via the API. This is done automatically on commit by drone. The `src/application` package has the definitions as well as commands to build to a discord server.

Use the `-b` flag with the values `create` or `delete` to manage commands locally.

### Build

The drone process will also build a binary, but you can use the provided Dockerfiles.

You can build the binary like so:

`./build/build.sh`

and then:

`./bin/sheeta -t <your discord token>`

You can provide a flag to only build the commands found in `src/appplication`
