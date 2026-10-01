---
id: collect-261001-general-networking/general-networking/keycloak-sso-open-source-en-14-etapes-2026-6
title: "Base de données"
domain: general-networking
role: reference
task: reference
actors: ["Microsoft", "Oracle"]
dates: []
keywords: []
source: docs/RAG/collect-261001-general-networking/keycloak-sso-open-source-en-14-etapes-2026.md
source_anchor: ""
source_lines: [410, 432]
sha256: 1f3d6570d2f4bdba87dae712e89db9bb51ecdb61f3c1bf2c854b7dbe6f45d6ea
---

# Base de données

### Combien d’utilisateurs Keycloak peut-il gérer ?

Le logiciel n’impose aucune limite fixe. Le nombre réel dépend surtout de votre base de données et de votre configuration de cache. Avec PostgreSQL correctement dimensionné et le cache distribué Infinispan activé, la communauté documente des déploiements de plusieurs centaines de milliers de comptes, une échelle cohérente avec les chiffres publiés par la CNCF en août 2026 : plus de 12 000 utilisateurs actifs du projet et plus de 2 millions d’utilisateurs à travers tout son écosystème. Le nombre de projets qui dépendent de Keycloak a lui aussi bondi de 68,6 %, passant de 4 947 à 8 342 projets recensés d’ici juin 2025 d’après les données Linux Foundation LFX Insights. L’enquête communautaire de mars 2026, qui a recueilli plus de 360 réponses, confirme cette dynamique d’adoption.

### Keycloak est-il conforme au RGPD et à la directive NIS2 ?

Keycloak est un outil, pas un service certifié : il ne délivre aucune conformité automatique. Mais l’auto-hébergement en Europe facilite le respect du RGPD, puisque les données d’identité restent sous votre contrôle direct, et répond à l’esprit de NIS2, qui pousse les opérateurs essentiels à maîtriser leur chaîne d’authentification plutôt qu’à la sous-traiter hors de l’UE.

### Faut-il Kubernetes pour faire tourner Keycloak en production ?

Non. Docker Compose, comme dans ce tutoriel, suffit largement pour une PME ou un projet interne. Kubernetes, via l’Operator Keycloak officiel, devient pertinent à partir du moment où vous avez besoin d’une haute disponibilité multi-nœuds avec bascule automatique.

### Comment migrer depuis Auth0 ou Okta vers Keycloak ?

La migration passe généralement par trois étapes : exporter les utilisateurs via l’API du fournisseur d’origine, les réimporter dans Keycloak via l’API Admin ou un script kcadm.sh, puis migrer progressivement les applications clientes vers les nouveaux endpoints OIDC. Le point le plus délicat concerne les mots de passe : la plupart des fournisseurs SaaS ne les exportent jamais en clair pour des raisons évidentes de sécurité, ce qui oblige généralement à forcer une réinitialisation groupée ou à s’appuyer sur un algorithme de hachage compatible si le fournisseur le documente. Prévoyez une période de coexistence des deux systèmes, avec les deux fournisseurs d’identité actifs en parallèle, pour éviter toute coupure de service pendant la bascule.

### Quelle est la meilleure base de données pour Keycloak en production ?

PostgreSQL est la base recommandée par le projet et la plus testée en production. MySQL, MariaDB, Oracle et Microsoft SQL Server sont également supportés. La base H2 embarquée livrée par défaut ne doit jamais servir au-delà d’un test local : elle ne supporte ni la montée en charge ni des sauvegardes fiables.

### À lire aussi

**Sources et documentation officielle :** Keycloak.org, documentation d’administration Keycloak, dépôt GitHub du projet, Fondation Eclipse, spécification OpenID Connect, RFC 7519 (JSON Web Token) et la directive NIS2 de la Commission européenne.
