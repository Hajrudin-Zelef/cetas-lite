---
id: collect-260926-rattrapage/rattrapage/r-localllama-comments-1rq7jtm-qwen3535ba3b-uncensored-aggressive-gguf-release-bfb2a4ae-1
title: "r-localllama-comments-1rq7jtm-qwen3535ba3b-uncensored-aggressive-gguf-release-bfb2a4ae"
domain: rattrapage
role: reference
task: reference
actors: ["Alibaba", "Hugging Face", "Meta"]
dates: []
keywords: ["gguf", "llama", "qwen", "attention", "jailbreak", "llama.cpp", "moe", "multimodal"]
source: docs/RAG/lot-rattrapage/ai-llm/r-localllama-comments-1rq7jtm-qwen3535ba3b-uncensored-aggressive-gguf-release-bfb2a4ae.md
source_anchor: ""
source_lines: [1, 65]
sha256: c30ec6c0975ce79b7a4fe3979f36a3ec24558198f2b99de88a9ef17c4caf3bdb
---

# r-localllama-comments-1rq7jtm-qwen3535ba3b-uncensored-aggressive-gguf-release-bfb2a4ae

Merci pour ton avis !
Explique-nous pourquoi ce contenu n’est pas utile.
Qwen3.5-35B-A3B Non censuré (Aggressif) — Sortie GGUF
Celui que tout le monde demandait. Qwen3.5-35B-A3B Aggressive est sorti !
Aggressive = pas de refus ; il n'y a AUCUN changement ou altération de personnalité, c'est la version ORIGINALE de Qwen, complètement non censurée
https://huggingface.co/HauhauCS/Qwen3.5-35B-A3B-Uncensored-HauhauCS-Aggressive
0/465 refus. Complètement débloqué sans perte de capacité.
Celle-ci a pris quelques jours supplémentaires. J'ai travaillé dessus 12 à 16 heures par jour (c'est littéralement vrai) et je voulais m'assurer que la sortie soit de la plus haute qualité possible. D'après mes propres tests : 0 problèmes. Pas de bouclage, pas de dégradation, tout fonctionne comme prévu.
Ce qui est inclus :
- BF16, Q8_0, Q6_K, Q5_K_M, Q4_K_M, IQ4_XS, Q3_K_M, IQ3_M, IQ2_M
- mmproj pour le support visuel
- Tous les quants sont générés avec imatrix
Spécifications rapides :
- 35B au total / ~3B actifs (MoE — 256 experts, 8+1 actifs par token)
- 262K contexte
- Multimodal (texte + image + vidéo)
- Attention hybride : Gated DeltaNet + softmax (ratio 3:1)
Params d'échantillonnage que j'ai utilisés :
temp=1.0, top_k=20, repeat_penalty=1, presence_penalty=1.5, top_p=0.95, min_p=0
Mais vérifiez aussi les recommandations officielles de Qwen car elles ont des réglages différents pour le mode réflexion et le mode non-réflexion :)
Note : Utilisez le drapeau --jinja avec llama.cpp. LM Studio peut montrer "256x2.6B" dans les paramètres pour celui en BF16, c'est juste cosmétique, le modèle fonctionne à 100% correctement.
Versions précédentes de Qwen3.5 :
Tous mes modèles : HuggingFace HauhauCS
J'espère que tout le monde apprécie la sortie. Dites-moi comment ça fonctionne pour vous.
La communauté a été super utile pour Ollama, merci de lire les discussions dans les autres modèles sur Huggingface pour des conseils sur la façon de le faire fonctionner.
Section des commentaires
À quel point ce processus de désocclusion est-il difficile ? Ce n'est pas une désocclusion destructrice comme on pouvait le voir dans le passé. Est-ce que tu t'es penché sur cette architecture particulière ou as-tu un genre de "manuel" qui fonctionne de la même manière sur la plupart des modes ?
Je suis juste curieux, je n'ai aucune idée de comment c'est fait.
Bonjour,
Il y a des projets assez 'conviviaux' que vous pouvez utiliser :)
Pour ma part, j'utilise actuellement mon propre projet pour déverrouiller des modèles. Je pensais que cette variante en particulier (35b-a3b) serait la même que les autres dans la gamme qwen3.5, mais je me suis trompé et par exemple, cela m'a pris juste un peu moins d'une semaine pour bien le faire :)
Hauhau, tu déchires ! J'ai utilisé ton 9b, j'espère que tu sais à quel point tu es apprécié !
Merci pour les gentils mots !
Encore une fois, je TE SUPPLIE d'au moins évaluer KLD pour vraiment soutenir l'affirmation de "pas de perte de capacité".
J'apprécie que tu sois poli, donc je vais te répondre cette fois-ci. La divergence KL est une métrique incomplète. Tu peux avoir une divergence KL identique avec 1 modèle complètement incohérent, 1 complètement non censuré et 1 partiellement non censuré.
De plus, la raison pour laquelle je n'aime pas répondre à ce genre de choses, c'est que c'est un terrain glissant. Les gens demanderont les valeurs, puis la "preuve", ensuite la méthodologie, puis la source.
La divergence KL pour ce modèle (et encore une fois, ce n'est pas aussi pertinent que tu le penses) était exactement de 0,00053. Et la raison pour laquelle elle enregistre même cette valeur de divergence KL dans mon approche est à cause de la non-censure elle-même.
J'espère que ça aide.
trucs sympas.
Haha, ne fais rien que je ne ferais pas :D Mais merci de dire que c'est vraiment complètement non censuré.
La dégradation de la qualité est inévitable, surtout sur de longs contextes. La question est de savoir à quel point ce modèle s'éloigne du modèle standard ?
100%. Mais je voulais m'assurer que la dégradation de la qualité est celle du modèle original publié par Qwen et non pas par mes processus. C'est ce qui a pris plusieurs jours de travail supplémentaires au final. Teste-le et fais-nous savoir ! Quelques autres l'ont déjà fait aussi.
Merci ! J'attendais celui-ci, les autres poids fonctionnent plutôt bien, super boulot ici
Si tout se passe bien (j'espère de tout cœur que ça le sera, aucun problème dans mes tests), cette version devrait même supprimer la plupart (voir même TOUS) des avertissements aussi.
Eh bien, je dois demander, quand la 3.5 122b sera-t-elle disponible ?
Et je télécharge ça aujourd'hui pour le comparer à Heretic v2, j'ai hâte de l'essayer !
Merci beaucoup ! Je pense que j'ai déjà ta 9b et ça fonctionne super.
J'attends avec impatience ton suivi !
Oui, j'aimerais faire 122b, mais ça pourrait me prendre un peu plus de quelques jours :)
Je préfère sous-promettre et sur-livrer.
Quelle est la différence entre hauhau et huihui ?
Des gens différents, différentes techniques de désembarrassage.
Que signifie agressif ici ?
Salut, " Variante Aggressive
Désinhibition plus forte — le modèle est entièrement débloqué et niera pas les demandes. Peut occasionnellement ajouter de courtes clauses de non-responsabilité (intégrées dans l'entraînement de base du modèle, pas des refus), mais le contenu complet est toujours généré.
Pour une désinhibition plus conservatrice qui maintient certaines mesures de sécurité, vérifiez la variante Équilibrée lorsqu'elle sera disponible. "
Hmm, je n'arrivais pas à faire en sorte que le modèle reconnaisse et décrive le contenu de l'image dans la variante 35B-A3B. J'ai utilisé l'IQ4_XS quant. Il a essentiellement halluciné sur l'image (en se référant à d'autres contextes aléatoires de la conversation).
C'est avec LMStudio 0.4.7-b1 (Mac M4Pro) qui fait tourner OpenClaw en discutant sur Whatsapp et Telegram.
Quelqu'un pourrait-il essayer rapidement les capacités multimodales ?
J'avais ça fonctionnant correctement auparavant avec le modèle huihui 35B-A3B abliterated-i1 (quants mradermacher)
S'il te plaît, ne me fais pas peur comme ça !
merci beaucoup ! Qwen 3.5 est le premier Qwen que je n'ai pas pu désactiver en dehors de ses garde-fous depuis quelques heures ; ils semblent l'avoir entraîné sur beaucoup d'essais de jailbreak. ironiquement, générer des variations sur les jailbreaks est une des choses pour lesquelles les modèles non censurés sont utiles. j'ai vraiment apprécié celui-ci et ta version 27B. hâte de voir le 122B si tu y arrives un jour.
Salut, je peux confirmer que 122b est bien en cours maintenant et j'espère pouvoir sortir l'une des qualités auxquelles les gens se sont habitués avec mes sorties BIENTÔT
Merci pour la sortie ! Je vais comparer ça avec le 35B-A3B officiel sur mon M4 Pro 64 Go. Actuellement, j'obtiens 73 tok/s sur LM Studio (MLX) et 31 tok/s sur Ollama (GGUF Q4_K_M) avec le modèle vanille. J'ai hâte de voir si la version non censurée garde la même vitesse — je ferai un retour !
J'utilise des modèles locaux depuis l'époque OG de LLaMA. J'ai testé tellement de finetunes soi-disant "non censurés" au fil des ans, et aucun d'entre eux n'a vraiment été sans restrictions. J'ai généralement fini par revenir aux grandes API fermées parce que les alternatives locales avaient encore des gardes cachés.
