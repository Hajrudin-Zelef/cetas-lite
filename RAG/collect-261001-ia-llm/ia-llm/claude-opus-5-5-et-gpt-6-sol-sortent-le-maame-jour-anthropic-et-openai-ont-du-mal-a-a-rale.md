---
id: collect-261001-ia-llm/ia-llm/claude-opus-5-5-et-gpt-6-sol-sortent-le-maame-jour-anthropic-et-openai-ont-du-mal-a-a-rale
title: "claude-opus-5-5-et-gpt-6-sol-sortent-le-maame-jour-anthropic-et-openai-ont-du-mal-a-a-ralentir-a-l-i"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "Google", "Meta", "OpenAI", "SpaceX", "xAI"]
dates: []
keywords: ["claude", "gpt-6", "sol", "agents", "agi", "astra", "benchmarks", "chatgpt", "fable 5", "gemini", "gemini 4", "gpt-5.6"]
source: docs/RAG/collect-261001-ia-llm/claude-opus-5-5-et-gpt-6-sol-sortent-le-maame-jour-anthropic-et-openai-ont-du-mal-a-a-ralentir-a-l-i.md
source_anchor: ""
source_lines: [1, 78]
sha256: cd1aa2eb0b2695780caa6ab3fbe233e48bc69085bda95839e59ff673b83333b0
---

# claude-opus-5-5-et-gpt-6-sol-sortent-le-maame-jour-anthropic-et-openai-ont-du-mal-a-a-ralentir-a-l-i

Il est rare de voir les deux leaders d’un secteur lancer simultanÃ©ment un nouveau produit. Pourtant, le 22 septembre, Anthropic a lancÃ© Claude Opus 5.5, qui prend la tÃªte de la plupart des classements, et OpenAI, qui estimait dÃ©but septembre Ãªtre entrÃ© dans Â« *l’Ã¨re de l’AGI* Â» avec GPT-6 Astra, a lancÃ© GPT-6 Sol et GPT-6 Luna. L’ironie est que les deux groupes ont appelÃ© Ã  ralentir la course Ã  l’IA la semaine derniÃ¨re et qu’ils n’avaient pas lancÃ© de nouveaux modÃ¨les en mÃªme temps depuis le 14 mars 2023, quand les LLM Ã©taient encore Ã  leurs dÃ©buts (GPT-4 et le premier Claude). 

Au programme : de meilleurs rÃ©sultats, Ã©videmment. Mais surtout des prix en baisse. Avec GPT-6 Luna, facturÃ© 0,10 dollar le million de tokens en entrÃ©e, OpenAI s’attaque mÃªme aux modÃ¨les open source et Ã l’IA locale : Ã ce tarif, seules les personnes en quÃªte de confidentialitÃ© ont intÃ©rÃªt Ã configurer une machine en local.

| Prix API (par million de tokens) | EntrÃ©e | Sortie | 
|---|---|---|
| **GPT-6 Luna** | 0,10 $ | 0,50 $ | 
| **GPT-6 Sol** | 2 $ (4 $ avant) | 10 $ (20 $ avant) | 
| **Claude Opus 5.5** | 4 $ (5 $ avant) | 20 $ (25 $ avant) | 
| **GPT-6 Astra et Claude Fable 5.1** | 10 $ | 50 $ | 

## Claude Opus 5.5 : le niveau de Fable 5.1 pour 40 % moins cher qu’Opus 5

Le premier modÃ¨le de la famille Claude 5.5 est Opus 5.5. Selon Anthropic, il atteint le niveau de Claude Fable 5.1 sur la plupart des tÃ¢ches, malgrÃ© un prix bien plus bas.

Sur Terminal-Bench 4.0 (programmation en ligne de commande), il obtient 66,4 %, contre 57,9 % pour GPT-6 Astra et 55,8 % pour Fable 5.1. Sur GDPval-AA, qui Ã©value des tÃ¢ches professionnelles dans 44 mÃ©tiers, il atteint 1 846 points Elo, contre 1 735 pour Fable 5.1 et 1 542 pour Astra. GPT-6 Astra garde l’avantage sur deux tests, l’automatisation de tÃ¢ches en entreprise (41,4 % contre 40 %) et la recherche scientifique (64,6 % contre 58,7 %). Mais les benchmarks ne veulent plus dire grand chose Ã ce niveau.

Opus coÃ»te 2,5 fois moins cher que Fable mais le dÃ©passe dans la plupart des tests, ce qui pose la question de l’intÃ©rÃªt du modÃ¨le haut de gamme d’Anthropic jusqu’Ã la sortie du prochain Fable 5.5 (et donc la course constante Ã l’IA). L’entreprise reconnaÃ®t elle-mÃªme que l’Ã©cart rÃ©el entre les deux est plus faible que ce que montrent les chiffres. Le prix, lui, baisse : 4 dollars par million de tokens en entrÃ©e et 20 dollars en sortie, contre 5 et 25 dollars pour Opus 5. Les abonnÃ©s Pro, Max et Team profitent aussi de limites d’usage relevÃ©es sur cinq heures et d’une rÃ©initialisation gratuite de ces limites qu’ils peuvent garder de cÃ´tÃ© et dÃ©clencher quand ils le souhaitent, Ã la maniÃ¨re de ce que propose dÃ©jÃ OpenAI.

Notons tout de mÃªme que, dans Claude, le sÃ©lecteur d’effort est dÃ©sormais rÃ©glÃ© sur Â« Moyen Â» par dÃ©faut, et non plus sur Â« ÃlevÃ© Â». Une partie des Ã©conomies annoncÃ©es vient de lÃ , alors que les benchmarks sont rÃ©alisÃ©s en effort maximal. Il n’est plus possible non plus de dÃ©sactiver le mode rÃ©flexion. Enfin, Opus 5.5 hÃ©rite des garde-fous de Fable 5.1 : les demandes de cybersÃ©curitÃ© sont redirigÃ©es vers Opus 4.8, celles de biologie vers Opus 5.

