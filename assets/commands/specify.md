# Specify - Spec-driven Development Workflow

## Summary
Guide the user through a comprehensive spec-driven development workflow for feature requirements gathering and implementation planning using the specware tool.

## Important Rules
- Always use the specware tool to track state and create artifacts
- Maintain one feature at a time in active development
- Support stopping and resuming at any point in the workflow
- Record all Q&A interactions in the appropriate q&a files
- Follow the existing codebase patterns and conventions
- Use actual file paths and component names in artifacts

### Q&A Rules
- ONLY yes/no questions with smart defaults
- ONE question at a time
- Write ALL questions to file BEFORE asking any
- Stay focused on requirements during requirements gathering, exclude implementation details
- Document WHY each default makes sense
- The answer "I don't know" is okay, use the default in this case.

### Specware usage Guidance
The specware tool is necessary to facilitate the basic operations of the workflow.

If the specware tool is not avaialble, immediately stop and instruct the user to install the tool.

**Feature Management**
  specware feature new-requirements <short-name>         # Add requirements to feature (creates dir if not exist)
  specware feature new-implementation-plan <short-name>  # Add implementation plan to feature (creates dir if not exist)
  specware feature update-state <short-name> <status>    # Update feature development status

**Directory Structure Created**

  .claude/commands/                          # Claude Code commands (includes specify.md workflow)
  .spec/                                     # Feature specifications (001-feature-name/, 002-another-feature/)
  .spec/001-feature-name/.spec-status.json    # Feature status tracking

**Typical Workflow**

  1. specware feature new-requirements user-auth - Start new feature
  2. specware feature new-implementation-plan user-auth - Add implementation planning
  3. specware feature update-state user-auth "implementation-complete" - Track progress

## Workflow

### Pre: Assess Input

$ARGUMENTS should be either:
1. Feature requirements, indiciating a new feature specification should be started
2. Feedback on requirements or implementation plans, indicating that finalization step is likely in progress and review is occurring asynchronously.
3. A short name, indicating the user wants to pick up where they left off in this workflow for the feature with the given short name.

If unsure, ask for clarification or request a feature name.

### Phase 1: Requirements Building

#### Step 1: Feature Specification File Setup
- Generate a descriptive short-name based on the feature description
- Use `specware feature new-requirements <short-name>` to create the feature directory and base `requirements.md` file based on template.

#### Step 2: Requirements Gathering
- Use `specware feature update-state <short-name> "Requirements Gathering"`
- Fill in the basic sections and metadata of the requirements spec
- Create initial content in both `requirements.md` and `context-requirements.md`
- Generate the five most important yes/no questions to understand the problem space:
  - Questions informed by codebase structure
  - Questions about user interactions and workflows
  - Questions about similar features users currently use
  - Questions about data/content being worked with
  - Questions about external integrations or third-party services
  - Questions about performance or scale expectations
  - Write all questions to `context-requirements.md` with smart defaults
  - Ask questions one at a time proposing the question with a smart default option
- Only after all questions are asked, record answers in `context-requirements.md` as received, not before.

#### Step 3: Context Gathering
- Use `specware feature update-state <short-name> "Requirements Context Gathering"`
- Research the codebase to become an "expert" on the relevant topics
- Deep dive into existing similar practices, patterns, and features
- Use web searches for best practices or library documentation
- Document findings in `context-requirements.md`

#### Step 4: Expert Requirements Questions
- Use `specware feature update-state <short-name> "Requirements Expert Q&A"`
- Now you are an expert on the codebase, a senior developer with the right knowledge.
- Write the top 3-5 most important questions yes/no questions to `context-requirements.md`
  - These questions should clarify expected behavior with a deep understanding of the code.
  - Ask questions one at a time proposing the question with a smart default option

#### Step 5: Finalize Requirements
- Generate comprehensive requirements based on the template in `requirements.md`. Do not delete or modify existing sections, honor the template.
- Fill in `requirements.md` with the final requirements.
- Use `specware feature update-state <short-name> "Requirements Complete"`
- Offer three options:
  1. Interactive review session of the requirements documentation
  2. Stop here for asynchronous review and feedback, the user being expected to review the requirements document.
  3. Move to implementation planning phase, skipping review (not reccomended).

#### Step 6: (optional) Interactive review session
- Use `specware feature update-state <short-name> "Requirements Interactive Review"`
- For each section of the `requirements.md` document, perform the following interactive review steps:
  1. Generate a 1-3 sentence summary of the section
  2. Display the section without modifications to the user and then the generated summary.
  3. Ask the user for any changes or amendments.
  4. If the user provides changes or amendments, make the changes and make any additional changes needed to other sections of the document.
  5. Once the user is satisfied, proceed to the next section.
- Once all sections are approved, use `specware feature update-state <short-name> "Requirements Complete"`

### Phase 2: Implementation Planning
When the user is ready for implementation planning:

#### Step 1: Implementation Plan File Setup
- Use `specware feature new-implementation-plan <short-name>` to create the plan spec
- Update status with `specware feature update-state <short-name> "Implementation Planning"`
- Read the `requirements.md` for this feature

#### Step 2: Codebase Analysis
- Understand the existing codebase structure and patterns
- Dive deep into the parts of the codebase that need to be modified to meet the requirements
- Review the existing code for patterns, best practices, and similar features to ensure existing patterns are followed
- Check for and understand CONTRIBUTING.md or similar contribution documentation
- Review any existing code style guidelines
- Record your findings in `context-implementation-plan.md`

#### Step 3: Implementation Plan Q&A
- Use `specware feature update-state <short-name> "Implementation Plan Q&A"`
- Generate the five most important yes/no questions to understand technical implementation details:
  - Questions about best practices or patterns to follow
  - Questions about packaging and file structure
  - Questions about data model details and conventions
  - Questions about error handling and logging requirements
  - Questions about performance and security concerns
  - Write all questions to `context-implementation-plan.md` with smart defaults and example code snippets where necessary
  - Ask questions one at a time proposing the question with a smart default option and examples
- Only after all questions are asked, record answers in `context-implementation-plan.md` as received, not before.

#### Step 5: Finalize Implementation Plan
- Generate a comprehensive implementation plan, breaking out large operations and changes into smaller tasks
- Write the complete implementation plan to `implementation-plan.md`
- Update status with `specware feature update-state <short-name> "Implementation Plan Generated"`

#### Step 6: Interactive Review
- Use `specware feature update-state <short-name> "Implementation Plan Interactive Review"`
- For each section or phase of the `implementation-plan.md` document, perform the following interactive review steps:
  1. Generate a 1-3 sentence summary of the section.
  2. Display the section without modifications to the user and then the generated summary.
  3. Ask the user for any changes or amendments.
  4. If the user provides changes or amendments, make the changes and make any additional changes needed to other sections of the document.
  5. Once the user is satisfied, proceed to the next section.
- Once all sections are approved, use `specware feature update-state <short-name> "Implementation Planning Complete"`

## Question format when displayed to user:

### Discovery Questions:
```
## Q1: Will users interact with this feature through a visual interface?
**Default if unknown:** Yes (most features have some UI component)

## Q2: Does this feature need to work on mobile devices?
**Default if unknown:** Yes (mobile-first is standard practice)
```

### Expert Questions:
```
## Q7: Should we extend the existing UserService at services/UserService.ts?
**Default if unknown:** Yes (maintains architectural consistency)

## Q8: Will this require new database migrations in db/migrations/?
**Default if unknown:** No (based on similar features not requiring schema changes)
```
