---
id: collect-261001-ia-llm/ia-llm/ia-agents-codeurs-6
title: "Les agents codeurs IA"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "OpenAI"]
dates: []
keywords: ["agent", "agents", "apache", "chatgpt", "claude", "mai", "mcp", "open source", "sandbox"]
source: docs/RAG/collect-261001-ia-llm/ia_agents_codeurs.md
source_anchor: ""
source_lines: [914, 1119]
sha256: 4796bc9d3e25f1a8750bb244678d04cf57f1d7e6e429beefbe793e11d998cd9d
---

# Les agents codeurs IA

**Si tu ne dois en choisir qu'un pour commencer : c'est celui-là.** Le plus
abouti, le mieux documenté, le plus prévisible en équipe. Prends la **route
abonnement Pro** si ton usage est régulier (facture fixe, zéro stress), la
**route API** si tu veux contrôler au token près ou brancher d'autres outils
sur la même clé. Pour ton RAG : le serveur MCP + les skills en font le
meilleur « collègue » pour nettoyer/indexer ton corpus documentaire.
Seule vraie réserve : la dépendance à un seul fournisseur — d'où l'intérêt
de garder OpenCode en second pour le jour où tu veux comparer ou basculer.

## 31. Codex (CLI) : c'est quoi, pour qui

**Codex CLI** est l'agent terminal officiel d'**OpenAI** (dépôt open source
**Apache 2.0** : github.com/openai/codex). Comme Claude Code, il lit/édite/
teste ton repo en local ; il tourne sur les modèles **GPT** (GPT-5 par
défaut en 2026).

Positionnement honnête : c'est **le pendant OpenAI de Claude Code**. Mêmes
gestes, philosophie légèrement différente — plus « exécutant » que
« pair collaboratif » : il fonce davantage, demande moins. D'où l'importance
de bien régler ses modes de permission.

Pour qui :

- Tu es déjà abonné **ChatGPT Plus/Pro** : Codex CLI est **inclus**, sans
  surcoût — c'est l'argument massue.
- Tu veux un agent **open source** adossé à OpenAI (auditable, forkable,
  contrairement au CLI fermé de Claude Code).
- Tu veux du **scriptable en CI** : `codex exec` est taillé pour ça.
- Tu travailles dans l'écosystème OpenAI (ou tu veux comparer GPT-5 vs
  Claude sur les mêmes tâches).

## 32. Codex CLI : installation pas à pas

```bash
# Méthode officielle (macOS / Linux) — vérifiée sept. 2026
curl -fsSL https://chatgpt.com/codex/install.sh | sh

# Alternatives
npm install -g @openai/codex
# macOS : brew install --cask codex
# Windows (PowerShell) :
#   powershell -ExecutionPolicy ByPass -c "irm https://chatgpt.com/codex/install.ps1 | iex"
#   (support natif Windows depuis mai 2026 — plus besoin de WSL)

# Vérifier
codex --version
```

Prérequis : **Node.js 22+** pour l'install npm. Sur Windows frais, installe
aussi le **Visual C++ Redistributable** (sinon Codex refuse de démarrer
silencieusement) et **Git for Windows**.

**Connexion — deux routes :**

```bash
cd /tmp/lab-agent
codex
# → choisis « Sign in with ChatGPT » : le navigateur s'ouvre,
#   le token revient dans le terminal.
#   Le login est mis en cache dans ~/.codex/auth.json (fichier à
#   traiter comme un mot de passe !) ou le trousseau de l'OS.

# Alternative : clé API (utile en CI, facturation au token)
printenv OPENAI_API_KEY | codex login --with-api-key

# Machine distante sans navigateur :
codex login --device-auth
# → code à usage unique à saisir ailleurs (bêta, à activer dans les
#   paramètres de sécurité ChatGPT)

# Vérifier l'état
codex login status
```

**Mise à jour** (le CLI évolue vite : ~1 release/semaine en sept. 2026) :

```bash
# Selon ta méthode d'install :
npm install -g @openai/codex
# ou : brew upgrade --cask codex
# ou : codex update   (auto-update si ta version le supporte)
```

## 33. Codex CLI : commandes et modes de permission

```bash
codex                          # session interactive dans le dossier courant
codex "ton prompt"             # exécute un prompt directement
codex exec "prompt"            # mode NON interactif (scripts, CI, cron)
codex exec --json --full-auto "prompt"   # sortie JSON, sans demander
codex login / codex login status
codex update
```

**Les trois modes de permission** — le réglage le plus important :

