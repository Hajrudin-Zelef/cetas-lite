---
id: collect-opencode-docs/opencode-docs/docs-fr-providers-md-e9e0eae6-3
title: "docs-fr-providers-md-e9e0eae6"
domain: opencode-docs
role: reference
task: reference
actors: ["Anthropic", "DeepSeek", "Fireworks AI", "Google", "Microsoft", "Moonshot"]
dates: []
keywords: ["agent", "agents", "claude", "copilot", "deepseek", "kimi"]
source: docs/RAG/Collect RAG/04_opencode_docs/docs-fr-providers-md-e9e0eae6.md
source_anchor: ""
source_lines: [431, 648]
sha256: ee09d2c2c50905eec8254b085001992a0fe84d90758ca0c7324988bc327b6753
---

# docs-fr-providers-md-e9e0eae6

   └ enter
   ```
4. Exécutez la commande `/models` pour sélectionner un modèle DeepSeek tel que _DeepSeek V4 Pro_.
   ```txt
   /models
   ```
---
### Deep Infra
1. Rendez-vous sur le [tableau de bord Deep Infra](https://deepinfra.com/dash), créez un compte et générez une clé API.
2. Exécutez la commande `/connect` et recherchez **Deep Infra**.
   ```txt
   /connect
   ```
3. Entrez votre clé Deep Infra API.
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
### FrogBot
1. Rendez-vous sur le [Tableau de bord du micrologiciel](https://app.frogbot.ai/signup), créez un compte et générez une clé API.
2. Exécutez la commande `/connect` et recherchez **FrogBot**.
   ```txt
   /connect
   ```
3. Entrez la clé API de votre micrologiciel.
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
### Fireworks AI
1. Rendez-vous sur la [console Fireworks AI](https://app.fireworks.ai/), créez un compte et cliquez sur **Créer une clé API**.
2. Exécutez la commande `/connect` et recherchez **Fireworks AI**.
   ```txt
   /connect
   ```
3. Entrez votre clé Fireworks AI API.
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
### GitLab Duo
GitLab Duo fournit un chat agent basé sur l'IA avec des capacités d'appel d'outils natives via le proxy Anthropic de GitLab.
1. Exécutez la commande `/connect` et sélectionnez GitLab.
   ```txt
   /connect
   ```
2. Choisissez votre méthode d'authentification :
   ```txt
   ┌ Select auth method
   │
   │ OAuth (Recommended)
   │ Personal Access Token
   └
   ```
   #### Utilisation de OAuth (recommandé)
Sélectionnez **OAuth** et votre navigateur s'ouvrira pour autorisation.
#### Utilisation d'un jeton d'accès personnel
1.  Accédez à [GitLab Paramètres utilisateur > Jetons d'accès](https://gitlab.com/-/user_settings/personal_access_tokens)
2.  Cliquez sur **Ajouter un nouveau jeton**
3.  Nom : `OpenCode`, Portées : `api`
4.  Copiez le jeton (commence par `glpat-`)
5.  Entrez-le dans le terminal
6.  Exécutez la commande `/models` pour voir les modèles disponibles.
    ```txt
    /models
    ```
````
Trois modèles basés sur Claude sont disponibles :
- **duo-chat-haiku-4-5** (Par défaut) - Réponses rapides pour des tâches rapides
- **duo-chat-sonnet-4-5** - Performances équilibrées pour la plupart des flux de travail
- **duo-chat-opus-4-5** - Le plus capable pour les analyses complexes
:::note
Vous pouvez également spécifier la variable d'environnement 'GITLAB_TOKEN' si vous ne souhaitez pas
pour stocker le jeton dans le stockage d'authentification opencode.
:::
##### GitLab auto-hébergé
:::note[note de conformité]
OpenCode utilise un petit modèle pour certaines tâches d'IA telles que la génération du titre de la session.
Il est configuré pour utiliser gpt-5-nano par défaut, hébergé par Zen. Pour verrouiller OpenCode
pour utiliser uniquement votre propre instance hébergée par GitLab, ajoutez ce qui suit à votre
Fichier `opencode.json`. Il est également recommandé de désactiver le partage de session.
```json
{
"$schema": "https://opencode.ai/config.json",
"small_model": "gitlab/duo-chat-haiku-4-5",
"share": "disabled"
}
````
:::
Pour les instances GitLab auto-hébergées :
```bash
export GITLAB_INSTANCE_URL=https://gitlab.company.com
export GITLAB_TOKEN=glpat-...
```
Si votre instance exécute une AI Gateway personnalisée :
```bash
GITLAB_AI_GATEWAY_URL=https://ai-gateway.company.com
```
Ou ajoutez à votre profil bash :
```bash title="~/.bash_profile"
export GITLAB_INSTANCE_URL=https://gitlab.company.com
export GITLAB_AI_GATEWAY_URL=https://ai-gateway.company.com
export GITLAB_TOKEN=glpat-...
```
:::note
Votre administrateur GitLab doit activer les éléments suivants :
1. [Duo Agent Platform](https://docs.gitlab.com/user/duo_agent_platform/turn_on_off/) pour l'utilisateur, le groupe ou l'instance
2. Indicateurs de fonctionnalités (via la console Rails) :
   - `agent_platform_claude_code`
   - `third_party_agents_enabled`
     :::
##### OAuth pour les instances auto-hébergées
Afin que Oauth fonctionne pour votre instance auto-hébergée, vous devez créer
une nouvelle application (Paramètres → Applications) avec le
rappel URL `http://127.0.0.1:8080/callback` et étendues suivantes :
- api (Accédez au API en votre nom)
- read_user (Lire vos informations personnelles)
- read_repository (Autorise l'accès en lecture seule au référentiel)
Exposez ensuite l'ID de l'application en tant que variable d'environnement :
```bash
export GITLAB_OAUTH_CLIENT_ID=your_application_id_here
```
Plus de documentation sur la page d'accueil [opencode-gitlab-auth](https://www.npmjs.com/package/opencode-gitlab-auth).
##### Configuration
Personnalisez via `opencode.json` :
```json title="opencode.json"
{
  "$schema": "https://opencode.ai/config.json",
  "provider": {
    "gitlab": {
      "options": {
        "instanceUrl": "https://gitlab.com"
      }
    }
  }
}
```
##### Outils GitLab API (facultatif, mais fortement recommandé)
Pour accéder aux outils GitLab (demandes de fusion, tickets, pipelines, CI/CD, etc.) :
```json title="opencode.json"
{
  "$schema": "https://opencode.ai/config.json",
  "plugin": ["opencode-gitlab-plugin"]
}
```
Ce plugin fournit des fonctionnalités complètes de gestion du référentiel GitLab, notamment les examens MR, le suivi des problèmes, la surveillance du pipeline, etc.
---
### GitHub Copilot
Pour utiliser votre abonnement GitHub Copilot avec opencode :
:::note
Certains modèles peuvent avoir besoin d'un [Pro+
abonnement](https://github.com/features/copilot/plans) à utiliser.
Certains modèles doivent être activés manuellement dans vos [GitHub Copilot paramètres](https://docs.github.com/en/copilot/how-tos/use-ai-models/configure-access-to-ai-models#setup-for-individual-use).
:::
1. Exécutez la commande `/connect` et recherchez GitHub Copilot.
   ```txt
   /connect
   ```
2. Accédez à [github.com/login/device](https://github.com/login/device) et entrez le code.
   ```txt
   ┌ Login with GitHub Copilot
   │
   │ https://github.com/login/device
   │
   │ Enter code: 8F43-6FCF
   │
   └ Waiting for authorization...
   ```
3. Exécutez maintenant la commande `/models` pour sélectionner le modèle souhaité.
   ```txt
   /models
   ```
---
### Google Vertex AI
Pour utiliser Google Vertex AI avec OpenCode :
1. Rendez-vous sur **Model Garden** dans Google Cloud Console et vérifiez les
   modèles disponibles dans votre région.
   :::note
   Vous devez disposer d'un projet Google Cloud avec Vertex AI API activé.
   :::
2. Définissez les variables d'environnement requises :
   - `GOOGLE_CLOUD_PROJECT` : ID de votre projet Google Cloud
   - `VERTEX_LOCATION` (facultatif) : région pour Vertex AI (par défaut : `global`)
   - Authentification (au choix) :
     - `GOOGLE_APPLICATION_CREDENTIALS` : chemin d'accès au fichier de clé JSON de votre compte de service
     - Authentifiez-vous à l'aide de gcloud CLI : `gcloud auth application-default login`
Définissez-les lors de l'exécution de opencode.
```bash
   GOOGLE_APPLICATION_CREDENTIALS=/path/to/service-account.json GOOGLE_CLOUD_PROJECT=your-project-id opencode
```
Ou ajoutez-les à votre profil bash.
```bash title="~/.bash_profile"
   export GOOGLE_APPLICATION_CREDENTIALS=/path/to/service-account.json
   export GOOGLE_CLOUD_PROJECT=your-project-id
   export VERTEX_LOCATION=global
```
:::tip
