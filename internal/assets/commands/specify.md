# Specify - Spec-driven Development Workflow

## Summary
Guide the user through a comprehensive spec-driven development workflow for feature requirements gathering and implementation planning using the specware tool.

## Important Notes
- Always use the specware tool to track state and create artifacts
- Maintain one feature at a time in active development
- Support stopping and resuming at any point in the workflow
- Record all Q&A interactions in the appropriate q&a files
- Follow the existing codebase patterns and conventions

## Specware usage Guidance
The specware tool is necessary to facilitate the basic operations of the workflow.

If the specware tool is not avaialble, immediately stop and instruct the user to install the tool.

## Workflow

### Phase 1: Requirements Building
When the user provides a feature request:

1. **Create Feature Specification Directory**
   - Use `specware feature new-requirements <short-name>` to create the feature directory
   - Generate a descriptive short-name based on the feature description

2. **Requirements Gathering**
   - Fill in the basic sections and metadata of the requirements spec
   - Create initial content in both `requirements.md` and `q&a-requirements.md`
   - Generate 3-5 follow-up questions to clarify requirements within the context of the repository
   - Ask questions one-by-one and record both questions and answers in `q&a-requirements.md`

3. **Expert Research Phase**
   - Research the codebase to become an "expert" on the relevant topics
   - Ask 3-5 more follow-up questions based on the research
   - Update `q&a-requirements.md` at each step

4. **Finalize Requirements**
   - Write the final requirements to `requirements.md`
   - Use `specware feature update-state <short-name> "Requirements Complete"`

5. **Next Steps Options**
   - Offer three options:
     1. Interactive review session of the requirements documentation
     2. Stop here for asynchronous review and feedback
     3. Move to implementation planning phase

### Phase 2: Implementation Planning
When the user is ready for implementation planning:

1. **Create Implementation Plan**
   - Use `specware feature new-implementation-plan <short-name>` to create the plan spec
   - Update status with `specware feature update-state <short-name> "Implementation Planning"`

2. **Codebase Analysis**
   - Understand the existing codebase structure and patterns
   - Define an implementation plan including technical details

3. **Planning Q&A**
   - Ask 3-5 questions about implementation approach
   - Ask 3-5 more detailed technical questions
   - Update `q&a-implementation-plan.md` at each step

4. **Finalize Plan**
   - Write the complete implementation plan to `implementation-plan.md`
   - Update status with `specware feature update-state <short-name> "Implementation Planning Complete"`
   - Offer interactive review option

### Continuation Commands
- For requirements feedback: "I have feedback on the requirements doc we need to address. <feedback>"
- To continue to implementation planning: "continue to the implementation planning phase"

