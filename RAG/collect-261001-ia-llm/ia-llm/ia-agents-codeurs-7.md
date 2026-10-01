---
id: collect-261001-ia-llm/ia-llm/ia-agents-codeurs-7
title: "Les agents codeurs IA"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "OpenAI"]
dates: ["2026-09-10"]
keywords: ["agent", "agents", "apache", "arr", "chatgpt", "claude", "mcp", "open source"]
source: docs/RAG/collect-261001-ia-llm/ia_agents_codeurs.md
source_anchor: ""
source_lines: [1120, 1283]
sha256: f5591815ecbb97cf1f9700b6b666b67a99b233d6f0e7a1195a1d375fd9534db9
---

# Les agents codeurs IA

| Route | Prix | Notes |
|---|---|---|
| ChatGPT **Plus** | **20 $/mois** | Codex CLI **inclus** |
| ChatGPT **Pro** | paliers **5× / 20×** (le 20× a vu ses nouvelles inscriptions **mises en pause** le 10/09/2026 — existants non impactés) | Quotas supérieurs |
| ChatGPT **Free / Go** | inclus avec **limites basses** | Pour tester, pas pour produire |
| **Clé API** | **au token**, tarifs API OpenAI standard | Pas de quota plan, mais pas de fonctions cloud (revue, Slack…) |
| Business / Edu / Entreprise | inclus selon l'offre | **À vérifier** selon ton contrat |

Limites : usage régi par des **fenêtres glissantes de 5 h et hebdomadaires**,
variables selon modèle et tâche — consulte avec `/status` dans le CLI.
Avec une clé API, tu paies au token et tu n'as pas ces quotas, mais tu dois
surveiller la facture (section 98).

Ordres de grandeur API (modèles GPT-5, sept. 2026 — **à vérifier**, ça bouge
vite) : l'input se compte en dizaines de $/M tokens, l'output en centaines
de $/M sur les plus gros modèles. Traduction : **une session d'1 h sur
GPT-5.x peut coûter de quelques $ à 20 $+** selon le modèle exact. En
abonnement Plus, c'est « inclus » dans les 20 $ — imbattable si tu utilises
déjà ChatGPT.

## 37. Codex CLI vs Claude Code : différences franches

| Critère | Claude Code | Codex CLI |
|---|---|---|
| Fournisseur / modèles | Anthropic uniquement (Sonnet/Opus/Haiku) | OpenAI uniquement (GPT) |
| Licence du CLI | **Fermé** | **Open source (Apache 2.0)** |
| Fichier mémoire | `CLAUDE.md` + `AGENTS.md` | `AGENTS.md` |
| Plan mode | `Shift+Tab` | `/plan` |
| Permissions | Default / Auto-accept / `--dangerously-skip-permissions` | **Suggest / Auto-edit / Full-auto** |
| MCP | Client **+ serveur** (mature) | Client (en 2026) |
| Skills | Oui (`.claude/skills/`) | Pas d'équivalent direct en 2026 |
| Scriptable CI | `claude "prompt"` | **`codex exec`** (pensé pour ça : `--json`, `--full-auto`) |
| Cloud | Sessions web + `--teleport` | `codex cloud` + réapplication locale |
| Prix d'entrée | Pro 20 $/mois ou API | **Plus 20 $/mois (inclus)** ou API |
| Personnalité | « Pair collaboratif » (questionne, nuance) | « Exécutant » (fonce, demande moins) |

**Comment choisir** : tu as déjà ChatGPT Plus ? Codex CLI ne te coûte rien de
plus — teste-le en parallèle. Tu veux le meilleur agent « qui réfléchit avec
toi » sur des tâches longues ? Claude Code. Tu veux du CI scriptable et de
l'open source ? Codex CLI. Beaucoup de devs en 2026 **utilisent les deux** :
Claude Code pour concevoir, Codex pour exécuter en boucle.

## 38. Codex CLI : bonnes pratiques

1. **Toujours `AGENTS.md`** à la racine avant la première session — Codex
   n'a pas de `/init` magique équivalent : écris-le toi (section 6).
2. **Reste en Suggest/Auto-edit** tant que tu n'as pas vu l'agent travailler
   10 fois. Le Full-auto se mérite.
3. **`codex exec` + `git diff --exit-code`** : en script, tout changement
   inattendu doit faire échouer le job.
4. **Prompts fermés en non-interactif** : objectif + critères d'arrêt +
   périmètre de fichiers explicites. Jamais de « améliore le projet ».
5. **`/status`** régulièrement pour voir où tu en es des quotas (route
   abonnement).
6. **`~/.codex/auth.json` = secret** : permissions 600, jamais copié,
   jamais dans un backup non chiffré.
7. Sur Windows : vérifie le **VC++ Redist** avant de déclarer « ça ne
   marche pas » (section 32).

## 39. Codex CLI : points forts / faiblesses honnêtes

