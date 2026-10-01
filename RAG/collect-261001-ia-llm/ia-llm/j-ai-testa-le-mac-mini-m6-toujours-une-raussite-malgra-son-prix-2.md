---
id: collect-261001-ia-llm/ia-llm/j-ai-testa-le-mac-mini-m6-toujours-une-raussite-malgra-son-prix-2
title: "j-ai-testa-le-mac-mini-m6-toujours-une-raussite-malgra-son-prix"
domain: ia-llm
role: reference
task: reference
actors: ["AMD", "Alibaba", "Apple", "Intel"]
dates: []
keywords: ["amd", "gpu", "humanoid", "panther lake"]
source: docs/RAG/collect-261001-ia-llm/j-ai-testa-le-mac-mini-m6-toujours-une-raussite-malgra-son-prix.md
source_anchor: ""
source_lines: [80, 110]
sha256: 0ce23179535d7467f514fa34b8531dba30c46e9abe8a1c452090beffbff1b78f
---

# j-ai-testa-le-mac-mini-m6-toujours-une-raussite-malgra-son-prix

Sur les performances, nous mesurons les tokens produits par seconde, une idÃ©e de la fluiditÃ© d’utilisation des modÃ¨les en local. Jusqu’Ã 9B paramÃ¨tres, les performances sont trÃ¨s correctes et dÃ©passent les mini-PC avec puce Panther Lake de nos autres tests avec des mesures Ã 69 tok/s en 3B, 33 tok/s en 7B et 25 tok/s en 9B.

Mais on remarque trÃ¨s vite un plafond de verre. Sur notre version testÃ©e avec 24 Go de mÃ©moire unifiÃ©e, seuls 17,8 Go sont adressables pour l’infÃ©rence, le reste Ã©tant rÃ©servÃ© au systÃ¨me. Dans nos tests, les modÃ¨les plus lourds comme Qwen3, qui rÃ©clament 18 Go, dÃ©bordent donc et provoquent des erreurs GPU. Il est possible de passer en force en s’affranchissant de cette limite VRAM pour faire passer ces plus gros modÃ¨les, au prix d’une certaine instabilitÃ© systÃ¨me. Pour la Â« *science* Â», nous avons pu mesurer les performances dans cette configuration : 

Qwen3 30B tourne Ã  61,6 tok/s et Qwen3-coder Ã  58 tok/s.

Globalement, ce Mac mini M6 est une machine parfaite pour faire tourner des modÃ¨les et assistants IA de petite Ã moyenne taille (jusqu’Ã 9B dans nos tests). De plus, il le fait avec une consommation minimale (9 W seulement au GPU), mais il nous faut encore ici rÃ©aliser d’autres tests pour mesurer prÃ©cisÃ©ment l’enveloppe thermique que gÃ©nÃ¨rent tous ces modÃ¨les.

La configuration 24 Go est donc un bon compromis entre efficacitÃ© IA et prix, mais elle vous bridera pour les modÃ¨les plus volumineux. Sur ce terrain, l’AMD Ryzen AI Max+ 395 s’en sort sans surprise bien mieux avec son pool de mÃ©moire unifiÃ©e pouvant aller jusqu’Ã 128 Go. Le Mac mini M6 peut Ãªtre configurÃ© jusqu’Ã 64 Go pour les utilisateurs les plus exigeants, mais les 128 Go (et mÃªme les 512 Go) sont rÃ©servÃ©s au Mac Studio.

## Refroidissement et bruit

C’Ã©tait l’une des grandes qualitÃ©s du modÃ¨le M4 : le Mac mini M6 se fait Ã peine entendre, y compris lors de compilations ou de rendus lourds. Les ventilateurs fonctionnent bien, mais ils sont quasiment inaudibles.

Niveau consommation globale, la machine ne dÃ©passera que rarement les 40 W en utilisation normale et pourra flirter avec les 46 W lors des charges lourdes. La puce M6 pourra grimper Ã 65 W, comme sur le M4, lorsque tous les cÅurs CPU et GPU sont sollicitÃ©s, ce qui est rarement le cas.

On sent ici que le passage au 2 nm permet de conserver une consommation mesurÃ©e malgrÃ© le gain de performances consÃ©quent par rapport Ã la gÃ©nÃ©ration prÃ©cÃ©dente.

## Prix et disponibilitÃ©

Le Mac mini M4 Ã©tait sÃ»rement la meilleure affaire de ces derniÃ¨res annÃ©es pour s’Ã©quiper avec une machine robuste et polyvalente, que ce soit pour la bureautique, la crÃ©ation ou, dÃ©sormais, l’IA. Mais la crise de la RAM et du stockage bat son plein et Apple a Ã©tÃ© obligÃ©, comme beaucoup d’autres constructeurs, d’augmenter ses prix.

Le Mac mini M6 dÃ©marre dÃ©sormais Ã 1 049 euros, soit 350 euros de plus que son prÃ©dÃ©cesseur. Avec 350 euros d’Ã©cart, l’entrÃ©e de gamme perd le statut de trÃ¨s bonne affaire qui faisait le succÃ¨s du M4. Les gains de la puce M6 se justifient surtout pour les crÃ©atifs et les usages IA lÃ©gers, moins pour la seule bureautique.

Ce contenu est bloquÃ© car vous n'avez pas acceptÃ© les cookies et autres traceurs. Ce contenu est fourni par Disqus.

Pour pouvoir le visualiser, vous devez accepter l'usage Ã©tant opÃ©rÃ© par Disqus avec vos donnÃ©es qui pourront Ãªtre utilisÃ©es pour les finalitÃ©s suivantes : vous permettre de visualiser et de partager des contenus avec des mÃ©dias sociaux, favoriser le dÃ©veloppement et l'amÃ©lioration des produits d'Humanoid et de ses partenaires, vous afficher des publicitÃ©s personnalisÃ©es par rapport Ã votre profil et activitÃ©, vous dÃ©finir un profil publicitaire personnalisÃ©, mesurer la performance des publicitÃ©s et du contenu de ce site et mesurer l'audience de ce site (en savoir plus)

En cliquant sur Â« Jâaccepte tout Â», vous consentez aux finalitÃ©s susmentionnÃ©es pour lâensemble des cookies et autres traceurs dÃ©posÃ©s par Humanoid et .

Vous gardez la possibilitÃ© de retirer votre consentement Ã tout moment. Pour plus dâinformations, nous vous invitons Ã prendre connaissance de notre Politique cookies.
