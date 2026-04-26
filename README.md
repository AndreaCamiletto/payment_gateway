# REST Service in GO

Il progetto realizza una semplice API REST in Go che simula un sistema di pagamenti in-memory.


---

## Tecnologie utilizzate

- Go
- GoLand IDE
- Gin Web Framework
- UUID
- Map in-memory come storage
- sync.RWMutex per gestione concorrenza

---

## Funzionalità

- Creazione pagamento
- Recupero pagamento per ID
- Aggiornamento stato pagamento
- Stati supportati: PENDING, SUCCESS
- Storage in-memory thread-safe

---

## Architettura

- **Handler**
    - Gestione delle richieste HTTP
    - Parsing input (JSON / path params)
    - Simile ai Controller in Spring

- **Service**
    - Logica di business
    - Gestione stato pagamento
    - Validazioni

- **Storage**
    - Map in-memory
    - Protezione con mutex per concorrenza

## API e esempi CURL


### Crea pagamento

POST /payments 
- curl -X POST http://localhost:8080/payments -H "Content-Type: application/json" -d "{\"amount\":1000,\"currency\":\"EUR\"}"

### Recupera pagamento

GET /payments/{id}
- curl http://localhost:8080/payments/{id}
 
### Aggiorna pagamento

PUT /payments/{id}

- curl -X PUT http://localhost:8080/payments/{id}



