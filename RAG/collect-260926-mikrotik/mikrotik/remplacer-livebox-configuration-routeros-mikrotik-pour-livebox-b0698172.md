---
id: collect-260926-mikrotik/mikrotik/remplacer-livebox-configuration-routeros-mikrotik-pour-livebox-b0698172
title: "remplacer-livebox-configuration-routeros-mikrotik-pour-livebox-b0698172"
domain: mikrotik
role: reference
task: reference
actors: []
dates: []
keywords: ["arr"]
source: docs/RAG/lot-mikrotik/forum/routing-bgp-ospf/remplacer-livebox-configuration-routeros-mikrotik-pour-livebox-b0698172.md
source_anchor: ""
source_lines: [1, 26]
sha256: d745edd4b2cae32fc54ace957c58eb76f256194e85a89f6d00e927262967f06c
---

# remplacer-livebox-configuration-routeros-mikrotik-pour-livebox-b0698172

Je me répond à moi-même car j'ai avancé dans mon investigation. Apparemment mon routeur ajoute 2 octets NULL de plus à la fin du paquet que la LB lors de l'envois d'un paquet CHAP Response : il en envois 4 alors que la LB n'en envois que 2.

En gros ils envoient tous les deux : 

1 octet Code

1 octet ID

2 octet Longueur CHAP

1 octet Longueur Challenge

16 octets Challenge

(dans mon cas) 11 Octets Name (contenant le user CHAP donc "fti/xxxxxxx" )

Soit un total de 32 octets. Comme il est bien indiqué 32 dans mon entête CHAP, ca devrait s'arrêter là, le reste c'est des octets de bourrage.

La livebox termine par deux 0x00 et le Mihrotik par quatre 0x00.

Normalement ça ne devrait pas poser de soucis parce que la 

RFC 1994
 dit que les octets situés après le nombre précisé par l'en-tête length du paquet CHAP doivent être ignorés., mais bon Orange, respect des RFC toussa...

Est-ce que quelqu'un qui aurait un setup sur BSD/Linux/Whatever OS qui permet de dump les paquets aurait la gentillesse de faire un petit PCAP pour mater le résultat qu'il obtient lors de l'authentification CHAP?
