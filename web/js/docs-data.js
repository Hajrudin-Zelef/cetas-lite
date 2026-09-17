// Centre d'aide Cetas Lite — contenu de référence public.
// Aucune information sensible d'infrastructure : uniquement l'usage produit.
export const CETAS_DOCS = {
  collections: [
    {
      id: "prise-en-main",
      title: "Prise en main",
      icon: "rocket",
      articles: [
        {
          id: "bienvenue",
          title: "Bienvenue dans Cetas Lite",
          sections: [
            { h: "Qu'est-ce que Cetas Lite ?", html: `<p>Cetas Lite est un assistant IA auto-hébergé : <strong>chat</strong> conversationnel, <strong>agents</strong> capables d'agir sur vos fichiers, et <strong>terminal intégré</strong>, réunis dans une seule application.</p><p>Contrairement aux services en ligne, l'application tourne sur votre propre serveur : vos conversations, vos fichiers et vos réglages restent chez vous.</p>` },
            { h: "Les trois espaces", html: `<ul><li><strong>Chat</strong> — discutez avec l'IA, posez des questions, faites relire ou résumer des documents.</li><li><strong>Agents</strong> — confiez une mission à un agent : il lit, écrit et exécute dans un espace isolé, avec votre approbation.</li><li><strong>Terminal</strong> — un vrai shell intégré, confiné à votre espace de travail.</li></ul>` },
            { h: "Confidentialité par conception", html: `<p>Vos conversations sont stockées localement. Seuls les messages que vous envoyez à l'IA transitent vers le fournisseur du modèle choisi — jamais vos fichiers ni vos réglages.</p><p>Vos clés API sont chiffrées avant d'être stockées sur votre serveur.</p>` },
          ],
        },
        {
          id: "premier-lancement",
          title: "Premier lancement et connexion",
          sections: [
            { h: "Créer votre compte", html: `<p>Au premier lancement, créez votre compte administrateur : il servira à sécuriser l'accès à l'application. Les lancements suivants se font par simple connexion.</p>` },
            { h: "Découvrir l'interface", html: `<p>La barre latérale gauche donne accès aux modules <strong>Agents</strong> et <strong>Terminal</strong>, à vos conversations, puis à votre menu utilisateur en bas. Le panneau de droite affiche les détails de la conversation active.</p><p>Commencez par connecter un fournisseur d'IA (article suivant), puis envoyez votre premier message dans le chat.</p>` },
          ],
        },
        {
          id: "connecter-fournisseur",
          title: "Connecter un fournisseur d'IA",
          sections: [
            { h: "Où renseigner vos clés", html: `<p>Ouvrez la <strong>Configuration</strong> (engrenage en bas de la barre latérale), onglet <strong>API et Modèles</strong>, puis renseignez la clé du fournisseur souhaité. Chaque clé est <strong>chiffrée</strong> avant d'être stockée sur votre serveur.</p>` },
            { h: "Fournisseurs cloud", html: `<p>Plusieurs fournisseurs cloud sont pris en charge (DeepSeek, OpenRouter, OpenCode, …). Vous pouvez en configurer plusieurs et changer de modèle à tout moment, même en cours de conversation.</p>` },
            { h: "Modèles locaux", html: `<p>Si vous disposez d'un serveur d'IA local (Ollama, LM Studio ou llama.cpp), Cetas Lite le <strong>détecte automatiquement</strong> : les modèles locaux apparaissent alors dans le sélecteur, sans clé API.</p>` },
          ],
        },
        {
          id: "choisir-modele",
          title: "Choisir un modèle",
          sections: [
            { h: "Le sélecteur de modèle", html: `<p>Le sélecteur, au-dessus de la zone de saisie, regroupe les modèles par <strong>familles</strong> : chaque famille correspond à un usage (discussion, code, génération locale).</p>` },
            { h: "Les familles", html: `<ul><li><strong>SamAgent Nano</strong> — léger et rapide, idéal pour les questions simples.</li><li><strong>SamAgent N4 / N8</strong> — polyvalents, en versions <em>flash</em> (rapidité) ou <em>standard</em> (équilibre).</li><li><strong>Code</strong> — optimisée pour le développement.</li><li><strong>SamGen</strong> — vos modèles <strong>locaux</strong> (Ollama, LM Studio, llama.cpp).</li></ul>` },
            { h: "Quel modèle choisir ?", html: `<p>Pour discuter, un modèle léger suffit. Pour du code ou des missions complexes, préférez une version <em>standard/elite</em>. Vous pouvez changer d'alias à chaque message sans casser le fil de la discussion. Le <strong>menu +</strong> liste les alias (Nano, N4, SamGen) en <strong>cascade</strong> : survolez un alias, puis une version. Au niveau de l'alias, <strong>Défaut</strong> = <em>routage par effort</em> (faible → Flash, moyen → Standard, max → Elite ; jamais Elite sans effort max explicite). Sous une version, <strong>Auto</strong> = le <em>fallback du mode</em> (tirage aléatoire sur son seul ensemble) et <strong>Models</strong> permet de choisir <em>manuellement</em> un modèle (Nano : <em>Free de OpenRouter</em>, <em>Zen Free</em>). Dans <strong>Configuration → Sélecteur de modèles</strong>, réglez chaque version sur <strong>1 modèle</strong> (fixe) ou <strong>Fallback</strong> (tirage aléatoire à chaque requête) — le bouton Fallback re-coche désormais tout l'ensemble.</p>` },
          ],
        },
        {
          id: "premier-message",
          title: "Envoyer votre premier message",
          sections: [
            { h: "Rédiger", html: `<p>Écrivez votre message dans la zone de saisie en bas. <strong>Entrée</strong> envoie, <strong>Maj + Entrée</strong> insère un saut de ligne. La zone grandit automatiquement avec votre texte.</p>` },
            { h: "Réponse en streaming", html: `<p>La réponse s'affiche au fur et à mesure de sa génération, avec un rendu riche (titres, code coloré, tableaux). Vous pouvez <strong>interrompre</strong> la génération à tout moment avec le bouton stop.</p>` },
            { h: "Affiner", html: `<p>Chaque réponse affiche le modèle utilisé, la durée et les jetons consommés. Vous pouvez <strong>copier</strong> une réponse ou demander à l'IA de la <strong>régénérer</strong>.</p>` },
          ],
        },
      ],
    },
    {
      id: "chat",
      title: "Chat",
      icon: "chat",
      articles: [
        {
          id: "conversations",
          title: "Gérer vos conversations",
          sections: [
            { h: "Nouvelle conversation", html: `<p>Le bouton <strong>Nouvelle conversation</strong> en haut de la barre latérale démarre une session vierge. Chaque session est isolée et conservée sur le serveur : cliquez sur une session de la liste pour la rouvrir, son historique complet est rejoué.</p>` },
            { h: "Favoris et sessions", html: `<p>Épinglez vos fils importants en <strong>favoris</strong> pour les retrouver instantanément. <strong>Supprimer</strong> efface définitivement la session, partout — elle ne réapparaîtra pas.</p>` },
            { h: "Exporter", html: `<p>Exportez la conversation active en <strong>Markdown</strong> ou <strong>JSON</strong> pour la conserver ou la partager hors de l'application.</p>` },
          ],
        },
        {
          id: "pieces-jointes",
          title: "Pièces jointes",
          sections: [
            { h: "Ajouter des fichiers", html: `<p>Cliquez sur <strong>+</strong> à gauche de la zone de saisie, ou <strong>glissez-déposez</strong> vos fichiers directement dans la fenêtre.</p>` },
            { h: "Formats pris en charge", html: `<ul><li><strong>Documents</strong> : PDF, texte brut, HTML — leur contenu est lu et peut être résumé, relu ou cité.</li><li><strong>Images</strong> : analysées visuellement par les modèles compatibles (vision).</li></ul><p>Convertissez les documents bureautiques (Word, Excel…) en PDF avant de les joindre.</p>` },
            { h: "Bon à savoir", html: `<p>Les pièces jointes sont conservées quelques jours puis nettoyées automatiquement. Retirer un fichier avant l'envoi l'annule ; après l'envoi, il fait partie du message.</p>` },
          ],
        },
        {
          id: "recherche-web",
          title: "Recherche web",
          sections: [
            { h: "Le globe", html: `<p>L'icône <strong>globe</strong> près de la zone de saisie contrôle la recherche web. Éteint, aucune recherche n'est effectuée ; allumé, l'IA peut chercher des informations à jour sur le web avant de répondre.</p>` },
            { h: "Les modes", html: `<ul><li><strong>Auto</strong> — l'IA utilise la recherche intégrée du fournisseur quand elle existe, sinon ses outils de recherche.</li><li><strong>Natif</strong> — force la recherche intégrée du fournisseur.</li><li><strong>Outils</strong> — force la recherche via les outils de l'agent.</li><li><strong>Désactivé</strong> — aucune recherche, réponse sur les seules connaissances du modèle.</li></ul><p>Le mode par défaut se règle dans Configuration → Fonctionnalités.</p>` },
            { h: "Recherche approfondie", html: `<p>Dans la vue Agents, le menu <strong>+</strong> propose une <strong>recherche approfondie</strong> : l'agent multiplie alors les requêtes et croise les sources pour les sujets complexes.</p>` },
          ],
        },
        {
          id: "thinking",
          title: "Raisonnement (Thinking) et effort",
          sections: [
            { h: "Afficher le raisonnement", html: `<p>Le bouton <strong>Thinking</strong> demande au modèle d'exposer son raisonnement avant de répondre. Dans le chat, il est désactivé par défaut ; dans la vue Agents, le raisonnement est <strong>obligatoire</strong> (pas d'interrupteur), avec un niveau d'effort réglable (Défaut / Faible / Moyen / Max).</p><p>Le raisonnement s'affiche dans un panneau latéral dédié, en temps réel pendant la génération : en haut, les informations de la requête (modèle, temps de génération, tokens d'entrée, tokens de sortie et débit, contexte utilisé, coût estimé), puis le raisonnement sous l'en-tête <strong>REASONING</strong>.</p><p>Quand le raisonnement est terminé, le panneau se masque automatiquement. Pour le reconsulter, cliquez le bouton <strong>Raisonnement</strong> affiché au-dessus de votre message.</p>` },
            { h: "Niveau d'effort", html: `<p>Le sélecteur d'<strong>effort</strong> (défaut, faible, moyen, max) règle la profondeur du raisonnement : un effort élevé donne des réponses plus fouillées, au prix d'un temps de génération plus long. En mode <strong>défaut</strong>, l'effort est résolu automatiquement selon la longueur du message (plus de 400 caractères → moyen, sinon faible) — l'effort réellement appliqué s'affiche sous chaque réponse.</p><p>L'effort pilote aussi le <strong>routage</strong> quand l'alias est sur <strong>Défaut</strong> (menu +) : faible → version Flash, moyen → Standard, max → Elite. Le tier le plus cher (Elite) n'est jamais choisi sans un effort max explicite.</p>` },
          ],
        },
        {
          id: "tokens-couts",
          title: "Jetons et coûts",
          sections: [
            { h: "Compteurs en direct", html: `<p>Sous la zone de saisie, une ligne affiche les <strong>jetons d'entrée et de sortie</strong> estimés pendant que vous tapez, ainsi que le <strong>coût estimé</strong> de la réponse à venir.</p><p>Chaque réponse terminée indique le modèle, la durée, les jetons réellement consommés et le débit.</p>` },
            { h: "Fenêtre de contexte", html: `<p>Un compteur montre la part de la <strong>fenêtre de contexte</strong> utilisée. Quand elle se remplit, l'application <strong>résume automatiquement</strong> les messages les plus anciens pour continuer sans perdre le fil (compactage).</p>` },
            { h: "Limiter la longueur", html: `<p>Le curseur <strong>Jetons max par réponse</strong> (menu +) borne la longueur des réponses — utile pour garder des réponses concises ou maîtriser les coûts.</p>` },
          ],
        },
        {
          id: "messages-avances",
          title: "Copier, régénérer, écouter",
          sections: [
            { h: "Actions sur les messages", html: `<ul><li><strong>Copier</strong> — récupère le texte brut d'une réponse.</li><li><strong>Régénérer</strong> — demande une nouvelle version de la dernière réponse.</li><li><strong>Lire</strong> — écoute la réponse en synthèse vocale (via votre navigateur).</li></ul>` },
            { h: "Saisie vocale", html: `<p>Le bouton <strong>micro</strong> dicte votre message à la voix (reconnaissance vocale du navigateur).</p>` },
            { h: "Mise en forme", html: `<p>Les réponses utilisent le Markdown : titres, listes, tableaux, blocs de code colorés et citations cliquables vers les sources.</p>` },
          ],
        },
      ],
    },
  ],
};
CETAS_DOCS.collections.push(
  {
    id: "agents",
    title: "Agents",
    icon: "agents",
    articles: [
      {
        id: "presentation-agents",
        title: "Présentation de la vue Agents",
        sections: [
          { h: "Une mission, un agent", html: `<p>La vue <strong>Agents</strong> (module en haut de la barre latérale) ouvre un espace plein écran dédié : vous décrivez une <strong>mission</strong>, et un agent l'exécute — il lit, écrit et lance des commandes dans un espace isolé.</p>` },
          { h: "Anatomie de la vue", html: `<ul><li><strong>Barre latérale</strong> — retour, métriques système, nouvelle conversation, recherche, projets, discussions.</li><li><strong>Zone centrale</strong> — fil de discussion avec l'agent.</li><li><strong>Compositeur</strong> — zone de saisie enrichie : projet, permissions, menu +, modèle, envoi.</li><li><strong>Panneau Raisonnement</strong> — le raisonnement de l'agent, en temps réel.</li></ul>` },
        ],
      },
      {
        id: "nouvelle-conversation-agent",
        title: "Lancer une mission",
        sections: [
          { h: "Décrire la mission", html: `<p>Cliquez <strong>Nouvelle conversation</strong>, choisissez éventuellement un <strong>projet</strong> (le dossier sur lequel l'agent travaillera), puis décrivez précisément le résultat attendu. Plus la mission est claire, meilleur est le résultat.</p>` },
          { h: "Choisir le modèle", html: `<p>Le sélecteur de modèle propose les familles compatibles avec l'agent. Pour du code, la famille <strong>Code</strong> est recommandée.</p>` },
          { h: "Modèles par défaut de l'agent", html: `<p>Dans <strong>Configuration → Sélecteur agent</strong>, réglez les modèles utilisés par la vue Agents, par alias (Nano / N4 / N8) et par niveau (Flash / Standard / Elite) : <strong>1 modèle</strong> fixe ou <strong>Fallback</strong> (tirage aléatoire à chaque requête, puis bascule séquentielle en cas d'échec) — comme le sélecteur du chat.</p>` },
          { h: "Suivi", html: `<p>Une <strong>pastille verte pulsante</strong> sur le bouton du module signale qu'au moins un agent travaille — même si vous avez fermé la vue. Chaque agent peut être suivi, stoppé ou supprimé depuis la liste des discussions.</p>` },
        ],
      },
      {
        id: "permissions",
        title: "Permissions de l'agent",
        sections: [
          { h: "Trois niveaux", html: `<ul><li><strong>Écriture (Write)</strong> — l'agent lit, écrit et exécute librement.</li><li><strong>Demander (Ask)</strong> — chaque action d'écriture ou d'exécution vous est soumise pour approbation.</li><li><strong>Lecture seule (Read only)</strong> — l'agent explore et analyse, sans rien modifier.</li></ul>` },
          { h: "Approuver ou refuser", html: `<p>En mode <strong>Demander</strong>, une carte d'approbation apparaît pour chaque action sensible : validez ou refusez en un clic. Sans réponse sous 10 minutes, la demande expire et l'action est annulée.</p>` },
          { h: "Changer en cours de route", html: `<p>Le sélecteur de permissions, dans le compositeur, peut être modifié à tout moment pendant la mission.</p>` },
        ],
      },
      {
        id: "plan-build",
        title: "Modes Plan et Build",
        sections: [
          { h: "Deux modes", html: `<ul><li><strong>Plan</strong> (barre jaune) — l'agent analyse et propose un plan d'action détaillé, sans rien exécuter.</li><li><strong>Build</strong> (barre bleue) — l'agent exécute le plan validé.</li></ul>` },
          { h: "Basculer", html: `<p>Appuyez sur <strong>Tab</strong> dans la zone de saisie, ou cliquez le badge <strong>PLAN / BUILD</strong>. Passer en Plan restaure automatiquement la lecture seule ; repasser en Build restaure vos permissions.</p>` },
          { h: "Valider le plan", html: `<p>En mode Plan, l'agent vous présente son plan et attend votre validation explicite avant toute exécution. C'est le workflow recommandé pour les missions délicates.</p>` },
        ],
      },
      {
        id: "menu-plus-agent",
        title: "Le menu + de l'agent",
        sections: [
          { h: "Fichiers et projet", html: `<p>Joignez des <strong>fichiers</strong> à la mission, ou rattachez un <strong>projet</strong> existant comme contexte de travail.</p>` },
          { h: "Compétences", html: `<p>Activez des <strong>compétences</strong> : des savoir-faire réutilisables injectés dans les instructions de l'agent pour la mission en cours.</p>` },
          { h: "Recherche", html: `<p>Choisissez une <strong>recherche standard</strong> ou <strong>approfondie</strong> : en mode approfondi, l'agent croise davantage de sources (le globe web s'allume automatiquement).</p>` },
          { h: "Plugins et worktree", html: `<p>Activez des <strong>plugins</strong> externes, ou isolez la mission dans un <strong>worktree</strong> : une copie de travail Git séparée, sans toucher à votre dossier principal.</p>` },
        ],
      },
      {
        id: "projets-agent",
        title: "Projets et dépôts",
        sections: [
          { h: "Sélecteur de projet", html: `<p>Le sélecteur de projet, à gauche du compositeur, indique sur quel dossier l'agent travaille. La barre au-dessus de la zone de saisie rappelle le projet actif et le mode en cours.</p>` },
          { h: "Projet actif", html: `<p>Le projet actif est conservé entre les sessions. L'agent ne voit que les chemins <strong>relatifs</strong> à ce projet : il ne peut pas sortir de son périmètre.</p>` },
        ],
      },
      {
        id: "worktree",
        title: "Worktree isolé",
        sections: [
          { h: "Travailler sans risque", html: `<p>Le <strong>worktree isolé</strong> exécute la mission dans une copie de travail Git séparée de votre dossier principal. Vos fichiers restent intacts pendant que l'agent expérimente.</p>` },
          { h: "Nettoyage", html: `<p>Le worktree est nettoyé automatiquement à la suppression de l'agent. Vous pouvez récupérer le travail produit avant de le supprimer.</p>` },
        ],
      },
      {
        id: "raisonnement-agent",
        title: "Panneau Raisonnement",
        sections: [
          { h: "Voir l'agent réfléchir", html: `<p>Le panneau <strong>Raisonnement</strong> s'ouvre automatiquement dès que l'agent commence à réfléchir, et se remplit en temps réel. Fermez-le d'un clic ; un bouton permet de le rouvrir à tout moment.</p>` },
          { h: "Niveau d'effort", html: `<p>Comme dans le chat, réglez l'<strong>effort</strong> de raisonnement (défaut, faible, moyen, max) selon la complexité de la mission.</p>` },
        ],
      },
      {
        id: "taches",
        title: "Panneau Tâches",
        sections: [
          { h: "Suivi en temps réel", html: `<p>Au-dessus du compositeur, le panneau <strong>Tâches</strong> liste les étapes de la mission avec leur statut : <strong>✓</strong> terminée, <strong>◐</strong> en cours, <strong>○</strong> à venir — et une progression <em>x / y</em>.</p>` },
          { h: "Reconstruction", html: `<p>La liste est reconstruite à l'ouverture d'une discussion, même après redémarrage : vous retrouvez toujours l'état d'avancement exact.</p>` },
        ],
      },
      {
        id: "discussions-agents",
        title: "Discussions d'agents",
        sections: [
          { h: "Liste et recherche", html: `<p>La barre latérale liste vos discussions d'agents avec leur statut (<em>en cours</em>, <em>terminé</em>, <em>stoppé</em>, <em>erreur</em>). La <strong>recherche</strong> filtre instantanément par titre ou contenu.</p>` },
          { h: "Favoris", html: `<p>Épinglez vos missions importantes en favoris pour les garder en tête de liste.</p>` },
          { h: "Supprimer", html: `<p>Supprimer un agent arrête son travail, nettoie son worktree éventuel et retire la discussion de la liste. Les agents stoppés sont restaurés comme tels après un redémarrage.</p>` },
        ],
      },
      {
        id: "brouillons-undo",
        title: "Brouillons et annuler / rétablir",
        sections: [
          { h: "Brouillon conservé", html: `<p>Votre texte en cours de rédaction est <strong>sauvegardé automatiquement</strong> par discussion : fermez la vue ou changez de discussion, votre brouillon vous attendra. Il est effacé à l'envoi.</p>` },
          { h: "Annuler / rétablir", html: `<p>Les boutons <strong>annuler / rétablir</strong> du compositeur permettent de revenir en arrière dans votre saisie, comme dans un éditeur de texte.</p>` },
        ],
      },
      {
        id: "outils-github",
        title: "Outils GitHub de l'agent",
        sections: [
          { h: "Neuf outils natifs", html: `<p>Connectez votre compte GitHub (Configuration → Connecteurs), et l'agent dispose de neuf outils :</p><ul><li><strong>Lecture</strong> (sans approbation) : lister vos dépôts, lister et lire les issues, lister les pull requests.</li><li><strong>Écriture</strong> (avec approbation) : créer un dépôt, créer une issue, commenter, créer et fusionner une pull request.</li></ul>` },
          { h: "En mode Plan", html: `<p>Les outils de lecture restent disponibles en mode Plan ; les outils d'écriture exigent le mode Build et votre approbation.</p>` },
          { h: "Sans connexion", html: `<p>Si aucun compte n'est connecté, l'agent vous indiquera comment le connecter au lieu d'échouer silencieusement.</p>` },
        ],
      },
      {
        id: "metriques",
        title: "Métriques système",
        sections: [
          { h: "Surveiller la machine", html: `<p>La section <strong>Métriques</strong> de la barre latérale affiche en temps réel le <strong>processeur</strong>, la <strong>mémoire</strong>, le <strong>disque</strong> et le <strong>réseau</strong> (débits descendant / montant) de la machine qui héberge l'application.</p>` },
          { h: "Modèle affiché", html: `<p>Le modèle actuellement sélectionné est rappelé en tête de la section.</p>` },
        ],
      },
    ],
  },
  {
    id: "projets",
    title: "Projets & distant",
    icon: "folder",
    articles: [
      {
        id: "creer-projet",
        title: "Créer un projet",
        sections: [
          { h: "Projet local", html: `<p>Dans la vue Agents, section <strong>Projets</strong>, créez un projet en choisissant un dossier de votre espace de travail. L'agent y aura accès pour ses missions.</p>` },
          { h: "Bonnes pratiques", html: `<p>Un projet = un périmètre (un dépôt de code, un dossier de documents…). L'agent ne manipule que des chemins relatifs à ce projet : impossible de sortir accidentellement du dossier.</p>` },
        ],
      },
      {
        id: "importer-fichiers",
        title: "Importer des fichiers",
        sections: [
          { h: "Upload", html: `<p>À la création d'un projet, importez des <strong>fichiers</strong> ou un <strong>dossier entier</strong>. Les fichiers volumineux sont acceptés (jusqu'à 200 Mo par fichier).</p>` },
          { h: "Formats", html: `<p>Tous les formats sont acceptés pour le stockage ; la lecture intégrée prend en charge le <strong>texte</strong>, le <strong>code</strong> et le <strong>PDF</strong>.</p>` },
        ],
      },
      {
        id: "explorer-lire",
        title: "Explorer et lire",
        sections: [
          { h: "Arborescence", html: `<p>Cliquez un projet pour déplier son <strong>arborescence</strong> : dossiers repliables, fichiers triés, navigation au clavier.</p>` },
          { h: "Lecteur intégré", html: `<p>Cliquez un fichier pour l'ouvrir dans le <strong>lecteur</strong> : coloration syntaxique pour le code, rendu paginé pour le PDF, aperçu pour le texte.</p>` },
        ],
      },
      {
        id: "sftp",
        title: "Projet distant (SFTP)",
        sections: [
          { h: "Connecter un serveur", html: `<p>Créez un projet de type <strong>distant</strong> en renseignant l'hôte, le port, l'utilisateur et le mode d'authentification. Le bouton <strong>Tester</strong> vérifie la connexion avant de créer le projet.</p>` },
          { h: "Sécurité de la clé d'hôte", html: `<p>À la première connexion, l'empreinte de la clé du serveur vous est présentée (<strong>TOFU</strong> : confiance à la première utilisation). Vérifiez-la avant d'accepter : elle protège contre l'usurpation de serveur.</p>` },
          { h: "Usage", html: `<p>Un projet distant s'utilise comme un projet local : arborescence, lecture de fichiers, et missions d'agent exécutées via SSH sur le serveur.</p>` },
        ],
      },
      {
        id: "projet-actif",
        title: "Projet actif",
        sections: [
          { h: "Sélection", html: `<p>Le <strong>projet actif</strong> est celui sur lequel travaillent vos nouvelles missions d'agent. Changez-le depuis le sélecteur du compositeur ou la section Projets.</p>` },
          { h: "Persistance", html: `<p>Le projet actif est mémorisé entre les sessions : vous retrouvez votre contexte de travail à chaque ouverture.</p>` },
        ],
      },
    ],
  },
);
CETAS_DOCS.collections.push(
  {
    id: "terminal",
    title: "Terminal intégré",
    icon: "terminal",
    articles: [
      {
        id: "terminal",
        title: "Utiliser le terminal",
        sections: [
          { h: "Ouvrir le terminal", html: `<p>Le module <strong>Terminal</strong>, en haut de la barre latérale, ouvre un tiroir avec un vrai shell — pas une simulation.</p>` },
          { h: "Sessions multiples", html: `<p>Jusqu'à <strong>6 sessions</strong> simultanées, chacune dans son onglet. Les programmes interactifs (éditeurs, outils en mode plein écran) fonctionnent, et le redimensionnement suit la taille du tiroir.</p>` },
          { h: "Espace confiné", html: `<p>Le terminal démarre dans votre <strong>espace de travail</strong> et y reste confiné : impossible de naviguer vers le reste du système depuis ce shell.</p>` },
          { h: "Fonctionne hors-ligne", html: `<p>Le terminal est entièrement embarqué dans l'application : il fonctionne même sans connexion internet.</p>` },
        ],
      },
    ],
  },
  {
    id: "extensions",
    title: "Compétences, plugins & MCP",
    icon: "puzzle",
    articles: [
      {
        id: "competences",
        title: "Compétences",
        sections: [
          { h: "Qu'est-ce qu'une compétence ?", html: `<p>Une <strong>compétence</strong> est un savoir-faire réutilisable (méthode, conventions, expertise métier) injecté dans les instructions de l'agent pour une mission.</p>` },
          { h: "Gérer", html: `<p>Créez et modifiez vos compétences dans <strong>Configuration → Compétences</strong>. Jusqu'à 64 compétences, chacune avec un nom, une description et un contenu détaillé.</p>` },
          { h: "Activer pour une mission", html: `<p>Dans la vue Agents, le menu <strong>+</strong> liste vos compétences : cochez celles utiles à la mission en cours. L'agent les reçoit en contexte.</p>` },
        ],
      },
      {
        id: "plugins",
        title: "Plugins externes",
        sections: [
          { h: "Principe", html: `<p>Les <strong>plugins</strong> ajoutent des outils sur mesure à l'agent : un script dans n'importe quel langage, ou un appel HTTP vers votre service.</p>` },
          { h: "Activer", html: `<p>Déclarez un plugin via son manifeste, rechargez depuis <strong>Configuration → Fonctionnalités</strong>, puis activez-le pour la mission depuis le menu <strong>+</strong> de l'agent.</p>` },
          { h: "Sécurité", html: `<p>Chaque appel de plugin exige votre <strong>approbation</strong>. Les plugins s'exécutent isolés du système, avec leurs ressources strictement limitées.</p>` },
        ],
      },
      {
        id: "mcp",
        title: "Serveurs MCP",
        sections: [
          { h: "Le protocole MCP", html: `<p><strong>MCP</strong> (Model Context Protocol) connecte l'agent à des outils externes standards : bases de données, API, services — via des serveurs déclarés en JSON.</p>` },
          { h: "Déclarer un serveur", html: `<p>Ajoutez vos serveurs (mode <strong>stdio</strong> ou <strong>HTTP</strong>) dans la configuration MCP. Leurs outils apparaissent alors pour l'agent sous la forme <em>serveur_outil</em>.</p>` },
          { h: "Isolation", html: `<p>Les serveurs MCP s'exécutent dans des processus isolés, tués proprement à l'arrêt, avec leurs journaux d'erreur bornés.</p>` },
        ],
      },
      {
        id: "outils-perso",
        title: "Outils HTTP personnalisés",
        sections: [
          { h: "Vos propres outils", html: `<p>Déclarez des <strong>outils HTTP</strong> personnels (URL + paramètres) : l'agent pourra les appeler comme ses outils natifs, avec approbation.</p>` },
          { h: "Cas d'usage", html: `<p>Idéal pour brancher vos API internes : recherche documentaire, tickets, déploiements — sans écrire de plugin.</p>` },
        ],
      },
    ],
  },
  {
    id: "configuration",
    title: "Configuration",
    icon: "settings",
    articles: [
      {
        id: "api-modeles",
        title: "API et Modèles",
        sections: [
          { h: "Clés API", html: `<p>Renseignez ici les clés de vos fournisseurs. Chaque clé est <strong>chiffrée</strong> avant stockage ; elle n'est jamais affichée en clair après enregistrement.</p>` },
          { h: "Familles de modèles", html: `<p>La section <strong>Familles</strong> liste les modèles disponibles tels que configurés côté serveur, avec leurs usages recommandés.</p>` },
        ],
      },
      {
        id: "fonctionnalites",
        title: "Fonctionnalités",
        sections: [
          { h: "Valeurs par défaut", html: `<p>Cet onglet règle les comportements par défaut des <strong>nouvelles</strong> conversations :</p><ul><li><strong>Recherche web</strong> — activée ou non par défaut.</li><li><strong>Mode de recherche</strong> — auto, natif, outils ou désactivé.</li><li><strong>Thinking</strong> — raisonnement affiché par défaut dans le chat.</li><li><strong>Effort</strong> — profondeur de raisonnement par défaut.</li></ul>` },
          { h: "Extensions", html: `<p>Activez les <strong>serveurs MCP</strong> par défaut et <strong>rechargez les plugins</strong> après modification de leurs manifestes.</p>` },
          { h: "MAREX.md", html: `<p>Éditez ici votre fichier <strong>MAREX.md</strong> : des instructions persistantes lues au démarrage de chaque session de chat ou d'agent (vos conventions, votre contexte, vos préférences).</p>` },
        ],
      },
      {
        id: "connecteurs",
        title: "Connecteurs",
        sections: [
          { h: "GitHub", html: `<p>Connectez votre compte GitHub avec un <strong>token personnel</strong> (portée <em>repo</em>). Le token est <strong>validé</strong> à la connexion et stocké chiffré.</p><p>Une fois connecté, l'agent dispose de ses outils GitHub (dépôts, issues, pull requests) et vos opérations Git utilisent ce compte.</p>` },
          { h: "Déconnexion", html: `<p>Vous pouvez déconnecter le compte à tout moment depuis cette même carte : l'agent perd alors l'accès à GitHub.</p>` },
        ],
      },
      {
        id: "remote-config",
        title: "Remote",
        sections: [
          { h: "Serveurs distants", html: `<p>Cet onglet centralise vos <strong>serveurs SFTP</strong> : ajoutez-les ici une fois, puis rattachez-les à autant de projets distants que nécessaire.</p><p>Chaque serveur peut être <strong>testé</strong> avant usage, et sa clé d'hôte vérifiée (TOFU).</p>` },
          { h: "Réutilisation", html: `<p>Un serveur défini ici peut alimenter <strong>plusieurs projets</strong> : modifiez ses paramètres une seule fois, tous les projets rattachés suivent.</p>` },
        ],
      },
      {
        id: "competences-config",
        title: "Compétences",
        sections: [
          { h: "Bibliothèque", html: `<p>Retrouvez ici l'ensemble de vos <strong>compétences</strong> : créez, modifiez, activez ou supprimez. Les modifications sont prises en compte immédiatement pour les nouvelles missions.</p>` },
          { h: "Bien les nommer", html: `<p>Donnez à chaque compétence un <strong>nom explicite</strong> et une description précise : c'est ce qui vous permettra de retrouver la bonne au moment de composer une mission.</p>` },
        ],
      },
      {
        id: "coffre",
        title: "Coffre",
        sections: [
          { h: "Principe", html: `<p>Le <strong>Coffre</strong> stocke vos secrets (clés API, tokens…) <strong>chiffrés</strong> dans un fichier unique (<em>vault.enc</em>), protégé par un <strong>mot de passe maître</strong> (12 caractères minimum).</p><p>Le chiffrement utilise des standards éprouvés (dérivation de clé à coût mémoire élevé, chiffrement authentifié) avec écriture atomique sur disque. Le mot de passe n'est conservé qu'en mémoire, jamais sur disque.</p>` },
          { h: "Verrouillage", html: `<p>Le coffre se <strong>verrouille</strong> manuellement (bouton Verrouiller) ou <strong>automatiquement après 15 minutes d'inactivité</strong>. Un mot de passe incorrect ne déconnecte pas votre session : seul le coffre reste fermé.</p>` },
          { h: "Mot de passe maître", html: `<p>Changez-le depuis le panneau dépliable <strong>« Changer le mot de passe maître »</strong> : le coffre est rechiffré avec un nouveau sel. En cas d'oubli, aucun mécanisme de récupération n'existe — conservez-le en lieu sûr.</p>` },
        ],
      },
      {
        id: "apparence",
        title: "Apparence",
        sections: [
          { h: "Thèmes", html: `<ul><li><strong>Océan</strong> — le thème signature, fond bleu profond animé.</li><li><strong>Sombre</strong> — gris profonds, reposant pour les yeux.</li><li><strong>Clair</strong> — fond clair pour les environnements lumineux.</li></ul><p>Le choix s'applique instantanément à toute l'interface, vue Agents comprise.</p>` },
          { h: "Thème rapide", html: `<p>Dans le chat, le menu utilisateur propose aussi un <strong>basculeur de thème</strong> direct, sans passer par la Configuration.</p>` },
        ],
      },
    ],
  },
);
CETAS_DOCS.collections.push(
  {
    id: "compte",
    title: "Compte & données",
    icon: "user",
    articles: [
      {
        id: "menu-avatar",
        title: "Menu du compte",
        sections: [
          { h: "Votre avatar", html: `<p>En bas de la barre latérale, votre <strong>avatar</strong> ouvre le menu du compte :</p><ul><li><strong>Paramètre</strong> — ouvre la Configuration.</li><li><strong>Obtenir de l'aide</strong> — ouvre ce centre d'aide.</li><li><strong>En savoir plus</strong> — ouvre ce centre d'aide.</li><li><strong>FAQ</strong> — ouvre la foire aux questions.</li><li><strong>Déconnexion</strong> — ferme votre session.</li></ul>` },
          { h: "Rôles et prompts", html: `<p>Le menu utilisateur du chat donne aussi accès à vos <strong>rôles</strong> et <strong>prompts enregistrés</strong> : des personas et des modèles de messages réutilisables en un clic.</p>` },
        ],
      },
      {
        id: "confidentialite",
        title: "Confidentialité des données",
        sections: [
          { h: "Ce qui reste chez vous", html: `<ul><li>Vos <strong>conversations</strong> et leur historique.</li><li>Vos <strong>fichiers</strong> et projets.</li><li>Vos <strong>clés API</strong> (chiffrées) et réglages.</li><li>Votre <strong>mémoire</strong> (pages personnelles et index).</li></ul>` },
          { h: "Ce qui transite vers le fournisseur", html: `<p>Uniquement le <strong>contenu nécessaire à la réponse</strong> : vos messages, les pièces jointes que vous envoyez explicitement, et le contexte de la conversation en cours. Rien d'autre.</p>` },
          { h: "Chiffrement", html: `<p>Les secrets (clés API, tokens) sont chiffrés en <strong>AES-256</strong> avant stockage. Les mots de passe ne sont jamais conservés en clair.</p>` },
        ],
      },
      {
        id: "sauvegardes",
        title: "Sauvegardes",
        sections: [
          { h: "Exporter vos données", html: `<p>Depuis le menu utilisateur, <strong>Sauvegardes</strong> permet d'exporter l'ensemble de vos données (conversations, réglages, mémoire) dans une archive.</p>` },
          { h: "Restaurer", html: `<p>Importez une archive pour retrouver votre environnement complet — utile lors d'un changement de machine ou après réinstallation.</p>` },
        ],
      },
      {
        id: "memoire",
        title: "Mémoire de l'assistant",
        sections: [
          { h: "Pages personnelles", html: `<p>L'assistant tient une <strong>mémoire</strong> sous forme de pages : ce qu'il apprend de vous et de vos préférences au fil des conversations.</p>` },
          { h: "Recherche", html: `<p>Avant de répondre, il peut <strong>rechercher</strong> dans sa mémoire les informations pertinentes — fini de répéter votre contexte à chaque fois.</p>` },
          { h: "Contrôle", html: `<p>Votre mémoire vous appartient : consultez et corrigez ce qui y est enregistré depuis l'interface.</p>` },
        ],
      },
    ],
  },
  {
    id: "reference",
    title: "Référence",
    icon: "book",
    articles: [
      {
        id: "raccourcis",
        title: "Raccourcis clavier",
        sections: [
          { h: "Partout", html: `<ul><li><strong>Entrée</strong> — envoyer le message.</li><li><strong>Maj + Entrée</strong> — saut de ligne.</li><li><strong>Échap</strong> — fermer le menu ou la fenêtre ouverte.</li></ul>` },
          { h: "Vue Agents", html: `<ul><li><strong>Tab</strong> (dans la zone de saisie) — basculer entre les modes <strong>Plan</strong> et <strong>Build</strong>.</li></ul>` },
          { h: "Conseil", html: `<p>Les boutons affichent leur action au survol : passez la souris pour découvrir les raccourcis disponibles.</p>` },
        ],
      },
      {
        id: "glossaire",
        title: "Glossaire",
        sections: [
          { h: "Termes courants", html: `<ul><li><strong>Agent</strong> — une IA qui agit : elle utilise des outils (fichiers, commandes, web) pour accomplir une mission.</li><li><strong>Alias / famille</strong> — un nom convivial regroupant des modèles par usage (ex. <em>Code</em>, <em>SamAgent</em>).</li><li><strong>Compactage</strong> — résumé automatique des messages anciens quand la fenêtre de contexte se remplit.</li><li><strong>Jeton (token)</strong> — unité de facturation des modèles ; un mot ≈ un à deux jetons.</li><li><strong>MCP</strong> — protocole standard pour brancher des outils externes à l'agent.</li><li><strong>Permissions</strong> — ce que l'agent a le droit de faire seul : écrire, demander, ou lecture seule.</li><li><strong>Plan / Build</strong> — deux phases d'une mission : concevoir le plan, puis l'exécuter.</li><li><strong>Streaming</strong> — affichage de la réponse au fur et à mesure de sa génération.</li><li><strong>TOFU</strong> — <em>Trust On First Use</em> : la clé d'un serveur distant est approuvée à la première connexion, après votre vérification.</li><li><strong>Worktree</strong> — copie de travail Git isolée où l'agent expérimente sans toucher à vos fichiers.</li></ul>` },
          { h: "Abréviations", html: `<ul><li><strong>PTY</strong> — pseudo-terminal : ce qui fait du terminal intégré un vrai shell.</li><li><strong>SSE</strong> — <em>Server-Sent Events</em> : le canal par lequel les réponses arrivent en streaming.</li><li><strong>STT / TTS</strong> — <em>Speech-to-Text</em> (dictée au micro) et <em>Text-to-Speech</em> (lecture à voix haute).</li></ul>` },
        ],
      },
      {
        id: "faq-doc",
        title: "Foire aux questions",
        sections: [
          { h: "Questions fréquentes", html: `<p>La <strong>FAQ</strong> intégrée répond aux questions les plus courantes : clés API, mode Agent, agents parallèles, terminal, confidentialité.</p><p>Ouvrez-la depuis le menu de votre avatar, ou via <strong>Configuration → FAQ</strong>.</p>` },
          { h: "Besoin d'aide ?", html: `<p>Ce centre d'aide couvre l'ensemble de l'application par modules. Utilisez la <strong>recherche</strong> en haut pour trouver un sujet, ou parcourez les collections dans la barre latérale.</p>` },
        ],
      },
    ],
  },
);
