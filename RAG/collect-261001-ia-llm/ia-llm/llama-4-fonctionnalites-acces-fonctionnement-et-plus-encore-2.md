---
id: collect-261001-ia-llm/ia-llm/llama-4-fonctionnalites-acces-fonctionnement-et-plus-encore-2
title: "llama-4-fonctionnalites-acces-fonctionnement-et-plus-encore"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "DeepSeek", "Google", "Hugging Face", "Meta", "Microsoft", "Mistral", "OpenAI"]
dates: []
keywords: ["llama", "benchmark", "benchmarks", "claude", "deepseek", "distillation", "gemini", "gpu", "mistral", "multimodal", "open-weight", "qwen"]
source: docs/RAG/collect-261001-ia-llm/llama-4-fonctionnalites-acces-fonctionnement-et-plus-encore.md
source_anchor: ""
source_lines: [58, 133]
sha256: bbd719fbcc84969405b5a8cb181faea769ef85d3a85986054a6c57fac0416c67
---

# llama-4-fonctionnalites-acces-fonctionnement-et-plus-encore

Behemoth compte 288 milliards de paramètres actifs, organisés via 16 experts, pour un total avoisinant les 2 billions de paramètres. Meta a conçu une infrastructure d’entraînement entièrement nouvelle pour supporter Behemoth à cette échelle. Elle introduit de l’apprentissage par renforcement asynchrone, un échantillonnage du programme basé sur la difficulté des prompts, et une nouvelle fonction de perte de distillation qui équilibre dynamiquement cibles « molles » et « dures ».

Le post-entraînement de Behemoth a également nécessité une recette différente. Meta a écarté plus de 95 % des exemples SFT pour se concentrer sur des prompts difficiles et a orienté l’apprentissage par renforcement vers des scénarios complexes de raisonnement, de codage et multilingues. L’échantillonnage à partir d’instructions système variées a aidé le modèle à généraliser, tandis qu’un filtrage dynamique éliminait les prompts à faible valeur durant l’entraînement RL.

## Benchmarks de Llama 4

Meta a publié des résultats de benchmarks internes pour chacun des modèles Llama 4, les comparant aux précédentes variantes Llama ainsi qu’à plusieurs modèles open-weight concurrents et modèles de pointe.

Dans cette section, je vous présente les points saillants des benchmarks pour Scout, Maverick et Behemoth, selon les chiffres de Meta. Comme toujours, prudence avec les benchmarks auto-reportés ; ils offrent néanmoins un premier aperçu utile des performances de chaque modèle selon les tâches et de leur position dans le paysage actuel. Commençons par Scout.

### Benchmarks de Llama Scout

Llama 4 Scout s’en sort très bien sur un mélange de benchmarks de raisonnement, de code et multimodaux — d’autant plus au vu de son nombre de paramètres actifs réduit et de son empreinte sur un seul GPU.

Source : MetaAI

Sur la compréhension d’images, Scout devance ses concurrents : 88,8 sur ChartQA et 94,4 sur DocVQA (test), mieux que Gemini 2.0 Flash-Lite (73,0 et 91,2) et au niveau ou légèrement au-dessus de Mistral 3.1 et Gemma 3 27B.

Sur les benchmarks de raisonnement visuel comme MMMU (69,4) et MathVista (70,7), il mène aussi la danse côté open-weight, devant Gemma 3 (64,9 ; 67,6), Mistral 3.1 (62,8 ; 68,9) et Gemini Flash-Lite (68,0 ; 57,6).

En code, Scout obtient 32,8 sur LiveCodeBench, devant Gemini Flash-Lite (28,9) et Gemma 3 27B (29,7), mais légèrement derrière Llama 3.3 (33,3). Ce n’est pas un modèle prioritairement axé code, mais il tient son rang.

Sur la connaissance et le raisonnement, Scout atteint 74,3 sur MMLU Pro et 57,2 sur GPQA Diamond, surpassant tous les autres modèles open-weight sur les deux. Ces benchmarks favorisent le raisonnement long et multi-étapes, la performance de Scout est donc notable, surtout à cette échelle.

Enfin, ses capacités long-contexte montrent un vrai potentiel. Sur MTOB (Massive Textual Overlap Benchmark), qui évalue la capacité à traduire entre l’anglais et le KGV, une langue à faibles ressources, il obtient 42,2/36,6 sur le test demi-livre et 39,7/36,3 sur le test livre entier. Sur le demi-livre, Gemini 2.0 Flash-Lite garde un léger avantage avec 42,3, mais Scout comble l’écart sur le livre entier, dépassant les 35,1/30,0 de Gemini.

### Benchmarks de Llama Maverick

Maverick est le modèle le plus équilibré de la gamme Llama 4 — et les benchmarks le confirment. Sans viser les extrêmes de contexte de Scout ni l’échelle brute de Behemoth, il reste constant dans toutes les catégories clés : raisonnement multimodal, codage, compréhension linguistique et rétention long-contexte.

