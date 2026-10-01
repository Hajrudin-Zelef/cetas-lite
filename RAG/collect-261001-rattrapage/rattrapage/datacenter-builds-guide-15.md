---
id: collect-261001-rattrapage/rattrapage/datacenter-builds-guide-15
title: "Datacenter Builds — Le guide des BOMs"
domain: rattrapage
role: reference
task: reference
actors: ["Nvidia"]
dates: ["2026-09-27"]
keywords: ["arr", "compute", "distribution", "hyperscaler", "incident", "nvidia", "open source", "rubin"]
source: docs/RAG/collect-261001-rattrapage/datacenter_builds_guide.md
source_anchor: ""
source_lines: [2618, 2819]
sha256: 879a79d01033371080636ea78317b15e20743fa51e1658f29e2b46b600520174
---

# Datacenter Builds — Le guide des BOMs

| Étape | Calcul | Résultat |
|---|---|---|
| Σ P_réaliste | 40 kW | 40 kW |
| × 1,25 | — | 50 kW |
| Onduleur | 60 kVA / 60 kW (fp=1), N+1 → 2× 60 kW | 2× 60 kW |
| Prix indicatif | Easy UPS 3M 60 kW ≈ 25–35 k€ pièce (à vérifier) | **≈ 60 k€** |
| Batteries | 15 min à 50 kW | incluses ou +15 k€ |

Architecture : **2 onduleurs en parallèle redondant** (si l'un tombe, l'autre
prend tout), bypass de maintenance externe, 2 arrivées A/B derrière.

---

## 164. Autonomie : le calcul batteries

```
Énergie (kWh) = P(kW) × autonomie(h) / rendement_onduleur
```

- 50 kW × 0,25 h (15 min) / 0,95 = **13,2 kWh** → 1 armoire batteries.
- 50 kW × 2 h / 0,95 = **105 kWh** → 4–6 armoires (cher !).
- **Règle PME** : 15 min d'autonomie + groupe électrogène si > 30 min requis.
  Les batteries pour 2 h coûtent plus cher que le groupe.
- Li-ion vs VRLA : Li-ion = 2× le prix, 3× la durée de vie, 2× moins de place
  (détail : guide onduleurs).

---

## 165. ATS et redondance des arrivées

- Chaque rack : **2 PDU (A+B)** sur 2 onduleurs (ou 2 jeux de batteries)
  différents.
- Équipements mono-alimentation (vieux switch, console) : **ATS** (Automatic
  Transfer Switch) 16 A ≈ 1 300 € (vérifié le 27/09/2026, gamme APC).
- **Testez le basculement** : coupez A un dimanche, vérifiez que tout tient
  sur B. Puis l'inverse. Tous les ans.

---

## 166. Pièges terrain — Énergie

1. **Dimensionner en kVA au lieu de kW** : un 60 kVA à fp 0,8 = 48 kW. Avec
   des serveurs à 50 kW, ça disjoncte. **Toujours en kW.**
2. **Oublier le courant d'appel** : les PSU appellent 2× au démarrage. Un
   redémarrage simultané de 20 serveurs fait déclencher l'onduleur. Démarrage
   séquencé (PDU switched, §129).
3. **Batteries jamais testées** : test mensuel auto + test de décharge annuel.
   Une batterie morte = 0 minute d'autonomie.
4. **Pas de délestage** : en fin d'autonomie, arrêtez proprement (NUT —
   voir guide Proxmox, scénario d'arrêt auto). Un arrêt brutal corrompt.

---

## 167. Mise en service pas à pas

| Étape | Action | Durée |
|---|---|---|
| 1. Réception | Check-list §137 + §168, photos des cartons | 1 j |
| 2. Montage | Rails, serveurs (lourds en bas), PDU, câblage §144 | 1–2 j/rack |
| 3. Électrique | Phases §131, test A/B, serrage | 0,5 j |
| 4. Burn-in | §169, 48–72 h | 3 j |
| 5. Firmware | BMC, BIOS, NIC, SSD — versions notées | 1 j |
| 6. Réseau | iperf3, VLAN, LACP, management | 1 j |
| 7. OS/hyperviseur | Install, durcissement (voir guides OS) | 1–2 j |
| 8. Applicatif | Ceph, DB, VDI… par workload | variable |
| 9. As-built | §172 | 1 j |
| 10. Prod | progressive, §173 | 1–2 sem |

---

## 168. Check-list de réception matériel

- [ ] Colis : nombre, état, n° de série vs bon de livraison
- [ ] Chaque serveur : modèle, CPU, RAM (quantité + part-number), disques
- [ ] Firmware notés (BMC/BIOS/NIC/SSD) — photo des étiquettes
- [ ] Câbles/optiques : quantités, types
- [ ] PDU : ampérage, type (metered/switched)
- [ ] Garantie : enregistrée, date de fin notée
- [ ] Photos : rack vide, rack plein, câblage arrière

---

## 169. Burn-in : le protocole 72 h

```
J1 : memtest86+ (1 passe) + smartctl (tous disques)
J1-2 : stress-ng --cpu 0 --vm 4 (24 h) + surveillance temp/conso
J2 : fio randwrite 4k sur TOUS les SSD (2 h) — révèle les SSD faibles
J2-3 : iperf3 24 h sur chaque NIC (débit + 0 perte)
J3 : reboot × 5 (POST, BMC, boot) + re-vérification firmware
```

