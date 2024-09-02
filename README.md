# Alignon

> [!NOTE]
> This project is nothing more than a showcase and is deprecated.
> 
> Also, lacking of OSS license is intentional.

`Alignon = anagram("Loaning")`

[Database diagram](https://dbdiagram.io/d/My-Alignon-Database-Schema-66c48707a346f9518c90ca7c)

## Endpoints

### Consumer

- [x] Middleware: Authenticate user
- [x] Apply loan
- [x] Get loan history
- [x] Get loan limit
- [x] Get current loan detail
- [x] Register user + file uploads

### Admin

- [ ] Middleware: Authenticate admin
- [ ] Approve loan application
- [ ] Reject loan application
- [ ] Get consumer detail
- [ ] Get consumers
- [ ] Get loan detail
- [ ] Get loans
- [ ] Register loan payment

## Stack

- Web framework: [Fiber](https://gofiber.io/)
- Query builder: [Jet](https://github.com/go-jet/jet)
- RDBMS: PostgreSQL