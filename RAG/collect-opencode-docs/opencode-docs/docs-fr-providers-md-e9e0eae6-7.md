---
id: collect-opencode-docs/opencode-docs/docs-fr-providers-md-e9e0eae6-7
title: "docs-fr-providers-md-e9e0eae6"
domain: opencode-docs
role: reference
task: reference
actors: ["AWS", "Cerebras", "OpenAI"]
dates: []
keywords: ["bedrock"]
source: docs/RAG/Collect RAG/04_opencode_docs/docs-fr-providers-md-e9e0eae6.md
source_anchor: ""
source_lines: [1336, 1418]
sha256: bf9777df5bc31f159943f4eb61c1ba3023c03dd08cd800607c08b52c2e813086
---

# docs-fr-providers-md-e9e0eae6

   ┌  Add credential
   │
   ▲  This only stores a credential for myprovider - you will need to configure it in opencode.json, check the docs for examples.
   │
   ◇  Enter your API key
   │  sk-...
   └
   ```
4. Créez ou mettez à jour votre fichier `opencode.json` dans le répertoire de votre projet :
   ```json title="opencode.json" ""myprovider"" {5-15}
   {
     "$schema": "https://opencode.ai/config.json",
     "provider": {
       "myprovider": {
         "npm": "@ai-sdk/openai-compatible",
         "name": "My AI ProviderDisplay Name",
         "options": {
           "baseURL": "https://api.myprovider.com/v1"
         },
         "models": {
           "my-model-name": {
             "name": "My Model Display Name"
           }
         }
       }
     }
   }
   ```
Voici les options de configuration :
- **npm** : package AI SDK à utiliser, `@ai-sdk/openai-compatible` pour les fournisseurs compatibles OpenAI
- **nom** : nom à afficher dans l'interface utilisateur.
- **modèles** : Modèles disponibles.
- **options.baseURL** : URL de l'endpoint API.
- **options.apiKey** : définissez éventuellement la clé API, si vous n'utilisez pas d'authentification.
- **options.headers** : définissez éventuellement des en-têtes personnalisés.
En savoir plus sur les options avancées dans l'exemple ci-dessous.
5. Exécutez la commande `/models` et votre fournisseur et vos modèles personnalisés apparaîtront dans la liste de sélection.
---
##### Exemple
Voici un exemple de définition des options `apiKey`, `headers` et modèle `limit`.
```json title="opencode.json" {9,11,17-20}
{
  "$schema": "https://opencode.ai/config.json",
  "provider": {
    "myprovider": {
      "npm": "@ai-sdk/openai-compatible",
      "name": "My AI ProviderDisplay Name",
      "options": {
        "baseURL": "https://api.myprovider.com/v1",
        "apiKey": "{env:ANTHROPIC_API_KEY}",
        "headers": {
          "Authorization": "Bearer custom-token"
        }
      },
      "models": {
        "my-model-name": {
          "name": "My Model Display Name",
          "limit": {
            "context": 200000,
            "output": 65536
          }
        }
      }
    }
  }
}
```
Détails de configuration :
- **apiKey** : défini à l'aide de la syntaxe de variable `env`, [en savoir plus](/docs/config#env-vars).
- **headers** : en-têtes personnalisés envoyés avec chaque requête.
- **limit.context** : nombre maximal de jetons d'entrée acceptés par le modèle.
- **limit.output** : nombre maximal de jetons que le modèle peut générer.
Les champs `limit` permettent à OpenCode de comprendre la quantité de contexte qu'il vous reste. Les fournisseurs standard les extraient automatiquement de models.dev.
---
## Dépannage
Si vous rencontrez des difficultés lors de la configuration d'un fournisseur, vérifiez les points suivants :
1. **Vérifiez la configuration de l'authentification** : exécutez `opencode auth list` pour voir si les informations d'identification
   pour le fournisseur sont ajoutés à votre configuration.
Cela ne s'applique pas aux fournisseurs comme Amazon Bedrock, qui s'appuient sur des variables d'environnement pour leur authentification.
2. Pour les fournisseurs personnalisés, vérifiez la configuration opencode et :
   - Assurez-vous que l'ID du fournisseur utilisé dans la commande `/connect` correspond à l'ID de votre configuration opencode.
   - Le bon package npm est utilisé pour le fournisseur. Par exemple, utilisez `@ai-sdk/cerebras` pour Cerebras. Et pour tous les autres fournisseurs compatibles OpenAI, utilisez `@ai-sdk/openai-compatible`.
   - Vérifiez que le point de terminaison API correct est utilisé dans le champ `options.baseURL`.
