---
id: collect-261001-general-networking/general-networking/a-propos-du-probleme-du-prix-du-millenaire-de-navierstokes-1
title: "a-propos-du-probleme-du-prix-du-millenaire-de-navierstokes"
domain: general-networking
role: reference
task: reference
actors: ["Anthropic", "OpenAI"]
dates: []
keywords: ["agent", "agents", "astra", "claude", "lean"]
source: docs/RAG/collect-261001-general-networking/a-propos-du-probleme-du-prix-du-millenaire-de-navierstokes.md
source_anchor: ""
source_lines: [1, 38]
sha256: 09607ed9d0d108c445f84bc9dba044f73a88df184920e1a58908d97b752100ce
---

# a-propos-du-probleme-du-prix-du-millenaire-de-navierstokes

Nous présentons une solution au problème d’existence et de régularité de Navier–Stokes, l’un des problèmes du prix du millénaire. Cette démonstration, produite par un système interne d’OpenAI, montre que la dynamique des équations de Navier–Stokes décrivant le mouvement des fluides peut engendrer une singularité en temps fini. Nous publions à la fois un exposé de la démonstration et une formalisation dans Lean.

Les problèmes du prix du millénaire(ouverture dans une nouvelle fenêtre) comptent parmi les questions les plus profondes à la pointe des mathématiques. La question de savoir si le mouvement régulier d’un fluide tridimensionnel peut cesser de l’être demeure sans réponse depuis environ 90 ans.

L’un des grands objectifs de nos travaux est de donner aux scientifiques les moyens de faire progresser la recherche et les technologies au bénéfice de toute l’humanité. Pour résoudre le problème de Navier–Stokes, nous avons utilisé un modèle interne nettement plus performant que GPT‑6 Astra. Nous estimons important d’informer le monde du rythme des progrès de l’IA et de ce qu’il peut attendre des prochains modèles.

Le problème

Les équations de Navier–Stokes s’appuient sur la deuxième loi du mouvement de Newton (« F=ma ») pour décrire le déplacement des fluides. Point important, elles considèrent un fluide comme un milieu continu au lieu de suivre chaque molécule. Ces équations servent à concevoir des avions, à prévoir la météo et à étudier la circulation sanguine.

Une question fondamentale restée ouverte pour ces équations dynamiques était de savoir si l’approximation continue du fluide pouvait cesser d’être valable. Plus précisément, les équations de Navier–Stokes pour un fluide tridimensionnel incompressible de densité constante peuvent-elles engendrer une « singularité », même si le mouvement est initialement régulier ? Ici, une singularité signifie que la dynamique conduit les vitesses du fluide à croître sans limite en un temps fini. Une singularité devrait apparaître malgré la viscosité, qui tend à régulariser le mouvement. Comme un fluide réel ne peut se déplacer à une vitesse infinie, cela marquerait une défaillance de la modélisation du fluide par les équations. Pour continuer à modéliser le système, il faudrait alors suivre individuellement le comportement de chaque particule.

Ces équations remontent aux travaux menés au XIXe siècle par Claude-Louis Navier et George Gabriel Stokes. En 1934, Jean Leray a démontré l’existence de solutions au sens généralisé, mais la question de leur régularité permanente est devenue un problème central non résolu. En 2000, le Clay Mathematics Institute a classé le problème d’existence et de régularité de Navier–Stokes parmi les sept problèmes du prix du millénaire.

Le résultat

Notre système a produit une démonstration analytique et une formalisation dans Lean établissant qu’un fluide initialement régulier et au repos peut engendrer une singularité en temps fini. Une force régulière est appliquée au fluide, dont l’énergie demeure finie pendant toute la dynamique, du repos jusqu’à la formation de la singularité. Cela résout le problème du prix du millénaire de Navier–Stokes en établissant l’énoncé « C » (ainsi que « D ») de la formulation officielle du prix du millénaire(ouverture dans une nouvelle fenêtre).

