---
id: collect-261001-ia-llm/ia-llm/ia-grand-dossier-12
title: "IA — Le grand dossier"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "DeepSeek", "Google", "Lambda", "Mistral", "Nvidia", "OpenAI", "vLLM", "xAI"]
dates: ["2025-01-27", "2026-09-27"]
keywords: ["agent", "agents", "attention", "attribution", "deepseek", "embeddings", "fine-tuning", "gpu", "mistral", "nvidia", "rlhf", "training"]
source: docs/RAG/collect-261001-ia-llm/ia_grand_dossier.md
source_anchor: ""
source_lines: [770, 804]
sha256: 54582443b316d22e969c3c0c7c4c23c6dc088e3c1e60cca0d4469f6a7e121f6a
---

# IA — Le grand dossier

| Mois | Compétence | Livrable |
|---|---|---|
| 1-2 | Serving local : Ollama + un 8B quantifié + RAG pgvector | RAG perso qui tourne sur une machine du service |
| 2-3 | vLLM en production : API OpenAI-compatible, monitoring tokens/latence | Endpoint interne documenté |
| 3-4 | Évaluation : jeux de questions métier, mesure avant/après (chunking, reranking) | Le RAG **mesuré**, pas « senti » |
| 4-5 | Agents : function calling, premier agent (lecture de doc + action) | Un agent qui fait une vraie tâche (ex. : synthèse d'alarmes) |
| 5-6 | FinOps IA : coût/token par cas d'usage, routage petit/gros modèle, budgets | Tableau de bord coûts + politique d'usage |

**Ce qui ne sert plus :** apprendre à entraîner un Transformer from scratch (sauf R&D). **Ce qui monte :** évaluation (evals), data pipelines, sécurité des prompts (injection), et — spécifique au métier de Zelef — **l'énergétique des baies IA**.

### 6.6. Risques à anticiper (le côté « énergies » du dossier)

1. **Dépendance fournisseurs** : un RAG 100 % API est otage des prix et des CGU. Stratégie : **abstraction** (interface OpenAI-compatible) + **modèle local de repli**.
2. **Fuite de données par les prompts** : les documents métier envoyés à une API tierce — politique de classification + DPA + option zéro-rétention, ou on-premise.
3. **Prompt injection** : un document du corpus RAG peut contenir des instructions malveillantes (« ignore les instructions précédentes... ») — **ne jamais laisser un agent RAG exécuter d'action critique sans validation humaine**.
4. **Emballement des coûts agents** : un agent en boucle = facture ×100. **Garde-fous obligatoires** : budget tokens par tâche, timeout, nombre max d'itérations.
5. **Énergie** : l'IA déplace la contrainte du SI vers le **tableau électrique** — dimensionner l'alimentation/ondulation/refroidissement AVANT d'acheter les GPU (cf. guide onduleurs de Zelef : un rack IA se traite comme une charge critique).

---

## 7. Sources et méthode de vérification

Recherches web effectuées le 27/09/2026 via browser.search / browser.open sur : AlexNet/ILSVRC 2012, « Attention Is All You Need » (auteurs, venue), lois de scaling Kaplan/Chinchilla, coûts de training (Lambda Labs, Epoch AI, déclarations Altman), DeepSeek V3/R1 (rapports techniques, analyses Stratechery), o1/Strawberry (annonces OpenAI), RLHF/InstructGPT, prix Turing 2018 (ACM), Nobel 2024 (physique : Hinton/Hopfield ; chimie : Hassabis/Jumper/Baker), fondations OpenAI/Anthropic/DeepMind/Mistral/xAI, crise NVIDIA du 27/01/2025, Epoch AI « data wall », carrières Karpathy/Amodei/Sutskever/Hassabis/Mensch.

**Limites assumées :** les chiffres de coûts non publiés par les labos sont des estimations d'analystes (marqués « à vérifier ») ; les numéros de versions de modèles 2026 évoluent vite (vérifier au jour près avant usage) ; les interprétations du « mur » relèvent du débat d'experts, présentées ici dans les deux camps avec attribution.

*Fin de la partie 1 — Histoire, révolutions, scaling, ralentissement du training.*
*Partie 2 (à venir) : architectures en profondeur — Transformer, embeddings, RAG, fine-tuning, évaluation.*

---

## 8. Les 25 papiers à lire absolument (avec résumé opérationnel)

> Classés par ordre chronologique. « Pourquoi le lire » = ce que ça change pour un praticien/RAG.

