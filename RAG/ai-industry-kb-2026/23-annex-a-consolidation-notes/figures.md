---
id: ai-industry-kb-2026/23-annex-a-consolidation-notes/figures
title: "Figures"
domain: appendix
role: appendix
task: reference
actors: ["Alibaba", "Anthropic", "China", "DeepSeek", "EU", "Google", "Groq", "Hugging Face", "MiniMax", "Moonshot", "Nebius", "Nvidia", "OpenRouter", "SGLang", "SpaceX", "United States", "Z.ai", "vLLM", "xAI"]
dates: ["2025-04", "2025-05", "2026-04", "2026-08-03", "2026-08-12", "2026-08-14", "2026-09-02", "2026-09-03", "2026-09-04", "2026-09-22"]
keywords: ["acquisition", "apache", "attention", "attribution", "benchmark", "benchmarks", "compute", "cost", "deepseek", "fp8", "glm", "gpu"]
source: docs/RAG/ai-industry-knowledge-base-2026.md
source_anchor: ""
source_lines: [11149, 11214]
section: "Annex A — Consolidation notes"
sha256: 287a34d716ddbaf78d665288193f3742ab2a4e0187e1fe3f72c3e4623b46e027
---

# Figures

| Wrong / prior value | Corrected value | § refs |
|---|---|---|
| "vLLM-Omni v0.28.0" (Wave 2) | Core vLLM v0.28.0; **Omni is a separate repo/product line** | §6, §7 |
| MiniMax H3 / Hailuo 3.0 / "Hailuo 03" as separate entries | **One model, three names** — merge; distinct from M3 | §3 |
| "SpaceXAI" (press naming drift) | **xAI** — no corporate confirmation | §1 |
| "Compressed Expert Dispatch" (CED) | **"Causal Encoder-Decoder"** — the only sanctioned expansion; grep-and-delete every occurrence of the former | §5 |
| "Sohgo" | **Sohu** | §15 |
| Source typo "Nevus" | **Nebius** | §15 |
| "Pax Sillica" | **Pax Silica** | §18 |
| Super Micro indictment naming | **"Super Micro"** = the company (not a defendant); defendants are **Liaw / Chang / Sun** (individuals only) | §18 |
| Manus / Butterfly Effect | **Manus** = the product, **Butterfly Effect** = the company | §18 |
| NVIDIA–Groq "acquisition" | **Non-exclusive license + talent** | §15, §20 |
| ExLlamaV2 (current) | Archived; **ExLlamaV3 / EXL3** (v1.4.6–v1.4.9, Sept 2026) is the live line | §8 |
| AutoAWQ / AutoGPTQ (current) | AutoAWQ deprecated (May 2025 → confirmed Sept 2026) → **llm-compressor**; **GPTQModel** replaces AutoGPTQ | §8 |
| Marlin compute floor "SM80" | **SM75** (Turing); Machete SM90-exact; Marlin-24 kernels removed from vLLM main | §8 |
| `VLLM_NVFP4_GEMM_BACKEND` | Deprecated → **`--linear-backend`** | §8 |
| optimum-quanto (active) | **Maintenance mode** (2026-09-22) | §8 |
| Qwen3.8-Max (single entry) | Three distinct index entries: **Qwen3.8-Max** (hosted GA 2026-08-03, multimodal, 1M ctx, $2/$6) ≠ **Qwen3.8-2.4T-A95B** (downloadable weights 2026-08-12, text-only, thinking-mode-only, custom license) ≠ **Qwen3.8-Max-0902** (checkpoint 2026-09-02). Also ≠ Qwen3-8B (April 2025). | §4 |
| Qwen3.8-Flash vs Flash-Next | **Qwen3.8-Flash** (open weights, 6B active) ≠ **Qwen3.8-Flash-Next** (hosted-only, 125B/6B, Qwen4-generation preview) | §4 |
| Qwen3.8-27B vs Qwen3.6-27B | **Qwen3.8-27B** (weights 2026-08-14, Apache 2.0) ≠ **Qwen3.6-27B** (date unpinned [DATE UNVERIFIED], benchmark-tables only) | §4 |
| MiMo-V2.6-Flash drafter | Model card says 5-layer DFlash-style MTP drafter; **shipped config `num_nextn_predict_layers: 3` is authoritative** for memory sizing | §11 |
| DeepSeek model ids | `deepseek-flash` (V4.1-Flash GA), `deepseek-v4-pro`, `deepseek-v4-flash`, `deepseek-v4-flash-vision-exp`, `deepseek-v4.1-flash-expires-on-0910` (beta probe) — code formatting, exact ids | §5 |
| "Agreed to acquire" vs "completed purchase" | Use **"agreed to acquire"** (NVIDIA–HF) — deal not closed as of cutoff | §16 |

### Figures

