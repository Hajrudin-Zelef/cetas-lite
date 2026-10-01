---
id: collect-261001-ia-llm/ia-llm/ia-grand-dossier-24
title: "IA — Le grand dossier"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "DeepSeek", "Google", "Meta", "Mistral", "United States"]
dates: []
keywords: ["agent", "agents", "agi", "attention", "benchmark", "benchmarks", "datacenter", "deepseek", "diffusion", "dpo", "embedding", "embeddings"]
source: docs/RAG/collect-261001-ia-llm/ia_grand_dossier.md
source_anchor: ""
source_lines: [1585, 1672]
sha256: 7e00e4daab0bf1f5c8a78bd09886e9cedbf40e763ee43f940804ef6c2b7758e4
---

# IA — Le grand dossier

| Acronyme | Signification | En bref |
|---|---|---|
| AGI | Artificial General Intelligence | IA de niveau humain général (objectif affiché des labos) |
| ASI | Artificial Superintelligence | IA dépassant l'humain partout (horizon débattu) |
| RLHF | Reinforcement Learning from Human Feedback | Alignement par préférences humaines (InstructGPT 2022) |
| RLVR | Reinforcement Learning with Verifiable Rewards | RL avec récompenses vérifiables (2025+) |
| SFT | Supervised Fine-Tuning | Affinage supervisé (étape 1 du RLHF) |
| DPO | Direct Preference Optimization | Alignement sans RL (2023) |
| PPO | Proximal Policy Optimization | Algo RL du RLHF classique (Schulman 2017) |
| CoT | Chain-of-Thought | Raisonnement par étapes explicites |
| MoE | Mixture of Experts | N'active que quelques experts par token |
| MLA | Multi-head Latent Attention | Attention à KV-cache compressé (DeepSeek) |
| GQA | Grouped-Query Attention | Variante d'attention économe (Llama 2+) |
| RAG | Retrieval-Augmented Generation | Réponses ancrées dans un corpus |
| LoRA | Low-Rank Adaptation | Fine-tuning à ~0,1 % des paramètres |
| QLoRA | Quantized LoRA | LoRA + quantification (fine-tune sur 1 GPU) |
| KV-cache | Key-Value cache | Mémoire des tokens déjà traités (le dimensionnant VRAM) |
| TTFT | Time To First Token | Latence avant le 1er token (UX) |
| TPS | Tokens Per Second | Débit de génération |
| MFU | Model FLOPs Utilization | % du débit théorique réellement atteint |
| FLOP | Floating-point Operation | Unité de calcul |
| FP16 / BF16 | Formats 16 bits | Précision standard d'entraînement/inférence |
| INT8 / INT4 | Quantification entière | Formats économes du self-hosting |
| TPU | Tensor Processing Unit | Chip IA de Google |
| HBM | High Bandwidth Memory | Mémoire des GPU IA (le goulot 2024-2026) |
| DDPM | Denoising Diffusion Probabilistic Models | La diffusion moderne (Ho 2020) |
| GAN | Generative Adversarial Network | Générateur vs discriminateur (Goodfellow 2014) |
| VAE | Variational Autoencoder | Ancêtre des modèles latents (DALL-E 1) |
| SSM | State Space Model | Alternative linéaire au Transformer (Mamba) |
| JEPA | Joint Embedding Predictive Architecture | La voie « world models » de LeCun |
| DPA | Data Processing Agreement | Contrat de traitement des données (RGPD) |
| PUE | Power Usage Effectiveness | Efficacité énergétique d'un datacenter |
| MTBF | Mean Time Between Failures | À 10 000 GPU, le cluster tombe en panne < 24h |
| MMLU | Massive Multitask Language Understanding | Benchmark QCM 57 matières |
| GSM8K | Grade School Math 8K | Maths niveau primaire (saturé) |
| AIME | American Invitational Mathematics Examination | Olympiades US (thermomètre du raisonnement) |
| GPQA | Graduate-level Google-Proof Q&A | Questions niveau PhD |
| SWE-bench | Software Engineering benchmark | Vrais bugs GitHub à corriger |
| ARC-AGI | Abstraction and Reasoning Corpus | Raisonnement abstrait anti-mémorisation (Chollet) |
| IFEval | Instruction-Following Eval | Suivi strict des instructions |
| RSP | Responsible Scaling Policy | Politique de sûreté d'Anthropic |

