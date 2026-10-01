---
id: collect-261001-ia-llm/ia-llm/ia-agents-codeurs-10
title: "Les agents codeurs IA"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "Microsoft", "OpenAI", "OpenRouter"]
dates: []
keywords: ["agent", "agents", "apache", "claude", "copilot", "mcp", "open source"]
source: docs/RAG/collect-261001-ia-llm/ia_agents_codeurs.md
source_anchor: ""
source_lines: [1605, 1765]
sha256: f5f439eb2d218bf08653fe0a5a640e48081ed61f01055cc295b573138348e4a3
---

# Les agents codeurs IA

## 57. Kilo Code côté éditeur : ce qui change vs Cline

Si tu connais Cline, voici ce que Kilo Code ajoute/changent (vérifié 2026) :

| Aspect | Cline | Kilo Code |
|---|---|---|
| Base de code | Historique (Cline) | **Reconstruit sur OpenCode** (avr. 2026) |
| Modes | Plan/Act (un agent) | **Architect / Code / Debug / Ask** + modes perso |
| Autocomplétion inline | Non | **Oui** |
| Sessions parallèles | Non (expérimental) | **Agent Manager + git worktrees** |
| MCP | Manuel (`cline_mcp_settings.json`) | **Marketplace intégré** |
| Routage modèles | Manuel | **Automatique** (par palier) |
| Config | `cline_mcp_settings.json` global | `kilo.jsonc` + `.kilo/` (partagée CLI/IDE) |
| Fichier d'instructions | Modes JSON perso | **`AGENTS.md`** (standard) |
| IDE | VS Code (+ JetBrains Enterprise) | **VS Code + JetBrains + CLI + Cloud** |

Migration Cline/Roo → Kilo : Kilo lit les `.roomodes` et `.roo/rules/` au
démarrage (compat ascendante annoncée) et migre les `.kilocode/` vers
`.kilo/`. En pratique : installe Kilo à côté de Cline, importe tes clés,
laisse les deux cohabiter une semaine, puis tranche.

## 58. Kilo Code : workflow type côté éditeur

```
1. Ouvre ton projet. La vue Kilo Code propose les modes : Architect, Code,
   Debug, Ask.
2. NOUVELLE FONCTIONNALITÉ :
   a. Mode Architect : « Conçois la fonctionnalité X. Contraintes : […].
      Livre un plan avec fichiers et tests. »
   b. Relis le plan (c'est le moment où tu gagnes de l'argent).
   c. Mode Code : « Implémente le plan validé. »
   d. Mode Debug si les tests échouent : « Les tests échouent avec [log].
      Diagnostique sans tout réécrire. »
3. GROSSE TÂCHE EN PARALLÈLE : Agent Manager → 2 sessions (feature + tests)
   dans 2 worktrees → revue → merge.
4. QUESTION RAPIDE : mode Ask, sans toucher au code.
```

## 59. Extensions agentiques : tableau comparatif

|  | Cline | Kilo Code (ext.) | Cursor (éditeur) | Continue (ext.) |
|---|---|---|---|---|
| Forme | Extension VS Code | Extension VS Code/JetBrains | Éditeur complet | Extension VS Code/JetBrains |
| Licence | MIT (OSS) | MIT (OSS) | Propriétaire | Apache 2.0 (OSS) |
| Modèles | BYOK (200+ via OpenRouter) | BYOK / Gateway / Pass | Multi (dont maison) | BYOK (tout fournisseur) |
| Agent autonome | Oui (Plan/Act) | Oui (modes) | Oui (Agent/Composer) | Agent (plus récent) |
| Tab inline | Non | Oui | Oui (excellent) | Oui |
| MCP | Manuel | Marketplace | Oui | Oui (mode Agent) |
| Prix outil | 0 € | 0 € | ~20 $/mois (à vérifier) | 0 € |
| Idéal pour | Apprendre, contrôle total | Agent OSS multi-surfaces | Confort max, tout intégré | Rester sur VS Code stock, OSS |

## 60. Choisir son extension : arbre de décision

```
Tu veux rester sur VS Code classique ?
├── OUI → Tu veux de l'open source ?
│   ├── OUI → Contrôle étape par étape ? → **Cline**
│   │         Sessions parallèles / modes ? → **Kilo Code**
│   └── NON → **Continue** (l'option OSS pragmatique, section 84)
└── NON (tu acceptes un nouvel éditeur) → **Cursor**
    └── …mais vérifie la question confidentialité (section 99) avant
        d'y mettre du code sensible.
```

## 61. Faut-il cumuler extension + CLI ?

**Oui, et c'est même le setup recommandé en 2026** — mais avec des rôles
distincts :

