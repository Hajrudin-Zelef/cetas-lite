---
id: collect-261001-ia-llm/ia-llm/llm-gateways-infra-3
title: "Passerelles et infra LLM — OpenRouter, Groq, Cerebras, gratuit, self-hosting, TEI, HF, location GPU"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "Cerebras", "Google", "Groq", "OpenAI", "OpenRouter"]
dates: ["2026-08-27", "2026-09-27"]
keywords: ["gpu", "agents", "benchmarks", "claude", "embeddings", "gemini", "gpt-6", "llama", "qwen", "transcription", "wafer"]
source: docs/RAG/collect-261001-ia-llm/llm_gateways_infra.md
source_anchor: ""
source_lines: [241, 365]
sha256: 9830fb5efbb9e2b118a5a34db02320fa8c66f8ab16b4fccf5551fb0c7329ec8c
---

# Passerelles et infra LLM — OpenRouter, Groq, Cerebras, gratuit, self-hosting, TEI, HF, location GPU

## 18. Groq : free tier — limites vérifiées

- **Toujours gratuit avec limites**, sans carte bancaire. Jamais facturé : au-delà du quota → HTTP 429.
- Ordres de grandeur (varient par modèle, limites **au niveau de l'organisation** — multiplier les clés ne contourne pas) :
  - **30 req/min** (RPM) sur les modèles texte courants ;
  - **1 000 à 14 400 req/jour** selon le modèle (ex. gpt-oss-120b : 30/1 000/8K/200K en RPM/RPD/TPM/TPD d'après le relevé du 27/08/2026) ;
  - Whisper : 20 RPM, 2 000 req/jour.
- Le tier payant (pay-as-you-go) multiplie les limites par ~10 et ajoute le batch à −50 %.
- Données : pas d'entraînement sur tes prompts ; rétention jusqu'à 30 jours (ZDR disponible selon l'offre — à vérifier dans les réglages du compte).

## 19. Groq : exemple d'appel

```bash
curl https://api.groq.com/openai/v1/chat/completions \
  -H "Authorization: Bearer gsk_FAKEFAKEFAKEFAKE" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "openai/gpt-oss-20b",
    "messages": [{"role": "user", "content": "Résume ce log en une ligne : ..."}],
    "temperature": 0.2,
    "max_tokens": 200
  }'
```

```python
from openai import OpenAI
client = OpenAI(base_url="https://api.groq.com/openai/v1", api_key="gsk_FAKEFAKEFAKEFAKE")
r = client.chat.completions.create(
    model="openai/gpt-oss-120b",
    messages=[{"role": "system", "content": "Tu es un assistant réseau."},
              {"role": "user", "content": "Diagnostique : ..."}],
)
print(r.choices[0].message.content)
```

## 20. Groq : cas d'usage (où il bat tout le monde)

1. **Agents temps réel / voix** : 500–1 000 tok/s → une réponse de 200 tokens arrive en < 1 s. Idéal pour un assistant vocal d'atelier.
2. **RAG rapide** : génération de la réponse finale sur tes chunks — la récupération (embeddings) est locale, la génération est instantanée.
3. **Classification / triage à haut volume** : `gpt-oss-20b` à 1 000 tok/s en gratuit, parfait pour étiqueter des tickets ou des logs.
4. **Transcription** : Whisper turbo à $0.04/h (ou gratuit dans les quotas) — comptes-rendus de réunions d'équipe.
5. **Dev itératif** : boucle code→test→fix sans attendre.

## 21. Groq : limites

1. **Contexte** : 131K max sur les modèles phares, souvent moins en pratique sur le free tier — pas de 1M tokens ici.
2. **Catalogue fermé aux poids ouverts** : pas de GPT-6, pas de Claude, pas de Gemini. Pour le raisonnement le plus dur, prévoir un second provider.
3. **Quotas gratuits modestes en TPM** (8K–30K/min selon modèle) : les gros batchs passent au payant.
4. **Modèles dépréciés sans préavis long** : épingler les IDs exacts et surveiller `console.groq.com/docs/models`.
5. **Pas de vision** sur la plupart des modèles texte (à vérifier par modèle).

## 22. Checklist Groq

- [ ] Compte sur `console.groq.com`, clé `gsk_...` (pas de CB).
- [ ] Tester `openai/gpt-oss-20b` (rapide) et `openai/gpt-oss-120b` (costaud).
- [ ] Mesurer le temps de réponse réel sur tes prompts (objectif < 2 s).
- [ ] Coder la gestion du 429 (retry avec backoff) — obligatoire en free tier.
- [ ] Si volume : comparer le pay-as-you-go (−25 % tier dev) vs OpenRouter.

---

# PARTIE C — CEREBRAS

## 23. Cerebras : présentation et Wafer-Scale Engine

Cerebras (cerebras.ai) est, comme Groq, un **fondeur d'inférence** : sa puce, le **Wafer-Scale Engine (WSE)**, est le plus grand processeur jamais gravé — **un wafer entier** (pas une puce découpée), avec des dizaines de Go de SRAM sur la puce et une bande passante mémoire de l'ordre de **21 Po/s sur le wafer** (contre ~80 To/s par rack côté LPU Groq et ~3,3 To/s par GPU H100 — ordres de grandeur constatés dans les comparatifs 2026, à prendre comme tels).
Résultat : **1 800 à 3 000 tokens/s** mesurés sur les petits modèles (Llama 3.1 8B à 1 800 tok/s dès 2024), et **2 522 tok/s** relevés sur Llama 4 Maverick en 2026. Cerebras reste en **précision 16 bits de bout en bout** (pas de sacrifice de précision pour la vitesse, d'après leurs benchmarks vérifiés par Artificial Analysis).
API compatible OpenAI : `https://api.cerebras.ai/v1`. **Free tier permanent sans carte** : 1M tokens/jour.

## 24. Cerebras : modèles servis vérifiés (sept 2026)

| Modèle | Contexte | Vitesse constatée | Prix payant /1M in/out (relevé sept 2026, sources tierces — à reconfirmer) |
|---|---|---|---|
| `gpt-oss-120b` | 131K | ~2 000 tok/s | $0.35 / $0.75 |
| `llama-3.1-8b` | 128K | ~1 800–3 000 tok/s | $0.10 / $0.10 |
| `llama-3.3-70b` | 128K | ~450–1 000 tok/s | $0.85 / $1.20 |
| `qwen-3-32b` | 128K | à vérifier | à vérifier |
| `gemma-4-31b` | à vérifier | à vérifier | à vérifier |
| `llama-4-maverick` | à vérifier | 2 522 tok/s (relevé tiers) | à vérifier |

Catalogue **plus étroit** que Groq/OpenRouter : poids ouverts uniquement (Llama, Qwen, gpt-oss, Gemma), **pas de vision**, pas de modèles fermés. Les fine-tunes communautaires exotiques mettent plus de temps à arriver (compilation spécifique au wafer).

## 25. Cerebras : free tier et tarification vérifiés

- **Free tier permanent, sans carte bancaire** : **1 000 000 tokens/jour** (reset quotidien), ~5 RPM, 30K–100K TPM selon le modèle (relevés convergents août–sept 2026).
- **Limite temporaire constatée** : contexte plafonné à **8K tokens** sur le free tier (annoncé comme temporaire par la communauté — à re-vérifier, bloquant pour du RAG à gros chunks).
- Tier développeur : dépôt min ~$10, facturation à l'usage (prix du tableau section 24).
- $5 de crédits d'essai à l'inscription constatés sur certaines pages (à vérifier — l'offre bouge).
- Documentation vie privée plus légère que chez les concurrents : pour des données sensibles, préférer un provider avec ZDR explicite.

## 26. Cerebras vs Groq : tableau comparatif

| Critère | Cerebras (WSE) | Groq (LPU) |
|---|---|---|
| Techno | Wafer entier, SRAM on-chip massive | Puces LPU déterministes |
| Débit max constaté | **1 800–3 000 tok/s** (8B), 2 522 tok/s (Maverick) | 500–1 000 tok/s |
| Latence premier token | Très faible | Très faible (~100 ms) |
| Free tier | **1M tokens/jour**, sans CB | 30 RPM, 1K–14,4K req/jour, sans CB |
| Contexte free | **8K cap** (temporaire, à surveiller) | 131K (selon modèle) |
| Prix payant (ex. 120B) | $0.35 / $0.75 | **$0.15 / $0.60** |
| Prix petit modèle | $0.10 / $0.10 (8B) | **$0.075 / $0.30** (20B) |
| Catalogue | Étro
...[truncated 12332 chars]
## 29. Exemple d'appel Cerebras

```bash
curl https://api.cerebras.ai/v1/chat/completions \
  -H "Authorization: Bearer csk-FAKEFAKEFAKEFAKE" \
  -H "Content-Type: application/json" \
  -d '{
    "model": "llama-3.1-8b",
    "messages": [{"role": "user", "content": "Liste les étapes de consignation électrique avant intervention sur un onduleur."}],
    "max_tokens": 500,
    "temperature": 0.3
  }'
```

Même client Python OpenAI, `base_url="https://api.cerebras.ai/v1"`.

---

# PARTIE D — COMPARATIF : QUAND UTILISER QUOI

## 30. Tableau de synthèse prix / latence / choix (prix vérifiés le 27/09/2026 sauf mention)

