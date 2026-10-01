---
id: collect-261001-rattrapage/rattrapage/huawei-ekit-guide-15
title: "GUIDE ULTRA-COMPLET — PLATEFORME HUAWEI eKit (PME / DISTRIBUTION)"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["capex"]
source: docs/RAG/collect-261001-rattrapage/huawei_ekit_guide.md
source_anchor: ""
source_lines: [1321, 1394]
sha256: 8893fdabe7b9b89b3364aa0e9f13c198a65dfd00964a4f7c98a36c3b1aec8343
---

# GUIDE ULTRA-COMPLET — PLATEFORME HUAWEI eKit (PME / DISTRIBUTION)

1. **802.3at (PoE+) recommandé** : la fiche technique précise qu'en 802.3af certaines fonctions peuvent être limitées. Avec 17,7 W max, on est au-delà des 15,4 W du af.
2. Calcul : 20×9 = 180 W ; 2×17,7 = 35,4 W ; 8×10 = 80 W → **295,4 W**. Avec 30 % de marge : 295,4 × 1,3 = **384 W > 380 W** : ça ne tient pas avec la marge. Il faut répartir sur 2 switchs ou réduire la charge.
3. **En amont des équipements** : alimentation du site, lien opérateur (box/ONT), passerelle AR. Le LAN local fonctionne → le problème est l'uplink Internet (ou le lien vers le cloud), pas les AP/switchs.
4. **Fit** (piloté par WAC — ex. AR280/AR281 en tête), **Fat** (autonome — petit site ou dépannage), **Cloud** (piloté par app/SNC — le standard eKit multi-sites).
5. Parce que c'est la **plage par défaut des box grand public** : dès qu'on branche une box, un routeur 4G ou un VPN, les conflits d'adressage sont garantis.
6. Chaque SSID consomme de l'airtime (beacons) et dégrade tout le monde : on propose **3 SSID max** (staff, invités, + 1 technique si vraiment nécessaire) et on fait le tri des usages réels avec le client.
7. (1) Identifier quels équipements ont été mis à jour et lesquels sont en panne ; (2) comparer la config actuelle à la **sauvegarde d'avant-MAJ** ; (3) restaurer/downgrader vers la version précédente stable. Et noter la leçon : site pilote + 48 h la prochaine fois.
8. Exemples : > 30-50 AP avec roaming fin ; besoin de 802.1X/NAC ; routage dynamique (OSPF/BGP) ; redondance matérielle contractuelle ; débits inter-VLAN multi-Gbit/s permanents ; équipe réseau voulant CLI/API/automatisation.
9. Pour des **raisons de confidentialité** (fiche constructeur) : la détection spycam fonctionne en **gestion locale uniquement**, les données ne partent pas dans le cloud.
10. **(a) Cuivre classique** : AR280/USG + switchs S310 PoE + AP361 (1 pour 2 chambres) — moins cher, s'appuie sur un câblage cuivre standard. **(b) MiniFTTO eKitOptix** : OLT + 1 F700D par chambre en fibre — plus cher à l'achat, esthétique et évolutivité supérieures, idéal en neuf. Arbitrage : CAPEX vs OPEX, esthétique, compétence fibre disponible.

## 113. Pour aller plus loin

