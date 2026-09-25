---
id: collect-250926-servers-hardware/servers-hardware/go
title: "Go"
domain: servers-hardware
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "DeepSeek", "LongCat", "Meta", "Microsoft", "MiniMax", "Moonshot", "Z.ai", "xAI"]
dates: []
keywords: ["agent", "agents", "claude", "copilot", "deepseek", "glm", "gpt-5.6", "gpt-6", "gpu", "grok", "grok 4", "kimi"]
source: docs/RAG/clean4/go.md
source_anchor: ""
source_lines: [1, 330]
sha256: 0654302af40c39e5786e1e9c67f745c5e1b64ba0a3501e8b51b1adc0025a4931
---

# Go

Abonnement Ã bas coÃ»t pour les modÃ¨les de codage ouverts.

OpenCode Go est un abonnement Ã  bas coÃ»t Ã  **10 $/mois** qui vous donne un accÃ¨s fiable aux modÃ¨les de codage ouverts populaires.

Go fonctionne comme nâimporte quel autre fournisseur dans OpenCode. Vous vous abonnez Ã  OpenCode Go et obtenez votre clÃ© dâAPI. Câest **totalement facultatif** et vous nâavez pas besoin de lâutiliser pour utiliser OpenCode.

Il est conÃ§u principalement pour les utilisateurs internationaux et offre un accÃ¨s mondial stable.

Les modÃ¨les ouverts sont devenus vraiment performants. Ils atteignent dÃ©sormais des performances proches de celles des modÃ¨les propriÃ©taires pour les tÃ¢ches de codage. Et comme de nombreux fournisseurs peuvent les proposer de maniÃ¨re compÃ©titive, ils sont gÃ©nÃ©ralement beaucoup moins chers.

Cependant, obtenir un accÃ¨s fiable et Ã faible latence Ã ces modÃ¨les peut Ãªtre difficile. Les fournisseurs varient en qualitÃ© et en disponibilitÃ©.

Pour remÃ©dier Ã cela, nous avons fait plusieurs choses :

1. Nous avons testÃ© un groupe sÃ©lectionnÃ© de modÃ¨les ouverts et discutÃ© avec leurs Ã©quipes de la meilleure faÃ§on de les exÃ©cuter.
2. Nous avons ensuite travaillÃ© avec quelques fournisseurs pour nous assurer quâils Ã©taient correctement servis.
3. Enfin, nous avons Ã©valuÃ© les performances de la combinaison modÃ¨le/fournisseur et avons dressÃ© une liste que nous nous sentons Ã lâaise de recommander.

OpenCode Go vous donne accÃ¨s Ã  ces modÃ¨les pour **10 $/mois**.

OpenCode Go fonctionne comme nâimporte quel autre fournisseur dans OpenCode.

1. Vous vous connectez Ã  **OpenCode Zen** , vous vous abonnez Ã  Go et copiez votre clÃ© dâAPI.
2. Vous exÃ©cutez la commande `/connect` dans la TUI, sÃ©lectionnez`OpenCode Go` et collez votre clÃ© dâAPI.
3. ExÃ©cutez `/models` dans la TUI pour voir la liste des modÃ¨les disponibles via Go.

La liste actuelle des modÃ¨les comprend :

- **Grok 4.7**
- **Grok 4.6**
- **GLM-5.3-Flash**
- **GLM-5.3**
- **GLM-5.2**
- **GLM-5.1**
- **GPT 6 Luna**
- **GPT 5.6 Luna**
- **Kimi K3**
- **Kimi K2.7 Code**
- **Kimi K2.6**
- **LongCat-2.0**
- **MiMo-V2.6-Flash**
- **MiMo-V2.6-Pro**
- **MiMo-V2.5**
- **MiMo-V2.5-Pro**
- **MiniMax M3**
- **MiniMax M2.7**
- **Muse Spark 1.3 Contributor** (rÃ©gions limitÃ©es)
- **Muse Spark 1.2 Contributor** (rÃ©gions limitÃ©es)
- **Qwen3.8 Max**
- **Qwen3.8 Flash**
- **Qwen3.7 Max**
- **Qwen3.7 Plus**
- **Qwen3.6 Plus**
- **DeepSeek V4.1 Flash**
- **DeepSeek V4 Pro**
- **DeepSeek V4 Flash**
- **DeepSeek V4 Flash Vision Exp**
- **Hy4 preview**
- **Hy3**
- **Space Bunny Free** (pour une durÃ©e limitÃ©e)

