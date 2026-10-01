---
id: collect-261001-ia-llm/ia-llm/ia-grand-dossier-14
title: "IA — Le grand dossier"
domain: ia-llm
role: reference
task: reference
actors: ["AWS", "DeepSeek", "Google", "Huawei", "Microsoft", "Nvidia", "OpenAI", "Samsung", "United States"]
dates: []
keywords: ["agent", "agents", "ascend", "aws", "blackwell", "chatgpt", "compute", "datacenter", "deepseek", "ethernet", "gpu", "hbm"]
source: docs/RAG/collect-261001-ia-llm/ia_grand_dossier.md
source_anchor: ""
source_lines: [868, 950]
sha256: 7214addf5fe3a7d3c3fb6b9555276d6f71a2976964eff71e4c1b5d8222139c10
---

# IA — Le grand dossier

**Lecture :** entre 2012 et 2024, le compute des runs frontières a été multiplié par ~10^7 (10 millions de fois). Entre 2020 (GPT-3) et 2024 (Llama 3 405B), par ~100. **C'est cette courbe que le camp « mur » dit en train de s'aplatir** — non pas que le compute n'augmente plus, mais que chaque ×10 rapporte moins de capacité visible.

### 9.3. L'efficacité matérielle : le vrai moteur silencieux

| GPU NVIDIA | Année | FP16 dense (TFLOPS, ordre) | Mémoire | Contexte |
|---|---|---|---|---|
| GTX 580 | 2010 | ~1,6 | 1,5 Go | AlexNet en utilisait 2 |
| P100 (Pascal) | 2016 | ~21 | 16 Go | Ère ResNet/AlphaGo |
| V100 (Volta) | 2017 | ~125 (Tensor) | 16/32 Go | Ère Transformer/BERT |
| A100 (Ampère) | 2020 | ~312 | 40/80 Go | Ère GPT-3/ChatGPT |
| H100 (Hopper) | 2022 | ~990 | 80 Go (HBM3) | Ère GPT-4/Llama |
| B200 (Blackwell) | 2024 | ~2 250 | 192 Go (HBM3e) | Ère raisonnement |
| Rubin (annoncé) | 2026+ | « à vérifier » | HBM4 | Ère agents |

Entre la GTX 580 et la B200 : **~1 400×** en débit FP16 en 14 ans — soit ~×2 tous les ~1,3 an, plus vite que la loi de Moore classique sur la période. **Une partie du « scaling » attribué aux idées est en réalité du scaling matériel.** Quand le camp « continuation » dit « le compute devient moins cher », c'est largement de ça qu'il parle — et quand le camp « mur » dit « ça ne suffit plus », il répond que le matériel seul ne compense plus les rendements décroissants algorithmiques.

### 9.4. Structure d'un budget d'entraînement (ordres Epoch AI)

Pour un run frontière type (2023-2024) :

| Poste | Part du total | Commentaire sysadmin |
|---|---|---|
| Hardware (amortissement GPU, réseau, DC) | 47-65 % | Le poste n° 1, et de loin |
| Personnel (chercheurs, equity incluse) | 29-49 % | Les talents coûtent presque autant que le silicium |
| Électricité | 2-6 % | Contre-intuitif : l'énergie est marginale **dans le budget training**... |
| ...mais l'énergie est le **goulot physique** | — | ...car sans MW disponibles, les GPU ne tournent pas (voir 4.6) |

