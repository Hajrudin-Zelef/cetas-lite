---
id: collect-261001-general-networking/general-networking/quasar-438b-vs-laguna-m-1-vs-flux-3-ia-open-source-ue-1
title: "quasar-438b-vs-laguna-m-1-vs-flux-3-ia-open-source-ue"
domain: general-networking
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "DeepSeek", "Google", "Hugging Face", "Mistral", "Nvidia", "OpenAI", "Poolside", "Stability AI", "Z.ai"]
dates: []
keywords: ["agent", "agents", "apache", "benchmark", "benchmarks", "deepseek", "glm", "mistral", "multimodal", "nvidia", "open source", "open-weight"]
source: docs/RAG/collect-261001-general-networking/quasar-438b-vs-laguna-m-1-vs-flux-3-ia-open-source-ue.md
source_anchor: ""
source_lines: [1, 28]
sha256: 82a8d10a0be27037c9675a9c9bf6986c3ea40b084e94b5ad2b9b963f0177d408
---

# quasar-438b-vs-laguna-m-1-vs-flux-3-ia-open-source-ue

Le 2 septembre 2026, une entreprise espagnole quasiment inconnue du grand public a revendiqué le titre de “nouveau leader européen” de l’intelligence artificielle. Trois semaines plus tôt, une start-up franco-américaine avait signé un chèque de 1 milliard de dollars avec Nvidia. Et depuis juillet, un laboratoire allemand fondé par d’anciens ingénieurs de Stability AI fait tourner les têtes avec un modèle qui génère vidéo, audio et actions robotiques dans un seul réseau de neurones. Trois annonces, trois pays, une seule question : l’Europe a-t-elle enfin ses propres champions de l’IA open source, ou s’agit-il d’habillages marketing autour de technologies venues d’ailleurs ?

Ce comparatif met face à face **Quasar 438B** (Multiverse Computing, Espagne), **Poolside Laguna M.1 et Laguna S 2.1** (Poolside, France/États-Unis) et **FLUX 3** de Black Forest Labs (Allemagne). Trois projets qui n’ont presque rien en commun sur le papier : un modèle de raisonnement propriétaire, une paire de modèles de code en partie open-weight, et un moteur génératif multimodal. Mais tous les trois se disputent la même promesse commerciale : offrir à l’Europe une alternative crédible à OpenAI, Anthropic, Google et aux modèles chinois comme GLM ou Qwen, au moment où la Commission européenne pousse sa stratégie de souveraineté technologique et son “AI stack” fondé sur l’IA open source.

L’enjeu dépasse le simple classement de benchmarks. Pour une entreprise française ou européenne qui doit choisir où héberger ses données, quel modèle intégrer dans son produit ou comment répondre aux exigences du RGPD et de l’AI Act, la question de savoir qui est vraiment “made in Europe” a des conséquences juridiques et financières concrètes. Ce guide détaille les caractéristiques techniques, les prix, les licences, les cas d’usage réels et propose un guide de migration pour les équipes qui envisagent de basculer vers l’une de ces trois options.

## Quasar 438B de Multiverse Computing : le pari espagnol de la compression

Quasar 438B est né à San Sebastián, dans le nord de l’Espagne, sous la marque Multiverse Computing. L’entreprise ne s’était jusque-là fait connaître que pour sa technologie de compression de modèles, baptisée CompactifAI. Le 2 septembre 2026, elle a changé de statut en lançant son premier grand modèle : un système de raisonnement de 438 milliards de paramètres, pensé pour les agents d’entreprise et le développement logiciel, doté d’une fenêtre de contexte d’un million de tokens et capable de répondre en environ 15 secondes pour 500 tokens générés, raisonnement inclus.

Le communiqué officiel de Multiverse Computing a été relayé par le site Frandroid, qui a résumé la promesse en une formule reprise partout en France : Quasar 438B serait le “nouveau leader européen” de l’IA. Sur l’Artificial Analysis Intelligence Index v4.1.1, un indice composite qui agrège neuf tests couvrant les tâches d’agent, le code et le raisonnement scientifique, le modèle obtient un score de 43, contre 30 pour Mistral Medium 3.5, soit un écart de 13 points en faveur du modèle espagnol. Sur Terminal-Bench v2.1, il atteint 69,3, et sur les tests de raisonnement à long contexte, 75,0.