La liste des modÃ¨les peut changer au fur et Ã mesure que nous en testons et en ajoutons de nouveaux.

OpenCode Go est conÃ§u pour OpenCode et dâautres agents de codage qui produisent des types de requÃªtes similaires. Le trafic est surveillÃ© afin de dÃ©tecter les abus qui dÃ©gradent lâexpÃ©rience des autres utilisateurs.

Votre client doit :

1. Envoyer le trafic habituel dâun agent de codage
2. Sâidentifier avec son propre agent utilisateur, tel que `my-coding-agent/1.0` , plutÃ´t
quâavec le nom gÃ©nÃ©rique dâun SDK ou dâune bibliothÃ¨que HTTP.
3. Envoyer un ID de session stable dans `x-opencode-session` pour chaque conversation afin que nous puissions optimiser le routage et
la mise en cache des prompts.

Outre OpenCode, le bon fonctionnement des clients suivants avec OpenCode Go a Ã©tÃ© validÃ©. Nous ne garantissons toutefois pas quâils continueront Ã fonctionner Ã lâavenir.

| Client | Prise en charge des sessions | 
|---|---|
| **Hermes** | Les builds contenant la PR #101864 envoient lâen-tÃªte sur les requÃªtes OpenCode principales et auxiliaires. Le correctif a Ã©tÃ© fusionnÃ© aprÃ¨s la v0.21.0 ; cette version seule ne lâinclut pas. | 
| **Claude Code** | Go reconnaÃ®t son en-tÃªte de session natif. Aucun wrapper dâen-tÃªte personnalisÃ© nâest nÃ©cessaire. | 
| **Codex** | Go reconnaÃ®t son en-tÃªte de session natif. Certaines versions et configurations de proxy lâomettent encore ; conservez lâen-tÃªte de session lors du transfert des requÃªtes. | 
| **ZCode** | Go reconnaÃ®t son en-tÃªte de session natif. Notre demande concernant `x-opencode-session` reste ouverte, mais il nâest plus nÃ©cessaire dâenvoyer spÃ©cifiquement cet en-tÃªte. | 
| **Pi** | Les builds actuels envoient les informations de session pour OpenCode. Mettez Ã jour les installations plus anciennes. | 
| **jcode** | Passez Ã  la version **v0.81.6 ou ultÃ©rieure** , qui inclut le correctif de lâen-tÃªte de session. | 
| **Kilo Code CLI** | Les builds contenant la PR #13752 rÃ©tablissent les en-tÃªtes de session OpenCode. Ce correctif concerne la CLI, pas lâextension VS Code. Consultez lâissue #13723. | 

Dans les versions que nous avons examinÃ©es, ces clients ne prennent pas en charge les sessions ou ne les prennent en charge que partiellement. Les rapports associÃ©s permettent de suivre les correctifs et les solutions de contournement.

| Client | Ãtat et suivi | 
|---|---|
| **DeepSeek Harness** | Les informations de session sont prÃ©sentes pour certains chemins de modÃ¨les, mais absentes pour dâautres. Nous reconnaissons son en-tÃªte natif ; il reste Ã lâenvoyer depuis tous les adaptateurs. Discussion #5495. | 
| **GitHub Copilot Chat** | La prise en charge automatique de lâen-tÃªte de session fait lâobjet de lâissue VS Code #334186. | 
| **Kimi Code** | La prise en charge automatique de lâen-tÃªte de session fait lâobjet de lâissue #3506. | 
| **MiMo Code** | Lâissue #2317 dispose dâun correctif proposÃ© dans la PR #2327, qui nâa pas encore Ã©tÃ© fusionnÃ©e. | 

Les limites dâutilisation sont dÃ©finies sous forme de montants mensuels en dollars. Le tableau ci-dessous indique la limite mensuelle et les prix des tokens pour chaque modÃ¨le.

