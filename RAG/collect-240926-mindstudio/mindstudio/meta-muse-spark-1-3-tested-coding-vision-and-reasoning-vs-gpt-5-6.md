---
id: collect-240926-mindstudio/mindstudio/meta-muse-spark-1-3-tested-coding-vision-and-reasoning-vs-gpt-5-6
title: "meta-muse-spark-1-3-tested-coding-vision-and-reasoning-vs-gpt-5-6"
domain: mindstudio
role: reference
task: reference
actors: ["AWS", "Anthropic", "Meta", "OpenAI"]
dates: []
keywords: ["muse", "reasoning", "agentic", "agents", "aws", "benchmark", "benchmarks", "context window", "cost", "distribution", "gpt-5.6", "multimodal"]
source: docs/RAG/clean_en/mindstudio/meta-muse-spark-1-3-tested-coding-vision-and-reasoning-vs-gpt-5-6.md
source_anchor: ""
source_lines: [1, 81]
sha256: 1195e97cbc63ae6969d37e7d5f407c09650eb302ed25db1b33bf4f0a7efafde4
---

# meta-muse-spark-1-3-tested-coding-vision-and-reasoning-vs-gpt-5-6

<!-- source: https://www.mindstudio.ai/blog/meta-muse-spark-1-3-open-weight -->

## What is Meta Muse Spark 1.3?

Muse Spark 1.3 is Meta’s newest multimodal reasoning model, built for long horizon agentic work, coding, and computer use, with a 1 million token context window. Meta priced it at $1.25 per million input tokens and $4.25 per million output tokens, a rate low enough that it undercuts most frontier competitors by a wide margin. Meta founder Mark Zuckerberg has confirmed the model will be released as open weight soon, which would put a frontier-class model in the hands of anyone willing to self-host it.

## TL;DR

- Meta’s **Muse Spark 1.3** targets frontier-level agentic and coding work with a full 1 million token context window and pricing of $1.25 in / $4.25 out per million tokens.
- The model is set to become **open weight soon** , confirmed by Mark Zuckerberg, which would make it one of the few open-weight models competing directly with closed frontier systems.
- In a hands-on test, the model **autonomously provisioned AWS infrastructure** (an S3 bucket and CloudFront distribution) from a single prompt and returned a working live URL with no step-by-step guidance.
- On a **long-context retrieval benchmark** , Muse Spark scored in the high 90s while GPT-5.6 and Opus 5 fell into the 60s and 70s, with Opus not posting a score at all.
- The model showed strong **situational awareness in a vision test** , correctly parsing a WhatsApp thread and catching a joke buried in a miscommunication about a boss’s wife.
- It handled a **combined chemistry and math reasoning problem** (buffer pH after adding strong acid) with a fully correct, well-justified chain of reasoning.
- A **multilingual test across 79 languages** produced accurate native-script answers and culturally specific picks, like soju for Korean and lassi for Punjabi, all for about $2.49 in total API cost.

## Remy doesn't build the plumbing. It inherits it.

Remy ships with all of it from MindStudio — so every cycle goes into the app you actually want.

## How does Muse Spark 1.3 perform on agentic and coding benchmarks?

According to benchmark charts referenced during testing, Muse Spark tops the field on professional tool use and agentic computer use, edging out both GPT-5.6 and Opus 5. On harder end-to-end business workflows, it runs essentially neck and neck with Opus 5. Coding is described as the closest race of all: Muse Spark, GPT-5.6, and Opus 5 cluster tightly together on long horizon agentic coding tasks, without a clear winner.

The standout number is long context retrieval. On the million-token retrieval test, Muse Spark sits in the high 90s while GPT-5.6 and Opus 5 fade into the 60s and 70s, and Opus reportedly doesn’t post a score at all. That gap matters more than it looks. Agentic workflows that run long (multi-step cloud deployments, large codebases, extended conversations) depend on a model staying coherent over huge amounts of context without losing track of earlier instructions. A model that holds up at native long-context retrieval avoids the silent failure mode where it “forgets” something buried deep in the prompt.

It’s worth noting Opus 5 isn’t uniformly behind. It reportedly wins on knowledge work and some agentic roles, so this isn’t a clean sweep for Meta. The picture that emerges is a model that trades blows with GPT-5.6 and Opus 5 on most fronts, but pulls distinctly ahead on tasks that require staying coherent across very long stretches of context.

## Can it actually deploy infrastructure on its own?

A real test of the model’s agentic claims involved handing it a live AWS account with no existing resources and one instruction: build a self-contained animated website and deploy it globally using an S3 bucket for storage and CloudFront for content delivery. No CloudFront distribution existed beforehand, and no step-by-step guidance was given. The model had to plan the entire sequence itself: provision the S3 bucket, write the HTML file, create the CloudFront distribution, tie it to the S3 origin, and return a working URL.

