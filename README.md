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

## Usage

> **NOTE**: This is only a research project. It should not be used to monitor the university service
peroformance without their notice. I am not responsible for any damage or complaints caused by my software.

## Copyright

MIT License.

_just do, what you'd like to do_