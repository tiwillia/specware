# Specware initial requirements

## Summary

Spec-driven workflow tooling and documentation. Provides tooling limited to Claude Code covering requirements gathering and implementation planning. Task management is not in-scope for the initial creation, but should be considered a future goal.

## Purpose

Facilitate a spec-driven workflow through the Claude Code AI Coding Assistant.

### Goals
- Absolute minimal first-pass implementation
- Allow for flexiblity and individualized workflow modifications
- Provide a dialogue enabling assisted iteration and review of spec artifacts
- Use reliable traditional code and tooling for any functionality that should be repeatable and does not strictly require LLM interaction
- Allow for template and command modifications to be persisted in the project
- Artifact creation should be minimal, especially at project root

### Non-Goals
- Automate or integrate git functionality 
- Publish commands or artifacts remotely for the tool to use
- Add task management or an implementation framework as part of initial requirements
- Handling of corrupted/broken spec files or directories.

## Definitions
**Spec**: Or "specification" refers to documents and related media that specify the makeup of a given feature or initiative.
**Tool**: A script or CLI command
**Commands**: Claude Code Commands, markdown files providing repeatable workflows in a prompt - specifically referring to Claude commands created by the tool.

## Requirements
- Commands guide LLM on tool usage to initiate new feature or initiative
- Tool creates a .spec/ directory and guidance within to initiate organize spec-drive-development artifact storage
- Tool creates a .spec/000-feature-name directory for each feature or initiative
- Tool stores all relevant spec artifacts in the feature directory
- Tool and commands allow for stopping anywhere in the process and picking up where left off.
- Tool follows naming convention of spec sub-directories including sequential numbering. Conflicts due to parallel work efforts are expected to be handled manually.
- Commands allow for and support updating requirements during implementation planning phase when necessary.
- Tool creates claude commands in initialization phase but does not maintain or update them, allowing for customizations.
- Tool creates spec files using built-in templates by default, or localized templates if they exist. Localized templates are preferred if they exist.
- Tool and Commands do not expect and enforce any specific content or sections in templates, allowing for full customization.
- Commands use Tool to create seperate spec artifacts containing Q&A results.
- Root-level README provides a very simple getting started flow and a complex "full walkthrough".

### Specific Technical Requirements
- The tool must be written in golang
- Claude Code is the only supported AI coding assistant, no others will be added.

## Example Usage

Throughout the following phases, the Claude Code command workflows instruct claude to use the tool to track the state. Examples:
```
$ specware feature update-state example-spec "Requirements Q&A"
```

The tool uses the `.spec/000-example-spec/.spec-status.json` file to track status.

### Phase 0: Initialization
User initiates the spec workflow in their repo using the tool. User is expected to specify a directory to inititialize.
```
$ specware init .
```

The tool creates the necessary resources in the repo:
```
.claude/commands/
  specify.md
.spec/
  README.md
  000-example-spec/
    .spec-status.json
```

Note that git ignores the .specware/.* and .spec/ files by default.

User can optionally initiate project-specific templates
```
$ specware localize-templates
```

Tool creates localized templates
```
.spec/templates/
  requirements.md
  implementation-plan.md
```

### Phase 1: Requirements Building
User begins feature specification in claude code:
```
$ claude
> /specify Add a new page "community" where users can create and view bulletin-board style posts from other users
```

Claude command instructs claude to create the new feature spec dir and a copy of the requirements template via the tool. The tool updates the current_feature tracker:
```
$ specware feature new-requirements community-page
.spec/
  001-community-page/
    .spec-status.json
    q&a-requirements.md
    requirements.md
```

Claude fills in the basic sections and metadata of the requirements spec. Claude fills in the basic sections of the q&a-requirements spec.

Claude generates 3-5 follow-up questions to clarify requirements within context of the repo and updates the q&a-requirements spec.

Claude asks these questions, one-by-one, and records both the questions and the answers in q&a-requirements.md.

Claude does research to become an "expert" on the topics and asks 3-5 more follow-up questions, updating q&a-requirements.md at each step.

Claude writes the final requirements to requirements.md.

The command instructs Claude to offer three options for next steps:
1. An interactive review session of the requirements documentation
2. Stop here, user to review the document asynchronously and provide details on any changes or amendments required.
3. Move onto the implementation planning phase

If they stop and need to open a new session at a later time to provide requirements feedback, they can continue with:
```
$ claude
> /specify I have feedback on the requirements doc we need to address. <feedback>
```

### Phase 2: Implementation Planning
Once the user is satisfied with the requirements documentation, they are expected to initiate the Implementation planning phase via natural lanugage with Claude. 

If they have exited to asynchronous review of the requirements doc, they can continue using the single command:
```
$ claude
> /specify continue to the implementation planning phase
```

Claude is instructed to use the tool to create the plan spec
```
$ specware feature new-implementation-plan community-page
.spec/
  001-community-page/
    .spec-status.json
    implementation-plan.md
    requirements.md
    q&a-implementation-plan.md
    q&a-requirements.md
```

Claude is instructed through the command to understand the existing codebase - if any - and define an implementation plan for the requirements including technical details.

Claude again asks 3-5 questions twice, again updating a q&a spec specific to the implementation planning phase.

Claude writes the plan and again offers interactive review.

The user is expected to use the requirements and plan to implement the feature using their coding assistant. There is no helper for the implementation phase.

## Future considerations (to be ignored by YOU, CLAUDE)
- Include a set of "janitorial" agents that are used as part of the workflow
  - Example: scope creep craig
  - Determine which parts of the workflow may be essential to include these agents
- Directory containing spec files should be configurable
