---
id: collect-261001-ia-llm/ia-llm/ia-agents-codeurs-3
title: "Les agents codeurs IA"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "OpenAI", "OpenRouter"]
dates: []
keywords: ["agent", "agents", "claude", "cloud agent", "gpu", "mai", "mcp", "open source", "pricing", "qwen"]
source: docs/RAG/collect-261001-ia-llm/ia_agents_codeurs.md
source_anchor: ""
source_lines: [342, 524]
sha256: 0bd9305d8a7a0cd7430931b50fc1ade12118bfa8bbabe549a8ed65a6acc06b29
---

# Les agents codeurs IA

| Touche / commande | Effet |
|---|---|
| `Tab` | Bascule entre l'agent **Plan** (lecture seule) et **Build** (écriture) |
| `@` | Référence un fichier/dossier dans ton prompt (`@app/main.py`) |
| `/init` | Génère un `AGENTS.md` à partir de l'analyse du projet |
| `/connect` | Connecter/configurer un fournisseur de modèle |
| `/login` | Connexion OAuth quand le fournisseur le propose |
| `!` | Préfixe pour lancer une commande shell directement |
| `opencode run "prompt"` | Mode non interactif (scriptable, CI) |
| `opencode run -m <modèle> "prompt"` | Choisir le modèle pour un run |

Exemple de choix de modèle par tâche (à adapter selon tes clés) :

```bash
# Exploration / questions : modèle rapide et pas cher
opencode run -m openrouter/qwen/qwen3-coder-flash "liste les modules et leurs dépendances"

# Implémentation sérieuse : modèle costaud
opencode run -m anthropic/claude-sonnet-4-6 "implémente la fonction dans @app/main.py"
```

> ⚠️ Les identifiants de modèles changent vite (les fournisseurs renomment).
> Liste ce qui est réellement disponible avec ta config avant de copier un
> nom de modèle d'un tuto de 3 mois.

## 12. OpenCode : premier workflow commenté

Scénario : tu veux qu'OpenCode **documente** ton bac à sable puis ajoute une
fonction utilitaire testée. On reste en mode Plan d'abord.

```bash
cd /tmp/lab-agent
opencode
```

Dans la TUI :

```
# 1. Contextualiser le projet (génère AGENTS.md)
> /init
# → OpenCode explore le repo et propose un AGENTS.md. RELIS-LE avant d'accepter :
#    corrige les commandes de test, ajoute tes interdictions (section 6).

# 2. Passer en mode Plan (Tab) pour une exploration sans écriture
> [Tab pour passer en Plan] @. explique l'architecture de ce projet en 10 lignes,
>   sans modifier aucun fichier

# 3. Demander un plan d'implémentation avant d'écrire
> propose un plan pour ajouter une fonction ping_hote(hote) avec retry,
>   sans écrire de code pour l'instant

# 4. Si le plan te convient : Tab pour passer en Build, puis
> implémente le plan validé, ajoute un test pytest, lance les tests
```

Ce que tu observes : l'agent lit les fichiers (`@`), propose, écrit, lance
`pytest`, corrige. **Tu valides chaque étape sensible** : c'est le workflow
« plan d'abord » qui évite 80 % des bêtises.

En mode non interactif (pour un script cron, par ex. génération de doc
nuit) :

```bash
opencode run "génère docs/API.md à partir des docstrings de app/, ne modifie rien d'autre" \
  && git status --porcelain
# Si git status montre des fichiers inattendus → git checkout -- . pour annuler
```

## 13. OpenCode : modèles supportés et configuration

Fichier de config : `~/.config/opencode/opencode.json` (global) ou
`opencode.json` à la racine du projet. Exemple réel :

```json
{
  "$schema": "https://opencode.ai/config.json",
  "agent": {
    "build": { "model": "openrouter/qwen/qwen3-coder-flash" },
    "plan":  { "model": "openrouter/qwen/qwen3-coder-flash" }
  },
  "mcp": {
    "filesystem": {
      "type": "local",
      "command": ["npx", "-y", "@modelcontextprotocol/server-filesystem", "/data/projets"]
    }
  }
}
```

Points clés :

- **Modèles différents par agent** : un modèle pas cher en Plan (exploration),
  un costaud en Build. C'est le levier n°1 d'économie.
- **Fournisseurs** : toute clé compatible OpenAI/Anthropic, OpenRouter
  (agrégateur, pratique pour tester plein de modèles avec une seule clé),
  Ollama local.
- **OpenCode Zen** : l'offre « modèles testés et curatés » d'OpenCode, avec
  un modèle gratuit intégré pour démarrer sans clé (les offres gratuites
  tournent — à vérifier sur opencode.ai/docs).
