# ETL de servicios
Esta es una aplicación desarrollada en Go y Vue.js, con la cual se pueden obtener los datos de los servicios que brinda la Universidad de Santiago de Chile que son registrados en una plataforma desarrollada en WordPress, 
con la cual también se pueden realizar ediciones sobre estos datos con el fin de asegurar la consistencia de la información proporcionada y finalmente actualizar la base de datos de la misma plataforma.

## Requisitos
- Docker

- De manera provisoria la aplicación para poder ser ejecutada de forma correcta se debe tener creada una base de datos MySQL llamada "test_usachatiendebd" de la cual se van a extraer los datos con los cuales trabaja la aplicación.

## Instrucciones de uso
Luego de clonar el repositorio se debe ubicar en la carpeta "Pingeso_2023"

`cd Pingeso_2023 `

Luego se procede a construir la imagen

` docker compose build `

Finalmente se construyen y levantan los contenedores de la aplicación

` docker compose up `







