---
id: collect-261001-rattrapage/rattrapage/huawei-ap761-guide-7
title: "Guide ultra-complet — Huawei eKit AP761"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: ["2026-09-27"]
keywords: ["ethernet"]
source: docs/RAG/collect-261001-rattrapage/huawei_ap761_guide.md
source_anchor: ""
source_lines: [542, 646]
sha256: 781dff19b6e8130d15a0f72e6413f143fdead2d6bf4a2bb320d90b6ead24349e
---

# Guide ultra-complet — Huawei eKit AP761

| Fonction | AP761 (Wi-Fi 6) [constructeur] | Wi-Fi 7 (théorie / modèles eKit Wi-Fi 7) |
|---|---|---|
| Standard | 802.11ax | 802.11be |
| MLO | ❌ Non | ✅ (AP771/AP772E/AP371/AP673) |
| Canaux 320 MHz | ❌ Non (max 80 MHz) | ⚠️ Seulement avec 6 GHz (AP673 indoor — à vérifier) ; **non sur l'outdoor eKit** |
| 4096-QAM | ❌ Non (max 1024-QAM) | ✅ (modèles Wi-Fi 7) |
| Bande 6 GHz | ❌ Non | ⚠️ AP673 indoor uniquement — à vérifier ; pas sur l'outdoor |
| Multi-RU / puncturing | ❌ Non | ✅ (modèles Wi-Fi 7) |
| 16 flux spatiaux | ❌ Non (2×2) | Standard : oui ; produits eKit : à vérifier par modèle |
| WPA3 | ✅ Supporté (SAE + Enterprise) | ✅ (obligatoire pour la certification Wi-Fi 7) |
| OFDMA / MU-MIMO | ✅ | ✅ (améliorés) |
| BSS Coloring / TWT | ✅ | ✅ |
| Port 2.5G / 10G | ❌ (1× GE + 1× SFP GE) | ✅ AP772E : 2.5GE + 10GE SFP+ |
| PoE++ 802.3bt | ❌ Non requis (at/af) | ✅ Requis/recommandé (AP772E) |

## 21. La gamme eKit Wi-Fi 7 réelle : AP771, AP772E, AP371, AP673

Si un jour tu dois répondre à « on veut du Wi-Fi 7 », voici les vrais modèles eKit (valeurs vérifiées le 27/09/2026 — les détails fins restent **à vérifier sur la fiche du modèle exact**) :

**Outdoor (les remplaçants « esprit AP761 » en Wi-Fi 7) :**
| | AP771 | AP772E |
|---|---|---|
| Wi-Fi | 7 (802.11be), bi-bande 2.4/5 GHz | 7 (802.11be), bi-bande 2.4/5 GHz |
| Débit max | 3.57 Gbps | **6.45 Gbps** (5.76 @ 5 GHz + 0.688 @ 2.4 GHz) |
| Ports | à vérifier | **1× 2.5GE RJ45 (PoE) + 1× 10GE SFP+** |
| PoE | à vérifier | **802.3af/at/bt** ou DC |
| Protection | à vérifier | IP68, −40 à +70 °C, 6 kA |
| Portée indicative | ~130 m | ~250 m (optimale) |
| Fonctions Wi-Fi 7 | MLO, 4096-QAM, multi-RU | MLO, 4096-QAM, OFDMA, MU-MIMO |

**Indoor :**
| | AP371 | AP673 |
|---|---|---|
| Wi-Fi | 7, bi-bande 2.4/5 GHz | 7, **tri-bande** (avec 6 GHz) |
| Débit max | 3.57 Gbps (BE3600) | à vérifier |
| Ports | 1× 2.5GE PoE | à vérifier |
| Particularité | Le Wi-Fi 7 d'entrée de gamme | **320 MHz sur 6 GHz**, MLO, 4096-QAM |

**Lecture stratégique :** Huawei a choisi de ne pas mettre de 6 GHz sur l'outdoor eKit (réglementation UE, coût). Le Wi-Fi 7 outdoor eKit = du Wi-Fi 7 « sans 320 MHz », dont les gains viennent du MLO et de l'efficacité. Si ton besoin outdoor est de la **couverture fiable**, l'AP761 Wi-Fi 6 reste pertinent et deux fois moins cher ; si c'est de la **densité/latence**, regarde l'AP772E — et refais tout ton budget PoE et câblage (chap. 25).

## 22. Installation physique : où poser un AP extérieur

Un AP761 mal placé ne sera jamais rattrapé par la configuration. Les règles de placement, dans l'ordre d'importance :

