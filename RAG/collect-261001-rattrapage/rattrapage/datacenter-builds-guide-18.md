---
id: collect-261001-rattrapage/rattrapage/datacenter-builds-guide-18
title: "Datacenter Builds — Le guide des BOMs"
domain: rattrapage
role: reference
task: reference
actors: []
dates: ["2026-09-27"]
keywords: ["arr", "attention", "ethernet", "gpu", "incident", "memory"]
source: docs/RAG/collect-261001-rattrapage/datacenter_builds_guide.md
source_anchor: ""
source_lines: [3132, 3328]
sha256: 4218d02be02290924d6db460abeae4ad3fc57bb08dfd7dc9d612a966962b8b8c
---

# Datacenter Builds — Le guide des BOMs

| Outil | Rôle | Coût |
|---|---|---|
| Zabbix (voir guide zabbix) | métriques serveurs, PDU, sondes | 0 € |
| NetBox (voir guide netbox) | inventaire, câblage, IP | 0 € |
| NUT | arrêt auto onduleur (voir guide proxmox) | 0 € |
| Caméra thermique (FLIR One) | audit annuel | ≈ 300 € |
| DCIM commercial | capacity planning | 10–30 k€ (à vérifier) |

**Alertes critiques** (SMS/astreinte) : température > 27 °C, PDU > 80 %,
onduleur sur batterie, fuite d'eau, porte ouverte hors heures. Tout le reste
en mail.

---

## 189. Fibre inter-racks et inter-salles

- **OM4** jusqu'à 100 m en 100G (SR4), **OS2** au-delà (LR4).
- Prévoyez des **tiroirs de brassage** avec 30 % de ports libres.
- **2 chemins physiques** pour les liens critiques (A et B ne passent pas
  dans le même caniveau — un coup de pelle coupe les deux sinon).
