---
id: etape4-trackb-local-inference/00-local-inference/part-1-lm-studio-element-labs
title: "PART 1 — LM STUDIO (Element Labs)"
domain: step-4-track-b-local-inference-stack-llama-cpp-ollama-lm-stu
role: deep-dive
task: reference
actors: ["AMD", "Alibaba", "Anthropic", "Apple", "Google", "Hugging Face", "Intel", "Meta", "Mistral", "Moonshot", "Nvidia", "OpenAI", "United States", "Z.ai"]
dates: ["2023-05", "2025-07", "2026-06-04", "2026-07", "2026-07-16", "2026-08-28", "2026-09-09", "2026-09-19"]
keywords: ["acquisition", "agent", "agentic", "amd", "embeddings", "gguf", "glm", "gpu", "inference", "inference engine", "intel", "kimi"]
source: docs/RAG/etape4_trackB_local_inference.md
source_anchor: ""
source_lines: [688, 745]
section: "Step 4 — Track B: Local Inference Stack (llama.cpp + Ollama + LM Studio)"
sha256: bf170034425c605338bc786d6d74f8a1c1ee7edcc258acec0f5d677eca090ce2
---

# PART 1 — LM STUDIO (Element Labs)

## PART 1 — LM STUDIO (Element Labs)

### 1.1 Latest version & release history (2026)

