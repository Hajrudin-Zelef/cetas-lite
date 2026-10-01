---
id: collect-261001-rattrapage/rattrapage/datacenter-reseau-guide-10
title: "RÉSEAU DATACENTER / IA — NICs, FIBRE, SPINE-LEAF, SWITCHES IA, DPU, CDN/EDGE"
domain: rattrapage
role: reference
task: reference
actors: ["AMD", "Broadcom", "Google", "Intel", "Nvidia", "Oracle"]
dates: ["2026-09-27"]
keywords: ["agents", "amd", "capex", "gpu", "intel", "nvidia"]
source: docs/RAG/collect-261001-rattrapage/datacenter_reseau_guide.md
source_anchor: ""
source_lines: [1326, 1484]
sha256: e14ac98d9c1c891b910242deed47b9cf500b1d4ede9050ea74c4ed4f607f8d30
---

# RÉSEAU DATACENTER / IA — NICs, FIBRE, SPINE-LEAF, SWITCHES IA, DPU, CDN/EDGE

## 70. Alternatives : Intel IPU, AMD Pensando, Marvell

| DPU | Base | Note |
|---|---|---|
| Intel IPU (ex-Mount Evans) | Intel + Google | ⚠️ Stratégie Intel réseau revue en 2024-2025 — vérifier la roadmap au 27/09/2026 |
| AMD Pensando | P4 programmable | Solide, écosystème plus petit |
| Marvell OCTEON | ARM + accélérateurs | Télécoms/edge, moins DC |
| Broadcom Stingray | — | ❌ Largement éclipsé par Thor — non recommandé en neuf |

**En 2026, le marché DPU = NVIDIA BlueField par défaut**, challengers de
niche. Ne pas acheter de DPU « exotique » sans valider le support OS
(driver inbox ?) et la roadmap 3 ans.

## 71. CAS CHIFFRÉ — TCO DPU sur 500 serveurs cloud privé

Hypothèses : 500 serveurs, OVS + firewall distribué, 100G.

| Poste | Sans DPU | Avec DPU (BF-3) |
|---|---|---|
| Cœurs CPU réservés réseau | 8/serveur × 500 = 4000 cœurs | 0 |
| Coût cœurs (⚠️ ~250 €/cœur serveur) | ~1 M€ | 0 |
| DPU (⚠️ ~3 000 €) | 0 | ~1,5 M€ |
| Conso réseau CPU (⚠️ 8 cœurs ~80 W) | 500 × 80 W = 40 kW | — |
| Conso DPU (⚠️ ~75 W) | — | 500 × 75 W = 37,5 kW |
| Licences au cœur (si VMware/Oracle) | Plein pot | **Économie majeure** |

**Verdict** : sans licences au cœur, le DPU se paie à peu près ; **avec**
licences au cœur ou contrainte d'isolation, il est très rentable. Toujours
chiffrer les **trois** : hardware, énergie, licences.

## 72. Pièges DPU — l'essentiel (détail §107)

- **P24** — DPU acheté pour « faire moderne » sans workload identifié.
- **P25** — DOCA développé sans plan de sortie : lock-in logiciel.
- **P26** — DPU sous-alimenté en airflow : 75 W dans un slot PCIe = hotspot.

# PARTIE G — CDN & SERVEURS EDGE

## 73. Principes : rapprocher le contenu, pas l'utilisateur

Un CDN = un réseau de **PoP** (points de présence) qui cachent le contenu
au plus près des utilisateurs. Trois effets :

1. **Latence** : 200 ms → 20 ms (le contenu vient du PoP local, pas de
   l'origine à l'autre bout du monde).
2. **Débit** : l'origine ne sert chaque objet qu'une fois (cache-fill), les
   PoP absorbent les pics (match de foot, lancement produit).
3. **Résilience** : si un PoP tombe, l'anycast reroute vers le suivant.

```
  Utilisateur ──→ PoP local (cache HIT 95 %) ──→ [contenu servi en 20 ms]
                         │ cache MISS (5 %)
                         ▼
                   PoP régional ──→ Origine (data center)
```

**Lien avec l'IA** : l'edge sert aussi l'**inférence** (LLM petits, vision)
— mêmes contraintes (latence, débit) + GPU edge (voir §77).

## 74. Dimensionnement du cache — règles de pouce

| Paramètre | Règle (⚠️ ordre de grandeur) |
|---|---|
| Taux de hit visé | 90-98 % (vidéo), 80-90 % (web) |
| Taille cache / débit | ~1 To de SSD par Gb/s de pic pour de la vidéo |
| Répartition | 80 % chaud sur NVMe, 20 % tiède sur SATA |
| TTL | Vidéo : jours ; API : secondes-minutes ; jamais « infini » sans purge |

**Calcul** : PoP 100 Gb/s vidéo → ~100 To utiles → ~25-30 SSD NVMe 4 To
(avec parité/erasure). **Le cache se dimensionne en To ET en IOPS** : 100 Gb/s
de petits objets = millions d'IOPS — le NVMe est obligatoire, le SATA seul
s'effondre (P27, §107).

## 75. Dimensionnement du débit — la méthode

