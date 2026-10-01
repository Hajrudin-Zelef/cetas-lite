---
id: collect-261001-ia-llm/ia-llm/nemotron-3-5-lightning-nvidia-lance-un-agent-ia-open-2026-1
title: "nemotron-3-5-lightning-nvidia-lance-un-agent-ia-open-2026"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "CoreWeave", "Google", "Nvidia", "OpenAI", "OpenRouter", "SGLang", "TensorRT-LLM", "vLLM"]
dates: []
keywords: ["agent", "nvidia", "agents", "apache", "attention", "datacenter", "fine-tuning", "moe", "nvfp4", "open source", "sglang", "tensorrt"]
source: docs/RAG/collect-261001-ia-llm/nemotron-3-5-lightning-nvidia-lance-un-agent-ia-open-2026.md
source_anchor: ""
source_lines: [1, 37]
sha256: 965d44d0e1c51ce2ddbc1c7915281efcd08c6816f2c0ffa5ffddd294d71f15b4
---

# nemotron-3-5-lightning-nvidia-lance-un-agent-ia-open-2026

Nvidia a lancé le 11 août 2026 deux briques logicielles qui changent la donne pour les développeurs d’agents IA en Europe : **Nemotron 3.5 Lightning**, un modèle ouvert de 30 milliards de paramètres taillé pour les tâches longues et autonomes, et **NeMo Switchyard**, une bibliothèque open source qui route le trafic entre plusieurs fournisseurs de modèles. L’annonce tombe six jours après l’entrée en application des règles de transparence de l’AI Act européen, le 2 août 2026, ce qui place immédiatement ces deux outils sous le regard des régulateurs bruxellois. Pour les équipes techniques françaises qui déploient des agents IA en production, la coïncidence de calendrier n’a rien d’anodin.

## Nemotron 3.5 Lightning : ce que Nvidia a réellement publié

Nemotron 3.5 Lightning repose sur une architecture hybride qui mêle blocs Mamba-2, mélange d’experts (MoE) et couches d’attention classique. Sur le papier, le modèle affiche 30 milliards de paramètres au total, mais seuls 3 milliards sont activés à chaque token traité (configuration dite « 30B A3B »). Cette parcimonie est le cœur de l’argument de vente : Nvidia ne cherche pas à dominer les classements de raisonnement pur, mais à proposer un modèle rapide et bon marché à faire tourner en boucle, comme le ferait un agent qui enchaîne des dizaines d’appels d’outils avant de rendre un résultat final.

La fenêtre de contexte annoncée grimpe jusqu’à 1 million de tokens, un chiffre qui rapproche Nemotron 3.5 Lightning des modèles à contexte long de Google ou d’Anthropic, mais avec un coût d’inférence bien inférieur grâce à l’activation partielle des paramètres. Nvidia distribue le modèle sous deux formats de poids, NVFP4 et BF16, et publie non seulement les poids mais aussi une partie des données d’entraînement et des recettes de fine-tuning, sous la licence **OpenMDW-1.1** (Open Model and Data Weights). Cette licence autorise un usage commercial, ce qui distingue Nemotron 3.5 Lightning d’un simple modèle de recherche publié pour la forme.

Autre détail qui parlera aux équipes francophones : la fiche technique du modèle liste explicitement le français, l’espagnol, l’allemand, l’italien et le japonais parmi les langues supportées, aux côtés de l’anglais et des langages de programmation courants. Ce n’est pas un support multilingue de façade : Nvidia positionne clairement Nemotron 3.5 Lightning comme un modèle déployable en Europe, du edge (cartes Jetson) jusqu’au datacenter (clusters DGX Spark), en passant par des piles de service standard comme vLLM, SGLang ou TensorRT-LLM.

## NeMo Switchyard : router les agents plutôt que les enfermer

Le deuxième outil publié le même jour, NeMo Switchyard, répond à un problème différent mais complémentaire. Une fois qu’une équipe a déployé plusieurs modèles (un modèle rapide pour les tâches simples, un modèle puissant pour les cas complexes, un modèle local pour les données sensibles), il faut un mécanisme pour aiguiller chaque requête vers le bon backend. C’est exactement le rôle de Switchyard : un proxy écrit en Rust, distribué sous licence Apache 2.0, qui expose un point d’entrée unique compatible avec l’API OpenAI tout en routant les appels, selon des règles définies par l’équipe, vers des modèles hébergés localement (vLLM, Ollama) ou chez des fournisseurs tiers.

