---
id: collect-261001-cisco/cisco/certbot-automatiser-le-https-en-13-etapes-2026-3
title: "Sur Ubuntu 24.04 / Debian 12 - installation via Snap (recommandée)"
domain: cisco
role: reference
task: reference
actors: []
dates: ["2026-09-23"]
keywords: ["incident", "mai"]
source: docs/RAG/collect-261001-cisco/certbot-automatiser-le-https-en-13-etapes-2026.md
source_anchor: ""
source_lines: [169, 258]
sha256: 1a7f526c0e4cd0599ffea87e9a221f4de5c1b7299120154221c26b9cf6fa3531
---

# Sur Ubuntu 24.04 / Debian 12 - installation via Snap (recommandée)

Par défaut, Certbot génère des clés RSA 2048 bits, un choix pensé avant tout pour la compatibilité maximale avec les anciens clients. En 2026, pour une nouvelle configuration sur des systèmes récents, ECDSA avec la courbe NIST P-256 est généralement recommandée : les clés et les signatures sont beaucoup plus compactes, les opérations cryptographiques sont plus rapides, ce qui réduit la latence de négociation TLS, sans compromis de sécurité pour l'immense majorité des usages web modernes. RSA 2048 bits reste pertinent si vous devez impérativement supporter des clients très anciens ou des systèmes embarqués atypiques qui ne gèrent pas encore ECDSA.

```
# Générer un certificat avec une clé ECDSA P-256
sudo certbot certonly --nginx \
  -d exemple.fr \
  --key-type ecdsa \
  --elliptic-curve secp256r1
```
Vous pouvez faire coexister les deux types de clés sur un même serveur (une chaîne ECDSA et une chaîne RSA de secours), une technique appelée agilité cryptographique, utile en période de transition ou si votre audience inclut des équipements hérités.

## Étape 10 : activer HSTS et l'OCSP Stapling

Obtenir un certificat n'est que la moitié du travail : encore faut-il configurer le serveur web pour exploiter pleinement les protections TLS disponibles. HTTP Strict Transport Security (HSTS) indique au navigateur de toujours utiliser HTTPS pour votre domaine, même si l'utilisateur tape volontairement une URL en HTTP, ce qui neutralise les attaques de rétrogradation de protocole. L'OCSP Stapling permet au serveur de fournir lui-même la preuve de non-révocation du certificat, plutôt que de laisser chaque navigateur interroger l'autorité de certification séparément, ce qui accélère la connexion et réduit la charge sur l'infrastructure de Let's Encrypt.

```
server {
    listen 443 ssl;
    server_name exemple.fr;
    ssl_certificate     /etc/letsencrypt/live/exemple.fr/fullchain.pem;
    ssl_certificate_key /etc/letsencrypt/live/exemple.fr/privkey.pem;
    ssl_protocols       TLSv1.2 TLSv1.3;
    # OCSP Stapling
    ssl_stapling on;
    ssl_stapling_verify on;
    ssl_trusted_certificate /etc/letsencrypt/live/exemple.fr/chain.pem;
    resolver 1.1.1.1 8.8.8.8 valid=300s;
    # HSTS - 1 an, inclut les sous-domaines
    add_header Strict-Transport-Security "max-age=31536000; includeSubDomains; preload" always;
}
```
Activez le préchargement HSTS (`preload`) uniquement une fois certain que tous vos sous-domaines fonctionnent correctement en HTTPS : la liste de préchargement des navigateurs est extrêmement difficile à quitter en cas d'erreur de configuration.

## Étape 11 : surveiller l'expiration et éviter la panne

Un timer systemd bien configuré échoue rarement, mais "rarement" ne veut pas dire "jamais" : changement de fournisseur DNS, token API expiré, quota de renouvellement dépassé après un test mal isolé, ou simplement une mise à jour système qui désactive le timer par erreur. La surveillance active de la date d'expiration reste indispensable, en complément de l'automatisation, pas à sa place.

```
# Vérifier rapidement la date d'expiration en ligne de commande
echo | openssl s_client -servername exemple.fr -connect exemple.fr:443 2>/dev/null \
  | openssl x509 -noout -enddate
# Sortie attendue :
# notAfter=Dec  1 08:14:22 2026 GMT
```
En production, préférez une surveillance externe automatisée plutôt qu'une vérification manuelle. L'exporter Prometheus Blackbox permet de sonder chaque domaine HTTPS et d'exposer le nombre de jours restants avant expiration sous forme de métrique, avec une alerte Grafana ou Alertmanager déclenchée par exemple 10 jours avant échéance. C'est une approche cohérente avec les principes de centralisation des journaux de sécurité déjà abordés dans notre tutoriel sur Graylog pour la centralisation des logs, où la même logique de détection proactive s'applique.

