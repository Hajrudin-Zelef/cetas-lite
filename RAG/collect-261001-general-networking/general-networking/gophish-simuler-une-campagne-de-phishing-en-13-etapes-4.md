---
id: collect-261001-general-networking/general-networking/gophish-simuler-une-campagne-de-phishing-en-13-etapes-4
title: "1. Mise à jour système"
domain: general-networking
role: reference
task: reference
actors: []
dates: []
keywords: ["apache", "arr", "exploit", "incident"]
source: docs/RAG/collect-261001-general-networking/gophish-simuler-une-campagne-de-phishing-en-13-etapes.md
source_anchor: ""
source_lines: [256, 303]
sha256: 6652544a9877640b48675c29b57c228618b88b8493ec93bd68f7832cc16bf7db
---

# 1. Mise à jour système

Ces avis renforcent un principe simple : ne jamais exposer l'interface d'administration de GoPhish sur Internet, restreindre les droits de création de modèles aux seuls administrateurs de confiance, et surveiller le dépôt GitHub pour toute mise à jour au-delà de la version 0.12.1. Ce dernier point mérite une vigilance particulière : contrairement à des outils commerciaux qui poussent des correctifs automatiques, GoPhish étant un projet communautaire, c'est à chaque administrateur de vérifier périodiquement l'existence d'une nouvelle version et de l'appliquer manuellement.

Une bonne pratique complémentaire consiste à isoler entièrement le serveur GoPhish dans un réseau dédié, séparé du reste de l'infrastructure de production, y compris du contrôleur de domaine ou des serveurs de fichiers. Ainsi, même dans le pire scénario où une des vulnérabilités listées ci-dessus serait exploitée par un tiers malveillant, l'impact resterait cantonné à l'outil de test lui-même plutôt que de se propager vers les systèmes critiques de l'entreprise.

## Dépannage : 8 problèmes courants et leurs solutions

**Le service systemd échoue au démarrage avec « unable to open database file ».** Vérifiez que WorkingDirectory dans l'unité systemd correspond exactement au dossier contenant gophish.db et config.json. Un chemin relatif mal résolu est la cause la plus fréquente.

**Impossible d'accéder à l'interface admin sur le port 3333.** Si vous avez limité listen_url à 127.0.0.1, ouvrez un tunnel SSH local avant de tenter la connexion depuis votre navigateur.

**Le navigateur affiche une alerte de certificat invalide.** Le certificat par défaut généré par GoPhish est auto-signé. Remplacez cert_path et key_path par un certificat Let's Encrypt valide, ou acceptez l'exception uniquement en environnement de test isolé.

**Le port 80 est déjà utilisé par un autre service (Apache, Nginx).** Modifiez phish_server.listen_url vers un autre port, par exemple 0.0.0.0:8080, et redirigez le trafic via un reverse proxy, ou arrêtez le service en conflit.

**Les emails de test n'arrivent jamais.** Vérifiez les enregistrements SPF, DKIM et DMARC du domaine d'expédition, testez l'envoi via Send Test Email dans le profil SMTP, et consultez les logs du relais pour un rejet ou une mise en quarantaine.

**Le taux de clic affiché semble anormalement élevé.** Certains scanners antivirus d'entreprise ou passerelles de sécurité email cliquent automatiquement sur les liens pour les analyser (URL rewriting). Filtrez ces IP connues dans vos statistiques finales pour ne pas fausser l'analyse.

**L'import CSV des cibles échoue silencieusement.** Vérifiez l'encodage du fichier (UTF-8 sans BOM) et l'orthographe exacte des en-têtes de colonnes attendues par GoPhish (Email est le seul champ strictement obligatoire).

**La page d'atterrissage importée via « Import Site » affiche des ressources cassées.** L'import clone le HTML mais pas toujours les chemins relatifs vers les images et scripts. Passez les URL des ressources en absolu manuellement après l'import.

**Le mot de passe admin initial est perdu.** Arrêtez le service, sauvegardez gophish.db, puis réinitialisez le mot de passe via la ligne de commande officielle documentée dans le guide utilisateur, ou supprimez la base pour repartir d'un état vierge si aucune campagne n'a encore été lancée. Dans tous les cas, effectuez d'abord une copie du fichier de base de données : une manipulation malheureuse à ce stade peut effacer l'historique complet de vos campagnes précédentes, y compris les résultats déjà exploités pour vos rapports internes.

