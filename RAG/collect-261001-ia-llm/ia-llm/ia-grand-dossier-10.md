---
id: collect-261001-ia-llm/ia-llm/ia-grand-dossier-10
title: "IA — Le grand dossier"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "DeepSeek", "Google", "Meta", "Microsoft", "Mistral", "Nvidia", "OpenAI", "United States", "xAI"]
dates: []
keywords: ["attention", "blackwell", "chatgpt", "claude", "compute", "copilot", "datacenter", "deepseek", "diffusion", "embeddings", "gemini", "gpu"]
source: docs/RAG/collect-261001-ia-llm/ia_grand_dossier.md
source_anchor: ""
source_lines: [570, 671]
sha256: 533856cff3249351782a8e2ec5f6ff25bd8c145dda4f127de27d96ce7df446f4
---

# IA — Le grand dossier

### 2012 — L'an zéro du deep learning
- **P/M** — Octobre : **AlexNet** (Krizhevsky, Sutskever, Hinton) gagne ImageNet : 15,3 % d'erreur top-5 vs 26,2 % (NeurIPS 2012).
- **P** — Hinton et al. : le **dropout** comme régularisation.
- **M** — Google X lab : réseau de 16 000 CPU qui « découvre » les chats dans YouTube (Le et al.) — l'apprentissage non supervisé à grande échelle.
- Contexte : le financement mondial des startups IA passe de 670 M$ (2011) vers des milliards (36 Md$ en 2020, 77 Md$ en 2021 — chiffres VentureBeat citant des études).

### 2013
- **P** — **word2vec** (Mikolov et al., Google) : les embeddings de mots efficaces — le sens devient un vecteur.
- **P** — **DQN** (Mnih et al., DeepMind) : deep reinforcement learning sur Atari.
- **E** — Yann LeCun rejoint Facebook pour créer **FAIR**.

### 2014
- **P** — **GAN** (Goodfellow et al., NeurIPS) : la génération d'images par affrontement.
- **P** — **seq2seq** (Sutskever, Vinyals, Le) : encodeur-décodeur pour la traduction — l'ancêtre des LLM.
- **M** — VGG (Oxford), GoogLeNet (Google) : la profondeur comme recette.
- **E** — Google rachète **DeepMind** (~400-650 M£/$ selon les sources).

### 2015
- **P** — **ResNet** (He et al., Microsoft Research) : connexions résiduelles, 152 couches — l'erreur ImageNet passe sous le niveau humain.
- **P** — Bahdanau et al. : l'**attention** en traduction (précurseur du Transformer).
- **E** — Décembre : fondation d'**OpenAI** (non-profit) par Altman, Musk, Sutskever, Brockman et al.
- **P** — Karpathy (thèse Stanford, sous Fei-Fei Li) ; il rejoint OpenAI.

### 2016
- **M/E** — Mars : **AlphaGo bat Lee Sedol 4-1** — choc mondial.
- **P** — « Deep Residual Learning » publié (CVPR 2016, best paper).
- **E** — NVIDIA lance la **Tesla P100** (architecture Pascal) — le GPU datacenter IA moderne commence ici.

### 2017
- **P** — 12 juin : **« Attention Is All You Need »** (Vaswani et al., Google, arXiv:1706.03762) — le Transformer.
- **M** — **AlphaGo Zero / AlphaZero** (DeepMind) : surhumain par pur auto-jeu, sans données humaines.
- **P** — **PPO** (Schulman et al., OpenAI) : l'algorithme qui servira au RLHF.
- **E** — Karpathy devient directeur IA de **Tesla** (Autopilot vision).
- **P** — « The Bitter Lesson » n'est pas encore publié (ce sera mars 2019, Rich Sutton) mais l'idée circule.

### 2018
- **P** — Juin : **GPT-1** (OpenAI, Radford et al.) : pré-entraînement génératif, 117M paramètres.
- **P** — Octobre : **BERT** (Devlin et al., Google) : 340M paramètres, écrase GLUE/SQuAD.
- **P** — **StyleGAN** (Karras et al., NVIDIA) : visages photoréalistes.
- **E** — Musk quitte le board d'OpenAI. **Prix Turing** (annoncé mars 2019) à Hinton, LeCun, Bengio pour 2018.

### 2019
- **M** — Février : **GPT-2** (1,5B) — OpenAI retient d'abord le modèle complet (« staged release » controversée).
- **P** — Mars : **« The Bitter Lesson »** (Rich Sutton) : seules l'apprentissage et la recherche (search) scalent indéfiniment.
- **E** — Mars : OpenAI passe en **capped-profit** ; Microsoft investit **1 Md$**.
- **P** — Kaplan et al. : premières **lois de scaling** (préprint janvier 2020, travaux 2019).

