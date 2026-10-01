---
id: collect-261001-general-networking/general-networking/use-the-web-ui
title: "Use the Web UI"
domain: general-networking
role: reference
task: reference
actors: ["DeepSeek", "OpenAI"]
dates: []
keywords: ["agent", "deepseek"]
source: docs/RAG/collect-261001-general-networking/use-the-web-ui.md
source_anchor: ""
source_lines: [1, 22]
sha256: 56b41f57b686362cbabf3a443d133ca657e355f3656f137d1001c642dcf8ddb5
---

# Use the Web UI 

Start the Web UI through the root README; the command prints its URL. This guide begins after that server is running. The `dsh` process uses its invoking directory as the default filesystem location, but a fresh Web UI has no selected workspace until you add one.

## Configure a model 

Open **Settings → Models**, enter a DeepSeek API key, and save it. The model route becomes usable immediately without restarting the server.

The model configuration guide covers other providers and custom OpenAI-compatible endpoints.

## Choose a workspace 

Click **Choose workspace**, add the project directory where you started `dsh`, and select it. The session composer remains unavailable until a workspace is selected.

## Run a task 

Start a session and send:

Summarize this repository and identify its main packages.


The agent can read and edit workspace files, run commands, delegate work, and maintain a plan. The Web UI asks before operations that require approval under the active permission policy.
