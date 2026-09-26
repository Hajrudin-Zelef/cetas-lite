---
id: collect-260926-rattrapage/rattrapage/r-localllama-comments-1i74mao-how-trusted-is-lm-studio-04a26e10
title: "r-localllama-comments-1i74mao-how-trusted-is-lm-studio-04a26e10"
domain: rattrapage
role: reference
task: reference
actors: ["Hugging Face"]
dates: []
keywords: ["llama", "deepseek", "gguf", "gpu", "llama.cpp", "mcp", "open source"]
source: docs/RAG/lot-rattrapage/ai-llm/r-localllama-comments-1i74mao-how-trusted-is-lm-studio-04a26e10.md
source_anchor: ""
source_lines: [1, 39]
sha256: cb38997d6096ecd331e4c17be249f4ea3667948213092da701e0f4ae467b4790
---

# r-localllama-comments-1i74mao-how-trusted-is-lm-studio-04a26e10

Merci pour ton avis !
Explique-nous pourquoi ce contenu n’est pas utile.
      À quel point LM Studio est-il fiable ? 
        
        
        
    
    
    Je m'excuse si ce n'est pas l'endroit pour demander et je supprimerai le post si ce n'est pas le cas. Je suis un véritable débutant avec ces choses-là, donc je m'excuse pour mon ignorance. J'avais vu que LM Studio était une introduction conviviale pour les débutants au monde des LLM locaux et je voulais l'essayer. Mon ami plus expérimenté a dit que je ne devrais pas m'y intéresser et de ne pas lui faire confiance. Est-ce que c'est quelque chose que la communauté soutient ou est-ce que ce n'est pas fiable ?
Section des commentaires
Si tu veux des alternatives open source, tu peux regarder GPT4ALL ou Koboldcpp ou LocalAI ou Ollama ou Open-WebUI ou tout simplement llama.cpp. Il y a beaucoup plus d'options.
Je vais m'en occuper. Merci pour les suggestions.
Pas de souci.
GPT4ALL est le plus proche en fonctionnalité de LM Studio.
J'ai été très heureux d'exécuter llama.cpp avec llama-swap pour gérer différents modèles, et Open-WebUI pour gérer l'interface web de chat.
Ajouteriez-vous Jan.ai à cette liste ? C'est très convivial pour les débutants, offre une belle interface et fonctionne à la fois sur Windows et Linux.
Je sais qu'Ollama ne prend pas en charge les modèles MLX. Parmi d'autres alternatives, y en a-t-il une qui prend en charge les modèles MLX ?
Avance d'un an plus tard, je pense que c'est le cas maintenant
Essayez kolosal.ai, c'est léger (seulement 20 Mo) et open source. Ils ont aussi une fonctionnalité serveur et on peut définir le nombre de couches déchargées sur le GPU.
Commentaire supprimé par un membre de l’équipe de modération
Yesss, Kolosal.ai est aussi une alternative à LMStudio, qui est open source
C'est un code source fermé, je ne vois pas comment nous pourrions éventuellement 'être de son côté', mais cela ne semblait pas pirater mon système.
Mon ami est très prudent face à tout ce qui n'est pas open source. Il m'apprend un peu à naviguer dans le domaine, donc je suis leurs conseils, mais tout ce que j'ai vu c'est que LM studio est souvent évoqué positivement sur YouTube.
Lm studio est un excellent point de départ, c'est probablement le plus facile pour un débutant, mais tu n'as pas à te limiter à un seul. Essaie tous. Ils semblent tous avoir leurs propres points forts. Ils utiliseront tous les mêmes fichiers .gguf de HuggingFace.com sauf Ollama, qui pour une raison quelconque te verrouille à utiliser leurs modèles propriétaires.
J'utilise koboldcpp et il ne m'a jamais déçu.
C'est bon, utilisez-le, si vous utilisez Windows, vous avez déjà 1000~ trucs fermés dans votre système.
Il n'y a pas de problèmes connus avec LLM Studio. Vous pouvez utiliser des alternatives plus ouvertes, mais au final, même les logiciels entièrement open source ne peuvent pas être totalement fiables à 100 %.
Évaluez votre risque. Quelle serait la gravité d'une violation ? Traitez-vous des données critiques ou confidentielles ? Ou êtes-vous juste en train d'expérimenter avec des LLM pour avoir un chatbot local ?
D'après votre message, il semble que ce soit ce dernier cas, et dans ce cas, je ne m'en inquiéterais pas.
LM Studio bloque l'ordinateur pendant l'exécution de Deepseek R1 32B, ouvrir webui + ollama fonctionne bien. Juste un point de données.
Mec, tu essaies de faire tourner un modèle 32B...
Salut ! Je suis en train de créer https://kolosal.ai , c'est une alternative open source à LM Studio, et c'est très léger, seulement 16 Mo d'installateur, et ça fonctionne super bien pour la plupart des GPU et des CPU. Ça a aussi des fonctionnalités serveur, et nous travaillons à ajouter MCP, l'augmentation de données et des fonctionnalités d'entraînement.
Ça a l'air génial - un ensemble de fonctionnalités vraiment solides. Y a-t-il des projets pour une version MacOS de cela ?
Merci. Nous prévoyons de supporter, mais nous faisons encore face à des bugs critiques que nous devons régler avant d'ajouter le support sur d'autres systèmes d'exploitation. On y arrivera éventuellement.
LM Studio est actuellement l'application la mieux notée pour faire fonctionner des llms localement. Elle vous permet d'ajouter des outils pour connecter votre IA locale à Internet, ainsi que de lui donner des capacités vocales. C'est privé et personnalisable. Votre ami est très mal informé.
non vede la seconda gpu di Tesla K80 come si abilita??
Oublie ça, après la pénurie de RAM, ce que je soupçonnais s'avère être vrai. Lorsque je charge des versions plus récentes du runtime CUDA, elles font planter tous les modèles avec plus de contexte. En essayant des moteurs de runtime CUDA de version inférieure, les modèles se chargent sans problème. C'est clairement une astuce et des conneries en cours.
c'est très fiable. je ne sais pas pourquoi l'autre a dit que c'était closed source. ce n'est pas le cas. https://github.com/lmstudio-ai
L'existence d'un compte github != Open source.
