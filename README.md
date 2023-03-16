# Arquitectura de Software 2 - Universidad Nacional de Quilmes

[Enunciado del Trabajo Práctico](https://github.com/cassa10/arq2-tp1/blob/main/doc/Arq2%20-%20Trabajo%20pr%C3%A1ctico.pdf)


## Swagger

Para actualizar la api doc de swagger, ejecutar en el folder root del repo:

```
swag init -g src/infrastructure/api/app.go
```

Luego de levantar la api, ir al endpoint:

```
http://localhost:<port>/docs/index.html
```


## Inicialización y ejecución del proyecto

### Prerequisitos:

- Docker

### Pasos:

1) Ir a la carpeta root del repositorio

2) Construir el Dockerfile (imagen) del servicio

```
docker build -t arq2-tp1 .
```

3) Ejecutar la imagen construida

```
docker run -p <port>:8080 arq2-tp1
```

Nota: agregar "-d" si se quiere ejecutar como deamon

```
docker run -d -p <port>:8080 arq2-tp1
```

Ejemplo:

```
docker run -d -p 8082:8080 arq2-tp1
```

4) En un browser, abrir swagger del servicio en el siguiente url: 
`http://localhost:<port>/docs/index.html`

Segun el ejemplo: 
`http://localhost:8082/docs/index.html`

5) Probar el endpoint health check y debe retornar ok

