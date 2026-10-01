---
id: collect-261001-rattrapage/rattrapage/huawei-ap361-guide-6
title: "Huawei eKit AP361 — Guide ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/huawei_ap361_guide.md
source_anchor: ""
source_lines: [764, 919]
sha256: ae9f38a31cf8ecea4e12ddcc69eff5277e32511de18ae3de3f9d97c814c9bfcf
---

# Huawei eKit AP361 — Guide ultra-complet

| Bande | Points forts | Points faibles | Usage conseillé |
|---|---|---|---|
| **2,4 GHz** | Porte loin, traverse (un peu) les murs | 3 canaux non chevauchants seulement, encombrée (Bluetooth, micro-ondes, voisins) | IoT, anciens terminaux, couverture de secours |
| **5 GHz** | Beaucoup de canaux, débits élevés, moins d'interférences | Portée plus courte, traverse mal le béton | Usage principal : bureautique, voix, vidéo |

**Règle d'or** : poussez les clients capables vers le **5 GHz** (band steering,
§52). Le 2,4 GHz est une ressource rare : ne la gaspillez pas.

## 46. Plan de canaux 2,4 GHz : 1, 6, 11 — et rien d'autre

En 2,4 GHz, seuls **3 canaux ne se chevauchent pas** en largeur 20 MHz : **1, 6, 11**
(en Europe ; le 13 existe mais chevauche partiellement — à éviter en planifié).

```
Canaux 2,4 GHz (20 MHz) :

 1        6        11      <- utilisez UNIQUEMENT ceux-ci
 |####|   |####|   |####|
 1  2  3  4  5  6  7  8  9  10 11 12 13
         ^^^^^^^^^^^^^^^^
         chevauchement = interférences
```

- **Largeur 20 MHz uniquement** en 2,4 GHz. Le 40 MHz en 2,4 GHz, c'est occuper
  2/3 de la bande tout seul : égoïste et contre-productif dès qu'il y a des voisins.
- En planifié : alternez 1/6/11 entre AP adjacents (motif en nid d'abeille).
- **Ne laissez pas « Auto » sans contrôle** : vérifiez après 48 h les canaux
  réellement choisis (voir §50).

## 47. Plan de canaux 5 GHz : le terrain de jeu

En Europe (ETSI), canaux 20 MHz utilisables : **36, 40, 44, 48** (UNII-1, sans DFS),
**52-64** (UNII-2, DFS), **100-140** (UNII-2e, DFS), **149-165** restreints selon pays.

| Groupe | Canaux 20 MHz | DFS (radar) | Usage |
|---|---|---|---|
| UNII-1 | 36, 40, 44, 48 | Non | ✅ Premier choix : pas de DFS, mise en service immédiate |
| UNII-2 | 52, 56, 60, 64 | Oui | OK si pas de radar à proximité (aéroports, météo, militaires) |
| UNII-2e | 100-140 | Oui | OK, plus de canaux pour les grands déploiements |
| UNII-3 | 149+ | Selon pays | 🔎 Vérifiez la réglementation locale |

**DFS** : si l'AP détecte un radar, il doit **changer de canal** (coupure de
quelques dizaines de secondes à 10 min). Près d'un aéroport ou d'une station
météo : **restez sur 36-48** pour éviter les bascules intempestives.

**Plan 80 MHz** (canaux « larges ») : 42 (36-48), 58 (52-64), 106, 122, 138, 155.
**Plan 40 MHz** : paires 36-40, 44-48, 52-56, etc.

## 48. Largeur de canal : 20, 40 ou 80 MHz — que choisir

| Largeur | Débit max théorique (2x2, 5 GHz) | Quand l'utiliser |
|---|---|---|
| 20 MHz | ~287 Mbit/s (ax) | 2,4 GHz (toujours) ; 5 GHz en zone très dense / beaucoup d'AP |
| 40 MHz | ~574 Mbit/s (ax) | 5 GHz : bon compromis densité/débit |
| 80 MHz | ~1201 Mbit/s (ax) | 5 GHz : peu d'AP voisins, besoin de débit brut |

**Recommandation terrain** :

- **2,4 GHz : 20 MHz, toujours.**
- **5 GHz : 40 MHz par défaut** en PME (bon équilibre). Passez en **80 MHz**
  uniquement si : peu d'AP (≤ 4), pas de voisins Wi-Fi agressifs, et besoin réel
  de > 400 Mbit/s par client (rare en bureautique).
- Le 80 MHz « mange » 4 canaux : en immeuble de bureaux, vous allez vous battre
  avec les voisins. Le 40 MHz est le choix adulte.

## 49. Puissance d'émission : moins, c'est souvent mieux

Réflexe débutant : « mettre la puissance à fond pour mieux couvrir ». **Erreur.**
Trop de puissance = :