Switchyard traduit également entre les formats d’API OpenAI et Anthropic, ce qui évite de réécrire l’intégration à chaque changement de fournisseur, et journalise des métriques opérationnelles utiles pour le pilotage des coûts. Le dépôt GitHub officiel, hébergé sous l’organisation NVIDIA-NeMo, précise toutefois que le projet est en phase « pré-alpha » et n’est pas recommandé pour un usage en production immédiat. Pour les équipes DevOps françaises qui expérimentent déjà des architectures multi-modèles, c’est un signal à suivre plutôt qu’à adopter dès aujourd’hui.

La logique stratégique de Nvidia devient plus claire une fois les deux annonces mises côte à côte, comme le détaille MarkTechPost dans son analyse du lancement. L’entreprise ne se contente plus de vendre les puces sur lesquelles tournent les modèles : elle publie désormais le modèle lui-même et l’infrastructure logicielle qui décide quel modèle traiter chaque requête. C’est une extension verticale de la chaîne de valeur de l’IA, du silicium jusqu’à la couche d’orchestration applicative, qui rappelle la stratégie déjà engagée avec CUDA, TensorRT et la gamme NIM.

## Tarification : combien coûte réellement Nemotron 3.5 Lightning

Comme pour tout modèle à poids ouverts, il n’existe pas de tarif unique fixé par Nvidia : le prix dépend du fournisseur d’hébergement choisi. DeepInfra facture le modèle 0,05 dollar par million de tokens en entrée et 0,20 dollar par million de tokens en sortie. CoreWeave propose une tarification légèrement supérieure, à 0,10 dollar en entrée et 0,25 dollar en sortie. OpenRouter, de son côté, propose un palier gratuit, avec toutefois une limitation notable : la génération en sortie plafonne à 65 536 tokens par appel, même si la fenêtre de contexte totale du modèle atteint 1 million de tokens. Pour un hébergement local, sur ses propres serveurs ou via un fournisseur cloud européen, le coût dépend uniquement du matériel utilisé, la licence OpenMDW-1.1 n’imposant aucune redevance.

| Fournisseur | Prix entrée (par million de tokens) | Prix sortie (par million de tokens) | Particularité | 
|---|---|---|---|
| DeepInfra | 0,05 $ | 0,20 $ | Tarif le plus bas identifié au lancement | 
| CoreWeave | 0,10 $ | 0,25 $ | Infrastructure Nvidia dédiée | 
| OpenRouter (gratuit) | 0 $ | 0 $ | Sortie plafonnée à 65 536 tokens par appel | 
| Auto-hébergement (vLLM/SGLang/TensorRT-LLM) | Coût matériel uniquement | Coût matériel uniquement | Aucune redevance de licence sous OpenMDW-1.1 | 

Cette fourchette de prix place Nemotron 3.5 Lightning dans la catégorie des modèles économiques, nettement en dessous des tarifs pratiqués par les modèles propriétaires de pointe. Le compromis est assumé : Nvidia ne prétend pas rivaliser avec les meilleurs modèles de raisonnement du marché sur la qualité brute, mais mise sur le rapport coût-débit pour les charges de travail agentiques répétitives, où un modèle plus cher n’apporterait pas de gain proportionnel.

## Le contexte réglementaire : l’AI Act change la donne pour les modèles ouverts

Le calendrier de cette double annonce ne pouvait pas mieux tomber pour illustrer un tournant réglementaire européen. Depuis le 2 août 2026, la Commission européenne applique de nouvelles règles de transparence issues de l’AI Act, pilotées par son AI Office en coordination avec les autorités nationales. Selon la Commission européenne, « à partir du 2 août 2026, l’AI Office de la Commission européenne, avec les autorités nationales, commencera à faire appliquer le règlement sur l’intelligence artificielle ».

La Commission précise également que « à la même date, de nouvelles règles de transparence commenceront à s’appliquer, exigeant de certains systèmes d’IA qu’ils informent les utilisateurs lorsqu’ils interagissent avec une IA et lorsqu’un contenu a été généré ou modifié par elle ». Concrètement, cela signifie que « les chatbots et autres systèmes d’IA interactifs devront indiquer aux utilisateurs qu’ils ont affaire à une IA, et non à un humain », et que « les deepfakes (images, vidéos ou fichiers audio modifiés ou générés par IA) devront être étiquetés ». La Commission ajoute que « le contenu généré ou modifié par IA devra également porter des marques lisibles par machine afin d’être détecté plus facilement ».

