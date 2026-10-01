---
id: collect-261001-huawei/huawei/fr-review-elevate-your-cloud-experience-hpe-greenlake-for-private-cloud-business-b0180b1a-2
title: "fr-review-elevate-your-cloud-experience-hpe-greenlake-for-private-cloud-business-b0180b1a"
domain: huawei
role: reference
task: reference
actors: ["AWS", "Microsoft"]
dates: []
keywords: ["aws"]
source: docs/RAG/collect-261001-huawei/fr-review-elevate-your-cloud-experience-hpe-greenlake-for-private-cloud-business-b0180b1a.md
source_anchor: ""
source_lines: [32, 85]
sha256: 1d6ac7ce03610ee0ac7138794a7e06649876641b1d100a81fff080c602814f9f
---

# fr-review-elevate-your-cloud-experience-hpe-greenlake-for-private-cloud-business-b0180b1a

- Fournisseur de services (pour cet examen, nous avons testé AWS ; Azure est également disponible)
- Surnom du compte – Le compte AWS auquel celui-ci sera attaché
- Région
- Nom de la paire de clés – AWS associera ce nom à la clé publique générée
Après avoir saisi les informations de base, l'écran suivant présentera plus de 2200 XNUMX images à sélectionner pour votre instance de machine virtuelle.
Le dernier écran consiste à choisir le type d'instance. Les options répertoriées sont la taille, l'architecture et si elle est disponible dans l'offre gratuite AWS.
Comme pour le cloud privé, pendant le provisionnement des machines virtuelles, vous pouvez afficher l'état et tous les messages de journal essentiels générés au cours du processus.
Mise à jour des systèmes
La fonctionnalité de mise à jour des systèmes permet aux administrateurs de mettre à jour sans interruption les versions de HPE Alletra Storage Array, ESXi et du micrologiciel du serveur pour un cluster entier. Ces mises à jour sont exécutées dans la section Systèmes de la console de gestion.
Dans cette section de la console de gestion, cliquez sur le bouton Logiciel dans le coin droit de la liste Systèmes pour afficher les catalogues de logiciels actuellement installés dans la pile et le catalogue disponible à mettre à jour dans la pile.
Un catalogue contient des versions spécifiques de Array OS, ESXi, HPE Storage Connection Manager et Service Pack pour ProLiant. Les versions individuelles du catalogue sont visibles en survolant le numéro de version du catalogue affiché dans la ligne de la pile que vous examinez.
Pour commencer le processus de mise à niveau, sélectionnez le cluster que vous souhaitez mettre à niveau et cliquez sur Exécuter la vérification préalable, ou explorez ce système et sélectionnez Exécuter la vérification préalable dans le menu d'actions. Ensuite, sélectionnez la version du catalogue de destination vers laquelle vous souhaitez mettre à niveau la pile de cluster.
Une fois la mise à niveau commencée, elle sera soumise à un ensemble complet de vérifications préalables pour garantir qu'elle est préparée et qu'elle n'entraînera pas de dommages ni de perte de données sur le système.
Une fois terminé, vous recevrez une notification et la vue principale du logiciel mettra également à jour son apparence pour refléter les modifications.
Bilans de santé
Maintenir et confirmer la santé du matériel physique est une responsabilité impérative de chaque administrateur système. L'application Health Checks fournit une automatisation permettant de gagner du temps, garantissant le bon fonctionnement des applications et de l'infrastructure et le respect des meilleures pratiques par l'environnement cloud. Pour valider cela, HPE GreenLake charge le système d'exécuter des vérifications par rapport à environ 75 règles ou directives de configuration pour garantir que l'infrastructure fonctionne dans des conditions optimales.
À titre d'exemple concret, en cas de problème d'accès à une banque de données de machine virtuelle, les contrôles d'intégrité guideront l'administrateur vers l'hôte concerné parmi tous les clusters présentant ce problème. Dans ce cas précis, au sein du portail HPE GreenLake, la méthode la plus efficace pour exécuter les contrôles d'intégrité nécessaires sur les piles d'infrastructure Alletra dHCI consiste à accéder à la page du portail du système concerné via le menu de gauche et à sélectionner « Systèmes » . Une fois sur cette page, choisissez la pile dHCI spécifique sur laquelle vous souhaitez effectuer un contrôle d'intégrité.
Intégrations avec l'automatisation de la console Cloud Data Services
Comprendre comment HPE GreenLake for Private Cloud Business Edition s'intègre dans l'architecture HPE Data Services est essentiel. Comme le montre le diagramme ci-dessous, la console Data Services Cloud (DSCC) comporte plusieurs couches. En commençant par le bas se trouve la couche d’infrastructure de données cloud native. Ceci est représenté par les clusters sur site exécutant des machines virtuelles telles que les clusters HPE Alletra dHCI ou les clusters HPE SimpliVity, y compris les clusters cloud exécutés dans l'un des hyperscalers pris en charge.
La couche intermédiaire est constituée des services d'infrastructure Cloud, auxquels s'intègre Private Cloud Business Edition. Il comprend également des services tels que le service de configuration pour automatiser les déploiements de clusters sur les sites afin de gagner du temps et Data Ops Manager qui peut être utilisé pour configurer des services de blocage tels que la réplication entre clusters ou l'analyse des performances pour détecter les problèmes de performances de stockage.
La couche supérieure est constituée des services de données Cloud, avec des exemples tels que HPE GreenLake pour la sauvegarde et la restauration.
Sauvegarde et récupération
Chaque ingénieur système sait que la sauvegarde et la restauration font partie des domaines d'administration les plus critiques. HPE le comprend également et consacre un portail entier à cet objectif.
Le tableau de bord principal fournit un aperçu du résumé de l'inventaire, du nombre de tâches et de tâches de protection, des avertissements et des problèmes, ainsi que de la consommation des données issues de la sauvegarde dans le cloud et sur site.
Dans le menu de gauche, vous disposez d'un large éventail d'options pour gérer les sauvegardes et les politiques de tous les systèmes gérés dans HPE GreenLake, notamment :
- Politiques de protection
- Amazon Web Services
  - Magasin de blocs élastiques
  - EC2
  - Grappes EKS
  - Service de base de données relationnelle
