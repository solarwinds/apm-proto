# Contributing to apm-proto

Thank you for your interest in contributing to `apm-proto`! This repository defines the Protocol Buffer (protobuf) contract between SolarWinds APM libraries and the collector. The following guidelines will help you contribute effectively.

## Table of Contents

- [Prerequisites](#prerequisites)
- [Repository Structure](#repository-structure)
- [Making Changes](#making-changes)
  - [Modifying the Proto File](#modifying-the-proto-file)
  - [Regenerating Code](#regenerating-code)
  - [Validating Changes](#validating-changes)
- [Pull Request Process](#pull-request-process)
- [Versioning & Compatibility](#versioning--compatibility)
- [Code Owners](#code-owners)

---

## Prerequisites

All code generation is done via Docker, so the hard requirements are:

| Requirement                                   | Purpose                                                                                           |
|-----------------------------------------------|---------------------------------------------------------------------------------------------------|
| amd64 architecture environment                | Running the required `amd64` only Docker images `namely/protoc-all` & `pseudomuto/protoc-gen-doc` |
| [Docker](https://docs.docker.com/get-docker/) | Running `protoc` and code generators                                                              |
| [Git](https://git-scm.com/)                   | Version control                                                                                   |
| `make`                                        | Running build targets                                                                             |

> **Note:** You do **not** need to install `protoc`, Go, or any other language toolchain locally — the Makefile pulls the required Docker images automatically.

---

## Repository Structure

```
apm-proto/
├── collector.proto        # The canonical proto definition (edit this)
├── Makefile               # Build, generate, and validation targets
├── README.md              # Auto-generated protocol documentation (do not edit manually)
├── cpp/                   # Generated C++ sources (do not edit manually)
└── go/
    └── collectorpb/       # Generated Go sources and mocks (do not edit manually)
```

---

## Making Changes

### Modifying the Proto File

All meaningful changes start in `collector.proto`. This file defines:

- **Messages** – data structures exchanged between APM libraries and the collector.
- **Enums** – typed constants used in messages.
- **Services** – the `TraceCollector` RPC service definition.

When editing `collector.proto`:

- Follow the existing commenting style (use `/** ... */` block comments for messages/enums).
- **Never reuse or reassign field numbers.** If a field is removed, reserve the number with a `reserved` statement.
- Prefer **backwards-compatible** changes (adding new optional fields) over breaking ones (removing or renumbering fields).
- Update or add comments that explain the purpose, constraints, and any cross-references (e.g., links to external docs such as AWS IMDS).

### Regenerating Code

After modifying `collector.proto`, regenerate all language bindings and documentation:

```bash
make all
```

Individual targets are also available:

| Target | Description |
|--------|-------------|
| `make cpp` | Regenerate C++ sources under `cpp/` |
| `make go`  | Regenerate Go sources and mocks under `go/collectorpb/` |
| `make doc` | Regenerate `README.md` from the proto comments |
| `make clean` | Remove all generated files |

> The generated files (`cpp/`, `go/`, `README.md`) are **checked in** to this repository so that GitHub action checks will pass. Always commit the regenerated files together with your proto change.

### Validating Changes

Run the following to confirm that the generated output is consistent with the proto source and that nothing was accidentally left out:

```bash
make check
```

This runs `make all` and then checks that there are no uncommitted diffs. GitHub action will run the same check on every pull request.

---

## Pull Request Process

1. Make sure `make check` passes locally with no diff.
2. Write a clear PR description that explains:
   - **What** changed in `collector.proto`.
   - **Why** the change is needed.
   - Any **backwards-compatibility** considerations.
3. Update inline proto comments if the change affects field semantics.
4. Request a review from the code owners (see [Code Owners](#code-owners)). At least **one approval** is required before merging.
5. Squash or rebase your commits to keep the history clean before the PR is merged.

---

## Versioning & Compatibility

This proto file is consumed by multiple APM language agents (Java, PHP, .Net) and [solarwinds-cloud/collector2cloudwatchlogs](https://github.com/solarwinds-cloud/collector2cloudwatchlogs). Backwards-incompatible changes (removing fields, changing field types, renumbering fields) are **breaking changes** and must be coordinated across all consumers before merging.

When in doubt, prefer:
- Adding a new **optional** field rather than modifying an existing one.
- Marking old fields as `reserved` rather than deleting them outright.
- Opening an issue first to discuss the impact before submitting a breaking change PR.

---

## Code Owners

This repository is maintained by:

- [@solarwinds/eng-pub-apm-instrumentation](https://github.com/orgs/solarwinds/teams/eng-pub-apm-instrumentation)
- [@solarwinds/con-pub-apm-instrumentation](https://github.com/orgs/solarwinds/teams/con-pub-apm-instrumentation)

For questions or design discussions, open a GitHub issue or reach out to the team directly.

