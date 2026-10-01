---
id: collect-261001-rattrapage/rattrapage/datacenter-builds-guide-11
title: "Datacenter Builds — Le guide des BOMs"
domain: rattrapage
role: reference
task: reference
actors: ["AMD", "Intel"]
dates: ["2026-09-27"]
keywords: ["arr", "benchmark", "compute", "dram", "gpu", "intel", "luna"]
source: docs/RAG/collect-261001-rattrapage/datacenter_builds_guide.md
source_anchor: ""
source_lines: [1823, 2016]
sha256: d1ca9beb404a4b099ab1cbb5c2856bc4e4244d7a8cadadcdd0c1ba17ae834c02
---

# Datacenter Builds — Le guide des BOMs

| Composant | Référence type | Qté | Prix |
|---|---|---|---|
| 1U EPYC 9124 (16c), 2× 25G SFP28 | AS-1115CS-TNR nu | 1 | 4 585 € |
| EPYC 9124 | — | 1 | 749 € |
| RAM 2× 32 Go = 64 Go ECC | DDR5 | 2 | 1 500 € |
| SSD M.2 512 Go (miroir) | 2× | 2 | 160 € |
| NIC 2× 25G SFP28 (si non incluse) | E810 | 1 | 450 € |
| **TOTAL** | ×2 pour HA | | **≈ 7 445 € × 2 = 14 890 €** |
| Conso | | | **≈ 250 W / unité** |

---

## 113. HA firewall : l'architecture qui ne tombe pas

```
        Internet
            |
     +------+------+
     |             |
  FW-1 (CARP)  FW-2 (CARP)   ← sync états pfsync (lien dédié 10G)
     |             |
     +------+------+
            |
      Switch LAN
```

- **2 nœuds mini**, CARP (OPNsense) ou équivalent, sync des états sur lien
  dédié (pas sur le LAN de prod).
- Testez le **failover réel** (débranchez FW-1) tous les 6 mois.
- Les 2 FW doivent avoir la **même version OPNsense** : MAJ l'un après
  l'autre, jamais les deux ensemble.

---

## 114. Pièges terrain — Firewall

1. **Dimensionner au débit du port** (10G) au lieu du débit réel (2 Gb/s
   opérateur) : on sur-paie ×5. Dimensionnez au débit facturé + 50 %.
2. **Oublier la chute Suricata** : 10 Gb/s annoncés → 2 Gb/s avec IPS.
   Testez AVEC votre ruleset.
3. **NIC Realtek** : pas de pilote stable sous FreeBSD. **Intel uniquement**
   (i226, X710, E810).
4. **HA sans lien de sync dédié** : la sync sur le LAN sature et le failover
   perd les états.
5. **Pas de console série** : quand le FW ne boot plus, l'IPMI/serial est le
   seul accès. Vérifiez avant la mise en rack.

---

## 115. Workload 11 — HSM : principes (pas de BOM interne)

Un HSM (Hardware Security Module) est une **appliance réseau** : on ne la
construit pas, on l'intègre. Rôle : génération, stockage et usage des clés
cryptographiques dans un boîtier inviolable (tamper-evident/resistant),
certifié **FIPS 140-3 niveau 3**.
- Acteurs : **Thales Luna** (leader), **Entrust nShield**, Utimaco, Atos.
- Modèles 2026 : **Thales Luna 8** (annoncé 2026, FIPS 140-3 L3, support
  post-quantique — vérifié le 27/09/2026), Entrust nShield (gammes XC).
- Prix : **sur devis uniquement** (non trouvé au 27/09/2026 en prix public ;
  ordre de grandeur historique : 15 000–50 000 €/appliance selon perfs —
  à vérifier).

---

## 116. Intégration au rack — prérequis

| Prérequis | Détail |
|---|---|
| Format | 1U rackable, 2× alim redondantes |
| Conso | 100–300 W par appliance (à vérifier selon modèle) |
| Réseau | 2× 1/10 GbE dédiés (réseau crypto isolé, VLAN ou physiquement séparé) |
| Redondance | **2 appliances mini** (cluster), jamais une seule |
| Accès | Console série + réseau de management isolé |
| Environnement | 0–35 °C, comme un serveur standard |

Comptez dans la PDU : **2× 0,3 kW** pour la paire HSM, et 2U de rack.

---

## 117. Mise en service HSM : la cérémonie des clés

1. **Cérémonie d'initialisation** : génération de la clé maîtresse avec
   quorum (ex : 3 porteurs sur 5). Filmez/documentez, témoins requis.
2. **Sauvegarde chiffrée** des clés (tokens/clés de récupération) dans **2
   coffres géographiques distincts**.
3. **Intégration applicative** : PKCS#11, JCE, ou API REST selon modèle.
   Prévoyez 2–5 j d'intégration par application.
4. **Plan de rotation** : clés applicatives tous les 1–2 ans, documenté.
5. **Fin de vie** : procédure de destruction (zéroïsation) écrite avant la
   mise en prod.

