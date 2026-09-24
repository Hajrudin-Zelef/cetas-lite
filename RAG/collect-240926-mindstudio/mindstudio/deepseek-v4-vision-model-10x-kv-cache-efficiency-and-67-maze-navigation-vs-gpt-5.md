---
id: collect-240926-mindstudio/mindstudio/deepseek-v4-vision-model-10x-kv-cache-efficiency-and-67-maze-navigation-vs-gpt-5
title: "deepseek-v4-vision-model-10x-kv-cache-efficiency-and-67-maze-navigation-vs-gpt-5"
domain: mindstudio
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "DeepSeek", "Google", "OpenAI", "OpenRouter"]
dates: ["2024-03", "2024-10", "2024-12", "2025-10"]
keywords: ["deepseek", "agent", "agents", "attention", "benchmark", "benchmarks", "claude", "cost", "distillation", "fine-tuning", "gemini", "gpu"]
source: docs/RAG/clean_en/mindstudio/deepseek-v4-vision-model-10x-kv-cache-efficiency-and-67-maze-navigation-vs-gpt-5.md
source_anchor: ""
source_lines: [1, 122]
sha256: 639afb8165f493e04aab225fbe93487ab2730378233416d0d5d2ea8fe00a20e5
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

Si vous exploitez un pipeline qui traite des milliers d’images — extraction de documents, questions-réponses visuelles sur des catalogues de produits, analyse de captures d’écran pour des agents — la différence entre 90 et 870 entrées de cache par image est la différence entre faire tenir 9× plus de requêtes dans le même budget mémoire ou avoir besoin de 9× moins de GPU pour atteindre le même débit.

DeepSeek V4 Flash est déjà tarifé à 1,74 $/M de tokens d’entrée et 3,48 $/M de tokens de sortie. Comparez cela à Claude Opus 4.7 à 5 $/M en entrée et 25 $/M en sortie, ou GPT-5.5 à 5 $/M en entrée et 30 $/M en sortie. Le modèle de vision hérite de cette tarification, et ensuite l’efficacité du cache KV multiplie davantage l’avantage de coût effectif lorsque vous servez des images en volume.

Des plateformes comme MindStudio qui prennent en charge plus de 200 modèles et vous permettent d’acheminer les requêtes entre fournisseurs rendent ce type d’arbitrage de coûts pratique sans écrire de logique de routage personnalisée — vous pouvez échanger le backend de vision et mesurer la différence de qualité sans reconstruire votre pipeline.

## La décision d’architecture à surveiller

Le pipeline d’entraînement du modèle de vision mérite un examen plus attentif car il est inhabituel.

### Tout le monde a construit un ouvrier du bâtiment.

Nous avons construit l’entrepreneur.

    Un fichier à la fois.

UI, API, base de données, déploiement.

DeepSeek a d’abord entraîné deux modèles spécialistes distincts : un pour « penser avec ancrage » en utilisant des boîtes englobantes, un pour « penser avec pointage » en utilisant des coordonnées de points. Chacun a eu sa propre passe de fine-tuning supervisé, puis son propre apprentissage par renforcement avec GRPO utilisant trois têtes de récompense (format, qualité et précision). Ensuite, une étape de distillation unifiée sans récompense les a fusionnés en un seul modèle étudiant.

Le résultat est un modèle qui peut faire à la fois du raisonnement visuel basé sur des boîtes et basé sur des points, mais qui a été entraîné comme deux spécialistes avant d’être consolidé. C’est une approche d’entraînement plus coûteuse que d’entraîner un seul modèle de bout en bout, mais cela évite le compromis qui découle de l’entraînement d’un seul modèle à faire deux choses différentes simultanément.

La même logique — découpler les parties difficiles, entraîner des spécialistes, puis fusionner — est apparue dans Janus avec la conception à double encodeur. Cela devient un schéma dans la façon dont DeepSeek aborde l’architecture multimodale.

