---
id: collect-260926-rattrapage/rattrapage/r-localllama-comments-1mj9e2y-underwhelmed-by-512gb-m3-ultra-mac-studio-3a62883c
title: "r-localllama-comments-1mj9e2y-underwhelmed-by-512gb-m3-ultra-mac-studio-3a62883c"
domain: rattrapage
role: reference
task: reference
actors: ["Apple", "Nvidia"]
dates: []
keywords: ["llama", "attention", "fine-tuning", "gpu", "llama.cpp", "mcp", "moe", "nvidia", "sol", "vllm"]
source: docs/RAG/lot-rattrapage/ai-llm/r-localllama-comments-1mj9e2y-underwhelmed-by-512gb-m3-ultra-mac-studio-3a62883c.md
source_anchor: ""
source_lines: [1, 164]
sha256: 1ee01fd7fe9e9cff63a5d293bb822b2d23151207db3ae36b66c13e8898a77d00
---

# r-localllama-comments-1mj9e2y-underwhelmed-by-512gb-m3-ultra-mac-studio-3a62883c

Merci pour ton avis !
Explique-nous pourquoi ce contenu n’est pas utile.
      Déçu par le Mac Studio M3 Ultra 512 Go 
        
        
        
    
    
    Je sais pas trop à quoi je m'attendais, mais mon nouveau Mac Studio 512 Go a pas l'air d'être la bête de course que j'espérais - je crois que je m'attendais à des perfs plus rapides.
