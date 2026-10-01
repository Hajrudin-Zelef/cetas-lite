---
id: collect-261001-ia-llm/ia-llm/chatgpt-gratuit-illimite-gpt-5-6-luna-remplace-o3-2026-1
title: "chatgpt-gratuit-illimite-gpt-5-6-luna-remplace-o3-2026"
domain: ia-llm
role: reference
task: reference
actors: ["Anthropic", "Hugging Face", "Moonshot", "OpenAI"]
dates: []
keywords: ["chatgpt", "luna", "attention", "benchmarks", "claude", "compute", "gpt-5.6", "kimi", "open source", "opus 5", "sol", "terra"]
source: docs/RAG/collect-261001-ia-llm/chatgpt-gratuit-illimite-gpt-5-6-luna-remplace-o3-2026.md
source_anchor: ""
source_lines: [1, 30]
sha256: c6bec11c911a49d8c7645d72a458efc28fcc77c8d6bc7977c1041e5b2fde825d
---

# chatgpt-gratuit-illimite-gpt-5-6-luna-remplace-o3-2026

Depuis le 6 août 2026, ChatGPT a changé de visage pour ses centaines de millions d’utilisateurs gratuits. OpenAI a basculé les comptes Free et Go vers GPT-5.6 Luna comme modèle par défaut, supprimé les quotas de conversation textuelle et, dans la foulée, retiré deux piliers historiques de son catalogue : o3 le 26 août, et GPT-4.5 dès le 26 juin. Une semaine plus tard, le 24 août, OpenAI a aussi déployé ChatGPT Ads dans 31 pays européens, dont la France, l’Allemagne, l’Espagne et l’Italie. Trois décisions prises en moins de trois semaines, qui redessinent l’économie du chatbot le plus utilisé au monde et posent une question simple : est-ce vraiment un cadeau, ou le prix à payer pour un service désormais financé par la publicité ?

## GPT-5.6 Luna : ce qui change concrètement pour les comptes gratuits

La famille GPT-5.6, composée de trois variantes — Sol, Terra et Luna — est passée de préversion à disponibilité générale le 9 juillet 2026, au terme d’un accès restreint de deux semaines ouvert à une vingtaine d’organisations partenaires. Les fiches techniques publiées sur Hugging Face à cette date créditent les trois modèles d’une fenêtre de contexte de 1,05 million de tokens et d’une sortie maximale de 128 000 tokens. Luna, la plus légère des trois, est taillée pour le débit conversationnel plutôt que pour le raisonnement long. C’est elle qui équipe désormais par défaut les comptes ChatGPT Free et ChatGPT Go, remplaçant GPT-5.5 Instant qui occupait ce rôle depuis plusieurs mois.

Le changement le plus visible pour l’utilisateur lambda, ce n’est pas le modèle en lui-même mais la fin des quotas. OpenAI a intégré cette bascule à sa mise à jour de sécurité d’août 2026, qui remplace GPT-5.5 Instant par la famille GPT-5.6 pour les comptes Free et Go dès le 6 août ; à partir de la semaine du 10 août 2026, la limite de messages textuels a été supprimée progressivement sur ces offres gratuites. Selon les relevés d’AI-Toolbox, les comptes Free et Go bénéficient toujours, en septembre 2026, de conversations textuelles illimitées avec GPT-5.6 Luna, sans compteur horaire visible. OpenAI a aussi introduit un bouton « Réfléchir » (Think) qui permet à Luna de prendre plus de temps sur les questions complexes, avec des garde-fous anti-abus pour éviter que cette fonction ne soit détournée pour contourner les limites de calcul.

Sur le papier, c’est une bonne nouvelle pour l’utilisateur français qui utilisait ChatGPT avec parcimonie pour ne pas tomber sur le mur de quota. Dans les faits, la gratuité illimitée d’un service qui coûte cher à faire tourner en inférence a toujours une contrepartie. Ici, elle s’appelle ChatGPT Ads.

## ChatGPT Ads débarque en France : 31 pays concernés

Le 24 août 2026, OpenAI a annoncé l’extension de ChatGPT Ads à 31 marchés européens, dont la France, l’Allemagne, l’Espagne, l’Italie, la Suède, la Norvège, le Danemark, les Pays-Bas et l’Autriche. Concrètement, les utilisateurs des offres gratuites verront désormais des publicités intégrées directement dans le flux de conversation, en complément des réponses générées par GPT-5.6 Luna.