- **Latest release: LM Studio 0.4.24 — September 9, 2026 (Build 1)** [official] (https://lmstudio.ai/changelog/lmstudio/lmstudio-v0.4.24)
  - Release notes: new LM Studio Engine Protocol with advanced llama.cpp argument overrides for GGUF model loading; improved DSpark and DFlash assistant-drafter compatibility heuristics; fixed Loaded Instances UI context-length display; fixed `/api/v1/chat` splitting text and image inputs into separate user messages; improved sign-in reliability. [official]
- Previous: **0.4.22 — August 28, 2026** [official] (https://lmstudio.ai/changelog/lmstudio/lmstudio-v0.4.22)
  - Support for DFlash, DSpark, and MTP assistant drafters (speculative decoding; requires llama.cpp engine ≥ 2.29.1); images returned by tools in OpenAI-compatible `/v1/responses`, `/v1/chat/completions`, and Anthropic-compatible `/v1/messages`; Linux fixes (AppImage without FUSE 2, Chromium sandbox, `lms` wake-up, .deb icons, in-app updates, Wayland keyboard shortcuts). [official]
- **0.4.11** added support for updated Gemma 4 chat template; **0.4.6 (Feb 27, 2026)** introduced LM Link (remote instance access via end-to-end encryption in partnership with Tailscale) and Qwen3.5 RAG jinja rendering fix. [official] (https://lmstudio.ai/changelog/lmstudio/lmstudio-v0.4.6)
- **0.4.16 — June 4, 2026**: shipped with "Locally", a native iPhone/iPad app paired with LM Link — a remote-access layer letting a phone use models running on the user's Mac over an end-to-end-encrypted Tailscale mesh; chat history stays on-device; only a device discovery list touches LM Studio servers; iPhone/iPad only at launch; free during preview. [secondary] (https://www.digitalapplied.com/blog/lm-studio-locally-lm-link-iphone-local-llm-2026)
- **0.4.0 (Jan 28, 2026)** — major re-architecture [official]: (https://lmstudio.ai/blog/0.4.0)
  - GUI separated from core: the core ("llmster" daemon) can run standalone on Linux boxes, cloud servers, GPU rigs, even Google Colab — i.e., a server-native, headless deployment of LM Studio.
  - Parallel requests with continuous batching (high-throughput serving instead of queueing).
  - New stateful native REST API (`/api/v1/*`), including MCP via API, stateful chats, API-token auth, model download/load/unload endpoints.
  - Refreshed UI: chat export, split view, developer mode, in-app docs.
  - Headless install: `curl -fsSL https://lmstudio.ai/install.sh | bash` / Windows `irm https://lmstudio.ai/install.ps1 | iex`; commands `lms daemon up`, `lms get <model>`, `lms server start`, `lms chat`, `lms runtime update llama.cpp` (and `lms runtime update mlx` on macOS). [official]
  - Self-hosted server mode via `lms server` — positioned against Ollama as the server option. [secondary] (http://www.howtogeek.com/you-can-make-a-self-hosted-ai-server-with-lm-studio-040/)

### 1.2 LM Studio Bionic (AI agent app)

- **LM Studio Bionic** — separate standalone desktop app (Mac, Windows, later Linux x64/ARM64 in 1.1.2), announced July 16, 2026 on the company blog [official/secondary]: (https://www.webpronews.com/lm-studios-bionic-agent-puts-open-models-to-work-on-your-desktop/), (https://9to5mac.com/2026/07/16/lm-studio-expands-beyond-chat-with-bionic-a-new-ai-agent-app-for-open-models/)
- An agentic app that works on code projects, documents, research with file access; uses local models or **LM Studio Secure Cloud** for heavier tasks (Zero Data Retention policy; US-based servers); cloud use requires an LM Studio account with billing configured. [secondary]
- **Bionic 1.1.5 — September 19, 2026** (latest at cutoff) [official]: (https://lmstudio.ai/changelog)
  - "[Mac only] Ultra fast Qwen3.8 inference using a new engine: Splash (by Inco AI)" — a new inference engine for Apple Silicon from Inco AI.
  - Prior versions: 1.1.4 (Sep 17, 2026) — agent introspection, llama.cpp 2.41.0 extension packs, tool-call argument streaming; 1.1.3 (Sep 15, 2026) — Linux voice transcription, MTP speculative decoding for more models, per-layer GGUF expert counts, MCP organization management; 1.1.0 (Aug 27, 2026) — "2–2.75× faster prompt processing on M5 Macs", Muse Glimmer tool calling for MLX models; 1.1.1/1.1.2 (Aug 31 / Sep 8, 2026) — Linux builds, `lms load --auto`, llama.cpp 2.31.x packs. [official]
- Bionic cloud models: GLM 5.2, Kimi K2.7 Code (code projects), Kimi K3, Z.ai GLM-5.3-Flash (320B MoE, 18B active, 1M context; served from US servers with ZDR; reported 9–10× cheaper than GLM-5.2) [secondary]: (https://9to5mac.com/2026/08/26/lm-studio-adds-glm-5-3-flash-to-bionic-with-image-support-and-1m-token-context/)
- Offline voice: ships with Mistral Voxtral model for multilingual offline transcription; a "voice keyboard" works across apps. [secondary]

### 1.3 SDKs & developer surfaces

- **lmstudio-js** (TypeScript): `npm install @lmstudio/sdk`; `import { LMStudioClient } from "@lmstudio/sdk"`. Source: https://github.com/lmstudio-ai/lmstudio-js. The app itself uses the same public APIs. Works from Node and (with CORS) the browser. [official] (https://lmstudio.ai/blog/introducing-lmstudio-sdk)
- **lmstudio-python**: `pip install lmstudio`; `import lmstudio as lms`. Sync convenience API + session-based async API. [official]
- Core SDK APIs: `.respond()` chat, `.act()` agentic tool use, structured output (Pydantic/zod/JSON schema), image input, speculative decoding (MLX and llama.cpp), `.complete()`, `.embed()`, model load/unload in memory, low-level config (GPU, context length). Auto engine selection (llama.cpp vs MLX) with per-hardware parameters chosen automatically. [official]
- REST: OpenAI-compatible `/v1/chat/completions`, `/v1/responses` (with `reasoning.effort` for `openai/gpt-oss-20b`), `/v1/embeddings`; Anthropic-compatible `/v1/messages` (added in 0.4.1); native stateful API `/api/v1/chat` with MCP support and auth tokens (added in 0.4.0). API changelog: [official] (https://lmstudio.ai/docs/developer/api-changelog)
- CLI `lms`: `lms load --estimate-only <model>` (GPU + total memory estimates, honoring `--context-length`/`--gpu`), `lms chat` (Ctrl+C interrupt), `lms ps --json`, `lms ls --variants`. [official]

### 1.4 Model catalog (notable additions)

- In-app model search wired to Hugging Face with memory-requirement estimates before download; supports GGUF (llama.cpp) and MLX formats. [secondary] (https://medium.com/@nishilbhave/lm-studio-in-2026-download-models-run-local-llms-vs-ollama-b30567df17f8)
- 2026 model landmarks referenced in official materials: `openai/gpt-oss-20b` (reasoning effort support in `/v1/responses`), Gemma 4 (0.4.11 chat-template support), Qwen3.5 (RAG fix in 0.4.6), Qwen3.8 (Splash engine in Bionic 1.1.5), NVIDIA Nemotron-Nano-v2 (0.3.25). [official]
- Company history context: launched May 2023 by Element Labs; old commercial license scrapped — free at home and at work since ~July 2025. [secondary] (https://www.aixploria.com/en/lm-studio-ai/), (https://medium.com/@nishilbhave/lm-studio-in-2026-download-models-run-local-llms-vs-ollama-b30567df17f8)
- One secondary source reports "LM Studio crossed millions of downloads worldwide" as of 2026 [secondary — no vendor-published user count found at cutoff; flag as unverified for exact figures]: (https://medium.com/@nishilbhave/lm-studio-in-2026-download-models-run-local-llms-vs-ollama-b30567df17f8)

### 1.5 Hardware acceleration

- Inference engines: **llama.cpp** (primary, CPU/GPU, all platforms) and **Apple MLX** (Apple Silicon M1–M5); both are independently downloadable runtimes — engine updates ship separately from app updates (`lms runtime update`). [secondary cross-checked with official 0.4.0 install docs] (https://github.com/aterrylu/autonomos/blob/HEAD/docs/research/desktop-shells/lm-studio.md)
- Backend coverage on desktop: CUDA (NVIDIA), Metal (Apple), Vulkan (cross-vendor, incl. AMD/Intel); changelog entries also reference ROCm builds and multi-GPU selection bugs on CUDA 12 / ROCm / Vulkan. [official] (https://lmstudio.ai/changelog/lmstudio/lmstudio-v0.4.22)
- NPU: LM Studio runs on the user's GPU or NPU per third-party coverage; NPU-specific offload reported in community NPU laptop tests (see §2.4). Flag: exact NPU backends in LM Studio are not enumerated in official docs found at cutoff [secondary/unverified]. [secondary] (http://www.howtogeek.com/you-can-make-a-self-hosted-ai-server-with-lm-studio-040/)

### 1.6 Pricing

- **Desktop app and local use: free** (home and work). [secondary — widely reported, incl. LM Studio comparison coverage] (https://www.aixploria.com/en/lm-studio-ai/), (https://www.kunalganglani.com/blog/lm-studio-vs-ollama)
- One comparison site reports an **Enterprise tier** for teams needing LM Link multi-device workload routing and priority support — but this is not confirmed on LM Studio's public pricing pages at cutoff; treat as [unverified]. (https://www.kunalganglani.com/blog/lm-studio-vs-ollama)
- **LM Studio Secure Cloud**: pay-as-you-go via Cloud Credits; exact pricing "not published" as of July 2026 per a Bionic developer guide; "Bionic Pass" subscription announced as "coming soon". LM Link free tier covers up to 5 devices. [secondary] (https://learning.christiandrapatz.de/lmstudio-en.pdf)
- LM Studio acquisition of Locally AI (mobile, Adrien Grondin) — strategic push to iPhone/iPad on-device AI; iOS local models realistically 1B–3B at practical speeds. [secondary] (https://evermx.com/case/lm-studio-acquires-locally-ai-mobile-on-device-ai)

---

