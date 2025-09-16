# Specify - Spec-driven Development Workflow

## Summary
Guide the user through a comprehensive spec-driven development workflow for feature requirements gathering and implementation planning based on $ARGUMENTS using the specware tool.

## Workflow

### Pre: Assess Input

$ARGUMENTS should be either:
1. Feature requirements, indiciating a new feature specification should be started
2. Feedback on requirements or implementation plans, indicating that finalization step is likely in progress and review is occurring asynchronously.
3. A short name, indicating the user wants to pick up where they left off in this workflow for the feature with the given short name.

If unsure, ask for clarification or request a feature name.

### Phase 1: Requirements Building
Guide the user through generating a requirements specification.

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
  - Questions about external integrations or third-party services
  - Questions about access control
  - Questions about performance or scale expectations
  - These questions should clarify expected behavior with a deep understanding of the code.
  - Ask questions one at a time proposing the question with a smart default option

#### Step 5: Finalize Requirements
- Generate comprehensive requirements based on the template in `requirements.md`. Do not delete or modify existing sections, honor the template.
- Fill in `requirements.md` with the final requirements.
- Use `specware feature update-state <short-name> "Requirements Complete"`
- Offer three options:
  1. Interactive review session of the requirements documentation
  2. Stop here for asynchronous review and feedback, the user being expected to review the requirements document.
  3. Move to the next phase, skipping review (not reccomended).

#### Step 6: (optional) Interactive review session
- Use `specware feature update-state <short-name> "Requirements Interactive Review"`
- For each section of the `requirements.md` document, perform the following interactive review steps:
  1. Generate a 1-3 sentence summary of the section
  2. Display the section in two parts:
    a) Show the exact section content from the file (verbatim, no modifications or paraphrasing)
    b) Then show the generated summary as a separate block below the original content
  3. Ask the user directly for any changes or amendments to this section or if they'd like to consider this section approved and move onto the next.
  4. If the user provides changes or amendments, make the changes and make any additional changes needed to other sections of the document.
  5. Once the user is satisfied, proceed to the next section.
- Once all sections are approved, use `specware feature update-state <short-name> "Requirements Complete"`

### Phase 2: Technical Specification Creation
Technical specifications, such as OpenAPI, CLI reference, API output, Diagrams, data models, etc often help refine requirements before planning the implementation.

#### Step 1: Determine Necessary Technical Specs
- Present three options to the user to generate technical specifications to further refine requirements:
  1. Provide specific technical specification types you'd like to generate and review
  2. Use an agent to determine the best technical specification to generate
  3. Skip creating specific technical specifications (not recommended)
- **Wait for user selection before proceeding**
- If user selects (1): Ask the user to provide the technical specification types, providing a single suggestion, then **wait for response** before moving to the next step
- If user selects (3): Move to the next phase
- If user selects (2): Use the tech-spec-beck agent to determine the top 1-2 technical specifications, then **present the agent's recommendations to the user and ask for approval** before proceeding

#### Step 2: Generate Technical Specifications
- Present the technical specification types (from user input or approved agent recommendations)
- **Ask user to confirm which specifications to generate before proceeding**
- **After user approval**, for each approved specification:
  - Review the requirements again to determine technical specification details
  - Generate the technical specifications content
  - Store the technical specification in the spec sub-directory for this feature
  - Use the file format that makes the most sense (OpenAPI: YAML, for example)
  - Keep the technical specification document limited to only the technical details - do not add summaries, descriptions, or other text
  - **Show the generated specification to the user**
  - **Ask for approval before proceeding to the next specification**

#### Step 3: Interactive Review
- Generate the most important 1-2 questions that would provide additional clarification needed for the specification.
  - Ask questions one at a time proposing the question with a smart default option
- Display the technical specification without modification to the user, then ask questions.
- Consider the answers and update the specification as needed.
- Display the technical specification without modification to the user again
- Ask the user directly if they'd like to make any changes or amendments or if they'd like to move onto the next step -> integrating the changes into the requirements documentation.

#### Step 4: Requirements Integration
- With the new technical specification(s), consider each section of the requirements file and update them as necessary to reflect changes in the requirements.
- Inform the user of the changes made.
- Offer two options:
  1. Stop here for asynchronous review and feedback of the updated requirements document.
  2. Move to implementation planning phase

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
  - Questions about testing requirements
  - Write all questions to `context-implementation-plan.md` with smart defaults and example code snippets where necessary
  - Ask questions one at a time proposing the question with a smart default option and examples
- Only after all questions are asked, record answers in `context-implementation-plan.md` as received, not before.

#### Step 4: Testing Q&A
- Review what testing exists for similar features and determine what unit, integration, and e2e tests may be necessary for this feature.
- Generate at least 2 of the most important yes/no questions to clarify testing requirements, considering:
  - Questions about unit, integration, and e2e testing
  - Questions about when to write which tests (Test Driven Development for example)
  - Questions about how to run specific tests if required and unclear
  - Write all questions to `context-implementation-plan.md` with smart defaults and example code snippets where necessary
  - Ask questions one at a time proposing the question with a smart default option and examples
- Only after all questions are asked, record answers in `context-implementation-plan.md` as received, not before.

#### Step 5: Finalize Implementation Plan
- Generate a comprehensive implementation plan, breaking out large operations and changes into smaller tasks
- Be detailed in steps regarding testing:
  - What tests specifically will be run?
  - What output are you expecting?
- Write the complete implementation plan to `implementation-plan.md`
- Update status with `specware feature update-state <short-name> "Implementation Plan Generated"`

#### Step 6: Identify Scope Creep
- Use the scope-creep-craig agent to determine any areas in the implementation plan that may exceed the approved requirements.
- Consider scope-creep-craig's most important feedback:
  - Ignore feedback that is minor.
  - Collect no more than 2-3 suggestions.
- Display the suggested changes to the implementation plan to the user and offer to make them all, none, or individual changes they desire
- Update the implementation plan as the user requested.

#### Step 7: Interactive Review
- Use `specware feature update-state <short-name> "Implementation Plan Interactive Review"`
- For each section or phase of the `implementation-plan.md` document, perform the following interactive review steps:
  1. Generate a 1-3 sentence summary of the section.
  2. Display the section in two parts:
    a) Show the exact section content from the file (verbatim, no modifications or paraphrasing)
    b) Then show the generated summary as a separate block below the original content
  3. Ask the user directly for any changes or amendments to this section or if they'd like to consider this section approved and move onto the next.
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
  .claude/agents/                            # Claude Code agents for specialized workflows
  .spec/                                     # Feature specifications (001-feature-name/, 002-another-feature/)
  .spec/001-feature-name/.spec-status.json    # Feature status tracking

**Typical Workflow**

  1. specware feature new-requirements user-auth - Start new feature
  2. specware feature new-implementation-plan user-auth - Add implementation planning
  3. specware feature update-state user-auth "implementation-complete" - Track progress
