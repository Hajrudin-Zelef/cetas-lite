---
id: collect-261001-cisco/cisco/certbot-automatiser-le-https-en-13-etapes-2026-1
title: "Sur Ubuntu 24.04 / Debian 12 - installation via Snap (recommandée)"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["apache", "mai"]
source: docs/RAG/collect-261001-cisco/certbot-automatiser-le-https-en-13-etapes-2026.md
source_anchor: ""
source_lines: [1, 40]
sha256: 7bd8251eb5a6a9d1bbfbb1f98a29ec073a64e1aa6c7199757faa28a3e35abcfc
---

# Sur Ubuntu 24.04 / Debian 12 - installation via Snap (recommandée)

Le 15 mars 2026, une règle discrète votée par le CA/Browser Forum est entrée en vigueur : la durée de vie maximale d’un certificat TLS public est passée à 200 jours. Ce n’est que la première marche d’un calendrier qui va bouleverser la gestion du HTTPS pour des millions de sites : 100 jours en mars 2027, puis 47 jours en mars 2029. Chez Let’s Encrypt, la trajectoire est encore plus serrée, avec un profil court à 45 jours déjà ouvert en test depuis le 13 mai 2026. Gérer manuellement un renouvellement de certificat tous les 45 jours au lieu de tous les 90 jours, ou pire tous les ans comme le faisaient certaines équipes avec des certificats commerciaux, devient tout simplement intenable sans automatisation. C’est exactement le problème que résout Certbot, le client ACME développé par l’Electronic Frontier Foundation, aujourd’hui en version 5.8.0. Ce tutoriel détaille, en 13 étapes concrètes, comment installer Certbot, obtenir vos premiers certificats Let’s Encrypt, automatiser leur renouvellement, et préparer votre infrastructure à l’ère des certificats à durée de vie courte.

## Pourquoi l’automatisation TLS devient incontournable en 2026

Pendant longtemps, un certificat TLS commercial durait un an, parfois deux ou trois ans avant l’interdiction des certificats multi-annuels par les navigateurs. Cette époque touche à sa fin. Le ballot SC-081v3, adopté par le CA/Browser Forum, impose une réduction progressive et irréversible de la durée de validité maximale des certificats publiquement approuvés, ainsi que de la durée de réutilisation des validations de contrôle de domaine (DCV). Le calendrier est désormais public et gravé dans les exigences de référence du secteur.

Cette réduction n’est pas une lubie administrative. Elle répond à un vrai problème de sécurité : plus un certificat vit longtemps, plus la fenêtre d’exploitation d’une clé privée compromise, d’une mauvaise configuration ou d’un domaine détourné reste ouverte. Un certificat de 47 jours limite mécaniquement les dégâts en cas de fuite. Le revers de la médaille, c’est la charge opérationnelle : une équipe qui gérait deux ou trois renouvellements par an devra bientôt en gérer huit, sur chaque domaine, sans interruption de service. Sans automatisation fiable, cela se traduit par des pages d’erreur “votre connexion n’est pas privée” chez les visiteurs et des alertes de sécurité chez les clients professionnels.

Let’s Encrypt, l’autorité de certification gratuite et automatisée lancée en 2015, a anticipé ce virage. Sa page de statistiques officielle, mise à jour le 18 août 2026, publie en continu le nombre de certificats actifs et le volume de délivrance quotidienne, confirmant qu’elle reste l’infrastructure de référence pour le chiffrement automatisé du Web. C’est dans ce contexte que Certbot, son client ACME historique, prend toute son importance : il ne s’agit plus d’un confort, mais d’une brique d’infrastructure critique.

## Certbot et Let’s Encrypt : comment fonctionne le protocole ACME

Certbot est un client du protocole ACME (Automatic Certificate Management Environment), une norme ouverte qui permet à un serveur de prouver, sans intervention humaine, qu’il contrôle bien un nom de domaine, puis de recevoir automatiquement un certificat signé. Concrètement, le dialogue se déroule en trois temps. D’abord, votre serveur contacte l’autorité de certification (Let’s Encrypt dans la majorité des cas) et demande un certificat pour un ou plusieurs domaines. Ensuite, l’autorité lance un défi de validation : elle vous demande de déposer un fichier accessible publiquement sur votre serveur web (défi HTTP-01) ou d’ajouter un enregistrement TXT temporaire dans votre zone DNS (défi DNS-01). Enfin, une fois le défi validé, le certificat est émis et téléchargé automatiquement par Certbot, qui peut aussi reconfigurer votre serveur web pour l’utiliser immédiatement.

