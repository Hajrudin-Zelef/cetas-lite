---
id: collect-250926-servers-hardware/servers-hardware/permissions-2
title: "Permissions"
domain: servers-hardware
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/clean4/permissions.md
source_anchor: ""
source_lines: [230, 251]
sha256: acf27bc39ad7a7f96d4068375b0186f407ff93b85b0a9bcbd2e93317e5de4dde
---

# Permissions

`shell: git status * â allow`
Saved approvals are durable, project-scoped `allow` rules. They never override
a configured `deny`. Tools choose the proposed saved pattern: some propose `*`,
shell proposes command prefixes, and skills and subagents propose their IDs.
Review broad approvals and remove those no longer needed.

Clients may attach feedback when rejecting. Non-interactive clients must decide
how to handle approval requests; configured `deny` rules always remain enforced.

## Policies

A policy can hard-deny a permission check after these rules and saved
approvals run. It turns `allow` or `ask` into `deny` and never grants access.

```
{
  "experimental": {
    "policies": [{ "action": "permission", "resource": "shell:sudo *", "effect": "deny" }],
  },
}
```
Global and Console-managed policies override project configuration, which is how an organization blocks a command that a repository would allow.