The reasoning traces showed the model working on multiple independent tasks in parallel, drafting the HTML animation while simultaneously provisioning the cloud infrastructure, then combining both once each piece was ready. Despite heavy server load and repeated throttling during testing, the model eventually completed the task end to end: it created the S3 bucket, correctly linked it as the CloudFront origin, and produced a live, working URL. The result was an animated rotisserie chicken simulation, complete with flickering flames, a rotating spit, a live fire cam overlay, and a running counter for cook time and internal temperature, all generated from a single prompt with zero manual intervention.

## Is Muse Spark 1.3 good at vision and situational reasoning?

The vision test involved a WhatsApp screenshot thread where an employee tells his boss about a work deployment, then unintentionally causes friction by mentioning the boss’s wife. The task for the model was to read the full conversation and determine whether the employee was still on for a planned evening event, correctly interpreting the double meaning behind the exchange.

## Other agents start typing. Remy starts asking.

Scoping, trade-offs, edge cases — the real work. Before a line of code.

The model reportedly got every message right: it correctly attributed messages to the right speaker based on bubble position, recognized that the wife had only seen part of the conversation without the surrounding context, and understood the joke underlying the boss’s frustrated reaction. It concluded correctly that, work-wise, the employee was still on for the evening. This kind of test isn’t a typical benchmark category, but it’s a good proxy for how a model handles ambiguous, socially loaded real-world content rather than clean synthetic prompts.

## How does it handle scientific and multilingual reasoning?

The scientific reasoning test combined chemistry and math: a buffer system with two acid dissociation stages, followed by the addition of a strong acid, with the model asked to calculate the exact final pH. This requires identifying which reaction dominates and carrying that assumption through multi-step math without introducing an error early that throws off the final answer.

The model correctly identified that the added acid would protonate the stronger conjugate base first, justified ignoring the second equilibrium stage based on the roughly 5.55 pKa gap between the two stages, and then verified that formation of the fully protonated species was negligible. That’s a coherent, self-checking chain of reasoning rather than a single guessed number.

The multilingual test asked the model to name a popular drink associated with speakers of 79 different languages, using the correct native script for each. The model covered all 79 languages, using correct scripts across Arabic, Cyrillic, and Devanagari, among others, and made culturally specific choices: soju for Korean, rakija for Serbo-Croatian, kumis for Kazakh, kava for Hawaiian, lassi for Punjabi, shaqaati for Somali, white wine for Basque, and rice water for Malagasy. Most answers defaulted sensibly to coffee or tea where no distinctive traditional drink applied. The total cost for all the tests combined, including the cloud deployment task, came to roughly $2.49.

## Is Muse Spark 1.3 worth using once it’s open weight?

Based on the tests here, the model holds up well outside of benchmark charts. It completed a genuinely autonomous cloud deployment task without hand-holding, showed strong situational and social reasoning in a vision test, produced a rigorous chemistry answer, and handled a large multilingual task with cultural nuance. Combined with aggressive pricing and a pending open-weight release, it’s positioned as a serious option for anyone building agentic tools or long-context applications who wants an alternative to closed models like GPT-5.6 and Opus 5. The one caveat from testing was infrastructure related rather than model related: heavy server load caused repeated throttling, which suggests demand may currently be outpacing capacity.

## Frequently Asked Questions

### What is Meta Muse Spark 1.3 used for?

It’s built for long horizon agentic tasks: multi-step coding, autonomous computer use, and workflows that require reasoning across very large amounts of context, such as a full codebase or an extended multi-turn task.

### Will Muse Spark 1.3 be open weight?

Yes. Meta founder Mark Zuckerberg confirmed on X that Muse Spark 1.3 will be released as open weight soon, following its initial hosted release.

### How much does Muse Spark 1.3 cost to use?

Meta has priced it at $1.25 per million input tokens and $4.25 per million output tokens, notably cheaper than many comparable frontier models.

### How does Muse Spark 1.3 compare to GPT-5.6 and Opus 5?

## Seven tools to build an app. Or just Remy.

Editor, preview, AI agents, deploy — all in one tab. Nothing to install.

It’s roughly matched with both on coding and most agentic benchmarks, but pulls ahead significantly on long-context retrieval (scoring in the high 90s versus the 60s and 70s for the others). Opus 5 reportedly still leads on some knowledge work tasks.

### Can Muse Spark 1.3 actually complete real-world agentic tasks unsupervised?

In testing, it independently provisioned AWS resources (S3 and CloudFront), wrote a working animated web page, and deployed it to a live URL from a single prompt, without step-by-step instructions.
