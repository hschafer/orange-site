# orange-site

A demo project for full-stack development

## Running

Everything should be contained in pre-defined Docker images. Should be able to get the service up and running with

```bash
docker-compose up
```

And then navigating to a browser and accessing port 80 (e.g., `localhost:80`).

### Configuration

In order to run, need to specify a `.env` file with mostly secret constants. An example of the constants and their meaning is shown in `dot.env`.

### Pre-populate

To give the app some initial data, there is a file `test_db.sql` that has some `INSERT` statements to give some initial data. Once the docker containers are started, access the database container and paste these insert statements there to run them. Note that it won't be possible to log in to those dummy accounts as they use dummy plain-text passwords, while the service assumes the passwords stored are hashed.

The following commands are an example for how to access the database container.

```bash
$> docker ps
# Output will look something like
# CONTAINER ID   IMAGE                  COMMAND                  CREATED       STATUS          PORTS                    NAMES
# f71ccda06eda   nginx                  "/docker-entrypoint.…"   2 hours ago   Up 2 hours      0.0.0.0:80->80/tcp       orange-site-nginx-1
# bc63fa32a0a4   orange-site-frontend   "docker-entrypoint.s…"   2 hours ago   Up 2 hours      0.0.0.0:8080->8080/tcp   orange-site-frontend-1
# f3a43e88e671   golang:1.21            "go run main.go"         2 hours ago   Up 22 minutes   0.0.0.0:3000->3000/tcp   backend
# 4c16a92e8475   orange-site-db         "docker-entrypoint.s…"   3 hours ago   Up 2 hours      5432/tcp                 orange-site-db-1

$> docker exec -it 4c16a92e8475 bash
4c16a92e8475:/# psql db user

# Should match the user name in the config file, now you can paste in INSERT statements
```

## Acknowledgments

- Some of the overall structure and infrastructure to get an app up and running is attributed to [this example project by ndabAP](https://github.com/ndabAP/vue-go-example/tree/master)
