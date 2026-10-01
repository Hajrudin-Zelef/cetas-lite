---
id: collect-261001-ia-llm/ia-llm/llm-gateways-infra-25
title: "Passerelles et infra LLM — OpenRouter, Groq, Cerebras, gratuit, self-hosting, TEI, HF, location GPU"
domain: ia-llm
role: reference
task: reference
actors: ["AMD", "Alibaba", "DeepSeek", "Google", "Intel", "Moonshot", "Nvidia", "OpenAI", "SGLang", "TensorRT-LLM", "vLLM"]
dates: ["2026-09-27"]
keywords: ["gpu", "agents", "amd", "apache", "attention", "benchmark", "benchmarks", "decode", "deepseek", "gpus", "intel", "kimi"]
source: docs/RAG/collect-261001-ia-llm/llm_gateways_infra.md
source_anchor: ""
source_lines: [2945, 3035]
sha256: 789ce36c59fcabba96a08f72841fbeec11f3aca1571359e5a792975873c363b9
---

# Passerelles et infra LLM — OpenRouter, Groq, Cerebras, gratuit, self-hosting, TEI, HF, location GPU

- Projet open source (Université de Berkeley / LMSYS à l'origine), licence Apache 2.0, comme vLLM.
- Architecture : **RadixAttention** — un arbre radix des KV caches qui réutilise automatiquement les préfixes communs entre requêtes (prompt système, few-shot, historique multi-tours, documents RAG réinjectés). Là où vLLM recalcule le préfixe à chaque requête (sauf cache explicite), SGLang le partage.
- **Langage de frontend structuré** : SGLang propose aussi un DSL pour la génération contrainte (JSON schema, choix multiples, raisonnement multi-étapes avec `gen`, `select`) — utile pour des agents qui doivent produire du structuré fiable.
- Installation typique :

```bash
pip install sglang
# Serveur OpenAI-compatible
python -m sglang.launch_server --model-path Qwen/Qwen3-8B --port 30000
# curl http://localhost:30000/v1/chat/completions  (même format qu'OpenAI)
```

### 163.2. État au 27/09/2026 (vérifié)

- Ligne de version : **SGLang 0.5.x** (été 2026). **vLLM 0.24.0** (août 2026) — attention au breaking change : vLLM ne définit plus `CUDA_VISIBLE_DEVICES` en interne, un nouvel argument `device_ids` le remplace ; les déploiements qui comptaient sur l'ancien comportement silencieux changent de binding GPU après upgrade.
- **MoE** : SGLang travaille en direct avec DeepSeek/Qwen/Kimi/Z AI sur des implémentations optimisées — sur DeepSeek V3, SGLang revendique **3,1x plus vite que vLLM** ; sur GB200 NVL72, **25x** (février 2026, chiffres vendeur/écosystème).
- **TPU** : backend SGLang-Jax natif — SGLang tourne sur TPU Google, pas seulement NVIDIA.
- **Disaggregation prefill/decode** : devenue une fonctionnalité de production (pas expérimentale) : jusqu'à 6,4x de débit et 20x moins de variance de latence sur les déploiements qui la configurent — vLLM et SGLang la proposent tous deux.
- Point de vigilance : les benchmarks « SGLang écrase vLLM » sont **dépendants du trafic**. Mesures indépendantes 2026 :
  - Trafic à préfixes partagés (agents, multi-tours, RAG) : SGLang ~+29 % de débit vs vLLM sur H100 (classe Llama-3-8B, début 2026) ; multi-tours jusqu'à ~1,75x dans certains setups.
  - **Trafic à prompts uniques** (génération créative one-shot, requêtes de recherche indépendantes) : l'avantage s'effondre à **~1–4 %** (bruit de mesure). RadixAttention n'a rien à réutiliser quand rien n'est partagé.
  - Sur petite GPU consommateur (8 Go, Qwen3-0.6B) : vLLM jusqu'à 2,35x plus rapide *out of the box* dans un test — l'écart venait d'une optimisation GPU que vLLM tenait en mémoire et pas SGLang. Moralité : **benchmarke sur TON matériel et TON trafic**, pas sur des slides.

### 163.3. vLLM vs SGLang — tableau comparatif (sept 2026)

| Critère | vLLM 0.24.0 | SGLang 0.5.x |
|---|---|---|
| Licence | Apache 2.0 | Apache 2.0 |
| Idée centrale | PagedAttention (gestion mémoire KV efficace) | RadixAttention (partage des préfixes entre requêtes) |
| Point fort | Le plus simple à adopter, **support modèles/matériel le plus large** (NVIDIA, AMD, Intel, TPU) | **MoE et gros déploiements multi-nœuds** ; trafic à préfixes répétés |
| API | OpenAI-compatible (`/v1/chat/completions`), port configurable | OpenAI-compatible (port 30000 par défaut) + DSL structuré natif |
| Tool calling | Via `--tool-call-parser` | Natif |
| Génération structurée | Oui (guidance/xgrammar selon version) | Oui, DSL intégré (`gen`, `select`, JSON) |
| Multimodal (« Omni ») | Large support | En retrait sur certains modèles vision |
| Maturité écosystème | La référence : intégrations partout (TGI-like, LiteLLM, etc.) | En forte croissance, support day-one des MoE chinois |
| Cas d'usage idéal | Serveur généraliste « ça marche avec (presque) tout » | Agents, RAG à gros prompts système partagés, serving de MoE |
| Piège 2026 | 0.24 : `device_ids` remplace le `CUDA_VISIBLE_DEVICES` implicite | Avantage nul sur prompts uniques ; perfs sensibles au tuning |

### 163.4. Que choisir pour ton RAG ?

- **Règle simple** : si ton trafic a un **gros préfixe partagé** (prompt système RAG + instructions + few-shot identiques à chaque requête), SGLang mérite un benchmark — c'est exactement son terrain de jeu, et ton RAG est un cas d'école (même system prompt, documents variables).
- Si tu sers des modèles variés, du multimodal, ou si tu veux le chemin le plus documenté : **vLLM** reste le choix par défaut.
- Les deux exposent une API OpenAI-compatible → ton routeur LiteLLM/FreeLLMAPI ne voit pas la différence : tu peux **A/B tester en changeant une URL**, mesurer tok/s et TTFT sur ton trafic réel, et décider avec des chiffres.
- Alternative « je ne veux rien tuner » : **NVIDIA NIM** (section 165) choisit le moteur (vLLM, TensorRT-LLM ou SGLang selon le profil) à ta place.

### 163.5. Génération structurée : le DSL SGLang en 10 lignes

Là où vLLM s'appuie sur des libs externes (xgrammar, guidance), SGLang intègre un mini-langage pour forcer la structure des sorties — décisif quand ton RAG doit produire du JSON valide à tous les coups :

```python
import sglang as sgl

@sgl.function
def diagnostic(s, symptome):
    s += sgl.user(f"Symptôme onduleur : {symptome}. Réponds en JSON strict.")
    s += sgl.assistant(sgl.gen("reponse",
                               max_tokens=256,
                               regex=r'\{"cause": ".+", "gravite": "(critique|majeure|mineure)", "action": ".+"\}'))
    # le regex est CONTRAINT pendant la génération : la sortie matche toujours

# Exécution contre ton serveur SGLang local
state = diagnostic.run(symptome="bypass statique en défaut, charge sur réseau",
                       backend=sgl.Runtime("http://localhost:30000"))
print(state["reponse"])  # JSON garanti valide, ex: {"cause": "...", "gravite": "majeure", "action": "..."}
```

Autres primitives : `sgl.select` (choix parmi une liste fermée — classification sans risque de sortie libre), `sgl.gen` avec `choices`, chaînage multi-étapes avec variables. Pour un RAG qui alimente des tickets GLPI ou des rapports : la génération contrainte élimine toute une classe de bugs de parsing en aval.

### 163.6. Prefill/decode disaggregation : l'idée en 30 secondes

L'inférence a deux phases aux besoins opposés : le **prefill** (lire tout le prompt d'un coup — gourmand en calcul) et le **decode** (générer token par token — gourmand en mémoire/bande passante). Les servir sur les mêmes GPU, c'est faire cohabiter deux workloads qui se gênent.

La disaggregation = des GPU dédiées au prefill, d'autres au decode, avec transfert du KV cache entre les deux. Gains mesurés (2026) : jusqu'à **6,4x de débit** et **20x moins de variance de latence** sur les bons setups. Le prix : de la complexité réseau (transfert KV cache rapide exigé — NVLink/InfiniBand, pas du 1 GbE) et un dimensionnement à deux pools. **Pour ton usage** : pertinent si tu sers un RAG à fort trafic avec de gros prompts (chunks longs) ; overkill pour un prototype ou un usage interne modéré. vLLM et SGLang la proposent tous deux en 2026 — ce n'est plus un argument différenciant, c'est une option d'archi.

### 163.7. Mise en route Docker et benchmark minimal

```bash
# Serveur SGLang via Docker (adapte le tag à la version du moment)
docker run -d --gpus all \
  -p 30000:30000 \
  -v ~/.cache/huggingface:/root/.cache/huggingface \
  --name sglang \
  lmsysorg/sglang:latest \
  python -m sglang.launch_server \
    --model-path Qwen/Qwen3-8B \
    --port 30000 \
    --enable-prefix-caching   # RadixAttention : le flag qui fait la différence

# Attendre la dispo
curl http://localhost:30000/v1/models