**Coût total de développement ≈ 2-3× le run final** (sweeps d'hyperparamètres, runs avortés, ablations). Quand un labo annonce « 5,576 M$ » (DeepSeek-V3), c'est **le run final uniquement** — le coût programme complet est un multiple, débattu (fourchettes d'analystes « à vérifier »).

### 9.5. Le coût de l'inférence : le calcul que tout sysadmin doit savoir faire

**Coût par token (self-hosting), formule simplifiée :**

```
coût par 1M tokens = coût horaire total des GPU / (tokens/seconde × 3600) × 1 000 000
```

Exemple : un 70B Q4 sur 2× H100 (location ~4 $/h/GPU, soit 8 $/h au total), débit ~150 tok/s :
production = 150 × 3600 = **540 000 tokens/h** ; coût = 8 / 540 000 × 1 000 000 ≈ **14,81 $/1M tokens**. À comparer aux ~8-15 $/1M tokens d'une API moyenne : dans cet exemple, **le self-hosting n'est pas moins cher au token** — il devient rentable quand le débit réel est élevé, les GPU bon marché (instances spot, cartes prosumer) et **l'utilisation soutenue** (le coût fixe tourne même à 3h du matin). D'où la règle honnête : **API en dessous de ~50-100M tokens/mois ou usage irrégulier ; self-hosting au-dessus, à condition d'un taux d'utilisation élevé** — et toujours avec du MLOps (monitoring, mises à jour, redondance). Tout chiffrage sérieux commence par mesurer VOTRE débit réel (tok/s) sur VOTRE modèle, pas par recopier un exemple.

**Le multiplicateur « raisonnement » :** une requête o1-like consomme 5-20× plus de tokens (thinking) qu'une requête standard. **Un agent autonome en boucle peut consommer 100-1 000×** une requête simple. C'est pour ça que le FinOps IA (budgets par cas d'usage, routage, plafonds) est devenu une discipline — voir 6.3 et 6.5.

---

## 10. Histoire parallèle : le hardware et l'infrastructure (2012 → 2026)

> L'IA moderne est une histoire de **silicium** autant que d'algorithmes. Cette annexe donne au sysadmin la frise matérielle qui sous-tend tout le dossier.

### 10.1. Frise hardware

- **2006 — CUDA (NVIDIA).** Jensen Huang parie que les GPU deviennent des processeurs généralistes programmables. À l'époque, idée excentrique ; rétrospectivement, le pari le plus rentable de la tech.
- **2012 — GTX 580.** AlexNet tourne sur 2 cartes gamer. Le message : le deep learning est à portée d'un budget labo.
- **2016 — P100 / NVLink.** Le GPU datacenter avec interconnect rapide : l'entraînement distribué multi-GPU devient praticable.
- **2017 — V100 + Tensor Cores.** Unités matricielles dédiées : le débit en précision mixte (FP16) explose. Le Transformer en profite immédiatement.
- **2019 — A100 (annoncé fin 2019, dispo 2020).** 40 puis 80 Go HBM2e, sparsity, TF32 : le GPU de l'ère GPT-3/ChatGPT. Les files d'attente commencent.
- **2020 — TPU v3/v4 (Google).** Google prouve qu'on peut entraîner des LLM sans NVIDIA — mais seulement en interne.
- **2022 — H100.** ~990 TFLOPS FP16, HBM3 : le GPU de la ruée 2023-2024. Pénurie mondiale, marché gris, files de 6-12 mois.
- **2022-2023 — Contrôles d'export US.** Interdiction de vendre H100/H800 à la Chine : naissance officielle des **deux écosystèmes** (NVIDIA bridé vs Ascend/SMIC). DeepSeek naît **dans** cette contrainte — son efficacité est une adaptation, pas un hasard.
- **2023 — InfiniBand / Spectrum-X.** Le réseau devient le goulot : entraîner sur 10 000 GPU exige une interconnexion à ~400 Gb/s par nœud. Le coût réseau ≈ 20-30 % d'un cluster (« à vérifier » selon topologie).
- **2024 — B200 (Blackwell).** ~2 250 TFLOPS FP16, 192 Go HBM3e ; racks NVL72 (72 GPU, ~120 kW le rack — refroidissement liquide obligatoire).
- **2024-2025 — Les hyperscalers sortent leurs chips.** TPU v6/Trillium puis Ironwood v7 (Google), Trainium/Inferentia (AWS), Maia (Microsoft) : la **désintermédiation** commence par l'inférence, moins exigeante que le training.
- **2025-2026 — Crise HBM et crise électrique.** La mémoire HBM (SK Hynix, Micron, Samsung) est en pénurie ; les datacenters reçoivent des GPU **qu'ils ne peuvent pas alimenter** (raccordements 2-4 ans). Le débat public passe du « manque de puces » au « manque de mégawatts ».
- **2026 — Rubin (NVIDIA, annoncé), HBM4.** La feuille de route continue, mais l'industrie a compris : **le prochain goulot est le réseau électrique, pas le transistor.**

### 10.2. Ce que le hardware change pour l'exploitation

| Sujet | Avant (2020) | Maintenant (2026) |
|---|---|---|
| Refroidissement | Air, ~10-15 kW/rack | Eau/immersion, 100+ kW/rack IA |
| Densité électrique | 1 baie = 1 armoire | 1 baie IA = 1 transformateur de quartier |
| Pannes | GPU qui meurt = job relancé | À 10 000 GPU, **des pannes par jour** : checkpointing obligatoire, MTBF du cluster < 24h (« à vérifier » selon l'opérateur) |
| Réseau | 10/25 GbE suffisait | 400 Gb/s IB/Ethernet, topologie rail-optimized |
| Stockage | NFS lent OK | 10-100 Go/s pour charger les checkpoints (un checkpoint 405B FP16 = ~810 Go) |
| Sécurité | Vol de données | **Vol de poids** (des centaines de Mo valent des centaines de M$) + attaques sur la supply chain CUDA |

**Lien avec le métier de Zelef (systèmes & énergies) :** un datacenter IA est une **charge électrique critique non linéaire** — les GPU passent de 100 W à 700 W+ en millisecondes (transitoires), ce qui torture les onduleurs et les groupes. Le dimensionnement électrique d'une baie IA suit les mêmes méthodes que celles de son guide onduleurs : foisonnement, redondance N+1, tests en charge. **L'IA a fait de l'énergéticien un acteur du SI.**

---

## 11. 15 idées reçues sur l'IA, passées au crible

> Chaque idée reçue : le mythe, la réalité vérifiée, la nuance.

