---
id: collect-261001-rattrapage/rattrapage/huawei-ekit-guide-2
title: "GUIDE ULTRA-COMPLET — PLATEFORME HUAWEI eKit (PME / DISTRIBUTION)"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-rattrapage/huawei_ekit_guide.md
source_anchor: ""
source_lines: [86, 160]
sha256: 4ead59e6fa4a38dcb552ca180f1ed2c9b375126e255339977ea9e63c48ab0c99
---

# GUIDE ULTRA-COMPLET — PLATEFORME HUAWEI eKit (PME / DISTRIBUTION)

| Modèle | Usage | Wi-Fi | Débit max agrégé | Ports | PoE-in | Particularités |
|---|---|---|---|---|---|---|
| **AP361** | Intérieur (plafond/mur) | Wi-Fi 6, 2,4 GHz 2×2 + 5 GHz 2×2 | **1,775 Gbit/s** (575 + 1200 Mbit/s) | 1× GE | 802.3af (~9 W) | Antenne intelligente intégrée, modes Fit/Fat/Cloud, USB pour IoT (ZigBee/RFID selon fiche) |
| **AP266** | Intérieur | Wi-Fi 6 | à vérifier sur la documentation officielle | à vérifier | à vérifier | Gestion cloud, PSK + portail captif intégrés **sans WAC ni serveur d'authentification** (gros avantage PME) |
| **AP761** | **Extérieur** | Wi-Fi 6 (AX1800 : 574 + 1201 Mbit/s) | 1,775 Gbit/s | 1× GE + **1× SFP** | 802.3at recommandé (af = fonctions limitées), conso max 17,7 W | BLE 5.2 (gestion, localisation, IoT), MU-MIMO, Beamforming, TWT |

### 5.2. Wi-Fi 7 (la vague 2025-2026)

| Modèle | Usage | Wi-Fi | Particularités |
|---|---|---|---|
| **AP572** | Intérieur haute densité (salles de classe, salles de réunion) | Wi-Fi 7 | Positionné « haute densité » par Huawei |
| **AP673E** | Intérieur haute densité | Wi-Fi 7 | Salles de classe / grandes salles de conférence |
| **AP673H (iGuard)** | Intérieur (bureaux, hôtels, santé, éducation) | Wi-Fi 7 **tri-bande** 2,4 (2×2) + 5 (2×2) + **6 GHz (4×4)**, 8 flux spatiaux, **13,66 Gbit/s** | **Détection de caméras espion** par IA (voir section 29) — la détection spycam ne fonctionne qu'en **gestion locale** (pas via cloud, pour des raisons de confidentialité) |
| **AP772** | Extérieur | Wi-Fi 7 | Antenne **omnidirectionnelle 360°** |
| **AP772E** | Extérieur longue distance | Wi-Fi 7 | Antenne **directionnelle** |

