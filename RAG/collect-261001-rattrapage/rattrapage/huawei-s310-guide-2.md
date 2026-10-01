---
id: collect-261001-rattrapage/rattrapage/huawei-s310-guide-2
title: "Guide ultra-complet — Huawei eKit S310"
domain: rattrapage
role: reference
task: reference
actors: ["Huawei"]
dates: []
keywords: []
source: docs/RAG/collect-261001-rattrapage/huawei_s310_guide.md
source_anchor: ""
source_lines: [145, 288]
sha256: 7ffbcd4e54f05018f09351764ad618e2bd2aa7f754465b64c26d67086d475660
---

# Guide ultra-complet — Huawei eKit S310

🔧 **Règle de dimensionnement uplink :** additionne les débits crête réalistes des ports
d'accès, applique un ratio de sursouscription de 4:1 à 8:1 en bureautique. 24 ports GE
avec 20 postes actifs à ~50 Mbit/s crête = ~1 Gbit/s → un uplink 1G suffit ; avec un NAS
ou 3 AP Wi-Fi 6, passe en 10G (modèle X) ou en Eth-Trunk (section 55).

---

## 6. Les 48 ports : S310-48P4S (et 48T4X / 48P4X)

| Caractéristique | S310-48P4S (valeurs constructeur) |
|---|---|
| Référence | 98012384 |
| Ports | 48× 10/100/1000BASE-T **PoE+** + 4× GE SFP |
| Capacité / débit | 104 Gbit/s / 77 Mpps |
| Budget PoE | **380 W** |
| Conso | 34,04 W (statique) / 48,64 W (typique) / 63,7 W (max sans PoE) / **462,8 W** (max pleine charge PoE) |
| Ventilateurs | 2, intégrés, vitesse intelligente — flux d'air **de gauche vers la droite** |
| Dimensions | 43,6 × 442 × 220 mm (227 mm avec les parties saillantes) — 1U |
| Poids | 3,24 kg (4,29 kg avec emballage) |
| Montage | Rack, bureau, **mural** |
| Température / humidité | –5 °C à +50 °C / 5–95 % HR sans condensation |
| Alimentation | AC intégrée 90–290 V AC (45–65 Hz) ; entrée DC haute tension 190–290 V DC supportée |
| Protection surtension | ±6 kV |

⚠️ **380 W pour 48 ports** : c'est moins que 24P4S en proportion (7,9 W/port en moyenne).
Tu ne peux pas mettre 48 équipements PoE+ à 15 W. La planification PoE (sections 66–70)
est **obligatoire** sur ce modèle. Priorités PoE à configurer dès le jour 1.

---

## 7. Modèles spécifiques : S310-24PN4X et S310-24ST4X

- **S310-24PN4X** : 24× **2,5GE** RJ45 PoE+ (2,5G/1G/100M/10M) + 4× 10GE SFP+. Capacité
  **200 Gbit/s**, 144 Mpps, PoE **400 W**. Le choix naturel pour des **AP Wi-Fi 6/6E**
  dont le débit réel dépasse 1 Gbit/s. Température de fonctionnement limitée à
  **–5 °C à +45 °C** (et non +50 °C) — à vérifier sur la fiche du modèle exact si ton
  local est chaud.
- **S310-24ST4X** : 24× **GE SFP** (dont 8 ports **combo** 10/100/1000BASE-T ou SFP) +
  4× 10GE SFP+. 128 Gbit/s, 96 Mpps. Pour les répartiteurs tout-fibre ou les environnements
  perturbés (ateliers, liaisons inter-bâtiments).
- **S310-24U4X** (vu dans certains datasheets) : variante PoE++ — **à vérifier sur la
  fiche du modèle exact**, ne pas confondre avec le 24P4X.

✅ **Combo, rappel :** sur un port combo, le cuivre **ou** la fibre est actif, jamais les
deux en même temps. La fibre est généralement prioritaire quand un SFP est inséré
(comportement à vérifier sur la version logicielle exacte).

---

## 8. Face avant : ports, LED, console — savoir lire le switch

De gauche à droite sur un S310-24Tx/24Px typique :

1. **24 ports RJ45 GE** (1 à 24), avec 2 LED par port (Link/Act).
2. **4 cages SFP/SFP+** (25 à 28), LED dédiées.
3. **Port console RJ45** : pour le câble console (première config, secours).
4. **Bouton MODE** : change le mode d'affichage des LED (à vérifier sur le modèle exact).
5. **Bouton RST** : reset (appui court = reboot, appui long ~5–10 s = retour usine —
   ⚠️ à vérifier sur le modèle exact, un appui long efface la configuration).
6. **LED système** : PWR (alimentation), SYS (état système).
7. Sur les modèles PoE : LED **PoE** par port ou globale.