Un HSM sans cérémonie documentée = un HSM dont personne ne peut prouver
l'intégrité devant un auditeur. **C'est un sujet conformité, pas juste
technique.**

---

## 118. Pièges terrain — HSM

1. **Un seul HSM** : panne = toutes les applis crypto à l'arrêt. Toujours 2.
2. **Clés de récupération dans le même bâtiment** : incendie = perte totale.
3. **Oublier la montée en charge** : un HSM fait 1 000–20 000 op/s selon
   modèle — dimensionnez aux pics (TLS, signatures).
4. **Post-quantique** : en 2026, exigez le support des algo PQC (ML-KEM,
   ML-DSA) — le Luna 8 l'annonce. Un HSM sans PQC est un investissement
   à durée de vie courte.

---

## 119. Workload 12 — Méthode : chiffrer soi-même (template)

Feuille de calcul type (une ligne par serveur, colonnes ci-dessous) :

| Serveur | CPU (TDP) | RAM (Go) | Disques | GPU (TDP) | NIC | P_max | P_réaliste | Prix |
|---|---|---|---|---|---|---|---|---|
| srv-01 | 280 W | 384 | 4×NVMe | — | 25G | 0,62 kW | 0,47 kW | 29 750 € |
| ... | | | | | | | | |

Puis :
1. **Σ P_réaliste** → dimensionnement onduleur + refroidissement (§155, §146).
2. **Σ P_max × 1,25** → dimensionnement PDU + PSU.
3. **Σ prix × 1,1** → budget (10 % intégration/câblage).
4. **TCO 5 ans** : + électricité (§14) + maintenance (10 %/an).

---

## 120. Exemple complet : rack de 8 serveurs mixtes

| Serveur | Type | P_réaliste | Prix |
|---|---|---|---|
| 2× Compute M | §20 | 2 × 0,95 = 1,9 kW | 114 k€ |
| 2× PG OLTP S | §28 | 2 × 0,6 = 1,2 kW | 93 k€ |
| 2× Ceph OSD S | §42 | 2 × 0,6 = 1,2 kW | 99 k€ |
| 1× Backup S | §102 | 0,5 kW | 25 k€ |
| 1× Firewall 10G (paire = 2×) | §111 | 2 × 0,1 = 0,2 kW | 6,7 k€ |
| 2× switch 25G | — | 2 × 0,3 = 0,6 kW | 16 k€ |
| **TOTAL rack** | 11 serveurs + 2 switch | **≈ 5,6 kW** | **≈ 354 k€** |

5,6 kW dans un 42U : **air cooling standard**, 1 PDU 3× 32 A suffit
(22 kW). C'est le rack « PME dense » type.

---

## 121. Marge et croissance : la règle des 3 ans

- Achetez la capacité **J+36 mois**, pas J+0 : un serveur vit 5 ans, le
  remplir à 90 % le jour 1 = racheter dans 18 mois.
- Mais n'achetez pas J+60 mois : la DRAM/GPU de 2029 sera moins chère et
  meilleure. **Le juste milieu : 30 % de marge.**
- Prévoyez **4U libres par rack** et **30 % de ports switch libres**.

---

## 122. Devis : ce qu'il faut exiger du fournisseur

- [ ] Références exactes (CPU stepping, RAM part-number, SSD modèle + firmware)
- [ ] Prix unitaires **et** remise par volume, valables 90 j
- [ ] Délais **écrits** par ligne (surtout GPU et DRAM en 2026)
- [ ] Garantie : durée, NBD vs 4 h, pièces sur site
- [ ] Burn-in inclus ? (24–72 h, rapport fourni)
- [ ] Firmware livré : versions notées
- [ ] Reprise de l'ancien matériel (DEEE)

---

## 123. Benchmark avant achat : les 4 tests

1. **CPU** : `stress-ng --cpu 0 --timeout 3600` + mesure conso à la PDU.
2. **RAM** : `memtester` ou memtest86+ (1 passe complète mini).
3. **Disques** : `fio` randread/randwrite 4k QD32 — comparez à la fiche
   (PM9A3 : ≈ 900 k/200 k IOPS).
4. **Réseau** : `iperf3` bidirectionnel sur chaque NIC à la vitesse nominale.

Un serveur qui ne passe pas ces 4 tests **ne va pas en prod**. Point.

---

## 124. Coûts unitaires de pilotage (€/vCPU, €/To, €/token)

| Métrique | Calcul (exemples du guide) | Valeur 2026 |
|---|---|---|
| €/cœur (compute) | 56 975 € / 192 cœurs | ≈ 297 € |
| €/To utile NVMe | §47 | ≈ 3 090 € |
| €/To utile hybride | §57 | ≈ 307 € |
| €/poste VDI/an | §67 | ≈ 418 € |
| €/M tokens (70B on-prem) | §87 | ≈ 0,9 € |

Suivez ces 5 métriques par trimestre : elles disent si votre infra reste
compétitive vs cloud.

---

## 125. Tableau récapitulatif — consommation de toutes les BOMs

