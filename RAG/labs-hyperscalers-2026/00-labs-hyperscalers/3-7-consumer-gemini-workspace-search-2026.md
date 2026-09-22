---
id: labs-hyperscalers-2026/00-labs-hyperscalers/3-7-consumer-gemini-workspace-search-2026
title: "3.7 Consumer Gemini / Workspace / Search (2026)"
domain: step-3-labs-hyperscalers-february-1-september-22-2026
role: deep-dive
task: reference
actors: ["AWS", "Anthropic", "Broadcom", "Google", "Microsoft", "Nvidia", "TSMC", "United States"]
dates: ["2025-05", "2026-03-31", "2026-05", "2026-05-12", "2026-05-19", "2026-06", "2026-06-18", "2026-07", "2026-07-30"]
keywords: ["consumer", "gemini", "acquisition", "agent", "agentic", "agents", "backlog", "benchmark", "benchmarks", "capex", "claude", "compute"]
source: docs/RAG/Labos  hyperscalersEN.md
source_anchor: ""
source_lines: [597, 647]
section: "Step 3 — Labs & Hyperscalers (February 1 → September 22, 2026)"
sha256: 544a657767fd3a53ab27f71a77d4845f99b75d1de6bad4d94c4cb4aa21c42e01
---

# 3.7 Consumer Gemini / Workspace / Search (2026)

### 3.7 Consumer Gemini / Workspace / Search (2026)

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

- **Project Genie (Genie 3 world model) public preview — Jan 29, 2026**: create/explore/remix interactive 3D worlds via text/image; Genie 3 = 11B-parameter autoregressive transformer, real-time navigable worlds at 720p/24fps; limited to Google AI Ultra subscribers, US, 18+ [secondary].
- **AlphaFold Protein Database — 3M researchers milestone (~Feb 17, 2026)**: >3M researchers across 190 countries; >240M structural predictions; >1/3 of users in low/middle-income nations [independent] https://dataconomy.com/2026/02/17/deepminds-alphafold-database-hits-3-million-researchers-milestone/. **AlphaFold 4**: SEO-style articles claim a 2026 "AlphaFold 4" — **no credible announcement found; NOT reported as fact** [unverified].
- **Gemini Robotics-ER 1.6 — Apr 14/15, 2026**: embodied-reasoning model; improved spatial/physical reasoning, pointing, counting, success detection, gauge/instrument reading; validated with Boston Dynamics Spot; available via Gemini API and AI Studio with published safety evaluation [secondary].
- **Gemini Robotics 2 suite — July 30, 2026**: three-model suite — Whole-Body VLA (feet-to-fingertips humanoid control), ER 2 (multi-step planning, multi-robot collaboration), On-Device 2 (adapts to new robot bodies "within hours"); ER 2: 91.3% moment-finding accuracy, 57.4% progress classification; demonstrated on Boston Dynamics Spot, Apptronik Apollo 2 humanoid, Franka arms; ASIMOV-Agentic safety benchmark published [secondary].
- **DeepMind × Fenris Creations (EVE Online) — Aug 21, 2026**: partnership to test SIMA 2 generalist agents in persistent game worlds for continual learning, long-term planning, memory, multi-agent interaction [secondary].
- **Decoupled DiLoCo — May 2026**: distributed-training architecture; 198 Gbps → 0.84 Gbps inter-DC bandwidth across 8 datacenters; 88% goodput vs 27% baseline [secondary — single weak source].
- **Co-Scientist cellular-aging-reversal — May 19, 2026**: DeepMind blog reporting novel rejuvenation factors in human cells [secondary — no peer-reviewed companion captured].

### 3.11 Alphabet × Anthropic: the defining deal of the window

- **Apr 6, 2026**: Anthropic announced agreement with Google and Broadcom for **multiple gigawatts of next-generation TPU capacity** online starting 2027 — Anthropic's "most significant compute commitment to date"; no dollar figure disclosed [secondary].
- **Apr 24, 2026** (during Cloud Next): Bloomberg/CNBC reported Alphabet will invest **up to $40B** in Anthropic — **$10B immediate cash at ~$350B valuation, +$30B contingent on performance milestones** [independent]. No joint official announcement was published; terms confirmed to reporters by sources. Valuation reported as $350B (Bloomberg) vs $380B (some outlets) — flagged. Prior Google stake: >$3B across tranches since 2023, ~14% ownership. Full exercise would take Google's invested capital to ~$43B [secondary].
- Compute dimension: separate reports describe **up to 1M TPUs / 5 GW Ironwood capacity over 5 years from 2027**, "well over a gigawatt coming online in 2026" [secondary — contractual caps/analyst estimates, not provisioned capacity].
- Parallel track: Amazon committed up to $33B to Anthropic days earlier [secondary].

**Isomorphic Labs — $2.1B Series B (May 12, 2026)**: Alphabet/DeepMind drug-discovery spinout; led by Thrive Capital; participants Alphabet, GV, MGX, Temasek, CapitalG, UK Sovereign AI Fund; largest AI-drug-discovery raise to date [independent — Reuters]. Funds to scale IsoDDE (claimed >2× AlphaFold 3 accuracy on protein-ligand benchmarks); first clinical trials now expected **end of 2026** (delayed from end 2025) [independent]. Demis Hassabis remains founder/CEO [independent].

**Alphabet capital raising (June 2026)**: filed to raise up to $80B in equity — upsized to **$84.75B** (largest US corporate equity raise; prior record $70B Petrobras 2010); $30B underwritten public offering + $40B at-the-market + $10B Berkshire Hathaway private placement (14,212,035 Class A @ ~$351.81; 14,359,656 Class C @ ~$348.20) [secondary]. Proceeds for AI compute infra; first Alphabet equity raise in 20+ years; debt topped $100B (from $28B Mar 2025) [secondary — Barron's].

**Alphabet capex trajectory**: 2025 actual ~$91B [secondary]. Q1 2026 (reported Apr 29): raised 2026 guidance to **$175–185B** (some outlets $180–190B — conflict). Q2 2026 (reported Jul 22): raised to **$195–205B**; Q2 capex $44.9B (~60% servers, ~40% DC/networking); first-ever negative free cash flow quarter; further increases signaled for 2027 [secondary].

**Alphabet earnings (AI-relevant)**: Q1 2026 (Apr 29): revenue $109.9B; Google Cloud **$20.02B (+63%)**; backlog $462B; Gemini API **16B tokens/min** direct customer use; 330 Cloud customers >1T tokens TTM; gen-AI product revenue +800% YoY [secondary]. Q2 2026 (Jul 22): revenue $119.8B (+24%); Cloud **$24.8B (+82%)**; Cloud backlog +$50B in quarter to $514B; demand outpacing supply per Pichai [secondary].

**Energy / data-center**: Intersect acquisition (announced Dec 22, 2025 — context): **$4.75B cash** + debt assumption, close expected H1 2026; multiple GW of energy/data-center projects incl. Haskell County, TX (640 MW solar + 1.3 GW storage, operational 2027) [independent]. Other moves: NV Energy 115 MW geothermal; Energy Dome CO2 batteries; Broadwing Energy CCS gas plant [secondary]. Alphabet AI infra footprint: 30+ data centers, 40+ cloud regions, 10M km fiber [vendor-reported].

