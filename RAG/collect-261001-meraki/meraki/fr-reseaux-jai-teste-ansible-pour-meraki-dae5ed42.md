---
id: collect-261001-meraki/meraki/fr-reseaux-jai-teste-ansible-pour-meraki-dae5ed42
title: "Installation d'Ansible"
domain: meraki
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-meraki/fr-reseaux-jai-teste-ansible-pour-meraki-dae5ed42.md
source_anchor: ""
source_lines: [1, 52]
sha256: 61d280f0e004d7e56fec53c20a595f56e7389a7643cc6ad2a179edf3d3aff9a8
---

# Installation d'Ansible

Meraki, la solution cloud-native Cisco, possède sa collection Ansible,
Pour une solution manageable en GUI via le Dashboard Meraki (“at scale”), on peut se demander l’intérêt de cette collection.
Partons d’un cas concret !
Prenons l’exemple d’un administrateur réseau, qui a besoin de mettre à jour les ports trunk de son parc, pour y ajouter des VLANs fraichement créés,
“Facile ! ” me direz vous, avec le Virtual Stacking de Meraki (à ne pas confondre avec StackWise Virtual sous IOS-XE), il suffit de sélectionner tous les ports que nous souhaitons mettre à jour, faire un edit, et rajouter la liste des VLANs dans la liste de VLANs autorisés :
C’est effectivement aussi simple que ça, si tout le LAN est en Meraki :
Oui mais voilà, même si notre administrateur réseau à choisi Meraki comme standard (convaincu de la simplicité et de la scalabilité de la solution), rendre son réseau homogène n’est pas toujours chose facile (récupération de matériels existants, rachats d’entreprises, décisions technologiques décentralisées, …)
Son environnement ressemblera plutôt à ça :
Même tâche, mais un déploiement beaucoup plus complexe…
Hors de question pour notre administrateur réseau de faire le tour des consoles de management de chacun des éditeurs, ou pire encore, d’aller mettre à jour la configuration des équipements l’un après l’autre, en CLI/GUI. L’opération serait trop laborieuse, risquée, coûteuse. C’est ici que la collection Ansible va prendre tout son sens.
Automatisons tout ça !
Installation
Nous avons besoin d’installer Ansible, mais pas seulement. Dans un contexte réseau, Ansible utilise souvent SSH pour intéragir avec les équipements (c’est le cas dans cet article). Néanmoins, Ansible peut aussi utiliser d’autres moyens de transports, comme les APIs REST. C’est comme ça qu’Ansible va piloter le dashboard Meraki.
Installons donc les éléments nécessaires pour utiliser les APIs REST Meraki via Ansible :
# Installation d'Ansible
xvalette@lab cisco-meraki-ansible % python3 -m pip install --user ansible
# Installation du SDK Meraki
xvalette@lab cisco-meraki-ansible % python3 -m pip install --user meraki
# Commande pour les utilisateurs MacOs
xvalette@lab cisco-meraki-ansible % export OBJC_DISABLE_INITIALIZE_FORK_SAFETY=YES
Créons un playbook
Nous souhaitons mettre à jour les VLANs autorisés sur les ports trunks,
En inspectant la documentation de la collection, on trouve le module devices_switch_ports pour répondre à notre besoin,
A partir de la documentation et des exemples, nous arrivons avec le playbook suivant :
- hosts: localhost
   vars:
      network_id: L_00000000000000 | à remplacer par votre id de site
      serial_number: XXXX-XXXX-XXXX | à remplacer par le SN de vos switches
      meraki_api_key: clé API | à remplacer par votre clé
   gather_facts: false
   tasks:
   - name: Update Allowed VLANs on trunk ports
      cisco.meraki.devices_switch_ports:
         meraki_api_key: "{{meraki_api_key}}"
         state: present
         serial: "{{serial_number}}"
         portId: 4
         allowedVlans: 42-44,100-200
Tout est prêt !
Nous pouvons lancer notre playbook. Pas besoin d’inventaire, puisque Ansible va piloter le dashboard Meraki, pas les équipements.
xvalette@lab cisco-meraki-ansible % ansible-playbook playbooks/updateAllowedVlanTrunk.yml
Pour conclure
Dans de précédents articles, nous avions utilisé :
- Ansible avec Catalyst Center
- Ansible avec des équipements IOS-XE (ou third-party via les collections respectives)
Aujourd’hui nous venons de voir comment utiliser Ansible avec Meraki,
Ainsi, la gestion d’un parc hétérogène pour notre administrateur réseau devient moins laborieuse, moins risquée, et moins coûteuse :
Références
- Collection Ansible pour Meraki : https://docs.ansible.com/ansible/latest/collections/cisco/meraki/index.html
- Exemples de playbook de la collection Ansible : https://github.com/meraki/dashboard-api-ansible/tree/main/playbooks
- Article de Florian GIRARD sur Ansible pour Catalyst Center : https://gblogs.cisco.com/fr/reseaux/accelerer-vos-operations-avec-ansible-et-cisco-dna-center/
- Article sur Ansible pour IOS-XE : https://gblogs.cisco.com/fr/reseaux/manager-vos-equipements-cisco-avec-ansible/
