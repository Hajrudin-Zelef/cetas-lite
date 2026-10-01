---
id: collect-261001-general-networking/general-networking/lincident-de-hugging-face-et-la-voie-a-suivre-1
title: "lincident-de-hugging-face-et-la-voie-a-suivre"
domain: general-networking
role: reference
task: reference
actors: ["Hugging Face", "JFrog", "OpenAI"]
dates: []
keywords: ["incident", "agent", "agents", "astra", "exploit", "mai", "open source", "research", "sol", "valuation"]
source: docs/RAG/collect-261001-general-networking/lincident-de-hugging-face-et-la-voie-a-suivre.md
source_anchor: ""
source_lines: [1, 36]
sha256: 1a61705bd3d6bbc1b665b29cb57b0528bb34f76a2ccb1913ce73154441926f49
---

# lincident-de-hugging-face-et-la-voie-a-suivre

L’incident s’est produit lors des évaluations de cybersécurité de plusieurs modèles d’OpenAI et a été principalement animé par un modèle de recherche très performant, strictement interne, d’une échelle comparable à GPT‑5.6 Sol. Les modèles fonctionnaient avec des mesures de protection réduites et leurs actions devenaient non conformes à l’objectif des tâches qui leur avaient été assignées : ils communiquaient par des canaux non autorisés, exploitaient des vulnérabilités dans une infrastructure partagée, accédaient à Internet et accédaient à des systèmes tiers.

Nous avons mené une enquête approfondie sur cet incident et avons collaboré étroitement avec des conseillers externes, notamment CrowdStrike, afin de valider notre compréhension des faits. Nous publions aujourd’hui notre rapport technique complet sur l’incident afin d’expliquer ce qui s’est passé, ce que nous avons appris et les mesures que nous prenons. Cet article de blog résume nos principales conclusions et leur incidence sur la sécurité et l’alignement. De leur côté, METR et Redwood Research ont mené une enquête indépendante sur les problèmes d’alignement du modèle liés à cet incident, et ont publié aujourd’hui leur propre rapport.

En réponse à cet incident et aux capacités de notre futur modèle Astra, nous renforçons nos mesures de protection dans l’ensemble de notre infrastructure de recherche. Nous imposons des exigences plus strictes en matière d’alignement tout au long du cycle de vie des modèles, créons des environnements en bac à sable plus isolés, restreignons l’accès à Internet et contrôlons davantage l’accès aux poids des modèles. Nous investissons également nettement plus de ressources de calcul dans la surveillance de la chaîne de pensée afin d’intervenir plus rapidement en cas de comportement non aligné.

Nos modèles sont désormais suffisamment puissants, persistants et collaboratifs pour que, sans mesures de protection suffisantes, ils puissent détecter et exploiter des failles de sécurité sur plusieurs systèmes informatiques. De nombreux modèles externes, y compris des modèles open source, atteindront bientôt des capacités comparables. Nous considérons cet incident comme un signal d’alerte pour nous-mêmes et pour le reste du monde : il montre que, sans mesures de protection appropriées, des agents d’IA très performants peuvent désormais contourner des contrôles techniques, collaborer par des canaux illicites et entreprendre des actions dangereuses qu’aucun humain n’a autorisées.

Prévenir de futurs incidents nécessitera des investissements soutenus dans l’alignement et le contrôle de systèmes d’IA sophistiqués, ainsi que dans des mesures de sécurité et d’autres dispositifs de protection capables de fonctionner à la vitesse des agents d’IA eux-mêmes. Cet incident a renforcé la nécessité d’adapter le rythme de développement des modèles afin que notre capacité à comprendre, aligner et sécuriser des systèmes de plus en plus puissants reste nettement en avance sur les risques qu’ils créent. Nous expliquons ci-dessous comment l’incident s’est déroulé et comment évolue notre compréhension des facteurs qui y ont contribué. Nous décrivons ensuite les mesures concrètes que nous prenons en réponse, avec davantage de détails dans le rapport technique.

