---
id: collect-261001-ia-llm/ia-llm/ia-grand-dossier-40
title: "IA — Le grand dossier"
domain: ia-llm
role: reference
task: reference
actors: ["AMD", "AWS", "Alibaba", "Anthropic", "Apple", "Cerebras", "DeepSeek", "Fireworks AI", "Google", "Groq", "MiniMax", "Mistral", "Moonshot", "Nvidia", "OpenAI", "OpenRouter", "Z.ai", "vLLM"]
dates: []
keywords: ["amd", "apache", "awq", "benchmarks", "datacenter", "decode", "deepseek", "embeddings", "fine-tuning", "gguf", "glm", "gpu"]
source: docs/RAG/collect-261001-ia-llm/ia_grand_dossier.md
source_anchor: ""
source_lines: [3033, 3152]
sha256: ca52b1ced013f0f005a58e2c248b7267cc0cc344e7a2dd9df227e8ad2e711885
---

# IA — Le grand dossier

- Fait : 70B Q4 entièrement en mémoire unifiée ; 100B+ Q4 en 128 Go ; <110 W, inaudible.
- Débit : 8B ~70–95 tok/s ; 70B ~12–28 tok/s (bande passante 546–1 200 Go/s).
- Limite : pas de CUDA (pas de vLLM classique, pas de fine-tuning PyTorch standard) ; écosystème MLX.
- Idéal pour : bureau, RAG d'équipe, poste du DSI.

### I.5. « Le cloud qui remplace l'achat » — RunPod/Vast.ai à l'heure

- 4090 à ~0,14–0,74 $/h (spot/secure) ; H100 à ~1,47–3,29 $/h.
- Fait : tester avant d'acheter ; pics temporaires ; fine-tuning pas cher (spot + checkpoints).
- Règle : si le besoin dépasse 300 h/mois récurrentes → acheter devient moins cher.

---

## Annexe J — Licences open-weight : le mémo juridique pratique

**Principe** : la licence suit le **modèle**, pas le format (un Qwen3 reste Apache 2.0 en GGUF). Vérifier sur le dépôt **d'origine**.

| Situation | Apache 2.0 / MIT | MIT modifiée / custom | Llama Community | Propriétaire |
|---|---|---|---|---|
| Usage interne entreprise | ✅ libre | ✅ (lire les seuils) | ✅ | selon contrat |
| Produit commercial qui embarque le modèle | ✅ libre | ⚠️ vérifier seuils (revenus/MAU/territoire) | ⚠️ clause 700M MAU + usage acceptable | ❌ sauf licence |
| Fine-tuner et redistribuer | ✅ (mention de licence) | ⚠️ selon licence | ⚠️ conditions Llama | ❌ |
| Entraîner un concurrent | ✅ | ⚠️ certains l'interdisent | ❌ au-delà des seuils | ❌ |

**Cas concrets (sept. 2026)** : Qwen3/3.5, GPT-OSS, Phi, Gemma 4 → ✅ sans réserve. DeepSeek V4/V3.2, GLM-4.5/4.7/5.2, Kimi K2.6, MiniMax M2.7, MiMo → ✅ MIT mais **revérifier modèle par modèle**. Mistral Medium 3.5, Kimi K2.5, MiniMax M2.5, Liquid LFM → ⚠️ seuils. Llama 3.3/4 → ⚠️ Llama Community. GLM-5.3, Kimi K3 → ❓ licence non-MIT à lire.

**Réflexe** : à chaque `pull`/`download` d'un nouveau modèle en entreprise, noter (modèle, version, licence, URL du dépôt, date) dans un registre — 5 minutes qui évitent un contentieux.

---

## Annexe K — FAQ : les 20 questions qu'on pose vraiment

**1. Par où commencer si je n'ai jamais touché à un LLM ?**
Ollama + `qwen3:8b` en 10 minutes (§5.5.1). Puis OpenRouter pour comparer avec le cloud. Puis LiteLLM quand tu as 2 apps.

**2. Quel est le modèle local le plus « rentable » en 2026 ?**
Qwen3-32B en Q4_K_M : le meilleur rapport qualité/VRAM (~22 Go). En dessous : Qwen3-8B pour le bulk.

**3. Mon 70B rame à 3 tok/s, c'est normal ?**
Si tu es en offload CPU partiel ou sur CPU seul : oui. En full-GPU 4090×2 : attends 15–30 tok/s. En dessous : vérifie `--n-gpu-layers 99` et le contexte.

**4. Pourquoi mon appel RAG coûte 10× plus que prévu ?**
Tokens invisibles : template de chat, historique, chunks trop gros, outils. Mesure le ratio avec les spend logs (§4.6).

**5. Faut-il un GPU pour les embeddings ?**
Non : `nomic-embed-text` tourne très bien sur CPU (~100 phrases/s). Le GPU ne sert qu'au-delà de ~1M de chunks/jour.

**6. GGUF ou AWQ ?**
GGUF = défaut (Ollama, llama.cpp, partout). AWQ = si tu sers avec vLLM sur NVIDIA et que tu veux le max de débit.

