---
id: collect-261001-general-networking/general-networking/keycloak-sso-open-source-en-14-etapes-2026-1
title: "Base de données"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["advisory", "apache", "cyber"]
source: docs/RAG/collect-261001-general-networking/keycloak-sso-open-source-en-14-etapes-2026.md
source_anchor: ""
source_lines: [1, 55]
sha256: e3d58cad4c588ef7862d8490e72a1eabfdbd97d18eaf8467e185b3d836e2d139
---

# Base de données

Les identifiants volés restent la porte d’entrée numéro un des cyberattaques en 2026. Après la fuite chez l’Agence nationale des titres sécurisés, qui a exposé 11,7 millions de comptes, une question revient sans cesse dans les équipes techniques : qui contrôle vraiment les accès à vos applications, et sur quels serveurs vivent ces identifiants ? Keycloak, la plateforme de gestion des identités et des accès (IAM) portée par la Fondation Eclipse, répond directement à cette question. Elle est gratuite, publiée sous licence Apache 2.0, et elle tourne sur votre propre infrastructure plutôt que sur les serveurs d’un fournisseur tiers.

Keycloak centralise la connexion unique (SSO), l’authentification multifacteur, la fédération d’annuaires LDAP et Active Directory, ainsi que les protocoles OIDC, OAuth 2.0 et SAML 2.0. La version stable actuelle, la 26.7.3, est listée comme la plus récente par endoflife.date au 31 août 2026 ; elle succède à la 26.7.0 du 9 juillet 2026, qui avait introduit un aperçu de l’API SCIM, la haute disponibilité multi-cluster v2 et l’authentification SAML step-up d’après Skycloak, elle-même précédée par la 26.6.3 du 4 juin 2026 selon le service de suivi de versions VeRSSion, le tout tournant sur OpenJDK 21 en conteneur. Le jalon GitHub du projet visait initialement la version majeure 27.0.0 pour le 31 mars 2026, mais l’équipe a préféré continuer d’itérer sur la branche 26.x plutôt que de précipiter une rupture de compatibilité. Rester sur une branche ancienne coûte cher en sécurité : la série 26.0, marquée comme terminée par endoflife.date dès le 15 janvier 2025 (dernier correctif 26.0.8 daté du 13 janvier 2025), est concernée par la CVE-2025-11538 publiée en décembre 2025 par GitHub Security Advisory pour toute version antérieure à 26.4.4, et Red Hat a dû publier en 2025 deux avis distincts, RHSA-2025:19925 pour la RHBK 26.0.17 et RHSA-2025:12016 listant sept CVE. C’est le projet qui sert de socle au Red Hat Build of Keycloak (RHBK), la version commerciale supportée par Red Hat, ce qui donne une idée du sérieux de sa base de code en production — et de l’intérêt de suivre le rythme des mises à jour plutôt que de rester figé sur une version ancienne.

Ce tutoriel vous fait construire, de zéro, un déploiement Keycloak complet avec Docker Compose : base PostgreSQL, royaume dédié, client OIDC, utilisateurs et rôles, authentification multifacteur avec passkeys, fédération LDAP, reverse proxy TLS et bascule en mode production. Quatorze étapes, un peu moins d’une heure, et un projet fonctionnel que vous pourrez brancher sur une vraie application à la fin.

Ce guide s’adresse aux développeurs backend, aux administrateurs système et aux équipes DevOps qui gèrent déjà Docker au quotidien mais n’ont jamais mis les mains dans un serveur d’identité. Aucune connaissance préalable de Keycloak n’est nécessaire : chaque commande est expliquée, chaque piège est signalé avant qu’il ne vous morde, et chaque étape part du principe que vous partez d’une machine vierge. Si vous gérez déjà un cluster Kubernetes en production, les principes restent identiques, seul le mode de déploiement change.

## Pourquoi Keycloak s’impose en 2026

