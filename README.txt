Tarea 3 Sistemas Distribuidos

Integrantes:
- Camilo Gonzalez Naranjo 201873550-7
- Zarko Kuljis Gaete 201823523-7
- Jose Montecinos Arismendi 201804541-1

Orden de ejecucion y en que servidor se ejecuta cada archivo: 
1. En dist13 se debe ejecutar el comando "make Servidor1"
2. En dist14 se debe ejecutar el comando "make Servidor2"
3. En dist15 se debe ejecutar el comando "make Servidor3"
4. En dist16 se debe ejecutar el comando "make Brocker"
5. En dist13 se debe ejecutar el comando "make Almirante"
6. En dist14 se debe ejecutar el comando "make Tano"
7. En dist15 se debe ejecutar el comando "make Leia"
Cada uno de estos comandos debe ser ejecutado dentro de la carpeta "INF343-Tarea-3".
Espere a que aparezca texto en la terminal después de cada comando para ejecutar el siguiente.

IMPORTANTE: No realizamos Consistencia Eventual ni Merge, por lo que los comandos UpdateName, UpdateNumber, DeleteCity y GetNumberRebelds podrían ser redirigidos a servidores Fulcrum que no tengan la información de los Planetas y/o Ciudades solicitados, por lo que la ejecucion de los scripts se vería interrumpida. A pesar de esto, los métodos Read your Writes y Monotonic Reads si están implementados.
Se recomienda crear varias veces el mismo par (Planeta, Ciudad) para asegurarse que todos los servidores Fulcrum tengan los datos.