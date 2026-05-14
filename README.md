# go-application-template

[![Continuous Integration Status](https://github.com/tombenke/go-application-template/workflows/Continuous%20Integration/badge.svg)](https://github.com/tombenke/go-application-template)
[![Release Status](https://github.com/tombenke/go-application-template/workflows/Release/badge.svg)](https://github.com/tombenke/go-application-template)

## About

This is the go-application-template application.

## How to create a new application from this repo

This repository holds the minimal boilerplate code that is needed to create a new Go microservice.
You can simply create a new application by executing the steps listed below:

1. Create a new repository using this repo as a template.
2. Clone the newly created repository, and step into it.
3. Edit the [`.kickoff.conf`](.kickoff.conf) file that holds the template parameters, and change them according to your needs.
4. Execute the replacing of template parameters, according to the content of the sed file:

   ```bash
       ./.kickoff.sh .
   ```

5. Install the dependencies and the git hook:

   ```bash
       task install
   ```

6. Create a local `.secrets` file that holds your PAT to the private github registry. See more details below.

7. Build the docker image, that is required for integration testing:

    ```bash
        task build-docker
    ```

8. Run the default task that will run the unit tests and the integration tests:

    ```bash
        task
    ```

9. Remove the helper files used for template parameter management:

   ```bash
       rm .kickoff.*
   ```

10. Configure the the GitHub pages for the repo with the __Settings/Pages__ (Source: `'Deploy from branch'`, Branch: `'main, docs/'`),
   then set the API docs link in the README.md file to point the newly created docs pages.

11. Commit the changes you made, and push it to the GitHub repo.

__NOTE__: If you do not have `sed` installed on your machine, you can do the replacement of parameters with your preferred tool, or IDE.

Keep this repository up-to-date and create the new actors using it, instead of always recreating them manually, from scratch.

__NOTE__: Please, do NOT forget to remove this block from the README, and write the right documentation for the newly created module!

__NOTE__: You may run `dos2unix deployments/docker-compose-int.yml` once, to fix endline issue.

## Development

Clone the repository, then install the dependencies and the development tools:

```bash
task install
```

List the tasks:

```bash
task list
```

Build:

```bash
task build
```

To run the application, first start the nats:

```bask
task dc-up
```

And start the application:

```bash
./main
```

### Building the docker image

To build the docker image a _.secrets_ file has to be created and a Github personal access token must be provided with the REPO_PULL key in it:

```
REPO_PULL=<your_ghp_token>
```

Then run the `build-docker` task:

```bash
task docker-build
```

### Running the integration tests

To run the integration tests the docker image must be [built](#building-the-docker-image) first. Then run the `test-int` task:

```bash
task test-int
```

_NOTE:_ The default `task` try to run the `test-int` task as well, so you need to build the docker image,
before run either the integration test or the default task with the standalone `task` command.

## Usage

Get help:

```bash
./main -h
Usage: ./main -h

  -f string
     The log format: json | text (default "json")
  -h Show help message
  -healthcheck-port int
     The HTTP port of the healthcheck endpoints (default 8080)
  -help
     Show help message
  -l string
     The log level: panic | fatal | error | warning | info | debug | trace (default "info")
  -liveness-check-path string
     The path of the liveness check endpoint (default "/live")
  -log-format string
     The log format: json | text (default "json")
  -log-level string
     The log level: panic | fatal | error | warning | info | debug | trace (default "info")
  -p Print configuration parameters
  -print-config
     Print configuration parameters
  -readiness-check-path string
     The path of the readiness check endpoint (default "/ready")

```

## API docs

See the API docs by starting the docs generator and server using the `task docs` task after installation.
