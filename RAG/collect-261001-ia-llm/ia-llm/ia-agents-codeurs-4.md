---
id: collect-261001-ia-llm/ia-llm/ia-agents-codeurs-4
title: "Les agents codeurs IA"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic"]
dates: []
keywords: ["agent", "agents", "claude", "cost", "mcp", "open source", "pricing", "sandbox"]
source: docs/RAG/collect-261001-ia-llm/ia_agents_codeurs.md
source_anchor: ""
source_lines: [525, 718]
sha256: 1698accb1828d5d43a3f380ab336006767bf113722db15bb2844e3cb6074050c
---

# Les agents codeurs IA

1. Ouvre le marketplace d'extensions (VS Code : `Ctrl+Shift+X`).
2. Cherche **« Kilo Code »** (éditeur : Kilo Code / Kilo-Org).
3. Installe, puis ouvre la vue Kilo Code dans la barre latérale.
4. Crée un compte ou branche ta clé : au choix **BYOK** (tes clés, sans
   marge), **Kilo Gateway** (crédits prépayés), ou **Kilo Pass** (abonnement).

**CLI (méthode documentée : npm) :**

```bash
npm install -g @kilocode/cli
kilo --version

# Dans un projet
cd /tmp/lab-agent
kilo
```

> La doc officielle du CLI ne documente que la méthode npm (des chemins
> curl/brew sont évoqués sur le site marketing mais non confirmés dans la
> doc — à vérifier sur kilo.ai/docs si tu préfères une autre méthode).

**Config partagée** : depuis la reconstruction sur OpenCode, la config vit
dans un **`kilo.jsonc`** central + dossiers **`.kilo/`** (les anciens
`.kilocode/` sont migrés automatiquement). L'agent lit **`AGENTS.md`**
(en cascade racine → sous-dossiers), plus les skills standard
`.claude/skills/` et `.agents/skills/`.

> ⚠️ La config exacte évolue vite depuis la reconstruction d'avril 2026 :
> pars de `kilo.ai/docs`, pas d'un tuto.

## 19. Kilo Code : premier workflow commenté (modes + Agent Manager)

Scénario : corriger un bug dans un script, en utilisant les **modes** comme
les anciens utilisateurs de Roo Code les aiment.

```
1. Dans VS Code, ouvre la vue Kilo Code, choisis le mode **Architect**.
2. Prompt : « Analyse @scripts/sauvegarde.py : la rotation des archives
   supprime parfois la dernière sauvegarde valide. Propose un plan de
   correction, sans écrire de code. »
   → Architect explore en lecture seule et rend un plan numéroté.
3. Relis le plan. S'il est bon, bascule en mode **Code**.
4. Prompt : « Applique le plan validé ci-dessus. »
   → Kilo édite, avec diff visible dans l'éditeur.
5. Bascule en mode **Ask** pour faire expliquer : « Explique-moi la condition
   limite que tu as corrigée, avec un exemple. »
```

**Agent Manager (sessions parallèles)** — le vrai différenciateur :

```
1. Dans la vue Kilo Code : « New session via Agent Manager ».
2. Kilo crée un **git worktree** séparé par session
   (ex. /tmp/lab-agent/.kilo/worktrees/fix-rotation/).
3. Lance 2 sessions en parallèle : l'une corrige le bug, l'autre écrit
   la doc — elles ne se marchent pas dessus car le worktree isole
   le working directory.
4. À la fin : revue des diffs, merge du worktree retenu, suppression des autres.
```

En CLI, le même réflexe :

```bash
cd /tmp/lab-agent
kilo "en mode plan : liste les risques de la fonction de rotation, sans modifier de fichier"
```

## 20. Kilo Code : prix, forces, faiblesses, verdict

**Tarifs (vérifiés sept. 2026, à re-vérifier sur kilo.ai/pricing) :**

| Poste | Coût |
|---|---|
| Agent (VS Code / JetBrains / CLI) | **0 €** — MIT |
| BYOK | 0 € de marge Kilo, tu paies ton fournisseur au token |
| Kilo Gateway | Crédits prépayés, tarifs = tarifs fournisseurs sans marge (annoncé) |
| Kilo Pass | Abonnement **à partir de ~19 $/mois** |
| Offre gratuite | Palier gratuit avec modèles « free » routés automatiquement (liste tournante, **à vérifier** : les modèles gratuits changent vite et les prompts gratuits peuvent servir à l'entraînement — voir section 99) |
| Teams | **~15 $/utilisateur/mois** (annoncé) |

**Points forts** : un agent, trois surfaces avec la même config ; modes
éprouvés (Architect/Code/Debug/Ask) ; Agent Manager/worktrees pour le
parallélisme réel ; inline autocomplete intégrée ; 500+ modèles, zéro marge
en BYOK ; lecture d'AGENTS.md + skills standards (interopérable avec tes
autres outils).

