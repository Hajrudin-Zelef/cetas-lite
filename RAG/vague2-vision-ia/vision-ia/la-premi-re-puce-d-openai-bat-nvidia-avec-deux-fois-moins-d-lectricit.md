---
id: vague2-vision-ia/vision-ia/la-premi-re-puce-d-openai-bat-nvidia-avec-deux-fois-moins-d-lectricit
title: "La première puce d'OpenAI bat Nvidia, avec deux fois moins d'électricité"
domain: vision-ia
role: reference
task: article
actors: ["AMD", "AWS", "Anthropic", "Broadcom", "China", "CoreWeave", "DeepSeek", "EU", "Google", "Meta", "Moonshot", "Nvidia", "OpenAI", "SpaceX", "Stability AI", "United States"]
dates: ["2025-02", "2025-09", "2025-11", "2026-05", "2026-08-26", "2026-09-23"]
keywords: ["nvidia", "accelerator", "agent", "agentic", "amd", "aws", "benchmarks", "blackwell", "chatgpt", "claude", "compute", "cost"]
source: docs/RAG/Collect RAG Vague 2/04_vision_ia/la-premi-re-puce-d-openai-bat-nvidia-avec-deux-fois-moins-d-lectricit.md
source_anchor: ""
source_lines: [1, 54]
sha256: 3cfb199718892a4affd17785a615ad24ec502d9d9483ff9f86df4cb556107821
---

# La première puce d'OpenAI bat Nvidia, avec deux fois moins d'électricité

## Metadata

- **Source** : https://vision-ia.beehiiv.com/p/la-premi-re-puce-d-openai-bat-nvidia-avec-deux-fois-moins-d-lectricit
- **Site** : Vision-IA (beehiiv)
- **Type** : Newsletter article
- **Language** : fr (original) / summary in English
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This August 26, 2026 issue leads with OpenAI unveiling "Jalapeño," its first in-house accelerator, designed with Broadcom, at the Hot Chips conference. Independent SemiAnalysis tests put it ahead of Nvidia's GB200 and GB300: 1.5–1.9x more throughput per kilowatt, 1.7–3.6x lower latency, and 700 watts consumption versus 1,400 watts for a GB300. It is an inference-only chip—it runs already-trained models, optimized for the two real LLM bottlenecks, the prefill phase and inter-chip communication. SemiAnalysis benchmarks (InferenceX suite) on GPT-OSS 120B, DeepSeek R1 670B, and Kimi K2.5 1T show ~1,400 tokens/second on GPT-OSS 120B, handling 2.1–4.1x more interactive load than a GB200/GB300, with 50% more HBM memory per watt. It did not directly face Vera Rubin but beats Nvidia/CoreWeave's July-published throughput-per-megawatt figures for Rubin. Design began mid-2024, frozen November 2025. Only engineering samples exist: first units late 2026, volume ramp in 2027. Dylan Patel of SemiAnalysis notes first-generation chips are rarely competitive, yet OpenAI beats Nvidia's Blackwell and even Rubin. OpenAI maintains its agreements with Nvidia, AMD, and AWS—a strong negotiating position.

Next, OpenAI announced it dismantled a Russian account network using ChatGPT to fuel a Western influence operation built around an entirely fictitious think tank, the "International Burke Institute." Of the 36 articles published between September 2025 and May 2026, 34 were plagiarized. The IBI site, registered February 2025, claimed to be based in Israel and cited Francis Fukuyama and Noam Chomsky as intellectual cover without any real link. Operators used VPNs to bypass OpenAI's geographic blocking; ChatGPT drafted, translated, produced Russian syntheses for a second operator, and erased linguistic markers betraying a Russian pen. The institute published a homemade "sovereignty index" ranking Russia ahead of Western democracies. Distribution was on X, LinkedIn, Facebook, Substack, and Telegram, including a German channel "Lahme Ente" hostile to the EU, Ukraine, and the German government, with 10,000–20,000 subscribers. Targets: Germany, US, France, Poland, Turkey. OpenAI classifies it category 3 of 6 on the Brookings scale: credible infrastructure, tiny real audience.

