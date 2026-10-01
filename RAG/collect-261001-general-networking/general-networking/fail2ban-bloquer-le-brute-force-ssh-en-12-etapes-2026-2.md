---
id: collect-261001-general-networking/general-networking/fail2ban-bloquer-le-brute-force-ssh-en-12-etapes-2026-2
title: "Vérifier la version installée"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["apache"]
source: docs/RAG/collect-261001-general-networking/fail2ban-bloquer-le-brute-force-ssh-en-12-etapes-2026.md
source_anchor: ""
source_lines: [60, 211]
sha256: cc9c0eda35827a91f80b6d9e761d9ebbcc51dd5a1757405e003f66248f6e5219
---

# Vérifier la version installée

Le piège classique : un administrateur modifie `jail.conf` en pensant changer le comportement global, alors que le fichier `defaults-debian.conf` livré avec le paquet continue de s’appliquer par-dessus. Résultat, la modification semble ne jamais prendre effet. La bonne pratique, documentée dans les pages de manuel Debian de jail.conf(5), est de ne jamais toucher aux fichiers upstream et de tout centraliser dans `jail.local` et `jail.d/`.

Pour visualiser concrètement cet ordre de lecture, voici ce que renvoie une inspection typique d’un serveur Ubuntu 24.04 fraîchement installé, avant toute personnalisation :

```
$ ls -la /etc/fail2ban/jail.d/
total 16
drwxr-xr-x 2 root root 4096 fail2ban-jail.d
-rw-r--r-- 1 root root  198 defaults-debian.conf
$ cat /etc/fail2ban/jail.d/defaults-debian.conf
[DEFAULT]
banaction = nftables
banaction_allports = nftables[type=allports]
```
Tant qu’aucun fichier `jail.local` n’existe, Fail2ban tourne avec les valeurs de `jail.conf` surchargées uniquement par ce fichier de defaults Debian. C’est précisément ce que vous allez changer à l’étape suivante.

## Étape 3 : Créer votre fichier jail.local

Copiez le modèle de base puis créez votre propre fichier de surcharge. Ne copiez jamais l’intégralité de `jail.conf` dans `jail.local` : cela fige une configuration qui ne recevra plus les mises à jour de sécurité du paquet amont. Ne reprenez que les valeurs que vous souhaitez réellement modifier.

`sudo nano /etc/fail2ban/jail.local`
Insérez la configuration de base suivante, qui définit les paramètres globaux (temps de bannissement, fenêtre d’observation, nombre d’essais toléré) et une liste blanche pour vos propres IP de confiance :

```
[DEFAULT]
# Adresses jamais bannies : loopback + votre IP fixe/VPN
ignoreip = 127.0.0.1/8 ::1 203.0.113.10/32
# Durée de bannissement : 1 heure
bantime = 3600
# Fenêtre d'observation : 10 minutes
findtime = 600
# Nombre d'échecs tolérés avant bannissement
maxretry = 5
# Action de bannissement moderne (nftables déjà par défaut sur Debian/Ubuntu)
banaction = nftables
banaction_allports = nftables[type=allports]
# Notification email (optionnel, nécessite un MTA configuré)
destemail = [email protected]
action = %(action_mwl)s
```
Le paramètre `ignoreip` est celui que vous ne devez jamais oublier. Ajoutez-y systématiquement l’adresse IP publique depuis laquelle vous administrez le serveur, sans quoi une simple série de fautes de frappe sur votre mot de passe SSH peut vous bannir vous-même.

## Étape 4 : Protéger SSH avec le jail sshd

Le jail SSH est le point de départ de la quasi-totalité des déploiements Fail2ban, car SSH reste le service le plus scanné sur Internet. Créez un fichier dédié dans `jail.d/` plutôt que d’entasser toute votre configuration dans `jail.local`.

`sudo nano /etc/fail2ban/jail.d/sshd.local.conf`
Deux approches coexistent en 2026 : le backend fichier classique et le backend systemd/journald, aujourd’hui recommandé sur les distributions récentes puisque rsyslog n’est plus toujours installé par défaut.

```
# Option A — backend systemd (recommandé sur Ubuntu 24.04 / Debian 12+)
[sshd]
enabled = true
port = ssh
filter = sshd
backend = systemd
journalmatch = _SYSTEMD_UNIT=ssh.service
maxretry = 5
bantime = 3600
findtime = 600
# Option B — backend fichier classique (si rsyslog écrit /var/log/auth.log)
[sshd]
enabled = true
port = ssh
filter = sshd
logpath = %(sshd_log)s
maxretry = 5
```
Notez que la manpage `jail.conf(5)` de Debian est explicite sur ce point : avec `backend = systemd`, le paramètre `logpath` n’est pas valide, il faut utiliser exclusivement `journalmatch`. Mélanger les deux est une erreur fréquente qui empêche le jail de démarrer correctement.

