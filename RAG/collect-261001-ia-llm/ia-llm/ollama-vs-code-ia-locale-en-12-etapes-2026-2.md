---
id: collect-261001-ia-llm/ia-llm/ollama-vs-code-ia-locale-en-12-etapes-2026-2
title: "macOS et Linux"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "DeepSeek", "Mistral", "Z.ai"]
dates: []
keywords: ["agent", "agents", "apache", "benchmark", "deepseek", "distribution", "glm", "llama", "mistral", "mixture of experts", "moe", "open source"]
source: docs/RAG/collect-261001-ia-llm/ollama-vs-code-ia-locale-en-12-etapes-2026.md
source_anchor: ""
source_lines: [40, 128]
sha256: 4325f2960c2f376b638e15f7987e16c79d286c2df30bfc320945f413c20fde6f
---

# macOS et Linux

- Un système d’exploitation à jour : Windows 10/11, macOS 13 ou plus récent, ou une distribution Linux récente comme Ubuntu 22.04+
- Ollama en dernière version stable (v0.33.1, publiée le 26 août 2026, avec un candidat v0.33.2 déjà en test le lendemain selon Local AI Master et PromptQuorum, plusieurs versions après la v0.32.0 du 11 juillet 2026 qui avait transformé la commande `ollama` en un véritable agent CLI interactif)
- Visual Studio Code en version récente, téléchargeable sur code.visualstudio.com
- L’extension Continue, disponible sur le VS Code Marketplace
- Entre 15 et 45 Go d’espace disque libre selon le nombre de modèles installés
- Une connexion internet uniquement nécessaire pour le téléchargement initial des modèles et de l’extension

## Étape 1 : Installer Ollama sur votre système

Sur macOS et Linux, l’installation tient en une seule ligne de commande dans un terminal. Sur Windows, un installeur graphique fait le travail.

```
# macOS et Linux
curl -fsSL https://ollama.com/install.sh | sh
# Windows : téléchargez OllamaSetup.exe depuis ollama.com
# ou, via winget en PowerShell :
winget install Ollama.Ollama
```
Une fois l’installation terminée, vérifiez la version installée :

```
ollama --version
# Sortie attendue, par exemple :
# ollama version is 0.30.10
```
Sur macOS et Windows, l’installeur ajoute aussi une icône dans la barre de menu ou la zone de notification qui garde Ollama actif en arrière-plan. Sur Linux, l’installation configure un service systemd qui démarre automatiquement avec la session. Le projet est entièrement open source et son code source est consultable sur le dépôt GitHub officiel, ce qui permet de vérifier exactement ce que fait le script d’installation avant de l’exécuter, une précaution raisonnable dès qu’une commande est exécutée avec les droits d’un compte utilisateur standard.

## Étape 2 : Démarrer le service et vérifier l’API

Si le service ne démarre pas automatiquement, lancez-le manuellement dans un terminal dédié :

`ollama serve`
Laissez ce terminal ouvert, ou fermez-le si vous avez déjà l’icône de la barre de menu active : les deux méthodes lancent le même serveur, qui écoute par défaut sur le port 11434. Vérifiez que l’API répond correctement depuis un second terminal :

```
curl http://localhost:11434/api/version
# Sortie attendue :
# {"version":"0.30.10"}
```
Si cette commande renvoie une erreur de connexion refusée, le service n’est pas encore démarré ou un pare-feu local bloque le port. La section dépannage plus bas couvre ce cas précis.

## Étape 3 : Choisir le bon modèle de code pour votre matériel

C’est l’étape la plus déterminante pour la qualité de l’expérience finale. Un modèle trop volumineux pour votre VRAM tournera au ralenti ou plantera au chargement. Un modèle trop petit donnera des suggestions moins pertinentes sur du code complexe. Voici les options les plus citées dans les classements et bancs d’essai indépendants publiés en 2026 pour un usage sur Ollama.

| Modèle | Taille | RAM / VRAM conseillée | Repère observé | 
|---|---|---|---|
| qwen2.5-coder:7b | 7B, dense | 8 Go (CPU possible) | Autocomplétion légère, machines modestes, plus de 20 langages supportés | 
| qwen3-coder:30b | 30B MoE, 3,3B actifs | ~19 Go de téléchargement (Q4_K_M), 24 Go VRAM conseillés | Fenêtre de contexte de 256K tokens, bon compromis chat + agent | 
| Devstral Small 2 24B (Mistral) | 24B, dense | 14 Go VRAM | 46,8 % SWE-bench Verified rapporté, conçu pour les agents type Cline, Aider, OpenHands | 
| Qwen3.6-27B | 27B, dense | 24 Go VRAM | 77,2 % SWE-bench Verified rapporté par certains bancs d’essai indépendants pour cette catégorie de taille | 

