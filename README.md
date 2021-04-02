# api-rest-golang

API REST hecha con Golang

Los endpoint son los siguientes(si se ejecuta desde un entorno local):
* http://localhost:3000/list : en este endpoint recibirá un JSON con un array de números con un formato como el del siguiente ejemplo:

 {"sin clasificar": [3,5,5,6,8,3,4,4,7,7,1,1,2]} 
 
y devolverá :

{
"sin clasificar": [3,5,5,6,8,3,4,4,7,7,1,1,2],
   "clasificado": [1,2,3,4,5,6,7,8,5,3,4,7,1]
            }
* http://localhost:3000/users : Este endpoint recibe dos métodos POST y GETcon el método POST se pueden crear usuarios en la base de datos usando un JSON con un formato como el siguiente :
 {    "name": "lucas",    "email": "pepe@mail.com",    "phone":"321654987"}
 con el método GET se consultan todos los usuarios creados en la base de datos
* http://localhost:3000/users/us?id=# : se puede consultar un usuario en específico usando su id  

* http://localhost:3000/note : que con el método POST permite crear notas por usuario usando JSON con el siguiente formato
{
    "title":"hola",
    "description":"haciendo pruebas",
    "user_id":1
}
* http://localhost:3000/notes ?id=# : con este endpoint se puede consultar las notas realizadas por un usuario 