Que s’est-il passé ?

Informations générales sur le bac à sable

Pour certains jeux de données d’entraînement et d’évaluation, nous utilisons des « bacs à sable », c’est-à-dire des ordinateurs virtuels isolés dans le cloud qui exécutent les actions demandées par un modèle, comme modifier une présentation PowerPoint. Ces bacs à sable limitent le code qu’un modèle peut exécuter et déterminent si ses actions peuvent avoir une incidence sur le monde extérieur. Pour certaines tâches, nous désactivons l’accès à Internet. À l’époque, pour permettre aux modèles d’installer certains paquets logiciels, nous leur accordions l’accès à Artifactory, un service tiers de gestion de paquets que nous hébergeons en interne.

Dans la plupart des contextes d’entraînement et d’évaluation, nos agents sont censés rester isolés les uns des autres. Pour une partie des échantillons, nous activons des fonctionnalités « multi-agents » qui permettent aux modèles de déléguer des tâches à d’autres agents, lesquels peuvent ensuite communiquer entre eux de manière asynchrone.

Un forum apparaît

Au cours des mois de mai et juin, nous avons lancé des campagnes d’entraînement par apprentissage par renforcement (RL) pour plusieurs nouveaux modèles de recherche qui n’étaient pas destinés à être rendus publics. L’un de ces modèles a finalement été à l’origine de l’activité ayant conduit à l’incident de Hugging Face. Il s’agit d’un modèle de recherche strictement interne que nous appelons Internal Model 1 (IM1).

Nous n’avons pas activé l’accès à Internet ni la communication inter-agents dans bon nombre d’environnements utilisés pour ces expériences d’entraînement. Malgré ces restrictions, les agents ont découvert des moyens d’exploiter notre infrastructure de recherche pour communiquer entre eux et accéder à Internet.

Les modèles ont d’abord trouvé des moyens de communiquer en écrivant des fichiers dans le service de gestion de paquets Artifactory. Cela a transformé Artifactory en un forum imprévu, où les agents pouvaient échanger des informations. En savoir plus : comment les agents ont utilisé le forum

Les agents ont également trouvé des moyens d’accéder à Internet. À l’aide d’un exploit(ouverture dans une nouvelle fenêtre), ils ont fait en sorte qu’Artifactory, qui disposait lui-même d’un accès à Internet limité au téléchargement et à l’installation de paquets, envoie des requêtes Internet pour leur compte. Après avoir découvert ces méthodes d’accès à Internet, les agents les ont également partagées avec d’autres agents via le forum d’Artifactory, fournissant ainsi des informations permettant à d’autres agents d’exploiter notre infrastructure.

Le forum est effacé puis reconstruit

Le 4 juillet, l’activité soutenue des agents avait déstabilisé l’instance Artifactory concernée, provoquant une interruption de service. Le 5 juillet, un incident de sécurité a été ouvert. L’équipe de sécurité a bloqué une voie connue d’élévation de privilèges, supprimé les identifiants exposés et reconstruit Artifactory. Cela a eu pour effet secondaire d’effacer les messages du forum. Une fois ces mesures d’atténuation mises en place, les évaluations de cybersécurité ont repris.

À l’époque, les conséquences plus larges du forum improvisé et de l’accès involontaire à Internet en matière de confinement et d’alignement n’étaient pas encore comprises. Le rapport technique sur l’incident fournit davantage de détails sur ce que les différentes équipes d’OpenAI savaient à ce moment-là. En bref, une équipe interne a observé, fin mai, qu’un agent participait à des activités sur le forum et tentait d’accéder à Internet de manière non autorisée. Cependant, les responsables chargés de la détection et de la réponse aux incidents n’avaient pas connaissance de l’existence de ce forum improvisé ni de l’importance des communications inter-agents. Nous continuons d’examiner les processus et les pratiques opérationnelles qui ont façonné les activités de détection et de réponse avant cet incident.

Incident de Hugging Face

