---
id: collect-261001-ia-llm/ia-llm/ia-grand-dossier-3
title: "IA — Le grand dossier"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "Apple", "Google", "Hugging Face", "Lambda", "Meta", "Microsoft", "OpenAI", "Stability AI"]
dates: []
keywords: ["chatgpt", "claude", "compute", "diffusion", "embedding", "fine-tuning", "gpu", "llama", "lora", "mai", "merger", "multimodal"]
source: docs/RAG/collect-261001-ia-llm/ia_grand_dossier.md
source_anchor: ""
source_lines: [155, 213]
sha256: 53c738e19a63928fe84cee92da8353f48aec64accca78520050c500b096603b9
---

# IA — Le grand dossier

**Pour le RAG de Zelef :** BERT et ses descendants (sentence-transformers, E5, BGE) sont **les modèles d'embedding** — dont `text-embedding-3-small` qu'il utilise. Comprendre BERT, c'est comprendre pourquoi son RAG retrouve les bons passages : l'encodeur bidirectionnel produit des vecteurs de sens comparables par similarité cosinus.

Autres jalons 2018-2019 :

- **Février 2019 — GPT-2 (OpenAI).** 1,5 milliard de paramètres. OpenAI refuse d'abord de publier le modèle complet par « crainte d'usage malveillant » (décision controversée, modèle publié intégralement fin 2019). GPT-2 montre des capacités de génération de texte cohérent sur plusieurs paragraphes qui impressionnent — et inquiètent — la presse.
- **2019 — Les lois de scaling de Kaplan (OpenAI).** Voir section 4 : premier formalisme sérieux de la relation entre taille, données, compute et performance.

### 1.7. Révolution n° 3 — GPT-3 et l'émergence du « few-shot » (2020)