1. **Vue directe sur la zone à couvrir.** Le Wi-Fi ne traverse pas bien les murs, et en extérieur les obstacles sont les bâtiments, les arbres (l'eau des feuilles absorbe le 5 GHz), les camions stationnés. Un AP à 6 m avec vue directe couvre 3× mieux qu'un AP à 3 m derrière un auvent métallique.
2. **Au centre du besoin, pas au centre du bâtiment.** On fixe souvent l'AP « où c'est facile » (à côté du local technique). Non : on le met **où sont les clients**, quitte à tirer 80 m de câble.
3. **Hauteur : 4 à 8 m.** En dessous de 4 m : vandalisme, obstacles humains/véhicules. Au-dessus de 8 m : zone morte au pied du mât (lobe vertical 20° en 5 GHz — chap. 9), et les clients voient l'AP sous un mauvais angle.
4. **Loin des sources de bruit :** moteurs, variateurs, postes de soudure, fours à micro-ondes des cuisines extérieures, mâts d'antennes 4G/5G (saturation du récepteur), enseignes LED à alimentation à découpage bon marché.
5. **Penser à la maintenance :** pourra-t-on y accéder dans 3 ans sans louer une nacelle ? Un AP inaccessible = un AP jamais dépoussiéré, jamais vérifié.

**Schéma de réflexion (à dessiner sur plan avant chaque chantier) :**
```
        [Bâtiment]
            |
   (zone morte = arrière de l'AP)
            |
    ┌───────AP761───────┐  <- 65° de secteur utile
    │                   │
    │   COUR / PARKING  │  <- clients ici, vue directe
    │                   │
```
Flécher le secteur de 65° sur le plan. Tout client hors du cône = non couvert. Si la zone fait plus de 65°, il faut 2 AP en secteurs opposés ou un autre modèle.

## 23. Montage mural et sur mât — procédure pas à pas

**Montage mural (façade) :**
1. Choisir un mur **porteur** (pas un bardage seul). Sonder si doute.
2. Présenter l'étrier, marquer les perçages au niveau (un AP penché = un secteur penché).
3. Percer, cheviller (chevilles adaptées au support : béton, brique creuse, parpaing — pas les mêmes).
4. Fixer l'étrier, **serrage au couple raisonnable** : trop serré sur brique creuse = éclatement.
5. Présenter l'AP, orienter le secteur vers la zone (chap. 22), régler le downtilt (quelques degrés vers le bas si la zone est en contrebas).
6. Passer le câble par le **presse-étoupe orienté vers le bas**, serrer le presse-étoupe **à la main + un quart de tour** (trop serré = écrasement du câble et du joint ; pas assez = entrée d'eau).
7. Faire une **boucle d'égouttage** : le câble descend sous l'AP puis remonte — l'eau suit la gravité et goutte au point bas au lieu de suivre le câble jusqu'au presse-étoupe.

**Montage sur mât :**
1. Vérifier le mât : verticalité, rigidité (pas de fléchissement au vent), mise à la terre existante.
2. Utiliser les colliers du kit (ou colliers inox adaptés au diamètre du mât — **à vérifier sur la fiche du modèle exact** pour la plage de diamètres).
3. Serrer en croix, progressivement. Revérifier après 48 h (tassement).
4. Même règle : presse-étoupe vers le bas, boucle d'égouttage.

**Erreurs vues sur le terrain :**
- ❌ AP fixé sur une gouttière ou un bardage léger : ça vibre, ça se desserre, le secteur bouge.
- ❌ Câble qui arrive **par le haut** dans le presse-étoupe : entonnoir à eau garanti.
- ❌ Serrage du presse-étoupe à la pince : joint écrasé, IP68 foutu.
- ❌ AP collé sous un auvent métallique : le métal fait écran, le secteur est haché.

## 24. Câblage : choix du câble, distance, étanchéité

**Le câble :** c'est le composant le moins cher du chantier et la cause n° 1 des pannes « mystères » 2 ans après. Ne pas lésiner.

| Critère | Exigence |
|---|---|
| Type | Câble **extérieur** (gaine PE résistante UV), **blindé F/UTP ou S/FTP** |
| Âme | **100 % cuivre**, 24 AWG (ou 23 AWG). **Jamais de CCA** (cuivre plaqué alu) |
| Catégorie | Cat6 minimum (le GE de l'AP761 passe sur Cat5e, mais le Cat6 laisse de la marge PoE et diaphonie) |
| Longueur max | **100 m** (90 m de câble + 10 m de cordons), norme Ethernet |
| Cheminement | Gaine ICTA ou chemin de câble ; **pas de câble nu au soleil** long terme même « extérieur » |

**Distances > 100 m :** c'est le cas d'usage du **port SFP** de l'AP761 — tirer une fibre optique jusqu'à l'AP (jarretière + SFP GE au bout), et alimenter en PoE via un injecteur local ou un switch PoE proche (chap. 6). La fibre règle aussi les problèmes de **différence de potentiel** entre bâtiments (pas de boucle de masse).

**Étanchéité des traversées :**
- Traversée de façade : fourreau + mastic/silicone extérieur, pente vers l'extérieur.
- Arrivée au local technique : le câble ne doit pas faire entrer l'eau dans la baie (boucle d'égouttage avant l'entrée).
- Connecteurs RJ45 : **ne jamais laisser un RJ45 nu en extérieur** — toujours le brancher sous le capot/presse-étoupe de l'AP ou dans un boîtier étanche. Un RJ45 oxydé = erreurs CRC = débit qui s'effondre par paliers (cas 4).