La solution est un vortex, un tourbillon de fluide en rotation, qui décrit une spirale vers l’intérieur et s’allonge toujours davantage, comme un spaghetti. Cette région centrale rétrécit tout en accélérant de telle sorte que son énergie demeure finie, conformément aux lois de la physique. La difficulté technique consiste à faire en sorte que les équations engendrent cette défaillance par le mouvement même du fluide, et non, par exemple, en introduisant artificiellement une force infinie. En termes plus mathématiques, les termes des équations de Navier–Stokes décrivant le mouvement — accélération, gradients de pression, transfert de quantité de mouvement et viscosité — doivent à la fois devenir grands et s’annuler de façon précise. Cet équilibre précis conserve une force externe régulière alors même que la vitesse du fluide croît sans limite.

Comment nous avons trouvé la démonstration

Depuis le 28 août, nous entraînons un nouveau modèle interne qui affiche des performances sans précédent dans nos évaluations, notamment en mathématiques. L’entraînement de ce modèle se poursuit et ses performances continuent de progresser.

Le mardi 1er septembre, nous avons entendu des rumeurs selon lesquelles deux problèmes du prix du millénaire avaient été résolus. Inspirés par ces rumeurs et le bond de performance de notre modèle interne, nous avons entrepris de l’évaluer sur tous les problèmes du prix du millénaire encore ouverts et sur quelques autres problèmes à fort impact.

Nous avons utilisé un système d’agents coordonnés propulsés par notre modèle interne. Les agents avaient accès à des outils leur permettant notamment de consulter une version en cache d’Internet et d’exécuter du code. Les agents étaient répartis en groupes et pouvaient communiquer au sein de leur groupe. La taille des groupes variait ; celui qui a résolu Navier–Stokes comptait environ 10 000 agents travaillant simultanément. Nous avons constamment maintenu les mêmes mesures de protection strictes que pour toutes nos évaluations de modèles de pointe, notamment la surveillance et l’isolement.

Pour chaque problème, nous avons soumis à différents groupes d’agents diverses variantes de son énoncé afin de toutes les couvrir. Pour le problème de Navier–Stokes, nous avons proposé à des groupes distincts les versions « A » et « B » (des formes particulières qui aboutiraient à une démonstration) ainsi que « C » et « D » (qui aboutiraient à une réfutation).

Outre les problèmes complets du prix du millénaire, nous avons demandé à notre système multi-agent d’essayer de résoudre une série de problèmes « plus faciles ». L’un d’eux portait sur une question similaire d’explosion pour la limite du problème de Navier–Stokes obtenue en supprimant le terme de viscosité. Cette question est connue comme le problème de régularité des équations d’Euler, et nos agents nous ont surpris en la résolvant. La variante précise qu’ils ont résolue était la version sans forçage, dans laquelle aucune force externe n’est appliquée au fluide. Près de 100 agents ont collaboré pendant environ 50 heures afin de produire notre réfutation de la régularité d’Euler1.

Après avoir vu la solution d’Euler, nous avons estimé que Navier–Stokes était le problème le plus prometteur. Nous avons donc décidé de consacrer nos ressources à Navier–Stokes. Pour cela, nous avons retiré des agents des autres problèmes du millénaire et leur avons fourni la résolution d’Euler. Lorsqu’une version plus entraînée de notre modèle interne est devenue disponible au cours du projet, nous avons fait passer nos agents à ce modèle.

Nous avons encouragé les différents groupes d’agents à explorer diverses approches. Après un certain temps, nous avons fait circuler les idées entre les groupes d’agents en utilisant Codex pour consolider les conclusions les plus utiles de chacun. Ces instructions de suivi s’appuyaient sur les résultats intermédiaires des agents eux-mêmes. Le groupe qui a trouvé la solution à Navier–Stokes a été guidé de cette manière.

Les agents sont parvenus à leur solution le samedi 5 septembre, environ 88 heures après le lancement des premiers agents. La formalisation et la vérification dans Lean ont nécessité 17 heures supplémentaires avec GPT‑6 Astra.

