---
id: ai-industry-kb-2026-wave6/15-nvidia-and-nemotron/part-10
title: "§15. NVIDIA and Nemotron (part 10)"
domain: nvidia-and-nemotron
role: deep-dive
task: actor-profile
actors: ["Groq", "Hugging Face", "Nvidia"]
dates: ["2025-08-25", "2025-12-19", "2026-03", "2026-06-01", "2026-06-03", "2026-07", "2026-07-20", "2026-07-21", "2026-08-05", "2026-08-06", "2026-08-11", "2026-08-12", "2026-08-14", "2026-08-22"]
keywords: ["nvidia", "agent", "agents", "benchmarks", "disaggregated", "fp8", "gguf", "humanoid", "inference", "intel", "license", "moe"]
source: docs/RAG/ai-industry-knowledge-base-2026-wave6.md
source_anchor: ""
source_lines: [7740, 7779]
section: "§15. NVIDIA and Nemotron"
delta_of: ai-industry-kb-2026
sha256: 883e12c77bb35d67c5b37539c9b26ceaf2612224831d28d1e6bea736cbfac2cd
---

# §15. NVIDIA and Nemotron (part 10)

- [S1] MarkTechPost: NVIDIA releases Nemotron 3.5 Lightning (30B MoE, 3B active) and NeMo Switchyard (2026-08-11) — https://www.marktechpost.com/2026/08/11/nvidia-ai-releases-nemotron-3-5-lightning-and-nemo-switchyard/
- [S2] The Register: NVIDIA's NeMo Switchyard software router (2026-08-12) — https://www.theregister.com/ai-and-ml/2026/08/12/nvidias-latest-solution-for-soaring-enterprise-costs-nemo-switchyard-software-router/5286911
- [S3] GitHub: nvidia-nemo/switchyard (open-source router library) — https://github.com/nvidia-nemo/switchyard
- [S4] GitHub: nvidia-nemo/switchyard AGENTS.md (Rust/Python architecture) — https://github.com/nvidia-nemo/switchyard/blob/HEAD/AGENTS.md
- [S5] GitHub: nvidia-nemo/switchyard README.md (Terminal-Bench 2.1 chart, integrations) — https://github.com/nvidia-nemo/switchyard/blob/HEAD/README.md
- [S6] NVIDIA Nemotron LLM information page — https://www.nvidia.com/en-us/ai-data-science/foundation-models/nemotron/llm-info/
- [S7] Hugging Face: ggml-org/NVIDIA-Nemotron-3.5-Lightning-30B-A3B-GGUF — https://huggingface.co/ggml-org/NVIDIA-Nemotron-3.5-Lightning-30B-A3B-GGUF
- [S8] GitHub: anyaa-labs/agent-skills — Nemotron model profiles (Omni vs text-only, catalog models) — https://github.com/anyaa-labs/agent-skills/blob/HEAD/agent-architect/model-profiles/nvidia-nemotron.md
- [S9] NVIDIA Developer Blog: creating the Nemotron-3-Ultra-NVFP4 checkpoint with NVIDIA Model Optimizer — https://developer.nvidia.com/blog/creating-the-nvidia-nemotron-3-ultra-nvfp4-checkpoint-with-nvidia-model-optimizer/
- [S10] Hugging Face: nvidia/Nemotron-3-Embed-8B-BF16 — https://huggingface.co/nvidia/Nemotron-3-Embed-8B-BF16
- [S11] NVIDIA Technical Blog: develop physical AI reasoning, world and action models with Cosmos 3 — https://developer.nvidia.com/blog/develop-physical-ai-reasoning-world-and-action-models-with-nvidia-cosmos-3/
- [S12] GitHub: pedro-bright/the-ledger — NVIDIA Cosmos 3 release event record (2026-06-01) — https://github.com/pedro-bright/the-ledger/blob/HEAD/content/events/2026/57-nvidia-cosmos-3-release.md
- [S13] MarkTechPost: NVIDIA releases Cosmos 3 — two-tower MoT unifying physical reasoning, world and action generation (2026-06-03) — https://www.marktechpost.com/2026/06/03/nvidia-releases-cosmos-3-a-two-tower-mixture-of-transformers-foundation-model-unifying-physical-reasoning-world-generation-and-action-generation/
- [S14] MarkTechPost: NVIDIA releases Cosmos 3 Edge — 4B on-device world model (2026-07-21) — https://www.marktechpost.com/2026/07/21/nvidia-releases-cosmos-3-edge-a-4b-parameter-open-world-model-that-reasons-and-generates-robot-actions-on-device/
- [S15] TPS Report: NVIDIA Cosmos 3 Edge 4B world model for edge AI and robotics (2026-07-20) — https://tpsreport.news/news/nvidia-cosmos-3-edge-4b-world-model
- [S16] GitHub: raasaar-org/robot-management-system — RES-001 Cosmos 3 research note (sizes, modes, access, licensing) — https://github.com/raasaar-org/robot-management-system/blob/HEAD/.mc/research/RES-001-nvidia-cosmos-3-omnimodal-world-model-for-physical-ai/RES-001.md
- [S17] Robotics Business News: NVIDIA unites global robotics leaders — Cosmos, Isaac and GR00T (N1.7 early access, N2 preview) — https://roboticsbusinessnews.com/news/73/2618/nvidia-unites-global-robotics-leaders-to-scale-physical-ai-with-cosmos-isaac-and-gr00t-models.html
- [S18] NVIDIA/GlobeNewswire (PDF): NVIDIA announces Isaac GR00T Reference Humanoid Robot (2026-06-01) — https://pnr-files.pro1.gus.wdc.dianum.io/globenewswire/articles/3303990/en/nvidia-announces-nvidia-isaac-gr00t-reference.pdf
- [S19] Morningstar (Dow Jones Newswires): LG, Nvidia to unveil jointly developed humanoid robot next year (2026-08-14) — https://www.morningstar.com/news/dow-jones/20260814851/lg-nvidia-to-unveil-jointly-developed-humanoid-robot-next-year
- [S20] TechCrunch: NVIDIA launches Alpamayo — open AI models for autonomous vehicles (CES 2026) — https://techcrunch.com/2026/01/05/nvidia-launches-alpamayo-open-ai-models-that-allow-autonomous-vehicles-to-think-like-a-human/
- [S21] MarkTechPost: NVIDIA releases Alpamayo 2 Super — 34B open VLA under OpenMDW-1.1 (2026-08-05) — https://www.marktechpost.com/2026/08/05/nvidia-alpamayo-2-super-open-vla-model-autonomous-driving/
- [S22] AdwaitX: NVIDIA Alpamayo open AI model for self-driving cars (CES 2026 launch bundle) — https://www.adwaitx.com/nvidia-alpamayo-autonomous-vehicle-ai-model/
- [S23] Brief.news: NVIDIA unveils Alpamayo 2 Super (2026-06-01) — https://www.brief.news/ai/2026/06/01/nvidia-unveils-alpamayo-2
- [S24] Highways Today: NVIDIA opens Alpamayo 2 Super for commercial autonomous driving (2026-08-22) — https://highways.today/2026/08/22/nvidia-alpamayo-2-super/
- [S25] GitHub: smfworks/aiclearinghouse-site — NVIDIA NemotronLabs VoiceChat 11B deep dive (2026-08-06) — https://github.com/smfworks/aiclearinghouse-site/blob/HEAD/content/blog/2026-08-06-nvidia-nemotronlabs-voicechat-11b.md
- [S26] GitHub: liunix61/vllm-omni — NemotronLabs VoiceChat recipe (12.5 Hz frame lock, architecture) — https://github.com/liunix61/vllm-omni/blob/HEAD/recipes/NVIDIA/NemotronLabs-VoiceChat.md
- [S27] GitHub: sgl-project/sglang-omni — Nemotron VoiceChat cookbook (44 GB checkpoint, H200 footprint) — https://github.com/sgl-project/sglang-omni/blob/HEAD/docs/cookbook/nemotron_voicechat.md
- [S28] GitHub: vllm-project/vllm-omni PR #5842 — offline speech-to-speech support for VoiceChat 11B — https://github.com/vllm-project/vllm-omni/pull/5842
- [S29] GitHub: flyteorg/flyte-sdk — Nemotron-Omni voice chat example (Nemotron-3-Nano-Omni-30B-A3B, FP8 ~33 GB) — https://github.com/flyteorg/flyte-sdk/blob/HEAD/examples/genai/nemotron_omni_voice/README.md
- [S30] NVIDIA Developer Forums: 5M downloads on Hugging Face (2025-12-19) — https://forums.developer.nvidia.com/t/5m-downloads-on-huggingface/355228
- [S31] TokenCost: NVIDIA Nemotron 3 Super pricing & benchmarks (2026) — https://tokencost.app/blog/nvidia-nemotron-3-super-pricing
- [S32] AwesomeAgents: NVIDIA Nemotron 3 Super 120B-A12B (specs, benchmarks, license) — https://awesomeagents.ai/models/nvidia-nemotron-3-super-120b-a12b/
- [S33] DEV Community: NVIDIA Nemotron 3 Super — the open AI model that just beat GPT on coding (March 2026) — http://dev.to/akaranjkar08/nvidia-nemotron-3-super-the-open-ai-model-that-just-beat-gpt-on-coding-march-2026-dg0
- [S34] GitHub: frank-1150/learning-ai-technical — Nemotron 3 Super hybrid architecture note (PinchBench 85.6%) — https://github.com/frank-1150/learning-ai-technical/blob/HEAD/docs/machine-learning/neural-networks/nemotron-3-super-hybrid-architecture.md
- [S35] GitHub: muhammadsaqlainaslam/my-llm-wiki — Nemotron 3 Super (arXiv 2604.12374, technical report) — https://github.com/muhammadsaqlainaslam/my-llm-wiki/blob/HEAD/wiki/Nemotron_3_Super.md
- [S36] GitHub: cpm0722/nvidia-nemotron-hackathon-2026 — Nemotron 3 Super 120B report (AA indices, RULER, HF card figures) — https://github.com/cpm0722/nvidia-nemotron-hackathon-2026/blob/HEAD/example_runs/nemotron-3-super-120b/report_nemotron-3-super-120b.md
- [S37] SiliconANGLE: NVIDIA launches Jetson Thor into general availability (2025-08-25) — https://siliconangle.com/2025/08/25/powering-age-robotics-nvidia-launches-jetson-thor-general-availability/
- [S38] Assured Systems/Vecow: EAC-7000 series with NVIDIA Jetson Thor (T5000 specs, carrier detail) — https://www.assured-systems.com/vecow-launches-cutting-edge-eac-7000-series-with-nvidia-jetson-thor-to-drive-the-future-of-edge-ai-robotics/
- [S39] GitHub: redhat-et/physical-ai-platform-intel — NVIDIA Jetson family analysis (July 2026, four tiers) — https://github.com/redhat-et/physical-ai-platform-intel/blob/HEAD/deliverables/references/nvidia-jetson-family.md
- [S40] TrendForce: NVIDIA expands portfolio — GTC 2026 disaggregated inference, Groq LP30/LP40 — https://www.trendforce.com/presscenter/news/20260318-12976.html
