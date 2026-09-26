---
id: collect-240926-mindstudio/mindstudio/dark-bloom-rent-out-your-mac-for-ai-inference-and-get-paid-2
title: "dark-bloom-rent-out-your-mac-for-ai-inference-and-get-paid"
domain: mindstudio
role: reference
task: reference
actors: ["Alibaba", "Apple"]
dates: []
keywords: ["qwen"]
source: docs/RAG/clean_en/mindstudio/dark-bloom-rent-out-your-mac-for-ai-inference-and-get-paid.md
source_anchor: ""
source_lines: [63, 79]
sha256: 79e36e2c1293ee1ae5770287ddf83e532c76cf3e48cf6b3ffcf79ea2d8869bb7
---

# dark-bloom-rent-out-your-mac-for-ai-inference-and-get-paid

C’est un réseau qui permet aux propriétaires de Mac de partager leur puissance de calcul inutilisée afin que d’autres puissent exécuter des modèles d’IA à poids ouverts comme Qwen, Gemma et GPT-OSS, fonctionnant comme une alternative distribuée à l’inférence cloud centralisée.

### Combien d’argent pouvez-vous gagner en faisant tourner Dark Bloom ?

Les gains dépendent de votre matériel et de l’utilisation de votre machine par le réseau. Une configuration Mac Studio à mémoire élevée a été estimée à environ quelques dizaines de dollars par mois, moins une modeste augmentation des coûts d’électricité.

### Dark Bloom est-il sûr à installer ?

Le projet publie son code ouvertement pour inspection, et un premier examen indépendant n’aurait trouvé aucun malware, minage ou vol d’identifiants. Cela dit, c’est un projet à un stade précoce sans audit de sécurité formel par un tiers, donc toute personne qui l’installe devrait évaluer ce risque par elle-même.

### Dark Bloom expose-t-il vos données à des inconnus utilisant la puissance de calcul de votre Mac ?

L’architecture déclarée du projet exécute l’inférence dans un seul processus renforcé utilisant le framework MLX d’Apple, spécifiquement conçu pour empêcher le propriétaire de la machine d’observer les prompts ou les réponses. Il s’agit d’une affirmation de conception appuyée par un livre blanc publié, et non d’une garantie vérifiée indépendamment.

### De quel Mac avez-vous besoin pour rejoindre le réseau ?

Les exigences actuelles nécessiteraient au moins 48 Go de RAM, ce qui couvre la plupart des Mac Studio, Mac mini et modèles haut de gamme de MacBook Pro, bien que ce seuil puisse changer à mesure que le projet se développe.
