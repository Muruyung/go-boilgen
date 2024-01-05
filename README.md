# Go Boilerplate Generator

Go Boilerplate Generator is used to generate core modules of boilerplate such as usecase, service, repository, models, and entity.
It is use Clean Architecture and Domain-Driven Design.
You can generate a whole modules just by running this command.

## Installation

The following command is used for install it into your local

```sh {"id":"01HK4A6PCACAJVWTMCQ8WCY8AP"}
go install github.com/Muruyung/go-boilgen@latest
```

## Requirement

If you want to generate code for a new project, you have to do init first

```sh {"id":"01HK4A6PCACAJVWTMCQ9M2TXCP"}
go mod init project-name
```

If you want to generate code for a new project or also an old project, you need to add this dependencies into your project

```sh {"id":"01HK4A6PCACAJVWTMCQAQM3ZGD"}
# Required for using some utilities, like converter, pagination, and query builder
go-boilgen init

# Required for generate mock (it will used for unit testing)
go install github.com/RanguraGIT/genut@v1.0.0-release
```

**\*P.S: For a mock generator (genut) is created by my co-worker. Thank you to [Agung Maulana Syahputra](https://github.com/RanguraGIT).**

## Command

**NEW**: Now you can use go-boilgen interactively using this command

```sh {"id":"01HK4A6PCACAJVWTMCQEMCPJKH"}
go-boilgen run
```

This is example basic command using default method

```bash {"id":"01HK4A6PCACAJVWTMCQEVTH1HK"}
go-boilgen -s service_name -n module_name -f "field_name1:data_type,field_name2:data_type" -m "method1,method2"
```

You can just generate entity only using flag **--entity-only**.

```bash {"id":"01HK4A6PCACAJVWTMCQJRV8PSM"}
go-boilgen -s service_name -n module_name -f "field_name1:data_type,field_name2:data_type" --entity-only
```

If you want to generate custom method, you must use the flag **-c** (or **--custom-method**), **-p** (or **--params**), and **-r** (or **--return**, optional)

```bash {"id":"01HK4A6PCACAJVWTMCQKM4AE1P"}
go-boilgen -s service_name -n module_name -c custom_method_name -p "field_name1:data_type,field_name2:data_type" -r "field_name1:data_type,field_name2:data_type"
```

If you want to use CQRS pattern, just use the flag **--cqrs**

```bash {"id":"01HK4A6PCACAJVWTMCQNFFQANK"}
go-boilgen --cqrs -s service_name -n module_name -f "field_name1:data_type,field_name2:data_type" -m "method1,method2"
```

Use this command for more information about Go-Boilgen.

```bash {"id":"01HK4A6PCACAJVWTMCQRV9QVKQ"}
go-boilgen --help
```

## License

This project is licensed under the MIT License.
