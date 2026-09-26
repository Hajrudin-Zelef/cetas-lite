---
id: collect-240926-mindstudio/mindstudio/deepseek-v4-vision-model-10x-kv-cache-efficiency-and-67-maze-navigation-vs-gpt-5-1
title: "deepseek-v4-vision-model-10x-kv-cache-efficiency-and-67-maze-navigation-vs-gpt-5"
domain: mindstudio
role: reference
task: reference
actors: ["Anthropic", "DeepSeek", "Google", "OpenAI"]
dates: ["2024-03", "2024-10", "2024-12", "2025-10"]
keywords: ["deepseek", "agents", "attention", "benchmark", "benchmarks", "claude", "gemini", "gpu", "inference", "kv cache", "memory", "multimodal"]
source: docs/RAG/clean_en/mindstudio/deepseek-v4-vision-model-10x-kv-cache-efficiency-and-67-maze-navigation-vs-gpt-5.md
source_anchor: ""
source_lines: [1, 66]
sha256: a28f4fbcba65db58a4028095ce7efb544ab06c92accbb06fefc55b9da4d81627
---

# deepseek-v4-vision-model-10x-kv-cache-efficiency-and-67-maze-navigation-vs-gpt-5

<!-- source: https://www.mindstudio.ai/blog/deepseek-v4-vision-model-kv-cache-efficiency-maze-navigation-benchmark -->

## DeepSeek’s Vision Model Uses 90 KV-Cache Entries Where Sonnet Uses 870

DeepSeek’s vision variant — built on the V4 Flash backbone — processes an 80×80 image using approximately 90 KV-cache entries. Claude Sonnet 4.6 uses around 870 for the same image. That’s a 10x difference in memory footprint per image, and it’s not an accident. It’s the result of a deliberate multi-stage compression pipeline that DeepSeek has been quietly building toward for two years.

If you’re building vision-heavy pipelines and you’re watching inference costs, this is the number that matters.

The maze navigation benchmark makes the efficiency story even stranger: DeepSeek’s vision model scores 67% on topological reasoning tasks, against GPT-5.4’s 50% and Gemini Flash 3’s 49%. A model that costs a tenth as much to run on images is also outperforming frontier models on the class of spatial reasoning tasks where visual grounding matters most.

## How DeepSeek Gets to 90 Cache Entries

The compression isn’t a single trick. It’s a pipeline of four stages, each multiplying the reduction.

Start with a 756×756 image. That’s 571,000 pixels. The DeepSeek vision transformer — which they call a “DeepSeek Vision Transformer” and built from scratch to support arbitrary resolution — processes this using 14×4 patches. That initial patch tokenization produces around 2,916 patch tokens.

Then a 3×3 spatial compression step runs along the channel dimension, collapsing nine adjacent patches into one. That brings the token count down to roughly 324.

Then the compressed sparse attention mechanism from the V4 paper applies another 4× compression to the KV cache.

End result: approximately 81 entries in the KV cache for a full image. The paper describes this as roughly a 7,000× total compression ratio from raw pixels to KV-cache entries.

The language backbone underneath all of this is DeepSeek V4 Flash — a 284B parameter mixture-of-experts model with 13B active parameters at inference. You’re getting a model that reasons at frontier quality but only activates 13B parameters per forward pass, combined with a vision encoder that represents each image in a fraction of the memory that competing models require.

## The “Thinking With Visual Primitives” Paper

The vision model isn’t just about efficiency. The paper — titled *Thinking with Visual Primitives* — argues that current multimodal models have two distinct gaps, not one.

The first is the **perception gap**: models can’t always see fine-grained detail. Most of the 2024 work on high-resolution cropping and dynamic patching was aimed at this.

The second is what the paper calls the **reference gap**: even when a model sees an image correctly, natural language is too imprecise to point at things reliably. If you ask a model to count the third bear from the left on a rocky ledge, it can describe what it sees, but it loses track of which entity it’s actually referring to as its reasoning chain gets longer. Humans solve this with a finger. Models, until now, didn’t have an equivalent.

