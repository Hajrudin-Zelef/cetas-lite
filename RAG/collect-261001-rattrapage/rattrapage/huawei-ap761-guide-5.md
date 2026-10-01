---
id: collect-261001-rattrapage/rattrapage/huawei-ap761-guide-5
title: "Guide ultra-complet — Huawei eKit AP761"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["throughput"]
source: docs/RAG/collect-261001-rattrapage/huawei_ap761_guide.md
source_anchor: ""
source_lines: [375, 462]
sha256: 5886210ff488dcdcedf52171b01f8fbae96cedc23a71314a6a5c171c90d5a628
---

# Guide ultra-complet — Huawei eKit AP761

| | **Fat (autonome)** | **Fit (contrôleur)** | **Cloud (eKit)** |
|---|---|---|---|
| Cerveau | L'AP lui-même | Un AC (contrôleur, ex. WAC eKit) | La plateforme cloud Huawei eKit |
| Config | Web/CLI sur chaque AP | Centralisée sur l'AC via CAPWAP | Centralisée sur le cloud via app/web |
| Idéal pour | 1–3 AP, site isolé | Parc homogène, roaming avancé, analyse de spectre | Multi-sites PME, gestion à distance, zéro licence |
| Analyse de spectre | Non | **Oui** (seul mode qui la supporte) | Non |
| SNMP | v1/v2c/v3 | Via l'AC | Limité (voir chap. 68) |
| Prérequis | Rien | Un AC compatible sur le réseau | Compte cloud + Internet |

**Recommandation terrain :**
- **1 AP isolé (terrasse, parking) → Fat ou Cloud.** Le cloud apporte la supervision à distance gratuite : si tu gères plusieurs sites, c'est imbattable.
- **Plusieurs AP avec roaming soigné → Cloud** (suffit en PME) ou **Fit** si tu as déjà un AC.
- **Site sans Internet fiable → Fat.** Le cloud a besoin d'Internet pour la gestion (le Wi-Fi local continue de fonctionner si le cloud tombe, mais tu ne peux plus administrer à distance).

**Changement de mode (principe) :**
```
# En CLI (mode Fat), basculer l'AP vers le mode cloud :
<AP761> system-view
[AP761] ap-mode cloud
# L'AP redémarre et cherche la plateforme cloud (QR code / app).
# Retour en Fat :
[AP761] ap-mode fat
```
> La syntaxe exacte de bascule (`ap-mode`) et les prérequis (firmware) sont **à vérifier sur la fiche du modèle exact** et la version logicielle. Ne jamais basculer un AP en production sans fenêtre de maintenance : ça redémarre et ça réinitialise la configuration locale.

## 12. Wi-Fi 6 (802.11ax) : ce que l'AP761 fait vraiment

Le Wi-Fi 6 n'est pas « le Wi-Fi plus vite » : c'est « le Wi-Fi **plus efficace quand il y a du monde** ». L'AP761 l'implémente en 2×2 sur deux bandes [constructeur]. Les briques qui comptent sur le terrain :

**OFDMA — le découpage de canal :** en Wi-Fi 5, un client monopolisait tout le canal même pour envoyer trois octets. En Wi-Fi 6, le canal est découpé en sous-porteuses (Resource Units) et **plusieurs clients émettent/reçoivent en même temps**. Résultat : latence divisée et bien meilleure tenue en charge quand 50 smartphones se battent pour le canal. C'est LA raison d'être du Wi-Fi 6 dans une cour d'école ou une terrasse bondée.

**MU-MIMO montant et descendant :** l'AP761 parle à plusieurs clients **simultanément** grâce à ses 2 flux spatiaux (2×2), en émission ET en réception [constructeur]. Limite physique : 2 flux = 2 clients en MU-MIMO plein débit en même temps ; au-delà, c'est l'OFDMA qui prend le relais.

**1024-QAM :** chaque symbole transporte 10 bits au lieu de 8 (256-QAM en Wi-Fi 5) → **+25 % de débit théorique** à signal égal [constructeur]. Mais ça exige un signal très propre (SNR élevé). En extérieur avec de la distance, les clients retomberont vite sur 256-QAM ou moins : le 1024-QAM ne sert qu'à courte distance.

**BSS Coloring :** l'AP761 « colore » son réseau [constructeur]. Quand il entend une trame d'un réseau voisin de couleur différente, il peut émettre quand même au lieu d'attendre — **moins d'attente, plus de débit utile** en zone dense (lotissement, zone commerciale). C'est invisible à configurer, mais c'est ce qui fait la différence entre un Wi-Fi 6 et un Wi-Fi 5 dans un environnement chargé.

**TWT (Target Wake Time) :** l'AP négocie avec les objets connectés leurs horaires de réveil [constructeur]. Pour des capteurs sur batterie en extérieur (station météo, compteur), c'est de l'autonomie en plus. Pour des smartphones, gain marginal.

