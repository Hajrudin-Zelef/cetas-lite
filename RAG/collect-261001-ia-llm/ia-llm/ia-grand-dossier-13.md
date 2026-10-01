---
id: collect-261001-ia-llm/ia-llm/ia-grand-dossier-13
title: "IA — Le grand dossier"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "DeepSeek", "Google", "Lambda", "Meta", "OpenAI"]
dates: []
keywords: ["attention", "chatgpt", "compute", "deepseek", "diffusion", "distillation", "dpo", "embedding", "embeddings", "fine-tuning", "fp8", "gpu"]
source: docs/RAG/collect-261001-ia-llm/ia_grand_dossier.md
source_anchor: ""
source_lines: [805, 867]
sha256: 516b609ec46b27421d7af5ce02740e4ae3d190282ec728686332b2805b16f26c
---

# IA — Le grand dossier

| # | Papier | Auteurs / année | Résumé en 2 lignes | Pourquoi le lire |
|---|---|---|---|---|
| 1 | Learning representations by back-propagating errors | Rumelhart, Hinton, Williams — *Nature*, 1986 | Popularise la rétropropagation pour les réseaux multicouches ; démontre expérimentalement son utilité. | Comprendre l'algorithme qui entraîne tout depuis 40 ans. |
| 2 | Gradient-Based Learning Applied to Document Recognition | LeCun, Bottou, Bengio, Haffner — 1998 | LeNet-5 : CNN entraîné par gradient pour lire les chèques ; le motif conv-pool-FC. | L'ancêtre de toute la vision moderne ; le pattern est toujours là. |
| 3 | A Neural Probabilistic Language Model | Bengio et al. — 2003 | Premier modèle neuronal de langage avec word embeddings appris. | L'idée « le sens = un vecteur » naît ici — base des embeddings du RAG. |
| 4 | Dropout: A Simple Way to Prevent Neural Networks from Overfitting | Srivastava, Hinton et al. — 2014 (travaux 2012) | Éteindre aléatoirement des neurones à l'entraînement réduit le surapprentissage. | Régularisation encore utilisée ; comprendre le surapprentissage. |
| 5 | ImageNet Classification with Deep CNNs (AlexNet) | Krizhevsky, Sutskever, Hinton — NeurIPS 2012 | 60M paramètres, ReLU, 2 GPU : 15,3 % top-5, divise l'erreur par ~2. | L'an zéro ; ReLU+GPU+dropout expliqués simplement. |
| 6 | Generative Adversarial Nets | Goodfellow et al. — NeurIPS 2014 | Générateur vs discriminateur : la génération d'images réalistes. | Le paradigme génératif avant la diffusion ; intuitions adversariales utiles. |
| 7 | Sequence to Sequence Learning with Neural Networks | Sutskever, Vinyals, Le — NeurIPS 2014 | Encodeur-décodeur LSTM pour la traduction ; l'ancêtre direct des LLM. | Voir d'où vient l'idée « encoder puis décoder ». |
| 8 | Deep Residual Learning for Image Recognition (ResNet) | He et al. — CVPR 2016 | Skip connections : on entraîne 152 couches ; erreur sous le niveau humain. | Les connexions résiduelles sont partout (y compris dans les Transformers). |
| 9 | Attention Is All You Need | Vaswani et al. — NeurIPS 2017 | Self-attention, plus de récurrence ; parallélisme total. | LE papier. À lire en entier une fois dans sa vie. |
| 10 | Improving Language Understanding by Generative Pre-Training (GPT-1) | Radford et al. — OpenAI, 2018 | Pré-entraînement génératif + fine-tuning ; 117M paramètres. | Le principe « pré-entraîner puis adapter » du RAG moderne. |
| 11 | BERT | Devlin et al. — Google, 2018 (NAACL 2019) | Encodeur bidirectionnel à masques ; SOTA compréhension. | Comprendre les modèles d'embedding (base du retrieval). |
| 12 | The Bitter Lesson | Rich Sutton — 2019 (essai) | Seuls l'apprentissage et la recherche scalent indéfiniment. | La philosophie du camp « continuation » en 3 pages. |
| 13 | Scaling Laws for Neural Language Models | Kaplan et al. — OpenAI, 2020 | Lois de puissance L(N), L(D), L(C) ; privilégier les paramètres. | Le cadre quantitatif qui a justifié GPT-3. |
| 14 | Language Models are Few-Shot Learners (GPT-3) | Brown et al. — NeurIPS 2020 | 175B paramètres ; capacités émergentes sans fine-tuning. | Le moment où l'échelle devient qualitative. |
| 15 | Denoising Diffusion Probabilistic Models (DDPM) | Ho, Jain, Abbeel — NeurIPS 2020 | Apprendre à débruiter : la génération d'images stable. | Base de Stable Diffusion / DALL-E 2 / Sora. |
| 16 | Training Compute-Optimal LLMs (Chinchilla) | Hoffmann et al. — DeepMind, 2022 | N et D à parts égales ; D ≈ 20×N ; GPT-3 sous-entraîné. | La correction qui a réorienté toute l'industrie. |
| 17 | Training language models to follow instructions with human feedback (InstructGPT) | Ouyang et al. — OpenAI, 2022 | RLHF en 3 étapes ; 1,3B préféré à 175B. | La recette qui a fait ChatGPT ; encore le standard. |
| 18 | Chain-of-Thought Prompting Elicits Reasoning | Wei et al. — Google, 2022 | Expliciter les étapes intermédiaires améliore le raisonnement. | L'ancêtre du test-time compute ; technique toujours utile. |
| 19 | Constitutional AI: Harmlessness from AI Feedback | Bai et al. — Anthropic, 2022 | Auto-critique selon une constitution ; moins d'annotation humaine. | L'alternative d'Anthropic au RLHF pur. |
| 20 | LLaMA: Open and Efficient Foundation Language Models | Touvron et al. — Meta, 2023 | Modèles ouverts 7B-65B compétitifs ; recette data + archi. | Le papier qui a lancé l'ère open-weight. |
| 21 | Llama 2: Open Foundation and Fine-Tuned Chat Models | Touvron et al. — Meta, 2023 | Détail du RLHF industriel (dont reward modeling à grande échelle). | Le RLHF « en vrai », avec les chiffres. |
| 22 | Direct Preference Optimization (DPO) | Rafailov et al. — 2023 | Aligner sans RL : le modèle de langage est « secrètement » un modèle de récompense. | L'alternative simple au PPO, très utilisée en open source. |
| 23 | Scaling LLM Test-Time Compute Optimally | Snell et al. — 2024 | Lois d'échelle de l'inférence : réfléchir plus = mieux, en loi de puissance. | Le fondement théorique du paradigme 2024-2026. |
| 24 | DeepSeek-V3 Technical Report | DeepSeek — déc. 2024 | MoE 671B/37B, MLA, FP8, 5,576 M$ le run final. | La transparence maximale sur un modèle frontière. |
| 25 | DeepSeek-R1 Technical Report | DeepSeek — janv. 2025 | Raisonnement par pur RL (R1-Zero) ; distillation ; niveau o1 en ouvert. | Le papier du choc de janvier 2025 ; le RL appliqué au raisonnement. |

