---
id: collect-261001-rattrapage/rattrapage/huawei-s310-guide-6
title: "Guide ultra-complet — Huawei eKit S310"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["voice"]
source: docs/RAG/collect-261001-rattrapage/huawei_s310_guide.md
source_anchor: ""
source_lines: [926, 1195]
sha256: a1076287a12c33a6948a9d28771bd6fed9055d5eff1399633ba90f633f348f30
---

# Guide ultra-complet — Huawei eKit S310

|  | Access | Trunk | Hybrid |
|---|---|---|---|
| Nombre de VLAN | 1 | Plusieurs | Plusieurs |
| Trames entrantes non taggées | → PVID (le VLAN du port) | → PVID du trunk | → PVID du port |
| Trames sortantes | Toujours non taggées | Taggées (sauf PVID) | Au choix par VLAN (tagged/untagged) |
| Usage typique | PC, imprimante, caméra | Uplink, AP multi-SSID | Téléphone+PC, cas mixtes |
| Commande clé | `port default vlan X` | `port trunk allow-pass vlan ...` | `port hybrid tagged/untagged vlan ...` |

---

## 36. Interface VLANIF et routage inter-VLAN (le « L2+ » du S310)

Le S310 est **L2+** : il sait router entre VLAN avec des **routes statiques**.
Pour que le VLAN 10 parle au VLAN 50, il faut une interface IP dans chaque VLAN
(sur le switch qui fait office de passerelle, souvent le switch cœur).

```
system-view
interface Vlanif10
 ip address 192.168.10.1 255.255.255.0
quit
interface Vlanif50
 ip address 192.168.50.1 255.255.255.0
quit
# Les postes utilisent .1 comme passerelle → le switch route entre VLAN
quit
save
```

Et vers Internet / le routeur principal :

```
system-view
ip route-static 0.0.0.0 0.0.0.0 192.168.50.254   # route par défaut vers le pare-feu
quit
save
```

🔧 Vérifications :

```
display ip interface brief
display ip routing-table
ping 192.168.50.10 -a 192.168.10.1   # test avec source précisée
```

⚠️ **Qui route ?** Un seul équipement doit être la passerelle d'un VLAN donné.
Deux passerelles (le S310 **et** le pare-feu) sur le même VLAN = routage
asymétrique et pannes intermittentes. Décide et documente.

⚠️ Le routage inter-VLAN **sans ACL** ouvre tout entre les VLAN : le VLAN
invités peut joindre les serveurs. Ajoute des **ACL** (section 86) dès que le
routage inter-VLAN est actif.

---

## 37. Exemple complet commenté : PME 3 VLAN (bureautique, invités, serveurs)

**Contexte :** un S310-24P4S, 20 PC (VLAN 10), 2 AP Wi-Fi invités (VLAN 30),
1 NAS (VLAN 50), uplink SFP vers le pare-feu en trunk.

```
system-view
sysname SW-ACC-01
#
vlan batch 10 30 50 99
vlan 10
 description VLAN_USERS
quit
vlan 30
 description VLAN_GUEST
quit
vlan 50
 description VLAN_SERVERS
quit
vlan 99
 description VLAN_MGMT
quit
#
# --- Ports utilisateurs 1-20 : access VLAN 10 ---
port-group pg-users
 group-member GigabitEthernet0/0/1 to GigabitEthernet0/0/20
 port link-type access
 port default vlan 10
 stp edged-port enable        # ports terminaux : pas d'attente STP (section 48)
quit
#
# --- AP invités 21-22 : trunk, SSID invités taggé 30, management 99 natif ---
interface GigabitEthernet0/0/21
 port link-type trunk
 port trunk allow-pass vlan 30 99
 port trunk pvid vlan 99
 description AP_GUEST_ENTREE
 stp edged-port enable
quit
interface GigabitEthernet0/0/22
 port link-type trunk
 port trunk allow-pass vlan 30 99
 port trunk pvid vlan 99
 description AP_GUEST_ATELIER
 stp edged-port enable
quit
#
# --- NAS port 23 : access VLAN 50 ---
interface GigabitEthernet0/0/23
 port link-type access
 port default vlan 50
 description NAS_SYNOLOGY
 undo shutdown
quit
#
# --- Uplink SFP 28 vers pare-feu : trunk tous les VLAN ---
interface GigabitEthernet0/0/28
 port link-type trunk
 port trunk allow-pass vlan 10 30 50 99
 port trunk pvid vlan 99
 description UPLINK_FW_PFSENSE
 undo shutdown
quit
#
# --- Management ---
interface Vlanif99
 ip address 192.168.99.10 255.255.255.0
quit
ip route-static 0.0.0.0 0.0.0.0 192.168.99.1
quit
save
```

[web] Tout est faisable en web (*VLAN → Ports*), mais la CLI est 5× plus rapide
dès qu'il y a plus de 5 ports.

---