**7. Puis-je mettre un LLM sur un NAS ou un mini-PC ?**
Oui pour les petits modèles : un N100 fait tourner un 3B Q4 à ~10–15 tok/s (à vérifier). Pour du 8B+ interactif : non.

**8. Comment empêcher les hallucinations ?**
Température 0–0,2, RAG avec citations obligatoires, structured output, et consigne « dis je ne sais pas ». Aucune méthode n'atteint 100 %.

**9. Quelle est la différence entre Ollama et llama.cpp ?**
Ollama = llama.cpp (entre autres backends) emballé pour être simple : `pull/run`, API sur 11434. llama.cpp = le moteur nu, plus rapide et plus réglable.

**10. vLLM ou Ollama pour servir une équipe ?**
vLLM : multi-utilisateurs, batching, monitoring. Ollama : 1–2 utilisateurs, simplicité. Les deux derrière LiteLLM.

**11. C'est quoi le « prompt caching » concrètement ?**
Le provider mémorise ton préfixe (system prompt, doc) et ne te facture la lecture qu'à -90 %. En RAG avec gros system prompt : économie massive (Anthropic, OpenAI, DeepSeek, Google).

**12. Pourquoi deux prix (input/output) ?**
Générer coûte plus cher que lire (decode séquentiel). Les modèles de raisonnement gonflent l'output (chaîne de pensée) : surveiller le ratio.

**13. C'est quoi un « reasoning effort » / « thinking budget » ?**
Le nombre de tokens de réflexion autorisés. Plus = meilleur raisonnement, plus cher et plus lent. À régler par cas d'usage.

**14. Mon provider a délisté mon modèle, que faire ?**
C'est pour ça qu'on utilise des noms logiques (§4.8) : changer une ligne du YAML LiteLLM, pas le code. Et toujours 2 providers par cas critique.

**15. Quelle est la vraie différence de prix entre un 8B et un frontier ?**
Facteur ~100–1 000 : ~0,03–0,30 $/1M (8B chez DeepInfra/Groq) vs 3–75 $/1M (frontier). D'où le routage par tâche.

**16. Puis-je fine-tuner sur mon GPU ?**
LoRA d'un 7B : oui sur 24 Go (qlora). Full fine-tune d'un 70B : non (il faut 8×H100 ou un cloud). Together/Fireworks le font en managé.

**17. Air-gap : vraiment aucun réseau ?**
Oui : poids + runtime sur machine isolée, modèles vérifiés au préalable. Prévoir une procédure de mise à jour manuelle (clé USB chiffrée, hash vérifiés).

**18. Comment choisir entre les 50 providers du top ?**
Filtre en 3 questions : (1) quel modèle ? (2) quel budget/latence ? (3) quelle conformité ? Il reste 2–3 candidats. Les tester 1 semaine chacun avec les mêmes evals.

**19. Les benchmarks (SWE-bench, MMLU) sont-ils fiables ?**
Partiellement : contamination possible, et ils ne mesurent pas TON usage. Utiles pour dégrossir, insuffisants pour décider. Toujours compléter par une eval métier (§12).

**20. Que faire quand les prix changent ?**
Rejouer le calcul de seuil (§8.2) chaque trimestre, et laisser le routeur coût de LiteLLM absorber les micro-variations automatiquement.

---

## Annexe L — Lexique express du hardware IA

| Acronyme | Signification | À retenir |
|---|---|---|
| VRAM | Video RAM (mémoire GPU) | Ce qui tient = ce qui tourne |
| HBM | High Bandwidth Memory | Mémoire des GPU datacenter (H100 : 80 Go HBM3) |
| GDDR | Graphics DDR | Mémoire des GPU grand public (4090 : 24 Go GDDR6X) |
| SRAM | Static RAM (on-chip) | Le secret vitesse de Groq (LPU) et Cerebras (WSE) |
| TFLOPS | Tera Floating-Point Ops/s | Puissance de calcul ; compte peu en decode solo |
| To/s (bande passante) | Tera-octets/s | **La métrique qui compte** en génération (4090 : ~1 To/s) |
| TDP | Thermal Design Power | 4090 : 450 W ; 5090 : 575 W ; Mac Studio : ~40–110 W |
| NVLink | Interconnexion GPU-GPU NVIDIA | Pour le multi-GPU rapide (serveurs) |
| PCIe | Bus carte mère | Le goulet entre 2 GPU grand public (plus lent que NVLink) |
| CUDA | Écosystème NVIDIA | Le standard ; absent sur Mac |
| ROCm | Écosystème AMD | Progresse, mais edge cases en inférence |
| MLX | Framework Apple | L'inférence Mac, sans CUDA |
| NPU | Neural Processing Unit | Puces d'inférence (datacenter : Trainium, TPU) |
| Q4/Q8 | Quantification 4/8-bit | Divise la mémoire par ~4/~2 |
| OOM | Out Of Memory | L'erreur n°1 du débutant local |
| TTFT | Time To First Token | La latence perçue (<1 s chez Groq) |
| tok/s | Tokens par seconde | Le débit ; ~15 tok/s = lecture confortable |

---

## Table des matières complète

