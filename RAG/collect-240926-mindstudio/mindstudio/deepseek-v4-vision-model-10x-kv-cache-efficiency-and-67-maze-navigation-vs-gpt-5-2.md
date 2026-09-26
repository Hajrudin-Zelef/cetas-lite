---
id: collect-240926-mindstudio/mindstudio/deepseek-v4-vision-model-10x-kv-cache-efficiency-and-67-maze-navigation-vs-gpt-5-2
title: "deepseek-v4-vision-model-10x-kv-cache-efficiency-and-67-maze-navigation-vs-gpt-5"
domain: mindstudio
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "DeepSeek", "OpenAI", "OpenRouter"]
dates: []
keywords: ["deepseek", "agent", "agents", "claude", "cost", "distillation", "fine-tuning", "gpu", "opus 4", "qwen", "reasoning"]
source: docs/RAG/clean_en/mindstudio/deepseek-v4-vision-model-10x-kv-cache-efficiency-and-67-maze-navigation-vs-gpt-5.md
source_anchor: ""
source_lines: [67, 122]
sha256: d1f7e2138ed9c54cd20577e6505a586787a3cf32420f10e1b2168bbe5f6f0cfe
---

# deepseek-v4-vision-model-10x-kv-cache-efficiency-and-67-maze-navigation-vs-gpt-5

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
