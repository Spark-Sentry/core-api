# SparkSentry API :zap:

Welcome to the SparkSentry API repository, an innovative energy management system designed to optimize your energy consumption.

## Features :star:

SparkSentry offers a robust platform for energy data collection and analysis, enabling users to monitor and optimize their energy consumption efficiently. It now supports flexible user-account associations, account creation, detailed management of buildings, areas, systems, and equipment, catering to a wide range of energy management needs.

## Setup :gear:

To set up and run the project locally, follow these steps:

1. Clone the repository to your local machine.
2. Ensure you have Go installed on your machine.
3. Install dependencies by running `go mod tidy`.
4. Configure the necessary environment variables in a `.env` file based on the `.env.example` file.
5. Launch the application with `go run ./cmd/sparksentry/main.go`.

## API Routes :world_map:

The SparkSentry API exposes the following routes:

### Authentication
- `POST /api/v1/login`: Log in to the application and receive a JWT token. :key:

### User Management (JWT token required)
- `POST /api/v1/register`: Register a new user into the system. :bust_in_silhouette:
- `GET /api/v1/users/me`: Get the authenticated user's information. :bust_in_silhouette:

### Account Management (JWT token required)
- `POST /api/v1/accounts`: Create a new account. :office:
- `POST /api/v1/accounts/users`: Associate an existing user to an account. :link:

### Building Management (JWT token required)
- `POST /api/v1/buildings`: Create a new building with areas. :house_with_garden:
- `GET /api/v1/buildings`: Retrieve all buildings associated with the authenticated account's ID. :houses:

### Area Management (JWT token required)
- `POST /api/v1/buildings/:building_id/areas`: Add a new area to a specific building. :park:
- `GET /api/v1/buildings/:building_id/areas`: List all areas within a specific building. :national_park:
- `PUT /api/v1/areas/:area_id`: Update details of a specific area. :construction:
- `DELETE /api/v1/areas/:area_id`: Remove a specific area. :no_entry:

### System Management (JWT token required)
- `POST /api/v1/buildings/:building_id/areas/:area_id/systems`: Add a new system to a specific area within a building. :gear:
- `GET /api/v1/buildings/:building_id/areas/:area_id/systems`: Retrieve all systems associated with a specific area within a building. :wrench:
- `PUT /api/v1/systems/:system_id`: Update details of a specific system. :hammer_and_wrench:
- `DELETE /api/v1/systems/:system_id`: Remove a specific system. :no_entry_sign:

### Equipment Management (JWT token required)
- `POST /systems/:system_id/equipments`: Add new equipment to a specific system. :heavy_plus_sign:
- `GET /systems/:system_id/equipments`: List all equipments within a specific system. :clipboard:
- `PUT /equipments/:equipment_id`: Update details of a specific piece of equipment. :memo:
- `DELETE /equipments/:equipment_id`: Remove a specific piece of equipment. :wastebasket:

## Hot Reloading with Air :fire:

This project uses [Air](https://github.com/cosmtrek/air) for hot reloading during development. Air automatically rebuilds and restarts your application when file changes in the directory are detected, making development faster and more efficient.

To use Air, ensure you have the `.air.toml` configuration file at the root of your project, then simply run `air` in your terminal within the project directory.

## Environment Variables :key:

- `DB_USER`: Your database username
- `DB_PASSWORD`: Your database password
- `DB_NAME`: Your database name
- `DB_HOST`: Your database host, e.g., localhost
- `DB_PORT`: Your database port, e.g., 5432 for PostgreSQL
- `JWT_SECRET_KEY`: A secret key for signing JWTs
- `USER_ADMIN_EMAIL`: Email for the initial superadmin user
- `USER_ADMIN_PWD`: Password for the initial superadmin user

## Contribution :handshake:

Contributions are welcome! If you have suggestions or improvements, feel free to open an issue or submit a pull request.

## License :page_facing_up:

This project is licensed under the MIT License. See the `LICENSE` file for more details.

## Contact :mailbox_with_mail:

For any questions or comments, please contact us at [Jbagostin@gmail.com](mailto:jbagostin@gmail.com).
