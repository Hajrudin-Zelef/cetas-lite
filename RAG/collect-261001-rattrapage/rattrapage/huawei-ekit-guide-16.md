---
id: collect-261001-rattrapage/rattrapage/huawei-ekit-guide-16
title: "GUIDE ULTRA-COMPLET — PLATEFORME HUAWEI eKit (PME / DISTRIBUTION)"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["distribution", "attention", "omni"]
source: docs/RAG/collect-261001-rattrapage/huawei_ekit_guide.md
source_anchor: ""
source_lines: [1395, 1476]
sha256: 8b185ff8b7be0504ecc2ef9dc7d6540125e296d23a19a7702715c38ba731711e
---

# GUIDE ULTRA-COMPLET — PLATEFORME HUAWEI eKit (PME / DISTRIBUTION)

- **Positionnement :** AP intérieur **Wi-Fi 7 haute densité**, pensé pour salles de classe et grandes salles de réunion (annonce MWC 2025).
- **Cas d'usage :** écoles, amphis, open spaces denses, salles de conférence.
- **À vérifier sur la documentation officielle :** specs radio exactes (flux, débit), PoE (probablement at/bt), prix et disponibilité — modèle récent.
- **Conseil :** ne le déploie pas « parce que c'est du Wi-Fi 7 » dans un bureau de 5 personnes — l'AP361 suffit et coûte moins cher. Le Wi-Fi 7 paie en **densité**, pas en prestige.

## 120. Fiche terrain — AP673E (Wi-Fi 7 haute densité, 2e option)

- **Positionnement :** Wi-Fi 7 intérieur haute densité, cité aux côtés de l'AP572 pour les mêmes usages (classes, conférences).
- **À vérifier sur la documentation officielle :** différences exactes avec l'AP572 (positionnement, prix, specs) — les deux modèles coexistent dans les annonces 2025.
- **Réflexe commercial :** demande au distributeur lequel est **stocké localement** — un modèle légèrement moins bon mais dispo en 48 h bat un modèle parfait en 6 semaines.

## 121. Fiche terrain — AP673H iGuard (le confidentiel)

- **Positionnement :** AP intérieur **Wi-Fi 7 tri-bande** (2,4 GHz 2×2 + 5 GHz 2×2 + **6 GHz 4×4**), 8 flux spatiaux, **13,66 Gbit/s** théoriques.
- **Différenciant :** **détection de caméras espion** par IA (analyse des ondes, y compris caméras non connectées au Wi-Fi) — voir section 29.
- **Contrainte officielle :** détection spycam en **gestion locale uniquement** (pas de cloud pour cette fonction).
- **Cas d'usage :** hôtels haut de gamme, salles de direction, cabinets sensibles, résidences VIP.
- **À vérifier :** disponibilité de la bande 6 GHz dans ton pays (réglementation), prix (positionné premium), et la procédure exacte de scan/rapport.

## 122. Fiche terrain — AP772 / AP772E (l'extérieur Wi-Fi 7)

- **AP772 :** extérieur Wi-Fi 7, antenne **omnidirectionnelle 360°** — couverture large (place, cour).
- **AP772E :** extérieur Wi-Fi 7, antenne **directionnelle** — couverture longue distance ciblée (allée, parking en longueur, liaison entre bâtiments proches).
- **À vérifier sur la documentation officielle :** specs détaillées, PoE, indice de protection, prix — annoncés en 2025.
- **Règle :** directionnel = on vise une zone précise ; omnidirectionnel = on couvre autour. Ne mets pas un omni là où il faut un directionnel « parce qu'il est moins cher » : tu arroseras les voisins et pas la cible.

## 123. Fiche terrain — AR180 / AR180 Plus / AR180 Pro

- **Concept :** passerelle tout-en-un (routage + switch + Wi-Fi + gestion Internet), 8 fonctions intégrées d'après Huawei.
- **AR180 / Plus :** 4× GE LAN + 1× 2,5GE WAN, Wi-Fi 7 double bande, ~100 terminaux, 8 tunnels IPsec, 200 Mbit/s VPN, 2 Gbit/s de contrôle de bande passante, **sans ventilateur**, pas de PoE, pas de port console.
- **AR180 Pro :** 5 ports commutables + 1× 2,5GE, ~150 terminaux, 16 tunnels IPsec, **8 AP max / 32 équipements gérés**, 256 VLAN, boîtier métal.
- **Antennes :** 4 externes intelligentes, non démontables, 23 dBm max (réglementation locale).
- **Cas d'usage :** AR180 = TPE/boutique/bureau < 15 personnes ; AR180 Pro = PME jusqu'à ~50 personnes avec quelques AP.
- **Limite dure :** 8 AP max sur le Pro — au-delà, passer à l'AR280.

## 124. Fiche terrain — AR280 (le chef d'orchestre PME)

