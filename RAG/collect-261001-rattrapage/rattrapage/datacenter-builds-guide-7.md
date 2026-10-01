---
id: collect-261001-rattrapage/rattrapage/datacenter-builds-guide-7
title: "Datacenter Builds — Le guide des BOMs"
domain: rattrapage
role: reference
task: reference
actors: ["AMD", "Nvidia"]
dates: []
keywords: ["amd", "capex", "gpu", "nvidia", "open source"]
source: docs/RAG/collect-261001-rattrapage/datacenter_builds_guide.md
source_anchor: ""
source_lines: [1075, 1257]
sha256: b3158c47da98e6152486cdc7a3b79627c045e9bdb4cbdc0894bdc4929c8c50d3
---

# Datacenter Builds — Le guide des BOMs

| Profil | vCPU | RAM | IOPS pic | GPU | Exemples |
|---|---|---|---|---|---|
| Tâche (call center) | 1–2 | 4 Go | 20 | non | saisie, téléphonie |
| Bureautique | 2 | 6–8 Go | 50 | non | Office, navigateur |
| Développeur | 4 | 16 Go | 80 | optionnel | IDE, builds |
| CAO/DAO | 4–8 | 16–32 Go | 100 | vWS 4–8 Go | AutoCAD, Revit |

Règle CPU : **surallocation 4:1** max en bureautique (4 vCPU pour 1 pCPU),
2:1 en CAO. Au-delà, la latence se ressent au clavier.

---

## 63. Serveur VDI S — BOM (200 users bureautiques)

Cible : 200 × 2 vCPU = 400 vCPU (ratio 4:1 → 100 pCPU), 200 × 8 Go = 1,6 To RAM.

| Composant | Référence type | Qté | Prix |
|---|---|---|---|
| Châssis 2U 2P EPYC | 2U 2P nu | 1 | 4 200 € |
| EPYC 9655P (96c) — 192 cœurs au total | 2× | 2 | 5 208 € × 2 = 10 416 € |
| RAM 24× 64 Go = 1,5 To | DDR5 ECC | 24 | 33 120 € |
| 2× M.2 960 Go RAID 1 (hyperviseur) | — | 2 | 360 € |
| 4× NVMe 7,68 To (images, linked clones) | PM9A3 | 4 | 3 915 € × 4 = 15 660 € |
| NIC 2× 25 GbE | E810 | 1 | 450 € |
| PSU 2× 2 000 W | incluses | — | — |
| **TOTAL / serveur (100 users)** | 2 serveurs requis | | **≈ 64 205 € × 2 = 128 410 €** |
| P_max / serveur | | | **≈ 1,35 kW** |

Note : 1,5 To RAM pour 100 users = 15 Go/user brut, soit 8 Go/user + marge N+1.
La RAM est **le poste n° 1** de la VDI en 2026.

---

## 64. Serveur VDI M — BOM (200 users CAO avec vGPU)

Cible : 200 × 4 vCPU, 200 × 16 Go = 3,2 To RAM, vGPU 4 Go/user.

| Composant | Référence type | Qté | Prix |
|---|---|---|---|
| Châssis 2U 4× GPU double-slot | ASUS ESC4000A-E12 nu | 1 | ≈ 5 000 € |
| EPYC 9555P (64c) | 2× | 2 | 4 830 € × 2 = 9 660 € |
| RAM 24× 128 Go = 3 To | DDR5 ECC | 24 | ≈ 2 750 € × 24 = 66 000 € |
| 2× M.2 960 Go RAID 1 | — | 2 | 360 € |
| 4× NVMe 7,68 To | PM9A3 | 4 | 15 660 € |
| 4× NVIDIA L40S 48 Go (vWS) | — | 4 | ≈ 8 000 € × 4 = 32 000 € |
| NIC 2× 25 GbE | E810 | 1 | 450 € |
| PSU 2× 2 600 W | incluses châssis | — | — |
| **TOTAL / serveur (100 users CAO)** | 2 serveurs | | **≈ 129 130 € × 2 = 258 260 €** |
| P_max / serveur | 720+240+120+1400+50+150 | | **≈ 2,68 kW** |

4× L40S = 192 Go VRAM → 48 users à 4 Go par GPU, 192 users par serveur à
vWS 4 Go. **2,68 kW par serveur : c'est du rack haute densité** (air limite,
RDHx recommandé — section 147).

---

## 65. vGPU NVIDIA : licences et découpage

| Licence | Usage | Prix indicatif |
|---|---|---|
| vPC | Bureau virtualisé | ≈ 50–100 €/user/an (à vérifier) |
| vWS | CAO/3D | ≈ 200–300 €/user/an (à vérifier) |
| vApps | RDSH applicatif | ≈ 30–50 €/user/an (à vérifier) |

- Découpage : L40S 48 Go → 12× vGPU 4 Go (vWS) ou 24× vGPU 2 Go (vPC).
- **Alternative sans licence** : GPU passthrough 1:1 (1 GPU = 1 VM puissante)
  ou AMD MxGPU/SR-IOV (pas de licence, écosystème plus restreint).
- Comptez les licences sur 5 ans dans le TCO : 200 users × 250 € × 5 ans =
  **250 000 €** — plus cher que les GPU eux-mêmes.

---

## 66. Stockage VDI : le boot storm décide

- 200 users × 100 IOPS pic = **20 000 IOPS** pendant 10 minutes à 8h55.
- 4× NVMe PM9A3 = 800 k IOPS écriture : largement couvert.
- **Linked clones / instant clones** : 1 image mère + deltas → divise le
  stockage par 5–10. Sans ça, 200 × 60 Go = 12 To par serveur.
