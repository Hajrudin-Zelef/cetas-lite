---
id: collect-261001-huawei/huawei/gpt6-sol-et-luna-des-performances-de-pointe-mais-a-moindre-cout-1
title: "gpt6-sol-et-luna-des-performances-de-pointe-mais-a-moindre-cout"
domain: huawei
role: reference
task: reference
actors: ["Anthropic", "OpenAI"]
dates: []
keywords: ["luna", "sol", "agents", "astra", "benchmark", "benchmarks", "claude", "fable 5", "opus 5", "valuation"]
source: docs/RAG/collect-261001-huawei/gpt6-sol-et-luna-des-performances-de-pointe-mais-a-moindre-cout.md
source_anchor: ""
source_lines: [1, 114]
sha256: f8cd296b2d4eecd6af7244f0612c27ad1a454fa40e7e7973a74a0ca233178712
---

# gpt6-sol-et-luna-des-performances-de-pointe-mais-a-moindre-cout

Cours

OpenAI vient d'élargir la famille GPT‑6 avec deux nouveaux modèles : GPT‑6 Sol et GPT‑6 Luna, tous deux positionnés sous le fleuron GPT‑6 Astra, dont nous parlions plus tôt ce mois-ci. Sol et Luna bénéficient tous deux d'importantes baisses de prix (50 % par rapport à leurs tarifs GPT‑5.6), avec en prime des gains en programmation et en automatisation des workflows.

