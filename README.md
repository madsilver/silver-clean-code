# Silver Clean Code

![Badge](https://img.shields.io/badge/Go-v1.19-blue)
[![Quality Gate Status](https://sonarcloud.io/api/project_badges/measure?project=madsilver_silver-clean-code&metric=alert_status)](https://sonarcloud.io/summary/new_code?id=madsilver_silver-clean-code)
[![Duplicated Lines (%)](https://sonarcloud.io/api/project_badges/measure?project=madsilver_silver-clean-code&metric=duplicated_lines_density)](https://sonarcloud.io/summary/new_code?id=madsilver_silver-clean-code)
[![Code Smells](https://sonarcloud.io/api/project_badges/measure?project=madsilver_silver-clean-code&metric=code_smells)](https://sonarcloud.io/summary/new_code?id=madsilver_silver-clean-code)
[![Security Rating](https://sonarcloud.io/api/project_badges/measure?project=madsilver_silver-clean-code&metric=security_rating)](https://sonarcloud.io/summary/new_code?id=madsilver_silver-clean-code)
[![Coverage](https://sonarcloud.io/api/project_badges/measure?project=madsilver_silver-clean-code&metric=coverage)](https://sonarcloud.io/summary/new_code?id=madsilver_silver-clean-code)

This project is a practical exercise in using the **_clean code pattern_** for application architecture.
Feel free to submit suggestions and pull requests.

## Configuration
### Environment variables
| variable       | default    | description                             |
|----------------|------------|-----------------------------------------|
| SERVER_NAME    | echo       | server http framework name (echo / gin) |
| SERVER_PORT    | 8000       | server port                             |
| MYSQL_USER     | silver     |                                         |
| MYSQL_PASSWORD | silver     |                                         |
| MYSQL_DATABASE | silverlabs |                                         |
| MYSQL_HOST     | localhost  |                                         |
| MYSQL_PORT     | 3306       |                                         |
| POSTGRES_USER  | silver     |                                         |
| POSTGRES_PASSWORD | silver     |                                         |
| POSTGRES_DATABASE | silverlabs |                                         |
| POSTGRES_HOST     | localhost  |                                         |
| POSTGRES_PORT     | 5432       |                                         |

## Usage
### Start using it
```shell
make run
```

### Makefile
Use ``make help`` or only ``make`` to check all the available commands.

## API Documentation
http://localhost:8000/swagger/
