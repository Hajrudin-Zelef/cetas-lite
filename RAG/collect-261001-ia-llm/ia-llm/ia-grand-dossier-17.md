---
id: collect-261001-ia-llm/ia-llm/ia-grand-dossier-17
title: "IA — Le grand dossier"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "DeepSeek", "Google", "Meta", "Mistral", "OpenAI", "vLLM", "xAI"]
dates: []
keywords: ["agent", "agents", "attention", "compute", "deepseek", "diffusion", "distillation", "embeddings", "fine-tuning", "gpu", "memory", "mistral"]
source: docs/RAG/collect-261001-ia-llm/ia_grand_dossier.md
source_anchor: ""
source_lines: [1081, 1153]
sha256: 829cdb0485b8421380148733d8cbdb1786e6636839cd0c1a4a9ad188a5e91fed
---

# IA — Le grand dossier

**Optimisations (dans l'ordre d'impact) :**
1. **Routage** : 80 % des alarmes sont triviales → petit modèle sans thinking (0,15 $/1M) ; 20 % au raisonnant. Économie : ~70 %.
2. **Distillé local** : un R1-Distill-14B sur RTX 4090 pour le premier tri, API pour les cas durs. Économie : ~85 %.
3. **Plafond de thinking** : `reasoning_effort=low` par défaut, `high` sur escalade. Économie : ~40 % sur la part raisonnante.
4. **Cache de préfixe** : le contexte système/consignes est identique → facturé une fois (remise ~50-90 % selon fournisseur sur les tokens cachés).

**Leçon :** le test-time compute est un **budget**, pas une fatalité. Sans garde-fous, un agent en boucle coûte 100× une requête simple (voir 6.6, risque n° 4).

### 14.3. Scénario 3 — Choisir la taille du modèle : le test A/B qui tranche

**Protocole (1 semaine, coût ~0) :**
1. Constituer **100 questions réelles** du métier avec les **bonnes réponses** (gold set) — extraites des guides existants.
2. Faire tourner le RAG avec **3 modèles** : 8B Q4 local, 32B Q4 local, API modèle moyen.
3. Mesurer : **exactitude** (réponse correcte ?), **latence P50/P99**, **coût/1 000 questions**.

**Résultat typique observé dans l'industrie (2024-2026) :** sur un RAG bien construit (bon chunking + reranking), l'écart 8B → 32B → API est de **quelques points** d'exactitude (ex. : 82 % → 85 % → 87 % — ordres de grandeur illustratifs), mais le coût est de **1× → 4× → 20×**. **La décision rationnelle est presque toujours : améliorer le retrieval, garder le petit modèle.** C'est la traduction opérationnelle de « l'efficacité bat la taille » (section 4).

### 14.4. Scénario 4 — Dimensionner électriquement une baie IA (lien avec le métier énergies)

**Besoin :** 4 serveurs × 8 GPU H100 (32 GPU) pour du serving interne.

| Poste | Calcul | Résultat |
|---|---|---|
| GPU | 32 × 700 W (TDP H100) | 22,4 kW |
| Serveurs (CPU, RAM, ventilos, alim) | 4 × ~2 kW | 8 kW |
| Réseau (switches IB 400G) | ~2 kW | 2 kW |
| **Total IT** | | **~32 kW** |
| Refroidissement (PUE 1,3) | 32 × 0,3 | ~10 kW |
| **Total à prévoir** | | **~42 kW** |

**Conséquences :** on ne branche pas 42 kW sur un tableau de bureaux — c'est un **départ TGBT dédié**, disjoncteurs courbe adaptée aux appels de courant (les GPU font des transitoires de 100 W → 700 W en ms), **onduleur 60 kVA minimum en N+1** (soit 2× 60 kVA ou modulaire), et groupe électrogène si la charge est critique. **La méthode est exactement celle du guide onduleurs de Zelef** (dimensionnement 10/20/40/60 kVA) — l'IA n'invente rien en électrotechnique, elle ajoute juste des MW au problème.

### 14.5. Scénario 5 — Politique d'usage IA pour une équipe (modèle de charte)

**Principes (à adapter) :**
1. **Classification des données** : public / interne / confidentiel. Le confidentiel ne va **jamais** sur une API sans DPA + zéro-rétention ; par défaut, il va sur le modèle **local**.
2. **Routage par défaut** : tâche simple → modèle local/économique ; tâche complexe → modèle puissant **avec validation humaine** ; agent autonome → **budget tokens + timeout + superviseur humain**.
3. **Interdit sans revue** : agent qui exécute des actions en production (redémarrer un service, modifier un firewall) sans validation — risque **prompt injection** via le corpus RAG (voir 6.6).
4. **FinOps** : tableau de bord mensuel **tokens et € par équipe/cas d'usage** ; alerte à +50 % de la baseline ; revue trimestrielle du routage.
5. **Évaluation continue** : le gold set de 100 questions (14.3) est rejoué à chaque changement de modèle/prompt — **on ne change jamais de modèle sans mesurer**.

