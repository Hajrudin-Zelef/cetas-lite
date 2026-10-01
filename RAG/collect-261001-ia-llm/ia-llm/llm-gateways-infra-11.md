---
id: collect-261001-ia-llm/ia-llm/llm-gateways-infra-11
title: "Passerelles et infra LLM — OpenRouter, Groq, Cerebras, gratuit, self-hosting, TEI, HF, location GPU"
domain: ia-llm
role: reference
task: reference
actors: ["AWS", "Alibaba", "Lambda", "OpenAI", "vLLM"]
dates: []
keywords: ["gpu", "arr", "awq", "aws", "embeddings", "lora", "quantization", "qwen", "vllm"]
source: docs/RAG/collect-261001-ia-llm/llm_gateways_infra.md
source_anchor: ""
source_lines: [1349, 1458]
sha256: 1be4cd56a103d41f545213d1f27d15ae366ec9c705d425c770417f6e30c0743b
---

# Passerelles et infra LLM — OpenRouter, Groq, Cerebras, gratuit, self-hosting, TEI, HF, location GPU

| Critère | Vast.ai | RunPod |
|---|---|---|
| Modèle | Marketplace P2P (enchères) | Cloud dev (community + secure) |
| RTX 4090 typique | **$0.29–0.59/h** (spot→on-demand) | $0.34/h (community) / $0.59–0.74/h (secure) |
| A100 80 Go | ~$0.75–1.45/h | ~$1.19–2.49/h |
| H100 80 Go | ~$1.65–2.90/h | ~$2.49–3.89/h |
| Facturation | **À la seconde** | À la milliseconde (serverless) / heure |
| Egress | **$0.01–0.02/Go** (à re-vérifier) | **Gratuit** |
| Fiabilité | Variable (noter l'hôte, `reliability`) | Secure > Community |
| Démarrage | CLI `vastai` (bien) | UI + templates 1-clic (mieux) |
| API | CLI/SDK correct | **Python SDK + GraphQL + runpodctl** (excellent) |
| Serverless | Non (instances) | **Oui** (scale-to-zero) |
| Idéal pour | **Prix plancher**, batch, fine-tune | Dev rapide, API serverless, équipe |

**Verdict** : Vast.ai pour le **prix le plus bas** (batch, LoRA, tests) ; RunPod pour la **simplicité** (templates, serverless, egress gratuit). Les deux battent AWS/GCP (~2–3× moins cher) et écrasent l'achat si < 1 000 h/an.

## 95. Cas d'usage location GPU (concrets pour toi)

1. **Fine-tune LoRA d'un 8B sur tes docs** (8 h sur 4090) : ~$2.80 (Vast spot) à $5.92 (RunPod secure). Le prix d'un sandwich pour un modèle à ta sauce.
2. **Batch d'embeddings massif** : 500M tokens sur TEI GPU loué 2 h ≈ $1 — vs $10 chez OpenAI par passe.
3. **Test 70B Q4** : 2 h sur A100 80 Go ≈ $2.40 — pour décider si ça vaut l'achat d'une 2ᵉ 4090.
4. **Inférence API temporaire** : RunPod serverless vLLM pour une démo équipe, scale-to-zero après.
5. **Entraînement complet 70B** : 24 h sur 8× H100 ≈ $420–630 — à ne faire qu'avec un vrai besoin (et des checkpoints fréquents).

## 96. Pièges Vast.ai / RunPod (15 points de vigilance)

1. **Oublier de détruire l'instance** : n°1 des factures surprises → script d'auto-destruction (section 98).
2. **Spot préempté en plein fine-tune** : checkpoints toutes les N steps + reprise auto, sinon tu paies pour rien.
3. **Egress Vast.ai** : rapatrier un checkpoint 70B (140 Go) ≈ $1.40–2.80 — prévoir, ou rester chez RunPod (gratuit).
4. **Fiabilité hôte** : sur Vast.ai, filtrer `reliability > 0.98` et éviter les offres anormalement pas chères (surchargées).
5. **Disque trop petit** : 60 Go min (poids + cache + datasets) ; le resize après coup est galère.
6. **CUDA mismatch** : image `cu12.4` sur hôte driver 12.1 → vérifier `cuda_max_good` dans l'offre.
7. **SSH par mot de passe** : interdit — clé uniquement, et firewall sur les ports exposés.
8. **Données sensibles** : machine d'un tiers → jamais de secrets/clés réelles, jamais de données clients ; instance jetable.
9. **Scripts de minage** : certains templates communautaires sont douteux → images officielles ou tes propres Dockerfiles.
10. **Facturation à la seconde ≠ gratuite à l'arrêt** : « stopped » ≠ « destroyed » — seul `destroy`/`terminate` stoppe la facture (le disque peut continuer à coûter).
11. **Volume persistant oublié** : le network volume RunPod / Vast.ai se facture même sans instance.
12. **Ports exposés par défaut** : Jupyter sans token sur Internet = cadeau — token + auth systématiques.
13. **Timezone/fuseau** : les jobs longs planifiés en heures creuses (moins de concurrence sur le spot).
14. **Double facturation régions** : vérifier la devise et les taxes (VAT) sur la facture.
15. **Pas de SLA en community/spot** : jamais de prod client critique dessus.

## 97. Template RunPod « RAG-lab » (exemple)

```dockerfile
# Dockerfile — image perso à pousser sur Docker Hub puis utiliser comme template
FROM runpod/pytorch:2.4.0-py3.11-cuda12.4.1-devel-ubuntu22.04
RUN pip install vllm transformers sentence-transformers
# Pré-télécharger les poids à la construction (pas à chaque location)
RUN python -c "from huggingface_hub import snapshot_download; \
    snapshot_download('Qwen/Qwen3-32B-AWQ')"
COPY start.sh /start.sh
EXPOSE 8000 22
CMD ["/start.sh"]
```

```bash
# start.sh
#!/bin/bash
vllm serve Qwen/Qwen3-32B-AWQ --quantization awq \
  --host 0.0.0.0 --port 8000 --enable-prefix-caching &
sshd -D
```

→ Template RunPod pointant sur cette image : **pod vLLM prêt en ~2 min**, coût ≈ $0.34/h en community.

## 98. Auto-destruction : le garde-fou obligatoire

```bash
# À la fin de TON script de job (fine-tune, batch...) — Vast.ai
vastai destroy instance $VAST_INSTANCE_ID

# RunPod : via l'API dans le script Python de fin de job
import runpod, os
runpod.api_key = os.environ["RUNPOD_API_KEY_FAKE"]
runpod.terminate_pod(os.environ["RUNPOD_POD_ID"])

# Ceinture + bretelles : cron local qui tue les instances de plus de X heures
# (ex. toutes les 30 min, tue les pods de test de plus de 4 h)
```

Et côté fournisseur : **plafonds de dépense** (RunPod : limites dans les réglages ; Vast.ai : solde prépayé = limite naturelle — ne charger que ce qu'il faut).

## 99. Quand louer plutôt qu'acheter (décision)

| Situation | Verdict |
|---|---|
| < 1 000 h GPU/an, besoins ponctuels | **Louer** |
| Tests avant achat (70B ? 2× 4090 ?) | **Louer** pour décider |
| Fine-tune occasionnel, batch mensuel | **Louer** (spot) |
| Inférence 24/7 pour ton RAG perso | **Acheter** (seuil ~3 800 h) ou… gratuit cloud |
| Données sensibles / secret industriel | **Acheter** (ou on-prem pro) — jamais de P2P |
| Équipe, SLA, prod | **Secure cloud** (RunPod Secure, Lambda) ou on-prem |

## 100. Checklist location GPU

- [ ] Compte Vast.ai + RunPod, solde prépayé modeste (pas de CB à débit auto si possible).
- [ ] Clé SSH dédiée (`~/.ssh/id_gpu`) ajoutée aux deux plateformes.
- [ ] Template/image perso prêt (section 97) — ne jamais partir d'une image inconnue.
- [ ] Script de job avec **checkpoint + auto-destruction** (sections 95, 98).
- [ ] Estimation du coût AVANT (`prix/h × heures`) notée quelque part.
- [ ] Vérifier l'egress si rapatriement de poids (Vast.ai).
- [ ] Aucun secret réel sur l'instance (variables d'env jetables).

---

# PARTIE J — STRATÉGIE BUDGET ET OPÉRATIONS

## 101. Architecture cible « facture ~0 € » pour ton RAG

