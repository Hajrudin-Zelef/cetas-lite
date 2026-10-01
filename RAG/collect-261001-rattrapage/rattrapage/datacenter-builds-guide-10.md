---
id: collect-261001-rattrapage/rattrapage/datacenter-builds-guide-10
title: "Datacenter Builds — Le guide des BOMs"
domain: rattrapage
role: reference
task: reference
actors: ["AMD", "Intel", "Nvidia"]
dates: ["2026-09-27"]
keywords: ["datacenter", "blackwell", "capex", "dram", "gpu", "intel", "nvidia", "nvlink", "open source"]
source: docs/RAG/collect-261001-rattrapage/datacenter_builds_guide.md
source_anchor: ""
source_lines: [1632, 1822]
sha256: e115ee21ba8d2102717ea69460a5890a4100dc018f198128413ddbedf95a6cf6
---

# Datacenter Builds — Le guide des BOMs

1. **Le 600 W du RTX PRO 6000** : beaucoup de châssis 4-GPU sont qualifiés
   350 W/GPU (génération L40S). Vérifiez la qualification thermique par écrit.
2. **PCIe 4.0 vs 5.0** : un GPU PCIe 5.0 dans un slot 4.0 perd peu en inférence,
   mais vérifiez quand même la négociation (GPU-Z/nvidia-smi).
3. **ECC VRAM** : activé par défaut sur les datacenter (L40S), pas sur les
   workstation. Pour de la prod, activez-le (perte ~6 % VRAM, fiabilité).
4. **Délais Blackwell** : 3–7 mois début 2026. Ne planifiez pas une mise en
   prod à J+30 avec du RTX PRO 6000.
5. **Refroidissement des 600 W** : à 30 °C ambiants, un 600 W air-cooled
   throttle. 24 °C max en salle, ou liquide.

---

## 100. Ce qu'il faut retenir — Inférence NVIDIA

- L40S = standard 48 Go (≈ 47 k€ le serveur 2×) ; RTX PRO 6000 = 70B sur
  1 carte (≈ 35 k€ le serveur 1×).
- 600 W/GPU : qualification thermique du châssis obligatoire.
- H100/H200 uniquement pour haute concurrence ou NVLink.
- Le PCIe suffit jusqu'à ~50 req/s sur 70B.

---

## 101. Workload 9 — Backup : principes

La règle **3-2-1** : 3 copies, 2 supports différents, 1 hors site.
Le serveur backup se dimensionne à :
1. **La volumétrie source** × taux de dédup × rétention (formule §104).
2. **La fenêtre de sauvegarde** → débit réseau/disques (formule §11).
3. **La volumétrie de restauration** → le restore doit tenir dans le RTO.

Logiciels : Veeam, Proxmox Backup Server (déduplication excellente),
Bareos, restic/Borg pour le fichier.

---

## 102. Serveur backup S — BOM (50 To utiles)

| Composant | Référence type | Qté | Prix |
|---|---|---|---|
| Châssis 2U 12 baies 3,5", 1P EPYC | 2U 12× 3,5" nu | 1 | 3 800 € |
| EPYC 9355P (32c) — la dédup aime les cœurs | — | 1 | 2 891 € |
| RAM 12× 32 Go = 384 Go (cache dédup) | DDR5 ECC | 12 | 9 000 € |
| 2× M.2 960 Go RAID 1 (OS) | — | 2 | 360 € |
| 2× NVMe 3,84 To (landing zone, cache) | PM9A3 | 2 | 6 070 € |
| 10× HDD 12 To CMR (RAID 60 ou ZFS raidz2) | Exos | 10 | ≈ 280 € × 10 = 2 800 € |
| NIC 2× 25 GbE | E810 | 1 | 450 € |
| PSU 2× 1 600 W | incluses | — | — |
| **TOTAL** | | | **≈ 25 370 €** |
| P_max / réaliste | | | **≈ 0,7 kW / 0,5 kW** |

Capacité : 10× 12 To en RAID 60 (2× raidz2 5 disques) ≈ **72 To utiles**.
La landing zone NVMe absorbe les pics d'écriture nocturnes.

---

## 103. Serveur backup M — BOM (200 To utiles)

| Composant | Référence type | Qté | Prix |
|---|---|---|---|
| Châssis 4U 24 baies 3,5", 1P EPYC | 4U24 nu | 1 | 4 500 € |
| EPYC 9455P (48c) | — | 1 | 3 406 € |
| RAM 12× 64 Go = 768 Go | DDR5 ECC | 12 | 16 560 € |
| 2× M.2 960 Go RAID 1 (OS) | — | 2 | 360 € |
| 4× NVMe 7,68 To (landing zone) | PM9A3 | 4 | 15 660 € |
| 20× HDD 20 To CMR (ZFS raidz2 ×2) | Exos X20 | 20 | 380 € × 20 = 7 600 € |
| NIC 2× 25 GbE | E810 | 1 | 450 € |
| PSU 2× 1 600 W | incluses | — | — |
| **TOTAL** | | | **≈ 48 535 €** |
| P_max / réaliste | | | **≈ 0,85 kW / 0,65 kW** |

Capacité : 20× 20 To en 2× raidz2 (8+2) ≈ **320 To bruts → ~290 To utiles**.
La RAM (768 Go) sert de cache ARC ZFS : le restore des petits fichiers est
transformé.

---

## 104. Dimensionnement à la volumétrie — la formule

```
Capacité = sources × (1 + incr_quot × jours + full_sup) × (1 + rétention_mois/12) / taux_dédup
```

