# ADK Go Extensions

A set of shared extensions and middleware for the ADK (Agent Development Kit) in Go.

## Overview

This library extends the core ADK with specialized tools for modern agent workflows:
- **Zep Session Provider**: Persistent session management powered by Zep, including RAG support.
- **Enhanced Middleware**: Built-in options for policy enforcement, privacy-focused persistence, and dynamic response formatting.

## Installation

```bash
go get go.naturallyfunny.dev/adk
```

## Usage

This project follows Go's idiomatic approach to documentation. For detailed API info, you can browse the source comments or use `godoc`.

### Examples
Check out the `examples/` directory for runnable code samples focusing on specific features:

- [Zep Provider](examples/zep/main.go): Setting up Zep as your session backend.
- [Session Options](examples/session/main.go): Using advanced middleware like Policies, Privacy controls, and Dynamic Formatting.

To run an example:
```bash
# Make sure ZEP_API_KEY is set in your environment
go run examples/zep/main.go
```

### API Documentation
```bash
go doc -all ./zep
go doc -all ./session
```

## Project Structure

- `zep/`: Zep-backed session service implementation.
- `session/`: Middleware wrappers and session decorators.
- `examples/`: Focused, runnable usage guides.
