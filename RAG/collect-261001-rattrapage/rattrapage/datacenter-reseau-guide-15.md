---
id: collect-261001-rattrapage/rattrapage/datacenter-reseau-guide-15
title: "RÉSEAU DATACENTER / IA — NICs, FIBRE, SPINE-LEAF, SWITCHES IA, DPU, CDN/EDGE"
domain: rattrapage
role: reference
task: reference
actors: []
dates: []
keywords: ["arr", "asic", "attention", "capex", "distribution", "ethernet", "gpu", "training"]
source: docs/RAG/collect-261001-rattrapage/datacenter_reseau_guide.md
source_anchor: ""
source_lines: [2090, 2242]
sha256: ead67c91ca558399afceaa4127b29dff2e65fe59bab518b491e617ac29797a58
---

# RÉSEAU DATACENTER / IA — NICs, FIBRE, SPINE-LEAF, SWITCHES IA, DPU, CDN/EDGE

**Conséquences** :
- Serveurs **courts** (≤ 600 mm) et switchs **DC −48 V** si site télécom —
  vérifier les options d'alimentation à la commande (P44, §122).
- **Tout en double** : pas de « on ira voir demain » à 300 km. Spares sur
  site (1 switch, 2 optiques, câbles).
- **OOB 4G/5G** : un routeur cellulaire pour l'accès distant quand le lien
  principal tombe — le seul moyen de réparer à distance.
- **Filtres à air** : en shelter, les filtres se colmatent → surchauffe
  lente. Maintenance trimestrielle, pas annuelle.

## 116. Water-cooling et réseau : cohabitation

En 2026, les racks GPU sont à eau (plaques froides), les switchs restent
à air. Trois architectures :

| Architecture | Principe | Note |
|---|---|---|
| Salles séparées | Réseau dans une salle air, GPU dans salle eau | Simple, + fibre inter-salles |
| Allées mixtes | Baies réseau à air entre rangées eau | Gérer les flux d'air (confinement) |
| CDU + air | Unité de distribution d'eau + clim d'appoint | Le plus courant en retrofit |

**Points d'attention** :
- **Fuite d'eau** : jamais de tuyauterie **au-dessus** des baies réseau —
  les switchs n'ont pas de bac de rétention. Cheminement latéral ou sous
  plancher technique avec détection de fuite.
- **Humidité** : < 60 % HR — la condensation sur les optiques = corrosion
  des contacts. Sonde d'humidité par salle.
