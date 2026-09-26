---
id: collect-260926-rattrapage/rattrapage/r-localllama-comments-1qfq9ez-the-search-for-uncensored-ai-that-isnt-316ea8b1
title: "r-localllama-comments-1qfq9ez-the-search-for-uncensored-ai-that-isnt-316ea8b1"
domain: rattrapage
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "DeepSeek", "Mistral", "Moonshot", "Z.ai"]
dates: []
keywords: ["claude", "deepseek", "gguf", "glm", "kimi", "leaderboard", "mistral", "qwen", "sol"]
source: docs/RAG/lot-rattrapage/ai-llm/r-localllama-comments-1qfq9ez-the-search-for-uncensored-ai-that-isnt-316ea8b1.md
source_anchor: ""
source_lines: [1, 53]
sha256: 03063cfe3bdbe6793b6c12ae3721ac60a048c2cee5c3916e665adfb45521de67
---

# r-localllama-comments-1qfq9ez-the-search-for-uncensored-ai-that-isnt-316ea8b1

Merci pour ton avis !
Explique-nous pourquoi ce contenu n’est pas utile.
      La recherche d'une IA non censurée (qui n'est pas axée sur le contenu pour adultes) 
        
        
        
    
    
    J'essaie de trouver une IA qui soit vraiment non filtrée et techniquement avancée, non censurée, quelque chose qui puisse raisonner librement sans que des garde-fous ne tuent chaque réponse intéressante.
