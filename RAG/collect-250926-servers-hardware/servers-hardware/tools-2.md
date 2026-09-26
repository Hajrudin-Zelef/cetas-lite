---
id: collect-250926-servers-hardware/servers-hardware/tools-2
title: "Tools"
domain: servers-hardware
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "mcp"]
source: docs/RAG/clean4/tools.md
source_anchor: ""
source_lines: [189, 202]
sha256: 5f61a0ecad753212788b77fb2088b9eb82105fe4f8c0f583d4fc2d44edc555e9
---

# Tools

`Open https://example.com, take a snapshot of that tab, and report the main heading.`
The namespace includes tab and navigation commands, page snapshots and search, clicking and form input, screenshots, file transfer, console and network inspection, performance traces, heap inspection, and Lighthouse audits. Screenshots require a focused visible tab. Upload paths are server-local; captures return server-local paths, and each file transfer is limited to 5 MiB.

Page content, logs, headers, and response bodies are untrusted data, not agent
instructions. A `browser` deny rule with resource `*` removes the browser catalog;
browser operations do not issue individual permission prompts.

## Extensions

MCP servers add tools whose names and inputs come from each
connected server. Their permission action is `<server>_<tool>` with resource
`*`. They are not a fixed part of the built-in catalog.

Skills add task instructions rather than new executable tools. Attachments put user-selected content into a prompt, while References provide named outside directories. These features complement tools without bypassing their permission checks.
