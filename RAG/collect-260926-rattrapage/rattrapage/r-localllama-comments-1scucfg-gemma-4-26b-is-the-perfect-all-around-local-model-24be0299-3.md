---
id: collect-260926-rattrapage/rattrapage/r-localllama-comments-1scucfg-gemma-4-26b-is-the-perfect-all-around-local-model-24be0299-3
title: "r-localllama-comments-1scucfg-gemma-4-26b-is-the-perfect-all-around-local-model-24be0299"
domain: rattrapage
role: reference
task: reference
actors: ["Alibaba", "OpenAI"]
dates: []
keywords: ["agent", "moe", "qwen"]
source: docs/RAG/lot-rattrapage/ai-llm/r-localllama-comments-1scucfg-gemma-4-26b-is-the-perfect-all-around-local-model-24be0299.md
source_anchor: ""
source_lines: [32, 54]
sha256: 3f30987a75abc966a9d141cf58398291476793e5f685d35de06e4496bfab5e6d
---

# r-localllama-comments-1scucfg-gemma-4-26b-is-the-perfect-all-around-local-model-24be0299

Dernier commentaire - Quel temps pour être en vie ! Avoir ce pouvoir sur votre machine locale ! Qu'est-ce qui est mieux que l'intelligence ? (même si elle est artificielle ?) Je suis très reconnaissant envers toute cette communauté.
Merci pour le partage. Pourrais-tu expliquer un peu plus comment tu as corrigé Qwen pour pi ? Mon expérience jusqu'à présent avec cline, roo et zed agent n'est pas géniale et je suis intéressé par l'essai de pi pour voir comment ça marcherait. J'ai essayé Qwen 3.5 122b et 27b.
As-tu aussi essayé Qwen3.5 27b (dense) et Gemma4 31b (dense) pour voir comment ils se comparent au modèle Qwe3.5 MoE et au modèle Gemma4 MoE ?
Je sais qu'ils sont bien sûr beaucoup plus lents en termes de tokens/seconde que les homologues MoE de taille similaire, mais certaines personnes disaient qu'ils sont nettement plus puissants que les versions MoE. Donc, en termes de temps total passé sur une tâche globale, ils peuvent parfois être "plus rapides", s'ils peuvent accomplir des tâches en moins d'essais (ou même être capables de le faire par rapport à ne pas le faire), comparés aux modèles MoE, même si les modèles MoE fonctionnent à des tokens/seconde plus rapides. Je veux dire, ça varie évidemment en fonction de la tâche spécifique et des types de cas d'utilisation (et parfois juste de la chance aussi, d'une tentative à l'autre, je suppose).
Quoi qu'il en soit, je suis curieux de savoir si tu les as essayés aussi et comment ils se comparent selon toi et pour ce que tu as essayé avec eux.
Le 31b est tout de même bien meilleur, je comprends que la vitesse est bien pire, mais à mon avis, je fais toujours fonctionner le modèle le plus intelligent que je peux.
J'avais lu que Qwen3.5-27b était toujours meilleur en codage que Gemma-4, donc c'est une super nouvelle !
Comment ça se compare au niveau de la conversation avec Gemma-3 ?
Gemma4 semble beaucoup plus intelligente et nuancée.
Gemma3 27B était déjà vraiment bonne, mais Gemma4 vous laisse avec un sentiment beaucoup plus fort que le modèle a de la profondeur et une intention derrière ce qu'il dit.
En termes de connaissance du monde, elles sont similaires. En termes de raisonnement, Gemma4 est d'un ordre de grandeur supérieur. C'est comme ce que GPT-4 était pour GPT-3.
Utiliser Gemma4 avec Hermes, mais c'est très en désordre.
Je suppose que je dois réessayer, parce que pour mes tests, c'était terrible en codage.
J'ai testé le même jour où il a été lancé à Q6 et 128k de contexte
Le contexte de 128k, c'est ce qui change la donne pour moi. Un contexte plus long signifie que vous pouvez transmettre plus d'état dans le pipeline sans le subdiviser - c'est vraiment utile pour les flux de travail d'agent. La capacité multimodale est aussi étonnamment solide pour un modèle de cette taille. Sur quel matériel l'exécutez-vous ?
Un MacBook Pro M5 Pro de 48 Go est-il suffisant ?
Je veux créer un assistant exécutif local
Le codage agentique de CLINE est plutôt mauvais avec ça.
Toutes les familles Qwen 3.5 s'en sortent bien
Et Qwen 3 Coder Next est au-dessus de tout.
Un 26b MoE sur un Mac de 64 Go est plutôt le bon choix en ce moment. Il ne charge que les poids experts actifs, donc tu as beaucoup plus de contexte utilisable que ce à quoi tu t'attendrais par rapport au nombre de paramètres. À mon avis, le qwen 3.5 27b est toujours meilleur pour le code pur, mais le gemma gère tout le reste sans ramer.
Il charge l'ensemble du modèle, donc l'espace disponible reste le même. Le gain est en vitesse, pas en taille. 27a4b pèsera le même poids que 27b.
J'ai un test visuel avec la photo d'une femme tenant un bouquet avec 3 types de fleurs (dahlias, renoncules, queue de lapin). Les renoncules ressemblent à une rose dense. Qwen 35B Q4 identifie correctement les fleurs, Gemma 26B Q6 les appelle des roses et se souvient des renoncules seulement après avoir été demandée si ce sont vraiment des roses ?
