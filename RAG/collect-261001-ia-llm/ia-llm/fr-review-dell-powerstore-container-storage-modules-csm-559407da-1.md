---
id: collect-261001-ia-llm/ia-llm/fr-review-dell-powerstore-container-storage-modules-csm-559407da-1
title: "fr-review-dell-powerstore-container-storage-modules-csm-559407da"
domain: ia-llm
role: reference
task: reference
actors: ["AWS"]
dates: []
keywords: ["arr", "open source"]
source: docs/RAG/collect-261001-ia-llm/fr-review-dell-powerstore-container-storage-modules-csm-559407da.md
source_anchor: ""
source_lines: [1, 36]
sha256: 8db6041f8bd8053c901115f7adbea28a52390d456528f85fedceae55ab6564a6
---

# fr-review-dell-powerstore-container-storage-modules-csm-559407da

Alors que le paysage numérique a continué d'évoluer, les solutions de stockage capables de prendre en charge à la fois un large éventail de charges de travail d'entreprise et les plates-formes de développement d'applications émergentes sont devenues de plus en plus importantes. Notamment, Dell Technologies PowerStore est une baie de stockage principale conçue pour répondre aux besoins des charges de travail d'entreprise traditionnelles et des applications conteneurisées modernes exécutées sur Kubernetes (K8).
Les conteneurs ont radicalement changé la façon dont les charges de travail sont fournies dans le centre de données moderne. À certains égards, il est difficile de croire que depuis leur création il y a moins de dix ans, ils ont établi une empreinte aussi forte et croissante dans le centre de données. Cependant, une fonctionnalité qui fait toujours défaut aux conteneurs au cours de cette période est le stockage persistant avec des fonctionnalités d'entreprise. Pour remédier à cette insuffisance, Dell Technologies a doté sa gamme PowerStore d'une interface de stockage de conteneur (CSI) et de modules de stockage de conteneur (CSM).
Stockage de conteneurs
Lorsque les conteneurs ont été publiés pour la première fois, ils étaient conçus pour être sans état et n'avaient pas de stockage persistant. Cependant, à mesure que les entreprises découvraient l'utilité des conteneurs et que leurs cas d'utilisation augmentaient, il est devenu évident qu'elles auraient besoin d'un stockage persistant. Cet écart a finalement conduit à l'introduction de deux solutions : les plug-ins de volume, suivis rapidement par des conteneurs fournis par CSI avec un autre niveau d'abstraction pour le stockage.
Le pilote CSI est un composant logiciel qui permet aux systèmes d'orchestration de conteneurs tels que les K8 de communiquer avec les systèmes de stockage. Il fournit une API standard pour créer, supprimer et gérer des volumes de stockage et des services de données, permettant aux fournisseurs de stockage d'interagir avec divers systèmes d'orchestration de conteneurs.
Dell PowerStore – Stockage prêt pour les conteneurs
Dell a publié un ensemble de pilotes CSI pour ses séries de périphériques de stockage PowerFlex, PowerScale, Unity, PowerMax et PowerStore. Le schéma ci-dessous montre comment un système de stockage Dell peut s'interfacer avec des K8 à l'aide de CSI.
Le CSI de Dell s'est avéré très populaire. Par exemple, le pilote PowerScale CSI a enregistré plus de cinq millions de téléchargements et PowerStore a enregistré 1.7 million de téléchargements en seulement deux ans.
Modules de stockage de conteneurs Dell PowerStore
Bien que l'utilisation des CSI soit extrêmement puissante, les utilisateurs d'entreprise en sont venus à s'attendre à des fonctionnalités de stockage d'entreprise pour leur stockage. Avec le lancement par Dell des modules de stockage de conteneurs (CSM), la gestion du stockage de conteneurs est devenue considérablement moins complexe.
Dans le cas de Dell PowerStore, les modules Dell CSM permettent aux applications K8 de fournir des services de données, tels que des instantanés de groupe, des clones légers, la réplication, le chiffrement, etc. Ces fonctionnalités sont prises en charge pour les objets de stockage, y compris les protocoles bloc et fichier sur Fibre Channel, iSCSI, NVMe/TCP et NFS.
Dell CSM pour PowerStore respecte deux principes fondamentaux. Tout d'abord, les développeurs d'applications peuvent gérer les tâches de stockage courantes via le plan de contrôle K8s, ce qui évite à l'équipe DevOps d'avoir à se rendre chez l'administrateur de stockage. Deuxièmement, l'équipe DevOps peut accéder aux performances et aux services de données approfondies offerts par PowerStore sans modifier les flux de travail actuels. Cela garantit que les administrateurs de stockage peuvent fournir des applications exactement comme ils en ont besoin pour respecter les SLA.
Dell a rendu open source l'ensemble de ses développements CSM et CSI, non seulement pour PowerStore, mais aussi pour nombre de ses autres systèmes de stockage, sur GitHub . De plus, Dell propose des modules fournisseurs Ansible et Terraform qui facilitent la gestion de diverses tâches telles que le provisionnement du stockage, la modification des configurations de stockage et les paramètres de réplication locale/distante.
Pour garantir la compatibilité, le CSM de Dell a été validé pour fonctionner avec Amazon EKS, Mirantis, VMware Tanzu, SUSE Rancher et d'autres plateformes K8 et d'orchestration de conteneurs.
Le CSM comprend plusieurs modules qui offrent des fonctionnalités que les entreprises clientes de Dell attendent.
Les modules de stockage conteneurisé Dell PowerStore sont régulièrement mis à jour. Pour consulter la liste complète des modules actuellement pris en charge, rendez-vous sur le dépôt GitHub du projet.
Avec cet arrière-plan présenté, examinons les spécificités de quelques-uns de ces modules.
Modules de stockage en conteneur Dell PowerStore pour la réplication et la résilience
CSM for Replication apporte les capacités de réplication et de reprise après sinistre de Dell PowerStore aux clusters K8s d'une organisation. Il utilise la technologie de réplication native disponible sur la baie PowerStore pour fournir un moyen de redémarrer les applications en cas de migration planifiée et non planifiée. La réplication est prise en charge sur les clusters K8 étendus et répliqués, et la fonctionnalité résiliente est activée dans le graphique Helm lors de l'installation du CSM.
La fonctionnalité de résilience utilise un moniteur de pod pour protéger les applications avec état de diverses défaillances. Il est conçu pour détecter les types de pannes suivants : pannes de nœud Kubernetes, pannes de plan de contrôle et pannes de réseau d'E/S. Une panne de nœud se produit lorsqu'un événement, tel qu'une panne de courant, arrête un nœud entier. Une défaillance du réseau IO est détectée en interrogeant la baie pour voir si elle dispose d'une connexion saine au nœud.
Actuellement, PowerStore Replication prend en charge les éléments suivants :
- Réplication des données à l'aide de la réplication basée sur la baie de stockage native
- Réplication d'un volume de bloc asynchrone
- Création d'objets PersistentVolume dans le cluster représentant le volume répliqué
- Création d'objets DellCSIReplicationGroup dans le cluster
- Création d'objets DellCSIReplicationGroup dans le cluster Fournit via un utilitaire de ligne de commande, repctl, qui est utilisé pour configurer et gérer les ressources liées à la réplication sur plusieurs clusters
La commande repctl est un outil CLI qui facilite les procédures liées à la réplication sur plusieurs clusters K8.
Pour démontrer comment cela fonctionne, nous avons configuré deux baies de stockage dans notre laboratoire StorageReview : RT-D0355 et RT-D0338. Nous avons ensuite exécuté les commandes repctl cluster list et repctl list rg sur la console, la sortie répertoriant les clusters et les groupes de réplication.
Notez qu'aucun objet n'est affiché dans les clusters.
Lorsqu'une application K8s est déployée, son stockage backend sera déployé sur les deux baies. Cela peut être vu en utilisant la ligne de commande.
Notez que chaque nœud du cluster contient un objet.
Cela peut également être consulté sur le portail Web PowerStore.
Si un basculement est déclenché, celui-ci sera indiqué comme étant en cours.
Cela peut également être vu dans le portail.
La commande repctl est également utilisée pour reprotéger les objets.
La commande repctl permet aux utilisateurs d'interagir avec leur stockage par programmation via la ligne de commande. Dell propose un document décrivant différentes manières d'utiliser la réplication dans les procédures de reprise après sinistre.
