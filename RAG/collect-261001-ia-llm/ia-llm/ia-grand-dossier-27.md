---
id: collect-261001-ia-llm/ia-llm/ia-grand-dossier-27
title: "IA — Le grand dossier"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "Cohere", "DeepSeek", "Google", "Meta", "Mistral", "Nvidia", "OpenAI", "xAI"]
dates: ["2022-11-30", "2025-01-20"]
keywords: ["apache", "attention", "benchmark", "capex", "chatgpt", "claude", "cohere", "compute", "deepseek", "diffusion", "distillation", "embedding"]
source: docs/RAG/collect-261001-ia-llm/ia_grand_dossier.md
source_anchor: ""
source_lines: [1807, 1889]
sha256: 63803b62bc3f013457da8ae1e85793096782c8fb79bf483d191eacc044539c26
---

# IA — Le grand dossier

> Synthèse des échecs les plus fréquents observés 2023-2026. À relire avant chaque mise en production.

**24.1. Chunker au hasard.** Découper tous les 500 caractères sans respecter la structure (titres, tableaux, procédures) détruit le sens. **Parade :** chunking **sémantique/structurel** (par section, avec recouvrement), taille adaptée au domaine ; tableaux et procédures = chunks dédiés.

**24.2. Un seul embedding pour tout.** Le modèle d'embedding générique rate le vocabulaire métier (« VRP », « U-codes », références produits). **Parade :** A/B tester 2-3 embeddings sur un gold set ; envisager un **fine-tuning d'embedding** sur paires question/réponse métier si le volume le justifie.

**24.3. Pas de reranker.** Le top-k vectoriel brut ramène du bruit ; le générateur s'y noie. **Parade :** retrieve large (top-50) → **rerank** (Cohere Rerank, bge-reranker, ou cross-encoder local) → top-5 au générateur. Le reranker est souvent le meilleur ROI du pipeline.

**24.4. Évaluer au feeling.** « Ça a l'air bon » n'est pas une métrique. **Parade :** gold set de 50-100 questions avec réponses attendues (méthode 14.3) ; rejouer à chaque changement ; suivre faithfulness (fidélité aux sources) et answer relevancy.

**24.5. Laisser le modèle répondre hors corpus.** Sans consigne stricte, le générateur complète avec ses connaissances — hallucinations « propres ». **Parade :** prompt système « réponds UNIQUEMENT d'après les passages fournis ; sinon dis "je ne sais pas" » + citation obligatoire des sources.

**24.6. Ignorer les images et tableaux.** Dans la doc technique, l'info critique est souvent dans un schéma ou un tableau (tensions, codes d'erreur). **Parade :** extraction dédiée (pdftotext -layout, OCR), légendes générées par modèle vision, indexation multimodale (voir 1.16).

**24.7. Pas de gestion des versions.** Le corpus change (nouveaux firmwares, docs révisées) mais l'index ne suit pas → réponses obsolètes. **Parade :** pipeline d'ingestion **versionné** (doc → version → chunks), réindexation incrémentale, date de validité affichée dans les réponses.

**24.8. Sécurité : le corpus comme cheval de Troie.** Documents scrapés = risque d'injection indirecte (voir 6.8). **Parade :** le RAG **informe**, ne **décide** pas ; aucune action critique sans validation humaine ; filtrage des documents ingérés.

**24.9. Scaler avant de mesurer.** Passer en production multi-utilisateurs sans connaître le coût/token ni la latence p95. **Parade :** modéliser (voir section 9, formule corrigée), budgets par cas d'usage, cache des requêtes fréquentes, routage petit/gros modèle.

**24.10. Choisir le modèle avant le corpus.** 80 % de la qualité d'un RAG vient des données et du retrieval, 20 % du générateur. **Parade :** investir d'abord dans le nettoyage du corpus (la méthode de Zelef : 5 étapes validées), puis tester les générateurs par A/B — dans cet ordre.