Mais la présentation officielle a rapidement soulevé des interrogations. Une mise à jour publiée par Multiverse Computing elle-même, le 5 septembre 2026, a précisé que Quasar 438B est en réalité un modèle compressé à partir de GLM-5.2, le modèle open-weight de la société chinoise Z.ai (Zhipu AI), sous licence MIT. Multiverse explique avoir utilisé CompactifAI pour rendre GLM-5.2 “plus petit et plus efficace”, avant de le spécialiser pour les agents et le code. Le site espagnol Xataka et le média spécialisé Signal Over Noise ont documenté cette filiation, notant que le changelog de CompactifAI décrit des capacités “identiques à GLM-5.2”.

Autre limite notable : Quasar 438B ne parle que deux langues, l’anglais et l’espagnol, ce qui restreint fortement son usage pour un marché francophone ou pour toute entreprise opérant dans plusieurs langues européennes. Le modèle n’est disponible qu’en API via CompactifAI, sans possibilité de téléchargement des poids ni de déploiement indépendant : il s’agit d’un modèle propriétaire, malgré son origine open-weight. Le prix annoncé est de 0,60 $ par million de tokens en entrée et 1,80 $ en sortie, un tarif nettement supérieur à celui de son propre modèle source GLM-5.2, disponible gratuitement en auto-hébergement sous licence MIT.

## Poolside Laguna M.1 et Laguna S 2.1 : le double jeu franco-américain

Poolside occupe une position ambiguë dans le paysage de l’IA européenne. Fondée en 2023 par Jason Warner, ancien directeur technique de GitHub, et Eiso Kant, entrepreneur du développement logiciel, l’entreprise a son siège social légal à San Francisco, mais dispose d’un bureau important à Paris, où une partie significative de la recherche est menée. Le classement “European LLMs in 2026” du site MRKT3.0 la classe néanmoins parmi les projets franco-américains, aux côtés de Mistral, en notant que la société a déplacé une partie de ses activités de San Francisco vers Paris dès 2023.

Sur le plan financier, Poolside a levé 1,626 milliard de dollars au total : 26 millions en amorçage, 100 millions en série A, 500 millions en série B en octobre 2024 (valorisation de 3 milliards de dollars), puis 1 milliard de dollars supplémentaires en série C en août 2026, apportés par Nvidia, qui valorise désormais l’entreprise à environ 13 milliards de dollars. Ce dernier tour s’accompagne d’un accord de “model factory” avec Nvidia, pour un montant rapporté à 6 milliards de dollars, et d’un recrutement de plus de 100 personnes chez Nvidia dédiées au projet depuis l’été 2026.

Sur le plan technique, Poolside a publié deux modèles distincts en juillet 2026. Laguna M.1 compte 225 milliards de paramètres, sous licence Apache 2.0, et vise les tâches de planification agentique : sur le benchmark WebBrain, qui teste la capacité d’un modèle à naviguer sur le web pour accomplir des tâches complexes, il atteint un taux de réussite de 73 %. Laguna S 2.1, plus petit avec 118 milliards de paramètres, est un modèle de code entièrement open-weight sous licence OpenMDW-1.1, actuellement en cours d’examen par l’Open Source Initiative. Sur SWE-Bench Multilingual, il obtient 78,5 %, devançant de justesse Qwen 3.7 Max (78,3 %), DeepSeek-V4-Pro-Max (76,2 %, pourtant un modèle de 1,6 billion de paramètres) et Tencent Hy3 (75,8 %). Sur SWE-Bench Pro, son score tombe à 59,4 %, selon la fiche modèle publiée par Poolside sur Hugging Face.

Côté tarifs, la situation est confuse : certains hébergeurs tiers facturent Laguna M.1 à 0,20 $ en entrée et 0,40 $ en sortie par million de tokens (avec un tarif réduit de 0,10 $ pour l’entrée mise en cache), tandis qu’un environnement de test développeur propose un accès gratuit à 0 $. Cette disparité illustre une réalité fréquente chez les modèles open-weight récents : le prix dépend entièrement de l’hébergeur choisi, et non d’un tarif officiel unique fixé par Poolside.

## Black Forest Labs et FLUX 3 : l’Allemagne mise sur le multimodal

