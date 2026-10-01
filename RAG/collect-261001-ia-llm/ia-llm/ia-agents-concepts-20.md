---
id: collect-261001-ia-llm/ia-llm/ia-agents-concepts-20
title: "Concepts : agents IA, agentic, autonomie"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "Google", "Hugging Face", "Microsoft", "OpenAI", "Z.ai"]
dates: ["2026-09-13"]
keywords: ["agent", "agentic", "agents", "apache", "benchmark", "benchmarks", "claude", "gemini", "glm", "guardrails", "incident", "mcp"]
source: docs/RAG/collect-261001-ia-llm/ia_agents_concepts.md
source_anchor: ""
source_lines: [2819, 2945]
sha256: 97725c1c5a95c0f0d2792e88441be417f9b1c6bd4f6da3c9bb1f5c2b19d5d5b9
---

# Concepts : agents IA, agentic, autonomie

- **LangGraph** : framework (LangChain) de graphes d'états pour agents : nœuds,
  arêtes conditionnelles, checkpoints, streaming, interrupts. Le choix par défaut
  pour de la prod auditable en 2026.
- **LlamaIndex (agents)** : framework orienté données/RAG ; pertinent si ton agent
  est d'abord un agent documentaire.
- **MCP (Model Context Protocol)** : protocole ouvert (Anthropic, fin 2024,
  standardisé en 2025-2026) pour brancher outils et données sur les agents via des
  « serveurs MCP ». Interopérabilité oui ; confiance non (voir alertes 2026).
- **Mémoire court terme** : historique de la session en cours, borné par la fenêtre
  de contexte ; à élaguer/résumer.
- **Mémoire long terme** : faits persistants entre sessions (fichier, SQLite,
  vecteurs) ; ton RAG peut en tenir lieu.
- **Moindre privilège** : chaque agent ne reçoit que les outils strictement
  nécessaires à son rôle.
- **Multi-agents** : plusieurs agents coordonnés (superviseur, graphe, débat…) ;
  justifié par : séparation des rôles, parallélisme, spécialisation des modèles.
- **Niveau d'autonomie** : voir « Autonomie (niveaux 0-4) ».
- **Observation** : sortie d'un outil réinjectée dans le contexte ; donnée non
  fiable par défaut (injection possible).
- **OpenAI Agents SDK** : SDK d'OpenAI (handoffs entre agents, guardrails) ;
  chemin court vers la prod dans son écosystème.
- **Open Interpreter** : outil open source (terminal) : langage naturel → code
  exécuté en local. Actif en sept 2026 (réécriture Rust) ; safe mode expérimental.
- **OSWorld** : benchmark d'usage d'un vrai système d'exploitation (369 tâches
  desktop) ; mesure l'état du système.
- **pass@k** : proba de réussir au moins 1 fois sur k essais (optimiste).
- **pass^k** : proba de réussir k fois sur k (pessimiste ; τ-bench). C'est la
  fiabilité, celle qui compte en prod.
- **Périmètre** : ensemble fermé de ce que l'agent a le droit de toucher
  (outils, dossiers, réseau, données). Flou = danger.
- **Plan-then-execute** : le modèle produit un plan complet (validable par un
  humain) puis un exécuteur le déroule ; opposé à l'entrelacé (ReAct).
- **Prompt injection** : instruction malveillante glissée dans une donnée pour
  détourner l'agent ; via outils quand elle arrive par une observation.
- **Pydantic AI** : framework d'agents Python typé (Pydantic) ; pour les amateurs
  de typage strict.
- **ReAct** : pattern fondateur (Yao et al., 2022) : alternance de raisonnement
  (Thought) et d'actions (Action → Observation), traçable et auditable.
- **Réflexion** : l'agent génère une critique verbale de ses échecs et la réutilise ;
  utile avec un signal d'échec clair, coûteuse, à borner.
- **Rollback** : capacité à annuler les effets d'une action (snapshot, backup,
  inverse idempotent) ; exigée pour toute écriture en prod.
- **Sandboxing** : exécution dans un périmètre isolé (compte dédié, conteneur,
  VM, environnement jetable) ; proportionné au pire outil autorisé.
- **Scaffold** : voir « Harness ».
- **Single-agent** : un modèle, une boucle, N outils ; le point de départ obligatoire.
- **Skill** : paquet de connaissances/procédures fourni à l'agent (docs, scripts,
  conventions) ; dans ZCode/Claude Code : unités d'extension configurables.
- **Superviseur** : pattern d'orchestration hub-and-spoke : un agent central
  planifie/délègue/synthétise, des workers spécialisés exécutent.
- **SWE-bench** : benchmark de correction de vrais bugs GitHub (le patch doit
  passer les tests) ; variantes Verified/Live contre la contamination.