| Mode | Comportement |
|---|---|
| **Suggest** | L'agent propose, **tu valides chaque action** (défaut prudent) |
| **Auto-edit** | Il édite les fichiers seul, **demande pour exécuter** des commandes |
| **Full-auto** | Il fait tout seul (édition + exécution) — sandboxé, mais tout seul |

```bash
# Choisir le mode au lancement (exemples)
codex --ask-for-approval on-request "prompt"     # proche de Suggest
codex --ask-for-approval on-failure "prompt"
codex --ask-for-approval never "prompt"          # proche de Full-auto ⚠️
```

> 🔒 **Règle** : `Suggest` pour découvrir, `Auto-edit` au quotidien,
> `Full-auto` uniquement dans un bac à sable jetable ou en CI avec un
> périmètre verrouillé (section 99). Le mode Full-auto + un prompt flou =
> la recette classique du « il a tout réécrit ».

**Plan mode** : commande `/plan` dans la session interactive — même
philosophie que les autres : explorer et proposer sans écrire.

Fichier mémoire : **`AGENTS.md`** (le standard multi-outils — pas de fichier
propriétaire chez Codex, bon point d'interopérabilité).

## 34. Codex CLI : premier workflow commenté

Scénario : écrire un script de **vérification d'onduleur** (ton métier !)
avec tests, en mode `codex exec` pour voir la voie scriptable.

```bash
cd /tmp/lab-agent
cat > AGENTS.md <<'EOF'
# AGENTS.md — lab-agent
- Python 3.12. Tests : `pytest -q` (doivent passer).
- Pas de dépendance externe sans validation humaine.
- Ne jamais écrire de secrets en dur ; exemples fictifs uniquement.
EOF

# 1. D'abord en interactif, en plan :
codex
```

```
> /plan
> Besoin : script check_ups.py qui interroge un onduleur en SNMP (hostname,
> communauté en variables d'environnement) et sort un code Nagios
> (0 OK / 1 WARNING / 2 CRITICAL) selon le niveau de batterie.
> Propose un plan sans écrire de code.
```

```
# 2. Plan validé ? On exécute en non-interactif (la voie CI) :
codex exec --full-auto \
  "implémente le plan validé pour check_ups.py : module + tests pytest + README d'usage. \
   Les tests doivent passer. Ne modifie rien d'autre."

# 3. Contrôle systématique :
git status --porcelain
git diff --stat
pytest -q
```

**En CI (GitHub Actions)** — le cas d'usage roi de `codex exec` :

```yaml
# .github/workflows/codex-review.yml — exemple
name: codex-review
on: [pull_request]
jobs:
  review:
    runs-on: ubuntu-latest
    steps:
      - uses: actions/checkout@v4
      - uses: actions/setup-node@v4
        with: { node-version: 22 }
      - run: npm install -g @openai/codex
      - run: printenv OPENAI_API_KEY | codex login --with-api-key
        env: { OPENAI_API_KEY: ${{ secrets.OPENAI_API_KEY }} }
      - run: >
          codex exec --json --full-auto --skip-git-repo-check
          "relis le diff de cette PR et liste les bugs potentiels par sévérité"
```

> 🔒 En CI : secret via `secrets.OPENAI_API_KEY`, jamais en dur ; périmètre
> `Full-auto` limité à de la **lecture/revue**, jamais d'écriture auto sur
> `main`.

## 35. Codex CLI : cloud, subagents, MCP, SDK

- **Codex cloud** : depuis le terminal, tu peux basculer une tâche vers le
  cloud (`codex cloud`) puis **rappliquer le résultat** dans ton repo local.
  Pratique pour les longues tâches : tu fermes ton laptop, ça continue.
- **Subagents** : Codex peut découper une investigation en agents parallèles
  et consolider leurs résultats dans la session principale.
- **MCP** : support **client** (tu branches des serveurs MCP). Config
  partagée dans `~/.codex/config.toml` — le même fichier est utilisé par le
  CLI, l'app desktop et l'extension IDE (bon point d'unification).

```toml
# ~/.codex/config.toml — extrait (schéma indicatif, vérifier la doc)
# [mcp.servers.docs]
# command = "npx"
# args = ["-y", "@modelcontextprotocol/server-filesystem", "/home/zelef/docs"]
```

- **Codex SDK** : embarque le même agent dans tes propres outils/scripts —
  la voie « je construis mon automatisation » au-delà du CLI.

## 36. Codex CLI : prix (vérifiés sept. 2026)

