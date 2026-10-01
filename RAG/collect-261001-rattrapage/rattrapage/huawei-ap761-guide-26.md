---
id: collect-261001-rattrapage/rattrapage/huawei-ap761-guide-26
title: "Guide ultra-complet — Huawei eKit AP761"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["incident"]
source: docs/RAG/collect-261001-rattrapage/huawei_ap761_guide.md
source_anchor: ""
source_lines: [2669, 2746]
sha256: c48907ea7c305b42d11ce73df24a7bb481d6e50d79a9e4392735456462d8bbef
---

# Guide ultra-complet — Huawei eKit AP761

**R9.** **Classifier avant de contre-attaquer :** (1) noter BSSID/SSID/canal/RSSI ; (2) chercher sa MAC dans les tables MAC des switches — **sur le filaire = interne** (remonter au port, aller voir : 90 % = équipement personnel) ; (3) si externe : usurpe-t-il ton SSID (evil twin → contre-mesure ciblée + logs + alerte direction) ou est-ce un voisin légitime (**ne rien faire d'agressif**, optimiser tes canaux) ; (4) noter l'incident dans le dossier de site.

**R10.** Réponse type : **(1) Besoin :** le Wi-Fi 7 apporte MLO (latence), 4K-QAM et efficacité — utile si on a des besoins de **densité/latence** que le Wi-Fi 6 ne satisfait pas ; sinon, un Wi-Fi 6 bien déployé suffit. **(2) Parc :** sans clients Wi-Fi 7, un AP Wi-Fi 7 = un AP Wi-Fi 6 cher — inventorier d'abord. **(3) Calendrier :** les AP761 ont 5–7 ans de vie devant eux ; la migration naturelle se fera au renouvellement (2029–2032), avec des **AP772E** (6.45 Gbps, 2.5G/10G, PoE bt) sur les zones denses et conservation des AP761 en couverture — en refaisant les budgets PoE et câblage.

**Barème indicatif :** 8+/10 = tu peux déployer et dépanner en autonomie. 5–7 = relis les chapitres de tes erreurs. < 5 = reprends le guide depuis le chapitre 12.

## 115. Cas pratique 1 : cour d'école / cour d'entreprise

**Contexte :** cour de 50 × 35 m, 150 élèves/salariés à la pause, usage = smartphones (réseaux sociaux, visio), 2 bâtiments en bordure.

**Solution :**
- **2× AP761** en angles opposés de la cour, secteurs pointés vers le centre (recouvrement 15–20 % à −67 dBm au milieu).
- Hauteur 5–6 m sur façades, downtilt léger.
- **Radio :** 2.4 GHz en 20 MHz (canaux 1 et 6), 5 GHz en 40 MHz (canaux 36 et 44 — non-DFS, pas de surprise à la récré).
- **SSID :** `ECOLE` (WPA2-PSK ou 802.1X selon le cas) + `ECOLE-INVITE` (portail, isolé). Pas de SSID « profs cachés » en plus — 2 SSID suffisent.
- **Sécurité :** débits de base relevés (pas de zombies à −85 dBm au fond de la cour), isolation client sur l'invité.
- **Capacité :** 150 clients / 2 AP = 75/AP en pic — dans la fourchette « navigation » (chap. 59). Si visio massive : passer à 3 AP.
- **PoE :** 2× 25 W = 50 W réservés sur le switch.

**Point de vigilance :** les pics sont **synchronisés** (la cloche sonne → 150 connexions en 2 minutes). Le DHCP doit encaisser : baux courts (2 h), scope dimensionné (un /23 si besoin, pas un /24 plein à 80 % en temps normal).

## 116. Cas pratique 2 : parking et contrôle d'accès

**Contexte :** parking de 120 × 40 m, barrières avec lecteurs de plaques connectés en Wi-Fi, 2 bornes de recharge VE connectées, smartphones des usagers.

**Solution :**
- **2–3× AP761** le long du parking, secteurs orientés dans l'axe (la forme allongée se prête aux secteurs directionnels).
- **SSID IOT** dédié en 2.4 GHz (canal 11, 20 MHz) pour les lecteurs de plaques et bornes : WPA2-PSK fort, VLAN 40 restreint (chap. 52), filtrage MAC en garde-fou (chap. 50).
- **SSID usagers** en 5 GHz pour les smartphones.
- **Criticité :** le contrôle d'accès ne doit pas tomber si le cloud est injoignable → vérifier le comportement local (le Wi-Fi continue, chap. 30), et **ne pas mettre les lecteurs sur un canal DFS** (une évacuation radar = barrières aveugles pendant la bascule).
- **Alimentation :** les lecteurs et l'AP sur le même onduleur ? **Décider** : si le parking doit rester contrôlé pendant une coupure, l'AP qui porte les lecteurs est critique (chap. 25).

**Point de vigilance :** les **carrosseries métalliques** font des réflexions et des zones d'ombre mouvantes (les voitures bougent). Le survey doit se faire **parking plein**, pas un dimanche vide — sinon les mesures sont optimistes.

## 117. Cas pratique 3 : entrepôt / zone logistique

**Contexte :** zone de chargement extérieure 80 × 50 m, douchettes/scanners Wi-Fi des caristes, 1 AP existant qui « ne passe pas » au fond.

**Solution :**
- **2× AP761** : un à chaque extrémité, secteurs vers le centre. Les scanners ont besoin de **roaming propre** (chap. 57) : 802.11r, même SSID/sécurité/VLAN.
- **Priorité à la fiabilité** : 5 GHz en 40 MHz sur canaux 36–48 (non-DFS), 2.4 GHz en 20 MHz pour les vieux scanners.
- **Scanners anciens** : beaucoup sont en Wi-Fi 4, voire avec des pilotes exotiques → **tester le modèle exact** avant de généraliser, prévoir un SSID dédié en WPA2-PSK si le 802.1X les fait trébucher.
- **Environnement :** racks métalliques = réflexions ; poussière = contrôle visuel semestriel renforcé (l'IP68 protège, mais les connecteurs et fixations s'encrassent).
- **Chariots élévateurs** : prévoir la **hauteur** (un AP à 4 m au-dessus d'une allée de chariots = collision possible) et les **vibrations** (serrage contrôlé).

**Point de vigilance :** les scanners envoient peu de données mais **n'acceptent pas les coupures** (une coupure = une palette non scannée = un litige). Le test de référence, c'est le **ping continu en chariot** sur tout le parcours (chap. 57), pas un speedtest.

## 118. Cas pratique 4 : terrasse d'hôtel / camping

**Contexte :** terrasse 40 × 25 m + piscine, 80 clients en été, usage = streaming vidéo, visio, réseaux sociaux. Équipement accessible au public.

**Solution :**
- **1–2× AP761** selon la forme (1 suffit en couverture, 2 pour la capacité en haute saison).
- **SSID `HOTEL`** (WPA2-PSK, clé changée chaque saison) + **portail captif** pour les visiteurs de passage (chap. 49) avec les mentions légales.
- **Isolation client** obligatoire (chap. 53), **limitation de débit par client** (10–15 Mbps : largement assez pour du streaming, ça protège des accapareurs).
- **Sécurité physique** : AP hors de portée (> 3 m), câble antivol, fixation vérifiée (chap. 108) — un AP arraché = un AP à remplacer + une config à révoquer.
- **Saisonnalité** : en hiver, baisser la puissance ou éteindre le SSID la nuit (économies + moins de pollution pour les voisins).

**Point de vigilance :** l'**eau** (piscine, arrosage) + le **soleil** (surchauffe, chap. 91) + le **sel** si bord de mer (corrosion, chap. 102). La visite semestrielle est **non négociable** sur ce type de site.

## 119. Cas pratique 5 : relais entre deux bâtiments (point à point)

**Contexte :** deux bâtiments distants de 150 m, pas de fourreau possible (route entre les deux), besoin de relier le réseau du bâtiment B.

**Solution honnête d'abord :** l'AP761 **n'est pas une antenne point-à-point dédiée** — ses antennes sont des secteurs 65°, pas des faisceaux étroits. Pour un vrai bridge 150 m, une paire d'antennes directionnelles dédiées (ou de la fibre aérienne) est plus propre. **Mais** si c'est du dépannage temporaire ou un faible débit :

- **2× AP761** en vis-à-vis, chacun pointé vers l'autre (lobe principal dans l'axe).
- **SSID dédié** au bridge, WPA2-PSK fort (ou WPA3-SAE), **masqué** (chap. 43 — ici justifié : ce n'est pas un SSID public).
- **VLAN trunké** sur le lien pour faire transiter plusieurs réseaux.
- **Débit attendu :** correct en 5 GHz 40 MHz à 150 m en vue directe (100–200 Mbps utiles), mais **sans garantie** : la pluie, le brouillard et les oiseaux sur l'axe dégradent.

**Limites à annoncer clairement :**
- Pas de redondance : si un AP tombe, le lien tombe.
- La latence/gigue est celle du Wi-Fi partagé, pas d'une fibre : **ne pas y faire passer de la téléphonie critique** sans test.
- **Solution pérenne = fibre** (aérienne sur poteaux ou forage dirigé sous la route). Le bridge Wi-Fi est un pansement, pas une artère.

## 120. Pour aller plus loin : docs, outils, formations

