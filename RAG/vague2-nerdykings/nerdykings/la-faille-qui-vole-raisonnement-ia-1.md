---
id: vague2-nerdykings/nerdykings/la-faille-qui-vole-raisonnement-ia-1
title: "La Faille Qui Vole Le Raisonnement Des IA"
domain: nerdykings
role: reference
task: article
actors: ["Anthropic", "Google", "Hugging Face", "Moonshot", "OpenAI", "Z.ai"]
dates: ["2026-05", "2026-09-23"]
keywords: ["agent", "agents", "claude", "distillation", "glm", "kimi", "reasoning", "training"]
source: docs/RAG/Collect RAG Vague 2/03_nerdykings/la-faille-qui-vole-raisonnement-ia.md
source_anchor: ""
source_lines: [1, 57]
sha256: 8ac60e06ca766722e91bab8b895249e5075d21b84a2223fb62b41cdf105d1582
---

# La Faille Qui Vole Le Raisonnement Des IA

## Metadata

- **Source** : https://www.nerdykings.com/blog/la-faille-qui-vole-raisonnement-ia.html
- **Site** : NerdyKings
- **Type** : Article
- **Language** : fr (original) / summary in English
- **Verification status** : ✅ reachable
- **Collection date** : 2026-09-23

## Full summary

AI labs go to great lengths to hide how their models reason. Yet researchers have proven that the hidden reasoning of OpenAI, Anthropic, and Google models can be recovered without a complex attack. As a bonus, the same flaw allowed recovery of API keys, passwords, and personal information developers thought were protected.

Why AIs hide reasoning: a reasoning model first generates a long chain of intermediate steps (testing leads, backtracking, checking calculations, correcting errors). This chain of thought is precious — it shows exactly how the model solves a problem. A competing lab could ask a powerful model to solve thousands of problems, harvest all these detailed demonstrations, and use them to train a smaller, cheaper model — **distillation**. To limit copying, major providers stopped showing full reasoning: users get the final answer, sometimes a small summary, while the real internal monologue stays hidden. But hiding information in the interface doesn't make it impossible to recover.

A first method, **Trace Inversion**, reconstructs part of the hidden reasoning by training a second model to guess the logical path between a question and its answer, using examples where full reasoning is known. The reconstructed chain approximates the original; in tests, small models trained on these reconstructed reasonings performed much better in math than those trained only on final answers. Another team went further by attacking the real encrypted reasoning.

The real flaw: certain APIs don't keep the whole conversation server-side — the app stores history and resends it each turn. A reasoning model may generate thousands of internal steps it must retrieve later, but can't send them in clear text. So the provider turns them into an unreadable **encrypted block**, which the app stores and returns next turn. The user can't read or modify it without the AI noticing — but can still move the whole block without touching its content. That's where it broke. In May 2026, security specialist **Matthew Green** found a block generated in one conversation could be reused elsewhere — sometimes transferred to another account, even another model of the same family. In one experiment, he created a block containing a reflection about a social security number, transferred it to another session, and the new model revealed the number.

Turning it into a full extraction method: stronger models of a company are better protected (ask directly and they refuse), but smaller, cheaper models of the same family are often less hardened while still able to understand their big sibling's thoughts. Protocol: ask a powerful model to solve a problem, retrieve the encrypted block containing its reasoning, transfer it to a weaker model of the same company, and convince it to transcribe it. The big model is never attacked directly — its little sibling unknowingly acts as a decryptor. In a demo, Claude Opus solved a math problem and produced hidden thought; transferred to Claude Haiku, it came out as perfectly readable text. Variants also worked with OpenAI and Google models. To verify it wasn't a fabricated explanation, researchers compared the number of reasoning tokens billed by the API with the length of reasonings recovered over 120 coding problems — they matched very closely.

The real danger: **315,000 blocks decrypted online**. An AI agent can handle an API key during work without ever showing it in the final answer — but it remains in the internal reasoning. Many developers publish full agent sessions to debug or share experiences. They delete visible passwords but keep the encrypted blocks, since they're theoretically unreadable. Researchers analyzed **6,708 agent sessions published on GitHub and Hugging Face** and decrypted **over 315,000 reasoning blocks** containing emails, API keys, passwords, access tokens, and even private keys. Data thought removed from the visible result still existed in the AI's encrypted thoughts. A second perverse effect: a full reasoning reveals the entire method used to reach a result, not just the result — a goldmine dataset for distilling a competing model.

On reasoning styles: with several proprietary AIs' reasonings in hand, researchers compared GPT and Claude styles with open models. Giving the first words of a Claude Opus thought to other models, some — like **Kimi K3** or **GLM 5.2** — came noticeably close to Claude's style. Kimi K3 was between 10,000 and 1 million times more likely to correctly continue a Claude or GPT reasoning than other tested models. The figure seems huge, but the actual probability remained very low — billions of attempts would be needed for a perfect match. Such similarities may also come from shared training data or natural convergence. Intriguing, but no proof one company copied another.

A paradoxical security problem: when a model receives a dangerous request, it may think in detail about how to execute it, then decide at the end to refuse. The visible answer stays clean, but dangerous details may remain in hidden reasoning and become recoverable. The flaw also works in reverse: instead of stealing a thought, an attacker can insert a fake one into the conversation to discreetly influence an agent's behavior, without the app being able to verify since the content is unreadable.

The author's view: the likely fix is to bind each encrypted block to the account, conversation, and model — a moved block is then refused. Researchers warned providers before publishing, and by publication the attacks were no longer reproducible, so it isn't an active flaw today. But the principle revealed is key: **a hidden thought is not a protected thought**. Once transported, stored, or reread by another AI, it becomes data like any other — and thus a target. For developers working with agents: stop publishing raw logs thinking the encrypted gibberish contains nothing important, and change credentials if they linger in old public logs. An AI's reasoning has become almost as valuable as the model itself.

## Key points

- Hidden reasoning of OpenAI, Anthropic, and Google models can be extracted without complex attacks.
- Motivation for hiding: preventing **distillation** of powerful models into cheaper ones.
- **Trace Inversion** reconstructs an approximation of hidden reasoning from question-answer pairs.
- Core flaw: encrypted reasoning blocks can be moved/reused across sessions, accounts, or models.
- Attack chains a strong model's encrypted thought into a weaker sibling model that transcribes it (e.g., Claude Opus → Claude Haiku).
- Discovered by security specialist **Matthew Green** in May 2026 (social security number test).
- Analysis of **6,708** public agent sessions (GitHub/Hugging Face) decrypted **315,000+** reasoning blocks containing secrets.
- Reasoning style similarities (Kimi K3, GLM 5.2 vs Claude) are intriguing but not proof of copying.
- Flaw works in reverse: injecting fake hidden thoughts can covertly influence an agent.
- Likely fix: bind encrypted blocks to account + conversation + model; attacks no longer reproducible at publication.

## Technical data / figures

| Item | Value |
|---|---|
| Agent sessions analyzed | 6,708 |
| Reasoning blocks decrypted | 315,000+ |
| Coding problems used for verification | 120 |
| Discovery date | May 2026 (Matthew Green) |
| Style-continuation odds (Kimi K3) | 10,000x to 1,000,000x more likely |
| Affected providers | OpenAI, Anthropic, Google |
| Demo models | Claude Opus → Claude Haiku |
| Publication status | Attacks no longer reproducible (fixed) |

