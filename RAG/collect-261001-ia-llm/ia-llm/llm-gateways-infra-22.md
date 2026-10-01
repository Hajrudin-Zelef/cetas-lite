---
id: collect-261001-ia-llm/ia-llm/llm-gateways-infra-22
title: "Passerelles et infra LLM — OpenRouter, Groq, Cerebras, gratuit, self-hosting, TEI, HF, location GPU"
domain: ia-llm
role: reference
task: reference
actors: ["AMD", "Alibaba", "Apple", "ByteDance", "DeepSeek", "Hugging Face", "Intel", "Nvidia", "OpenAI", "SGLang", "vLLM"]
dates: ["2026-06-04"]
keywords: ["gpu", "agents", "amd", "benchmark", "cost", "deepseek", "gguf", "gpt-6", "intel", "llama", "llama.cpp", "nvidia"]
source: docs/RAG/collect-261001-ia-llm/llm_gateways_infra.md
source_anchor: ""
source_lines: [2587, 2684]
sha256: 78dcd4adf2c785c3baaf2b4830c2c8f31414a75f8edd4515996ce9d829a34a05
---

# Passerelles et infra LLM — OpenRouter, Groq, Cerebras, gratuit, self-hosting, TEI, HF, location GPU

Lecture : DeepSeek reste imbattable au prix ; Doubao se place entre DeepSeek et les occidentaux avec un argument « agents + cache » ; Sol est le joker si son prix publié confirme la promesse. **Aucun ne gagne sans passer tes 30 questions d'éval en français** — le prix ne dit rien de la qualité sur tes procédures d'onduleurs.

### 160.8. Pièges opérationnels avant la prod

1. **Facturation CNY** : ta compta veut de l'EUR/USD — prévois la conversion et la TVA applicable aux services cloud chinois ; un agrégateur USD simplifie la facture mais ajoute sa marge.
2. **Souveraineté des données** : prompts envoyés sur des serveurs en Chine (région `cn-beijing` par défaut). Pour des docs d'exploitation internes, valide avec ta politique de sécurité — sinon, réserve Doubao aux usages non sensibles.
3. **Endpoint ≠ modèle** : côté Ark, tu déploies un modèle sur un **endpoint** (avec son propre ID et son propre débit alloué). En pic, c'est le débit de TON endpoint qui limite, pas le modèle — dimensionne-le comme tu dimensionnerais une VM.
4. **Dérive des IDs** : ByteDance a retiré Seed 2.0 Pro du flagship en 3 mois (juin 2026). Épingle des IDs datés et garde la section 159 à jour — un alias qui change de modèle sous-jacent, tu as déjà vu le cas DeepSeek le 14/09.

### 160.9. Intégrer Doubao dans ton routeur LiteLLM

```yaml
# litellm config — Doubao via Ark (endpoint OpenAI-compatible)
model_list:
  - model_name: doubao-seed-2-1-pro
    litellm_params:
      model: openai/doubao-seed-2-1-pro
      api_base: https://ark.cn-beijing.volces.com/api/v3
      api_key: os.environ/VOLC_ARK_API_KEY   # clé Ark fictive en exemple
  - model_name: doubao-seed-2-1-turbo
    litellm_params:
      model: openai/doubao-seed-2-1-turbo
      api_base: https://ark.cn-beijing.volces.com/api/v3
      api_key: os.environ/VOLC_ARK_API_KEY

router_settings:
  routing_strategy: cost-based  # Turbo pour le volume, Pro pour les agents
```

Stratégie de routage suggérée : questions RAG simples → Turbo ; tâches agentiques multi-étapes ou code → Pro ; et un fallback `doubao-seed-2-1-pro → deepseek-v4-pro → gpt-6-sol` pour la résilience. Mesure le TTFB réel depuis ton infra avant de mettre Doubao sur un chemin synchrone exposé aux utilisateurs — la latence réseau Chine→Europe peut annuler l'avantage prix sur un chat interactif.

### 160.10. Veille prix : le rituel mensuel (car ça bouge vite)

