---
id: collect-261001-ia-llm/ia-llm/ollama-executer-un-llm-en-local-12-etapes-2026-5
title: "macOS (via Homebrew)"
domain: ia-llm
role: reference
task: reference
actors: ["AMD", "Nvidia", "OpenAI", "vLLM"]
dates: []
keywords: ["chatgpt", "embeddings", "gguf", "gpu", "llama", "llama.cpp", "memory", "nvidia", "open source", "vllm"]
source: docs/RAG/collect-261001-ia-llm/ollama-executer-un-llm-en-local-12-etapes-2026.md
source_anchor: ""
source_lines: [412, 489]
sha256: 377ff4e34c6d391171fa8e316a1c1f94bd3e23fb7e5a1fe80b4d2856de34fc26
---

# macOS (via Homebrew)

Plusieurs variables d’environnement règlent finement le serveur, notamment `OLLAMA_NUM_PARALLEL`, dont l’intérêt a été chiffré par un banc d’essai SitePoint de mars 2026 : sur Llama 3.1 8B, le débit passe d’environ 62 tokens par seconde pour un utilisateur unique à 155 tokens par seconde cumulés avec 50 utilisateurs simultanés. Ces variables se définissent avant le lancement, ou via `systemctl edit ollama` sous Linux.

```
# Garder le modèle en mémoire 30 minutes (au lieu de 5 par défaut)
export OLLAMA_KEEP_ALIVE=30m
# Autoriser plusieurs requêtes simultanées
export OLLAMA_NUM_PARALLEL=4
# Exposer Ollama au réseau local (et non au seul localhost)
export OLLAMA_HOST=0.0.0.0:11434
# Cibler un GPU NVIDIA précis
export CUDA_VISIBLE_DEVICES=0
```
Le deuxième levier est la **quantification**. Elle compresse les poids du modèle, réduisant la mémoire et accélérant l’inférence au prix d’une légère perte de précision. Le tableau ci-dessous résume les compromis courants. `Q4_K_M` est le défaut recommandé : il offre le meilleur équilibre pour la plupart des usages.

| Format | Bits / poids | Taille relative | Qualité | 
|---|---|---|---|
| Q4_K_M | ~4,5 | La plus compacte (défaut) | Très bonne | 
| Q5_K_M | ~5,5 | Intermédiaire | Excellente | 
| Q6_K | ~6,5 | Plus lourde | Quasi optimale | 
| Q8_0 | ~8 | Lourde | Quasi sans perte | 
| FP16 | 16 | Maximale | Référence | 

Enfin, la **fenêtre de contexte** (`num_ctx`) influe directement sur la mémoire consommée. Augmentez-la pour traiter de longs documents, mais sachez qu’un contexte de 32 768 tokens consomme bien plus de VRAM qu’un contexte de 4 096. Ajustez selon votre matériel.

## Ollama vs llama.cpp vs LM Studio : quelle solution choisir ?

Ollama n’est pas seul sur le terrain de l’IA locale. Deux alternatives reviennent souvent pour un usage individuel : **llama.cpp**, le moteur d’inférence bas niveau sur lequel Ollama lui-même s’appuie, et **LM Studio**, une application de bureau avec interface graphique. Pour un déploiement en production à forte charge, un banc d’essai Red Hat d’août 2025 pointait déjà un écart net face à **vLLM** – 41 tokens par seconde de crête contre 793 pour vLLM, avec une latence p99 de 673 ms contre seulement 80 ms –, et les 79 résultats publics agrégés par OpenBenchmarking.org depuis janvier 2026 (dernière mise à jour le 22 mars 2026) confirment qu’Ollama reste avant tout optimisé pour l’usage local mono-utilisateur plutôt que pour le service à grande échelle. Le bon choix dépend donc de votre profil.

| Critère | Ollama | llama.cpp | LM Studio | 
|---|---|---|---|
| Interface | CLI + API | CLI / bibliothèque | Graphique (GUI) | 
| Courbe d’apprentissage | Faible | Élevée | Très faible | 
| Gestion des modèles | Automatique (pull) | Manuelle (GGUF) | Catalogue intégré | 
| API serveur | Oui (REST + OpenAI) | Oui (serveur léger) | Oui (OpenAI) | 
| Licence | MIT (libre) | MIT (libre) | Propriétaire (gratuit) | 
| Public visé | Développeurs, intégration | Experts, embarqué | Débutants, exploration | 

En résumé : choisissez **llama.cpp** si vous voulez le contrôle maximal, l’optimisation fine ou un déploiement embarqué – c’est le moteur, mais il demande de gérer manuellement les fichiers GGUF. Choisissez **LM Studio** si vous préférez cliquer plutôt que taper des commandes et explorer des modèles dans une interface. Choisissez **Ollama** si vous développez des applications : son API, sa gestion automatique des modèles et sa compatibilité OpenAI en font l’option la plus productive pour intégrer un LLM dans un projet.

