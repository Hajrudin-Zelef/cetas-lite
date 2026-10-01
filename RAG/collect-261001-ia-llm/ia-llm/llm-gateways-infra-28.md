---
id: collect-261001-ia-llm/ia-llm/llm-gateways-infra-28
title: "Passerelles et infra LLM — OpenRouter, Groq, Cerebras, gratuit, self-hosting, TEI, HF, location GPU"
domain: ia-llm
role: reference
task: reference
actors: ["AMD", "Anthropic", "DeepSeek", "Google", "Huawei", "Hugging Face", "Intel", "Nvidia", "OpenAI", "SGLang", "vLLM"]
dates: ["2026-07-16", "2026-09-02", "2026-09-15", "2026-09-27"]
keywords: ["gpu", "amd", "benchmark", "deepseek", "fp8", "gemini", "intel", "kv cache", "llama", "nvidia", "open-weight", "sglang"]
source: docs/RAG/collect-261001-ia-llm/llm_gateways_infra.md
source_anchor: ""
source_lines: [3303, 3381]
sha256: 2b317bb371545178ce9ad39bca90532b9252a71e29ecfc83b5f2e5014d37b30a
---

# Passerelles et infra LLM — OpenRouter, Groq, Cerebras, gratuit, self-hosting, TEI, HF, location GPU

Scénario réaliste : ton routeur LiteLLM pointe `/v1` vers un NIM local pour les modèles open-weight (zéro coût marginal au token, données on-prem), et vers les API cloud (OpenAI/Anthropic/DeepSeek/Doubao) pour le reste. Le jour où tu changes de modèle, tu changes de conteneur — l'API ne bouge pas. Pour ton cas (données techniques internes, contrainte de confidentialité potentielle), c'est l'option « LLM local en prod sans devenir expert CUDA » la plus crédible du marché en sept 2026.

### 165.5. Catalogue et clés : le parcours sans surprise

