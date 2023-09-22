# Go Boilerplate Generator

Go Boilerplate Generator is used to generate core modules of boilerplate such as usecase, service, repository, models, and entity.
It is use Clean Architecture and Domain-Driven Design.
You can generate a whole modules just by running this command.

## Installation

The following command is used for install it into your local

```sh
go install github.com/Muruyung/go-boilgen@latest
```

## Command

This is example basic command using default method

```bash
go-boilgen -s service_name -n module_name -f "field_name1:data_type,field_name2:data_type" -m "method1,method2"
```

You can just generate entity only using flag **--entity-only**.

```bash
go-boilgen -s service_name -n module_name -f "field_name1:data_type,field_name2:data_type" --entity-only
```

If you want to generate custom method, you must use the flag **-c** (or **--custom-method**), **-p** (or **--params**), and **-r** (or **--return**, optional)

```bash
go-boilgen -s service_name -n module_name -c custom_method_name -p "field_name1:data_type,field_name2:data_type" -r "field_name1:data_type,field_name2:data_type"
```

Use this command for more information about Go-Boilgen.

```bash
go-boilgen --help
```
