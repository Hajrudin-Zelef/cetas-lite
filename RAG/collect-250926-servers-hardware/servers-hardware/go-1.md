---
id: collect-250926-servers-hardware/servers-hardware/go-1
title: "Go"
domain: servers-hardware
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "DeepSeek", "LongCat", "Meta", "Microsoft", "MiniMax", "Moonshot", "Z.ai", "xAI"]
dates: []
keywords: ["agent", "agents", "claude", "copilot", "deepseek", "glm", "grok", "grok 4", "kimi", "luna", "muse", "muse spark"]
source: docs/RAG/clean4/go.md
source_anchor: ""
source_lines: [1, 108]
sha256: 3917431ce8e28bd5a3ef35e949bffa05a600faf5f25e29fae5420fa64eff8e95
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

