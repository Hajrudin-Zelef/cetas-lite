---
id: collect-261001-ia-llm/ia-llm/llm-gateways-infra-15
title: "Passerelles et infra LLM — OpenRouter, Groq, Cerebras, gratuit, self-hosting, TEI, HF, location GPU"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "Cerebras", "Groq", "OpenAI", "OpenRouter", "vLLM"]
dates: []
keywords: ["gpu", "agent", "astra", "awq", "deepseek", "fp8", "gpt-6", "kv cache", "llama", "opus 5", "parameters", "quantization"]
source: docs/RAG/collect-261001-ia-llm/llm_gateways_infra.md
source_anchor: ""
source_lines: [1837, 1993]
sha256: 0f690cae547e106e867b5ca8f82f9157cb919614ce1b3a4e88882e6a1444e2cb
---

# Passerelles et infra LLM — OpenRouter, Groq, Cerebras, gratuit, self-hosting, TEI, HF, location GPU

```yaml
# litellm_config.yaml
model_list:
  - model_name: rapide
    litellm_params:
      model: groq/openai/gpt-oss-120b
      api_key: os.environ/GROQ_API_KEY_FAKE
  - model_name: local
    litellm_params:
      model: openai/Qwen/Qwen3-32B-AWQ   # ton vLLM
      api_base: http://localhost:8000/v1
      api_key: local
router_settings:
  fallbacks: [{"rapide": ["local"]}]
```

```bash
pip install 'litellm[proxy]'
litellm --config litellm_config.yaml --port 4000
# Puis : base_url=http://localhost:4000, model="rapide"
```

Pour toi : LiteLLM = le **routeur central** de ton labo (un seul endpoint pour tout ton code), avec cache Redis et budgets. FreeLLMAPI = orienté gratuit ; LiteLLM = orienté contrôle/SLA.

## 132. LiteLLM vs OpenRouter vs FreeLLMAPI (tableau)

| Critère | LiteLLM (self-hosted) | OpenRouter (SaaS) | FreeLLMAPI (self-hosted) |
|---|---|---|---|
| Hébergement | Chez toi (Docker/pip) | Cloud | Chez toi (Docker) |
| Catalogue | 100+ providers (tes clés) | 300+ modèles (routés) | Tiers gratuits (~28 providers) |
| Fallback / retry | Oui (config) | Oui (`allow_fallbacks`) | Oui (auto) |
| Cache | Oui (Redis) | Non (côté provider) | Non |
| Budgets par clé | Oui | Oui (garde-fous) | Suivi conso |
| Coût | 0 (+ ta machine) | Marge sur tokens | 0 |
| Idéal | Labo/équipe, contrôle total | Simplicité, comparaison | Empiler le gratuit |

## 133. Structured outputs : forcer du JSON propre

Indispensable quand un LLM pilote du code (inventaire, tickets). Trois voies :

```python
# 1. OpenAI / Groq / Cerebras (support variable — vérifier par modèle)
r = client.chat.completions.create(
    model="openai/gpt-oss-120b",
    messages=[{"role": "user", "content": "Liste 2 onduleurs : nom, puissance_kva"}],
    response_format={"type": "json_object"},
)

# 2. JSON Schema strict (modèles compatibles)
r = client.chat.completions.create(
    model="deepseek/deepseek-v4.1-flash",
    messages=[{"role": "user", "content": "..."}],
    response_format={"type": "json_schema",
        "json_schema": {"name": "ond", "schema": {
            "type": "object",
            "properties": {"nom": {"type": "string"},
                           "puissance_kva": {"type": "number"}},
            "required": ["nom", "puissance_kva"],
            "additionalProperties": False}}},
)

# 3. vLLM local : guided decoding (grammaire garantie, moteur xgrammar)
# vllm serve ... --guided-decoding-backend xgrammar
# puis extra_body={"guided_json": {...schema...}}
```

Sur OpenRouter, ajoute `"provider": {"require_parameters": true}` pour ne router que vers des fournisseurs qui supportent le JSON Schema.

## 134. Le seuil du local : calcul de rentabilité (méthode)

```
Coût cloud mensuel  = tokens_mois/1e6 × prix_1M
Coût local mensuel  = (amortissement GPU / mois) + élec + ton temps

Exemple : 200M tokens/mois sur gpt-oss-120b @ $0.15/$0.60
  ≈ 150M in + 50M out → 150×0.15 + 50×0.60 = $22.5 + $30 = $52.5/mois
Local : RTX 4090 (1900 € / 36 mois ≈ 53 €/mois) + élec 16 € = ~69 €/mois
→ À ce volume, le cloud GAGNE encore. Le local gagne quand :
  - la machine existe déjà (serveur Proxmox allumé 24/7) → coût marginal ~0
  - le volume dépasse ~500M tokens/mois
  - ou la confidentialité l'exige (alors le prix ne compte plus)
```

**Règle** : ne pas acheter de GPU « pour le RAG » tant que le gratuit + un petit payant couvrent le besoin. Acheter quand le besoin est prouvé par les logs (section 103).

