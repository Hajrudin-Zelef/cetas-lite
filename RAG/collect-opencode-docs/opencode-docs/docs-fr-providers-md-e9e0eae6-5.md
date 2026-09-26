---
id: collect-opencode-docs/opencode-docs/docs-fr-providers-md-e9e0eae6-5
title: "docs-fr-providers-md-e9e0eae6"
domain: opencode-docs
role: reference
task: reference
actors: ["AWS", "Alibaba", "Anthropic", "Google", "Meta", "MiniMax", "Mistral", "Moonshot", "Nebius", "OpenAI", "OpenRouter"]
dates: []
keywords: ["chatgpt", "kimi", "llama", "mistral", "pricing", "qwen"]
source: docs/RAG/Collect RAG/04_opencode_docs/docs-fr-providers-md-e9e0eae6.md
source_anchor: ""
source_lines: [844, 1086]
sha256: 2b0e054bc811d6de2582bed7002a94305913db639b937315a9730451ee41b4ad
---

# docs-fr-providers-md-e9e0eae6

          "name": "Gemma 3n-e4b (local)"
        }
      }
    }
  }
}
```
Dans cet exemple :
- `lmstudio` est l'ID du fournisseur personnalisé. Cela peut être n’importe quelle chaîne de votre choix.
- `npm` spécifie le package à utiliser pour ce fournisseur. Ici, `@ai-sdk/openai-compatible` est utilisé pour tout API compatible OpenAI.
- `name` est le nom d'affichage du fournisseur dans l'interface utilisateur.
- `options.baseURL` est le point de terminaison du serveur local.
- `models` est une carte des ID de modèle avec leurs configurations. Le nom du modèle sera affiché dans la liste de sélection du modèle.
---
### Moonshot AI
Pour utiliser Kimi K2 de Moonshot AI :
1. Rendez-vous sur la [console Moonshot AI](https://platform.moonshot.ai/console), créez un compte et cliquez sur **Créer une clé API**.
2. Exécutez la commande `/connect` et recherchez **Moonshot AI**.
   ```txt
   /connect
   ```
3. Entrez votre clé Moonshot API.
   ```txt
   ┌ API key
   │
   │
   └ enter
   ```
4. Exécutez la commande `/models` pour sélectionner _Kimi K2_.
   ```txt
   /models
   ```
---
### MiniMax
1. Rendez-vous sur la [Console MiniMax API](https://platform.minimax.io/login), créez un compte et générez une clé API.
2. Exécutez la commande `/connect` et recherchez **MiniMax**.
   ```txt
   /connect
   ```
3. Entrez votre clé MiniMax API.
   ```txt
   ┌ API key
   │
   │
   └ enter
   ```
4. Exécutez la commande `/models` pour sélectionner un modèle tel que _M2.1_.
   ```txt
   /models
   ```
---
### Nebius Token Factory
1. Rendez-vous sur la [console Nebius Token Factory](https://tokenfactory.nebius.com/), créez un compte et cliquez sur **Ajouter une clé**.
2. Exécutez la commande `/connect` et recherchez **Nebius Token Factory**.
   ```txt
   /connect
   ```
3. Entrez votre clé Nebius Token Factory API.
   ```txt
   ┌ API key
   │
   │
   └ enter
   ```
4. Exécutez la commande `/models` pour sélectionner un modèle tel que _Kimi K2 Instruct_.
   ```txt
   /models
   ```
---
### Ollama
Vous pouvez configurer opencode pour utiliser des modèles locaux via Ollama.
:::tip
Ollama peut se configurer automatiquement pour OpenCode. Voir les [documents d'intégration Ollama](https://docs.ollama.com/integrations/opencode) pour plus de détails.
:::
```json title="opencode.json" "ollama" {5, 6, 8, 10-14}
{
  "$schema": "https://opencode.ai/config.json",
  "provider": {
    "ollama": {
      "npm": "@ai-sdk/openai-compatible",
      "name": "Ollama (local)",
      "options": {
        "baseURL": "http://localhost:11434/v1"
      },
      "models": {
        "llama2": {
          "name": "Llama 2"
        }
      }
    }
  }
}
```
Dans cet exemple :
- `ollama` est l'ID du fournisseur personnalisé. Cela peut être n’importe quelle chaîne de votre choix.
- `npm` spécifie le package à utiliser pour ce fournisseur. Ici, `@ai-sdk/openai-compatible` est utilisé pour tout API compatible OpenAI.
- `name` est le nom d'affichage du fournisseur dans l'interface utilisateur.
- `options.baseURL` est le point de terminaison du serveur local.
- `models` est une carte des ID de modèle avec leurs configurations. Le nom du modèle sera affiché dans la liste de sélection du modèle.
:::tip
Si les appels d'outils ne fonctionnent pas, essayez d'augmenter `num_ctx` dans Ollama. Commencez vers 16k - 32k.
:::
---
### Ollama Cloud
Pour utiliser Ollama Cloud avec OpenCode :
1. Rendez-vous sur [https://ollama.com/](https://ollama.com/) et connectez-vous ou créez un compte.
2. Accédez à **Paramètres** > **Clés** et cliquez sur **Ajouter une clé API** pour générer une nouvelle clé API.
3. Copiez la clé API à utiliser dans OpenCode.
4. Exécutez la commande `/connect` et recherchez **Ollama Cloud**.
   ```txt
   /connect
   ```
5. Entrez votre clé Ollama Cloud API.
   ```txt
   ┌ API key
   │
   │
   └ enter
   ```
6. **Important** : Avant d'utiliser des modèles cloud dans OpenCode, vous devez extraire les informations du modèle localement :
   ```bash
   ollama pull gpt-oss:20b-cloud
   ```
7. Exécutez la commande `/models` pour sélectionner votre modèle Ollama Cloud.
   ```txt
   /models
   ```
---
### OpenAI
Nous vous recommandons de vous inscrire à [ChatGPT Plus ou Pro](https://chatgpt.com/pricing).
1. Une fois inscrit, exécutez la commande `/connect` et sélectionnez OpenAI.
   ```txt
   /connect
   ```
2. Ici, vous pouvez sélectionner l'option **ChatGPT Plus/Pro** et cela ouvrira votre navigateur.
   et vous demande de vous authentifier.
   ```txt
   ┌ Select auth method
   │
   │ ChatGPT Plus/Pro
   │ Manually enter API Key
   └
   ```
3. Désormais, tous les modèles OpenAI devraient être disponibles lorsque vous utilisez la commande `/models`.
   ```txt
   /models
   ```
##### Utilisation des clés API
Si vous disposez déjà d'une clé API, vous pouvez sélectionner **Entrer manuellement la clé API** et la coller dans votre terminal.
---
### OpenCode Zen
OpenCode Zen est une liste de modèles testés et vérifiés fournie par l'équipe OpenCode. [En savoir plus](/docs/zen).
1. Connectez-vous à **<a href={console}>OpenCode Zen</a>** et cliquez sur **Créer une clé API**.
2. Exécutez la commande `/connect` et recherchez **OpenCode Zen**.
   ```txt
   /connect
   ```
3. Entrez votre clé OpenCode API.
   ```txt
   ┌ API key
   │
   │
   └ enter
   ```
4. Exécutez la commande `/models` pour sélectionner un modèle tel que _Qwen 3 Coder 480B_.
   ```txt
   /models
   ```
---
### OpenRouter
1. Rendez-vous sur le [tableau de bord OpenRouter](https://openrouter.ai/settings/keys), cliquez sur **Créer une clé API** et copiez la clé.
2. Exécutez la commande `/connect` et recherchez OpenRouter.
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
4. De nombreux modèles OpenRouter sont préchargés par défaut, exécutez la commande `/models` pour sélectionner celui que vous souhaitez.
   ```txt
   /models
   ```
Vous pouvez également ajouter des modèles supplémentaires via votre configuration opencode.
```json title="opencode.json" {6}
{
  "$schema": "https://opencode.ai/config.json",
  "provider": {
    "openrouter": {
      "models": {
        "somecoolnewmodel": {}
      }
    }
  }
}
```
5. Vous pouvez également les personnaliser via votre configuration opencode. Voici un exemple de spécification d'un fournisseur
   ```json title="opencode.json"
   {
     "$schema": "https://opencode.ai/config.json",
     "provider": {
       "openrouter": {
         "models": {
           "moonshotai/kimi-k2": {
             "options": {
               "provider": {
                 "order": ["baseten"],
                 "allow_fallbacks": false
               }
             }
           }
         }
       }
     }
   }
   ```
---
### SAP AI Core
SAP AI Core donne accès à plus de 40 modèles de OpenAI, Anthropic, Google, Amazon, Meta, Mistral et AI21 via une plateforme unifiée.
1. Accédez à votre [SAP BTP Cockpit](https://account.hana.ondemand.com/), accédez à votre instance de service SAP AI Core et créez une clé de service.
   :::tip
   La clé de service est un objet JSON contenant `clientid`, `clientsecret`, `url` et `serviceurls.AI_API_URL`. Vous pouvez trouver votre instance AI Core sous **Services** > **Instances et abonnements** dans le cockpit BTP.
   :::
2. Exécutez la commande `/connect` et recherchez **SAP AI Core**.
   ```txt
   /connect
   ```
3. Entrez votre clé de service JSON.
   ```txt
   ┌ Service key
   │
   │
   └ enter
   ```
Ou définissez la variable d'environnement `AICORE_SERVICE_KEY` :
```bash
   AICORE_SERVICE_KEY='{"clientid":"...","clientsecret":"...","url":"...","serviceurls":{"AI_API_URL":"..."}}' opencode
```
Ou ajoutez-le à votre profil bash :
```bash title="~/.bash_profile"
