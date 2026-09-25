---
id: collect-250926-servers-hardware/servers-hardware/how-to-use-mcp-servers-with-local-llms-unsloth-documentation
title: "how-to-use-mcp-servers-with-local-llms-unsloth-documentation"
domain: servers-hardware
role: reference
task: reference
actors: ["AMD", "Alibaba", "Hugging Face", "Intel", "Nvidia", "OpenAI", "Unsloth"]
dates: []
keywords: ["mcp", "amd", "cost", "embedding", "embeddings", "gguf", "inference", "intel", "llama", "llama.cpp", "memory", "model context protocol"]
source: docs/RAG/clean4/how-to-use-mcp-servers-with-local-llms-unsloth-documentation.md
source_anchor: ""
source_lines: [1, 138]
sha256: 7b289dcff23d5465eee5ef9843f4fc5b6fe2f79db0e37b05d0e97ffd84ae3b7b
---

# how-to-use-mcp-servers-with-local-llms-unsloth-documentation

This step-by-step guide shows you how to connect **Model Context Protocol (MCP)** servers to local LLMs like Qwen or Gemma, so any model you run can call external tools and services via MCP. Connecting MCP to a local model lets it securely use your local files, apps, databases, and tools instead of only chatting from memory, to build a more useful, private, and interchangeable AI assistant that can act on your real environment.

We'll use the open-source repos Unsloth and llama.cpp as they are popular frameworks for local model inference/deployment. MCP works for local GGUF models and cloud provider models. **We'll also show how** **multiple MCP Servers** **can be utilized.**

MCP tools work alongside other model capabilities such as code execution and web search, so a single model can search the web, run code, and call your connected services in the same thread.

Once an MCP server is connected, you can ask your local model to do many automated tasks. A few examples:

- **Search docs:** “Find the relevant docs and summarize the setup steps.” - Context7 can be used.
- **Analyze a codebase:** “Map out this repo and explain where authentication, billing, and data access happen.” - The GitHub official MCP and GitMCP can be used to analyze repos.
- **Search the web with embeddings** - Exa's MCP Server can be used for semantic web searches with contextual embedding support.
- **Debug websites UIs** - The Playwright and Chrome DevTools MCP Servers can be used to drive websites to find fixes to issues.

We'll be using two ways to connect your local model on your device to MCP Servers. Both use open-source packages: Unsloth and llama.cpp to run, serve and deploy your model.

In this example we'll use Unsloth to connect any local model like Qwen3.6 or Gemma 4 with the MCP Servers: Vercel, Context7, Exa and Hugging Face. We'll then ask a model what it can do with it. The same steps work for any MCP server.

The easiest way to get started is by installing the Unsloth Desktop app. It supports MacOS, Linux, Windows, NVIDIA, AMD, Intel and CPU setups.

Or, if you prefer manual installation:

**MacOS, Linux, WSL:**

`curl -fsSL https://unsloth.ai/install.sh | sh`
**Windows PowerShell:**

`irm https://unsloth.ai/install.ps1 | iex`
See here for detailed instructions.

Click on "MCP" in the chat toolbar.

Unsloth Studio by default has MCP support for Context7, Exa and Hugging Face. Turning on Exa web search will disable the default search tool which we have.

To add the Vercel MCP server, click on "Add custom MCP", and you will get a pop-up:

Fill in the server details:

1. **Display name** : a friendly label, e.g.`Vercel` .
2. **URL** : the server's base endpoint, e.g.`https://mcp.vercel.com` .
3. Choose an authentication method below.

For servers that require browser-based authentication (GitHub, Linear, Vercel, etc.), turn on **Use OAuth sign-in**. A browser window will open on first connect so you can authorize Unsloth.

For servers that authenticate with a token, leave OAuth off and click **Add header** under **Custom headers**. Add an `Authorization` header with your token:

The server now appears in the MCP Servers list. Unsloth fetches its tools automatically and shows a confirmation, e.g. *Refreshed "Vercel" (18 tools)*.

Each server has controls to **toggle** it on/off, **refresh** its tools, **edit** it, or **delete** it. Make sure both the server's toggle and the **Use MCP Servers** master toggle are on, then close the dialog.

It'll be highlighted if enabled. You can also disable them by simply clicking on them again.

Pick any model from the **Select model** dropdown and start chatting. The model can now call the server's tools on its own when your request calls for it.

Above, a local `gemma-4-E2B-it-GGUF` was asked *"Can you use Vercel MCP server?"* and reported the actions it can take: managing projects, analysing logs, listing teams, generating access links, checking domains, and searching Vercel's documentation.

How about calling multiple MCP servers like 3? We'll use Unsloth Studio's default Exa, Context7 & Hugging Face provided MCP servers and enable all 3.

If you ask "Can Unsloth support Qwen finetuning", Exa will provide great details on it:

Then as a follow up "Search Unsloth docs on how to do this", and Context7 is used for docs:

Then type "Search Hugging Face for unsloth/Qwen models", and Hugging Face's MCP Server will be called:

Once a server is connected, ask the model to do real work in plain language. A few examples with the Vercel server:

- **Debug a failing build** :*"Pull the build logs for my latest deployment and tell me why it failed."*
- **Check deployment status** :*"List my most recent deployments and their status."*
- **Search the docs** :*"Search the Vercel docs for how to set up a custom domain."*
- **Domain research** :*"Is* *myproject.dev*
 *available, and how much would it cost?"*

We are using Gemma 4 E4B GGUF in this example:

For the bigger Gemma 4 26B-A4B model:

For more inference parameter adjustments, see our Gemma 4 guide.

You can test the server:

Copy the absolute path. Create a separate project folder for your MCP host:

Create `server_config.json`:

Replace the path with your real workspace path. IBM’s `mcp-cli` docs use this same `server_config.json` shape and the same `npx -y @modelcontextprotocol/server-filesystem /path/to/allowed/files` filesystem configuration.

Use IBM’s `mcp-cli`; it is a command-line MCP client/host with chat mode, tool discovery, and custom OpenAI-compatible provider support. Its docs recommend `uvx mcp-cli --help`, support project `server_config.json`, and support runtime custom OpenAI-compatible providers via `--api-base` and `--api-key`.

`mcp-cli` needs a config file to start up. So it's necessary to add a step to create it. For example:

Then run:

Then try:

Then:

`mcp-cli` has tool-call confirmation enabled by default, so you should see prompts before tool execution.

**Full code example:**

If a server fails to connect or its tools don't appear, check the URL is the server's base endpoint (e.g. `https://mcp.vercel.com`), not a docs or dashboard page. For OAuth servers, complete the browser sign-in when it opens; for token-based servers, verify the `Authorization` header and that the token is valid.

Click **Refresh** if tools don't show up after connecting, and make sure both the individual server toggle and the **Use MCP Servers** master toggle are on.

Only connect MCP servers you trust. Review requested permissions and keep human confirmation enabled for actions that read private data, change deployments, purchase domains, or modify projects. Be especially careful when combining MCP servers with web search or other tools, because prompt-injected content can try to trigger unwanted tool calls.

Here is a list of some popular and useful MCP Servers you can connect to:

**GitHub MCP**

Repos, issues, PRs, code search, Actions

**Context7**

Fresh library docs and examples

**Notion MCP**

Docs, notes, tasks, project knowledge

**Slack MCP**

Team conversation search

**Linear MCP**

Issues, projects, product workflows

**Vercel MCP**

Deployments, logs, docs, domains

**Sentry MCP**

Production debugging

Last updated

Was this helpful?
