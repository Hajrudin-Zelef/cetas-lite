---
id: collect-261001-ia-llm/ia-llm/tutoriel-ollama-2026-llm-local-en-13-etapes-7
title: "Vérifier la version du pilote NVIDIA"
domain: ia-llm
role: reference
task: reference
actors: ["AMD", "Alibaba", "Meta", "Mistral", "Nvidia", "OpenAI"]
dates: []
keywords: ["nvidia", "agents", "apache", "arr", "attention", "benchmarks", "gguf", "gpu", "license", "llama", "llama.cpp", "lora"]
source: docs/RAG/collect-261001-ia-llm/tutoriel-ollama-2026-llm-local-en-13-etapes.md
source_anchor: ""
source_lines: [579, 670]
sha256: 84155117171ec6c3066ebfc40fa9b0426fc58523dc15ff1d78efc9e0991f0e54
---

# Vérifier la version du pilote NVIDIA

| Symptôme | Cause probable | Solution | 
|---|---|---|
| « connection refused » sur :11434 | Serveur non démarré | `systemctl start ollama` ou`ollama serve` | 
| « CUDA out of memory » | Modèle trop gros pour VRAM | Quantification inférieure (Q4 → Q3) ou `num_gpu` réduit | 
| Génération < 5 tok/s sur RTX | Modèle sur CPU (partial offload) | Vérifier nvidia-smi, augmenter VRAM libre | 
| Pull s’arrête à 99 % | Hash mismatch, registry timeout | `ollama rm` puis`ollama pull` à nouveau | 
| « model requires more system memory » | RAM système insuffisante | Activer swap 16 Go ou quantif. plus légère | 
| Réponse tronquée à 200 tokens | `num_predict` par défaut bas | Ajouter `"options": {"num_predict": 2048}` | 
| Contexte oublié après ~2 000 tokens | `num_ctx` par défaut 2048 | Définir `OLLAMA_CONTEXT_LENGTH=8192` | 
| Service ne démarre pas au boot | systemd non activé | `sudo systemctl enable ollama` | 

Pour les erreurs GPU plus subtiles, l’outil `OLLAMA_DEBUG=1 ollama serve` active des logs verbeux qui montrent exactement combien de couches sont chargées en VRAM versus RAM. Si vous voyez « offload to GPU : 28/33 layers », cela signifie que 5 couches sont restées en RAM système, ralentissant l’inférence. Solution : libérer 1 à 2 Go de VRAM (fermer le navigateur, désactiver l’accélération desktop) ou descendre à une quantification plus légère.

## Astuces avancées : LoRA, function calling, vision

Trois fonctionnalités avancées méritent d’être mentionnées pour les utilisateurs qui dépassent le cas d’usage de base. D’abord les **adaptateurs LoRA** : depuis la v0.4, Ollama peut charger un LoRA fine-tuné par-dessus un modèle de base via `ADAPTER ./mon-lora.gguf` dans le Modelfile. Cela permet de spécialiser un modèle 8B sur votre domaine (jargon médical, juridique, technique interne) sans re-télécharger les 4,7 Go du modèle complet : seul le LoRA de 50 à 200 Mo est ajouté.

Ensuite le **function calling** (tool use) : les modèles compatibles (Llama 3.1+, Qwen 2.5+, Mistral Nemo) acceptent un schéma d’outils JSON dans le prompt, et renvoient des appels de fonction structurés. C’est la brique de base pour construire des agents qui interagissent avec votre SI (consultation CRM, envoi d’emails, requêtes SQL). L’endpoint `/api/chat` accepte le champ `tools` au format OpenAI standard.

```
# Function calling avec /api/chat
curl http://localhost:11434/api/chat -d '{
  "model": "llama3.1:8b",
  "messages": [
    {"role": "user", "content": "Quelle est la météo à Paris ?"}
  ],
  "tools": [{
    "type": "function",
    "function": {
      "name": "get_weather",
      "description": "Obtenir la météo actuelle d une ville",
      "parameters": {
        "type": "object",
        "properties": {
          "city": {"type": "string", "description": "Nom de la ville"}
        },
        "required": ["city"]
      }
    }
  }],
  "stream": false
}'
# Réponse contient :
# "tool_calls": [{
#   "function": {"name": "get_weather", "arguments": {"city": "Paris"}}
# }]
```
Enfin la **vision multimodale** : depuis la v0.4, Ollama supporte les modèles vision comme `llava:13b`, `llama3.2-vision:11b` ou `qwen2-vl:7b`. Vous pouvez envoyer une image en base64 dans le champ `images` de la requête. Cas d’usage typiques : OCR de documents, analyse de captures d’écran pour automatisation RPA, description d’images pour accessibilité. Attention, ces modèles consomment 2 à 4 Go de VRAM supplémentaires pour l’encodeur visuel.

## Benchmarks 2026 : tokens/seconde par GPU

