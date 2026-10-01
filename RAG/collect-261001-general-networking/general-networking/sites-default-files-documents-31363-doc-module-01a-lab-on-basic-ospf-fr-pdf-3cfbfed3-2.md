---
id: collect-261001-general-networking/general-networking/sites-default-files-documents-31363-doc-module-01a-lab-on-basic-ospf-fr-pdf-3cfbfed3-2
title: "sites-default-files-documents-31363-doc-module-01a-lab-on-basic-ospf-fr-pdf-3cfbfed3"
domain: general-networking
role: reference
task: reference
actors: []
dates: ["2016-08-12"]
keywords: ["cyber", "memory"]
source: docs/RAG/collect-261001-general-networking/sites-default-files-documents-31363-doc-module-01a-lab-on-basic-ospf-fr-pdf-3cfbfed3.md
source_anchor: ""
source_lines: [144, 277]
sha256: 54297b7050b1f7462b0daa882ee48442506d69714559b55f69b8af1f7ff10cd0
---

# sites-default-files-documents-31363-doc-module-01a-lab-on-basic-ospf-fr-pdf-3cfbfed3

Remarque B: pour IOS antérieures à la version 12.3, la paire nom d'utilis ateur / secret n'est pas 
disponible, et les opérateurs devront configurer le nom d'utili s a t e u r  /  m o t  d e  p a s s e .  C e  d e r n i e r  
format utilise le cryptage de type 7, alors que le premier est un cryptage basé sur  md5 un peu 
mieux sécurisé. IOS 15.1 et version ultérieure  utilisent  SHA256 pour remplacer  MD5. 
 
7. Activation de l'accès de connexion pour les autres équipes. Afin de permettre à d'autres équipes 
telnet accès sur le routeur pour de futurs modules de cet ateli er, vous devez configurer un mot de 
passe pour toutes les lignes de terminal virtuel. 
 
Router1 (config)# aaa new-model 
   Router1 (config)# aaa authentication login default local 
 Router1 (config)# aaa authentication enable default enable 
 
