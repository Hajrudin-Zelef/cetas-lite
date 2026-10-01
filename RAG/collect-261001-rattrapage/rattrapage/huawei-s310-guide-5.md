---
id: collect-261001-rattrapage/rattrapage/huawei-s310-guide-5
title: "Guide ultra-complet — Huawei eKit S310"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["voice"]
source: docs/RAG/collect-261001-rattrapage/huawei_s310_guide.md
source_anchor: ""
source_lines: [677, 925]
sha256: 9f9b75d8fc90f48b1d771fc2e5abe89b9ef7cc4d6d69ea8ca859fd9a246eb445
---

# Guide ultra-complet — Huawei eKit S310

## 27. Comprendre les vues et la sauvegarde (VRP en 2 minutes)

| Commande | Effet |
|---|---|
| `system-view` | Passe en mode configuration |
| `quit` | Remonte d'un niveau (ou quitte le mode config) |
| `return` | Remonte directement en vue utilisateur `<...>` |
| `display current-configuration` | Config **active** (en RAM) |
| `display saved-configuration` | Config **sauvegardée** (en flash) |
| `save` | Copie la config active vers la flash |
| `display this` | Config de la vue courante (pratique dans une interface) |
| `?` | Aide contextuelle — ton meilleur ami |
| `Tab` | Complétion |

⚠️ **Piège n°1 des débutants Huawei :** tout ce que tu configures est **perdu au
reboot** tant que tu n'as pas fait `save`. Prends le réflexe : **chaque session de
config se termine par `save`**.

⚠️ **Piège n°2 :** `display current-configuration` peut être très longue.
Filtre avec `| include` : `display current-configuration | include vlan`.

---

## 28. Les ports du S310 dans la CLI : numérotation

Les interfaces se nomment `GigabitEthernet0/0/X` (abrégé `GE0/0/X`) pour les ports
cuivre, et `XGigabitEthernet0/0/X` pour les SFP+ 10G.

- `0/0/1` à `0/0/24` : ports RJ45 (sur un 24 ports).
- `0/0/25` à `0/0/28` : cages SFP/SFP+ (uplinks).
- Sur un 48 ports : `0/0/1` à `0/0/48`, SFP en `0/0/49` à `0/0/52`
  (**à vérifier sur le modèle exact** via `display interface brief`).

**Configurer plusieurs ports d'un coup (gain de temps énorme) :**

```
system-view
port-group pg-utilisateurs
 group-member GigabitEthernet0/0/1 to GigabitEthernet0/0/20
 port link-type access
 port default vlan 10
quit
save
```

🔧 `display interface brief` : la commande à taper en premier devant **tout**
problème de connectivité. Elle montre l'état physique (up/down) et protocole de
chaque port en un écran.

---

## 29. Mise à jour immédiate ? La question à se poser au déballage

Avant de configurer 50 VLAN, regarde la version logicielle :

```
display version
```

- Si le switch sort d'usine avec une version ancienne et que le site a Internet,
  le **Smart Upgrade** via la plateforme HOUP (Huawei Online Upgrade Platform)
  propose le chemin de mise à jour standardisé, en un clic (valeur datasheet).
- ✅ **Fais la mise à jour AVANT la configuration** sur un switch neuf : tu
  configures une fois, sur la version cible, sans risque de changement de
  comportement entre deux versions.
- ⚠️ Sur un switch **en production**, jamais de mise à jour sans sauvegarde
  (section 99) et sans plan de rollback (section 103).

---

## 30. VLAN : rappels 802.1Q (le minimum vital)

- Un VLAN = un **domaine de broadcast** séparé. Sans routage inter-VLAN, deux
  VLAN ne se parlent pas. C'est la base de la segmentation.
- Une trame **taggée** porte un en-tête 802.1Q avec l'ID du VLAN (1–4094).
  Une trame **non taggée** appartient au VLAN natif/PVID du port.
- **PVID** (Port VLAN ID) : le VLAN attribué aux trames non taggées qui entrent
  par le port.
- Sur le S310 : **4 094 VLAN** supportés, table MAC **16 000 entrées**
  (valeurs catalogue — à vérifier sur la fiche du modèle exact).

**Plan de VLAN type PME (exemple, à adapter) :**

| VLAN | Nom | Usage | Réseau exemple |
|---|---|---|---|
| 1 | default | Ne pas utiliser en production | — |
| 10 | USERS | Postes bureautiques | 192.168.10.0/24 |
| 20 | VOICE | Téléphonie IP | 192.168.20.0/24 |
| 30 | GUEST | Wi-Fi invités | 192.168.30.0/24 |
| 40 | CAMERA | Vidéosurveillance | 192.168.40.0/24 |
| 50 | SERVERS | Serveurs / NAS | 192.168.50.0/24 |
| 99 | MGMT | Management | 192.168.99.0/24 |

