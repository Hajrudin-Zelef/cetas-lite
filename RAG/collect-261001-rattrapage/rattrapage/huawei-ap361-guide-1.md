---
id: collect-261001-rattrapage/rattrapage/huawei-ap361-guide-1
title: "Huawei eKit AP361 — Guide ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: ["2026-09-27"]
keywords: ["ethernet"]
source: docs/RAG/collect-261001-rattrapage/huawei_ap361_guide.md
source_anchor: ""
source_lines: [1, 145]
sha256: bdef6cf82ad31945b82111a8afa660f29cc463c8a9e532d87880dfaa99e534cb
---

# Huawei eKit AP361 — Guide ultra-complet

**Point d'accès Wi-Fi 6 d'intérieur — Gamme eKit PME**
**Rédigé pour Zelef, chef de service Systèmes & Énergies — ton terrain, direct, sans blabla.**
**Version du guide : 2026-09-27**

> ⚠️ **Avertissement de méthode** : les caractéristiques marquées ✅ sont tirées de la
> fiche produit officielle Huawei eKit (datasheet AP361) et de fiches distributeurs
> recoupées en septembre 2026. Les points marqués 🔎 **« à vérifier sur la fiche du
> modèle exact »** n'ont pas pu être confirmés par une source officielle accessible :
> vérifiez-les sur l'étiquette du produit, le *Hardware Installation and Maintenance
> Guide* WLAN Huawei, ou ekit.huawei.com avant de les appliquer en production.
> Les commandes CLI sont syntaxiquement plausibles pour VRP Huawei ; la syntaxe exacte
> dépend de la version logicielle — l'app eKit / le cloud eKit restent le chemin
> officiel de configuration. Tous les mots de passe et clés du guide sont **fictifs**.

---

## Sommaire

- **A. Fiche produit** — sections 1 à 10
- **B. Montage mécanique** — sections 11 à 17
- **C. Alimentation PoE** — sections 18 à 25
- **D. Première mise en route** — sections 26 à 32
- **E. Gestion : eKit app, cloud, Fat/Fit/Cloud** — sections 33 à 44
- **F. Configuration radio** — sections 45 à 58
- **G. SSID et sécurité d'accès** — sections 59 à 72
- **H. VLAN** — sections 73 à 80
- **I. Itinérance (roaming)** — sections 81 à 88
- **J. Sécurité Wi-Fi (WIDS, rogue, isolation)** — sections 89 à 96
- **K. QoS / WMM** — sections 97 à 103
- **L. Supervision** — sections 104 à 111
- **M. Sauvegarde, restauration, firmware** — sections 112 à 119
- **N. Dépannage terrain (16 cas)** — sections 120 à 137
- **O. Maintenance préventive** — sections 138 à 143
- **P. Durcissement** — sections 144 à 151
- **Q. Pense-bête de poche** — sections 152 à 156
- **R. Glossaire** — section 157
- **S. Quiz (10 questions + réponses)** — sections 158 à 159
- **T. Pour aller plus loin** — section 160

---

# A. FICHE PRODUIT

## 1. L'AP361, c'est quoi exactement

Le **Huawei eKitEngine AP361** est un point d'accès Wi-Fi 6 (802.11ax) d'intérieur,
positionné sur le segment **PME / SOHO** via la gamme **Huawei eKit** (anciennement
« eKitEngine »). C'est l'entrée de gamme « plafond » de la famille eKit : petit,
léger, alimenté en PoE, pensé pour être déployé en nombre sans contrôleur physique
dédié.

Cas d'usage typiques cités par le constructeur ✅ :

| Environnement | Pourquoi l'AP361 colle |
|---|---|
| Bureaux PME / SOHO | Densité moyenne, 1 AP pour 15-25 utilisateurs |
| Petits/moyens hôpitaux | Couverture chambres/couloirs, coût maîtrisé |
| Commerces, retail | SSID invité + SSID caisse/terminaux |
| Hôtellerie économique | 1 AP pour 4-8 chambres selon cloisonnement |
| Écoles primaires/secondaires | Salles de classe, faible densité simultanée |

Ce n'est **pas** un AP de très haute densité (pas de 4x4, pas de 6 GHz) : pour des
amphithéâtres, halls de gare ou stades, viser la gamme AirEngine (ex. AP5760) ou
l'équivalent eKit haut de gamme.

## 2. Caractéristiques hardware vérifiées (constructeur)

Tableau de synthèse — valeurs recoupées entre la datasheet officielle et plusieurs
fiches distributeurs (septembre 2026) :