**Débits réels à attendre (ordres de grandeur terrain) :**
| Client | Condition | Débit utile réaliste |
|---|---|---|
| Wi-Fi 6, 2×2, 80 MHz, proche (< 15 m, vue directe) | 1024-QAM | 400–600 Mbps |
| Wi-Fi 6, 2×2, 80 MHz, 30–50 m | 256-QAM | 150–300 Mbps |
| Wi-Fi 5, 2×2, 80 MHz | 256-QAM | 150–350 Mbps |
| Wi-Fi 4 / vieux smartphone, 20 MHz | 64-QAM | 20–70 Mbps |

Le 1.775 Gbps constructeur est un **débit PHY agrégé théorique** (les deux radios à fond en même temps). En pratique, un client seul ne verra jamais ce chiffre : le débit utile, c'est 50–70 % du débit PHY dans le meilleur cas.

## 13. Rétrocompatibilité Wi-Fi 5/4 et vieux clients

L'AP761 parle **802.11a/b/g/n/ac/ac Wave 2/ax** [constructeur] : un client de 2012 se connecte sans problème. Mais la rétrocompatibilité a un coût, et il faut le connaître :

**Ce qui change pour les vieux clients :**
- Un client Wi-Fi 4 (802.11n) ne comprend ni l'OFDMA ni le MU-MIMO : il est servi en mode « classique », un par un. Il ne casse rien, il ne profite juste pas des optimisations.
- Les débits 802.11b (1, 2, 5.5, 11 Mbps) : si tu les laisses activés, les vieux clients peuvent s'accrocher de loin avec un débit minable et **monopoliser l'airtime**. En extérieur, c'est le piège n° 1 (voir cas 13).
- Le 2.4 GHz reste la bande « refuge » des objets connectés (caméras, portiers, capteurs) : la plupart sont en Wi-Fi 4, parfois même en 802.11g.

**Modes mixtes — réglages recommandés :**
| Paramètre | Réglage conseillé | Pourquoi |
|---|---|---|
| Débits de base 2.4 GHz | Désactiver 1, 2, 5.5, 11 Mbps (garder 6, 12, 24 Mbps min) | Force les clients à ne s'associer que s'ils ont un signal correct ; libère l'airtime |
| Débits de base 5 GHz | Garder 6, 12, 24 Mbps | Le 5 GHz n'a jamais eu de 802.11b, moins critique |
| 802.11ax activé | Oui, toujours | Rétrocompatible par nature |
| WPA3-only | **Non** si tu as des vieux clients | Ils ne connaissent pas SAE (voir chap. 47) |

**Règle terrain :** en extérieur, un client accroché à −85 dBm en 802.11b à 1 Mbps peut à lui seul diviser le débit utile de la cellule par deux. **Mieux vaut un client déconnecté (qui basculera sur la 4G) qu'un client zombie qui pourrit l'airtime de tout le monde.** C'est violent à dire, mais c'est la physique.

## 14. Le Wi-Fi 7 (802.11be) expliqué simplement

> 📌 Ce chapitre est de la **théorie** (standard IEEE 802.11be, alias EHT — Extremely High Throughput, certifié « Wi-Fi CERTIFIED 7 » par la Wi-Fi Alliance en janvier 2024). **L'AP761 ne supporte AUCUNE de ces fonctions.** C'est ici pour que tu saches ce qui existe, et le chapitre 21 dit quels modèles eKit les supportent vraiment.

Le Wi-Fi 6 a optimisé **l'efficacité** (plus de clients bien servis). Le Wi-Fi 7 attaque trois autres fronts : **le débit crête**, **la latence**, et **la fiabilité**. Les quatre briques :

1. **MLO (Multi-Link Operation)** — utiliser plusieurs bandes en même temps (chap. 15).
2. **Canaux 320 MHz** — des canaux deux fois plus larges qu'en Wi-Fi 6 (chap. 16).
3. **4096-QAM (4K-QAM)** — 12 bits par symbole au lieu de 10 (chap. 17).
4. **Preamble puncturing / multi-RU** — utiliser les morceaux de canal non brouillés (chap. 18).

**Ordre de grandeur :** le débit PHY max du standard monte à **46 Gbps** (16 flux, 320 MHz, 4096-QAM). En pratique sur un AP eKit Wi-Fi 7 réel (AP772E) : **6.45 Gbps agrégés**. Et comme toujours : le débit utile sera 50–70 % du PHY, partagé entre les clients.

**Ce que le Wi-Fi 7 change VRAIMENT pour un chef de service :**
- **Latence** : le MLO fait chuter la latence et la gigue — intéressant pour la voix sur Wi-Fi, la visio, et plus tard la réalité augmentée terrain.
- **Densité** : le multi-RU rend l'OFDMA encore plus efficace en environnement très chargé.
- **Backbone** : un AP Wi-Fi 7 a besoin d'un uplink **2.5G minimum** (l'AP772E a du 2.5GE + 10GE SFP+) et de **PoE++ (802.3bt)** : le budget électrique et câblage change d'échelle (chap. 25).

## 15. MLO — Multi-Link Operation, en clair