Doubao a augmenté ses prix flagship entre Seed 2.0 et 2.1 (¥3,20 → ¥6 l'entrée) tout en divisant par deux via Turbo — la gamme se repositionne chaque semestre. Rituel à caler sur ta relecture mensuelle (section 159.17) :

1. Ouvre la page pricing Volcano (docs.volcengine.com/docs/82379/1544106) — note les prix Pro/Turbo/cache.
2. Compare avec ton coût réel du mois (tes logs LiteLLM : tokens × prix).
3. Si l'écart avec DeepSeek/GPT-6 Sol dépasse 30 % sur ton mix réel entrée/sortie, relance un benchmark d'éval.
4. Mets à jour le tableau 160.3 avec la date — un prix sans date dans ce guide est un prix périmé.

---

## 161. LM Studio : l'alternative desktop à Ollama (avec GUI)

**En une phrase :** le même moteur que tout le monde (llama.cpp) mais avec une vraie application desktop : navigateur de modèles Hugging Face intégré, historique de chat persistant, comparaisons côte à côte, serveur OpenAI-compatible en un clic — gratuit, propriétaire, y compris pour usage commercial.

### 161.1. Ce que c'est (état : v0.4.16, juin 2026)

- App native **Windows 10+, macOS (Apple Silicon et Intel), Linux (AppImage)**. Sous le capot : llama.cpp, accélération GPU (NVIDIA CUDA, Apple Metal/MLX, AMD Vulkan/ROCm).
- Installation : téléchargeur sur lmstudio.ai → installe → l'onboarding te propose un modèle recommandé. **Temps jusqu'à la première inférence : ~10–15 min** (install + téléchargement d'un modèle 8B).
- Licence : propriétaire, **gratuite pour tout usage y compris commercial** (l'obligation de licence business a été supprimée en 2025). Ce n'est pas open source — à noter si ta politique interne l'exige.
- Nouveautés 2026 notables : **v0.4.14** — décodage spéculatif MTP (Multi-Token Prediction) ; **v0.4.15** — parallélisme tensoriel multi-GPU en un clic ; **v0.4.16 (04/06/2026)** — app iPhone/iPad **Locally** + **LM Link** : ton téléphone devient un terminal chiffré (mesh Tailscale, E2E) vers les modèles de ton Mac — rien ne transite par un relais cloud, seul un annuaire de découverte touche les serveurs LM Studio. Lancement iPhone/iPad uniquement.
- Panneau **Mission Control** (2026) : gestion headless — les modèles se chargent à la demande sans UI, utile pour un poste de dev qui sert aussi de backend local.

### 161.2. Installation et usage

```text
1. Télécharge depuis lmstudio.ai (version du site officiel uniquement)
2. Installe (Windows : Setup.exe / macOS : glisser dans Applications / Linux : AppImage)
3. Onglet Search : cherche un modèle (ex. "Qwen3"), filtre par ton matériel —
   LM Studio détecte ta RAM/VRAM et indique quelles quantifications passent
4. Clique Download (un 8B en Q4 ≈ 5 Go, compte ~10 min)
5. Onglet Chat : sélectionne le modèle dans la liste, discute
```

**Mode serveur** (le pont vers tes outils) : onglet *Local Server* → *Start* → API OpenAI-compatible sur `http://localhost:1234`. N'importe quel outil qui accepte `OPENAI_BASE_URL` s'y branche :

```bash
# Exemple : pointer un client OpenAI vers LM Studio (clé fictive, valeur ignorée en local)
export OPENAI_BASE_URL="http://localhost:1234/v1"
export OPENAI_API_KEY="lm-studio"
```

Cas d'usage concrets : comparer plusieurs quantifications d'un même modèle (Q4 vs Q5 vs Q8) pour ton matériel ; tester un modèle vision (LLaVA, Qwen-VL) en glissant une image dans le chat ; servir de backend local à Continue.dev / Open WebUI sans toucher au terminal.

### 161.3. LM Studio vs Ollama — tableau comparatif (vérifié sept 2026)

| Critère | LM Studio 0.4.16 | Ollama 0.24.0 |
|---|---|---|
| Interface | GUI native complète (chat, sliders température/top-p, comparaisons) | CLI (`ollama run`) ; GUI via tiers (Open WebUI) |
| Historique de chat | Oui, persistant | Non en natif |
| Découverte de modèles | Navigateur Hugging Face intégré + indicateur compatibilité matérielle | `ollama pull` + catalogue ollama.com (~100+ modèles) |
| Formats | GGUF (tout Hugging Face) | GGUF via registre Ollama (curé) |
| API REST | OpenAI-compatible, port **1234** (l'app doit tourner) | OpenAI-compatible, port **11434** (daemon système) |
| Headless / serveur | Non — desktop requis (Mission Control = semi-headless) | Oui — daemon, idéal serveurs |
| Multi-GPU | Parallélisme tensoriel en 1 clic (0.4.15) | Oui, configurable |
| Mobile | Locally (iOS) + LM Link via Tailscale | Non (tiers uniquement) |
| macOS Apple Silicon | Metal + MLX | Metal + MLX (depuis 0.24.0, plus rapide) |
| Mises à jour | Auto in-app | Auto en arrière-plan |
| Licence | Propriétaire, gratuit (y c. commercial) | MIT (open source) |
| Temps 1re inférence | ~12–15 min | ~5 min |

### 161.4. Doctrine d'usage pour ton labo

- **Les deux cohabitent** : ports différents (11434 vs 1234), aucun conflit. Ollama = backend API/headless de référence ; LM Studio = l'outil d'évaluation interactive (« ce nouveau modèle vaut-il quelque chose ? » en 5 minutes, avec historique et comparaison côte à côte).
- Pour Zelef : LM Studio est ce que tu donnes à un collègue non-technique qui veut « essayer un modèle local » sans installer Python ni apprendre ce que veut dire GGUF. Ollama reste ton choix pour tout ce qui est scripté, conteneurisé ou servi (RAG local, Vast.ai, RunPod).
- Limite honnête : LM Studio ne remplace ni vLLM ni SGLang (section 163) pour servir en production — c'est un outil desktop, pas un serveur d'inférence à haut débit.