**Les statistiques de la campagne restent bloquées à zéro plusieurs heures après le lancement.** Vérifiez d'abord que le job d'envoi n'est pas resté en file d'attente côté relais SMTP en consultant ses propres logs. Si les emails sont bien partis mais qu'aucun clic n'est enregistré, assurez-vous que le lien de tracking généré par GoPhish pointe vers une adresse IP ou un nom de domaine réellement joignable depuis l'extérieur, et non vers une adresse interne au réseau du serveur.

## Astuces avancées pour des campagnes plus réalistes

Segmentez vos campagnes par métier plutôt que d'envoyer le même email à toute l'entreprise. Un email de fausse facture fournisseur sera plus efficace auprès de la comptabilité, tandis qu'une fausse alerte RH fonctionnera mieux auprès des équipes support. Cette segmentation reproduit fidèlement la manière dont un attaquant cible réellement une organisation après une phase de reconnaissance, où il adapte son scénario au métier de la victime plutôt que d'envoyer un message générique à toute l'entreprise.

Programmez des campagnes récurrentes espacées de plusieurs mois plutôt qu'un exercice unique annuel. La répétition régulière ancre mieux les réflexes que des sessions ponctuelles, même très élaborées. Un rythme trimestriel avec des scénarios différents à chaque fois (fausse facture, fausse alerte de sécurité, fausse notification de livraison) évite également l'effet d'usure où les employés reconnaissent le même modèle d'un test à l'autre et faussent artificiellement les statistiques de vigilance.

Couplez GoPhish à un webhook vers votre plateforme de tickets pour déclencher automatiquement un module de micro-formation dès qu'un utilisateur soumet des données sur la page piégée, plutôt que d'attendre un débriefing collectif en fin de campagne. L'API REST de GoPhish permet d'interroger l'état d'une campagne en continu, ce qui rend possible ce type d'automatisation sans intervention manuelle de l'équipe sécurité à chaque événement.

Variez les vecteurs : email classique, lien raccourci, QR code imprimé sur une fausse note de service. Les attaquants de 2026 ne se limitent plus au seul email, et une campagne de sensibilisation complète devrait refléter cette diversification, même si GoPhish reste centré sur le vecteur email et landing page.

Documentez chaque campagne dans un registre interne : date, scénario utilisé, périmètre, taux de clic, taux de soumission, taux de signalement. Ce registre devient un historique précieux pour justifier des investissements en formation auprès de la direction, et il constitue également une preuve de diligence en cas de contrôle ou d'incident réel de phishing qui toucherait l'entreprise entre deux campagnes.

Enfin, formez un petit groupe de « champions sécurité » dans chaque service, chargés de relayer les bonnes pratiques et de répondre aux questions de leurs collègues après une campagne. Cette approche par les pairs a généralement plus d'impact qu'une communication descendante unique venant de la DSI, car elle normalise la vigilance comme un comportement d'équipe plutôt qu'une contrainte imposée d'en haut.

## Mesurer le retour sur investissement d'un programme de sensibilisation

Justifier le temps passé à administrer GoPhish auprès de la direction demande de traduire les résultats bruts en indicateurs compréhensibles pour un comité de direction qui ne lit pas de logs. Trois métriques suffisent généralement à construire ce dossier.

La première est l'évolution du taux de clic sur plusieurs campagnes consécutives. Une baisse progressive, même modeste, démontre que l'effort de sensibilisation porte ses fruits et justifie la poursuite du programme plutôt que son abandon au bout d'un ou deux exercices.

La deuxième est le taux de signalement, souvent plus parlant que le taux de clic lui-même. Un employé qui clique par réflexe mais signale ensuite l'email suspect a compris le message de fond, même s'il n'a pas eu le bon réflexe au premier contact. Suivre ce taux dans le temps permet de mesurer la maturité réelle de l'organisation plutôt que sa seule capacité à éviter un piège ponctuel.

