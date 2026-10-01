---
id: collect-261001-general-networking/general-networking/fr-review-tracking-the-untracked-with-fibre-channel-vm-id-27f85376-2
title: "fr-review-tracking-the-untracked-with-fibre-channel-vm-id-27f85376"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["exploit"]
source: docs/RAG/collect-261001-general-networking/fr-review-tracking-the-untracked-with-fibre-channel-vm-id-27f85376.md
source_anchor: ""
source_lines: [20, 47]
sha256: f9a8eef185a2fb0044df765021aec6b1ec8855f30bed7f88112a702cd9893fac
---

# fr-review-tracking-the-untracked-with-fibre-channel-vm-id-27f85376

Marvell QLogic Fibre Channel excelle en termes de performances et de fonctionnalités pour les réseaux de stockage dans VMware. Il rationalise le déploiement de VM à l'aide de VM-ID et prend en charge plusieurs ports avec des configurations FC et FC-NVMe simultanées pour une flexibilité optimale.
Les ID de machine virtuelle fournissent les détails
Il est essentiel de comprendre comment l'intégration de VM-ID à une solution de gestion SAN donne aux administrateurs une vue plus granulaire de l'ensemble de l'infrastructure SAN. SANnav et VM-ID répondent à des objectifs différents au sein de la structure de gestion d'un réseau de stockage, se complétant mutuellement en fournissant une approche complète de la gestion du réseau de stockage dans un environnement virtualisé.
Un VM-ID différencie et identifie les machines virtuelles individuelles, permettant aux plates-formes de virtualisation d'allouer des ressources, de gérer le cycle de vie des machines virtuelles, de faciliter la mise en réseau et de s'intégrer aux outils de gestion. VM-ID permet aux administrateurs de suivre et de gérer les machines virtuelles, d'attribuer des configurations spécifiques et de surveiller les performances.
ID de machine virtuelle Marvell QLogic et SANnav
Il ne fait aucun doute que la virtualisation des serveurs a grandement profité à la plupart des organisations, mais elle a posé plusieurs défis à l'équipe d'infrastructure et aux propriétaires d'applications. Au départ, les propriétaires d'applications étaient sceptiques quant à la capacité d'une plate-forme virtualisée à répondre aux besoins de leurs applications, et il y avait une résistance persistante à l'abandon des serveurs autonomes au profit d'un environnement virtualisé.
Il est juste de dire que les développeurs d'applications n'ont pas eu d'impact sur la virtualisation des serveurs. Cependant, des plaintes persistaient quant au besoin de plus de visibilité sur les mesures réelles, principalement en ce qui concerne les E/S. Le manque de visibilité était dû au fait que l'hyperviseur (dans le cas du serveur VMware vSphere ESXi) faisait abstraction du disque physique vers des disques virtuels placés sur un magasin de données, avec toutes les E/S vers le magasin de données de toutes les machines virtuelles de l'hyperviseur. agrégat. Ainsi, même si les performances globales du sous-système d’E/S sur le serveur étaient visibles, le niveau granulaire de visibilité sur la machine virtuelle réelle et l’application était inconnu.
Pour gagner en visibilité sur ces flux de VM individuels, la structure FC fournit une balise d'identification d'application de machine virtuelle basée sur des normes, VM-ID, à chaque VM. Une fois l'ID d'application attribué à une VM, la VM et les HBA Marvell QLogic 32GFC et 64GFC sur l'hyperviseur utilisent l'ID de la VM pour baliser toutes les trames de cette VM.
Le VM-ID identifie l'instance de VM spécifique qui lance les E/S et toutes les E/S ultérieures destinées à la cible. La balise VM-ID ne peut être appliquée que si la matrice de stockage prend en charge VM-ID. Les informations contenues dans chaque ID de VM permettent à SANnav de les corréler avec des mesures de performances, permettant ainsi aux administrateurs de surveiller des VM individuelles, de suivre l'utilisation des ressources et d'identifier rapidement les goulots d'étranglement potentiels en matière de performances. Les informations fournies par VM-ID fournissent aux administrateurs les détails nécessaires pour identifier et dépanner une VM affectée et prendre des mesures pour résoudre rapidement le problème.
Travaillant de concert, SANnav utilise les informations VM-ID intégrées dans chaque paquet FC par le Marvell QLogic FC-HBA pour suivre efficacement les performances des machines virtuelles individuelles.
VMware ESXi dans l'entreprise
Les organisations utilisent VMware ESXi pour créer et gérer des machines virtuelles. ESXi est un hyperviseur nu qui permet aux organisations de consolider plusieurs machines virtuelles sur un seul serveur, offrant ainsi flexibilité, optimisation des ressources et gestion plus facile d'une infrastructure informatique.
Le déploiement d'ESXi dans l'entreprise présente de nombreux avantages, ainsi que certains inconvénients. Grâce à la possibilité d'exécuter plusieurs machines virtuelles sur un seul serveur physique, les organisations peuvent réduire les coûts matériels et optimiser l'utilisation du serveur. Cela permet d'économiser sur la consommation d'énergie, le refroidissement et l'espace requis. Cependant, il s’agit du facteur le plus important contribuant à la prolifération des machines virtuelles.
Avec ESXi, les administrateurs peuvent allouer des ressources informatiques, comme le processeur, la mémoire et le stockage, aux machines virtuelles en fonction des besoins. Cette flexibilité garantit une utilisation efficace des ressources et évite le surprovisionnement et la sous-utilisation des ressources du serveur. ESXi prend en charge des fonctionnalités telles que vSphere High Availability (HA) et vSphere Fault Tolerance (FT), offrant ainsi une disponibilité et une résilience accrues du serveur.
ESXi utilise VMware vCenter Server comme interface de gestion centralisée pour surveiller, provisionner et gérer les environnements virtualisés. Sans VM-ID, gérer une infrastructure virtualisée comme celle-ci serait difficile, voire impossible.
Propagation des machines virtuelles
ESXi peut sans aucun doute contribuer à la prolifération des machines virtuelles s'il n'est pas correctement géré. VMware répond à ces préoccupations en fournissant des outils et des bonnes pratiques aux administrateurs informatiques. Ces outils incluent la planification des ressources et des capacités, la gestion du cycle de vie et l'automatisation basée sur des politiques.
VMware ESXi et vCenter sont essentiels au déploiement de la virtualisation d'entreprise, permettant aux organisations de répondre aux exigences de consolidation des serveurs, d'optimisation des ressources, de haute disponibilité et de gestion. Cependant, VM-ID est fondamental pour identifier et différencier les machines virtuelles individuelles, permettant ainsi aux administrateurs de gérer et d'optimiser efficacement leur infrastructure virtualisée.
Méfiez-vous de l'effet de mixage d'E/S
L'effet de mélange d'E/S se produit dans les environnements virtualisés, y compris ESXi, lorsque les modèles d'entrée/sortie (E/S) de stockage deviennent aléatoires et moins prévisibles. Cela peut être dû au fonctionnement simultané de plusieurs machines virtuelles (VM) partageant le même hôte physique et accédant aux ressources de stockage.
Dans un environnement virtualisé, plusieurs machines virtuelles exécutées sur un seul hôte peuvent envoyer des requêtes d'E/S à l'infrastructure de stockage sous-jacente à des moments différents, avec différents niveaux d'intensité et de fréquence. Lorsque l'hyperviseur reçoit ces requêtes d'E/S, elles sont agrégées et sérialisées avant d'être envoyées au système de stockage. En conséquence, les modèles d'E/S générés par les machines virtuelles deviennent « mélangés » ou mélangés. Cela rend difficile et chronophage l'identification des voisins bruyants et des coupables qui provoquent des embouteillages ou des blocages de tête de ligne, ce qui entraîne souvent le non-respect des SLA.
L’atténuation des effets de ce phénomène peut être obtenue grâce à diverses techniques telles que :
- Implémentation de la hiérarchisation du stockage
- Utiliser les mécanismes de QoS
- Optimisation des E/S
- Technologie VM-ID dans les HBA FC
Quel rôle joue le VM-ID ?
Bien que le VM-ID n’affecte pas directement l’effet du mélangeur d’E/S, il atténue considérablement son impact. Le VM-ID peut être exploité en implémentant les éléments suivants :
