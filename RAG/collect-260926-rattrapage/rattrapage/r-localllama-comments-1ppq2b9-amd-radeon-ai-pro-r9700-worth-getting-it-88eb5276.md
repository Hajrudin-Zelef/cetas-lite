---
id: collect-260926-rattrapage/rattrapage/r-localllama-comments-1ppq2b9-amd-radeon-ai-pro-r9700-worth-getting-it-88eb5276
title: "r-localllama-comments-1ppq2b9-amd-radeon-ai-pro-r9700-worth-getting-it-88eb5276"
domain: rattrapage
role: reference
task: reference
actors: ["AMD", "Nvidia", "vLLM"]
dates: []
keywords: ["amd", "fp8", "gpu", "nvidia", "vllm"]
source: docs/RAG/lot-rattrapage/ai-llm/r-localllama-comments-1ppq2b9-amd-radeon-ai-pro-r9700-worth-getting-it-88eb5276.md
source_anchor: ""
source_lines: [1, 142]
sha256: 8aaae2791693a8d7ea2d1b684448d3e60537a43bfa2b6556bddc18d7eeb09e3c
---

# r-localllama-comments-1ppq2b9-amd-radeon-ai-pro-r9700-worth-getting-it-88eb5276

Merci pour ton avis !
Explique-nous pourquoi ce contenu n’est pas utile.
      AMD Radeon AI PRO R9700, ça vaut le coup de l'acheter ? 
        
        
        
    
    
    Donc, on dirait que c'est la seule carte 32 Go qui n'est pas hors de prix et disponible, et qui n'est pas en fin de vie logiciellement parlant. Quelqu'un qui a une vraie expérience personnelle et pratique avec, surtout dans une configuration multi-cartes ?
Aussi, le grand frère de 48 Go : Radeon Pro W7900 AI 48G ?
Meilleurs
            Ouvrir les options de tri des commentaires
            
        
    
       
            
          
      
      
        Meilleurs
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Top
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Nouvelles
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Controversées
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Anciennes
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Questions et réponses
        
      
    
    
    
  
  
      
 
    
Section des commentaires
J'ai 2 R9700, et je prévois d'en ajouter 2 autres l'année prochaine. C'est des GPU corrects et ça marche bien avec vLLM. J'ai acheté ces GPU parce que je veux caser 128 Go dans mon serveur monté en rack, et ils ont des ventirads double slot. Si vous faites tourner un rig de minage en plein air et que vous avez de la place pour plus de GPU, les 3090 pourraient être un meilleur choix.
J'ai un système avec une RTX 5000 Ada que j'utilise pour les trucs Cuda. Mais pour tout ce qui est LLM, j'utilise mon système R9700. Il y a aussi un truc à dire sur le fait de pouvoir acheter du nouveau matos avec une garantie. Et l'architecture RDNA4 est plus économe en énergie qu'Ampere.
Avant de vous décider, allez faire des recherches sur les logiciels que vous voulez faire tourner et voyez s'ils supportent ROCm et Vulkan. Si c'est le cas, prenez les cartes AMD. Sinon, vous voudrez peut-être rester avec Nvidia.
L'architecture RDNA4 a aussi un support natif pour FP8, ce que les 3090 n'ont pas. Mais la plupart des logiciels qui supportent ROCm n'en profitent pas encore. Il y a une PR ouverte dans le projet AITER d'AMD pour ajouter un support natif pour le R9700. Donc les utilisateurs de vLLM pourront bientôt profiter de l'accélération FP8 de ces GPU.
T'as acheté quelle marque ? J'ai deux ASRock R9700. Elles sont bien, mais j'ai aussi renvoyé deux modèles, un ASRock et un PowerColor, qui avaient de gros problèmes de bruit de ventilateur, genre un truc qui grince ou qui siffle. Mes deux ASRock ont encore un problème de bruit de ventilateur quand ça tourne à fond pendant longtemps, mais c'est supportable…
Donc, j'ai maintenant 4 cartes ASRock creator R9700. J'avais aussi 2 de leurs cartes graphiques 7900XTX de type soufflante. J'ai toutes mes cartes graphiques installées dans un boîtier monté en rack dans une pièce à distance, mais je ne crois pas avoir jamais entendu les ventilateurs vibrer quand je suis près du serveur.
J'ai une W7900 (mon gros bébé). Elle est géniale. Ce n'est pas le GPU le plus rapide du marché (mais il n'est pas lent non plus), mais les 48 Go, le support ROCM qui mûrit rapidement et la qualité professionnelle, c'est un bon package.
Est-ce que £1500 pour un W7900 d'occasion est un bon achat en ce moment ?
Cela dépend de votre cas d'utilisation. Le R9700 AI Pro est plus rapide, moins cher et prend en charge le fp8 mais avec moins de mémoire. Tous les GPU Nvidia comparables avec 48 Go coûtent beaucoup plus cher.
Si vous voulez être capable d'exécuter des modèles de taille moyenne (30B et moins en q8 ou 70B et moins en q4) avec un long contexte, c'est une offre incroyable. Le mien coûtait presque deux fois plus cher et je l'adore. ROCM 7.2 est maintenant assez mature et chaque grande plateforme d'inférence le prend en charge.
Le plus grand point faible d'AMD est la vitesse d'entraînement - si vous ne prévoyez pas de beaucoup entraîner de modèles avec cela et que vous voulez la capacité de taille de modèle et de contexte plus grande, je dirais que oui, c'est une bonne affaire.
J'ai fait un post récemment sur le fait d'avoir eu une expérience positive en achetant une w6800 à 500$ via les offres eBay. Elle a toujours un support officiel et fonctionne très près de la R9700 en génération de tokens (la vraie force de la R9700 est le traitement des prompts où elle prend une avance plus notable).
À ce prix, prends une 3090.
Ouais, mais vu que OP envisage une R9700 contre une 4090 ou deux 3090 d'occasion, je suppose que sa priorité, c'est la VRAM par slot.
Où trouves-tu une 3090 pour 500 dollars, monsieur ?
T'en as eu un ?
Non, je voulais en prendre une autre 4090, maintenant j'ai du mal à avoir une 3090 mais au final, je vais probablement en prendre une.
Pourquoi t'as décidé de prendre une 3090 au lieu d'une 4090 ?
Je suggérerais certainement AMD Radeon Pro W7900 si ton budget le permet. https://amzn.to/3NCCg6o
C'est une bête avec 48 Go de mémoire, ça gère les moniteurs 12K et deux 8K comme un pro, encodage/décodage AV1, et ça consomme peu aussi !
