---
id: collect-261001-huawei/huawei/fr-review-elevate-your-cloud-experience-hpe-greenlake-for-private-cloud-business-b0180b1a-1
title: "fr-review-elevate-your-cloud-experience-hpe-greenlake-for-private-cloud-business-b0180b1a"
domain: huawei
role: reference
task: reference
actors: ["AWS"]
dates: []
keywords: ["aws"]
source: docs/RAG/collect-261001-huawei/fr-review-elevate-your-cloud-experience-hpe-greenlake-for-private-cloud-business-b0180b1a.md
source_anchor: ""
source_lines: [1, 31]
sha256: 7944bf604a08702a1e0699dbe0568f968a6728e602b77dc75a7debeea291b4e5
---

# fr-review-elevate-your-cloud-experience-hpe-greenlake-for-private-cloud-business-b0180b1a

En juin 2020, Hewlett Packard Enterprise annonçait le lancement de HPE GreenLake Cloud Services afin de répondre aux besoins des entreprises hybrides et multicloud. Cette année, l'offre s'est encore enrichie avec HPE GreenLake for Private Cloud Business Edition. Offrant une flexibilité et des options encore plus grandes, HPE GreenLake for Private Cloud Business Edition permet aux entreprises de bénéficier d'une expérience hautement personnalisée et efficace pour choisir l'infrastructure et la capacité les mieux adaptées à leurs besoins.
Cette expérience et cette offre uniques rendent HPE GreenLake plus attrayante en tant qu'option d'infrastructure cloud hybride que les offres cloud plus omniprésentes, telles qu'Amazon Web Services et Azure. Il est spécialement conçu pour les entreprises recherchant l'agilité des modèles de consommation cloud combinée au contrôle de l'infrastructure sur site. Il offre un déploiement simplifié avec des configurations prédéfinies, permettant une configuration plus rapide et une complexité réduite. Grâce à son optimisation pour des charges de travail spécifiques, les entreprises peuvent garantir l'efficacité des performances sans les frais généraux liés à la gestion d'une vaste gamme de services souvent présents dans les grands cloud publics. De plus, pour les entreprises soucieuses de la souveraineté des données, de la conformité réglementaire ou de mesures de sécurité spécifiques, l'édition Business offre l'avantage de la localisation des données et un contrôle accru sur leur environnement.
En bref, HPE GreenLake for Private Cloud Business Edition est une solution de cloud privé autogérée avec une interface unifiée pour simplifier la gestion des VM vers l'infrastructure. Il vous permet de créer votre cloud en libre-service à la demande là où vous en avez besoin, avec un choix de facturation mensuelle prévisible ou de paiement initial.
Infrastructure
Pour démarrer avec HPE GreenLake Private Cloud Business Edition, les utilisateurs doivent d'abord provisionner le matériel pour créer leur cloud privé. Pour les sites sur site, HPE utilise principalement une infrastructure cloud native composée de l'infrastructure hyperconvergée désagrégée HPE Alletra (dHCI) pour les charges de travail critiques pour l'entreprise ou de HPE SimpliVity (véritable HCI) pour les sites de périphérie distribués. Il fournit aux entreprises un package de cloud privé complet, fidèle au concept de flexibilité de HPE GreenLake pour les organisations ayant des charges de travail variables.
HPE s'est associé à Amazon Web Services pour compléter son offre sur site afin d'utiliser ses instances EC2 en tant que segment de cloud public. La combinaison de ces deux solutions permet à HPE GreenLake de présenter un cloud hybride unifié aux entreprises à la recherche de la solution adaptée à leurs besoins.
Une fois que vous avez provisionné le matériel et préparé la configuration cloud via Amazon Web Services, vous pouvez commencer à provisionner des machines virtuelles, des politiques, des sauvegardes, des contrôles de santé et toute autre myriade de services proposés par HPE GreenLake.
SimpliVité HPE
Pour un rappel rapide sur HPE SimpliVity et une mise à jour de ses capacités, il s'agit d'une véritable appliance d'infrastructure hyperconvergée (HCI) qui combine toutes les fonctionnalités et services principaux nécessaires aux fonctionnalités du serveur (comme l'informatique, la mise en réseau et le stockage) dans une seule grande boîte. Il améliore considérablement l’efficacité et la sécurité des données ainsi que la résilience intégrée.
Auparavant, il était accessible et géré via un portail centralisé. Avec l'introduction de HPE GreenLake Private Cloud Business Enterprise, HPE SimpliVity peut être géré au sein de la plateforme HPE GreenLake.
Dans HPE GreenLake, l'accessibilité d'un système HPE SimpliVity par rapport à un système HPE Alletra est affichée dans une vue commune. Il se trouve dans le portail Systèmes et n'est délimité que par le type de système.
Politiques de provisionnement des machines virtuelles
Les politiques de provisionnement des machines virtuelles sont essentielles à la réussite d’une stratégie globale de déploiement hybride. Pour accéder au portail VM Provisioning Policy, sélectionnez le bouton de menu des barres horizontales en haut à gauche pour ouvrir la liste des options d'accès aux portails. À partir de là, choisissez Politiques de provisionnement de VM, ce qui ouvrira la page principale de toutes les politiques actuellement disponibles dans votre environnement.
Pour créer un nouveau profil de provisionnement de VM, cliquez sur l'icône « + » en haut à gauche et une fenêtre s'ouvrira pour vous permettre de commencer à remplir les informations requises. Vous devrez nommer le profil et fournir une description de base. Sélectionnez le bouton radio approprié pour Déduplication, Chiffrement des données ou All-Flash si ce sont les options dont vous avez besoin. En passant votre curseur sur l'icône d'information à droite de chaque option, vous obtenez des informations supplémentaires sur l'utilisation de ces options. Enfin, vous pouvez laisser les paramètres de performances QoS par défaut, soit la limite d'un million d'IOPS, ou les ajuster en fonction de l'environnement.
Surveillance des machines virtuelles
Une fois que vous avez créé une stratégie de provisionnement, revenez à la page Machines virtuelles pour créer les instances souhaitées.
Dans le menu supérieur, vous pouvez afficher vos VM de cloud privé (physiques) ou les VM de cloud public (AWS).
La page principale vous donne un bref aperçu des machines virtuelles de votre environnement et d'autres données fondamentales les concernant, y compris les détails de leur état et de la protection des données. Vous pouvez les trier selon l'une des caractéristiques répertoriées et même avoir la possibilité d'afficher leurs statistiques de performances et d'utilisation actuelles à partir des boutons sur le côté droit.
La vue Public Cloud n'est pas différente, si ce n'est qu'elle n'a pas autant d'informations à afficher en raison du manque de matériel physique nécessaire à la surveillance.
Création d'une machine virtuelle – Cloud privé
Pour créer une nouvelle machine virtuelle, cliquez sur l'icône (+) dans le coin supérieur gauche pour placer la fenêtre de création au premier plan.
La première section contient les informations générales : noms de machines virtuelles, nombre de machines virtuelles et si vous souhaitez qu'elles soient allumées après leur création.
Dans la section suivante, sélectionnez le cluster Hyperviseur où ces VM seront créées. Vous pouvez choisir des clusters HPE SimpliVity ou HPE Alletra dHCI.
Dans la section 3, choisissez la banque de données cible.
Ensuite, sélectionnez le modèle de système d'exploitation souhaité.
Enfin, la politique de provisionnement des machines virtuelles spécifie la politique de protection des données de classe entreprise et exploite HPE GreenLake pour la sauvegarde et la restauration.
Une fois toutes les différentes options et configurations sélectionnées, cliquez sur Créer pour soumettre la VM pour création à votre pile d'infrastructure.
Vous pouvez également vérifier l'état de mise à disposition de la VM en bas, ce qui fera également apparaître une fenêtre pour voir les tâches en cours et en attente ainsi que l'état et une brève sortie de journal de base du processus.
Création d'une machine virtuelle – Cloud public
Créer une VM dans le Cloud Public est aussi simple que d'en créer une dans le Cloud Privé. Cliquez sur l'icône (+) pour afficher la fenêtre des options de création. Sur la première page se trouvent les options initiales de la VM :
- Nom de la machine virtuelle
