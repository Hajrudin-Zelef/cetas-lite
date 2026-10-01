---
id: collect-261001-rattrapage/rattrapage/datacenter-cpu-ram-guide-23
title: "CPU SERVEUR, RAM SERVEUR, CHÂSSIS & VENTILATION"
domain: rattrapage
role: reference
task: reference
actors: ["AMD", "Intel"]
dates: ["2026-09-27"]
keywords: ["amd", "arr", "capex", "clearwater forest"]
source: docs/RAG/collect-261001-rattrapage/datacenter_cpu_ram_guide.md
source_anchor: ""
source_lines: [3697, 3888]
sha256: 2884609b5795de53053fe9cbbd4ba514e416bc3df3a395e9af8f1858f4609ba0
---

# CPU SERVEUR, RAM SERVEUR, CHÂSSIS & VENTILATION

```
FICHE SERVEUR
  Nom ............  U baie ............  Salle ............
  Modele .........  N° serie ...........  Garantie jusqu'au ............
  CPU ............  x ....  (ref, TDP)
  RAM ............  .... x .... Go .... (ref QVL, 1R/2R, frequence)
  Disques ........  .... x .... (modele, FW)
  Reseau .........  .... (cartes, MAC)
  BMC ............  IP ............  FW ............
  BIOS ...........  version ............  date ............
  PDU ............  A prise ....  B prise ....
  Role ...........  Criticite ............
  OS .............  Hyperviseur ............
  Supervision ....  [ ] Zabbix  [ ] PDU  [ ] NUT
  Interventions :
    date ........  action ..........................  par ............
```

Une fiche par serveur, à jour. C'est le document que l'on sort en premier
devant toute panne — et celui que l'auditeur demande en premier.

## 210. Annexe : template de comparatif de devis

| Critère | Offre A (OEM 1) | Offre B (OEM 2) | Offre C (intégrateur) |
|---|---|---|---|
| Config exacte (ref CPU/RAM) | | | |
| QVL mémoire confirmée | oui/non | | |
| Redondance fans N+1 (écrit) | | | |
| PSU (puissance, 80 PLUS) | | | |
| Garantie (durée, SLA) | | | |
| CAPEX | | | |
| Licences 5 ans | | | |
| Énergie 5 ans (estimée) | | | |
| Maintenance 5 ans | | | |
| **TCO 5 ans** | | | |
| Délai de livraison | | | |
| Reprise ancien parc | | | |

**On ne compare jamais des CAPEX seuls** : remplissez la ligne TCO avant
toute décision (sections 101, 102).

## 211. Lexique anglais-français des termes courants

| Anglais | Français | Section |
|---|---|---|
| Air shroud | déflecteur d'air | 64 |
| Blank / blanking panel | cache de baie | 63, 112 |
| Burn-in | test de rodage à réception | 187 |
| Cable management arm | bras de câbles | 74, 79 |
| Derating | déclassement (en température/altitude) | 67 |
| Hot-swap | remplaçable à chaud | 58 |
| Power capping | plafonnement de puissance (cTDP) | 133 |
| Rack unit (U) | unité de baie (44,45 mm) | 69 |
| Rails (ball-bearing) | rails télescopiques à roulement | 74 |
| Throttling | bridage thermique | 67 |
| Top-of-rack (ToR) | commutateur haut de baie | 193 |
| Wattmeter | wattmètre | 85 |
| Write endurance (DWPD) | endurance en écriture | 188 |

---

*Fin du guide — 211 sections. Rédigé et vérifié le 27/09/2026.*
*Pour Zelef — chiffres, BOMs, énergie. Les faits datés se revérifient,*
*les méthodes restent.*

---

## 212. Comparatif synthétique final : Turin vs Xeon 6 vs Xeon 6+

Tableau de décision ultime (données vérifiées le 27/09/2026, prix publics
au lancement, performances = ordres de grandeur à valider par POC) :

| Critère | EPYC 9005 Turin | Xeon 6 6900P | Xeon 6+ Clearwater Forest |
|---|---|---|---|
| Cœurs max/socket | 192 (Zen 5c) | 128 (P-core) | 288 (E-core Darkmont) |
| Threads max | 384 (SMT) | 256 (HT) | 288 (pas de HT) |
| Fréquence boost max | 5,0 GHz (réf. F) | ~3,9 GHz | non publié |
| TDP max | 500 W | 500 W | 450 W |
| Canaux mémoire | 12 | 12 | 12 |
| Débit mémoire max | DDR5-6400 | MRDIMM-8800 | DDR5-8000 |
| BP théorique max | 614 Go/s | 845 Go/s | 768 Go/s |
| PCIe | Gen5, 128 lignes | Gen5, 96 lignes | Gen5, 96 lignes |
| CXL | 2.0 | 2.0 | 2.0 |
| Accélérateurs IA | AVX-512/VNNI | **AMX** + DSA/IAA/QAT/DLB | (hérités E-core) |
| Socket | SP5 (mature) | LGA 7529 (mature) | LGA 7529 (compatible) |
| Prix amiral | 14 813 $ (9965) | 17 800 $ (6980P) | 14 995 $ (6990E+) |
| $/cœur amiral | ~77 $ | ~139 $ | ~52 $ |
| Point fort | densité + maturité SP5 | BP mémoire + AMX | cœurs/watt, DDR5-8000 |
| Point faible | pas d'AMX | prix, conso MRDIMM | E-core : perf/thread |

