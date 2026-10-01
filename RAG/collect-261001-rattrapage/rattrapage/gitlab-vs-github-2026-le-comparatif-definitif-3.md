---
id: collect-261001-rattrapage/rattrapage/gitlab-vs-github-2026-le-comparatif-definitif-3
title: ".gitlab-ci.yml"
domain: rattrapage
role: reference
task: reference
actors: ["AWS", "Apple", "Google", "Microsoft"]
dates: []
keywords: ["aws", "cyber", "open source"]
source: docs/RAG/collect-261001-rattrapage/gitlab-vs-github-2026-le-comparatif-definitif.md
source_anchor: ""
source_lines: [173, 216]
sha256: 90b42aca4c0afcf5baddd3ec35fb7ec274f48bf8ad9a6de4e02c4e9b3cee011b
---

# .gitlab-ci.yml

Pour les entreprises européennes soumises à la directive **NIS2** et au **Cyber Resilience Act**, GitLab offre un avantage réglementaire significatif avec ses fonctionnalités de compliance intégrées : frameworks de conformité personnalisables, politiques d’approbation de merge basées sur la sécurité, et rapports de conformité automatisés. GitHub offre des fonctionnalités similaires mais avec une configuration plus manuelle et une dépendance à des outils tiers.

**ThePrimeagen** résume bien la situation dans sa vidéo de comparaison DevSecOps : « Si vous faites du DevSecOps sérieux, GitLab Ultimate est difficile à battre. Tout est intégré, tout fonctionne ensemble. Avec GitHub, vous finissez par assembler une douzaine d’outils différents. C’est faisable, mais c’est du travail en plus. »

## Auto-hébergement et Souveraineté des Données

L’auto-hébergement est un critère décisif pour de nombreuses organisations européennes en 2026. Avec l’entrée en vigueur progressive du **RGPD renforcé**, de la directive **NIS2** et des exigences croissantes de souveraineté numérique en Europe, la capacité à héberger ses outils de développement sur son propre infrastructure ou dans un cloud souverain européen est devenue un avantage concurrentiel majeur.

**GitLab Community Edition (CE)** est entièrement open source et peut être auto-hébergé gratuitement sans limitation de nombre d’utilisateurs. Cette version inclut la gestion de code source, le CI/CD de base, le registre de conteneurs et les fonctionnalités de gestion de projet. Pour les entreprises nécessitant des fonctionnalités avancées (LDAP, haute disponibilité, audit logs), **GitLab Enterprise Edition** est disponible aux mêmes tarifs que la version SaaS (29 $ Premium, 99 $ Ultimate par utilisateur/mois).

**GitHub Enterprise Server** est la seule option d’auto-hébergement côté GitHub, disponible à **21 $/utilisateur/an**. C’est une application complète qui peut être déployée sur site ou dans un cloud privé. Cependant, contrairement à GitLab CE, il n’existe pas de version gratuite auto-hébergée de GitHub. De plus, GitHub Enterprise Server a historiquement eu un délai de publication des nouvelles fonctionnalités par rapport à GitHub.com.

Pour les organisations françaises qui suivent les recommandations de l’ANSSI en matière de souveraineté numérique, GitLab CE auto-hébergé sur un cloud souverain certifié SecNumCloud (comme **Bleu**, **S3NS** ou **NumSpot**) représente la solution la plus conforme. GitHub Enterprise Server est également déployable sur ces infrastructures, mais le surcoût de licence s’ajoute au coût du cloud souverain.

Un point notable en 2026 : GitLab propose désormais des **instances SaaS dédiées hébergées en Europe** (région UE), permettant de bénéficier du SaaS tout en gardant les données dans l’Espace économique européen. GitHub, via Microsoft Azure, propose également des régions européennes, mais les données transitent par l’infrastructure Microsoft soumise au **CLOUD Act américain**, ce qui peut poser problème pour certains secteurs réglementés (défense, santé, finance).

## Écosystème et Intégrations