- Testez chaque lien au **photomètre/OTDR** à la réception (pas « ça ping,
  c'est bon »).

---

## 190. Sécurité incendie : gaz, pas d'eau

- **Gaz inerte (IG-541, Novec)** : extinction sans résidu, sans danger pour
  l'électronique. Prix : 30 000–50 000 € pour 100 m² (à vérifier).
- Détection **aspiration (VESDA)** : détecte à l'état de pré-combustion.
- **Jamais de sprinkler à eau** sur les racks (sauf pré-action à double
  détection si le code l'impose).
- Exercice d'évacuation annuel, coupure électrique d'urgence identifiée et
  **testée** (bouton coup de poing).

---

## 191. Contrôle d'accès et traçabilité

- Badge nominatif, zones (salle serveurs ≠ local batteries ≠ TGBT).
- **Registre** : toute intervention notée (qui, quoi, quand).
- Caméras : allées + entrée, rétention 30 j mini.
- Fournisseurs : accompagnés, jamais seuls. **Sans exception.**

---

## 192. Exploitation : la check-list mensuelle

- [ ] Températures (min/max du mois) — dérive ?
- [ ] PDU : pic de charge < 80 % ?
- [ ] Onduleur : test batterie OK, autonomie mesurée
- [ ] Groupe : test en charge 30 min
- [ ] Disques : SMART, 0 erreur non corrigée
- [ ] Sauvegardes : 1 restore test
- [ ] Firmware : CVE critiques ?
- [ ] Consommables : filtres, étiquettes, DAC de rechange
- [ ] Registre d'accès : relu, anomalies ?

---

## 193. Négociation fournisseurs : les 7 leviers

1. **3 devis minimum**, même intégrateur habituel (le concurrent fait baisser).
2. **Fin de trimestre** : les commerciaux ont des objectifs — −5 à 10 %.
3. **Paiement comptant** vs 60 j : −2 à 3 % si vous pouvez.
4. **Garantie 5 ans** négociée au lieu de 3 : coûte moins cher à l'achat
   qu'en extension plus tard.
5. **Pièces détachées** : 1 lot (disques, PSU, ventilos) offert dès 10 nœuds.
6. **Burn-in + rapport** inclus par écrit.
7. **Reprise DEEE** de l'ancien parc : 0 € de coût caché.

---

## 194. Occasion et reconditionné : quand c'est pertinent

| Composant | Occasion intéressante ? | Prix constaté |
|---|---|---|
| H100 80 Go SXM | oui (garantie vendeur) | 18 000–22 000 $ (vérifié le 27/09/2026) |
| RTX 6000 Ada 48 Go | oui | ≈ 5 000–6 000 $ (à vérifier) |
| A100 80 Go | oui si budget serré | 12 000–18 000 $ (vérifié le 27/09/2026) |
| Serveurs complets N−1 | oui (garantie 1 an) | −40 à −60 % vs neuf |
| SSD d'occasion | **non** (endurance inconnue) | — |
| Batteries onduleur | **non** | — |

Règle : occasion = **GPU et serveurs** (testés, garantis), jamais = SSD,
batteries, câbles fibre (fragiles).

---

## 195. Fin de vie : DEEE et effacement

- **Effacement certifié** (NIST 800-88) avant toute sortie : disques, SSD
  (les SSD s'effacent mal — destruction physique pour le très sensible).
- DEEE : filière agréée, bordereau de suivi (obligation légale FR).
- Données : le disque qui part au recyclage **sans** effacement = fuite de
  données. Procédure écrite, registre d'effacement.

---

## 196. Retours terrain — 8 histoires vraies (anonymisées, ordres de grandeur)

1. **Le rack qui ne fermait pas.** 12 serveurs 2U commandés avec un rack
   1 000 mm « en promo ». Les PDU 0U + les bras de câbles dépassaient de
   12 cm. Solution : racheter 2 racks 1 200 mm (5 000 €) et 2 jours de
   déménagement. **Leçon : §127.**
2. **Le neutre qui chauffait.** 3 racks sur PDU tri 32 A, phases à 28/12/9 A.
   Le neutre à 65 °C, odeur de chaud un vendredi soir. Rééquilibrage le
   week-end, 0 incident. **Leçon : §131.**
3. **Le boot storm du lundi.** 250 VDI sur Ceph HDD (pas NVMe). À 8h55,
   latence 8 s, helpdesk saturé. Migration des images sur NVMe local :
   problème disparu. Coût : 12 000 € de NVMe. **Leçon : §66.**
4. **Le firmware qui divergeait.** 8 nœuds Ceph, 3 versions de BIOS. Un nœud
   rebootait aléatoirement (bug C-state corrigé dans la version récente).
   3 semaines de debug pour 20 min de MAJ. **Leçon : §170.**
5. **La clim « au total ».** Salle de 60 kW avec 2× 40 kW de clim… mais tout
   le froid soufflé d'un côté. Le rack du fond à 34 °C en entrée, throttling
   CPU. Ajout d'un confinement : −8 °C. **Leçon : §158.**
6. **L'onduleur en kVA.** 50 kW de serveurs sur 60 kVA fp 0,8 = 48 kW.
   Premier test de basculement : surcharge, bypass. Échange standard contre
   80 kW. **Leçon : §166.**
7. **Le backup jamais restauré.** 2 ans de backups Veeam « verts ». Le jour
   du ransomware : le repo n'était pas immuable, chiffré aussi. 3 semaines
   de reconstruction. **Leçon : §105, §107.**
8. **Le GPU non qualifié.** 4× RTX PRO 6000 dans un châssis 350 W/GPU.
   Throttling à 2 400 W dissipés au lieu de 2 400 W… en fait 600 W × 4 que le
   châssis ne pouvait pas extraire : −30 % de perfs. Changement de châssis.
   **Leçon : §99.**

---

## 197. Scripts — burn-in CPU/RAM

```bash
# CPU : 24 h, tous les cœurs
stress-ng --cpu 0 --cpu-method matrixprod --timeout 86400 --metrics-brief

# RAM : 1 passe complète (8 h pour 768 Go environ)
memtester 700G 1

# Surveillance pendant le test (autre terminal)
watch -n 5 'sensors | grep -E "Package|Tdie"; ipmitool sdr | grep -i fan'
```

---

## 198. Scripts — burn-in disques (fio)

```bash
# NVMe : 4k aléatoire, QD32, 2 h — comparez à la fiche (PM9A3: ~900k/200k)
fio --name=randread --ioengine=libaio --rw=randread --bs=4k \
    --iodepth=32 --numjobs=4 --runtime=7200 --time_based \
    --filename=/dev/nvme0n1 --direct=1 --group_reporting

fio --name=randwrite --ioengine=libaio --rw=randwrite --bs=4k \
    --iodepth=32 --numjobs=4 --runtime=7200 --time_based \
    --filename=/dev/nvme0n1 --direct=1 --group_reporting

# ATTENTION : randwrite use le SSD — ne le faites qu'au burn-in, jamais en prod.
# Vérifiez l'endurance restante après : nvme smart-log /dev/nvme0n1
```

---

## 199. Scripts — burn-in réseau (iperf3)

```bash
# Serveur (sur la cible)
iperf3 -s -D

# Client : 24 h, bidirectionnel, 4 flux
iperf3 -c <IP> -t 86400 -P 4 --bidir --logfile iperf.log

# Objectif 25G : ≥ 23 Gb/s utiles, 0 retransmit anormal
# Surveillez : ethtool -S eth0 | grep -i -E "drop|error|crc"
```

---

## 200. Scripts — inventaire rapide d'un nœud

```bash
#!/bin/bash
# inventaire.sh — à lancer sur chaque nœud, sortie vers l'as-built
echo "=== $(hostname) $(date -I) ==="
dmidecode -t processor | grep -E "Version|Core Count"
dmidecode -t memory | grep -E "Size|Speed|Part Number" | head -40
lsblk -d -o NAME,SIZE,MODEL | grep -v loop
lspci | grep -i -E "ethernet|vga|3d"
ipmitool fru | head -20
nvme list 2>/dev/null | awk '{print $1, $2, $13}'
ethtool eth0 2>/dev/null | grep -E "Speed|Duplex"
```

---

## 201. Gabarit — BOM vierge (à copier par projet)

