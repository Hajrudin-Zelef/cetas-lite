---
id: collect-260926-rattrapage/rattrapage/r-localllama-comments-1rcjzsk-is-opencode-the-best-free-coding-agent-currently-19078b8c-2
title: "r-localllama-comments-1rcjzsk-is-opencode-the-best-free-coding-agent-currently-19078b8c"
domain: rattrapage
role: reference
task: reference
actors: ["Alibaba", "Anthropic", "Google", "Mistral", "Moonshot", "Z.ai"]
dates: []
keywords: ["agent", "llama", "agents", "benchmarks", "claude", "copilot", "gemini", "glm", "gpu", "kimi", "leaderboard", "llama.cpp"]
source: docs/RAG/lot-rattrapage/ai-llm/r-localllama-comments-1rcjzsk-is-opencode-the-best-free-coding-agent-currently-19078b8c.md
source_anchor: ""
source_lines: [121, 165]
sha256: e74f1a14f2238ff6c55d77af6a51a841a19c4bbcccffefb7a45a50b1f3d0dd4e
---

# r-localllama-comments-1rcjzsk-is-opencode-the-best-free-coding-agent-currently-19078b8c

Section des commentaires
J'ai remarqué qu'Opencode demande un reprocessing, ce qui gaspille des ressources, ce qui n'est pas tenable lorsqu'on exécute des modèles localement.
Le code Qwen fonctionne bien, pas de reprocessing, il fait tout ce qu'Opencode fait.
on dirait qu'opencode pourrait résoudre l'une des raisons pour lesquelles cela se produithttps://github.com/anomalyco/opencode/issues/5224
> no prompt reprocessing
qu'est-ce que ça veut dire ?
sur llama.cpp, cela signifie que le contexte est invalidé et doit être retravaillé à partir de zéro.
Ce qui peut être acceptable sur une configuration GPU complète, mais c'est une expérience absolument atroce sur une configuration hybride GPU + CPU.
Même moi, avec 600+ pp tk/s, je trouve l'expérience horrible parce que quand vous dépassez 70k tokens (ce qui ne prend pas beaucoup dans des scénarios de codage normaux), il faut une minute au mieux pour chaque tâche à accomplir.
pas d'idée comment vérifier ça - tu sais si kilocode souffre du même problème ? On dirait que c'est construit sur opencode, mais en regardant les opérations du modèle, on dirait qu'il utilise cache cache.
J'apprécie OpenCode mais j'ai trouvé que Kilo CLI fonctionnait beaucoup mieux grâce au mode orchestrateur.
Il peut vraiment produire de bons résultats associé à Qwen Code Next sur Q5_K_L et 128K context ! Sur des scénarios IRL React + .NET aussi, pas de scénario "je construis cette atrocité codée", on parle de vrai code, d'entreprise, DÉSORDONNÉ.
OpenCode peut produire le même résultat (après tout c'est le même système sous le capot) mais le mode orchestrateur fait mieux fonctionner les choses et OpenCode ne l'a pas par défaut.
Oh-My-OpenCode peut-être ?
Bizarre, parce que kilo cli est un fork d'opencode et maintenant la dernière version de l'extension est juste un client
OpenCode demande étonnamment beaucoup de ressources CPU sur un portable même si ce n'est qu'une TUI. J'ai utilisé un portable avec Ryzen AI 350 et je vois les cœurs monter à 100 % et les ventilateurs s'accélérer chaque fois qu'OpenCode est en train de faire une tâche. Même au repos, ça pousse encore la consommation d'énergie au repos de l'ordinateur portable de 5W à presque 12W. Ce n'est pas un problème pour un desktop, mais pour un portable, chaque W compte. Et cela suggère une certaine inefficacité dans la façon dont ils gèrent la boucle d'événements.
Personnellement, j'utilise soit Qwen Code soit Gemini CLI (c'est la même chose), selon si je veux profiter de la fonctionnalité de mise à la terre de Google ou non. Les charges CPU restent stables à 5W, sans pics, et le fonctionnement de l'agent est très transparent. Pas à pas est directement dans le terminal.
Probablement qu'OpenCode a besoin d'une sorte d'optimisation du cache de prompts pour correspondre à Qwen Code / Gemini CLI / tout ce que GLM ou Mistral ou Kimi font. Quel autre agent aurait une qualité comme ça ?
Je sais que c'est bizarre de le relever, mais je suis passé à la version bureau (tauri) et cela a réduit de manière significative l'utilisation de la RAM (3 Go à 400 Mo).
Cela n'a pas de sens puisque la version bureau permet aux utilisateurs de créer des espaces de travail pour des agents parallèles (ce que je fais souvent)
Je dois aussi souligner que le bureau est seulement un client qui envoie des requêtes au serveur opencode.
je suis passé àhttps://buildwithpi.ai/ et j'adore, https://github.com/knoopx/pi
J'utilise opencode mais Crush est plutôt... https://github.com/charmbracelet/crush
Crush est nul Droid est bon mais pas ouvert Vibe préfère Mistral et manque de fonctionnalités Opencode est bon mais a parfois des problèmes d'appel d'outils mystérieux
En ce moment, j'utilise Oh-my-pi (omp). C'est super rapide mais a des problèmes d'appel d'outils avec GLM 4.7 à moins que tu crées une liste de tâches pour ton projet, alors ça va 😝
Oui, opencode a un quota gratuit pour les modèles open source avancés, c'est suffisant pour un usage quotidien ; mais si tu veux développer une application avec toutes les fonctionnalités, tu dois apporter ta clé API ou acheter un plan tarifaire avancé avec opencode zen.
OpenCode n'est même pas à la hauteur en ce qui concerne les références d'agents. Jette un œil à ça https://www.tbench.ai/leaderboard/terminal-bench/2.0
ForgeCode est open-source, écrit en Rust, et mène des benchmarks publics sur la précision.
Pour moi, c'est sûr. Mais il faut un petit ajustement au départ. Les agents intégrés n'ont pas les meilleures suggestions, mais tu peux les remplacer.
Le meilleur, c'est que ça fonctionne avec opus ou sonnet grâce à l'abonnement GitHub copilot.
La différence, c'est que copilot facture pour les suggestions et non les tokens. Tu peux l'ajuster davantage pour le rendre plus efficace au niveau des prompts.
J'utilise OpenCode et OpenChamber comme extension de barre latérale dans Antigravity depuis environ quatre mois.
C'est vraiment le meilleur cadre de codage agentique entièrement gratuit !
OpenCode fait tourner une variété de grands modèles frontier de plusieurs fournisseurs, principalement destinés à être une période de test pour ces modèles. En général, ils y restent environ un mois, jusqu'à ce qu'ils soient retirés d'OpenCode et remplacés par un autre, généralement plus récent. Je parie que : Ces fournisseurs ajoutent leurs modèles là-bas (et peut-être paient-ils OpenCode pour cela), afin que les utilisateurs puissent les tester gratuitement pendant un certain temps, puis espérer se tourner directement vers ce fournisseur (ou un abonnement OpenCode pour garder l'accès à ce modèle particulier).
Mais voici la bonne nouvelle. OpenCode propose une gamme de modèles toujours gratuits, comme BigPickle (un modèle frontier de milieu de gamme, qui est assez bon pour coder en bash, python, ainsi que pour des tâches de complexité moyenne), MiMo V2 (la version inférieure de MiMo, mais toujours assez capable), et actuellement il y a Nemotron 3 Ultra, qui surpasse Gemini 3.1 Pro High et atteint presque Opus en termes de complexité et de maintien du contexte. Cela fait déjà un moment qu'il est là et bien que je l'ai mis à l'épreuve littéralement pendant des jours dans une grande base de code avec des choses complexes, il ne s'est jamais plaint de quotas ou de limites. Jamais !
Cela a vraiment affecté ma stratégie et réduit mes coûts globaux de plus de 50 % !
Voici comment :
En plus de mon plan Gemini Pro, OpenCode offre une grande opportunité d'utiliser de grands modèles capables entièrement gratuitement. Cela m'a même amené à renoncer à mon abonnement Claude Code. Bien configuré, en utilisant des compétences personnalisées et des améliorations comme l'Antigravity-Kit et Get-Shit-Done ainsi qu'un bon prompt système, OpenCode dépasse même Claude Code par moments.
Récemment, j'ai ajouté (@chopretajas)/headroom (Github) au mélange, ce qui réduit le débit des tokens jusqu'à 65% !!
J'ai créé une compétence Multi-Model-Planning pour Antigravity, qui vérifie d'abord la disponibilité des modèles dans le chat interactif (principalement pour le cas où Claude manquerait de quotas dans Antigravity, ce qui arrive très vite, surtout avec Opus), puis crée un plan basé sur la tâche donnée, où il assigne les tâches requises au modèle qui y correspond le mieux. Ensuite, il crée des fichiers de tâches séparés pour chaque modèle, nommés avec un ordre d'exécution.
Exemple :
- T01-Logics_Gen_Gemini-Pro-High.md
- T02-Verification_and_Review_Nemotron.md
et ainsi de suite...
J'utilise principalement cette compétence avec un modèle comme Opus ou Nemotron 3 Ultra, car ce sont les meilleurs planificateurs et examinators/correcteurs de bogues, et ensuite je donne toutes les tâches créées à leurs modèles assignés, soit un par un soit simultanément (selon les conflits d'exécution).