\n
Hasard du calendrier ou non, Anthropic a publié Claude Opus 5.5 le même jour (aujourd'hui), avec un positionnement similaire : même approche d'entraînement que le modèle phare, optimisée pour le coût et la vitesse.

\n
Pour en savoir plus sur le modèle de référence, consultez notre tutoriel API GPT‑6 Astra et notre comparaison GPT‑6 Astra vs Claude Fable 5.1.

\n
## En bref

\n
- \n
- **GPT‑6 Sol et Luna sont les deux niveaux sous Astra** , entraînés avec la même recette et proposés à moitié prix de leurs prédécesseurs GPT‑5.6. \n
- **Le principal gain concerne le coût par tâche pour les agents et le code** , plus que la capacité brute. Presque tous les résultats communiqués par OpenAI sont ajustés au coût. \n
- **GPT‑6 Sol commet environ deux fois moins d'erreurs factuelles que GPT‑5.6 Sol** ; Luna progresse également, mais reste derrière Sol en fiabilité. \n
- **Privilégiez Sol pour les workflows agents et le code, et Luna pour les volumes importants de tâches répétitives** . Ne choisissez Astra que si la tâche justifie son prix. \n
- **Si vous utilisez déjà GPT‑5.6 Sol ou Luna, passez-y pour le prix** ; dans notre test pratique, le gain de capacité était modeste et s'est manifesté par une honnêteté face à une entrée défectueuse. \n

## Que sont GPT‑6 Sol et Luna ?

\n
Astra demeure le modèle le plus performant et le mieux aligné d'OpenAI, réservé aux travaux les plus exigeants. Sol et Luna visent à dériver l'essentiel de cette même approche d'entraînement, et une grande partie des performances obtenues en programmation, pilotage d'ordinateur, etc., vers des niveaux plus abordables et plus rapides.

\n
On peut l'assimiler au même rapport qu'entretient Opus 5.5 avec Fable 5.1 : pas une nouvelle frontière, mais des performances proches de l'état de l'art à une fraction du coût.

\n
## Fonctionnalités clés de GPT‑6 Sol et Luna

\n
Quelques points à retenir :

\n
### Vraiment moins chers, sur toute la ligne

\n
Le prix API de Sol est divisé par deux. Celui de Luna est même légèrement inférieur à la moitié. Vous trouverez plus de détails dans la section tarification.

\n
### Solides résultats sur les workflows métiers

\n
Sur AutomationBench, qui évalue des agents sur des tâches de vente, marketing, opérations, support, finance, etc., Sol au niveau d'effort le plus élevé dépasse Claude Opus 5, pour une fraction minime de son coût par tâche. Luna s'améliore nettement par rapport à son prédécesseur GPT‑5.6 tout en devenant sensiblement moins cher par tâche. OpenAI a inclus un graphique que j'ai reconstitué ici :

\n
J'ai remarqué que le graphique n'intégrait pas les chiffres d'Opus 5.5 publiés par Anthropic dans leur annonce Opus 5.5, j'ai donc ajouté une nouvelle série. On voit que les dernières versions de Sol et Luna restent les plus efficaces. Probablement, si OpenAI ne l'a pas inclus dans son annonce de lancement, c'est uniquement une question de timing : Opus 5.5 est sorti une heure avant.

\n
### Des gains en code qui suivent le coût, pas seulement la précision

\n
Sur FrontierCode, Sol est décrit comme égalant à peu près le meilleur score de Claude Fable 5.1, mais à un prix bien inférieur. Sur un autre benchmark de code (DeepSWE), Sol se rapproche du meilleur score de Fable 5 avec une forte remise côté coût, et Luna est présenté comme comparable à Opus 5 et Fable 5 à des niveaux d'effort intermédiaires.

\n
### Moins d'erreurs factuelles

\n
OpenAI indique que Sol divise approximativement par deux le taux d'erreurs de son prédécesseur sur une évaluation interne de factualité issue de conversations réelles où les utilisateurs avaient signalé des erreurs. Luna progresse aussi ; OpenAI affirme qu'à un niveau d'effort plus élevé, elle égale la factualité de GPT‑5.6 Sol pour une fraction de son coût. Ce point était plus difficile à vérifier.

\n
### Un style d'écriture moins verbeux

\n
Le style de communication d'Astra, plus concis et moins jargonneux, a été transposé sur Sol et Luna. OpenAI affine en continu le style conversationnel de ses modèles depuis que 4o a été critiqué pour son côté obséquieux.

\n
### Un caching plus malin et moins cher pour les agents

\n
Au-delà du prix par token, OpenAI met en avant des améliorations du cache de prompts destinées à aider les agents et les longues conversations à réutiliser le contexte plus efficacement, avec une remise d'environ 90 % sur les lectures d'entrées en cache. De nouveaux tableaux de bord et diagnostics permettent aux développeurs d'identifier où le cache fonctionne ou non.

\n
### Moins d'affirmations trompeuses sur leur propre travail

\n
Les deux modèles sont moins susceptibles que leurs homologues GPT‑5.6 de déformer ce qu'ils ont fait sur une tâche de codage.

\n
Les évaluations d'alignement d'OpenAI, menées au niveau d'effort maximal, créent délibérément des situations propices à la malhonnêteté, et Sol et Luna affichent des taux de tromperie inférieurs à ceux de GPT‑5.6 Sol et Luna sur les tests de désinformation en codage, recherche cassée, contournement de relecteur, détour des avertissements et interactions non autorisées.

\n
OpenAI précise qu'il s'agit de configurations adversariales, et non de taux d'échec en usage courant, et publie les résultats complets dans la carte système GPT‑6.

\n
## Comment GPT‑6 Sol et Luna s'en sortent-ils sur les benchmarks ?

\n
La communication d'OpenAI s'appuie fortement sur des comparaisons ajustées au coût. J'évoquais plus haut la performance de Sol sur AutomationBench. OpenAI ne se contente pas de dire que c'est \"mieux qu'Opus 5\" : c'est mieux *et* environ 11 fois moins cher par tâche.

Il faut aussi noter que, selon les propres chiffres d'OpenAI, Astra reste devant Sol et Luna pour les tâches d'utilisation d'ordinateur, et que certaines comparaisons avec les modèles Claude utilisent des niveaux d'effort différents (par ex. : Sol en \"xhigh\" contre Opus 5 en \"medium\"), ce qui renforce vraiment les arguments d'efficacité, mais rend les comparaisons de capacité brute plus délicates.

\n
Avec ces nuances en tête, voici ce qu'OpenAI rapporte, regroupé par type de travail représenté par chaque benchmark.

\n
### Workflows métiers et agents

\n
Le meilleur résultat de Sol est sur AutomationBench, le test de Zapier pour des agents travaillant de bout en bout sur 47 outils en vente, marketing, opérations, support, finance et RH. GPT‑6 Sol au niveau d'effort xhigh obtient 33,2 % pour 0,27 $ par tâche, devant Claude Opus 5 à l'effort maximal et même devant GPT‑6 Astra à faible effort, pour une fraction du coût par tâche d'Opus 5.

\n
La ligne Fable 5.1 est à lire avec précaution. OpenAI indique Claude Fable 5.1 avec fallback Opus 5 juste sous Sol, mais le fallback s'est déclenché sur environ 40 % des tâches, et son coût n'est pas inclus dans le chiffre annoncé.

\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n
| Modèle (et effort) | Score AutomationBench | Coût par tâche | 
|---|---|---|
| GPT‑6 Sol (xhigh) | 33,2 % | 0,27 $ | 
| GPT‑6 Astra (low) | 30,3 % | 3,9× GPT‑6 Sol | 
| Claude Fable 5.1 avec fallback Opus 5 (max) | 31,4 % | Plus de 8,9× GPT‑6 Sol (coût du fallback non reporté) | 
| Claude Opus 5 (max) | 26,9 % | 11,1× GPT‑6 Sol | 