- Tiering : images sur NVMe local (pas sur Ceph — latence), profils utilisateurs
  sur CephFS/NAS.

---

## 67. Cas chiffré : 300 postes bureautiques VDI complets

| Poste | Détail | Total |
|---|---|---|
| 3× serveurs VDI S (N+1 : 2+1) | 64 205 € × 3 | 192 615 € |
| Licences VMware Horizon ou Citrix (300 users) | ≈ 200 €/user/an | 60 000 €/an |
| Stockage profils (NAS 20 To) | — | 15 000 € |
| Switches 25G (mutualisés) | — | — |
| Clients légers (300× 400 €) | thin clients | 120 000 € |
| **CAPEX année 1** | | **≈ 327 600 € + 60 k€/an** |
| P_max serveurs | 3 × 1,35 kW | **≈ 4 kW** |
| **Coût/poste/an (5 ans)** | (327 600 + 300 000) / 300 / 5 | **≈ 418 €** |

À comparer à un PC fixe à 800 € renouvelé tous les 5 ans (160 €/an) : la VDI
ne se justifie **que** par la sécurité, le nomadisme ou la CAO centralisée.
Ne vendez jamais la VDI comme « moins cher » — vendez-la comme « plus sûr ».

---

## 68. Réseau VDI : le protocole d'affichage

- **Blast/PCoIP/HDX** : 150–500 kb/s par session bureautique, 2–5 Mb/s en CAO.
- 300 sessions × 300 kb/s = 90 Mb/s : un lien 1G suffit… mais la **latence**
  doit être < 50 ms (RTT). Au-delà, les utilisateurs détestent.
- Séparez le trafic d'affichage du trafic de management (VLAN dédié).

---

## 69. Pièges terrain — VDI

1. **Sous-dimensionner la RAM** : la VDI est un workload RAM. Le CPU peut
   attendre, la RAM non.
2. **Oublier le boot storm** : tout marche en test à 14h, tout s'écroule à
   8h55 le lundi. Testez à 100 % de connexions simultanées.
3. **vGPU sans licences comptées** : le coût des licences dépasse le matériel
   sur 5 ans.
4. **Profils itinérants de 20 Go** : le login prend 10 minutes. FSLogix ou
   équivalent, quotas stricts.
5. **VDI sur Wi-Fi d'hôtel** : 200 ms de latence = inutilisable. Prévoyez un
   mode dégradé ou des PC locaux pour les nomades extrêmes.

---

## 70. Ce qu'il faut retenir — VDI

- 8 Go RAM + 2 vCPU par user bureautique, 16 Go + 4 vCPU + vGPU en CAO.
- 100 IOPS/user au boot storm : NVMe local obligatoire.
- Coût 2026 : ≈ 420 €/poste/an tout compris — la sécurité, pas le prix.
- 2,7 kW par serveur vGPU : prévoyez le refroidissement (RDHx).

---

## 71. Workload 6 — SOC : principes

Un SOC se dimensionne **au volume de logs par jour** (Go/j ou EPS —
événements/seconde). La chaîne : collecte → parsing/normalisation → indexation
→ détection → rétention.
- **1 To/j ≈ 12 000 EPS** en moyenne (ordre de grandeur, dépend des sources).
- Rétention typique : 90 j chauds (recherche rapide) + 1 an tiède/froid.
- Les SIEM open source (Wazuh, ELK) se dimensionnent comme les commerciaux,
  la licence en moins.

---

## 72. Tableau volumétrie → sizing

| Volume/jour | EPS moyen | Nœuds indexation | Stockage 90 j chauds | Profil |
|---|---|---|---|---|
| 100 Go | 1 200 | 1 | 9 To | PME |
| 500 Go | 6 000 | 2–3 | 45 To | ETI |
| 1 To | 12 000 | 3–4 | 90 To | Grand compte |
| 2 To | 24 000 | 6–8 | 180 To | SOC mutualisé |
| 5 To | 60 000 | 12–16 | 450 To | Opérateur |

Stockage chaud = volume × 90 j × facteur_index (≈ 1,1–1,5 selon parsing).
Le tiède/froid (objet S3/Ceph) coûte 10× moins cher : **ne gardez en chaud que
ce que vous recherchez vraiment**.

---

## 73. Serveur SIEM S — BOM (500 Go/j, Wazuh/ELK)

| Composant | Référence type | Qté | Prix |
|---|---|---|---|
| Châssis 2U 1P EPYC, 12 baies | 2U 1P nu | 1 | 3 800 € |
| EPYC 9455P (48c) — l'indexation aime les cœurs | — | 1 | 3 406 € |
| RAM 12× 64 Go = 768 Go | DDR5 ECC | 12 | 16 560 € |
| 2× M.2 960 Go RAID 1 (OS) | — | 2 | 360 € |
| 8× NVMe 7,68 To (hot tier 61 To) | PM9A3 | 8 | 3 915 € × 8 = 31 320 € |
| NIC 2× 25 GbE | E810 | 1 | 450 € |
| PSU 2× 1 600 W | incluses | — | — |
| **TOTAL / nœud** | | | **≈ 55 895 €** |
| P_max / réaliste | | | **≈ 0,95 kW / 0,7 kW** |

---

## 74. Serveur SIEM M — BOM (2 To/j, cluster 3 nœuds)

Par nœud (×3, N+1) :

