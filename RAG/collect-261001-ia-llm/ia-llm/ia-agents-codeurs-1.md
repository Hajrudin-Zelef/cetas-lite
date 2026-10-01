---
id: collect-261001-ia-llm/ia-llm/ia-agents-codeurs-1
title: "Les agents codeurs IA"
domain: ia-llm
role: reference
task: reference
actors: ["AWS", "Anthropic", "Google", "Microsoft"]
dates: []
keywords: ["agent", "agents", "agentic", "claude", "context window", "copilot", "gemini", "mcp", "model context protocol", "sandbox"]
source: docs/RAG/collect-261001-ia-llm/ia_agents_codeurs.md
source_anchor: ""
source_lines: [1, 141]
sha256: 0f3b353e0fe564e1e4415c9db3e4253e8c6a02ceac8362e497d64b65c8ca9470
---

# Les agents codeurs IA

**Guide pratique — état des lieux fin septembre 2026.**
Pour Zelef, sysadmin qui construit son RAG et veut maîtriser l'outillage IA/dev.

> Règle d'or de ce guide : tout ce qui suit a été vérifié par recherche web en
> septembre 2026. Les prix et limites changent vite — quand une valeur est
> incertaine, elle est marquée **« à vérifier »** avec la source officielle à
> consulter. Aucun mot de passe réel dans ce document : les secrets sont des
> exemples fictifs.

---

## Sommaire rapide

| Partie | Sections | Contenu |
|---|---|---|
| A — Les CLI agentiques | 1–45 | OpenCode, Claude Code, Codex CLI |
| B — Éditeurs & extensions agentiques | 46–72 | Cursor, Cline, Kilo Code |
| C — L'agent personnel autonome | 73–80 | OpenClaw (le « open clow ») |
| D — Plugins & complétion | 81–95 | Continue, Supermaven, Copilot, Tabnine, Amazon Q, JetBrains AI, Zed, Windsurf, Trae, Aider, Gemini CLI, Kiro… |
| E — Transversal | 96–115 | Comparatif, choix, équipe, sécurité, coûts/tokens, MCP, 20 pièges, cas pratiques, glossaire, quiz, pense-bête, pour aller plus loin |

---

## 1. Comment utiliser ce guide

1. **Lis d'abord les sections 2 à 5** : tu poses le vocabulaire (agent, MCP,
   AGENTS.md, contexte) et la carte du territoire.