### Claude se soumet Ã la rÃ¨glementation europÃ©enne : Opus 5.5 tatoue ses textes

Autre changement, Opus 5.5 est le premier Opus lancÃ© avec un marquage de ses contenus, comme l’impose l’AI Act depuis le 2 aoÃ»t 2026. Anthropic dit appliquer ce marquage partout dans le monde, et pas seulement en Europe, faute de savoir le limiter par rÃ©gion. L’outil de dÃ©tection est rÃ©servÃ© aux rÃ©gulateurs, aux mÃ©dias, aux fact-checkers et aux chercheurs. Sans surprise, Anthropic prÃ©vient qu’il ne s’agit pas d’une preuve infaillible.

## GPT-6 Sol et Luna : OpenAI divise ses prix par deux

Trois semaines aprÃ¨s Astra, OpenAI complÃ¨te sa famille GPT-6 avec GPT-6 Sol, son modÃ¨le haut/milieu de gamme et GPT-6 Luna, en entrÃ©e de gamme, tandis qu’Astra reste le modÃ¨le le plus puissant. Contrairement Ã la gÃ©nÃ©ration GPT-5.6, il n’y a pas de version Terra pour l’instant.

Comme chez Anthropic, les prix baisses. Les deux nouveaux modÃ¨les coÃ»tent deux fois moins cher que leurs prÃ©dÃ©cesseurs dans l’API : 2 et 10 dollars pour Sol, 0,10 et 0,50 dollar pour Luna.

OpenAI compare ses modÃ¨les Ã ceux d’Anthropicâ¦ mais pas au nouveau Opus 5.5, Ã©videmment.

Sur AutomationBench, Sol en effort xhigh atteint 33,2 % pour 0,27 dollar par tÃ¢che, devant Fable 5.1 (31,4 %) et Opus 5 (26,9 %), qui coÃ»te 11 fois plus cher par tÃ¢che. Opus 5.5 affiche 40 % sur ce mÃªme test selon Anthropic : il reste devant, mais pour un prix plus Ã©levÃ©. En programmation, Sol obtient 68,8 % sur DeepSWE, Ã 1,1 point de Fable 5 pour environ 80 % de coÃ»t en moins. Luna atteint 66,6 %, au niveau d’Opus 5 en effort moyen. OpenAI annonce aussi deux fois moins d’erreurs factuelles qu’avec GPT-5.6 Sol, selon une Ã©valuation interne.

OpenAI glisse au passage un tacle Ã son rival : le coÃ»t de Fable 5.1 serait sous-estimÃ©, car il n’inclut pas les bascules vers Opus 5, survenues sur environ 40 % des tÃ¢ches de ce test. Un mÃ©canisme qu’utilise dÃ©sormais aussi Opus 5.5.

GPT-6 Sol et GPT-6 Luna arrivent dÃ¨s aujourd’hui dans ChatGPT Work et Codex pour les abonnÃ©s Plus, Pro, Business, Enterprise et Edu. Les utilisateurs gratuits et Go ont accÃ¨s Ã Luna dans l’application de bureau.

## Grok 4.7 publiÃ© la veille, Google aux abonnÃ©s absents

La veille, xAI, dÃ©sormais intÃ©grÃ© Ã SpaceX, avait lancÃ© Grok 4.7â¦ alors qu’Elon Musk appelait aussi Ã ralentir la course Ã l’IA. FacturÃ© 2 dollars en entrÃ©e et 6 dollars en sortie, il reste un cran en dessous des modÃ¨les d’Anthropic et OpenAI. Elon Musk lui-mÃªme le situait Ã peu prÃ¨s au niveau de Claude Opus 5.

Reste Google, portÃ© disparu. L’entreprise a promis que Gemini 3.5 Pro arriverait en juinâ¦ mais ses mauvaises performances le condamnent progressivement Ã ne jamais sortir. L’entreprise prÃ©fÃ¨re dÃ©sormais parler de Gemini 4, en cours d’entraÃ®nement, mais personne ne sait si elle arrivera Ã revenir au niveau d’Anthropic ou OpenAI, qui mise dÃ©sormais tout sur les agents, ou mÃªme Meta et son Muse qui en convainc plus d’un.

+ rapide, + pratique, + exclusif

ZÃ©ro publicitÃ©, fonctions avancÃ©es de lecture, articles rÃ©sumÃ©s par l'I.A, contenus exclusifs et plus encore.

DÃ©couvrez les nombreux avantages de Numerama+.

Vous avez lu **0 articles** sur Numerama ce mois-ci

            **Tout le monde n'a pas les moyens de payer** pour l'information. 

            C'est pourquoi nous maintenons notre journalisme ouvert Ã  tous.
        

            **Mais si vous le pouvez,**

            voici trois bonnes raisons de soutenir notre travail :
        

- 
                1
                Numerama+ contribue Ã  offrir **une expÃ©rience gratuite Ã  tous les lecteurs de Numerama** .
- 
                2
                Vous profiterez d'une **lecture sans publicitÃ©** , de nombreuses**fonctions avancÃ©es de lecture et des contenus exclusifs** .
- 
                3
                **Aider Numerama dans sa mission** : comprendre le prÃ©sent pour anticiper l'avenir.

**Si vous croyez en un web gratuit** et Ã  une information de qualitÃ© accessible au plus grand nombre, rejoignez Numerama+.

Toute l'actu tech en un clin d'Åil

Ajoutez Numerama Ã votre Ã©cran d'accueil et restez connectÃ©s au futur !
