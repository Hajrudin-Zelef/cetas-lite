---
id: collect-261001-ia-llm/ia-llm/ia-agents-concepts-19
title: "Concepts : agents IA, agentic, autonomie"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "Google", "Meta", "Microsoft"]
dates: ["2026-06-16"]
keywords: ["agent", "agentic", "agents", "arr", "benchmark", "benchmarks", "claude", "copilot", "leaderboard", "mcp", "research", "valuation"]
source: docs/RAG/collect-261001-ia-llm/ia_agents_concepts.md
source_anchor: ""
source_lines: [2758, 2818]
sha256: 73bf6c26aef7c717bd873d3a75ea92465f3d5590b0731ce536c4e4c184a54d68
---

# Concepts : agents IA, agentic, autonomie

- **ACP (Agent Client Protocol)** : protocole ouvert permettant à des éditeurs
  et interfaces de piloter des agents de code ; supporté par la version récente
  d'Open Interpreter.
- **Agent** : système qui poursuit un objectif en bouclant plan → action →
  observation via des outils, jusqu'à l'objectif ou l'épuisement du budget.
- **Agentic (workflow)** : mode de travail où le modèle décide lui-même des étapes,
  par opposition à un pipeline fixé à l'avance.
- **AgentDojo** : benchmark de **sécurité** des agents (résistance aux prompt
  injections : succès de la tâche + attaques bloquées).
- **AutoGen** : framework multi-agents de Microsoft, conversationnel et
  async (v0.4+) ; adapté aux charges événementielles.
- **Autonomie (niveaux 0-4)** : 0 manuel, 1 assisté, 2 supervisé, 3 semi-autonome,
  4 autonome. La prod sérieuse vit aux niveaux 2-3.
- **BFCL (Berkeley Function-Calling Leaderboard)** : benchmark de référence pour
  la qualité du function calling (appels simples, multi-tours, multi-étapes).
- **Budget (dur)** : plafond imposé par le harness (étapes, tokens, temps) que le
  modèle ne peut pas négocier ; à distinguer d'une consigne dans le prompt.
- **Checkpointer** : mécanisme (LangGraph) qui sauvegarde l'état du graphe après
  chaque nœud → reprise, audit, time-travel.
- **Claude Agent SDK** : SDK d'Anthropic pour construire des agents (MCP natif,
  computer use) dans son écosystème.
- **Claude Cowork** : fonctionnalité « agentic desktop » d'Anthropic (research
  preview janv. 2026, à revérifier) : agent fichiers sans exécution de code.
- **Contamination (benchmark)** : quand les données d'évaluation ont fuité dans
  l'entraînement → scores gonflés. D'où les variantes « Verified » / « Live ».
- **Copilot Cowork** : agent Microsoft 365 (GA 16/06/2026) qui exécute des tâches
  transverses Outlook/Teams/Word/Excel/SharePoint, facturé à l'usage.
- **Correcteur (eval)** : script déterministe qui dit OK/KO pour une tâche d'eval
  set, sans LLM (ex : le fichier existe-t-il ? le service répond-il ?).
- **CrewAI** : framework multi-agents par rôles (« crews ») ; proto rapide,
  moins de contrôle fin que LangGraph.
- **Dérive** : comportement de l'agent qui s'écarte de l'objectif ou du périmètre
  (boucle, contournement, action hors sujet).
- **Dry-run** : mode d'un outil qui décrit ce qu'il ferait sans le faire ;
  indispensable avant toute action sensible.
- **Eval set** : jeu de tâches maison avec correcteurs automatiques ; la « CI »
  de ton agent, à relancer à chaque changement.
- **Fail-closed** : en cas de doute/timeout/panne, le choix par défaut est le
  **refus** (l'inverse, fail-open, est une faute).
- **Fan-out / fan-in** : parallélisation (un superviseur répartit N sous-tâches)
  puis synthèse des résultats.
- **Function calling** : mécanisme par lequel le modèle émet un appel de fonction
  structuré (JSON) ; c'est le harness qui l'exécute réellement.
- **GAIA / GAIA2** : benchmarks d'assistant généraliste multi-outils (Meta/HF) ;
  GAIA2 ajoute des environnements dynamiques.
- **Garde-fou** : dispositif qui borne l'agent : budgets, approbations, sandboxing,
  vérificateurs, logs. Toujours dans le code, jamais seulement dans le prompt.
- **Google ADK** : kit de développement d'agents de Google (Agent Development Kit),
  natif Vertex AI.
- **Harness** : le code qui fait tourner la boucle agentique (ton `agent.py`,
  LangGraph, un CLI…). Synonyme : scaffold.
- **Human-in-the-loop** : un humain valide les étapes sensibles (approbations,
  interrupts LangGraph) ; le contraire du « fire and forget ».
- **Idempotence** : propriété d'une action qui, répétée, produit le même effet
  qu'une seule exécution (clé pour la reprise sur crash).
- **Interrupt** : pause programmable d'un graphe LangGraph (avant/après un nœud)
  pour validation humaine puis reprise.
- **Kill-switch** : mécanisme d'arrêt d'urgence global, testé à l'avance.

## 130. Glossaire (N–Z)

