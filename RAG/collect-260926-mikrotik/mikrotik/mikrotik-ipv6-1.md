---
id: collect-260926-mikrotik/mikrotik/mikrotik-ipv6-1
title: "in/out-bridge-port matcher not possible when interface (vlan832-orange-internet) is not slave"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["ethernet"]
source: docs/RAG/lot-mikrotik/forum/misc/mikrotik-ipv6.md
source_anchor: ""
source_lines: [1, 34]
sha256: 1fecb38e447f757b939d2e48f87d21d6e665fe8619cc6495a8a911c9064c5b5a
---

# in/out-bridge-port matcher not possible when interface (vlan832-orange-internet) is not slave

Bonjour,


J'avais réussi à configurer IPv4 et IPv6 sur mon routeur Mikrotik RB5009UPr+S+IN avec un module optique GPON-ONU-34-20BI pour remplacer une Iivebox fibre Sosh en Juillet 2023.


En préparant une doc pour partager sur github, je me suis aperçu que ma configuration IPv6 ne fonctionne plus. ça a cessé de fonctionner entre Juillet 2023 et aujourd'hui (Novembre 2023).


Y a t il eu des modifications récentes pour les options DHCPv6 ou toute configuration IPv6 récemment entre juillet et novembre 2023 ?


En regardant sur wireshark (mirroring de l'interface SFP vers l'une des interfaces ethernet), je retrouve bien ma requête DHCPv6 avec mes options. Je trouve juste curieux de recevoir ça dans la réponse d'Orange à ma requête (je sais pas si c'est lié à mon problème IPv6):


```
DHCPv6
...
    Domain Search List
        Option: Domain Search List (24)
        Length: 34
        Domain name suffix search list
            List entry: PUT.access.orange-multimedia.net.
```

"PUT.access.orange-multimedia.net." Qu'est-ce que ça peut bien vouloir dire ? Est-ce normal ?


Sinon voici ma configuration Mikrotik:


- avec "0xXXXXXXXXXXXXXXXXXXXXXX..." comme valeur identique entre DHCPv4 option 90 et DHCPv6 option 11... (pour rappel l'IPv4 fonctionne, j'ai accès à internet)


