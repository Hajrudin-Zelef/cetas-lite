---
id: collect-261001-cisco/cisco/github-copilot-enterprise-spaces-et-api-des-metriques-dusage-1
title: "github-copilot-enterprise-spaces-et-api-des-metriques-dusage"
domain: cisco
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: ["copilot"]
source: docs/RAG/collect-261001-cisco/github-copilot-enterprise-spaces-et-api-des-metriques-dusage.md
source_anchor: ""
source_lines: [1, 130]
sha256: 6d07d1b1db148ea8141b123c7f9c230fd05c0a688d83e0e7e2d2d87bbf759a95
---

# github-copilot-enterprise-spaces-et-api-des-metriques-dusage

Cours

Vous avez déployé GitHub Copilot Enterprise dans l’organisation, attribué les licences, configuré les politiques, et vos développeurs l’utilisent déjà dans leurs IDE. Il faut maintenant répondre aux questions difficiles :

- Comment optimiser Copilot pour qu’il apprenne mieux le contexte d’ingénierie propre à votre entreprise ?
- Comment mesurer la valeur de Copilot ? Quels départements l’adoptent avec succès et lesquels l’ignorent totalement ?

C’est là que GitHub Copilot Spaces et l’API des métriques d’usage entrent en jeu. Spaces permet à Copilot d’ingérer le savoir technique de votre organisation. L’API des métriques d’usage aide les administrateurs à mesurer l’adoption, la rétention et les tendances de productivité à l’échelle de l’entreprise.

Dans cet article, nous verrons :

- Ce que comprend GitHub Copilot Enterprise
- Comment fonctionnent Copilot Spaces
- Comment configurer Spaces à l’échelle
- Les endpoints de l’API des métriques d’usage de GitHub Copilot
- Les workflows d’authentification et de reporting
- Des stratégies concrètes pour mesurer le ROI

Si vous n’êtes pas à l’aise avec les organisations GitHub, les pull requests et les modèles d’autorisations, le cours Intermediate GitHub Concepts couvre ces fondamentaux. Si vous découvrez aussi Copilot, notre tutoriel How to Use GitHub Copilot présente les fonctionnalités de base sur lesquelles ce guide s’appuie.

## Renforcez la confidentialité et la gouvernance de vos données

Garantissez la conformité et protégez votre entreprise avec DataCamp for Business. Des cours spécialisés et un suivi centralisé pour protéger vos données.

## Qu’est-ce que GitHub Copilot Enterprise ?

GitHub Copilot Enterprise se situe au sommet de la gamme des offres Copilot de GitHub.

Comparé à GitHub Copilot Business ou Pro+, Enterprise met l’accent sur la gouvernance, le contexte organisationnel et les capacités de mesure. Il est conçu pour les entreprises gérant de vastes environnements d’ingénierie, plutôt que pour des développeurs individuels ou de petites équipes.

Deux capacités comptent surtout dans la pratique :

1. Contexte organisationnel personnalisé via les **Spaces**
2. Télémétrie à l’échelle de l’organisation via l’**API des métriques d’usage**

Ces deux fonctionnalités font passer Copilot d’un simple « autocomplétion intelligente » à quelque chose qui s’apparente à une véritable plateforme interne d’ingénierie propulsée par l’IA.

Les entreprises qui tirent le plus de valeur de GitHub Copilot Enterprise l’intègrent comme un élément clé de leur infrastructure interne. Elles soignent le contexte organisationnel, mesurent l’adoption en continu et ajustent les politiques à partir des données d’usage, pas d’hypothèses.

Pour un tour d’horizon plus large de l’écosystème GitHub, nous vous recommandons notre guide Introduction to GitHub Products.

### En quoi Enterprise diffère de Business et Pro+

GitHub Copilot Enterprise étend l’offre Business avec notamment :

- Métriques d’usage au niveau de l’organisation
- Contrôles de gouvernance renforcés
- Héritage des politiques à l’échelle de l’entreprise
- Quotas supérieurs pour les requêtes premium (1 000 contre 300 en offre Business)
- Accès et gestion de modèles supplémentaires

Enterprise nécessite GitHub Enterprise Cloud en plus de l’abonnement Copilot Enterprise. Cela ajoute un coût supplémentaire par utilisateur : assurez-vous donc que votre organisation a réellement besoin de gouvernance, de télémétrie et d’administration au niveau entreprise.

