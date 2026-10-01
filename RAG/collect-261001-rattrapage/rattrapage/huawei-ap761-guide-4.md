---
id: collect-261001-rattrapage/rattrapage/huawei-ap761-guide-4
title: "Guide ultra-complet — Huawei eKit AP761"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["attention", "energy"]
source: docs/RAG/collect-261001-rattrapage/huawei_ap761_guide.md
source_anchor: ""
source_lines: [291, 374]
sha256: 53ed962ebacae82b5657f5bf5d954de538a8edccd1e874d81395b313f19dc33b
---

# Guide ultra-complet — Huawei eKit AP761

**Le point critique : ce sont des ports COMBO.** Un seul des deux est actif à la fois. **Si les deux sont branchés, le port optique (SFP) est prioritaire** [constructeur]. Conséquences terrain :

1. Tu ne peux pas faire « cuivre = uplink, fibre = downlink » : c'est l'un OU l'autre.
2. Si tu branches la fibre pour un test sans débrancher le cuivre, l'AP bascule sur la fibre. Si le SFP est mal enfiché, tu perds l'AP au lieu de tester.
3. **L'alimentation arrive TOUJOURS par le RJ45** (PoE-In). La fibre ne transporte pas d'énergie : un AP uplinké en fibre a quand même besoin du cuivre pour le PoE (ou d'une alimentation locale — à vérifier sur la fiche du modèle exact si un adaptateur secteur est supporté).

**Scénarios de câblage :**

| Scénario | Câblage | Remarque |
|---|---|---|
| Classique PME | RJ45 seul vers switch PoE | Le plus courant. 100 m max en cuivre. |
| Distance > 100 m | Fibre via SFP + RJ45 vers injecteur PoE local | Le SFP prend l'uplink, le RJ45 ne sert qu'au PoE. |
| Liaison inter-bâtiments | Fibre via SFP + PoE local | Idéal avec parafoudre des deux côtés. |

**Négociation :** auto-négociation du débit et du duplex, MDI/MDI-X automatique [constructeur] — un câble droit suffit dans tous les cas.

> ⚠️ **Piège classique :** brancher le SFP « au cas où » pendant que l'uplink cuivre fonctionne. L'AP bascule sur l'optique, et si le module SFP n'est pas compatible ou la fibre pas encore brassée, l'AP disparaît du réseau. **Ne branche le SFP que quand la fibre est prête.**

## 7. Hardware : alimentation PoE — 802.3at/af, 17.7 W

**Valeurs constructeur :** PoE **IEEE 802.3at/af**, consommation max **17.7 W**. Point capital du datasheet : **en mode 802.3af, l'AP fonctionne avec des fonctions restreintes** [constructeur].

**Traduction terrain :**
- **802.3af** = 15.4 W au port du switch (12.95 W utiles à l'AP après pertes câble). C'est **en dessous** des 17.7 W max de l'AP. Résultat : l'AP démarre mais bride des fonctions (typiquement : puissance radio réduite, USB/BLE désactivés, un débit plafonné — le détail exact des restrictions est **à vérifier sur la fiche du modèle exact**).
- **802.3at (PoE+)** = 30 W au port (25.5 W utiles). C'est le **minimum recommandé** : l'AP tourne à plein régime.
- **802.3bt** : non requis par l'AP761 (pas de 2.5G, pas de radios Wi-Fi 7 gourmandes). Un switch 802.3bt alimente très bien l'AP761 en négociant en 802.3at — c'est rétrocompatible.

**Règle d'or du budget PoE :**
> Budget par port = 17.7 W (AP) + pertes câble (~2.5 W pour 100 m en PoE+) + marge 20 % ≈ **25 W réservés par AP761** sur le budget du switch.

**Exemple :** un switch 8 ports PoE+ avec budget total 130 W → 130 / 25 ≈ **5 AP761 max** en sécurité, pas 8. Le nombre de ports PoE ne fait pas le budget : c'est le **budget total en watts** qui compte. Toujours lire la ligne « PoE budget » de la fiche du switch.

**Injecteur PoE :** si le switch n'est pas PoE, un injecteur **802.3at 30 W** fait l'affaire. Vérifier qu'il est **gigabit** (certains injecteurs bas de gamme ne passent que du 100 Mbps et tu brides l'AP sans le savoir).

**Câble et PoE :** le PoE chauffe le câble. En extérieur au soleil + PoE, prends du **câble extérieur blindé (F/UTP ou S/FTP)**, âme **100 % cuivre** (pas de CCA — cuivre plaqué aluminium — qui chauffe et casse). Section 24 AWG minimum.

## 8. Hardware : LED d'état et bouton reset

**LED** [constructeur] : une LED unique (ou un ensemble selon la version) indique les états **power-on, démarrage, fonctionnement, alarme, défaut**. En pratique sur les AP Huawei eKit :

| État LED | Signification probable |
|---|---|
| Éteinte | Pas d'alimentation — vérifier PoE/injecteur |
| Rouge fixe | Démarrage / bootloader |
| Vert clignotant | Démarrage en cours, chargement du firmware |
| Vert fixe | Fonctionnement normal |
| Orange/rouge clignotant | Alarme ou défaut (perte cloud, erreur radio…) |

> Les codes exacts variant selon la version logicielle : **à vérifier sur la fiche du modèle exact** (guide de démarrage rapide fourni dans la boîte). En dépannage, la LED est ton premier indicateur à distance quand quelqu'un est sur place au téléphone.

**Bouton reset** (revendeurs) : appui long (≈ 10 s, à vérifier sur la fiche du modèle exact) = retour aux paramètres d'usine. **Attention :** sur un AP en mode cloud déjà adopté, un reset d'usine le fait perdre du site — il faudra le ré-adopter (chap. 29). Ne reset jamais sans avoir la configuration sauvegardée (chap. 72).

**Port de verrouillage** (revendeurs) : emplacement pour câble antivol type Kensington (câble vendu séparément). Utile sur les AP accessibles au public (terrasses, campings).

## 9. Hardware : antennes intégrées directionnelles

**Valeurs constructeur :** antennes intégrées **directionnelles** —
- 2.4 GHz : **10 dBi**, ouverture **65° horizontale / 40° verticale**
- 5 GHz : **11 dBi**, ouverture **65° horizontale / 20° verticale**
- BLE : 5 dBi

**Ce que « directionnel 65° » veut dire sur le terrain :**
- L'AP761 n'arrose pas à 360°. Il couvre un **secteur d'environ 65°** devant lui — comme un projecteur, pas comme une ampoule.
- C'est un choix assumé : en extérieur, on veut couvrir une **zone ciblée** (une cour, une terrasse, un parking) en portant loin, pas arroser les voisins et ramasser leurs interférences.
- **Conséquence n° 1 : l'orientation compte.** L'AP doit être pointé vers la zone à couvrir, avec un léger **tilt vers le bas** (downtilt) si la zone est en contrebas. Un AP directionnel monté à l'envers ou de travers = une zone morte là où tu voulais couvrir.
- **Conséquence n° 2 : l'arrière est sourd.** Ne compte pas sur l'AP761 pour couvrir derrière lui (le bâtiment sur lequel il est fixé, par exemple). Prévoir un AP indoor séparé pour l'intérieur.
- **Conséquence n° 3 : le gain élevé (10–11 dBi) concentre l'énergie.** À puissance radio égale, ça porte plus loin dans l'axe, mais la zone couverte est plus étroite. Pour une couverture large (place publique), il faut soit plusieurs AP761 en secteurs, soit un modèle omnidirectionnel — **à vérifier sur la fiche du modèle exact** s'il existe une variante.

**Hauteur de montage :** 4 à 8 m typique pour une cour/un parking. Plus haut = plus de portée mais zone morte juste en dessous (lobe vertical étroit, surtout en 5 GHz : 20°). Règle de pouce : **ne pas monter plus haut que nécessaire** — chaque mètre de hauteur en trop creuse la zone d'ombre au pied du mât.

## 10. Hardware : BLE 5.2 et USB IoT

**BLE 5.2** [constructeur] : le Bluetooth Low Energy intégré sert à l'**O&M via l'app mobile** (onboarding, diagnostic de proximité) sans avoir besoin d'être sur le réseau. En pratique :
- L'app HUAWEI eKit détecte l'AP en BLE pour l'adoption initiale (chap. 29).
- Ça dépanne quand l'AP n'a pas encore d'IP ou que le réseau est en vrac : tu te connectes en BLE à côté de l'AP et tu lis son état.

**USB** (source revendeur, **à vérifier sur la fiche du modèle exact**) : un port USB permettrait l'extension IoT (dongles ZigBee, RFID). Si ton usage est purement Wi-Fi, ignore-le. Si un jour tu veux de la localisation d'actifs ou des capteurs, c'est une piste — mais vérifie d'abord la liste des dongles supportés par la version logicielle.

**Point d'attention PoE :** en alimentation 802.3af (fonctions restreintes), l'USB et le BLE sont typiquement les premières fonctions coupées. Encore une raison de rester en 802.3at.

## 11. Les 3 modes de fonctionnement : Fat, Fit, Cloud

L'AP761 supporte trois modes [constructeur]. Le choix du mode conditionne TOUTE l'administration : choisis-le avant de configurer.

