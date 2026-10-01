---
id: collect-261001-cisco/cisco/tutoriel-n8n-2026-automatiser-vos-workflows-ia-5
title: "Vérifier le statut du conteneur"
domain: cisco
role: reference
task: reference
actors: ["OpenAI"]
dates: []
keywords: ["agent"]
source: docs/RAG/collect-261001-cisco/tutoriel-n8n-2026-automatiser-vos-workflows-ia.md
source_anchor: ""
source_lines: [489, 562]
sha256: 6ec1b02ba8caf4dbce6834bc5423f6257dc3ac06bd56b67e8e13ad30491f91e1
---

# Vérifier le statut du conteneur

```
#!/bin/bash
# Script de backup automatique des workflows n8n
# À exécuter via cron : 0 2 * * * /opt/n8n-backup/backup.sh
BACKUP_DIR="/opt/n8n-backup/workflows"
N8N_URL="http://localhost:5678"
API_KEY="votre-api-key-n8n"
DATE=$(date +%Y-%m-%d_%H%M)
mkdir -p "$BACKUP_DIR/$DATE"
# Exporter tous les workflows
curl -s -H "X-N8N-API-KEY: $API_KEY" \
  "$N8N_URL/api/v1/workflows?limit=100" | \
  python3 -c "
import sys, json, os
data = json.load(sys.stdin)
for wf in data.get('data', []):
    filename = f\"$BACKUP_DIR/$DATE/{wf['id']}_{wf['name'].replace(' ', '_')}.json\"
    with open(filename, 'w') as f:
        json.dump(wf, f, indent=2, ensure_ascii=False)
    print(f'Exported: {wf[\"name\"]} -> {filename}')
"
# Exporter les credentials (chiffrés)
curl -s -H "X-N8N-API-KEY: $API_KEY" \
  "$N8N_URL/api/v1/credentials" > "$BACKUP_DIR/$DATE/credentials_list.json"
# Commit et push vers Git
cd "$BACKUP_DIR"
git add -A
git commit -m "Backup n8n workflows - $DATE"
git push origin main
# Nettoyage des backups locaux de plus de 30 jours
find "$BACKUP_DIR" -maxdepth 1 -type d -mtime +30 -exec rm -rf {} +
echo "Backup terminé : $(ls $BACKUP_DIR/$DATE/*.json | wc -l) workflows exportés"
```
Configurez ce script en tâche cron quotidienne pour ne jamais perdre vos workflows. Pour restaurer un workflow, utilisez l'endpoint POST `/api/v1/workflows` avec le contenu JSON du fichier de backup. Cette approche Infrastructure-as-Code garantit la traçabilité complète des modifications et permet de revenir à n'importe quelle version antérieure.

Pour les équipes qui travaillent sur les mêmes workflows, n8n 2.0 propose un système de tags et de dossiers qui, combiné avec le versionnement Git, permet d'organiser les workflows par domaine métier (marketing, support, data, DevOps) et de suivre les modifications par équipe. Associé aux GitHub Actions, vous pouvez même automatiser le déploiement de workflows depuis un dépôt Git vers votre instance n8n de production.

## Les 7 Pièges les Plus Courants à Éviter avec n8n

Après avoir accompagné des dizaines d'équipes dans leur adoption de n8n, voici les erreurs les plus fréquentes et comment les éviter. Ces pièges sont particulièrement sournois car les workflows semblent fonctionner correctement en développement avant d'échouer en production.

**Piège 1 : Ignorer les rate limits des API tierces.** Chaque API impose des limites de requêtes (Slack : 1 requête/seconde par méthode, GitHub : 5 000 requêtes/heure, OpenAI : variable selon le plan). Sans le nœud « Split In Batches » avec un délai entre les lots, votre workflow sera bloqué par des erreurs 429. Solution : ajoutez systématiquement un « Wait » de 1-2 secondes entre les appels à des API tierces.

**Piège 2 : Stocker des secrets en clair dans les nœuds Code.** Il est tentant de coder en dur une clé API dans un nœud Code pour un test rapide. Mais ces valeurs sont visibles par tous les utilisateurs de l'instance et exportées en clair dans les backups JSON. Solution : utilisez toujours les credentials n8n et accédez-y via `$credentials` dans les expressions.

