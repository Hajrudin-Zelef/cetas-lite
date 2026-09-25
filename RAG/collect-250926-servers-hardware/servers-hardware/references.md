---
id: collect-250926-servers-hardware/servers-hardware/references
title: "References"
domain: servers-hardware
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "agents"]
source: docs/RAG/clean4/references.md
source_anchor: ""
source_lines: [1, 225]
sha256: 87b7464a315db4904f5a380ad2065a4b9affca3e7003af43907cb969fc17f4fc
---

# References

References give OpenCode named access to directories outside the current
project. Add one in `opencode.json` or `opencode.jsonc`, then attach its alias in
your client when you need it.

After attaching `docs`, ask the agent to inspect the relevant path:

`Read the attached docs reference and summarize api/authentication.md.`
Use references for documentation, shared libraries, examples, or source from another repository.

## Local

Use `path` for a local directory:

```
{
  "references": {
    "design-system": {
      "path": "../design-system",
      "description": "Use when working with components or design tokens",
    },
  },
}
```
Paths behave as follows:

- Relative paths resolve from the directory containing the config file.
- Absolute paths are supported.
- Home-relative paths such as `~/docs` are supported.

Use the string shorthand when no other fields are needed:

```
{
  "references": {
    "docs": "../docs",
    "shared": "~/work/shared",
  },
}
```
## Git

Use `repository` for a remote Git repository:

```
{
  "references": {
    "effect": {
      "repository": "Effect-TS/effect",
      "branch": "main",
    },
    "internal-sdk": {
      "repository": "git@gitlab.example.com:platform/sdk.git",
      "branch": "release/v2",
    },
  },
}
```
Supported remote forms include:

- GitHub `owner/repo` shorthand
- Git URLs
- Host/path forms
- SCP-style remotes

Local `file:` repositories are not supported.

Use the string shorthand to follow the remoteâs default branch:

```
{
  "references": {
    "effect": "Effect-TS/effect",
    "sdk": "gitlab.com/platform/sdk",
  },
}
```
Without `branch`, OpenCode checks out and refreshes the remoteâs default branch.
Branch names may contain letters, numbers, `/`, `_`, `.`, and `-`, but cannot
start with `-` or contain `..`.

```
{
  "references": {
    "sdk": {
      "repository": "example/sdk",
      "branch": "release/v2.1",
    },
  },
}
```
## Storage

OpenCode normalizes each remote and stores one checkout per remote and branch under its global data directory. Without an explicit branch, the path uses this shape:

`opencode/repos/<host>/<repository-path>`
On a typical Linux installation, `Effect-TS/effect` is stored at:

`~/.local/share/opencode/repos/github.com/Effect-TS/effect`
An explicit branch adds an encoded `@<branch>` suffix:

`~/.local/share/opencode/repos/github.com/Effect-TS/effect@main`
## Refresh

Missing repositories are cloned asynchronously. Existing checkouts are checked in the background when references load or reload and after a new user prompt is admitted in their Location.

`Prompt admitted â refresh check starts in background â agent continues`
A checkout is eligible when its last refresh attempt was at least 24 hours ago, or when no attempt has been recorded. A refresh fetches and resets to the configured branch or the remoteâs default branch.

```
Last attempt: 25 hours ago â eligible
Last attempt: 2 hours ago  â skipped
```
Refresh behavior follows these rules:

- Timestamps persist across service restarts.
- Locations using the same checkout share its timestamp.
- Failed attempts are logged and remain subject to the 24-hour limit.
- There is no periodic polling while a Location is unused.
- Clone or refresh failures do not stop other references from loading.

Prompts do not wait for cloning or refreshing. An attachment can contain older content even if a later tool read sees the updated checkout, and a new reference can appear before its checkout is ready.

```
Attachment created â cached content
Background refresh completes â later tool read sees newer content
```
## Guidance

`description` tells agents when a reference is relevant. References with a
description appear in agent instructions with their alias and resolved path.

```
{
  "references": {
    "docs": {
      "path": "../docs",
      "description": "Use for product behavior and terminology",
    },
  },
}
```
References without a description remain available to clients but are not advertised automatically.

```
{
  "references": {
    "archive": {
      "path": "../archive",
    },
  },
}
```
## Visibility

Set `hidden` to `true` to remove a reference from interactive client selectors:

```
{
  "references": {
    "internal": {
      "path": "../internal",
      "description": "Use for internal service behavior",
      "hidden": true,
    },
  },
}
```
`hidden` affects only interactive visibility. The reference remains in the
reference API and, when it has a description, in agent instructions.

## Usage

Clients attach a reference by its root alias. The attachment contains a non-recursive listing of the rootâs immediate files and directories.

```
Attached alias: docs
Ask: Inspect docs/guides/deployment.md and explain the deployment steps.
```
Ask the agent to inspect a specific path when you need content below the root.

## Permissions

References do not grant extra tool permissions. Access outside the active
Location still follows normal tool rules and the `external_directory`
permission; editing also requires the applicable edit permission.

`Read ../product-docs/api.md`
The request succeeds only when the active permissions allow that external read.

## Fields

| Field | Local | Git | Description | 
|---|---|---|---|
| `path` | Required | No | Local directory path | 
| `repository` | No | Required | Remote Git repository | 
| `branch` | No | Optional | Branch to fetch and check out | 
| `description` | Optional | Optional | Guidance describing when agents should use it | 
| `hidden` | Optional | Optional | Hide it from interactive client selectors | 

A complete Git entry can use every Git-compatible field:

```
{
  "references": {
    "sdk": {
      "repository": "example/sdk",
      "branch": "main",
      "description": "Use for SDK implementation details",
      "hidden": false,
    },
  },
}
```
## Aliases

An alias cannot be empty or contain `/`, `\`, whitespace, a backtick, or a
comma.

```
Valid:   docs
Invalid: product docs
Invalid: docs/api
```
