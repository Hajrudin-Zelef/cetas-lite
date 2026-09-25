---
id: collect-250926-servers-hardware/servers-hardware/websearch
title: "Websearch"
domain: servers-hardware
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/clean4/websearch.md
source_anchor: ""
source_lines: [1, 47]
sha256: fca9fb40f8c4146856e4a6fc05e05f9666ac99b20ee7a78fd93bb7346bf072e9
---

# Websearch

Ask OpenCode for current information and it can search the web, read the results, and include source links in its response.

`Find the latest Bun release and summarize the changes.`
The first search asks you to allow web search and select a provider. OpenCode remembers your choice for later sessions.

## Providers

OpenCode includes four search providers:

| Provider | ID | Environment variable | 
|---|---|---|
| Exa | `exa` | `EXA_API_KEY` | 
| Firecrawl | `firecrawl` | `FIRECRAWL_API_KEY` | 
| Parallel | `parallel` | `PARALLEL_API_KEY` | 
| Tavily | `tavily` | `TAVILY_API_KEY` | 

Connect an account from the TUI with `/connect`, or set the providerâs environment variable before starting OpenCode.

`$ TAVILY_API_KEY=your-key opencode`
## Selection

Set a provider in `opencode.jsonc` to skip the selection prompt.

Use `"random"` to choose an available provider automatically. Each session keeps using its selected provider until that provider is rate limited.

## Limits

When a provider returns HTTP 429, random selection retries the search with another available provider. The rate-limited provider waits for its `Retry-After` period, or 60 seconds when the response does not include a valid period.

```
Exa returns HTTP 429
â Exa enters cooldown
â OpenCode retries with another provider
```
If every provider is cooling down, the search fails immediately. Moving a session or restarting its server clears its remembered provider; cooldowns are shared by sessions in the same workspace.

## Permissions

Web searches use the `websearch` permission action and the search query as the resource. Ask before searches that are not explicitly allowed.

See Permissions for rule order and matching.

## Disable

Set `websearch` to `false` to remove the web search tool from model requests.
