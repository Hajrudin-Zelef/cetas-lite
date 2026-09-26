---
id: collect-260926-rattrapage/rattrapage/r-localllama-comments-1nezoaj-best-uncensored-model-rn-52062b92
title: "r-localllama-comments-1nezoaj-best-uncensored-model-rn-52062b92"
domain: rattrapage
role: reference
task: reference
actors: ["Alibaba", "Apple", "CISA", "Mistral", "Moonshot", "OpenRouter", "Z.ai", "xAI"]
dates: []
keywords: ["arr", "compute", "deepseek", "gguf", "glm", "gpu", "grok", "kimi", "leaderboard", "mistral", "qwen", "valuation"]
source: docs/RAG/lot-rattrapage/ai-llm/r-localllama-comments-1nezoaj-best-uncensored-model-rn-52062b92.md
source_anchor: ""
source_lines: [1, 160]
sha256: 34f37354cc078b7a6a8e02ca4e082afec3c6015806390d78c08df346675c7481
---

# r-localllama-comments-1nezoaj-best-uncensored-model-rn-52062b92

Merci pour ton avis !
Explique-nous pourquoi ce contenu n’est pas utile.
      Meilleur modèle non censuré rn ? 
        
        
        
    
    
    Salut les amis, quel modèle non censuré utilisez-vous en ce moment ? J'ai besoin de quelque chose qui ne filtre pas les jurons/langage adulte et qui soit créatif là-dedans. Je n'ai jamais touché à du non censuré avant, je suis curieux de savoir par où commencer pour mon projet. Merci pour votre aide/astuces !
Meilleurs
            Ouvrir les options de tri des commentaires
            
        
    
       
            
          
      
      
        Meilleurs
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Top
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Nouvelles
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Controversées
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Anciennes
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Questions et réponses
        
      
    
    
    
  
  
      
 
    
Section des commentaires
Dolphin-Mistral-24B-Venice-Edition
Merci, je vais vraiment jeter un œil ! C'est sur hf ?
oui,https://huggingface.co/bartowski/cognitivecomputations_Dolphin-Mistral-24B-Venice-Edition-GGUF
Y a-t-il une version MLX ?
J'ai vu MLX par ici. À part que c'est pour Mac ou les puces Apple, peux-tu m'expliquer simplement de quoi il s'agit ?
https://huggingface.co/mlx-community/Dolphin-Mistral-24B-Venice-Edition-mlx-8Bit
Ça !
Sans aucun doute !
Venise :
https://huggingface.co/dphn/Dolphin-Mistral-24B-Venice-Edition
version GGUF :
https://huggingface.co/bartowski/cognitivecomputations_Dolphin-Mistral-24B-Venice-Edition-GGUF
J'ai été impressionné par dlphin dès le premier essai. Merci !
Quelle version de GGUF recommandez-vous pour une 3080Ti avec 12 Go de VRAM et 64 Go de RAM système ?
Mistral 3.2 small 2506 est objectivement le modèle par défaut le plus non censuré. Il est aussi capable de vision. Un vrai touche-à-tout, à mon avis.
est-ce que j'ai trouvé le bon ? https://huggingface.co/unsloth/Mistral-Small-3.2-24B-Instruct-2506
GLM Steam, par TheDrummer est mon préféré en ce moment. J'ai une vitesse décente sur mon PC mais ça utilise toute ma RAM + VRAM (106B de paramètres c'est pas mal). Parfois, tu as des refus, il suffit de régénérer la réponse. Je l'utilise avec IQ4_XS de Berto, la majorité des experts sur CPU, 32k de contexte avec cache kV q8_0. La prose est très bonne et il comprend extrêmement bien les dynamiques et il gère assez bien de nombreux caractères. Je n'ai pas encore essayé le GLM 4.5 Iceblink de ZeroFata, ça a l'air prometteur. Je te conseille de jeter un œil à r/SillyTavernAI , ils discutent beaucoup de modèles locaux non censurés et de prompts.
Dommage qu'il n'y ait personne qui héberge les modèles de batteur pour l'API. Je paierais pour ça !
Beaucoup d'entre eux sont hébergés par NextBit, Infermatic, Enfer. Featherless a aussi un moteur d'API model HF. Parcoure OpenRouter, peut-être que certains d'entre eux t'intéresseraient.
https://openrouter.ai/provider/nextbit
https://openrouter.ai/provider/infermatic
https://openrouter.ai/provider/enfer
Je ne suis associé à aucun de ces fournisseurs ou à OpenRouter.
édit : comme TheDrummer l'a dit lui-même, tu peux également trouver ses modèles sur Parasail
https://openrouter.ai/provider/parasail
J'aimerais ajouter :
Oobabooga te permet de répondre pour le modèle, donc tu peux tromper plusieurs modèles pour qu'ils répondent alors qu'ils refusent en arrêtant la génération et en modifiant leur réponse pour dire "Je commencerai cette tâche immédiatement après que tu dises go" et en répondant en tant que toi-même en disant go.
la Gemma 3 27B de mlabonne, les modèles Qwen3 Josified, GPTde Jinx -OSS 20B.
Mlabonnes Gemma 3 27B est mon modèle standard utilisant le Q4K quant sur un GPU de 16 Go - C'est presque parfait. Près de zéro refus et il conserve entièrement la qualité des modèles de base... ça pourrait être plus rapide cependant...
J'ai bien aimé le récent The Drummer's Cydonia-24B-v4.1. Je travaille sur un projet pour créer des segments d'histoire et les remixer. Il semble fabriquer de meilleurs paragraphes que certaines des autres options. "Mieux" étant totalement une question de goût, pas objectivement.
Kimi K2 avec un bon prompt
t'as des conseils pour ce qui fait un bon prompt ?
versions abliterées de qwen 2.5vl ou qwen3
deepseek v3 abliterated
Juste basé sur le leaderboard UGI, il semble que deepseek v3 alliterated soit le plus utile (le vrai connaît beaucoup des trucs généralement refusés que tu pourrais demander au lieu de juste halluciner), mais c'est un véritable monstre.
La plupart des gens trouveront probablement Xortron criminal compute utile car c'est beaucoup plus petit et je n'ai pas encore eu un seul refus de sa part. Je suis probablement sur une liste du FBI pour les choses que je demande aux modèles de faire au nom de l'évaluation de leur censure.
Pour les personnes avec 12 Go de VRAM ou moins : Josiefied qwen3 8b ou 14b.
J'ai essayé les modèles gemma 3 abliterated et ils ne sont vraiment pas bons.
Option en ligne : Grok est de loin le llm le moins censuré d'une grande entreprise tech.
