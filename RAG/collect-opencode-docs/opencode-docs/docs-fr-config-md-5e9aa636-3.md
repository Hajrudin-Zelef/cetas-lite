---
id: collect-opencode-docs/opencode-docs/docs-fr-config-md-5e9aa636-3
title: "docs-fr-config-md-5e9aa636"
domain: opencode-docs
role: reference
task: reference
actors: []
dates: []
keywords: ["gemini", "mcp"]
source: docs/RAG/Collect RAG/04_opencode_docs/docs-fr-config-md-5e9aa636.md
source_anchor: ""
source_lines: [354, 483]
sha256: 4f6a9ad99978264179681ae8b58cf32572d5c855a7511bf9c1730daed02cc02d
---

# docs-fr-config-md-5e9aa636

- `reserved` - Tampon de jetons pour le compactage. Laisse suffisamment de marge pour éviter le débordement lors du compactage.
---
### Observateur
Vous pouvez configurer les modèles d'ignorance de l'observateur de fichiers via l'option `watcher`.
```json title="opencode.json"
{
  "$schema": "https://opencode.ai/config.json",
  "watcher": {
    "ignore": ["node_modules/**", "dist/**", ".git/**"]
  }
}
```
Les modèles suivent la syntaxe glob. Utilisez ceci pour exclure les répertoires bruyants de la surveillance des fichiers.
---
### Serveurs MCP
Vous pouvez configurer les serveurs MCP que vous souhaitez utiliser via l'option `mcp`.
```json title="opencode.json"
{
  "$schema": "https://opencode.ai/config.json",
  "mcp": {}
}
```
[En savoir plus ici](/docs/mcp-servers).
---
### Extensions
[Plugins](/docs/plugins) étendent OpenCode avec des outils, des hooks et des intégrations personnalisés.
Placez les fichiers du plugin dans `.opencode/plugins/` ou `~/.config/opencode/plugins/`. Vous pouvez également charger des plugins depuis npm via l'option `plugin`.
```json title="opencode.json"
{
  "$schema": "https://opencode.ai/config.json",
  "plugin": ["opencode-helicone-session", "@my-org/custom-plugin"]
}
```
[En savoir plus ici](/docs/plugins).
---
### Instructions
Vous pouvez configurer les instructions pour le modèle que vous utilisez via l'option `instructions`.
```json title="opencode.json"
{
  "$schema": "https://opencode.ai/config.json",
  "instructions": ["CONTRIBUTING.md", "docs/guidelines.md", ".cursor/rules/*.md"]
}
```
Cela prend un tableau de chemins et de modèles globaux vers les fichiers d'instructions. [En savoir plus sur les règles ici](/docs/rules).
---
### Fournisseurs désactivés
Vous pouvez désactiver les fournisseurs chargés automatiquement via l'option `disabled_providers`. Ceci est utile lorsque vous souhaitez empêcher le chargement de certains fournisseurs même si leurs informations d'identification sont disponibles.
```json title="opencode.json"
{
  "$schema": "https://opencode.ai/config.json",
  "disabled_providers": ["openai", "gemini"]
}
```
:::note
Le `disabled_providers` est prioritaire sur `enabled_providers`.
:::
L'option `disabled_providers` accepte un tableau d'ID de fournisseur. Lorsqu'un fournisseur est désactivé :
- Il ne sera pas chargé même si des variables d'environnement sont définies.
- Il ne sera pas chargé même si les clés API sont configurées via la commande `/connect`.
- Les modèles du fournisseur n'apparaîtront pas dans la liste de sélection des modèles.
---
### Fournisseurs activés
Vous pouvez spécifier une liste autorisée de fournisseurs via l'option `enabled_providers`. Une fois défini, seuls les fournisseurs spécifiés seront activés et tous les autres seront ignorés.
```json title="opencode.json"
{
  "$schema": "https://opencode.ai/config.json",
  "enabled_providers": ["anthropic", "openai"]
}
```
Ceci est utile lorsque vous souhaitez restreindre OpenCode à l'utilisation de fournisseurs spécifiques plutôt que de les désactiver un par un.
:::note
Le `disabled_providers` est prioritaire sur `enabled_providers`.
:::
Si un fournisseur apparaît à la fois dans `enabled_providers` et `disabled_providers`, le `disabled_providers` est prioritaire pour la compatibilité ascendante.
---
### Expérimental
La clé `experimental` contient des options en cours de développement actif.
```json title="opencode.json"
{
  "$schema": "https://opencode.ai/config.json",
  "experimental": {}
}
```
:::caution
Les options expérimentales ne sont pas stables. Elles peuvent changer ou être supprimées sans préavis.
:::
---
## Variables
Vous pouvez utiliser la substitution de variables dans vos fichiers de configuration pour référencer les variables d'environnement et le contenu des fichiers.
---
### Variables d'environnement
Utilisez `{env:VARIABLE_NAME}` pour remplacer les variables d'environnement :
```json title="opencode.json"
{
  "$schema": "https://opencode.ai/config.json",
  "model": "{env:OPENCODE_MODEL}",
  "provider": {
    "anthropic": {
      "models": {},
      "options": {
        "apiKey": "{env:ANTHROPIC_API_KEY}"
      }
    }
  }
}
```
Si la variable d'environnement n'est pas définie, elle sera remplacée par une chaîne vide.
---
### Fichiers
Utilisez `{file:path/to/file}` pour remplacer le contenu d'un fichier :
```json title="opencode.json"
{
  "$schema": "https://opencode.ai/config.json",
  "instructions": ["./custom-instructions.md"],
  "provider": {
    "openai": {
      "options": {
        "apiKey": "{file:~/.secrets/openai-key}"
      }
    }
  }
}
```
Les chemins de fichiers peuvent être :
- Relatifs au répertoire du fichier de configuration
- Ou des chemins absolus commençant par `/` ou `~`
Elles sont utiles pour :
- Conserver les données sensibles telles que les clés API dans des fichiers séparés.
- Inclure de gros fichiers d'instructions sans encombrer votre configuration.
- Partager des extraits de configuration communs sur plusieurs fichiers de configuration.
