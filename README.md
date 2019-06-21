# snmpbeat
Custom beat for snmp monitoring in firewall
## Instalacion de GO
Link : https://linuxize.com/post/how-to-install-go-on-centos-7/

Descargar goland usado "wget" o "curl"

> wget https://dl.google.com/go/go1.10.3.linux-amd64.tar.gz

Verificar el tarball

> sha256sum go1.10.3.linux-amd64.tar.gz

El checksum deberia arrojar lo siguiente:

> fa1b0e45d3b647c252f51f5e1204aba049cde4af177ef9f2181f43004f901035  go1.10.3.linux-amd64.tar.gz

Extraer el archivo en "/usr/local/"

> fa1b0e45d3b647c252f51f5e1204aba049cde4af177ef9f2181f43004f901035  go1.10.3.linux-amd64.tar.gz

Ajustar la variable de entorno

> export GOROOT=/usr/local/go
> export PATH=$PATH:/usr/local/go/bin

Guardando el archivo y cargar el new PATH  a las variables de entorno

> source ~/.bash_profile

## Compilación de snmpbeat


### Árbol de carpetas

>- $GOPATH: ruta al espacio de trabajo

>   - snmpbeat.yml : Archivo de configuración requerido por snmpbeat

>   - [src] : Carpetas que contiene el codigo fuente a ser compilado

>       - [github.com]  :Carpeta adicional para agregar el codigo de github

>          - [hoat23]

>             - [snmpbeat] :Carpeta con el codigo de <snmpbeat
>

>          - [elastic]

>             - [beats] :Carpeta con el codigo adicional para compilar el beats

>                - [libbeat]      :Libreria muy importante para la compilación del código

>                - [auditbeat]

>                - [heartbeat]

>                - [metricbeat]

>                - [packetbeat]

>                - [winlogbeat]

>                - [deploy]

>                - [dev-tools]

>                - [docs]

>                - [filebeat]

>                - [generator]

>                - [journalbeat]

>                - [licenses]

>                - [script]

>                - [testing]

>                - [vendor]

>                - .editorconfigj

>                - .gitattributes

>                - .gitnore

>                - .go-version

>                - .travis.yml

>                - .PYLITRC

>                - Jenkisfile

>                - LICENCE.txt

>                - magefile.go    : Archivo make para la compilación con go

>                - MAKE.BAT

>                - mAKEFILE

>                - NOTICE.txt

>                - reviewdog.yml

>                - setup.yml


                


   
   


