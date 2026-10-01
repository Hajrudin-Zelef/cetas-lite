---
id: collect-261001-ia-llm/ia-llm/tutoriel-n8n-2-20-workflows-ia-en-13-etapes-2026-4
title: "Nœud Code – Langage Python"
domain: ia-llm
role: reference
task: reference
actors: ["OpenAI"]
dates: []
keywords: ["arr", "datacenter", "sandbox"]
source: docs/RAG/collect-261001-ia-llm/tutoriel-n8n-2-20-workflows-ia-en-13-etapes-2026.md
source_anchor: ""
source_lines: [300, 404]
sha256: dcedaba843cc62a8dce9f1c00adc43d35df62e63683af5dfc072b78d548ce402
---

# Nœud Code – Langage Python

```
Classifie cet e-mail entrant.
Retourne UNIQUEMENT un JSON conforme au schéma :
{
  "category": "urgent" | "commercial" | "spam",
  "priority": 1 | 2 | 3 | 4 | 5,
  "summary": "string FR max 200 caractères",
  "action": "string FR description de l'action recommandée"
}
E-mail :
Sujet : {{$json.subject}}
De : {{$json.from}}
Corps : {{$json.text}}
```
Le nœud **Switch** évalue ensuite `{{$json.category}}` et oriente le flux. Pour la branche *urgent*, ajoutez un Slack avec mention `@channel` et un Notion avec statut *P1*. Pour la branche *commercial*, créez un lead dans HubSpot via le nœud dédié. Pour *spam*, le workflow s’arrête après l’écriture dans la table de log.

Ce workflow remplace en moyenne 30 minutes de tri manuel par jour pour une équipe support de 5 personnes, soit l’équivalent d’un mois homme par an selon les estimations partagées dans la communauté n8n. Le coût marginal en tokens OpenAI s’élève à environ 0,30 € par tranche de 1 000 e-mails sur gpt-4.1-mini.

## Étape 11 : Tester, débugger et observer vos workflows

n8n dispose d’un onglet **Executions** qui conserve l’historique de chaque exécution avec timestamps, durée, statut et entrée/sortie de chaque nœud. Vous pouvez rejouer une exécution en cliquant sur *Retry > From this node*, ce qui économise un temps précieux pour itérer sur la dernière étape sans relancer l’intégralité.

Pour les workflows critiques, activez la rétention prolongée avec ces variables d’environnement :

```
EXECUTIONS_DATA_PRUNE: "true"
EXECUTIONS_DATA_MAX_AGE: 336        # heures, soit 14 jours
EXECUTIONS_DATA_PRUNE_MAX_COUNT: 50000
N8N_LOG_LEVEL: info
N8N_METRICS: "true"
N8N_METRICS_INCLUDE_DEFAULT_METRICS: "true"
```
La variable `N8N_METRICS` active un endpoint `/metrics` compatible Prometheus, indispensable pour brancher Grafana et alerter sur les workflows en échec. C’est l’approche standard dans tout cluster Kubernetes mature, et elle s’intègre parfaitement à une stack DevOps. Les équipes opérations disposent ainsi d’une vue temps réel sur les workflows.

Pour les erreurs récurrentes, configurez un workflow **Error Trigger** dédié. Ce trigger spécial capte toutes les exécutions échouées de l’instance et permet d’envoyer une alerte Slack, créer un ticket Jira ou simplement écrire dans un fichier de log centralisé. C’est l’équivalent du *dead-letter queue* en architecture événementielle.

## Étape 12 : Sécuriser n8n pour la production en Europe

La **CVE-2025-65964**, publiée le 8 décembre 2025 selon le NIST, a affecté toutes les versions de n8n entre 0.123.1 et 1.119.1. La vulnérabilité permettait une exécution de code à distance via le nœud Git en clonant un dépôt malveillant. Le correctif est inclus dans la version 1.119.2 et, a fortiori, dans la branche 2.x actuelle ; l’équipe n8n-io a d’ailleurs continué à maintenir la ligne 1.123.x en parallèle, avec un tag GitHub 1.123.20 publié le 6 février 2026, pour les instances qui n’ont pas encore basculé vers la 2.x.

Cette CVE rappelle quelques règles d’or pour exposer n8n en production :

