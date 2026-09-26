---
id: collect-opencode-docs/opencode-docs/docs-fr-providers-md-e9e0eae6-4
title: "docs-fr-providers-md-e9e0eae6"
domain: opencode-docs
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "Groq", "Hugging Face", "Moonshot", "OpenAI"]
dates: []
keywords: ["claude", "glm", "inference", "kimi", "llama", "llama.cpp"]
source: docs/RAG/Collect RAG/04_opencode_docs/docs-fr-providers-md-e9e0eae6.md
source_anchor: ""
source_lines: [649, 843]
sha256: ecc194e4301d2289c85238ce9f1f8faef02783b21a9a4cd616f3edbea2d2c5f7
---

# docs-fr-providers-md-e9e0eae6

La région `global` améliore la disponibilité et réduit les erreurs sans frais supplémentaires. Utilisez des points de terminaison régionaux (par exemple, `us-central1`) pour les exigences de résidence des données. [En savoir plus](https://cloud.google.com/vertex-ai/generative-ai/docs/partner-models/use-partner-models#regional_and_global_endpoints)
:::
3. Exécutez la commande `/models` pour sélectionner le modèle souhaité.
   ```txt
   /models
   ```
---
### Groq
1. Rendez-vous sur la [console Groq](https://console.groq.com/), cliquez sur **Créer une clé API** et copiez la clé.
2. Exécutez la commande `/connect` et recherchez Groq.
   ```txt
   /connect
   ```
3. Saisissez la clé API du fournisseur.
   ```txt
   ┌ API key
   │
   │
   └ enter
   ```
4. Exécutez la commande `/models` pour sélectionner celle que vous souhaitez.
   ```txt
   /models
   ```
---
### Hugging Face
[Hugging Face Inference Providers](https://huggingface.co/docs/inference-providers) donne accès à des modèles ouverts pris en charge par plus de 17 fournisseurs.
1. Rendez-vous sur [Hugging Face settings](https://huggingface.co/settings/tokens/new?ownUserPermissions=inference.serverless.write&tokenType=fineGrained) pour créer un jeton avec l'autorisation de passer des appels aux fournisseurs d'inférence.
2. Exécutez la commande `/connect` et recherchez **Hugging Face**.
   ```txt
   /connect
   ```
3. Entrez votre jeton Hugging Face.
   ```txt
   ┌ API key
   │
   │
   └ enter
   ```
4. Exécutez la commande `/models` pour sélectionner un modèle comme _Kimi-K2-Instruct_ ou _GLM-4.6_.
   ```txt
   /models
   ```
---
### Helicone
[Helicone](https://helicone.ai) est une plate-forme d'observabilité LLM qui fournit la journalisation, la surveillance et l'analyse de vos applications d'IA. L'Helicone AI Gateway achemine automatiquement vos demandes vers le fournisseur approprié en fonction du modèle.
1. Rendez-vous sur [Helicone](https://helicone.ai), créez un compte et générez une clé API à partir de votre tableau de bord.
2. Exécutez la commande `/connect` et recherchez **Helicone**.
   ```txt
   /connect
   ```
3. Entrez votre clé Helicone API.
   ```txt
   ┌ API key
   │
   │
   └ enter
   ```
4. Exécutez la commande `/models` pour sélectionner un modèle.
   ```txt
   /models
   ```
Pour plus de fournisseurs et de fonctionnalités avancées telles que la mise en cache et la limitation de débit, consultez la [documentation Helicone](https://docs.helicone.ai).
#### Configurations facultatives
Si vous voyez une fonctionnalité ou un modèle d'Helicone qui n'est pas configuré automatiquement via opencode, vous pouvez toujours le configurer vous-même.
Voici le [Répertoire des modèles d'Helicone](https://helicone.ai/models), vous en aurez besoin pour récupérer les identifiants des modèles que vous souhaitez ajouter.
```jsonc title="~/.config/opencode/opencode.jsonc"
{
  "$schema": "https://opencode.ai/config.json",
  "provider": {
    "helicone": {
      "npm": "@ai-sdk/openai-compatible",
      "name": "Helicone",
      "options": {
        "baseURL": "https://ai-gateway.helicone.ai",
      },
      "models": {
        "gpt-4o": {
          // Model ID (from Helicone's model directory page)
          "name": "GPT-4o", // Your own custom name for the model
        },
        "claude-sonnet-4-20250514": {
          "name": "Claude Sonnet 4",
        },
      },
    },
  },
}
```
#### En-têtes personnalisés
Helicone prend en charge les en-têtes personnalisés pour des fonctionnalités telles que la mise en cache, le suivi des utilisateurs et la gestion des sessions. Ajoutez-les à la configuration de votre fournisseur en utilisant `options.headers` :
```jsonc title="~/.config/opencode/opencode.jsonc"
{
  "$schema": "https://opencode.ai/config.json",
  "provider": {
    "helicone": {
      "npm": "@ai-sdk/openai-compatible",
      "name": "Helicone",
      "options": {
        "baseURL": "https://ai-gateway.helicone.ai",
        "headers": {
          "Helicone-Cache-Enabled": "true",
          "Helicone-User-Id": "opencode",
        },
      },
    },
  },
}
```
##### Suivi des sessions
La fonctionnalité [Sessions](https://docs.helicone.ai/features/sessions) d'Helicone vous permet de regrouper les requêtes LLM associées. Utilisez le plugin [opencode-helicone-session](https://github.com/H2Shami/opencode-helicone-session) pour enregistrer automatiquement chaque conversation OpenCode en tant que session dans Helicone.
```bash
npm install -g opencode-helicone-session
```
Ajoutez-le à votre configuration.
```json title="opencode.json"
{
  "plugin": ["opencode-helicone-session"]
}
```
Le plugin injecte les en-têtes `Helicone-Session-Id` et `Helicone-Session-Name` dans vos requêtes. Sur la page Sessions d'Helicone, vous verrez chaque conversation OpenCode répertoriée comme une session distincte.
##### En-têtes Helicone communs
| En-tête                    | Descriptif                                                                           |
| -------------------------- | ------------------------------------------------------------------------------------ |
| `Helicone-Cache-Enabled`   | Activer la mise en cache des réponses (`true`/`false`)                               |
| `Helicone-User-Id`         | Suivre les métriques par utilisateur                                                 |
| `Helicone-Property-[Name]` | Ajouter des propriétés personnalisées (par exemple, `Helicone-Property-Environment`) |
| `Helicone-Prompt-Id`       | Associer les requêtes aux versions d'invite                                          |
Consultez le [Helicone Header Directory](https://docs.helicone.ai/helicone-headers/header-directory) pour tous les en-têtes disponibles.
---
### llama.cpp
Vous pouvez configurer opencode pour utiliser des modèles locaux via l'utilitaire llama-server de [llama.cpp's](https://github.com/ggml-org/llama.cpp)
```json title="opencode.json" "llama.cpp" {5, 6, 8, 10-15}
{
  "$schema": "https://opencode.ai/config.json",
  "provider": {
    "llama.cpp": {
      "npm": "@ai-sdk/openai-compatible",
      "name": "llama-server (local)",
      "options": {
        "baseURL": "http://127.0.0.1:8080/v1"
      },
      "models": {
        "qwen3-coder:a3b": {
          "name": "Qwen3-Coder: a3b-30b (local)",
          "limit": {
            "context": 128000,
            "output": 65536
          }
        }
      }
    }
  }
}
```
Dans cet exemple :
- `llama.cpp` est l'ID du fournisseur personnalisé. Cela peut être n’importe quelle chaîne de votre choix.
- `npm` spécifie le package à utiliser pour ce fournisseur. Ici, `@ai-sdk/openai-compatible` est utilisé pour tout API compatible OpenAI.
- `name` est le nom d'affichage du fournisseur dans l'interface utilisateur.
- `options.baseURL` est le point de terminaison du serveur local.
- `models` est une carte des ID de modèle avec leurs configurations. Le nom du modèle sera affiché dans la liste de sélection du modèle.
---
### IO.NET
IO.NET propose 17 modèles optimisés pour différents cas d'utilisation :
1. Rendez-vous sur la [console IO.NET](https://ai.io.net/), créez un compte et générez une clé API.
2. Exécutez la commande `/connect` et recherchez **IO.NET**.
   ```txt
   /connect
   ```
3. Entrez votre clé IO.NET API.
   ```txt
   ┌ API key
   │
   │
   └ enter
   ```
4. Exécutez la commande `/models` pour sélectionner un modèle.
   ```txt
   /models
   ```
---
### LM Studio
Vous pouvez configurer opencode pour utiliser des modèles locaux via LM Studio.
```json title="opencode.json" "lmstudio" {5, 6, 8, 10-14}
{
  "$schema": "https://opencode.ai/config.json",
  "provider": {
    "lmstudio": {
      "npm": "@ai-sdk/openai-compatible",
      "name": "LM Studio (local)",
      "options": {
        "baseURL": "http://127.0.0.1:1234/v1"
      },
      "models": {
        "google/gemma-3n-e4b": {
