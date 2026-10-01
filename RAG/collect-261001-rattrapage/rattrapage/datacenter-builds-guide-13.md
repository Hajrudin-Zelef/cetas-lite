---
id: collect-261001-rattrapage/rattrapage/datacenter-builds-guide-13
title: "Datacenter Builds — Le guide des BOMs"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["arr", "capex", "incident", "sol"]
source: docs/RAG/collect-261001-rattrapage/datacenter_builds_guide.md
source_anchor: ""
source_lines: [2195, 2416]
sha256: f85b3042bae2fca4762c30a2c142142c1dbff1e2f5d44262602051a0b5613ea7
---

# Datacenter Builds — Le guide des BOMs

## 134. KVM, IPMI, console : l'accès hors-bande

- Chaque serveur : **port BMC/IPMI dédié** sur un VLAN/switch de management
  séparé (pas sur le LAN de prod).
- Licences : iDRAC Enterprise / iLO Advanced / XCC Enterprise se paient
  (200–400 €/serveur — à vérifier). **Négociez-les dans le devis**, pas après.
- 1× console série (OpenGear, Digi) par rangée de racks : quand le réseau
  est down, le série reste. ≈ 1 500–3 000 € (à vérifier).
- Documentez **tous** les mots de passe BMC dans le coffre (HSM §117 ou
  équivalent) — pas sur un post-it.

---

## 135. Sécurité physique du rack

