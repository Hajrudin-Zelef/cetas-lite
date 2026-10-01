---
id: collect-261001-rattrapage/rattrapage/datacenter-builds-guide-12
title: "Datacenter Builds — Le guide des BOMs"
domain: rattrapage
role: reference
task: reference
actors: []
dates: ["2026-09-27"]
keywords: ["datacenter", "capex", "compute", "gpu", "sol"]
source: docs/RAG/collect-261001-rattrapage/datacenter_builds_guide.md
source_anchor: ""
source_lines: [2017, 2194]
sha256: dc8613573c985aff95f36d97c67154353bffbbdfce6c0125b07a216c51beddb7
---

# Datacenter Builds — Le guide des BOMs

| Build | Section | P_max | P_réaliste |
|---|---|---|---|
| Compute S | §18–19 | 0,62 kW | 0,47 kW |
| Compute M | §20–21 | 1,26 kW | 0,95 kW |
| PG OLTP S | §28 | 0,79 kW | 0,60 kW |
| PG OLTP M | §29 | 1,43 kW | 1,05 kW |
| PG Warehouse | §30 | 1,54 kW | 1,10 kW |
| Ceph OSD NVMe S | §42 | 0,78 kW | 0,60 kW |
| Ceph OSD NVMe M | §43 | 1,49 kW | 1,10 kW |
| Ceph hybride S | §54 | 0,74 kW | 0,55 kW |
| Ceph hybride M | §55 | 1,53 kW | 1,10 kW |
| VDI S (100 users) | §63 | 1,35 kW | 1,00 kW |
| VDI M vGPU (100 users) | §64 | 2,68 kW | 2,00 kW |
| SIEM S | §73 | 0,95 kW | 0,70 kW |
| SIEM M | §74 | 1,60 kW | 1,20 kW |
| 2× MI300X | §82 | 2,30 kW | 1,80 kW |
| 4× MI300X | §83 | 4,40 kW | 3,50 kW |
| 2× L40S | §92 | 1,38 kW | 1,05 kW |
| 4× L40S | §93 | 2,69 kW | 2,00 kW |
| 2× RTX PRO 6000 | §94 | 1,88 kW | 1,40 kW |
| 4× RTX PRO 6000 | §95 | 3,74 kW | 2,80 kW |
| Backup S / M | §102–103 | 0,70 / 0,85 kW | 0,50 / 0,65 kW |
| Firewall 1G / 10G / 25G | §110–112 | 0,02 / 0,12 / 0,25 kW | idem |
| HSM (paire) | §116 | 0,60 kW | 0,40 kW |

**C'est LE tableau d'entrée du dimensionnement électrique (partie D).**

---

## 126. Synthèse budgets par workload (ordre de grandeur)

| Workload | Build d'entrée | Budget CAPEX | kW/rack typique |
|---|---|---|---|
| Compute | S : 30 k€ | 90–190 k€ (3–6 nœuds) | 3–6 kW |
| DB HA | 2× S : 93 k€ | 100–210 k€ | 2–4 kW |
| Ceph NVMe 150 To | 5× S : 312 k€ | 300–500 k€ | 6–10 kW |
| Ceph hybride 2 Po | 8× M : 615 k€ | 600–800 k€ | 12–18 kW |
| VDI 300 postes | 328 k€ + 60 k€/an | 350–500 k€ | 4–6 kW |
| SOC 1 To/j | 320 k€ | 300–450 k€ | 5–8 kW |
| Inférence 70B | 35–97 k€ | 40–150 k€ | 2–5 kW/serveur |
| Backup 100 To (3-2-1) | 74 k€ + 36 k€/an | 75–120 k€ | 1–2 kW |
| Firewall 10G HA | 6,7 k€ | 7–15 k€ | 0,3 kW |
| HSM (paire) | devis (15–50 k€/u) | 30–100 k€ | 0,6 kW |

---

# PARTIE C — MONTAGE DATACENTER COMPLET

---

## 127. Rack : le choix du 42U

| Critère | Recommandation PME/datacenter | Pourquoi |
|---|---|---|
| Hauteur | 42U (rarement 47U) | standard, pièces détachées |
| Largeur | 600 mm (800 mm si câblage dense) | 800 mm = gestion verticale des câbles |
| Profondeur | **1 100–1 200 mm** | les serveurs GPU font 800+ mm + câbles |
| Charge statique | ≥ 1 500 kg | un rack plein pèse 800–1 200 kg |
| Charge dynamique | ≥ 1 000 kg | déplacement chargé (à éviter quand même) |
| Portes | perforées ≥ 80 % (avant+arrière) | le plein panneau tue le flux d'air |
| Prix | ≈ 2 700 $ (APC NetShelter AR3100, vérifié le 27/09/2026) | 1 500–3 000 € selon options |

**Ne prenez jamais un rack de 1 000 mm de profondeur** pour des serveurs :
avec les câbles et les PDU 0U, il vous manquera 15 cm. C'est l'erreur n° 1.

---

## 128. Charge au sol : le calcul qu'on oublie

- Rack plein : ≈ 1 000 kg sur 0,72 m² (600×1200) = **1 390 kg/m²**.
- Dalle technique standard : 500–1 000 kg/m² → **insuffisant** !
- Solutions : répartiteurs de charge, racks sur dalle béton (rez-de-chaussée),
  ou étude structurelle. **Faites vérifier par un BE** au-delà de 800 kg/rack.
