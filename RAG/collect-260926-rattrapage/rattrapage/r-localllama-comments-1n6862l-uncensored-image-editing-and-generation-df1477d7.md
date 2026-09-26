---
id: collect-260926-rattrapage/rattrapage/r-localllama-comments-1n6862l-uncensored-image-editing-and-generation-df1477d7
title: "r-localllama-comments-1n6862l-uncensored-image-editing-and-generation-df1477d7"
domain: rattrapage
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "Apple", "Meta"]
dates: []
keywords: ["claude", "diffusion", "gpu", "qwen"]
source: docs/RAG/lot-rattrapage/ai-llm/r-localllama-comments-1n6862l-uncensored-image-editing-and-generation-df1477d7.md
source_anchor: ""
source_lines: [1, 150]
sha256: d373db9008d0bb82ec138e24534a979f15a5d048b88254532ee43dc572138c91
---

# r-localllama-comments-1n6862l-uncensored-image-editing-and-generation-df1477d7

Uncensored image editing and generation ? 
        
        
        
    
    
    I have been enjoying Imagen for image editing a lot but it' is heavily censored which can be very annoying. What is the best uncensored local image editing and generation tool?
 
     Publication verrouillée. Il n’est pas possible de publier de nouveaux commentaires.  
        
          
        
        
        
        
         Meilleurs
            Ouvrir les options de tri des commentaires
            
        
    
       
            
          
      
      
        Meilleurs
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Top
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Nouvelles
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Controversées
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Anciennes
        
      
    
    
    
  
  
      
 
    
        
    
       
            
          
      
      
        Questions et réponses
        
      
    
    
    
  
  
      
 
    
Section des commentaires
In terms of image generation,Chroma 
civit.ai
all the loras, models, guides etc.
last I checked it was the best resource, worth making an account.
Getcomfyui 
there might be easier options, but that's what im using right now. started with A1111
everyone out here recommending the biggest and most hardware intense new models, whenstable diffusion 1.5 
I must say, SD1.5 and SDXL are not easier to get you started, in fact they are harder to use and requires a lot of experience to use correctly.Qwen image/edit 
way easier to get started. Download invoke.ai and it holds your hand. Or https://github.com/lllyasviel/Fooocus , which is even easier.
Getting the perfect image and working on it tediously? yea they arent gonna be as "easy", but i have yet to find a non-comfyui way to run the models others recommend, and comfyui is everything but easay to get started in.
SD has no real image editing.
have you ever used it??? what are you talking about
I built a really simple image editor tool with Qwen Image Edit, a FastAPI Python backend and a Swift UI frontend for iOS. Pull an image from your camera, type in a prompt for how you want the image changed ('take the main person and put them on a forest trail', 'change the persons shirt to red') in an hour or so using Claude Code. Since it was for my own use internal to my network (actually accessible via my Tailsnet) I didn't add security protections, which is why I haven't released it.
Generation should be easily doable with Qwen Image (I'm considering making it so that if you _haven't_ selected an image, it switches to a Qwen Image model, but that'd be a bit of a switching delay). The key thing is that it's not a service I run, or something like that. (I wouldn't DARE run something like that on a public endpoint.) It's a Python service you would run on your home system with enough GPU, and the app has a spot where you can put the URL for your server, so it can talk to it. So it's pretty appropriate for LocalLLaMa. 🤣
Truthfully, just using the sample code for Qwen Image Edit, and a bit of time with Claude Code or some other good coding model, and I imagine you could replicate it easily.
If you can't share the code, could you share an MD file claude makes that has the detailed architecture and description of the project, stack used etc. Should provide a head start to people who want to replicate your project, specially the vibe coders.
Yeah, that was the prodding I needed to just put the damn thing out there. :)
https://www.reddit.com/r/LocalLLaMA/comments/1n6hk90/image_editing_app_with_qwen_image_edit_and_an_ios/
Enjoy, and feel free to let me know what you think!
Try out qwen image
Oi has Mrs. Potts given you a licence for that?
Imagen is slick but yeah, the censorship makes it frustrating. I switched over tokalon ai 
That's a good idea
