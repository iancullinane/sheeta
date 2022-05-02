[![Generic badge](https://img.shields.io/badge/development-in&nbsp;progress-orange.svg)](https://shields.io/)

<img align="left" width="225" height="200" src="https://i.postimg.cc/HLGxXbdy/sheeta.png">

# Sheeta

CDK deployment infrastructure can be found here: [link](https://github.com/iancullinane/sheeta-infrastructure)

Cloud bot for discord. Has features for slash commands (auto-generated even), direct chat, cloud deployment, and more to come. Everything in various states of progress. This is currently in heavy development. Originally I utilized a websocket approach, but have switched to a lambda. Use caution.

## Usage

This repo will automatically build to a configured ecr repository via the `drone.yaml`.

The drone process will also build a binary, but you can also use the provided Dockerfiles.

You can build the binary like so:

`./build/build.sh`

and then:

`./bin/sheeta -t <your discord token>`

You can provide a flag to only build the commands found in `src/appplication`