### 2020
- **P** — Mai : **GPT-3** (Brown et al., 175B, 3,14×10^23 FLOP, ~4,6 M$ estimés) — le few-shot.
- **M** — Juin : API GPT-3 (accès fermé, payant).
- **P** — **DDPM** (Ho, Jain, Abbeel) : la diffusion moderne pour les images.
- **M** — Décembre : **AlphaFold 2** (DeepMind) résout le repliement des protéines.
- **I** — NVIDIA **A100** (Ampère) : le GPU de l'ère GPT-3/ChatGPT.

### 2021
- **M** — Janvier : **DALL-E** (OpenAI) : texte → image par Transformer.
- **E** — Les frères Amodei + ~10 chercheurs quittent OpenAI et fondent **Anthropic**.
- **M** — **GitHub Copilot** (OpenAI Codex + Microsoft) : premier produit dev mainstream.
- **P** — « On the Dangers of Stochastic Parrots » (Bender, Gebru et al., FAccT) : le contre-discours éthique des grands modèles.

### 2022
- **P** — Janvier/mars : **InstructGPT / RLHF** (Ouyang et al., OpenAI) — 1,3B préféré à 175B.
- **P** — Mars : **Chinchilla** (Hoffmann et al., DeepMind) corrige Kaplan : D ≈ 20×N.
- **M** — Avril : **DALL-E 2** ; juillet : **Midjourney** ; août : **Stable Diffusion** (ouvert !).
- **P** — **Chain-of-Thought** (Wei et al., Google) : le raisonnement par prompting.
- **M** — **30 novembre : ChatGPT** — 100M utilisateurs en 2 mois.
- **E** — Décembre : Google déclare un « code rouge ».
- **P** — **Constitutional AI** (Anthropic, Bai et al., décembre).
- **I** — NVIDIA **H100** (Hopper) : le GPU de la ruée 2023-2024.

### 2023
- **E** — Janvier : Microsoft remet **~10 Md$** dans OpenAI.
- **M** — Février : **LLaMA** (Meta) — fuite, puis écosystème ouvert.
- **M** — Mars : **GPT-4** (multimodal, sans détails publiés) ; **Claude** (Anthropic).
- **E** — Avril : fondation de **Mistral AI** (Paris) ; **DeepMind + Google Brain = Google DeepMind**.
- **M** — Juillet : **Llama 2** (licence commerciale) ; **Claude 2** (100k contexte).
- **E** — Juillet : fondation de **xAI** (Musk).
- **M** — Septembre : **Mistral 7B** ; novembre : **Grok** (xAI).
- **E** — 1-2 nov. : **sommet de Bletchley Park** sur la sûreté de l'IA (déclaration internationale).
- **E** — 17-22 nov. : **limogeage puis réintégration de Sam Altman** ; départs, fusion Anthropic envisagée.
- **M** — Décembre : **Gemini 1.0** (Google) ; **Mixtral 8x7B** (Mistral, MoE).
- **I** — Pénurie de H100 : files d'attente de plusieurs mois, marché gris.

### 2024
- **M** — Février : **Gemini 1.5** (contexte 1M tokens) ; **Sora** (OpenAI, vidéo).
- **M** — Mai : **GPT-4o** (omni, multimodal natif, ~320 ms de latence).
- **E** — Mars : Mustafa Suleyman rejoint **Microsoft** (CEO Microsoft AI).
- **M** — Juillet : **Llama 3.1 405B** (15 000B tokens — le sur-entraînement assumé).
- **M** — **12 septembre : o1-preview** (« Strawberry ») — le test-time compute devient un produit.
- **E** — Octobre : **Nobel de physique à Hinton & Hopfield** ; **Nobel de chimie à Hassabis, Jumper & Baker** ; essai « Machines of Loving Grace » (Amodei).
- **M** — Décembre : **o1** complet ; **DeepSeek-V3** (671B MoE, 5,576 M$ le run final) ; **Gemini 2.0** (« ère agentique »).
- **I** — NVIDIA **Blackwell B200** ; TPU v6/Trillium chez Google.

### 2025
- **M** — **20 janvier : DeepSeek-R1** (poids MIT, niveau o1) — choc mondial.
- **E** — **27 janvier : NVIDIA -17 %, ~593 Md$** de capitalisation évaporés en une séance (record US à l'époque).
- **M** — o3 (96,7 % AIME 2024), **o4-mini** (avril), **GPT-4.5 « Orion »** (mars — le symbole du mur), **GPT-5** (août — « intelligence unifiée », lancement mitigé).
- **M** — Claude raisonnant, Gemini « thinking », **Veo 3** (vidéo + audio).
- **E** — Janvier : annonce **Stargate** (~500 Md$, « à vérifier » la réalisation) ; contrôles d'export renforcés ; crise énergie/datacenters (« physicality crisis »).
- **P** — Snell et al. formalisent le **scaling du test-time compute** ; études sur l'« overthinking » ; RLVR.
- **I** — La part du compute en **inférence** dépasse celle du training dans les projections fin d'année (débat « 80/20 reversal »).

