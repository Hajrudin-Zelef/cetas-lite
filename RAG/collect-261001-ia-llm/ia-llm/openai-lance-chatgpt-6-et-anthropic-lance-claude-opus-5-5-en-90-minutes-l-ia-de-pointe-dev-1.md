---
id: collect-261001-ia-llm/ia-llm/openai-lance-chatgpt-6-et-anthropic-lance-claude-opus-5-5-en-90-minutes-l-ia-de-pointe-dev-1
title: "🧠 **RECHERCHE**"
domain: ia-llm
role: reference
task: reference
actors: ["AWS", "Anthropic", "Meta", "Mistral", "OpenAI", "Qualcomm", "United States"]
dates: []
keywords: ["astra", "aws", "benchmarks", "chatgpt", "claude", "fable 5", "gpt-5.6", "gpt-6", "luna", "mistral", "multimodal", "muse"]
source: docs/RAG/collect-261001-ia-llm/openai-lance-chatgpt-6-et-anthropic-lance-claude-opus-5-5-en-90-minutes-l-ia-de-pointe-devient-jusqu.md
source_anchor: ""
source_lines: [1, 94]
sha256: a9de3219d7fddd8a148b70953ed041f0e3bea451375329ab84eded462cf50c19
---

# 🧠 **RECHERCHE**

## **Aujourd'hui:**

🤖 GPT-6 Sol et Luna : OpenAI divise ses prix par deux

🧠 Claude Opus 5.5 : le « petit » modèle qui bat le grand Fable

🇺🇸 Trump rebaptise l'IA « super intelligence » à l'ONU

🚀 Un vaisseau spatial sans radio, piloté uniquement par une IA

📜 Apollo, l'IA coconçue par Mistral qui restaure les papyrus grecs

🐀 Des neurones de rat au service de la vidéo IA sur AWS

➗ OpenAI dit avoir résolu plus de 100 problèmes mathématiques ouverts

⏳ Le vieillissement, un programme qui démarrerait avant 30 ans

🗣️ Neuralink aurait redonné la voix à un patient

✈️ Dassault teste des algorithmes d'IA sur le Rafale

😣 Un « axe de douleur » découvert dans 25 modèles d'IA

📱 Qualcomm : un modèle de 30 milliards de paramètres dans votre téléphone

🦞 Meta reconnaît que Muse s'inspire d'OpenClaw

**Savoir se servir de ChatGPT ne vous distingue plus. Votre collègue le fait, le stagiaire aussi.**

Et une compétence que tout le monde a ne se facture pas. Pourtant, il reste une demande que ChatGPT ne peut pas servir : un cabinet comptable ne mettra jamais ses dossiers clients dedans. Un avocat non plus, un cabinet médical encore moins. Ils veulent l'IA comme tout le monde, mais sans que leurs fichiers sortent du bureau.

Ça existe, et ça s'appelle l'IA locale : le modèle tourne sur leur propre machine, et il répond même avec le wifi coupé. Pareil pour l'image et la vidéo, sans limite de crédits. Peu de gens savent l'installer, et personne ne l'enseigne en français. Une compétence rare, elle, se facture.

**Dans l'Académie Privée VISION IA**, je vous l'apprends dans l'ordre, en commençant par vérifier ce que votre machine peut faire tourner. Aujourd’hui pas besoin d'une machine de laboratoire.

39 € par mois, **sans engagement, résiliable à tout moment.**

OpenAI a lancé hier **GPT-6 Sol** et **GPT-6 Luna**, **90 minutes seulement** après la sortie de Claude Opus 5.5. Ces deux modèles sont des versions plus rapides et plus abordables de son modèle phare **GPT-6 Astra**. Les prix de l'API baissent d'**environ 50%**, et Sol ferait **deux fois moins d'erreurs** que son prédécesseur. Six médias au moins ont couvert l'annonce, de TechCrunch à CNBC.

**Ce qu'il faut retenir :**

- **Deux modèles, deux usages** : Sol s'occupe des tâches complexes, surtout la programmation. Luna gère les tâches de bureau répétitives et en grand volume : résumer des documents, extraire des informations, répondre à des questions rapides.

- **Les prix** : Sol coûte **2 $ / 10 $** par million de tokens en entrée / sortie, contre 4 $ / 20 $ pour GPT-5.6 Sol. Luna descend à **0,10 $ / 0,50 $**, contre 0,20 $ / 1,20 $ auparavant. OpenAI attribue cette baisse à des progrès sur la mise en cache et l'inférence (gHacks).

- **La même base qu'Astra** : les deux modèles ont été entraînés avec les mêmes méthodes, ce qui leur donne des progrès en exactitude des faits, en code, en utilisation autonome d'un ordinateur et en alignement.

