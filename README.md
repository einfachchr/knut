# Knut #
![picture alt](https://de.wikipedia.org/wiki/Knut_(Eisb%C3%A4r)#/media/File:Knut_IMG_8095.jpg "Das Maskottchen")


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
  [X] red
    [X] Greeter kopieren
    [X] Routing auf / ändern
    [X] Template auf Rahmen ändern
    [X] HEADER dumpen
    [X] Request-Parameter dumpen
    [X] Dockerfile
    [X] docker-compose.yml

  [ ] cp red -> blue
    [ ] Namensänderung

  [ ] nginx

  [ ] Integration
    [ ] iframe muss auf den entsprechenden Content-Endpunkt zeigen
