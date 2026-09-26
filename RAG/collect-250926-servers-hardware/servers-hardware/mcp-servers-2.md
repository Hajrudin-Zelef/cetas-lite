---
id: collect-250926-servers-hardware/servers-hardware/mcp-servers-2
title: "MCP servers"
domain: servers-hardware
role: reference
task: reference
actors: []
dates: []
keywords: ["mcp"]
source: docs/RAG/clean4/mcp-servers.md
source_anchor: ""
source_lines: [201, 243]
sha256: cf9a011ba2fac3f8d8c93c719e5afc73273f37371c01cc94a72c7c2b98081f52
---

# MCP servers

| Rule | Behavior | 
|---|---|
| Identity | It identifies the invoking OpenCode session, not the MCP transport session. | 
| Presence | It can be absent for calls without session context. | 
| Security | Do not use it by itself for authentication or authorization. | 
| Privacy | Remote servers receive the raw ID and may log or retain it. | 

## Management

List servers and their current connection state from any project:

`opencode mcp list`
Use `/mcps` in OpenCode to view, connect, disconnect, or authenticate servers. Use the CLI to add servers and manage OAuth credentials:

```
opencode mcp add sentry --url https://mcp.sentry.dev/mcp
opencode mcp auth sentry
opencode mcp logout sentry
```
To remove a server, delete its entry from the project or global configuration where it was added:

```
{
  "mcp": {
    "servers": {},
  },
}
```
Edit configuration directly for OAuth client settings, timeouts, working directories, or persistent enablement:

```
{
  "mcp": {
    "servers": {
      "sentry": {
        "type": "remote",
        "url": "https://mcp.sentry.dev/mcp",
        "disabled": true,
      },
    },
  },
}
```
