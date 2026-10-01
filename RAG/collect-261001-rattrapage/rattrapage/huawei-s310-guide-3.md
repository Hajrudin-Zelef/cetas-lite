---
id: collect-261001-rattrapage/rattrapage/huawei-s310-guide-3
title: "Guide ultra-complet — Huawei eKit S310"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: ["arr"]
source: docs/RAG/collect-261001-rattrapage/huawei_s310_guide.md
source_anchor: ""
source_lines: [289, 457]
sha256: 793892df2481b1dcb573f2c074298d4b2380c5acd9776a3cac5a30576bf12a27
---

# Guide ultra-complet — Huawei eKit S310

1. Dégagement minimum : 10 cm à l'avant (câbles), 10 cm sur les côtés pour les
   modèles à flux latéral (48P4S : entrée d'air à gauche, sortie à droite).
2. Ne jamais empiler directement deux switches sans espace ni les coincer entre
   des classeurs : l'air doit circuler.
3. Température ambiante cible : **< 35 °C** en continu. Au-delà de 40 °C, les
   ventilateurs tournent à fond (bruit) et la durée de vie des condensateurs chute.
4. En baie : 1U de vide ou panneau de brassage ventilé entre deux switches PoE
   chargés si la baie est dense.
5. ⚠️ Poussière (ateliers, BTP, menuiserie) : le S310 est **IP20**, sans filtre.
   Prévois un nettoyage à l'air comprimé (bombe, pas de compresseur d'atelier
   humide) lors de la maintenance (section 114), ou un coffret étanche ventilé.

---

## 13. Déballage : contenu du carton et checklist de réception

**Contenu typique (à vérifier sur le bordereau du modèle exact) :**

- [ ] Le switch S310 (référence conforme au bon de commande : T/P, S/X, 24/48)
- [ ] Kit d'oreilles rack 19" + visserie
- [ ] Cordon secteur (avec sangle de maintien sur certains modèles)
- [ ] Câble console (RJ45 vers série/USB selon version — à vérifier)
- [ ] Pieds adhésifs pour pose bureau
- [ ] Guide de démarrage rapide / carte de garantie

**Checklist de réception (5 minutes qui évitent des semaines de galère) :**

- [ ] Référence sur l'étiquette = référence commandée (un 24T4S au lieu d'un 24P4S,
      ça arrive, et ça se voit **après** le câblage PoE).
