# AOJ CLI
A command-line tool for [Aizu Online Judge (AOJ)](https://onlinejudge.u-aizu.ac.jp/)

This is a minimal CLI tool for just doing 3 things:

- Create a project directory with a boilerplate code and test cases
- Test your solution with the test cases
- Submit the solution to AOJ

## Installation

```shell
```

## Usage

Initialize configuration (Required only once)

```shell
aoj init
# > Username?
# > Password?
# > Coding Language? (cpp, java, py)
```

Generate a boilerplate code and test cases

```shell
aoj gen <PROBLEM-ID>
# Create related files under "./${PROBLEM-ID}" 
```

Run tests

```shell
aoj test
```

Submit a source code

```shell
aoj submit
```

## Configuration

Default config file path: `~/.aoj-cli/config.toml`
Default template file path: `~/.aoj-cli/template.txt`

```toml
[general]
language = "cpp"
username = "username"
password = "password"

[gen]
template_file = "$HOME/.aoj-cli/template.txt"
destination_file_name = "main.cpp"

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
| gen.template_file | path to a template file | |
| gen.destination_file_name | | |
| test.before_all | command to be executed once before all tests | |
| test.before_each | command to be executed before each test | |
| test.test_commands | command to execute a solution | |
| test.after_each | command to be executed after each test | |
| test.after_all | command to be executed once after all tests | |
| submit.source_file_name | source file to be submitted | |

Available options of `general.language`:

- `c`
- `cpp`
- `cpp11`
- `cpp14`
- `java`
- `py`
- `py3`
- (Feel free to open an issue when you need other options)

## TODO

- [ ] Support available languages
- [ ] Remove username and password from configuration file (security)
