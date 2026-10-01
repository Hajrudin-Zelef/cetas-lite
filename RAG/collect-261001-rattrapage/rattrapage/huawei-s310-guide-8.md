---
id: collect-261001-rattrapage/rattrapage/huawei-s310-guide-8
title: "Guide ultra-complet — Huawei eKit S310"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/huawei_s310_guide.md
source_anchor: ""
source_lines: [1449, 1704]
sha256: d9366de21a709e797f0413877ce864e686d1c555bd55f62855494c4b9fb92527
---

# Guide ultra-complet — Huawei eKit S310

```
system-view
stp mode mstp
stp enable
stp region-configuration
 region-name PME_NORD
 revision-level 1
 instance 1 vlan 10 50
 instance 2 vlan 20 30
 active region-configuration
quit
stp instance 1 root primary
stp instance 2 root secondary
quit
save
```

⚠️ **Les 3 paramètres de région doivent être IDENTIQUES sur tous les switches :**
nom, révision, mapping VLAN→instance. Un seul qui diffère = les switches ne se
reconnaissent pas comme même région = comportements erratiques. `display stp
region-configuration` pour comparer.

✅ En PME simple : **reste en RSTP**. MSTP ne se justifie qu'avec un vrai besoin
de répartition de charge par VLAN.

---

## 53. Agrégation de liens (Eth-Trunk) : le principe

Un **Eth-Trunk** regroupe plusieurs liens physiques en **un seul lien logique** :
- **Plus de débit** (2× 1G = 2 Gbit/s agrégés).
- **Redondance** : si un brin tombe, le trafic bascule sur l'autre sans
  reconvergence STP (le trunk est vu comme un seul port).
- La répartition se fait par **hachage** (MAC/IP source/destination) : **une
  seule conversation** ne dépassera jamais la vitesse d'un brin. C'est normal,
  pas une panne.

**Deux modes :**

| Mode | Négociation | Usage |
|---|---|---|
| Manuel (`mode manual`) | Aucune | Équipements sans LACP, ou besoin déterministe |
| **LACP** (`mode lacp`) | Oui (802.3ad) | **Recommandé** : détection de panne, ajout/retrait dynamique |

---

## 54. Eth-Trunk manuel : configuration

**Exemple :** 2 liens (ports 25-26) vers le NAS en mode manuel.

```
system-view
interface Eth-Trunk1
 description TRUNK_VERS_NAS
 mode manual load-balance        # mode manuel (selon version : "mode manual")
quit
interface GigabitEthernet0/0/25
 eth-trunk 1
quit
interface GigabitEthernet0/0/26
 eth-trunk 1
quit
# Le trunk hérite ensuite de la config VLAN :
interface Eth-Trunk1
 port link-type trunk
 port trunk allow-pass vlan 10 50
quit
save
```

> La syntaxe exacte (`mode manual` vs `mode manual load-balance`) est **à vérifier
> sur la version logicielle du modèle exact** (`interface Eth-Trunk1` puis `mode ?`).

---

## 55. LACP : la configuration recommandée

**Exemple :** 2× 1G (ports 27-28) vers le switch cœur, en LACP actif.

**Sur SW-ACC-01 :**

```
system-view
interface Eth-Trunk10
 description LACP_VERS_CORE
 mode lacp
 port link-type trunk
 port trunk allow-pass vlan 10 20 30 50 99
quit
interface GigabitEthernet0/0/27
 eth-trunk 10
 lacp priority 100
quit
interface GigabitEthernet0/0/28
 eth-trunk 10
quit
save
```

**Sur le cœur :** config miroir (même numéro de trunk conseillé, pas obligatoire).

🔧 Vérifications :

```
display eth-trunk 10
```

État attendu : les deux membres en **Selected** (actifs). Un membre en
**Unselected** = problème de négociation (vérifie le mode LACP des deux côtés,
la vitesse/duplex, le câble).

⚠️ **LACP des deux côtés** : un côté en LACP et l'autre en manuel = le trunk ne
monte pas (ou monte à moitié). Toujours vérifier le mode en face.

---

## 56. Eth-Trunk : bonnes pratiques et limites

1. **Même vitesse, même duplex** sur tous les membres (négociation auto des
   deux côtés, ou forcé identique des deux côtés — jamais de mélange).
2. **Même type de port** : que du cuivre ou que de la fibre, pas de mix.
3. Maximum de membres : **à vérifier sur la fiche du modèle exact**
   (souvent 8 sur cette gamme).
4. Le trunk se comporte comme **un seul port** pour STP et les VLAN : configure
   le VLAN sur l'interface Eth-Trunk, **pas** sur les membres.
