---
id: labs-hyperscalers-2026/00-labs-hyperscalers/3-8-tpus-ironwood-7th-gen-ga-tpu-8t-8i-8th-gen
title: "3.8 TPUs: Ironwood (7th gen) GA; TPU 8t/8i (8th gen)"
domain: step-3-labs-hyperscalers-february-1-september-22-2026
role: deep-dive
task: reference
actors: ["Anthropic", "Google", "Microsoft", "Nvidia", "TSMC", "United States"]
dates: ["2025-05", "2026-03-31", "2026-05-19", "2026-06-18", "2026-07"]
keywords: ["tpu", "agent", "agentic", "agents", "claude", "compute", "diffusion", "fp4", "gemini", "hbm", "inference", "mcp"]
source: docs/RAG/Labos  hyperscalersEN.md
source_anchor: ""
source_lines: [599, 622]
section: "Step 3 — Labs & Hyperscalers (February 1 → September 22, 2026)"
sha256: f38a838480e6d16da7bc621e0f3143edf4fd79b75e1aafd12bc043a43387e867
---

# 3.8 TPUs: Ironwood (7th gen) GA; TPU 8t/8i (8th gen)

- **Gemini Spark (I/O, May 19, 2026)**: "24/7 personal AI agent" running on Google Cloud VMs (works after device closed), built on Gemini 3.5 Flash + Antigravity harness. Workspace integrations (Gmail/Docs/Slides) at launch, MCP third-party expansion over summer 2026 (Canva, OpenTable, Instacart confirmed). Available to Google AI Ultra subscribers in the US within a week of I/O; macOS desktop integration planned summer 2026. [secondary] https://github.com/aicoachellavalley/homepage/blob/HEAD/src/content/briefs/2026-05-19-google-gemini-spark.mdx
- **Google AI Ultra pricing**: reported at **$100/month** at I/O 2026 [secondary]; an older (Nov 2025) source lists $250/month — the $100 figure is presented as a price cut, but timing/terms unverified.
- **AI Mode in Search**: passed **1 billion MAU** globally by May 19, 2026; upgraded to Gemini 3.5 Flash at I/O; "information agents" for web monitoring rolling out summer 2026 for Pro/Ultra subscribers. [secondary] https://ppc.land/inside-google-i-o-2026-the-agentic-ai-shift-no-one-saw-coming/
- **Gemini app**: passed **950M MAU** (per Pichai note reported by Axios, ~Aug 2026) [secondary]. **Gemma open models: >900M downloads** [secondary].
- **Workspace AI subscriptions passed 11M paid seats, +76% YoY** (reported ~July 2026) [secondary]. **Gemini Enterprise pay-as-you-go** added ~Aug 2026 for select workloads (pooled daily allowances, consolidated billing) [secondary].
- **Veo 3 (video)**: image-to-video added to the Gemini app, now in **150+ countries** for AI Ultra/Pro subscribers, 3 videos/day cap; **>40M videos generated in the first seven weeks**; visible "Veo" + invisible SynthID watermarks [secondary]. (Veo 3 itself launched May 2025 — outside window; 2026 news is feature expansion.)
- **Agent Payments Protocol (AP2)** extension announced at I/O: open protocol for agent-led transactions; user approval required, merchant/amount limits, permanent transaction trail; integration into Google products "in the coming months," starting with Spark [secondary].
- Pichai stat (Cloud Next, Apr 2026): **~75% of all new Google code is now AI-generated and engineer-reviewed** [vendor-reported].

### 3.8 TPUs: Ironwood (7th gen) GA; TPU 8t/8i (8th gen)

**Ironwood / TPU7x — GA March 31, 2026:** "TPU7x is generally available (GA). TPU7x is the first release within the Ironwood family, Google Cloud's seventh generation TPU" for large-scale training/inference incl. LLMs/MoEs/diffusion models [official — Google Cloud TPU release notes] https://docs.cloud.google.com/tpu/docs/release-notes. Pods scale to 9,216 chips [independent — SiliconANGLE]. Pricing (press citing Google's TPU pricing page): on-demand **$12.00/chip-hour** in us-central1 (Iowa), $13.20 in europe-west2 (London); Flex-start $6.00/hr; 1-year commit $8.40/$9.24; 3-year $5.40/$5.94 [secondary — not verified against live pricing page]. SemiAnalysis (InferenceX): at 100 tokens/s/user interactivity, Ironwood ≈ **$0.181/M total tokens vs $0.222 (B200) and $0.276 (B300)** — ~19% cheaper than B200, ~34% cheaper than B300 per token at equal interactivity [independent].