Chaque modÃ¨le est soumis aux limites dâutilisation suivantes : 5 heures â 20 % de la limite mensuelle ; hebdomadaire â 50 % ; et mensuelle â 100 %.

Par exemple, si un modÃ¨le a une limite mensuelle de $60, vous pouvez utiliser jusquâÃ :

- **Limite de 5 heures** â $12 dâutilisation
- **Limite hebdomadaire** â $30 dâutilisation
- **Limite mensuelle** â $60 dâutilisation

Les prix des tokens sont indiquÃ©s par million de tokens.

| ModÃ¨le | Input | Output | Cached Read | Cached Write | Limite mensuelle | 
|---|---|---|---|---|---|
| GLM-5.3-Flash | $0.15 | $0.50 | $0.03 | - | **$60** | 
| GLM-5.3 | $1.40 | $4.40 | $0.26 | - | **$15** | 
| GLM-5.2 | $1.40 | $4.40 | $0.26 | - | **$60** | 
| GLM-5.1 | $1.40 | $4.40 | $0.26 | - | **$60** | 
| Kimi K3 | $3.00 | $15.00 | $0.30 | - | **$15** | 
| Kimi K2.7 Code | $0.95 | $4.00 | $0.19 | - | **$60** | 
| Kimi K2.6 | $0.95 | $4.00 | $0.16 | - | **$60** | 
| LongCat-2.0 | $0.30 | $1.20 | $0.006 | - | **$60** | 
| MiMo-V2.6-Flash | $0.14 | $0.28 | $0.0028 | - | **$60** | 
| MiMo-V2.6-Pro | $0.435 | $0.87 | $0.003625 | - | **$15** | 
| MiMo-V2.5 | $0.14 | $0.28 | $0.0028 | - | **$60** | 
| MiMo-V2.5-Pro | $0.435 | $0.87 | $0.003625 | - | **$15** | 
| MiniMax M3 | $0.30 | $1.20 | $0.06 | - | **$60** | 
| MiniMax M2.7 | $0.30 | $1.20 | $0.06 | $0.375 | **$60** | 
| MiniMax M2.5 | $0.30 | $1.20 | $0.06 | $0.375 | **$60** | 
| Muse Spark 1.3 Contributor | $0.10 | $0.20 | $0.002 | - | **$60** | 
| Muse Spark 1.2 Contributor | $0.10 | $0.20 | $0.002 | - | **$60** | 
| Qwen3.8 Max | $2.00 | $6.00 | $0.25 | $2.50 | **$15** | 
| Qwen3.8 Flash | $0.15 | $0.47 | $0.016 | $0.20 | **$30** | 
| Qwen3.7 Max | $2.50 | $7.50 | $0.50 | $3.125 | **$30** | 
| Qwen3.7 Plus (â¤ 256K tokens) | $0.40 | $1.60 | $0.04 | $0.50 | **$60** | 
| Qwen3.7 Plus (> 256K tokens) | $1.20 | $4.80 | $0.12 | $1.50 | **$60** | 
| Qwen3.6 Plus (â¤ 256K tokens) | $0.50 | $3.00 | $0.05 | $0.625 | **$60** | 
| Qwen3.6 Plus (> 256K tokens) | $2.00 | $6.00 | $0.20 | $2.50 | **$60** | 
| DeepSeek V4.1 Flash (Off-Peak) | $0.15 | $0.60 | $0.003 | - | ~~$15~~**$60**4x Â· Fin le 27 sept. | 
| DeepSeek V4.1 Flash (Peak) | $0.30 | $1.20 | $0.006 | - | ~~$15~~**$60**4x Â· Fin le 27 sept. | 
| DeepSeek V4 Pro (Off-Peak) | $0.66 | $1.98 | $0.022 | - | **$15** | 
| DeepSeek V4 Pro (Peak) | $1.32 | $3.96 | $0.044 | - | **$15** | 
| DeepSeek V4 Flash (Off-Peak) | $0.15 | $0.60 | $0.003 | - | **$30** | 
| DeepSeek V4 Flash (Peak) | $0.30 | $1.20 | $0.006 | - | **$30** | 
| DeepSeek V4 Flash Vision Exp (Off-Peak) | $0.15 | $0.60 | $0.003 | - | **$15** | 
| DeepSeek V4 Flash Vision Exp (Peak) | $0.30 | $1.20 | $0.006 | - | **$15** | 
| Hy4 preview | $0.834 | $2.501 | $0.042 | - | **$30** | 
| Hy3 | $0.14 | $0.58 | $0.035 | - | **$60** | 
| Space Bunny Free | Free | Free | Free | - | **IllimitÃ©**pour une durÃ©e limitÃ©e | 
| Grok 4.7 (â¤ 200K tokens) | $2.00 | $6.00 | $0.50 | - | **$15** | 
| Grok 4.7 (> 200K tokens) | $4.00 | $12.00 | $1.00 | - | **$15** | 
| Grok 4.6 (â¤ 200K tokens) | $2.00 | $6.00 | $0.50 | - | **$15** | 
| Grok 4.6 (> 200K tokens) | $4.00 | $12.00 | $1.00 | - | **$15** | 
| GPT 6 Luna (â¤ 272K tokens) | $0.10 | $0.50 | $0.01 | $0.125 | **$15** | 
| GPT 6 Luna (> 272K tokens) | $0.20 | $0.75 | $0.02 | $0.25 | **$15** | 
| GPT 5.6 Luna (â¤ 272K tokens) | $0.20 | $1.20 | $0.02 | $0.25 | **$15** | 
| GPT 5.6 Luna (> 272K tokens) | $0.40 | $1.80 | $0.04 | $0.50 | **$15** | 

