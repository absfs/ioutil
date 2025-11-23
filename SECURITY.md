# Security Policy

## Supported Versions

We actively support the following versions of this project:

| Version | Supported          |
| ------- | ------------------ |
| latest  | :white_check_mark: |

Since this is a library package, we recommend always using the latest version from the `master` branch.

## Reporting a Vulnerability

We take the security of our software seriously. If you believe you have found a security vulnerability in this project, please report it to us responsibly.

### How to Report

**Please do NOT report security vulnerabilities through public GitHub issues.**

Instead, please report security vulnerabilities by:

1. Opening a [security advisory](https://github.com/absfs/ioutil/security/advisories/new) on GitHub
2. Or by opening a private issue and marking it as a security concern

Please include the following information in your report:

- Type of issue (e.g., buffer overflow, path traversal, etc.)
- Full paths of source file(s) related to the manifestation of the issue
- The location of the affected source code (tag/branch/commit or direct URL)
- Any special configuration required to reproduce the issue
- Step-by-step instructions to reproduce the issue
- Proof-of-concept or exploit code (if possible)
- Impact of the issue, including how an attacker might exploit it

### What to Expect

- **Acknowledgment**: We will acknowledge receipt of your vulnerability report within 48 hours
- **Communication**: We will keep you informed about the progress of addressing the vulnerability
- **Timeline**: We aim to address critical vulnerabilities within 7 days
- **Credit**: We will credit you in the security advisory (unless you prefer to remain anonymous)

## Security Considerations

When using this library, keep in mind:

1. **Path Traversal**: Always validate and sanitize user-provided paths before passing them to filesystem operations
2. **Permissions**: Be cautious with file permissions when using `WriteFile`, `TempFile`, and `TempDir`
3. **Temporary Files**: Always clean up temporary files and directories to avoid information disclosure
4. **Abstract Filesystems**: When implementing custom filesystem types, ensure proper access controls and validation

## Best Practices

- Use appropriate file permissions (avoid 0777)
- Always close files after use
- Clean up temporary files and directories
- Validate user input before using it in filesystem operations
- Be aware that `TempFile` and `TempDir` create predictable names - ensure your temp directory has proper permissions
- When using with network-based or custom filesystems, implement additional security checks as needed

## Dependencies

This project minimizes dependencies to reduce the attack surface. We only depend on:

- `github.com/absfs/absfs` - Core abstract filesystem interface
- `github.com/absfs/memfs` - In-memory filesystem (testing only)
- `github.com/absfs/osfs` - OS filesystem wrapper (testing only)

We regularly update dependencies and monitor for security vulnerabilities.

## Disclosure Policy

When we receive a security bug report, we will:

1. Confirm the problem and determine affected versions
2. Audit code to find any similar problems
3. Prepare fixes for all supported versions
4. Release new versions as quickly as possible

## Comments on This Policy

If you have suggestions on how this process could be improved, please submit a pull request or open an issue.
