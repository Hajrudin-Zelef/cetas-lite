---
id: collect-261001-general-networking/general-networking/crowdsec-tutoriel-en-12-etapes-30-min-2026-3
title: "Mettre à jour la liste des paquets et le système"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["acquisition"]
source: docs/RAG/collect-261001-general-networking/crowdsec-tutoriel-en-12-etapes-30-min-2026.md
source_anchor: ""
source_lines: [161, 271]
sha256: d01c0a8547739a1b1f7480e4712bbed615f5b93bb46423cb0e3c8b24806442be
---

# Mettre à jour la liste des paquets et le système

```
# Bannir manuellement une IP de test pendant 4 heures
sudo cscli decisions add --ip 203.0.113.42 --duration 4h --reason "test SSH"
# Vérifier qu'elle apparaît dans la liste des décisions
sudo cscli decisions list
# Lever le ban une fois le test terminé
sudo cscli decisions delete --ip 203.0.113.42
```
Sortie attendue : `cscli decisions list` affiche un tableau avec la source (Cscli), la valeur de l’IP, le motif, l’action (ban) et la durée restante. Si vous gérez l’accès SSH avec un VPN, pensez à compléter ce tutoriel par notre guide Tutoriel WireGuard : VPN Linux en 12 étapes pour réduire encore davantage la surface d’attaque.

## Étape 7 – Protéger un serveur web Nginx

Puisque 70 % des attaques détectées par la communauté sont de type HTTP, protéger votre couche web est prioritaire. La démarche comporte deux volets : faire *analyser* les logs Nginx par le moteur, puis installer un bouncer Nginx capable de *bloquer* les requêtes au niveau applicatif (renvoyer un 403 ou une page captcha).

```
# Installer la collection de détection Nginx
sudo cscli collections install crowdsecurity/nginx
# Déclarer les logs Nginx à analyser (acquisition)
sudo nano /etc/crowdsec/acquis.yaml
```
Ajoutez (ou vérifiez) un bloc d’acquisition pointant vers vos journaux Nginx. La clé `type: nginx` indique au moteur quel parseur appliquer.

```
filenames:
  - /var/log/nginx/access.log
  - /var/log/nginx/error.log
labels:
  type: nginx
```
Rechargez le moteur, puis installez le bouncer Nginx, qui s’intègre via un module Lua dans votre configuration Nginx :

```
# Recharger pour prendre en compte la nouvelle acquisition
sudo systemctl reload crowdsec
# Installer le bouncer Nginx (remédiation applicative)
sudo apt install crowdsec-nginx-bouncer -y
# Vérifier la lecture des logs Nginx
sudo cscli metrics | grep -i nginx
```
Le paquet du bouncer Nginx ajoute automatiquement sa configuration et génère sa propre clé API auprès de la LAPI. Après un `systemctl reload nginx`, toute IP frappée d’une décision se verra refuser l’accès directement par le serveur web, avant même d’atteindre votre application. C’est la double protection : réseau (firewall-bouncer) *et* applicatif (nginx-bouncer).

## Étape 8 – Brancher la Console et le CTI communautaire

La Console (app.crowdsec.net) est un tableau de bord SaaS **gratuit** qui transforme l’expérience CrowdSec. Vous y visualisez vos alertes en temps réel, la cartographie des attaquants, l’état de vos moteurs et la blocklist appliquée – sans plus dépendre uniquement de la ligne de commande. L’inscription est gratuite et n’impose aucun engagement.

Créez un compte sur la Console, copiez la clé d’enrôlement (*enrollment key*) affichée dans l’interface, puis liez votre moteur :

```
# Vérifier d'abord la connexion à la CAPI
sudo cscli capi status
# Enrôler ce moteur dans la Console (collez votre clé)
sudo cscli console enroll VOTRE_CLE_D_ENROLEMENT
# Recharger pour activer la liaison
sudo systemctl reload crowdsec
```
De retour dans la Console, validez la demande d’enrôlement : votre serveur apparaît en quelques secondes. Vous bénéficiez alors pleinement du renseignement communautaire. En partageant vos signaux d’attaque (anonymisés), vous « gagnez » l’accès à la blocklist mutualisée – un échange de bons procédés qui constitue le cœur du modèle CrowdSec. Cette logique de défense collective rejoint les ambitions de la stratégie nationale de cybersécurité de la France 2026-2030 en matière de mutualisation du renseignement.

