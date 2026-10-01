---
id: collect-261001-cisco/cisco/tutoriel-n8n-2026-automatiser-vos-workflows-ia-2
title: "Vérifier le statut du conteneur"
domain: cisco
role: reference
task: reference
actors: ["AWS", "Anthropic", "Google", "OpenAI", "Stripe"]
dates: []
keywords: ["agents", "aws", "claude"]
source: docs/RAG/collect-261001-cisco/tutoriel-n8n-2026-automatiser-vos-workflows-ia.md
source_anchor: ""
source_lines: [154, 229]
sha256: 91cb40d3c221183ea78c7a3532e94efe8eca5f38c63d2041e48107d62e5d8a1a
---

# Vérifier le statut du conteneur

Maintenant que n8n est installé et configuré, créons un premier workflow fonctionnel. Ce workflow surveillera un flux RSS de nouvelles tech, filtrera les articles pertinents et enverra un résumé par email chaque matin — un cas d'usage classique mais immédiatement utile pour toute équipe technique.

Ouvrez l'interface n8n à `http://localhost:5678` et suivez ces étapes :

**1. Créer un nouveau workflow** : Cliquez sur « New Workflow » dans le menu principal. Nommez-le « Veille Tech Quotidienne ».

**2. Ajouter un déclencheur Schedule** : Cliquez sur le bouton « + » au centre du canvas. Recherchez « Schedule Trigger » et ajoutez-le. Configurez-le pour s'exécuter chaque jour à 8h00 (heure de Paris). Le cron expression sera `0 8 * * *`.

**3. Ajouter un nœud RSS Feed Read** : Connectez un nœud « RSS Feed Read » au déclencheur. Entrez l'URL du flux RSS souhaité, par exemple `https://hnrss.org/frontpage` pour Hacker News ou `https://www.lemondeinformatique.fr/flux-rss/thematique/toutes-les-actualites/rss.xml` pour Le Monde Informatique.

**4. Filtrer avec un nœud IF** : Ajoutez un nœud « IF » pour ne garder que les articles contenant des mots-clés spécifiques. Configurez la condition : `{{ $json.title.toLowerCase().includes('kubernetes') || $json.title.toLowerCase().includes('docker') || $json.title.toLowerCase().includes('ia') }}`.

**5. Formater avec un nœud Code** : Ajoutez un nœud « Code » connecté à la sortie « true » du IF pour formater les articles en HTML lisible. Ce nœud utilise JavaScript natif dans n8n 2.0 :

```
// Nœud Code : Formater les articles en HTML
const articles = $input.all();
const today = new Date().toLocaleDateString('fr-FR', {
  weekday: 'long',
  year: 'numeric',
  month: 'long',
  day: 'numeric'
});
let html = `<h2>Veille Tech du ${today}</h2>`;
html += `<p>${articles.length} articles pertinents trouvés :</p>`;
html += '<ul>';
for (const article of articles) {
  html += `<li>
    <a href="${article.json.link}">${article.json.title}</a>
    <br/><small>${article.json.pubDate || 'Date inconnue'}</small>
  </li>`;
}
html += '</ul>';
return [{ json: { htmlContent: html, articleCount: articles.length } }];
```
**6. Envoyer par email** : Connectez un nœud « Send Email » (SMTP ou Gmail) au nœud Code. Utilisez `{{ $json.htmlContent }}` comme corps HTML de l'email et `Veille Tech - {{ $json.articleCount }} articles` comme sujet. Pour Gmail, vous devrez configurer les credentials OAuth2 (voir l'étape 4).

Cliquez sur « Test Workflow » pour exécuter le workflow manuellement et vérifier que chaque nœud produit la sortie attendue. Le panneau de droite affiche les données de sortie de chaque nœud, ce qui facilite le débogage visuel. Une fois satisfait, activez le workflow avec le toggle en haut à droite.

## Étape 4 : Configurer les Credentials et Intégrations Tierces

La puissance de n8n réside dans ses plus de 200 intégrations natives. Chaque service externe nécessite la configuration de credentials sécurisés. En 2026, n8n 2.0 supporte le stockage chiffré des secrets via AWS Secrets Manager, Google Cloud Secret Manager, Azure Key Vault et HashiCorp Vault pour les environnements d'entreprise.

