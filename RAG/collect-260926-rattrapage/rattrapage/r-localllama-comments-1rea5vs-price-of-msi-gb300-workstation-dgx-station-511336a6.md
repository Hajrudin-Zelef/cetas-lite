---
id: collect-260926-rattrapage/rattrapage/r-localllama-comments-1rea5vs-price-of-msi-gb300-workstation-dgx-station-511336a6
title: "r-localllama-comments-1rea5vs-price-of-msi-gb300-workstation-dgx-station-511336a6"
domain: rattrapage
role: reference
task: reference
actors: ["Alibaba", "DeepSeek", "Meta", "Nvidia"]
dates: []
keywords: ["blackwell", "deepseek", "gpu", "kimi", "lpddr5x", "memory", "nvidia", "nvlink", "sol", "vllm"]
source: docs/RAG/lot-rattrapage/servers-reviews/r-localllama-comments-1rea5vs-price-of-msi-gb300-workstation-dgx-station-511336a6.md
source_anchor: ""
source_lines: [1, 161]
sha256: df7e5744532874303b8eec97e7bf9bf694b38bb843e321fc8dcf72718cd2cd77
---

# r-localllama-comments-1rea5vs-price-of-msi-gb300-workstation-dgx-station-511336a6

Merci pour ton avis !
Explique-nous pourquoi ce contenu n’est pas utile.
      Prix de la station de travail MSI GB300 (DGX Station) apparu en ligne ~ 97k $ 
        
        
        
     Meilleurs
            Ouvrir les options de tri des commentaires
            
        
    
       
            
          
      
      
        Meilleurs
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Top
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Nouvelles
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Controversées
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Anciennes
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Questions et réponses
        
      
    
    
    
  
  
      
 
    
Section des commentaires
Trouvé le même article ici pour 85k : https://www.centralcomputer.com/msi-ct60-s8060-nvidia-dgx-station-cpu-memory-up-to-496gb-lpddr5x-nvidia-blackwell-ultra-gpu-1x-10-gbe-2x-400-gbe.html
Quelle est la bande passante combinée des modules RAM ?
Mémoire CPU 496 Go LPDDR5X | 396 Go/s
Quand on utilise vllm, ça voit 288 ou 784 ? Ça pourra servir DeepSeek ? Si ça donne un bon débit avec le batch, ça va être un achat facile pour mon équipe.
Et beaucoup ont reçu des votes négatifs ici l'année dernière en disant que cela coûterait plus de 56 000 $ (56 000 $ était le prix de son prédécesseur).
Au moins, il a un support complet pour Blackwell SM 10.3, contrairement au DGX Spark.
Est-ce que j'en veux un ? OUI. Qui n'en veut pas.
Peut-on en acheter un ? Non, sans prendre un prêt sur ma maison. 😥
Je sais pas mec, j'ai un peu pas envie. Qu'est-ce que je suis censé faire avec un GPU de 20 To ?
Edit : ma faute, la capacité totale de VRAM est de 784 Go
Où tu as vu un GPU de 20 To ? Il a 775 Go combinés, ce qui est bon pour faire tourner un modèle de plus de 700B comme Deepseek R1 à de très bonnes vitesses et quantifications, ou un modèle de 1 To en Q4. 🤔
Même si vous m'en donniez un gratuitement, je ne pourrais pas me le permettre en Europe, l'électricité est trop chère ici.
Si vous admettiez enfin que le truc du climat est un canular, vous devriez être en pleine révolution là-bas. Vos gouvernements vous détestent manifestement et veulent que vous échouiez.
Ils préfèrent faire du "virtue-signaling" aux islamistes qui veulent tous vous décapiter et vous humilier. Pourquoi la gauche est-elle aussi stupide à ce sujet ? Les faits sont indéniables à ce stade.
Éviter d'utiliser les combustibles fossiles tout en important de la technologie dite "verte" du régime le plus méprisable de Chine, qui pollue 10 000 fois plus que n'importe quelle nation occidentale, est la définition même de la stupidité.
Oh attendez, on va éteindre nos centrales au gaz naturel et au charbon incroyablement propres pour pouvoir importer des panneaux solaires qui ont été construits en empoisonnant le sol pendant des milliers d'années et qui ont nécessité 109 fois plus d'empreinte carbone à créer, parce qu'on est VERTS !
Comment tout cela n'est-il pas la chose la plus évidente au monde ? Êtes-vous tous payés par des acteurs chinois pour vous détruire sciemment ou quoi ? Ça n'a aucun sens.
Les faits sont objectifs et incontestables. La vérité existe, mais "votre" vérité est une connerie subjective qui, par définition, n'est PAS VRAIE. Alors réveillez-vous et je déteste ces sociopathes qui vous gouvernent.
C'est plutôt cool que tu aies un chez-toi, même si je ne vais pas mentir.
Vous pouvez construire un système 8xRTX PRO 6000 avec 768 de VRAM. À ce prix.
Tu pourrais, mais tu serais alors limité par l'interface PCIe5.
L'autre option serait un système quad H200NVL avec un pont nvlink (mais qui n'atteint pas la capacité en vram) - et ça coûterait probablement autour de 130k $.
Honnêtement, montrer une liste des prix des configurations qui pourraient faire tourner, disons, le kimi k2.5 à quelles vitesses serait intéressant.
Le seul inconvénient que je vois avec les systèmes multi-GPU, ce sont les exigences en puissance.
La RTX 6000 pro n'est pas vraiment Blackwell. Il y a donc une énorme différence
$97k c’est fou. Pour donner un peu de contexte, tu peux monter une configuration avec deux RTX 4090 pour environ $5-6k, ou même prendre 4x 3090 d’occasion pour environ $4k au total. Tu n’auras pas la mémoire unifiée ni la bande passante NVLink, mais pour la plupart des cas d’utilisation locale de LLM — inférence, ajustement de modèles plus petits, RAG — c’est largement suffisant.
J’utilise une seule RTX 4080 Super (16 Go) et je peux faire du Qwen3 30B en Q4 avec des vitesses convenables. Si j’avais $97k à dépenser, je préférerais construire 15+ de ces machines et faire de l’inférence distribuée, ou simplement remplir un rack avec des 3090.
Le GB300 a du sens pour les entreprises/recherches où tu as besoin de 288 Go de mémoire unifiée pour des modèles massifs, mais pour la communauté LocalLLaMA, c’est un peu un achat pour frimer. La voie des GPU grand public s’améliore à chaque génération.
Dual 4090/Dual 5090 à 5000-6000 $ n'a pas de sens quand à 7500 $ on peut obtenir un seul RTX6000 96 Go. 🤔
En ce qui concerne le GB300, je suis anti-Nvidia mais c'est une super machine pour un serveur costaud faisant tourner des modèles à 700B+ à la maison à grande vitesse ou dans un petit bureau avec 20-30 utilisateurs simultanés faisant tourner des modèles de taille moyenne à grande.
Ça pourrait même devenir rentable si tu l'appliques au bon business (c'est-à-dire le trading).
Pourquoi tu les compares même ? Tu ne peux littéralement pas acheter l'un d'eux sans passer par une entreprise. J'ai essayé d'envoyer un e-mail à MSI, supermicro, ASUS, Dell et les autres, et TU DOIS avoir un e-mail d'entreprise et une "carte de crédit" pour les acheter, ils ne vont littéralement pas te les vendre d'une autre manière parce qu'ils ne vendent pas ces produits un à un.
Ils viennent aussi par sets de 8 : https://servers.asus.com/products/detail/overview/XA-NB3I-E12
Si cela devient la norme dans les centres de données au cours des prochaines années, je sais que je ne pourrai pas rivaliser, donc je vais me concentrer sur le fait de tirer le meilleur parti des Pro 6000 (toutes choses considérées, je vois ceux-ci comme le meilleur choix pour les utilisateurs sérieux, du moins pour l'instant).
Ce n'est pas une DGX Station. C'est un rack de serveurs avec des B300.
La DGX Station est un produit autonome avec 2 CPU Grace et 1 Blackwell 300.
Éditer :
C'est en fait 1 CPU G, le gb300 du centre de données a les 2 CPU.
Je vois ça comme une inférence pour une petite entreprise tech. Mais à quel moment ça a plus de sens d'utiliser le cloud computing, je ne sais pas.
97 000 $ c’est beaucoup d'argent.
Tu pourrais probablement obtenir ça facilement en vendant juste trois reins.
Cela peut être plus attrayant que d'acheter un threadripper avec 4xblackwell rtx 6000 maxq (environ 80 000 $). Pour l'inférence, cela peut ne pas avoir autant d'importance, mais pour l'entraînement ou la création de petits modèles depuis zéro, je me demande comment ils se comparent pour ce cas d'utilisation ?
Pour être honnête, l'idée d'avoir 748 Go de VRAM me fait rêver un peu.
