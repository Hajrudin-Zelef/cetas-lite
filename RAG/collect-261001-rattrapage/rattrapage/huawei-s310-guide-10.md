---
id: collect-261001-rattrapage/rattrapage/huawei-s310-guide-10
title: "Guide ultra-complet — Huawei eKit S310"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["incident"]
source: docs/RAG/collect-261001-rattrapage/huawei_s310_guide.md
source_anchor: ""
source_lines: [1930, 2159]
sha256: 60265039ec0737871be63d939cdc0422bc2aff79b778d0ffb1a26d28ade87310
---

# Guide ultra-complet — Huawei eKit S310

La visio est plus gourmande que la voix et moins bien marquée (souvent en
best effort par défaut). Deux options :

1. **Marquer côté poste** (stratégie Windows/GPO, à vérifier) puis `trust dscp`.
2. **Classifier par ports/applications** sur le switch (ex. plages UDP de Teams)
   — complexe et fragile (les plages changent).

✅ **Recommandation terrain :** pour la visio en PME, la vraie solution est
rarement la QoS fine : c'est **de la bande passante** (uplink 10G, section 5)
et un **VLAN visio séparé** si besoin. La QoS ne sauve pas un lien saturé à 100 %.

---

## 70. Vérifications QoS : la routine

```
display qos interface GigabitEthernet0/0/28     # files et compteurs
display traffic classifier user-defined
display traffic behavior user-defined
```

**Compteurs à surveiller :** les paquets **droppés** par file. Des drops en file
voix = la priorité ne suffit pas = il faut plus de bande passante, pas plus de QoS.

---

## 71. Panorama sécurité du S310 (tableau)

Fonctions annoncées au datasheet — disponibilité fine **à vérifier sur la version
logicielle du modèle exact** :

| Fonction | Rôle | Section |
|---|---|---|
| Port security + sticky MAC | Limite les MAC par port, mémorise | 72 |
| Storm control | Plafonne broadcast/multicast/unicast inconnu | 73 |
| Port isolation | Isole des ports entre eux (L2) | 74 |
| DHCP snooping | Filtre les serveurs DHCP pirates | 75 |
| Binding IP+MAC+port+VLAN | Table de correspondance anti-spoofing | 76 |
| DAI (Dynamic ARP Inspection) | Vérifie les ARP (via binding) | 77 |
| IPSG (IP Source Guard) | Filtre le trafic selon le binding | 77 |
| ACL | Filtrage L2/L3/L4 | 78 |
| 802.1X / MAC auth / Portal | Authentification d'accès | 79 |
| AAA / RADIUS / HWTACACS | Centralisation des comptes | 79 |
| Anti-attaques DoS/ARP/ICMP | Protection CPU et usurpation | 80 |
| Blackhole MAC | Bloque une MAC partout | 80 |
| Apprentissage ARP strict | Anti-épuisement de la table ARP | 80 |

---

## 72. Port security et sticky MAC : un port = N équipements max

**Objectif :** empêcher qu'on branche un hub/switch sauvage (ou qu'on usurpe une
MAC) sur une prise murale. Limite le nombre de MAC apprises par port.

```
system-view
interface GigabitEthernet0/0/5
 port-security enable
 port-security max-mac-num 2          # 2 MAC max (PC + téléphone en daisy-chain)
 port-security mac-address sticky     # mémorise les MAC vues (sticky)
 port-security protect-action shutdown # action si dépassement : shutdown
quit
save
```

**Actions possibles en cas de violation** (à vérifier sur la version exacte) :
`protect` (bloque les nouvelles MAC, log), `restrict` (bloque + alarme),
`shutdown` (coupe le port).

⚠️ Avec `sticky`, les MAC sont mémorisées : si tu **changes le PC** d'une prise,
pense à effacer (`undo port-security mac-address sticky` puis réactiver) sinon
le nouveau PC est bloqué. Documente la procédure pour le support niveau 1.

🔧 Vérifications :

```
display port-security
display mac-address sticky
```

---

## 73. Storm control : plafonner les tempêtes

Même avec STP, une boucle sur un équipement non-STP ou un équipement défaillant
peut inonder le réseau. Le storm control **plafonne** les broadcast, multicast
et unicast inconnu par port.

```
system-view
interface GigabitEthernet0/0/5
 storm-control broadcast min-rate 1000 max-rate 2000    # en pps (exemple)
 storm-control multicast min-rate 1000 max-rate 2000
 storm-control unicast min-rate 1000 max-rate 2000
 storm-control action shutdown                           # ou block
quit
save
```

En série sur tous les ports d'accès :

