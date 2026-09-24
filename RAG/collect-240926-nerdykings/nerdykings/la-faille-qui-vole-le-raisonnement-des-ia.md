---
id: collect-240926-nerdykings/nerdykings/la-faille-qui-vole-le-raisonnement-des-ia
title: "The Flaw That Steals AI Reasoning"
domain: nerdykings
role: reference
task: reference
actors: ["Anthropic", "Google", "Hugging Face", "Moonshot", "OpenAI", "Z.ai"]
dates: ["2026-05"]
keywords: ["reasoning", "agent", "agents", "claude", "distillation", "glm", "kimi", "training"]
source: docs/RAG/clean_en/nerdykings/la-faille-qui-vole-le-raisonnement-des-ia.md
source_anchor: ""
source_lines: [1, 59]
sha256: a92dcc2f8af2a5af9039a07ebde27daa66fb28f360013433665fb3061c18895f
---

# The Flaw That Steals AI Reasoning

<!-- source: https://www.nerdykings.com/blog/la-faille-qui-vole-raisonnement-ia.html -->

# The Flaw That Steals AI Reasoning

AI labs are doing absolutely everything to hide how their models think. And yet, rather clever researchers have just proven that these secret reasonings can be recovered from OpenAI, Anthropic, and Google — without any complex attack. As a bonus: the same flaw made it possible to recover API keys, passwords, and personal information that developers thought they had protected. Here's how it works, and what it changes.

## Why AIs hide their reasoning

When you pose a complex problem to a reasoning model, it doesn't just blurt out its answer directly. It first generates a long series of intermediate steps: it tests leads, backtracks, checks its calculations, corrects its errors. This chain of thought is precious — it shows exactly *how* the model solves a problem.

The problem is that a competing lab could ask a powerful model to solve thousands of problems, recover all these detailed demonstrations, then use them to train a smaller and cheaper model. This is what's called **distillation**. And it's precisely to limit this kind of copying that the major providers no longer show the full reasoning of their models — you get the final answer, sometimes a short summary, and the real internal monologue stays hidden.

Except that hiding information in the interface doesn't mean it becomes impossible to recover.

## A first method: reconstructing reasoning by approximation

Even before the discovery of the main flaw, a team had already shown that part of the hidden reasoning could be reconstructed — a method dubbed **Trace Inversion**. The principle: train a second model to guess the logical path between a question and its answer, relying on examples where the full reasoning is known.

Result: this model produces a chain of reasoning that resembles the one that could have been followed. It's not a literal reading of thoughts, but an approximation — and in the tests, small models trained on these reconstructed reasonings obtained far better results in math than those trained only on the final answers. Already not bad. But another team went much further by attacking the real encrypted reasoning directly.

## The real flaw: an encrypted block that can simply be moved

To understand, you have to look at how certain APIs work. A provider doesn't necessarily keep the whole conversation on its servers: it's often the application that retains the history and sends it back with each new message. A reasoning model can generate thousands of internal steps that it must find again later — but it's impossible to send them in plaintext, the user would read them. The provider therefore transforms them into an **encrypted and unreadable block of text**, which the application stores and sends back on the next turn.

The user can neither read it nor modify it without the AI noticing. But they can still do one thing: **move the entire block without touching its content**. That's exactly where things started to go wrong. In May 2026, security specialist Matthew Green discovered that a block generated in one conversation can be reused elsewhere — sometimes transferred to another account, or even another model of the same family. In one experiment, he created a block containing a reflection about a social security number, transferred it to another session, and the new model eventually revealed the number.

## How they turned this into a complete extraction method

The attack relies on a fairly logical security difference: a company's most powerful models are better protected — ask it directly to reveal its reasoning, and it refuses. But the smaller and cheaper models of the same family are often less well armored against this kind of workaround, while still being capable of understanding their big brother's thoughts.

The protocol: ask a powerful model to solve a problem, recover the encrypted block containing its reasoning, then transfer it to a weaker model from the same company by convincing it to transcribe it. The big model is never attacked directly — it's its little brother that unwittingly serves as a decryptor. In a demo, Claude Opus solved a math problem and produced a hidden thought; transferred to Claude Haiku, it came out as perfectly readable text. Variants also worked with models from OpenAI and Google.

To verify that this wasn't just an explanation invented by the small model, the researchers compared the number of reasoning tokens billed by the API with the length of the reasonings recovered across 120 coding problems — the two matched very closely.

## The real danger: 315,000 decrypted blocks found online

This is where it gets frankly worrying. An AI agent can handle an API key during its work without ever displaying it in its final answer — except that it remains present in its internal reasoning. Yet many developers publish their agents' complete sessions to debug them or share their experience. They delete the visible passwords, but keep the encrypted blocks — since in theory, they're unreadable.

The researchers analyzed **6,708 agent sessions published on GitHub and Hugging Face**, and thanks to this flaw, they decrypted more than **315,000 reasoning blocks**. Inside: email addresses, API keys, passwords, access tokens, and even private keys. Data that was thought to have been deleted from the visible output still existed in the AI's encrypted thoughts.

And there's a second perverse effect: a complete reasoning trace reveals the entire *method* used to arrive at a result, not just the result. Enough to constitute a gold dataset for training a competing model through distillation.

## Strangely similar reasoning styles

Once in possession of reasoning traces from several proprietary AIs, the researchers were able to compare the styles of GPT and Claude with open models. By giving the first few words of a Claude Opus thought to other models, some — like Kimi K3 or GLM 5.2 — came noticeably close to Claude's style. Kimi K3 in particular was between 10,000 and 1 million times more likely to correctly continue a Claude or GPT reasoning trace than other models tested.

That number seems enormous, but the actual probability remained very low — it would have taken billions of attempts for a perfect match. And these similarities can also come from shared training data, or simply from the fact that several models naturally converge on the same way of solving a problem. Intriguing, but it absolutely does not prove that one company copied another.

## A paradoxical security problem

When a model receives a dangerous request, it can think in detail about how to execute it, then decide only at the end to refuse. The visible response remains clean — but the dangerous details can remain in the hidden reasoning, and therefore become recoverable via this flaw. And the flaw also works in reverse: instead of stealing a thought, an attacker can **insert a fake one** into the conversation to discreetly influence an agent's behavior, without the application being able to verify anything since the content is unreadable.

## My take

The researchers warned the providers before publishing, and at the time of publication, the attacks were no longer reproducible — the most likely solution: linking each encrypted thought to the account, the conversation, and the model that created it. So breathe easy, it's not an active flaw today. But what strikes me most in this story is the principle it reveals: **a hidden thought is not a protected thought**. As soon as it is transported, stored, or read back by another AI, it becomes data like any other — and therefore a target. And for developers who work with agents, the lesson is simple: stop publishing raw logs thinking that the encrypted gibberish contains nothing important. Also change your credentials if they're lying around in old public journals. An AI's reasoning has become almost as valuable as the model itself.

### 🛠️ Tools you can test related to this article

A selection of my tested tools, relevant for going further.
