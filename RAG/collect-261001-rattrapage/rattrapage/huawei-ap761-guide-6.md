---
id: collect-261001-rattrapage/rattrapage/huawei-ap761-guide-6
title: "Guide ultra-complet — Huawei eKit AP761"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei", "United States"]
dates: ["2026-09-27"]
keywords: ["arr"]
source: docs/RAG/collect-261001-rattrapage/huawei_ap761_guide.md
source_anchor: ""
source_lines: [463, 541]
sha256: d263c1f6ed1dd50a2f4ccd4608be97c72e231b959607d52c34b75726c904bda5
---

# Guide ultra-complet — Huawei eKit AP761

**Le concept :** jusqu'ici, un client Wi-Fi n'était connecté que sur **UNE** bande à la fois (2.4 OU 5 GHz). Avec le MLO, un client Wi-Fi 7 peut être connecté **simultanément sur deux (ou trois) bandes** — par exemple 5 GHz + 6 GHz — et l'AP répartit les trames sur les deux liens en temps réel.

**Les deux modes de MLO :**
| Mode | Principe | Usage |
|---|---|---|
| **STR** (Simultaneous Transmit and Receive) | Les deux liens émettent/reçoivent en même temps | Débit max : agrégation pure. Exige une bonne isolation entre radios (matériel haut de gamme). |
| **EMLSR** (Enhanced Multi-Link Single Radio) | Un seul lien actif à la fois, mais bascule instantanée vers le lien libre | Latence min : si le 5 GHz est occupé, ça part sur le 6 GHz sans attendre. Moins gourmand en matériel. |

**Analogie terrain :** avant, tu avais **une voie** entre l'AP et le client (avec des bouchons). Le MLO, c'est **deux voies parallèles** avec un aiguillage intelligent : si une voie est bouchée (interférence, radar DFS), le trafic passe par l'autre **sans couper la connexion**.

**Limites honnêtes :**
- Il faut un **client Wi-Fi 7** aussi : un smartphone Wi-Fi 6 ne fera jamais de MLO, même face à un AP Wi-Fi 7.
- Les premiers clients Wi-Fi 7 gèrent mieux l'EMLSR que le STR : les gains en débit agrégé réel sont variables.
- En extérieur dans l'UE, sans 6 GHz (chap. 19), le MLO se fait entre 2.4 et 5 GHz : utile pour la latence, modeste pour le débit.

> **Sur l'AP761 :** pas de MLO. Point. Si un commercial te parle de MLO sur un AP761, c'est une erreur ou une tromperie.

## 16. Canaux 320 MHz : pourquoi l'outdoor eKit n'en a pas

**La théorie :** le Wi-Fi 7 double la largeur max de canal : **320 MHz** (contre 160 MHz en Wi-Fi 6). Un canal deux fois plus large = deux fois plus de débit PHY à modulation égale.

**La réalité physique :**
- Les canaux 320 MHz n'existent en pratique que dans la **bande 6 GHz** (il faut 320 MHz de spectre contigu — introuvable en 2.4 GHz, très rare en 5 GHz).
- En 5 GHz, le max réaliste reste **160 MHz** (et encore : avec le DFS, trouver 160 MHz propres en extérieur relève du défi — chap. 41).

**Le fait vérifié qui compte :** la fiche d'un AP eKit **outdoor Wi-Fi 7** (bimel.com.tr, vérifiée le 27/09/2026) précise noir sur blanc : **« 320 MHz bandwidth (not supported by this product) »**. Les AP outdoor Wi-Fi 7 eKit (AP771/AP772E) sont **bi-bande 2.4/5 GHz** : pas de 6 GHz, donc pas de 320 MHz. Leurs gains Wi-Fi 7 viennent du **MLO**, du **4096-QAM** et du **multi-RU**, pas des canaux géants.

**Seul l'AP673 (indoor tri-bande)** annonce du 320 MHz sur la bande 6 GHz — **à vérifier sur la fiche du modèle exact**.

**Implication pour toi :** ne dimensionne jamais un projet outdoor sur du 320 MHz. En extérieur UE, le plan de canaux réaliste reste : **20 MHz en 2.4 GHz, 40 ou 80 MHz en 5 GHz** (chap. 37–39).

## 17. 4096-QAM (4K-QAM) : +20 % de débit, à quel prix

**La théorie :** le 4096-QAM encode **12 bits par symbole** contre 10 en 1024-QAM → **+20 % de débit théorique** à largeur de canal égale.

**Le prix à payer :** pour distinguer 4096 points de constellation, il faut un signal **extrêmement propre** (SNR très élevé, typiquement > 40 dB — valeur indicative, à vérifier sur la fiche du modèle exact). En pratique :
- Ça ne marche qu'à **courte distance**, en vue directe, sans interférence.
- En extérieur (distance, vent dans les arbres, humidité), les clients retomberont sur 1024-QAM ou 256-QAM 95 % du temps.
- Le 4K-QAM est donc un **bonus de courte portée**, pas un argument de couverture.