- Serveurs Microsoft SQL
  - Bases de données
  - Cas
  - Groupes de protection
  - Hôtes d'applications
- VMware
  - Machines virtuelles
  - Banques de données
  - Groupes de protection
  - Serveurs vCenter
- Volumes de baies HPE
- Rapports
- Configuration sur site
  - Orchestrateurs de données
  - Passerelles du magasin de protection
  - Magasins de protection
  - MagasinUne fois
Réflexions finales
HPE GreenLake pour Private Cloud Business Edition, s'appuyant sur HPE GreenLake Cloud Services, offre aux entreprises une solution de cloud hybride sur mesure. Contrairement aux fournisseurs traditionnels, à savoir Amazon Web Services et Microsoft Azure, HPE GreenLake combine l'agilité de la consommation du cloud avec le contrôle de l'infrastructure sur site, mettant l'accent sur la localisation des données pour les entreprises soucieuses de conformité et de sécurité. Il s'appuie sur une infrastructure de données cloud native composée de HPE Alletra dHCI pour les charges de travail critiques pour l'entreprise ou de HPE SimpliVity pour les sites périphériques distribués. Il s'intègre parfaitement aux instances EC2 et Microsoft Azure d'Amazon. Les utilisateurs bénéficient d'une configuration, d'un déploiement, d'un provisionnement de machines virtuelles rationalisés, d'outils de sauvegarde/récupération complets et de contrôles efficaces de l'état du système, garantissant une expérience cloud robuste et efficace.
* HPE GreenLake a déployé l'accès au cloud public aux machines virtuelles Azure en décembre. Cela signifie que vous pouvez désormais créer et gérer des machines virtuelles dans l'écosystème Microsoft de la même manière que vous les gérez actuellement avec AWS EC2.
HPE GreenLake pour l'édition professionnelle du cloud privé
Démo minute 8
Démo minute 3