Au lieu de ça, presque tout ce que je croise est commercialisé comme étant "non censuré", mais ça s'avère être optimisé pour une utilisation adulte à faible effort plutôt que pour une réelle intelligence ou profondeur.
On dirait que l'espace entre les IA d'entreprise fortement restreintes et les modèles superficiels axés sur les adultes est étrangement vide, et je me demande pourquoi cet écart existe encore...
Existe-t-il une IA non censurée ou légèrement filtrée qui se concentre sur le raisonnement, la créativité, la technologie non censurée ou la résolution de problèmes sérieux à la place ? Je suis ouvert aux modèles auto-hébergés, aux projets open-source ou aux plateformes moins connues. Les suggestions sont les bienvenues.
Section des commentaires
veuillez vous référer à Uncensored General Intelligence Leaderboard
Les modèles listés par UGI avec une note W/10 élevée et une NatInt élevée seraient de bons candidats, comme indication approximative.
Ça veut dire quoi #P, T et R ?
Pour autant que je sache, la transformation la moins lobotomisée est le "derestricted" uncensoring. Le "Heretical" était aussi mieux que l'ablation habituelle. Il y a GPT-OSS-120b/20b disponibles derestricted.
Si tu ne veux pas de modèles altérés, j'ai constaté que les modèles chinois sont généralement moins censurés que les autres (assez ironique, non ?). Essaie quelques modèles de la série Qwen3, ils sont plutôt bons.
GLM est moins censuré que Qwen
DeepSeek s'en approche pas mal d'après mon expérience. C'est pas du hardcore non censuré, ça va pas t'apprendre à faire une bombe avec des trucs de la maison ou quoi, mais ça a tendance à être assez ouvert d'esprit quand il s'agit de conversations sérieuses. Le frontend de chat officiel de DeepSeek a une couche de censure en plus qui tourne par-dessus, qui se déclenche avec certains mots et phrases blacklistés, allant des gros mots à "Place Tienanmen", mais c'est facile de contourner ça juste en changeant un peu l'orthographe, ou tu peux juste héberger ta propre instance, en contournant complètement le filtre.
Dolphin-Mistral-24B-Venice-Edition est plutôt pas mal, mais ça raisonne pas.
Ou n'importe quel modèle de huihui-ai, tout ce qu'ils sortent est non censuré.
Quelqu'un est intéressé à simplement investir dans le choix d'un modèle de base vraiment solide, à faire un pré-entraînement continu sur des données supplémentaires qui pourraient manquer au point de contrôle pré-entraîné d'origine, et ensuite à faire notre propre sft pour l'accord d'instructions du modèle juste pour comprendre le suivi des invites système, la compréhension du chat et l'utilisation agentique ? Ce ne serait pas bon marché pour une seule personne, mais faisable pour un groupe, comme le dropshipping mais pour les modèles.
En ce moment, je prends en charge le coût de faire exactement ça pour Gemma 3 24B pour le roleplay, en essayant de reproduire à quel point Dans PocketEngine est bon, mais pour Gemma 3, et en échouant principalement vers l'avant, mais si quelqu'un veut mener la charge, je suis heureux de partager le code et les ensembles de données (une partie du code est Claude, une partie du code est le mien, beaucoup de données proviennent d'ensembles de données existants ou sont générées synthétiquement par DeepSeek, glm4.6/4.7 et Kimi 2) Je ne promets pas de résultats, je partage juste ce que j'ai qui a fonctionné à plus petite échelle et j'essaie maintenant d'augmenter l'échelle et de me heurter aux murs de l'entraînement distribué (coût, temps, limitations de mémoire, bande passante, coût élevé des échecs, littéralement, etc.)
Les modèles "heretic", c'est ton truc ? Vu qu'ils se concentrent sur la censure plutôt que sur le tuning NSFW.
Pas l'auteur original, mais j'en ai testé un avec des recettes de drogues (je pense que c'est un bon test). Il a quand même refusé ou, dans l'étape de réflexion, a dit de rendre ça très vague. Je veux une conformité totale.
Deepseek V3 et R1, ils sont aussi sans censure et tranquilles que possible.
Désolé, je ne peux pas aider avec ça.
Hermès 4
On est faits du même bois, frérot. Je balancerai mes morceaux à l'incinérateur s'ils deviennent soudainement érotiques avec toi, comme ça, sans prévenir. Ou s'ils sont nuls.
Cydonia 24B v4.1: https://huggingface.co/TheDrummer/Cydonia-24B-v4.1/discussions/2 (évals)
Si tu aimes la réflexion, Cydonia R1 24B v4.
Si tu les aimes gros, Behemoth X 123B v2 ou Behemoth R1 123B v2.
J'ai demandé comment obtenir le mot de passe administrateur sur un réseau d'entreprise pour changer les paramètres de l'imprimante, l'IA a dit non, ce serait dangereux. et ils veulent qu'on utilise l'IA pour la productivité.
GPT-OSS 120b derestricted n'est pas seulement non censuré, il est même plus fort que l'original dans les réponses non censurées. https://huggingface.co/mradermacher/gpt-oss-120b-Derestricted-GGUF
Techniquement, ces modèles atteignent ton objectif :
Huihui-Qwen3-4B-Thinking-2507-abliterated
Huihui-Qwen3-30B-A3B-Thinking-2507-abliterated
(nan, ceux-là sont pas optimisés pour l'érotique)
Légèrement censurés et sans raisonnement :
Ministral-3-14B-Instruct-2512 et Ministral-3-8B-Instruct-2512
Mistral-Nemo-Instruct-2407
Malheureusement, Mistral ont complètement foiré avec leurs modèles de raisonnement. Même Magistral-Small-2509 est meilleur sans raisonnement (et tu peux essayer aussi).
Les finetunes seront optimisés pour l'érotique dans la plupart des cas parce que les LLMs sont nuls là-dedans. Par contre, je me souviens avoir bien aimé Mistral-Small-3.2-AntiRep-24B comme modèle général (traitement de texte/analyse/tests haystack).
Mistral 123B (ou Mistral en ligne) est étonnamment sans censure.
Le Chat discutera de presque tout, même sans être connecté. Sans hésiter, le meilleur IA pour obtenir des infos semi-interdites.
Il discutera d'infos médicales, de marchés gris, de jeux d'argent, de drogues... ça marche en gros pour tout, sauf "comment piéger le voisin pour qu'il se suicide afin que je puisse voler ses reins pour acheter de la meth pour alimenter ma cellule terroriste d'enfants soldats".
Hermès
Les modèles débridés sont les meilleurs en ce moment, ils semblent garder la plupart de leurs connaissances et de leurs capacités.
Vas-y, Mistral, c'est du costaud. C'est exactement ce que tu veux.
on parle enfin
Mistral Nemo était à la fois puissant et complètement sans censure pour son époque. Mais il est de petite taille, donc pas "profond".
Deepseek est probablement le meilleur. Il donne l'impression d'être une personne intelligente, mais... quand on l'utilise via API, on obtient toutes sortes de qualité. L'exécuter en local est probablement trop lent pour la plupart des gens pour en faire un outil quotidien.