L’écosystème d’intégrations est l’un des points forts historiques de GitHub. Le **GitHub Marketplace** compte plus de 20 000 actions pour GitHub Actions et des milliers d’applications tierces. Cette richesse permet d’intégrer facilement GitHub avec pratiquement n’importe quel outil : Jira, Slack, Datadog, Sentry, Vercel, Netlify, AWS, Azure, Google Cloud – la liste est quasi infinie.

L’intégration native avec l’écosystème **Microsoft** est particulièrement puissante : Azure DevOps, Visual Studio, VS Code, Teams, et Azure Cloud forment un ensemble cohérent. Pour les organisations déjà investies dans l’écosystème Microsoft, GitHub s’intègre de manière quasi transparente, réduisant considérablement les frictions.

GitLab adopte une philosophie « **single application** » : plutôt que de s’appuyer sur un marketplace d’intégrations tierces, la plateforme intègre nativement la majorité des fonctionnalités DevOps. Le registre de packages, le registre de conteneurs, le système de monitoring, la gestion d’infrastructure (Terraform), et même un wiki et un système de pages statiques sont tous inclus dans la plateforme de base. Cette approche réduit la dépendance aux outils tiers mais peut limiter la flexibilité pour les équipes qui préfèrent choisir chaque outil indépendamment.

GitLab propose néanmoins des intégrations avec les outils les plus courants : **Jira, Slack, Kubernetes, Terraform, Prometheus, Grafana**. L’intégration avec Atlassian est particulièrement mise en avant comme alternative au stack Microsoft de GitHub. En 2026, GitLab a également renforcé ses intégrations avec les fournisseurs cloud européens, un argument de poids pour les entreprises de l’UE.

Comme le note **Fireship** : « GitHub, c’est le Android de la gestion de code – un écosystème ouvert avec un marketplace infini. GitLab, c’est le Apple – tout est intégré, tout fonctionne ensemble, mais vous êtes dans leur jardin. Les deux approches ont du mérite, ça dépend de votre philosophie. »

## Gestion de Projet et Collaboration

Les deux plateformes proposent des outils de gestion de projet qui ont considérablement évolué en 2025-2026, réduisant le besoin d’outils tiers comme Jira ou Linear pour de nombreuses équipes.

**GitHub** offre un système d’**Issues** simple mais efficace, enrichi par les **GitHub Projects** (tableaux Kanban et vues tableur), les **Discussions** pour les échanges communautaires, et les **Milestones** pour le suivi des versions. En 2026, GitHub Projects a gagné en maturité avec l’ajout de vues timeline, de champs personnalisés et d’automatisations basées sur les workflows. La force de GitHub réside dans la simplicité : le système est intuitif et fonctionne bien pour les projets open source avec de nombreux contributeurs externes.

**GitLab** propose un système plus riche et plus structuré : **Issues** avec poids et estimation de temps, **Boards** (Kanban) multi-projets, **Epics** pour regrouper les issues en initiatives stratégiques, **Roadmaps** visuels pour la planification à long terme, et **Iterations** (sprints). La possibilité d’assigner **plusieurs responsables** par issue (disponible sur les plans payants) est un avantage pour les équipes pratiquant le pair programming ou le co-ownership de fonctionnalités.

Pour les **code reviews**, les deux plateformes sont comparables. GitHub a popularisé le concept de Pull Request et offre une expérience de revue de code fluide avec suggestions in-line, revues requises et protection de branches. GitLab utilise le terme **Merge Request** et ajoute des fonctionnalités avancées comme les approbation rules multi-niveaux, les merge trains (pour fusionner en séquence sans conflits) et les merge request pipelines qui exécutent le CI/CD sur la branche mergée avant la fusion effective.

Un avantage notable de GitLab pour les grandes organisations est la **gestion multi-projets**. Les groupes et sous-groupes permettent d’organiser des dizaines ou centaines de projets avec des permissions héritées, des pipelines cross-projets et des tableaux Kanban agrégés. GitHub offre des fonctionnalités similaires via les Organisations et Teams, mais la granularité des permissions est historiquement moins fine.

## Performances et Fiabilité en 2026

La **disponibilité** et les performances sont critiques pour des plateformes qui hébergent le code source de millions de projets. En 2026, les deux plateformes affichent des bilans contrastés.

