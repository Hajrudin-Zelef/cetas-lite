---
id: collect-261001-ia-llm/ia-llm/ia-grand-dossier-63
title: "IA — Le grand dossier"
domain: ia-llm
role: reference
task: reference
actors: ["AWS", "Anthropic", "Baseten", "Google", "Nvidia", "OpenAI", "United States", "vLLM", "xAI"]
dates: []
keywords: ["agent", "agi", "alignment", "attention", "awq", "aws", "benchmark", "benchmarks", "chatgpt", "claude", "compute", "decode"]
source: docs/RAG/collect-261001-ia-llm/ia_grand_dossier.md
source_anchor: ""
source_lines: [5345, 5392]
sha256: 1a686fc305257995764b8d173b8016adf33a38da634c585985bba74db1c24df8
---

# IA — Le grand dossier

1. **Agent (IA)** — Système qui combine un LLM avec des outils (recherche web, code, API), une mémoire et une boucle planifier-agir-observer pour accomplir des tâches multi-étapes de façon autonome.
2. **AGI (Artificial General Intelligence)** — Intelligence artificielle générale : système capable d'accomplir toute tâche cognitive humaine. Objectif affiché d'OpenAI, Anthropic, DeepMind, xAI. Aucun consensus sur sa date (les prédictions vont de 2026 à « jamais » selon les camps — voir 2.5).
3. **AI Act** — Règlement européen sur l'IA (application progressive 2024-2026) : obligations par niveaux de risque, première régulation complète au monde.
4. **Alignement (alignment)** — Faire en sorte que le modèle poursuive les intentions de l'utilisateur et respecte des garde-fous (sécurité, honnêteté). Techniques : RLHF, Constitutional AI, DPO.
5. **Alucination… non : hallucination** — voir n° 20. (Rappel volontaire : le terme le plus galvaudé du domaine.)
6. **API (modèle)** — Accès distant payant au modèle (OpenAI, Anthropic, Google...) : zéro infra, facturation au token, données chez un tiers.
7. **Attention (mécanisme)** — Opération qui pondère l'importance relative des éléments d'une séquence entre eux. Formule : softmax(QK^T/√d_k)V. Cœur du Transformer (2017).
8. **Auto-régressif** — Génération token par token, chacun conditionné par les précédents. Tous les LLM actuels fonctionnent ainsi.
9. **AWQ** — Format de quantification 4-bit « activation-aware » ; le standard pour servir vite sur vLLM/NVIDIA.
10. **Backpropagation (rétropropagation)** — Algorithme (1986) qui calcule comment ajuster chaque paramètre pour réduire l'erreur, en propageant le gradient à rebours.
11. **Batch API** — Mode différé des providers (résultat en heures) à -50 % du prix — idéal pour l'indexation RAG.
12. **Benchmark** — Jeu de tests standardisé (MMLU, GSM8K, HumanEval, AIME, GPQA...) servant à comparer les modèles. À manier avec prudence : contamination possible (le test a fuité dans l'entraînement) et « benchmark hacking ».
13. **Bitter Lesson (la leçon amère)** — Essai de Rich Sutton (2019) : à long terme, seules les méthodes qui exploitent le calcul brut (apprentissage + recherche/search) gagnent ; les connaissances humaines codées en dur finissent toujours par être dépassées.
14. **BM25** — Algorithme classique de recherche lexicale (mots-clés, fréquence). Complément indispensable du vectoriel pour les termes techniques exacts.
15. **Chain-of-Thought (CoT)** — Technique (prompting puis entraînement) où le modèle explicite des étapes intermédiaires de raisonnement avant la réponse finale. Papier fondateur : Wei et al., 2022.
16. **Chat template** — Mise en forme imposée du prompt (rôles system/user/assistant) ; chaque famille a la sienne.
17. **Checkpoint** — Sauvegarde des poids pendant l'entraînement (un checkpoint 405B FP16 ≈ 810 Go). Indispensable car les clusters tombent en panne.
18. **Chinchilla (lois)** — Voir 4.1 : N et D doivent scaler à parts égales ; règle pratique D ≈ 20×N tokens par paramètre (Hoffmann et al., DeepMind, 2022).
19. **Chunk / chunking** — Découpage des documents en morceaux avant embedding. Paramètres clés : taille (500-1000 tokens), chevauchement (10-20 %), respect de la structure.
20. **Chunking** — Découpage des documents en morceaux (chunks) pour l'indexation RAG.
21. **Code rouge** — Alerte interne Google (déc. 2022) après ChatGPT : réorganisation autour de l'IA générative.
22. **Compute** — Quantité de calcul (en FLOP) utilisée pour entraîner ou servir un modèle. Devenue l'unité de compte de l'industrie (« combien de compute ? »).
23. **Constitutional AI** — Méthode d'alignement d'Anthropic (2022) : le modèle s'auto-critique et se corrige selon une « constitution » de principes, réduisant le besoin d'annotateurs humains.
24. **Contamination** — Quand les données de test ont fuité dans l'entraînement : les scores sont gonflés. Fléau des benchmarks.
25. **Contexte (fenêtre de)** — Nombre maximal de tokens que le modèle peut traiter d'un coup (4k en 2022 → 128k en 2024 → 1-2M ensuite). Le contexte long coûte cher en VRAM (KV-cache).
26. **Continuous batching** — Technique (vLLM) qui regroupe les requêtes en continu pour maximiser le débit GPU.
27. **Cross-encoder (reranker)** — Modèle qui relit en profondeur (question, document) pour reclasser les candidats du retrieval. Coûteux mais précis — utilisé sur le top-20 → top-5.
28. **Data residency** — Lieu géographique de traitement/stockage des données ; enjeu réglementaire (UE, secteur public).
29. **Decode** — Phase de génération token par token ; limitée par la bande passante mémoire.
30. **Dedicated endpoint** — Instance GPU réservée à ton modèle (Baseten, Together…) : perf stable, coût fixe.
31. **Deepfake** — Contenu synthétique (image, audio, vidéo) imitant une personne réelle. Régulé : TAKE IT DOWN Act (US), art. 5 AI Act (UE), étiquetage obligatoire (Chine).
32. **DePIN** — Réseaux décentralisés de GPU (Akash, io.net…) : moins chers, moins garantis.
33. **Diffusion (modèle de)** — Famille de modèles génératifs (images, vidéo, audio) qui apprennent à débruiter progressivement un bruit pur. DDPM (Ho et al., 2020) ; Stable Diffusion, DALL-E 2/3, Sora, Veo.
34. **DistBelief / Distribué (entraînement)** — Répartir l'entraînement sur des milliers de GPU (data/tensor/pipeline parallelism). Le métier des clusters 2020-2024.
35. **Distillation** — Entraîner un petit modèle « élève » à reproduire le comportement (et les raisonnements) d'un grand modèle « professeur ». Le levier n° 1 du déploiement économique (ex. : R1-Distill).
36. **Egress** — Données sortant du cloud, facturées au Go (0,09 $/Go chez AWS) ; piège classique des checkpoints.
37. **Embedding** — Vecteur dense représentant le sens d'un texte (ou d'une image). Base du RAG : on compare les embeddings de la question et des chunks par similarité cosinus. Modèles : BERT-like, E5, BGE, text-embedding-3-small.
38. **Emergent abilities** — Capacités qui apparaissent brutalement au-delà d'une certaine échelle sans avoir été explicitement entraînées (ex. : le few-shot de GPT-3). Concept débattu (certains y voient un artefact de mesure).
39. **Eval (évaluation)** — Mesure systématique d'un modèle sur des tâches : la discipline qui sépare le pro de l'amateur. « Sans eval, pas de déploiement. »
40. **Fallback** — Bascule automatique vers un autre (modèle, provider) en cas d'échec ; cœur des gateways.
41. **Fine-tuning (affinage)** — Réentraînement (complet ou partiel) d'un modèle pré-entraîné sur des données spécifiques à un domaine ou une tâche.
42. **FLOP** — Floating-point operation : une opération sur nombres flottants. 1 PFLOP = 10^15. L'entraînement de GPT-3 ≈ 3,14×10^23 FLOP.
43. **Fonction objectif** — Ce que l'entraînement optimise (ex. : prédire le mot suivant). Elle définit le comportement — d'où les hallucinations (M5).
44. **Foundation model (modèle de fondation)** — Grand modèle pré-entraîné, généraliste, adaptable à de nombreuses tâches en aval (GPT-4, Claude, Llama, Gemini...). Terme popularisé par Stanford HAI (2021).
45. **FP16 / BF16** — Précision 16-bit des poids « pleine qualité » ; ~2 octets par paramètre.
46. **FP8** — Précision 8-bit native des GPU récents ; -50 % mémoire pour ~0 perte (vLLM, serveurs).
47. **Frontière (modèle)** — Les 2-3 meilleurs modèles du moment sur les benchmarks durs. La frontière bouge tous les 3-6 mois.
48. **FSDP / DeepSpeed ZeRO** — Techniques de parallélisme d'entraînement : répartissent poids, gradients et états d'optimiseur sur plusieurs GPU pour entraîner des modèles géants.