- Sismique : ancrage au sol dans les zones concernées.

---

## 129. PDU : metered vs switched, bien choisir

| Type | Mesure | Pilotage prises | Prix indicatif 22 kW 3-ph | Usage |
|---|---|---|---|---|
| Basic | non | non | ≈ 300–500 € | à éviter en prod |
| Metered | globale (+ par phase) | non | ≈ 2 050 £ (~2 400 €, vérifié le 27/09/2026) | standard mini |
| Metered per-outlet | par prise | non | ≈ 3 000–4 000 € (à vérifier) | facturation, debug |
| Switched | globale | oui (on/off distant) | ≈ 3 782 $ (~3 480 €, vérifié le 27/09/2026) | redémarrage distant |
| Switched per-outlet | par prise | oui | ≈ 5 200 $ (~4 800 €, vérifié le 27/09/2026) | le top |

**Recommandation : 2× PDU metered par rack** (une par arrivée A/B), switched
si vous n'avez pas d'IPMI fiable partout. La mesure par prise paie son
surcoût dès le premier debug de surcharge.

---

## 130. Tableau PDU : quelle puissance par rack

| PDU | Tension | Courant | Puissance | Prix indicatif |
|---|---|---|---|---|
| Mono 16 A | 230 V | 16 A | 3,7 kW | ≈ 400–600 € |
| Mono 32 A | 230 V | 32 A | 7,4 kW | ≈ 600–900 € |
| Tri 16 A | 400 V | 3× 16 A | 11 kW | ≈ 1 500–2 000 € |
| Tri 32 A | 400 V | 3× 32 A | 22 kW | ≈ 2 400 € (metered, §129) |
| Tri 63 A | 400 V | 3× 63 A | 43,5 kW | ≈ 4 000–5 000 € (à vérifier) |

Règle : **PDU chargée à 80 % max** en nominal. Un rack à 5,6 kW (§120) →
PDU tri 16 A (11 kW) confortable, ou tri 32 A pour la croissance.

---

## 131. Répartition des charges par phase — exemple chiffré

Rack §120 (5,6 kW) sur PDU tri 16 A (3× 16 A = 11 kW) :
- Phase L1 : 2× Compute M (1,9 kW) + 1 switch (0,3) = **2,2 kW (9,6 A)**
- Phase L2 : 2× PG S (1,2 kW) + 2× Ceph S (1,2 kW) = **2,4 kW (10,4 A)**
- Phase L3 : backup (0,5) + FW (0,2) + switch (0,3) = **1,0 kW (4,3 A)**
- Déséquilibre : 10,4 vs 4,3 A → **rééquilibrez** : déplacez un Ceph S en L3.
- Cible : **écart < 20 %** entre phases (neutre non surchargé, disjoncteur
  différentiel content).

**Les 2 arrivées A/B** : chaque serveur double-alimenté prend 1 cordon sur
chaque PDU. En cas de perte d'une arrivée, l'autre porte 100 % : donc chaque
PDU doit pouvoir porter **le rack entier** (pas la moitié !).

---

## 132. Schéma ASCII — rack 42U type (§120)

```
+--------------------------------------------------+
| 42                                               |  <- panneau de brassage
| 41  [ Patch panel fibre 24 ports ]               |
| 40                                               |
| 39  [ Switch ToR-1 25/100G ]                     |  0,3 kW (A/B)
| 38  [ Switch ToR-2 25/100G ]                     |  0,3 kW
| 37                                               |
| 36  [ FW-1 ]  [ FW-2 ]  (1U côte à côte)          |  0,2 kW
| 35                                               |
| 34  [ Backup S 2U ]                              |  0,5 kW
| 33                                               |
| 32  [ Ceph OSD S 2U ]                            |  0,6 kW
| 31                                               |
| 30  [ Ceph OSD S 2U ]                            |  0,6 kW
| 29                                               |
| 28  [ PG OLTP S 2U ]                             |  0,6 kW
| 27                                               |
| 26  [ PG OLTP S 2U ]                             |  0,6 kW
| 25                                               |
| 24  [ Compute M 2U ]                             |  0,95 kW
| 23                                               |
| 22  [ Compute M 2U ]                             |  0,95 kW
| 21                                               |
| 20  [ KVM / console ]                            |
| 19                                               |
|  1-18  LIBRE (croissance : 18U = 40 % du rack)   |
+--------------------------------------------------+
  | PDU-A (gauche, 0U) | | PDU-B (droite, 0U) |
```

Règles de placement : switchs en haut (câbles courts vers le brassage),
serveurs lourds en bas (stabilité), 1U libre entre groupes pour l'air.

---

## 133. Rails, gestion verticale, accessoires

| Accessoire | Prix indicatif | Pourquoi |
|---|---|---|
| Rails coulissants (par serveur) | 100–150 € | inclus souvent — vérifiez |
| Gestionnaire de câbles vertical 0U | 80–150 € | sans ça, c'est le chaos en 6 mois |
| Panneau obturateur 1U (×10) | 50 € | **obligatoire** : l'air passe par les trous sinon |
| Brosse passe-câbles | 30 € | étanchéité + passage |
| Tablette 1U | 40 € | clavier/écran crash-cart |

**Les panneaux obturateurs sont le meilleur investissement thermique du rack**
(50 € pour −3 °C en entrée serveur). Les racks « gruyère » surchauffent par le
haut.

---