---

## 6.9. Feuille de route RAG : recommandations spécifiques au projet de Zelef

> Zelef construit son RAG personnel : embeddings `text-embedding-3-small`, corpus scrapé/nettoyé par ses soins (méthode validée : 5 étapes, voir mémoire), guides longs en français. Voici ce que la partie 1 du dossier implique pour **son** projet.

**R1. Ne pas entraîner, servir (validé par la section 4).**
Le training est l'ère révolue ; le RAG vit à l'ère de l'inférence. Budget : 0 € de training, 100 % sur la **qualité du corpus** et le **serving**.

**R2. Le corpus est l'actif, pas le modèle (validé par 4.2 et section 16).**
Le mur des données concerne le pré-entraînement général — pas un corpus métier curé. Chaque guide nettoyé (2500+ lignes) est une **donnée de haute qualité** que les labos n'ont pas. C'est l'avantage compétitif du projet.

**R3. Générateur : petit et local d'abord (validé par 4.4 et 14.3).**
Un 8B-32B quantifié (Qwen, Llama, Mistral) suffit si le retrieval est bon. Protocole : gold set de 100 questions métier → A/B 8B vs 32B vs API → décider sur chiffres. Le français du corpus n'est pas un problème : les modèles multilingues actuels gèrent très bien le français technique (point d'attention noté par Zelef le 26/09 — à valider empiriquement sur le gold set).

**R4. Multimodalité : prévoir dès maintenant (validé par 1.16).**
Les guides contiennent schémas, tableaux, photos : indexer aussi des **légendes d'images** générées par un modèle vision, et stocker les images sources pour restitution. Un RAG texte-seul sera legacy.

**R5. Évaluation continue (validé par section 15).**
Rejouer le gold set à chaque changement (modèle, chunking, prompt). C'est la seule façon de savoir si « c'est mieux » — les benchmarks publics ne disent rien de VOS documents.

**R6. Sécurité : le corpus est une surface d'attaque (validé par 6.8).**
Documents scrapés du web → risque d'injection indirecte. Règle : le RAG **informe**, il ne **décide** pas seul ; toute action (surtout via futurs agents) validée par un humain.

**R7. Coûts : modéliser avant de scaler (validé par 6.3 et 14.1-14.2).**
À usage personnel/équipe, le self-hosting local est imbattable (~100 $/mois tout compris pour un 8B). Si usage API : budgets par cas d'usage, routage petit/gros modèle, et **jamais** de boucle agent sans plafond.

**R8. Veille : suivre les bons signaux (validé par 1.15, pattern 2).**
Les capacités « fermées » d'aujourd'hui sont « ouvertes » dans 12-24 mois : ne pas s'enfermer sur une API propriétaire pour le cœur du système. Suivre : les rapports techniques (section 17), les sorties Qwen/Mistral/DeepSeek (les plus utiles en ouvert), et les benchmarks de **retrieval** (pas seulement de génération).

---

## 6.10. Le datacenter IA vu par un énergéticien : ce que le rack GPU change au tableau électrique

> Zelef est chef de service systèmes & énergies : cette section traduit l'ère de l'inférence (6.1) en **kW, en ampères et en PUE**. C'est ici que l'IA rencontre son métier.

**6.10.1. Le changement d'échelle : du rack informatique au rack industriel.**

| Génération de rack | Densité typique | Refroidissement | Alimentation |
|---|---|---|---|
| Serveurs CPU classiques | 5-15 kW/rack | Air (clim de salle) | 2× 16-32 A mono/tri |
| Serveur GPU 8× (2022-2024) | 30-60 kW/rack | Air forcé renforcé ou hybride | Tri 63-125 A |
| Rack NVL72 / équivalent (2024-2026) | **~120 kW/rack** | **Liquide direct-to-chip obligatoire** | Barres DC / tri 400 A+ |

Un rack IA 2026 consomme comme **30 à 40 foyers**. La salle serveur d'entreprise « 20 kW au total » ne peut pas accueillir un nœud 8×H100 sans travaux : c'est un **projet électrique**, pas un achat informatique.

