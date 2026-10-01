---
id: collect-261001-general-networking/general-networking/fr-review-the-fast-path-to-hybrid-cloud-dell-technologies-cloud-704fa971-1
title: "fr-review-the-fast-path-to-hybrid-cloud-dell-technologies-cloud-704fa971"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["arr", "valuation"]
source: docs/RAG/collect-261001-general-networking/fr-review-the-fast-path-to-hybrid-cloud-dell-technologies-cloud-704fa971.md
source_anchor: ""
source_lines: [1, 25]
sha256: b497c10c5425d14d3365ea2cd43e81818ce22c5a4e83d41441eff0a827c0cece
---

# fr-review-the-fast-path-to-hybrid-cloud-dell-technologies-cloud-704fa971

86 % plus rapide à déployer VMware Cloud Foundation qu'un modèle à faire soi-même
VMware Cloud Foundation sur Dell EMC VxRail (Dell Technologies APEX Hybrid Cloud) offre une solution simplifiée et rapide pour le déploiement d'un cloud hybride. Non seulement la solution est opérationnelle sur site en 14 jours [1] , mais la mise en service de VMware Cloud Foundation sur VxRail est également beaucoup plus rapide qu'une installation sur serveur dédié. En effet, lors de notre évaluation, nous avons constaté que le déploiement de VCF sur VxRail était 86 % plus rapide que le modèle d'installation sur site.
Cependant, les avantages de VCF sur VxRail se poursuivent en ce qui concerne les mises à jour du cycle de vie (matériel et logiciel). VCF continue de montrer une valeur énorme, étant 15% plus rapide, avec moins d'étapes. En fin de compte, la solution VCF sur VxRail est plus facile que le bricolage à déployer et à gérer avec des centaines de tâches automatisées. Dans cet article, nous détaillons l'installation et la configuration de VCF dans ces deux modalités et soulignons les avantages techniques du Cloud hybride Dell Technologies APEX.
Introduction
Au sein du portefeuille d'entreprise de Dell Technologies se trouve la principale solution de la société pour le centre de données défini par logiciel (SDDC) moderne. Dell EMC VxRail est une appliance d'infrastructure hyperconvergée (HCI) intégrée conçue conjointement par Dell Technologies et VMware. Alors que VMware propose des logiciels HCI via des partenaires qui vendent vSAN ReadyNodes™ (que Dell EMC propose également), VxRail va encore plus loin. Avec VxRail, les clients reçoivent un système entièrement intégré, préconfiguré et prétesté qui offre la virtualisation, le calcul et le stockage dans une seule appliance. Le fait que tous les éléments (y compris le logiciel VMware et le matériel et la mise en réseau Dell EMC PowerEdge) soient réunis dans une seule unité offre aux clients un chemin plus fluide vers le déploiement de VMware HCI.
Cependant, Dell Technologies ne s'arrête pas là avec VxRail. Les clients souhaitant adopter une vision de cloud hybride, exploiter les conteneurs pour les applications modernes ou déployer un véritable centre de données défini par logiciel (SDDC) peuvent déployer VMware Cloud Foundation (VCF) sur Dell EMC VxRail. VxRail est le premier système d'infrastructure hyperconvergée entièrement intégré à VMware Cloud Foundation (VCF) SDDC Manager [2] , offrant un ensemble de composants logiciels VMware entièrement intégrés, notamment vSphere, vRealize, NSX, vSAN et SDDC Manager.
L'intégration de VCF à VxRail offre aux clients une plate-forme unifiée qui constitue une expérience complète et automatisée sur l'ensemble de la pile matérielle et logicielle. Grâce à cette intégration étroite, les clients bénéficieront d'un déploiement fluide et rapide et d'une expérience de gestion simplifiée tout en bénéficiant de l'agilité de l'infrastructure qui peut accélérer la capacité de leur organisation à fournir des applications. De plus, en raison de l'intégration profonde entre le matériel et les logiciels, VxRail offre également des avantages opérationnels cruciaux en matière de gestion du cycle de vie.
Bien que la mise en ligne rapide d'un cluster VxRail présente un avantage immédiat en termes d'impact commercial, les avantages opérationnels continus offrent les résultats les plus impressionnants. Celles-ci vont de l'évidence, comme ne plus rechercher les derniers pilotes pris en charge pour des éléments tels que les cartes réseau, les SSD et d'autres composants installés, à la recherche de correctifs/mises à jour logiciels de VMware. Mais il y a aussi le fait que Dell Technologies inclut de nouvelles fonctionnalités VMware dans les 30 jours à VxRail et Dell Technologies sert de point de contact unique pour tous les problèmes de support.
Apporter rapidement de nouvelles fonctionnalités aux clients est également un avantage considérable. Par exemple, à la mi-2020, VMware a publié plusieurs mises à jour autour de son logiciel Tanzu qui permet aux clients d'exécuter Kubernetes à partir d'un seul plan de contrôle. Pour les clients VxRail qui adoptent la fourniture d'applications modernes, VCF sur VxRail propose un processus clé en main pour mettre Tanzu en ligne. Sur le plan opérationnel, cela offre aux clients un moyen cohérent de déployer et de gérer des machines virtuelles traditionnelles aux côtés de conteneurs.
Compte tenu de l'étendue des technologies VMware SDDC et de leur adoption dans toute l'entreprise, les clients disposent de deux options distinctes en matière de déploiement. En tant que tel, nous avons cherché à comparer les avantages du VCF sur VxRail par rapport à l'alternative du "Do-It-Yourself". Nous avons commencé par déployer VCF sur VxRail, en suivant le processus de déploiement des hôtes, du générateur de cloud et, finalement, des mises à jour du cycle de vie du matériel et des logiciels.
À la fin, nous avons réaffecté exactement le même matériel aux serveurs PowerEdge vanille, en installant les composants individuels comme le ferait une organisation si elle déployait des composants vSphere, vSAN, NSX et vRealize. Pour la phase de gestion du cycle de vie, nous avons effectué manuellement la mise à niveau de VVD 5.1.1 vers 5.1.2.
Le tableau ci-dessous met en évidence ces trois segments définissables, mais il convient de noter que les actions du cycle de vie seront une tâche perpétuelle dans laquelle les organisations s'engageront régulièrement.
Bien que nous constations des avantages de déploiement immédiats (ce qui signifie que les clients seront en ligne et livreront plus rapidement avec VxRail), les avantages de la gestion continue du cycle de vie feront gagner du temps et permettront à l'entreprise de concentrer ses efforts ailleurs tout au long du cycle de vie du cluster.
Bien qu'il s'agisse de chiffres de comparaison de haut niveau, le rapport suivant décrit ces résultats en détail avec les processus techniques et le temps nécessaire pour terminer chaque étape. Bien que les résultats cumulés en accéléré racontent l'histoire, le détail du processus et la liste des tâches qui l'accompagne clarifient les différences entre l'achat d'une appliance conçue dans VxRail auprès de Dell EMC et les serveurs bare metal x86 standard.
Enfin, la comptabilisation des avantages de VxRail néglige de prendre en compte le délai entre la commande et la livraison. Dell Technologies propose plusieurs configurations fixes de l'usine qui peuvent être sur site et entièrement déployées en deux semaines.
Présentation technique
L'objectif principal était de quantifier la valeur apportée par la solution clé en main conçue conjointement avec VCF sur VxRail par rapport à la construction pièce par pièce dans l'approche DIY à l'aide de Vmware Validated Design. Les tests ont été divisés en trois parties : construction du domaine de gestion, du domaine de charge de travail et du LCM. En procédant ainsi, nous avons pu rassembler un temps comparable pour compléter les résultats sur les deux méthodes de résolution.
À des fins de test, nous avons utilisé un cluster à huit nœuds ; quatre nœuds ont été configurés pour le domaine de gestion et quatre nœuds pour le domaine de charge de travail.
Serveurs – 8 Dell EMC PowerEdge R640
- Mémoire - 576 Go de RAM
- Mise en réseau – 2 x Mellanox25GbE 2P ConnectX4LX
- Stockage – 4 disques SSD de 3.84 To Capacité totale 15.4 To
- Contrôleur de stockage – Dell EMC HBA330 Mini
Tous les tests ont été effectués en commençant par tous les hôtes dans un état hors tension. Les documentations d'installation et de configuration de Dell Technologies et de VMware ont été suivies.
Versions du logiciel VMware