**Faiblesses honnêtes** : jeune (2025) et déjà reconstruit une fois — la doc
et les chemins de config ont changé en avril 2026 ; l'écosystème est plus
petit que Claude Code/Cursor ; le rachat par Anaconda (juil. 2026) pose la
question classique : quelle indépendance à 2 ans ? Le routage auto de modèles
est pratique mais opaque : vérifie ce qui a réellement tourné si tu dois
justifier un coût.

**Verdict Zelef** : excellent **deuxième agent** après un CLI de référence.
Si tu vis dans VS Code et que tu veux de l'open source sans abonnement
obligatoire, c'est le candidat n°1. Pour du 100 % terminal scripté,
OpenCode/Claude Code restent plus simples.

## 21. Claude Code : c'est quoi, pour qui

**Claude Code** est l'agent codeur officiel d'**Anthropic**. C'est la
référence du marché en 2026 : celui contre lequel tout le monde se compare.
Il ne tourne **que sur les modèles Claude** (Sonnet/Opus/Haiku) — c'est à la
fois sa force (le harnais est co-développé avec le modèle) et sa limite
(verrouillage Anthropic total).

Cinq surfaces, un seul moteur (vérifié juil.–sept. 2026) :

- **Terminal** : `cd ton-projet && claude` — la surface la plus complète.
- **Extension VS Code** (fonctionne aussi dans Cursor) : diffs inline,
  `@`-mentions, revue de plan.
- **Plugin JetBrains** (IntelliJ, PyCharm, WebStorm ; nécessite le CLI à côté).
- **App desktop** (macOS/Windows) : revue visuelle des diffs, sessions
  côte à côte, tâches planifiées.
- **Web** : sessions cloud sans installation locale, récupérables dans le
  terminal via `claude --teleport`.

Pour qui : tu veux **l'agent le plus abouti** et tu acceptes l'écosystème
Anthropic (abonnement ou API). Le choix par défaut pour les longues sessions
autonomes (30+ min) : Sonnet/Opus tiennent le cap mieux que la concurrence
sur les tâches longues, d'après les retours terrain 2026.

## 22. Claude Code : installation pas à pas

```bash
# Méthode officielle — vérifiée sept. 2026
npm install -g @anthropic-ai/claude-code
# (un installeur curl existe aussi : https://claude.ai/install.sh — à vérifier)

# Vérifier
claude --version

# Premier lancement (dans ton bac à sable !)
cd /tmp/lab-agent
claude
```

À l'invite de connexion, deux routes :

**Route A — abonnement Claude (le plus simple) :**

```
/login
```
→ ouvre le navigateur, connecte ton compte Claude **Pro (20 $/mois)** ou
**Max (100 ou 200 $/mois)**. Le plan gratuit **n'inclut pas** Claude Code
(vérifié le 10 sept. 2026).

**Route B — clé API (paiement au token) :**

```bash
export ANTHROPIC_API_KEY="sk-ant-ta-cle-fictive-ici"
claude
```

> 🔒 Ne mets jamais ta vraie clé dans un fichier versionné. `export` dans ton
> `~/.bashrc`/`~/.zshrc`, ou gestionnaire de secrets.

Vérifier que ça marche :

```
> résume en 5 lignes ce que contient ce dossier : @.
```

## 23. Claude Code : commandes essentielles

| Commande | Effet |
|---|---|
| `claude` | Lance la session interactive dans le dossier courant |
| `claude "prompt"` | Exécute un prompt directement (une fois) |
| `claude --teleport` | Récupère une session cloud dans ton terminal |
| `/login`, `/logout` | (Re)connexion au compte / déconnexion |
| `/init` | Génère `CLAUDE.md` à partir du projet |
| `/agents` | Gérer les subagents personnalisés |
| `/skills` | Lister/gérer les skills (`.claude/skills/`) |
| `/mcp` | Gérer les serveurs MCP |
| `/hooks` | Configurer les hooks (ex. : vérif auto après écriture) |
| `/status` | État session, modèle, conso |
| `/cost` | Coût estimé de la session (utile en API !) |
| `/compact` | Résume la session pour libérer du contexte |
| `/clear` | Nouvelle conversation, même dossier |
| `Shift+Tab` | Bascule en **Plan mode** (lit, propose, n'écrit pas) |
| `Esc` | Interrompt l'agent en cours (ton frein d'urgence) |
| `--dangerously-skip-permissions` | ⚠️ Désactive les demandes de permission — **à ne jamais utiliser** sauf sandbox jetable |

Fichier mémoire : **`CLAUDE.md`** à la racine (généré par `/init`), relu à
chaque session. Claude Code lit aussi **`AGENTS.md`** — tu peux n'en
maintenir qu'un et le dupliquer/symlinker selon l'outil dominant de l'équipe.

```bash
# Astuce : un seul fichier source de vérité
ln -s AGENTS.md CLAUDE.md
```

