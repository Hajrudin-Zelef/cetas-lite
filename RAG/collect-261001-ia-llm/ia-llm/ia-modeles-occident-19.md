---
id: collect-261001-ia-llm/ia-llm/ia-modeles-occident-19
title: "Encyclopédie des modèles IA — Volume 2 : l'Occident"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "Google", "MiniMax", "Moonshot", "OpenAI", "OpenRouter", "Z.ai"]
dates: ["2026-09-27"]
keywords: ["agent", "glm", "kimi", "open source", "research"]
source: docs/RAG/collect-261001-ia-llm/ia_modeles_occident.md
source_anchor: ""
source_lines: [1563, 1579]
sha256: 91e723df97aaa413866e15c4b1dd25521088ea7c921aaf9414c5ea702a57ec8d
---

# Encyclopédie des modèles IA — Volume 2 : l'Occident

| Champ | Valeur vérifiée |
|---|---|
| Nom exact | Hermes Agent (repo `NousResearch/hermes-agent`) |
| Sortie | **février 2026** (eweek) ; v0.10.0 citée en juin 2026 |
| Statut | disponible — open source **MIT** ; 95,6K–135K+ stars GitHub selon les sources |
| Description | agent harness générique auto-évolutif : **closed learning loop (GEPA)** — après une tâche complexe, phase réflexive : analyse, extraction de patterns, écriture de nouveaux skill files (standard agentskills.io) réutilisés ensuite. **Mémoire 3 couches** (session, épisodique SQLite, skills procédurales). 40+ à 118 skills bundlés selon les sources (40+ eweek, 70+ Ollama, 118 Medium — **divergence, dépend de la version**). Model-agnostic (OpenRouter, Anthropic, OpenAI, Nous Portal, Kimi, MiniMax, GLM, endpoint custom). Gateways messaging : Telegram, Discord, Slack, WhatsApp, Signal, Email, Google Chat (20 plateformes annoncées). Installateur one-line curl |
| Repo compagnon | `hermes-agent-self-evolution` : pipeline DSPy + GEPA qui fait évoluer les skills/prompts/code de l'agent |
| Infra associée | **atropos** (environnements RL pour collecter/évaluer les trajectoires LLM), **Psyche** (réseau d'entraînement distribué), **Nous Portal** (API/subscription, OAuth, Tool Gateway avec image-gen/TTS/web-search/browser), Nous Chat, simulateurs |
| Prix | gratuit (open source) ; Nous Portal = subscription (tarifs **non vérifiés au 27/09/2026**) |
| Note | le mécanisme GEPA est décrit comme accepté en **oral ICLR 2026** (Berkeley/Stanford), surpassant le RL de 6–20 points avec 35× moins de rollouts — source : rapport Delphi Digital (secondaire, non recoupée ; expansion exacte de l'acronyme GEPA **non vérifiée**). Hermes 4 est explicitement **déconseillé par Nous Research pour les boucles tool-calling rapides** de Hermes Agent (mieux : modèles agentiques frontier du catalogue) |

---

# PARTIE XII — Synthèse : comparer, dater, retenir

## 136. Tableau comparatif géant — l'Occident au 27/09/2026