Source : MetaAI

En raisonnement visuel, Maverick obtient 73,4 sur MMMU et 73,7 sur MathVista, devant Gemini 2.0 Flash (71,7 et 73,1) et GPT-4o (69,1 et 63,8). Sur ChartQA (compréhension d’images), il atteint 90,0, légèrement au-dessus des 88,3 de Gemini et nettement au-dessus des 85,7 de GPT-4o. Sur DocVQA, Maverick atteint 94,4, à l’égal de Scout et devant les 92,8 de GPT-4o.

En codage, Maverick obtient 43,4 sur LiveCodeBench, au-dessus de GPT-4o (32,3), de Gemini Flash (34,5) et proche des 45,8 de DeepSeek v3.1.

Sur le raisonnement et la connaissance, Maverick atteint 80,5 sur MMLU Pro et 69,8 sur GPQA Diamond, là encore devant Gemini Flash (77,6 et 60,1) et GPT-4o (pas de score MMLU Pro, 53,6 sur GPQA). DeepSeek v3.1 garde 0,7 point d’avance sur MMLU Pro.

Maverick se montre également performant en compréhension multilingue, avec 84,6 sur Multilingual MMLU, légèrement au-dessus des 81,5 de Gemini. Un atout pour les développeurs qui travaillent sur plusieurs langues ou zones géographiques.

Sur les évaluations long-contexte (MTOB), Maverick obtient 54,0/46,4 sur le demi-livre et 50,8/46,7 sur le livre entier — nettement devant les 48,4/39,8 et 45,5/39,6 de Gemini. Ces scores suggèrent que, même s’il ne met pas autant en avant sa longueur de contexte que Scout, il profite bel et bien de sa fenêtre étendue.

### Benchmarks de Llama Behemoth

Behemoth n’est pas encore publié, mais ses chiffres de benchmark valent le détour.

Source : MetaAI

Sur les benchmarks à forte composante STEM, Behemoth excelle. Il atteint 95,0 sur MATH-500 — au-dessus de Gemini 2.0 Pro (91,8) et nettement devant Claude Sonnet 3.7 (82,2). Sur MMLU Pro, Behemoth marque 82,2, quand Gemini Pro obtient 79,1 (pas de score rapporté pour Claude). Et sur GPQA Diamond, un autre benchmark qui valorise la profondeur factuelle et la précision, Behemoth atteint 73,7, devant Claude (68,0), Gemini (64,7) et GPT-4.5 (71,4).

En compréhension multilingue, Behemoth obtient 85,8 sur Multilingual MMLU, légèrement devant Claude Sonnet (83,2) et GPT-4.5 (85,1). Des scores importants pour les développeurs à l’échelle mondiale, et Behemoth mène actuellement cette catégorie.

En raisonnement visuel, Behemoth atteint 76,1 sur MMMU, devant Gemini (71,8), Claude (72,7) et GPT-4.5 (74,4). Même si ce n’est pas son axe principal, il reste compétitif face aux modèles multimodaux leaders.

En génération de code, Behemoth obtient 49,4 sur LiveCodeBench. Un score nettement supérieur à Gemini 2.0 Pro (36,0).

## Comment accéder à Llama 4

Llama 4 Scout et Llama 4 Maverick sont tous deux disponibles dès maintenant sous la licence open-weight de Meta. Vous pouvez les télécharger directement depuis le site Llama officiel ou via Hugging Face.

Pour accéder aux modèles via les services de Meta, vous pouvez interagir avec Meta AI sur plusieurs plateformes : WhatsApp, Messenger, Instagram et Facebook. L’accès nécessite actuellement une connexion avec un compte Meta, et il n’existe pas d’endpoint API autonome pour Meta AI — du moins pas encore.

Si vous envisagez d’intégrer les modèles à vos propres applications ou à votre infrastructure, gardez en tête la clause de licence : si votre produit ou service dépasse les 700 millions d’utilisateurs actifs mensuels, vous devrez obtenir une autorisation distincte de Meta. Les modèles restent par ailleurs utilisables pour la recherche, l’expérimentation et la plupart des usages commerciaux.

## Conclusion

Scout introduit une longueur de contexte inédite sur un seul GPU. Maverick rivalise avec des modèles plus grands sur les tâches de raisonnement, de code et multimodales. Et Behemoth, encore en entraînement, montre comment des modèles enseignants peuvent façonner des variantes plus efficaces et déployables.

L’espace open-weight n’a jamais été aussi concurrentiel. DeepSeek, Qwen, Gemma, et bientôt OpenAI, avancent tous avec des sorties solides. Llama 4 s’inscrit dans la continuité de l’effort de Meta pour proposer des modèles extensibles et ouverts pour une large gamme d’usages.

## FAQs

### Y a-t-il une API pour Llama 4 ?

Meta n’a pas publié d’API officielle pour Llama 4. Toutefois, des fournisseurs tiers peuvent proposer un accès API à Llama 4.

