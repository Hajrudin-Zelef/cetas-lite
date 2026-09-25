---
id: collect-250926-servers-hardware/servers-hardware/mcp-servers
title: "MCP servers"
domain: servers-hardware
role: reference
task: reference
actors: []
dates: ["2025-11-25", "2026-07-28"]
keywords: ["mcp", "agent", "model context protocol"]
source: docs/RAG/clean4/mcp-servers.md
source_anchor: ""
source_lines: [1, 243]
sha256: 8f691fefbc263dc06eaed91a996ee0e406e93f61f675cf3a55be34eec37f702d
---

# MCP servers

OpenCode connects to Model Context Protocol servers and exposes capabilities such as tools, prompts, resources, and instructions. MCP tools consume model context, so add only the servers you need.

## Setup

Add a remote server from the project that should use it, then check its connection:

```
opencode mcp add context7 --url https://mcp.context7.com/mcp
opencode mcp list
```
The command writes the server to the project configuration. Add `--global` to make it available in every project:

`opencode mcp add context7 --global --url https://mcp.context7.com/mcp`
Remote servers use OAuth by default. If the list shows `needs authentication`, open OpenCode, run `/mcps`, select the server, and sign in. A connected server is ready for an agent to use:

`â context7  connected`
## Config

To configure a server by hand, give it a unique name under `mcp.servers`. V2 does not place server names directly under `mcp`.

Servers connect automatically. Use `disabled`, not an `enabled` field, to keep one configured without connecting it:

```
{
  "mcp": {
    "servers": {
      "my-server": {
        "type": "local",
        "command": ["npx", "-y", "example-mcp-server"],
        "disabled": true,
      },
    },
  },
}
```
A higher-precedence project config replaces the entire server object with the same name. Use different names for separate connections or accounts; otherwise repeat every required field in the override:

## Local

A local server is a command that OpenCode starts over the MCP stdio transport. Add one with a command after `--`:

`opencode mcp add everything -- npx -y @modelcontextprotocol/server-everything`
Use configuration for process options such as a working directory or environment variables:

| Field | Required | Description | 
|---|---|---|
| `type` | Yes | Must be `"local"` . | 
| `command` | Yes | Executable followed by its arguments. | 
| `cwd` | No | Process directory. Relative paths resolve from the workspace, which is also the default. | 
| `environment` | No | String variables added to OpenCodeâs inherited process environment. | 
| `disabled` | No | Prevents connection when `true` . Defaults to`false` . | 
| `codemode` | No | Set to `false` to expose tools directly instead of through Code Mode. Defaults to`true` . | 
| `timeout` | No | Per-server timeout overrides. | 
| `protocol` | No | `legacy` (default),`auto` , or`2026-07-28` . See Protocol version. | 

Use `{env:NAME}` for environment substitution. Shell expressions such as `$NAME` are not expanded in JSON strings:

```
{
  "environment": {
    "MCP_API_KEY": "{env:MCP_API_KEY}",
  },
}
```
## Remote

A remote server uses the MCP Streamable HTTP transport and requires an absolute URL:

`opencode mcp add context7 --url https://mcp.context7.com/mcp`
Use configuration when the server needs headers or other options. Store secrets in environment variables rather than in the file:

| Field | Required | Description | 
|---|---|---|
| `type` | Yes | Must be `"remote"` . | 
| `url` | Yes | Absolute Streamable HTTP endpoint. | 
| `headers` | No | String HTTP headers sent to the endpoint. | 
| `oauth` | No | OAuth settings, or `false` to disable OAuth. | 
| `disabled` | No | Prevents connection when `true` . Defaults to`false` . | 
| `codemode` | No | Set to `false` to expose tools directly instead of through Code Mode. Defaults to`true` . | 
| `timeout` | No | Per-server timeout overrides. | 
| `protocol` | No | `legacy` (default),`auto` , or`2026-07-28` . See Protocol version. | 

Use `oauth: false` only when the server exclusively uses an API key or another header credential:

