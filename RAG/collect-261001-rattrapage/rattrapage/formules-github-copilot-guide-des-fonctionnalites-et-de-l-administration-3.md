---
id: collect-261001-rattrapage/rattrapage/formules-github-copilot-guide-des-fonctionnalites-et-de-l-administration-3
title: "Ignore the /src/some-dir/kernel.rs file in this repository."
domain: rattrapage
role: reference
task: reference
actors: ["Anthropic", "Microsoft"]
dates: ["2026-05-01", "2026-05-31"]
keywords: ["agents", "attribution", "claude", "copilot", "open source"]
source: docs/RAG/collect-261001-rattrapage/formules-github-copilot-guide-des-fonctionnalites-et-de-l-administration.md
source_anchor: ""
source_lines: [160, 265]
sha256: 3a6590195fc71790b4cf767c7e64ea529fb92d7fc3827ad8a63aec9a579687e3
---

# Ignore the /src/some-dir/kernel.rs file in this repository.

```
REPOSITORY-REFERENCE:
  - "/PATH/TO/DIRECTORY/OR/FILE"
  - "/PATH/TO/DIRECTORY/OR/FILE"
  - …
```
Conserver le `REPOSITORY-REFERENCE` fait partie intégrante du paramétrage. Les bases de configuration courantes doivent prioriser les identifiants sensibles, les profils d’orchestration de production, les modules algorithmiques propriétaires sensibles ou les dossiers fortement réglementés.

### Comment les exclusions s’appliquent aux fonctionnalités de Copilot

Lorsqu’une correspondance d’exclusion est détectée, l’isolation des données est totale sur tous les sous‑systèmes Copilot :

- **Complétions en ligne :** impossibilité de générer du contexte à l’intérieur du fichier ou d’en tirer pour alimenter les fichiers adjacents.
- **Copilot chat / agents :** le système renvoie un avis indiquant que le fichier ne peut pas être examiné en raison des politiques de l’organisation.

Les moteurs IDE locaux standards fonctionnent de la même manière. Les fonctions de confort comme le parsing de texte, la coloration syntaxique et l’IntelliSense localisée compilent normalement, car la couche d’exclusion s’applique explicitement aux flux de télémétrie externes de Copilot.

Les administrateurs doivent tester soigneusement les modèles de chemins dans des dépôts de préproduction ; des jokers mal formés peuvent échouer en mode « ouvert » et exposer des données que vous souhaitiez isoler.

## Gestion des politiques à l’échelle de l’organisation

Appliquer des politiques GitHub Copilot au niveau de l’organisation garantit que la sécurité de l’entreprise est définie par l’équipe d’administration, et non par les préférences individuelles des développeurs.

### Paramètres de politique disponibles

Les organisations peuvent contrôler plusieurs paramètres pour les développeurs :

- **Activation des fonctionnalités :** activer ou désactiver globalement Copilot Chat dans les environnements de développement, les interfaces en ligne de commande (via Copilot CLI) ou les systèmes avancés de revue de code agentiques.
- **Filtre de code public :** mécanisme juridique qui empêche Copilot de proposer des suggestions trop proches de dépôts open source publics sur GitHub, réduisant les risques de non‑conformité aux licences.
- **Restrictions de choix de modèles :** limiter les modèles (par exemple, variantes spécifiques de GPT ou Claude) que les développeurs peuvent sélectionner, afin de gérer la latence, la consommation de crédits et la performance. Pour un aperçu des modèles disponibles sur la plateforme GitHub, consultez ce guide pratique des GitHub Models.
- **Instructions personnalisées d’organisation :** injecter des fichiers markdown standard qui ajoutent les conventions de code, cadres de sécurité et paradigmes d’architecture à chaque requête émise par vos développeurs.

Si votre équipe maîtrise moins bien le modèle d’organisation et d’autorisations de GitHub, le cours Intermediate GitHub Concepts fournit un bon socle. Pour les équipes d’ingénierie qui généralisent les outils en ligne de commande, voyez notre Tutoriel GitHub Copilot CLI.

### Héritage des politiques au niveau Enterprise

Dans les grands environnements d’entreprise, le moteur de politiques suit une cascade d’héritage stricte : politique Enterprise > politique d’organisation > préférences utilisateur