- **Documentation officielle** : ekit.huawei.com (portail), support.huawei.com (fiches techniques, guides d'installation, release notes — à consulter **avant chaque déploiement** pour le modèle exact).
- **Vidéos constructeur** : le distributeur eKit local publie souvent des tutoriels d'onboarding via l'app — demande à ton Gold Partner ses liens.
- **Formations** : renseigne-toi auprès de ton distributeur sur les formations/certifications eKit (argument commercial + montée en compétence de ton équipe).
- **Communauté** : forums Huawei Enterprise (e.huawei.com) — retours d'autres intégrateurs sur les firmwares et les cas tordus.
- **Dans ton workspace** : croise ce guide avec `onduleurs_ups_guide.md` (dimensionnement électrique des baies), `debian_ubuntu_guide.md` / `proxmox_guide.md` (si tu héberges un NMS), et les guides `*_guide.md` réseau pour la supervision complémentaire.
- **Prochaine étape terrain** : monte une **maquette** (1 AR + 1 switch + 2 AP) dans ton atelier et déroule les sections 19-24 : c'est en onboardant pour de vrai que l'app n'aura plus de secret.

## 114. Sources et avertissement de fraîcheur

Informations constructeur vérifiées par recherche web (septembre 2026) :
- Fiches techniques **eKitEngine AR180 / AR280** (hardware description, datasheets ES — via e-file.huawei.com et distributeurs).
- Fiches techniques **AP361, AP761, AP266, AP673H (iGuard)** (datasheets constructeur et distributeurs).
- Fiches techniques **S310-48P4S / 48P4X / 24J4X / 48J4X, S220-8P4S, S620** (distributeurs, datasheet S310).
- Fiche technique **USG6000F-S** (S125/S150/S200, V600R025).
- Communiqués Huawei **MWC 2025** (20+ produits eKit, SNC, troubleshooting 2.0) et **Intelligent Office Solution 2.0** (2026 : AR281, DF10).
- Portail **ekit.huawei.com** (cité dans les fiches).
**Avertissement :** la gamme eKit évolue vite (nouveautés 2025-2026 : AR281, AP772/AP772E, F700D, USG6000F-S). Les références, les fonctions de l'app/SNC et les conditions de licence peuvent changer : **revérifie la fiche exacte du modèle avant chaque chiffrage**, et traite tout « à vérifier sur la documentation officielle » de ce guide comme une action, pas comme une formule.

## 115. Note de l'auteur : comment utiliser ce guide sur le terrain

1. **Avant-vente** : sections 1-15 (périmètre) + 83 (grille de décision) + un scénario (46-60) = ta proposition technique.
2. **Jour J** : fiche site (65) + pense-bêtes (107-109) imprimés + check-list (74).
3. **Exploitation** : rituel hebdo (43), plan annuel (101), KPI mensuels (80).
4. **Panne** : méthode (91) + cas (92-100) + pense-bête dépannage (108).
5. **Fin de contrat** : transfert documenté (103) + registre (106) à jour.
Ce guide vit : annote-le à chaque déploiement (un cas tordu = un cas n°17 à ajouter). Dans un an, il vaudra de l'or.

---

## 116. Fiche terrain — AP361 (l'ouvrier de base)

- **Positionnement :** AP intérieur Wi-Fi 6 d'entrée de gamme, le plus déployé en PME.
- **Radio :** 2,4 GHz 2×2 (575 Mbit/s) + 5 GHz 2×2 (1,2 Gbit/s), 1,775 Gbit/s agrégés, MU-MIMO, OFDMA, 1024-QAM.
- **Antenne :** intelligente intégrée (ajustement auto de la couverture) — pas d'antenne externe.
- **Ports :** 1× GE (données + PoE-in 802.3af). Pas de 2e port sur certaines déclinaisons (la fiche de base cite 1× GE ; une variante cite uplink + downlink — **à vérifier sur la documentation officielle** du modèle exact).
- **Alimentation :** PoE 802.3af, ~8,8-9,4 W. Un injecteur PoE fait l'affaire pour un AP isolé.
- **Modes :** Fit, Fat, Cloud. **Montage :** plafond ou mural (kit inclus selon distributeur).
- **Bonus :** port USB pour extension IoT (ZigBee/RFID selon fiche), Bluetooth pour O&M.
- **Cas d'usage :** bureaux, chambres d'hôtel, boutiques, écoles, cliniques.
- **Limite :** 1 Gbit/s d'uplink — ne pas le mettre là où il faudrait du 2,5G (forte densité Wi-Fi 7).
- **Prix indicatif :** le moins cher de la gamme AP eKit — **à vérifier auprès du distributeur** (les prix varient fortement par pays).

## 117. Fiche terrain — AP761 (l'extérieur)

- **Positionnement :** AP extérieur Wi-Fi 6 (AX1800 : 574 + 1201 Mbit/s).
- **Ports :** 1× GE PoE-in + **1× SFP** (fibre — parfait pour les longues distances : parking, portail, dépendance).
- **Alimentation :** 802.3at **recommandé** (en af, fonctions limitées), 17,7 W max.
- **Particularités :** BLE 5.2 (gestion, localisation, IoT), MU-MIMO, Beamforming, TWT.
- **Installation :** mât ou mur, **mise à la terre + parafoudre** si exposé (section 72), presse-étoupes bien serrés (l'eau ne pardonne pas).
- **Cas d'usage :** terrasses, piscines, parkings, cours d'école, quais semi-ouverts.
- **Piège :** l'oublier dans le plan de nommage (« l'AP dehors » n'est pas un nom). Le nommer `SITE-EXT-AP01` dès l'onboarding.

## 118. Fiche terrain — AP266 (le simple et efficace)

- **Positionnement :** AP intérieur Wi-Fi 6, gestion cloud, **authentification intégrée** (PSK + portail) sans WAC ni serveur externe.
- **Déploiement :** par Wi-Fi (connexion au Wi-Fi de gestion) ou par scan du SN — les deux modes documentés par la fiche.
- **Cas d'usage :** quand le client veut un portail captif simple sans ajouter d'USG : le couple AP266 + cloud suffit.
- **À vérifier sur la documentation officielle :** débit exact, nombre de ports, puissance PoE — la fiche distributeur reste sommaire.

## 119. Fiche terrain — AP572 (la haute densité Wi-Fi 7)