**Méthode de lecture conseillée :** 9 → 13 → 16 → 17 → 23 → 25 = le fil rouge « scaling » en 6 papiers. Le reste en fonction des besoins (vision : 5, 8 ; génération : 6, 15 ; RAG/embeddings : 11 ; alignement : 17, 19, 22).

---

## 9. Les chiffres du compute, 2012 → 2026 : tables et formules

### 9.1. La formule fondamentale : C ≈ 6ND

Pour un Transformer, le compute d'entraînement (en FLOP) vaut approximativement :

```
C = 6 × N × D
```

- **N** = nombre de paramètres, **D** = nombre de tokens d'entraînement.
- Le facteur 6 vient de : ~2ND pour la passe avant + ~4ND pour la passe arrière.
- **Usage sysadmin :** estimer le coût d'un run. Exemple : entraîner un 8B sur 1 000B tokens → C = 6 × 8×10^9 × 10^12 = 4,8×10^22 FLOP. Sur des H100 (~10^15 FLOP/s utiles à MFU 50 %) : ~4,8×10^7 secondes-GPU ≈ 13 300 heures-GPU ≈ 554 jours-GPU → ~18 jours sur 32 GPU. À ~2-3 $/h/GPU (cloud) : **~27-40 k$** pour le run seul (hors essais). C'est le calcul à faire avant tout projet de (pré-)entraînement.

### 9.2. Table d'évolution du compute des modèles repères

| Modèle | Année | Paramètres | Tokens (ordre) | Compute (FLOP) | Coût run (estim.) |
|---|---|---|---|---|---|
| AlexNet | 2012 | 60M | 1,2M images | ~10^18 (« à vérifier ») | ~quelques k$ |
| Transformer originel | 2017 | 65M | ~36M phrases | ~10^19 | ~930 $ |
| GPT-1 | 2018 | 117M | ~1B | ~10^20 (« à vérifier ») | ~quelques k$ |
| BERT-Large | 2018 | 340M | ~3,3B mots | ~10^20-10^21 (« à vérifier ») | ~7 k$ (TPU, ordre public) |
| GPT-2 | 2019 | 1,5B | ~10B | ~10^21 (« à vérifier ») | ~43 k$ (ordre public) |
| GPT-3 | 2020 | 175B | 300B | 3,14×10^23 | ~4,6 M$ (Lambda Labs) |
| Chinchilla | 2022 | 70B | 1 400B | ~5,9×10^23 (« à vérifier ») | ~2-5 M$ (ordres publics) |
| Llama 2 70B | 2023 | 70B | 2 000B | ~8,4×10^23 (« à vérifier ») | « à vérifier » |
| GPT-4 | 2023 | non publié | non publié | non publié | > 100 M$ (Altman) |
| Llama 3 405B | 2024 | 405B | 15 000B | ~3,6×10^25 (« à vérifier ») | ~100-200 M$ (analystes — « à vérifier ») |
| DeepSeek-V3 | 2024 | 671B (37B actifs) | 14 800B | ~2,8×10^24 (« à vérifier ») | 5,576 M$ (run final, rapport) |
| Génération GPT-5 | 2025 | non publié | non publié | non publié | ~500 M$ (analystes — « à vérifier ») |

