---
id: collect-mindstudio/mindstudio/docker-sandboxes-ai-agents
title: "Docker Sandboxes: How to Safely Run AI Coding Agents"
domain: mindstudio
role: reference
task: article
actors: ["AWS", "Anthropic", "Google", "OpenAI", "OpenRouter"]
dates: ["2026-08", "2026-09-23"]
keywords: ["agent", "agents", "sandbox", "aws", "claude", "memory"]
source: docs/RAG/Collect RAG/02_mindstudio/docker-sandboxes-ai-agents.md
source_anchor: ""
source_lines: [1, 52]
sha256: 5e8c50a744cf224bdc9f6a44edc6469fd695921c028760d387442b8f390ebc1d
---

# Docker Sandboxes: How to Safely Run AI Coding Agents

## Metadata

- **Source**: https://www.mindstudio.ai/blog/docker-sandboxes-ai-agents
- **Site**: MindStudio
- **Type**: Article
- **Language**: en
- **Verification status**: ✅ reachable
- **Collection date**: 2026-09-23

## Full summary

This article explains **Docker Sandboxes** — isolated micro VMs built to run AI coding agents like Claude Code, Codex, or custom agent frameworks without giving them free rein over the host machine. Each sandbox boots its own Linux kernel, runs a Docker engine inside it, and lets users set policies for network access, file read/write permissions, and credential handling. When done, the whole sandbox is discarded and the host system remains untouched.

**Why agents need a sandbox.** Coding agents with elevated permissions are efficient but risky. Tools like Claude Code or Codex offer a "skip permissions" mode that stops approve/deny prompts, meaning the agent can in principle touch anything: delete files, modify configs, make unauthorized network calls. Risk compounds with more autonomous or custom agents (frameworks like Hermes or OpenClaw) where each step may not be fully controlled or audited. The tension is autonomy vs safety: you want an agent that installs packages and runs commands without babysitting, but don't want it wiping a directory or leaking an API key because it was prompt-injected by something read on the web. Sandboxing draws a hard boundary: the agent gets a contained environment, and anything outside either fails or requires an explicit rule.

**How isolation works.** Docker Sandboxes are NOT standard Docker containers. Containers package applications but share more of the host kernel and aren't a strict security boundary for autonomous processes. A sandbox is a **micro VM**: a small, fast virtual machine with its own separate Linux kernel. Isolation is enforced by a **hypervisor** — a layer between physical hardware and the VM that doles out CPU/memory slices — making it hardware-enforced rather than a software convention, so it's harder for anything inside to reach outside. Full VMs offer similar isolation but are heavy and slow to start/stop. Micro VMs aim for the middle ground: real VM-level separation with boot/teardown times fast enough to spin up and discard for a single agent task. Inside the VM, a Docker engine lets the agent install packages, run scripts, and make a mess; when the session ends, the sandbox is deleted and nothing persists on the host.

**File and folder permissions.** Sandboxes are managed through a command-line tool. Running a coding agent inside one is as simple as pointing the command at a working directory (e.g., running Codex inside a specific project folder). The agent can read/write freely within that folder, but writes anywhere else are silently blocked. In testing, an agent instructed to create a file one directory above its sandboxed folder appeared to complete the task from its own perspective, but the file never appeared outside the permitted directory. Scoping isn't all-or-nothing: directories can be configured read-only (reference documents, config files) while a separate scratch folder gets full write access (notes/memory space).

**Network policies.** Three modes: **open** (no restrictions), **closed** (all network traffic blocked), **balanced** (practical default — pre-approves a broad list of hosts needed for common AI tools and cloud services: OpenAI, Anthropic, AWS, Google API endpoints; blocks everything else by default). Rules are inspectable and editable per sandbox. A fresh shell sandbox rejects outbound requests to arbitrary domains until the host is explicitly added. This lets users run an untrusted/experimental agent with network locked to only the specific APIs it needs rather than trusting it with the open internet.

**Credential proxying.** Instead of dropping a raw API key into the sandbox environment where a compromised or prompt-injected agent could read and exfiltrate it, the secret is registered with the sandbox tooling. Inside the sandbox, the agent only sees a placeholder marked proxy-managed. The real key is substituted only when the request actually leaves the sandbox toward an approved destination. An agent can be configured to use a service like OpenRouter without direct access to the underlying API key. Combined with network policy restrictions limiting outbound traffic to that service, it becomes much harder for a malicious/compromised agent to redirect credentials. Secrets can be set per sandbox or globally.

**Not just for coding agents.** Sandboxes also support a plain **shell mode** with no agent framework attached — for people building their own agents who want an isolated environment to install dependencies, run scripts, and test behavior without granting broad filesystem/network access. Shell mode is a practical way to evaluate new agents from GitHub before trusting them on a real system: run an unfamiliar agent for an extended period (connected to a local model rather than a paid API), then check what it did rather than worrying in real time. For teams wanting standardized configurations, an early-access system called **kits** bundles preset tools, environment variables, credentials, and network rules so a given sandbox setup can be reproduced consistently.

**FAQ highlights.** Regular containers share more host kernel; sandboxes are micro VMs with hardware-level hypervisor isolation. An agent can make authorized changes within its own working directory but writes outside are blocked and network is policy-restricted. Real API keys never go into the sandbox (proxy substitution on exit to approved destinations). Setup and individual use of the tooling is free.

## Key points

- Docker Sandboxes are micro VMs — not standard containers — with hardware-enforced hypervisor isolation and their own Linux kernel.
- File access is scoped by folder: writes outside the working directory are silently blocked even if the agent believes it succeeded.
- Three network policy modes: open, closed, and balanced (default whitelists OpenAI, Anthropic, AWS, Google API hosts).
- Credential proxying keeps secrets out of the sandbox: the agent sees only a placeholder; the real key is substituted on approved outbound requests.
- Micro VMs combine VM-level separation with fast boot/teardown, suitable for single-agent tasks.
- A shell-only sandbox mode extends protection to custom agents, frameworks, or scripts from GitHub.
- Kits (early access) package reusable sandbox templates with preset tools, env vars, credentials, and network rules.
- The command-line tool (sbx) handles login, creation, listing, and policy management; setup is free.

## Technical data / figures

- Isolation: micro VM with separate Linux kernel; hypervisor-enforced (hardware-level) separation vs container kernel sharing.
- File permissions: per-folder scoping; read-only vs scratch (read/write) directories.
- Network policies: open / closed / balanced (default whitelist: OpenAI, Anthropic, AWS, Google API endpoints; editable per sandbox).
- Credential proxying: placeholder in sandbox; real key substituted at exit to approved destinations; per-sandbox or global secrets.
- Use cases: coding agents (Claude Code, Codex), custom frameworks (Hermes, OpenClaw), shell mode for testing untrusted code.
- Tools: command-line tool `sbx`; kits for reusable templates; free to set up for individual use.

## Why this source matters for the RAG

Provides current (August 2026) technical detail on Docker Sandboxes for safely running AI coding agents — micro-VM architecture, file/network policy modes, credential proxying, and practical usage. This is relevant to agent security and prompt-injection defense, giving the RAG accurate up-to-date knowledge on agent sandboxing and isolation techniques.