**Points forts** : inclus dans ChatGPT Plus (coût marginal nul si déjà
abonné) ; **open source Apache 2.0** ; `codex exec` excellent pour CI/scripts ;
`codex cloud` pour les longues tâches ; config unifiée CLI/desktop/IDE ;
maturité OpenAI sur l'exécution.

**Faiblesses honnêtes** : **verrouillage OpenAI** (comme Claude Code avec
Anthropic) ; personnalité « fonceuse » = plus de bêtises si mal cadré ;
MCP moins mature que Claude Code (client seul en 2026) ; pas de système de
skills équivalent ; les quotas d'abonnement sont opaques (fenêtres
glissantes non publiées en détail) ; la doc bouge vite (releases
hebdomadaires).

## 40. Codex CLI : verdict Zelef

**Le meilleur rapport qualité/prix si tu as déjà ChatGPT Plus.** Pour un
sysadmin : `codex exec` dans un cron qui génère un rapport d'inventaire, ou
dans une GitHub Action qui relit tes playbooks Ansible, c'est du pain bénit.
Garde-le en **exécutant** et Claude Code (ou OpenCode) en **concepteur** :
c'est le duo le plus rentable de 2026. Seul bémol structurel : tout repose
sur OpenAI — ne mets pas 100 % de tes automatisations critiques dessus sans
plan B.

---

# PARTIE B — ÉDITEURS & EXTENSIONS AGENTIQUES

---

## 41. Cursor : c'est quoi, pour qui

**Cursor** (cursor.com, édité par **Anysphere**) est un **éditeur de code IA
natif** : un fork de VS Code repensé autour de l'IA, pas « VS Code + un
plugin ». En 2026 c'est l'éditeur IA le plus utilisé (plusieurs centaines de
milliers de développeurs, ARR revendiqué en milliards).

Les trois couches, du plus passif au plus autonome :

1. **Tab** : complétion inline ultra-rapide pendant que tu tapes (le modèle
   `cursor-tab`, renforcé par le rachat de **Supermaven** — voir section 86).
2. **Composer** : tu décris un changement en langage naturel, il l'applique
   **sur plusieurs fichiers** (refacto, feature). En 2026, Cursor a aussi
   lancé son **propre modèle de code « Composer »**, réglé pour la vitesse
   dans l'éditeur.
3. **Agent** : exécution autonome multi-étapes (lit, édite, lance des
   commandes, ouvre des PR), avec **Background Agents** (tournent sans
   bloquer ton éditeur) et **Cloud Agents** (tournent sur l'infra Cursor).

Pour qui : tu vis dans l'éditeur et tu veux **le confort maximal** —
indexation du repo, diffs élégants, règles projet, agents d'arrière-plan —
sans assembler toi-même un puzzle CLI + plugins. Le prix : l'enfermement
dans l'éditeur (ton workflow ne te suit pas dans le terminal) et une
facturation par crédits parfois opaque.

## 42. Cursor : installation pas à pas

```bash
# Téléchargement : https://cursor.com/download
# (Windows, macOS, Linux — installeur graphique, pas de CLI d'install officiel)
```

1. Télécharge et installe depuis **cursor.com/download**.
2. Au premier lancement : importe tes extensions et réglages VS Code
   (proposé automatiquement — tes réflexes clavier suivent).
3. Connecte-toi (compte Cursor) et choisis ton plan (section 45).
4. Ouvre ton bac à sable : `File > Open Folder > /tmp/lab-agent`.
5. Vérifie l'indexation : la barre de statut indique l'état d'indexation du
   repo (l'agent « connaît » ton code via cet index — voir section 99 pour
   l'angle confidentialité).

**Fichiers de pilotage** (l'équivalent du CLAUDE.md, version Cursor) :

```
// .cursorrules  (à la racine — format historique, toujours lu)
// ou le dossier moderne :
.cursor/
  rules/          # règles projet (markdown)
  mcp.json        # serveurs MCP : { "mcpServers": { ... } }
```

Exemple de règle projet :

```markdown
# .cursor/rules/python.mdc
---
description: Standards Python du projet
globs: ["**/*.py"]
---
- Python 3.12, type hints obligatoires sur les fonctions publiques.
- Tests pytest dans tests/, un fichier de test par module.
- Docstrings en français. Pas de print() en prod : logging.
```

## 43. Cursor : Tab, Composer, Agent — mode d'emploi

**Tab (complétion)** — le plus rentable au quotidien :

- Tu tapes, une suggestion grise apparaît, `Tab` pour accepter.
- Réglages : `Cursor Settings > Tab` — active/désactive par langage.
- Astuce : Tab est **d'autant meilleur** que ton code est conventionné
  (noms explicites, types) — l'IA devine la suite logique.

**Composer (`Ctrl+I` / `Cmd+I`)** — le multi-fichiers :