Redémarrez le service pour appliquer la configuration :

```
sudo fail2ban-client reload
sudo systemctl restart fail2ban
sudo fail2ban-client status sshd
```
## Étape 5 : Vérifier l’action de bannissement (nftables vs iptables vs ufw)

En 2026, la tendance est nette : Debian et Ubuntu basculent l’action de bannissement par défaut vers **nftables**, le successeur moderne d’iptables intégré au noyau Linux depuis plusieurs années. Sur Ubuntu 24.04, le fichier `/etc/fail2ban/jail.d/defaults-debian.conf` définit déjà `banaction = nftables`, même si le modèle amont `jail.conf` affiche encore `iptables-multiport` par habitude historique.

| Backend | Statut en 2026 | Cas d’usage recommandé | 
|---|---|---|
| nftables | Par défaut sur Debian 12/13 et Ubuntu 24.04 | Nouveaux déploiements, meilleure gestion IPv6 | 
| iptables-multiport | Toujours supporté, présent dans jail.conf amont | Anciens systèmes, compatibilité avec scripts existants | 
| ufw | Surcouche à iptables/nftables | Serveurs déjà administrés via ufw pour la simplicité | 

Pour vérifier que les règles nftables sont bien injectées par Fail2ban, listez la table dédiée après un bannissement de test :

`sudo nft list ruleset | grep -A5 fail2ban`
Si vous préférez rester sur iptables pour des raisons de compatibilité avec d’autres scripts déjà en place, forcez explicitement l’action dans `jail.local` avec `banaction = iptables-multiport`. Les deux fonctionnent, mais nftables offre une gestion IPv6 plus cohérente nativement, sans avoir à dupliquer les règles comme c’était le cas avec ip6tables.

## Étape 6 : Configurer les jails web (Nginx, Apache, WordPress)

Si votre serveur héberge un site web, les attaques par force brute ne se limitent pas à SSH. Les tentatives de connexion sur `/wp-login.php`, `/xmlrpc.php` ou les zones d’authentification HTTP basique sont tout aussi fréquentes. Créez un fichier dédié pour ces jails.

`sudo nano /etc/fail2ban/jail.d/web.local.conf````
[nginx-http-auth]
enabled = true
filter = nginx-http-auth
logpath = /var/log/nginx/error.log
maxretry = 5
[nginx-botsearch]
enabled = true
filter = nginx-botsearch
logpath = /var/log/nginx/access.log
maxretry = 10
bantime = 86400
[apache-auth]
enabled = true
filter = apache-auth
logpath = /var/log/apache2/error.log
maxretry = 5
```
Pour WordPress spécifiquement, aucun filtre officiel n’est fourni de série avec le paquet Fail2ban. La pratique courante consiste à créer un filtre personnalisé qui recherche les échecs de connexion dans les journaux Nginx ou Apache, en ciblant les requêtes POST vers `wp-login.php` suivies d’une réponse HTTP 200 (échec silencieux typique de WordPress qui ne renvoie pas de 401/403).

`sudo nano /etc/fail2ban/filter.d/wordpress-login.conf````
[Definition]
failregex = ^<HOST> -.*"POST /wp-login\.php HTTP/.*" 200
ignoreregex =
```
Puis activez le jail correspondant dans votre fichier `web.local.conf` en pointant vers ce filtre et votre journal d’accès Nginx ou Apache.

## Étape 7 : Protéger le mail (Postfix et Dovecot)

Si votre serveur héberge également un service de messagerie, les serveurs SMTP et IMAP sont des cibles privilégiées pour le relais de spam et les attaques par dictionnaire. Fail2ban fournit des filtres prêts à l’emploi pour Postfix et Dovecot.

`sudo nano /etc/fail2ban/jail.d/mail.local.conf````
[postfix-sasl]
enabled = true
filter = postfix-sasl
logpath = /var/log/mail.log
maxretry = 4
bantime = 7200
[dovecot]
enabled = true
filter = dovecot
logpath = /var/log/mail.log
maxretry = 5
bantime = 7200
[recidive]
enabled = true
filter = recidive
logpath = /var/log/fail2ban.log
bantime = 604800
findtime = 86400
maxretry = 3
```
Le jail `recidive` mérite une mention à part : il surveille les propres journaux de Fail2ban et bannit pour une semaine entière (604 800 secondes) toute IP qui a déjà été bannie trois fois sur une période de 24 heures, toutes catégories confondues. C’est un filet de sécurité efficace contre les attaquants persistants qui changent de tactique.

## Étape 8 : Les commandes fail2ban-client indispensables

