---
id: collect-opencode-docs/opencode-docs/docs-fr-config-md-5e9aa636-2
title: "docs-fr-config-md-5e9aa636"
domain: opencode-docs
role: reference
task: reference
actors: ["AWS"]
dates: []
keywords: ["agent", "agents", "aws", "bedrock", "claude"]
source: docs/RAG/Collect RAG/04_opencode_docs/docs-fr-config-md-5e9aa636.md
source_anchor: ""
source_lines: [156, 353]
sha256: 0149e08278d8fe88469262f5df577805df420ded3effc3cce63f76f3fc2bde50
---

# docs-fr-config-md-5e9aa636

  "$schema": "https://opencode.ai/config.json",
  "provider": {},
  "model": "anthropic/claude-sonnet-4-5",
  "small_model": "anthropic/claude-haiku-4-5"
}
```
L'option `small_model` configure un modèle distinct pour les tâches légères comme la génération de titres. Par défaut, OpenCode essaie d'utiliser un modèle moins cher s'il est disponible auprès de votre fournisseur, sinon il revient à votre modèle principal.
Les options du fournisseur peuvent inclure `timeout` et `setCacheKey` :
```json title="opencode.json"
{
  "$schema": "https://opencode.ai/config.json",
  "provider": {
    "anthropic": {
      "options": {
        "timeout": 600000,
        "setCacheKey": true
      }
    }
  }
}
```
- `timeout` - Délai d'expiration de la demande en millisecondes (par défaut : 300 000). Réglez sur `false` pour désactiver.
- `setCacheKey` - Assurez-vous qu'une clé de cache est toujours définie pour le fournisseur désigné.
Vous pouvez également configurer [modèles locaux](/docs/models#local). [En savoir plus](/docs/models).
---
#### Options spécifiques au fournisseur
Certains fournisseurs prennent en charge des options de configuration supplémentaires au-delà des paramètres génériques `timeout` et `apiKey`.
##### Amazon Bedrock
Amazon Bedrock prend en charge la configuration spécifique à AWS :
```json title="opencode.json"
{
  "$schema": "https://opencode.ai/config.json",
  "provider": {
    "amazon-bedrock": {
      "options": {
        "region": "us-east-1",
        "profile": "my-aws-profile",
        "endpoint": "https://bedrock-runtime.us-east-1.vpce-xxxxx.amazonaws.com"
      }
    }
  }
}
```
- `region` - Région AWS pour Bedrock (par défaut : `AWS_REGION` env var ou `us-east-1`)
- `profile` - Profil nommé AWS de `~/.aws/credentials` (par défaut : `AWS_PROFILE` env var)
- `endpoint` - Point de terminaison personnalisé URL pour les points de terminaison d'un VPC. Il s'agit d'un alias pour l'option générique `baseURL` utilisant la terminologie spécifique à AWS. Si les deux sont spécifiés, `endpoint` est prioritaire.
:::note
Les jetons du porteur (`AWS_BEARER_TOKEN_BEDROCK` ou `/connect`) ont priorité sur l'authentification basée sur le profil. Voir [précédence d'authentification](/docs/providers#authentication-precedence) pour plus de détails.
:::
[En savoir plus sur la configuration d'Amazon Bedrock](/docs/providers#amazon-bedrock).
---
### Thèmes
Définissez votre thème d'interface utilisateur dans `tui.json`.
```json title="tui.json"
{
  "$schema": "https://opencode.ai/tui.json",
  "theme": "tokyonight"
}
```
[En savoir plus ici](/docs/themes).
---
### Agents
Vous pouvez configurer des agents spécialisés pour des tâches spécifiques via l'option `agent`.
```jsonc title="opencode.jsonc"
{
  "$schema": "https://opencode.ai/config.json",
  "agent": {
    "code-reviewer": {
      "description": "Reviews code for best practices and potential issues",
      "model": "anthropic/claude-sonnet-4-5",
      "prompt": "You are a code reviewer. Focus on security, performance, and maintainability.",
      "tools": {
        // Disable file modification tools for review-only agent
        "write": false,
        "edit": false,
      },
    },
  },
}
```
Vous pouvez également définir des agents à l'aide de fichiers markdown dans `~/.config/opencode/agents/` ou `.opencode/agents/`. [En savoir plus ici](/docs/agents).
---
### Agent par défaut
Vous pouvez définir l'agent par défaut à l'aide de l'option `default_agent`. Ceci détermine quel agent est utilisé lorsqu'aucun n'est explicitement spécifié.
```json title="opencode.json"
{
  "$schema": "https://opencode.ai/config.json",
  "default_agent": "plan"
}
```
L'agent par défaut doit être un agent principal (et non un sous-agent). Il peut s'agir d'un agent intégré tel que `"build"` ou `"plan"`, ou d'un [agent personnalisé](/docs/agents) que vous avez défini. Si l'agent spécifié n'existe pas ou est un sous-agent, OpenCode reviendra à `"build"` avec un avertissement.
Ce paramètre s'applique à toutes les interfaces : TUI, CLI (`opencode run`), application de bureau et GitHub Action.
---
### Partage
Vous pouvez configurer la fonctionnalité [share](/docs/share) via l'option `share`.
```json title="opencode.json"
{
  "$schema": "https://opencode.ai/config.json",
  "share": "manual"
}
```
Elle accepte :
- `"manual"` – Autoriser le partage manuel via des commandes (par défaut)
- `"auto"` – Partager automatiquement de nouvelles conversations
- `"disabled"` – Désactiver complètement le partage
Par défaut, le partage est défini en mode manuel où vous devez partager explicitement les conversations à l'aide de la commande `/share`.
---
### Commandes
Vous pouvez configurer des commandes personnalisées pour les tâches répétitives via l'option `command`.
```jsonc title="opencode.jsonc"
{
  "$schema": "https://opencode.ai/config.json",
  "command": {
    "test": {
      "template": "Run the full test suite with coverage report and show any failures.\nFocus on the failing tests and suggest fixes.",
      "description": "Run tests with coverage",
      "agent": "build",
      "model": "anthropic/claude-haiku-4-5",
    },
    "component": {
      "template": "Create a new React component named $ARGUMENTS with TypeScript support.\nInclude proper typing and basic structure.",
      "description": "Create a new component",
    },
  },
}
```
Vous pouvez également définir des commandes à l'aide de fichiers markdown dans `~/.config/opencode/commands/` ou `.opencode/commands/`. [En savoir plus ici](/docs/commands).
---
### Raccourcis clavier
Personnalisez les raccourcis clavier dans `tui.json`.
```json title="tui.json"
{
  "$schema": "https://opencode.ai/tui.json",
  "keybinds": {}
}
```
[En savoir plus ici](/docs/keybinds).
---
### Mise à jour automatique
OpenCode téléchargera automatiquement toutes les nouvelles mises à jour au démarrage. Vous pouvez désactiver cela avec l'option `autoupdate`.
```json title="opencode.json"
{
  "$schema": "https://opencode.ai/config.json",
  "autoupdate": false
}
```
Si vous ne souhaitez pas de mises à jour mais souhaitez être averti lorsqu'une nouvelle version est disponible, définissez `autoupdate` sur `"notify"`.
Notez que cela ne fonctionne que s'il n'a pas été installé à l'aide d'un gestionnaire de packages tel que Homebrew.
---
### Formateurs
Vous pouvez configurer les formateurs de code via l'option `formatter`.
```json title="opencode.json"
{
  "$schema": "https://opencode.ai/config.json",
  "formatter": {
    "prettier": {
      "disabled": true
    },
    "custom-prettier": {
      "command": ["npx", "prettier", "--write", "$FILE"],
      "environment": {
        "NODE_ENV": "development"
      },
      "extensions": [".js", ".ts", ".jsx", ".tsx"]
    }
  }
}
```
[En savoir plus sur les formateurs ici](/docs/formatters).
---
### Autorisations
Par défaut, opencode **autorise toutes les opérations** sans nécessiter d'approbation explicite. Vous pouvez modifier cela en utilisant l'option `permission`.
Par exemple, pour garantir que les outils `edit` et `bash` nécessitent l'approbation de l'utilisateur :
```json title="opencode.json"
{
  "$schema": "https://opencode.ai/config.json",
  "permission": {
    "edit": "ask",
    "bash": "ask"
  }
}
```
[En savoir plus sur les autorisations ici](/docs/permissions).
---
### Compactage
Vous pouvez contrôler le comportement de compactage du contexte via l'option `compaction`.
```json title="opencode.json"
{
  "$schema": "https://opencode.ai/config.json",
  "compaction": {
    "auto": true,
    "prune": false,
    "reserved": 10000
  }
}
```
- `auto` - Compacte automatiquement la session lorsque le contexte est plein (par défaut : `true`).
- `prune` - Supprimez les anciennes sorties de l'outil pour économiser des tokens (par défaut : `false`).
