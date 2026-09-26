---
id: collect-opencode-docs/opencode-docs/docs-fr-config-md-5e9aa636-1
title: "docs-fr-config-md-5e9aa636"
domain: opencode-docs
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "agents", "claude", "mcp"]
source: docs/RAG/Collect RAG/04_opencode_docs/docs-fr-config-md-5e9aa636.md
source_anchor: ""
source_lines: [1, 155]
sha256: 361a61f2fdc3acaf0216203c7ad216d9e929ee696ca2a116a1f74169afcd5785
---

# docs-fr-config-md-5e9aa636

Vous pouvez configurer OpenCode à l'aide d'un fichier de configuration JSON.
---
## Format
OpenCode prend en charge les formats **JSON** et **JSONC** (JSON avec commentaires).
```jsonc title="opencode.jsonc"
{
  "$schema": "https://opencode.ai/config.json",
  "model": "anthropic/claude-sonnet-4-5",
  "autoupdate": true,
  "server": {
    "port": 4096,
  },
}
```
---
## Emplacements
Vous pouvez placer votre configuration à plusieurs emplacements différents et ils ont un
ordre de priorité différent.
:::note
Les fichiers de configuration sont **fusionnés**, pas remplacés.
:::
Les fichiers de configuration sont fusionnés et non remplacés. Les paramètres des emplacements de configuration suivants sont combinés. Les configurations ultérieures remplacent les précédentes uniquement en cas de clés en conflit. Les paramètres non conflictuels de toutes les configurations sont conservés.
Par exemple, si votre configuration globale définit `autoupdate: true` et que la configuration de votre projet définit `model: "anthropic/claude-sonnet-4-5"`, la configuration finale inclura les deux paramètres.
---
### Ordre de priorité
Les sources de configuration sont chargées dans cet ordre (les sources ultérieures remplacent les précédentes) :
1. **Configuration distante** (à partir de `.well-known/opencode`) - paramètres par défaut de l'organisation
2. **Configuration globale** (`~/.config/opencode/opencode.json`) - préférences utilisateur
3. **Configuration personnalisée** (`OPENCODE_CONFIG` env var) - remplacements personnalisés
4. **Configuration du projet** (`opencode.json` dans le projet) - paramètres spécifiques au projet
5. **Répertoires `.opencode`** - agents, commandes, plugins
6. **Configuration en ligne** (`OPENCODE_CONFIG_CONTENT` env var) - remplacements d'exécution
Cela signifie que les configurations de projet peuvent remplacer les valeurs par défaut globales, et que les configurations globales peuvent remplacer les valeurs par défaut de l'organisation distante.
:::note
Les répertoires `.opencode` et `~/.config/opencode` utilisent des **noms au pluriel** pour les sous-répertoires : `agents/`, `commands/`, `modes/`, `plugins/`, `skills/`, `tools/` et `themes/`. Les noms singuliers (par exemple, `agent/`) sont également pris en charge pour une compatibilité ascendante.
:::
---
### Configuration distante
Les organisations peuvent fournir une configuration par défaut via le point de terminaison `.well-known/opencode`. Ceci est récupéré automatiquement lorsque vous vous authentifiez auprès d’un fournisseur qui le prend en charge.
La configuration distante est chargée en premier, servant de couche de base. Toutes les autres sources de configuration (globales, projet) peuvent remplacer ces valeurs par défaut.
Par exemple, si votre organisation fournit des serveurs MCP qui sont désactivés par défaut :
```json title="Remote config from .well-known/opencode"
{
  "mcp": {
    "jira": {
      "type": "remote",
      "url": "https://jira.example.com/mcp",
      "enabled": false
    }
  }
}
```
Vous pouvez activer des serveurs spécifiques dans votre configuration locale :
```json title="opencode.json"
{
  "mcp": {
    "jira": {
      "type": "remote",
      "url": "https://jira.example.com/mcp",
      "enabled": true
    }
  }
}
```
---
### Globale
Placez votre configuration globale OpenCode dans `~/.config/opencode/opencode.json`. Utilisez la configuration globale pour les préférences de l'utilisateur telles que les fournisseurs, les modèles et les autorisations.
Pour les paramètres spécifiques à TUI, utilisez `~/.config/opencode/tui.json`.
La configuration globale remplace les paramètres par défaut de l'organisation distante.
---
### Par projet
Ajoutez `opencode.json` à la racine de votre projet. La configuration du projet a la priorité la plus élevée parmi les fichiers de configuration standard : elle remplace les configurations globales et distantes.
Pour les paramètres TUI spécifiques au projet, ajoutez `tui.json` à côté.
:::tip
Placez la configuration spécifique au projet à la racine de votre projet.
:::
Lorsque OpenCode démarre, il recherche un fichier de configuration dans le répertoire actuel ou remonte jusqu'au répertoire Git le plus proche.
Il peut également être archivé en toute sécurité dans Git et utilise le même schéma que le schéma global.
---
### Chemin personnalisé
Spécifiez un chemin de fichier de configuration personnalisé à l'aide de la variable d'environnement `OPENCODE_CONFIG`.
```bash
export OPENCODE_CONFIG=/path/to/my/custom-config.json
opencode run "Hello world"
```
La configuration personnalisée est chargée entre les configurations globales et celles du projet dans l'ordre de priorité.
---
### Répertoire personnalisé
Spécifiez un répertoire de configuration personnalisé à l'aide de `OPENCODE_CONFIG_DIR` variable d'environnement. Ce répertoire sera recherché pour les agents, les commandes, modes et plugins tout comme le répertoire standard `.opencode`, et devrait suivre la même structure.
```bash
export OPENCODE_CONFIG_DIR=/path/to/my/config-directory
opencode run "Hello world"
```
Le répertoire personnalisé est chargé après les répertoires de configuration globale et `.opencode`, il **peut donc remplacer** leurs paramètres.
---
## Schéma
Le fichier de configuration a un schéma défini dans [**`opencode.ai/config.json`**](https://opencode.ai/config.json).
La configuration TUI utilise [**`opencode.ai/tui.json`**](https://opencode.ai/tui.json).
Votre éditeur doit être capable de valider et de compléter automatiquement en fonction du schéma.
---
### TUI
Utilisez un fichier dédié `tui.json` (ou `tui.jsonc`) pour les paramètres spécifiques à TUI.
```json title="tui.json"
{
  "$schema": "https://opencode.ai/tui.json",
  "scroll_speed": 3,
  "scroll_acceleration": {
    "enabled": true
  },
  "diff_style": "auto"
}
```
Utilisez `OPENCODE_TUI_CONFIG` pour pointer vers un fichier de configuration TUI personnalisé.
Les anciennes clés `theme`, `keybinds` et `tui` dans `opencode.json` sont obsolètes et migrées automatiquement lorsque cela est possible.
[En savoir plus sur l'utilisation du TUI ici](/docs/tui#configure).
---
### Serveur
Vous pouvez configurer les paramètres du serveur pour les commandes `opencode serve` et `opencode web` via l'option `server`.
```json title="opencode.json"
{
  "$schema": "https://opencode.ai/config.json",
  "server": {
    "port": 4096,
    "hostname": "0.0.0.0",
    "mdns": true,
    "mdnsDomain": "myproject.local",
    "cors": ["http://localhost:5173"]
  }
}
```
Options disponibles :
- `port` - Port d'écoute.
- `hostname` - Nom d'hôte d'écoute. Lorsque `mdns` est activé et qu'aucun nom d'hôte n'est défini, la valeur par défaut est `0.0.0.0`.
- `mdns` - Activer la découverte du service mDNS. Cela permet à d'autres appareils du réseau de découvrir votre serveur OpenCode.
- `mdnsDomain` - Nom de domaine personnalisé pour le service mDNS. La valeur par défaut est `opencode.local`. Utile pour exécuter plusieurs instances sur le même réseau.
- `cors` - Origines supplémentaires pour autoriser CORS lors de l'utilisation du serveur HTTP à partir d'un client basé sur un navigateur. Les valeurs doivent être des origines complètes (schéma + hôte + port facultatif), par exemple `https://app.example.com`.
[En savoir plus sur le serveur ici](/docs/server).
---
### Outils
Vous pouvez gérer les outils qu'un LLM peut utiliser via l'option `tools`.
```json title="opencode.json"
{
  "$schema": "https://opencode.ai/config.json",
  "tools": {
    "write": false,
    "bash": false
  }
}
```
[En savoir plus sur les outils ici](/docs/tools).
---
### Modèles
Vous pouvez configurer les fournisseurs et les modèles que vous souhaitez utiliser dans votre configuration OpenCode via les options `provider`, `model` et `small_model`.
```json title="opencode.json"
{
