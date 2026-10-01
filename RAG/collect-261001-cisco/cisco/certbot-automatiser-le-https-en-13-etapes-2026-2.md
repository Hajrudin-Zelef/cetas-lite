---
id: collect-261001-cisco/cisco/certbot-automatiser-le-https-en-13-etapes-2026-2
title: "Sur Ubuntu 24.04 / Debian 12 - installation via Snap (recommandée)"
domain: cisco
role: reference
task: reference
actors: ["Google"]
dates: ["2026-09-23", "2026-12-01"]
keywords: ["apache"]
source: docs/RAG/collect-261001-cisco/certbot-automatiser-le-https-en-13-etapes-2026.md
source_anchor: ""
source_lines: [41, 168]
sha256: bdb3c9e14fcf66afd0efb0d49874a6c0a52d508a8b1e75bca8744f8674c4070f
---

# Sur Ubuntu 24.04 / Debian 12 - installation via Snap (recommandée)

```
# Sur Ubuntu 24.04 / Debian 12 - installation via Snap (recommandée)
sudo apt update
sudo apt install snapd -y
sudo snap install core
sudo snap refresh core
sudo snap install --classic certbot
sudo ln -s /snap/bin/certbot /usr/bin/certbot
# Vérifier la version installée
certbot --version
# Sortie attendue : certbot 5.8.0
```
Si vous préférez éviter Snap, le paquet `python3-certbot-nginx` ou `python3-certbot-apache` via `apt` fonctionne aussi, mais vérifiez toujours la version livrée avec `apt-cache policy certbot` avant de vous y fier pour de la production, car les dépôts Debian/Ubuntu accusent parfois plusieurs mois de retard sur la dernière release.

## Étape 3 : configurer le pare-feu et le DNS de votre domaine

Avant de lancer la moindre demande de certificat, vérifiez que votre domaine pointe correctement vers le serveur et que les ports nécessaires sont ouverts. C’est l’étape la plus souvent négligée, et la première source d’échec en production.

```
# Vérifier la résolution DNS depuis l'extérieur
dig +short A exemple.fr
dig +short AAAA exemple.fr
# Ouvrir les ports 80 et 443 avec ufw (Ubuntu)
sudo ufw allow 80/tcp
sudo ufw allow 443/tcp
sudo ufw reload
# Avec firewalld (Rocky Linux / RHEL)
sudo firewall-cmd --permanent --add-service=http
sudo firewall-cmd --permanent --add-service=https
sudo firewall-cmd --reload
```
Si votre serveur est hébergé derrière un reverse proxy comme Traefik, vérifiez que les règles de routage laissent bien passer les requêtes vers `/.well-known/acme-challenge/` sans redirection HTTPS forcée, sinon la validation HTTP-01 échouera systématiquement. Ce point est d’ailleurs directement lié aux bonnes pratiques détaillées dans notre tutoriel Traefik pour le reverse proxy HTTPS, une configuration très répandue en environnement Docker.

## Étape 4 : obtenir votre premier certificat avec le plugin Nginx

Le plugin Nginx de Certbot est le plus intégré : il détecte automatiquement vos blocs `server` existants, obtient le certificat, puis modifie lui-même la configuration Nginx pour activer le HTTPS et rediriger le trafic HTTP.

```
# Installer le plugin Nginx
sudo snap install certbot-plugin-nginx
# Obtenir et installer automatiquement le certificat
sudo certbot --nginx -d exemple.fr -d www.exemple.fr
# Sortie attendue (extrait) :
# Successfully received certificate.
# Certificate is saved at: /etc/letsencrypt/live/exemple.fr/fullchain.pem
# Key is saved at:         /etc/letsencrypt/live/exemple.fr/privkey.pem
# This certificate expires on 2026-12-01.
# Deploying certificate
# Successfully deployed certificate for exemple.fr to /etc/nginx/sites-enabled/exemple.fr
# Congratulations! You have successfully enabled HTTPS
```
Certbot vous proposera aussi de forcer la redirection HTTP vers HTTPS automatiquement (option recommandée en production). Si vous préférez garder le contrôle total sur votre fichier de configuration Nginx plutôt que de laisser Certbot le modifier, utilisez l’option `certonly` à la place de `--nginx`, puis référencez manuellement les chemins des certificats dans votre bloc `server`.

## Étape 5 : obtenir un certificat avec le plugin Apache

Le fonctionnement est identique côté Apache, avec le plugin dédié `certbot-plugin-apache`.

