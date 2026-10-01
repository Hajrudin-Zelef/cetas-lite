---
id: collect-261001-rattrapage/rattrapage/datacenter-builds-guide-22
title: "Datacenter Builds — Le guide des BOMs"
domain: rattrapage
role: reference
task: reference
actors: []
dates: ["2026-09-27"]
keywords: ["arr", "dram", "gpu"]
source: docs/RAG/collect-261001-rattrapage/datacenter_builds_guide.md
source_anchor: ""
source_lines: [3954, 4003]
sha256: 2c41ddcff7b53a9a338ad35fb4fb0d2877641cea87ad9cdabdece422c28676fd
---

# Datacenter Builds — Le guide des BOMs

1. 1 500 $ — le 64 Go DDR5 en contrat (09/2026).
2. 300 m³/h — le débit d'air par kW (ΔT 12 °C).
3. 85 % — le remplissage max d'un pool Ceph.
4. 80 % — la charge max d'une PDU en nominal.
5. 20 % — l'écart max entre phases.
6. 1 752 € — le coût annuel d'1 kW à 0,20 €/kWh.
7. 28 h — le rebuild d'un HDD 20 To.
8. 72 h — le burn-in avant prod.
9. 15 min — l'autonomie onduleur standard.
10. 3 — le nombre de devis minimum.

---

## 241. Pour aller plus loin (guides liés de Zelef)

- `onduleurs_ups_guide.md` — le détail onduleurs/batteries/groupes.
- `proxmox_guide.md` — grappe Proxmox + NUT (arrêt auto).
- `zabbix_guide.md` — supervision (sondes, PDU, températures).
- `netbox_guide.md` — inventaire et plan de câblage.
- `debian_ubuntu_guide.md` — OS des serveurs.
- `ansible_guide.md` — automatisation du provisionning.

---

## 242. Crédits et méthode de recherche

Prix relevés sur le web public le 27/09/2026 : configurateurs européens
(rect.coreto.de, ahead-it.eu, scan.co.uk), snapshots de marché GPU
(sourcebyspec.com, gpusmith.com), boutiques (tech-america.com, dell.com,
superwarehouse.com), constructeurs (shop.opnsense.com, info.netgate.com),
presse spécialisée (sedaily.com/TrendForce, datacentremagazine.com,
upsite.com/AFCOM 2026). Les prix « à vérifier » sont des ordres de grandeur
2026 sans source directe ce jour-là. Aucun prix n'a été inventé.

---

## 243. Note de maintenance du guide

- Relire la section 3 (marché) avant chaque budget annuel : DRAM, GPU et
  délais bougent par trimestre en 2026.
- Mettre à jour le tableau §125 à chaque nouvel achat : c'est la référence
  électrique de la salle.
- Archiver chaque devis reçu avec sa date : c'est l'historique de prix qui
  rend les négociations suivantes crédibles.
- Ce guide ne remplace ni l'électricien (TGBT, sélectivité), ni le BET
  (structure, froid), ni le devis écrit du fournisseur (délais GPU/DRAM).
- En cas de doute sur un dimensionnement : surdimensionnez la marge, pas le
  matériel — +30 % de place libre coûte moins cher qu'un rack à refaire.

*Document de travail — version 1.0 — 27/09/2026.*
