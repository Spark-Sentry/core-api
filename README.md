# SparkSentry API :zap:

Welcome to the SparkSentry API repository, an innovative energy management system designed to optimize your energy consumption.

## Features :star:

SparkSentry offers a robust platform for energy data collection and analysis, enabling users to monitor and optimize their energy consumption efficiently. It now supports flexible user-account associations, account creation, and detailed management of buildings, systems, and equipment, catering to a wide range of energy management needs.

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

### System Management (JWT token required)
- `POST /api/v1/buildings/:building_id/systems`: Add a new system to a specific building. :gear:
- `GET /api/v1/buildings/:building_id/systems`: Retrieve all systems associated with a specific building's ID. :wrench:

## Environment Variables :key:

Ensure the following environment variables are set in your `.env` file:

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
