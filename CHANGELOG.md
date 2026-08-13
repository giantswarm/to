# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

### Added

- Add `go.mod`, making this a proper Go module.
- Add unit tests covering all exported functions.
- Add `renovate.json5` to pick up the shared Giant Swarm Renovate presets.
- Add this changelog.

### Changed

- Replace the hand-rolled CircleCI config with the `giantswarm/architect` orb.
  The old config downloaded `architect` v1.0.0 from the GitHub API using
  `RELEASE_TOKEN`, which is no longer set on the project, so every build failed.

## [0.4.2] - 2024-11-26

### Changed

- Revert the deprecation notice added in 0.4.1.

## [0.4.1] - 2024-11-22

### Deprecated

- Mark the package as deprecated in favour of idiomatic Go.

## [0.4.0] - 2021-11-18

### Added

- Add support for floats and 32-bit ints.

## [0.3.0] - 2020-05-04

### Added

- Add `Duration` and `DurationP`.

## [0.2.0] - 2019-10-22

### Added

- Initial set of pointer/value conversion helpers.

[Unreleased]: https://github.com/giantswarm/to/compare/v0.4.2...HEAD
[0.4.2]: https://github.com/giantswarm/to/compare/v0.4.1...v0.4.2
[0.4.1]: https://github.com/giantswarm/to/compare/v0.4.0...v0.4.1
[0.4.0]: https://github.com/giantswarm/to/compare/v0.3.0...v0.4.0
[0.3.0]: https://github.com/giantswarm/to/compare/v0.2.0...v0.3.0
[0.2.0]: https://github.com/giantswarm/to/releases/tag/v0.2.0
