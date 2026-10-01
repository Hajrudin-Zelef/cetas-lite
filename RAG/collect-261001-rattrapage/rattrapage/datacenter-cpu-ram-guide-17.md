---
id: collect-261001-rattrapage/rattrapage/datacenter-cpu-ram-guide-17
title: "CPU SERVEUR, RAM SERVEUR, CHÂSSIS & VENTILATION"
domain: rattrapage
role: reference
task: reference
actors: ["AMD", "Intel"]
dates: ["2026-09-27"]
keywords: ["attention", "capex", "dram", "hyperscaler", "intel", "memory"]
source: docs/RAG/collect-261001-rattrapage/datacenter_cpu_ram_guide.md
source_anchor: ""
source_lines: [2633, 2813]
sha256: f165b54f9b22a0b1987788e74f0f34d06f083ad498043edef5dfc7c46ead0ea7
---

# CPU SERVEUR, RAM SERVEUR, CHÂSSIS & VENTILATION

Formule générale : **BP = canaux × débit × 8 octets** (DDR = 8 octets par
transfert et par canal ; en DDR5, 2 sous-canaux × 32 bits + ECC).

| Plateforme | Canaux × débit | BP théorique |
|---|---|---|
| SP5 Turin (12× 6400) | 12 × 51,2 Go/s | **614 Go/s** |
| SP5 Turin (12× 5600) | 12 × 44,8 Go/s | 538 Go/s |
| Xeon 6900P (12× 6400) | 12 × 51,2 Go/s | **614 Go/s** |
| Xeon 6900P (12× MRDIMM-8800) | 12 × 70,4 Go/s | **845 Go/s** |
| Xeon 6700P (8× 6400) | 8 × 51,2 Go/s | 410 Go/s |
| Xeon 6+ (12× 8000) | 12 × 64 Go/s | **768 Go/s** |
| SP7 Venice (16× 8000) | 16 × 64 Go/s | **1 024 Go/s** |
| SP7 Venice (16× MRDIMM-12800) | 16 × 102,4 Go/s | **1 638 Go/s** |

Débit soutenu mesuré (STREAM) : comptez **70-80 %** de ces valeurs. Le reste
part en overhead de protocole et en contention.

## 143. Population avancée : schémas par plateforme

**SP5 (12 canaux, 1 DPC)** — le cas standard :
```
 12 barrettes identiques, une par canal -> 768 Go (12x64), pleine BP
```
**SP5 (24 slots, 2 DPC)** — quand 1,5 To+ sont requis :
```
 24 barrettes identiques -> frequence reduite d'1 a 2 crans (section 47)
 Alternative : 12x 128 Go en 1 DPC si la carte a 12 slots
```
**LGA-4710 (8 canaux)** — Xeon 6700 :
```
 8 barrettes identiques -> 512 Go (8x64). Ne JAMAIS mettre 6 ou 10.
```
**LGA-7529 (12 canaux)** — Xeon 6900 : même logique que SP5.

Règle d'or : **le nombre de barrettes = un multiple du nombre de canaux.**
Tout autre nombre = canaux asymétriques = bande passante perdue.

## 144. Rangs et fréquences : le tableau des compromis

| Config par canal | Fréquence atteignable (indicatif) | Note |
|---|---|---|
| 1× 1R | max JEDEC | idéal |
| 1× 2R | max JEDEC | idéal (capacité ++) |
| 2× 1R | -1 cran | acceptable |
| 2× 2R | -1 à -2 crans | dernier recours |
| 1× MRDIMM | 8800 (si supporté) | HPC/IA |

En pratique : à capacité totale égale, **12× 2R en 1 DPC > 24× 1R en
2 DPC** — même débit par barrette, mais fréquence max conservée et 12 slots
libres pour l'avenir.

## 145. Mémoire persistante (PMem) : état en 2026

La mémoire persistante (ex-Optane, abandonnée par Intel) n'a pas de
successeur grand public au 27/09/2026. Les alternatives :
- **CXL-attached memory** : extension de capacité avec persistance possible
  selon le device (section 56) — l'héritier logique, en déploiement
  hyperscaler.
- **NVMe + DAX** : persistance via stockage, latence en µs (pas en ns).
- **Batteries + DRAM** (NVDIMM-N) : niche, coûteux.

Verdict : **ne dimensionnez pas de PMem en 2026.** Si votre application
exigeait Optane, migrez vers CXL (trajectoire) ou vers plus de DRAM.

## 146. Diagnostic mémoire : méthode pas à pas

Symptômes : erreurs ECC corrigées en hausse, crashs aléatoires, corruption
de données, VM qui « plantent » sans raison.

1. **Lire les compteurs** : BMC/IPMI → SEL → erreurs ECC par DIMM. La
   barrette fautive est presque toujours identifiée.
2. **Isoler** : si plusieurs barrettes signalent, suspectez le **slot ou le
   canal** (poussière, mauvais contact), pas les barrettes.
3. **Permuter** : déplacez la barrette suspecte sur un autre canal. Si
   l'erreur suit la barrette → barrette HS. Si elle reste → slot/carte mère.
4. **Tester** : memtest86+ (hors production) ou test intégré du BIOS sur
   la barrette isolée.
5. **Remplacer** : par la **référence exacte** (QVL), puis revérifier la
   symétrie (section 45).

**Ne jamais** : « attendre de voir » avec des erreurs ECC qui augmentent —
c'est le signe avant-coureur d'une erreur non corrigible.