- **MCP natif** : serveurs MCP déclarés dans la config, outils disponibles
  dans la TUI.
- **Plugins** : écosystème de plugins (hooks, skills) documenté sur
  opencode.ai/docs/plugins.

## 14. OpenCode : prix (vérifiés sept. 2026)

| Poste | Coût |
|---|---|
| Le logiciel | **0 €** — MIT, gratuit, à vie |
| OpenCode Zen / offre hébergée | Formule autour de **~10 $/mois** citée dans les comparatifs (à vérifier sur opencode.ai — l'offre a évolué en 2026) |
| BYOK (ta clé API) | Tu paies le fournisseur au token : ex. quelques $ pour des sessions légères, 20–50 $/mois en usage intensif sur modèles frontier |
| 100 % local (Ollama) | 0 € de tokens, coût = ton électricité et ton GPU |

Règle pratique : **en BYOK avec un modèle « flash/mini » pour le plan et un
gros modèle pour le build, un mois d'usage régulier coûte souvent moins cher
qu'un abonnement fixe à 20 $** — mais sans plafond : mets une alerte de
facturation chez ton fournisseur (section 98).

## 15. OpenCode : points forts / faiblesses honnêtes

**Points forts**

- Liberté totale de modèle : tu n'es marié à personne, tu arbitres
  qualité/prix/confidentialité par tâche.
- Open source MIT : auditable, auto-hébergeable, énorme communauté
  (plugins, skills, docs communautaires).
- TUI soignée + `opencode run` scriptable : aussi à l'aise en interactif
  qu'en cron/CI.
- Plan/Build via `Tab` : le garde-fou est dans le geste, pas dans la doc.
- `/init` → `AGENTS.md` : l'onboarding d'un repo existant prend 2 minutes.

**Faiblesses honnêtes**

- **La qualité dépend du modèle que TU choisis.** Un mauvais modèle =
  un mauvais agent, et OpenCode ne te protège pas de ton propre choix.
- Jeune projet qui bouge très vite : commandes, config et offres changent
  entre deux versions — la doc de 6 mois est déjà périmée.
- Pas de « tout inclus » : pas d'indexation codebase magique offerte, pas de
  cloud agent clé en main comme Cursor.
- L'écosystème plugins/skills est riche mais inégal : à trier.

## 16. OpenCode : bonnes pratiques

1. **Fige tes modèles dans `opencode.json` versionné** : toute l'équipe (ou
   toi dans 3 mois) rejoue avec les mêmes modèles.
2. **Sépare plan/build par modèle** (section 13) : l'économie la plus simple.
3. **Toujours `/init` sur un repo existant**, puis corrige le AGENTS.md généré
   (il invente parfois des commandes de test).
4. **En CI** : `opencode run` avec un modèle fixe, prompt versionné, et
   `git diff --exit-code` après pour détecter les changements inattendus.
5. **MCP avec parcimonie** : chaque serveur MCP ajoute des outils donc des
   tokens à chaque tour. Ne branche que ce qui sert la tâche.

## 17. Kilo Code : c'est quoi, pour qui

**Kilo Code** (kilo.ai) est un agent codeur **open source (MIT)** né en mars
2025 comme **fork de Roo Code** (lui-même dérivé de Cline). En avril 2026, il
a été **reconstruit sur la base d'OpenCode** : le même agent tourne désormais
dans **VS Code, JetBrains et en CLI**, avec une config partagée.

Repères (vérifiés sept. 2026) :

- **Éditeur** : Kilo-Org ; **rachat par Anaconda en juillet 2026**.
- **Adoption** : >1,5 million d'utilisateurs revendiqués, l'un des agents
  open source qui croît le plus vite ; point d'atterrissage naturel des
  utilisateurs de **Roo Code (archivé en mai 2026)**.
- **Modèle éco** : agent gratuit et open source ; tu paies l'usage modèle
  (BYOK sans marge, ou crédits prépayés « Kilo Gateway », ou abonnement
  **Kilo Pass à partir de ~19 $/mois** — à vérifier sur kilo.ai/pricing).
- **Particularités** : modes personnalisés (Architect, Code, Debug, Ask
  hérités de Roo), **autocomplétion inline** (rare pour un agent VS Code),
  **Agent Manager** (sessions parallèles isolées dans des git worktrees
  séparés), marketplace MCP intégré, routage automatique de modèles
  (planification/codage/debug vers le bon modèle selon le palier choisi).

Pour qui : tu veux **un seul agent open source sur tous tes environnements**
(éditeur + terminal), avec une vraie gestion de sessions parallèles, sans
t'enfermer chez un fournisseur.

## 18. Kilo Code : installation pas à pas

**Extension VS Code / JetBrains :**