Lecture : **il n'y a pas de « meilleur CPU » absolu.** Il y a le meilleur
CPU pour votre workload, vos licences et votre salle. Ce tableau + la
section 195 donnent la réponse en 5 minutes ; le POC la confirme en
2 semaines.

## 213. Nomenclature OEM : s'y retrouver (bref)

Chaque constructeur renomme les plateformes. Repères (indicatif, gammes
2024-2026) :

| OEM | 1U 2S | 2U 2S | Logique |
|---|---|---|---|
| Dell PowerEdge | R660/R6615 | R760/R7615 | R = rack, 6 = génération |
| HPE ProLiant | DL360 | DL380 | DL = rack dense |
| Lenovo ThinkSystem | SR630 | SR650 | SR = rack |

Règles :
- Le **chiffre des centaines** = la génération (x60 = 2023+, à vérifier
  par OEM).
- Les références « 15 » (ex. R6615) = **mono-socket AMD**, « 45 » = bi-socket.
- **Ne comparez jamais des prix entre OEM sans aligner les options**
  (BMC, garantie, rails, cache) : 2 000 € d'écart viennent souvent des
  options, pas du châssis.

## 214. Checklist de mise en service (jour J)

**Avant allumage :**
- [ ] Baie : profondeur, PDU A/B branchées, phases équilibrées (section 90).
- [ ] Caches de baie posés, allées froide/chaude respectées.
- [ ] Câbles : longueurs justes, chemins séparés, étiquetés (section 166).

**Premier boot :**
- [ ] BIOS/BMC à jour (section 104), réglages documentés (section 181).
- [ ] RAM : `dmidecode` → 12/12 canaux, fréquence attendue, ECC actif.
- [ ] Ventilateurs : tous présents, N+1 vérifié, courbe adaptée.
- [ ] Disques : tous vus, firmware à jour, SMART propre.
- [ ] Réseau : liens négociés au bon débit (`ethtool`), 2 chemins testés.

**Recette (section 187) :**
- [ ] memtest86+, stress-ng 1 h, fio, iperf3 : résultats archivés.
- [ ] Supervision : BMC, PDU, ECC remontés dans Zabbix/Prometheus.
- [ ] NUT : arrêt ordonné testé (section 87).
- [ ] Fiche serveur créée (section 209), photo de la baie archivée.

**Un serveur mis en production sans passer cette checklist est un serveur
dont on ne connaît ni l'état initial, ni la configuration réelle.**

## 215. Modèle de calcul TCO : colonnes du tableur

Pour construire votre comparatif chiffré (1 ligne par offre) :

```
A. CAPEX
   A1 Serveurs (chassis + CPU + RAM + disques + cartes)
   A2 Rails, PDU, câbles (quote-part)
   A3 Onduleur (quote-part au prorata des kW)
   A4 Installation / mise en service
B. OPEX annuel x 5 ans
   B1 Energie : P_moy x 8760 x PUE x prix_kWh
   B2 Froid : inclus via PUE (ou ligne separee si PUE=1)
   B3 Licences logicielles (annuel x 5)
   B4 Maintenance / garantie (extension annees 4-5)
   B5 Administration (heures x cout horaire)
   B6 Batteries onduleur (remplacement an 4-5, quote-part)
C. Fin de vie
   C1 Destruction certifiee des disques
   C2 Recyclage DEEE / revente (negatif si revente)
TCO = A + B + C
```

Renseignez **toutes** les lignes, même à zéro (zéro documenté ≠ oubli).
Le comparatif qui gagne est rarement celui qu'on croyait au départ —
c'est tout l'intérêt de l'exercice.

## 216. Pour finir : les 5 décisions qui comptent vraiment

Après 216 sections, tout se résume à cinq décisions :

1. **Le workload** : latence ou débit ? (Tout le dimensionnement en découle.)
2. **Les licences** : chiffrées avant le CPU. (Le poste n°1 du TCO.)
3. **La mémoire** : 12 canaux remplis, QVL, 1 DPC. (Le poste n°1 des perfs.)
4. **L'énergie** : pic + 25 % pour l'onduleur, PUE suivi, TCO 5 ans.
   (Le poste n°1 des surprises.)
5. **L'exploitation** : supervision jour 1, firmwares par vagues, tests
   annuels. (Le poste n°1 de la tranquillité.)

Le reste — la marque du CPU, la fréquence exacte de la RAM, le modèle de
la PDU — est de l'optimisation. Ces cinq décisions sont de la stratégie.
Prenez-les bien, dans cet ordre, et le reste suivra.

---

*Fin du guide — 216 sections. Rédigé et vérifié le 27/09/2026.*
*Pour Zelef — chiffres, BOMs, énergie. Les faits datés se revérifient,*
*les méthodes restent.*

---

## 217. Ressources et références web (vérifiées le 27/09/2026)

Pour revérifier ou approfondir les faits datés de ce guide :