Meilleurs
            Ouvrir les options de tri des commentaires
            
        
    
       
            
          
      
      
        Meilleurs
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Top
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Nouvelles
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Controversées
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Anciennes
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Questions et réponses
        
      
    
    
    
  
  
      
 
    
Section des commentaires
Tu t'y prends mal. J'ai fait le même achat. 512 Go Studio 4 To SSD.
C'est pas une bête de course pour l'inférence, mais ça peut affiner des modèles avec mlx.lm.
Commence à affiner des modèles, mon pote.
T'as besoin de tonnes de vram pour le fine-tuning ?
Je pense que tu as juste besoin de pouvoir charger le modèle et le jeu de données sur lequel tu fais du fine-tuning.
Quelle partie est trop lente ? Le remplissage préalable ou la génération de texte ?
En général, les LLM locaux ont du sens en termes de sécurité, pas en termes de jetons par dollar.
Le pré-remplissage est lent. La vitesse de génération de R1 est correcte.
Note pour l'auteur - assure-toi d'utiliser un serveur d'inférence avec le caching activé, comme llama.cpp avec --cache-reuse activé. Ça rend l'Ultra parfait pour discuter avec des modèles plus intelligents, et rend l'utilisation agentique beaucoup plus faisable.
C'est parce qu'anthropoc est en train de perdre un max. Quand les vrais prix vont arriver, ça coûtera moins cher d'acheter un rig à 30k en mensualités que de payer pour un LLM.
Ça fait quoi ?
Traitement de beaucoup de contexte, probablement. Désolé que tu n'aies pas reçu le mémo, OP.
Que voulais-tu en faire et qu'est-ce que tu espérais ? Je suis content du mien, ce n'est pas le plus rapide, mais je ne veux pas faire tourner plusieurs 3090 et les performances semblent correctes pour des trucs asynchrones comme paperless-ai et discuter.
Et le bruit des fans...
J'en ai acheté un et il est super pour le développement de mon produit SaaS d'IA. Il est vraiment excellent avec les modèles MoE, à mon avis. Il me donne 75% de la performance d'un serveur B200 que j'ai en production, avec un contexte raisonnable.
Pour info, l'OP, ce mec fait des comparaisons approfondies entre Nvidia et M3/M2 Ultra (ou d'autres puces Studio) sur sa chaîne.
Pas affilié, juste fan de son contenu pertinent et de qualité : https://www.youtube.com/@AZisk
Ses comparaisons sont basiques et pas terribles, à mon avis.
Il ne met pas assez de contexte dans ses prompts, comme si tu utilisais du code agentique comme le code Cline / Roo, etc. (32k de contexte minimum).
Quand tu commences à utiliser des modèles au-delà du simple chat. Genre, utiliser RAG, MCP, la recherche web, etc… Il devrait faire des trucs comme ça dans ses comparaisons, parce que ça remplit le contexte et ça met à rude épreuve le matos / la bande passante, au-delà d'un simple prompt de chat LLM court.
Il compare pas la qualité entre les LLMs locaux et les modèles de pointe. Je trouve pas que les tokens par seconde soient super utiles.
C'est un dev, mais il donne pas son avis sur les LLMs locaux pour coder.
Mon expérience, qui est courante, c'est que pour coder, la différence entre les locaux et les gros LLMs est plus grande que ce que les scores des tests indiquent. Cet écart finira par se réduire, je pense.
C'est pour ça que j'ai pris le 96 Go. Mais si la prochaine génération envoie du lourd, j'achèterai un modèle avec beaucoup de mémoire sans hésiter.
inferencer, c'est genre 2 fois plus rapide que LM Studio en pré-remplissage... essaie ça, utilise aussi la variante MLX. Ou alors, achète un DGX pour le pré-remplissage et utilise le M3U pour la génération... EXO montre que ce combo déchire... mais c'est limité à 128 Go de DGX.
Essayez d'utiliser vMLX et des modèles optimisés
https://vmlx.net/
Ça fonctionne très bien sur mon MacBook Pro M2 avec 64 Go de RAM Gemma4 fait 50-70 t/s
N'oubliez pas de mettre à jour régulièrement
Sinon, essayez oMLX https://omlx.ai/
Dans tous les cas, utilisez uniquement MLX sur Apple et vérifiez qu'il prend en charge Metal 4
N'utilisez pas llama, vllm, ollama et d'autres qui ne sont pas profondément intégrés au framework Apple
De plus, consultez l'IA pour les meilleurs réglages de chaque modèle Cela vous donnera un boost de 10 à 30 %
Tu prends un Mac parce que tu veux un Mac, et tu en prends un avec beaucoup de RAM parce que tu veux faire tourner des trucs qui demandent beaucoup de RAM ou parce que tu veux jouer avec des gros modèles, quoi. L'inférence LLM, c'est surtout un truc marrant que tu peux faire avec, même si de plus en plus, les modèles MoE peuvent être assez rapides.
Si tu t'attendais à des performances Nvidia d'un Mac sans carte graphique, t'as jamais vraiment cherché ce que t'achetais.
À chaque fois qu'on me demande "Est-ce que je devrais acheter un Mac pour l'inférence ?", je réponds un "non" catégorique, à moins que tu sois OK avec le fait que ce soit super lent.
Il n'a pas de vrai GPU. Il n'a pas de PCIe pour ajouter une carte graphique. Il ne peut pas faire de traitement de prompt qui vaille la peine.
Dans le contexte des LLMs, le Mac est malheureusement un jouet, pas un outil. Désolé que tu l'aies appris à tes dépens.
Merci d'avoir partagé cette expérience. En tant que quelqu'un qui a envisagé un Studio dans ce but, je pense que c'est une anecdote utile et concrète.
Combien de RAM utilises-tu ? Ça pourrait faire tourner un modèle assez costaud, même si la bande passante limite un peu la vitesse de sortie.
C'est l'inverse. Bonne bande passante et bonne vitesse de sortie, mais traitement des requêtes nul. Une fois qu'il y aura des algorithmes d'attention plus efficaces en utilisation courante, ça commencera à devenir plus utile, à mon avis.
Pour ceux qui savent, comment est le support Mac pour d'autres types d'inférence, comme l'audio et la vidéo ? Vitesse mise à part, y a-t-il un réel support ?
Combien tu as payé pour ça ? Quels modèles et quants utilises-tu ? T'as combien de contexte et quelle vitesse pp et tg tu obtiens ? Merci.
