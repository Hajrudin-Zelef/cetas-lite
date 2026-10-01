---
id: collect-261001-rattrapage/rattrapage/datacenter-reseau-guide-16
title: "RÉSEAU DATACENTER / IA — NICs, FIBRE, SPINE-LEAF, SWITCHES IA, DPU, CDN/EDGE"
domain: rattrapage
role: reference
task: reference
actors: ["xAI"]
dates: []
keywords: ["datacenter", "arr", "asic", "capex", "cpo", "gpu", "training"]
source: docs/RAG/collect-261001-rattrapage/datacenter_reseau_guide.md
source_anchor: ""
source_lines: [2243, 2365]
sha256: aebc39d78529d0a20582869a6e4ce55ad8dfd323f6c904da8770fc6b8ba61270
---

# RÉSEAU DATACENTER / IA — NICs, FIBRE, SPINE-LEAF, SWITCHES IA, DPU, CDN/EDGE

**Q8.** Quand un DPU est-il rentable ?
**R.** Quand il libère des cœurs CPU chers (licences au cœur), quand
l'isolation multi-tenant est requise, ou pour le NVMe-oF / la sécurité
inline à 400G. Pas pour un serveur généraliste < 100G (§69, §71).

**Q9.** Que change le CPO (Tomahawk 6-Davisson) pour l'énergie ?
**R.** 3,5 W/port à 800G contre ~15 W en pluggable (−77 %) : les optiques,
qui représentent ~50 % de la conso d'un switch 800G, sont divisées par 4
(§101).

**Q10.** Citez 3 vérifications avant de signer la réception d'un fabric IA.
**R.** (au choix) : NCCL all-reduce ≥ 80 % du théorique ; ECMP équilibré à
±10 % sous charge ; perte d'un spine < 5 % de perte et < 1 s de convergence ;
72 h de burn-in sans erreur FEC ; aucune optique > 75 °C (§107).

## 122. Les 45 pièges terrain — récapitulatif complet

| # | Piège | Symptôme | Parade (§) |
|---|---|---|---|
| P1 | NIC 400G en slot PCIe Gen4 | Débit plafonné à ~200G | Vérifier génération × lanes (§8) |
| P2 | Offloads désactivés (vieux driver) | CPU 100 % à 100G | `ethtool -k`, MAJ driver (§9) |
| P3 | Optiques tierces refusées (FW lock) | Lien down, log « unsupported » | Tester 1 avant 100 (§11) |
| P4 | Firmwares NIC hétérogènes | Micro-pertes aléatoires | Triplet FW/driver/OS figé (§13) |
| P5 | NIC OCP 3.0 pour châssis CEM | Ne rentre pas physiquement | Valider le format avant commande (§12) |
| P6 | OSFP flat-top sur switch | Surchauffe +15-20 °C | Finned-top sur switch (§16) |
| P7 | Jarretière OS2 (jaune) sur port SR | Lien muet | Couleurs : aqua=OM, jaune=OS2 (§18) |
| P8 | Polarité MPO non documentée | Heures de debug nocturne | Étiqueter méthode A/B/C (§26) |
| P9 | DAC + transceiver sur le même lien | Pas de link | Même média des 2 côtés (§22) |
| P10 | Breakout non supporté par le port | Ports morts après achat | Matrice de breakout du switch (§23) |
| P11 | FEC qui masque une fibre sale | Panne brutale 6 mois plus tard | Superviser le pré-FEC (§27) |
| P12 | Backend IA + front-end même fabric | Le training noie tout | 2 fabrics séparés (§32) |
| P13 | Zéro port libre | Croissance = nouveau pod | +20 % de ports (§35) |
| P14 | MTU VXLAN oublié (+50 o.) | Fragmentation silencieuse | MTU 9050+ sur l'underlay (§37) |
| P15 | ECMP sans hash seed | Un spine saturé | Seed aléatoire par switch (§36) |
| P16 | PFC global au lieu de par file | Pauses storms, cascade | PFC sur la seule file RoCE (§45) |
| P17 | Leaf shallow-buffer en spine | Drops sous incast | Deep-buffer au spine (§44) |
| P18 | RoCE sans ECN/DCQCN | 60 % du débit (cas xAI) | Checklist §46 |
| P19 | Comparer les switchs sur l'ASIC seul | Mauvais choix | Évaluer le NOS + la validation (§50) |
| P20 | Airflow switch inversé | Hotspot 20 kW/rangée | Cohérence allées chaudes/froides (§59) |
| P21 | DAC 800G de 3 m | Ne monte pas | Portées §59 (DAC ≤ 2 m) |
| P22 | Optiques 18 W sur switch 12 W | Throttling thermique | Spec « max power per port » (§60) |
| P23 | Disjoncteurs au nominal | Déclenchement au boot | Calibre ≥ 125 % (§61) |
| P24 | DPU « pour faire moderne » | Surcoût sans gain | Matrice de décision §69 |
| P25 | 50 k lignes de DOCA sans plan de sortie | Lock-in logiciel | Évaluer le coût de sortie (§68) |
| P26 | DPU sans airflow (75 W) | Hotspot slot PCIe | Vérifier le cooling du slot (§72) |
| P27 | Cache CDN sur SATA seul | IOPS saturées | NVMe obligatoire (§74) |
| P28 | Serveur 800 mm en baie edge 600 mm | Ne rentre pas | Vérifier profondeur + alim (§76) |
| P29 | Anycast sur TCP longues sans persistance | Coupures au flap BGP | Persistance ou DNS (§78) |
| P30 | TTL infini sans purge | Contenu périmé servi | Purge + TTL adaptés (§74) |
| P31 | Fibre coincée sous un rail | BER élevée, invisible | Mou aux extrémités (§83) |
| P32 | Goulotte à 120 % | Dizaines de liens dégradés | Remplissage ≤ 50 %, audit annuel (§86) |
| P33 | Câble non étiqueté | 1 h de debug à 150 €/h | TIA-606 systématique (§85) |
| P34 | Plan d'adressage trop juste | Renumérotation nocturne | Prévoir la croissance (§87) |
| P35 | Salle réseau à 3 kW/baie | Surchauffe en 800G | 8-12 kW/baie (§94) |
| P36 | Disjoncteurs au nominal | Déclenchement au boot | (voir P23) calibre ≥ 125 % (§95) |
| P37 | Réseau pas sur UPS prioritaire | Arrêt à l'aveugle | UPS réseau séparé, en dernier (§117) |
| P38 | PUE sans les optiques | −50 % du poste oublié | Compter switch + optiques + NIC (§93) |
| P39 | Tests à vide seulement | Échec en all-to-all | Tester sous charge (§107) |
| P40 | NCCL sur le front-end 100G | 4× moins vite | `NCCL_SOCKET_IFNAME` (§108) |
| P41 | Stockage dimensionné à la moyenne | Checkpoint = 10 s de pause | Dimensionner pour le checkpoint (§109) |
| P42 | Automatiser un process cassé | Casser plus vite | Stabiliser puis automatiser (§111) |
| P43 | Fabric IA étendu inter-sites | All-reduce effondré | 1 fabric par site (§114) |
| P44 | Switch AC sur site −48 V DC | Ne s'alimente pas | Vérifier l'option DC (§115) |
| P45 | GPU et réseau sur le même UPS | La panne GPU emporte la supervision | UPS séparés (§117) |