---

## Conclusion de la partie 1

De 2012 à 2026, l'IA a connu **six révolutions** (deep learning, Transformer, LLM few-shot, diffusion, RLHF, raisonnement/test-time compute), portées par des **chercheurs** (Hinton, LeCun, Bengio, l'équipe Vaswani, Karpathy, Sutskever...), des **fondateurs** (Altman, Amodei, Hassabis, Musk, Liang Wenfeng, Mensch, Huang) et des **labos** aux stratégies divergentes (maximalisme fermé d'OpenAI, confiance d'Anthropic, intégration de Google, ouverture de Meta, efficacité de DeepSeek, souveraineté de Mistral).

Le débat central — **la fin du scaling naïf du training** — se tranche ainsi en septembre 2026 : **le pré-entraînement massif montre des rendements décroissants** (mur des données, coûts, GPT-4.5/5 en illustration), mais **le progrès a changé d'axe** (test-time compute, RL avec récompenses vérifiables, MoE, distillation, quantification, agents). Les deux camps ont partiellement raison ; ils ne parlent simplement plus du même objet.

Pour le sysadmin, la conséquence est claire : **l'ère du training héroïque est finie, l'ère de l'inférence industrielle commence**. Le métier se déplace vers le serving (vLLM), le FinOps par token, la sécurité des agents, l'évaluation sur données métier — et, spécifiquement pour un chef de service systèmes & énergies, vers **l'électrotechnique des baies IA**. Le RAG d'entreprise n'a jamais été aussi accessible : un 8B quantifié sur une carte gamer, un corpus bien nettoyé, un gold set de 100 questions — et on bat des solutions à 2 000 $/mois d'API.

*Suite : partie 2 — architectures en profondeur (Transformer décortiqué, embeddings, chunking, vector DB, reranking, fine-tuning, évaluation, sécurité du RAG).*

---

## 1.14. Les chaînons manquants : ce que le récit principal a sauté (mais qu'il faut connaître)

> Entre les grandes révolutions, des avancées moins célèbres ont rendu la suite possible. Les voici, pour ne jamais être perdu dans une discussion technique.

**1997 — LSTM (Hochreiter & Schmidhuber).** Le Long Short-Term Memory résout (partiellement) le problème de la mémoire longue des RNN grâce à des « portes » (gates) qui décident quoi garder/oublier. Il domine le NLP de 1997 à 2017 — vingt ans de règne. Tout système de traduction/reconnaissance vocale pré-Transformer en utilise.

**2014 — seq2seq (Sutskever, Vinyals, Le).** Architecture encodeur-décodeur : un RNN lit toute la phrase source, un autre génère la traduction. C'est le **moule** dont le Transformer hérite (encodeur + décodeur), en remplaçant juste les RNN par de l'attention.

**2014-2015 — L'attention avant le Transformer (Bahdanau et al., 2014 ; Luong et al., 2015).** L'idée « le décodeur regarde les bons mots de la source » existe déjà greffée sur des RNN. Le papier de 2017 ne fait « que » la généraliser en architecture complète — ce « que » valant une révolution.

**2019 — T5 (Raffel et al., Google).** « Text-to-Text Transfer Transformer » : **tout** est formulé comme de la génération de texte (traduction, résumé, classification = « traduire vers le label »). Philosophie reprise par GPT : l'uniformité des tâches simplifie le scaling.

**2020 — ViT (Dosovitskiy et al., Google).** Vision Transformer : on découpe l'image en **patchs** traités comme des « mots ». Un Transformer pur bat les CNN en vision **à condition d'avoir assez de données** — preuve que l'architecture est vraiment universelle, le biais inductif des convolutions devenant inutile à grande échelle.

**2020 — Les lois de scaling multimodales.** Kaplan/Chinchilla portent sur le texte, mais les mêmes lois de puissance s'observent en vision, audio, et multimodal — c'est ce qui autorise les labos à dimensionner les runs Sora/Veo comme des runs LLM.

**2021 — CLIP (Radford et al., OpenAI).** Apprentissage contrastif texte-image sur 400M de paires : l'encodeur qui permet à DALL-E 2 puis à Stable Diffusion de « comprendre » les prompts. Sans CLIP, pas de génération texte→image moderne.

