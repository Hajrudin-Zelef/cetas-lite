---
id: collect-261001-rattrapage/rattrapage/formules-github-copilot-guide-des-fonctionnalites-et-de-l-administration-2
title: "Ignore the /src/some-dir/kernel.rs file in this repository."
domain: rattrapage
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: ["agent", "attention", "copilot"]
source: docs/RAG/collect-261001-rattrapage/formules-github-copilot-guide-des-fonctionnalites-et-de-l-administration.md
source_anchor: ""
source_lines: [74, 159]
sha256: 40684b1a9c1c590daf8b3655475e9dc64a80ab818b2071fb95112d2b61f3b3f2
---

# Ignore the /src/some-dir/kernel.rs file in this repository.

### Traitement des données et paramètres d’entraînement par défaut

Pour les équipes qui gèrent des systèmes propriétaires, la confidentialité des données est généralement le facteur décisif entre des formules personnelles et une souscription Business.

En avril 2026, GitHub a modifié la collecte des données d’interaction pour les formules individuelles Copilot. Pour les utilisateurs Free, Pro et Pro+, les données d’interaction peuvent désormais servir à l’entraînement des modèles par défaut, sauf si l’utilisateur se désinscrit explicitement.

Clarifions la différence entre le code au repos et les données d’interaction, afin de savoir ce qui est utilisé pour l’entraînement de l’IA :

- **Code au repos :** le code brut présent dans votre dépôt privé n’est pas lu ni intégré à des jeux d’entraînement publics.
- **Données d’interaction :** inclut les invites, requêtes de chat, contexte du curseur, blocs de code environnants transmis via l’API de l’IDE pendant les sessions actives, métriques d’acceptation des suggestions et journaux de feedback.

Les contrats Business et Enterprise garantissent strictement que les données d’interaction ne sont jamais utilisées à des fins d’entraînement, quelles que soient les circonstances. Aucune action manuelle de l’utilisateur n’est requise.

Pour approfondir l’usage des données et la résolution des problèmes dans Copilot, lisez notre guide GitHub Copilot : confidentialité et dépannage.

### Indemnisation en propriété intellectuelle

GitHub Copilot Business et Enterprise incluent une indemnisation en propriété intellectuelle (PI) pour le code généré. Les formules individuelles n’en bénéficient pas.

Concrètement, l’indemnisation signifie que GitHub s’engage contractuellement à fournir une protection juridique dans des circonstances spécifiées si le code généré entraîne des litiges de PI. Cela ne supprime pas tous les risques juridiques, mais modifie la discussion sur la responsabilité pour les équipes qui livrent des logiciels commerciaux.

Un freelance qui livre du code à des clients doit y prêter attention. La différence entre « outil de productivité personnel » et « plateforme de développement soutenue par l’organisation » devient très concrète dès lors qu’entrent en jeu des contrats et des livraisons commerciales.

### Facturation, sièges et passage aux AI Credits

La facturation individuelle est en libre-service et liée aux comptes personnels. Les formules Business centralisent la facturation avec des sièges attribués par l’administrateur. De plus, au lieu que chaque utilisateur gère un lot de crédits indépendant, l’organisation mutualise ses AI Credits mensuels selon le nombre d’utilisateurs.

Les formules Enterprise vont plus loin, avec des limites d’application budgétaire granulaires, des regroupements par centres de coûts et des allocations par département afin d’éviter qu’un seul groupe de développement, via des workflows agentiques intensifs, n’épuise tout le stock de crédits de l’entreprise.

## SKUs et considérations de confidentialité

Comprendre les protections de confidentialité et les SKUs est essentiel. Les frontières architecturales régissant les flux de données, les protections juridiques et le suivi selon les différents paliers sont résumées ci‑dessous :

