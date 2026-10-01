---
id: collect-261001-rattrapage/rattrapage/huawei-ap361-guide-15
title: "Huawei eKit AP361 — Guide ultra-complet"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["ethernet"]
source: docs/RAG/collect-261001-rattrapage/huawei_ap361_guide.md
source_anchor: ""
source_lines: [2365, 2519]
sha256: 5c1aa01a5334a9021202dc8a26134e2f389621d3e70c9afa2bd696fcbe198a3c
---

# Huawei eKit AP361 — Guide ultra-complet

| Rôle | Nom | Téléphone |
|---|---|---|
| Responsable IT | __________ | __________ |
| Astreinte réseau | __________ | __________ |
| Fournisseur / RMA Huawei | __________ | __________ |
| Électricien (baie) | __________ | __________ |
| Gardien / accès bâtiment | __________ | __________ |

## 155. Check-list « jour J » d'une intervention (à imprimer)

- [ ] Sac : testeur câble + testeur PoE + trombone + AP spare + étiqueteuse
- [ ] Téléphone chargé (app eKit + analyseur Wi-Fi)
- [ ] Accès : badge, clé du local, autorisation
- [ ] Consignation si intervention électrique en baie
- [ ] Photo « avant », photo « après »
- [ ] Test client complet après chaque modif (§31)
- [ ] Note au journal d'exploitation

## 156. Phrases à dire (et à ne pas dire)

- ✅ « Le signal est bon mais le débit est mauvais : je regarde le câble et les
      interférences. » (Ne confondez jamais signal et débit.)
- ✅ « Je change une variable à la fois et je note. »
- ⛔ « J'ai mis la puissance à fond, ça devrait mieux passer. » (Non.)
- ⛔ « C'est le Wi-Fi qui rame. » (Lequel ? Quel AP ? Quel client ? Quelle heure ?
      Mesurez d'abord.)

---
---

# R. GLOSSAIRE

## 157. Glossaire (les termes qui reviennent tout le temps)