Notez la différence entre un modèle “dense”, qui active tous ses paramètres à chaque calcul, et un modèle “MoE” (mixture of experts) comme qwen3-coder:30b, qui n’active qu’une fraction de ses paramètres par requête. C’est ce qui permet à un modèle de 30 milliards de paramètres de tourner à une vitesse proche d’un modèle bien plus petit. Ollama gère nativement les deux familles depuis la mise à jour 0.30.10, qui a aussi ajouté le support de Gemma 4 et de Llama 4 aux côtés des modèles Qwen. Le registre s’est encore étoffé en juillet 2026 avec l’arrivée de laguna-xs-2.1, un modèle MoE 33B (3B de paramètres actifs) qui pousse le contexte à 256K tokens et revendique 70,9 % sur le benchmark SWE-bench Verified selon Prompt Quorum, une option à surveiller pour les profils matériel les plus généreux du tableau ci-dessus. Le 9 juillet 2026, Ollama a par ailleurs étoffé son offre cloud avec l’ajout de modèles ouverts plus lourds comme GLM et DeepSeek, selon MentorCruise, une piste pour les développeurs qui veulent dépasser les limites de VRAM locale sans changer d’outil.

Le nom de chaque modèle sur Ollama porte en général un suffixe de quantification, comme `q4_K_M` ou `q8_0`, qui indique sur combien de bits chaque paramètre est représenté en mémoire. Une quantification plus agressive (4 bits) réduit fortement l’espace disque et la VRAM nécessaires, au prix d’une perte de précision généralement faible sur du code classique mais plus sensible sur des raisonnements longs. Une quantification plus légère (8 bits) se rapproche davantage du modèle d’origine mais double presque la mémoire requise. Sans préférence particulière, la variante `q4_K_M` proposée par défaut lors d’un `ollama pull` reste le compromis le plus raisonnable pour démarrer.

## Étape 4 : Télécharger et tester un modèle en ligne de commande

Téléchargez le modèle choisi avec la commande `ollama pull`. Pour ce tutoriel, les exemples utilisent qwen2.5-coder:7b comme base sûre pour la plupart des machines, avec qwen3-coder:30b en option pour le matériel plus généreux.

```
ollama pull qwen2.5-coder:7b
# Sur une machine avec 24 Go de VRAM ou plus :
ollama pull qwen3-coder:30b
```
Une fois le téléchargement terminé, testez directement le modèle dans le terminal, sans passer par VS Code, pour confirmer qu’il répond correctement :

`ollama run qwen2.5-coder:7b "Écris une fonction Python qui calcule la suite de Fibonacci de façon itérative, avec un commentaire de type docstring."`
Exemple de sortie obtenue :

```
def fibonacci(n):
    """Retourne le n-ième terme de la suite de Fibonacci (itératif)."""
    a, b = 0, 1
    for _ in range(n):
        a, b = b, a + b
    return a
```
Si la réponse arrive en plusieurs dizaines de secondes sur un modèle censé être léger, c’est le premier signal que le modèle tourne sur le processeur plutôt que sur la carte graphique. La section optimisation plus loin détaille comment vérifier et corriger ce point.

## Étape 5 : Installer VS Code et l’extension Continue

Si VS Code n’est pas encore installé, téléchargez-le depuis code.visualstudio.com. Installez ensuite l’extension Continue, soit depuis l’onglet Extensions de VS Code en cherchant “Continue”, soit en ligne de commande :

`code --install-extension continue.continue`
Un point de contexte important : Continue a été racheté par Cursor le 16 juin 2026, et sa version 2.0.0, sortie le 19 juin, est la dernière officielle. Le dépôt GitHub est désormais marqué comme non maintenu et en lecture seule. Cela ne bloque en rien l’installation ni l’usage décrit ici : le code reste sous licence Apache 2.0, l’extension continue de fonctionner normalement dans VS Code, et la connexion à un modèle local via Ollama ne dépend d’aucun service cloud de l’éditeur d’origine. Si vous utilisiez la synchronisation cloud de vos configurations Continue, sachez que ces données ont été supprimées après le 15 juillet 2026 : seule la configuration locale, celle que ce tutoriel met en place, reste pertinente désormais.

## Étape 6 : Connecter Continue à Ollama

