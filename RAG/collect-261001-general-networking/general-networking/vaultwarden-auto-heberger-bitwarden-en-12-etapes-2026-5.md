---
id: collect-261001-general-networking/general-networking/vaultwarden-auto-heberger-bitwarden-en-12-etapes-2026-5
title: "Mise à jour complète du système"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["mai"]
source: docs/RAG/collect-261001-general-networking/vaultwarden-auto-heberger-bitwarden-en-12-etapes-2026.md
source_anchor: ""
source_lines: [366, 439]
sha256: a92a464bf4dea72338749b4fb5df11b2804ad75457623f7f0db1759362d0177e
---

# Mise à jour complète du système

```
cd /opt/vaultwarden
# Sauvegarde AVANT toute mise à jour
./backup.sh
# Récupérer la nouvelle image et redéployer
docker compose pull
docker compose up -d
# Nettoyer les anciennes images
docker image prune -f
```
Surveillez les notes de version sur GitHub avant chaque montée : elles signalent les changements de comportement et les éventuels correctifs de sécurité – Release-monitoring.org retrace d’ailleurs cette cadence dès les correctifs mineurs 1.33.2 du 9 février 2025 et 1.34.1 du 27 mai 2025, avant que la 1.34.0 (web vault v2025.5.0, mai 2025) n’introduise l’inscription avec vérification e-mail obligatoire, tandis que la 1.35.0 (web vault v2025.12.0), sortie le 27 décembre 2025, a ajouté le support de l’authentification unique OpenID Connect (SSO), une attestation de release renforçant l’intégrité des binaires publiés, et étendu la compatibilité aux clients mobiles 2026.1.0 et suivants, selon WinterFlow.io, Tweakers et SourceForge – une cadence de publication toujours aussi soutenue, puisque le dépôt communautaire Debian livrait déjà le paquet 1.37.2-1 avec le web vault 2026.7.0-1 dès le 28 août 2026, quelques jours seulement après sa sortie officielle, selon le dépôt vaultwarden-deb. Évitez l’automatisation aveugle des mises à jour majeures sur un service aussi sensible – une sauvegarde fraîche suivie d’un `pull` manuel reste la voie la plus sûre.

## Choisir un hébergeur européen pour la souveraineté des données

L’argument de souveraineté ne tient que si le VPS est lui-même hébergé en Europe. Plusieurs fournisseurs proposent des machines adaptées à Vaultwarden, dont l’empreinte minimale autorise les offres les plus modestes. Les tarifs ci-dessous sont indicatifs au moment de la rédaction et doivent être vérifiés sur les sites des fournisseurs ; les spécifications, elles, sont stables.

| Hébergeur | Localisation | Config d’entrée typique | Tarif indicatif | Atout | 
|---|---|---|---|---|
| OVHcloud | France (Gravelines, Roubaix) | 1 vCPU / 2 Go | ≈ 5–7 €/mois | Acteur français, RGPD natif | 
| Scaleway | France (Paris) | 1–2 vCPU / 2 Go | ≈ 5–8 €/mois | Souveraineté, API moderne | 
| Infomaniak | Suisse | 2 vCPU / 2 Go | ≈ 6–9 €/mois | Neutralité, énergie renouvelable | 
| Hetzner | Allemagne / Finlande | 2 vCPU / 4 Go | ≈ 4–6 €/mois | Rapport perf/prix imbattable | 
| Auto-hébergement | Domicile / entreprise | Raspberry Pi 5 / mini-PC | Matériel ponctuel | Contrôle physique total | 

Pour un usage strictement personnel ou familial, un Raspberry Pi 5 à domicile suffit amplement, à condition d’avoir une IP publique stable ou un tunnel. Pour un usage professionnel ou multi-utilisateurs, un VPS européen géré offre une meilleure disponibilité. Vous pouvez vérifier l’empreinte réelle de votre instance en production avec une seule commande :

```
# Mesurer la consommation réelle du conteneur
docker stats --no-stream vaultwarden
# Sortie (exemple) :
# CONTAINER     CPU %   MEM USAGE / LIMIT   MEM %
# vaultwarden   0.00%   18.4MiB / 1.94GiB   0.93%
```
Avec moins de 20 Mo consommés au repos, l’empreinte mémoire dérisoire de Vaultwarden signifie que le facteur limitant sera presque toujours le réseau et la stratégie de sauvegarde, jamais le CPU ou la RAM. C’est précisément ce qui rend l’auto-hébergement accessible à quasiment n’importe quel matériel.

## 5 pièges courants à éviter

