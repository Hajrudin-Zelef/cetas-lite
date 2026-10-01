---
id: collect-261001-ia-llm/ia-llm/fr-review-dell-powerstore-container-storage-modules-csm-559407da-2
title: "fr-review-dell-powerstore-container-storage-modules-csm-559407da"
domain: ia-llm
role: reference
task: reference
actors: []
dates: []
keywords: ["agent", "open source"]
source: docs/RAG/collect-261001-ia-llm/fr-review-dell-powerstore-container-storage-modules-csm-559407da.md
source_anchor: ""
source_lines: [37, 48]
sha256: acc1fee8841e7adf9998ced7a962961e3195d80769e355c49038f6c5eb1e5485
---

# fr-review-dell-powerstore-container-storage-modules-csm-559407da

Modules de stockage en conteneur Dell PowerStore pour l'observabilité
La suite open source d'outils de visibilité et de création de rapports sur le stockage K8s de Dell s'appelle CSM for Observability, qui utilise des composants open source courants que l'on trouve fréquemment dans les déploiements K8s. Il dispose d'un agent OpenTelemetry qui collecte des mesures au niveau de la baie pour Dell PowerStore et les place dans une base de données Prometheus. Cela permet aux administrateurs de K8 de collecter des métriques au niveau de la baie pour vérifier la capacité et les performances globales directement à partir des outils Prometheus/Grafana plutôt que de s'interfacer directement avec le système de stockage lui-même.
CSM for Observability offre une visibilité sur la capacité et les performances des volumes et des partages de fichiers sur PowerStore qui sont gérés avec les pilotes Dell CSM CSI. Le module comprend également des tableaux de bord Grafana pré-emballés pour analyser les métriques historiques et voir la topologie entre un PV K8s et sa traduction en tant que LUN ou partage de fichiers dans la baie backend.
Déploiement
Il est possible de déployer les modules CSI et CSM avec Helm ou en utilisant les opérateurs CSI et CSM (aperçu technique pour CSM).
Réflexions finales
Dell reconnaît la position et l'importance continue des conteneurs et des K8 dans le centre de données moderne d'aujourd'hui. À ce titre, Dell a ajouté de nouvelles fonctionnalités à la gamme éprouvée d'appliances de stockage PowerStore pour répondre à ces besoins. Ces fonctionnalités permettent aux charges de travail modernes d'avoir bon nombre des mêmes fonctionnalités de stockage que les charges de travail traditionnelles. Mais Dell ne se contente pas de s'asseoir sur ses lauriers ; il continue de permettre son intégration CSI/CSM et ajoutera bientôt des modules de mobilité d'application (actuellement en préversion technique), de chiffrement et de placement de volume à ses offres.
Parmi ces solutions, la mobilité des applications nous semble la plus intéressante car elle permet aux administrateurs Kubernetes de cloner leurs charges de travail et données applicatives avec état vers d'autres clusters, sur site ou dans le cloud. La mobilité des applications utilise Velero et son intégration de Restic pour copier les métadonnées et les données applicatives vers un stockage objet.
Le travail de Dell permet aux développeurs d'applications et aux équipes DevOps de gérer davantage de provisionnement et de maintenance, économisant ainsi les ressources informatiques. Les équipes informatiques savent qu'elles fournissent les services de données et les applications de performance dont les équipes ont besoin. Un excellent exemple d'apport d'outils de niveau entreprise au stockage K8s est son modèle de réplication, qui permet de protéger automatiquement les données des objets de stockage et de la ligne de commande comme les autres services K8s.
Dell continue d'être le leader de la communauté K8 et a acquis une crédibilité considérable en intégrant son stockage à des outils couramment utilisés tels que Grafana et Prometheus. De plus, Dell travaille avec toutes les plates-formes K8 les plus populaires (VMware Tanzu, EKS, etc.), ce qui est essentiel dans le monde multi-cloud d'aujourd'hui.
Modules de stockage de conteneurs Dell
Ce rapport est parrainé par Dell Technologies. Tous les points de vue et opinions exprimés dans ce rapport sont basés sur notre vision impartiale du ou des produits à l'étude.
