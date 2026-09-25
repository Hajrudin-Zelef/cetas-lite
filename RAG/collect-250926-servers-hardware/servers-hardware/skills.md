---
id: collect-250926-servers-hardware/servers-hardware/skills
title: "Skills"
domain: servers-hardware
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "agents", "claude", "license"]
source: docs/RAG/clean4/skills.md
source_anchor: ""
source_lines: [1, 165]
sha256: bc6862c358a4f01f78fe69641fe00abd72ed0d6d2996641c093d22e99f37ce93
---

# Skills

Create a skill to give an agent reusable instructions for a specific task. Put a `SKILL.md` file in `.opencode/skills/<skill-id>` and describe when to use it in `description`.

OpenCode advertises this skill when it is relevant. The agent can then load its instructions with the `skill` tool instead of adding every skill to every prompt.

## Create

Keep related scripts, references, and templates beside `SKILL.md`:

```
.opencode/skills/
âââ git-release/
    âââ SKILL.md
    âââ scripts/
    â   âââ changelog.ts
    âââ references/
        âââ release-policy.md
```
Paths written inside the skill are relative to the directory containing `SKILL.md`. The directory form is recommended because it gives supporting files a private base directory.

## Discovery

OpenCode automatically searches these locations:

| Scope | Sources | 
|---|---|
| Global | `~/.config/opencode/skills` | 
| Global compatibility | `~/.claude/skills` ,`~/.agents/skills` | 
| Project | `.opencode/skills` | 
| Project compatibility | `.claude/skills` ,`.agents/skills` | 

For project sources, OpenCode searches from the current directory up to the project root and includes matching directories at every level.

Each source can contain either form:

```
skills/
âââ review.md
âââ git-release/
    âââ SKILL.md
```
- Markdown files must be at the source root, such as `skills/review.md` .
- Files named exactly `SKILL.md` can be at any depth, such as`skills/git-release/SKILL.md` .

## Sources

Add more local directories or HTTP catalogs with the `skills` array in any `opencode.json` or `opencode.jsonc`:

| Value | Resolution | 
|---|---|
| Relative path | From the active OpenCode working directory, not the config file | 
| `~/` path | From the current userâs home directory | 
| Absolute path | Used as written | 
| `http://` or`https://` URL | Loaded as an HTTP catalog | 

Every discovered config file contributes its entries. `skills` arrays are combined rather than replaced.

## Catalogs

An HTTP catalog is a base URL with an `index.json` file:

For that entry, OpenCode downloads each file from `<base-url>/git-release/<file>`.

| Rule | Requirement | 
|---|---|
| Paths | Must be safe, relative, and same-origin | 
| Entry file | Include `SKILL.md` or`<name>.md` , such as`git-release.md` | 
| Updates | Increment `version` when files change so OpenCode refreshes its cache | 

Prefer the named Markdown form in an HTTP catalog. Each downloaded skill directory becomes a source root, so `git-release.md` has the ID `git-release`. A root-level `SKILL.md` currently has the literal ID `SKILL` in V2.

## Frontmatter

Use frontmatter to name the skill and control where it appears:

| Field | Behavior | 
|---|---|
| `name` | Display name; defaults to the path-derived ID | 
| `description` | Summary used to show the skill to the model | 
| `slash` | Set to `false` to hide the skill from interactive command catalogs | 
| `metadata.opencode/slash` | Boolean or `"true"` /`"false"` ; overrides`slash` | 
| `metadata.opencode/autoinvoke` | Set to `false` to omit the skill from the modelâs available list | 

All frontmatter is optional at runtime. Add a clear `description` when the model should discover the skill; skills without one are not advertised.

`opencode/autoinvoke: false` only hides the skill from the modelâs available list. The skill remains registered and can still be loaded explicitly by ID. V2 accepts portability fields such as `license` and `compatibility` but does not interpret them.

## IDs

The file path determines the skill ID. The frontmatter `name` is only a display label.

| File | ID | 
|---|---|
| `<source>/git-release.md` | `git-release` | 
| `<source>/git-release/SKILL.md` | `git-release` | 
| `<source>/teams/release/SKILL.md` | `release` | 

IDs are exact and case-sensitive. For portable skills, use a unique lowercase kebab-case ID of 1â64 characters and keep it aligned with the directory name:

`^[a-z0-9]+(-[a-z0-9]+)*$`
V2 currently does not enforce this pattern, the 1â64 character recommendation, a match between `name` and the directory, or a maximum description length.

## Precedence

Skills are selected by ID. If two sources define `git-release`, the source registered later supplies the skill that OpenCode loads:

```
~/.config/opencode/skills/git-release/SKILL.md  â lower precedence
.opencode/skills/git-release/SKILL.md           â loaded
```
Sources are registered from lower to higher precedence:

1. Built-in skills
2. `.claude/skills` , global first and then from the farthest ancestor toward the current directory
3. `.agents/skills` , global first and then from the farthest ancestor toward the current directory
4. `~/.config/opencode/skills`
5. Project `.opencode/skills` , from the project root toward the current directory
6. Explicit `skills` config entries, in config priority and array order

Avoid duplicate IDs unless you intend to override an earlier skill.

## Loading

At each model step, OpenCode lists permitted skills that have a description and do not set `opencode/autoinvoke` to `false`. The list includes only the ID, name, and description, not the full Markdown body.

The model loads a skill by calling the `skill` tool with its exact ID:

OpenCode then:

1. Selects the current definition for that ID.
2. Checks the selected agentâs `skill` permission.
3. Adds the Markdown body, without frontmatter, to the conversation.
4. Provides the skillâs base directory and a sample of up to ten supporting file paths.

Supporting file contents are not loaded automatically. The agent reads them when the skill directs it to do so. The file sample is available for directory-based `SKILL.md` skills; flat Markdown skills do not receive a neighboring file list.

## Permissions

Use the `skill` action and the skill ID as the resource. Rules run in order, and the last matching rule wins:

| Effect | Behavior | 
|---|---|
| `allow` | Advertises and loads matching skills without approval | 
| `ask` | Advertises matching skills and asks before loading them | 
| `deny` | Hides matching skills from the model and rejects loading | 

Place the same rules under `agents.<id>.permissions` to apply them only to one agent.

## Troubleshooting

For example, this skillâs ID is `release`, not its frontmatter name `Git Release`:

```
.opencode/skills/release/SKILL.md
                      ââ ID: release
```
If a skill is missing or loads the wrong content:

1. Confirm the file is either a root-level `*.md` or a nested file named exactly`SKILL.md` .
2. Check the path-derived, case-sensitive ID rather than the frontmatter `name` .
3. Add a `description` if the model should discover the skill.
4. Check `opencode/autoinvoke` and the selected agentâs`skill` permissions.
5. Look for a later source that defines the same ID.
6. For HTTP catalogs, verify `index.json` , same-origin file paths, and a changed`version` .
