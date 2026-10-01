---
id: collect-261001-ia-llm/ia-llm/ia-agents-codeurs-17
title: "Les agents codeurs IA"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic"]
dates: []
keywords: ["agent", "agents", "apache", "benchmark", "chatgpt", "claude", "cloud agent", "cost", "mcp", "open source", "pricing", "sandbox"]
source: docs/RAG/collect-261001-ia-llm/ia_agents_codeurs.md
source_anchor: ""
source_lines: [2762, 2883]
sha256: c5cc91a17de7032aacd4c9bd47981a255f612088bfd4a90bf618714dc12c6be4
---

# Les agents codeurs IA

| Terme | Définition courte |
|---|---|
| Agent | Programme qui boucle : raisonne → appelle des outils → observe → recommence |
| Boucle agentique | Le cycle raisonnement → action → observation jusqu'à l'objectif |
| Contexte | Tout ce que le modèle « voit » d'un coup (fenêtre limitée, en tokens) |
| Token | Unité de facturation (≈ 4 caractères anglais) ; input et output tarifés séparément |
| MCP | Standard ouvert pour brancher des outils/serveurs sur un agent |
| Skill | Savoir-faire packagé et réutilisable (dossier + instructions) |
| Subagent | Agent lancé par l'agent principal pour une sous-tâche, avec son propre contexte |
| AGENTS.md | Fichier d'instructions projet lu par (presque) tous les agents |
| CLAUDE.md | Équivalent propriétaire chez Claude Code (généré par `/init`) |
| Plan mode | Phase d'exploration/proposition sans écriture, avant exécution |
| BYOK | Bring Your Own Key : tu branches ta clé API au lieu de l'abonnement |
| Sandbox | Environnement d'exécution restreint (pas de réseau, pas d'écriture hors périmètre…) |
| Permission | Ce que l'agent peut faire sans demander (lire/écrire/exécuter) |
| Auto-approve | Validation automatique par type d'action (à régler avec prudence) |
| TUI | Interface texte dans le terminal (ex. : OpenCode) |
| Complétion | Suggestion de code pendant la frappe (Tab) |
| Composer | Mode multi-fichiers de Cursor (décris → il applique sur N fichiers) |
| Background Agent | Agent qui tourne sans bloquer ton éditeur (Cursor) |
| Cloud Agent | Agent qui tourne sur l'infra du fournisseur, pas sur ta machine |
| Worktree | Clone git léger d'une branche dans un dossier séparé (parallélisme) |
| Checkpoint | Point de restauration avant une action d'agent (annulation d'urgence) |
| Cache de prompt | Contexte stable facturé ~10× moins cher (input en cache) |
| Fair use | Limite d'usage « raisonnable » des abonnements, souvent floue |
| Crédits | Unité de facturation interne (Cursor, Kilo Gateway…) |
| ClawHub | Registre communautaire de skills pour OpenClaw |
| Prompt injection | Attaque : des instructions malveillantes cachées dans des données lues par l'agent |
| Typosquatting | Paquet malveillant au nom proche d'un paquet légitime |
| ACP | Agent Client Protocol (Zed) : brancher n'importe quel agent externe à l'éditeur |
| Headless | Mode sans interface, pour scripts/CI (`codex exec`, `opencode run`) |
| Idempotence | Propriété : rejouer l'action ne change rien la 2ᵉ fois (crucial en Ansible) |

## 106. Pense-bête de poche (à imprimer)

