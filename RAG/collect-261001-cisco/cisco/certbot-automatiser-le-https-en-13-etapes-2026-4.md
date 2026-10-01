---
id: collect-261001-cisco/cisco/certbot-automatiser-le-https-en-13-etapes-2026-4
title: "Sur Ubuntu 24.04 / Debian 12 - installation via Snap (recommandée)"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["apache", "cyber"]
source: docs/RAG/collect-261001-cisco/certbot-automatiser-le-https-en-13-etapes-2026.md
source_anchor: ""
source_lines: [259, 356]
sha256: 7a3f2acac717e8e9361491d449a30543e74964949294b370246fc037303c1be1
---

# Sur Ubuntu 24.04 / Debian 12 - installation via Snap (recommandée)

| Client | Langage / dépendances | Plateforme cible | Wildcard (DNS-01) | Cas d'usage idéal | 
|---|---|---|---|---|
| Certbot | Python 3 | Linux (Nginx, Apache) | Oui, via plugins DNS | Serveurs Linux classiques, plugins officiels nombreux | 
| acme.sh | Shell pur, zéro dépendance lourde | Linux, BSD, conteneurs légers | Oui, très large éventail de fournisseurs DNS | Environnements minimalistes, scripts d'automatisation custom | 
| win-acme | .NET | Windows Server, IIS | Oui, via plugins | Infrastructures Windows/IIS | 
| lego | Go, binaire unique | Multiplateforme (Linux, macOS, Windows) | Oui | Intégration dans des outils tiers (Traefik, Caddy s'en inspirent) | 
| Caddy (natif) | Go, serveur web intégré | Linux, macOS, Windows | Oui, HTTPS automatique par défaut | Nouveaux projets acceptant de changer de serveur web | 

Pour une infrastructure Linux déjà construite autour de Nginx ou Apache, Certbot reste le choix le plus documenté et le plus intégré, avec la communauté la plus large pour le support. acme.sh séduit les équipes qui veulent un script auto-contenu sans dépendance Python, notamment dans des images de conteneurs très légères. win-acme est incontournable dès qu'IIS entre en jeu. Caddy, de son côté, ne se contente pas d'automatiser Certbot : il intègre nativement la gestion ACME dans le serveur web lui-même, ce qui supprime une bonne partie de la configuration manuelle, au prix d'un changement de serveur web complet.

## Projet complet : stack Docker Compose Nginx + Certbot en production

Voici une stack fonctionnelle complète, pensée pour un déploiement Docker Compose, avec un conteneur Nginx et un conteneur Certbot dédié au renouvellement, sans dépendance à Snap sur l'hôte.

```
# docker-compose.yml
services:
  nginx:
    image: nginx:1.27-alpine
    ports:
      - "80:80"
      - "443:443"
    volumes:
      - ./nginx/conf.d:/etc/nginx/conf.d:ro
      - ./certbot/conf:/etc/letsencrypt:ro
      - ./certbot/www:/var/www/certbot:ro
    restart: unless-stopped
  certbot:
    image: certbot/certbot:latest
    volumes:
      - ./certbot/conf:/etc/letsencrypt
      - ./certbot/www:/var/www/certbot
    entrypoint: >
      /bin/sh -c "trap exit TERM;
      while :; do
        certbot renew --webroot -w /var/www/certbot --quiet;
        sleep 12h & wait $${!};
      done"
```
```
# nginx/conf.d/exemple.fr.conf
server {
    listen 80;
    server_name exemple.fr;
    location /.well-known/acme-challenge/ {
        root /var/www/certbot;
    }
    location / {
        return 301 https://$host$request_uri;
    }
}
server {
    listen 443 ssl;
    server_name exemple.fr;
    ssl_certificate     /etc/letsencrypt/live/exemple.fr/fullchain.pem;
    ssl_certificate_key /etc/letsencrypt/live/exemple.fr/privkey.pem;
    location / {
        proxy_pass http://app:8080;
    }
}
```
Pour l'obtention initiale du certificat (avant que le conteneur ne puisse tourner en boucle de renouvellement), lancez une commande ponctuelle avec le même conteneur, en mode webroot, avant de démarrer la boucle permanente.

```
docker compose run --rm certbot certonly \
  --webroot -w /var/www/certbot \
  -d exemple.fr --email [email protected] \
  --agree-tos --no-eff-email
docker compose up -d
```
Cette architecture sépare clairement le renouvellement (conteneur Certbot autonome) du service (conteneur Nginx), une approche cohérente avec les principes d'infrastructure comme code déjà décrits dans notre tutoriel Traefik Docker, une alternative à Nginx qui gère nativement le renouvellement ACME sans conteneur Certbot séparé.

## Erreurs courantes à éviter avec Certbot et Let's Encrypt

La majorité des incidents liés à Certbot en production ne viennent pas d'un bug de l'outil, mais d'erreurs de configuration récurrentes, que voici.

- **Tester en production plutôt qu'en staging :** Let's Encrypt applique des limites de taux strictes (nombre de certificats par domaine par semaine). Multiplier les tentatives ratées en environnement réel épuise ce quota et vous bloque parfois plusieurs jours. Utilisez systématiquement`--staging` pour tout test de configuration.
- **Oublier le hook de rechargement du service :** un certificat renouvelé mais jamais rechargé par Nginx ou Apache ne sert à rien. Le service continue de présenter l'ancien certificat jusqu'à son redémarrage manuel, ce qui annule tout le bénéfice de l'automatisation.
- **Confondre HTTP-01 et wildcard :** un certificat`*.exemple.fr` ne peut être obtenu qu'avec une validation DNS-01. Tenter d'utiliser`--nginx` ou`--apache` seul pour un wildcard échoue systématiquement.
- **Permissions trop larges sur les identifiants API DNS :** un jeton API donnant un accès en écriture à l'ensemble de votre compte DNS, au lieu d'être restreint à la seule zone concernée, est un risque de sécurité disproportionné par rapport au bénéfice.
- **Ignorer la propagation DNS avant validation :** un enregistrement TXT tout juste créé peut mettre plusieurs minutes à se propager selon le TTL configuré. Certbot propose un délai d'attente paramétrable (`--dns-cloudflare-propagation-seconds` par exemple) qu'il faut ajuster si les validations échouent de façon intermittente.
- **Ne pas surveiller les emails d'expiration :** Let's Encrypt envoie des alertes automatiques à l'adresse renseignée lors de la première demande. Si cette adresse n'est plus valide ou n'est jamais consultée, le dernier filet de sécurité avant une panne disparaît.

## Dépannage : 8 problèmes fréquents et leurs solutions

Voici les erreurs les plus fréquemment rencontrées lors du déploiement de Certbot, avec leur cause probable et la correction associée.

- **"Timeout during connect (likely firewall problem)" :** le port 80 est bloqué par votre pare-feu, votre groupe de sécurité cloud, ou un load balancer en amont. Vérifiez avec`curl -I http://votredomaine.fr` depuis une machine externe.
- **"Too many certificates already issued for exact set of domains" :** vous avez dépassé la limite de taux de production de Let's Encrypt. Attendez la fin de la fenêtre (généralement une semaine glissante) ou continuez vos tests en`--staging` .
- **"DNS problem: NXDOMAIN looking up TXT" :** l'enregistrement`_acme-challenge` n'existe pas encore ou n'a pas fini de se propager. Vérifiez avec`dig TXT _acme-challenge.exemple.fr` et augmentez le délai de propagation du plugin DNS.
- **Certificat renouvelé mais l'ancien reste actif dans le navigateur :** le service web n'a pas été rechargé. Ajoutez ou corrigez le hook dans`/etc/letsencrypt/renewal-hooks/deploy/` .
- **"Permission denied" sur /etc/letsencrypt :** les commandes Certbot doivent presque toujours s'exécuter avec`sudo` , car les clés privées sont protégées avec des permissions restrictives par conception. Ne modifiez jamais ces permissions pour "simplifier" un script.
- **"Failed authorization procedure" sur un wildcard :** une tentative de validation HTTP-01 a été utilisée par erreur au lieu de DNS-01. Relancez la commande avec le plugin DNS approprié.
- **Le timer systemd est actif mais rien ne se passe :** vérifiez les journaux avec`journalctl -u snap.certbot.renew.service` pour identifier une erreur silencieuse, souvent liée à un identifiant API expiré ou changé sans mise à jour du fichier de credentials.
- **OCSP Stapling signalé comme non fonctionnel par un test SSL externe :** vérifiez que la directive`resolver` pointe vers un DNS accessible depuis le serveur, et que le fichier`chain.pem` (et non`fullchain.pem` ) est bien référencé dans`ssl_trusted_certificate` .

## Conformité NIS2 et Cyber Resilience Act : le rôle du chiffrement automatisé