Pour vous aider à choisir votre matériel, voici les débits mesurés sur Ollama v0.6.2 avec Llama 3.1 8B Q4_K_M, prompt de 512 tokens, génération de 256 tokens, batch size 1. Ces chiffres proviennent de benchmarks reproductibles publiés par la communauté en avril 2026. Notez que les performances dépendent fortement du driver GPU, du contexte effectif et de la température ambiante (throttling).

| Matériel | VRAM/RAM | Llama 3.1 8B Q4 | Llama 3.3 70B Q4 | Prix indicatif | 
|---|---|---|---|---|
| RTX 4090 24 Go | 24 Go | 105 t/s | N/A (insuffisant) | 1 900 € | 
| RTX 4080 Super 16 Go | 16 Go | 88 t/s | N/A | 1 100 € | 
| RTX 4070 Ti Super 16 Go | 16 Go | 82 t/s | N/A | 850 € | 
| RTX 4060 Ti 16 Go | 16 Go | 52 t/s | N/A | 500 € | 
| RX 7900 XTX 24 Go (ROCm 6.3) | 24 Go | 78 t/s | N/A | 1 000 € | 
| MacBook Pro M3 Max 128 Go | 128 Go unifié | 62 t/s | 12 t/s | 4 200 € | 
| MacBook Pro M3 Pro 36 Go | 36 Go unifié | 48 t/s | N/A | 2 600 € | 
| Mac Mini M4 Pro 64 Go | 64 Go unifié | 55 t/s | 9 t/s | 2 000 € | 
| Ryzen 9 7950X (CPU only) | 64 Go DDR5 | 11 t/s | 1,8 t/s | 1 500 € (config) | 
| 2× RTX 4090 (NVLink off) | 48 Go | 180 t/s parallel | 22 t/s | 3 800 € (GPU seuls) | 

Pour un développeur solo, la RTX 4070 Ti Super 16 Go représente le meilleur rapport performance/prix en 2026 : suffisamment de VRAM pour Llama 3.1 8B avec contexte 32k, environ 82 tokens/seconde, et un prix sous les 900 €. Pour une petite équipe, le Mac Mini M4 Pro 64 Go séduit par sa mémoire unifiée qui permet de faire tourner des modèles 32B impossibles sur la plupart des cartes grand public. Pour les budgets serrés, une RTX 4060 Ti 16 Go reste très honnête à 500 €.

## Foire aux questions

### Ollama est-il vraiment gratuit ?

Oui, Ollama est entièrement gratuit et open source sous licence MIT. Le runtime, l’API, le client CLI et l’application desktop sont gratuits. Les modèles eux-mêmes ont leurs propres licences : Llama 3.x est sous licence Meta Community License (gratuit jusqu’à 700 millions d’utilisateurs actifs), Mistral 7B sous Apache 2.0, Qwen sous Tongyi Qianwen License. Vérifiez la licence du modèle si vous comptez l’utiliser commercialement.

### Peut-on faire tourner Llama 3.3 70B sur un MacBook ?

Oui, sur un MacBook Pro M3 Max ou M4 Max avec au moins 64 Go de mémoire unifiée. Le modèle 70B en Q4_K_M occupe environ 40 Go. Le débit observé est de 9 à 12 tokens/seconde, ce qui reste utilisable pour des prompts courts. Sur un Mac Mini M4 Pro 64 Go, c’est aussi possible avec environ 9 tokens/seconde. En dessous de 48 Go de RAM unifiée, Llama 3.3 70B n’est pas réaliste.

### Quelle différence entre Ollama et LM Studio ?

Ollama est orienté CLI et API HTTP, conçu pour s’intégrer dans des scripts, applications et services. LM Studio est une application desktop graphique avec interface chat intégrée, idéale pour les utilisateurs non-développeurs. Les deux utilisent llama.cpp sous le capot et supportent le format GGUF. Pour un développeur qui construit une intégration, Ollama est le bon choix. Pour un usage personnel ponctuel, LM Studio est plus convivial.

### Ollama envoie-t-il des données vers le cloud ?

Non. Une fois un modèle téléchargé via `ollama pull`, toute l’inférence reste 100 % locale sur votre machine. Aucune télémétrie n’est envoyée par défaut. Les seules connexions externes sont : le téléchargement initial des modèles depuis ollama.com, et la vérification optionnelle des mises à jour du binaire. Vous pouvez désactiver cette dernière avec `OLLAMA_NO_AUTOUPDATE=1`.

### Quel modèle choisir pour le français en 2026 ?

Pour un usage généraliste en français, trois modèles se distinguent : **Mistral 7B Instruct v0.3** (excellent niveau de français, contexte 32k, créé en France), **Llama 3.1 8B** (très bon français, contexte 128k, large écosystème), et **Qwen 3 14B** (raisonnement supérieur, multilingue solide). Pour les modèles plus petits (1-3B), `gemma3:4b` reste le plus convaincant en français.

### Comment fine-tuner un modèle avec Ollama ?

