---
id: collect-261001-ia-llm/ia-llm/ia-grand-dossier-18
title: "IA — Le grand dossier"
domain: ia-llm
role: reference
task: reference
actors: ["DeepSeek", "Falcon", "Google", "Hugging Face", "Meta", "Microsoft", "OpenAI", "TII", "United States"]
dates: []
keywords: ["agent", "agents", "agi", "attention", "benchmark", "benchmarks", "deepseek", "diffusion", "dpo", "fine-tuning", "leaderboard", "llama"]
source: docs/RAG/collect-261001-ia-llm/ia_grand_dossier.md
source_anchor: ""
source_lines: [1154, 1234]
sha256: 3f548d1fc4047f115f8227dd60f7ce498365db6abd92606d02b1ee7a6f33c43b
---

# IA — Le grand dossier

**2021 — LoRA (Hu et al., Microsoft).** Fine-tuning par matrices de rang faible : on n'entraîne que ~0,1 % des paramètres. **C'est LoRA qui a rendu le fine-tuning accessible aux entreprises** (une RTX 4090 suffit pour adapter un 7B à un domaine).

**2022 — Chinchilla (déjà couvert en 4.1)** + **GATO (DeepMind)** : un seul Transformer pour jouer, parler, contrôler un robot — la démonstration « généraliste » qui préfigure les agents.

**2023 — FlashAttention (Dao et al., Stanford).** Réécriture IO-aware de l'attention : 2-4× plus rapide, mémoire en O(N) au lieu de O(N²). **Sans FlashAttention, les contextes 128k+ seraient impraticables.** Exemple parfait d'innovation « invisible » qui débloque une capacité produit.

**2023 — DPO (Rafailov et al.).** Alignement par préférences **sans RL** : plus simple et stable que le PPO, massivement adopté en open source.

**2024 — Mamba (Gu & Dao).** Architecture à **espaces d'états sélectifs** (SSM) : complexité linéaire en séquence au lieu de quadratique — le premier concurrent crédible au Transformer pour le contexte très long. En 2025-2026, les hybrides Transformer+Mamba/SSM se multiplient (« à vérifier » les parts d'adoption).

**2024-2025 — MLA, DeepSeekMoE, Multi-Token Prediction (DeepSeek).** Voir 3.2 : la boîte à outils de l'efficacité — attention latente (KV-cache compressé), experts fins partagés, prédiction de plusieurs tokens d'un coup (×1,5-2 de débit).