✅ Documente ce plan **par écrit** et affiche-le dans la baie. Un VLAN créé
« de tête » et oublié, c'est une panne dans 6 mois.

---

## 31. Créer des VLAN : la méthode propre

```
system-view
vlan batch 10 20 30 40 50 99
vlan 10
 description VLAN_USERS_BUREAUTIQUE
quit
vlan 20
 description VLAN_VOICE_TELEPHONIE
quit
vlan 99
 description VLAN_MANAGEMENT
quit
save
```

🔧 Vérification :

```
display vlan
display vlan 10
```

[web] *Fonctions L2 → VLAN → créer*, avec description. Le `vlan batch` n'a pas
toujours d'équivalent web : la CLI reste plus rapide pour les séries.

---

## 32. Ports access : la configuration de tous les jours

Un port **access** appartient à **un seul** VLAN, en non taggé. C'est le cas standard :
PC, imprimante, caméra, AP en mode simple.

```
system-view
interface GigabitEthernet0/0/5
 port link-type access
 port default vlan 10
 description PC_COMPTA_03
 undo shutdown
quit
save
```

**En série avec un port-group (20 postes d'un coup) :**

```
system-view
port-group pg-users
 group-member GigabitEthernet0/0/1 to GigabitEthernet0/0/20
 port link-type access
 port default vlan 10
quit
save
```

🔧 Vérification :

```
display port vlan GigabitEthernet0/0/5
```

⚠️ `undo shutdown` : sur Huawei, les ports sont **up par défaut**, mais après
certaines manips (ou un `shutdown` de test oublié), un port peut rester
administrativement down. Devant un port « mort », `display interface` dit si
c'est physique ou administratif.

---

## 33. Ports trunk : faire passer plusieurs VLAN

Un port **trunk** transporte **plusieurs VLAN taggés** (+ éventuellement un VLAN
natif non taggé). Usage : uplinks vers le cœur, vers un autre switch, vers un
pare-feu/routeur, vers un AP multi-SSID.

```
system-view
interface GigabitEthernet0/0/28        # cage SFP, uplink vers le cœur
 port link-type trunk
 port trunk allow-pass vlan 10 20 30 40 50 99
 port trunk pvid vlan 99              # VLAN natif (non taggé) = management
 description UPLINK_VERS_SW_CORE
 undo shutdown
quit
save
```

🔧 Vérifications :

```
display port vlan GigabitEthernet0/0/28
display trunkmembership
```

⚠️ **Erreur classique :** oublier un VLAN dans `allow-pass` → le VLAN « ne passe
pas » alors que tout semble configuré (voir dépannage section 121). **Des deux
côtés du lien**, les listes doivent correspondre.

⚠️ Le **PVID du trunk** (VLAN natif) doit être **identique des deux côtés**,
sinon : fuite de broadcast d'un VLAN vers l'autre et comportements bizarres.
En pratique : mets le VLAN de management en natif des deux côtés, ou `vlan 1`
si tu assumes (déconseillé, section 20).

---

## 34. Ports hybrid : la spécificité Huawei (à maîtriser absolument)

C'est **LA** différence avec Cisco/HP : le port **hybrid** peut envoyer certains
VLAN **taggés** et d'autres **non taggés**, au choix, port par port. C'est le mode
le plus souple — et celui qu'il faut comprendre pour la **voice VLAN** et les
cas mixtes (téléphone + PC).

```
system-view
interface GigabitEthernet0/0/8
 port link-type hybrid
 port hybrid pvid vlan 10                    # VLAN non taggé par défaut = users
 port hybrid tagged vlan 20                  # VLAN 20 (voix) sort taggé
 port hybrid untagged vlan 10                # VLAN 10 sort non taggé
 description TEL_IP_+_PC_BUREAU_08
 undo shutdown
quit
save
```

**Logique de fonctionnement :**

| Trame entrante | Traitement |
|---|---|
| Non taggée | Marquée avec le **PVID** du port |
| Taggée VLAN 20 | Acceptée (VLAN autorisé en tagged) |
| Taggée VLAN 30 | **Refusée** (non autorisée sur ce port) |

| Trame sortante (VLAN 10) | Non taggée (untagged) |
|---|---|
| Trame sortante (VLAN 20) | Taggée (tagged) |

✅ **Quand utiliser hybrid plutôt que trunk ?**
- Téléphone IP + PC branché derrière le téléphone (section 38).
- Équipement qui attend un VLAN en non taggé **et** un autre en taggé.
- Interconnexion avec du matériel hétérogène aux attentes exotiques.

⚠️ **Ne mélange pas les philosophies sans raison :** si tout ton parc est en
access/trunk classique, n'introduis pas du hybrid « pour voir ». Chaque
exception de config est une heure de dépannage en plus dans 2 ans.

---

## 35. Tableau récapitulatif : access / trunk / hybrid