Face arrière : **vis de terre**, **passage pour la sangle du cordon secteur**,
**prise AC** (alimentation intégrée, pas de module amovible sur cette gamme).

🔧 **Réflexe terrain :** avant de toucher à quoi que ce soit, photographie la face avant
(LED allumées) et note les ports occupés. En cas de panne, tu compares.

---

## 9. Les LED : signification complète

| LED | État | Signification |
|---|---|---|
| PWR | Vert fixe | Alimentation OK |
| PWR | Éteinte | Pas d'alimentation — vérifier cordon, prise, disjoncteur |
| PWR | Rouge / orange | Défaut d'alimentation — à vérifier sur le modèle exact |
| SYS | Vert fixe | Système démarré et fonctionnel |
| SYS | Vert clignotant | Démarrage en cours |
| SYS | Rouge fixe/clignotant | Défaut système — consulter les logs, puis support |
| Port Link/Act | Vert fixe | Lien établi (link up) |
| Port Link/Act | Vert clignotant | Activité (émission/réception) |
| Port Link/Act | Éteinte | Pas de lien — câble, carte réseau distante, port désactivé |
| Port Link/Act | Orange (si dispo) | Lien à 10/100 Mbit/s ou défaut — à vérifier sur le modèle exact |
| PoE (par port) | Vert fixe | PD alimenté en PoE |
| PoE (par port) | Clignotant / orange | Négociation ou défaut PoE (surcharge, court-circuit, PD incompatible) |
| SFP | Vert fixe | Module inséré et lien OK |
| SFP | Éteinte | Pas de module ou pas de lien |

> Les couleurs exactes et les modes du bouton MODE sont **à vérifier sur le guide
> matériel du modèle exact** : Huawei a des variantes selon les séries. En cas de doute,
> la CLI (`display interface brief`, `display poe`) dit toujours la vérité.

---

## 10. Première mise sous tension : ce qui est normal (et ce qui ne l'est pas)

1. Branche le cordon secteur **avec la sangle de maintien** (elle évite les
   débranchements accidentels — un classique en baie).
2. Le ventilateur tourne à plein régime quelques secondes, puis ralentit
   (vitesse intelligente). ✅ Normal.
3. La LED SYS clignote pendant le boot (~1 à 2 minutes sur cette gamme, ordre de
   grandeur à vérifier), puis passe au vert fixe. ✅ Normal.
4. Sur un modèle PoE, les PD sont alimentés en **quelques secondes** (Fast PoE),
   sans attendre la fin du boot. ✅ Normal — ne t'inquiète pas si une caméra
   s'allume « avant » le switch.
5. ⚠️ **Anormal :** SYS rouge, ventilateur à fond en permanence après 5 minutes,
   odeur de chaud, LED PWR éteinte → coupe l'alimentation et vérifie
   l'environnement (section 13) avant de re-tester.

---

## 11. Alimentation électrique : caractéristiques et dimensionnement

- **Entrée :** 100–240 V AC, 50/60 Hz (plage de fonctionnement 90–290 V AC, 45–65 Hz).
  Le switch accepte donc les réseaux 110 V comme 230 V sans sélecteur.
- **Protection :** ±6 kV en mode différentiel et commun (parafoudre intégré de base —
  ça ne remplace pas un parafoudre T1+T2 en tête d'installation, voir section 114).
- **Dimensionnement du circuit :**

| Modèle | Conso max à prévoir | Disjoncteur conseillé | Commentaire |
|---|---|---|---|
| 24T4S / 24T4X | ~35 W | 10 A courbe C | Négligeable, mais dédié quand même |
| 24P4S / 24P4X | ~492 W (pleine charge PoE) | 10 A courbe C | Prévoir la charge PoE réelle, pas le max théorique |
| 48P4S | ~463 W (pleine charge PoE) | 10 A courbe C | Idem |

- ✅ Chaque switch sur un **départ protégé identifié** au tableau (« Baie — SW1 »).
- ✅ En environnement critique : **onduleur** en amont (voir le guide onduleurs de
  Zelef : un 24P4S pleine charge ≈ 500 W → prévoir ~1 kVA d'UPS pour 30+ min
  d'autonomie selon la batterie).
- ⚠️ Ne jamais brancher le switch sur une prise commandée par interrupteur mural.

---

## 12. Ventilation, bruit, température : où poser le switch

| Modèle | Bruit normal | Bruit haute temp. | Conséquence pratique |
|---|---|---|---|
| 24T4S / 24T4X | 47 dB(A) | 51 dB(A) | Acceptable en local technique, limite en open space |
| 24P4S | 49,3 dB(A) | **63 dB(A)** | Local technique recommandé si PoE chargé |
| 48P4S | à vérifier | à vérifier | 2 ventilateurs — prévoir dégagement gauche/droite (flux d'air latéral) |

**Règles d'installation thermique :**