Cette série de commandes indique au routeur de regarder localem ent pour la connexion utilisateur 
standard (le paire nom d'utilisateur /mot de passe configuré pr écédemment), et au niveau local 
configuré enable secret pour le enable login. Par défaut, le lo gin sera activé sur tous les vtys pour 
que d'autres équipes y accèdent. 
 
8. Configurer system logging. Une partie essentielle de tout système d'exploitation Internet est 
d'enregistrer les logs. Le route ur affiche par défaut les logs système sur la console du routeur. 
Toutefois, cela n'est pas souhaitable pour les routeurs Interne t opérationnelles, comme la console 
est une connexion 9600 bauds, et peut placer une charge process eur élevée d'interruption au 
moment de trafic intense sur le réseau. Cependant, les logs du routeur peuvent également être 
enregistrées dans une mémoire t ampon sur le routeur - ne prend pas d’interrupt load et permet 
également à l'opérateur de vérifier l'historique de ce qui s'es t passé sur le routeur. Dans un module 
futur, le laboratoire consistera  à configurer le routeur pour envoyer les messages log vers un 
serveur SYSLOG. 
 
                                                            
1 Cette phrase doit être soulignée. Les cyber-attaquants peuvent fréquemment avoir accès à des réseaux tout simplement 
parce que les opérateurs ont utilisé des mots de passe familiers ou faciles à deviner.

Atelier Labo ISP 
  5   
  
Router1 (config)# no logging console 
   Router1 (config)# logging buffer 8192 debug 
 
qui désactive les console logs et enregistre à la place  tous l es logs dans un tampon 8192 Byte mis 
de côté sur le routeur. Pour voir le contenu de ce tampon loggi n g  i n t e r n e ,  à  t o u t  m o m e n t ,  l a  
commande "sh log" doit être utilisé sur la ligne de commande. 
 
9. Sauvegarder  la  configuration. Grâce à la configuration de base en place, sauv egardez  la 
configuration. Pour faire ça, sortir du mode enable  en tapant «end» ou « <ctrl> Z », et sur la ligne 
de commande, entrez “write memory”. 
 
Router1(config)#^Z 
Router1# write memory 
Construire la configuration... 
[OK] 
Router1# 
 
Il est fortement recommandé que la configuration soit sauvegard ée dans NVRAM assez 
fréquemment, en particulier dans l'environnement atelier où il est possible pour les câbles 
d'alimentation de se détacher. Si la configuration n'est pas en registrée dans NVRAM, toutes les 
modifications apportées à la configuration courante seront perdues après un cycle d'alimentation.  
 
Déconnectez-vous au routeur en tapant exit, puis connectez-vous  à nouveau. Remarquez comment 
la séquence de login a changé,  demandant  l'utilisateur à entr er un «nom d'utilisateur» et «mot de 
passe». Remarquez qu'à chaque point de contrôle dans l'atelier,  vous devez sauvegarder la 
configuration de la mémoire - se rappeler que l’interruption d' alimentation du routeur se traduira 
par revenir à la dernière configuration enregistrée dans la mémoire NVRAM. 
 
10. Adresses IP. Ce module présente les concepts de base pour mettre sur pied u n plan d’adressage 
pour  un backbone ISP. Nous mettons en place un système autonom e sur les 14 routeurs que nous 
avons dans le laboratoire. Les RIR distribuent généralement de l’espace adresses IPv4 en 
morceaux / 20 (dépend de la région RIR) - on suppose pour les b esoins de ce labo que notre ISP a 
reçu un /20. Plutôt que d'utiliser l'espace d'adressage public,  nous allons utiliser une partie du 10/8 
(RFC1918 ou espace d'adressage privé) pour ce labo. Dans le mon de réel de l’Internet, nous 
pouvons utiliser l'espace d'adressage public pour notre infrastructure réseau. 
 
La manière typique dont les ISP divisent leur espace d'adressag e alloué est de le découper en trois 
morceaux. Une pièce est utilisée pour des assignations à la cli entèle, la deuxième pièce est utilisée 
pour les liens infrastructure point-à-point, et le dernier morc eau est utilisé pour les adresses 
d'interface loopback pour l'ensemble de leurs routeurs backbone .   L e  s c h é m a  d e  l a  F i g u r e  2  2  
montre ce  qui se fait habituellement. 
 
 
 
 
 
 
 
 
Figure 2 2 - division du bloc alloué de / 20 entre la clientèle, l'Infrastructure et les Loopbacks 
 10.0.15.223 10.0.14.255 10.0.15.224 10.0.15.255 
Loopbacks Espace d'adressage client  
10.0.0.0/20 Bloc réseau 
10.0.0.0 10.0.15.0 
Infrastructure réseau

Friday, August 12, 2016   
 6 
 
Étudier le plan d'adressage qui a été distribué comme un additi f au présent module d'atelier. 
Remarquez comment l'adressage d’infrastructure commence à 10.0. 15.0 et porte sur un maximum 
de 10.0.15.70 - ce nous laisser de l'espace pour agrandir le ré seau avec plus de liens point à point, 
jusqu'à 10.0.15.223. Remarquez comment nous avons mis de côté u n seul / 27 pour les loopback 
des routeurs - mais nous avons seulement utilisé les 14 adresse s à partir de 241 jusqu'à 254 pour 
notre réseau, ce qui laisse une certaine réserve pour la croiss ance future (non pas que nous avons 
une croissance future prévue pour l’atelier), une proposition t out à fait réaliste pour un backbone 
ISP. En effet, les ISP ont tendance à documenter leurs plans d' adressage dans des fichiers texte ou 
dans des feuilles de calcul (spreadsheets)  - Figure 3  3 ci-dessous  montre un extrait d'un exemple 
typique (à l'aide de notre schéma d'adressage ici). 
 
 
Figure 3 3 - Extrait d'un plan d'adressage ISP 
 
11. Connexions en série Back-to-Back.  Connecter les connexions en série comme dans la figure 1 
Figure 1. TLe côté DCE d'une connexion en série Back-to-Back est config uré avec la commande  
clock rate qui anime le circuit en série. (Les anciennes versions de l'IOS 
utilisaient la commande clockrate, maintenant caché mais toujours fonctionnel.) Vérifier 
le câble physiquement pour voir de quel côté est DCE et lequel est DTE. Sur certains routeurs, la 
commande show controller <interface>  montrera l’état DCE / DTE. Par exemple, sur un 
routeur Cisco 3620, show controllers serial 0/0 va produire un résultat qui affichera si 
le câble connecté au port en série 0/0 est DTE ou DCE. 
 
Une fois que les câbles DTE et DCE ont été déterminées et la co mmande clock rate<  a  été  
appliquée, configurer l'adresse IP (selon le plan d'adressage d iscuté plus tôt) et d'autres 
commandes BCP recommandées qui sont recommandés pour chaque Interface de l’ISP:

Atelier Labo ISP 
  7   
  
Router2(config)# interface serial 1/0 
Router2(config-if)# ip address 100.1.17.1 255.255.255.252 
Router2(config-if)# description 2 Mbps Link to Router4 via DTE/DCE Serial 
Router2(config-if)# bandwidth 2000 
Router2(config-if)# clock rate 2000000 
Router2(config-if)# no ip redirects 
Router2(config-if)# no ip directed-broadcast 
Router2(config-if)# no ip proxy-arp 
Router2(config-if)# no shutdown 
 