| **Fonctionnalité** | **Pro+** | **Business** | **Enterprise** | 
| Usage individuel | Oui | Non | Non | 
| Gestion centralisée des licences | Non | Oui | Oui | 
| Journaux d’audit | Non | Oui | Oui | 
| Exclusions de fichiers | Non | Oui | Oui | 
| Prise en charge de Spaces | Oui, avec Copilot | Oui, visibilité admin limitée | Oui, gestion complète au niveau entreprise | 
| API des métriques d’usage | Non | Niveau organisation | Niveau entreprise + organisation | 
| Héritage des politiques entreprise | Non | Non | Oui | 

**Remarque :** Les abonnés Business accèdent à l’API des métriques d’usage au niveau de l’organisation (`/orgs/{org}/…`). Les abonnés Enterprise disposent en plus de rapports agrégés au niveau de l’entreprise (`/enterprises/{enterprise}/…`) couvrant toutes les organisations dans une vue unique.

### Pour qui est GitHub Copilot Enterprise

GitHub Copilot Enterprise s’adresse aux organisations disposant d’environnements GitHub matures.

Clients Enterprise typiques :

- Grandes organisations d’ingénierie
- Secteurs réglementés
- Équipes plateformes multi-équipes
- Entreprises avec standards internes de développement
- Organisations nécessitant une gouvernance centralisée

Notez que cela n’améliore pas intrinsèquement les performances de Copilot. Cette nuance est importante : beaucoup d’équipes surdimensionnent initialement leur achat en pensant qu’Enterprise = « meilleur Copilot », alors qu’Enterprise ajoute surtout des outils de gouvernance et de mesure.

## Copilot Spaces : un contexte sur mesure pour votre organisation

Copilot Spaces résout l’une des principales limites des assistants de code généralistes.

Par défaut, Copilot maîtrise correctement le savoir public en programmation. Il ne comprend pas automatiquement vos API internes, vos décisions d’architecture, vos conventions de code, vos workflows de déploiement ou votre documentation d’onboarding.

Spaces fournit un contexte organisationnel sélectionné que Copilot peut exploiter en conversation et pour l’assistance au codage.

Concrètement, Spaces aide Copilot à répondre à des questions telles que :

- « Comment structurons-nous en interne nos gestionnaires d’API ? »
- « Quelle bibliothèque d’authentification notre équipe plateforme recommande-t-elle ? »
- « Quel workflow de déploiement ce microservice doit-il utiliser ? »
- « Quelles conventions de nommage notre équipe backend applique-t-elle ? »

### Ce que prennent en charge les Spaces

Spaces couvre une palette de contenus organisationnels plus large que l’ancien système Knowledge Bases.

Types de contenus pris en charge :

- Fichiers de code
- Documentation Markdown
- Fichiers JSON
- Fichiers téléversés
- Images
- GitHub Issues
- Pull requests

Chaque type de contenu apporte une valeur différente.

Les fichiers de code aident Copilot à comprendre les schémas d’implémentation. Les fichiers Markdown détaillent l’architecture et l’onboarding. Les pull requests exposent les discussions de revue et les décisions d’ingénierie passées. L’ensemble améliore la compréhension des pratiques de développement de votre organisation.

Point subtil mais important : Spaces n’est pas simplement une base vectorielle adossée à GitHub. Elle inclut des contrôles de partage et des workflows de gouvernance pensés pour l’entreprise.

### La fin de Knowledge Bases

GitHub a mis fin à l’ancienne fonctionnalité Copilot Knowledge Bases le 1er novembre 2025.

Spaces remplace Knowledge Bases avec :

- Un support de contenu plus large
- De meilleurs contrôles de partage
- Une administration améliorée
- Une gestion plus flexible au niveau de l’organisation

Vous trouverez encore de la documentation et des billets de blog obsolètes mentionnant Knowledge Bases. Soyez prudent avec les anciens tutoriels : beaucoup d’endpoints et de workflows ont changé entre 2025 et 2026.

## Créer et configurer des Copilot Spaces

Côté administration, créer un Copilot Space est assez simple. Le défi, c’est d’en gérer des dizaines, voire des centaines, à travers les équipes.