| Terme | Définition courte |
|---|---|
| **802.11ax** | Norme Wi-Fi 6 : OFDMA, MU-MIMO montant/descendant, 1024-QAM, BSS Coloring |
| **802.11k** | Le client reçoit la liste des AP voisins → roaming plus rapide |
| **802.11v** | L'AP peut suggérer au client de changer d'AP (BSS Transition) |
| **802.11r** | Ré-authentification rapide lors du roaming (Fast BSS Transition) |
| **802.11w (PMF)** | Protection des trames de management contre la falsification |
| **AC (contrôleur)** | Équipement central pilotant des AP en mode Fit (pas le positionnement eKit) |
| **Airtime fairness** | Partage équitable du temps d'antenne entre clients |
| **AP** | Access Point — point d'accès |
| **Band steering** | Mécanisme poussant les clients vers le 5 GHz |
| **Beacon** | Trame émise ~10x/s par SSID annonçant le réseau |
| **BSS Coloring** | Marquage des trames (Wi-Fi 6) pour ignorer les transmissions voisines |
| **Captive portal** | Portail captif : page d'authentification avant accès Internet |
| **DCA** | Dynamic Channel Assignment — choix automatique du canal |
| **DFS** | Dynamic Frequency Selection — changement de canal obligatoire si radar détecté |
| **DSCP** | Marquage de QoS dans l'en-tête IP |
| **EAP** | Extensible Authentication Protocol — cadre d'authentification 802.1X |
| **EIRP** | Puissance isotrope rayonnée équivalente (puissance + gain d'antenne) |
| **Evil twin** | Faux AP diffusant votre SSID pour piéger les clients |
| **Fat / Fit** | Mode autonome / mode piloté par contrôleur |
| **GE** | Gigabit Ethernet (10/100/1000) |
| **MIMO** | Multiple-Input Multiple-Output — plusieurs flux spatiaux simultanés |
| **MU-MIMO** | MIMO multi-utilisateurs : l'AP parle à plusieurs clients à la fois |
| **OFDMA** | Découpage du canal en sous-porteuses allouées à plusieurs clients (Wi-Fi 6) |
| **OKC** | Opportunistic Key Caching — réutilisation de clé au roaming |
| **PVID** | Port VLAN ID — VLAN natif (non tagué) d'un port trunk |
| **QAM** | Modulation d'amplitude en quadrature (1024-QAM = plus de bits/symbole) |
| **RADIUS** | Serveur d'authentification (802.1X, WPA-Enterprise) |
| **RMA** | Return Merchandise Authorization — retour garantie |
| **Roaming** | Passage d'un client d'un AP à un autre en se déplaçant |
| **Rogue AP** | Point d'accès non autorisé |
| **RSSI** | Received Signal Strength Indicator — niveau de signal (dBm, négatif) |
| **SAE** | Simultaneous Authentication of Equals — handshake WPA3-Personal |
| **SSID** | Nom du réseau Wi-Fi diffusé |
| **Sticky client** | Client qui reste accroché à un AP lointain |
| **TPC** | Transmit Power Control — ajustement automatique de puissance |
| **Trunk** | Port transportant plusieurs VLAN tagués |
| **VAP** | Virtual AP — instance de SSID sur une radio |
| **VLAN** | Réseau local virtuel — séparation logique au niveau 2 |
| **WIDS / wIPS** | Détection / prévention d'intrusion sans fil |
| **WMM** | Wi-Fi Multimedia — QoS sans fil (4 files) |
| **WPA2 / WPA3** | Protocoles de sécurité Wi-Fi (PSK = clé partagée, EAP = entreprise) |

---
---

# S. QUIZ — 10 QUESTIONS

## 158. Questions (réfléchissez avant de regarder les réponses)

**Q1.** L'AP361 est alimenté en 802.3af et consomme 8,8 W max. Vous déployez
12 AP361 sur un switch PoE. Quel budget PoE minimum prévoyez-vous (avec 20 % de
marge), et quel palier commercial choisissez-vous ?

**Q2.** En 2,4 GHz, quels canaux utilisez-vous en planifié, et quelle largeur de
canal ? Pourquoi pas autre chose ?

**Q3.** Un client est associé au SSID Bureau mais obtient une adresse en
169.254.x.x. Citez 3 causes possibles, dans l'ordre de vérification.

**Q4.** Pourquoi ne faut-il pas mettre la puissance d'émission « à fond » ?
Citez au moins 3 effets négatifs.

**Q5.** Un utilisateur se plaint que ses appels Wi-Fi coupent quand il marche
dans les couloirs. Quelles sont les 3 vérifications prioritaires ?

**Q6.** Vous activez l'isolation client sur le SSID Bureau un vendredi soir.
Que risque-t-il de se passer lundi matin ? Sur quels SSID l'isolation est-elle
recommandée ?

**Q7.** Quelle est la différence entre 802.11k, 802.11v et 802.11r ? Lequel est
indispensable pour la voix sur Wi-Fi ?

**Q8.** Un AP est « en ligne » dans eKit mais aucun client ne s'y connecte alors
que l'AP voisin est saturé. Citez 4 causes possibles.

**Q9.** Pendant un upgrade firmware, que ne faut-il jamais faire ? Citez la
procédure de déploiement recommandée sur un parc (ordre des étapes).

**Q10.** Votre switch a un budget PoE de 185 W. Pouvez-vous y brancher 12 AP361
+ 4 caméras consommant 12 W chacune ? Détaillez le calcul.

## 159. Réponses

**R1.** 12 × 10 W (arrondi avec marge unitaire) = 120 W ; × 1,2 = **144 W
minimum**. Palier commercial : **185 W** (le palier 130-150 W est trop juste).
Et on vérifie que le budget PoE est bien la valeur « power budget », pas
24 ports × 15,4 W.

**R2.** Canaux **1, 6, 11** uniquement (les seuls non chevauchants en 20 MHz en
Europe), largeur **20 MHz toujours**. Le 40 MHz en 2,4 GHz occupe deux tiers de
la bande et génère des interférences avec les voisins ; tout autre canal
chevauche 1/6/11.

**R3.** (1) Le VLAN du SSID n'est pas autorisé sur le trunk du port AP
(`allow-pass`) ; (2) pas de serveur DHCP sur ce VLAN / relai DHCP oublié ;
(3) scope DHCP plein (baux trop longs). Ordre : d'abord tester en filaire sur le
même VLAN pour isoler Wi-Fi vs réseau.

**R4.** (1) L'AP arrose les voisins et l'étage → interférences ; (2) liaison
asymétrique : le client entend l'AP mais l'AP n'entend pas le client →
pertes ; (3) sticky clients : les terminaux restent accrochés de loin au lieu
de roamer. Bonus : gaspillage énergétique.

**R5.** (1) Couverture : mesurer en marchant, visez **-67 dBm minimum** partout ;
(2) **802.11r activé** (sans lui, ré-authentification lente à chaque changement
d'AP) ; (3) QoS voix (marquage EF + WMM) pour que l'appel ne soit pas noyé.
Vérifier aussi le sticky client (puissances équilibrées).

**R6.** Lundi : « l'imprimante réseau ne marche plus », « je ne vois plus les
PC du bureau » — l'isolation bloque les communications **entre clients** du
même SSID (partage Windows, impression locale, Chromecast). Elle est
recommandée sur **Invité** (toujours) et **IoT**, pas sur Bureau sans analyse
des usages.

**R7.** **k** : le client reçoit la liste des AP voisins (scan plus rapide) ;
**v** : l'AP suggère au client de migrer (BSS Transition) ; **r** : la
ré-authentification est accélérée lors du changement d'AP. Pour la voix,
c'est le **r (802.11r)** qui est indispensable (coupure < 100 ms).