- [ ] Numéro de série noté et archivé (photo de l'étiquette) — indispensable pour
      le support et la garantie.
- [ ] Version logicielle affichée au boot notée (pour planifier la mise à jour,
      section 102).
- [ ] Aucun choc sur le châssis, ventilateurs qui tournent librement (test à
      l'oreille au premier démarrage).
- [ ] Test de base : 2 PC en DHCP ou IP fixe sur 2 ports → ping OK avant montage.

🔧 **Astuce stock :** garde un carton d'origine pour chaque switch en spare.
Un retour SAV sans emballage d'origine, c'est un risque de casse refusée.

---

## 14. Montage en rack 19" : procédure pas à pas

1. Fixe les **oreilles rack** de chaque côté avec les vis fournies (ne force pas :
   un filetage foiré = oreille qui ne tient plus).
2. Présente le switch dans la baie à la hauteur prévue (prévois le brassage :
   panneau de brassage **juste au-dessus** ou **juste en dessous**).
3. Visse les 4 vis cage (2 par côté) sans serrer à fond, ajuste l'alignement,
   puis serre en croix.
4. ✅ Le switch 3,2–3,8 kg ne nécessite pas de rails, mais en baie profonde
   (> 80 cm) des **étagères 1U** ou rails légers évitent le porte-à-faux.
5. Passe la **terre** (section 15) **avant** le cordon secteur.
6. Câble les ports (section 16), puis alimente (section 10).

**Plan de baie type PME (exemple) :**

| U | Équipement |
|---|---|
| 42 | Panneau de brassage 24 ports (arrivées postes) |
| 41 | S310-24P4S (SW-ACC-01) |
| 40 | 1U vide (ventilation) ou second S310 |
| 39 | Onduleur / PDU |

---

## 15. Mise à la terre : pourquoi et comment (ne la saute pas)

La vis de terre en face arrière n'est pas décorative : avec la protection ±6 kV et
des liaisons cuivre longues, une terre correcte évite les reboots fantômes et les
ports qui « tombent » pendant les orages.

1. Utilise un fil **vert/jaune ≥ 2,5 mm²** (6 mm² si la baie est loin de la barrette).
2. Visse sur la borne de terre du switch, l'autre extrémité sur la **barrette de
   terre de la baie**, elle-même reliée à la terre du bâtiment.
3. Serre correctement : une cosse qui bouge = une terre qui ne sert à rien.
4. 🔧 Vérifie à l'ohmmètre si tu as un doute (continuité borne ↔ barrette).
5. ⚠️ Ne jamais utiliser le neutre ou une tuyauterie comme terre.

---

## 16. Câblage : choisir et poser les bons câbles

**Cuivre (ports RJ45 GE) :**

| Usage | Câble conseillé |
|---|---|
| Postes bureautiques, téléphones IP | Cat 6 (ou Cat 6A si neuf — pérenne 10G) |
| AP Wi-Fi 6 (surtout 24PN4X en 2,5G) | **Cat 6A minimum** pour le 2,5G/5G propre |
| Caméras PoE extérieures | Cat 6 extérieur (gaine UV) ou fibre + convertisseur |
| Jarretières en baie | 0,5–1 m, Cat 6, de couleur par usage (voir code couleur) |

**Code couleur suggéré (à afficher dans la baie) :**

| Couleur | Usage |
|---|---|
| Bleu | Postes utilisateurs |
| Jaune | AP Wi-Fi |
| Rouge | Uplinks / inter-switch |
| Vert | Téléphonie IP |
| Gris/noir | Serveurs / imprimantes |
| Orange | Caméras |

**Fibre (uplinks SFP/SFP+) :**

- SFP GE : modules **1000BASE-SX** (multimode, ≤ 550 m) ou **1000BASE-LX** (monomode, ≤ 10 km).
- SFP+ 10G : **10GBASE-SR** (multimode OM3/OM4) ou **10GBASE-LR** (monomode).
- ⚠️ N'achète que des modules **compatibles Huawei / codés Huawei** ou des compatibles
  tiers de qualité pro : un SFP non reconnu = uplink mort (voir dépannage, section 124).
- Nettoie les connecteurs (stylo de nettoyage fibre) avant insertion. Un connecteur
  sale = 3 dB de perte = lien instable.

**Règles de brassage :**

1. Longueur max d'un lien cuivre : **100 m** (90 m de câble fixe + 10 m de jarretières).
2. Sépare les chemins cuivre et 230 V (30 cm minimum en parcours parallèle).
3. Étiquette **les deux extrémités** de chaque câble (ex. `BRAS-12 ↔ SW1-P07`).
4. Laisse du mou (boucle de service) derrière la baie, jamais de câble tendu.
5. ✅ Photographie le brassage final : c'est ta doc « plan de câblage » v1.

---

## 17. Les 4 modes de management : choisir le bon

| Mode | Quand l'utiliser | Prérequis |
|---|---|---|
| **Cloud eKit** (app + portail) | Multi-sites, PME sans admin réseau à temps plein, supervision centralisée | Accès Internet sur le switch, compte Huawei eKit |
| **Web local** (HTTP/HTTPS) | Site unique, configuration ponctuelle, technicien sur place | IP de management joignable |
| **CLI** (console/SSH) | Scripts, dépannage avancé, tout ce que le web ne fait pas | Câble console ou SSH activé |
| **SNMP** | Supervision (Zabbix, PRTG, Centreon...), pas de configuration | SNMP activé + superviseur |

✅ **En pratique PME :** onboarding en **cloud eKit** pour le déploiement (rapide, guidé),
puis **CLI/SSH** pour le réglage fin et le dépannage. Les trois modes coexistent :
le cloud ne verrouille pas le CLI.

> Le cloud eKit utilise **NETCONF/YANG** en sous-jacent (valeur datasheet). Les deux
> modes « cloud » et « on-premise » sont commutables selon le besoin.

---

## 18. Branchement console : paramètres et matériel

Le port console RJ45 sert à la première configuration et au **secours** (quand le réseau
est en vrac, la console répond toujours).

**Paramètres série (standard Huawei) :**

| Paramètre | Valeur |
|---|---|
| Débit | 9 600 bit/s |
| Bits de données | 8 |
| Parité | Aucune |
| Bits d'arrêt | 1 |
| Contrôle de flux | Aucun |

**Matériel :**

- Câble console RJ45 ↔ USB (la plupart des PC récents n'ont plus de port série DB9).
  Installe le pilote du chipset (Prolific PL2303 / FTDI / CH340 selon le câble).
- Logiciels : PuTTY, Tera Term, ou `screen /dev/ttyUSB0 9600` sous Linux.

🔧 **Test rapide :** allume le switch avec la console branchée : tu dois voir les
messages de boot défiler. Si rien ne s'affiche : mauvais port COM, mauvais débit,
ou câble console HS (en avoir 2 en sacoche, c'est du vécu).

---

## 19. Première connexion console : pas à pas