| Rôle | Outil conseillé | Pourquoi |
|---|---|---|
| Complétion au fil de la frappe | Tab (Cursor) / Supermaven / Copilot | Zéro friction, gain immédiat |
| Refacto cadrée multi-fichiers | Composer (Cursor) / Kilo Code (mode Code) | Diffs relisibles |
| Exploration / plan | CLI (OpenCode/Claude Code) en Plan mode | Moins de distraction, scriptable |
| Exécution longue / CI | `codex exec` / `opencode run` | Non-interactif, JSON, cron-friendly |
| Revue | Subagent reviewer / Bugbot | Deuxième paire d'yeux |

Anti-pattern : **trois agents autonomes sur le même repo en même temps**
(conflits d'édition, factures ×3, aucun responsable). Un seul agent
« écrivain » à la fois par working directory — les autres en lecture.

## 62. Raccourcis et réflexes à ancrer (pense-bête éditeur)

```
PLAN D'ABORD ......... Architect (Kilo) / Plan mode / périmètre validé (Cursor)
JAMAIS SANS DIFF ...... relis chaque diff avant Apply/commit
UN SEUL ÉCRIVAIN ...... un agent actif en écriture par dossier
COMMIT AVANT .......... git commit propre avant chaque session d'agent
SECRETS ............... .env.example, jamais .env réel sous les yeux de l'agent
INTERROMPRE ........... Esc / Stop dès que ça dérive (corriger tôt = 10× moins cher)
SESSIONS COURTES ...... >45 min → /compact ou nouvelle session résumée
```

## 63. Ce que les extensions voient de ton code (rappel vie privée)

Une extension agentique lit **tout ton workspace ouvert** par construction.
Points à vérifier par outil (section 99 détaillée) :

- Où sont traités les prompts (local vs cloud du fournisseur) ?
- Le code sert-il à l'entraînement (opt-in/opt-out) ?
- L'indexation (Cursor) est-elle désactivable par repo ?
- Les logs de session contiennent-ils des extraits de code en clair
  (`~/.claude/`, `~/.codex/`, dossiers de stockage VS Code) ?

## 64. Verdict extensions : la recommandation pragmatique

- **Tu débutes avec les agents** → Cline (2 semaines, tout valider).
- **Tu veux de l'OSS au quotidien dans VS Code** → Kilo Code.
- **Tu veux le confort maximal et tu acceptes l'abonnement + le cloud** →
  Cursor.
- **Tu veux rester sur VS Code stock en OSS** → Continue (section 84).
- Dans tous les cas : **un CLI à côté** (OpenCode ou Claude Code) pour le
  scriptable et les longues sessions.

---

# PARTIE C — L'AGENT PERSONNEL AUTONOME : OPENCLAW

---

## 65. « Open clow » : c'est bien OpenClaw

Le nom entendu (« open clow ») correspond à **OpenClaw** (openclaw.ai,
github.com/openclaw/openclaw) — il n'existe **aucun outil sérieux nommé
« OpenClow »** en septembre 2026. La confusion est fréquente à l'oral :
*claw* (griffe) se prononce comme *clow*.

⚠️ Ne confonds pas non plus avec :

- **Claude** (Anthropic) — le modèle, pas l'agent.
- **Clawdbot / Moltbot** — les **anciens noms** d'OpenClaw (voir section 66).

## 66. OpenClaw : c'est quoi — et son histoire folle

**OpenClaw** est un **agent IA personnel, autonome, auto-hébergé et open
source (MIT)**. La différence fondamentale avec tout ce qui précède : il ne
vit pas dans ton terminal ou ton éditeur, il vit **sur une machine qui
tourne en permanence** (ton serveur, un Mini-PC, un VPS) et tu le pilotes
depuis **tes apps de messagerie** (WhatsApp, Telegram, Discord, Slack,
Signal, iMessage…) ou via le web.

Ce qu'il fait, concrètement : tu lui écris « résume-moi mes mails urgents
et prépare les réponses », il **agit** — lit les mails, rédige, te propose,
envoie si tu l'autorises. Avec des **crons intégrés**, il travaille aussi
**sans qu'on lui demande** (ex. : chaque matin à 7h, briefing du jour).

**L'histoire** (vérifiée, c'est documenté partout en 2026) :

| Date | Événement |
|---|---|
| Nov. 2025 | **Clawdbot** : prototype publié par Peter Steinberger (écrit en ~1 h) |
| Janv. 2026 | Renommé **Moltbot** (plainte marque d'Anthropic — le thème du homard reste) |
| 30 janv. 2026 | Renommé **OpenClaw** (« Moltbot » était imprononçable) |
| Fév. 2026 | **145 000+ stars**, croissance explosive |
| 14 fév. 2026 | Steinberger **rejoint OpenAI** |
| Mars 2026 | **250 000+ stars**, dépasse React sur GitHub — l'un des projets OSS à la croissance la plus rapide de l'histoire |

Écosystème : **ClawHub**, un registre de **skills** (13 000+ skills début
2026 — « le npm des capacités d'agent ») : navigateur, e-mail, fichiers,
API, domotique, etc. 100+ skills préconfigurées à l'install.

