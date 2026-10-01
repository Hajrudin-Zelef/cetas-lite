---
id: collect-261001-general-networking/general-networking/fail2ban-bloquer-le-brute-force-ssh-en-12-etapes-2026-3
title: "Vérifier la version installée"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["distribution", "incident"]
source: docs/RAG/collect-261001-general-networking/fail2ban-bloquer-le-brute-force-ssh-en-12-etapes-2026.md
source_anchor: ""
source_lines: [212, 314]
sha256: 74d7503dc431a9f96f59ffb6f64e0d2e110f3324bddac2530ce6893f0e8c20ce
---

# Vérifier la version installée

Une fois vos jails actifs, vous aurez besoin d’un jeu de commandes courant pour l’exploitation quotidienne : consulter le statut, bannir ou débannir manuellement une IP, et diagnostiquer un filtre qui ne fonctionne pas comme prévu.

```
# Voir tous les jails actifs
sudo fail2ban-client status
# Voir le détail d'un jail précis (IP bannies, compteurs)
sudo fail2ban-client status sshd
# Bannir manuellement une IP suspecte
sudo fail2ban-client set sshd banip 203.0.113.42
# Débannir une IP (utile après un faux positif)
sudo fail2ban-client set sshd unbanip 203.0.113.42
# Recharger la configuration sans interrompre le service
sudo fail2ban-client reload
# Tester un filtre contre les journaux réels avant activation
sudo fail2ban-regex /var/log/auth.log /etc/fail2ban/filter.d/sshd.conf
# Tester un filtre contre le journal systemd
sudo fail2ban-regex systemd-journal /etc/fail2ban/filter.d/sshd.conf --journalmatch="_SYSTEMD_UNIT=ssh.service"
```
La commande `fail2ban-regex` est particulièrement utile avant de mettre en production un filtre personnalisé : elle vous montre exactement combien de lignes de log matchent votre expression régulière, sans bannir qui que ce soit, ce qui vous permet de valider le filtre en toute sécurité.

## Étape 9 : Se débannir soi-même en cas d’erreur

C’est l’incident le plus courant chez les débutants : après quelques essais de mot de passe SSH ratés, vous vous retrouvez bloqué hors de votre propre serveur. Deux méthodes de récupération existent, selon que vous avez encore un accès ou non.

**Si vous avez encore un accès** (une deuxième session SSH ouverte, ou un autre compte utilisateur) :

`sudo fail2ban-client set sshd unbanip VOTRE_IP_PUBLIQUE`
**Si vous êtes complètement bloqué**, passez par la console de secours de votre hébergeur (KVM chez OVHcloud, console Scaleway, ou VNC chez Hetzner), connectez-vous localement, puis exécutez la même commande de débannissement, ou ajoutez temporairement votre IP à `ignoreip` dans `jail.local` avant de recharger le service.

La meilleure prévention reste en amont : configurez toujours `ignoreip` avec votre IP fixe ou celle de votre VPN d’entreprise avant même d’activer le jail SSH, et testez la configuration sur une session secondaire avant de fermer votre session principale.

## Étape 10 : Surveiller les journaux et les alertes email

Fail2ban journalise chaque action dans `/var/log/fail2ban.log`. C’est votre première source pour comprendre ce qui se passe réellement sur le serveur et affiner vos seuils de `maxretry` et `findtime` au fil du temps.

```
# Suivre les bannissements en temps réel
sudo tail -f /var/log/fail2ban.log
# Compter les bannissements par jail sur les dernières 24h
sudo grep "Ban " /var/log/fail2ban.log | awk '{print $NF}' | sort | uniq -c | sort -rn | head -20
```
Si vous avez configuré `destemail` et `action = %(action_mwl)s` comme dans l’étape 3, vous recevrez un email à chaque bannissement, avec un extrait du journal correspondant. Sur un serveur très exposé, cela peut vite devenir bruyant : dans ce cas, limitez les notifications aux jails critiques (sshd, recidive) plutôt qu’à l’ensemble des jails web.

## Étape 11 : Durcir la configuration pour la production

Une fois la configuration de base validée, plusieurs réglages supplémentaires renforcent la protection sans complexifier excessivement l’administration.

