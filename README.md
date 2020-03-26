# AOJ CLI [![CircleCI](https://circleci.com/gh/travelist/aoj-cli.svg?style=svg)](https://circleci.com/gh/travelist/aoj-cli)

[<p dir='rtl' align="right;">日本語(Japanese)</p>](https://github.com/travelist/aoj-cli/#%E6%97%A5%E6%9C%AC%E8%AA%9Ejapanese)

A command-line tool for [Aizu Online Judge (AOJ)](https://onlinejudge.u-aizu.ac.jp/)

This is a minimal CLI tool for just doing 3 things:

- Create a project directory with a boilerplate code and test cases
- Test your solution with the test cases
- Submit the solution to AOJ

<img src="https://raw.githubusercontent.com/travelist/aoj-cli/master/doc/animation.gif">

## Installation

```shell
brew tap travelist/homebrew-aoj-cli
brew install aoj
```

## Usage

Initialize configuration (Required only once)

```shell
aoj init
```

Generate a boilerplate code and test cases

```shell
aoj gen [PROBLEM-ID]
```

Test the solution

```shell
# cd ./[PROBLEM-ID]
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

---

# 日本語(Japanese)

AOJ 用のコマンドラインツールです。主に以下のことができます:

- テストサンプルのダウンロードと、ディレクトリの作成
- プログラムのテスト
- プログラムの提出

## インストール

```shell
brew tap travelist/homebrew-aoj-cli
brew install aoj
```

## 利用方法

AOJ CLI の設定 (初回のみ実行)

```shell
aoj init
```

テストサンプルのダウンロードとディレクトリの作成

```shell
aoj gen [PROBLEM-ID]
```

ソースコードのテスト

```shell
# cd ./[PROBLEM-ID]
aoj test
```

ソースコードの提出

```shell
aoj submit
```

## Configuration

デフォルトの設定ファイル: `~/.aoj-cli/config.toml`
デフォルトのテンプレートファイル: `~/.aoj-cli/template.txt`

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
| gen.template_file | テンプレートファイルへのパス | |
| gen.destination_file_name | 生成するコードのファイル名 | |
| test.before_all | テスト全体の前に実行するコマンド | |
| test.before_each | 各テスト前に実行するコマンド | |
| test.test_commands | テストで実行するコマンド | |
| test.after_each | 各テスト後に実行するコマンド | |
| test.after_all | すべてのテスト後に実行するコマンド | |
| submit.language | 提出するプログラムの言語 | ex: `JAVA`, `C++` |
| submit.source_file_name | 提出時するファイル名 | |

## TODO

- [ ] テスト時の `TLE` チェック