## Étape 9 – Déployer le pare-feu applicatif (AppSec / WAF)

Introduit puis consolidé dans les versions récentes, le composant **AppSec** fait de CrowdSec un véritable pare-feu applicatif web (WAF), capable d’inspecter le contenu des requêtes HTTP en temps réel – pas seulement les logs *a posteriori*. Il s’appuie sur des règles de type OWASP et sur du *virtual patching* pour bloquer les exploits connus (injections, traversées de répertoire, CVE web). Depuis la version **1.8.0**, publiée en août 2026, ce WAF embarque également une détection de bots native qui vient compléter le virtual patching pour repérer directement les comportements automatisés (scrapers, credential stuffing) sans attendre la publication d’une CVE. La réactivité de ce mécanisme s’est illustrée avec **CVE-2025-14528** : CrowdSec a publié une règle de détection dès le 18 février 2026, soit deux jours seulement avant que l’exploitation de cette faille ne soit observée massivement sur Internet – un exemple concret de virtual patching qui protège avant même que l’attaque ne se généralise. C’est une fonctionnalité absente de fail2ban et l’un des grands arguments de la migration.

```
# Installer les collections AppSec (WAF)
sudo cscli collections install crowdsecurity/appsec-virtual-patching
sudo cscli collections install crowdsecurity/appsec-generic-rules
# Déclarer l'écouteur AppSec dans l'acquisition
sudo nano /etc/crowdsec/acquis.d/appsec.yaml
```
Créez le fichier d’acquisition AppSec qui fait écouter le moteur sur un port local dédié. Le bouncer Nginx transmettra les requêtes entrantes à ce point d’inspection.

```
listen_addr: 127.0.0.1:7422
appsec_config: crowdsecurity/appsec-default
name: appsec_serveur
source: appsec
labels:
  type: appsec
```
Configurez ensuite le bouncer Nginx pour activer le mode AppSec (paramètres `APPSEC_URL` et `ENABLE_INTERNAL_APPSEC` dans son fichier de configuration), rechargez le moteur puis Nginx. Désormais, une requête contenant une charge malveillante connue est rejetée *avant* d’atteindre votre application, même si l’IP émettrice n’a jamais été vue auparavant. Validez l’écoute avec `sudo cscli metrics`, section **Appsec Metrics**.

## Étape 10 – Gérer décisions, listes blanches et durées de ban

Une protection trop agressive est aussi dangereuse qu’une protection absente : se bannir soi-même ou bloquer un partenaire légitime est l’erreur la plus coûteuse. CrowdSec gère cela via les **listes blanches** (whitelists) et les **profils** de décision. Commençons par protéger vos adresses d’administration.

Créez un parseur de liste blanche pour vos IP de confiance (poste d’admin, monitoring, plages internes) :

`sudo nano /etc/crowdsec/parsers/s02-enrich/mes-whitelists.yaml````
name: crowdsecurity/whitelists
description: "Liste blanche des IP de confiance"
whitelist:
  reason: "IP d'administration et reseau interne"
  ip:
    - "198.51.100.7"
  cidr:
    - "10.0.0.0/8"
    - "192.168.0.0/16"
```
Rechargez avec `sudo systemctl reload crowdsec`. Pour ajuster la durée des bannissements, éditez `/etc/crowdsec/profiles.yaml` : la valeur `duration` par défaut est de 4 heures. Vous pouvez aussi remplacer un simple *ban* par un *captcha* (si un bouncer compatible est installé), une approche moins frontale pour les faux positifs potentiels :

```
# Allonger un ban a 24 h pour un recidiviste
sudo cscli decisions add --ip 203.0.113.42 --duration 24h --reason "recidive"
# Appliquer un captcha plutot qu'un blocage sec
sudo cscli decisions add --ip 203.0.113.99 --type captcha --duration 1h
# Purger toutes les decisions manuelles (prudence !)
sudo cscli decisions delete --all
```
## Étape 11 – Simuler une attaque pour valider la protection

Un dispositif de sécurité non testé est un dispositif théorique. Validons que la chaîne complète – détection, décision, remédiation – fonctionne réellement. La méthode la plus sûre consiste à utiliser `cscli explain`, qui rejoue un fichier de log à travers les parseurs et scénarios sans rien bloquer, idéal pour le débogage.

