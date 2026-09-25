---
id: collect-250926-servers-hardware/servers-hardware/muse-glimmer-from-meta-superintelligence-labs-is-now-available-ollama-blog
title: "Muse Glimmer from Meta Superintelligence Labs is now available"
domain: servers-hardware
role: reference
task: reference
actors: ["Anthropic", "Apple", "Meta", "Microsoft"]
dates: ["2026-08-10"]
keywords: ["muse", "agent", "agentic", "agents", "apache", "claude", "copilot", "latency", "license", "multimodal", "reasoning", "tool calling"]
source: docs/RAG/clean4/muse-glimmer-from-meta-superintelligence-labs-is-now-available-ollama-blog.md
source_anchor: ""
source_lines: [1, 55]
sha256: d4515baa4c214ea635d7f572a28b12ed55d4584b4ba10f3168642154a5321cdc
---

# Muse Glimmer from Meta Superintelligence Labs is now available

## August 10, 2026

**Muse Glimmer**, Meta’s newest open model and the first released by Meta Superintelligence Labs, is now available on Ollama. It’s a 30B multimodal model purpose-built for agent workloads that run locally with a 128K+ context length, released under the Apache 2.0 license.

With Ollama, you can now use Muse Glimmer to power coding agent applications such as Claude Code, Codex, Pi and more, as well as long-running personal assistants such as OpenClaw and Hermes.

On Apple Silicon, it is available via Ollama’s MLX engine, which provides state-of-the-art performance with new DFlash and image input support.

To run Muse Glimmer, download the latest release of Ollama and run:

```
ollama run muse-glimmer
```
For state of the art performance using Ollama’s MLX engine on Apple Silicon, use:

```
ollama run muse-glimmer:30b-mlx
```
## Power coding agents locally

To run Muse Glimmer with Claude Code, download Ollama and run:

```
ollama launch claude --model muse-glimmer
```
For a lighter-weight coding agent, try Pi:

```
ollama launch pi --model muse-glimmer
```
For personal assistant frameworks such as OpenClaw and Hermes, use:

```
ollama launch openclaw --model muse-glimmer
```
```
ollama launch hermes --model muse-glimmer
```
`ollama launch` also works with Codex, OpenCode, GitHub Copilot, and more. See all integrations.

Muse Glimmer also supports controllable reasoning strength: `low`, `medium`, `high`, and `xhigh`. Use `high` or `xhigh` for complex coding and agentic tasks, and lower strengths when speed matters more.

## State-of-the-art performance on Apple Silicon with new DFlash and image input support

Ollama’s MLX engine now supports DFlash, building on previous multi-token prediction (MTP) support. With DFlash, Muse Glimmer runs 1.5×–1.8× faster on Apple Silicon.

To run Muse Glimmer using Ollama’s MLX engine, use the `muse-glimmer:30b-mlx` model.

Muse Glimmer’s dedicated 1.8B-parameter perception encoder gives it native image understanding. With new image input support in the MLX engine, coding agents can use it for low-latency, back-to-back tool calling on tasks that require images, such as:

- Building websites or applications from a drawing or mockup
- Computer use applications powered by screenshots
- Reading documents, receipts, and charts
