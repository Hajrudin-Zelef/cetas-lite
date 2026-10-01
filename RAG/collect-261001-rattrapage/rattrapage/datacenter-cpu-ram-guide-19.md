---
id: collect-261001-rattrapage/rattrapage/datacenter-cpu-ram-guide-19
title: "CPU SERVEUR, RAM SERVEUR, CHÂSSIS & VENTILATION"
domain: rattrapage
role: reference
task: reference
actors: ["AMD", "Oracle"]
dates: []
keywords: ["arr", "capex", "distribution", "gpu", "kv cache"]
source: docs/RAG/collect-261001-rattrapage/datacenter_cpu_ram_guide.md
source_anchor: ""
source_lines: [2989, 3165]
sha256: 003197abf90a0e8f505d69f5f4a28eb02966abd91a20e52bff45616e75474b38
---

# CPU SERVEUR, RAM SERVEUR, CHÂSSIS & VENTILATION

Après chaque vague : vérifiez les courbes de ventilation et la fréquence
mémoire négociée (les défauts changent parfois, sections 60 et 148).

## 165. Gestion thermique de la salle : confinement

Sans confinement, l'air chaud se mélange à l'air froid : les serveurs du
haut de baie respirent l'air chaud de ceux du bas (+5 à +15 °C).

| Solution | Efficacité | Coût (indicatif) |
|---|---|---|
| Caches de baie (section 112) | base obligatoire | ~5 €/U |
| Confinement allée froide (portes/toit) | +15-25 % | quelques k€/allée |
| Confinement allée chaude (cheminée) | +15-25 % | quelques k€/allée |
| Dalles perforées bien placées | +10 % | — (réglage) |

Règle : **confinez d'abord, climatisez ensuite.** Chaque degré gagné à
l'entrée des serveurs, c'est des ventilateurs qui tournent moins vite
(loi cubique, section 62) et des CPU qui ne throttlent pas l'été.

## 166. Plan de câblage : méthode

1. **Schéma unifilaire** : arrivées, onduleur, tableaux, PDU A/B, baies —
   affiché dans le local (et dans la GED).
2. **Repérage** : chaque câble étiqueté aux deux extrémités (tenant/
   aboutissant), chaque prise PDU numérotée, chaque serveur identifié
   (nom + U).
3. **Chemins** : énergie d'un côté, réseau de l'autre (section 79) ;
   goulottes dimensionnées pour +30 % (extensions).
4. **Réserve** : 2U libres par baie + 20 % de prises PDU libres — la
   croissance ne prévient pas.
5. **Photo** : photographiez chaque baie après câblage. La photo « propre »
   est la référence pour détecter les dérives.

## 167. Documentation d'exploitation minimale

Pour chaque serveur : fiche d'identité (modèle, n° série, CPU, RAM
détaillée, firmware), rôle, criticité, contrats (garantie, support),
procédures (redémarrage, arrêt ordonné, remplacement à chaud), schéma
réseau et électrique.

Pour la salle : plan, unifilaire électrique, PUE suivi, planning de
maintenance, contacts d'astreinte (électricien, climaticien, OEM).

**La documentation n'est pas de la bureaucratie : c'est ce qui permet à
quelqu'un d'autre que vous de gérer la panne un dimanche à 3h du matin.**

## 168. PRA/PCA : le rôle des serveurs

- **PRA (reprise)** : RTO/RPO définis par application, pas par serveur.
  Les serveurs du PRA doivent être **testés** (bascule réelle annuelle) —
  un PRA jamais testé est une fiction.
- **PCA (continuité)** : redondance géographique, réplication synchrone/
  asynchrone selon le RPO.
- Les serveurs du site de secours consomment aussi (même à l'arrêt, le
  socle idle, section 85) : intégrez-les au dimensionnement onduleur et
  au TCO.

## 169. Pièces de rechange : stock minimal

| Pièce | Stock conseillé | Pourquoi |
|---|---|---|
| Modules ventilateurs | 1-2 par modèle | panne la plus fréquente |
| Alimentations | 1 par modèle | criticité |
| Barrettes RAM (QVL) | 2 par capacité | ECC/diagnostic |
| Disques NVMe | 1-2 par modèle | usure |
| Rails/caches | 1 kit | casse au montage |

Coût total du stock : quelques milliers d'euros — contre des heures
d'indisponibilité à chaque panne sans pièce. **Le stock se dimensionne
avec le taux de panne observé**, pas au doigt mouillé : suivez vos
remplacements (section 163).

## 170. Fin de vie : procédure type

1. **Sauvegarde et vérification** des données à conserver.
2. **Effacement certifié** des disques (norme type NIST 800-88, attestation).
3. **Destruction physique** si données sensibles (broyeur, attestation).
4. **Dépollution DEEE** : écrans, batteries, cartes (prestataire agréé,
   bordereau de suivi).
5. **Revente/don** du matériel banalisé (serveurs sans disques).
6. **Sortie d'inventaire** : mise à jour CMDB, fin des contrats de
   maintenance (ne pas payer un an de plus pour un serveur parti).

**Anticipez** : négociez la reprise à l'achat (section 122), planifiez
les sorties par vagues annuelles, jamais « quand on aura le temps ».

---

## 171. Cas d'école A : PME, 3 nœuds Proxmox (BOM + TCO indicatifs)

