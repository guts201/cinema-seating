

**Cinema Project Documentation**

**Overview**

**Architecture Overview**
The architecture is based on a microservices pattern using gRPC for communication between services. Here is an overview of the key components:
+ Cinema Service (gRPC Server): Handles seat reservation and cancellation requests.
+ MySQL Database: Stores cinema, screenings, seat reservations, and other entities.
+ Redis: Caches seat availability for faster access.

**Features**
Multiple cinemas with different row/column layouts and distancing rules
One screening per (cinema, movie) pair

* **Users can:**

+ Find available seat groups (respecting distancing rules)
+ Reserve specific seats for a screening
+ Cancel reservations
+ ...

* **Admin APIs for:**

+ Creating cinemas
+ Creating screenings
+ ...


**Assumptions**
The system supports multiple cinemas, each with:
+ A unique layout of rows, columns
+ A custom min_distance for enforcing social distancing
+ Each (cinema, movie) pair has one screening
+ A screening is a showtime of a movie at a cinema

+ Users book seats directly on a screening, not on a cinema or movie. It's considered reserved immediately (no payment process for now)
+ Users can cancel reservations without restriction



### Cinema Service

**Project Structure**

internal/
&nbsp;&nbsp;mapper/             # Converts DB entities <-> API
&nbsp;&nbsp;entity/             # Domain models
&nbsp;&nbsp;schema/             # Entgo schema definitions
&nbsp;&nbsp;usecase/            # Business logic
&nbsp;&nbsp;repository/         # Interfaces + implementations (DB + Redis)
&nbsp;&nbsp;server/
&nbsp;&nbsp;&nbsp;&nbsp;clientcinema/           # gRPC server logic for client
&nbsp;&nbsp;&nbsp;&nbsp;backofficecinema/           # gRPC server logic for admin
&nbsp;&nbsp;&nbsp;&nbsp;serve.go
&nbsp;&nbsp;tx/                 # Transaction helpers
api/
&nbsp;&nbsp;cinema.proto        # gRPC service definition
&nbsp;&nbsp;cinema_config.proto
pkg/
&nbsp;&nbsp;genkit/             # Shared generator helpers
&nbsp;&nbsp;carbon/             # Config, secrets
&nbsp;&nbsp;grpc/               # gRPC connection setup
&nbsp;&nbsp;logging/            # Logger setup
cmd/
&nbsp;&nbsp;main.go             # Entry point


## 📁 Key Components

###  `api/`
- `cinema.proto`: gRPC services for user-side seat operations
- `cinema_config.proto`: Admin APIs (create cinema, screening, etc.)
- `cinema_code.proto`: Response codes / enums

###  `entity/`
- Contains domain logic and clean models:
  - `Cinema`, `Screening`, `Seat`, etc

###  `schema/` (Entgo)
- Database schema definitions:
  - `cinema.go`, `screening.go`, `seat.go`, `seat_reservation.go`

###  `repository/`
- Contains repository interfaces and DB/Redis implementations:
  - `seatRepository`, `screeningRepository`, `cinemaRepository`, etc.

###  `mapper/`
- Converts between:
  - DB models ↔ domain entities  
  - domain entities ↔ gRPC models

###  `usecase/`
- Business logic layer:
  - `ReserveSeats`, `CancelSeats`, `GetAvailableSeats`, etc.
- Handles validation, domain rules, and coordinates DB access

###  `server/`
- gRPC server handlers for user and admin APIs

###  `tx/`
- Helpers to wrap use cases in database transactions

###  `cmd/`
- `main.go`: Entry point to start the gRPC server

###  `pkg/`
- Shared utility packages:
  - `genkit/`: Generator helpers
  - `carbon/`: Config & secrets
  - `grpc/`: gRPC client setup
  - `logging/`: Logger setup

## 🚀 How to Run
### 1. Run the gRPC server

```bash
go run cmd/main.go -c configs/config.yaml
```

Configuration

+ DATABASE: 
+ REDIS_ADDR: Redis address (e.g. localhost:6379)