```
system-view
port-group pg-users
 group-member GigabitEthernet0/0/1 to GigabitEthernet0/0/20
 storm-control broadcast min-rate 1000 max-rate 2000
 storm-control action block
quit
save
```

> Les unités (pps vs kbit/s) et la syntaxe exacte sont **à vérifier sur la version
> logicielle du modèle exact**. Valeurs de départ raisonnables en bureautique :
> broadcast ≤ 1–2 % de la bande passante du port.

✅ **À activer sur TOUS les ports d'accès**, en complément de BPDU guard
(section 47). Les deux ensemble = quasi-immunité aux boucles « de bureau ».

---

## 74. Port isolation : isoler sans VLAN

**Objectif :** empêcher des ports de communiquer **entre eux** en L2, tout en
les laissant joindre l'uplink (ex. : chambres d'hôtel, box clients, bornes
publiques). Plus simple qu'un VLAN par client.

```
system-view
interface GigabitEthernet0/0/5
 port-isolate enable
quit
interface GigabitEthernet0/0/6
 port-isolate enable
quit
# L'uplink (port 28) N'est PAS isolé : les ports isolés le joignent normalement
save
```

Par défaut, tous les ports isolés rejoignent le même **groupe d'isolation** :
ils ne se voient plus entre eux, mais voient tous les ports non isolés.

🔧 Vérification : `display port-isolate group all`.

⚠️ L'isolation est **L2 uniquement** : si un routeur fait du inter-VLAN, deux
ports isolés peuvent se reparler via le routeur (aller-retour L3). Pour une
vraie étanchéité, combine avec des ACL (section 78).

---

## 75. DHCP snooping : traquer les serveurs DHCP pirates

**Le problème :** n'importe quelle box/routeur branché au réseau peut se mettre
à distribuer des adresses DHCP → pannes aléatoires « j'ai Internet puis plus ».

**Principe :** on déclare **trusted** uniquement les ports vers le vrai serveur
DHCP (ou le relais). Les réponses DHCP (OFFER/ACK) arrivant par un port
**untrusted** sont jetées.

```
system-view
dhcp enable
dhcp snooping enable
#
vlan 10
 dhcp snooping enable          # active le snooping dans le VLAN users
quit
#
interface GigabitEthernet0/0/28
 dhcp snooping trusted         # uplink vers le serveur DHCP / le cœur
 description UPLINK_VERS_DHCP
quit
# Les ports 1-20 restent untrusted par défaut : c'est ce qu'on veut
save
```

**Table de binding :** le switch mémorise (IP, MAC, port, VLAN, bail) de chaque
client légitime. C'est cette table qui alimente DAI et IPSG (sections 76–77).

🔧 Vérifications :

```
display dhcp snooping user-bind all
display dhcp snooping statistics
```

⚠️ **Cause n°1 de « plus de DHCP » après activation :** le port vers le vrai
serveur DHCP oublié en `trusted`. Et en cascade : **chaque** switch traversé
doit avoir son uplink en trusted. Teste l'obtention d'un bail **après**
activation, sur chaque VLAN.

⚠️ Si le serveur DHCP est **sur le switch lui-même** (DHCP server local),
adapte : le snooping et le serveur local cohabitent, mais vérifie le
comportement sur ta version (à vérifier sur le modèle exact).

---

## 76. Binding IP + MAC + port + VLAN : la table de vérité

Le datasheet annonce la fonction de **binding de l'adresse IP, de l'adresse MAC,
de l'ID de port et de l'ID de VLAN** (valeur datasheet). Deux sources :

1. **Dynamique** : via DHCP snooping (section 75) — le plus simple.
2. **Statique** : pour les équipements en IP fixe (imprimantes, serveurs).

```
system-view
user-bind static ip-address 192.168.10.50 mac-address 0011-2233-4455 interface GigabitEthernet0/0/23 vlan 10
#  ^ exemple fictif : adapte IP, MAC, port
quit
save
```

🔧 `display dhcp snooping user-bind all` montre les deux types.

✅ **Inventaire vivant :** cette table, c'est ton inventaire IP/MAC/prise à jour
automatique. Exporte-la régulièrement (`display dhcp snooping user-bind all`
dans un fichier) : en cas d'incident, tu sais **qui était où**.

---

## 77. DAI et IPSG : anti-ARP-spoofing et anti-usurpation IP

**DAI (Dynamic ARP Inspection)** : vérifie chaque paquet ARP reçu contre la
table de binding. Un ARP qui ment (IP/MAC incohérents) est jeté. Fini les
« attaques » ARP qui coupent le réseau ou interceptent le trafic.

**IPSG (IP Source Guard)** : filtre **tout le trafic IP** d'un port selon la
table de binding. Un poste qui change d'IP manuellement pour « passer outre »
est bloqué.