Anthropic merged Claude's chat memory with that of Claude Cowork, its office agent. What you say in a conversation, the agent already knows when acting. Deployment was the same day on Free, Pro, and Max plans across web, desktop, and mobile. Memory fills during conversation rather than by an end summary, so you can switch from chat to agent mid-task. Everything retained is viewable, editable, and deletable by topic. By default, Claude doesn't store health data, ethnicity, gender identity, religion, or political opinions; an "include sensitive topics" setting re-enables them. IDs, social security numbers, criminal records, and immigration status are never stored. Limitation: chat and Cowork memory cannot be compartmentalized without two separate accounts; Claude Code keeps separate memory for now.

SpaceX filed a "Starmind" dossier with the FCC for an orbital data center constellation in low Earth orbit (500–2,000 km) that could reach one million satellites. Each unit carries 72 Nvidia chips—equivalent to a full server rack—and the first, AI1, targets a Q4 2027 launch. AI1 runs on the Vera Rubin NVL72 "Space-1" platform (Rubin GPU + Vera CPU), with Nvidia claiming up to 25x an H100's AI compute. Heliosynchronous orbit faces the Sun 98% of the time; solar panels produce ~210 kW, for 175 kW average compute per satellite. Heat dissipates through ~1,700 sq ft (~160 m²) liquid radiators. Results return via laser links relayed by Starlink. It is a filed dossier, not an authorization; no FCC decision or public budget. Research items cover an mRNA adjuvant shrinking tumors in mice, MIT Media Lab's finding that four weeks of chatbot use reduced fake-news detection by 15%, Multiverse's 4-bit Quantization-Aware Healing, MIT's extreme weather generation, underwater sonar navigation, Inherent's Faraday, EleutherAI's lie-detection retrospective, Ukraine opening 5 million combat images to the UK, and student model preferences. Briefs cover Chinese hackers doubling attacks with DeepSeek, ChatGPT task scheduling for free accounts, OpenAI's Codex 5-hour limit, Google buying Spirit Airlines data, Meta's Hatch agent ($199.99/month), Waymo in Munich, Einride's 500 Tesla Semis, BlackBerry/QNX, Kiwi Health, hiring overload, agentic automation skepticism, Gartner's agent project forecast, OpenAI's data center chief departure, SpaceX's Louisiana spaceport, and Stability AI's $76M raise with music majors.

## Key points

- OpenAI's Jalapeño inference chip, built with Broadcom, beats Nvidia's GB200/GB300 on throughput-per-kW and latency while using half the power.
- Jalapeño is inference-only, with engineering samples only; volume ramp in 2027.
- OpenAI dismantled a Russian influence operation built on a fake think tank, with 34 of 36 articles plagiarized.
- Anthropic unified Claude chat and Cowork memory across Free/Pro/Max, with per-topic controls.
- SpaceX filed for a one-million-satellite orbital data center constellation (Starmind), first launch Q4 2027.
- Research: four weeks of chatbot use reduced fake-news detection by 15%; a 4-bit compressed model beat its own bf16 original on 7/9 benchmarks.

## Technical data / figures

| Item | Value |
| --- | --- |
| Jalapeño throughput per kW | 1.5–1.9x GB200/GB300 |
| Jalapeño latency | 1.7–3.6x lower |
| Jalapeño power | 700 W (vs 1,400 W GB300) |
| Jalapeño tokens/sec (GPT-OSS 120B) | ~1,400 |
| Jalapeño HBM per watt | +50% |
| Jalapeño design freeze | November 2025; volume 2027 |
| Fake think tank articles | 34 of 36 plagiarized |
| German channel subscribers | 10,000–20,000 |
| Starmind satellites (target) | up to 1,000,000 |
| Starmind chips per unit | 72 Nvidia (Vera Rubin NVL72) |
| Starmind satellite power | ~210 kW solar, 175 kW compute |
| Starmind first launch | Q4 2027 |
| MIT chatbot fake-news effect | +21% initially, -15% after 4 weeks |
| Quantization-Aware Healing | GPT-OSS 120B 60B pruned, 4-bit, beats bf16 on 7/9 benchmarks |
| Meta Hatch price | up to $199.99/month |
| Stability AI raise | $76M (total $232M) |

## Why this source matters for the RAG

It documents a major hardware milestone—OpenAI's first custom inference chip beating Nvidia—with precise efficiency figures, central to AI-infrastructure and cost queries. It also captures AI-enabled influence operations, agent memory unification, and orbital data centers. These are high-value for hardware, safety, and product-trend retrieval.
