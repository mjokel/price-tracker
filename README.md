# price-tracker
GitHub mirror of an application for tracking prices at German petrol stations, written in Go.

### First Interaction

Build Docker image from Dockerfile with
```sh
docker build --tag price_tracker .
```

Navigate to preferred mount-binding directory. Make script executable and initialize SQLite database with 
```sh
chmod +x init_db.sh
./init_db.sh
```

Finally, run application with 
```sh
docker run --rm -d \
    -v /tmp/price-tracker:/app/db \
    -e DB_PATH="./db/data.db" \
    -e API_KEY="..." \
    price_tracker
```

Note: 
- ensure to specify an **absolute** path on the host-side of the mount-binding statement (`-v`, left side)!
- `DB_PATH` is the path to the SQLite database in the container, which depends on the mounting point inside the container (`-v`, right side)
- `API_KEY` is the [Tankerkönig](https://creativecommons.tankerkoenig.de/) Creative Commons API key


Note the other (optional) flags:
- `--rm`: automatically remove the container when it exits
- `-d`: container starts in the background and does not attach its standard input, output, and error streams to the terminal ("detached")
- (`-i`: keeps STDIN open even if not attached ("interactive"))
- (`-t`: allocates a pseudo-TTY and allows the interaction with the running container ("tty"))

To check what is going on inside the container, use the following command and with the returned `id`:
```
docker exec -it <id> /bin/bash
```