## 135. Prompt caching : diviser la facture par 2 à 10

Le prompt système + les chunks RAG se répètent : le **cache de prompt** évite de les re-facturer plein pot.

| Provider | Cache read /1M | Cache write /1M | Note (vérifié sept 2026) |
|---|---|---|---|
| Anthropic Opus 5.5 | **$0.20** | $5 | −60 % vs Opus 5 |
| Anthropic Sonnet 5 | $0.20 | $2.50 | |
| OpenAI gpt-6-astra | $1.00 | $12.50 | Contexte court |
| Groq | **−50 %** input caché | — | Réduction simple |
| vLLM local | 0 (prefix caching) | — | `--enable-prefix-caching` |

En pratique : ordre stable des messages (système → chunks → question), et le cache fait le reste. Sur un agent qui porte 50K tokens de contexte sur 20 tours, le cache **divise la facture par ~5**.

## 136. Batch API : −50 % quand le temps réel ne compte pas

Groq et d'autres offrent le **batch à −50 %** : tu soumets un fichier de requêtes, réponse en heures creuses. Idéal : indexation d'annotations, étiquetage de 100K tickets, évals massives. Workflow : `jobs` en SQLite (section 106) → soumission batch la nuit → résultats au matin. Le local (vLLM) reste imbattable si tu as le GPU, mais le batch cloud évite d'immobiliser ta machine.

## 137. Ollama avancé : variables d'environnement utiles

```bash
# /etc/systemd/system/ollama.service.d/override.conf (ou EnvironmentFile)
OLLAMA_HOST=0.0.0.0:11434        # exposer sur le LAN (avec reverse proxy !)
OLLAMA_NUM_PARALLEL=4            # requêtes simultanées (défaut 1 → à monter en équipe)
OLLAMA_MAX_LOADED_MODELS=2       # modèles gardés en VRAM
OLLAMA_KEEP_ALIVE=30m            # garder le modèle chargé 30 min après usage
OLLAMA_MAX_QUEUE=512
```

```bash
ollama ps                        # modèles chargés + VRAM utilisée
ollama show qwen3:32b            # paramètres, template, quantization
```

`OLLAMA_KEEP_ALIVE` évite le rechargement (10–30 s) entre deux questions espacées — essentiel pour un usage interactif sans latence de cold start.

## 138. vLLM avancé : ce qui change la donne en prod

- **`--quantization fp8`** : 2× moins de VRAM que FP16, perte quasi nulle sur GPU récents (Ada/Hopper) — le meilleur choix 2026 si ton GPU le supporte.
- **Tool calling** : vLLM supporte le function calling OpenAI-compatible sur les modèles adaptés (Qwen3, Llama 3.x) — ton agent local peut appeler des outils.
- **`/metrics`** : exporter vers Prometheus → Grafana (latence, file d'attente, KV cache utilisé).
- **Multi-modèles** : un vLLM = un modèle ; pour 2 modèles, 2 conteneurs (ou Ollama qui multiplexe).
- **Versioning** : épingler `vllm/vllm-openai:vX.Y.Z` — les régressions entre versions existent.

## 139. Stockage des vecteurs : Chroma vs pgvector vs Qdrant

| Solution | Install | Perf | Quand |
|---|---|---|---|
| **Chroma** | `pip install chromadb` (embarqué) | OK < 1M vecteurs | Proto, RAG perso |
| **pgvector** (Postgres) | Extension PG | Bon, SQL + vecteurs | Tu as déjà Postgres ; filtres métier en SQL |
| **Qdrant** | Docker 1 conteneur | Excellent, filtres riches | Volume, prod, payloads complexes |

Pour ton RAG : **Chroma pour démarrer**, **pgvector** si tu veux croiser vecteurs + métadonnées métier (site, équipement, date) en SQL, **Qdrant** si > 1M chunks. Tous acceptent les 1024 dim de bge-m3. Sauvegarde = le dossier de données (ou mieux : le script de réindexation, section 109).

## 140. Chunking : la qualité du RAG se joue ici

1. **Taille** : 500–1 000 tokens pour de la doc technique (un chunk = une idée : une procédure, un tableau).
2. **Overlap** : 10–15 % (100–150 tokens) pour ne pas couper une phrase clé.
3. **Respecter la structure** : découper sur les titres (Markdown `##`), jamais au milieu d'un tableau ou d'une commande.
4. **Métadonnées par chunk** : source, page, section, date, équipement → filtres + citations.
5. **Tester** : même corpus, 3 stratégies de chunking, mesurer précision@5 (section 34) — le chunking bat souvent le changement de modèle.

```python
# Exemple : découpage simple sur titres Markdown
import re
def chunk_md(text, max_tokens=800):
    parts = re.split(r'(?m)^## ', text)
    # ... regrouper les petites sections, couper les trop grosses sur les paragraphes
```

## 141. Évaluer un RAG : les métriques qui comptent