- **Bannissement progressif** : augmentez la durée de bannissement à chaque récidive avec`bantime.increment = true` et`bantime.factor = 2` dans la section`[DEFAULT]` , ce qui double la peine à chaque nouvelle infraction de la même IP.
- **Changement du port SSH** : bien que ce ne soit pas une mesure de sécurité à part entière, déplacer SSH d’un port standard réduit fortement le volume de scans automatisés non ciblés, et donc la charge du jail.
- **Authentification par clé uniquement** : désactivez`PasswordAuthentication` dans`/etc/ssh/sshd_config` . Fail2ban reste utile même avec des clés, car il ralentit aussi les tentatives d’énumération d’utilisateurs.
- **Whitelist dynamique** : pour les équipes avec IP variable, envisagez un VPN type WireGuard et n’exposez SSH qu’à l’intérieur du tunnel.

```
[DEFAULT]
bantime.increment = true
bantime.factor = 2
bantime.maxtime = 604800
bantime = 3600
```
Avec ce réglage, une IP récidiviste verra sa durée de bannissement grimper de 1 heure à 2, puis 4, puis 8 heures, jusqu’à un plafond d’une semaine, sans intervention manuelle de votre part.

## Étape 12 : Vérifier les CVE et suivre les mises à jour de sécurité

Fail2ban n’échappe pas aux vulnérabilités logicielles. L’agrégateur OpenCVE référence **CVE-2025-45311**, publiée le 15 avril 2026 avec un score CVSS de **8,8 (élevé)** pour le vendor Fail2ban. La bonne pratique reste la même que pour tout composant de sécurité exposé : maintenez le paquet à jour via le gestionnaire de paquets système plutôt qu’une installation manuelle depuis les sources, afin de recevoir automatiquement les correctifs de sécurité distribués par Debian et Ubuntu.

```
# Vérifier si une mise à jour de sécurité est disponible
sudo apt update
apt list --upgradable | grep fail2ban
# Mettre à jour
sudo apt upgrade fail2ban
```
Consultez régulièrement la fiche CVE du projet sur OpenCVE ou abonnez-vous aux notifications de sécurité de votre distribution pour être alerté dès qu’un correctif est disponible.

## Fail2ban et les journaux de conteneurs Docker

De plus en plus de services tournent aujourd’hui dans des conteneurs Docker plutôt que directement sur l’hôte, ce qui complique la lecture des journaux par Fail2ban : par défaut, les logs applicatifs restent dans le conteneur et n’apparaissent pas dans `/var/log/` sur la machine hôte. Deux approches fonctionnent en 2026.

La première consiste à monter les journaux du conteneur vers un volume accessible depuis l’hôte, puis à pointer `logpath` vers ce fichier comme pour un service natif. La seconde, plus propre sur des hôtes qui utilisent déjà le pilote de journalisation journald pour Docker, consiste à filtrer directement le journal systemd sur le nom du conteneur.

```
# /etc/fail2ban/jail.d/docker-nginx.local.conf
[docker-nginx-auth]
enabled = true
filter = nginx-http-auth
backend = systemd
journalmatch = CONTAINER_NAME=nginx-proxy
maxretry = 5
bantime = 3600
```
Ce fonctionnement suppose que Docker est configuré avec le pilote de journalisation `journald` plutôt que le pilote `json-file` par défaut. Vérifiez ce réglage dans `/etc/docker/daemon.json` avant de vous fier à cette méthode, sans quoi `journalmatch` ne trouvera simplement aucune entrée correspondante et le jail restera silencieusement inactif.

## Empreinte système et impact sur les performances

Un argument fréquemment invoqué contre les outils de sécurité périphériques est leur coût en ressources. Dans la pratique, Fail2ban reste très léger sur un serveur correctement dimensionné : le démon Python tourne en tâche de fond et ne consomme des cycles CPU que lors de l’analyse des nouvelles lignes de journal, pas en continu.

| Indicateur | Valeur typique observée | Condition | 
|---|---|---|
| Mémoire résidente | 15 à 40 Mo | 2 à 5 jails actifs, trafic modéré | 
| Latence de bannissement | 500 ms à 5 s | Entre l’écriture du log et l’ajout de la règle nftables | 
| Charge CPU au repos | Quasi nulle | Aucune tentative en cours | 
| Charge CPU en pic | Notable mais transitoire | Backend systemd sans journalmatch, lecture du journal entier | 