```
{
  "type": "remote",
  "url": "https://mcp.example.com/mcp",
  "oauth": false,
  "headers": { "Authorization": "Bearer {env:MCP_API_KEY}" },
}
```
## OAuth

OAuth is enabled for remote servers unless `oauth` is `false`. OpenCode discovers the authorization server, uses PKCE, refreshes tokens, and attempts dynamic client registration when supported; credentials stay outside project configuration.

For dynamic registration, configure only the server URL:

If the server needs authentication, run `/mcps`, select it, and complete authorization in the browser. The CLI can start the same flow:

`opencode mcp auth sentry`
When a provider gives you client credentials, use V2âs snake_case OAuth fields:

| Field | Description | 
|---|---|
| `client_id` | Pre-registered client ID. Omit it to attempt dynamic registration. | 
| `client_secret` | Secret for a pre-registered client. | 
| `scope` | Space-delimited scopes to request. | 
| `callback_port` | Local callback port from `1` through`65535` . An available ephemeral port is the default. | 
| `redirect_uri` | Pre-registered loopback URI whose path and port reach the local callback listener. | 
| `auth_server_metadata_url` | URL of the authorization serverâs OAuth or OpenID Connect metadata document. Set it when the MCP server does not publish protected resource metadata that names its authorization server. | 

Remove stored OAuth credentials when you need to sign in again or switch accounts:

`opencode mcp logout sentry`
## Timeouts

Timeouts are positive integer milliseconds. Set defaults under `mcp.timeout`; a serverâs `timeout` object overrides matching defaults.

| Timeout | Default | Applies to | 
|---|---|---|
| `startup` | 30 seconds | Transport connection and server initialization. | 
| `catalog` | 30 seconds | Listing tools, prompts, resources, and resource templates. | 
| `execution` | 12 hours | Tool calls, prompt retrieval, and resource reads. | 

## Protocol version

OpenCode opens every server with the classic MCP `initialize` handshake by default. Set a serverâs `protocol` to talk to one built on the 2026-07-28 revision:

| Value | Behavior | 
|---|---|
| `legacy` | Default. Sends `initialize` and speaks protocol revisions up to 2025-11-25. | 
| `auto` | Probes with `server/discover` for the 2026-07-28 revision and falls back to`legacy` when the server does not support it. | 
| `2026-07-28` | Requires the 2026-07-28 revision. The connection fails against older servers. | 

Probing a local `legacy` server adds a short-lived extra process and can wait up to the startup timeout when the server ignores unknown requests, so keep `auto` scoped to servers that need it.

## Names

OpenCode names a tool `<server>_<tool>`. It replaces characters other than letters, numbers, `_`, and `-` with `_`:

```
server: context 7
tool:   resolve.library/id
name:   context_7_resolve_library_id
```
MCP prompts become commands named `<server>:<prompt>` with the same normalization. For example:

`/context_7:find_docs`
Choose short server names that remain unique after normalization. Under the default Code Mode, tools are grouped by the normalized server name:

`tools.context_7.resolve_library_id(...)`
## Permissions

Code Mode is the default. Set `codemode` to `false` when a serverâs tools must stay on the providerâs native tool list:

```
{
  "mcp": {
    "servers": {
      "context7": {
        "type": "remote",
        "url": "https://mcp.context7.com/mcp",
        "codemode": false,
      },
    },
  },
}
```
Use permission actions to hide or deny tools without disconnecting their server. Match the normalized `<server>_<tool>` name:

```
{
  "permissions": [
    {
      "action": "context7_*",
      "resource": "*",
      "effect": "deny",
    },
  ],
}
```
## Context

For calls made on behalf of a session, OpenCode sends the session ID in `CallToolRequest.params._meta.sessionID`. This applies to direct tools and Code Mode over stdio and Streamable HTTP:

```
{
  "method": "tools/call",
  "params": {
    "name": "lookup",
    "arguments": { "query": "example" },
    "_meta": { "sessionID": "ses_..." }
  }
}
```
The ID is request metadata, not a tool argument, so it is absent from the model-visible schema. Treat it as an opaque correlation value:

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
