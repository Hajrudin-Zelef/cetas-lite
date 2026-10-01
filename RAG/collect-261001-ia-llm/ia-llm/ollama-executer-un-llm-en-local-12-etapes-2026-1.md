---
id: collect-261001-ia-llm/ia-llm/ollama-executer-un-llm-en-local-12-etapes-2026-1
title: "macOS (via Homebrew)"
domain: ia-llm
role: reference
task: reference
actors: ["AMD", "Alibaba", "Anthropic", "Apple", "DeepSeek", "Google", "Mistral", "Nvidia", "OpenAI"]
dates: []
keywords: ["agent", "amd", "benchmark", "chatgpt", "claude", "deepseek", "distribution", "gemini", "gpu", "llama", "llama.cpp", "mai"]
source: docs/RAG/collect-261001-ia-llm/ollama-executer-un-llm-en-local-12-etapes-2026.md
source_anchor: ""
source_lines: [1, 50]
sha256: a8ec62f53033fe3d0196ebc54f894ce3cee25546d1e1f68cca96408ebfb20f8e
---

# macOS (via Homebrew)

Exécuter un grand modèle de langage (LLM) sur votre propre machine, sans abonnement, sans clé d’API et sans envoyer la moindre donnée à un serveur américain : c’est exactement ce que permet **Ollama**. En quelques minutes, vous lancez Llama 3, DeepSeek-R1 ou Qwen 3 directement sur votre ordinateur. Ce tutoriel Ollama vous guide pas à pas, de l’installation jusqu’à un projet RAG complet, en 12 étapes et environ 30 minutes.

Au 17 août 2026, Ollama est devenu la passerelle de référence vers l’IA locale. Le projet, écrit en Go et publié sous licence MIT, dépasse les **175 000 étoiles sur GitHub**, revendique **8,9 millions de développeurs actifs mensuels** – dont **85 % des entreprises du Fortune 500**, selon TechFundingNews – et plus de **67 000 intégrations communautaires** selon les chiffres publiés début juillet 2026, et a levé **88 millions de dollars** au total, dont un tour de série B de **65 millions de dollars** mené par Theory Ventures le 9 juillet 2026, aux côtés de Benchmark, 8VC et Y Combinator. Cette base d’utilisateurs a doublé depuis janvier 2026, avec près d’un million de nouvelles installations chaque semaine mi-2026, et le projet a livré sa dernière version stable, la **0.34.1**, publiée le 14 septembre 2026 sur GitHub. Construit au-dessus du moteur d’inférence llama.cpp, il masque toute la complexité de la quantification et de l’accélération GPU derrière une seule commande. Le résultat : `ollama run llama3.2` suffit à dialoguer avec une IA.

Pour la France et l’Europe, l’enjeu dépasse la simple commodité. Exécuter un **LLM en local** signifie que vos prompts, vos documents et vos données clients ne quittent jamais votre poste. Aucun transfert hors UE, aucune dépendance à un fournisseur soumis au *Cloud Act* : c’est la voie la plus directe vers une IA conforme au RGPD et alignée sur la souveraineté numérique. Voici ce que vous allez apprendre.

- Installer Ollama sur macOS, Linux et Windows et vérifier l’installation ;
- Télécharger, lancer et gérer des modèles open source (Llama, DeepSeek, Qwen, Gemma, gpt-oss) ;
- Interroger l’API REST locale (port 11434) et l’API compatible OpenAI ;
- Intégrer Ollama en Python, créer un *Modelfile* et générer des sorties JSON structurées ;
- Activer l’appel d’outils (tool calling), les modèles de raisonnement et Ollama Cloud ;
- Construire un assistant RAG complet qui répond sur vos documents locaux.

## Qu’est-ce qu’Ollama et pourquoi exécuter un LLM en local en 2026 ?

Ollama est un logiciel libre qui télécharge, gère et sert des modèles de langage en local via une interface en ligne de commande et une API HTTP. Pensez-y comme au « Docker des LLM » : une commande `pull` pour récupérer un modèle, une commande `run` pour l’exécuter, et un serveur qui écoute en arrière-plan. Depuis la version **0.32.0 du 11 juillet 2026** – consolidée depuis par les versions **0.33.1 du 26 août 2026** (citée comme référence par PromptQuorum) et **0.34.0 du 5 septembre 2026** (repérée comme dernière version par Traceary et releases.sh) –, la simple commande `ollama` lance même un agent interactif capable de dialoguer, de coder et d’effectuer des recherches web, brouillant un peu plus la frontière avec les assistants cloud. Sous le capot, Ollama empaquette les poids du modèle, son *template* de prompt et ses paramètres dans un format unifié, puis délègue le calcul à llama.cpp.