- **Tout défaut pendant le burn-in = retour fournisseur**, pas « on verra ».
- Rapport de burn-in archivé par n° de série (utile pour la garantie).
- Les 3XS et intégrateurs sérieux le font d'usine (24 h) — refaites 48 h
  quand même : le transport secoue.

---

## 170. Firmware : l'ordre des MAJ

1. **BMC/IPMI** d'abord (un BMC planté = plus d'accès distant).
2. BIOS/UEFI (notez les réglages : NUMA, virtualization, power profile).
3. NIC (firmware + option ROM).
4. SSD (firmware — voir §51 sur l'uniformité).
5. Backplane/expandeurs.

**Ne mettez jamais à jour tous les nœuds d'un cluster en même temps.**
Un par un, avec validation entre chaque (rolling upgrade).

---

## 171. Tests réseau à la mise en service

- [ ] iperf3 : débit nominal sur chaque lien (25G → ≥ 23 Gb/s utiles)
- [ ] 0 % de perte sur 24 h (sinon : câble, optique, ou port)
- [ ] LACP : débranchez un lien, le trafic continue
- [ ] VLAN : isolez, vérifiez qu'il n'y a pas de fuite
- [ ] Latence : ping < 0,2 ms en intra-rack (sinon : buffer, duplex)

---

## 172. Documentation as-built : le livrable

- Plan de rack (qui est où, §132) + plan de câblage (§147).
- Tableau électrique : phases, PDU, disjoncteurs, onduleur.
- Versions firmware par n° de série.
- Mots de passe (coffre, §134).
- Procédures : arrêt/démarrage, failover FW, restore backup.
- **Sans as-built, la salle n'est pas réceptionnée.** C'est contractuel.

---

## 173. Mise en production progressive

1. Semaine 1 : 10 % de la charge, surveillance renforcée.
2. Semaine 2 : 50 %, premier test de failover réel (coupez un nœud).
3. Semaine 3 : 100 %, exercice de restauration backup.
4. Mois 2 : revue thermique (caméra IR) + resserrage électrique.
5. **Ne déclarez jamais « c'est fini » avant le premier incident géré.**

---

## 174. Pièges terrain — Mise en service

1. **Sauter le burn-in** « pour gagner 3 jours » : le disque qui meurt à J+15
   en prod coûte 10× plus cher.
2. **Firmware hétérogènes** : 3 versions de BIOS sur 10 nœuds = comportements
   différents. Uniformisez.
3. **Mots de passe par défaut** sur les BMC : le premier scan Shodan les
   trouve. Changez AVANT de brancher le réseau.
4. **As-built « on le fera plus tard »** : plus tard = jamais.
5. **Mise en prod un vendredi** : non. Mardi ou mercredi matin, équipe complète.

---

## 175. PME vs hyperscaler : ce qui change à l'échelle

| Sujet | PME (1–10 racks) | Hyperscaler (1 000+ racks) |
|---|---|---|
| Serveurs | Intégrateurs (Supermicro, Dell) | Design propre (OCP) |
| Refroidissement | Air + RDHx | DLC généralisé, immersion |
| Puissance/rack | 5–30 kW | 100–240 kW (Rubin : 100–240 kW/rack) |
| Distribution élec | 400 V AC → PDU | **800VDC** (produits H2 2026) |
| Réseau | ToR 25/100G | 400/800G, topologies Clos |
| Achat | Devis, remise volume | Contrats pluriannuels directs |
| Staff | 1–5 personnes | Équipes dédiées par domaine |
| PUE cible | 1,4–1,6 | 1,1–1,2 |

---

## 176. 800VDC : l'état au 27/09/2026

- Constat : à 54 VDC (standard actuel in-rack), 1 MW = **18 500 A** → des
  barres de cuivre qui occuperaient 64U. À 800 VDC : **1 250 A**, −45 % de
  cuivre, rendement 83 % → 92 %+ (vérifié le 27/09/2026).
- Produits commerciaux (Vertiv, Schneider Electric, Eaton, Delta) :
  **2ᵉ semestre 2026**, alignés sur le rack NVIDIA Kyber.
- Pour la PME : **pas avant 2027–2028** en pratique. Mais prévoyez des
  chemins de câbles et des arrivées électriques **surdimensionnés** : le
  retrofit 800VDC passera par là.

---

## 177. OCP et designs ouverts : ce que la PME peut copier

- **Open Compute Project** : designs de racks, serveurs et PDU en open source.
- À copier : les principes (barres de puissance, airflow, gestion câbles),
  pas les pièces (pas de supply chain OCP en petite série).
- Les intégrateurs « OCP-inspired » (Supermicro notamment) apportent 80 %
  des gains pour 20 % de l'effort.

---

## 178. Ce que la PME doit copier / ne pas copier

**À copier :**
- Le monitoring exhaustif (tout est mesuré chez les hyperscalers).
- L'automatisation (provisioning, firmware — Ansible, voir guide ansible).
- Le burn-in systématique et les procédures écrites.
- Le PUE comme KPI financier.

