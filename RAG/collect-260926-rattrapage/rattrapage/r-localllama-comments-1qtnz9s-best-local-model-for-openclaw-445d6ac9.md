---
id: collect-260926-rattrapage/rattrapage/r-localllama-comments-1qtnz9s-best-local-model-for-openclaw-445d6ac9
title: "Best Local Model for Openclaw"
domain: rattrapage
role: reference
task: reference
actors: ["AMD", "Alibaba", "Meta", "Z.ai"]
dates: []
keywords: ["glm", "llama", "llama.cpp", "memory", "opus 4", "qwen"]
source: docs/RAG/lot-rattrapage/ai-llm/r-localllama-comments-1qtnz9s-best-local-model-for-openclaw-445d6ac9.md
source_anchor: ""
source_lines: [1, 92]
sha256: ee499d7f97b1452b404a0fed728f1748ba6a7dbf310de8329e720d5cfe58d793
---

# Best Local Model for Openclaw

*Source : https://www.reddit.com/r/LocalLLaMA/comments/1qtnz9s/best_local_model_for_openclaw/?tl=fr*
*Auteur : u/FeiX7 | Score : 11 | r/LocalLLaMA*

I have recently tried gpt-oss 20b for openclaw and it performed awfully...

openclaw requires so much context and small models intelligence degrades with such amount of context.

any thoughts about it and any ideas how to make the local models to perform better?

---

## Commentaires

**Prior-Combination473** (score 1):

Yeah the context degradation is brutal with smaller models - have you tried chunking the context or using a sliding window approach? 🤔 Might help keep the important stuff in focus without overwhelming the model 💀

  **FeiX7** (score 0):

  tried but still, model gets too confused
  
  I am now experimenting with my own project how to make use of context and tools effectively.  
  my core idea is to distill knowledge from bigger model to small one on-go  
  like for example if I ask openclaw for simple task, like tweet this message or translate this thing or text someone on whatsup, why I should use the opus 4.5 to do that, when even 4b model can do that?  
  so basically pattern is a "how-to-do thing with step by step instructions" so model should not think about usage of the skills and tools, he just reads instructions, extracts context from the query user send. and after success of the task we just compress the information about the instruction into the new chat and that's it )))
  
  I am interested what other thinks about it.  
  I wanted to make plugin for openclaw, but I guess experimenting from scratch will be better

    **Single_Foundation_40** (score 2):

    exactly what i want, an LLM that can follow instructions. I dont need it to be a math genius or able to code a simulation for a fractioning tower for heavy crude oil.

**iliaghp** (score -2):

Lmao I was just searching google for this.

**FPham** (score 12):

LOL, looking at what people pay for openclaw per day in API fees, it seems it sends so much data that anything local would just get lost in the sea of instructions. I tried it with lm studio. Qwen 3 was reasonably ok-ish - by that I mean it talked to me and didn't get in a total loop (GLM flash was lost). it could read file from workspace, but it would not write anything no matter how much I bribed it with bananas. .I really don't know what I would use it for in this state, it's going to mess up everything it touches. I'd say 70b and up, maybe that would work?

  **FeiX7** (score -5):

  with LMstudio it is even worse, I tried with it too and meh,  
  I guess maintainers don't care about local or small models, or even about documentation how to use claw with them...  
  biggest issue for me.
  
  giving API providers such a unique data, unbelievable.

    **DataGOGO** (score 2):

    Bro, you can’t honestly expect for this to work well on windows and LMstudio…

    **FPham** (score 2):

    I thought it is clawdbod interpretting the tools not LM Studio, because LM studio is just remote text LLM.

  **Klutzy-Snow8016** (score 2):

  I came here to recommend GLM 4.7 Flash, since it seems competent enough so far, but I see it performed really poorly for you, so I guess YMMV? I haven't used it for anything serious, though.

    **OkAbroad3112** (score 1):

    qwen3-coder-next responds better to calling tools and, in case of error, it can retry better.

  **Holiday_Purpose_3166** (score 5):

  The issue with LM Studio it's it always up to date with latest llama.cpp.
  
  GLM 4.7 Flash has been an amazing performer.

    **lolwutdo** (score 2):

    Did you mean "isn't always up to date?"
    
    Cause LMstudio definitely does not have the latest runtime.

    **Potential_Block4598** (score 2):

    How is update with llama.cpp is a bad thing exactly ?!

    **FeiX7** (score 2):

    same happens for runtimes as well, like ROCm.

  **mjuevos** (score 2):

  thanks. good info. glm4.7 flash has been a good chatbot, but cant do anything.. including committing things to memory [ it actually lies about this over and over, saying it made a file but when you go and look, nope ]. 
  
  will try gpt-oss:120b but i only have a 4070super, will it work?