## 123. BOMs types — trois projets prêts à chiffrer

### BOM 1 — Pod IA 256 GPU, backend 800G, 1:1 (prix ⚠️ à vérifier)

| Poste | Qté | PU (⚠️) | Total (⚠️) |
|---|---|---|---|
| Switch 32×800G (8 leafs + 16 spines) | 24 | 100-200 k€ | 2,4-4,8 M€ |
| NIC 800G (ConnectX-8) | 2048 | 3-6 k€ | 6,1-12,3 M€ |
| DAC/AOC 800G intra-rack | 2048 | 0,4-2 k€ | 0,8-4,1 M€ |
| Optiques 800G-DR8 inter-rangées | 768 | 4-10 k€ | 3,1-7,7 M€ |
| Fibre OS2/OM4 + brassage MDA | 1 lot | — | 150-300 k€ |
| Kit nettoyage + supervision | 1 lot | — | 30-60 k€ |
| Câblage, étiquetage, installation | 1 lot | — | 200-400 k€ |
| **Total** | | | **~12,8-29,7 M€** |

### BOM 2 — Datacenter cloud 100G, 512 serveurs, 3:1 (prix ⚠️)

| Poste | Qté | PU (⚠️) | Total (⚠️) |
|---|---|---|---|
| Leaf 48×25G+8×100G | 12 | 15-30 k€ | 180-360 k€ |
| Spine 32×100G | 2 | 25-50 k€ | 50-100 k€ |
| NIC 2×25G/100G (E810) | 1024 | 300-800 € | 307-819 k€ |
| DAC 25G/100G | 1024 | 50-150 € | 51-154 k€ |
| Optiques 100G-SR4/LR4 | 48 | 300-1 500 € | 14-72 k€ |
| **Total** | | | **~0,6-1,5 M€** |

### BOM 3 — PoP edge CDN 200 Gb/s (prix ⚠️, voir §79)

| Poste | Qté | PU (⚠️) | Total (⚠️) |
|---|---|---|---|
| Serveurs edge 2×100G, 12× NVMe | 2 | 25-40 k€ | 50-80 k€ |
| Switch 8×100G | 1 | 8-15 k€ | 8-15 k€ |
| Routeur cellulaire OOB | 1 | 300-600 € | 0,3-0,6 k€ |
| Spares (optiques, câbles, 1 switch) | 1 lot | — | 10-20 k€ |
| **CAPEX total** | | | **~68-116 k€** |

## 124. Checklist d'achat — à joindre à chaque appel d'offres

**NIC**
- [ ] Débit vs PCIe : génération × lanes **électriques** vérifiées (lspci)
- [ ] Format : CEM / OCP 3.0 / mezzanine — compatible châssis
- [ ] Firmware minimal pour le débit annoncé (ex. CX-8 : 40.46.x)
- [ ] Driver : inbox ou propriétaire ? Qui le maintient, quel cycle ?
- [ ] Offloads requis : SR-IOV (__ VF), VXLAN, RoCEv2, DPDK, TLS inline
- [ ] Optiques : lock constructeur ? Tester 1 tierce avant 100
- [ ] Conso **avec optiques**, pas carte seule
- [ ] UEC-ready / RoCEv2 validé sur le switch cible

**Switch**
- [ ] Capacité : non-bloquant ? Ratio down/up calculé (pas promis)
- [ ] Buffers : shallow (leaf) / deep (spine) selon position
- [ ] ECN, PFC par file, DCQCN / CC : supportés et validés
- [ ] Breakout : matrice des modes par port (doc constructeur)
- [ ] Max power per port optique (12 W ? 18 W ?)
- [ ] Airflow : sens cohérent avec les allées
- [ ] PSU : 1+1, 2 arrivées A/B, calibre disjoncteur ≥ 125 %
- [ ] NOS : EOS / NX-OS / SONiC durci — qui l'opère chez nous ?