- **Ports :** 4× GE (1 LAN dédié + 3 LAN/WAN configurables) + 1× 2,5GE (WAN/LAN configurable). **Pas de Wi-Fi intégré** d'après la fiche — le Wi-Fi passe par les AP.
- **Capacité :** ~150 terminaux, **16 AP max, 32 équipements gérés**, 16 tunnels IPsec, 200 Mbit/s VPN.
- **Gestion :** planification, déploiement, inspection et O&M via la plateforme cloud Huawei SME.
- **Cas d'usage :** la tête de réseau standard des déploiements 20-100 utilisateurs (bureau, hôtel, école).
- **Point d'attention :** la fiche mentionne une capacité PoE (41 W) — modèle ou variante à clarifier (**à vérifier sur la documentation officielle** : ne compte pas dessus pour alimenter des AP sans confirmation).

## 125. Fiche terrain — AR281 (le tout-en-un 2026)

- **Annonce 2026** (Intelligent Office Solution 2.0) : « un appareil = 5 fonctions » — routage, **NVR** (vidéosurveillance), **WAC** (contrôleur Wi-Fi), accès ligne privée, gestion du comportement en ligne.
- **Stockage vidéo :** compression « SuperCoding 3.0 » annoncée à -85 % (1 disque = 6× la capacité) — à valider en conditions réelles.
- **Cas d'usage visé :** petite structure voulant réseau + Wi-Fi + vidéosurveillance dans un seul boîtier.
- **À vérifier sur la documentation officielle :** TOUT (specs, prix, dispo) — produit très récent au moment de la rédaction. Ne le chiffre pas sans fiche en main.

## 126. Fiche terrain — S220-8P4S (le petit switch qui dépanne)

- **Ports :** 8× GE + 4× SFP 1G, **125 W PoE+** (802.3af/at), 24 Gbit/s, 18 Mpps.
- **Fonctions :** Fast PoE, Perpetual PoE, SNMP v1/v2c/v3, CLI, web, SSH, DHCP snooping, isolation de ports, sticky MAC, 802.1X.
- **Environnement :** -5 à +50 °C, ventilation intelligente (silencieux ajusté).
- **Cas d'usage :** boutique, agence, petit bureau (8 ports PoE suffisent pour 2-3 AP + 3-4 caméras), switch de secours dans ton stock.
- **Limite :** 8 ports — dès 6-7 équipements PoE, passe au S310 24 ports pour la marge.

## 127. Fiche terrain — S310-48P4S (le cheval de labour)

- **Ports :** 48× GE PoE+ + 4× SFP 1G, **380 W PoE+**, 104 Gbit/s, 77 Mpps.
- **Gestion :** app eKit (cloud) **ou** locale (CLI/web/SNMP/SSH) — commutable.
- **Particularités :** HOUP (MAJ intelligente en 1 clic), stacking zéro-config, protection ±7 kV / ±6 kV, MTBF ~48 ans.
- **Cas d'usage :** LE switch des hôtels, écoles, bureaux 30-100 utilisateurs.
- **Règle PoE :** 380 W = 24 ports à 15,4 W (af) ou 12 ports à 30 W (at) en pleine charge — **au-delà, le switch gère par priorité** : configure les priorités PoE (AP critiques d'abord).

## 128. Fiche terrain — S310-48P4X (le 10G)

- **Ports :** 48× GE PoE+ + **4× SFP+ 10G**, 380 W PoE+, 176 Gbit/s, 131 Mpps.
- **Différence avec le 48P4S :** les uplinks passent à 10 Gbit/s — pour agréger plusieurs switchs ou un serveur/NAS gourmand.
- **Cas d'usage :** tête de distribution d'un hôtel/école avec 2-3 switchs d'étage, ou site avec NAS de sauvegarde.
- **Surcoût vs 48P4S :** à chiffrer — si aucun besoin 10G à 3 ans, le 48P4S suffit. Si le client « grandit vite », le 48P4X est l'assurance (section 85).

## 129. Fiche terrain — S310 tout-optique (24J4X / 48J4X)

- **Ports :** 24 ou 48× **SFP 2,5GE** (tout-optique) + 4× SFP+ 10G. Capacités : 200 Gbit/s (24) / 320 Gbit/s (48).
- **Cas d'usage :** colonnes montantes en fibre (hôtel, campus), environnements perturbés électromagnétiquement, longues distances inter-bâtiments.
- **Contrainte :** 100 % fibre = 100 % modules SFP à acheter + compétence fibre (soudures, réflectométrie). Ne le propose que si tu maîtrises ou sous-traites la fibre.
- **Astuce :** en hôtel, une dorsale fibre + des S310-48P4S cuivre par étage = le meilleur des deux mondes.

## 130. Fiche terrain — USG6000F-S (le pare-feu IA)

