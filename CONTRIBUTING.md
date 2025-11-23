# Contributing to ioutil

First off, thank you for considering contributing to ioutil! It's people like you that make the absfs ecosystem better.

## Code of Conduct

This project and everyone participating in it is governed by our [Code of Conduct](CODE_OF_CONDUCT.md). By participating, you are expected to uphold this code.

## How Can I Contribute?

### Reporting Bugs

Before creating bug reports, please check the existing issues to avoid duplicates. When you create a bug report, include as many details as possible:

- **Use a clear and descriptive title**
- **Describe the exact steps to reproduce the problem**
- **Provide specific examples** to demonstrate the steps
- **Describe the behavior you observed** and what you expected to see
- **Include Go version and OS information**

### Suggesting Enhancements

Enhancement suggestions are tracked as GitHub issues. When creating an enhancement suggestion:

- **Use a clear and descriptive title**
- **Provide a detailed description** of the suggested enhancement
- **Explain why this enhancement would be useful** to most users
- **List any similar features** in other libraries if applicable

### Pull Requests

1. Fork the repo and create your branch from `master`
2. If you've added code, add tests that cover your changes
3. Ensure the test suite passes (`go test -v -race ./...`)
4. Run `go vet ./...` and `gofmt -w .` to ensure code quality
5. Update documentation as needed
6. Write a clear commit message describing your changes

#### Pull Request Guidelines

- **Keep PRs focused** - One feature or bug fix per PR
- **Write tests** - All new code should have test coverage
- **Follow Go conventions** - Use `gofmt`, follow standard Go style
- **Document public APIs** - All exported functions need documentation
- **Maintain backward compatibility** - Don't break existing APIs without discussion

## Development Setup

1. **Install Go 1.22 or later**
   ```bash
   go version
   ```

2. **Clone the repository**
   ```bash
   git clone https://github.com/absfs/ioutil.git
   cd ioutil
   ```

3. **Download dependencies**
   ```bash
   go mod download
   ```

4. **Run tests**
   ```bash
   go test -v -race ./...
   ```

5. **Check code quality**
   ```bash
   go vet ./...
   gofmt -l .
   ```

## Testing

We maintain high test coverage. When adding new features:

- Write unit tests for all new functions
- Include table-driven tests where appropriate
- Test error conditions and edge cases
- Use the `setup()` helper in tests to create test filesystems

Run tests with:
```bash
# Run all tests
go test -v ./...

# Run tests with race detection
go test -v -race ./...

# Run tests with coverage
go test -v -coverprofile=coverage.out ./...
go tool cover -html=coverage.out
```

## Code Style

- Follow standard Go formatting (use `gofmt`)
- Keep functions focused and small
- Document all exported functions with GoDoc comments
- Use meaningful variable and function names
- Prefer clarity over cleverness

## Commit Messages

- Use the present tense ("Add feature" not "Added feature")
- Use the imperative mood ("Move cursor to..." not "Moves cursor to...")
- Limit the first line to 72 characters or less
- Reference issues and pull requests after the first line

Examples:
```
Fix TempFile creation on Windows

Add support for custom temp directory permissions

Update dependencies to latest versions
Fixes #123
```

## License

By contributing, you agree that your contributions will be licensed under the same BSD-style license that covers this project (the Go license). Since this project is derived from the Go standard library, all contributions must be compatible with that license.

## Questions?

Feel free to open an issue with your question or reach out to the maintainers.

Thank you for contributing!
