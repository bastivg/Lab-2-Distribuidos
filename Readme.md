# Grupo26-Laboratorio-2
Los integrantes son:
- Daniel Parraguez - 202073542-5
- Iván Oyarzún - 202073612-k
- Bastian Varas - 201856629-2


# Información para la Ejecución de Archivos
Para la correcta ejecucion de este laboratorio, cada datanode, continentes, ONU y OMS, sus contenedores ya se encuentran en las maquinas
entregadas (dist129, dist102, dist103, dist104), por lo tanto, solo se deben ejecutar los siguientes comandos:

- `make docker-ONU`: Para ejecutar la consulta de la ONU.
- `make docker-OMS`: Para ejecutar el servidor OMS.
- `make docker-datanode`: Para ejecutar los servidores de los dos DataNodes
- `make docker-continente`: Para ejecutar a los continentes.

De todas formas cada carpeta tiene su nombre indicativo que contiene los archivos necesarios para cada parte del programa, de todas formas, se tienen ramas con los nombres de cada máquina con sus respectivos archivos.

# Cosas a tener en cuenta
Al ejecutar el programa con el make ya descrito anteriormente, si aparece un error como este: `Error response from daemon: Conflict. The container name "/oms" is already in use by container`,
se debe ejecutar en este caso `sudo docker rm <nombre>`, y luego volver a ejecutar el make. Además, hay que considerar que la OMS y los DataNodes en gRPC actuan como servidores, mientras que los continentes
y la ONU actuan como clientes, por lo tanto, para su correcta ejecución se deben iniciar primero los servidores OMS y DataNodes y al final los continentes con la ONU.

Nombres de las imagenes creadas:

- Servidor DataNode1: datanodeone
- Servidor DataNode2: datanodetwo
- Servidor OMS: oms
- ONU: onu
- Europa: europa
- Asia: asia
- Australia: australia
- Latinoamerica: latam

En caso de que los contenedores se borren se puede ejecutar el comando make build-docker para crear los contenedores de los servidores respectivos.

# Ubicación de los servidores:
Cada archivo tiene una carpeta especifica si se desea ejecutar la OMS esta estara ubicada en la carpeta "OMS", se dejara entre parentesis el nombre de la carpeta donde este ubicada:

- dist129: OMS(OMS) y Europa(Europa)
- dist102: Latinoamerica(Latinoamerica) y ONU(ONU)
- dist103: DataNode1(DataNode1) y Australia(Australia)
- dist104: DataNode2(DataNode2) y Asia(Asia)

