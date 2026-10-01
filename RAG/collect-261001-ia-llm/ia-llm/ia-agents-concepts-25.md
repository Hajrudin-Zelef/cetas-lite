---
id: collect-261001-ia-llm/ia-llm/ia-agents-concepts-25
title: "Concepts : agents IA, agentic, autonomie"
domain: ia-llm
role: reference
task: reference
actors: ["OpenAI", "OpenRouter"]
dates: ["2026-09-20", "2026-09-21", "2026-09-22", "2026-09-27"]
keywords: ["agents", "chatgpt", "distribution", "rlhf"]
source: docs/RAG/collect-261001-ia-llm/ia_agents_concepts.md
source_anchor: ""
source_lines: [3480, 3563]
sha256: c9cb68123e7bc4b98101b7be884ffa8324e3c4f1ed79741531a94bff407fdad8
---

# Concepts : agents IA, agentic, autonomie

- **Fiche signalétique (vérifié le 27/09/2026)** : **Jev** est le premier
  modèle « System One » de **TypeSafe AI** (San Francisco, fondée 2024 par
  **Diogo Almeida** — ex-OpenAI, RLHF / InstructGPT / ChatGPT / GPT-4 —,
  Erik Gafni et Sasha Sheng). Sorti de **2 ans de stealth le 15 septembre
  2026** en early access, en même temps qu'une **levée de 40 M$ menée par
  DCVC** (valorisation ~200 M$ selon Forbes). Modèle courant : **jev-1.13.0**
  (alias `jev-latest`), propriétaire, **poids fermés, API hébergée aux USA**.
  Sources : Wikipedia `Jev_(AI_model)`, `docs.typesafe.ai`, presse Forbes /
  TechCrunch.
- **Correction par rapport au brief** : le brief disait « early access ».
  Vérifié : TypeSafe a **supprimé la waitlist le 20/09/2026** (5 $ de crédit
  offert) et est passé en **GA le 21/09/2026**, avant de **mettre en pause les
  nouvelles inscriptions le 22/09/2026 à 06:19 UTC** pour protéger la qualité
  de service (comptes existants conservés). Au 27/09/2026, l'accès direct
  reste en pause ; les routes tierces (Vercel, Cloudflare, OpenRouter)
  servaient encore le modèle selon la presse.
- **Le nom** : « Jev » viendrait de **W. S. Jevons** (paradoxe de Jevons),
  « System One » de **Kahneman** — les deux sources indépendantes
  convergent sur ce point.

### 137.1. Le concept : System One vs System Two, appliqué aux agents

- **Thèse de TypeSafe** (Kahneman, *Thinking, Fast and Slow*) : le System 1
  est rapide et intuitif, le System 2 est lent et délibératif. L'industrie a
  construit des modèles **System Two** (les LLM : génération lente, coût
  élevé), alors que la plupart du logiciel a besoin de réponses **System
  One** : un jugement rapide, calibré, consommable par du code.
- **Jev n'est PAS un LLM** : il **ne génère pas de texte**, ne converse pas,
  n'écrit pas de code, n'appelle pas d'outils. Il prend un **état** (texte,
  JSON ou tableau : ticket, log, blob) + des **questions typées**, et renvoie
  des **décisions typées avec probabilités** — en un seul passage
  non-autorégressif.
- Ce que le modèle « comprend » en langage naturel, il le rend sous forme de
  **nombres, pas de phrases** (Simon Willison parle de « decision models »).
  La forme de sortie est fixée à l'avance : **rien à parser, pas de JSON
  malformé à réessayer**.

### 137.2. Les trois primitives : Choice, Score, Noul

- **`Choice`** — choisir une option dans une liste fournie par l'appelant.
  Renvoie `choice` (l'option gagnante), `probabilities` (distribution sur
  toutes les options) et `confidence`. Jusqu'à **255 options** (schéma
  two-stage score-then-choose selon le blog TypeSafe).
- **`Score`** — positionner l'état sur une **rubrique ordonnée** définie par
  l'appelant (2 à 10 niveaux). Renvoie `score` (position pondérée par les
  probabilités), `probabilities`, `confidence` et la `legend` des niveaux.
- **`Noul`** — « cette affirmation est-elle vraie ? » (le nom vient de la
  loi de Bernoulli). Renvoie une probabilité `noul` entre 0 et 1.
- **Toutes les questions d'un appel sont évaluées en parallèle et en
  isolation contre le même état** : ajouter des questions ne change presque
  pas la latence et ne coûte que des tokens d'entrée supplémentaires
  (baratins). Règle de conception de la doc : **une question = un gut-check
  atomique** (« ce qu'une personne compétente trancherait en quelques
  secondes avec le bon contexte »). Si le jugement demande du raisonnement
  étendu ou pèse plusieurs facteurs indépendants, **décompose** : une
  question par facteur, puis combine les résultats **dans ton code** (avec tes
  propres coefficients). Chaque Score reste **unidimensionnel**.

### 137.3. RLCD : ce que ça change par rapport à RLHF

- **RLHF** (Reinforcement Learning from Human Feedback) optimise pour la
  **préférence humaine** : le modèle apprend à plaire, à écrire du texte que
  les humains aiment. Effet secondaire documenté : **surconfiance** et
  flatterie. Citation d'Almeida (Forbes) : *« We've been optimizing for
  humans and we're super human at pleasing humans. »*
- **RLVR** (Reinforcement Learning with Verifiable Rewards) optimise pour
  des récompenses vérifiables (maths, code qui passe des tests) : excellent
  pour le raisonnement, mais toujours au service de la **génération**.
- **RLCD** (Reinforcement Learning for Calibrated Decisions) optimise pour
  des **probabilités honnêtes** : si le modèle dit « 70 % de confiance », il
  doit avoir raison environ 70 % du temps, sur un grand nombre de
  prédictions. Appliqué **au-dessus d'un modèle de base pré-entraîné**
  (confirmé par le CEO sur Hacker News). Architecture non publiée, **aucun
  papier au 27/09/2026** — on prend le mécanisme sur parole.
- **Conséquence pratique** : parce que l'espace de sortie est fixé à
  l'avance, le modèle **ne peut pas inventer une option qui n'existe pas**.
  TypeSafe parle de « zéro hallucination structurelle ». Traduction honnête :
  c'est une **validité de schéma** (pas d'erreur de type), pas une
  **correction** — il peut se tromper d'option en toute confiance. Et la
  calibration est une propriété de **groupes** de prédictions, pas une
  garantie sur une réponse isolée.

### 137.4. L'API (forme exacte vérifiée dans `docs.typesafe.ai`, 27/09/2026)

