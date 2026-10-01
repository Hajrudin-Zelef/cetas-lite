---
id: collect-261001-cisco/cisco/c-fr-ca-support-docs-smb-switches-cisco-business-switching-kmgmt-2237-vlan-confi-77bf9959
title: "c-fr-ca-support-docs-smb-switches-cisco-business-switching-kmgmt-2237-vlan-confi-77bf9959"
domain: cisco
role: reference
task: reference
actors: []
dates: []
keywords: ["diffusion", "ethernet"]
source: docs/RAG/collect-261001-cisco/c-fr-ca-support-docs-smb-switches-cisco-business-switching-kmgmt-2237-vlan-confi-77bf9959.md
source_anchor: ""
source_lines: [1, 45]
sha256: 09b9b591f3773aeee2e183e189a505c53e29b27b209789972b33105fc070eb96
---

# c-fr-ca-support-docs-smb-switches-cisco-business-switching-kmgmt-2237-vlan-confi-77bf9959

Les VLAN vous permettent d’effectuer une segmentation logique d’un réseau local en différents domaines de diffusion. Dans les scénarios où des données sensibles peuvent être diffusées sur un réseau, des VLAN peuvent être créés pour améliorer la sécurité en dirigeant une diffusion sur un VLAN spécifique. Seuls les utilisateurs appartenant à un VLAN peuvent accéder aux données sur ce VLAN et les manipuler. Les VLAN peuvent également être utilisés pour améliorer les performances en réduisant le besoin d’acheminer des diffusions et des multidiffusions vers des destinations inutiles.
L'objectif de ce document est de vous montrer comment configurer un VLAN de base via l'interface de ligne de commande (CLI) sur les commutateurs gérés des gammes CBS250 et CBS350.
Étape 1. Connectez-vous à l’interface de ligne de commande (CLI) du commutateur.
Étape 1. Saisissez les commandes suivantes pour créer un réseau local virtuel :
| Commande | Objectif | 
| configuration | Passez en mode de configuration. | 
| vlan database | Passer en mode de base de données de VLAN. | 
| vlan <ID> | Créer un nouveau VLAN avec un ID spécifié. | 
| tranche | Quitter le mode de configuration. | 
La capture d’écran ci-dessous montre les étapes nécessaires pour créer un VLAN avec un ID de 200.
Étape 2. (Facultatif) Saisissez la commande suivante pour afficher les renseignements sur le VLAN :
| Commande | Objectif | 
| show vlan | Afficher les renseignements sur le VLAN. | 
Remarque : Le tableau de renseignements sur le VLAN varie en fonction du type de commutateur que vous utilisez. Par exemple, les commutateurs de type SF peuvent avoir un champ Type et Authorization (autorisation) plutôt qu’un champ Creators (créateurs). Le champ Ports varie également, car différents commutateurs comportent des types de ports et des schémas de numérotation différents.
Le VLAN créé s’affiche :
Remarque : Le VLAN 1 est le VLAN par défaut et comprend tous les ports possibles sur le commutateur par défaut. Les ports numérotés avec gi sont des ports Gigabit Ethernet (liaisons individuelles). Les ports numérotés avec fa sont Fast Ethernet (liaisons individuelles). Les ports numérotés avec Po sont des canaux de ports (un groupe de liaisons Ethernet).
Une fois les VLAN créés, vous devez attribuer les ports au VLAN approprié. Vous pouvez configurer les ports à l’aide de la commande switchport et spécifier si un port doit être en mode access (accès) ou trunk (liaisons).
Les modes des ports sont définis comme suit :
• Accès – Les cadres reçus sur l’interface sont supposées ne pas avoir de balise de VLAN et sont attribués au VLAN spécifié par la commande. Les ports d’accès sont principalement utilisés pour les hôtes et ne peuvent acheminer du trafic que pour un seul VLAN.
• Liaisons – Les cadres reçus sur l’interface sont supposées avoir des balises de VLAN. Les ports de liaisons sont destinés aux liaisons entre les commutateurs ou d’autres appareils réseau et peuvent transporter du trafic pour plusieurs VLAN.
Remarque : Par défaut, toutes les interfaces sont en mode de liaisons, ce qui signifie qu’elles peuvent acheminer le trafic pour tous les VLAN.
Étape 1. Saisissez les commandes suivantes pour configurer un port d’accès :
| Commande | Objectif | 
| conf t | Passez en mode de configuration. | 
| int <port number> | Passer en mode de configuration d’interface pour le numéro de port spécifié. Gigabit Ethernet, Fast Ethernet et les canaux de ports sont valides. | 
| switchport mode access | Définir l’interface en tant qu’interface Ethernet à VLAN unique non balisée. Un port d’accès peut acheminer le trafic dans un seul VLAN. | 
| switchport access vlan <ID> | Préciser le VLAN pour lequel ce port d’accès acheminera le trafic. | 
| no shut | Activer le port. | 
| tranche | Quitter le mode de configuration. | 
La capture d’écran suivante indique les étapes nécessaires pour configurer le port Gigabit Ethernet (gi2) en tant que port d’accès et l’attribuer au VLAN 200.
Étape 2. (Facultatif) Saisissez la commande show vlan pour voir le port que vous avez attribué.
Étape 3. Saisissez les commandes suivantes pour configurer un port de liaisons et spécifiez que seuls certains VLAN sont autorisés sur la solution de liaisons spécifiée :
| Commande | Objectif | 
| conf t | Passez en mode de configuration. | 
| int <port number> | Passer en mode de configuration d’interface pour le numéro de port spécifié. Gigabit Ethernet, Fast Ethernet et les canaux de ports sont valides. | 
| switchport mode trunk | Faire que le numéro de port spécifié soit informé de tous les réseaux locaux virtuels. | 
| switchport trunk allowed vlan add <ID> | Faire du port un membre dans l’ID de VLAN spécifié et lui attribue une règle de sortie : Tagged (balisé). Cela signifie que les paquets sont marqués avec l’ID de VLAN lorsqu’ils quittent ce port sur l’appareil. | 
| no shut | Activer le port. | 
| tranche | Quitter le mode de configuration. | 
Remarque : En mode de liaisons, tous les VLAN sont autorisés par défaut. L’utilisation de la commande switchport trunk allowed vlan add vous permet de configurer les VLAN autorisés sur la solution de liaisons.
La capture d’écran suivante montre les étapes nécessaires pour définir le port Gigabit Ethernet (gi3) comme port de liaisons et l’ajouter au VLAN 200 :
Étape 4. (Facultatif) Saisissez la commande show vlan pour afficher vos modifications.
Étape 5. (Facultatif) Saisissez la commande ci-dessous pour afficher les renseignements sur un port :
| Commande | Objectif | 
| show interfaces switchport <port number> | Afficher des renseignements comme l’appartenance au VLAN, la règle de sortie et les VLAN interdits pour le port spécifié. |
