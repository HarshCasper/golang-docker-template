# [Golang-Docker template](https://harshcasper.github.io/golang-docker-template/)

A template project to create a Docker image for a Go application. The example application exposes an HTTP endpoint through a REST API packaged as a static binary. View on [Docker Hub](https://hub.docker.com/repository/docker/harshcasper/golang-docker-template) or [GitHub Container Registry](https://github.com/HarshCasper/golang-docker-template/pkgs/container/golang-docker-template).

## Usage

### Local development

This example project uses [Docker Compose](https://docs.docker.com/compose/overview/) to build and run the example application.

```shell
docker-compose up --build -d
docker-compose logs -f
```

After a kick-start, you can access the application through the following URL: `http://localhost:3000`. You can test the application by running the following command:

```shell
curl http://localhost:3000
```

You can do a local development of the application by running the following command: `go run main.go`.

### Remote development using GitHub Codespaces

[GitHub Codespaces](https://github.com/features/codespaces) is a new feature in GitHub that allows you to run your code in a sandboxed environment on cloud. To run the project on GitHub Codespaces, follow these steps:

1. Fork the repository.
2. Click on the green `Code` button beside the `Go to file` and `Add file` buttons.
3. Click on `Codespaces` in the dropdown menu and click `New Codespace`.
4. Build will start automatically and you can run `go run main.go` to start the application.

### Local development using Devcontainer

[Devcontainer](https://code.visualstudio.com/docs/remote/containers) is a feature in Visual Studio Code that allows you to run and execute your code in a Docker container. To run the project on Devcontainer, follow these steps:

1. Fork and clone the repository.
2. Open the repository in Visual Studio Code and download the Devcontainer extension.
3. Visual Studio Code will prompt you to open the folder in the Devcontainer.
4. On clicking, Visual Studio Code will automatically start the container and you can run `go run main.go` to start the application.

## Tools

The relevant tools for building, linting, and testing the application are used:

1. [Docker](https://www.docker.com/) and [Docker Compose](https://docs.docker.com/compose/overview/) for building and running the application.
2. [Pre-Commit](https://pre-commit.com/) for verifying the linting and formatting rules.
3. [Shellcheck](https://www.shellcheck.net/) for linting the Shell scripts.
4. [Make](https://www.gnu.org/software/make/) for controlling the build, linting, and testing processes.
5. [Hurl](https://hurl.dev/) for end-to-end testing by running HTTP requests against the application.
6. [GitHub Actions](https://github.com/features/actions) for running the build, linting, and testing processes on CI/CD.
7. [Devcontainer](https://code.visualstudio.com/docs/remote/containers) for running the application on a Docker container for local development.
8. [GitHub Codespaces](https://github.com/features/codespaces) for running the application on GitHub Codespaces for remote development.

## License

[MIT](LICENSE)