**Space Bunny Free:** Gratuit pour une durÃ©e limitÃ©e.

**DeepSeek V4.1 Flash / V4 Pro / V4 Flash / V4 Flash Vision Exp:** Les heures Peak sont 01:00-04:00 et 06:00-10:00 UTC, du lundi au vendredi ; toutes les autres heures, y compris le week-end, sont Off-Peak. En savoir plus.

**DeepSeek V4 Flash Vision Exp:** Les images sont converties en tokens selon leurs dimensions et facturÃ©es comme tokens dâentrÃ©e avec les tokens de texte. En savoir plus.

Le tableau ci-dessous fournit une estimation du nombre de requÃªtes basÃ©e sur les habitudes dâutilisation typiques de Go :

| Model | requÃªtes par 5 heures | requÃªtes par semaine | requÃªtes par mois | 
|---|---|---|---|
| GLM-5.3-Flash | 6,320 | 15,790 | 31,580 | 
| GLM-5.3 | 220 | 540 | 1,080 | 
| GLM-5.2 | 880 | 2,150 | 4,300 | 
| GLM-5.1 | 880 | 2,150 | 4,300 | 
| Kimi K3 | 110 | 250 | 490 | 
| Kimi K2.7 Code | 1,350 | 3,380 | 6,750 | 
| Kimi K2.6 | 1,150 | 2,880 | 5,750 | 
| LongCat-2.0 | 11,400 | 28,600 | 57,200 | 
| MiMo-V2.6-Flash | 30,100 | 75,200 | 150,400 | 
| MiMo-V2.6-Pro | 3,250 | 8,150 | 16,300 | 
| MiMo-V2.5 | 30,100 | 75,200 | 150,400 | 
| MiMo-V2.5-Pro | 3,250 | 8,150 | 16,300 | 
| MiniMax M3 | 3,200 | 8,000 | 16,000 | 
| MiniMax M2.7 | 3,400 | 8,500 | 17,000 | 
| Muse Spark 1.3 Contributor | 45,300 | 113,300 | 226,600 | 
| Muse Spark 1.2 Contributor | 45,300 | 113,300 | 226,600 | 
| Qwen3.8 Max | 160 | 400 | 810 | 
| Qwen3.8 Flash | 5,400 | 13,500 | 27,000 | 
| Qwen3.7 Max | 170 | 420 | 840 | 
| Qwen3.7 Plus | 4,300 | 10,800 | 21,600 | 
| Qwen3.6 Plus | 3,300 | 8,200 | 16,300 | 
| DeepSeek V4.1 Flash 4x Â· Fin le 27 sept. | ~~6,500~~**26,000** | ~~16,250~~**65,000** | ~~32,500~~**130,000** | 
| DeepSeek V4 Pro | 1,050 | 2,600 | 5,200 | 
| DeepSeek V4 Flash | 13,000 | 32,500 | 65,000 | 
| DeepSeek V4 Flash Vision Exp | 6,500 | 16,250 | 32,500 | 
| Hy4 preview | 1,350 | 3,380 | 6,770 | 
| Hy3 | 4,300 | 10,750 | 21,500 | 
| Space Bunny Free | IllimitÃ© | IllimitÃ© | IllimitÃ© | 
| Grok 4.7 | 169 | 423 | 845 | 
| Grok 4.6 | 169 | 423 | 845 | 
| GPT 6 Luna | 4,230 | 10,560 | 21,130 | 
| GPT 5.6 Luna | 2,050 | 5,100 | 10,250 | 