- Mettre à jour **chaque semaine** au minimum, suivant la cadence des releases
- Activer **HTTPS uniquement** via un reverse proxy Caddy ou Traefik avec Let’s Encrypt
- Configurer `N8N_ENCRYPTION_KEY` avec une chaîne aléatoire de 32 caractères, jamais commitée en clair
- Restreindre l’éditeur aux IP de bureau via `N8N_BLOCK_ENV_ACCESS_IN_NODE=true`
- Désactiver les nœuds inutiles via `NODES_EXCLUDE=["n8n-nodes-base.executeCommand","n8n-nodes-base.git"]`
- Activer l’**authentification multi-facteurs** dans Settings > Personal > MFA
- Stocker les credentials externes dans un coffre comme **HashiCorp Vault** ou Doppler

Pour le RGPD, configurez les logs pour anonymiser les données personnelles avant rotation. n8n permet de définir des *data filters* par credential, ce qui masque automatiquement les valeurs sensibles dans l’onglet Executions. Si vous opérez en Europe, hébergez votre instance dans un datacenter Tier III certifié – OVHcloud Strasbourg, Scaleway DC5 Paris ou Hetzner Falkenstein restent les options privilégiées pour la souveraineté.

## Étape 13 : Déployer en production avec haute disponibilité

Pour passer à l’échelle au-delà d’une instance unique, n8n propose un mode **queue** qui sépare l’éditeur (main) des workers d’exécution. Les workflows passent par une file Redis, ce qui permet d’ajouter des workers à la demande sans interrompre le service.

```
services:
  redis:
    image: redis:7-alpine
    restart: unless-stopped
  n8n-main:
    image: n8nio/n8n:2.20.9
    environment:
      EXECUTIONS_MODE: queue
      QUEUE_BULL_REDIS_HOST: redis
      QUEUE_BULL_REDIS_PORT: 6379
  n8n-worker:
    image: n8nio/n8n:2.20.9
    command: worker
    deploy:
      replicas: 4
    environment:
      EXECUTIONS_MODE: queue
      QUEUE_BULL_REDIS_HOST: redis
```
Cette configuration permet d’exécuter quatre workers en parallèle, soit potentiellement quatre fois plus de tâches simultanées. Le main reste léger, dédié à l’éditeur et à l’API REST. Pour les pics de charge, ajoutez un auto-scaler Kubernetes ou un script Docker Swarm qui adapte le nombre de réplicas selon la profondeur de la file Redis.

En façade, déployez Caddy 2.8 ou Traefik 3.2 comme reverse proxy. Caddy s’avère particulièrement simple grâce à son auto-TLS intégré. Voici un Caddyfile minimal :

```
n8n.exemple.fr {
    reverse_proxy n8n-main:5678
    encode gzip zstd
    header {
        Strict-Transport-Security "max-age=31536000; includeSubDomains; preload"
        X-Frame-Options "SAMEORIGIN"
        X-Content-Type-Options "nosniff"
    }
}
```
## Cinq pièges courants à éviter avec n8n 2.x

Les utilisateurs rencontrent souvent les mêmes écueils en démarrant avec n8n. Voici les cinq erreurs les plus fréquentes observées dans le forum officiel, complétées par les retours du guide sur les workflows de réponse aux incidents assistée par IA, publié le 9 juillet 2026 sous le tag « tutorial » du blog officiel n8n.

- **Piège 1 : Conserver SQLite en production.** SQLite plafonne rapidement avec des verrouillages écriture-écriture. Basculez impérativement sur PostgreSQL dès la mise en production, comme montré dans notre docker-compose.
- **Piège 2 : Ignorer la rotation des données d’exécution.** Sans`EXECUTIONS_DATA_PRUNE` , votre base PostgreSQL grossit indéfiniment. Comptez 1-3 Mo par exécution riche en données, soit potentiellement plusieurs gigas par jour pour des workflows volumineux.
- **Piège 3 : Oublier de définir `N8N_ENCRYPTION_KEY`.** Sans clé fixe, n8n en génère une à chaque démarrage, rendant illisibles toutes vos credentials existantes.
- **Piège 4 : Mélanger éditions web et webhooks de production.** L’URL de test cesse de fonctionner dès la fermeture de l’éditeur. Pour le production, activez le workflow et utilisez l’URL*Production* .
- **Piège 5 : Surcharger un workflow avec trop de nœuds Code.** Chaque nœud Code instancie un sandbox isolé. Privilégiez un seul nœud Code consolidé plutôt que cinq petits nœuds en chaîne – gain typique : 40-60 % de temps d’exécution.

## Huit problèmes de dépannage les plus signalés

La communauté community.n8n.io regroupe des milliers de discussions sur le dépannage. Voici les huit symptômes les plus récurrents avec leur solution éprouvée.