2. **Choisis UN outil pour commencer** (section 97 t'aide) et suis son bloc
   d'installation pas à pas. N'installe pas tout le même jour : chaque agent a
   ses réflexes, ses fichiers de config et sa facturation.
3. **Fais le premier workflow commenté** de l'outil choisi, sur un projet
   jetable (`/tmp/lab-agent/`), jamais sur un repo de production.
4. **Reviens aux sections transversales** (96–115) avant d'utiliser un agent
   sur du vrai code : sécurité, coûts, revue du code généré.
5. **Le quiz (section 112)** sert de contrôle : si tu as moins de 8/10,
   relis les sections correspondantes.

Convention de notation dans ce guide :

- `$ commande` → à taper dans ton terminal (shell).
- Les blocs `json`/`toml`/`yaml` sont des fichiers de configuration réels.
- ⚠️ = piège fréquent. 💰 = point de coût. 🔒 = point de sécurité.

## 2. Ce qu'est (vraiment) un agent codeur

Un **agent codeur** n'est pas un chatbot qui répond du code. C'est un programme
qui tourne **en boucle** :

```
objectif (ton prompt)
   → le modèle raisonne
   → il appelle des OUTILS (lire un fichier, exécuter un test, chercher dans le repo)
   → il observe le RÉSULTAT de l'outil
   → il recommence jusqu'à atteindre l'objectif ou abandonner
```

Cette boucle s'appelle la **boucle agentique** (agentic loop). Trois
conséquences pratiques :

1. **Il agit, pas seulement il parle** : il modifie des fichiers, lance des
   commandes, crée des commits. D'où l'importance des garde-fous.
2. **Il consomme du contexte** : chaque tour de boucle envoie l'historique au
   modèle. Une session qui tourne 30 minutes peut brûler des centaines de
   milliers de tokens.
3. **Il peut se tromper en cascade** : une mauvaise lecture au tour 2 produit
   un mauvais patch au tour 5, « validé » par un test mal écrit au tour 8.
   La revue humaine reste obligatoire (section 100).

Distinguer trois familles, car on les confond souvent :

| Famille | Ce qu'elle fait | Exemples |
|---|---|---|
| **Complétion** | Suggère du code pendant que tu tapes | Supermaven, Copilot (mode tab), Tabnine |
| **Chat in-IDE** | Répond à des questions sur ton code | Continue (chat), Copilot Chat |
| **Agent** | Exécute des tâches multi-étapes de façon autonome | Claude Code, Codex CLI, OpenCode, Cline, Cursor Agent, Kilo Code |

Un même produit mélange souvent les trois (Cursor fait les trois). Ce guide
les traite séparément pour que tu saches ce que tu paies et ce que tu
autorises.

## 3. Le vocabulaire minimal (à connaître par cœur)

- **Contexte (context window)** : la quantité de texte (en tokens) que le
  modèle peut « voir » d'un coup. Un repo de 200 000 lignes ne rentre jamais
  en entier : l'agent doit choisir quoi lire. Les fenêtres « 1M tokens »
  existent mais coûtent cher et se dégradent en qualité sur les très longs
  contextes.
- **Token** : unité de facturation. ≈ 4 caractères en anglais, ≈ 2–3 en
  français pour du texte ; le code est plus dense. On paie l'**input**
  (ce qu'on envoie) et l'**output** (ce que le modèle génère), à des tarifs
  différents.
- **MCP (Model Context Protocol)** : standard ouvert (initié par Anthropic)
  pour brancher des **outils** sur un agent : un « serveur MCP » expose des
  fonctions (ex. : interroger ta base, piloter ton navigateur, lire tes
  tickets). En 2026, MCP est le standard de fait : presque tous les agents de
  ce guide le supportent.
- **Skills** : paquets de savoir-faire réutilisables (ex. : « comment écrire
  un playbook Ansible à la mode Zelef »). Claude Code les lit dans
  `.claude/skills/`, OpenClaw dans son ClawHub, Kilo Code dans `.agents/skills/`
  (standard ouvert agentskills.io).
- **AGENTS.md / CLAUDE.md** : fichier d'instructions placé à la racine du
  projet, lu automatiquement par l'agent à chaque session. C'est TA doc de
  cadrage : conventions, commandes de test, choses interdites. Le standard
  `AGENTS.md` (multi-outils) a supplanté les fichiers propriétaires.
- **Plan mode** : l'agent explore et propose un plan **sans écrire** de code.
  Tu valides, puis il exécute. Le réflexe n°1 anti-catastrophe.
- **Sandbox / permissions** : ce que l'agent a le droit de faire sans te
  demander (lire ? écrire ? exécuter du shell ? le réseau ?). Chaque outil a
  son modèle ; tous se règlent.
- **BYOK (Bring Your Own Key)** : tu branches ta propre clé API au lieu de
  payer l'abonnement de l'outil. Moins cher si tu consommes peu, sans plafond
  « fair use », mais facturation au token à surveiller.
- **Subagent** : un agent lancé par l'agent principal pour une sous-tâche
  (ex. : explorer 3 dossiers en parallèle), avec son propre contexte. Très
  puissant, très gourmand en tokens.

## 4. Carte du territoire (septembre 2026)

```
                        ┌─────────────────────────┐
                        │   CE QUE TU VEUX FAIRE  │
                        └────────────┬────────────┘
                                     │
            ┌────────────────────────┼────────────────────────┐
            │                        │                        │
     TERMINAL PUR              DANS L'ÉDITEUR           AUTONOME / PERSO
            │                        │                        │
  ┌─────────┴─────────┐   ┌──────────┴──────────┐    ┌────────┴────────┐
  │ Claude Code       │   │ Cursor (éditeur)    │    │ OpenClaw        │
  │ Codex CLI         │   │ Cline (extension)   │    │ (WhatsApp/      │
  │ OpenCode (TUI)    │   │ Kilo Code (ext.)    │    │  Telegram/      │
  │ Kilo Code (CLI)   │   │ Continue (extension)│    │  cron…)         │
  │ Gemini CLI, Aider │   │ Copilot, Windsurf…  │    │                 │
  └───────────────────┘   └─────────────────────┘    └─────────────────┘
        ▲                         ▲                         ▲
        │        MCP / AGENTS.md / Skills : le socle commun  │
        └─────────────────────────┴─────────────────────────┘
```