Les estimations utilisent les nombres de tokens suivants par requÃªte ; lâutilisation rÃ©elle varie.

- Grok 4.7/4.6 â 390 tokens en entrÃ©e, 32,500 en cache, 120 tokens en sortie par requÃªte
- GLM-5.3-Flash â 1,000 tokens en entrÃ©e, 55,000 en cache, 200 tokens en sortie par requÃªte
- GLM-5.3/5.2/5.1 â 700 tokens en entrÃ©e, 52,000 en cache, 150 tokens en sortie par requÃªte
- GPT 6 Luna â 1,000 tokens en entrÃ©e, 50,000 en cache, 220 tokens en sortie par requÃªte
- GPT 5.6 Luna â 1,000 tokens en entrÃ©e, 50,000 en cache, 220 tokens en sortie par requÃªte
- Kimi K3 â 1,050 tokens en entrÃ©e, 76,500 en cache, 300 tokens en sortie par requÃªte
- Kimi K2.7/K2.6 â 870 tokens en entrÃ©e, 55,000 en cache, 200 tokens en sortie par requÃªte
- LongCat-2.0 â 920 tokens en entrÃ©e, 88,900 en cache, 200 tokens en sortie par requÃªte
- DeepSeek V4.1 Flash â 410 tokens en entrÃ©e, 71,300 en cache, 310 tokens en sortie par requÃªte
- DeepSeek V4 Pro â 750 tokens en entrÃ©e, 82,000 en cache, 290 tokens en sortie par requÃªte
- DeepSeek V4 Flash â 410 tokens en entrÃ©e, 71,300 en cache, 310 tokens en sortie par requÃªte
- DeepSeek V4 Flash Vision Exp â 410 tokens en entrÃ©e, 71,300 en cache, 310 tokens en sortie par requÃªte
- MiniMax M3 â 510 tokens en entrÃ©e, 56,000 en cache, 190 tokens en sortie par requÃªte
- MiniMax M2.7 â 300 tokens en entrÃ©e, 55,000 en cache, 125 tokens en sortie par requÃªte
- Muse Spark 1.3 Contributor â 620 tokens en entrÃ©e, 71,400 en cache, 300 tokens en sortie par requÃªte
- Muse Spark 1.2 Contributor â 620 tokens en entrÃ©e, 71,400 en cache, 300 tokens en sortie par requÃªte
- Qwen3.8 Max â 420 tokens en entrÃ©e, 66,000 en cache, 200 tokens en sortie par requÃªte
- Qwen3.8 Flash â 600 tokens en entrÃ©e, 58,000 en cache, 200 tokens en sortie par requÃªte
- Qwen3.7 Max â 420 tokens en entrÃ©e, 66,000 en cache, 200 tokens en sortie par requÃªte
- Qwen3.7 Plus â 500 tokens en entrÃ©e, 57,000 en cache, 190 tokens en sortie par requÃªte
- Qwen3.6 Plus â 500 tokens en entrÃ©e, 57,000 en cache, 190 tokens en sortie par requÃªte
- Hy4 preview â 830 tokens en entrÃ©e, 71,500 en cache, 295 tokens en sortie par requÃªte
- Hy3 â 830 tokens en entrÃ©e, 71,500 en cache, 295 tokens en sortie par requÃªte
- MiMo-V2.6-Flash â 830 tokens en entrÃ©e, 71,500 en cache, 295 tokens en sortie par requÃªte
- MiMo-V2.6-Pro â 790 tokens en entrÃ©e, 86,000 en cache, 305 tokens en sortie par requÃªte
- MiMo-V2.5 â 830 tokens en entrÃ©e, 71,500 en cache, 295 tokens en sortie par requÃªte
- MiMo-V2.5-Pro â 790 tokens en entrÃ©e, 86,000 en cache, 305 tokens en sortie par requÃªte

