# ORUM - Transfer Management System

This is a simple Go application for managing transfers and accounts. It provides RESTful APIs for creating accounts and retrieving transfer information. The application uses SQLite as its database to store transfer and account data.

## Prerequisites

Before running the application, ensure you have the following installed:

- Go (Golang) - You can download it from [golang.org](https://golang.org/dl/).
- SQLite - A lightweight, embedded database. No installation is needed as SQLite is included in the standard Go library.
- Third-party packages: You can install these packages using Go modules.

```bash
go get github.com/google/uuid
go get github.com/gorilla/mux
go get modernc.org/sqlite
```

## Setup

1. Clone this repository to your local machine.

   ```bash
   git clone git@github.com:90cigars/code-interview.git

2. Navigate to the project directory: `code-interview/go`
3. Run the application `go run main.go`
   1. The application will start and listen on port 8080 by default.


## Endpoints

1. **Get Transfers**
    - **Endpoint:** `/transfers`
    - **Method:** `GET`
    - **Description:** Retrieve a list of all transfers, including source and destination customer information.
    - **Response:** A JSON array of transfer objects.

2. **Get Transfer by ID**
    - **Endpoint:** `/transfers/{id}`
    - **Method:** `GET`
    - **Description:** Retrieve a specific transfer by its unique ID.
    - **Response:** A JSON object representing the transfer.

3. **Create Account**
    - **Endpoint:** `/accounts`
    - **Method:** `POST`
    - **Description:** Create a new account with the provided customer information.
    - **Request Body:** A JSON object containing customer details, including account holder name, account number, and routing number.
    - **Response:** A JSON object containing the newly created account details, including the generated account ID.

## Database

The application uses SQLite as its database. The database file (`transfers.db`) is expected to be located in the project's parent directory. You can modify the database connection settings in the main function if needed.

## Routing Number Validation

The application includes a simple routing number validation function to ensure the validity of routing numbers before creating an account. It follows the algorithm for validating routing numbers.

## Dependencies

- `github.com/google/uuid`: Package for generating UUIDs.
- `github.com/gorilla/mux`: A powerful HTTP router and URL matcher for building web applications.
- `modernc.org/sqlite`: SQLite database driver for Go (Golang).

## Error Handling

The application includes basic error handling for database operations and API requests. Detailed error messages are provided in responses to assist with debugging.


## Testing

The project includes test cases to ensure the functionality of the application. These tests are written in Go and utilize the testing framework provided by the Go standard library. Below are the key test cases included in the project:

### TestGetTransfers

- **Description:** This test case checks the behavior of the `/transfers` endpoint for retrieving a list of all transfers.
- **Test Steps:**
    1. Create a new router.
    2. Register the handler function for the `/transfers` endpoint.
    3. Create a new GET request for the `/transfers` endpoint.
    4. Create a response recorder.
    5. Serve the request using the router.
    6. Check the HTTP status code of the response (Expected: 200).
- **Additional Assertions:** You can add more assertions to validate the response as needed.
- **Usage:** This test verifies that the `/transfers` endpoint returns the expected response.

### TestGetTransfer

- **Description:** This test case checks the behavior of the `/transfers/{id}` endpoint for retrieving a specific transfer by its ID.
- **Test Steps:**
    1. Create a new router.
    2. Register the handler function for the `/transfers/{id}` endpoint.
    3. Create a new GET request for the `/transfers/{id}` endpoint (replace `TRANSFER_ID` with an actual transfer ID).
    4. Create a response recorder.
    5. Serve the request using the router.
    6. Check the HTTP status code of the response (Expected: 200).
- **Usage:** This test verifies that the `/transfers/{id}` endpoint returns the expected response for a specific transfer.

### TestCreateAccount

- **Description:** This test case checks the behavior of the `/accounts` endpoint for creating a new account.
- **Test Steps:**
    1. Create a new router.
    2. Register the handler function for the `/accounts` endpoint.
    3. Create an `AccountRequest` object with the required data.
    4. Marshal the `AccountRequest` object into JSON and create a POST request.
    5. Set the `Content-Type` header to `application/json`.
    6. Create a response recorder.
    7. Serve the request using the router.
    8. Check the HTTP status code of the response (Expected: 200).
- **Usage:** This test verifies that the `/accounts` endpoint correctly processes account creation requests.

### Database Setup

- **Description:** The testing environment includes an in-memory SQLite database (`file::memory:?cache=shared`) to isolate the tests from the production database. The database setup is performed in the `setupTestDB` function.
- **Usage:** The in-memory database ensures that tests can be run without affecting the production database.

To run the tests, use the `go test` command:

```bash
go test
```

## Potential Improvements

While the current version of the project meets its core requirements, there are several potential improvements and enhancements that can be considered for future development:

1. **Authentication and Authorization:** Implement user authentication and authorization mechanisms to secure the API endpoints, ensuring that only authorized users can perform certain actions.

2. **Data Validation:** Enhance data validation by implementing stricter validation rules and error handling for incoming requests to prevent invalid data from being processed.

3. **Logging:** Implement a robust logging system to record important events, errors, and transactions within the application. Logging can greatly aid in debugging and monitoring.

4. **Pagination:** Add pagination support to endpoints that return lists of data, such as transfers, to improve performance and usability when dealing with a large number of records.

5. **Unit Testing:** Expand the test suite to cover more edge cases and scenarios. Consider using testing frameworks like `testify` to simplify assertions and enhance test coverage.

6. **Database Migrations:** Implement a database migration tool to manage changes to the database schema as the application evolves. This ensures smooth updates without data loss.

7. **API Versioning:** Consider adding versioning to the API to maintain backward compatibility while making changes to the endpoints or response structures.

8. **Input Validation Middleware:** Implement middleware functions to handle input validation, reducing code duplication and enhancing maintainability.

9. **Error Handling Middleware:** Create a centralized error-handling middleware to handle errors uniformly across all endpoints and provide meaningful error responses.

10. **Containerization:** Dockerize the application to facilitate deployment and ensure consistency in different environments. Container orchestration with tools like Kubernetes can also be explored.

11. **Caching:** Introduce caching mechanisms for frequently requested data to improve response times and reduce the load on the database.

12. **Documentation:** Enhance the project's documentation by generating API documentation using tools like Swagger or documenting endpoints and functions thoroughly.

13. **Security Auditing:** Conduct security audits and vulnerability assessments regularly to identify and address potential security risks.

14. **Performance Optimization:** Profile and optimize critical parts of the codebase to improve application performance, especially when handling a high volume of requests.

15. **Monitoring and Metrics:** Implement monitoring and metric collection to gain insights into the application's health and performance in production.

16. **Load Testing:** Conduct load testing to ensure that the application can handle the expected traffic and identify bottlenecks or areas that need optimization.

17. **Feedback and User Experience:** Gather feedback from users and stakeholders to continuously improve the user experience and prioritize feature requests.


## Authors

- Sagar Khasnis

## License

This project is licensed under the MIT License - see the LICENSE file for details.
