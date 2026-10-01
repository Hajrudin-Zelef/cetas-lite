---
id: collect-261001-rattrapage/rattrapage/huawei-s310-guide-1
title: "Guide ultra-complet — Huawei eKit S310"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/huawei_s310_guide.md
source_anchor: ""
source_lines: [1, 144]
sha256: a21723fbf50306edb3e53142aaf3d739baa16490214f596ab64dd1afb7e47028
---

# Guide ultra-complet — Huawei eKit S310
## Switch manageable PME : installation, configuration, exploitation et dépannage

> **Public visé :** techniciens et responsables systèmes & réseaux en PME / multi-sites.
> **Ton :** terrain, concret, direct. Chaque section donne le pourquoi, le comment (CLI VRP + web),
> et le piège à éviter.
> **Périmètre :** gamme Huawei eKit S310 (eKitEngine S310) — L2+ manageable, gestion cloud eKit,
> web, CLI, SNMP. Les valeurs constructeur citées viennent de la fiche officielle de la série
> (datasheet eKitEngine S310). Quand une valeur dépend du modèle exact ou de la version logicielle,
> c'est écrit noir sur blanc : **« à vérifier sur la fiche du modèle exact »**.
>
> **Règle d'or de ce guide :** les mots de passe et clés figurant dans les exemples sont
> **fictifs** (ex. `Exemple_MotDePasse_2026!`). Ne jamais les utiliser en production.

**Sources constructeur consultées (recherche web, septembre 2026) :**
- Fiche produit officielle : `https://ekit.huawei.com/ekit/front/ssr/en/product/1239139538459211136` (S310-48P4S)
- Datasheet série eKitEngine S310 (PDF) : `https://i.pcdeacitec.com/datasheets/syscom/S310-24U4X/f2112277465062598bb47bdc547534d7.pdf`
- Datasheet miroir : `https://resource.ewe.rs/media/documents/2026/08/2026-08-07573726.pdf`
- Fiche S310-24T4S (datasheet PDF) : `https://shop.aodatacloud.es/medias/datasheet_s310-24t4s.pdf`

---

## 1. À qui s'adresse ce guide (et comment l'utiliser)

1. **Technicien de terrain** : tu déballes, tu montes en rack, tu câbles, tu fais la première
   configuration. Va directement aux sections 14 à 29, puis au pense-bête (section 119).
2. **Administrateur réseau PME** : VLAN, STP, PoE, QoS, sécurité. Cœur du guide : sections 30 à 95.
3. **Responsable / chef de service** : dimensionnement PoE (sections 66 à 70), maintenance
   préventive (sections 113 à 116), durcissement (sections 117 à 118), pièces de rechange.
