# gsa - Go Static Analysis

`gsa` is a simple static analysis tool for Go source code. It helps identify common issues and enforce coding standards in your projects.

## Features

`gsa` currently checks for the following issues:

*   **TODOs:** Detects `//TODO` comments, reminding you of pending tasks.
*   **`time.Sleep`:** Flags the use of `time.Sleep`, which can be problematic in production code.
*   **Concurrency:** Identifies basic concurrency issues.
*   **Unchecked Errors:** Finds error variables that are not checked.
*   **Cyclomatic Complexity:** Measures the complexity of functions and reports if they exceed a certain threshold.

## Installation

To install `gsa`, you need to have Go installed on your system. You can then use `go get` to install the tool:

```bash
go get github.com/rajiknows/gsa
```

## Usage

To analyze your Go project, run the `gsa analyze` command from your project's root directory:

```bash
gsa analyze
```

You can also specify a path to a specific directory or file:

```bash
gsa analyze ./...
```

`gsa` will print any found issues to the console.

## Example Output

```
[todo] /home/user/project/main.go:10 TODO: Refactor this function
[sleep] /home/user/project/main.go:15 Use of time.Sleep detected
[unchecked-error] /home/user/project/main.go:20 Unchecked error
```

## Contributing

Contributions are welcome! If you have any ideas for new rules or improvements, please open an issue or submit a pull request.


