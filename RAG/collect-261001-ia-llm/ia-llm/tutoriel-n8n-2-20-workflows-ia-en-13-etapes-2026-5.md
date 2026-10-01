---
id: collect-261001-ia-llm/ia-llm/tutoriel-n8n-2-20-workflows-ia-en-13-etapes-2026-5
title: "Nœud Code – Langage Python"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "Mistral", "United States"]
dates: []
keywords: ["agents", "attention", "claude", "llama", "mai", "memory", "mistral", "pruning"]
source: docs/RAG/collect-261001-ia-llm/tutoriel-n8n-2-20-workflows-ia-en-13-etapes-2026.md
source_anchor: ""
source_lines: [405, 495]
sha256: f23fe7adf7787f7188f35028921803a2fd7e9aa5cf5b4e7f096015981c1810bd
---

# Nœud Code – Langage Python

| Symptôme | Cause probable | Solution | 
|---|---|---|
| “Cannot read property of undefined” | Sortie nœud précédent vide | Ajouter un IF avec `{{$json}}` non vide | 
| Webhook 404 en production | Workflow inactif | Activer le toggle Active dans l’éditeur | 
| OAuth callback échoue | `WEBHOOK_URL` incorrect | Définir l’URL publique HTTPS exacte | 
| Lenteur après 1 000 executions | Pruning désactivé | Activer `EXECUTIONS_DATA_PRUNE` | 
| “Database is locked” | SQLite saturé | Migrer vers PostgreSQL 17 | 
| Memory leak du conteneur | Workflows infinis | Ajouter `NODE_OPTIONS=--max-old-space-size=4096` | 
| Variables d’env non prises en compte | Mauvaise indentation YAML | Valider avec `docker compose config` | 
| Credentials illisibles après upgrade | Clé chiffrement modifiée | Restaurer la clé d’origine ou recréer les credentials | 

Pour les problèmes plus exotiques, consultez les release notes officielles qui détaillent les breaking changes version par version. La migration de n8n 1.x vers 2.x demande notamment une attention particulière aux nœuds personnalisés tiers, dont le contrat d’API a évolué.

## Astuces avancées pour des workflows performants

Une fois les bases maîtrisées, plusieurs techniques permettent de tirer le maximum de n8n. Premièrement, exploitez les **sous-workflows** via le nœud *Execute Workflow*. Cette modularisation rend les workflows lisibles et réutilisables, à l’image des fonctions en programmation. Un sous-workflow *Send Slack Notification* peut être appelé depuis dix workflows parents avec des paramètres différents.

Deuxièmement, utilisez les **variables d’environnement personnalisées**. n8n expose toute variable préfixée `N8N_` aux expressions via `{{$env.N8N_MA_VARIABLE}}`. Idéal pour les clés API moins sensibles ou les flags de feature.

Troisièmement, tirez parti du nœud **Workflow Static Data**. Il permet de persister un état entre exécutions sans passer par une base externe. Pratique pour stocker un cursor de pagination, un timestamp de dernière exécution ou un compteur de retries.

```
// Stocker un timestamp
const staticData = $getWorkflowStaticData('global');
const last = staticData.lastRun || 0;
const now = Date.now();
staticData.lastRun = now;
return [{ json: { since: last, until: now } }];
```
Quatrièmement, employez **Split Out** et **Aggregate** pour les manipulations d’arrays. Ces deux nœuds, introduits dans n8n 2.x, remplacent avantageusement les anciens Split In Batches dans la majorité des cas, avec une syntaxe plus claire.

Cinquièmement, créez vos propres **nœuds communautaires** en TypeScript. Le SDK officiel `n8n-nodes-starter` sur GitHub fournit un modèle prêt à publier sur npm. Plusieurs entreprises françaises – Doctolib, Qonto, BlaBlaCar – utilisent cette approche pour exposer leurs API internes à leurs équipes ops sans écrire de code.

## Comparatif des coûts mensuels pour 100 000 exécutions

| Déploiement | Coût mensuel | Données restent en Europe | Effort opérationnel | 
|---|---|---|---|
| n8n self-hosted (Hetzner CX31) | ~8 € | Oui (Falkenstein) | Moyen | 
| n8n self-hosted (Scaleway DEV1-M) | ~13 € | Oui (Paris) | Moyen | 
| n8n Cloud Starter | 20 € | Oui (Francfort) | Faible | 
| n8n Cloud Pro | 50 € | Oui (Francfort) | Faible | 
| Zapier Professional | ~73 € | Non (US) | Faible | 
| Make Pro | ~16 € | Partiel | Faible | 