Pour configurer vos premiers credentials, rendez-vous dans **Settings → Credentials** dans l'interface n8n. Voici les intégrations les plus courantes et leur méthode d'authentification :

| Service | Type d'authentification | Configuration requise | Cas d'usage | 
|---|---|---|---|
| Gmail / Google Workspace | OAuth2 | Google Cloud Console → API credentials | Envoi/réception d'emails automatisés | 
| Slack | OAuth2 / Bot Token | Slack API → Create App → Bot Token Scopes | Notifications, alertes, chatbots | 
| GitHub | Personal Access Token | Settings → Developer Settings → PAT (fine-grained) | Issues, PRs, webhooks CI/CD | 
| Notion | API Key (Integration) | Notion Integrations → New Integration | Base de connaissances, gestion de projet | 
| OpenAI | API Key | platform.openai.com → API Keys | Workflows IA, génération de texte | 
| Anthropic (Claude) | API Key | console.anthropic.com → API Keys | Agents IA avancés, analyse de documents | 
| PostgreSQL | Connection string | Host, port, user, password, database | Lecture/écriture base de données | 
| Webhook | Aucune / Header Auth | URL générée automatiquement par n8n | Réception d'événements en temps réel | 

Pour configurer Gmail avec OAuth2, créez d'abord un projet dans la Google Cloud Console, activez l'API Gmail, puis créez des identifiants OAuth 2.0 de type « Application Web ». L'URL de redirection à configurer est `http://localhost:5678/rest/oauth2-credential/callback`. Dans n8n, créez un nouveau credential Gmail, collez le Client ID et le Client Secret, puis cliquez sur « Connect my account » pour autoriser l'accès.

Un point crucial pour la sécurité : n'utilisez jamais de credentials de production dans votre environnement de développement. Créez des comptes de service dédiés avec des permissions minimales. Pour les environnements d'entreprise, n8n 2.0 propose le RBAC (Role-Based Access Control) qui permet de limiter l'accès aux credentials par rôle utilisateur. Consultez notre guide sur l'architecture Zero Trust pour comprendre les meilleures pratiques de sécurité applicables à vos workflows automatisés.

## Étape 5 : Maîtriser les Webhooks et les Déclencheurs Avancés

Les webhooks sont le cœur battant de n8n pour les automatisations en temps réel. Contrairement aux déclencheurs planifiés (cron), les webhooks permettent à n8n de réagir instantanément à des événements externes : un push sur GitHub, un message Slack, un paiement Stripe, ou une soumission de formulaire. En 2026, n8n gère des millions de webhooks par jour sur ses instances cloud les plus actives.

Pour créer un webhook, ajoutez un nœud « Webhook » comme déclencheur de votre workflow. n8n génère automatiquement deux URLs : une URL de test (active uniquement pendant le test) et une URL de production (active quand le workflow est activé). L'URL suit le format `http://votre-domaine:5678/webhook/identifiant-unique`.

Voici un exemple pratique : créons un webhook qui reçoit des événements GitHub (push, PR, issue) et envoie une notification formatée sur Slack. Ce pattern est extrêmement courant dans les équipes DevOps françaises qui cherchent à centraliser leurs notifications.

Configurez le nœud Webhook avec la méthode POST et le chemin `/github-events`. Dans GitHub, allez dans Settings → Webhooks → Add webhook, et entrez l'URL de production n8n. Sélectionnez les événements « Push », « Pull requests » et « Issues ». N'oubliez pas de configurer un secret partagé pour valider l'authenticité des requêtes.

Après le webhook, ajoutez un nœud « Switch » pour router les événements selon leur type. Le header `X-GitHub-Event` contient le type d'événement (`push`, `pull_request`, `issues`). Chaque branche du Switch se connecte à un nœud Code qui formate le message Slack différemment selon l'événement. Ce pattern de routage événementiel est fondamental dans n8n et s'applique à tous les cas d'intégration webhook.

Pour les environnements de développement local, votre instance n8n n'est pas accessible depuis Internet. Utilisez un tunnel comme `ngrok` ou `cloudflared` pour exposer temporairement votre port 5678. En production, un reverse proxy Nginx ou Caddy avec HTTPS est indispensable. Nous détaillons cette configuration dans l'étape 9 sur le déploiement en production.

## Étape 6 : Créer des Workflows IA avec les Agents n8n

