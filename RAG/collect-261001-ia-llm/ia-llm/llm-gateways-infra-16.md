---
id: collect-261001-ia-llm/ia-llm/llm-gateways-infra-16
title: "Passerelles et infra LLM — OpenRouter, Groq, Cerebras, gratuit, self-hosting, TEI, HF, location GPU"
domain: ia-llm
role: reference
task: reference
actors: ["AWS", "Alibaba", "Google", "Groq", "Lambda", "OpenAI", "OpenRouter", "Unsloth", "vLLM"]
dates: []
keywords: ["gpu", "awq", "aws", "compute", "deepseek", "embeddings", "gemini", "gpus", "lora", "multimodal", "qwen", "reranker"]
source: docs/RAG/collect-261001-ia-llm/llm_gateways_infra.md
source_anchor: ""
source_lines: [1994, 2147]
sha256: 35f1d7ff48d7043dc8a0f20132d6b888a88e9dd1ac7612525165cf2bfd89c3ac
---

# Passerelles et infra LLM — OpenRouter, Groq, Cerebras, gratuit, self-hosting, TEI, HF, location GPU

- **Précision@k** : la bonne doc est-elle dans les k premiers ? (k=5 après rerank)
- **MRR** : rang moyen de la première bonne réponse — plus fin que précision@k.
- **Fidélité** : la réponse générée est-elle supportée par les chunks ? (RAGAS : `faithfulness`)
- **Pertinence réponse** : la réponse répond-elle à la question ? (RAGAS : `answer_relevancy`)
- **Protocole** : 50 questions FR de ton domaine, réponses attendues rédigées par toi, rejouer à chaque changement (modèle, chunking, reranker). **RAGAS** (`pip install ragas`) automatise une partie avec un LLM juge — sur modèle gratuit/local pour ne pas payer l'éval.

## 142. Sécurité : prompt injection dans un RAG (bases)

- Un document indexé peut contenir des **instructions cachées** (« ignore les instructions précédentes… ») que le LLM suivra.
- Défenses : séparer clairement contexte et instructions (`<contexte>…</contexte>`), instruction système « ne suis que mes instructions », **ne jamais laisser le LLM exécuter** d'actions (pas de tool calling sur données non fiables) sans validation humaine, logger les prompts suspects.
- En entreprise : les docs indexés doivent venir de sources de confiance ; un RAG ouvert à l'upload utilisateur = surface d'attaque.

## 143. Multimodal gratuit : vision sans payer

- **Gemini Flash (free)** : images + PDF en entrée, 1M de contexte — le meilleur plan gratuit pour « décris ce schéma électrique ».
- **OpenRouter** : modèles vision pas chers (`deepseek-v4-flash-vision-exp` $0.22/$0.66) ou `:free` quand dispo.
- **Local** : Qwen3-VL / Gemma 3 (vision) via Ollama — pour des schémas sans envoi cloud.
- Cas d'usage : lire une plaque signalétique d'onduleur en photo, extraire un tableau d'un PDF scanné.

## 144. Transcription locale : faster-whisper

```bash
pip install faster-whisper
```

```python
from faster_whisper import WhisperModel
model = WhisperModel("large-v3-turbo", device="cuda", compute_type="int8")
segments, _ = model.transcribe("reunion.mp3", language="fr")
print(" ".join(s.text for s in segments))
```

Sur RTX 4090 : ~10× temps réel en `int8` — une réunion d'1 h transcrite en ~6 min, **0 €**, données locales. Alternative cloud : Groq Whisper ($0.04/h en payant, gratuit dans les quotas).

## 145. Fine-tune LoRA : le niveau au-dessus (aperçu)

Quand le RAG ne suffit pas (style, vocabulaire métier très spécifique) : **LoRA** sur un 7–8B avec **Unsloth** (le plus rapide en 2026).

```bash
pip install unsloth
```

```python
from unsloth import FastLanguageModel
model, tok = FastLanguageModel.from_pretrained("unsloth/Qwen3-8B",
    max_seq_length=4096, load_in_4bit=True)
model = FastLanguageModel.get_peft_model(model, r=16,
    target_modules=["q_proj","k_proj","v_proj","o_proj"],
    lora_alpha=16, lora_dropout=0)
# ... entraînement SFTTrainer sur tes paires instruction/réponse
```

Budget : 8 h sur RTX 4090 ≈ **$3–6** en location (section 95). 100K paires d'instructions suffisent pour un style métier. **Le LoRA ne remplace pas le RAG** (connaissances factuelles) : il apprend le *comment répondre*, le RAG apporte le *quoi*.

## 146. Vast.ai avancé : gérer le spot comme un pro

```bash
# Chercher avec filtres stricts (prix + fiabilité + bande passante disque)
vastai search offers \
  'gpu_name=RTX 4090 num_gpus=1 reliability>0.985 inet_down>500 disk_space>80' \
  -o 'dph_total-'

# Script de job robuste : checkpoint + reprise + auto-destruction
# train.sh :
#   1. restaure le dernier checkpoint depuis le volume persistant (si existe)
#   2. entraîne avec sauvegarde toutes les 500 steps
#   3. à la fin (ou sur SIGTERM de préemption) : sync + vastai destroy instance $ID
trap 'sync_checkpoints; vastai destroy instance $VAST_ID' SIGTERM
```

