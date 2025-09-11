# Specware

Specware creates structured directories, commands, and templates and provides a tool to Claude Code to facilitate feature requirements gathering and implementation planning.

## Quick Start

### 1. Install
```bash
# go install github.com/tiwillia/specware@latest
```

Or to install from a local clone/fork:
```bash
# go install .
```

### 2. Initialize Your Project
Navigate to your project directory, then initialize: 
```bash
# Initialize spec workflow in current directory
./specware init .

# Optional: Create localized templates for customization
./specware localize-templates
```

### 3. Start Feature Specification with Claude Code
```bash
# Open Claude Code and begin the spec-driven workflow
claude

# Use the /specify command to start gathering requirements
> /specify Add user authentication with email and password
```

The `/specify` command will guide you through:
- **Requirements gathering** - Interview-style Q&A to understand the feature
- **Expert research** - Codebase analysis and technical planning  
- **Implementation planning** - Detailed technical specifications with interview-style Q&A
- **(optional) Interactive Review** - Interactive review and amendment of each section of generated specifications
- **Status tracking** - Progress monitoring throughout development

Claude Code automatically uses the `specware` tool to create directories, track progress, and maintain organized specifications for your features.

## How it works

`specware` tool is used to set up a project repository for spec-driven development with un-intrusive directories and files ignored by git by default, then is used during interactive specification generation by Claude Code to facilitate filesystem and state tracking operations.

`/specify` claude command created in the project repository facilitates the interactive specification generation. The command is localized in the project to allow for project-specific customization.

`assets/templates/` are used as the base specifications Claude Code will fill out during specification generation. Templates can optionally be localized for project-specific customization.

`.spec/` directory created in the project repository provides a basic structure to organize generated specifications.

### Specware Tool

The `specware` tool is used for both project setup and during interactive specification generation.

After initialization, Specware creates in your project:

```
.claude/
  commands/
    specify.md                 # Claude Code workflow command
.spec/
  README.md                    # Spec workflow documentation  
```

#### Project Setup
These commands are intended to be run by a user:
- `init <directory>` - Initialize project with spec-driven workflow support
- `localize-templates` - Copy embedded templates to `.spec/templates/` for customization, not required.

#### Feature Management
These commands are intended to be run by Claude Code to facilitate feature specification:
- `feature new-requirements <short-name>` - Create new feature specification directory with requirements template
- `feature new-implementation-plan <short-name>` - Add implementation plan to existing feature
- `feature update-state <short-name> <status>` - Update feature development status

### Claude Command (/specify)

Interactive Claude Command with two primary workflows:

1. **Requirements Building:**
- Requirements gathering via Q&A
- Context gathering through codebase analysis  
- Expert Q&A with technical insights
- Requirements finalization
- Optional interactive review

2. **Implementation Planning:**
- Codebase analysis for technical approach
- Implementation Q&A for technical details
- Plan generation with detailed tasks
- Optional interactive review

Features state tracking and a one-question-at-a-time interview style with smart defaults.

The `/specify` command is a glorified prompt, you are expected to take advantage of the flexibility granted by Claude Code to modify the workflow, continue where you left off, skip steps, etc as needed.

### Templates (assets/templates/)

Embedded templates used by Claude Code during specification generation:

- **`requirements.md`** - Structure for feature requirements with sections for problem statement, solution overview, functional/technical requirements, acceptance criteria, and constraints
- **`implementation-plan.md`** - Structure for technical plans with milestones, phases, tasks, code examples, and deployment considerations
- **`context.md`** - Template for context gathering sessions used to create both `context-requirements.md` and `context-implementation-plan.md` files

Templates can be localized to `.spec/templates/` for project-specific customization using `specware localize-templates`.

### Specification Artifacts (.spec/)

Generated specification files organized in numbered feature directories:

**Generated Files:**
- **`requirements.md`** - Final requirements specification filled from template
- **`implementation-plan.md`** - Final implementation plan with detailed tasks and code examples  
- **`context-requirements.md`** - Q&A context and codebase research for requirements phase
- **`context-implementation-plan.md`** - Q&A context and technical analysis for implementation phase
- **`.spec-status.json`** - Current workflow status and progress tracking

**Directory Structure:**
```
.spec/
  001-user-auth/           # Sequential numbering
  002-dashboard/
  003-notifications/
```

All files are ignored by git by default to keep specifications separate from source code.

#### Status Tracking

Features are tracked through `.spec-status.json` files within a feature spec directory (`.spec/000-example-feature/`) with suggested workflow phases:
- `"Requirements Gathering"`
- `"Requirements Context Gathering"`
- `"Requirements Expert Q&A"`
- `"Requirements Complete"`
- `"Requirements Interactive Review"`
- `"Implementation Planning"`
- `"Implementation Plan Q&A"`
- `"Implementation Plan Generated"`
- `"Implementation Plan Interactive Review"`
- `"Implementation Planning Complete"`

There is no validation or expectation that state values match this list.

## Guiding Principles

**Reduce reliance on the LLM**
- Extract operations that need to be consistent and repeatable to traditional code
- Extract output expectations to templates
- Extract state tracking to structured metadata
- Reduce token consumption and time spent waiting on inference

**Be un-intrusive**
- Do not create any artifacts that will not be ignored by git in default configurations
- Create the minimal amount of spec artifacts, commands, templates, etc.
- Do not force specific workflows

**Allow for customization**
- Teams and individuals will have specific preferences, enable project-specific customization where possible
- Provide sensible defaults

**Single assistant focus**
- Enable a single coding assistant (Claude Code) to avoid complexity
