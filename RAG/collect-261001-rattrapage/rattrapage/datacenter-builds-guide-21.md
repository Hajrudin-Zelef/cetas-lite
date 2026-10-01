---
id: collect-261001-rattrapage/rattrapage/datacenter-builds-guide-21
title: "Datacenter Builds — Le guide des BOMs"
domain: rattrapage
role: reference
task: reference
actors: ["AMD", "Intel"]
dates: ["2026-09-27"]
keywords: ["datacenter", "amd", "arr", "compute", "distribution", "dram", "gpu", "intel"]
source: docs/RAG/collect-261001-rattrapage/datacenter_builds_guide.md
source_anchor: ""
source_lines: [3750, 3953]
sha256: 62a421e8d6fc2c3e6c83251ffb4fd8a73b407a19719df5ac7a15bd31dff25a14
---

# Datacenter Builds — Le guide des BOMs

1. **Identifier** le circuit (étiquette, plan TGBT).
2. **Séparer** : ouvrir le disjoncteur, verrouiller avec cadenas personnel.
3. **Condamner** : étiquette « NE PAS MANŒUVRER » nominative.
4. **Vérifier** : VAT (vérificateur d'absence de tension) sur chaque phase.
5. **Travailler**, puis déconsigner en sens inverse.

**Un seul cadenas par intervenant.** Le chef ne retire jamais le cadenas d'un
autre. C'est la règle qui sauve des vies — pas une option.

---

## 229. Gestion des changements (change management)

| Changement | Procédure |
|---|---|
| Standard (MAJ firmware planifiée) | ticket, fenêtre, rollback |
| Normal (nouveau serveur) | demande, validation, as-built MAJ |
| Urgent (panne) | intervention, ticket a posteriori sous 24 h |

**Fenêtres de maintenance** : 1×/mois, nuit ou week-end, communiquées 1 sem
avant. **Jamais de changement non tracé en prod** — le « petit fix rapide »
est la cause n° 1 des incidents majeurs.

---

## 230. Faut-il secourir la climatisation sur onduleur ?

| Scénario | Réponse |
|---|---|
| Coupure < 15 min + arrêt auto (NUT) | **Non** : les serveurs tiennent thermiquement 10–15 min sans froid |
| Groupe électrogène présent | **Non** : le groupe reprend le froid en < 2 min |
| Pas de groupe, autonomie > 30 min | **Oui** : sinon la salle surchauffe avant la fin des batteries |
| Salle GPU dense (> 20 kW/rack) | **Oui** : l'inertie thermique est de quelques minutes seulement |

Règle : **l'inertie thermique d'une salle air = 10–20 min**, d'une salle
GPU dense = 3–5 min. Dimensionnez en conséquence (§164).

---

## 231. Scénario blackout — la procédure

```
T+0    : coupure. Onduleurs sur batteries, alerte SMS.
T+2    : groupe démarre (auto). Vérifier prise de charge.
T+5    : si groupe OK → rien à faire, surveillance.
T+5    : si groupe KO → délestage : arrêt NUT des non-critiques.
T+10   : si toujours pas de groupe → arrêt propre général (runbook).
T+30   : batteries vides. Tout est arrêté proprement. On attend.
Retour : groupe ou réseau → démarrage séquencé (§166), vérif services.
```

**Testez ce scénario 1×/an** (un dimanche). Le runbook non testé est un vœu
pieux. Chronométrez chaque étape.

---

## 232. Connecteurs fibre — tableau

| Connecteur | Usage | Remarque |
|---|---|---|
| LC duplex | 10/25/100G SR4 (via MPO→LC) | standard 2026 |
| MPO-12 | 100/400G SR4 | 1 connecteur = 4–8 fibres |
| SC | ancien, management | en voie de disparition |
| E2000 | laser haute puissance | rare en datacenter |

**Nettoyage** : stylo de nettoyage ou cassette avant CHAQUE connexion.
Une fibre sale = −3 dB = lien instable. Le microscope d'inspection (200 €)
est rentabilisé au premier lien douteux.

---

## 233. Brassage fibre — bonnes pratiques

- Tiroirs coulissants avec **gestion du mou** (jamais de fibre tendue).
- Étiquette à chaque extrémité : `R01-T1-P03 → R02-T1-P03`.
- **Ne regardez jamais** dans une fibre active (laser 850/1310 nm, invisible,
  dangereux pour la rétine).
- Stock : 10 % de jarretières de rechange, 2 cassettes de nettoyage.

---

## 234. Glossaire complémentaire (12 termes)

| Terme | Définition |
|---|---|
| **LOTO** | Lockout/Tagout : consignation électrique par cadenas personnel. |
| **VAT** | Vérificateur d'Absence de Tension : test avant intervention. |
| **TGBT** | Tableau Général Basse Tension : arrivée électrique du bâtiment. |
| **THD** | Taux de distorsion harmonique : les PSU à PFC le maintiennent < 5 %. |
| **Cos φ / fp** | Facteur de puissance : kW/kVA. 1,0 = parfait (PFC actif). |
| **MPO** | Multi-fiber Push On : connecteur 12/24 fibres pour 100/400G. |
| **OTDR** | Réflectomètre : mesure les défauts d'une fibre (distance, perte). |
| **VESDA** | Détection incendie par aspiration : alerte précoce. |
| **IG-541** | Gaz inerte d'extinction (azote/argon/CO2), sans résidu. |
| **MTBF** | Mean Time Between Failures : fiabilité théorique (ex : 2 Mh SSD). |
| **RPO / RTO** | Perte de données max / durée d'arrêt max acceptées. |
| **DEEE** | Déchets d'équipements électriques et électroniques : filière obligatoire. |

---

## 235. La phrase à retenir par partie

| Partie | Phrase |
|---|---|
| A — Méthode | Mesurez d'abord, la DRAM 2026 ne pardonne pas l'à-peu-près. |
| B — BOMs | Chaque workload a sa BOM ; le tableau §125 les relie à l'électrique. |
| C — Datacenter | 2 PDU A/B, phases équilibrées, rack 1 200 mm, obturateurs partout. |
| Câblage | DAC en rack, fibre dehors, velcro toujours, étiquettes des 2 côtés. |
| Froid | kW × 300 = m³/h ; au-delà de 20 kW/rack, le liquide s'impose. |
| Énergie | Onduleur en kW, batteries testées, groupe testé en charge. |
| MES | Burn-in 72 h, firmware uniformes, as-built livré. |
| Échelle | Copiez les méthodes des hyperscalers, pas leurs densités. |

---

## 236. Fin du guide

Ce guide est un **outil de travail**, pas une lecture : imprimez l'antisèche
(§205), la check-list devis (§122) et le top 15 des pièges. Le reste vit dans
le classeur de la salle, à côté de l'as-built.

Prochaines étapes suggérées : faire chiffrer 3 devis sur la BOM Compute S
(§18) pour étalonner VOS prix réels — c'est le test qui valide tout le reste.

*— Fin du guide datacenter_builds — 27/09/2026.*

---

## 237. FAQ express — les questions qu'on pose toujours

**Faut-il du bi-socket en 2026 ?**
Rarement. Un EPYC 9655P 96 cœurs mono-socket bat la plupart des bi-socket
d'hier. Le 2P ne se justifie que par > 1,5 To RAM ou > 96 lanes PCIe utiles.

**AMD ou Intel ?**
AMD par défaut (prix/cœur, lanes). Intel si l'appli/l'appliance l'exige
(AMX, QAT, certification).

**Combien de RAM par cœur ?**
4 Go mini (HPC), 8 Go standard (virt/K8s), 12–16 Go (IA/data).

**SATA a-t-il encore un sens ?**
Pour le boot en miroir (bon marché) et l'archivage froid. Jamais en perf.

**Ceph ou ZFS ?**
Ceph = distribué, multi-nœuds, S3. ZFS = local, 1–2 nœuds, simple.
En dessous de 100 To sur 3 nœuds, ZFS répliqué (Proxmox) est souvent plus
simple que Ceph.

**Cloud ou on-premise ?**
Charge stable > 18 mois → on-premise (2–3× moins cher, §185).
Charge variable ou < 18 mois → cloud.

**Quelle température en salle ?**
24–27 °C en entrée rack (ASHRAE A1). 21 °C = habitude chère.

**Faut-il un groupe électrogène ?**
Oui dès que l'indisponibilité coûte plus cher que 55 k€ + fioul.
En pratique : oui pour toute salle > 20 kW.

**Onduleur : quelle autonomie ?**
15 min + arrêt auto (NUT). Plus = groupe, pas des batteries.

**Par où commencer ?**
Un rack §120 (5,6 kW, 354 k€) : c'est le « datacenter minimum viable ».

---

## 238. Où acheter — les circuits (France/Europe)

| Circuit | Pour | Remarque |
|---|---|---|
| Intégrateurs (Serveur/boutiques pro) | serveurs complets | devis, burn-in, garantie — le circuit recommandé |
| Distributeurs (Tech Data, Ingram) | volumes, licences | via revendeur |
| Constructeurs direct (Dell, HPE) | parc homogène | cher mais support intégré |
| Occasion pro (garantie 1 an) | GPU, serveurs N−1 | §194 — jamais SSD/batteries |
| Grossistes câble/fibre | câblage | FS.com et équivalents (optiques codées) |

**Toujours 3 devis** (§193). Le prix public de ce guide sert d'étalon pour
détecter les devis anormaux (±30 % = question à poser).

---

## 239. Lexique anglais-français rapide

| EN | FR |
|---|---|
| Bill of Materials | nomenclature |
| Burn-in | test d'endurance |
| Hot-swap | remplaçable à chaud |
| Cold plate | plaque froide (DLC) |
| Rear door heat exchanger | échangeur porte arrière |
| Top of Rack | haut de baie |
| Power distribution unit | unité de distribution |
| Uninterruptible power supply | onduleur / ASI |
| Erasure coding | codage à effacement |
| Scrubbing | contrôle d'intégrité |
| Failover | bascule |
| As-built | dossier de récolement |

---

## 240. Les 10 chiffres à connaître par cœur