C’est aussi un contexte pertinent pour quiconque construit des agents à forte composante visuelle. Le modèle n’est pas un modèle de vision polyvalent qui se trouve faire du raisonnement spatial. C’est un modèle de raisonnement spatial qui a été explicitement entraîné à ancrer sa chaîne de pensée dans des coordonnées. Si votre cas d’usage implique du comptage, du traçage de chemins, ou toute tâche où « de quelle chose spécifique est-ce que je parle » est la partie difficile, c’est la cible de conception.

Pour les équipes évaluant des alternatives à poids ouverts pour des tâches de vision, la comparaison Gemma 4 vs Qwen 3.5 à poids ouverts couvre les autres principaux concurrents dans l’espace à poids ouverts sous-frontière — aucun des deux n’a publié de chiffres comparables d’efficacité du cache KV.

## Le statut de déploiement

Depuis fin avril 2025, DeepSeek a commencé à déployer le mode vision dans l’application et sur le web, au moins en test limité aux côtés de leurs modes rapide et expert. L’article lui-même est devenu difficile à trouver peu après sa publication — il est apparu brièvement puis a été retiré, ce qui est inhabituel.

Le modèle derrière l’article semble être en déploiement progressif. L’ID de modèle OpenRouter pour DeepSeek V4 Flash est `deepseek/deepseek-v4-flash`, qui est la colonne vertébrale textuelle. La variante vision n’est pas encore listée séparément sur OpenRouter à l’heure d’écrire ces lignes, mais l’architecture sous-jacente est la même base V4 Flash.

Si vous construisez quelque chose qui doit passer en production avant que la variante vision ne soit largement disponible, le modèle textuel est accessible maintenant. Les capacités de vision sont la couche au-dessus.

## Ce qu’il faut surveiller

La première chose à tester, une fois que le modèle de vision sera plus largement déployé, est de savoir si l’efficacité du cache KV tient à des résolutions plus élevées. L’article rapporte environ 90 entrées pour 80×80. L’architecture prend en charge une résolution arbitraire, mais les taux de compression dépendent de la résolution. Une image de 1024×1024 produira plus de tokens de patch avant compression, et la taille finale du cache sera plus grande — la question est de savoir si le taux de compression reste à peu près constant ou se dégrade.

La deuxième chose à surveiller est la limitation « le mode de primitives visuelles doit être déclenché explicitement ». Pour l’instant, le modèle ne décide pas automatiquement quand utiliser le raisonnement par boîtes englobantes plutôt que le traitement visuel standard. C’est une lacune d’utilisabilité pour les cas d’usage d’agents où vous voulez que le modèle auto-sélectionne le bon mode de raisonnement. Quand cela deviendra automatique, les chiffres de navigation dans les labyrinthes s’amélioreront probablement encore.

## 
      Plans first.
      *Then code.*
    

    Remy writes the spec, manages the build, and ships the app.

For anyone building vision agents specifically, the GPT-5.4 Mini vs Claude Haiku 4.5 sub-agent comparison is a useful reference for thinking about where a highly efficient vision model like this fits in a multi-tier agent architecture — the cost profile of V4 Flash makes it a plausible sub-agent for visual tasks even when you’re using a more capable model as the orchestrator.

The broader question the DeepSeek OCR paper raised — whether pixels are better inputs than tokens for certain tasks — is now being answered in practice. A model that represents an image in 90 KV-cache entries and outperforms GPT-5.4 on spatial reasoning is a concrete data point in that argument. The tokenizer may not be going anywhere soon, but the case for pixel-native representations is getting harder to dismiss.

If you’re building applications where the spec is “process images, reason about spatial relationships, do it cheaply at scale,” the architecture described in *Thinking with Visual Primitives* is worth understanding in detail. Tools like Remy take a different approach to the spec-to-production problem — you write annotated markdown and compile a full TypeScript stack from it — but the underlying principle is similar: the right representation at the right level of abstraction produces better results than forcing everything through a single bottleneck.

The DeepSeek vision team has been asking the same question for two years. The answer keeps getting more interesting.