- **Exposer Vaultwarden en HTTP.** Les clients exigent un contexte sécurisé (Web Crypto API). Sans HTTPS valide, le déverrouillage échoue. Le reverse proxy avec certificat n’est pas optionnel.
- **Oublier de doubler les `$` du jeton Argon2.** Dans`.env` et`docker-compose.yml` , chaque`$` doit devenir`$$` , sinon Compose corrompt le hachage et le panneau admin devient inaccessible.
- **Laisser `SIGNUPS_ALLOWED=true` en production.** Après création de votre compte, verrouillez les inscriptions, sinon votre URL devient un point d’entrée ouvert.
- **Copier le fichier SQLite à chaud.** En mode WAL, un simple`cp` peut produire une sauvegarde incohérente. Utilisez toujours`sqlite3 .backup` .
- **Configurer un port WebSocket 3012 séparé.** Obsolète depuis la 1.29 : les WebSockets passent par le port HTTP principal. Cette config héritée des vieux tutoriels génère des erreurs.

## Dépannage : 8 erreurs fréquentes

| Symptôme | Cause probable | Solution | 
|---|---|---|
| Certificat invalide / non sécurisé | DNS non propagé ou port 80 bloqué | Vérifier `dig` , ouvrir le port 80 pour le défi ACME | 
| Le panneau `/admin` refuse le jeton | `$` non doublés dans le hachage Argon2 | Remplacer chaque `$` par`$$` , redéployer | 
| « Username or password is incorrect » côté client | URL serveur non configurée | `bw config server` / réglage « Auto-hébergé » | 
| Pas de synchronisation temps réel | WebSocket non relayé par le proxy | Proxy vers le port 80, supprimer toute config 3012 | 
| Erreur 502 Bad Gateway | Conteneur Vaultwarden non démarré | `docker compose logs vaultwarden` | 
| Boucle de renouvellement Caddy | Domaine ne pointe pas vers le VPS | Corriger l’enregistrement DNS `A` | 
| Inscription impossible (voulue plus tard) | `SIGNUPS_ALLOWED=false` | Inviter l’utilisateur depuis l’admin/organisation | 
| E-mails de vérification non envoyés | SMTP non ou mal configuré | Renseigner les variables `SMTP_*` dans`.env` | 

En cas de doute, le premier réflexe est `docker compose logs -f` : les journaux de Vaultwarden et de Caddy y sont explicites. Pour les erreurs de certificat, les journaux de Caddy précisent toujours l’étape exacte du défi ACME qui a échoué, ce qui pointe presque systématiquement vers le DNS ou le pare-feu.

## Astuces avancées pour aller plus loin

Une fois la base solide, plusieurs optimisations valent le détour. Côté **collaboration**, créez une *organisation* pour partager des collections de mots de passe entre membres d’une équipe ou d’une famille, avec des permissions granulaires. L’**accès d’urgence** permet de désigner un contact de confiance qui pourra récupérer votre coffre après un délai d’attente – une fonctionnalité premium chez Bitwarden, gratuite ici.

Pour une instance à forte charge ou multi-utilisateurs, migrez la base de SQLite vers **PostgreSQL ou MySQL** via la variable `DATABASE_URL` ; SQLite reste néanmoins parfait pour la plupart des usages. Si vous préférez Traefik à Caddy, notre tutoriel Traefik reverse proxy HTTPS montre comment le brancher via des labels Docker. Enfin, pour ne *jamais* exposer le coffre sur l’internet public, restreignez l’accès à un réseau privé : couplez Vaultwarden à un tunnel WireGuard ou à un réseau mesh Tailscale, de sorte que seuls vos appareils authentifiés atteignent le serveur. C’est l’approche la plus sûre pour un coffre familial.

Activez aussi les **notifications push mobiles** en enregistrant un identifiant d’installation gratuit auprès du relais Bitwarden, ce qui évite aux applications mobiles de scruter le serveur en continu. Et pensez à activer le *fail2ban* spécifique au panneau d’administration en plus de la prison principale : un attaquant qui découvre `/admin` sera ainsi banni après quelques tentatives.

## FAQ – Auto-héberger Vaultwarden

### Vaultwarden est-il sûr pour stocker des mots de passe sensibles ?

Oui. Vaultwarden réimplémente l’API Bitwarden mais conserve son modèle *zero-knowledge* : le chiffrement et le déchiffrement se font côté client, votre mot de passe maître ne quitte jamais l’appareil. Le serveur ne stocke que des données déjà chiffrées. La sécurité dépend surtout de votre durcissement : HTTPS, 2FA, jeton admin Argon2, sauvegardes et mises à jour régulières.

### Quelle différence avec le serveur Bitwarden officiel ?

