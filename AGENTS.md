# AGENTS.md

Naming guidance for contributors and coding agents working in `snippetbox`.

This project follows idiomatic Go naming conventions in the spirit of Alex Edwards' guidance from:
`https://www.alexedwards.net/blog/go-naming-conventions`.

## Scope

- Apply these rules to packages, files, variables, functions, methods, structs, interfaces, constants, and tests.
- Prefer readability and consistency with the Go standard library over personal preference.
- Keep names simple and predictable.

## Core Rules

- Use short, meaningful names.
- Avoid stutter (for example, `snippet.Snippet` is usually unnecessary).
- Keep local variable names concise (`i`, `id`, `err`, `w`, `r` are fine when context is clear).
- Use longer names only when they improve clarity.
- Do not add type information to names when the type already communicates it (avoid names like `userMap map[string]User` when `users` is enough).

## Package Naming

- Package names should be:
  - lowercase
  - short
  - singular where sensible
  - descriptive of behavior/domain
- Do not use underscores, hyphens, or mixed casing in package names.
- Avoid generic package names like `common`, `utils`, or `helpers` unless the package has a clear, narrow purpose.
- Avoid repeating the package name in exported identifiers.
  - Good: package `snippet` with `Create`, `Store`, `Validate`
  - Less good: `snippet.CreateSnippet`, `snippet.SnippetStore`

## Function and Method Naming

- Name functions by behavior, not implementation details.
- Use verb or verb-phrase names for actions (`create`, `validate`, `parse`).
- Prefer consistent naming patterns for related handlers and operations.

Current examples in `cmd/web/handlers.go`:

- Good: `snippetView`, `snippetCreate`, `snippetCreatePost`
- Avoid introducing inconsistent variants like `createSnippetHandler` beside these names unless you rename the whole group for consistency.

## Variable Naming

- Keep variable names short for tight/local scope and longer for broader scope.
- Use `err` for errors unless multiple errors are in scope (then use `parseErr`, `dbErr`, etc.).
- Avoid single-letter names outside of small scopes or conventional cases.
- Avoid ambiguous abbreviations.

## Receiver Naming

- Use short receiver names, usually 1-2 letters.
- Receiver names should be consistent for a given type.
- Receiver names should hint at the type name.
  - Example: `func (s *SnippetStore) Insert(...)` not `func (this *SnippetStore) Insert(...)`
- Do not use `self` or `this`.

## Interface Naming

- Prefer small interfaces.
- Single-method interfaces should often use an `-er` suffix when natural (`Reader`, `Writer`, `Fetcher`).
- Multi-method interfaces should be named for capability, not implementation.
- Do not create interfaces prematurely; define them where they are consumed.

## Acronyms and Initialisms

- Keep common initialisms consistently cased: `ID`, `URL`, `HTTP`, `JSON`, `SQL`.
- Examples:
  - Good: `userID`, `requestURL`, `httpServer`
  - Avoid: `userId`, `requestUrl`, `HttpServer`

## Exported vs Unexported Names

- Export only what other packages must use.
- Keep package internals unexported by default.
- Exported names should be intentionally stable and clearly documented.

## Handler and Route Naming for This Repo

Given current routes in `cmd/web/main.go`:

- `GET /` -> `home`
- `GET /snippet/view/{id}` -> `snippetView`
- `GET /snippet/create` -> `snippetCreate`
- `POST /snippet/create` -> `snippetCreatePost`

When adding routes, follow the same naming style unless you perform a project-wide rename.

## Code Review Checklist (Naming)

Use this checklist in PR reviews:

- [ ] Are package names short, lowercase, and purpose-driven?
- [ ] Are function/method names behavior-based and consistent with nearby code?
- [ ] Is stutter avoided (`pkg.Type`, `pkg.Function`)?
- [ ] Are variable names concise but clear for their scope?
- [ ] Are receiver names short and consistent per type?
- [ ] Are initialisms cased correctly (`ID`, `URL`, `HTTP`)?
- [ ] Are interfaces minimal and capability-focused?
- [ ] Are exported identifiers only used when needed?

## Decision Rule

When in doubt, pick the name that a new Go developer would understand fastest when scanning the file once.