1. **Pic mesuré ou estimé** : trafic moyen × facteur de pic (2-5× selon
   l'usage ; un match = 5×).
2. **Marge** : pic × 1,5 (croissance + panne d'un lien).
3. **Exemple** : 20 000 spectateurs × 8 Mb/s (1080p) = 160 Gb/s de pic →
   **2× 100G** (200 Gb/s) par PoP, soit 2 serveurs edge 100G en actif/actif.

**95e percentile** : la facturation transit se fait au 95e pct — lisser les
pics via le cache réduit directement la facture. Chaque point de hit-rate
gagné = du transit en moins.

## 76. Hardware d'un serveur edge type

| Composant | Spec type (⚠️ config usuelle) |
|---|---|
| CPU | 1-2× Xeon/EPYC 16-32 cœurs (le cache est I/O-bound, pas CPU-bound) |
| RAM | 256-512 Go (cache chaud en RAM pour les top objets) |
| Stockage | 8-24× NVMe 4-8 To (ou 2 To si petits objets) |
| NIC | 2× 100G (ou 1× 200G) — DPDK/VPP pour le TLS |
| Châssis | 1-2U, rails courts (baies edge peu profondes !) |

**Contrainte edge** : les baies télécoms font souvent **600 mm de profondeur**
(contre 1000-1200 mm en DC) — vérifier la profondeur serveur **et** le
câblage avant (P28, §107). Refroidissement : clim de confort, pas de salle
froide — prévoir des serveurs tolérants à 35 °C.

## 77. Cas d'usage : vidéo, gaming, IoT, inférence IA

| Usage | Contrainte n°1 | Architecture |
|---|---|---|
| Vidéo live/VOD | Débit massif, hit-rate | Gros caches NVMe, 100G+ |
| Gaming / paris | Latence < 20 ms | PoP denses, anycast |
| IoT / télémétrie | Millions de connexions | L7 léger, pas de gros cache |
| Inférence IA edge | Latence + GPU | Serveurs avec GPU L4/L40S, modèles quantifiés |
| Sécurité (WAF/DDoS) | Absorption | Scrubbing au PoP, anycast |

**Inférence edge (2026)** : le pattern « petit modèle au PoP, gros modèle
au DC » se généralise (agents, RAG local). Dimensionner le PoP avec **2
postes** : le CDN classique + 1-2 GPU d'inférence (voir guide IA).

## 78. Anycast, DNS et routage

- **Anycast** : la même IP annoncée depuis tous les PoP — BGP envoie
  l'utilisateur au PoP le plus proche (en AS-path). Panne d'un PoP =
  reroutage automatique en ~minutes (BGP), pas de DNS à attendre.
- **DNS** : le sélecteur fin (GeoDNS / EDNS Client Subnet) pour les cas où
  l'anycast ne suffit pas (contraintes légales, PoP saturé).
- **Ne pas anycaster** : les sessions avec état (TCP long) sans mécanisme
  de persistance — un flap BGP coupe les connexions (P29, §107).

## 79. CAS CHIFFRÉ — PoP edge 100 Gb/s

| Poste | Qté | Prix (⚠️) | Total (⚠️) |
|---|---|---|---|
| Serveurs edge 2×100G, 12× NVMe 4 To | 2 | 25-40 k€ | 50-80 k€ |
| Switch 8×100G (agrégation) | 1 | 8-15 k€ | 8-15 k€ |
| Routeurs / peering (mutualisé) | — | — | (hébergeur) |
| Transit 100G 95e pct (⚠️ ~1-3 €/Mb/s/mois en Europe) | 60 Gb/s utiles | 60-180 k€/an | **60-180 k€/an** |
| Hébergement baie edge | 1 | 12-24 k€/an | 12-24 k€/an |
| **CAPEX** | | | **~60-95 k€** |
| **OPEX annuel** | | | **~72-204 k€/an** |

**Leçon** : en CDN, **l'OPEX transit domine** — chaque To servi depuis le
cache plutôt que l'origine, c'est de la marge. Le dimensionnement du cache
(§74) est une décision financière.

## 80. Pièges CDN/edge — l'essentiel (détail §107)

- **P27** — Cache sur SATA seul pour des petits objets : IOPS saturées.
- **P28** — Serveur 800 mm dans une baie edge 600 mm : ne rentre pas.
- **P29** — Anycast sur sessions TCP longues sans persistance : coupures.
- **P30** — TTL infini sans purge : contenu périmé servi pendant des jours.

# PARTIE H — PLANS D'ADRESSAGE, CÂBLAGE, ÉTIQUETAGE

## 81. Plan d'adressage : sobre, hiérarchique, documenté

| Usage | Plage type | Masque | Note |
|---|---|---|---|
| Loopbacks | 10.0.0.0/24 (ou /32 dans un /16) | /32 | 1 par équipement, jamais réutilisée |
| Liens leaf-spine | 10.0.1.0/24 | **/31** (RFC 3021) | 2 adresses utiles, pas de broadcast gaspillé |
| Management / OOB | 192.168.0.0/20 | /24 par salle | Séparé physiquement si possible |
| Overlay (VTEP) | = loopbacks | /32 | Annoncées en underlay |
| Services (VIP, anycast) | 10.0.10.0/24 | /32 | — |

