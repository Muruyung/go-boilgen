# Go Boilerplate Generator

Go Boilerplate Generator is used to generate core modules of boilerplate such as usecase, service, repository, models, and entity.
It is use Clean Architecture and Domain-Driven Design.
You can generate a whole modules just by running this command.

## Installation

The following command is used for install it into your local

```sh
go install github.com/Muruyung/go-boilgen@latest
```

## Requirement

If you want to generate code for a new project, you have to do init first

```sh
go mod init project-name
```

If you want to generate code for a new project or also an old project, you need to add this dependencies into your project

```sh
# Required for using some utilities, like converter, pagination, and query builder
go-boilgen init

# Required for generate mock (it will used for unit testing)
go install github.com/RanguraGIT/genut@v1.0.0-release
```

**\*P.S: For a mock generator (genut) is created by my co-worker. Thank you to [Agung Maulana Syahputra](https://github.com/RanguraGIT).**

## Command

You can use go-boilgen interactively using this command

```bash
go-boilgen run
```

Use this command for more information about Go-Boilgen.

```bash
go-boilgen --help
```

## License

This project is licensed under the MIT License.