## Étape 12 : préparer la migration vers des sous-domaines et applications multiples

Dans un environnement réel, vous gérez rarement un seul domaine. Certbot permet d'étendre un certificat existant à de nouveaux noms sans repartir de zéro, ou de gérer plusieurs certificats indépendants selon vos besoins d'isolation.

```
# Ajouter un nouveau sous-domaine à un certificat existant
sudo certbot --nginx -d exemple.fr -d www.exemple.fr -d api.exemple.fr --expand
# Lister tous les certificats gérés sur le serveur
sudo certbot certificates
```
Pour des architectures avec de nombreux sous-domaines gérés dynamiquement (SaaS multi-tenant, environnements de préproduction générés à la volée), il est souvent plus robuste de combiner un certificat wildcard unique avec une architecture de reverse proxy centralisée, plutôt que de multiplier les certificats individuels et le risque d'oubli de renouvellement sur l'un d'entre eux.

## Étape 13 : passer aux certificats short-lived et aux profils ACME

Let's Encrypt a ouvert depuis le 13 mai 2026 un accès anticipé et volontaire à un profil de certificat à durée de vie réduite, baptisé `tlsserver`, valable 45 jours. La feuille de route officielle prévoit un abaissement à 64 jours pour le profil par défaut le 10 février 2027, puis 45 jours le 16 février 2028, en cohérence avec le calendrier plus large du CA/Browser Forum. Les clients ACME récents, dont Certbot, permettent de demander explicitement un profil via le paramètre dédié.

```
# Demander un certificat avec le profil short-lived (test recommandé en staging d'abord)
sudo certbot certonly --nginx \
  -d test.exemple.fr \
  --preferred-profile tlsserver \
  --staging
```
Testez systématiquement ce profil sur un sous-domaine non critique avant de l'appliquer en production. Un cycle de renouvellement à 45 jours laisse une marge de manœuvre beaucoup plus faible en cas d'incident : toute panne du timer, toute rotation de clé API DNS mal gérée, se traduit par une expiration réelle bien plus vite qu'avec un certificat classique de 90 jours. C'est précisément pour cette raison que l'étape de surveillance décrite plus haut devient non négociable dès lors que vous adoptez ces profils courts.

## Calendrier de réduction de la durée de vie des certificats TLS (2026-2029)

Le tableau ci-dessous résume les échéances officielles adoptées via le ballot SC-081v3 du CA/Browser Forum, qui s'appliquent à tous les certificats publiquement approuvés, au-delà du seul cas de Let's Encrypt.

| Date d'entrée en vigueur | Durée max. du certificat | Réutilisation validation domaine (DCV) | Statut au 23/09/2026 | 
|---|---|---|---|
| 15 mars 2026 | 200 jours | 200 jours | En vigueur | 
| 15 mars 2027 | 100 jours | 100 jours | À venir | 
| 15 mars 2029 | 47 jours | 10 jours | À venir | 
| 13 mai 2026 (Let's Encrypt) | 45 jours (profil tlsserver, opt-in) | Non concerné (profil test) | Disponible en accès anticipé | 
| 10 février 2027 (Let's Encrypt) | 64 jours (profil par défaut) | Non communiqué | Planifié | 
| 16 février 2028 (Let's Encrypt) | 45 jours (profil par défaut) | Non communiqué | Planifié | 

Notez que la première étape de mars 2026 touche surtout les certificats commerciaux classiques : les certificats Let's Encrypt étaient déjà limités à 90 jours depuis leur lancement, donc cette première marche ne réduit pas immédiatement leur cycle de vie. C'est la trajectoire propre à Let's Encrypt, avec ses profils courts optionnels puis généralisés, qui impactera le plus tôt les utilisateurs de Certbot.

## Certbot vs acme.sh vs win-acme vs lego vs Caddy : quel client choisir

Certbot n'est pas le seul client ACME disponible, et selon votre environnement, une alternative peut être plus adaptée. Voici les principales options du marché en 2026.