## 147. Tests d'injection d'erreurs ECC

Pour valider que la chaîne ECC fonctionne **avant** la production :
- Certains BIOS/BMC proposent l'**injection d'erreurs** (1 bit, 2 bits)
  pour vérifier la correction, la détection et la remontée d'alerte.
- Procédure : injectez 1 bit → vérifiez la correction silencieuse et le
  compteur ; injectez 2 bits → vérifiez l'alerte (pas le crash en test).
- À faire **une fois** à la réception de chaque modèle de serveur, pas en
  routine.

C'est le seul moyen de prouver que votre supervision mémoire est
opérationnelle avant le jour où elle comptera vraiment.

## 148. Firmware des barrettes et SPD

Chaque barrette embarque un SPD (Serial Presence Detect) : timings,
capacité, fabricant. Points d'attention :
- Des barrettes **contrefaites** (SPD reprogrammé) circulent sur le marché
  gris : achetez via des canaux agréés, vérifiez les étiquettes et les
  lots.
- Après une mise à jour BIOS, **revérifiez la fréquence négociée** (dmidecode
  sous Linux) : certaines versions « sécurisent » en baissant la fréquence
  des références non QVL.
- `dmidecode -t memory` + les compteurs EDAC du kernel (`/sys/devices/
  system/edac/mc`) : vos deux sources de vérité sous Linux.

## 149. DDR5 et température : le point de vigilance

La DDR5 consomme plus par barrette que la DDR4 à capacité égale (puces plus
denses, PMIC embarqué qui chauffe). Repères (indicatif) :
- Température DIMM normale : 40-60 °C ; alerte à ~75 °C ; erreurs au-delà
  de ~85 °C.
- En 2 DPC et avec MRDIMM, ajoutez 5-10 °C : prévoyez le flux d'air
  (shroud en place, section 64) et surveillez les sondes.
- Le PMIC (circuit d'alimentation de la barrette) est un nouveau point
  chaud : une barrette qui chauffe anormalement au repos peut avoir un
  PMIC défectueux.

## 150. Tableau récapitulatif : choisir sa RAM en 5 questions

| Question | Si oui | Si non |
|---|---|---|
| > 1,5 To/socket ? | LRDIMM/3DS | RDIMM standard |
| Workload sensible à la BP ? | fréquence max + 1 DPC | 5600, budget capacité |
| HPC/IA avec Xeon 6900 ? | évaluer MRDIMM-8800 (+100 $/barrette, +énergie) | RDIMM-6400 |
| Budget serré ? | 64/96 Go (sweet spot), tout en une fois | — |
| Extension future probable ? | 1 DPC + slots libres, ou grosses barrettes | remplir au besoin |

---

## 151. Cas chiffré : cluster 5 nœuds virtualisation (indicatif)

5 nœuds 2U 2S EPYC (1100 W max, 350 W idle, 30 % de charge → ~550 W moyens).

| Poste | Calcul | Résultat |
|---|---|---|
| IT moyen | 5 × 550 W | 2,75 kW |
| IT max | 5 × 1100 W | 5,5 kW |
| Soutiré (PUE 1,5) | 2,75 × 1,5 | **4,1 kW** |
| Énergie annuelle | 4,1 × 8760 | **36 000 kWh** |
| Facture (0,18 €/kWh) | — | **~6 500 €/an** |
| Onduleur requis | 5,5 × 1,25 / 0,9 | **~7,6 kVA → 10 kVA** |
| PDU | < 10 kW | 2× tri 16 A suffisent (11 kW) |

Sur 5 ans : ~32 500 € d'énergie pour ~125 000 € de CAPEX serveurs. Le ratio
énergie/CAPEX (~26 %) est typique d'un cluster virtualisation bien rempli.

## 152. Cas chiffré : salle complète de 42U (indicatif)

Remplissage : 20× 2U 2S à 1100 W max = 22 kW IT max, ~11 kW moyens.

| Poste | Calcul | Résultat |
|---|---|---|
| IT max | 20 × 1100 W | **22 kW** |
| IT moyen (30 %) | — | ~11 kW |
| Froid requis (PUE 1,5) | 11 × 0,5 | 5,5 kW frigorifiques moyens |
| Soutiré total | 11 × 1,5 | **16,5 kW** |
| Énergie annuelle | 16,5 × 8760 | **144 500 kWh** |
| Facture | — | **~26 000 €/an** |
| Onduleur | 22 × 1,25 / 0,9 | **~30 kVA triphasé** |
| PDU | 22 kW | 2× tri 32 A par baie |
| Groupe électrogène | (22 + froid) × 1,5 | **~50 kVA** |

C'est à cette échelle que l'énergie devient un sujet de direction : 130 000 €
sur 5 ans, un groupe de 50 kVA, une vraie salle avec confinement.

## 153. Câbles et disjoncteurs : dimensionnement (indicatif)

Règles électriques (normes locales à vérifier par un électricien qualifié) :

| Charge | Section cuivre (indicatif) | Disjoncteur |
|---|---|---|
| PDU mono 16 A | 2,5 mm² | 16 A courbe C/D |
| PDU mono 32 A | 6 mm² | 32 A |
| PDU tri 16 A | 5G2,5 mm² | 16 A tétrapolaire |
| PDU tri 32 A | 5G6 mm² | 32 A tétrapolaire |
| Arrivée onduleur 20 kVA | 5G10 mm² | selon constructeur |