- Serrures à badge ou code (pas la clé universelle constructeur).
- Caméra sur l'allée, détection d'ouverture de porte (contact sec → SNMP).
- **Personne ne travaille seul** dans la salle (règle des 2 personnes pour
  l'électrique).
- Registre d'accès : qui, quand, pourquoi. En cas d'audit, c'est demandé.

---

## 136. Pièges terrain — Rack / PDU

1. **Rack 1 000 mm** pour des serveurs de 800 mm + PDU 0U : ça ne ferme pas.
2. **Une seule PDU** : la maintenance électrique = arrêt total. Toujours A+B.
3. **Phases déséquilibrées** : le neutre chauffe, le différentiel saute.
4. **Pas de panneaux obturateurs** : recyclage d'air chaud, +5 °C en haut.
5. **Câbles d'alim qui pendent devant les ventilos** : −30 % de flux d'air.
6. **Charge au sol non vérifiée** : 1 390 kg/m² sur une dalle à 500 kg/m².

---

## 137. Check-list réception rack

- [ ] Profondeur ≥ 1 100 mm vérifiée au mètre
- [ ] 2× PDU (A+B), type metered mini, ampérage ≥ P_max × 1,25 / 0,8
- [ ] Panneaux obturateurs sur tous les U libres
- [ ] Rails : 1 jeu par serveur, montés
- [ ] Ancrage au sol si > 800 kg ou zone sismique
- [ ] Étiquette du rack (nom, PDU A/B, phases)

---

## 138. Budget rack complet (vide → prêt)

| Poste | Prix |
|---|---|
| Rack 42U 600×1200 | ≈ 2 500 € |
| 2× PDU metered tri 32 A | 2 × 2 400 € = 4 800 € |
| Accessoires (obturateurs, gestion câbles) | 500 € |
| Brassage + jarretières (cuivre+fibre) | 1 500 € |
| **TOTAL / rack** | **≈ 9 300 €** |

À multiplier par le nombre de racks — et à ne pas oublier dans le budget
global (c'est 10–15 % du CAPEX qui disparaît si on l'oublie).

---

## 139. Câblage : cuivre vs fibre — tableau de choix

| Lien | Cuivre (DAC/RJ45) | Fibre (optique) | Conseil 2026 |
|---|---|---|---|
| 1G management | RJ45 Cat6 | — | cuivre |
| 10G < 5 m | DAC SFP+ (~30 €) | — | DAC |
| 25G < 5 m | DAC SFP28 (~40 €) | — | DAC |
| 25G > 5 m | — | SR (~80 €) | fibre |
| 100G < 3 m | DAC QSFP28 (~80 €) | — | DAC |
| 100G > 3 m | — | SR4 (~250 €) | fibre |
| 400G | — | SR4 (~800 €, à vérifier) | fibre |

Règle : **DAC en intra-rack** (pas cher, pas de panne d'optique), **fibre dès
que ça sort du rack**. Un DAC qui traverse 3 racks est un incident en attente.

---

## 140. AOC vs DAC vs optique+jarretière

| Solution | Prix 100G 5 m | Avantages | Inconvénients |
|---|---|---|---|
| DAC | ≈ 80 € | pas cher, fiable | rigide, ≤ 5 m |
| AOC (câble optique actif) | ≈ 200 € (à vérifier) | souple, 30 m | électronique aux 2 bouts |
| Optique + jarretière | ≈ 250 € + 30 € | standard, réparable | 2 points de panne |

En 2026, l'AOC gagne du terrain pour le ToR→serveur quand le DAC est trop
court : un seul composant, pas de nettoyage de connecteurs.

---

## 141. Chemins de câbles : l'infra invisible

- **Chemin de câbles** au-dessus des racks (300–600 mm de large) : cuivre d'un
  côté, fibre de l'autre (séparation EMI + rayon de courbure).
- **Fibre** : rayon de courbure mini 30 mm (OM4) — un angle à 90° sec = fibre
  morte dans 2 ans.
- **Descente** dans le rack par le haut (pas par le bas — l'eau et la
  poussière montent moins bien).
- Prévoyez **30 % de place libre** dans les chemins : le recâblage d'un chemin
  plein est un cauchemar.

---

## 142. Étiquetage : la norme maison (à écrire et afficher)

Format proposé : `RACK-SWITCH-PORT → RACK-SERVEUR-PORT`

```
R01-TOR1-ETH25 → R01-SRV03-ETH0
R01-TOR1-ETH26 → R01-SRV03-ETH1
```

- Étiquettes **des 2 côtés** de chaque câble, lisibles sans débrancher.
- Code couleur : **bleu** = data, **jaune** = management/IPMI,
  **rouge** = arrivée A, **noir** = arrivée B (ou votre code — mais UN code).
- Imprimante d'étiquettes pro (Brady/Dymo Rhino) : 200–400 € — rentabilisée
  au premier dépannage nocturne.

---

## 143. Brassage : l'architecture ToR

```
    [ Core / Agg switches ]
           | 100G (MLAG)
    +------+------+
    | ToR-1      ToR-2 |   ← 1 paire par rack (ou par 2 racks)
    +------+------+
     | 25G DAC | 25G DAC
  [srv-01] [srv-02] ...
```

- **ToR (Top of Rack)** : 1–2 switchs par rack, uplinks en MLAG vers
  l'agrégation. Pas de switch « fin de rangée » en 2026 sauf très gros site.
- Chaque serveur : **2 liens** (1 vers chaque ToR) en LACP/bond actif-passif.
- Management : 1 switch 1G 48 ports par 2–4 racks, réseau isolé.

---

## 144. Schéma ASCII — câblage d'un serveur type

```
Façade arrière serveur (2U) :

 [PSU-A]====rouge====>[ PDU-A ]   (alim redondante)
 [PSU-B]====noir=====>[ PDU-B ]

 [BMC ]-----jaune---->[ SW-MGMT ]  (IPMI, VLAN 99)
 [ETH0]-----bleu----->[ TOR1-ETH25 ]  (DAC 25G)
 [ETH1]-----bleu----->[ TOR2-ETH25 ]  (DAC 25G)
 [ETH2]~~~~~fibre~~~~>[ TOR1-ETH1 ]  (100G, si double NIC)
 [ETH3]~~~~~fibre~~~~>[ TOR2-ETH1 ]
```

`====` = cordon d'alim C13/C19, `-----` = cuivre/DAC, `~~~~~` = fibre.

---

## 145. Longueurs et stocks : le kit du câbleur

| Article | Stock conseillé (par rack) |
|---|---|
| DAC 25G 1 m / 3 m | 10 / 10 |
| DAC 100G 3 m | 4 |
| Jarretières OM4 LC-LC 3 m / 10 m | 10 / 4 |
| Cordons d'alim C13 2 m / C19 2 m | 10 / 4 |
| Étiquettes + ruban | 1 rouleau |
| Attaches velcro (pas de colliers !) | 1 paquet |

**Velcro, jamais de colliers de serrage** : un collier trop serré écrase le
cuivre et casse la fibre. Le jour où il faut recâbler, le velcro se rouvre.

---

## 146. Pièges terrain — Câblage

1. **DAC de 7 m** : n'existe pas en fiable. Fibre au-delà de 5 m.
2. **Fibre pliée à 90°** dans le gestionnaire : micro-cassures, erreurs CRC
   dans 6 mois.
3. **Une seule étiquette** (ou pire, au marqueur) : illisible dans 2 ans.
4. **Cuivre et fibre dans le même chemin sans séparation** : le jour où on
   tire un câble, on arrache les connecteurs LC.
5. **Pas de plan de câblage** : le « on s'en souvient » ne survit pas au
   premier départ en retraite.
6. **Optiques non codées** pour le switch : un transceiver « générique » non
   reconnu = lien down. Vérifiez la compatibilité (FS.com code sur demande).

---

## 147. Plan de câblage : le document

Un tableur (ou NetBox — voir guide netbox) avec : câble ID, de, vers, type,
longueur, couleur, date. **Chaque câble posé = 1 ligne.** 30 minutes de saisie
par rack évitent 4 h de debug par incident.

---

## 148. Budget câblage par rack

| Poste | Prix |
|---|---|
| 20× DAC 25G | 800 € |
| 4× DAC/optiques 100G | 800 € |
| Jarretières + cordons | 400 € |
| Étiquettes, velcro, divers | 200 € |
| **TOTAL** | **≈ 2 200 €** |

Comptez-le : sur 10 racks, c'est 22 000 € « invisibles ».

---

## 149. Refroidissement : les seuils qui décident de tout

| Densité/rack | Solution | Statut 2026 |
|---|---|---|
| < 15 kW | Air seul (allées chaude/froide) | standard |
| 15–30 kW | Air optimisé + confinement | limite haute de l'air |
| 30–50 kW | **Rear-door heat exchanger (RDHx)** | dominant (37 % adoption/envisagé) |
| 50–100 kW | RDHx + DLC partiel (hybride) | mainstream IA |
| > 100 kW | **DLC direct-to-chip obligatoire** | racks NVL72 > 120 kW |
| > 240 kW | Immersion (émergent) | hyperscalers |

