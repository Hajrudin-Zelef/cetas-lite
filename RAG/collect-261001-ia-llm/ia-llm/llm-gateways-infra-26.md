---
id: collect-261001-ia-llm/ia-llm/llm-gateways-infra-26
title: "Passerelles et infra LLM — OpenRouter, Groq, Cerebras, gratuit, self-hosting, TEI, HF, location GPU"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "DeepSeek", "Hugging Face", "Moonshot", "Nvidia", "OpenAI", "SGLang", "TensorRT-LLM", "Unsloth", "vLLM"]
dates: ["2026-07-08", "2026-09-02", "2026-09-27"]
keywords: ["gpu", "attention", "benchmark", "decode", "deepseek", "embedding", "fine-tuning", "fp8", "gguf", "int4", "kimi", "kv cache"]
source: docs/RAG/collect-261001-ia-llm/llm_gateways_infra.md
source_anchor: ""
source_lines: [3036, 3170]
sha256: d24654f8a9a657723062452441d1765ecf2f64e64787d8c0e5578ac23d592ea5
---

# Benchmark de TON trafic : le script officiel, avec TES prompts
python -m sglang.bench_serving \
  --backend sglang \
  --dataset-name random \
  --random-input-len 3000 --random-output-len 300 \
  --num-prompts 100 --max-concurrency 32
# Refais pareil côté vLLM (benchmark_serving.py), compare tok/s, TTFT p50/p99, VRAM.
# Règle : même GPU, même modèle, même prompts, 3 runs — sinon le benchmark ne vaut rien.
```

Flags utiles : `--tp 2` (tensor parallel sur 2 GPU), `--quantization fp8` (si supporté par ton modèle/GPU), `--tool-call-parser qwen25` (tool calling natif). La doc de référence : docs.sglang.ai — vérifie les flags à chaque upgrade mineur, ça bouge vite.

### 163.8. Check-list : dois-tu migrer de vLLM vers SGLang ?

Ne migre pas sur un benchmark lu en ligne. Migre si tu coches au moins 3 cases :

- [ ] Ton trafic a un **gros préfixe partagé** (prompt système RAG > 1K tokens, identique à chaque requête)
- [ ] Tu sers un **MoE** (DeepSeek V3, Qwen3-235B, Kimi K2…) où SGLang a le support day-one
- [ ] Tu as besoin de **génération structurée intensive** (JSON garanti, classification fermée)
- [ ] Ton benchmark maison (même GPU, même modèle, même prompts, 3 runs) montre **> 15 %** de gain
- [ ] Tu acceptes un écosystème plus jeune (moins de réponses StackOverflow, doc qui bouge)

Si tu ne coches rien : reste sur vLLM, ton temps vaut plus que 3 % de débit. Si tu coches 3+ : fais tourner les deux en parallèle derrière ton routeur pendant 2 semaines (SGLang sur 20 % du trafic), compare TTFT p99 et taux d'erreur réels, puis décide. Le coût de la migration = zéro changement de code côté appelant (API OpenAI-compatible des deux côtés) — c'est un changement d'URL et de monitoring, pas un rewrite.

---

## 164. PyTorch : les bases pour l'IA (tensors, inférence, écosystème)

**En une phrase :** la fondation sur laquelle tournent (presque) tous les modèles que tu utilises — transformers, vLLM, SGLang, Unsloth ne sont que des couches au-dessus de `torch`. Comprendre les tensors et l'inférence te rend autonome quand ça casse.

### 164.1. État au 27/09/2026

- Version stable : **PyTorch 2.13** (08/07/2026) ; **2.14** sortie le 02/09/2026. Cycle : une mineure tous les ~2 mois. Les images RunPod listent 2.9 → 2.13 ; CUDA supporté : 12.8/12.9/13.0.
- Points à connaître : `torch.compile` est mature (accélère l'inférence en compilant le graphe), l'export ONNX utilise désormais `dynamo=True` par défaut (pipeline `torch.export`), et `torchaudio.io` a été supprimé au profit de `torchcodec` (depuis 2.9).
- Installation standard (adapte `cu130` à ton driver — `nvidia-smi` te dit la version CUDA max) :

```bash
pip install torch --index-url https://download.pytorch.org/whl/cu130
python -c "import torch; print(torch.__version__, torch.cuda.is_available())"
```

### 164.2. Les tensors : 5 minutes pour comprendre

Un tensor = un tableau n-dimensionnel typé, sur CPU ou GPU. Tout le reste (modèles, gradients, KV cache) n'est que des tensors et des opérations dessus.

```python
import torch

# Création
x = torch.tensor([[1.0, 2.0], [3.0, 4.0]])          # 2x2, CPU, float32
w = torch.randn(2, 2, device="cuda")                # aléatoire, sur GPU
print(x.shape, x.dtype, x.device)                   # torch.Size([2, 2]) torch.float32 cpu