| Caractéristique | Valeur ✅ | Source |
|---|---|---|
| Norme Wi-Fi | 802.11ax (Wi-Fi 6), rétrocompatible a/b/g/n/ac | Datasheet |
| Bande 2,4 GHz | 2x2 MIMO, jusqu'à **575 Mbit/s** | Datasheet |
| Bande 5 GHz | 2x2 MIMO, jusqu'à **1 200 Mbit/s** | Datasheet |
| Débit agrégé max | **1,775 Gbit/s** | Datasheet |
| Antennes | Intégrées « smart antenna » (commutation intelligente) | Datasheet |
| Gain d'antenne | **5 dBi** (valeur distributeur, à confirmer) | 🔎 fiche distributeur |
| Ports Ethernet | **1 x GE RJ45** (10/100/1000BASE-T) | Datasheet |
| Alimentation | **PoE 802.3af** (rétrocompatible 802.3at) | Datasheet + fiches |
| Consommation max | **8,8 W** (datasheet) ; certains revendeurs annoncent 9,4 W | Datasheet |
| Dimensions | **Ø 180 mm x 35 mm** (hauteur) | Fiches distributeurs |
| Poids | **≈ 0,45 kg** | Fiches distributeurs |
| Montage | Plafond et mur | Datasheet |
| Température de fonctionnement | **-10 °C à +50 °C** (fiche Digitec/Galaxus) ; un revendeur annonce 0 à +40 °C | 🔎 à vérifier selon révision |
| Humidité | 5 % à 95 % HR, sans condensation | Fiche revendeur |
| Altitude | jusqu'à 5 000 m | Fiche revendeur |
| Modes de fonctionnement | **Fit, Fat, Cloud** | Datasheet |
| Sécurité Wi-Fi | WPA/WPA2/WPA3 (PSK et EAP selon fiches) | Fiches distributeurs |
| Gestion | App mobile **HUAWEI eKit**, cloud eKit, Web | Datasheet + fiches |
| Référence fabricant | **50086871** | Fiches distributeurs |

Technologies radio annoncées ✅ : MU-MIMO (montant et descendant, 2,4 + 5 GHz),
OFDMA, 1024-QAM (+25 % d'efficacité vs 256-QAM du Wi-Fi 5), beamforming,
BSS Coloring (réduction des interférences inter-BSS).

## 3. Ce qui est incertain — à vérifier sur la fiche du modèle exact

Soyons honnête, voici ce que la datasheet publique **ne détaille pas** clairement :

- 🔎 **Port console** : la datasheet ne mentionne qu'**1 port GE**. La présence d'un
  port console (RJ45 ou micro-USB) n'est pas confirmée pour l'AP361. Si vous
  prévoyez une procédure de secours par console, vérifiez physiquement l'appareil.
- 🔎 **Bouton reset** : très probable (trou « Reset » comme sur les autres AP
  Huawei), mais non documenté dans l'extrait de datasheet consulté.
- 🔎 **Comportement exact des LED** : les codes couleur/clignotement (boot, upgrade,
  défaut) sont décrits dans le *Hardware Installation and Maintenance Guide* WLAN
  (chapitre AP361) — pas dans la datasheet commerciale.
- 🔎 **Température de fonctionnement** : -10/+50 °C (Digitec, Galaxus) vs 0/+40 °C
  (revendeur Maroc). Écart probable entre révision hardware ou marge commerciale :
  retenez **0 à +40 °C** comme plage sûre en intérieur climatisé, et vérifiez
  l'étiquette.
- 🔎 **Version logicielle** : l'AP361 tourne sous le logiciel système eKit
  (dérivé VRP). Le numéro exact (ex. V200Rxxx) évolue — vérifiez dans l'app eKit
  ou sur ekit.huawei.com au moment du déploiement.
- 🔎 **PoE OUT / passthrough** : non mentionné pour l'AP361 (contrairement à
  certains AP muraux « wall plate »). Considérez qu'il n'y en a pas.
- 🔎 **USB / IoT / BLE** : non mentionnés → considérez comme absents.

> 💡 **Réflexe terrain** : avant d'acheter en volume, demandez au fournisseur la
> *datasheet PDF officielle* du lot (elle porte la révision hardware) et le
> *Hardware Installation and Maintenance Guide* correspondant. C'est 10 minutes qui
> évitent les surprises de chantier.

## 4. Contenu du carton (typique, à vérifier à réception)

D'après les pratiques Huawei eKit (🔎 à vérifier à l'ouverture, ça varie selon les lots) :

- [ ] L'AP361 lui-même (film de protection sur le dôme)
- [ ] Kit de fixation plafond (platine + vis/chevilles)
- [ ] Kit de fixation murale (souvent la même platine)
- [ ] Guide de démarrage rapide (QR code vers l'app eKit)
- [ ] Étiquette avec **numéro de série (SN)** et adresse **MAC** — photographiez-la,
      c'est votre identifiant d'onboarding cloud

**Non fourni** (à commander à part) : injecteur PoE, câble Ethernet, switch PoE.
L'AP361 n'a **pas** d'adaptateur secteur fourni : il est conçu pour être alimenté
exclusivement en PoE (802.3af).

## 5. Ports, LED, boutons — lecture terrain

### 5.1. Le port GE (unique)

