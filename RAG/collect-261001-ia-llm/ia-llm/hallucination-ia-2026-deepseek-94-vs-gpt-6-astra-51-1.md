---
id: collect-261001-ia-llm/ia-llm/hallucination-ia-2026-deepseek-94-vs-gpt-6-astra-51-1
title: "Comparaison basique du taux de réponses \"je ne sais pas\" entre deux fournisseurs"
domain: ia-llm
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "DeepSeek", "Google", "Mistral", "OpenAI", "xAI"]
dates: []
keywords: ["agent", "agents", "agi", "astra", "benchmark", "claude", "deepseek", "fable 5", "gemini", "gpt-5.6", "gpt-6", "grok"]
source: docs/RAG/collect-261001-ia-llm/hallucination-ia-2026-deepseek-94-vs-gpt-6-astra-51.md
source_anchor: ""
source_lines: [1, 26]
sha256: 57fa411f94f05e37c6fe683579ce7a6ba9af4b49066e8aa99ed962f749abcd10
---

# Comparaison basique du taux de réponses "je ne sais pas" entre deux fournisseurs

DeepSeek V4 Pro invente une réponse fausse dans 94 % des cas où il ne connaît pas la solution. GPT-6 Astra fait la même chose 51 % du temps. Claude Opus 5 se situe à 61 %. Ces trois chiffres, mesurés par le même laboratoire indépendant avec la même méthodologie, racontent une histoire que les fiches marketing des éditeurs passent sous silence : en 2026, aucun grand modèle généraliste ne dit vraiment “je ne sais pas” quand il le faudrait — même si un autre banc d’essai, HALC-Bench, montre qu’à l’inverse, des modèles plus petits et spécialisés comme finix_s1_32b d’Ant Group peuvent tomber sous 2 % d’erreur factuelle (1,8 % précisément, selon AI Multiple en août 2026). Le 3 septembre 2026, OpenAI a lancé GPT-6 Astra en promettant une chute spectaculaire de son taux d’hallucination interne, de 12,2 % à 4,2 %. Quatre jours plus tard, Fortune révélait que ce chiffre avait été discrètement modifié à deux reprises sur la page officielle. Ce comparatif s’appuie sur les données publiques d’Artificial Analysis, du Vectara Hallucination Leaderboard, du benchmark HALC-Bench relayé par AI Multiple, et des pages tarifaires officielles d’OpenAI, Anthropic, Google, DeepSeek, Mistral AI et xAI pour répondre à une question simple : quel modèle d’intelligence artificielle ment le moins en 2026, et à quel prix.

## Qu’est-ce qu’une hallucination d’IA et pourquoi ça change tout en 2026

Une hallucination désigne le moment où un modèle de langage génère une information fausse, inventée ou non vérifiable, mais la présente avec la même assurance qu’un fait exact. Le phénomène n’est pas nouveau, mais son coût a changé de nature depuis que les entreprises françaises et européennes branchent ces modèles directement sur des workflows de production : support client, rédaction de contrats, synthèse de dossiers médicaux ou automatisation d’agents autonomes. Une hallucination dans un chatbot de démonstration fait sourire. La même hallucination dans un agent qui répond à un client sur un remboursement, ou qui résume un dossier de conformité pour une PME soumise au RGPD, engage une responsabilité juridique.

La difficulté, en 2026, tient au fait que les modèles les plus récents répondent à presque toutes les questions, y compris celles pour lesquelles ils n’ont aucune base factuelle solide. Cette tendance à toujours produire une réponse plutôt qu’à admettre une incertitude explique pourquoi le taux d’hallucination ne baisse pas au même rythme que les scores de raisonnement ou de code. GPT-6 Astra sature le benchmark FrontierMath Tier 4 à 97,6 % et ARC-AGI-3 à 99,9 %, des scores quasi parfaits, alors que son taux d’hallucination sur un test de connaissances générales reste supérieur à 50 % selon Artificial Analysis. Un modèle peut donc exceller en mathématiques et en code tout en inventant des faits avec une confiance totale dès qu’il sort de son domaine d’expertise.

