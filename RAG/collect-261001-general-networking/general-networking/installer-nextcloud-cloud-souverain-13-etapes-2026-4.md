---
id: collect-261001-general-networking/general-networking/installer-nextcloud-cloud-souverain-13-etapes-2026-4
title: "Mise à jour complète du système"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["apache", "exploit", "open source"]
source: docs/RAG/collect-261001-general-networking/installer-nextcloud-cloud-souverain-13-etapes-2026.md
source_anchor: ""
source_lines: [306, 379]
sha256: 12786ba2c18b19432051981bb0d738f863d427aea7e27f24c0844e99a29af009
---

# Mise à jour complète du système

Un cloud souverain doit aussi être un cloud sûr. Commencez par activer l’authentification à deux facteurs : installez l’application « Two-Factor TOTP Provider » depuis l’App Store, puis imposez-la aux administrateurs et aux groupes sensibles via **Paramètres → Sécurité**. Cette mesure bloque l’écrasante majorité des compromissions de comptes par vol de mot de passe.

Pour une organisation de plus de 20 utilisateurs, envisagez un fournisseur d’identité (IdP) centralisé via SSO. **Authentik**, une solution open source européenne, s’intègre à Nextcloud en OpenID Connect et permet une gestion centralisée des comptes, des groupes et des politiques de mot de passe. C’est l’architecture modulaire – reverse proxy, IAM, base de données, stockage, bureautique – recommandée par les sources officielles 2026 sur la souveraineté numérique. Le principe du moindre privilège s’applique partout : créez des groupes par service, attribuez des quotas de stockage, et désactivez l’auto-inscription.

```
# Forcer le chiffrement HTTPS au niveau applicatif
docker compose exec -u www-data app php occ config:system:set \
  overwriteprotocol --value=https
# Activer l'app de sécurité TOTP en ligne de commande
docker compose exec -u www-data app php occ app:enable twofactor_totp
# Vérifier la note de sécurité depuis l'extérieur :
# rendez-vous sur https://scan.nextcloud.com et saisissez votre domaine
```
Utilisez le scanner officiel `scan.nextcloud.com` pour obtenir une note de A+ à F sur la configuration de votre serveur (en-têtes HTTP, version à jour, etc.). Caddy ajoute déjà l’en-tête HSTS ; visez le grade A. Enfin, pensez au RGPD : Nextcloud propose des applications de gestion du consentement et de journalisation des accès. La logique de conformité par conception rejoint celle que nous avons décrite pour l’hébergement local de modèles d’IA conformes au RGPD.

## Étape 12 – Sauvegardes et réversibilité des données

La réversibilité est un pilier de la souveraineté : vous devez pouvoir récupérer et migrer vos données à tout moment. Une sauvegarde Nextcloud complète comporte trois éléments : le **dump de la base PostgreSQL**, le **répertoire de données** (volume `nc_data`) et les fichiers de **configuration** (le `config.php` et le `docker-compose.yml`). Voici un script de sauvegarde robuste qui place d’abord Nextcloud en mode maintenance pour garantir la cohérence.

```
#!/bin/bash
# backup-nextcloud.sh
set -e
DEST=/srv/backups/nextcloud/$(date +%Y-%m-%d)
mkdir -p "$DEST"
cd ~/nextcloud-souverain
# 1. Mode maintenance pour figer les données
docker compose exec -T -u www-data app php occ maintenance:mode --on
# 2. Dump de la base PostgreSQL
docker compose exec -T db pg_dump -U nextcloud nextcloud > "$DEST/db.sql"
# 3. Copie du volume de données et de la config
docker run --rm -v nextcloud-souverain_nc_data:/data \
  -v "$DEST":/backup alpine tar czf /backup/nc_data.tar.gz -C /data .
# 4. Sortie du mode maintenance
docker compose exec -T -u www-data app php occ maintenance:mode --off
echo "Sauvegarde terminée : $DEST"
```
Rendez le script exécutable (`chmod +x backup-nextcloud.sh`) et planifiez-le chaque nuit via cron. Pour une vraie résilience, appliquez la règle **3-2-1** : 3 copies, sur 2 supports, dont 1 hors site – par exemple un stockage objet S3 chez un hébergeur français. Testez régulièrement la restauration : une sauvegarde jamais restaurée n’est qu’une hypothèse. Les plans d’action 2026 pour PME recommandent d’ailleurs explicitement de simuler un « scénario de rupture fournisseur » dans son PCA/PRA.

