---
id: collect-261001-rattrapage/rattrapage/datacenter-reseau-guide-22
title: "RÉSEAU DATACENTER / IA — NICs, FIBRE, SPINE-LEAF, SWITCHES IA, DPU, CDN/EDGE"
domain: rattrapage
role: reference
task: reference
actors: ["Nvidia"]
dates: ["2026-09-27"]
keywords: ["datacenter", "capex", "cpo", "dci", "distribution", "gpu", "memory", "nvidia", "serdes"]
source: docs/RAG/collect-261001-rattrapage/datacenter_reseau_guide.md
source_anchor: ""
source_lines: [3169, 3315]
sha256: 32c37cbd347e4cf416e74fefeb119767391b4a0148168aa2918e92dac33a3e14
---

# RÉSEAU DATACENTER / IA — NICs, FIBRE, SPINE-LEAF, SWITCHES IA, DPU, CDN/EDGE

| Câble | Débit | Longueur | Conso | Usage | Prix (⚠️) |
|---|---|---|---|---|---|
| DAC SFP28 25G 3 m | 25G | 3 m | 0,1 W | Serveur→ToR | 30-60 € |
| DAC QSFP28 100G 3 m | 100G | 3 m | 0,3 W | Intra-rack | 50-120 € |
| DAC QSFP28 100G 5 m | 100G | 5 m | 0,3 W | Racks adjacents | 60-150 € |
| AOC QSFP28 100G 20 m | 100G | 20 m | 1 W | Inter-racks | 200-400 € |
| DAC QSFP-DD 400G 2 m | 400G | 2 m | 0,5 W | Intra-rack | 150-350 € |
| AEC QSFP-DD 400G 5 m | 400G | 5 m | 3 W | Racks adjacents | 400-800 € |
| AOC QSFP-DD 400G 30 m | 400G | 30 m | 2 W | Inter-rangées | 600-1 200 € |
| DAC OSFP 800G 2 m | 800G | 2 m | 0,5 W | Intra-rack | 400-900 € |
| AEC OSFP 800G 5 m | 800G | 5 m | 4 W | Racks adjacents | 1 000-2 000 € |
| AOC OSFP 800G 30 m | 800G | 30 m | 3 W | Inter-rangées | 1 500-3 000 € |
| Breakout 400G→4×100G 2 m | 400G | 2 m | 0,5 W | Migration | 200-500 € |
| Breakout 800G→2×400G 2 m | 800G | 2 m | 0,5 W | Migration | 500-1 200 € |
| Jarretière OM4 LC 5 m | — | 5 m | — | Brassage | 15-30 € |
| Jarretière OM4 MPO-12 10 m | — | 10 m | — | Trunk court | 40-80 € |
| Jarretière OS2 LC 20 m | — | 20 m | — | Brassage | 25-50 € |
| Trunk OS2 MPO-24 50 m | — | 50 m | — | Rocade | 150-400 € |
| Trunk OM4 MPO-12 30 m | — | 30 m | — | Rocade | 100-250 € |
| Cat6A RJ45 10 m | 10G | 10 m | — | OOB | 10-20 € |
| Fibre armée OS2 100 m | — | 100 m | — | Chemin exposé | 300-800 € |
| Câble console RJ45 | — | 3 m | — | OOB | 5-10 € |

## 155. Protocole d'acceptation détaillé — 40 points de contrôle

**Physique (1-10)** : [ ] baies fixées [ ] airflow cohérent [ ] PDU A/B
branchées [ ] disjoncteurs calibrés [ ] terre < 5 Ω [ ] étiquettes TIA-606
[ ] goulottes ≤ 50 % [ ] bouchons sur ports libres [ ] kit nettoyage présent
[ ] sondes température/humidité.

**Câblage (11-20)** : [ ] DAC ≤ portée [ ] MPO polarité documentée
[ ] budgets optiques calculés [ ] puissances Rx mesurées [ ] marges ≥ 3 dB
[ ] BER pré-FEC < 1e-9 [ ] 0 FEC uncorrectable [ ] breakout validés
[ ] longueurs conformes au plan [ ] photos du brassage archivées.

**Réseau (21-30)** : [ ] BGP up partout [ ] BFD < 300 ms [ ] ECMP ±10 %
[ ] MTU 9000 bout-en-bout [ ] ECN/DCQCN actifs [ ] PFC par file seule
[ ] EVPN : routes échangées [ ] anycast gateway ping [ ] OOB joignable
[ ] AAA fonctionnel.

**Charge (31-40)** : [ ] ib_write_bw ≥ 95 % [ ] NCCL ≥ 80 % [ ] latence
p99 < 5 µs [ ] perte 1 spine < 5 % [ ] perte 1 leaf OK [ ] burn-in 72 h
[ ] thermique < 75 °C [ ] conso mesurée vs dimensionnée [ ] rollback testé
[ ] doc d'exploitation livrée.

**Signature** : __________  Date : __________  Réserves : __________

## 156. Modèle de RFP réseau — structure complète

1. **Contexte** : sites, effectifs, workloads (IA ? cloud ?), horizon 3 ans.
2. **Périmètre** : backend IA / front-end / DCI / edge — lots séparés.
3. **Exigences techniques** : débits, ratio (1:1 / 3:1), RDMA/UEC, buffers,
   ECN/PFC, breakout, MTU, telemetry (gNMI), NOS, ZTP.
4. **Exigences physiques** : format (2U/1U), airflow, PSU 1+1, conso max
   **avec optiques**, niveau sonore (bureaux proches ?).
5. **Interopérabilité** : UEC-ready exigé, optiques tierces acceptées,
   standards ouverts (SONiC si whitebox).