5. Pour un trunk **inter-sites en fibre** : utilise des paires de fibres
   identiques (même longueur d'onde, même type).
6. ⚠️ Ne mets jamais les deux brins d'un trunk sur le **même chemin physique**
   (même goulotte écrasée par un transpalette = les deux brins coupés).

---

## 57. PoE : rappels 802.3af/at et classes de puissance

Le switch **détecte** l'équipement (signature résistive), **classifie** son besoin,
puis alimente. Pas de détection = pas de courant : brancher un PC non-PoE sur un
port PoE **ne risque rien**.

| Classe | Standard | Puissance max fournie (PSE) | Puissance utile (PD) | Exemples |
|---|---|---|---|---|
| Classe 1 | 802.3af | 4 W | 3,84 W | Téléphone IP basique |
| Classe 2 | 802.3af | 7 W | 6,49 W | Téléphone IP écran couleur |
| Classe 3 | 802.3af | 15,4 W | 12,95 W | AP Wi-Fi 6 d'intérieur, caméra fixe |
| Classe 4 | 802.3at (PoE+) | 30 W | 25,5 W | AP extérieur, caméra PTZ, visio |

Le S310-P est **PoE+** (802.3at) : **30 W max par port**, dans la limite du
**budget total** (400 W sur 24 ports, 380 W sur 48 ports).

⚠️ La puissance « fournie » inclut les pertes dans le câble : un équipement de
12,95 W consomme ~15,4 W côté switch. **Toujours calculer côté switch.**

---

## 58. Budget PoE par modèle (valeurs constructeur)

| Modèle | Budget total | Max par port | Conso max pleine charge |
|---|---|---|---|
| S310-24P4S | **400 W** | 30 W | 491,66 W |
| S310-24P4X | **400 W** | 30 W | à vérifier sur la fiche du modèle exact |
| S310-24PN4X | **400 W** | 30 W | à vérifier sur la fiche du modèle exact |
| S310-48P4S | **380 W** | 30 W | 462,8 W |

🔧 Commandes :

```
display poe power-state          # budget total, consommé, disponible
display poe interface all        # état PoE port par port
```

---

## 59. Fast PoE et Perpetual PoE : ce que ça change au quotidien

Deux fonctions constructeur (valeur datasheet) qui évitent des incidents :

- **Fast PoE** : après une coupure de courant, le switch réalimente les
  équipements PoE en **quelques secondes**, sans attendre la fin du boot
  (1 à 3 minutes sur un switch classique). Tes caméras et AP repartent
  presque immédiatement.
- **Perpetual PoE** : pendant un **redémarrage logiciel** (ex. mise à jour
  firmware), l'alimentation PoE **n'est pas coupée**. Les AP restent associés,
  les caméras continuent d'enregistrer.

✅ **Conséquence pratique :** une mise à jour firmware un soir de semaine devient
envisageable sans couper la vidéosurveillance. (Ça ne dispense pas de la faire
en heure creuse — section 102.)

---

## 60. Configurer le PoE par port

```
system-view
interface GigabitEthernet0/0/7
 poe enable                      # active le PoE (souvent actif par défaut)
 poe priority critical           # priorité : critical / high / low
 poe power-off time-range ...    # extinction planifiée (si supporté)
 description AP_ETAGE1
quit
save
```

**Forcer une réalimentation (le « débranche/rebranche » à distance) :**

```
system-view
interface GigabitEthernet0/0/7
 poe power-off
 poe power-on
quit
```

🔧 C'est la commande magique quand une caméra ou un AP est « planté » :
couper/remettre le PoE **à distance**, sans monter à l'échelle. À documenter
dans tes procédures de niveau 1.

**Limiter la puissance max d'un port (sécurité + budget) :**

```
system-view
interface GigabitEthernet0/0/7
 poe power 15000                 # limite à 15 W (en milliwatts selon version)
quit
save
```

> L'unité et la syntaxe exacte (`poe power`, `poe max-power`) sont **à vérifier
> sur la version logicielle du modèle exact**.

---

## 61. Priorités PoE : éviter que l'AP directeur tombe avant la caméra du parking

Quand le budget est dépassé, le switch coupe les ports **par priorité croissante**
(d'abord `low`, en dernier `critical`). **Configure les priorités dès le jour 1**,
pas le jour où ça coupe.

**Politique type :**

| Priorité | Équipements |
|---|---|
| `critical` | AP du bureau de direction ? Non — **téléphones de secours, AP du hall**... En vrai : ce qui est vital pour la sécurité |
| `high` | AP Wi-Fi intérieurs, téléphones IP |
| `low` | Caméras extérieures non critiques, AP invités |

```
system-view
interface GigabitEthernet0/0/7
 poe priority high
quit
interface GigabitEthernet0/0/12
 poe priority low
quit
save
```