4. **Conventions de lecture** :
   - Les blocs ` ``` ` sont des commandes CLI VRP Huawei à taper en mode système (`system-view`).
   - `[web]` signale l'équivalent dans l'interface web quand il existe.
   - ⚠️ = piège classique de terrain. ✅ = bonne pratique. 🔧 = commande de vérification.

---

## 2. La gamme eKit S310 en un coup d'œil

La série eKitEngine S310 est la gamme de switches manageables d'entrée/milieu de gamme Huawei
pour les PME, pilotable en cloud via l'app **Huawei eKit**, en web local, en CLI (VRP) ou en SNMP.
Tous les modèles sont **L2+** (commutation L2 + routage statique de base), 1U, châssis métal,
alimentation AC intégrée.

| Modèle | Ports cuivre | PoE | Uplinks | Capacité | Débit | PoE budget |
|---|---|---|---|---|---|---|
| S310-24T4S | 24× GE RJ45 | Non | 4× GE SFP | 56 Gbit/s | 42 Mpps | — |
| S310-24P4S | 24× GE RJ45 | PoE+ | 4× GE SFP | 56 Gbit/s | 42 Mpps | **400 W** |
| S310-24T4X | 24× GE RJ45 | Non | 4× 10GE SFP+ | 128 Gbit/s | 96 Mpps | — |
| S310-24P4X | 24× GE RJ45 | PoE+ | 4× 10GE SFP+ | 128 Gbit/s | 96 Mpps | **400 W** |
| S310-24PN4X | 24× 2,5GE RJ45 | PoE+ | 4× 10GE SFP+ | 200 Gbit/s | 144 Mpps | **400 W** |
| S310-24ST4X | 24× GE SFP (8 combo) | Non | 4× 10GE SFP+ | 128 Gbit/s | 96 Mpps | — |
| S310-48T4X | 48× GE RJ45 | Non | 4× 10GE SFP+ | 176 Gbit/s | 131 Mpps | — |
| S310-48P4S | 48× GE RJ45 | PoE+ | 4× GE SFP | 104 Gbit/s | 77 Mpps | **380 W** |
| S310-48P4X | 48× GE RJ45 | PoE+ | 4× 10GE SFP+ | à vérifier sur la fiche du modèle exact | à vérifier | à vérifier |

> ⚠️ Les références **48P4X / 48T4X** existent dans la gamme mais leurs valeurs détaillées
> (bruit, poids, conso exacte) sont **à vérifier sur la fiche du modèle exact** avant tout
> dimensionnement électrique ou thermique.

**Logique de nommage (à retenir pour commander sans se tromper) :**

| Lettre | Signification |
|---|---|
| 24 / 48 | Nombre de ports d'accès |
| T | Ports cuivre RJ45 **sans** PoE (ex. 24T4S) |
| P | Ports cuivre RJ45 **avec PoE+** (ex. 24P4S) |
| 4S | 4 uplinks **GE SFP** (1 Gbit/s) |
| 4X | 4 uplinks **10GE SFP+** (10 Gbit/s) |
| N | Ports d'accès **2,5GE** (ex. 24PN4X : 24 ports 2,5G PoE+) |
| ST | Ports d'accès **SFP fibre** (ex. 24ST4X) |

---

## 3. S310-24T4S — fiche technique détaillée (valeurs constructeur)

Le modèle de référence « sans PoE », idéal comme switch d'accès bureautique ou d'agrégation
cuivre quand l'alimentation des équipements est déjà gérée (ou inutile).

| Caractéristique | Valeur constructeur |
|---|---|
| Référence | 98012202 |
| Ports fixes | 24× 10/100/1000BASE-T + 4× GE SFP |
| Capacité de commutation | 56 Gbit/s |
| Débit de transfert | 42 Mpps |
| Table MAC | 16 000 entrées (valeur catalogue distributeur, à vérifier sur la fiche du modèle exact) |
| VLAN max | 4 094 |
| Dimensions (H×L×P) | 43,6 × 442 × 220 mm — 1U |
| Poids (avec emballage) | 3,44 kg |
| Alimentation | AC intégrée, 100–240 V AC 50/60 Hz (plage max 90–290 V AC, 45–65 Hz) |
| Consommation max | **34,04 W** |
| Bruit (puissance acoustique) | 47 dB(A) à température normale / 51 dB(A) à haute température |
| Pression acoustique | 35 dB(A) à température normale |
| Température de fonctionnement | –5 °C à +50 °C (0–1 800 m ; au-delà, –1 °C par tranche de 220 m jusqu'à 5 000 m) |
| Humidité | 5 % à 95 % HR, sans condensation |
| Refroidissement | Air, vitesse des ventilateurs intelligente |
| Protection surtension | ±6 kV (mode différentiel et mode commun) |
| Montage | Rack 19", bureau, mural |
| Management | Cloud eKit (NETCONF/YANG), web, CLI, SNMPv1/v2c/v3, SSH v2.0, Telnet |
| Stacking | iStack (jusqu'à 4 unités de la même série) |

✅ **Bon réflexe d'achat :** si tu n'as besoin d'aucun PoE aujourd'hui mais que des AP ou
caméras sont prévus dans 2 ans, prends directement le **24P4S**. Le surcoût est inférieur au
coût d'un remplacement plus un injecteur PoE par équipement.

---

## 4. S310-24P4S — fiche technique (le modèle PoE+ 24 ports)

Le cheval de bataille des PME : 24 ports PoE+ pour AP Wi-Fi, caméras, téléphones IP,
plus 4 SFP pour la fibre vers le cœur.

| Caractéristique | Valeur constructeur |
|---|---|
| Ports fixes | 24× 10/100/1000BASE-T **PoE+** + 4× GE SFP |
| Capacité / débit | 56 Gbit/s / 42 Mpps |
| Budget PoE total | **400 W** |
| Standard PoE | IEEE 802.3af/at (PoE/PoE+, 30 W max par port) |
| Consommation max | **491,66 W** à pleine charge PoE / 47,1 W sans PoE |
| Poids (avec emballage) | 3,79 kg |
| Bruit | 49,3 dB(A) normal / **63 dB(A)** à haute température |
| Pression acoustique | 37,3 dB(A) à température normale |
| Autres | Identique au 24T4S (dimensions, température, ±6 kV, iStack, management) |
| Fonctions PoE avancées | **Fast PoE** (alimentation des PD en quelques secondes après mise sous tension) et **Perpetual PoE** (le PoE reste alimenté pendant un redémarrage / une mise à jour logicielle) |

⚠️ **63 dB(A) à haute température**, c'est le niveau sonore d'une conversation forte :
ne mets pas un 24P4S chargé en PoE dans un bureau occupé sans réfléchir. Prévois un local
technique ou un placard ventilé (voir section 13).

---

## 5. Les versions 10G : S310-24T4X et S310-24P4X

Quand les uplinks 1G deviennent le goulot (NAS, serveurs, inter-sites, Wi-Fi 6 dense),
les modèles **X** apportent 4× 10GE SFP+.

| Caractéristique | S310-24T4X | S310-24P4X |
|---|---|---|
| Ports | 24× GE (sans PoE) | 24× GE **PoE+** |
| Uplinks | 4× 10GE SFP+ | 4× 10GE SFP+ |
| Capacité / débit | 128 Gbit/s / 96 Mpps | 128 Gbit/s / 96 Mpps |
| Budget PoE | — | **400 W** |
| Conso max | 35,04 W | à vérifier sur la fiche du modèle exact |
| Bruit | 47 dB(A) / 51 dB(A) | à vérifier sur la fiche du modèle exact |

