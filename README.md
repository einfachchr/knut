# Knut #
![alt text][logo]
[logo]: https://de.wikipedia.org/wiki/Knut_(Eisbär)#/media/File:Knut_IMG_8095.jpg "Das Maskottchen"


## 1. Stufe: ##
* 3 Services
* natürlich alle in Docker, Start über docker-compose
* red & blue
  * Kopien voneinander
  * lauschen auf / , Port 80
  * Dumpen Header und Parameter in die Reponse
* green
  * nginx
  * /red   -> proxy_pass
  * /blue  -> proxy_pass
  * /      -> statische Welcome-Seite
* alle 3 liefern dasselbe Grundgerüst, Menü oben, Content unten
* fixe Links tauschen den iframe

### Schritte ###
  FERTIG red
    FERTIG Greeter kopieren
    FERTIG Routing auf / ändern
    FERTIG Template auf Rahmen ändern
    FERTIG HEADER dumpen
    FERTIG Request-Parameter dumpen
    FERTIG Dockerfile
    FERTIG docker-compose.yml

  FERTIG cp red -> blue
    FERTIG Namensänderung

  OFFEN nginx

  OFFEN Integration
    OFFEN iframe muss auf den entsprechenden Content-Endpunkt zeigen
