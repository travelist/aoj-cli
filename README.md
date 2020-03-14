# aoj
A command-line tool for Aizu Online Judge (AOJ)

This is a primitive CLI tool for just doing 3 things:

(1) Construct a project directory consisting of a boilerplate source code and sample test cases
(2) Test your solution with the given sample cases
(3) Submit the source code to AOJ system.

## Installation

```shell
```

## Usage

Generate test cases

```shell
aoj gen [PROBLEM-ID]
```

Run tests

```shell
aoj test
```

Submit a source code

```shell
aoj submit (-f filenpath)
```

## Configuration

default config path: `$HOME/.aoj/config.toml`

```toml
username = "username"
password = "password"
language = "cpp"

[gen]
template_directory = "$HOME/.aoj/templates"
workspace_directory = "/path/to/workspace/directory"
source_file_name = "main.cpp"

[test]
before_all=""
before_each=""
test_command=""
after_each=""
after_all=""

[submit]
source_file_name = "main.cpp"
```

| parameter| description |  |
|----------|-------------|--|
| username | Username | |
| password | password | |
| language | programming language | ex: `java`, `cpp` - See the list bellow. |
| gen.template_directory | path to the template directory |  |
| gen.workspace_directory | path to a workspace | source files are generated under this directory |
| gen.source_file_name | file name of source code | |
| test.before_all | command to be executed once before all tests | |
| test.before_each | command to be executed before each test | |
| test.test_commands | command to execute a solution | |
| test.after_each | command to be executed after each test | |
| test.after_all | command to be executed once after all tests | |
| submit.source_file_name | source file to be submitted | |

## TODO

[ ] Remove username and password from configuration file (security)