Exemple : 100 To sources, incr 3 %/j, full hebdo, rétention 3 mois, dédup 2:1
(PBS sur VMs : 2–4× réaliste) :
- 100 × (1 + 0,03×6 + 1) × 1,25 / 2 = **136 To utiles**.
→ Serveur S (72 To) trop petit, **serveur M (290 To) avec marge**.

Taux de dédup réalistes 2026 :
- VMs bureautiques (PBS) : 3–5×.
- DB (données uniques) : 1,2–1,5×.
- Fichiers bureautiques : 1,5–2×.
- **Ne comptez jamais plus de 2× sans mesure sur VOS données.**

---

## 105. Fenêtre de sauvegarde → débit (exemple)

100 To en 8 h (voir §11) : **≈ 10 Gb/s** → NIC 25 GbE.
Règles :
- Full hebdo le week-end (fenêtre 48 h), incr la nuit (8 h).
- **Backup et Ceph scrub jamais en même temps** (§50).
- Immuabilité (object lock S3 / WORM) contre le ransomware : le backup doit
  survivre à un attaquant avec les droits admin. **Sans immuabilité, le 3-2-1
  est une illusion.**

---

## 106. Cas chiffré : 100 To sources, 3-2-1 complet

| Poste | Détail | Total |
|---|---|---|
| Serveur backup M local | §103 | 48 535 € |
| Copie 2 (NAS secondaire ou Ceph objet) | 150 To utiles | 25 000 € |
| Copie 3 hors site (S3 immuable, 150 To) | ≈ 20 €/To/mois (à vérifier) | 36 000 €/an |
| Logiciel (PBS = 0 € ; Veeam ≈ 15 k€/an) | open source | 0 € |
| **CAPEX année 1** | | **≈ 73 500 € + 36 k€/an** |
| P_max local | | **≈ 1 kW** |

---

## 107. Pièges terrain — Backup

1. **Backup non testé** : restaurez 1 VM par mois en exercice. Le jour du
   sinistre n'est pas le jour de la découverte.
2. **Dédup promise par le commercial** : mesurez sur vos données (PBS a un
   mode dry-run).
3. **Fenêtre dépassée** : le backup qui finit à 10h ralentit la prod. Alertez
   sur la durée, pas seulement sur l'échec.
4. **Chiffrement des clés perdues** : un repo chiffré sans escrow des clés =
   perte définitive. 2 copies des clés, 2 lieux.
5. **Ransomware** : immuabilité obligatoire (object lock). Testez la
   restauration depuis la copie immuable.

---

## 108. Workload 10 — Firewall OPNsense : dimensionner AU DÉBIT

Un firewall se dimensionne sur **4 métriques** : débit firewall (Gbps), PPS,
débit IPsec, débit avec IDS/IPS activé (Suricata divise par 5–10).
Le tableau ci-dessous donne le **CPU/NIC minimal** par palier de débit
Internet réel (pas le débit du port — le débit facturé par l'opérateur).

---

## 109. Tableau débit → CPU/NIC (bare-metal OPNsense)

| Débit Internet | CPU minimal | NIC | RAM | Build type |
|---|---|---|---|---|
| ≤ 1 Gb/s | 4c Atom/Celeron récent (N100+) | 2× 2,5 GbE Intel i226 | 8 Go | Mini-PC |
| 2–5 Gb/s | 4–6c Core i3/i5 ou Xeon D-1700 | 2× 10G SFP+ (X710) | 16 Go | 1U léger |
| 10 Gb/s | 8c Xeon D-1733NT / EPYC 9124 | 2–4× 10G SFP+ | 32 Go | 1U (DEC3862 / Netgate 8300) |
| 25 Gb/s | 16c EPYC 9124/9355P | 2× 25G SFP28 (E810) | 32–64 Go | 1U EPYC |
| 40–100 Gb/s | 32c+ EPYC 9355P/9555P | 2× 100G QSFP28 | 64 Go | 1U/2U EPYC |

Repères constructeur (vérifiés le 27/09/2026) :
- Deciso DEC3842 : 14,4 Gbps firewall, 2,3 Gbps IPsec (AES-256-GCM).
- Netgate 8300 (Xeon D-1733NT 8c) : 36,7 Gbps L3, 26,8 Gbps firewall,
  14,6 Gbps IPsec — **3 599 $** base.
- **Avec Suricata IPS** : divisez ces chiffres par 5 environ.

---

## 110. Build firewall 1G — BOM (PME, 1 Gb/s)

| Composant | Référence type | Qté | Prix |
|---|---|---|---|
| Mini-PC 4× 2,5 GbE Intel i226, N100+ | barebone | 1 | ≈ 350 € |
| RAM 16 Go DDR5 | — | 1 | 80 € |
| SSD M.2 256 Go | — | 1 | 40 € |
| **TOTAL** | ×2 pour HA | | **≈ 470 € × 2 = 940 €** |
| Conso | | | **≈ 15 W / unité** |

---

## 111. Build firewall 10G — BOM (ETI, 10 Gb/s)

| Composant | Référence type | Qté | Prix |
|---|---|---|---|
| 1U EPYC 8c, 4× SFP+ | Deciso DEC3862 ou 1U EPYC nu | 1 | ≈ 2 500 € (appliance) |
| RAM 32 Go ECC | — | 1 | ≈ 750 € (prix 2026 !) |
| SSD 512 Go | — | 1 | 80 € |
| **TOTAL** | ×2 pour HA | | **≈ 3 330 € × 2 = 6 660 €** |
| Conso | | | **≈ 80–120 W / unité** |

Note : 32 Go de RAM à 750 € — la flambée DRAM frappe aussi les firewalls.

---

## 112. Build firewall 25/40G — BOM (datacenter)

