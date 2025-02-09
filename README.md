# Golang CRUD API with Search Functionality

This is a sample Golang project implementing CRUD (Create, Read, Update, Delete) APIs along with a search API using the Gin framework. It uses MySQL as the database, GORM as the ORM, and various other libraries for logging, environment management, and data migration.

## Technologies Used

- **Gin** - HTTP web framework
- **MySQL** - Relational database
- **GORM** - ORM for Golang
- **godotenv** - Environment variable management
- **sirupsen/logrus** - Logging library
- **zerolog** - High-performance structured logging
- **reflections** - Reflection utilities for Golang
- **dbmate** - Database migration tool

## Models

### Car
Represents a car entity with attributes such as:
- ID (Primary Key)
- Make
- CarModel
- Year
- BodyType

### CarVariants
Represents different variants of a car with attributes such as:
- ID (Primary Key)
- CarID (Foreign Key referencing Car)
- VariantCode
- VariantName
- Transmission
- Color
- Engine

## Features
- CRUD operations for `Car` and `CarVariants`
- Search API to filter `Car` and `CarVariants` based on various attributes
- Structured logging with `zerolog` and `logrus`
- Environment variables managed with `godotenv`
- Database migrations using `dbmate`

## Installation

1. Clone the repository:
   ```sh
   git clone git@github.com:arvindnikam/go-cars.git
   cd go-cars
   ```

2. Install dependencies:
   ```sh
   go mod tidy
   ```

3. Configure the `.env` file with database credentials:
   ```env
   DB_USER=root
   DB_PASSWORD=yourpassword
   DB_NAME=yourdbname
   DB_HOST=localhost
   DB_PORT=3306
   DATABASE_URL="mysql://root:yourpassword@localhost:3306/yourdbname"
   ```

4. Run database migrations:
   ```sh
   dbmate up
   ```

5. Start the application:
   ```sh
   go run main.go
   ```

## API Endpoints

### Cars
- `POST /cars/create` - Create a new car
- `POST /cars/search` - Search cars based on conditions
- `GET /cars/:car_id` - Get car details
- `PUT /cars/:car_id/update` - Update car details
- `DELETE /cars/:id/delete` - Delete a car

### Car Variants
- `POST /cars/:car_id/car_variants/create` - Add a variant to a car
- `POST /cars/:car_id/car_variants/search` - Search car variants based on conditions
- `GET /cars/:car_id/car_variants/:car_variant_id` - Get car variant details
- `PUT /cars/:car_id/car_variants/:car_variant_id/update` - Update variant details
- `DELETE /cars/:car_id/car_variants/:car_variant_id/delete` - Delete a variant

### Sample Search API
   ```sh
   curl --location 'localhost:8880/api/v1/cars/search' \
   --header 'Content-Type: application/json' \
   --data '{
      "Conditions": {
         "Make": {
               "operator": "eq",
               "value": "Honda"
         },
         "CarModel": {
               "operator": "eq",
               "value": "Accord"
         }
      }
   }'
   ```

## Logging
This project uses `logrus` and `zerolog` for structured logging.

## Contributions
Feel free to fork the repository and submit pull requests!

## License
This project is licensed under the MIT License.
