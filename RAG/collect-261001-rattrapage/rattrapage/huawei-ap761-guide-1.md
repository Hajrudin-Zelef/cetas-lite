---
id: collect-261001-rattrapage/rattrapage/huawei-ap761-guide-1
title: "Guide ultra-complet — Huawei eKit AP761"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: ["2026-09-27"]
keywords: ["distribution"]
source: docs/RAG/collect-261001-rattrapage/huawei_ap761_guide.md
source_anchor: ""
source_lines: [1, 37]
sha256: 403d7c6a823e602e6d4c60f6d27e91c1fa30d5922e8df4d35dca0c33098ed400
---

# Guide ultra-complet — Huawei eKit AP761

**Sous-titre :** point d'accès extérieur Wi-Fi 6 — installation, configuration, supervision, dépannage terrain, et tout ce qu'il faut savoir sur le Wi-Fi 7 pour préparer la suite.

**Public :** Zelef, chef de service systèmes & énergies — ton terrain, concret, direct.
**Date de rédaction :** 27 septembre 2026. **Langue :** français.

> ⚠️ **AVERTISSEMENT DE FONDATION — À LIRE EN PREMIER**
>
> La mission initiale décrivait l'AP761 comme un « point d'accès Wi-Fi 7 ». **Vérification faite auprès des fiches constructeur et des distributeurs (recherche web du 27/09/2026) : c'est FAUX.**
>
> - **Le Huawei eKit AP761 (réf. 02355VFB) est un point d'accès EXTÉRIEUR Wi-Fi 6 (802.11ax)**, bi-bande 2.4 GHz + 5 GHz, 2×2, 1.775 Gbps max. Pas de Wi-Fi 7, pas de MLO, pas de 320 MHz, pas de 6 GHz, pas de 4096-QAM sur ce modèle.
> - **Les modèles eKit Wi-Fi 7 d'extérieur sont l'AP771 (3.57 Gbps) et l'AP772E (6.45 Gbps)** ; en intérieur : AP371 (3.57 Gbps) et AP673 (tri-bande avec 6 GHz / 320 MHz).
>
> Ce guide est donc construit honnêtement en deux volets :
> 1. **L'AP761 tel qu'il est vraiment** (Wi-Fi 6 outdoor) : hardware, installation, configuration, dépannage — c'est le gros du document, et c'est du concret vérifié.
> 2. **Le Wi-Fi 7 expliqué simplement** (théorie solide + modèles eKit qui le supportent réellement) pour que tu saches exactement ce que l'AP761 ne fait PAS, et vers quoi migrer si tu veux du Wi-Fi 7 en extérieur.
>
> Chaque fois qu'une valeur vient du constructeur, c'est indiqué. Chaque fois qu'une valeur n'a pas pu être vérifiée, c'est marqué **« à vérifier sur la fiche du modèle exact »** — sans exception, sans invention.

## Sources vérifiées (recherche web du 27/09/2026)

| Source | Ce qu'elle confirme |
|---|---|
| Datasheet constructeur Huawei eKitEngine AP761 (hw-trade.ru, PDF officiel) | Wi-Fi 6 802.11ax ; 2×2 2.4 + 2×2 5 GHz ; 1.775 Gbps ; 1× GE RJ45 + 1× SFP GE combo ; PoE 802.3at/af ; 17.7 W max ; IP68 ; −40 °C à +65 °C ; antennes 10/11 dBi ; 1024 STA ; 16 SSID/radio |
| Fiche distributeur alltronstamm.ch (EAN 6901443451289, réf 02355VFB) | Dimensions 200×200×69 mm ; 1.91 kg ; montage mur/mât ; pas de Mesh ; PoE+ |
| Fiche distributeur 4-electronics.ch | AX1800 ; canaux 20/40/80 MHz ; WPA3-SAE/EAP, WAPI ; BLE 5.2 ; 802.11k/v/r |
| Fiche distributeur technerve.co.ke / microviewng.com | Combo SFP/GE (l'optique est prioritaire) ; 802.3af = fonctions restreintes ; antennes directionnelles 65°H |
| Manuel HUAWEI eKit Scenario-based Products (absolut-distribution.ch) | Comparatif AP761 vs AP361/AP362/AP661 : AP761 = outdoor 1.775 Gbps, 17.7 W ; AP361 = indoor 1.775 Gbps, 8.8 W |
| Fiche bimel.com.tr (AP Wi-Fi 7 eKit outdoor) | Wi-Fi 7 802.11be : 320 MHz **non supporté** sur ce produit outdoor ; 4096-QAM, multi-RU, MLO supportés |
| Fiche microtech.net.pk (AP772E) | Wi-Fi 7 outdoor : 6.45 Gbps, 1× 2.5GE + 1× 10GE SFP+, PoE af/at/bt ou DC, IP68, −40 à +70 °C, bi-bande 2.4/5 GHz |
| Comparatif shop.itegy.com.eg | AP761 = Wi-Fi 6, 1800 Mbps, 1024 users max, 500 m ; AP772 = Wi-Fi 7, 6.45 Gbps ; AP771 = Wi-Fi 7, 3.57 Gbps |

---

## Sommaire