| **Niveau de formule** | **Données d’interaction utilisées pour l’entraînement ?** | **Indemnisation PI contractuelle ?** | **Exclusions de contenu / fichiers ?** | **Accès aux journaux d’audit ?** | 
| Free | Oui (opt‑out possible) | Non | Non | Non | 
| Student | Oui (opt‑out possible) | Non | Non | Non | 
| Pro | Oui (opt‑out possible) | Non | Non | Non | 
| Pro+ | Oui (opt‑out possible) | Non | Non | Non | 
| Business | Non | Oui | Oui | Oui | 
| Enterprise | Non | Oui | Oui | Oui | 

### Changements de politique d’entraînement d’avril 2026

Le passage d’un modèle en opt‑in à un cadre en opt‑out pour les formules individuelles constitue un vecteur majeur de fuite de conformité. La charge utile des données d’interaction capturée automatiquement pendant une session IDE active inclut :

- Des historiques de chat détaillés et le contexte des invites.
- Des suggestions de code multi‑lignes et des taux d’acceptation locaux.
- Le contexte du curseur de l’éditeur actif, qui récupère souvent le contexte des fichiers adjacents, les instructions d’import et les déclarations de variables des onglets ouverts.

Imaginez qu’un développeur utilise un compte personnel Copilot Pro dans un dépôt d’entreprise. Si l’entraînement reste activé, les données d’interaction liées à cette session peuvent entrer dans l’écosystème d’entraînement de GitHub. C’est une raison courante pour laquelle les organisations adoptent les formules Business.

### Choisir le bon SKU selon vos exigences de confidentialité

Selon la nature des travaux, vous n’aurez pas besoin du même SKU.

- **Développeur solo / projets personnels :** les formules Free ou Pro offrent une flexibilité maximale. Désactivez simplement l’entraînement dans vos paramètres de confidentialité si vous travaillez sur du code propriétaire.
- **Freelances / sous‑traitants :** la formule Business fournit une barrière de protection. Les contrats clients interdisent souvent l’envoi de données à des fournisseurs de LLM externes ; un siège dédié au sein de l’organisation protège vos engagements.
- **Équipes en entreprise avec obligations de conformité** : la formule Business constitue la base standard, garantissant l’isolation des flux de données et la gouvernance administrative.
- **Secteurs réglementés (finance, santé) :** la formule Enterprise est généralement indispensable, permettant l’intégration à des configurations de sécurité spécialisées, des exigences strictes de résidence des données et des couches d’affinement localisées.

## Exclure des fichiers spécifiques de Copilot

La mise en place de règles d’exclusion de fichiers dans GitHub Copilot est l’un des moyens les plus efficaces de sécuriser votre environnement. L’exclusion de contenu empêche l’agent IDE local de traiter certains fichiers, les rendant totalement invisibles pour les complétions en ligne, les boîtes de dialogue et les opérations agentiques en arrière‑plan.

Notez que GitHub Copilot CLI, l’agent cloud de Copilot et le mode Agent dans Copilot Chat dans les IDE ne prennent pas en charge l’exclusion de contenu.

### Configurer les règles d’exclusion

Les équipes d’administration peuvent appliquer des exclusions au niveau global des paramètres de l’organisation ou au niveau des dépôts ciblés. Il suffit d’ouvrir les paramètres du dépôt ou de l’organisation en cliquant sur le bouton Settings en haut à droite.

Choisissez « Code and automation » dans les paramètres Copilot de la barre latérale. Renseignez ensuite vos exclusions dans la zone « Paths to exclude in this repository » comme suit :

```
# Ignore the /src/some-dir/kernel.rs file in this repository.
- "/src/some-dir/kernel.rs"
# Ignore files called secrets.json anywhere in this repository.
- "secrets.json"
# Ignore all files whose names begin with secret anywhere in this repository.
- "secret*"
# Ignore files whose names end with .cfg anywhere in this repository.
- "*.cfg"
# Ignore all files in or below the /scripts directory of this repository.
- "/scripts/**"
```
**Le paramétrage au niveau de l’organisation est similaire, sauf que l’option se trouve sous « Repositories and Paths to exclude » et utilise le format suivant :**

