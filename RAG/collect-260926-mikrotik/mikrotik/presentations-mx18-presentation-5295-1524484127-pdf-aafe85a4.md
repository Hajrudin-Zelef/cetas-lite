---
id: collect-260926-mikrotik/mikrotik/presentations-mx18-presentation-5295-1524484127-pdf-aafe85a4
title: "presentations-mx18-presentation-5295-1524484127-pdf-aafe85a4"
domain: mikrotik
role: reference
task: reference
actors: ["Microsoft"]
dates: []
keywords: ["mcp"]
source: docs/RAG/lot-mikrotik/RouterOS/presentations-mx18-presentation-5295-1524484127-pdf-aafe85a4.md
source_anchor: ""
source_lines: [1, 149]
sha256: 485aac333b46045f2673b1a0cadf26f4b04f013045459efb12eeebc468117a9f
---

# presentations-mx18-presentation-5295-1524484127-pdf-aafe85a4

Hotspot + CAPsMAN + 
VPN en la nube
MTI David Ricardo López Aldret
MUM México 2018

Acerca de mí
Mágister IT 
Mikrotik MTCNA, MTCTCE, MTCWE, MTCRE certiﬁed 
Ubiquiti UAC certiﬁed 
Microsoft MCSE, MCSA, MCP certiﬁed 
Cambium Networks ePMP , PtP 650 certiﬁed 
CompTIA A+, Security+, Network+, Server+ certiﬁed  
ETA-I CST, CNST certiﬁed 
Access Data ACE certiﬁed 
Denwa DeECP certiﬁed

Acerca de mí

Objetivo
Mostrar de una forma clara lo sencillo que es 
implementar uno o múltiples hotspots que sean 
controlados en la nube mediante herramientas propias 
de Mikrotik (RouterOS, User Manager y CAPsMAN) y 
accesibles a través de una VPN. 
Se omiten conﬁguraciones obvias y personalizadas.

Qué es CAPsMan
Es un controlador Wireless, a través de el, los AP’s que 
se encuentren registrados en el, serán administrados y 
manejados de forma centralizada. Permite despliegues 
masivos de equipos inalámbricos con conﬁguraciones 
homogéneas.

Qué es Hotspot
Sencillamente, es un servicio que ofrece conexión a 
internet a través de una red inalámbrica, pueden ser 
gratuitos o de paga.

Qué es VPN
“Red privada virtual” , permite una conexión segura 
entre dos puntos distantes entre si y que permiten 
interactuar como si se tratara de una red local. Existen 
varios tipos de VPN que han surgido con los años.

Qué es UserManager
Es una aplicación RADIUS desarrollada por Mikrotik que 
cumple con los estándares de AAA (Autorización, 
Autenticación y Contabilización). A través de el es 
posible administrar servicios PPP , Hotspot, DHCP , 
Wireless y RouterOS

Escenario

Ventajas de un escenario así
Cada Cap+Hotspot tiene una salida propia a internet 
Cada Cap+Hotspot está centralizado a un solo CAPsMAN 
y UserManager, ergo, solo es una generación de 
vouchers 
Roaming de usuarios con el mismo voucher 
Cambios en caliente en toda la red

Desventajas
CAPsMAN se encuentra estacionado a nivel de 
desarrollo, no han habido mejoras ni agregados.

Desventajas
¿Se cayó EL CHR? eso es algo malo, malo….malo

¿Qué necesitamos?
Licencia 
 AP
VPS

Instalamos RouterOS
https://www.digitalocean.com/community/questions/
installing-mikrotik-routeros 
Creamos un droplet y pegamos el script

Instalamos RouterOS
Usuario: root  y contraseña: xxxxxx 
Upgradeamos 
Realizamos las conﬁguraciones básicas (cambio de 
password, dns, etc) 
Descargamos el paquete UserManager y lo instalamos

Configuramos nuestros AP’s
Damos salida a internet (DNS, Gateway, IP , etc) 
Actualizamos y aseguramos nuestro RouterOS 
Personalizamos

Establecemos VPN (server)

Establecemos VPN (server)

Establecemos VPN (cliente)

Establecemos VPN (cliente)

Establecemos VPN (server)

Configuramos CAPsMAN 
(VPS)

Configuramos CAP

Establecemos conexión y 
personalizamos

Establecemos conexión y 
personalizamos

Establecemos conexión y 
personalizamos

Hotspot
MUY IMPORTANTE: 
La conﬁguración del HS DEBE hacerse sobre un BRIDGE 
Al ser la interfaz Wlan esclava del CAPsMAN no 
permitirá crearse ahí.

Hotspot

Hotspot
Generamos hotspot   
Upgradeamos 
Realizamos las conﬁguraciones básicas (cambio de 
password, dns, etc) 
Descargamos el paquete UserManager y lo instalamos

Hotspot 
vinculamos RADIUS

Hotspot vinculamos router a 
RADIUS

Hotspot 
Creamos perfil

Hotspot 
creamos perfil

Hotspot 
Usamos RADIUS en hs

Hotspot

¡Listo!

¿Dudas?