**Le fait.** Mai-juin 2020 : OpenAI publie le papier **« Language Models are Few-Shot Learners »** (Brown et al., NeurIPS 2020) présentant **GPT-3 : 175 milliards de paramètres**, entraîné sur ~300 milliards de tokens, pour un coût estimé à **4,6 millions de dollars** (estimation Lambda Labs : 3,14×10^23 FLOP, soit 355 années-GPU sur V100 à l'époque).

**La vraie nouveauté n'est pas la taille, c'est le comportement :** sans aucun réentraînement (sans fine-tuning), GPT-3 peut accomplir une tâche simplement en voyant **quelques exemples dans le prompt** (few-shot learning), voire **aucun** (zero-shot). Traduction, résumé, questions-réponses, génération de code : il suffit de décrire la tâche en langage naturel.

C'est le moment où le grand public technique comprend que **l'échelle fait émerger des capacités qualitativement nouvelles** — ce que la littérature appellera les « emergent abilities ». C'est aussi le fondement empirique de l'« hypothèse du scaling » (voir section 4).

**Accès :** GPT-3 n'est d'abord accessible que via API payante (juin 2020), pas en open source. Ce choix commercial — répété ensuite — structure tout le débat ouvert vs fermé (section 3).

### 1.8. Révolution n° 4 — La diffusion : l'image, puis la vidéo (2020-2023)

Pendant que le texte explose avec GPT-3, la génération d'images change de paradigme :

- **2020 — DDPM (Ho, Jain, Abbeel, UC Berkeley).** « Denoising Diffusion Probabilistic Models » : au lieu de l'affrontement des GAN, on apprend à **débruiter progressivement** une image de bruit pur. Plus stable à entraîner que les GAN, meilleurs résultats.
- **Janvier 2021 — DALL-E (OpenAI).** Génération d'images à partir de descriptions textuelles, via une approche Transformer + quantification (VQ-VAE). Le nom (Dali + WALL-E) devient iconique.
- **Avril 2022 — DALL-E 2.** Saut qualitatif majeur : images photoréalistes à partir d'une phrase.
- **Août 2022 — Stable Diffusion (Stability AI + LMU Munich, Rombach et al.).** Modèle de diffusion **ouvert** (poids publics). N'importe qui peut le faire tourner sur une carte gamer. C'est l'équivalent, pour l'image, de ce que LLaMA sera pour le texte en 2023 : la démocratisation.
- **Juillet 2022 — Midjourney** (version grand public), **2023 — Adobe Firefly** (modèle « propre » juridiquement, entraîné sur Adobe Stock).

**La vidéo suit :** 2022-2023, les premiers modèles texte→vidéo crédibles (Make-A-Video de Meta, Imagen Video de Google, Gen-2 de Runway, Pika). **Février 2024 — Sora (OpenAI)** : vidéos d'une minute d'un réalisme inédit. **2025 — Veo 3 (Google)** avec audio natif. La trajectoire est claire : **texte → image → vidéo → monde 3D/4D**, chaque média tombant l'un après l'autre sous le même paradigme (Transformer + diffusion).

**Note technique pour sysadmin :** l'inférence de diffusion est gourmande en VRAM mais parallélisable par batch ; c'est un cas d'usage typique des GPU « prosumer » (RTX 4090 24 Go) en self-hosting.

### 1.9. Révolution n° 5 — RLHF : d'un prédicteur de texte à un assistant (2022)

**Le problème.** GPT-3 sait « continuer » du texte, pas « aider » un utilisateur. Demandez-lui une question : il peut répondre par une autre question, ou divaguer. Le modèle brut n'est pas un produit.

**La solution — InstructGPT / RLHF.** OpenAI publie en janvier 2022 (papier arXiv 2203.02155, Ouyang et al.) la méthode **RLHF (Reinforcement Learning from Human Feedback)** en trois étapes :

1. **SFT (Supervised Fine-Tuning).** Des annotateurs humains écrivent des réponses idéales ; on affine le modèle dessus.
2. **Modèle de récompense.** Des humains classent plusieurs réponses du modèle par paire (« laquelle est la meilleure ? ») ; on entraîne un second modèle à prédire ces préférences.
3. **PPO (Proximal Policy Optimization).** On optimise le modèle de langage par renforcement contre le modèle de récompense, avec une pénalité KL pour ne pas trop s'éloigner du modèle initial (anti « reward hacking »).

**Résultat choc du papier :** un modèle **InstructGPT de 1,3 milliard de paramètres** (100× plus petit que GPT-3) est **préféré par les évaluateurs humains** au GPT-3 brut de 175 milliards. L'alignement bat la taille brute, à tâche d'assistance égale.

**30 novembre 2022 — ChatGPT.** OpenAI lance publiquement ChatGPT, « modèle frère d'InstructGPT », avec RLHF conversationnel. **100 millions d'utilisateurs en deux mois** — le produit grand public à la croissance la plus rapide de l'histoire à l'époque. C'est le « moment iPhone » de l'IA : la technologie existait (GPT-3, 2020), mais le produit aligné change tout.

**Conséquence industrielle :** Google déclare un « code rouge » interne (décembre 2022). Microsoft investit 10 milliards de dollars supplémentaires dans OpenAI (janvier 2023, portant le partenariat à ~13 milliards au total — chiffres annoncés publiquement). La course est lancée.

**Anthropic et la variante « Constitutional AI » (décembre 2022).** Anthropic publie sa méthode d'alignement par **IA constitutionnelle** : au lieu d'annotateurs humains pour chaque jugement, le modèle s'auto-critique selon une « constitution » (liste de principes). Moins coûteux en annotation humaine, plus scalable. C'est la marque de fabrique de Claude.

### 1.10. 2023 : la course LLM — GPT-4, l'open source contre-attaque

**Mars 2023 — GPT-4 (OpenAI).** Multimodal (accepte les images en entrée), barre d'examen du barreau américain dans le top 10 %, raisonnement nettement supérieur à GPT-3.5. OpenAI ne publie **aucun détail** sur l'architecture, la taille ou les données (« en raison du paysage concurrentiel ») — rupture avec la tradition académique, très critiquée par la recherche publique.

**Février 2023 — LLaMA (Meta).** Meta publie les poids de **LLaMA** (7B à 65B paramètres) sous licence « recherche ». Fuite sur 4chan une semaine après la publication : tout le monde y a accès. **C'est l'événement le plus structurant de l'open source IA** : pour la première fois, un modèle de niveau GPT-3 est téléchargeable et affinable par n'importe quel labo ou entreprise.

**Mars-juillet 2023 — l'explosion des dérivés.** Alpaca (Stanford, fine-tune de LLaMA 7B pour ~600 $ — « à vérifier » pour le montant exact, l'ordre de grandeur est documenté), Vicuna, Koala, puis **juillet 2023 — Llama 2** (licence commerciale ouverte) et ses versions « Chat » entraînées en RLHF (détaillé dans le papier). L'écosystème Hugging Face explose : des milliers de modèles dérivés.

**Le mémo « We have no moat » (mai 2023).** Un mémo interne Google fuité affirme que ni Google ni OpenAI n'ont de « fossé » défensif : l'open source innove plus vite (LoRA, quantization 4-bit, etc.). Thèse débattue, mais elle capte l'esprit de 2023 : **l'efficacité et l'ouverture contre la taille et le secret**.

