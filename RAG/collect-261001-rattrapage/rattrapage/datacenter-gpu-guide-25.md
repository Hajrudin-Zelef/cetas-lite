---
id: collect-261001-rattrapage/rattrapage/datacenter-gpu-guide-25
title: "Les GPU datacenter / IA — présent et futur vérifié"
domain: rattrapage
role: reference
task: reference
actors: ["AMD", "DeepSeek", "Mistral", "Nvidia", "SGLang", "vLLM"]
dates: ["2026-09-27"]
keywords: ["datacenter", "gpu", "amd", "attention", "blackwell", "deepseek", "distribution", "dpo", "ethernet", "fine-tuning", "fp4", "fp8"]
source: docs/RAG/collect-261001-rattrapage/datacenter_gpu_guide.md
source_anchor: ""
source_lines: [3778, 3933]
sha256: 44d2d0fbad3b9b830b5e1c623206055fa879686bc3139fadf6eff81e3eba0d39
---

# Les GPU datacenter / IA — présent et futur vérifié

```
1.  VRAM > TFLOPS : dimensionner sur la mémoire d'abord.
2.  70B FP8 = 80 Go → RTX PRO 6000 (96 Go) à 12 k$ = sweet spot.
3.  70B FP16 = 161 Go → H200 NVL ou MI300X (1 GPU).
4.  Contexte 128K : le KV cache dépasse les poids → MI325X/355X.
5.  Training : H100 (standard) / B200 (frontier) / MI355X (alternatif).
6.  NVLink intra-nœud, IB inter-nœuds, Ethernet UEC en 2027.
7.  B200 = 1 000 W → DLC. MI355X = 1 400 W → DLC obligatoire.
8.  Nœud 8× B200 = 10,6 kW au mur. Rack 4 nœuds = 42 kW.
9.  PUE 1,1 (DLC) vs 1,5 (air) = 35 k€/nœud/5 ans.
10. Onduleur 1,25× charge IT, 10 min d'autonomie, VFI.
11. Acheter si usage > 60-70 %, louer sinon. Break-even ~15 mois.
12. 3 devis OEM minimum, référence exacte exigée.
13. Burn-in 24 h avant réception. DCGM dès le jour 1.
14. CUDA = risque zéro. ROCm = POC 4 semaines.
15. FP8 par défaut, FP4 validé au cas par cas.
16. Batching continu = levier n°1 du $/token (×36 à batch 64).
17. Rubin (H2 2026) et Helios/MI450 (fin Q3 2026) : attendre les POC.
18. Rumeurs = rumeurs : ne jamais acheter dessus.
19. Le MWh est la vraie monnaie du datacenter IA.
20. Mesurer, comparer, ajuster — le terrain a toujours raison.
```

---

*Fin du guide — 214 sections, 4 000+ lignes, données vérifiées le 27/09/2026.*
## 215. Cas pratique : salle complète 8× B200, de zéro à la prod