Cette mécanique entièrement scriptable est ce qui rend possible le renouvellement toutes les 45 ou 47 jours sans intervention manuelle. Let’s Encrypt travaille également sur l’extension ACME Renewal Information (ARI), qui permet à l’autorité de certification d’indiquer proactivement au client ACME la fenêtre de temps recommandée pour renouveler un certificat donné, plutôt que de laisser chaque client deviner. L’objectif est d’éviter les pics de renouvellement simultanés et de mieux répartir la charge sur l’infrastructure de l’autorité. Cette évolution s’inscrit dans une feuille de route plus large, incluant des certificats à durée de vie courte accessibles via des profils ACME, et même des certificats couvrant des adresses IP directement dans leurs noms alternatifs (SAN), une nouveauté distincte des certificats classiques de 90 jours pour noms de domaine.

## Prérequis techniques et versions nécessaires

Avant de commencer, assurez-vous de disposer des éléments suivants. Le tutoriel a été vérifié avec Certbot 5.8.0 (publié le 1er septembre 2026), mais reste valable pour toute version de la branche 5.x.

- Un serveur Linux avec accès root ou sudo : Ubuntu 24.04 LTS, Debian 12/13, ou Rocky Linux 9 sont couverts nativement par les paquets officiels
- Certbot 5.8.0 ou une version récente de la branche 5 (support maintenu depuis le 19 septembre 2026 pour cette branche)
- Un nom de domaine dont vous contrôlez la zone DNS, avec les enregistrements A/AAAA déjà pointés vers votre serveur
- Un serveur web installé : Nginx 1.24+ ou Apache 2.4+ selon le plugin choisi
- Les ports 80 et 443 ouverts sur votre pare-feu pour la validation HTTP-01 (ou un accès API à votre fournisseur DNS pour la validation DNS-01)
- Python 3.9 ou supérieur, requis par Certbot en interne
- Docker et Docker Compose (version 2.x) si vous suivez la partie sur le projet conteneurisé

Comptez environ 60 minutes pour dérouler l’ensemble des 13 étapes sur un serveur déjà opérationnel, incluant les tests de renouvellement et la mise en place de la surveillance.

## Étape 1 : choisir entre validation HTTP-01 et DNS-01

Le choix de la méthode de validation conditionne tout le reste de votre configuration. La validation HTTP-01 est la plus simple : Certbot dépose un fichier temporaire dans le répertoire servi par votre serveur web, et Let’s Encrypt vérifie qu’il est bien accessible publiquement sur le port 80. Elle fonctionne pour un domaine ou un sous-domaine précis, mais ne permet jamais d’obtenir un certificat wildcard (type `*.exemple.fr`), et elle échoue si le port 80 est bloqué ou si le serveur est derrière un load balancer qui ne route pas correctement la requête de validation.

La validation DNS-01, elle, consiste à ajouter un enregistrement TXT temporaire (`_acme-challenge.exemple.fr`) dans votre zone DNS. Elle est indispensable pour les certificats wildcard, fonctionne même si le port 80 est fermé, mais nécessite soit un accès API à votre fournisseur DNS (OVH, Cloudflare, Gandi, Route 53 sont couverts par des plugins officiels ou communautaires), soit une intervention manuelle à chaque renouvellement, ce qui casse l’automatisation si vous n’avez pas de plugin dédié. Pour un site vitrine classique avec un seul sous-domaine, HTTP-01 suffit largement. Pour une infrastructure avec plusieurs sous-domaines dynamiques ou un besoin de wildcard, DNS-01 avec un plugin API est la seule option viable à long terme.

## Étape 2 : installer Certbot 5.8.0 sur votre serveur

La méthode d’installation recommandée par l’EFF elle-même est le paquet Snap, qui garantit une mise à jour automatique vers la dernière version stable, contrairement aux dépôts de certaines distributions qui peuvent livrer une version plus ancienne (Fedora propose par exemple encore la 5.7.0 dans ses dépôts standards au moment de la rédaction).

