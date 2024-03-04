# SparkSentry API :zap:

Welcome to the SparkSentry API repository, an innovative energy management system designed to optimize your energy consumption.

## Features :star:

SparkSentry offers a robust platform for energy data collection and analysis, enabling users to monitor and optimize their energy consumption efficiently. It now supports flexible user-account associations and account creation, catering to a wide range of energy management needs.

## Setup :gear:

To set up and run the project locally, follow these steps:

1. Clone the repository to your local machine.
2. Ensure you have Go installed on your machine.
3. Install dependencies by running `go mod tidy`.
4. Configure the necessary environment variables in a `.env` file based on the `.env.example` file.
5. Launch the application with `go run main.go`.

## API Routes :world_map:

The SparkSentry API exposes the following routes:

### Authentication

- `POST /api/v1/login`: Log in to the application and receive a JWT token. :key:

### User Management

- `POST /api/v1/register`: Register a new user into the system. :bust_in_silhouette:
- `POST /api/v1/associateUserToAccount`: Associate an existing user to an account. :link:

### Account Management

- `POST /api/v1/accounts`: Create a new account. :office:

### Protected Routes (JWT token required)

- `GET /api/v1/securedata`: Access secure data after authentication. :lock:

## Contribution :handshake:

Contributions are welcome! If you have suggestions or improvements, feel free to open an issue or submit a pull request.

## License :page_facing_up:

This project is licensed under the MIT License. See the `LICENSE` file for more details.

## Contact :mailbox_with_mail:

For any questions or comments, please contact us at [Jbagostin@gmail.com](mailto:jbagostin@gmail.com).