```
sudo snap install certbot-plugin-apache
sudo a2enmod ssl
sudo certbot --apache -d exemple.fr -d www.exemple.fr
# Pour ne récupérer que le certificat sans toucher à la config Apache
sudo certbot certonly --apache -d exemple.fr
```
Sur les serveurs mutualisés ou les environnements où vous ne gérez pas directement Apache ou Nginx (par exemple derrière un CDN ou un load balancer géré), la sous-commande `certonly --standalone` ou `certonly --webroot` permet d’obtenir le certificat sans plugin de serveur web, à charge pour vous de le déployer manuellement à l’endroit voulu.

## Étape 6 : générer un certificat wildcard via le défi DNS-01

Pour couvrir tous les sous-domaines d’un coup (`*.exemple.fr`), seule la validation DNS-01 fonctionne. L’exemple ci-dessous utilise le plugin DNS Cloudflare, mais des plugins équivalents existent pour OVH, Gandi, Route 53, Google Cloud DNS et la plupart des grands fournisseurs, en plus du mode manuel pour les autres.

```
# Installer le plugin DNS Cloudflare
sudo snap install certbot-dns-cloudflare
# Créer le fichier d'identifiants API (permissions restreintes obligatoires)
sudo mkdir -p /etc/letsencrypt/secrets
sudo nano /etc/letsencrypt/secrets/cloudflare.ini
# Contenu : dns_cloudflare_api_token = VOTRE_TOKEN_API
sudo chmod 600 /etc/letsencrypt/secrets/cloudflare.ini
# Obtenir le certificat wildcard
sudo certbot certonly \
  --dns-cloudflare \
  --dns-cloudflare-credentials /etc/letsencrypt/secrets/cloudflare.ini \
  -d "exemple.fr" -d "*.exemple.fr"
```
Si votre fournisseur DNS n’a pas de plugin dédié, une solution courante consiste à déléguer la sous-zone `_acme-challenge` vers un service tiers comme acme-dns, ce qui évite de donner à Certbot un accès API complet à l’ensemble de votre zone DNS de production, une bonne pratique de moindre privilège à ne pas négliger.

## Étape 7 : automatiser le renouvellement avec systemd timer

Depuis l’installation via Snap ou les paquets récents, Certbot installe automatiquement un timer systemd qui exécute `certbot renew` deux fois par jour. C’est cette mécanique, et non une intervention manuelle, qui doit gérer le renouvellement au quotidien, y compris quand la durée de vie des certificats descendra à 47 jours.

```
# Vérifier que le timer est actif
sudo systemctl status snap.certbot.renew.timer
# Sortie attendue (extrait) :
# ● snap.certbot.renew.timer - Timer for snap application certbot.renew
#      Loaded: loaded (/etc/systemd/system/snap.certbot.renew.timer; enabled)
#      Active: active (waiting)
#     Trigger: Tue 2026-09-23 18:12:44 UTC; 6h left
# Lister les prochaines exécutions programmées
systemctl list-timers | grep certbot
```
Certbot ne renouvelle un certificat que s’il lui reste 30 jours ou moins avant expiration, quelle que soit sa durée de vie initiale. Cette logique reste valable pour les futurs certificats de 45 ou 47 jours : le seuil de déclenchement est une fenêtre relative, pas une date fixe, ce qui simplifie grandement la transition vers des cycles plus courts.

## Étape 8 : tester avec –dry-run et ajouter des hooks de déploiement

Ne faites jamais confiance à un timer sans l’avoir testé. La commande de simulation permet de vérifier que tout le processus de renouvellement fonctionne, sans consommer votre quota réel auprès de Let’s Encrypt (les limites de taux de production sont volontairement strictes, l’environnement de simulation utilise une autorité de test séparée).

```
# Simuler un renouvellement complet
sudo certbot renew --dry-run
# Sortie attendue (extrait) :
# Simulating renewal of an existing certificate for exemple.fr
# Congratulations, all simulated renewals succeeded:
#   /etc/letsencrypt/live/exemple.fr/fullchain.pem (success)
```
Le renouvellement ne sert à rien si le service web n’est pas rechargé pour prendre en compte le nouveau certificat. Ajoutez un hook de déploiement qui s’exécute uniquement quand un certificat a réellement été renouvelé, plutôt qu’à chaque exécution du timer.

```
# Créer un hook exécuté après chaque renouvellement réussi
sudo mkdir -p /etc/letsencrypt/renewal-hooks/deploy
sudo tee /etc/letsencrypt/renewal-hooks/deploy/reload-nginx.sh > /dev/null <<'EOF'
#!/bin/bash
systemctl reload nginx
EOF
sudo chmod +x /etc/letsencrypt/renewal-hooks/deploy/reload-nginx.sh
```
## Étape 9 : choisir entre ECDSA P-256 et RSA pour vos clés