Trois forces poussent les équipes techniques vers l’auto-hébergement de leur IAM cette année. D’abord la directive européenne NIS2, dont l’application s’est durcie en 2025 et 2026 : elle pousse les opérateurs essentiels et importants à démontrer qu’ils maîtrisent leur chaîne d’authentification, pas seulement à la sous-traiter. Le Cyber Resilience Act, avec ses amendes pouvant atteindre 15 millions d’euros, ajoute une pression réglementaire similaire côté produits connectés. Ensuite, la généralisation des passkeys et de WebAuthn change la manière dont les utilisateurs s’attendent à se connecter, sans mot de passe, avec une clé biométrique ou matérielle. Keycloak a mûri son support de ces standards au fil des versions 22 à 26 : la 26.6.0, publiée le 8 avril 2026 d’après le dépôt GitHub du projet, a introduit à la fois le JWT Authorization Grant (RFC 7523) et un moteur de workflows, avant que la branche ne reçoive sa quatrième mise à jour mineure, la 26.6.4 du 26 juin 2026 rapportée par OpenStandia, preuve d’un rythme de maintenance soutenu sur cette seule lignée. Enfin, la souveraineté des données : héberger son propre serveur d’identité en Europe évite de faire transiter les données personnelles de connexion par des clouds soumis à des juridictions étrangères.

Sur le plan financier, l’argument est tout aussi direct. Keycloak ne facture rien par utilisateur : vous ne payez que l’infrastructure. Face à des solutions SaaS dont la facture grimpe avec chaque compte créé, ce modèle change complètement l’équation pour une startup qui passe de 500 à 50 000 utilisateurs.

Il y a aussi un changement de philosophie derrière cette bascule. Dans une architecture Zero Trust, l’identité remplace le pare-feu périmétrique comme véritable frontière de sécurité : chaque requête, qu’elle vienne du bureau ou d’un café à l’autre bout du pays, doit prouver qui elle est avant d’accéder à quoi que ce soit. Un serveur d’authentification que vous maîtrisez de bout en bout, avec ses journaux, ses règles de force brute et sa politique de sessions entièrement sous votre contrôle, s’intègre nettement mieux dans ce modèle qu’un service tiers où vous ne voyez qu’un tableau de bord.

## Prérequis : versions et matériel nécessaires

Avant de lancer la première commande, vérifiez que votre environnement correspond à ces exigences. Rien d’exotique : un serveur Linux standard ou une machine de développement suffisent pour suivre ce tutoriel de bout en bout.

| Composant | Version ou exigence minimale | 
|---|---|
| Docker Engine | 25.0 ou supérieur | 
| Docker Compose | v2.20 ou supérieur (plugin intégré à Docker) | 
| Image Keycloak | quay.io/keycloak/keycloak:26.7.0 | 
| Java embarqué dans l’image | OpenJDK 21 (conteneur), OpenJDK 25 en mode natif | 
| Base de données | PostgreSQL 15 ou supérieur (recommandée par le projet) | 
| RAM | 1 Go minimum, 2 Go ou plus en production | 
| CPU | 2 vCPU recommandés | 
| Espace disque | 5 Go minimum pour les images et les volumes | 
| Nom de domaine + DNS | requis pour le mode production et le TLS | 
| Certificat TLS | Let’s Encrypt via Traefik, ou certificat interne | 

Vérifiez vos versions installées avec deux commandes rapides avant de continuer :

```
docker --version
docker compose version
```
Si l’une des deux commandes échoue ou affiche une version trop ancienne, mettez à jour Docker avant d’aller plus loin. Les variables d’environnement Keycloak utilisées dans ce tutoriel (KC_BOOTSTRAP_ADMIN_USERNAME notamment) n’existent qu’à partir de la branche 26, une image plus ancienne ne fonctionnera pas de la même façon.

## Étape 1 : Préparer l’arborescence du projet

Créez un dossier de travail avec une structure claire. Elle sépare la configuration, les volumes de données et les certificats, ce qui facilite les sauvegardes et évite de mélanger les données persistantes avec le code de configuration versionné dans git.

```
mkdir -p keycloak-stack/{data/postgres,data/keycloak,certs}
cd keycloak-stack
touch docker-compose.yml .env
git init
```
Le dossier `data/postgres` contiendra les fichiers de la base, `data/keycloak` stockera les exports de royaume et les thèmes personnalisés, et `certs` accueillera vos certificats si vous ne passez pas par un reverse proxy pour le TLS. Ajoutez un fichier `.gitignore` qui exclut `data/` et `.env` avant de committer quoi que ce soit : ces dossiers contiennent des données sensibles et des identifiants.

## Étape 2 : Configurer les variables d’environnement

