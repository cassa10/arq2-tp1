# Arquitectura de Software 2 - Universidad Nacional de Quilmes

[Enunciado del Trabajo Práctico](https://github.com/cassa10/arq2-tp1/blob/main/doc/Arq2%20-%20Trabajo%20pr%C3%A1ctico.pdf)


## Test y coverage

### Prerequisitos:

- Go 1.20 or up

### Pasos:

1) Ir al folder root del repositorio

2) Ejecutar los comandos
```
> go test -coverprofile="coverage.out" -covermode=atomic ./...
> go install gitlab.com/fgmarand/gocoverstats@latest
> gocoverstats -v -f coverage.out > coverage_rates.out
```

3) Se generará el file coverage.out el cual contedrá el info del coverage, y 
overage_rates.out que contendrá los porcentajes de coverage en decimales [0.0 - 1].

**Nota**: En la construcción de la imagen, se realiza la ejecucion de test y la generacion del coverage.
Los archivos de coverage son almacenados dentro del container en la carpeta "/app". Es decir, que encontraremos:
- /app/coverage.out
- /app/coverage_rates.out

Dichos archivos se pueden acceder desde docker desktop (recomendado), o bien, vinculando el filesystem del container con un volume y accediendo a dicho volume.

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
docker run -p <port>:8080 --name arq2-tp1 arq2-tp1
```

Nota: agregar "-d" si se quiere ejecutar como deamon

```
docker run -d -p <port>:8080 --name arq2-tp1 arq2-tp1
```

Ejemplo:

```
docker run -d -p 8082:8080 --name arq2-tp1 arq2-tp1
```

4) En un browser, abrir swagger del servicio en el siguiente url:

`http://localhost:<port>/docs/index.html`

Segun el ejemplo:

`http://localhost:8082/docs/index.html`

5) Probar el endpoint health check y debe retornar ok

6) La API esta disponible para ser utilizada