Pour un public professionnel, la question hallucination IA ne se limite plus à un débat académique. Elle conditionne le choix d’un fournisseur pour des cas d’usage réglementés, la conception des garde-fous applicatifs (RAG, vérification croisée, human-in-the-loop) et, de plus en plus, la conformité à l’AI Act européen, qui impose une documentation des risques pour les systèmes à haut risque. Ce comparatif détaille les chiffres disponibles pour GPT-6 Astra, Claude Opus 5, Claude Fable 5.1, Gemini 3.1 Pro, Gemini 3 Deep Think, DeepSeek V4, Mistral Large 3, Grok 4.6 et Qwen3.8-Max, avec leurs limites méthodologiques.

## Comment mesure-t-on l’hallucination : AA-Omniscience, Vectara HHEM et les chiffres maison

Il n’existe pas un test unique de l’hallucination IA, ce qui explique pourquoi deux articles peuvent citer des pourcentages très différents pour le même modèle. Trois familles de mesures coexistent en 2026, et elles ne mesurent pas la même chose.

La première famille regroupe les chiffres internes publiés par les éditeurs eux-mêmes sur leurs pages de lancement. OpenAI a par exemple annoncé, pour GPT-6 Astra, un taux d’hallucination interne de 4,2 %, contre 12,2 % pour son prédécesseur GPT-5.6 Sol. Ces chiffres sont calculés sur un jeu de test choisi par l’éditeur, avec une méthodologie rarement documentée en détail, ce qui limite leur comparabilité d’un fournisseur à l’autre.

La deuxième famille est le AA-Omniscience, un benchmark indépendant conçu par le cabinet d’évaluation Artificial Analysis pour mesurer à la fois les connaissances factuelles d’un modèle et sa tendance à halluciner lorsqu’il ne sait pas. Le taux d’hallucination AA-Omniscience correspond au pourcentage de réponses fausses données avec assurance parmi les cas où le modèle choisit de répondre plutôt que d’admettre son ignorance. C’est cette méthodologie qui produit les écarts les plus spectaculaires de ce comparatif, avec des taux allant de 51 % pour GPT-6 Astra à 96 % pour DeepSeek V4 Flash — une fourchette déjà large en mai 2026, quand GPT-5.5 plafonnait à 86 % d’hallucination sur ce même indicateur selon les données publiées par Suprmind, avant même l’arrivée d’Astra.

La troisième famille est le Vectara Hallucination Leaderboard, un projet open source qui utilise un modèle d’évaluation maison (HHEM) pour vérifier si un LLM invente des faits lorsqu’il résume un article de presse. C’est une tâche beaucoup plus étroite que AA-Omniscience : elle ne teste pas la culture générale du modèle, seulement sa fidélité au texte source. Or, au moment de la rédaction de cet article, le classement public de Vectara, mis à jour le 11 mai 2026, ne référence encore aucun des modèles phares sortis depuis l’été 2026 : ni GPT-6 Astra, ni Claude Opus 5, ni Gemini 3.1 Pro, ni DeepSeek V4, ni Mistral Large 3 n’y figurent par leur nom. Les modèles les plus récents du classement Vectara sont des versions antérieures ou des variantes allégées, comme gemini-2.5-flash-lite (3,3 % d’hallucination sur la tâche de résumé) ou mistral-large-2411 (4,5 %). Ce décalage de plusieurs mois entre la sortie d’un modèle et son apparition sur un classement indépendant réputé est un point que peu de comparatifs mentionnent, alors qu’il change la lecture des scores marketing publiés au lancement.

Conséquence pratique : un chiffre de 4 % annoncé par un éditeur et un chiffre de 51 % publié par un laboratoire indépendant sur le même modèle ne sont pas contradictoires, ils mesurent deux choses différentes. C’est précisément ce qui s’est produit avec GPT-6 Astra, détaillé plus loin dans cet article.

## Tableau comparatif 2026 : specs et taux d’hallucination des principaux LLM

Le tableau ci-dessous rassemble les données publiques disponibles au 12 septembre 2026. L’indice AA-Omniscience combine précision factuelle et calibration (capacité à s’abstenir plutôt que d’inventer). Un indice plus élevé est meilleur, tandis qu’un taux d’hallucination plus bas est préférable. Les cases marquées “non testé publiquement” signalent l’absence de donnée indépendante trouvée à la date de publication, information qui a elle-même une valeur pour un acheteur professionnel.

