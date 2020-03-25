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
```

Generate a boilerplate code and test cases

```shell
aoj gen <PROBLEM-ID>
# Create a problem directory under current dir
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
language = "C++"
source_file_name = "main.cpp"
```

| parameter| description |  |
|----------|-------------|--|
| gen.template_file | path to a template file | |
| gen.destination_file_name | file name of a generated code | |
| test.before_all | command to be executed once before all tests | |
| test.before_each | command to be executed before each test | |
| test.test_commands | command to execute a solution | |
| test.after_each | command to be executed after each test | |
| test.after_all | command to be executed once after all tests | |
| submit.language | programming language | ex: `JAVA`, `C++` |
| submit.source_file_name | source file to be submitted | |

## TODO

- [ ] `TLE` Check