- **Vannes** : pouvoir isoler hydrauliquement une rangée sans couper le
  réseau (le réseau survit aux GPU, pas l'inverse en supervision).

## 117. Dimensionner l'onduleur pour le lot réseau (lien guide onduleurs)

Le réseau = **charge prioritaire** : il porte la supervision, l'OOB et les
ordres d'arrêt (NUT). Règles (complément du guide onduleurs_ups_guide.md) :

| Paramètre | Valeur |
|---|---|
| Autonomie visée | **≥ 15 min** (le temps d'un arrêt propre des baies GPU) |
| Puissance | Max switchs + NICs + optiques + 30 % (pas le typique) |
| Batches | 1 onduleur par **plan réseau** (pas 1 pour tout) : la panne d'un UPS ne doit pas couper 2 plans |
| Bypass | Bypass de maintenance **externe** : on peut shunter l'UPS sans couper |
| Test | Test batterie **trimestriel** avec le réseau en charge réelle |

**Calcul** : pod §62 = 181 kW réseau → 2 UPS 120 kVA (N+1) ou 4× 60 kVA
(2 par plan). **Ne pas mettre le réseau sur le même UPS que les GPU** :
un training qui fait disjoncter son UPS ne doit pas emporter la supervision
(P45, §122).
**Séquence d'arrêt** (coupure longue) : 1. Arrêt GPU (gros consommateurs),
2. Stockage (flush), 3. **Réseau en dernier** (il faut superviser 1 et 2),
4. Management/OOB = le tout dernier (il envoie l'alerte « tout est éteint »).

## 118. Contrats de maintenance réseau : SLA, spares, escalade

| Élément | Standard DC | Backend IA |
|---|---|---|
| SLA remplacement | J+1 ouvré | **4 h, 24/7** |
| Spares sur site | 1 switch, 5 % optiques | **1 switch par plan**, 10 % optiques, câbles |
| Support | Heures ouvrées | 24/7 avec accès lab (repro) |
| Firmware | Trimestriel | Figé par campagne (pas de « latest ») |
| Escalade | Niveau 1 → 2 | Accès direct niveau 3 (TAC) sur P1 |

**Chiffrage** : maintenance 24/7 4 h ≈ **12-18 % du CAPEX/an** (⚠️) vs
8-12 % en J+1. Sur un fabric à 5 M€, l'écart = 200-300 k€/an — à comparer
au coût d'une heure de cluster idle (6-7 chiffres, §22 du brief initial).
**Pièce critique** : les **optiques** — 80 % des pannes sont optique/câble,
pas switch. Un spare d'optiques bien garni vaut plus qu'un switch de spare.
**Contrat TAC** : exiger un numéro direct + un ingénieur nommé pour les
comptes IA (les files génériques noient les P1).

## 119. Retours terrain — 5 histoires (anonymisées, ⚠️ retours d'expérience)

**R1. Le DAC de 5 m qui « devrait passer ».**
Cluster 400G, DAC de 5 m commandés « parce que c'était moins cher que
l'AOC ». La moitié des liens ne montaient pas (limite du DAC 400G : 3 m).
3 semaines de retard, 200 AOC en urgence à prix ×2. **Leçon** : respecter
les portées (§59), le cuivre ne pardonne pas.

**R2. La poussière de chantier.**
Nouveau DC, brassage fait pendant les travaux. 40 % des liens 100G avec
BER pré-FEC > 1e-6. Cause : poussière de plâtre dans les MPO. 2 jours de
nettoyage systématique (§24). **Leçon** : brasser **après** la fin du
chantier, bouchons partout entre-temps.

**R3. Le firmware qui a tué un training.**
Mise à jour FW NIC « mineure » sur 1/4 du cluster un vendredi soir.
Le lundi : NCCL −30 % sur ces nœuds (changement de comportement CC).
Rollback le mardi, 4 jours de GPU à 70 %. **Leçon** : FW figé par campagne
(§13), jamais de mise à jour partielle sur un fabric.

**R4. L'ECMP qui polarisait.**
Fabric RoCE 3 tiers, 2 spines sur 8 qui saturaient, les autres à 30 %.
Cause : même hash seed sur tous les leafs (défaut constructeur). Seed
aléatoire → répartition ±5 %. **Leçon** : tester l'ECMP sous charge
(§107), pas à vide.

**R5. L'onduleur du réseau.**
Coupure longue, les GPU s'arrêtent proprement… puis plus rien : le switch
de management était sur le même UPS que les GPU, parti en même temps.
Arrêt à l'aveugle, 2 baies non redémarrées proprement. **Leçon** : réseau
sur UPS prioritaire séparé (§117).

## 120. Erreurs de débutant vs erreurs d'expert

| Débutant | Expert (plus coûteuses) |
|---|---|
| Oublier le /31, gaspiller des /24 | Sur-optimiser l'oversubscription « parce que ça passe en moyenne » |
| Brancher sans nettoyer la fibre | Faire confiance au FEC sans superviser le pré-FEC |
| Acheter la NIC avant de vérifier le PCIe | Choisir le switch sur l'ASIC sans évaluer le NOS |
| Pas d'étiquettes | Automatiser avant de stabiliser |
| Un seul plan d'alimentation | Un seul fabric pour IA + prod « pour mutualiser » |
| Tester à vide seulement | Ne pas tester la perte d'un spine sous charge |
| MTU 1500 partout « par défaut » | MTU 9000 sans vérifier **tout** le chemin (un équipement à 1500 = fragmentation) |

**La plus chère** : le fabric unique IA+prod. Elle économise 15 % de CAPEX
et coûte des semaines de debug par an + des trainings ralentis en
permanence. **Deux fabrics valent toujours le coup.**

## 121. Quiz — 10 questions + réponses

**Q1.** Une NIC 400G est installée dans un slot PCIe Gen4 x16. Quel débit
maximal peut-elle réellement atteindre, et pourquoi ?
**R.** ~200 Gb/s (la moitié). PCIe Gen4 x16 = 256 Gb/s théoriques ; la NIC
ne peut pas pousser plus que le bus. C'est le bridage invisible P1 (§8).

**Q2.** Quelle est la différence fondamentale entre RoCEv2 et Ultra Ethernet
(UET) concernant le réseau ?
**R.** RoCEv2 exige un réseau lossless (PFC+ECN à régler) et des paquets en
ordre ; UET est conçu pour un réseau lossy : livraison non ordonnée,
retransmission sélective, multipath natif (§10, §54).

**Q3.** Un lien 400G-DR4 « monte » mais affiche un BER pré-FEC de 5e-6.
Que faire en premier ?
**R.** Nettoyer les connecteurs (inspecter → nettoyer → ré-inspecter, §24),
pas changer le module. 5e-6 = 🟡 : surveiller après nettoyage ; si ça
persiste, vérifier le budget optique (§20).

**Q4.** Calculez l'oversubscription d'un leaf 48×25G avec 6×100G d'uplinks.
**R.** Down : 48×25 = 1200 Gb/s. Up : 6×100 = 600 Gb/s. Ratio = 2:1 (§32).

**Q5.** Avec des switchs 64 ports, combien de serveurs en 2 tiers
non-bloquant, et combien de spines ?
**R.** n²/4 = 1024 serveurs, n/2 = 32 spines (§33).

**Q6.** Pourquoi le rail-optimized sépare-t-il les plans par indice de NIC ?
**R.** Parce que le trafic dominant (all-reduce NCCL en anneau) circule
entre les mêmes indices de GPU : séparer les plans évite la contention
inter-rails et le head-of-line blocking (§39).

**Q7.** Un switch 64×800G tout-optique consomme ~2 kW. Quelle part vient
des optiques ?
**R.** Environ la moitié : 64 × 15 W ≈ 960 W d'optiques + ~940 W de switch
(SN5600 : 940 W typique, 2,08 kW à 64 optiques, §90).

