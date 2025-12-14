# Copilot Instructions (Go · templ · SQL)

## Core Philosophy

Optimize for clarity, locality, and long-term maintainability.

The primary audience for the code is a future developer reading it for the first time.
The secondary audience is the compiler.

Prefer simple, explicit solutions over clever, abstract, or flexible ones.

---

## General Coding Principles

- Keep the code simple, explicit, and easy to understand.
- Prioritize readability and maintainability over cleverness, brevity, or reuse.
- Avoid clever code. If a solution feels smart, simplify it.
- Use clear, descriptive, and explicit naming, even if names become long.
- Prefer longer names over comments.
- Write hardened, production-ready Go code:
  - Explicit error handling
  - No ignored errors
  - Clear failure paths
- Favor high cohesion:
  - Code that changes together should live together.
- Favor low abstraction:
  - Introduce abstractions only when they solve a concrete, existing problem.
  - Do not abstract for hypothetical reuse.

---

## Comments & Documentation

- Do not add comments unless the code is not obvious from naming and structure.
- Do not explain what the code does if it is clear.
- Only comment why something non-obvious is done.
- Prefer clear naming and structure over explanatory comments.

If comments seem necessary, reconsider the structure or naming first.

---

## Refactoring Rules

- Follow Martin Fowler’s Rule of Three:
  - One instance: do nothing.
  - Two instances: duplication is acceptable.
  - Three instances: refactor.
- Prefer duplication over premature abstraction.
- Do not extract helpers, interfaces, or packages unless duplication or complexity already exists.

---

## Go-Specific Guidelines

- Follow standard Go conventions where they improve clarity.
- Prefer concrete types over interfaces.
- Introduce interfaces only at the point of use.
- Avoid “helpers”, “utils”, or “common” packages.
- Prefer explicit parameters over hidden state.
- Avoid magic behavior.

---

## File & Folder Structure

- Keep the folder structure shallow.
- Avoid nesting unless it clearly improves understanding.
- Folder names must be clear and intention-revealing.
- The root directory should immediately communicate what the project does and how it is structured.

Only create subfolders when:
- A feature contains multiple closely related files, or
- Grouping clearly improves comprehension.

Do not create folders for single files.

---

## Feature-Based Structure (templ)

Group by feature, not by technical role.

Simple feature:
slides/simple_slide.tmpl

Complex feature:
slides/more_complex_slide/
├── more_complex_slide.tmpl
└── more_complex_slide_helper.go

Keep all logic for a feature together.

---

## templ + SQL Workflow

- Keep templates close to their data access logic.
- If a template both reads from and writes to the database, keep both operations in the same file or feature folder when practical.
- Avoid splitting read/write logic across packages when they belong to the same UI workflow.

---

## SQL Usage

- Write explicit SQL.
- Prefer clarity over clever queries.
- Keep SQL close to where it is used.
- Handle all database errors explicitly.
- Make failure cases obvious.

---

## Over-Engineering Guardrails

- Do not introduce interfaces, patterns, or abstractions unless required by existing duplicated usage.
- Do not optimize for hypothetical future needs.
- Do not split code across files purely for cleanliness if it reduces locality.

If deviating from these rules, briefly explain why.

## Template Instructions
Here is a link to the go templ template llm instructions: https://templ.guide/llms.md
