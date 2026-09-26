---
id: collect-260926-rattrapage/rattrapage/r-localllama-comments-1n27vhg-asus-ascent-gx10-nvidia-dgx-spark-s-this-genuine-bef5fe1c
title: "r-localllama-comments-1n27vhg-asus-ascent-gx10-nvidia-dgx-spark-s-this-genuine-bef5fe1c"
domain: rattrapage
role: reference
task: reference
actors: ["AMD", "Nvidia", "United States"]
dates: []
keywords: ["nvidia", "amd", "benchmarks", "gpu", "lpddr5x"]
source: docs/RAG/lot-rattrapage/servers-reviews/r-localllama-comments-1n27vhg-asus-ascent-gx10-nvidia-dgx-spark-s-this-genuine-bef5fe1c.md
source_anchor: ""
source_lines: [1, 141]
sha256: 55f371d623cbbb7541e252621268219646bf844136d7fe7f604614745235b288
---

# r-localllama-comments-1n27vhg-asus-ascent-gx10-nvidia-dgx-spark-s-this-genuine-bef5fe1c

Merci pour ton avis !
Explique-nous pourquoi ce contenu n’est pas utile.
      ASUS Ascent GX10 (NVIDIA DGX™ Spark): est-ce que c'est authentique et correct pour les llms ? 
        
        
        
    
    
    Je ne suis pas sûr si ce mini-PC basé sur NVIDIA en vaut la peine. Quelqu'un l'a-t-il testé ? Est-ce qu'il fonctionne correctement pour les inférences LLM ? Quelles vitesses PP et TG sont attendues pour des modèles comme gpt-oss-120b ?
Meilleurs
            Ouvrir les options de tri des commentaires
            
        
    
       
            
          
      
      
        Meilleurs
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Top
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Nouvelles
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Controversées
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Anciennes
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Questions et réponses
        
      
    
    
    
  
  
      
 
    
Section des commentaires
Elle a de la mémoire DDR5 standard, pas la version rapide. Sans aucun test, je peux déjà prédire que ce truc sera plus rapide qu'un CPU grand public, mais plus lent qu'un GPU normal. Cherchez les tests des mac m4 et ryzen ai max pour référence, dgx spark va être pareil en termes de vitesse de génération de tokens, avec peut-être un traitement des prompts plus rapide.
La bande passante mémoire n'est que de 270 Go/s, la RTX5060ti est presque à 500 Go/s, donc tu vas profiter d'une grande VRAM, mais la vitesse ne s'annonce pas terrible ici.
Pour l'inférence CPU, c'est bien, donc selon le prix, ça pourrait être une bonne alternative.
Donc ça dépend purement du prix. Et je peux déjà dire ça, ça ne vaudra pas le coup.
Juste en termes de prix, c'est une offre pourrie. Genre, un module V100 32GB SXM2 avec un adaptateur PCIe te coûterait environ 700$ pièce, soit 2100$ pour 128GB de VRAM - mais ce serait du vrai Nvidia avec de la mémoire HBM2, qui surclassera le dgx spark au moins dix fois. Ou, si tu veux économiser, tu peux prendre des Mi50 32GB - tu perds en vitesse par rapport aux V100, mais tu surclasses quand même le dgx spark pour 200$ par carte (800$ + les autres frais PC). Les mini PC comme ça, ça n'a de sens que dans trois cas : quand l'espace ou le bruit posent problème, quand ton électricité coûte un bras, ou quand c'est ton patron qui paie la facture, et qui fait probablement une défiscalisation en plus.
CPU basé sur ARM pour mobile, GPU RTX5070 avec une bande passante théorique de 273 Go/s (donc beaucoup plus lent que le dGPU) utilisant de la LPDDR5X soudée, bloqué dans le système d'exploitation Linux basé sur Ubuntu de NVIDIA avec un prix supérieur à 3000 $ (le prix de vente conseillé est de 4000 $).
Donc, comparé à des trucs comme l'AMD AI MAX 395, c'est 50 à 100 % plus cher, pour des perfs similaires* (astérisque ici), et on peut pas faire grand-chose avec à cause du CPU ARM. Et on peut pas faire tourner Windows et jouer à des jeux par exemple.
C'est MOYEN au mieux.
* Perfs. faut rester dans un TOUT PETIT panier ici parce que j'ai un mauvais pressentiment à ce sujet.
Y'a un hic : c'est un système ARM, donc tu peux pas faire tout ce que tu veux.
Étant donné que l'inférence LLM est limitée par la bande passante mémoire, ce truc a environ 273 Go/s contre 256 Go/s dans l'AMD AI MAX+ 395. Nvidia coûte environ deux fois plus cher qu'AMD, et c'est aussi pas x86, donc AMD pourrait être un meilleur choix.
Quand même, on n'a pas encore de benchmarks.
Y'a une chance que ça puisse être plusieurs fois plus rapide que l'option AMD : par exemple, la bande passante mémoire pratique sur le système AMD est seulement de 190 - 200 Go/s. Si la vitesse pratique sur le Nvidia finit par être beaucoup plus proche de sa bande passante théorique, et que les modules de calcul sont aussi plus rapides, alors peut-être que ça vaudrait le coup de le considérer malgré le prix plus élevé.
Le système utilise un SoC ARM Mediatek bon marché, une carte réseau Nvidia excessivement chère et une carte graphique 5070 dépouillée qui partage 128 Go de RAM relativement lente. À un prix de 1000 dollars US, deux unités combinées pourraient offrir une solution raisonnable qui tire parti de la carte réseau hors de prix. Cependant, dans sa forme actuelle, et en l'absence de données de référence convaincantes, il semble que ce soit une autre tentative de Nvidia pour extraire un maximum de marge et probablement aussi la marge des détaillants.
C'est autant un superordinateur que je suis un super-héros.
Et dans mon coin, Asus m'a cité 4 570 dollars US, ce qui m'a littéralement fait éclater de rire.
Commentaire supprimé par un membre de l’équipe de modération
C'est ça que beaucoup de gens ne comprennent pas avec ça.
Et que c'est fondamentalement une plateforme de développement pour les POC (Proof of Concept) et le prototypage, qui peut être basculée sur de gros systèmes (NVL72 GB200/300) avec un minimum de boulot (d'où les ports QSFP56 200gbe qui permettent de prototyper la logique système à système).
C'est pas vraiment pour le consommateur lambda.