Pour les volumes intermédiaires (entre 5 000 et 50 000 exécutions mensuelles), le self-hosted Hetzner reste imbattable avec une économie de l’ordre de 60-80 % sur trois ans, à condition de maîtriser Docker et le reverse proxy. Au-delà de 200 000 exécutions, le modèle queue avec workers s’impose et l’on bascule vers un cluster Kubernetes managé type Scaleway Kapsule ou OVH Managed Kubernetes.

## Intégration avec l’écosystème français et européen

n8n dispose de plus de 400 intégrations officielles plus une bibliothèque communautaire conséquente. Plusieurs nœuds répondent spécifiquement aux besoins des entreprises françaises et européennes :

- **Mistral AI** : appel direct aux modèles Mistral Large, Codestral et Pixtral via leur API hébergée en Europe
- **OVH AI Endpoints** : intégration HTTP simple avec les modèles hébergés à Strasbourg ou Gravelines
- **Pennylane** : automatisation comptable pour les TPE/PME françaises
- **Qonto** : récupération des transactions bancaires pour les workflows finance
- **Doctolib** : gestion automatisée des rendez-vous médicaux (via webhook)
- **Sellsy, Axonaut, Furious Squad** : trois CRM français nativement supportés

Cette densité d’intégrations locales fait de n8n l’orchestrateur naturel pour les ETI françaises qui veulent éviter la dépendance aux outils américains. C’est également la raison pour laquelle plusieurs intégrateurs spécialisés – agences low-code, ESN – recommandent désormais n8n par défaut dans leurs missions d’automatisation.

## FAQ – questions fréquentes sur n8n en 2026

### n8n est-il vraiment gratuit ?

n8n est **fair-code**, ce qui autorise un usage interne illimité gratuit en auto-hébergement. La revente du produit ou l’usage en SaaS commercial sous votre marque demande un accord commercial avec n8n GmbH. Pour 99 % des cas d’usage internes en entreprise, le coût se limite à celui du serveur d’hébergement.

### Quelle est la dernière version stable en mai 2026 ?

Selon la documentation officielle n8n Docs, la version **2.35** est sortie en août 2026, tandis que le suivi GitHub communautaire de n8n-assistant place la **2.36.8** comme dernière stable du même mois (la bêta 2.37.4 étant la ligne expérimentale la plus récente) ; la branche 2.20, elle, avait été lancée en mai 2026 avec la vérification du Netlify Trigger parmi ses nouveautés, avant la 2.23 du 27 mai 2026. Une nouvelle release mineure sort en moyenne chaque semaine. Consultez toujours docs.n8n.io/release-notes avant de mettre à jour en production.

### Peut-on utiliser n8n sans connaître de code ?

Oui, environ 80 % des workflows peuvent être créés sans une ligne de code grâce à l’éditeur visuel et aux expressions Mustache. Cependant, le nœud Code reste utile pour les transformations complexes ou les calls API non couverts par un nœud natif.

### Comment migrer depuis Zapier ou Make ?

Il n’existe pas d’importeur automatique. La méthode recommandée consiste à recréer les Zaps ou Scenarios un par un dans n8n, en profitant de la migration pour les optimiser. Un Zap moyen se traduit par 5 à 10 nœuds n8n et demande 30 minutes de portage par un opérateur expérimenté.

### n8n est-il compatible avec le RGPD ?

En auto-hébergement chez un cloud européen (OVHcloud, Scaleway, Hetzner), n8n est totalement compatible RGPD puisque les données ne quittent jamais le territoire européen. Documentez les credentials utilisées, anonymisez les logs d’exécution et tenez à jour un registre des traitements.

### Quels modèles IA fonctionnent le mieux avec n8n ?

Pour le tri et la classification simple, **gpt-4.1-mini** ou **Claude Haiku 4.5** offrent le meilleur rapport qualité/prix. Pour les agents avec outils, **Claude Sonnet 4.6** et **gpt-4.1** dominent. Pour le 100 % local, **Mistral-Small 3.2** et **Llama 4 8B** via Ollama sont les options les plus stables.

### Combien de workflows peut-on lancer en parallèle ?

Le mode *main* standard limite à environ 100 exécutions simultanées sur un serveur 4 vCPU / 8 Go. Le mode *queue* avec workers Redis lève cette limite : comptez environ 50-100 exécutions parallèles par worker, scalable horizontalement à l’infini.

### Comment sauvegarder ses workflows ?