## Pièges courants à éviter avec Ollama

Voici les erreurs les plus fréquentes que rencontrent les débutants, et comment les contourner avant qu’elles ne vous fassent perdre du temps.

- **Choisir un modèle trop gros pour sa mémoire.** Tenter de lancer un modèle 70B avec 16 Go de RAM provoque un débordement vers le disque (swap) et des réponses désespérément lentes. Respectez le tableau de configuration matérielle et commencez petit.
- **Oublier l’étiquette de taille.**`ollama run qwen3` télécharge la taille par défaut, qui peut être bien plus lourde que prévu. Précisez toujours`:8b` ,`:4b` , etc.
- **Croire qu’Ollama est forcément lent sur CPU.** Sans GPU, privilégiez les modèles 3B à 8B ; un 14B sur CPU sera frustrant. La vitesse dépend surtout de la taille du modèle, pas d’Ollama.
- **Confondre fenêtre de contexte et mémoire du chat.** Un`num_ctx` trop faible tronque silencieusement les longs documents : le modèle « oublie » le début. Augmentez-le pour le RAG, en surveillant la VRAM.
- **Exposer Ollama sans protection.** Définir`OLLAMA_HOST=0.0.0.0` ouvre le serveur à tout le réseau. Ne le faites que derrière un pare-feu ou un reverse proxy authentifié, jamais directement sur Internet.
- **Négliger le format des sorties.** Pour intégrer un LLM dans du code, n’analysez pas du texte libre : utilisez les sorties structurées (étape 9). Vous éviterez des heures de débogage de regex.

## Dépannage : 8 problèmes fréquents et leurs solutions

Même bien installé, Ollama peut présenter des comportements déroutants. Ce tableau couvre les incidents les plus courants signalés par la communauté et la marche à suivre.

| Symptôme | Cause probable | Solution | 
|---|---|---|
| « connection refused » sur le port 11434 | Serveur non démarré | `ollama serve` ou`systemctl start ollama` | 
| Port 11434 déjà utilisé | Conflit avec un autre service | Changer le port via `OLLAMA_HOST=127.0.0.1:11500` | 
| Génération très lente | Exécution sur CPU | Vérifier `ollama ps` ; installer les pilotes GPU | 
| « out of memory » au chargement | Modèle trop grand | Choisir une taille inférieure ou un format plus quantifié | 
| Le modèle « oublie » le contexte | `num_ctx` trop faible | Augmenter le contexte dans le Modelfile | 
| GPU non détecté (Linux) | Pilotes CUDA/ROCm absents | Installer les pilotes puis `systemctl restart ollama` | 
| Réponses incohérentes | Température trop élevée | Baisser `temperature` à 0,3–0,6 | 
| Disque saturé | Modèles accumulés | `ollama rm` ; déplacer`OLLAMA_MODELS` | 

En cas de doute persistant, les journaux du serveur sont votre meilleur allié. Sous Linux, consultez-les avec `journalctl -u ollama -f` ; sous macOS et Windows, l’application de barre des tâches propose un accès aux logs. La plupart des erreurs de chargement y sont explicitées en clair.

## Astuces avancées et bonnes pratiques

Vous maîtrisez les bases ? Voici comment passer au niveau supérieur et intégrer Ollama dans un véritable flux de travail.

- **Ajouter une interface web avec Open WebUI.** Cette interface open source, lançable via Docker, offre une expérience proche de ChatGPT (historique, multi-utilisateurs, gestion des modèles) en se connectant à votre serveur Ollama local.
- **Conteneuriser Ollama.** L’image officielle`ollama/ollama` facilite le déploiement reproductible. Pour les fondamentaux, suivez notre tutoriel Docker avant de monter un service Ollama isolé.
- **Précharger les modèles au démarrage.** Combinez`OLLAMA_KEEP_ALIVE=-1` et un appel initial pour garder un modèle résident en mémoire en permanence, éliminant la latence de premier chargement sur un serveur de production.
- **Utiliser des embeddings de qualité.** Pour un RAG en français, testez`nomic-embed-text` et`mxbai-embed-large` ; le second, plus lourd, capture souvent mieux les nuances sémantiques.
- **Surveiller la consommation.**`ollama ps` et`nvidia-smi` (sur GPU NVIDIA) permettent de suivre l’occupation mémoire en temps réel et d’anticiper les saturations.
- **Versionner vos Modelfiles.** Traitez vos assistants personnalisés comme du code : un dépôt Git de Modelfiles documente vos prompts système et vos paramètres, et rend l’équipe autonome.

