# LeetCode Go Solutions

A comprehensive collection of LeetCode problem solutions implemented in Go, organized by topic.

---

## Table of Contents

* [Description](#description)
* [Repository Structure](#repository-structure)
* [Getting Started](#getting-started)
* [Running Tests](#running-tests)
* [Utilities](#utilities)
* [Contributing](#contributing)
* [License](#license)

---

## Description

This repository houses Go implementations of LeetCode problems, grouped into high-level categories such as Arrays, Strings, Linked Lists, Trees & Graphs, and Dynamic Programming. Each problem resides in its own folder with a dedicated `README.md`, solution file (`*.go`), and test file (`*_test.go`). The goal is to provide a clear, consistent structure that helps you navigate, study, and contribute solutions efficiently.

## Repository Structure

```
leetcode-go/
├── README.md                # This file
├── go.mod                   # Go module definition
├── Makefile                 # Utility commands: fmt, vet, test, coverage
│
├── arrays/                  # Arrays & Slices solutions
│   ├── README.md            # Overview of this section
│   ├── 001_two_sum/         # Problem #1: Two Sum
│   │   ├── README.md
│   │   ├── two_sum.go
│   │   └── two_sum_test.go
│   └── ... more problems ...
│
├── strings/                 # Strings solutions
├── linked_lists/            # Linked Lists solutions
├── trees/                   # Trees & Graphs solutions
└── dynamic_programming/     # Dynamic Programming solutions
```

Each section folder starts with a `README.md` that links to the LeetCode topic page and outlines common patterns and tips for that category.

## Getting Started

1. **Clone the repository**

   ```bash
   git clone https://github.com/Perico6255/LeetCodeGo.git
   cd leetcode-go
   ```

## Running Tests

This project uses a `Makefile` to simplify common commands:

| Command             | Description                                      |
| ------------------- | ------------------------------------------------ |
| `make fmt`          | Format all Go files with `go fmt`                |
| `make vet`          | Run `go vet` for static analysis                 |
| `make lint`         | Run `golangci-lint` (if installed)               |
| `make test`         | Run all tests                                    |
| `make test-verbose` | Run tests with verbose output                    |
| `make test-single`  | Run a specific test using `TESTARGS`             |
| `make coverage`     | Generate and display test coverage report        |
| `make clean`        | Remove generated artifacts (e.g., coverage file) |

Examples:

```bash
make test
make test-single TESTARGS='-run ^TestTwoSum$'
make coverage
```

## Utilities

Common data structures and helper functions (e.g., linked list or tree constructors) live under the `utils/` package. Import them in your solutions when needed:

```go
import "github.com/your_username/leetcode-go/utils"
```

## Contributing

Contributions are welcome! Please read [CONTRIBUTING.md](CONTRIBUTING.md) for guidelines on:

* Problem folder naming conventions
* Code style and testing standards
* Pull request process