- **τ-bench / τ²-bench** : benchmarks de service client qui notent le **respect
  des règles** en plus du succès (Sierra).
- **Terminal-Bench** : 89 tâches terminal (Stanford) : sysadmin, sécu, data,
  SWE ; le harness y fait varier le score de 30-50 points à modèle égal.
- **ToolNode / tools_condition** : utilitaires LangGraph : nœud d'exécution des
  outils / routage « le modèle veut-il appeler un outil ? ».
- **Vérificateur (externe)** : code déterministe qui contrôle qu'un objectif est
  atteint ou qu'une action est autorisée, indépendamment du discours du modèle.
- **WebArena** : benchmark de navigation web réaliste (état fonctionnel des sites).
- **ZCode** : environnement de développement agentique de Z.ai/Zhipu (GLM) :
  desktop + web + CLI, sessions longues pilotables à distance, Apache 2.0 (récent).

---

# PARTIE N — Pour aller plus loin

## 131. Sources et références vérifiées (sept 2026)

Projets et outils cités dans ce guide (vérifie l'URL officielle au jour de ta visite —
les noms sont stables, les chemins exacts peuvent bouger) :

- Open Interpreter : dépôt `openinterpreter/openinterpreter` (GitHub) — actif,
  dernier commit relevé le 13/09/2026 ; doc du safe mode dans `docs/SAFE_MODE.md`
  du dépôt. Recherche : « openinterpreter openinterpreter github ».
- ZCode : site officiel rapporté `zcode.z.ai` ; dépôt open source rapporté
  `zai-org/ZCode` (Apache 2.0) ; docs d'intégration chez les providers
  (ex : guides « ZCode » des intégrateurs API). À vérifier avant téléchargement.
- LangGraph : `langchain-ai/langgraph` (GitHub) — à vérifier ; la doc officielle
  LangChain couvre `create_react_agent`, `StateGraph`, les checkpoints et les
  interrupts (noms d'API à revérifier selon version installée).
- Benchmarks : GAIA (Hugging Face), SWE-bench (leaderboards 2026), τ-bench
  (Sierra), Terminal-Bench (`benchmarkingagents.com/terminal-bench`), HAL
  (Princeton, Pareto précision/coût), AgentDojo (sécurité).
- Incidents cités : AI Incident Database (cas Gemini CLI, n°1178) ; étude Cyera
  « Agent-Inflicted Damage » (avril 2026, 188 cas sans attaquant) ; OECD.AI
  (incident ROME, mars 2026) ; déclarations sept. 2026 (Australie, Transluce) ;
  alerte Microsoft sur les descriptions d'outils MCP (juin 2026).
- Papiers fondateurs : « ReAct: Synergizing Reasoning and Acting in Language
  Models » (Yao et al., 2022) ; « Reflexion » (Shinn et al., 2023).

## 132. Feuille de route proposée (6 semaines)

- **Semaine 1** : fais tourner `agent.py` (partie G) sur 5 tâches en lecture seule.
  Lis un `.jsonl` en entier. Provoque un dépassement de budget.
- **Semaine 2** : ajoute `rag_search` branché sur ton RAG. Mesure le taux de
  réponses « avec source » vs « improvisées ».
- **Semaine 3** : construis ton eval set (20 tâches dont 5 piégées, section 103).
  Vise 100 % trois fois de suite sur le nominal + 100 % de refus sur les pièges.
- **Semaine 4** : réécris l'agent sous LangGraph (partie H) avec vérificateur et
  interrupts. Compare les métriques avec la version maison.
- **Semaine 5** : mets en service UN cas pratique (le tri de tickets ou la
  supervision, niveau 2). Relis les logs chaque jour.
- **Semaine 6** : décide si tu montes au niveau 3 sur un périmètre précis,
  avec la checklist de la section 42. Sinon, reste au niveau 2 : c'est déjà
  90 % du bénéfice.

## 133. Le mot de la fin

Un agent, c'est 10 % de modèle et 90 % de harness : boucle, outils, budgets,
vérifications, logs. Les gens qui « maîtrisent l'IA » en 2026 ne sont pas ceux
qui connaissent le plus de noms d'outils — ce sont ceux qui ont déjà vu une boucle
dérailler à 3h du matin et qui ont le kill-switch, les logs et le rollback pour
raconter l'histoire le lendemain.

Construis petit, mesure tout, monte en autonomie lentement. Et quand un commercial
te vendra un « agent 100 % autonome », demande-lui : où sont les budgets, où sont
les approbations, où sont les logs ? S'il bafouille, tu sauras.

---

*Fin du guide — Concepts : agents IA, agentic, autonomie.*
*Généré le 27 septembre 2026. Les versions, prix et scores marqués « à vérifier »*
*évoluent vite : revérifie avant d'agir.*
---