Vous pouvez suivre votre utilisation actuelle dans la **console**.

Les limites dâutilisation peuvent changer au fur et Ã mesure que nous tirons des enseignements des premiÃ¨res utilisations et des retours.

Si vous avez Ã©galement des crÃ©dits sur votre solde Zen, vous pouvez activer lâoption **Use balance** dans la console. Lorsquâelle est activÃ©e, Go se rabattra sur votre solde Zen aprÃ¨s que vous ayez atteint vos limites dâutilisation au lieu de bloquer les requÃªtes.

Avec Go, vous payez 10 $/mois, et lâutilisation mensuelle incluse varie selon le modÃ¨le.

Pour la plupart des modÃ¨les, nous y parvenons grÃ¢ce Ã des remises sur volume et Ã une capacitÃ© GPU rÃ©servÃ©e. Nous vous faisons ensuite bÃ©nÃ©ficier de ces Ã©conomies sous la forme dâune utilisation mensuelle plus Ã©levÃ©e.

Pour certains modÃ¨les, nous nâavons pas encore eu lâoccasion de nÃ©gocier une remise ou de les hÃ©berger Ã moindre coÃ»t, soit parce que le modÃ¨le est nouveau, soit parce que son tarif public est dÃ©jÃ rÃ©duit.

Pour ces modÃ¨les, vous obtenez tout de mÃªme un peu plus que si vous payiez directement les fournisseurs de modÃ¨les ; câest pourquoi leur utilisation mensuelle incluse est plus faible.

Vous pouvez Ã©galement accÃ©der aux modÃ¨les Go via les points de terminaison dâAPI suivants.

