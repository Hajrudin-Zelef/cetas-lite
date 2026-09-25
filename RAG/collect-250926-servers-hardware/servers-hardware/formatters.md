---
id: collect-250926-servers-hardware/servers-hardware/formatters
title: "Formatters"
domain: servers-hardware
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/clean4/formatters.md
source_anchor: ""
source_lines: [1, 83]
sha256: 77c3f7e2391b8feffe1c9fd40974c8d2cfca906f166f6c7edbd1337dfa1f2a0e
---

# Formatters

OpenCode can format files after its `write`, `edit`, or `patch` tools change
them. Formatters are disabled by default, so enable them in your configuration:

## Enable

Set `formatter` to `true` to enable every built-in formatter. OpenCode runs a
built-in only when its executable and any project-specific requirements are
available.

An object also enables the built-ins and lets you override them or add custom
formatters. An empty object is therefore equivalent to `true`.

## Builtins

OpenCode includes these formatter definitions. Most require the named command to be available; definitions with extra detection rules list them below.

| Formatter | Extensions | Requirement | 
|---|---|---|
| `gofmt` | `.go` | `gofmt` command | 
| `mix` | `.ex` ,`.exs` ,`.eex` ,`.heex` ,`.leex` ,`.neex` ,`.sface` | `mix` command | 
| `oxfmt` | `.js` ,`.jsx` ,`.mjs` ,`.cjs` ,`.ts` ,`.tsx` ,`.mts` ,`.cts` | `oxfmt` dependency in`package.json` | 
| `prettier` | `.js` ,`.jsx` ,`.mjs` ,`.cjs` ,`.ts` ,`.tsx` ,`.mts` ,`.cts` ,`.html` ,`.htm` ,`.css` ,`.scss` ,`.sass` ,`.less` ,`.vue` ,`.svelte` ,`.json` ,`.jsonc` ,`.yaml` ,`.yml` ,`.toml` ,`.xml` ,`.md` ,`.mdx` ,`.graphql` ,`.gql` | `prettier` dependency in`package.json` | 
| `biome` | Same extensions as `prettier` above | `biome.json` or`biome.jsonc` and an installed`@biomejs/biome` binary | 
| `zig` | `.zig` ,`.zon` | `zig` command | 
| `clang-format` | `.c` ,`.cc` ,`.cpp` ,`.cxx` ,`.c++` ,`.h` ,`.hh` ,`.hpp` ,`.hxx` ,`.h++` ,`.ino` ,`.C` ,`.H` | `clang-format` command and`.clang-format` | 
| `ktlint` | `.kt` ,`.kts` | `ktlint` command | 
| `ruff` | `.py` ,`.pyi` | `ruff` command and a Ruff config or dependency declaration | 
| `air` | `.R` | `air` command that identifies itself as the R formatter | 
| `uv` | `.py` ,`.pyi` | `uv` command with`uv format` support | 
| `rubocop` | `.rb` ,`.rake` ,`.gemspec` ,`.ru` | `rubocop` command | 
| `standardrb` | `.rb` ,`.rake` ,`.gemspec` ,`.ru` | `standardrb` command | 
| `htmlbeautifier` | `.erb` | `htmlbeautifier` command | 
| `dart` | `.dart` | `dart` command | 
| `ocamlformat` | `.ml` ,`.mli` | `ocamlformat` command and`.ocamlformat` | 
| `terraform` | `.tf` ,`.tfvars` | `terraform` command | 
| `latexindent` | `.tex` | `latexindent` command | 
| `gleam` | `.gleam` | `gleam` command | 
| `shfmt` | `.sh` ,`.bash` | `shfmt` command | 
| `nixfmt` | `.nix` | `nixfmt` command | 
| `rustfmt` | `.rs` | `rustfmt` command | 
| `pint` | `.php` | `laravel/pint` in`composer.json` | 
| `ormolu` | `.hs` | `ormolu` command | 
| `cljfmt` | `.clj` ,`.cljs` ,`.cljc` ,`.edn` | `cljfmt` command | 
| `dfmt` | `.d` | `dfmt` command | 

For example, enabling built-ins lets OpenCode discover and run a project-local Prettier dependency for matching files:

## Customize

Add a named entry to change a built-in or define a custom formatter. A custom
formatter needs both `command` and `extensions` to run.

| Field | Type | Behavior | 
|---|---|---|
| `disabled` | `boolean` | Removes the named formatter when `true` . | 
| `command` | `string[]` | Replaces the built-in command or defines a custom command. | 
| `environment` | `Record<string, string>` | Adds environment variables while preserving the parent environment. | 
| `extensions` | `string[]` | Replaces the built-in extension list or defines the custom list. Include the leading dot. | 

All fields are optional. A built-in entry inherits omitted values, while a new entry without a command or extensions cannot run.

## Commands

`command` is an argument array, not a shell command string. OpenCode replaces
`$FILE` with the fileâs absolute path and runs the command from the active
project directory.

## Matching

OpenCode compares the fileâs final extension with `extensions`. Matching is
case-sensitive, and compound entries such as `.part.md` do not match
`notes.part.md` because its final extension is `.md`.

When several formatters match, OpenCode tries them in registered order and stops after the first successful command. Built-ins retain their built-in order; custom entries follow in object order. If one exits unsuccessfully, OpenCode logs the failure and tries the next match.

## Disable

Omit `formatter` or set it to `false` to disable all formatting. An explicit
`false` can override a lower-priority configuration that enabled formatters.

To disable one built-in while leaving the others enabled, mark its named entry as disabled.