6. **Tests** : protocole d'acceptation §155 imposé, lab de validation.
7. **SLA** : 4 h 24/7 (IA) / J+1 (standard), spares, TAC nommé.
8. **Formation** : 2 sessions admin + 1 troubleshooting, en français.
9. **Prix** : CAPEX détaillé par poste, OPEX 5 ans (maintenance, énergie),
   options CPO / 1,6T.
10. **Planning** : jalons, pénalités de retard, clause de sortie.

**Critères de notation** (exemple) : technique 40 %, TCO 5 ans 30 %,
SLA/support 20 %, énergie 10 %. **Ne jamais noter sur le seul CAPEX.**

## 157. Pannes typiques — fréquence et MTTR

Ordres de grandeur constatés (⚠️ retours d'expérience, varient par site) :

| Panne | Fréquence | MTTR avec spares | MTTR sans spares |
|---|---|---|---|
| Optique / jarretière sale | Mensuelle | 15 min (nettoyage) | 4 h (attente pièce) |
| DAC défectueux | Trimestrielle | 10 min | 4 h |
| Ventilateur switch | Annuelle | 5 min (hot-swap) | J+1 |
| PSU switch | Annuelle | 5 min (hot-swap) | J+1 |
| Switch complet | Tous les 3-5 ans | 1-4 h (RMA + reconfig) | J+1 à J+7 |
| Bug NOS (crash) | Rare | 30 min (reboot + rollback) | — |
| Erreur humaine (config) | Mensuelle | 15 min (rollback Git) | Heures |
| Coupure électrique | Annuelle | 0 (UPS) à heures (sans) | — |
| Fibre sectionnée (travaux) | Rare | 2-8 h (épissure) | Jours |

**Lecture** : 80 % des pannes = optique/câble/humain, MTTR < 30 min
**avec** spares et procédures. Le stock de spares (§118) est l'assurance
la moins chère du datacenter.

## 158. Glossaire 2 — 40 termes complémentaires

| Terme | Définition |
|---|---|
| Anycast | Même IP annoncée depuis N sites (BGP choisit le plus proche). |
| Atténuation | Perte de puissance en dB le long de la fibre. |
| Backplane | Fond de panier : liaisons électriques internes au châssis. |
| BFD | Détection de panne en ~300 ms (vs 180 s BGP). |
| Breakout (câble) | 1 connecteur rapide → N connecteurs lents. |
| Burn-in | Charge continue 72 h avant mise en production. |
| CDU | Coolant Distribution Unit (water-cooling). |
| Châssis | Switch modulaire à cartes (vs fixe 1-2U). |
| Cladding | Gaine optique autour du cœur de la fibre. |
| Cœur (fibre) | Zone centrale qui guide la lumière (9 ou 50 µm). |
| Deadlock PFC | Boucle de pauses qui fige le fabric. |
| Dispersion | Étalement du signal (limite la distance). |
| Downlink | Vers les serveurs (vs uplink vers les spines). |
| EMI | Interférences électromagnétiques. |
| Épissure | Soudure de 2 fibres (0,05-0,1 dB). |
| ER | 40 km (optique). |
| Fan-out | Répartition d'un signal vers N destinations. |
| Férule | Embout céramique du connecteur (2,5/1,25 mm). |
| Flap | Lien qui monte/descend en boucle. |
| Golden config | Configuration de référence validée. |
| Hot-swap | Remplacement à chaud (PSU, ventilos). |
| In-network computing | Calcul dans le switch/NIC (ex. AllReduce). |
| Insertion loss | Perte d'insertion d'un connecteur (dB). |
| Jumbo frame | MTU 9000 (vs 1500 standard). |
| Latence p99 | Latence sous laquelle sont 99 % des paquets. |
| Link budget | Budget optique : Tx − sensibilité Rx. |
| Microburst | Pic de trafic µs-ms invisible au polling. |
| MTTR | Mean Time To Repair. |
| NUMA | Non-Uniform Memory Access (topologie CPU). |
| Oversubscription | Ratio downlink/uplink. |
| Packet spraying | Répartition paquet par paquet. |
| Polarisation (ECMP) | Hash identique → un lien saturé. |
| Radix | Nombre de ports d'un switch. |
| Retimer | Régénère le signal électrique (vs redriver). |
| SerDes | Serializer/Deserializer. |
| Réordonnancement | Remise en ordre des paquets (coûteux pour RDMA). |
| Snake (câble) | Testeur de continuité simplifié : à bannir sur fibre. |
| Store-and-forward | Commutation après réception complète (vs cut-through). |
| Trunk (fibre) | Câble multi-brins pour rocades. |
| Uplink | Vers les spines (vs downlink). |
| Xon/Xoff | Contrôle de flux historique (vs PFC). |

## 159. Fiches switches complémentaires — 5 acteurs

**Cisco (Nexus 9000, Silicon One G200)** : l'existant enterprise. NX-OS
mature, ACI pour l'overlay managé. G200 (51,2T, 800G) = l'entrée de Cisco
dans l'IA — ⚠️ détails non vérifiés au 27/09/2026, exiger une fiche à jour.
Verdict : rester si le parc est Cisco et l'équipe formée ; sinon Arista.

**Juniper (QFX5240, PTX)** : solide en SP et DCI, Apstra pour l'intent-based.
Moins « IA-first » qu'Arista/NVIDIA en 2026. Verdict : pertinent si
l'équipe est Juniper, ou pour le DCI.

**HPE Aruba CX** : enterprise, bon support, pas le backend IA. Verdict :
front-end et campus, pas les clusters GPU.

