# Oracle Relayer

Oracle Relayer is responsible for relaying events of bsc to binance chain.

## Build

Build binary:

```shell script
$ make build
```

Build docker image:

```shell script
$ make build_docker
```

## Config

There is a config template in config directory, you should create your own config to run your relayer correctly. 
You can refer to [config doc](./docs/config.md) for more details.

### Recommendations
If you are going to deploy your oracle relayer in production, AWS Secret Manager is recommended. You can use AWS Secret
Manager to host your mnemonic. 

For BBC and BSC providers, you should use trusted nodes and TLS connection is recommended.

## Run

Run locally:

```shell script
$ ./build/relayer --bbc-network [0 for testnet, 1 for mainnet] --config-type [local or aws] --config-path config_file_path --aws-region [aws region or omit] --aws-secret-key [aws secret key for config or omit]
```

Run docker:
```shell script
$ docker run -it -v /your/data/path:/relayer -e BBC_NETWORK={0 or 1} -e CONFIG_TYPE="local" -e CONFIG_FILE_PATH=/your/config/file/path/in/container -d oracle_relayer
```

## License

Distributed under the [GNU Lesser General Public License v3.0](https://www.gnu.org/licenses/lgpl-3.0.en.html). See [LICENSE](LICENSE) for more information.