DeepSeek’s solution is to make spatial coordinates first-class tokens in the chain of thought. When the model reasons about an image, it emits bounding boxes inline — a reference tag with a label, followed by a box tag with two corner coordinates. These are special tokens in the model’s vocabulary, not function calls, not a separate tool. The model literally writes `<ref>person_3</ref><box>(x1,y1),(x2,y2)</box>` mid-thought, then continues reasoning with that anchor in place.

This is why the maze navigation benchmark matters. Maze path tracing and topological reasoning are exactly the tasks where language is uniquely bad at trajectory description. When the model can point to a cell, mark it, and reason forward from that mark, it doesn’t lose its place. The 67% vs 50% gap over GPT-5.4 isn’t surprising once you understand the mechanism — it’s the expected result of having a reference primitive that GPT-5.4 lacks.

## Two Years of the Same Story

The paper didn’t come out of nowhere. DeepSeek has shipped roughly seven vision-related models in 24 months, and the through-line across all of them is the same question: *what’s the cheapest representation that still works?*

In March 2024, DeepSeek VL used a hybrid SigLIP and SAM encoder. Nothing flashy, but it set the foundation. In October 2024, Janus decoupled the visual encoder for understanding versus generation — most unified multimodal models at the time had a single encoder bottleneck, and Janus said no to that. In December 2024, the VL2 model ported mixture-of-experts and multi-head latent attention from V2 and V3 into vision. A 1B activated-parameter version was scoring 80.9 on OCR bench.

## Seven tools to build an app. Or just Remy.

Editor, preview, AI agents, deploy — all in one tab. Nothing to install.

Then in October 2025, the DeepSeek OCR paper landed. The framing was strange — they called it an OCR paper, but the actual idea was: take 1,000 text tokens, render them as an image, encode the image, and get back 100 vision tokens that reconstruct the original text at 97% accuracy. That’s 10× compression on long context. Andrej Karpathy’s reaction was: *“the tokenizer must go, pixels may be better inputs to language models than text.”* That quote spread fast, and it’s what put DeepSeek’s vision team on a lot of people’s radar.

The *Thinking with Visual Primitives* paper is the next chapter in that same story. Each release has been asking the same question and finding a more aggressive answer.

## What the Benchmarks Actually Claim (and Don’t)

The paper is honest about scope in a way that a lot of coverage will skip. There’s a footnote that says the reported scores cover only a subset of evaluation dimensions directly relevant to the research focus, and are therefore not indicative of the model’s overall capabilities.

DeepSeek n’affirme pas que cela dépasse GPT-5.4 sur tous les plans. Ils affirment que cela dépasse GPT-5.4 sur les tâches de raisonnement visuellement ancrées — navigation dans des labyrinthes, traçage de chemins, comptage dans des scènes denses. C’est une affirmation plus étroite et plus défendable.

Sur les questions de comptage brut, Gemini Flash 3 est encore devant. Sur les benchmarks de vision générale, l’article ne fait pas d’affirmations générales. Les trois limites qu’ils signalent explicitement : le modèle est limité par la résolution (les scènes fines peuvent encore le faire échouer), le mode de primitives visuelles doit être déclenché explicitement plutôt que d’être auto-sélectionné, et le raisonnement topologique basé sur les points ne se généralise pas bien à tous les scénarios.

Ce genre d’honnêteté mérite d’être salué. L’article vous indique où le modèle fonctionne et où il ne fonctionne pas. C’est plus utile qu’une capture d’écran de classement.

Pour comparaison, si vous évaluez où cela se situe par rapport aux autres modèles de pointe sur des tâches générales, la comparaison GPT-5.4 vs Claude Opus 4.6 couvre le paysage plus large des capacités — DeepSeek V4 Flash se situe en dessous des deux sur les benchmarks généraux mais les surclasse significativement en termes de coût.

## Pourquoi l’écart d’efficacité a des conséquences pratiques

La réduction de 10× du cache KV n’est pas seulement une curiosité de benchmark. La taille du cache KV affecte directement combien de requêtes d’images simultanées vous pouvez servir sur un GPU donné, combien de mémoire vous avez besoin pour traiter des lots de requêtes visuelles, et donc à quoi ressemble votre coût réel par image à grande échelle.

