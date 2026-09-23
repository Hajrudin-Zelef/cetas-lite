---
id: collect-mindstudio/mindstudio/dark-bloom-distributed-ai-inference
title: "Dark Bloom: Rent Out Your Mac for AI Inference and Get Paid"
domain: mindstudio
role: reference
task: article
actors: ["Alibaba", "Apple", "OpenRouter", "Stripe"]
dates: ["2026-09-23"]
keywords: ["inference", "compute", "cost", "gpu", "inference engine", "memory", "open-weight", "qwen", "revenue", "tokens per second"]
source: docs/RAG/Collect RAG/02_mindstudio/dark-bloom-distributed-ai-inference.md
source_anchor: ""
source_lines: [1, 54]
sha256: d6e05e317eff997dbec8af600b86ae80eac444bc9370042991080654c11a3487
---

# Dark Bloom: Rent Out Your Mac for AI Inference and Get Paid

## Metadata

- **Source** : https://www.mindstudio.ai/blog/dark-bloom-distributed-ai-inference
- **Site** : MindStudio
- **Type** : Article
- **Language** : en
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

This article profiles Dark Bloom, a distributed inference network that lets Mac owners share unused compute with people who want to run AI models, in exchange for money. Instead of routing requests through a centralized data center, Dark Bloom splits inference across a network of individual Macs, each running open-weight models like Qwen, Gemma, and GPT-OSS. Anyone with a compatible Mac can install the software, enroll their device, and earn based on how much inference work their machine handles. The article notes Dark Bloom had processed billions of tokens within its first week or so of public availability and serves models through OpenRouter, reportedly at lower prices than other providers.

The basic mechanic: install software, enroll as a provider node, and the network routes inference requests to your hardware when needed. Your Mac's GPU processes the request using Apple's MLX inference engine, and results return to the requester. This differs from traditional cloud inference because Dark Bloom aggregates thousands of individually owned machines into something functioning like a distributed data center — no single building, no centralized ownership, hardware owned by regular people. The models served are open-weight: Qwen, Gemma, and GPT-OSS.

On privacy: the article notes this raises the most questions, and Dark Bloom published a white paper. The stated problem is that running inference on a stranger's computer means that stranger has root access and physical possession and could in theory observe prompts and responses. Dark Bloom's answer is architectural: inference runs entirely inside a single hardened Swift process, with no subprocesses, no local server, and no interprocess communication that could expose data, using MLX Swift LM (Apple's inference library for the Apple silicon GPU). The goal is to eliminate software pathways through which an owner could inspect traffic. This doesn't make it provably unbreakable, but the design narrows leak paths and the code is published for review. Independent AI-assisted code review cited during early coverage reportedly found no hidden malware, no cryptocurrency mining, and no credential theft, though automated review has limits and isn't a formal security audit.

On earnings: expectations should be tempered. Estimates for a well-equipped Mac Studio with large memory point to earnings in the tens of dollars per month, not a replacement for a day job. Apple silicon is efficient, so incremental electricity cost is small but not zero; actual profit depends on local rates and network utilization. For those interested in distributed compute as an alternative to concentrated data centers, the calculus changes: traditional AI infrastructure depends on massive centralized facilities drawing enormous power and becoming a point of local controversy, whereas Dark Bloom uses computers already idle in homes and offices. Two caveats: the project currently pays 100% of revenue to node operators (not guaranteed to stay), and it's genuinely early software with rough edges and shifting hardware requirements.

Enrollment happens through a CLI tool (a Mac app is planned), involving creating an account, installing the CLI, and enrolling the device through macOS device management. Setup includes a hardware verification step testing inference speed in tokens per second before fully activating. The project reportedly raised its minimum memory requirement to 48GB of RAM after more demand than expected, using the higher bar as a quality filter — excluding lower-memory Macs but covering most Mac Studios, Mac minis, and higher-end MacBook Pros. Payment flows through Stripe rather than direct bank access, so Dark Bloom can only send payouts through Stripe as an intermediary.

## Key points

- Dark Bloom is a peer-to-peer inference network turning idle Macs into provider nodes serving Qwen, Gemma, and GPT-OSS.
- Enrollment via CLI at darkbloom.dev (Mac app in development); payouts through a connected Stripe account.
- Privacy architecture: inference runs in a single hardened Swift process using Apple's MLX framework to prevent owners from viewing prompts/responses.
- Already serves models through OpenRouter, reportedly cheaper than competitors; billions of tokens processed in the first week.
- Minimum 48GB RAM required (raised due to demand), covering most Mac Studios/minis and higher-end MacBook Pros.
- Very early stage, self-described as days old, with openly available code; AI-assisted review found no malware/mining/credential theft.
- Earnings estimate: high-memory Mac Studio in the tens of dollars per month, minus modest electricity.
- Providers currently keep 100% of revenue, though the split may change.

## Technical data / figures

| Item | Value |
|---|---|
| Network type | Distributed / peer-to-peer inference |
| Models | Qwen, Gemma, GPT-OSS |
| Framework | Apple MLX (MLX Swift LM) |
| Privacy design | Single hardened Swift process |
| Enrollment | CLI (Mac app planned) |
| Payouts | Stripe (intermediary) |
| Min RAM | 48GB |
| Early token volume | Billions within ~1 week |
| OpenRouter | Models served, lower prices claimed |
| Earnings estimate | Tens of $/month (high-memory Mac Studio) |
| Revenue split | 100% to providers (may change) |

## Why this source matters for the RAG

It documents an emerging distributed-compute model that could reshape inference economics and privacy, with concrete architecture, payout, and hardware requirements. It is a key reference for peer-to-peer AI inference and the broader decentralization trend in the RAG.

