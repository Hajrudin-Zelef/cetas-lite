---
id: collect-opencode-docs/opencode-docs/docs-fr-providers-md-e9e0eae6-6
title: "docs-fr-providers-md-e9e0eae6"
domain: opencode-docs
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "Google", "Mistral", "Moonshot", "OpenAI", "Together AI", "Z.ai", "xAI"]
dates: []
keywords: ["claude", "glm", "grok", "kimi", "llama", "mistral", "qwen"]
source: docs/RAG/Collect RAG/04_opencode_docs/docs-fr-providers-md-e9e0eae6.md
source_anchor: ""
source_lines: [1087, 1335]
sha256: cd5ffb5a1151cfc3173f2e65ccce4cc35dcf00ee5932e79e9e75d20860eccd7d
---

# docs-fr-providers-md-e9e0eae6

   export AICORE_SERVICE_KEY='{"clientid":"...","clientsecret":"...","url":"...","serviceurls":{"AI_API_URL":"..."}}'
```
4. Définissez éventuellement l'ID de déploiement et le groupe de ressources :
   ```bash
   AICORE_DEPLOYMENT_ID=your-deployment-id AICORE_RESOURCE_GROUP=your-resource-group opencode
   ```
   :::note
   Ces paramètres sont facultatifs et doivent être configurés en fonction de votre configuration SAP AI Core.
   :::
5. Exécutez la commande `/models` pour sélectionner parmi plus de 40 modèles disponibles.
   ```txt
   /models
   ```
---
### STACKIT
STACKIT AI Model Serving fournit un environnement d'hébergement souverain entièrement géré pour les modèles d'IA, se concentrant sur les LLM comme Llama, Mistral et Qwen, avec une souveraineté maximale des données sur l'infrastructure européenne.
1. Rendez-vous sur le [portail STACKIT](https://portal.stackit.cloud), accédez à **AI Model Serving** et créez un jeton d'authentification pour votre projet.
   :::tip
   Vous avez besoin d'un compte client STACKIT, d'un compte utilisateur et d'un projet avant de créer des jetons d'authentification.
   :::
2. Exécutez la commande `/connect` et recherchez **STACKIT**.
   ```txt
   /connect
   ```
3. Entrez votre jeton d'authentification STACKIT AI Model Serving.
   ```txt
   ┌ API key
   │
   │
   └ enter
   ```
4. Exécutez la commande `/models` pour sélectionner parmi les modèles disponibles tels que _Qwen3-VL 235B_ ou _Llama 3.3 70B_.
   ```txt
   /models
   ```
---
### OVHcloud AI Endpoints
1. Rendez-vous sur le [Panneau OVHcloud](https://ovh.com/manager). Accédez à la section `Public Cloud`, `AI & Machine Learning` > `AI Endpoints` et dans l'onglet `API Keys`, cliquez sur **Créer une nouvelle clé API**.
2. Exécutez la commande `/connect` et recherchez **OVHcloud AI Endpoints**.
   ```txt
   /connect
   ```
3. Saisissez votre clé OVHcloud AI Endpoints API.
   ```txt
   ┌ API key
   │
   │
   └ enter
   ```
4. Exécutez la commande `/models` pour sélectionner un modèle tel que _gpt-oss-120b_.
   ```txt
   /models
   ```
---
### Scaleway
Pour utiliser [Scaleway Generative APIs](https://www.scaleway.com/en/docs/generative-apis/) avec Opencode :
1. Rendez-vous dans les [Paramètres IAM de la console Scaleway](https://console.scaleway.com/iam/api-keys) pour générer une nouvelle clé API.
2. Exécutez la commande `/connect` et recherchez **Scaleway**.
   ```txt
   /connect
   ```
3. Entrez votre clé Scaleway API.
   ```txt
   ┌ API key
   │
   │
   └ enter
   ```
4. Exécutez la commande `/models` pour sélectionner un modèle tel que _devstral-2-123b-instruct-2512_ ou _gpt-oss-120b_.
   ```txt
   /models
   ```
---
### Together AI
1. Rendez-vous sur [Together AI console](https://api.together.ai), créez un compte et cliquez sur **Ajouter une clé**.
2. Exécutez la commande `/connect` et recherchez **Together AI**.
   ```txt
   /connect
   ```
3. Entrez votre clé Together AI API.
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
### Venice AI
1. Rendez-vous sur la [Venice AI console](https://venice.ai), créez un compte et générez une clé API.
2. Exécutez la commande `/connect` et recherchez **Venice AI**.
   ```txt
   /connect
   ```
3. Entrez votre clé Venise AI API.
   ```txt
   ┌ API key
   │
   │
   └ enter
   ```
4. Exécutez la commande `/models` pour sélectionner un modèle tel que _Llama 3.3 70B_.
   ```txt
   /models
   ```
---
### Vercel AI Gateway
Vercel AI Gateway vous permet d'accéder aux modèles de OpenAI, Anthropic, Google, xAI et plus encore via un point de terminaison unifié. Les modèles sont proposés au prix catalogue sans majoration.
1. Rendez-vous sur le [tableau de bord Vercel](https://vercel.com/), accédez à l'onglet **AI Gateway** et cliquez sur **API Keys** pour créer une nouvelle clé API.
2. Exécutez la commande `/connect` et recherchez **Vercel AI Gateway**.
   ```txt
   /connect
   ```
3. Entrez votre clé Vercel AI Gateway API.
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
Vous pouvez également personnaliser les modèles via votre configuration opencode. Voici un exemple de spécification de l'ordre de routage du fournisseur.
```json title="opencode.json"
{
  "$schema": "https://opencode.ai/config.json",
  "provider": {
    "vercel": {
      "models": {
        "anthropic/claude-sonnet-4": {
          "options": {
            "order": ["anthropic", "vertex"]
          }
        }
      }
    }
  }
}
```
Quelques options de routage utiles :
| Options             | Descriptif                                                                                  |
| ------------------- | ------------------------------------------------------------------------------------------- |
| `order`             | Séquence de fournisseur à essayer                                                           |
| `only`              | Restreindre à des fournisseurs spécifiques                                                  |
| `zeroDataRetention` | Utilisez uniquement des fournisseurs avec des politiques de conservation des données nulles |
---
### xAI
1. Rendez-vous sur la [console xAI](https://console.x.ai/), créez un compte et générez une clé API.
2. Exécutez la commande `/connect` et recherchez **xAI**.
   ```txt
   /connect
   ```
3. Entrez votre clé xAI API.
   ```txt
   ┌ API key
   │
   │
   └ enter
   ```
4. Exécutez la commande `/models` pour sélectionner un modèle tel que _Grok Beta_.
   ```txt
   /models
   ```
---
### Z.AI
1. Rendez-vous sur la [console Z.AI API](https://z.ai/manage-apikey/apikey-list), créez un compte et cliquez sur **Créer une nouvelle clé API**.
2. Exécutez la commande `/connect` et recherchez **Z.AI**.
   ```txt
   /connect
   ```
Si vous êtes abonné au **Plan de codage GLM**, sélectionnez **Plan de codage Z.AI**.
3. Entrez votre clé Z.AI API.
   ```txt
   ┌ API key
   │
   │
   └ enter
   ```
4. Exécutez la commande `/models` pour sélectionner un modèle tel que _GLM-4.7_.
   ```txt
   /models
   ```
---
### ZenMux
1. Rendez-vous sur le [tableau de bord ZenMux](https://zenmux.ai/settings/keys), cliquez sur **Créer une clé API** et copiez la clé.
2. Exécutez la commande `/connect` et recherchez ZenMux.
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
4. De nombreux modèles ZenMux sont préchargés par défaut, exécutez la commande `/models` pour sélectionner celui que vous souhaitez.
   ```txt
   /models
   ```
Vous pouvez également ajouter des modèles supplémentaires via votre configuration opencode.
```json title="opencode.json" {6}
{
  "$schema": "https://opencode.ai/config.json",
  "provider": {
    "zenmux": {
      "models": {
        "somecoolnewmodel": {}
      }
    }
  }
}
```
---
## Fournisseur personnalisé
Pour ajouter un fournisseur **OpenAI-compatible** qui n'est pas répertorié dans la commande `/connect` :
:::tip
Vous pouvez utiliser n'importe quel fournisseur compatible OpenAI avec opencode. La plupart des fournisseurs d'IA modernes proposent des API compatibles OpenAI.
:::
1. Exécutez la commande `/connect` et faites défiler jusqu'à **Autre**.
   ```bash
   $ /connect
   ┌  Add credential
   │
   ◆  Select provider
   │  ...
   │  ● Other
   └
   ```
2. Saisissez un identifiant unique pour le fournisseur.
   ```bash
   $ /connect
   ┌  Add credential
   │
   ◇  Enter provider id
   │  myprovider
   └
   ```
   :::note
   Choisissez un identifiant mémorable, vous l'utiliserez dans votre fichier de configuration.
   :::
3. Entrez votre clé API pour le fournisseur.
   ```bash
   $ /connect