Le calendrier n’est pas un hasard. OpenAI avait déjà annoncé le 15 août 2026 l’arrivée de ces nouveaux modèles dans l’Espace économique européen et en Suisse, avec une exigence explicite de consentement préalable pour toute personnalisation basée sur les données utilisateur — une contrainte directement liée au RGPD et aux discussions en cours avec les autorités de protection des données. En clair : la publicité arrive, mais son ciblage devra composer avec un cadre juridique européen plus strict qu’aux États-Unis.

Pour les entreprises françaises qui envisageaient ChatGPT comme un outil professionnel gratuit, ce virage publicitaire change la donne. Les comptes Plus et Pro restent, pour l’instant, épargnés par les publicités, ce qui pousse mécaniquement une partie des usages professionnels vers les offres payantes — une stratégie de monétisation classique du freemium, mais rarement assumée aussi frontalement par un acteur de l’IA générative.

## Pourquoi OpenAI retire o3 et GPT-4.5 maintenant

Le retrait de modèles n’est pas nouveau chez OpenAI, mais le calendrier de cet été 2026 est particulièrement chargé. GPT-4.5, lancé comme modèle de recherche premium en 2025, a été retiré de ChatGPT le 26 juin 2026 après une période de transition de 30 jours. OpenAI o3, modèle de raisonnement qui avait marqué les esprits sur les benchmarks scientifiques et mathématiques, a suivi le 26 août 2026, après une fenêtre de transition beaucoup plus longue de 90 jours — signe qu’il était encore largement utilisé par une base d’utilisateurs techniques au moment de l’annonce. La purge a aussi touché les alias API historiques : dès le 10 août 2026, OpenAI a redirigé les identifiants gpt-5.2-chat-latest et gpt-5.3-chat-latest vers gpt-5.6-sol, forçant au passage une bascule silencieuse pour tous les intégrateurs qui appelaient encore ces alias par défaut.

Cette purge du catalogue répond à une logique d’ingénierie autant qu’à une logique commerciale. Maintenir plusieurs générations de modèles en production coûte cher en infrastructure d’inférence, complique le support et dilue l’attention des équipes produit. En consolidant l’offre autour de la seule famille GPT-5.6 (Sol pour le raisonnement approfondi, Terra en position intermédiaire, Luna pour le conversationnel rapide), OpenAI simplifie sa pile technique au moment même où elle doit absorber un afflux massif d’utilisateurs gratuits libérés de leurs quotas.

Reste un point de friction pour les développeurs : les tarifs de l’API ont eux aussi bougé. Au lancement de juillet 2026, OpenAI facturait Sol 5 $ le million de tokens en entrée et 30 $ en sortie, Terra 2,50 $ et 15 $, et Luna 1 $ et 6 $. Ces grilles n’ont pas tenu longtemps : selon Capital & Compute, OpenAI a réduit ses tarifs dès le 30 juillet 2026, avec une baisse de 20 % sur Terra (ramené à 2 $/12 $ le million de tokens) et une chute spectaculaire de 80 % sur Luna (désormais 0,20 $/1,20 $), avant un second ajustement sur Sol le 21 août. Contrairement à l’interface grand public, l’API n’est pas concernée par la gratuité illimitée. Les intégrateurs qui construisaient sur o3 doivent désormais migrer leurs workflows vers GPT-5.6 Sol ou Terra, avec les ajustements de prompt que cela implique.

## Où se situe GPT-5.6 face à la concurrence

La famille GPT-5.6 n’a pas le monopole de l’attention en cette fin août 2026. Le classement quotidien Artificial Analysis, mis à jour au 26 août 2026 et cité par plusieurs observatoires français, place Claude Opus 5 d’Anthropic en tête sur l’intelligence générale, tandis que GPT-5.6 Luna (max) est identifié comme le meilleur rapport prix-performance de sa catégorie. Sur les benchmarks de codage, Tech-Insider rapporte que GPT-5.6 Sol atteint 88,8 % sur TerminalBench dès juillet 2026, pour un tarif d’entrée de 5 $ le million de tokens à l’époque — un score qui explique pourquoi Sol reste privilégié par les équipes de développement malgré la concurrence. Kimi K3 (max), modèle chinois open source, se distingue comme la meilleure option en accès libre.

