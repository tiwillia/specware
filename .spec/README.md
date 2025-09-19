# Spec-driven Development

This directory contains feature specifications and related artifacts for spec-driven development workflow.

## Purpose

Spec-driven development guides AI code generation through structured specifications, leading to more successful implementations. These artifacts help with feature iteration, understanding implementation decisions, and replicating similar work.

### No Maintenance
Specs are point-in-time artifacts - they are not maintained or used in automation outside of building features. Specs are not expected to be updated to reflect changes in a feature post-implementation.

## Getting Started

1. Use the Claude Code command "/specify" to begin feature specification
2. The workflow will guide you through requirements gathering and implementation planning
3. Each feature gets its own numbered directory (e.g., "001-feature-name/")

## Directory Structure

- "NNN-feature-name/" - Individual feature directories
- "templates/" - Project-specific templates (optional)

## Usage

### Start a new feature specification:
```
$ claude
> /specify Add a new feature for user authentication
```

### Continue working on an existing feature:
```
$ claude  
> /specify continue to where we left off for the user authentication feature
```

For more information, see the main project README.md.