Les administrateurs Enterprise peuvent verrouiller des politiques globalement sur toutes les entités, autoriser des dérogations sélectives par organisation, ou déléguer totalement le contrôle dans la hiérarchie. Par exemple, l’entreprise peut imposer des paramètres globaux pour restreindre l’usage de certains modèles.

Au niveau d’une équipe, elle peut imposer au pôle services financiers des filtres de code public stricts, tout en autorisant davantage d’expérimentation au pôle R&D logiciel interne.

## Journaux d’audit

Quand les auditeurs de conformité demandent une vérification de votre chaîne d’approvisionnement logicielle, ou que les équipes sécurité doivent remonter une fuite de données, GitHub Copilot journalise les modifications de la plateforme.

### Événements Copilot dans le journal d’audit

Le système consigne un registre complet des opérations d’administration, notamment :

- Attributions et révocations explicites de sièges, changements de groupes de facturation.
- Modifications du filtre de duplication de code public.
- Changements des modèles d’exclusion de fichiers et de répertoires.
- États d’activation des fonctionnalités (ex. : activation des modes de revue de code agentiques)

Le niveau de granularité dépend de votre abonnement. Les formules Business se concentrent sur les flux d’actions au périmètre de l’organisation, tandis que les comptes Enterprise donnent accès à une télémétrie forensique inter‑organisations.

### Recherche, filtrage et export

Les flux de journaux d’audit sont accessibles nativement via le panneau Organization Settings. Les administrateurs peuvent interroger l’interface avec des qualificateurs d’action spécifiques :

```
# Filtrer pour identifier qui a modifié les droits d'accès à Copilot
action:copilot.cfb_seat_assignment_created
# Identifier les changements des exclusions globales dans une période
action:copilot.content_exclusion_updated created:2026-05-01..2026-05-31
```
Les comptes Enterprise permettent de diffuser ces événements d’audit vers des SIEM externes (comme Splunk ou Datadog) pour l’alerte automatisée et la conservation centralisée et immuable.

## Gérer les sièges Copilot avec l’API REST

L’attribution manuelle de sièges depuis un tableau de bord convient aux petites équipes, mais ne passe pas à l’échelle avec des flux d’onboarding massifs. Utiliser les endpoints seats de l’API REST de GitHub Copilot vous permet de traiter l’identité et les accès entièrement en tant que code.

C’est l’une de mes parties préférées de l’administration Copilot, car elle transforme la gestion des licences en un processus que les équipes d’ingénierie peuvent automatiser proprement.

### Principaux endpoints API

Les workflows API courants incluent :

- Lister les attributions de sièges
- Attribuer des sièges
- Retirer des sièges
- Récupérer des métriques d’usage
- Lire les paramètres Copilot de l’organisation

L’authentification requiert généralement :

- Des jetons d’accès personnels à privilèges fins
- Des autorisations d’application GitHub
- Des privilèges d’administrateur d’organisation

Pour accéder à ces fonctions d’administration, vos scripts d’intégration doivent s’authentifier avec un Personal Access Token (PAT) doté des étendues admin:org ou s’exécuter via une application GitHub autorisée avec des privilèges explicites de gestion Copilot au niveau de l’organisation.

Pour approfondir les intégrations programmatiques de la plateforme, je vous recommande de suivre notre parcours de compétences GitHub Foundations.

### Schémas d’automatisation courants

Parmi les schémas pratiques :

- 
**Onboarding identitaire automatisé :** connecter un SIRH (Workday, Okta…) directement à GitHub via des webhooks. Lorsqu’un ingénieur rejoint une équipe donnée, un script déclenche une requête`POST` pour lui provisionner automatiquement Copilot.
- 
**Récupération des sièges inactifs :** un script Cron planifié interroge l’utilisation active via l’API. Si un utilisateur n’a pas utilisé Copilot depuis plus de 30 jours, le script exécute un`DELETE` pour récupérer la licence et préserver le pool de crédits.
- 
**Tableaux de bord financiers :** extraction quotidienne des données d’allocation et de consommation pour les injecter dans des plateformes de BI internes (telles que Tableau) et faciliter la refacturation par centre de coûts.

