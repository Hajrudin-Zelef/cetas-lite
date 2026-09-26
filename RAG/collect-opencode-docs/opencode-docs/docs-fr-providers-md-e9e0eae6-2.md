---
id: collect-opencode-docs/opencode-docs/docs-fr-providers-md-e9e0eae6-2
title: "docs-fr-providers-md-e9e0eae6"
domain: opencode-docs
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "Baseten", "Cerebras", "DeepSeek", "Moonshot", "OpenAI"]
dates: []
keywords: ["claude", "deepseek", "foundry", "inference", "kimi", "qwen", "sol", "tool calling"]
source: docs/RAG/Collect RAG/04_opencode_docs/docs-fr-providers-md-e9e0eae6.md
source_anchor: ""
source_lines: [208, 430]
sha256: c5129b862df435bacb34a1deac99540e1bc84d12d43a27eba3a3555c331873a2
---

# docs-fr-providers-md-e9e0eae6

   ┌ Select auth method
   │
   │ Claude Pro/Max
   │ Create an API Key
   │ Manually enter API Key
   └
   ```
3. Désormais, tous les modèles Anthropic devraient être disponibles lorsque vous utilisez la commande `/models`.
   ```txt
   /models
   ```
:::info
L'utilisation de votre abonnement Claude Pro/Max dans OpenCode n'est pas officiellement prise en charge par [Anthropic](https://anthropic.com).
:::
##### Utilisation des clés API
Vous pouvez également sélectionner **Créer une clé API** si vous n'avez pas d'abonnement Pro/Max. Il ouvrira également votre navigateur et vous demandera de vous connecter à Anthropic et vous donnera un code que vous pourrez coller dans votre terminal.
Ou si vous disposez déjà d'une clé API, vous pouvez sélectionner **Entrer manuellement la clé API** et la coller dans votre terminal.
---
### Atomic Chat
Vous pouvez configurer opencode pour utiliser des modèles locaux via [Atomic Chat](https://atomic.chat), une application de bureau qui exécute des LLM locaux derrière un serveur API compatible OpenAI (point de terminaison par défaut `http://127.0.0.1:1337/v1`).
```json title="opencode.json" "atomic-chat" {5, 6, 8, 10-14}
{
  "$schema": "https://opencode.ai/config.json",
  "provider": {
    "atomic-chat": {
      "npm": "@ai-sdk/openai-compatible",
      "name": "Atomic Chat (local)",
      "options": {
        "baseURL": "http://127.0.0.1:1337/v1"
      },
      "models": {
        "<your-model-id>": {
          "name": "<your-model-name>"
        }
      }
    }
  }
}
```
Dans cet exemple :
- `atomic-chat` est l'identifiant personnalisé du fournisseur. Il peut s'agir de n'importe quelle chaîne.
- `npm` spécifie le paquet à utiliser pour ce fournisseur. Ici, `@ai-sdk/openai-compatible` est utilisé pour toute API compatible OpenAI.
- `name` est le nom du fournisseur affiché dans l'interface.
- `options.baseURL` est le point de terminaison du serveur local. Modifiez l'hôte et le port selon votre configuration Atomic Chat.
- `models` est une carte d'ID de modèles vers leurs noms d'affichage. Chaque ID doit correspondre à l'`id` renvoyé par `GET /v1/models` — exécutez `curl http://127.0.0.1:1337/v1/models` pour lister les ID actuellement chargés dans Atomic Chat.
:::tip
Si les appels d'outils ne fonctionnent pas bien, choisissez un modèle chargé avec un bon support du tool calling (par exemple, une variante Qwen-Coder ou DeepSeek-Coder).
:::
---
### Azure OpenAI
:::note
Si vous rencontrez des erreurs « Je suis désolé, mais je ne peux pas vous aider avec cette demande », essayez de modifier le filtre de contenu de **DefaultV2** à **Default** dans votre ressource Azure.
:::
1. Rendez-vous sur le [portail Azure](https://portal.azure.com/) et créez une ressource **Azure OpenAI**. Vous aurez besoin de :
   - **Nom de la ressource** : cela fait partie de votre point de terminaison API (`https://RESOURCE_NAME.openai.azure.com/`)
   - **Clé API** : soit `KEY 1` ou `KEY 2` de votre ressource
2. Accédez à [Azure AI Foundry](https://ai.azure.com/) et déployez un modèle.
   :::note
   Le nom du déploiement doit correspondre au nom du modèle pour que opencode fonctionne correctement.
   :::
3. Exécutez la commande `/connect` et recherchez **Azure**.
   ```txt
   /connect
   ```
4. Entrez votre clé API.
   ```txt
   ┌ API key
   │
   │
   └ enter
   ```
5. Définissez le nom de votre ressource comme variable d'environnement :
   ```bash
   AZURE_RESOURCE_NAME=XXX opencode
   ```
Ou ajoutez-le à votre profil bash :
```bash title="~/.bash_profile"
   export AZURE_RESOURCE_NAME=XXX
```
6. Exécutez la commande `/models` pour sélectionner votre modèle déployé.
   ```txt
   /models
   ```
---
### Azure Cognitive Services
1. Rendez-vous sur le [portail Azure](https://portal.azure.com/) et créez une ressource **Azure OpenAI**. Vous aurez besoin de :
   - **Nom de la ressource** : cela fait partie de votre point de terminaison API (`https://AZURE_COGNITIVE_SERVICES_RESOURCE_NAME.cognitiveservices.azure.com/`)
   - **Clé API** : soit `KEY 1` ou `KEY 2` de votre ressource
2. Accédez à [Azure AI Foundry](https://ai.azure.com/) et déployez un modèle.
   :::note
   Le nom du déploiement doit correspondre au nom du modèle pour que opencode fonctionne correctement.
   :::
3. Exécutez la commande `/connect` et recherchez **Azure Cognitive Services**.
   ```txt
   /connect
   ```
4. Entrez votre clé API.
   ```txt
   ┌ API key
   │
   │
   └ enter
   ```
5. Définissez le nom de votre ressource comme variable d'environnement :
   ```bash
   AZURE_COGNITIVE_SERVICES_RESOURCE_NAME=XXX opencode
   ```
Ou ajoutez-le à votre profil bash :
```bash title="~/.bash_profile"
   export AZURE_COGNITIVE_SERVICES_RESOURCE_NAME=XXX
```
6. Exécutez la commande `/models` pour sélectionner votre modèle déployé.
   ```txt
   /models
   ```
---
### Baseten
1. Rendez-vous sur [Baseten](https://app.baseten.co/), créez un compte et générez une clé API.
2. Exécutez la commande `/connect` et recherchez **Baseten**.
   ```txt
   /connect
   ```
3. Entrez votre clé Baseten API.
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
### Cerebras
1. Rendez-vous sur la [console Cerebras](https://inference.cerebras.ai/), créez un compte et générez une clé API.
2. Exécutez la commande `/connect` et recherchez **Cerebras**.
   ```txt
   /connect
   ```
3. Entrez votre clé Cerebras API.
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
### Cloudflare AI Gateway
Cloudflare AI Gateway vous permet d'accéder aux modèles de OpenAI, Anthropic, Workers AI et bien plus encore via un point de terminaison unifié. Avec [Unified Billing](https://developers.cloudflare.com/ai-gateway/features/unified-billing/), vous n'avez pas besoin de clés API distinctes pour chaque fournisseur.
1. Rendez-vous sur le [tableau de bord Cloudflare](https://dash.cloudflare.com/), accédez à **AI** > **AI Gateway** et créez une nouvelle passerelle.
2. Définissez votre ID de compte et votre ID de passerelle comme variables d'environnement.
   ```bash title="~/.bash_profile"
   export CLOUDFLARE_ACCOUNT_ID=your-32-character-account-id
   export CLOUDFLARE_GATEWAY_ID=your-gateway-id
   ```
3. Exécutez la commande `/connect` et recherchez **Cloudflare AI Gateway**.
   ```txt
   /connect
   ```
4. Entrez votre jeton Cloudflare API.
   ```txt
   ┌ API key
   │
   │
   └ enter
   ```
Ou définissez-le comme variable d'environnement.
```bash title="~/.bash_profile"
   export CLOUDFLARE_API_TOKEN=your-api-token
```
5. Exécutez la commande `/models` pour sélectionner un modèle.
   ```txt
   /models
   ```
Vous pouvez également ajouter des modèles via votre configuration opencode.
```json title="opencode.json"
{
  "$schema": "https://opencode.ai/config.json",
  "provider": {
    "cloudflare-ai-gateway": {
      "models": {
        "openai/gpt-4o": {},
        "anthropic/claude-sonnet-4": {}
      }
    }
  }
}
```
---
### Cortecs
1. Rendez-vous sur la [console Cortecs](https://cortecs.ai/), créez un compte et générez une clé API.
2. Exécutez la commande `/connect` et recherchez **Cortecs**.
   ```txt
   /connect
   ```
3. Entrez votre clé Cortecs API.
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
### DeepSeek
1. Rendez-vous sur la [console DeepSeek](https://platform.deepseek.com/), créez un compte et cliquez sur **Créer une nouvelle clé API**.
2. Exécutez la commande `/connect` et recherchez **DeepSeek**.
   ```txt
   /connect
   ```
3. Entrez votre clé DeepSeek API.
   ```txt
   ┌ API key
   │
   │