**Règle de choix rapide :**
- Bureau/boutique standard → **AP361** (le best-seller, prix contenu, largement suffisant).
- Extérieur (terrasse, parking, piscine d'hôtel) → **AP761** (Wi-Fi 6) ou **AP772/AP772E** (Wi-Fi 7).
- Salle de réunion dense, école → **AP572/AP673E**.
- Hôtel haut de gamme / salle de direction (argument confidentialité) → **AP673H iGuard**.
- Fibre jusqu'au bureau/chambre → **eKitOptix F700D** (MiniFTTO, section 14).

## 6. Les switchs eKitEngine : S110, S220, S310, S620

### 6.1. Tableau de la gamme réelle

| Série | Positionnement | Modèles connus | PoE |
|---|---|---|---|
| **S110** | Entrée de gamme, petits sites | modèles à vérifier sur la documentation officielle | Budget PoE inférieur au S310 (le distributeur cdr.pl le cite comme référence basse) |
| **S220** | PME, L2 managé | **S220-8P4S** : 8× GE + 4× SFP 1G, **125 W PoE+**, 24 Gbit/s, 18 Mpps | 802.3af/at, Fast PoE, Perpetual PoE |
| **S310** | Cœur de gamme PME, L2+ | **S310-48P4S** : 48× GE PoE+ + 4× SFP 1G, **380 W**, 104 Gbit/s / 77 Mpps ; **S310-48P4X** : 48× GE PoE+ + **4× SFP+ 10G**, 380 W, 176 Gbit/s / 131 Mpps ; **S310-24J4X / 48J4X** : 24 ou 48× **2.5GE SFP** (tout-optique) + 4× SFP+ 10G | 802.3af/at sur tous les ports (48P), 380 W |
| **S620** | Haut de gamme eKit | **S620-24T16X8Y2CZ** : 24× GE + 16× 10GE + 8× 25GE + 2× 100GE | à vérifier |

**Détails terrain sur le S310 (le switch que tu déploieras le plus) :**
- Les 4 uplinks SFP/SFP+ acceptent modules optiques GE/10GE, jarretières cuivre DAC et câbles de stacking dédiés (**stacking zéro-configuration** supporté d'après la fiche).
- Mise à jour « intelligente » via **HOUP** (Huawei Online Upgrade Repository) : le switch récupère lui-même son chemin de mise à jour — mise à jour en un clic, avec pré-chargement pour réduire la coupure.
- Gestion : **app eKit (cloud) OU locale** (CLI, web, SNMP, SSH) — les deux modes sont **commutables** selon la fiche. C'est important : tu n'es pas enfermé dans le cloud.
- Protection surtension : ±7 kV en mode commun sur les ports de service, ±6 kV sur l'alimentation (S310-48P4X). Bien pour les sites à réseau électrique capricieux, mais ça ne remplace pas un parafoudre en tête d'installation.
- MTBF annoncé : ~48 ans, dispo > 99,999 %. Valeurs constructeur, à prendre comme ordre de grandeur.

### 6.2. Lecture d'une référence : S310-48P4X

- **S310** = série ; **48** = 48 ports descendants ; **P** = PoE ; **4X** = 4 uplinks SFP+ 10G (un **S** final = SFP 1G ; **J** = ports SFP 2.5GE tout-optique ; **T** = ports cuivre).
- Réflexe : devant une référence inconnue, décode-la au lieu de la googler en panique.

## 7. Le pare-feu eKitEngine USG6000F-S : S125 / S150 / S200

Huawei positionne l'USG6000F-S comme le **pare-feu IA pour PME** : « un boîtier = routeur + switch + antivirus » (formule distributeur). D'après la fiche technique V600R025 :

| Modèle | Débit pare-feu IPv4 | Débit NGFW (FW+SA+IPS) | Débit protection complète (+AV) | IPsec | Utilisateurs recommandés |
|---|---|---|---|---|---|
| **S125** | 8/8/3,6 Gbit/s | 1,7 Gbit/s | 1,4 Gbit/s (HTTP 100K) / 1 Gbit/s (mix réel) | 3,7 Gbit/s, 2000 tunnels | ~600 |
| **S150** | 12/12/4 Gbit/s | 2,1 Gbit/s | 1,8 / 1,2 Gbit/s | 3,5 Gbit/s, 4000 tunnels | ~1000 |
| **S200** | 16/16/5 Gbit/s | 3 Gbit/s | 2,5 / 1,8 Gbit/s | 5,6 Gbit/s, 4000 tunnels | ~2500 |

**À retenir pour le dimensionnement :** le débit qui compte, c'est le **débit « enterprise mix » avec toutes les protections activées** (1 à 1,8 Gbit/s selon modèle), pas le débit pare-feu brut. Si le client a une fibre 2 Gbit/s et active tout, le S125 sera le goulot. Fonctions : IPS, antivirus, filtrage URL (base de 500 M d'URL), anti-spam, SSL VPN (100 utilisateurs par défaut, jusqu'à 1000-2000 en licence), IPsec, portail captif local, SSO RADIUS/AD. Mises à jour de signatures via le centre de sécurité Huawei.

## 8. eKitOptix (MiniFTTO) : la fibre jusqu'à la chambre / au bureau

Le **MiniFTTO** (Fiber To The Office), c'est l'argument « hôtel » d'eKit : remplacer le faisceau de câbles (TV + téléphone + Wi-Fi) par **une seule fibre** jusqu'à chaque chambre/bureau, terminée par un boîtier discret qui fait Wi-Fi + ports Ethernet (+ parfois POTS/TV selon modèle).
- **eKitOptix F700D** : AP Wi-Fi 7 optique « 3-en-1 » (fibre + Wi-Fi + switch local).
- **eKitOptix FG736** : nouvelle génération d'AP Wi-Fi 7 optique.
- Cas d'usage roi : **hôtellerie** (une fibre par chambre, esthétique, pas de goulotte), bureaux neufs, cliniques.
- Contrainte : ça suppose une infrastructure optique passive (OLT en tête, splitters) — ce n'est pas « juste un AP », c'est une architecture. Ne le propose que si le câblage est à créer ou à refaire, ou si le client veut du tout-fibre. Détail d'architecture en section 52.

## 9. Wi-Fi Shield, iGuard et les arguments sécurité « différenciants »

Huawei pousse trois arguments marketing que tu dois savoir **expliquer honnêtement** au client :
1. **Wi-Fi Shield** : présenté comme une protection « zéro-écoute » intégrée au Wi-Fi 7 eKit. Concrètement : chiffrement et mécanismes anti-espionnage renforcés. Ne le vends pas comme « inviolable » — c'est une couche de plus, pas une baguette magique.
2. **iGuard / AP673H** : détection de **caméras espion** par analyse des ondes (CSI — Channel State Information) + IA, y compris des caméras **non connectées au Wi-Fi** (qui stockent en local ou transmettent en 4G/5G). La fiche constructeur précise : la détection spycam fonctionne en **gestion locale uniquement** (confidentialité des données). Il existe aussi un détecteur portable **DF10** avec appli mobile iGuard (mode « meeting guard », ~1,5 h d'autonomie annoncée). Argument fort pour hôtellerie haut de gamme et salles de direction.
3. **App/cloud « 100 % gratuits, zéro licence »** : c'est l'argument économique central (voir nuances section 33).

## 10. Où acheter, où se documenter (sources officielles)

