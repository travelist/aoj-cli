# AOJ CLI
A command-line tool for Aizu Online Judge (AOJ)

This is a primitive CLI tool for just doing 3 things:

- Construct a project directory consisting of a boilerplate source code and sample test cases
- Test your solution with the sample cases
- Submit the source code to AOJ System

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

Default config path: `~/.aoj/config.toml`
Templates: `~/.aoj-cli/template.txt`

```toml
[general]
language = "cpp"
username = "username"
password = "password"

[gen]
template_directory = "$HOME/.aoj/templates"
workspace_directory = "$HOME/aoj-workspace"
source_file_name = "main.cpp"

[test]
before_all="g++ main.cpp -o a.out"
before_each=""
command="a.out"
after_each=""
after_all=""

[submit]
source_file_name = "main.cpp"
```

| parameter| description |  |
|----------|-------------|--|
| general.language | programming language | ex: `java`, `cpp` - See the list bellow. |
| general.username | Username | |
| general.password | password | |
| gen.workspace_directory | path to a workspace | source files are generated under this directory |
| gen.source_file_name | file name of boilerplate source code | |
| test.before_all | command to be executed once before all tests | |
| test.before_each | command to be executed before each test | |
| test.test_commands | command to execute a solution | |
| test.after_each | command to be executed after each test | |
| test.after_all | command to be executed once after all tests | |
| submit.source_file_name | source file to be submitted | |

Available options of `general.language`:

- `c`
- `cpp11`
- `cpp14`
- `cpp`
- `java`
- `py`
- `py3`
- (Feel free to open an issue when you need other options)

## TODO

- [ ] Support available languages
- [ ] Remove username and password from configuration file (security)
