# How to Run
```
# At root project
docker-compose build --no-cache
docker-compose up
```

# Frontend
```
http://localhost:5173
```

# Backend
```
http://localhost:8080/api/swagger/index.html
```

# Strategy
- Before generating voucher, double check existing voucher by flight number and flight date
- Unique composite index on column flight number and flight date, preventing data duplication at database level 