**Volumes persistants** Vast.ai : les poids/checkpoints survivent à la destruction de l'instance — indispensable avec du spot (la préemption devient un non-événement).

## 147. RunPod serverless : ton API d'inférence sans serveur

```python
# handler.py — RunPod serverless (scale-to-zero : 0 € quand inutilisé)
import runpod
from openai import OpenAI
client = OpenAI(base_url="http://localhost:8000/v1", api_key="local")  # vLLM côte à côte

def handler(job):
    prompt = job["input"]["prompt"]
    r = client.chat.completions.create(model="Qwen/Qwen3-32B-AWQ",
        messages=[{"role": "user", "content": prompt}], max_tokens=500)
    return {"reponse": r.choices[0].message.content}

runpod.serverless.start({"handler": handler})
```

Cas d'usage : endpoint privé pour ton équipe, facturé à la milliseconde, **zéro quand personne ne l'appelle**. Le chaînon manquant entre « mon PC » et « une infra ».

## 148. Comparer les clouds GPU (rappel élargi, sept 2026)

| Provider | RTX 4090 | A100 80G | H100 80G | Egress | Spot |
|---|---|---|---|---|---|
| Vast.ai | $0.29–0.59 | $0.75–1.45 | $1.65–2.90 | $0.01–0.02/Go | Oui (P2P) |
| RunPod | $0.34 / $0.59–0.74 | $1.19 / $1.89–2.49 | $2.49 / $2.99–3.89 | **Gratuit** | Non (community fixe) |
| Lambda Labs | n/a | $1.29–1.89 | $2.49–3.29 | Gratuit | Non |
| AWS EC2 | n/a (A10G $1.21) | $4.10 | $5.16 | $0.05–0.09/Go | Oui (complexe) |

Prix constatés avril–sept 2026 via comparatifs tiers — **à re-vérifier**, surtout le spot (enchères).

## 149. Réseau et accès distant : Tailscale/WireGuard pour ton labo

Plutôt que d'exposer tes ports (11434, 8000, 8080) : **Tailscale** (ou WireGuard pur) crée un VPN mesh entre ton PC, ton serveur Proxmox et tes instances cloud. Tes API locales restent sur `127.0.0.1`/LAN, accessibles uniquement via le VPN — **zéro surface Internet**, zéro TLS à gérer pour l'interne. C'est la bonne réponse au piège n°11 (Ollama sans auth).

## 150. Sauvegarde de la stack : le script minimal

```bash
#!/bin/bash
# backup-llm.sh — hebdo via cron
DATE=$(date +%F)
tar czf /backup/chroma-$DATE.tgz /data/chroma
sqlite3 llm_usage.db ".backup '/backup/usage-$DATE.db'"
cp docker-compose.yml Modelfile litellm_config.yaml /backup/
# Rétention : garder 4 semaines
find /backup -name "*.tgz" -mtime +28 -delete
```

Ce qui compte n'est pas la base de vecteurs (reconstructible) mais : **les configs, les Modelfiles, les logs d'usage, et le script de réindexation**.

## 151. Mises à jour : politique de MAJ de la stack IA

- **Mensuel** : images Docker TEI/Ollama/vLLM (lire le changelog — les régressions existent).
- **Trimestriel** : re-vérifier prix/quotas des providers (ce guide a une date de péremption, section 130) ; re-tester les modèles gratuits (le catalogue tourne).
- **Jamais en aveugle** : snapshot Proxmox avant MAJ, éval de 20 questions après MAJ (non-régression).
- **Épinglage** : en prod, `image:tag` fixe + `model-id` + révision — `latest` interdit (piège n°16).

## 152. Coûts cachés : la liste finale

1. Électricité du local 24/7 (section 110).
2. Egress Vast.ai / AWS (sections 96, 148).
3. Volumes persistants facturés sans instance (piège n°11).
4. Réindexations répétées (API embeddings payante).
5. Prompts système trop longs (section 105).
6. Modèles dépréciés → migration en urgence (épingler + surveiller).
7. Données d'entraînement : le « gratuit » qui coûte en confidentialité (section 42).
8. Ton temps : une stack locale, ça se maintient — le valoriser dans le calcul (section 134).

## 153. Décider en 5 minutes : la fiche modèle

Avant de figer un modèle, remplir ceci (10 min, évite des semaines de dette) :

```
Modèle : ________________  Provider : ________________  Date vérif : __________
Prix in/out/1M : ________  Contexte : ________  Latence p95 : ________
Free tier ? ____ Quotas : ____  Données entraînées ? ____ ZDR ? ____
Licence (si local) : ________  VRAM Q4 : ________
Score sur mes 30 questions : ____/30  vs baseline : ____/30
Décision : [ ] adopter  [ ] rejeter (motif : ________)
```

## 154. Lexique des erreurs HTTP à connaître