| Étape | Durée | Coût indicatif |
|---|---|---|
| Étude élec + froid (bureau d'études) | 4 semaines | 8–15 k€ |
| Travaux élec (TGBT 60 kVA, onduleur 60 kVA, câblage) | 8–12 semaines | 60–90 k€ |
| DLC (CDU 2× 40 kW, dry cooler, tuyauterie) | 8–12 semaines | 80–120 k€ |
| Commande serveur 8× B200 (Dell/HPE/SMC) | 12–20 semaines | 350–450 k$ |
| Réseau (IB NDR, câbles, switch) | 4 semaines | 40–60 k$ |
| Installation + burn-in + recette | 3 semaines | 15–25 k€ |
| Logiciel (vLLM, monitoring, runbook) | 4 semaines | 1 ETP |
| **Total** | **6–9 mois** | **~600–750 k€** |

Le goulot n'est pas le GPU : c'est l'électricien et le frigoriste. Les commander
**le même jour** que le serveur.

## 216. Cas pratique : migration air → DLC sur nœuds existants

1. **Audit** : mesurer la charge réelle par nœud (pas la plaque) — souvent
   30 % sous le TDP nominal en inférence.
2. **CDU** : 1 CDU 80 kW couvre 4–6 nœuds 8× H100.
3. **Plaques froides** : retrofit possible sur HGX (kits OEM) ou remplacement
   des nœuds à la prochaine génération.
4. **Bascule** : nœud par nœud, jamais toute la salle d'un coup.
5. **ROI** : sur 50 kW IT, PUE 1,5→1,15 = 150 MWh/an = 22 k€/an — le retrofit
   (150–250 k€) se paie en 7–10 ans… ou en 3 ans si la salle était saturée en
   froid et que le DLC libère de la capacité pour de nouveaux nœuds.

---

*Fin du guide — 216 sections, 4 000 lignes, données vérifiées le 27/09/2026.*
## 217. Pièges supplémentaires (série 5)

1. **Le « 8× GPU » sans switch IB** : 2 nœuds 8× B200 en Ethernet 100G pour du
   training = 5× moins vite qu'en IB. Le réseau n'est pas une option.
2. **Le driver trop récent** : un driver NVIDIA tout frais peut casser vLLM
   (incompatibilité CUDA). Figer les versions driver/CUDA/framework ensemble.
3. **Le NVMe plein pendant un run** : un checkpoint qui échoue à 3 h du matin
   parce que le disque est plein = run perdu. Alerte à 80 %, purge auto.
4. **Le BIOS en mode « éco »** : certains serveurs livrent avec le profil
   énergétique « balanced » qui bride les GPU à 80 %. Passer en « performance ».
5. **Le câble 12VHPWR mal enfoncé** : la cause n°1 des Xid 79 sur les cartes
   PCIe. Le connecteur doit être **complètement** inséré (pas de jour visible).
6. **L'horloge système dérivée** : NTP en panne → les logs sont inexploitables
   et les certificats TLS tombent. Monitorer l'offset NTP.
7. **Le partage de GPU sans MIG** : 2 process sur le même GPU sans isolation =
   OOM aléatoires. MIG ou time-slicing, jamais du « à l'arrache ».
8. **L'oubli du firmware BMC** : un BMC vulnérable (CVE) expose tout le nœud.
   Patch BMC + BIOS dans le plan de maintenance trimestriel.

## 218. Glossaire : 20 termes supplémentaires

| Terme | Définition |
|---|---|
| NVFP4 | Format FP4 de NVIDIA (échelle fine), génération Rubin/Blackwell |
| MXFP4 | Format FP4 du standard OCP (échelle partagée), AMD CDNA 4 |
| UALink | Interconnexion scale-up ouverte (alternative à NVLink) |
| UEC | Ultra Ethernet Consortium — Ethernet pour l'IA/HPC |
| NVL72 | Rack NVIDIA de 72 GPU interconnectés (GB200, Rubin) |
| Helios | Rack AMD de 72 GPU MI450 (fin 2026) |
| Sidecar 800VDC | Module de distribution 800 V continu (racks Rubin) |
| TTFT | Time To First Token — latence du premier token |
| TPOT | Time Per Output Token — cadence de génération |
| RadixAttention | Cache de préfixes arborescent (SGLang) |
| PagedAttention | Gestion paginée du KV cache (vLLM) |
| DPO | Direct Preference Optimization — alignement sans RL |
| ORPO | Odds Ratio Preference Optimization — variante DPO |
| QLoRA | LoRA sur base quantifiée 4 bits — fine-tuning économe |
| ZeRO | Zero Redundancy Optimizer — sharding des états d'entraînement |
| FSDP | Fully Sharded Data Parallel — ZeRO natif PyTorch |
| MFU | Model FLOPs Utilization — efficacité du training (30–60 %) |
| GQA | Grouped Query Attention — KV cache réduit (Llama 3, Mistral) |
| MLA | Multi-head Latent Attention — KV cache compressé (DeepSeek) |
| MoE | Mixture of Experts — seuls quelques experts actifs par token |

---

*Fin du guide — 218 sections numérotées, 4 000 lignes, données vérifiées le 27/09/2026.*
## 219. Synthèse exécutive (pour la direction, 1 page)

**Le marché (27/09/2026)** : NVIDIA domine (H100/B200), AMD est une alternative
crédible (MI350X), la génération Rubin/MI450 arrive fin 2026. Les prix H100
baissent (-40 % depuis 2023).

**L'investissement type** : 1 nœud 8× GPU = 250–450 k$ de matériel + 150–250 k€
d'élec/froid/réseau. Compter 6–9 mois de projet.

**La règle d'or** : acheter si utilisation > 60–70 % et horizon > 18 mois,
louer sinon. Le cloud néocloud est moins cher que le self-hosted sous-utilisé.

**Le risque n°1** : l'énergie. 30–45 % du budget projet, 20 % du TCO 5 ans.
Le dimensionnement électrique et le PUE décident de la rentabilité.

**La recommandation** : commencer petit (1 nœud ou 1 GPU), mesurer, puis
étendre. Chaque étape se rentabilise avant la suivante.

## 220. Checklist ultime : « ai-je tout ? » (avant de signer)

- [ ] Références GPU exactes + format + vBIOS (§87)
- [ ] 3 devis OEM comparés au $/GPU tout compris (§75)
- [ ] BOM complète : serveur, réseau, élec, froid, licences (§34–36)
- [ ] Bilan de puissance signé par bureau d'études (§113)
- [ ] PUE cible et plan de mesure (§29, §170)
- [ ] Onduleur 1,25× + 10 min + groupe si critique (§58)
- [ ] DLC : CDU, dry cooler, commissioning (§135)
- [ ] POC logiciel 4 semaines si AMD ou nouveau framework (§197)
- [ ] Burn-in 24 h + PV de recette (§139, §211)
- [ ] Runbook + monitoring + journal de bord (§171, §170, §199)
- [ ] Budget formation équipe (§62)
- [ ] Assurance actifs à jour (§205)

*Si les 12 cases sont cochées, signer. Sinon, compléter d'abord.*

---

*FIN — Guide GPU Datacenter/IA, 220 sections, 4 000 lignes, vérifié le 27/09/2026.*
## 221. Annexe : gabarits de câbles par puissance (rappel §58)

| Charge IT | Tri 400 V | DJ tête | Câble (cuivre, 30 m) | Onduleur |
|---|---|---|---|---|
| 5 kW | 9 A | 16 A courbe D | 5G6 | 10 kVA |
| 11 kW | 18 A | 32 A courbe D | 5G10 | 15 kVA |
| 22 kW | 36 A | 63 A courbe D | 5G16 | 30 kVA |
| 44 kW | 71 A | 100 A | 5G35 | 60 kVA |
| 85 kW | 137 A | 200 A | 5G70 | 120 kVA |

Valeurs indicatives — la note de calcul du bureau d'études fait foi.

## 222. Annexe : planning type Gantt (projet 8× B200)