| Wrong / prior value | Corrected value | § refs |
|---|---|---|
| GLM-5.2 "744B total" | **753B total / ~40B active** (743.4B backbone config-derived; 753.3B with 9.95B MTP block; 39.3B active); "744B-A40B" = Z.ai's FP8-build shorthand, footnoted | §2, §3 |
| GLM-5.3 "743B" as the headline total | 743B (announcement figure) = config-derived backbone — the number Z.ai used for the "same base" claim; 753B/40B weights (756GB FP8) = full-weight total. Keep the roles straight. | §2, §3 |
| Opus 4.6 "$15/$75" (businessworld.in) | Dropped — extended-tier/editorial error vs $5/$25 consensus | §1 |
| Vals Index delta for V4-Pro-0813 "+10.6" | **+9.48** | §5 |
| AA Intelligence Index "53" | **53 (v4.1-era)** vs **≈36 (v4.3, rebased 2026-09-04)** — never compare across methodology versions | §5 |
| Vals Index rank formatting | **"52.37%, #18 (2026-08-12)"** / **"12th (mid-September press)"** — never a bare rank | §5 |
| "Inference ≈70% of AI cloud spend" | **55% global** (Gartner audited, 2026; 59% in 2027); 70–80% only for GPU-cloud spend among production AI teams | §14 |
| MLA "exactly 4× batch/GPU" | Use **93.3%** (DeepSeek-V2 paper) / **5.76×** throughput / 14.2× toy arithmetic [DIRECTIONAL]; "4×" = [UNVERIFIED] shorthand | §7 |
| "~90% MLA reduction" | **93.3%** (DeepSeek-67B baseline); ablation ~96%; raw element ratio ~98.2% | §7 |
| HybridKV "7.9×" | VERIFIED — "up to/best-case across 11 benchmarks" (Qwen2.5-VL-7B; ACL 2026 long paper) | §7 |
| TurboQuant "6×/8×" | VERIFIED — 6× memory; 8× attention speedup (4-bit vs 32-bit unquantized keys, H100; Google, independently evaluated by Red Hat llmkube#308) | §7 |
| Hugging Face "927K+ datasets" | Dropped [UNVERIFIED]; 500K (2026-09-03 deal coverage) + 730K–1M methodology range | §16 |
| "SGLang 2.5x cache hit rate vs competition" | [UNVERIFIED] — drop the formulation | §6 |
| "400,000+ GPUs" (SGLang) | [UNVERIFIED] — repeated again in RadixArk coverage, still no primary citation | §6 |
| GB300 "25×" (SGLang) | [VENDOR, baseline unaudited] — never quote without attribution and baseline caveat | §6 |
| "4x batch size per GPU" | [PARTIALLY VERIFIED] — derived planning figure only | §6 |
| GLM-5.3-Flash context "300K" (LumaDock evals) | **1M (config)** used; basis of the 300K figure unclear; OpenRouter lists 1,310,720 | §9 |
| DeepSeek V4 efficiency at 1M | 27%/10% (V4-Pro), 10%/7% (V4-Flash) vs V3.2 — via felloai, **secondary not vendor-primary**; MLA paper figures 93.3%/5.76x attributed to the DeepSeek-V2 MLA paper | §6, §7 |
| Mooncake "+115%/+107%/+75%" | Paper-era — replaced by verified 2026 measurement: TCP 17–22%, RDMA 26–33% vs Redis (Tsinghua ToS 2025 evaluation of Mooncake Store) | §7 |
| Unitree "$51B" context | $51B = **market cap**, not the ¥6.1B raise | §22 |
| Tesla Optimus | 1M/yr = **design capacity**, not production (Musk's Jan 2026 "not in usage in our factories in a material way" governs all dating) | §22 |
| ">97% Chinese-vendor share" | Tracker definition with contested denominator — corrected at every appearance | §22 |
| Kimi K2.6 context "256K" | **256K** shorthand = **262,144** exact tokens | §3 |
| Kimi K2.6 / K3 expert counts | K2.6: **384 total**, 8 routed + 1 shared per token. K3: **896 routed + 2 shared**, 16 active per token — never mixed. | §3 |
| Kimi K3 total | **2.8T** canonical; 2.72T routed and 33.0M/expert carry the flagged inconsistency — never headline facts | §3 |
| MiMo-V2.6 (309B/15B, 1M ctx) | [UNVERIFIED] — sources found MiMo V2.5 Pro at 1,020B total (April 2026 size chart), not the stated V2.6 specs | §7 |
| MiniMax H3 open weights | Territorial restriction: excludes US/EU/UK/South Korea local deployment; weights Aug 2–3 | §11 |
| Kimi K2.6 SWE-Bench Pro | **58.6%** — ahead of, not "near," Opus 4.6 (53.4%) | §3 |
| V4 "890 B/token" claim | No independent reproduction as of 2026-09-22 | §5 |
| HBM "sold out through 2028" / HBM prices | Financial-press/secondary-sourced (ainvest, momoview, TechTimes, AI Cost Estimator), not vendor disclosures — kept qualified; Goldman counter-view is the only bearish datapoint | §7 |
| NVIDIA $12.9B Hugging Face bid | **Excluded entirely** — single-source [UNVERIFIED] | §14 |

---

## A.4 Superseded claims

Claims replaced by later-wave information; the old wording must not be carried forward.