1. **build.nvidia.com** : le catalogue des NIM — chaque fiche modèle indique les GPU supportées, la précision (FP8/FP16…), le débit attendu. **Vérifie la tienne avant tout** : un NIM Llama-70B exige ~2×H100 ou équivalent, pas ta 4090.
2. **Clé NGC** : gratuite sur org.ngc.nvidia.com (compte NVIDIA) — elle sert à la fois à tirer le conteneur (`docker login nvcr.io`) ET à télécharger les poids au premier démarrage. Mets-la dans un secret (Docker secret, Vault, variable d'env du compose — jamais en clair dans un repo).
3. **Versionnement** : épingle le tag du conteneur (`:1.8.3`, pas `:latest`) — un NIM qui se met à jour tout seul un dimanche soir, c'est une loterie de perf et de comportement.
4. **Licence des poids** : le conteneur est de NVIDIA, mais les poids restent sous la licence du modèle (Llama Community, etc.) — vérifie la compatibilité avec ton usage commercial.

### 165.6. Supervision : les métriques que tu branches à ton monitoring

Chaque NIM expose `/metrics` au format Prometheus — branche-le à ta stack existante (tu as déjà Zabbix/Prometheus/Grafana dans tes guides) :

```yaml
# prometheus.yml — scrape du NIM
scrape_configs:
  - job_name: 'nim-llm'
    static_configs:
      - targets: ['nim-serveur:8000']
    metrics_path: '/metrics'
```

Métriques à alerter : **time-to-first-token** (dégradation = GPU saturée ou KV cache plein), **tokens/s**, **nombre de requêtes en file** (queue depth qui monte = sous-dimensionné), **santé** via `/v1/health/ready` (sonde Kubernetes `readinessProbe` — le pod ne reçoit du trafic que quand le modèle est chargé). Sans ces 4 signaux, tu pilotes un service d'inférence à l'aveugle — inacceptable en prod, même interne.

### 165.7. NIM vs « vLLM à la main » : l'arbitrage honnête

| Critère | NIM | vLLM / SGLang manuel |
|---|---|---|
| Temps jusqu'au 1er token en prod | ~1 h (pull + run) | 1–3 jours (tuning, tests) |
| Perf | Validée par NVIDIA pour ce couple modèle/GPU | À toi de la trouver (flags, quants, batching) |
| Contrôle fin | Faible (profils imposés) | Total |
| Poids du conteneur | Lourd (plusieurs Go + poids) | Léger (image Python + poids HF) |
| Modèles supportés | Catalogue NVIDIA uniquement | Tout Hugging Face |
| GPU non-NVIDIA | Non | Oui (vLLM : AMD/Intel/CPU) |
| Coût d'expérimentation | Faible | Faible aussi |
| Coût de sortie | Changer de moteur = changer de produit | Tu restes maître de ta stack |

Verdict : NIM quand tu veux un **résultat** (une API qui sert vite et bien) ; vLLM/SGLang manuel quand tu veux une **compétence** (comprendre et contrôler toute la chaîne). Les deux ne s'excluent pas : NIM pour le service stable, vLLM en labo pour tester les nouveautés du catalogue HF avant qu'elles n'arrivent en NIM.

---

## 166. NotebookLM (Google) : le RAG-as-a-product — et ce qu'il t'apprend pour ton RAG

**En une phrase :** le carnet de recherche de Google où tu déposes tes sources (PDF, Docs, liens, YouTube) et obtiens des réponses **ancrées avec citations cliquables** — un RAG clé en main, gratuit, excellent comme référence UX… mais qui ne remplace pas ton pipeline.

### 166.1. Présentation (état au 27/09/2026 — vérifié)

- **Nouveau nom depuis le 16/07/2026 : « Gemini Notebook »** — l'ancienne adresse notebooklm.google redirige vers notebook.google. Même produit, rattachement marketing à l'app Gemini.
- Principe : tu crées un carnet, tu ajoutes des sources, tu interroges. Chaque réponse pointe vers des **citations = extraits textuels directs** de tes sources (survol → texte complet, clic → saut au passage). Si la réponse n'est pas dans le corpus, il s'abstient — c'est le comportement que ton RAG devrait viser.
- Nouveautés 2026 : **exécution de code** dans chaque carnet (un ordinateur cloud sécurisé pour analyser tes données) ; **fonctionnalités d'étude** (15/09/2026 : conversations vocales temps réel, enregistreur de cours sur mobile) ; **Audio/Video Overviews** (le « podcast » généré depuis tes sources, 80+ langues), mind maps, flashcards, quiz, rapports, infographies, slide decks.
- Modèle sous le capot : Gemini (fenêtre 1M tokens côté chat du carnet).

### 166.2. Limites et quotas (sept 2026 — chiffres tiers, à vérifier sur l'aide officielle)

Google a remplacé les plafonds journaliers fixes par un **budget de calcul qui se recharge toutes les 5 heures** (02/09/2026) — plus flexible, mais rend les volumes prévisibles difficiles à anticiper. Ordres de grandeur publiés par des tiers (pas de page officielle unique) :

| Élément | Ordre de grandeur |
|---|---|
| Carnets | ~100 (gratuit) à 500 (payant) par utilisateur |
| Sources par carnet | ~50 (gratuit) à 300 (payant) |
| Taille par source | **500 Mo ou 500 000 mots** (seul plafond documenté par Google) |
| Chats | ~500 / jour / utilisateur |
| Audio/Video Overviews, mind maps, rapports | ~20 / jour chacun |
| Infographies | ~3 / jour observés en gratuit (non officiel) |

- **Tarif** : gratuit pour un compte Google standard (« free forever ») ; quotas élevés via Google One AI Premium (~$19,99/mois) ou Workspace/Éducation.
- **Confidentialité** : pour les comptes Workspace, Google indique que les uploads, requêtes et réponses ne sont ni relus par des humains ni utilisés pour entraîner les modèles. En compte grand public gratuit, pars du principe inverse pour des documents sensibles — **ne dépose jamais de doc interne confidentiel** sans validation de ta politique de sécurité.

### 166.3. Lien direct avec ton RAG : l'utiliser comme banc d'essai

NotebookLM est un excellent **étalon** pour ton pipeline, pas un remplaçant :

1. **Référence UX** : les citations inline avec extraits + abstention hors-corpus, c'est le gold standard. Si tes réponses RAG n'atteignent pas ce niveau de traçabilité, c'est ta spec d'amélioration.
2. **Benchmark qualité** : dépose le même corpus (ex. 20 PDF de docs Huawei/onduleurs) dans NotebookLM et dans ton RAG, pose les mêmes 30 questions d'éval (section 127), compare. Ça te donne un plafond « RAG industriel gratuit » à battre.
3. **Prototypage rapide** : avant d'ingérer un nouveau corpus dans ton pipeline, teste en 10 minutes dans NotebookLM si le corpus répond bien aux questions — ça évite d'indexer des documents inutiles.
4. **Idées de features** : Audio Overview = une façon de « consommer » tes guides ; les mind maps = une visualisation de la structure de ton corpus.

### 166.4. Limites honnêtes — pourquoi tu gardes ton RAG