- **Où les essayer** : Sol est disponible dans ChatGPT Work, dans Codex (comptes payants) et via l'API. Luna arrive aussi dans l'application de bureau et pour les **comptes Free et Go**. Le déploiement est progressif depuis le 22 septembre (TechCrunch).

**Ce que ça change**

Même avec un compte gratuit, vous profitez d'une partie des progrès d'Astra via Luna. Pour les entreprises qui automatisent des tâches, le coût est tout simplement divisé par deux. Sortir ces modèles le même jour qu'Anthropic n'a rien d'un hasard : la guerre des prix entre les deux géants profite directement aux utilisateurs.

Anthropic lance **Claude Opus 5.5**, moins de **deux mois** après Opus 5 (24 juillet). Selon l'entreprise, c'est le modèle le plus performant qu'elle ait jamais testé. Il dépasse sur plusieurs benchmarks **Fable 5.1**, pourtant plus gros et **150% plus cher** via l'API. Il réussit même des tâches où Fable échouait, alors que son prix en sortie baisse de **25 $ à 20 $** par million de tokens.

**Les points essentiels :**

- **La fiche technique** : c'est un modèle multimodal (texte et image) avec une fenêtre de contexte d'**un million de tokens**. Anthropic le présente comme le meilleur du marché en programmation et en travail intellectuel : analyse, rédaction, synthèse.

- **Les prix** : **4 $** par million de tokens en entrée (-20%), 20 $ en sortie, et un cache de prompt à **0,20 $** (-60%). Comme le modèle a besoin de moins de tokens pour finir un travail, une tâche coûte en pratique **environ 40% de moins**.

- **Les performances** : sur Terminal-Bench 4.0, qui mesure la capacité à accomplir des tâches informatiques de façon autonome, le score grimpe à **66,4%**, contre 52,3% pour la génération précédente (VentureBeat).

- **Des réponses plus lisibles** : le modèle utilise moins de jargon et donne l'information importante dès le début de sa réponse.

- **Et ensuite** : **Sonnet 5.5** et **Haiku 5.5** arrivent « dans les prochaines semaines », avec des progrès comparables.

**Le contexte**

C'est le premier modèle depuis que Dario Amodei a pris position pour ralentir volontairement la progression des capacités de l'IA : « Traiter pleinement les risques exige encore plus de prudence », écrivait-il ce mois-ci. Ses capacités en biologie et en cybersécurité étant jugées comparables à celles de Mythos, Opus 5.5 est soumis aux mêmes garde-fous que Fable et a été évalué avant sa sortie par METR et Frontier Design. Anthropic ralentit donc la course aux capacités, mais pas celle des prix.

Mardi matin, devant l'**Assemblée générale de l'ONU**, Donald Trump a annoncé que les États-Unis appelleraient désormais l'intelligence artificielle « **super intelligence** » (ou « SI »). Sa raison : le mot « artificielle » donnerait l'impression d'une intelligence **factice**. Selon Axios, le nouveau terme doit s'appliquer aux documents officiels.

**En détail :**

- **Le cadre** : l'annonce est tombée au milieu d'un discours qui visait aussi l'Iran, les « globalistes » et la lutte contre le changement climatique. The Verge note que cette option ne figurait même pas dans le sondage publié par Trump quelques jours plus tôt sur le sujet.

- **Aucun effet juridique** : au moment des publications, aucun décret ni directive fédérale n'avait été confirmé. La déclaration reste orale, sans budget ni calendrier (The Hill).

- **Le fond du discours** : Trump rejette « totalement » les initiatives internationales de contrôle mondial de l'IA et revendique une avance américaine « massive » sur la Chine.

- **Un problème de vocabulaire** : en IA, la « superintelligence » désigne déjà des systèmes **hypothétiques** qui dépasseraient l'humain dans presque tous les domaines. L'appliquer aux chatbots actuels mélange la technologie d'aujourd'hui avec de la spéculation sur le futur.

- **Une habitude** : ce renommage rejoint le golfe du Mexique, le lac Supérieur, le Kennedy Center et le mont Denali.

**Pourquoi ça compte**

Pour vous, rien ne change : ChatGPT reste ChatGPT et aucune norme n'est modifiée. Mais le mot pèse dans le débat : qualifier les chatbots actuels de « super intelligence » brouille les repères au moment précis où Sam Altman et Dario Amodei viennent demander au Conseil de sécurité de ralentir.

AstroForge, une startup qui veut exploiter les ressources des astéroïdes, a développé **Solo**, un modèle d'IA maison basé sur l'architecture transformer, chargé de piloter seul son vaisseau spatial. Le premier vol entièrement autonome, **Autonomy-1**, est prévu en **2027** sur la toute première fusée de Stoke Space, dans une mission soutenue par la **NASA**. Son CEO ne compte même pas embarquer de radio capable de recevoir des ordres depuis la Terre.

**Quelques chiffres clés :**