**TPU 8t / 8i — announced Apr 22, 2026 (Cloud Next):** first split of the TPU line into workload-specialized chips: **TPU 8t** (training) and **TPU 8i** (inference/RL) [official — Google Cloud Next '26 recap] https://blog.google/innovation-and-ai/infrastructure-and-cloud/google-cloud/google-cloud-next-26-recap/. TPU 8t: 12.6 FP4 PFLOPS/chip, 216GB HBM3e, TSMC N3; 9,600-chip superpod = 121 FP4 exaflops with 2PB shared HBM; 2.7× better training price-performance vs Ironwood [secondary]. TPU 8i: 384MB on-chip SRAM (3×), 288GB HBM (+50%), new Collectives Acceleration Engine; 80% better performance-per-dollar [official via recap]. **Virgo Network**: new megascale DC fabric; up to 47 petabits/sec non-blocking bisection; links 9,600-TPU pods and up to one million TPUs across data centers [secondary]. Google "goodput" target ~97% [secondary]. Google Cloud will be among the first to offer **NVIDIA Vera Rubin NVL72** systems [official via recap].

**TPU competitive positioning (analysis):** SemiAnalysis InferenceX — Ironwood doesn't beat B200/B300 on raw performance across most of the curve, but lower TCO yields the per-token advantage [independent]. The Register: NVIDIA Rubin offers more per-chip compute and memory bandwidth than TPU 8t, but Google's edge is scale — optical circuit switches link 9,600 TPUs per pod (vs 576 accelerators per NVLink domain), with Virgo extending to multi-DC million-TPU clusters [independent]. Spheron: TPU Ironwood is GCP-captive (JAX stack, single-vendor lock-in); migrating off requires rewriting the serving stack [secondary]. SemiAnalysis notes Google has begun **selling** TPUs (not only renting via Cloud) [independent].

### 3.9 Gemini Enterprise Agent Platform (Cloud Next '26, Apr 22–25)

- Vertex AI rebranded/reorganized into the **Gemini Enterprise Agent Platform** — "the end-to-end system for the agentic era": low-code Agent Studio, code-first Agent Development Kit, agent registries, shared context (memory layer), runtime engines, Agent Identity, Agent Gateway, observability; **Model Garden with 200+ models** incl. Gemini 3.1 Pro, Gemini 3.1 Flash Image ("Nano Banana 2"), Lyria 3, Gemma 4, and third-party models incl. Anthropic's Claude [official/secondary]. **Agentic Defense**: Google Threat Intelligence + Security Operations + **Wiz** folded into an AI Application Protection Platform (Wiz $32B deal announced Mar 2025 — outside window; Next confirms absorption) [secondary]. **$750M partner innovation fund** committed [secondary]. **Agent2Agent (A2A) protocol** in production at 150 organizations, with native integrations in ServiceNow, Salesforce, Atlassian, SAP [secondary]. Gemini Enterprise app additions: Agent Designer (no-code), Agent Inbox, long-running background agents in secure sandboxes, Skills, Projects (persistent memory), Deep Think, Microsoft 365 interoperability [official via recap].
- **Google I/O 2026 developer platform (May 19):** **Antigravity 2.0** (standalone agent-orchestration desktop app + CLI + SDK); Managed Agents in the Gemini API; **AgentKit 2.0** (16 pre-built agents, A2A protocol); Gemini CLI sunsets for Free/Pro/Ultra tiers **June 18, 2026** — migration to Antigravity CLI [secondary]. **Model Armor**: built-in security layer scanning/blocking unsafe AI chats [secondary].

### 3.10 DeepMind research milestones (in-window)

