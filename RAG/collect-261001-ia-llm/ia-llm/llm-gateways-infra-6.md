---
id: collect-261001-ia-llm/ia-llm/llm-gateways-infra-6
title: "Passerelles et infra LLM — OpenRouter, Groq, Cerebras, gratuit, self-hosting, TEI, HF, location GPU"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "DeepSeek", "Google", "Groq", "Nvidia", "OpenAI", "OpenRouter", "vLLM"]
dates: ["2026-09-27"]
keywords: ["gpu", "agent", "agents", "attention", "awq", "claude", "deepseek", "embeddings", "gpus", "llama", "llama.cpp", "memory"]
source: docs/RAG/collect-261001-ia-llm/llm_gateways_infra.md
source_anchor: ""
source_lines: [568, 759]
sha256: 36ac9225e0eec9851b775e2da4db68edc0bc8f0e8b445822f5d430e81fed5143
---

# Passerelles et infra LLM — OpenRouter, Groq, Cerebras, gratuit, self-hosting, TEI, HF, location GPU

- **Repo** : `github.com/deepseek-ai/deepseek-harness` — **réel, vérifié le 27/09/2026**.
- **Licence** : MIT. **Statut** : *developer preview* (changements cassants possibles — ne pas baser une prod dessus aujourd'hui).
- **Lancement** : 13 août 2026, en même temps que DeepSeek V4 Pro. ~95k stars en 2 jours, > 225k mi-septembre 2026 (d'après la presse tech — ordre de grandeur).
- **Nature** : gros monorepo **TypeScript**, architecture **« everything is a plugin »** (tout est un plugin : UI, providers, outils), motorisé par **Cordis** (framework de composabilité).
- **Ce n'est PAS un hébergeur de modèle** : c'est un **agent de code** (concurrent d'OpenCode/Claude Code) avec UI web locale, runner headless, sandbox d'exécution, sous-agents, recherche web, plugins.
- **Multi-provider** : route DeepSeek native + adaptateur multi-fournisseurs (compatible avec d'autres APIs/gateways) — tu peux le brancher sur OpenRouter, Groq, ou ton Ollama local.

## 47. DeepSeek Harness : installation et démarrage

Prérequis : **Node.js 22.19+ ou ≥ 24** (la 23.x n'est **pas** supportée — dépend de `node:sqlite` et du type-stripping natif).

```bash
# Option 1 : sans installation (recommandé pour tester)
npx @deepseek-ai/dsh web

# Option 2 : global
npm install -g @deepseek-ai/dsh
dsh web

# Option 3 : depuis les sources
git clone https://github.com/deepseek-ai/deepseek-harness.git
cd deepseek-harness
pnpm install && pnpm run build
pnpm dsh web
```

- L'UI web démarre sur **`http://127.0.0.1:3080`** (s'ouvre dans le navigateur).
- Config/API key : dans les réglages de l'UI (clé DeepSeek ou autre provider).
- Données locales : `~/.dsh` (plugins, profils, historique de sessions).
- **Sécurité** : l'agent tourne avec **tous tes droits utilisateur** (lecture fichiers, exécution commandes) → lance-le dans un dossier poubelle d'abord, lis `SAFETY.md` du repo.

## 48. DeepSeek Harness : usages et service systemd

- **Mode web** : `dsh web` — assistant de code dans le navigateur, sous-agents, outils terminaux.
- **Mode headless** : une tâche sans UI (à brancher sur CI/automation — vérifier la syntaxe exacte dans la doc, elle bouge en preview).
- **Plugins** : topic GitHub `dsh-plugin` ; extension VS Code communautaire existante (`deepseek-harness-for-vscode`).
- Service utilisateur systemd (toujours actif) :

```ini
# ~/.config/systemd/user/dsh-web.service
[Unit]
Description=DeepSeek Harness Web UI
Wants=network-online.target
After=network-online.target

[Service]
Type=simple
ExecStart=/usr/bin/dsh web --no-open
Restart=always
RestartSec=3

[Install]
WantedBy=default.target
```

```bash
systemctl --user daemon-reload
systemctl --user enable --now dsh-web.service
loginctl enable-linger $USER   # optionnel : actif même sans session ouverte
```

## 49. DeepSeek Harness vs alternatives (sept 2026)

| Outil | Éditeur | Langage | Prix | Point fort |
|---|---|---|---|---|
| DeepSeek Harness (`dsh`) | DeepSeek (MIT) | TypeScript | Gratuit (+ coût API du modèle) | Architecture tout-plugin, multi-provider |
| OpenCode | SST (MIT) | TypeScript/Go | Gratuit (+ API) | Léger, terminal |
| Claude Code | Anthropic | — | Abonnement/API | Intégration Claude |
| Codex CLI | OpenAI | TypeScript | Abonnement/API | Intégration GPT |

Pour toi : `dsh` est intéressant comme **labo d'agents** (comprendre l'architecture d'un harness réel), pas comme choix de prod aujourd'hui (preview). Le vrai gain, c'est d'étudier son découpage plugin pour ton propre scaffolding RAG.

## 50. Self-hosting : panorama des options

| Outil | Idéal pour | GPU requis | API OpenAI-compatible |
|---|---|---|---|
| **Ollama** | Démarrer vite, dev, laptop | Non (CPU ok) | Oui (`:11434/v1`) |
| **llama.cpp** (`llama-server`) | CPU, embarqué, contrôle fin | Non | Oui |
| **vLLM** | Servir en prod, débit | Oui (NVIDIA) | Oui (`:8000/v1`) |
| **TEI** | Embeddings uniquement | Optionnel | Oui (`/v1/embeddings`) |
| **TGI** | Alternative HF au serving | Oui | Partiel |

Règle : **Ollama pour expérimenter**, **vLLM pour servir** (débit, batching), **llama.cpp quand pas de GPU NVIDIA**.

## 51. Ollama : installation

```bash
# Linux (script officiel)
curl -fsSL https://ollama.com/install.sh | sh

# Vérifier
ollama --version
systemctl status ollama   # service actif par défaut

# macOS / Windows : installateur sur ollama.com
# Docker :
docker run -d --gpus=all -v ollama:/root/.ollama -p 11434:11434 \
  --name ollama --restart unless-stopped ollama/ollama
```

Par défaut Ollama écoute sur `127.0.0.1:11434`. Pour l'exposer sur le LAN : `OLLAMA_HOST=0.0.0.0:11434` (variable d'environnement du service — attention, **pas d'auth native** : ne l'expose jamais sur Internet sans reverse-proxy + auth).

## 52. Ollama : modèles DeepSeek et autres (commandes réelles)

```bash
# DeepSeek R1 distillés (raisonnement) — tailles vérifiées sur le catalogue Ollama
ollama run deepseek-r1:1.5b    # ~1,1 Go, CPU ok
ollama run deepseek-r1:7b      # ~4,7 Go
ollama run deepseek-r1:8b      # ~4,9 Go
ollama run deepseek-r1:14b     # ~9 Go
ollama run deepseek-r1:32b     # ~20 Go
ollama run deepseek-r1:70b     # ~43 Go

# Autres utiles pour un sysadmin
ollama run qwen3:32b           # bon équilibre FR/code (vérifier le tag exact : ollama.com/library/qwen3)
ollama run gemma3:27b          # Google, multilingue
ollama run mistral:7b          # léger
ollama run nomic-embed-text    # embeddings locaux 768 dim (alternative à TEI, simple)

# Lister / supprimer
ollama list
ollama rm deepseek-r1:1.5b
```

Note : `deepseek-r1:671b` (le modèle complet) existe au catalogue mais demande un cluster (~400 Go+ de VRAM) — pas pour un poste. Les distillés 7b→32b sont le sweet spot local.

## 53. Ollama : API locale et Python

```bash
# Génération (API native)
curl http://localhost:11434/api/generate -d '{
  "model": "qwen3:32b",
  "prompt": "Explique la différence entre LACP actif et passif.",
  "stream": false
}'

# Chat OpenAI-compatible (même code que Groq/OpenAI, autre base_url)
curl http://localhost:11434/v1/chat/completions \
  -H "Content-Type: application/json" \
  -d '{"model": "qwen3:32b", "messages": [{"role": "user", "content": "Bonjour"}]}'
```

```python
from openai import OpenAI
client = OpenAI(base_url="http://localhost:11434/v1", api_key="ollama")  # clé ignorée en local
r = client.chat.completions.create(model="qwen3:32b",
    messages=[{"role": "user", "content": "Commande VRP pour un trunk ?"}])
print(r.choices[0].message.content)
```

→ **Ton RAG peut pointer sa variable `LLM_BASE_URL` vers Ollama** : 0 €, 0 donnée sortante, même code que le cloud.

## 54. Ollama : Modelfile (personnaliser)

```dockerfile
# Modelfile
FROM qwen3:32b
PARAMETER temperature 0.2
PARAMETER num_ctx 16384
SYSTEM """Tu es un assistant expert réseau et systèmes pour un chef de service.
Réponds en français, direct, avec les commandes exactes. Si tu ne sais pas, dis-le."""
```

```bash
ollama create reseau-assistant -f Modelfile
ollama run reseau-assistant
```

`num_ctx` règle la fenêtre de contexte (défaut souvent 4K–8K selon le modèle) — à monter pour du RAG (16K–32K si la VRAM suit).

## 55. vLLM : installation et serving OpenAI-compatible

vLLM = le serveur de prod (PagedAttention, batching continu, débit). NVIDIA recommandé.

```bash
# Installer (dans un venv ou conteneur CUDA)
pip install vllm

# Servir un modèle (téléchargé depuis le Hub HF)
vllm serve Qwen/Qwen3-32B --port 8000 --gpu-memory-utilization 0.9

# Avec quantization AWQ (4-bit, 2× moins de VRAM)
vllm serve Qwen/Qwen3-32B-AWQ --quantization awq --port 8000

# Multi-GPU (70B sur 2× RTX 4090 par ex.)
vllm serve meta-llama/Llama-3.3-70B-Instruct \
  --tensor-parallel-size 2 \
  --gpu-memory-utilization 0.92 \
  --max-model-len 32768 \
  --port 8000
```