```
┌─────────────────────────────────────────────────┐
│           AGENTS CODEURS — PENSE-BÊTE           │
├─────────────────────────────────────────────────┤
│ AVANT CHAQUE SESSION                            │
│ □ git status propre (commit avant)              │
│ □ .env / secrets hors de portée                 │
│ □ AGENTS.md à jour                              │
│                                                 │
│ PENDANT                                         │
│ □ Plan d'abord (Tab / Shift+Tab / /plan)        │
│ □ Un seul écrivain par dossier                  │
│ □ Esc dès que ça dérive                         │
│ □ /cost ou /status de temps en temps            │
│                                                 │
│ APRÈS                                           │
│ □ git diff relu ligne par ligne                 │
│ □ pytest vert (vrai vert)                       │
│ □ Pas de secret, pas de dépendance sauvage      │
│ □ Commit avec message conventionnel             │
│                                                 │
│ INTERDIT (sauf validation explicite)            │
│ ✗ rm -rf / sudo / pip install / curl|bash       │
│ ✗ --dangerously-skip-permissions / Full-auto    │
│ ✗ Secrets dans les prompts                     │
│ ✗ Toucher la prod directement                   │
│                                                 │
│ INSTALLS (sept. 2026)                           │
│ $ curl -fsSL https://opencode.ai/install | bash │
│ $ npm i -g @anthropic-ai/claude-code            │
│ $ npm i -g @openai/codex                        │
│ $ npm i -g @kilocode/cli                        │
│ Extensions : Cline / Kilo Code / Continue       │
│ Éditeur : cursor.com/download                   │
└─────────────────────────────────────────────────┘
```

## 107. Checklist : ta première semaine avec les agents

- [ ] **Jour 1** : prérequis (section 5) + bac à sable `/tmp/lab-agent`.
- [ ] **Jour 1** : installe **UN** outil (recommandé : Claude Code ou OpenCode).
- [ ] **Jour 2** : `/init`, corrige le AGENTS.md/CLAUDE.md généré (section 6).
- [ ] **Jour 2** : workflow plan-d'abord sur une micro-feature (sections 12/24).
- [ ] **Jour 3** : installe **Cline**, travaille 1 journée en « tout valider ».
- [ ] **Jour 4** : configure le pense-bête sécurité (sections 92–93) dans ton AGENTS.md.
- [ ] **Jour 4** : active les **alertes de facturation** chez ton fournisseur.
- [ ] **Jour 5** : premier vrai usage (petit script utilitaire), revue en 4 passes (section 90).
- [ ] **Jour 5** : fais le **quiz** (section 112) — vise 8/10 minimum.
- [ ] **Jour 6–7** : écris ta première **skill** (le prompt que tu as réécrit 3 fois).

## 108. Pour aller plus loin

**Documentation officielle (références à jour — préfère-les aux tutos) :**

- OpenCode : opencode.ai/docs (config, plugins, modèles)
- Claude Code : docs.anthropic.com (CLI, hooks, MCP, skills)
- Codex : github.com/openai/codex (open source — lis le code !) + learn.chatgpt.com/docs (auth, pricing)
- Cursor : cursor.com/docs + cursor.com/pricing (vérifie la grille en vigueur)
- Kilo Code : kilo.ai/docs (kilo.jsonc, Agent Manager, CLI)
- Cline : docs.cline.bot (modes, MCP, auto-approve)
- OpenClaw : openclaw.ai + github.com/openclaw/openclaw (README, sécurité)
- Continue : docs.continue.dev (config.yaml, modèles locaux)
- MCP : modelcontextprotocol.io (le standard, ses serveurs officiels)

**Pour creuser :**

- Le dépôt **openai/codex** (Apache 2.0) : lire le code d'un vrai agent est
  la meilleure école d'architecture agentique.
- Les **awesome-lists** : awesome-ai-coding-tools, awesome-openclaw (panorama
  communautaire, à trier).
- **SWE-bench** : le benchmark de référence des agents codeurs — pour
  comprendre ce que « bon en code » veut dire mesurablement.
- Tes propres **logs de session** (`~/.claude/`, `~/.codex/`) : relis une
  session d'il y a 15 jours, note ce que tu ferais différemment — c'est là
  que naissent tes skills.

**Prochaines étapes pour ton profil :**

1. Serveur MCP « docs » sur ton corpus RAG (section 97) — le pont entre tes
   guides et tes agents.
2. Skill « revue-ansible » versionnée (section 98).
3. `codex exec` en CI sur tes playbooks (sections 34, 101).
4. OpenClaw durci en vigie (lecture seule) sur une machine dédiée (sections
   67–72).

## 109. Ce que ce guide ne couvre pas (honnêteté)

