---
id: collect-261001-huawei/huawei/r-networking-comments-1sd2kvy-absence-de-la-route-par-d-faut-dans-la-table-de-2293fd10
title: "r-networking-comments-1sd2kvy-absence-de-la-route-par-d-faut-dans-la-table-de-2293fd10"
domain: huawei
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["distribution", "energy"]
source: docs/RAG/collect-261001-huawei/r-networking-comments-1sd2kvy-absence-de-la-route-par-d-faut-dans-la-table-de-2293fd10.md
source_anchor: ""
source_lines: [1, 48]
sha256: 987bf689e9aebaf2fb7c6bf78adae34b33098719a433d36ec2a01ff0b9d454b4
---

# r-networking-comments-1sd2kvy-absence-de-la-route-par-d-faut-dans-la-table-de-2293fd10

Absence de la route par défaut dans la table de routage OSPF du Pare-Feu (Huawei USG) 
        
        
        
    
    
    Bonjour,
J'ai un reseau constitué:
- 
      Au coeur un routeur 8000 et un pare-Feu USG6000
- 
      A la distribution un core switch 12800
- 
      A accès des switchs TOR et accès.
le routeur (ASBR & ABR), le pare-Feu et le core-switch son dans la meme zone OSPF.
les neighbors adjency sont établies et les communications entre les équipements de la Zone OSPF et de mes réseaux locaux sont oéprationelles.
mon soucis est le suivant:
Mon routeur génère et redistribut le LSA de type 5 au Pare-Feu et Switch Core et ce LSA type 5 est bien présent dans leur LSBD.
Dans la table routage général et OSPF du Switch Coeur, on voit bien la route par défaut provenant du routeur (champs Nexhop) active mais sur le pare-Feu, cette route par défaut est également bien présente dans la table de routage OSPF mais inactive. Au contraire, je vois plutôt (dans le RIB général du pare-feu), une route par défaut avec la mention UNR dans la colone protocole avec comme next-hop le routeur.
Après quelques analyse:
- 
      je n'ai que la security policy par défaut qui est activé
- 
      je n'ai pas de route par défaut statique défini sur le pare-feu
- 
      je n'ai pas de PBR défini sur le pare-Feu
- 
      Aucune ACL défini sur le Pare-Feu
Quelqu'un peux avoir une idée du pourquoi la route par défaut obtenu par OSPF est désactivé au détriment de cette route (UNR) par défaut présent dans la table de routage général du pare-Feu ??
Merci d'avance,
Section des commentaires
You'll have more success posting in english about this.
Quebec energy right there lol
Ouais, j'aurai peu être du faire ainsi.
UNR is a user network route, it has an admin distance of 60 and a LSA type 5 has an admin distance/preference of 150 on Huawei. The lower admin distance wins.
Check with display ip routing-table protocol unr and see where it's coming from and determine if you need to keep it or delete it.
Avec la commande (display ip routing-table protocol unr), le next hop est mon routeur et je n'ai aucune route static crée ou importér dans mon process OSPF depuis mon routeur.
il n'y a que la redistribution qui a été faite sur le routeur. cette route (UNR) est crée dans quel circonstance s'il vous plait ?
Check
display ip routing-table verbose 0.0.0.0
UNR comes from either a VPN/tunnel protocol configured on the router, a configured DHCP server, NAT server mapping or PPPoE. Given what you've described I'm heavily assuming that it's coming from a configured tunneling protocol but the command I provided should show you.
UNR meaning unreachable? Do you have a working route to the next hop?
un itinéraire pour le prochain vers tous les adresses (0.0.0.0/0), oui mais il est inactif car cette route par défaut (avec UNR) est prioritaire.
Where are you getting your default route from in your router?
From redistribution?
oui, il y a de la redistribution au niveau de mon routeur (c'est un ASBR).
Your router being an ASBR means it's redistributing from either a static route or another routing process.
Also, do you have "default information originate" enabled in ospf? It won't redistribute a default route into the domain without it.