**Piège 3 : Ne pas configurer la purge des données d'exécution.** Par défaut, n8n conserve toutes les données d'exécution indéfiniment. Sur une instance active avec 100 workflows, la base de données peut atteindre plusieurs gigaoctets en quelques semaines, ralentissant l'interface et saturant le disque. Solution : activez `EXECUTIONS_DATA_PRUNE=true` et définissez une durée de rétention raisonnable (7 à 30 jours).

**Piège 4 : Utiliser SQLite en production.** SQLite est parfait pour le développement, mais il ne supporte pas les accès concurrents et ses performances se dégradent significativement au-delà de 10 000 exécutions stockées. Solution : migrez vers PostgreSQL dès que votre instance dépasse le stade du prototypage (voir étape 2).

**Piège 5 : Oublier la variable WEBHOOK_URL.** Si `WEBHOOK_URL` n'est pas définie ou pointe vers `localhost` en production, tous vos webhooks généreront des URLs inaccessibles depuis Internet. Les services tiers ne pourront pas atteindre votre instance. Solution : définissez toujours `WEBHOOK_URL` avec votre domaine complet incluant le protocole (`https://n8n.votredomaine.fr/`).

**Piège 6 : Ne pas définir de N8N_ENCRYPTION_KEY.** Sans clé de chiffrement personnalisée, n8n en génère une automatiquement. Si vous recréez le conteneur sans le même volume de données, la clé change et tous vos credentials deviennent illisibles. Solution : générez une clé avec `openssl rand -hex 32` et stockez-la dans votre fichier `.env`.

**Piège 7 : Créer des boucles infinies avec les webhooks.** Un workflow qui déclenche une action qui elle-même déclenche le même webhook crée une boucle infinie qui peut saturer votre instance en quelques secondes. Solution : ajoutez toujours une condition de sortie et activez la limite d'exécutions simultanées dans les paramètres du workflow.

## Dépannage Complet : 10 Problèmes et Solutions

Cette section couvre les problèmes les plus fréquemment rapportés par la communauté n8n en 2026 et leurs solutions testées. Conservez cette référence pour le diagnostic rapide en production.

| Problème | Cause probable | Solution | 
|---|---|---|
| Erreur « ECONNREFUSED » au démarrage | PostgreSQL pas encore prêt | Ajoutez un `healthcheck` PostgreSQL et`depends_on: condition: service_healthy` | 
| Webhooks non accessibles depuis Internet | Variable WEBHOOK_URL incorrecte ou pas de reverse proxy | Vérifiez que `WEBHOOK_URL` pointe vers votre domaine HTTPS public | 
| Credentials illisibles après mise à jour | Clé de chiffrement modifiée ou perdue | Restaurez le volume `n8n_data` ou redéfinissez`N8N_ENCRYPTION_KEY` | 
| Erreur 429 « Too Many Requests » | Rate limit API tierce atteint | Utilisez « Split In Batches » avec un délai de 1-2s entre les lots | 
| Workflow actif mais ne s'exécute pas | Timezone incorrect dans le Schedule Trigger | Vérifiez `GENERIC_TIMEZONE=Europe/Paris` dans les variables d'environnement | 
| Mémoire du conteneur saturée (OOM Kill) | Workflow traitant trop de données en une seule exécution | Limitez les données par batch, augmentez la mémoire Docker : `mem_limit: 2g` | 
| Interface web lente / timeout | Base SQLite trop volumineuse | Migrez vers PostgreSQL et activez la purge des exécutions | 
| Erreur SSL « certificate verify failed » | Certificat auto-signé ou expiré sur le service cible | Ajoutez `NODE_TLS_REJECT_UNAUTHORIZED=0` (dev uniquement) ou corrigez le certificat | 
| Nœud AI Agent timeout après 60s | Modèle IA trop lent ou prompt trop complexe | Augmentez le timeout du nœud HTTP et simplifiez le prompt système | 
| Données perdues après `docker compose down` | Volumes Docker non persistés | Vérifiez que les volumes nommés sont déclarés dans `docker-compose.yml` | 

**Problème récurrent : l'agent IA renvoie du JSON malformé.** C'est l'un des problèmes les plus fréquents avec les workflows IA dans n8n. Le modèle de langage peut occasionnellement renvoyer du texte supplémentaire avant ou après le JSON attendu. Solution robuste : ajoutez un nœud Code après l'agent IA qui extrait le JSON avec une regex :