Besoin : 80 VM, 1 To utile, budget serré, pas de licences coûteuses.

| Poste | Choix | Qté/nœud |
|---|---|---|
| Châssis | 2U 1S, N+1 fans | 3 |
| CPU | EPYC 9455P (48 c., 300 W, ~4 300 $ tarif) | 3 |
| RAM | 12× 64 Go DDR5-5600 RDIMM = 768 Go | 36 |
| Stockage | 6× NVMe 3,84 To (Ceph) | 18 |
| Réseau | 2× 25 GbE | 3 cartes |
| PSU | 2× 1200 W Titanium | 6 |

TCO 5 ans (indicatif) : CAPEX ~60 000 € (3 nœuds) ; énergie : 3 × 700 W
moyens × 8760 × 1,5 × 0,18 ≈ 5 000 €/an → 25 000 € ; maintenance ~12 000 €.
**Total ~97 000 €**, soit ~20 000 €/an pour 80 VM = **~250 €/VM/an**.
C'est le chiffre à comparer avec le cloud public.

## 172. Cas d'école B : base Oracle, licences optimisées (indicatif)

Besoin : Oracle SE, 400 Go chauds, minimiser les licences (facteur 0,5).

| Poste | Choix | Justification |
|---|---|---|
| CPU | 1× EPYC 9375F (32 c., 4,8 GHz) | 32 × 0,5 = 16 licences vs 64 |
| RAM | 12× 64 Go = 768 Go | dataset + 50 % |
| Stockage | NVMe faible latence, redo séparés | I/O log |
| Châssis | 2U 1S | pas de NUMA |

Économie licences : 32 cœurs au lieu de 64 = **2× moins de licences** pour
une performance OLTP souvent supérieure (fréquence). Sur 5 ans, l'économie
licences (dizaines de k€) dépasse le surcoût du CPU « F ». **C'est le cas
d'école du « CPU cher qui coûte moins cher ».**

## 173. Cas d'école C : inférence IA sur CPU (indicatif)

Besoin : inférence LLM 8B quantifié, 500 req/s, pas de GPU (contrainte
budgétaire ou de disponibilité).

| Poste | Choix | Justification |
|---|---|---|
| CPU | 2× Xeon 6980P (AMX) | AMX = 2×+ le débit EPYC (section 22) |
| RAM | 24× 64 Go MRDIMM-8800 = 1,5 To | modèle + KV cache en RAM, BP max |
| Réseau | 2× 100 GbE | distribution des requêtes |

Note : dès que le budget le permet, **un GPU modeste remplace ce serveur**
pour l'inférence (10-50× le débit). L'inférence CPU n'est pertinente qu'en
transition ou pour des modèles petits/moyens à faible volumétrie.

## 174. Cas d'école D : baie HPC 8 nœuds (indicatif)

8 nœuds type section 33 (2× 6980P + MRDIMM, ~2000 W chacun en charge).

| Poste | Calcul | Résultat |
|---|---|---|
| IT max | 8 × 2000 W | **16 kW** |
| Rack | 8× 2U = 16U + switch IB | ~20U |
| Froid | 16 kW à PUE 1,3 (DLC partiel) | ~21 kW à évacuer |
| Onduleur | 16 × 1,25 / 0,9 | **~22 kVA** |
| Réseau | 200 GbE par nœud + switch | budget réseau ≈ 30 % du CAPEX |

Leçon : en HPC, **le réseau et le froid coûtent aussi cher que les
serveurs**. Un devis « serveurs seuls » sous-estime le projet de 40-60 %.

## 175. Tableau comparatif final : que choisir en 2026 ?

| Profil | CPU recommandé | RAM | Format | Budget énergie/nœud |
|---|---|---|---|---|
| PME virtualisation | EPYC 9455P/9655P 1S | 768 Go RDIMM | 2U | ~700 W |
| ETI virtualisation dense | EPYC 9655 2S ou Xeon 6+ | 1,5 To RDIMM | 2U | ~1100 W |
| DB OLTP (licences) | EPYC 9375F/9575F | 768 Go-1,5 To | 2U 1S | ~600 W |
| DB analytique | Xeon 6980P + MRDIMM | 1,5 To MRDIMM | 2U | ~1700 W |
| HPC | Xeon 6980P + MRDIMM | 1,5 To MRDIMM | 2U DLC | ~2000 W |
| Cloud/scale-out | Xeon 6990E+ ou EPYC 9965 | 768 Go-1,5 To | 1U/2U | ~1000-1300 W |
| Inférence CPU | Xeon 6980P (AMX) | 1,5 To MRDIMM | 2U | ~1700 W |
| Edge | EPYC 8004 / Xeon 6500 | 256-512 Go | 1U/2U court | ~300-500 W |

## 176. FAQ acheteur (1/2)

**Faut-il attendre Venice/Diamond Rapids ?** Si votre besoin est immédiat
et standard : non, Turin/Xeon 6 sont matures et négociables. Si vous visez
le haut de gamme fin 2026 : oui, attendez prix et tests.

**1S ou 2S ?** 1S si : licences au socket, pas de NUMA voulu, redondance
par cluster. 2S si : densité, VM nombreuses, budget/baie contraint.

**DDR5-5600 ou 6400 ?** 5600 par défaut ; 6400 si workload sensible à la
bande passante (section 30) ; MRDIMM si HPC/IA sur Xeon 6900.

