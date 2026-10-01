---
id: collect-261001-cisco/cisco/certbot-automatiser-le-https-en-13-etapes-2026-5
title: "Sur Ubuntu 24.04 / Debian 12 - installation via Snap (recommandée)"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["cyber", "incident", "mai"]
source: docs/RAG/collect-261001-cisco/certbot-automatiser-le-https-en-13-etapes-2026.md
source_anchor: ""
source_lines: [357, 399]
sha256: c937c3153f96796504fbbdd03e508c47cf9eb87c9298c2fd8d432307e2ec72df
---

# Sur Ubuntu 24.04 / Debian 12 - installation via Snap (recommandée)

Au-delà de la simple hygiène technique, l'automatisation du renouvellement TLS s'inscrit désormais dans un cadre réglementaire européen plus large. La directive NIS2, dont la transposition française reste suivie de près par l'ANSSI, impose aux entités essentielles et importantes de démontrer une gestion rigoureuse de leurs mesures de sécurité technique, ce qui inclut la disponibilité et l'intégrité des communications chiffrées. Un certificat expiré sur un service exposé n'est pas qu'un désagrément esthétique : c'est un indicateur direct de défaillance dans la gestion du cycle de vie de la sécurité, exactement le type de lacune que les audits de conformité recherchent. Nous détaillons les obligations concrètes de ce texte dans notre guide technique sur la directive NIS2.

Le Cyber Resilience Act, de son côté, pousse les fabricants de produits comportant des éléments numériques à documenter et sécuriser l'ensemble du cycle de vie de leurs logiciels, y compris les mécanismes de mise à jour et de gestion des secrets cryptographiques. Une automatisation robuste du renouvellement de certificats, avec surveillance et hooks de déploiement testés, correspond exactement à l'esprit de traçabilité et de résilience qu'exige ce règlement, sujet que nous approfondissons dans notre article sur le Cyber Resilience Act et la conformité SBOM. Dans les deux cas, la réduction imposée de la durée de vie des certificats agit comme un accélérateur : elle transforme une bonne pratique auparavant facultative en une nécessité opérationnelle documentée.

## Conseils avancés pour vos environnements de production

Une fois les bases posées, plusieurs raffinements permettent de fiabiliser durablement votre gestion des certificats. Pour les environnements avec plusieurs serveurs derrière un load balancer, centralisez l'obtention des certificats sur une seule machine, puis synchronisez les fichiers `fullchain.pem` et `privkey.pem` vers les autres nœuds via un outil comme rsync sécurisé par clé SSH, ou un gestionnaire de secrets centralisé, une approche que nous détaillons dans notre tutoriel sur HashiCorp Vault pour la gestion des secrets. Évitez à tout prix qu'un même domaine tente d'obtenir des certificats indépendamment depuis plusieurs serveurs sans coordination : cela multiplie inutilement la consommation de votre quota Let's Encrypt.

Adoptez également l'extension ACME Renewal Information (ARI) dès que votre client la supporte pleinement : elle laisse l'autorité de certification piloter le moment optimal de renouvellement plutôt que de vous fier uniquement à un seuil fixe de 30 jours, ce qui devient précieux à mesure que les cycles se raccourcissent. Documentez enfin une procédure de reprise manuelle pour le jour où l'automatisation échouera malgré tout : qui contacter, où se trouve le token API de secours, et comment obtenir un certificat en urgence via `certonly --manual` si toute votre chaîne d'automatisation tombe en panne simultanément. Cette procédure de secours n'est jamais du temps perdu : elle transforme un incident de plusieurs heures en correctif de quelques minutes.

## Questions fréquentes

**Certbot et Let's Encrypt sont-ils entièrement gratuits ?**

Oui. Let's Encrypt est une autorité de certification à but non lucratif financée par des dons et des sponsors, et Certbot est un logiciel libre développé par l'Electronic Frontier Foundation. Aucune limite de nombre de domaines gratuits n'existe, seules des limites de taux techniques s'appliquent pour éviter les abus.

**Combien de temps dure un certificat Let's Encrypt aujourd'hui ?**

Les certificats standards restent valables 90 jours en septembre 2026. Un profil optionnel à 45 jours (`tlsserver`) est disponible en test depuis le 13 mai 2026, avec une généralisation progressive prévue en 2027 et 2028.

**Certbot fonctionne-t-il avec Cloudflare, OVH ou Gandi ?**

Oui, via des plugins DNS-01 dédiés pour la plupart des grands fournisseurs, ce qui permet d'automatiser entièrement les certificats wildcard sans intervention manuelle à chaque renouvellement.

**Que se passe-t-il si mon certificat expire malgré tout ?**

Les navigateurs affichent un avertissement de sécurité bloquant pour les visiteurs, et la plupart des intégrations API tierces rejettent également la connexion. C'est pour cette raison que la surveillance active (étape 11) doit toujours accompagner l'automatisation, jamais la remplacer.

**Faut-il préférer un certificat wildcard ou plusieurs certificats individuels ?**

Un wildcard simplifie la gestion si vous ajoutez fréquemment de nouveaux sous-domaines, mais concentre le risque : la compromission de sa clé privée expose tous les sous-domaines simultanément. Des certificats individuels ou un certificat multi-domaines (SAN) limité offrent un meilleur cloisonnement.

**Certbot est-il adapté à une utilisation en entreprise, face à des solutions payantes ?**

Pour la très grande majorité des sites web et API publiques, Certbot avec Let's Encrypt couvre entièrement le besoin sans coût de licence. Les certificats commerciaux (EV, OV) restent pertinents dans des cas spécifiques où l'identité juridique de l'organisation doit apparaître directement dans le certificat, un besoin que Let's Encrypt ne couvre pas par conception.

**La directive NIS2 impose-t-elle explicitement Certbot ou Let's Encrypt ?**

Non, la directive ne cite aucun outil précis. Elle impose en revanche des obligations de sécurité technique et de gestion des risques dont une automatisation fiable du chiffrement TLS constitue une brique naturelle, sans imposer de fournisseur particulier.

**Comment tester le profil short-lived sans risquer une panne en production ?**

Créez d'abord un sous-domaine de test dédié, appliquez-y le profil `tlsserver` en environnement `--staging`, validez que votre hook de rechargement et votre surveillance fonctionnent correctement sur ce cycle court, puis étendez progressivement la configuration aux domaines de production.