**Checklist de mise en production (à cocher) :**
- [ ] Gold set ≥ 50 questions, rejoué automatiquement
- [ ] Reranker en place, top-k calibré
- [ ] Prompt « sources uniquement » + citations
- [ ] Images/tableaux traités, pas ignorés
- [ ] Ingestion versionnée, réindexation planifiée
- [ ] Budgets tokens + alertes FinOps
- [ ] Revue sécurité (injection, exfiltration, secrets)
- [ ] Plan de rollback (version précédente de l'index)

---

## Fiche de révision express : la partie 1 en 40 points

> À relire avant une réunion, un entretien ou pour tester sa mémoire. Chaque point = une section du dossier.

1. 2012 : AlexNet, 15,3 % d'erreur top-5 (contre 26,2 %) — le choc qui lance le deep learning (1.2).
2. Les 3 ingrédients de 2012 : algorithmes (années 80-90), données (ImageNet 2009), compute (GPU CUDA) (1.1).
3. 2016 : AlphaGo bat Lee Sedol — le RL + search prouve sa force (1.3).
4. 2017 : « Attention Is All You Need » — 8 auteurs, suppression de la récurrence, parallélisme total (1.5).
5. 2018 : GPT-1 (génératif) et BERT (bidirectionnel) — le pré-entraînement devient la norme (1.6).
6. 2020 : GPT-3, 175B paramètres, few-shot ; lois de Kaplan : scaler les paramètres (1.7, 4.1).
7. 2020 : DDPM — la diffusion supplante les GAN en image (1.8).
8. 2022 : InstructGPT/RLHF (SFT → reward model → PPO) ; 1,3B aligné préféré à 175B brut (1.9).
9. 30/11/2022 : ChatGPT, 100M utilisateurs en 2 mois (1.9).
10. 2022 : Chinchilla corrige Kaplan — D ≈ 20 tokens/paramètre, scaler N et D à parts égales (4.1).
11. 2023 : LLaMA (fuite puis ouverture), Stable Diffusion, l'année de l'open-weight (1.10).
12. 2024 : o1 — le test-time compute devient un produit (AIME 13 % → 83 %) (4.5).
13. 20/01/2025 : DeepSeek-R1 — niveau o1 en poids MIT ; NVIDIA -17 % en un jour (1.12).
14. 2026 : l'inférence ≈ 80 % du compute ; le training ne rapporte plus assez par dollar (4.2, 6.1).
15. Parrains : Hinton, LeCun, Bengio — Turing 2018 ; Hinton Nobel physique 2024, Hassabis Nobel chimie 2024 (2.1).
16. Sutskever : AlexNet → seq2seq → OpenAI → Safe Superintelligence (2024) (2.3).
17. Karpathy : OpenAI → Tesla → OpenAI → Eureka Labs → Anthropic (05/2026) (2.4).
18. Amodei : quitte OpenAI (2020), fonde Anthropic (2021), Constitutional AI, Claude (2.3).
19. Hassabis : DeepMind (2010) → AlphaGo → AlphaFold → Google DeepMind (2.3).
20. LeCun vs Hinton : le sceptique du scaling contre l'alerteur — le débat en deux personnes (2.5).
21. OpenAI : le labo frontière fermé, le plus financé (Stargate ~500 Md$ annoncés) (3.2).
22. Anthropic : la sûreté comme produit (Constitutional AI, RSP), le roi du code agentique (3.2).
23. Google DeepMind : le plus complet (modèles + TPU + recherche), Gemini, Veo (3.2).
24. Meta : l'ouvert stratégique (Llama), l'ads qui paie le compute (3.2).
25. DeepSeek : l'efficacité comme arme géopolitique — V3 : 671B/37B actifs, 5,576 M$ (3.2).
26. Mistral : le champion européen, Apache 2.0, 3 Md€ levés (09/2026) (3.2).
27. Camp du mur : rendements décroissants, Epoch AI (données 2026-2032), Musk, coûts ×2-3 (4.2).
28. Camp de la continuation : Gwern, Bitter Lesson (Sutton), Snell (test-time), les sauts o1/R1 (4.3).
29. MoE : capacité d'un géant, coût d'un moyen — 37B actifs sur 671B (4.4).
30. Distillation : le raisonnement o1-class dans un 7B local — le levier n° 1 du déploiement (4.4).
31. Sur-entraînement : Llama 3 8B à 1 875 tok/param — l'optimum économique ≠ l'optimum training (4.4).
32. Quantification : FP16 → INT4 = VRAM ÷ 4-8 — ce qui rend le self-hosting possible (6.2).
33. Formule C ≈ 6ND : le coût d'un run en une multiplication (9.1).
34. Coût inférence : $/h ÷ (tok/s × 3600) × 1M — mesurer VOTRE débit réel (9.5, corrigé).
35. KV-cache : batch × contexte × couches — le vrai dimensionnant VRAM du serving (6.2).
36. RAG : 80 % de la qualité = corpus + retrieval, 20 % = générateur (24, S.10).
37. Gold set : 50-100 questions métier, rejoué à chaque changement — le seul benchmark qui compte (14.3).
38. Sécurité : injection indirecte via le corpus, exfiltration RGPD, secrets dans le fine-tuning (6.8).
39. Pattern historique : percée fermée → réplique ouverte en 12-24 mois — ne pas parier sur une avance temporaire (1.15).
40. Stratégie no-regret : abstraction multi-fournisseur, données propres, OPEX > CAPEX, compétences transférables (4.9).

---

## Index de la partie 1