**Support eKit :** le 4096-QAM est annoncé sur les AP eKit Wi-Fi 7 (AP772E, AP673). **L'AP761 s'arrête au 1024-QAM** [constructeur].

**Règle terrain :** ne jamais promettre un débit basé sur le 4K-QAM à un client ou à ta direction. Dimensionne sur du 256-QAM (Wi-Fi 5/6 robuste), et considère tout ce qui est au-dessus comme du bonus.

## 18. Preamble puncturing et multi-RU

**Le problème :** en Wi-Fi 6, si une petite portion d'un canal 80 MHz est brouillée (un radar, un réseau voisin sur un sous-canal), **tout le canal 80 MHz est inutilisable** — l'AP retombe sur 40 ou 20 MHz. C'est du gâchis : 60 MHz propres jetés à cause de 20 MHz brouillés.

**La solution Wi-Fi 7 :**
- **Preamble puncturing** : l'AP « perfore » le canal — il utilise les morceaux propres et ignore le morceau brouillé. Un 80 MHz avec 20 MHz brouillés devient un **60 MHz utiles** au lieu de retomber à 40 MHz.
- **Multi-RU** : en Wi-Fi 6, un client ne recevait qu'**un seul** bloc de sous-porteuses (RU) par trame OFDMA. En Wi-Fi 7, un client peut recevoir **plusieurs RU** à la fois → l'ordonnanceur remplit le canal beaucoup plus efficacement, surtout avec peu de clients gourmands.

**Pourquoi c'est important en extérieur :** l'extérieur est un environnement **sale** en radio (radars météo/aviation en 5 GHz, réseaux voisins, Bluetooth, fours à micro-ondes des food-trucks…). Le puncturing, c'est la différence entre « le canal 80 MHz tient » et « on retombe à 40 MHz dès qu'un radar passe ». C'est une fonction **de robustesse**, pas de débit crête — et c'est pour ça qu'elle compte plus sur le terrain que le 4K-QAM.

> **Sur l'AP761 :** pas de puncturing, pas de multi-RU. En cas d'interférence sur un sous-canal, il retombe à la largeur inférieure (80 → 40 → 20 MHz). D'où l'importance du plan de canaux (chap. 37–39).

## 19. 6 GHz et réglementation : le point pour l'Europe

**La théorie :** le Wi-Fi 7 adore la bande **6 GHz** (5945–7125 MHz) : beaucoup de spectre propre, canaux 320 MHz possibles, pas de clients legacy qui pourrissent l'airtime (le 6 GHz n'accepte que du WPA3).

**La réglementation (situation au 27/09/2026 — à vérifier auprès du régulateur local avant tout projet) :**
| Zone | Statut 6 GHz |
|---|---|
| **Union européenne** | Bande **5945–6425 MHz** ouverte au Wi-Fi en mode **LPI** (Low Power Indoor — intérieur, faible puissance) et VLP par décision UE 2021/1067. La partie haute **6425–7125 MHz** : statut en discussion, usage Wi-Fi non généralisé. |
| **France (ARCEP)** | Suit le cadre UE : 6 GHz cantonné à l'**intérieur** en faible puissance. **Pas d'usage extérieur** en 6 GHz dans le cadre actuel. |
| **USA (FCC)** | 6 GHz ouvert en intérieur ET extérieur (avec AFC pour l'extérieur en standard power). |
| **Reste du monde** | Très variable : certains pays n'ont pas ouvert le 6 GHz du tout. |

**Conséquence directe pour l'AP761 et l'outdoor en Europe :**
1. L'AP761 n'a pas de radio 6 GHz de toute façon — le sujet est clos pour lui.
2. Même les AP eKit Wi-Fi 7 **outdoor** (AP771/AP772E) sont bi-bande 2.4/5 GHz : les constructeurs ne mettent pas de 6 GHz sur de l'outdoor destiné au marché UE, parce que la réglementation ne le permet pas (ou pas encore).
3. **Ne jamais activer/importer un AP 6 GHz configuré en région US pour l'utiliser en France** : non-conformité réglementaire + risque de brouiller des services prioritaires (liaisons fixes, radioastronomie selon les sous-bandes).

**À vérifier sur la fiche du modèle exact** : le « Country Code & Channel Compliance Table » de Huawei liste les canaux autorisés par pays — c'est LA référence avant d'activer une bande ou une largeur de canal.

## 20. Ce que l'AP761 ne fait PAS (tableau Wi-Fi 6 vs Wi-Fi 7)

Pour couper court aux confusions commerciales, le tableau à garder sous la main :