**Le motif commun :** chaque « chaînon » supprime un goulot (mémoire, parallélisme, coût d'adaptation, complexité quadratique) et **déplace la frontière du possible sans changer le paradigme**. C'est exactement ce que le camp « continuation » appelle « le scaling par l'efficacité ».

---

## 15. Les 20 benchmarks qui structurent le débat (et leurs limites)

> Un benchmark = un examen standardisé. Indispensable pour comparer, dangereux si pris pour la vérité.

| Benchmark | Ce qu'il mesure | Repères (ordres, « à vérifier » au jour près) | Limite connue |
|---|---|---|---|
| **MMLU** | 57 matières, QCM niveau licence | GPT-3 : ~44 % ; GPT-4 : ~86 % ; 2025 : > 90 % (saturé) | QCM ≠ usage réel ; contamination probable |
| **MMLU-Pro** | Version durcie (10 choix, raisonnement) | Frontière 2025-2026 : ~85-90 % | Encore jeune |
| **GSM8K** | 8 500 problèmes de maths école primaire | o1/R1 : > 95 % (saturé) | Trop facile pour la frontière |
| **MATH** | 12 500 problèmes de maths concours | GPT-4 : ~53 % ; o1 : ~94 % | — |
| **AIME 2024** | 30 problèmes d'olympiades US | GPT-4o : 13 % ; o1 : 83 % ; o3 : 96,7 % ; R1 : 79,8 % | Le thermomètre du raisonnement 2024-2025 |
| **GPQA Diamond** | 448 questions de science niveau PhD | o1 : 78 % (> moyenne des PhD humains du domaine) | Petit échantillon ; variance |
| **HumanEval** | 164 problèmes de code Python | GPT-4 : ~67 % ; 2025 : > 95 % (saturé) | Saturé ; remplacé par SWE-bench |
| **SWE-bench** | Vrais bugs GitHub à corriger | 2024 : ~30 % ; agents 2025-2026 : 60-70 % (« à vérifier ») | Coûteux à évaluer |
| **ARC-AGI** | Raisonnement abstrait visuel (Chollet) | Longtemps < 10 % ; o3 (2024) : percée ~75-87 % (« à vérifier » le protocole) | Le test « anti-mémorisation » par design |
| **Codeforces** | Programmation compétitive (Elo) | o1 : 89e percentile ; o3 : 99e | — |
| **HELM** | Méta-benchmark Stanford (multi-tâches) | Référence académique | Lourd |
| **Chatbot Arena (LMArena)** | Votes humains en aveugle (Elo) | Le plus corrélé à la préférence réelle | Biais de style (les réponses longues plaisent) |
| **IFEval** | Suivi d'instructions strictes | — | — |
| **Needle in a Haystack** | Retrouver une info dans un long contexte | Teste le contexte 128k-1M | Ne teste pas le raisonnement long |
| **RULER** | Suite de tests de contexte long | — | — |
| **MMMU** | Multimodalité (image+texte) niveau expert | — | — |
| **SWE-Lancer** | Tâches freelance ($ réels) | 2025 : les agents gagnent des milliers de $ simulés | Protocole débattu |
| **Terminal-Bench** | Tâches terminal/système autonomes | — | **Pertinent pour les sysadmins** |
| **GAIA** | Assistant général multi-outils | — | Référence « agent » |
| **EQ-Bench / AlignBench** | Alignement, refus calibrés | — | Subjectif par construction |

**Règles d'hygiène benchmark (à appliquer dans le RAG de Zelef comme ailleurs) :**
1. **Saturation** : un benchmark > 95 % ne discrimine plus — il faut passer au suivant (cycle MMLU → MMLU-Pro → ...).
2. **Contamination** : si le test a fuité dans le training, le score est gonflé. Préférer les benchmarks **privés/dynamiques** pour l'évaluation interne.
3. **Gaming** : optimiser pour le benchmark ≠ améliorer le produit (loi de Goodhart).
4. **Le meilleur benchmark, c'est le vôtre** : 100 questions métier avec bonnes réponses (voir 14.3) > n'importe quel leaderboard public pour décider d'un déploiement.

---

## 16. Les datasets qui ont fait l'IA (et le mur qu'ils annoncent)

> Pas de données, pas d'IA. Cette annexe recense les corpus fondateurs — et montre **pourquoi le mur des données (4.2) est structurel**.

| Dataset | Année | Contenu (ordres) | Rôle historique |
|---|---|---|---|
| **ImageNet** | 2009 (Li et al.) | ~15M images, 22k catégories | Rend AlexNet possible (2012) |
| **Common Crawl** | 2007→ (mensuel) | Pétaoctets de web brut | La matière première de (presque) tous les LLM |
| **C4** | 2019 (Raffel et al.) | ~750 Go de web nettoyé | Base de T5 ; standard académique |
| **The Pile** | 2020 (EleutherAI) | ~800 Go, 22 sources | Le corpus ouvert de référence (GPT-Neo/J) |
| **RefinedWeb** | 2023 (Falcon/TII) | ~5 000 Md tokens | Montre que le web filtré à grande échelle suffit (Falcon 40B) |
| **Dolma** | 2024 (AI2) | ~3 000 Md tokens | Corpus ouvert documenté (OLMo) |
| **LAION-5B** | 2022 | 5,85 Md paires image-texte | Rend Stable Diffusion possible |
| **The Stack** | 2022 (BigCode) | ~6 To de code | StarCoder, puis tous les modèles code |
| **RedPajama** | 2023 | ~1 200 Md tokens | Réplique ouverte du corpus LLaMA |
| **FineWeb** | 2024 (Hugging Face) | ~15 000 Md tokens | Le plus grand corpus ouvert nettoyé |
| **BookCorpus / PG-19** | 2015/2020 | Livres | La « qualité » des premiers LLM |

**Pourquoi le mur est structurel :**
1. **Le web n'est pas infini en qualité.** Epoch AI estime le stock de texte public filtrable à ~300 000 Md tokens ; les runs frontières consomment déjà des dizaines de milliers de Md tokens chacun. À ce rythme : **épuisement 2026-2032** (fourchette Epoch AI, mise à jour périodiquement).
2. **Le web se ferme.** Data Provenance Initiative (MIT, 2024) : 5 % des données et **25 % des sources de haute qualité restreintes en un an** (robots.txt, CGU, API payantes, procès — NYT vs Microsoft/OpenAI, 2023→).
3. **Le web se pollue.** Une part croissante du texte en ligne est **générée par des IA** — réentraîner dessus sans précaution cause le **model collapse** (Shumailov et al., 2023 : dégradation démontrée en l'absence de données fraîches).
4. **Les échappatoires** (voir 4.3) : données synthétiques **filtrées par vérificateurs** (pas du brut), RL sans données humaines (R1-Zero, AlphaGo Zero), **multimodal** (vidéo, capteurs, robotique : des ordres de grandeur au-delà du texte).

**Leçon pour le RAG d'entreprise :** le mur des données concerne le **pré-entraînement général**. Pour un RAG métier, le problème inverse existe : **vos données internes sont rares et précieuses** — c'est votre avantage compétitif. Les nettoyer, les structurer, les évaluer (la méthode de Zelef) est le vrai « training » d'une entreprise en 2026.

---

## 3.4. Les acteurs secondaires mais importants (ne pas les ignorer)