- L'AP arrose l'étage du dessus et les voisins (interférences).
- Les clients « entendent » l'AP de loin mais l'AP n'entend pas leurs réponses
  faibles → **liaison asymétrique**, paquets perdus.
- Les clients restent « collés » à un AP lointain au lieu de roamer (sticky client).

| Contexte | Puissance 5 GHz conseillée | Puissance 2,4 GHz conseillée |
|---|---|---|
| Open space dense (AP tous les 15 m) | 11-14 dBm | 11-14 dBm |
| Bureaux cloisonnés | 14-17 dBm | 14-17 dBm |
| Entrepôt / grand volume | 17-20 dBm | 17-20 dBm |
| Maximum réglementaire (ETSI) | 23 dBm (200 mW) en 5 GHz UNII-1… selon bande | 20 dBm (100 mW) en 2,4 GHz |

> 💡 **Méthode** : commencez à **14 dBm** sur les deux bandes, mesurez (-60 dBm
> cible en zone de travail), ajustez par pas de 3 dB. Notez chaque changement.

🔎 Les plages exactes réglables sur l'AP361 (pas de 1 dB ? max ?) dépendent du
firmware : vérifiez dans l'app eKit.

## 50. DCA (sélection dynamique de canal) : l'auto qui se surveille

Le DCA choisit automatiquement le meilleur canal selon les interférences mesurées.
C'est utile, mais :

- **Vérifiez le résultat** : après 48 h, contrôlez les canaux attribués. Un DCA
  mal réglé met deux AP voisins sur le même canal.
- **Figez si c'est bon** : une fois le plan validé, passez en manuel pour éviter
  les changements intempestifs (un changement de canal = microcoupure pour les
  clients).
- **Planifiez les scans** : si le DCA est actif, limitez les scans aux heures
  creuses (un scan = l'AP quitte son canal quelques ms).

## 51. TPC (contrôle de puissance) : l'auto, bis

Même logique que le DCA : l'AP ajuste sa puissance selon la couverture mesurée.
En PME, **préférez le réglage manuel** (§49) : vous avez 5-15 AP, pas 500.
Le TPC auto a du sens dans les très grands parcs homogènes.

## 52. Band steering : pousser vers le 5 GHz

Le band steering incite les clients dual-band à s'associer en **5 GHz** plutôt
qu'en 2,4 GHz (en retardant/ignorant leurs sondes en 2,4 GHz).

- ✅ Activez-le sur le SSID bureautique : ça désengorge le 2,4 GHz.
- ⚠️ Certains objets connectés **anciens** (imprimantes, IoT) ne supportent que
  le 2,4 GHz et supportent mal le steering agressif → prévoyez un **SSID IoT
  dédié en 2,4 GHz uniquement** (voir §60).
- Si un client « ne voit pas » le SSID : premier suspect = band steering trop
  agressif + client 2,4 GHz only (voir §122).

## 53. Airtime fairness : contre les clients boulets

Un vieux client en 802.11b à 1 Mbit/s peut **monopoliser le temps d'antenne**
et plomber tout le monde. L'airtime fairness donne à chaque client un temps
d'émission équitable plutôt qu'un nombre de paquets équitable.

- ✅ Activez-le sur les SSID à forte mixité de clients (invités, écoles).
- C'est transparent pour les bons clients, salvateur pour le débit global.

## 54. Recouvrement de cellules : le dosage du roaming

Pour un roaming fluide, les cellules doivent se **recouvrir de 15 à 20 %** :

- Trop peu de recouvrement → trous de couverture, coupures.
- Trop de recouvrement → interférences co-canal, sticky clients.

**Méthode terrain** : marchez entre deux AP avec l'analyseur Wi-Fi. Le client
doit voir les deux AP à **-67 dBm ou mieux** dans la zone de transition. Si ce
n'est pas le cas : rapprochez les AP ou montez légèrement la puissance.

## 55. Configurer la radio via l'app eKit (pas à pas)

1. `eKit > Site > Paramètres Wi-Fi > Radio`.
2. Par AP (ou par groupe) : bande 2,4 GHz → canal **1/6/11**, largeur **20 MHz**,
   puissance **14 dBm** (point de départ).
3. Bande 5 GHz → canal **36/40/44/48** (ou plan DFS validé), largeur **40 MHz**,
   puissance **14 dBm**.
4. Activez **band steering** sur le SSID bureautique.
5. Appliquez, attendez la propagation (1-2 min), vérifiez dans l'app que les
   valeurs sont bien prises en compte.
6. **Mesurez** : tour du site à l'analyseur, notez les niveaux par zone.

## 56. Exemples CLI (mode autonome — syntaxe indicative VRP)

> ⚠️ Syntaxe **plausible** pour un AP Huawei en mode Fat ; vérifiez sur votre
> version. En mode cloud, utilisez l'app.

