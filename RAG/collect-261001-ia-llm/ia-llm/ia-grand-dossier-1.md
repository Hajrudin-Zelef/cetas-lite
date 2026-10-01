---
id: collect-261001-ia-llm/ia-llm/ia-grand-dossier-1
title: "IA — Le grand dossier"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "Nvidia", "OpenRouter", "United States", "vLLM"]
dates: ["2026-09-27"]
keywords: ["agents", "benchmarks", "claude", "compute", "diffusion", "distillation", "embeddings", "gpu", "kv cache", "llama", "llama.cpp", "mcp"]
source: docs/RAG/collect-261001-ia-llm/ia_grand_dossier.md
source_anchor: ""
source_lines: [1, 97]
sha256: c77c6492dbc4f79440490641c6bea04a1c41843a85d9272081e37b33e7ecf90f
---

# IA — Le grand dossier

*Document unique assemblé le 27 septembre 2026 pour Zelef — chef de service systèmes & énergies, sysadmin, construit son application RAG personnelle.*

Trois parties : **1.** l'histoire et le débat sur le scaling · **2.** les providers, le routage et l'IA locale · **3.** le débat pour/contre, les dangers, Claude, les stacks et les communautés. Glossaire et quiz fusionnés en fin de document.

## Table des matières

- [Partie 1 — Histoire, révolutions, scaling : du deep learning à la fin du training](#partie-1--histoire-révolutions-scaling)
- [Partie 2 — Providers, passerelles/routage, IA locale](#partie-2--providers-passerellesroutage-ia-locale)
- [Partie 3 — Débats, dangers, Claude, stacks, communautés](#partie-3--débats-dangers-claude-stacks-communautés)
- [Glossaire général](#glossaire-général)
- [Quiz général — 30 questions corrigées](#quiz-général--30-questions-corrigées)

## Comment utiliser ce dossier

Ce dossier fait 6 000+ lignes : ne le lisez pas d'une traite. Trois parcours selon votre objectif :

**Parcours A — Sysadmin pressé (2 h).** Partie 1, sections 4 (le débat scaling) et 6 (ce que ça change pour un sysadmin) → Partie 2, sections 5-8 (IA locale, hardware chiffré, local vs cloud) → Partie 3, section 6 (20 conseils). Vous repartez avec : une position sur le scaling, un dimensionnement VRAM, et un plan 90 jours.

**Parcours B — Builder RAG (1 journée).** Partie 2 en entier (providers, routage LiteLLM/OpenRouter, IA locale) → Partie 3, section 4.4 (stack RAG chaînée avec code) → Partie 1, section 4.4 (pourquoi la distillation rend le RAG local viable). Compléments naturels : `ia_modeles_specialises.md` (embeddings, rerankers, KV cache) et `outils_dev_rag.md` (chunk & corpus avec code Python).

**Parcours C — Culture IA complète (1 semaine).** Dans l'ordre, avec le quiz de chaque partie avant de passer à la suivante. Le glossaire général (122 termes) sert de référence pendant la lecture.

**Règle d'hygiène.** Les prix API, les versions de modèles et les chiffres réglementaires sont datés du **27 septembre 2026** et bougent vite. Tout fait marqué « à vérifier » doit être recoupé avant usage en production ou en réunion. Les opinions sont attribuées à leurs auteurs : ce dossier présente les positions, il n'en choisit aucune.

## Journal de maintenance du dossier

L'IA bouge par trimestres. Pour garder ce dossier utile, revérifiez à chaque trimestre :

1. **Prix et modèles** — tableaux de la Partie 2 (prix API, nouveaux modèles, providers disparus/fusionnés). Priorité : le Top 50 et les fiches des 12 providers clés.
2. **Régulation** — section 2.4 de la Partie 3 (AI Act : calendrier d'application ; US : patchwork des États ; Chine : nouvelles mesures). Les intitulés de lois changent vite.
3. **Scaling** — section 4 de la Partie 1 : le débat évolue à chaque génération de modèles (nouveaux arguments des deux camps, nouvelles lois d'échelle).
4. **Versions des stacks** — Partie 3, section 4 et Partie 2, section 5 : vLLM, Ollama, llama.cpp, frameworks d'agents (versions majeures = breaking changes possibles sur les exemples de code).
5. **Glossaire** — ajouter les termes apparus dans le trimestre (les buzzwords naissent vite et meurent parfois aussi vite : n'ajouter que ce qui survit 3 mois).

Ajoutez une ligne datée ci-dessous à chaque révision :

| Date | Révisé par | Sections touchées | Notes |
|---|---|---|---|
| 2026-09-27 | assemblage initial | — | v1 : 3 parties + glossaire 122 termes + quiz 30 Q |
| | | | |

**Documents associés** (même bibliothèque) : `ia_modeles_chine.md`, `ia_modeles_occident.md`, `ia_modeles_specialises.md` (les modèles) · `ia_agents_codeurs.md`, `ia_agents_concepts.md`, `mcp_lsp_protocoles.md`, `llm_gateways_infra.md`, `outils_dev_rag.md`, `ia_generatif_media.md` (l'outillage).

## Conventions de lecture

- **« à vérifier »** : fait plausible mais non recoupé au 27/09/2026 — à confirmer avant usage.
- **Prix en USD** sauf mention contraire, datés du 27/09/2026 ; les tarifs API changent sans préavis.
- Les blocs de code sont prêts à adapter, jamais à exécuter tels quels en production sans relecture.
- Les opinions citées sont attribuées à leurs auteurs ; leur présence ici ne vaut pas adhésion.
---

# PARTIE 1 — Histoire, révolutions, scaling : du deep learning à la fin du training
## Histoire, révolutions, scaling : du deep learning à la fin du training

**Destinataire :** Zelef — chef de service systèmes & énergies, sysadmin, construit son application RAG personnelle.
**Périmètre :** cette partie 1 couvre 2012 → 2026 : les révolutions techniques, les personnes, les laboratoires, puis le débat central — le ralentissement du training et la fin (partielle) du scaling naïf.
**Méthode :** tout fait daté, chiffré ou biographique ci-dessous a été vérifié par recherche web préalable (27/09/2026). Les opinions sont attribuées à leurs auteurs. Les chiffres non confirmés par une source sont marqués **« à vérifier »**. Les prédictions sont présentées comme des prédictions, pas comme des faits.

---

## Sommaire de la partie 1

1. Les grandes révolutions IA 2012→2026 (deep learning, transformers, GPT et la course LLM, diffusion, RLHF, agents, modèles de raisonnement, multimodalité)
2. Les grands noms : qui a fait quoi (vérifié)
3. Les grands laboratoires : positionnement de chacun
4. Le ralentissement du training / la fin du scaling naïf : le débat central, les deux camps, les chiffres, les maths, les scénarios
5. Timeline détaillée 2012→2026
6. Ce que ça change concrètement pour un sysadmin (serving, coûts, sécurité, énergie, feuille de route RAG)
7. Sources et méthode de vérification
8. → 24. Compléments : papiers, chiffres du compute, hardware, idées reçues, portraits, voix du débat, cas sysadmin, benchmarks, datasets, lecture de rapports, frise des modèles, FAQ, géopolitique, acronymes, checklist annonces, fiches modèles 2026, erreurs RAG
25. Glossaire (60 termes)
26. Quiz : 10 questions + corrigés (+ 5 bonus) — à la fin

---

## 1. Les grandes révolutions IA, 2012 → 2026

### 1.1. Le point de départ : pourquoi 2012 est l'an zéro

Trois ingrédients convergent autour de 2012 :

| Ingrédient | Détail | Effet |
|---|---|---|
| **Données massives** | ImageNet de Fei-Fei Li (2009) : ~15 millions d'images étiquetées, 22 000 catégories | Enfin assez de données pour entraîner des réseaux profonds sans surapprentissage immédiat |
| **GPU généralistes** | Cartes NVIDIA (GTX 580 à l'époque) programmables en CUDA | 10 à 50× plus rapides que les CPU pour les multiplications matricielles, le cœur des réseaux de neurones |
| **Algorithmes matures** | Rétropropagation (Rumelhart, Hinton, Williams, 1986), ReLU, dropout (Hinton et al., 2012) | Des réseaux de 60 millions de paramètres deviennent entraînables en pratique |

Avant 2012, les réseaux de neurones étaient une niche méprisée : le champ était dominé par les SVM, les arbres de décision et les modèles à traits « faits main ». Après 2012, tout bascule. Comprendre cette date, c'est comprendre que **le deep learning moderne n'est pas né d'une idée neuve, mais de la rencontre d'idées anciennes avec du compute et des données**.

### 1.2. Révolution n° 1 — AlexNet et le deep learning (2012)

**Le fait.** En octobre 2012, Alex Krizhevsky, Ilya Sutskever et leur directeur de thèse Geoffrey Hinton (Université de Toronto) remportent le concours ImageNet (ILSVRC 2012) avec un réseau de neurones convolutionnel (CNN) profond nommé **AlexNet**. Résultat : **15,3 % d'erreur top-5**, contre 26,2 % pour le second. Un écart jamais vu : ils divisent quasiment l'erreur par deux.

**Le papier :** « ImageNet Classification with Deep Convolutional Neural Networks », présenté à NeurIPS 2012.

