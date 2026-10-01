---
id: collect-261001-fortinet/fortinet/pare-feu-fortigate-comment-configurer-une-interface-vlan-58463c2f-2
title: "Création de l'interface VLAN et configurations de base"
domain: fortinet
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-fortinet/pare-feu-fortigate-comment-configurer-une-interface-vlan-58463c2f.md
source_anchor: ""
source_lines: [52, 86]
sha256: 5d3d025ca5cb78fec73818976e1689996a908b2ef4584b0340e6dfed91f5da2f
---

# Création de l'interface VLAN et configurations de base

Meister-Gate (interface) # edit "Mon_VLAN_1"
new entry 'Mon_VLAN_1' added
Meister-Gate (Mon_VLAN_1) # set type vlan 
Meister-Gate (Mon_VLAN_1) # set alias "Mon_VLAN"
Meister-Gate (Mon_VLAN_1) # set vdom root 
Meister-Gate (Mon_VLAN_1) # set vlanid 200
Meister-Gate (Mon_VLAN_1) # set role lan
Meister-Gate (Mon_VLAN_1) # set ip 192.168.200.10 255.255.255.0
Meister-Gate (Mon_VLAN_1) # set allowaccess https http ping ssh 
Meister-Gate (Mon_VLAN_1) # set vrf 14
Meister-Gate (Mon_VLAN_1) # set vlan-protocol 8021q 
Meister-Gate (Mon_VLAN_1) # set device-identification enable
Meister-Gate (Mon_VLAN_1) # set interface port1 
Meister-Gate (Mon_VLAN_1) # end
# Configuration d'un serveur DHCP sur la nouvelle interface VLAN
Meister-Gate # config system dhcp server
Meister-Gate (server) # edit 1
Meister-Gate (1) # set dns-service default
Meister-Gate (1) # set default-gateway 192.168.200.10
Meister-Gate (1) # set netmask 255.255.255.0
Meister-Gate (1) # set interface Mon_VLAN_1
Meister-Gate (1) # config ip-range
Meister-Gate (ip-range) # edit 1
Meister-Gate (1) # set start-ip 192.168.200.11
Meister-Gate (1) # set end-ip 192.168.200.254
Meister-Gate (1) # next
Meister-Gate (ip-range) # end
Meister-Gate (1) # set timezone-option default
Meister-Gate (1) # end
En images :
Les commandes show system interface Mon_VLAN_1 et show system dhcp server 1 permettent de vérifier nos configurations.
IV. Conclusion
Vous savez désormais configurer des interfaces VLAN sur un pare-feu FortiGate. Il reste plus qu'à relier l'équipement à un ou plusieurs switchs et/ou configurer ces derniers, pour que vos différents terminaux puissent s'y connecter.
Il est par ailleurs possible d'établir une communication inter-VLAN ou d'accéder à Internet depuis un VLAN précis. Cela requiert la configuration de routes statiques et de règles de pare-feu, des procédures qui seront détaillées dans un prochain tutoriel.
Voilà qui marque la fin du présent tutoriel. N'hésitez pas à tester cela chez vous et nous faire un retour en commentaire.