# Opérations : broadcasting, matmul (le cœur des transformers)
y = x @ x.T                                         # multiplication matricielle
z = torch.softmax(y / 8 ** 0.5, dim=-1)             # softmax — oui, c'est l'attention en miniature

# Le gradient : la base de l'entraînement (et de rien d'autre en inférence)
p = torch.tensor(2.0, requires_grad=True)
loss = (p ** 2 - 4) ** 2
loss.backward()
print(p.grad)                                       # d(loss)/dp — l'optimiseur s'en sert pour mettre à jour p
```

Trois règles d'or :
1. **dtype** : `float32` (précis, lent, 4 octets), `float16`/`bfloat16` (2 octets — le standard inférence/entraînement moderne ; `bfloat16` préféré sur Ampere+ car plus stable), `int8`/`int4` (quantification — section 162).
2. **device** : un tensor GPU et un tensor CPU ne se mélangent pas — `x.to("cuda")` avant de calculer. L'erreur `Expected all tensors to be on the same device` = 90 % des bugs débutants.
3. **`.item()` / `.detach()`** : pour sortir un scalaire vers Python ou couper le graphe de gradient (utile en debug, jamais en prod chaude).

### 164.3. L'inférence : le pattern à connaître par cœur

```python
import torch
from transformers import AutoModelForCausalLM, AutoTokenizer

model_id = "Qwen/Qwen3-8B"
tok = AutoTokenizer.from_pretrained(model_id)
model = AutoModelForCausalLM.from_pretrained(
    model_id,
    torch_dtype=torch.bfloat16,   # 2 octets/poids : ~16 Go pour un 8B
    device_map="auto",            # répartit sur GPU(s) dispo, sinon CPU
)
model.eval()                      # désactive dropout & co — OBLIGATOIRE en inférence

prompt = "Explique la différence entre kVA et kW pour un onduleur."
inputs = tok(prompt, return_tensors="pt").to(model.device)

with torch.no_grad():             # pas de graphe de gradient : 2 à 3x moins de VRAM
    out = model.generate(**inputs, max_new_tokens=200, do_sample=False)

print(tok.decode(out[0], skip_special_tokens=True))
```

Ce que ce code t'apprend, transposable partout :
- `torch.no_grad()` + `model.eval()` : le duo de l'inférence. Les oublier = VRAM qui explose + résultats non déterministes.
- `torch_dtype=torch.bfloat16` : divise la VRAM par 2 vs float32, sans perte visible sur un LLM.
- `device_map="auto"` (de la lib `accelerate`) : le pont entre « ça tourne sur mon laptop » et « ça tourne sur 2 GPU ».
- `torch.compile(model)` (une ligne, après le chargement) : peut accélérer l'inférence de 10–30 % après un warmup — gratuit, à tester.

Estimation VRAM d'un modèle : `params × octets/param` + KV cache. Un 8B en bfloat16 ≈ 16 Go de poids ; le KV cache ajoute ~1–2 Mo par token de contexte (d'où l'importance de la section 167 sur les tokens).

### 164.4. PyTorch dans ton écosystème : qui utilise quoi

```
PyTorch (tensors, autograd, torch.compile)
 ├── transformers (Hugging Face) : chargement/inférence/fine-tuning — la porte d'entrée
 │    └── TRL : SFTTrainer, GRPOTrainer (utilisé par Unsloth, section 162)
 ├── vLLM / SGLang : serveurs d'inférence (section 163) — kernels CUDA custom AU-DESSUS de torch
 ├── Unsloth : kernels Triton + backprop manuelle au-dessus de torch/TRL
 ├── NVIDIA NIM (section 165) : conteneur qui embarque TensorRT-LLM/vLLM/SGLang compilés
 └── llama.cpp / Ollama / LM Studio : monde C++ à part (GGUF) — PAS du PyTorch
```

Traduction opérationnelle : quand ton RAG rame, la question n'est presque jamais « PyTorch est-il lent ? » mais « quel étage de la pile est mal configuré ? » — dtype, batching, KV cache, quantification. Et quand tu dois debugger un script d'embedding ou de fine-tuning, c'est `torch.cuda.memory_summary()` et `nvidia-smi` qui te disent la vérité, pas les logs.

### 164.5. Debugger la VRAM : les 3 commandes qui disent la vérité

```python
import torch

# 1. Vue d'ensemble : alloué / réservé / libre
print(torch.cuda.memory_summary(device=0, abbreviated=True))

# 2. Le coupable classique : le cache allocator garde de la VRAM "réservée"
#    Si reserved >> allocated : torch.cuda.empty_cache() (ne corrige pas une fuite,
#    mais libère ce que l'allocator retient)
torch.cuda.empty_cache()

# 3. Trouver QUI mange la mémoire : snapshot (à activer AVANT le run)
torch.cuda.memory._record_memory_history(max_entries=100000)
# ... ton code d'inférence ...
torch.cuda.memory._dump_snapshot("vram.pkl")  # à ouvrir dans https://pytorch.org/memory_viz
torch.cuda.memory._record_memory_history(enabled=None)
```

