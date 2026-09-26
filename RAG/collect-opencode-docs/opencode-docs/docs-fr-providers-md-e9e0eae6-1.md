---
id: collect-opencode-docs/opencode-docs/docs-fr-providers-md-e9e0eae6-1
title: "docs-fr-providers-md-e9e0eae6"
domain: opencode-docs
role: reference
task: reference
actors: ["AWS", "Anthropic"]
dates: []
keywords: ["aws", "bedrock", "claude", "inference"]
source: docs/RAG/Collect RAG/04_opencode_docs/docs-fr-providers-md-e9e0eae6.md
source_anchor: ""
source_lines: [1, 207]
sha256: 1c33182b0a39aa92fb1c5037ee59856904905a2d494c7674fd149dd4d33989bb
---

# docs-fr-providers-md-e9e0eae6

import config from "../../../../config.mjs"
export const console = config.console
OpenCode utilise [AI SDK](https://ai-sdk.dev/) et [Models.dev](https://models.dev) pour prendre en charge **75+ fournisseurs LLM** et prend en charge l'exécution de modèles locaux.
Pour ajouter un fournisseur, vous devez :
1. Ajoutez les clés API pour le fournisseur à l'aide de la commande `/connect`.
2. Configurez le fournisseur dans votre configuration OpenCode.
---
### Informations d'identification
Lorsque vous ajoutez les clés API d'un fournisseur avec la commande `/connect`, elles sont stockées
en `~/.local/share/opencode/auth.json`.
---
### Configuration
Vous pouvez personnaliser les fournisseurs via la section `provider` de votre OpenCode
configuration.
---
#### Socle URL
Vous pouvez personnaliser la base URL pour n'importe quel fournisseur en définissant l'option `baseURL`. Ceci est utile lors de l'utilisation de services proxy ou de points de terminaison personnalisés.
```json title="opencode.json" {6}
{
  "$schema": "https://opencode.ai/config.json",
  "provider": {
    "anthropic": {
      "options": {
        "baseURL": "https://api.anthropic.com/v1"
      }
    }
  }
}
```
---
## OpenCode Zen
OpenCode Zen est une liste de modèles fournis par l'équipe OpenCode qui ont été
testé et vérifié pour fonctionner correctement avec OpenCode. [En savoir plus](/docs/zen).
:::tip
Si vous êtes nouveau, nous vous recommandons de commencer par OpenCode Zen.
:::
1. Exécutez la commande `/connect` dans le TUI, sélectionnez opencode et dirigez-vous vers [opencode.ai/auth](https://opencode.ai/auth).
   ```txt
   /connect
   ```
2. Connectez-vous, ajoutez vos informations de facturation et copiez votre clé API.
3. Collez votre clé API.
   ```txt
   ┌ API key
   │
   │
   └ enter
   ```
4. Exécutez `/models` dans le TUI pour voir la liste des modèles que nous recommandons.
   ```txt
   /models
   ```
Il fonctionne comme n’importe quel autre fournisseur dans OpenCode et son utilisation est totalement facultative.
---
## OpenCode Go
OpenCode Go est un plan d'abonnement à faible coût qui offre un accès fiable aux modèles de codage ouverts populaires fournis par l'équipe OpenCode qui ont été
testé et vérifié pour fonctionner correctement avec OpenCode.
1. Exécutez la commande `/connect` dans le TUI, sélectionnez `OpenCode Go` et rendez-vous sur [opencode.ai/auth](https://opencode.ai/zen).
   ```txt
   /connect
   ```
2. Connectez-vous, ajoutez vos informations de facturation et copiez votre clé API.
3. Collez votre clé API.
   ```txt
   ┌ API key
   │
   │
   └ enter
   ```
4. Exécutez `/models` dans le TUI pour voir la liste des modèles que nous recommandons.
   ```txt
   /models
   ```
Il fonctionne comme n’importe quel autre fournisseur dans OpenCode et son utilisation est totalement facultative.
---
## Annuaire
Examinons certains fournisseurs en détail. Si vous souhaitez ajouter un fournisseur au
liste, n'hésitez pas à ouvrir un PR.
:::note
Vous ne voyez pas de fournisseur ici ? Soumettez un PR.
:::
---
### 302.AI
1. Rendez-vous sur la [console 302.AI](https://302.ai/), créez un compte et générez une clé API.
2. Exécutez la commande `/connect` et recherchez **302.AI**.
   ```txt
   /connect
   ```
3. Saisissez votre clé 302.AI API.
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
### Amazon Bedrock
Pour utiliser Amazon Bedrock avec OpenCode :
1. Rendez-vous sur le **Catalogue de modèles** dans la console Amazon Bedrock et demandez
   accédez aux modèles que vous souhaitez.
   :::tip
   Vous devez avoir accès au modèle souhaité dans Amazon Bedrock.
   :::
2. **Configurez l'authentification** à l'aide de l'une des méthodes suivantes :
   #### Variables d'environnement (démarrage rapide)
Définissez l'une de ces variables d'environnement lors de l'exécution de opencode :
```bash
   # Option 1: Using AWS access keys
   AWS_ACCESS_KEY_ID=XXX AWS_SECRET_ACCESS_KEY=YYY opencode
   # Option 2: Using named AWS profile
   AWS_PROFILE=my-profile opencode
   # Option 3: Using Bedrock bearer token
   AWS_BEARER_TOKEN_BEDROCK=XXX opencode
```
Ou ajoutez-les à votre profil bash :
```bash title="~/.bash_profile"
   export AWS_PROFILE=my-dev-profile
   export AWS_REGION=us-east-1
```
#### Fichier de configuration (recommandé)
Pour une configuration spécifique au projet ou persistante, utilisez `opencode.json` :
```json title="opencode.json"
{
  "$schema": "https://opencode.ai/config.json",
  "provider": {
    "amazon-bedrock": {
      "options": {
        "region": "us-east-1",
        "profile": "my-aws-profile"
      }
    }
  }
}
```
**Options disponibles :**
- `region` - Région AWS (par exemple, `us-east-1`, `eu-west-1`)
- `profile` - Profil nommé AWS de `~/.aws/credentials`
- `endpoint` - URL de point de terminaison personnalisée pour les endpoints VPC (alias de l'option générique `baseURL`)
:::tip
Les options du fichier de configuration sont prioritaires sur les variables d'environnement.
:::
#### Avancé : points de terminaison d'un VPC
Si vous utilisez des points de terminaison d'un VPC pour Bedrock :
```json title="opencode.json"
{
  "$schema": "https://opencode.ai/config.json",
  "provider": {
    "amazon-bedrock": {
      "options": {
        "region": "us-east-1",
        "profile": "production",
        "endpoint": "https://bedrock-runtime.us-east-1.vpce-xxxxx.amazonaws.com"
      }
    }
  }
}
```
:::note
L'option `endpoint` est un alias pour l'option générique `baseURL`, utilisant la terminologie spécifique à AWS. Si `endpoint` et `baseURL` sont spécifiés, `endpoint` est prioritaire.
:::
#### Méthodes d'authentification
- **`AWS_ACCESS_KEY_ID` / `AWS_SECRET_ACCESS_KEY`** : créez un utilisateur IAM et générez des clés d'accès dans la console AWS
- **`AWS_PROFILE`** : utilisez les profils nommés de `~/.aws/credentials`. Configurez d'abord avec `aws configure --profile my-profile` ou `aws sso login`
- **`AWS_BEARER_TOKEN_BEDROCK`** : Générez des clés API à long terme à partir de la console Amazon Bedrock
- **`AWS_WEB_IDENTITY_TOKEN_FILE` / `AWS_ROLE_ARN`** : pour EKS IRSA (rôles IAM pour les comptes de service) ou d'autres environnements Kubernetes avec fédération OIDC. Ces variables d'environnement sont automatiquement injectées par Kubernetes lors de l'utilisation des annotations de compte de service.
#### Priorité d'authentification
Amazon Bedrock utilise la priorité d'authentification suivante :
1.  **Bearer Token** - Variable d'environnement `AWS_BEARER_TOKEN_BEDROCK` ou jeton de la commande `/connect`
2.  **AWS Credential Chain** - Profil, clés d'accès, informations d'identification partagées, rôles IAM, jetons d'identité Web (EKS IRSA), métadonnées d'instance
:::note
Lorsqu'un jeton de porteur est défini (via `/connect` ou `AWS_BEARER_TOKEN_BEDROCK`), il est prioritaire sur toutes les méthodes d'identification AWS, y compris les profils configurés.
:::
3. Exécutez la commande `/models` pour sélectionner le modèle souhaité.
   ```txt
   /models
   ```
:::note
Pour les profils d'inférence personnalisés, utilisez le nom du modèle et du fournisseur dans la clé et définissez la propriété `id` sur l'arn. Cela garantit une mise en cache correcte :
```json title="opencode.json"
{
  "$schema": "https://opencode.ai/config.json",
  "provider": {
    "amazon-bedrock": {
      // ...
      "models": {
        "anthropic-claude-sonnet-4.5": {
          "id": "arn:aws:bedrock:us-east-1:xxx:application-inference-profile/yyy"
        }
      }
    }
  }
}
```
:::
---
### Anthropic
1. Une fois inscrit, exécutez la commande `/connect` et sélectionnez Anthropic.
   ```txt
   /connect
   ```
2. Ici, vous pouvez sélectionner l'option **Claude Pro/Max** et cela ouvrira votre navigateur.
   et vous demande de vous authentifier.
   ```txt
