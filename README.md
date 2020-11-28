# rwth-ping

:smirk: collecting sexy data

## Installation

If you plan a non-development environment, it is enough to execute `./scripts/first_time.sh`.
It will configure influxdb and start all docker services. This process only needs to be
done once.

If you plan to do further development you will use `./start.sh` to execute the server in debug
mode. You will need to have golang installed and start influxdb manually.

```bash
# get code
$ git clone https://github.com/oltdaniel/rwth-ping.git
$ cd rwth-ping
# start influxdb
$ docker-compose up -d influxdb
# setup influxdb
$ ./scripts/setup_influxdb.sh
# start server
$ ./start.sh
```

### `docker-compose`

```yaml
version: '3.3'

services: 
    app:
        image: docker.pkg.github.com/oltdaniel/rwth-ping/server:main
        environment:
            D: 1
            HOST: https://somehost.oltdaniel.at
        env_file:
            - influxdb.env
        volumes:
            - "./config.yml:/config.yml"
        ports:
            - 4001:4001

    influxdb:
        image: quay.io/influxdb/influxdb:2.0.0-rc
        restart: always
        ports:
            - 8086:8086
```

```bash
# create the dokcer-compose.yml
$ nano docker-compose.yml
# create influxdb.env file
$ touch influxdb.env
# start influxdb
$ docker-compose up -d influxdb
# visit HOST:8086
# create buckets and token
# add details to influxdb.env
# SEE: influxdb_example.yml
$ nano influxdb.env
# start other services
$ docker-compose up -d
```

## Usage

> **NOTE**: This is only a research project. It should not be used to monitor the university service
peroformance without their notice. I am not responsible for any damage or complaints caused by my software.

## Copyright

MIT License.

_just do, what you'd like to do_