| ModÃ¨le | ID de modÃ¨le | Point de terminaison | Package AI SDK | 
|---|---|---|---|
| Grok 4.7 | grok-4.7 | `https://opencode.ai/zen/go/v1/responses` | `@ai-sdk/openai` | 
| Grok 4.6 | grok-4.6 | `https://opencode.ai/zen/go/v1/responses` | `@ai-sdk/openai` | 
| GPT 6 Luna | gpt-6-luna | `https://opencode.ai/zen/go/v1/responses` | `@ai-sdk/openai` | 
| GPT 5.6 Luna | gpt-5.6-luna | `https://opencode.ai/zen/go/v1/responses` | `@ai-sdk/openai` | 
| GLM-5.3-Flash | glm-5.3-flash | `https://opencode.ai/zen/go/v1/chat/completions` | `@ai-sdk/openai-compatible` | 
| GLM-5.3 | glm-5.3 | `https://opencode.ai/zen/go/v1/chat/completions` | `@ai-sdk/openai-compatible` | 
| GLM-5.2 | glm-5.2 | `https://opencode.ai/zen/go/v1/chat/completions` | `@ai-sdk/openai-compatible` | 
| GLM-5.1 | glm-5.1 | `https://opencode.ai/zen/go/v1/chat/completions` | `@ai-sdk/openai-compatible` | 
| Kimi K3 | kimi-k3 | `https://opencode.ai/zen/go/v1/chat/completions` | `@ai-sdk/openai-compatible` | 
| Kimi K2.7 Code | kimi-k2.7-code | `https://opencode.ai/zen/go/v1/chat/completions` | `@ai-sdk/openai-compatible` | 
| Kimi K2.6 | kimi-k2.6 | `https://opencode.ai/zen/go/v1/chat/completions` | `@ai-sdk/openai-compatible` | 
| LongCat-2.0 | longcat-2.0 | `https://opencode.ai/zen/go/v1/chat/completions` | `@ai-sdk/openai-compatible` | 
| DeepSeek V4.1 Flash | deepseek-v4.1-flash | `https://opencode.ai/zen/go/v1/chat/completions` | `@ai-sdk/openai-compatible` | 
| DeepSeek V4 Pro | deepseek-v4-pro | `https://opencode.ai/zen/go/v1/chat/completions` | `@ai-sdk/openai-compatible` | 
| DeepSeek V4 Flash | deepseek-v4-flash | `https://opencode.ai/zen/go/v1/chat/completions` | `@ai-sdk/openai-compatible` | 
| DeepSeek V4 Flash Vision Exp | deepseek-v4-flash-vision-exp | `https://opencode.ai/zen/go/v1/chat/completions` | `@ai-sdk/openai-compatible` | 
| MiMo-V2.6-Flash | mimo-v2.6-flash | `https://opencode.ai/zen/go/v1/chat/completions` | `@ai-sdk/openai-compatible` | 
| MiMo-V2.6-Pro | mimo-v2.6-pro | `https://opencode.ai/zen/go/v1/chat/completions` | `@ai-sdk/openai-compatible` | 
| MiMo-V2.5 | mimo-v2.5 | `https://opencode.ai/zen/go/v1/chat/completions` | `@ai-sdk/openai-compatible` | 
| MiMo-V2.5-Pro | mimo-v2.5-pro | `https://opencode.ai/zen/go/v1/chat/completions` | `@ai-sdk/openai-compatible` | 
| MiniMax M3 | minimax-m3 | `https://opencode.ai/zen/go/v1/messages` | `@ai-sdk/anthropic` | 
| MiniMax M2.7 | minimax-m2.7 | `https://opencode.ai/zen/go/v1/messages` | `@ai-sdk/anthropic` | 
| MiniMax M2.5 | minimax-m2.5 | `https://opencode.ai/zen/go/v1/messages` | `@ai-sdk/anthropic` | 
| Muse Spark 1.3 Contributor | muse-spark-1.3-contributor | `https://opencode.ai/zen/go/v1/responses` | `@ai-sdk/openai` | 
| Muse Spark 1.2 Contributor | muse-spark-1.2-contributor | `https://opencode.ai/zen/go/v1/responses` | `@ai-sdk/openai` | 
| Qwen3.8 Max | qwen3.8-max | `https://opencode.ai/zen/go/v1/messages` | `@ai-sdk/anthropic` | 
| Qwen3.8 Flash | qwen3.8-flash | `https://opencode.ai/zen/go/v1/messages` | `@ai-sdk/anthropic` | 
| Qwen3.7 Max | qwen3.7-max | `https://opencode.ai/zen/go/v1/messages` | `@ai-sdk/anthropic` | 
| Qwen3.7 Plus | qwen3.7-plus | `https://opencode.ai/zen/go/v1/messages` | `@ai-sdk/anthropic` | 
| Qwen3.6 Plus | qwen3.6-plus | `https://opencode.ai/zen/go/v1/messages` | `@ai-sdk/anthropic` | 
| Hy4 preview | hy4-preview | `https://opencode.ai/zen/go/v1/chat/completions` | `@ai-sdk/openai-compatible` | 
| Hy3 | hy3 | `https://opencode.ai/zen/go/v1/chat/completions` | `@ai-sdk/openai-compatible` | 
| Space Bunny Free | space-bunny-free | `https://opencode.ai/zen/go/v1/chat/completions` | `@ai-sdk/openai-compatible` | 

LâID de modÃ¨le dans votre configuration OpenCode utilise le format `opencode-go/<model-id>`. Par exemple, pour Kimi K3, vous utiliseriez `opencode-go/kimi-k3` dans votre configuration.

Vous pouvez rÃ©cupÃ©rer la liste complÃ¨te des modÃ¨les disponibles et leurs mÃ©tadonnÃ©es Ã partir de :