Pourquoi se donner la peine d’exécuter un LLM en local alors que ChatGPT ou Claude sont accessibles en un clic ? Trois raisons dominent en 2026. D’abord la **confidentialité** : tout reste sur votre machine, ce qui élimine la fuite de données sensibles et simplifie la conformité RGPD. Ensuite le **coût** : après le téléchargement, l’inférence est gratuite et illimitée, sans facturation au token ni quota. Enfin l’**indépendance** : pas de coupure d’API, pas de modèle déprécié du jour au lendemain, et un fonctionnement même hors ligne.

Le moment est idéal car les modèles ouverts ont rattrapé une grande partie de l’écart avec les API propriétaires. DeepSeek-R1 (cumulant plus de 88 millions de téléchargements sur la bibliothèque Ollama), Qwen 3, Gemma 4 – dont l’appel d’outils profite du moteur MLX, passé en version stable dès la **0.30 le 13 mai 2026** avec un gain de vitesse d’environ **2x sur Apple Silicon** selon l’analyse d’Andrew.ooo, puis dont la gestion du cache a encore été affinée dans la version 0.32.1 du 16 juillet 2026 – ou encore `gpt-oss` – le modèle ouvert publié par OpenAI – offrent désormais du raisonnement, de la vision et l’appel d’outils. Le modèle Qwen 3.6 27B illustre ce rattrapage : il atteint 77,2 % sur le benchmark de codage SWE-bench tout en tenant dans 24 Go de VRAM en quantification Q4, ce qui en fait mi-2026 l’un des meilleurs modèles de code disponibles sur Ollama. Vous pouvez faire tourner un assistant de codage compétent sur un simple ordinateur portable récent.

| Critère | Ollama (LLM en local) | API cloud (ChatGPT / Claude) | 
|---|---|---|
| Coût d’inférence | Gratuit, illimité | Facturation au token | 
| Confidentialité des données | 100 % local, rien ne sort | Envoi vers serveurs externes | 
| Conformité RGPD | Native (aucun transfert) | Dépend du fournisseur | 
| Fonctionnement hors ligne | Oui | Non | 
| Niveau brut du meilleur modèle | Très bon (ouvert) | État de l’art | 
| Matériel requis | Le vôtre (RAM / GPU) | Aucun | 

Si vous hésitez encore sur le modèle à viser, notre comparatif Claude vs ChatGPT vs Gemini vs Mistral 2026 situe le niveau des références propriétaires que les modèles ouverts cherchent à égaler.

## Prérequis et configuration matérielle pour Ollama

La bonne nouvelle : Ollama fonctionne sur du matériel grand public. La règle d’or porte sur la mémoire. En quantification `Q4_K_M` (le format par défaut, qui réduit la taille des poids à environ 4 bits par paramètre), comptez grossièrement 1 Go de mémoire vive ou vidéo par milliard de paramètres, plus une marge pour le contexte. Un GPU n’est pas obligatoire – Ollama tourne sur CPU – mais il accélère la génération d’un facteur 5 à 20.

| Taille du modèle | RAM / VRAM recommandée | Exemple de modèle | Usage type | 
|---|---|---|---|
| 1 à 3 milliards | 4 à 8 Go | llama3.2:3b, gemma3:1b | Tâches simples, edge, mobile | 
| 7 à 8 milliards | 8 à 16 Go | llama3.1:8b, qwen3:8b | Assistant généraliste | 
| 13 à 14 milliards | 16 à 24 Go | phi4:14b, deepseek-r1:14b | Raisonnement, qualité accrue | 
| 30 à 34 milliards | 24 à 48 Go | qwen3:32b, gpt-oss:20b | Production exigeante | 
| 70 milliards et plus | 48 Go et plus | llama3.1:70b | Station de travail / serveur | 

Côté système, Ollama exige **macOS 12 (Monterey) ou plus récent**, **Windows 10/11** (64 bits) ou une distribution Linux récente (Ubuntu 22.04+ recommandé). L’accélération GPU couvre quatre familles : **NVIDIA CUDA** (la plus mature), **AMD ROCm**, **Apple Metal** (automatique sur les Mac Apple Silicon) et **Vulkan**, ce dernier ouvrant la voie aux iGPU et à un éventail plus large de cartes. Sur un Mac M-series, l’accélération est transparente : la mémoire unifiée sert à la fois de RAM et de VRAM.

Conseil pratique : commencez avec un modèle de 3 à 8 milliards de paramètres. Vous obtiendrez des réponses fluides même sur CPU, et vous monterez en gamme une fois familiarisé. Inutile d’investir dans une carte à 2 000 € avant d’avoir mesuré vos besoins réels.

## Étape 1 – Installer Ollama sur macOS, Linux et Windows

