# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.0.0] - 2025-11-15

### Added
- Initial release of echo-server
- Complete echo server implementation using Fiber v3
- Support for GET, POST, PUT, PATCH, DELETE methods
- Customizable responses via headers or query parameters
- Custom HTTP status codes
- Custom response bodies
- Custom response headers
- Custom response latency
- Environment variable access
- File/folder exploration
- Health check endpoint
- Docker support with multi-stage builds
- GitHub Actions workflow for building and pushing to GHCR
- Comprehensive test coverage

### Changed
- Migrated Docker image publishing from AWS ECR to GitHub Container Registry (GHCR)
- Updated Go version from 1.24.0 to 1.24.7
- Upgraded testify from v1.10.0 to v1.11.1
- Upgraded golang.org/x/crypto from v0.31.0 to v0.44.0
- Upgraded golang.org/x/net from v0.31.0 to v0.47.0
- Upgraded golang.org/x/sys from v0.28.0 to v0.38.0
- Upgraded golang.org/x/text from v0.21.0 to v0.31.0

### Fixed
- Fixed Dockerfile COPY path from /app/echo-server to /build/echo-server to match the build WORKDIR

[1.0.0]: https://github.com/DonsWayo/echo-server/releases/tag/v1.0.0