## Étape 13 – Mettre à jour et exploiter dans la durée

La maintenance dans le temps fait la différence entre un projet pérenne et une dette technique. Nextcloud suit un cycle de publication soutenu (plusieurs versions majeures par an dans la nouvelle ère Hub) et un support limité dans le temps : la version 31 (Hub 10), sortie le 25 février 2025, a reçu son correctif de maintenance 31.0.6 dès le 12 juin 2025 (checksum publié par Nextcloud GmbH) avant d’atteindre 31.0.8 le 14 août 2025, puis sa fin de vie en février 2026, soit tout juste 12 mois après sa sortie, d’après le wiki GitHub de Nextcloud. Entre ces deux jalons, la version 32.0.0 est sortie le 27 septembre 2025 chez Nextcloud GmbH, cette branche recevant encore un correctif de maintenance 32.0.15 le 13 août 2026 ; puis Hub 9 a introduit la version 33.0.0 le 18 février 2026, une refonte majeure de l’interface selon WinterFlow.io, avant que la branche stable n’atteigne la 34.x. Il est donc impératif de **ne jamais sauter de version majeure** : pour passer de la 33 à la 35, il faut transiter par la 34.0.4 avant d’atteindre la 35.0.0, diffusée le 16 septembre 2026 par Nextcloud GmbH et devenue depuis la version stable de référence. Avec Docker, la procédure est sûre et reproductible.

```
# Procédure de mise à jour (toujours sauvegarder avant !)
cd ~/nextcloud-souverain
# 1. Sauvegarde préalable
./backup-nextcloud.sh
# 2. Modifier le tag d'image dans docker-compose.yml
#    Exemple : nextcloud:33-apache -> nextcloud:34-apache
nano docker-compose.yml
# 3. Récupérer la nouvelle image et recréer le conteneur
docker compose pull app
docker compose up -d app
# 4. Lancer la mise à niveau de la base
docker compose exec -u www-data app php occ upgrade
docker compose exec -u www-data app php occ db:add-missing-indices
```
Surveillez l’exploitation au quotidien : la page **Vue d’ensemble** de l’administration, l’espace disque (`df -h`), et les journaux (`docker compose logs`). Activez les notifications par e-mail pour les alertes critiques. Pour la supervision avancée, Nextcloud expose un point de terminaison `serverinfo` compatible avec des outils comme Prometheus ou Zabbix. Une infrastructure souveraine bien exploitée tient des années sans surprise.

## 5 pièges courants à éviter lors de l’installation

Certaines erreurs reviennent systématiquement chez ceux qui débutent l’**auto-hébergement Nextcloud**. Les anticiper vous fera gagner des heures.

- **Oublier OVERWRITEPROTOCOL=https** : sans ce paramètre, Nextcloud derrière un reverse proxy génère des liens en HTTP, ce qui casse le chargement des ressources et déclenche des avertissements « contenu mixte ». C’est l’erreur n°1.
- **Exposer la base de données ou Redis** : ne mappez jamais les ports 5432 ou 6379 sur l’hôte. Ces services doivent rester sur le réseau Docker interne. Les exposer ouvre une faille de sécurité majeure.
- **Négliger la tâche cron** : laisser le mode AJAX par défaut entraîne des notifications manquantes, des fichiers non nettoyés et une indexation incomplète. Configurez toujours le cron système.
- **Sous-dimensionner la RAM avec Collabora** : Collabora Online est gourmand. Sur un serveur à 2 Go, il provoquera des plantages. Comptez au minimum 8 Go pour une équipe utilisant la bureautique.
- **Stocker les secrets en clair dans le Compose** : utilisez toujours un fichier`.env` avec`chmod 600` . Coder les mots de passe en dur dans le YAML, surtout s’il est versionné, est une fuite assurée.

## Dépannage : 8 erreurs fréquentes et leurs solutions

Voici les problèmes les plus rencontrés et la façon de les résoudre rapidement. Gardez cette table sous la main lors de votre premier déploiement.