| ModÃ¨le | EntraÃ®nement des modÃ¨les | Conservation des donnÃ©es | 
|---|---|---|
| Grok 4.7 | Non utilisÃ© | 30 jours | 
| Grok 4.6 | Non utilisÃ© | 30 jours | 
| GPT 6 Luna | Non utilisÃ© | 30 jours | 
| GPT 5.6 Luna | Non utilisÃ© | 30 jours | 
| GLM-5.3-Flash | Non utilisÃ© | 0 jour | 
| GLM-5.3 | Non utilisÃ© | 0 jour | 
| GLM-5.2 | Non utilisÃ© | 0 jour | 
| GLM-5.1 | Non utilisÃ© | 0 jour | 
| Kimi K3 | Non utilisÃ© | 0 jour | 
| Kimi K2.7 Code | Non utilisÃ© | 0 jour | 
| Kimi K2.6 | Non utilisÃ© | 0 jour | 
| LongCat-2.0 | Non utilisÃ© | 0 jour | 
| MiMo-V2.6-Pro | Non utilisÃ© | 0 jour | 
| MiMo-V2.6-Flash | Non utilisÃ© | 0 jour | 
| MiMo-V2.5-Pro | Non utilisÃ© | 0 jour | 
| MiMo-V2.5 | Non utilisÃ© | 0 jour | 
| Qwen3.8 Max | Non utilisÃ© | 0 jour | 
| Qwen3.8 Flash | Non utilisÃ© | 0 jour | 
| Qwen3.7 Max | Non utilisÃ© | 0 jour | 
| Qwen3.7 Plus | Non utilisÃ© | 0 jour | 
| Qwen3.6 Plus | Non utilisÃ© | 0 jour | 
| MiniMax M3 | Non utilisÃ© | 0 jour | 
| MiniMax M2.7 | Non utilisÃ© | 0 jour | 
| Muse Spark 1.3 Contributor | Oui | Pas de ZDR | 
| Muse Spark 1.2 Contributor | Oui | Pas de ZDR | 
| DeepSeek V4.1 Flash | Non utilisÃ© | 0 jour | 
| DeepSeek V4 Pro | Non utilisÃ© | 0 jour | 
| DeepSeek V4 Flash | Non utilisÃ© | 0 jour | 
| DeepSeek V4 Flash Vision Exp | Non utilisÃ© | 0 jour | 
| Hy4 preview | Non utilisÃ© | 0 jour | 
| Hy3 | Non utilisÃ© | 0 jour | 
| Space Bunny Free | Non utilisÃ© | 0 jour | 

- **Grok 4.7/4.6:** Le ZDR dÃ©sactive dâimportantes fonctionnalitÃ©s API qui dÃ©pendent des donnÃ©es stockÃ©es, notamment Responses API avec Ã©tat, Files and Collections et Batch API. En savoir plus.
- **GPT 6 Luna / GPT 5.6 Luna:** Des journaux de surveillance des abus sont gÃ©nÃ©rÃ©s pour toute utilisation des fonctionnalitÃ©s API et conservÃ©s pendant un maximum de 30 jours. En savoir plus.
- **Muse Spark 1.3 Contributor:** Des tarifs de tokens fortement rÃ©duits en Ã©change de lâautorisation dâutiliser vos prompts et vos complÃ©tions pour entraÃ®ner de futurs modÃ¨les Meta. La disponibilitÃ© est limitÃ©e aux rÃ©gions autorisÃ©es par la Politique dâutilisation gÃ©ographique de Meta. En savoir plus.
- **Muse Spark 1.2 Contributor:** Des tarifs de tokens fortement rÃ©duits en Ã©change de lâautorisation dâutiliser vos prompts et vos complÃ©tions pour entraÃ®ner de futurs modÃ¨les Meta. La disponibilitÃ© est limitÃ©e aux rÃ©gions autorisÃ©es par la Politique dâutilisation gÃ©ographique de Meta. En savoir plus.
- **DeepSeek:** Lâaccord ZDR est renouvelÃ© chaque mois. Lâaccord actuel est valable jusquâau 30 septembre 2026.

Nous avons crÃ©Ã© OpenCode Go pour :

1. Rendre le codage par IA **accessible** Ã  un plus grand nombre de personnes avec un abonnement Ã  bas coÃ»t.
2. Fournir un accÃ¨s **fiable** aux meilleurs modÃ¨les de codage ouverts.
3. SÃ©lectionner des modÃ¨les qui sont **testÃ©s et Ã©valuÃ©s** pour une utilisation en tant quâagent de codage.
4. Nâavoir **aucun verrouillage exclusif** en vous permettant dâutiliser nâimporte quel autre fournisseur avec OpenCode Ã©galement.
