# snmpbeat
Custom beat for snmp monitoring in firewall
## Instalacion de GO
Link : https://linuxize.com/post/how-to-install-go-on-centos-7/
Descargar goland usado "wget" o "curl"
$ wget https://dl.google.com/go/go1.10.3.linux-amd64.tar.gz

Verificar el tarball

$ sha256sum go1.10.3.linux-amd64.tar.gz

El checksum deberia arrojar lo siguiente:

> fa1b0e45d3b647c252f51f5e1204aba049cde4af177ef9f2181f43004f901035  go1.10.3.linux-amd64.tar.gz

Extraer el archivo en "/usr/local/"

$ fa1b0e45d3b647c252f51f5e1204aba049cde4af177ef9f2181f43004f901035  go1.10.3.linux-amd64.tar.gz

Ajustar la variable de entorno

$ export GOROOT=/usr/local/go
$ export PATH=$PATH:/usr/local/go/bin

Guardando el archivo y cargar el new PATH  a las variables de entorno

$ source ~/.bash_profile

## Compilación de snmpbeat


### Árbol de carpetas
-> $GOPATH: ruta al espacio de trabajo
   |_ snmpbeat.yml : Archivo de configuración requerido por snmpbeat
   |_ [src] : Carpetas que contiene el codigo fuente a ser compilado
       |_ [github.com]  :Carpeta adicional para agregar el codigo de github
          |_ [hoat23]
             |_ [snmpbeat] :Carpeta con el codigo de <snmpbeat>
          |_ [elastic]
             |_ [beats] :Carpeta con el codigo adicional para compilar el beats
                |_[libbeat]      :Libreria muy importante para la compilación del código
                |_[auditbeat]
                |_[heartbeat]
                |_[metricbeat]
                |_[packetbeat]
                |_[winlogbeat]
                |_[deploy]
                |_[dev-tools]
                |_[docs]
                |_[filebeat]
                |_[generator]
                |_[journalbeat]
                |_[licenses]
                |_[script]
                |_[testing]
                |_[vendor]
                |_.editorconfigj
                |_.gitattributes
                |_.gitnore
                |_.go-version
                |_.travis.yml
                |_.PYLITRC
                |_Jenkisfile
                |_LICENCE.txt
                |_magefile.go    : Archivo make para la compilación con go
                |_MAKE.BAT
                |_mAKEFILE
                |_NOTICE.txt
                |_reviewdog.yml
                |_setup.yml
                


   
   