## 38. Voice VLAN : le téléphone IP sans se prendre la tête

**Principe :** le téléphone IP s'alimente en PoE, taggue sa voix dans le VLAN 20,
et laisse passer le PC branché derrière lui en non taggé dans le VLAN 10.
Le switch reconnaît le téléphone via son **OUI MAC** (préfixe constructeur).

```
system-view
vlan 20
 description VLAN_VOICE
quit
#
voice-vlan 20 enable
voice-vlan mac-address 001B-2A00-0000 mask ffff-ff00-0000 description YEALINK
#  ^ OUI du constructeur de téléphones (exemple fictif : remplace par le vrai OUI)
#
interface GigabitEthernet0/0/8
 port link-type hybrid
 port hybrid pvid vlan 10
 port hybrid tagged vlan 20
 port hybrid untagged vlan 10
 voice-vlan 20 enable
 description TEL_YEALINK_BUREAU08
 undo shutdown
quit
save
```

**Ce qui se passe :** le téléphone envoie ses trames taggées VLAN 20 → le switch
les accepte et les fait suivre ; le PC derrière envoie en non taggé → PVID 10.
**Zéro configuration sur le téléphone** (il doit juste être en mode VLAN auto /
LLDP-MED, réglage par défaut sur la plupart des Yealink/Grandstream — à vérifier
sur le modèle exact du téléphone).

🔧 Vérifications :

```
display voice-vlan status
display mac-address vlan 20
```

⚠️ **OUI inconnu :** si les téléphones n'apparaissent pas dans le VLAN 20,
vérifie l'OUI (`display mac-address` puis compare les 3 premiers octets) et
ajoute-le. C'est la cause n°1 des « la voix ne passe pas ».

---

## 39. Exemple : téléphone + PC en daisy-chain, avec QoS voix

On combine voice VLAN (section 38) et priorité voix (section 73) :

```
system-view
interface GigabitEthernet0/0/8
 port link-type hybrid
 port hybrid pvid vlan 10
 port hybrid tagged vlan 20
 port hybrid untagged vlan 10
 voice-vlan 20 enable
 trust dscp                       # fait confiance au marquage du téléphone
 description TEL+PC_BUREAU08
 undo shutdown
quit
save
```

✅ Le téléphone marque déjà ses paquets (DSCP EF en général) : avec `trust dscp`,
le switch les priorise sans autre config. Vérifie que le téléphone est bien
configuré pour marquer (souvent défaut usine, à vérifier sur le modèle exact).

---

## 40. MUX VLAN : isoler sans multiplier les VLAN

Le S310 supporte la fonction **MUX VLAN** (valeur datasheet) : un VLAN
**principal** + des VLAN **subordonnés** (groupes ou isolés). Les ports isolés
ne se parlent pas entre eux mais parlent au principal (ex. : imprimante
partagée, box Internet, serveur).

```
system-view
vlan 60
 mux-vlan
 subordinate separate 61
 subordinate group 62
quit
interface GigabitEthernet0/0/15
 port link-type access
 port default vlan 61          # port isolé : ne voit que le VLAN principal
quit
interface GigabitEthernet0/0/16
 port mux-vlan enable
quit
save
```

> La syntaxe exacte et la disponibilité dépendent de la **version logicielle** :
> **à vérifier sur le modèle exact** (`display version`, puis `?` dans la vue VLAN).
> Alternative universelle : **port isolation** (section 82), plus simple.

---

## 41. Interopérabilité : trunk vers un switch non-Huawei

Rien de magique : 802.1Q est un standard. Points de vigilance :

1. **VLAN natif identique** des deux côtés (PVID Huawei = native VLAN Cisco).
2. **Liste des VLAN autorisés identique** (`allow-pass` ↔ `switchport trunk allowed vlan`).
3. **STP compatible** : RSTP des deux côtés, ou MSTP avec même nom de région,
   même révision, même mapping VLAN→instance (section 52).
4. Côté Cisco, le trunk négocie parfois en DTP : force `switchport mode trunk`
   + `switchport nonegotiate` côté Cisco pour éviter les surprises.

🔧 Test : `ping` entre deux PC du même VLAN de part et d'autre du trunk,
puis `display mac-address vlan X` des deux côtés pour voir les MAC apprises.

---

## 42. Vérifications VLAN : la routine

```
display vlan                          # tous les VLAN + ports membres
display vlan 10                       # détail d'un VLAN
display port vlan                     # PVID et mode de tous les ports
display port vlan GigabitEthernet0/0/8
display mac-address vlan 10           # quelles MAC sont vues dans le VLAN
display mac-address interface GigabitEthernet0/0/8
```

✅ **Réflexe :** après chaque modification VLAN, `display port vlan` sur les
ports touchés. 80 % des pannes VLAN se voient là en 10 secondes.

---

## 43. STP/RSTP : pourquoi on ne peut pas s'en passer

