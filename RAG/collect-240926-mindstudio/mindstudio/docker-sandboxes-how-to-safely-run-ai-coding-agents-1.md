---
id: collect-240926-mindstudio/mindstudio/docker-sandboxes-how-to-safely-run-ai-coding-agents-1
title: "docker-sandboxes-how-to-safely-run-ai-coding-agents"
domain: mindstudio
role: reference
task: reference
actors: ["AWS", "Anthropic", "Google", "OpenAI", "OpenRouter"]
dates: []
keywords: ["agent", "agents", "sandbox", "aws", "claude", "memory", "packaging"]
source: docs/RAG/clean_en/mindstudio/docker-sandboxes-how-to-safely-run-ai-coding-agents.md
source_anchor: ""
source_lines: [1, 54]
sha256: 7f20b3a7e9586ae4744f158eecadeb4b2ec0843d62ea4019d664932bde9b1d30
---

# docker-sandboxes-how-to-safely-run-ai-coding-agents

<!-- source: https://www.mindstudio.ai/blog/docker-sandboxes-ai-agents -->

## What are Docker Sandboxes?

Docker Sandboxes are isolated micro VMs built to run AI coding agents like Claude Code, Codex, or custom agent frameworks without giving them free rein over your machine. Each sandbox boots its own Linux kernel, runs a Docker engine inside it, and lets you set policies for network access, file read/write permissions, and credential handling. When you’re done, you throw the whole sandbox away and your host system is untouched.

## TL;DR

- **Micro VMs** give hardware-enforced isolation through a hypervisor, sitting between the speed of a container and the heaviness of a full virtual machine.
- **File access is scoped by folder** , so an agent running in a sandbox can write inside its working directory but gets silently blocked from touching anything outside it, even if it thinks the write succeeded.
- **Network policies come in three modes** (open, closed, balanced), and the balanced default already whitelists a wide set of hosts for services like OpenAI, Anthropic, AWS, and Google APIs.
- **Credential proxying keeps secrets out of the sandbox entirely** : you set a secret like an API key, the agent only ever sees a placeholder, and the real key gets substituted only when the request leaves the sandbox toward an approved destination.
- **Custom agents get the same treatment as coding agents** through a shell-only sandbox mode, useful for testing agent frameworks or scripts pulled from GitHub before trusting them on your real filesystem.
- **Kits are an early-access system** for packaging your own sandbox templates with preset tools, environment variables, credentials, and network rules.
- **The command-line tool (sbx)** handles login, creation, listing, and policy management, and the setup itself is free to use.

## Why does running AI agents need a sandbox in the first place?

Coding agents that operate with elevated permissions are efficient but risky. Tools like Claude Code or Codex offer a “skip permissions” mode that stops the constant approve/deny prompts, but that convenience means the agent can, in principle, touch anything on your system: delete files, modify configs, or make network calls you didn’t authorize. The risk compounds with more autonomous or custom agents, including frameworks like Hermes or OpenClaw, where you may not fully control or audit what the agent does at each step.

The tension is between autonomy and safety. You want an agent that can install packages, run commands, and make real progress on a task without you babysitting every action. But you also don’t want to discover it wiped a directory or leaked an API key because it got prompt-injected by something it read on the web. Sandboxing solves this by drawing a hard boundary: the agent gets a contained environment to work in, and anything it tries to do outside that boundary either fails or requires an explicit rule you set.

## How does the isolation actually work?

Docker Sandboxes are not the same thing as a standard Docker container. Containers are good at packaging applications, but they share more of the host’s kernel and aren’t designed as a strict security boundary for autonomous processes. A sandbox, by contrast, is a micro VM: a small, fast virtual machine with its own separate Linux kernel.

The isolation is enforced by a hypervisor, a layer that sits between your physical hardware and the VM and doles out a slice of CPU and memory to it. Because this separation is hardware enforced rather than purely a software convention, it’s harder for anything running inside the sandbox to reach outside of it. Full VMs offer similar isolation but are heavy and slow to start and stop. Micro VMs aim for the middle ground: real VM-level separation with boot and teardown times fast enough to spin up and discard for a single agent task.

Inside the VM, there’s a Docker engine where the agent can install packages, run scripts, and generally make a mess. When the session ends, you delete the sandbox and nothing that happened inside it persists or affects your host machine.

## How do you set file and folder permissions for an agent?

Once installed, sandboxes are managed through a command-line tool. Running a coding agent inside one is as simple as pointing the command at a working directory, for example running Codex inside a specific project folder. From that point, the agent can read and write freely within that folder, but if it attempts to write anywhere else on the filesystem, the write is silently blocked. In testing, an agent instructed to create a file one directory above its sandboxed folder appeared to complete the task from its own perspective, but the file never actually appeared outside the permitted directory.

This scoping isn’t limited to a single all-or-nothing folder either. You can configure some directories as read-only, useful for reference documents or configuration files you want the agent to consult but never modify, while giving it full write access to a separate scratch folder it can use as a notes or memory space.

## How do network policies control what an agent can reach?

Network access follows the same logic. When you first configure sandboxes, you choose one of three network policies: open (no restrictions), closed (all network traffic blocked), or balanced. Balanced is the practical default for most use cases: it pre-approves a broad list of hosts needed for common AI tools and cloud services, covering things like OpenAI, Anthropic, AWS, and Google’s API endpoints, while blocking everything else by default.

You can inspect and edit these rules per sandbox. A freshly created shell sandbox with no coding agent attached, for instance, will reject an outbound request to an arbitrary domain until you explicitly add that host to the policy. Once added, the same request goes through. This makes it straightforward to run an untrusted or experimental agent with network access locked down to only the specific APIs it actually needs, rather than trusting it with the open internet.

## How does credential proxying keep API keys safe?

One of the more useful protections is how sandboxes handle secrets. Instead of dropping a raw API key into the sandbox’s environment where a compromised or prompt-injected agent could read and exfiltrate it, you register the secret with the sandbox tooling. Inside the sandbox, the agent only ever sees a placeholder value marked as proxy-managed. The real key is substituted only at the point the request actually leaves the sandbox, addressed to an approved destination.

This means an agent can be configured to use a service like OpenRouter without ever having direct access to the underlying API key. Combined with network policy restrictions that limit outbound traffic to only that service, it becomes much harder for a malicious or compromised agent to redirect credentials somewhere unintended. Secrets can be set per sandbox or globally if you’re reusing the same key across multiple sandboxes.

## Is this only useful for coding agents?

No. While coding agents like Claude Code and Codex are the most immediate use case, sandboxes also support a plain shell mode with no agent framework attached. This is aimed at people building their own agents, whether with agent frameworks or fully custom code, who want an isolated environment to install dependencies, run scripts, and test behavior without granting broad filesystem or network access on the host machine.

This shell mode is also a practical way to evaluate new agents you find on GitHub or elsewhere before trusting them on your real system. You can let an unfamiliar agent run for an extended period, connected to a local model rather than a paid API, and check afterward what it actually did rather than worrying in real time about damage to your file system or credentials.

