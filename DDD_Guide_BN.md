# Go প্রজেক্টে DDD শুরু করার সম্পূর্ণ গাইড (বাংলা)

এই ফাইলটি তোমার বর্তমান প্রজেক্ট স্ট্রাকচার (`domain/`, `user/`, `product/`, `repo/`, `rest/handlers/`) ধরে তৈরি করা হয়েছে, যাতে তুমি **নতুন ডোমেইনের API** বানাতে step-by-step শুরু করতে পারো।

---

## 1) DDD আসলে কী?

**DDD (Domain-Driven Design)** হলো সফটওয়্যার ডিজাইন করার একটি পদ্ধতি যেখানে মূল ফোকাস থাকে ব্যবসার নিয়ম (business rules) এবং ডোমেইন জ্ঞান (domain knowledge)-এর উপর।

সহজভাবে:

1. আগে business problem বোঝো
2. domain model বানাও (Entity, Value Object, Rule)
3. তারপর API, DB, framework এগুলো domain-এর চারপাশে বসাও

অর্থাৎ, DDD-তে **database-first না**, বরং **domain-first**।

---

## 2) তোমার প্রজেক্টে DDD mapping (কোন ফোল্ডার কী)

তোমার বর্তমান কোডে DDD/clean layering এর কাছাকাছি mapping:

| Layer | এই প্রজেক্টে কোথায় | কাজ |
|---|---|---|
| Domain Model | `domain/` | Core entity (`User`, `Product`) |
| Domain/Application Service | `user/service.go`, `product/service.go` | Business use-case orchestration |
| Port (Interface) | `user/port.go`, `product/port.go`, `rest/handlers/*/port.go` | Contract define করে (dependency inversion) |
| Infrastructure (Repository impl) | `repo/` | DB query/ persistence |
| Interface/Delivery | `rest/handlers/` | HTTP request/response handle |
| Composition Root | `cmd/serve.go` | সব dependency wire-up করে |
| Technical Infra | `infra/db`, `config`, `migrations` | DB connection, migration, config |

---

## 3) DDD এর core building blocks

### Entity
যে object-এর identity থাকে (যেমন `User.ID`, `Product.ID`)।

### Value Object
Identity ছাড়াই immutable concept (যেমন Email, Money)।  
তোমার প্রজেক্টে এখনো আলাদা VO নেই, পরে যোগ করতে পারো।

### Repository
Domain object persist/retrieve করার abstraction (interface + implementation)।

### Domain Service
যে business logic একাধিক entity/use-case coordinate করে।

### Application Service (Use Case)
Client/API request কে domain operation-এ translate করে; transaction/orchestration করে।

### Bounded Context
Domain boundary। যেমন `User` context, `Product` context, ভবিষ্যতে `Order` context।

---

## 4) নতুন প্রজেক্টে DDD দিয়ে শুরু করার সঠিক sequence

নতুন Go API project শুরু করলে এই order follow করো:

1. **Domain বোঝা ও boundary ঠিক করা**  
   - কাদের নিয়ে কাজ? (User, Product, Order, Payment)
   - কোন context কোথায় শেষ?

2. **Ubiquitous Language লিখে ফেলা**  
   - business term-এর dictionary বানাও (যেমন “Order Placed”, “Paid”, “Cancelled”)  
   - team-এর সবাই একই শব্দ ব্যবহার করবে।

3. **Domain model (Entity/VO) তৈরি করা**  
   - fields + invariants (rule) define করো।

4. **Use case list করা**  
   - যেমন: CreateOrder, GetOrder, CancelOrder, ListOrders।

5. **Port interface define করা**  
   - service interface, repository interface আগে define করো।

6. **Service implementation লিখো**  
   - business logic domain-centric রাখো।

7. **Repository implementation (DB)**  
   - SQL/query/ORM code infra layer-এ রাখো।

8. **HTTP Handler + Route**  
   - request validation, response mapping, error mapping করো।

9. **Dependency wiring (composition root)**  
   - `cmd/serve.go`-এর মতো জায়গায় repo -> service -> handler inject করো।

10. **Migration + run + test**  
   - schema migration, endpoint test, edge-case test।

---

## 5) তোমার current project-এ নতুন domain API যোগ করার exact sequence

ধরা যাক তুমি নতুন domain **`order`** যোগ করবে।

### Step A: Domain entity
`domain/order.go`

```go
package domain

type Order struct {
	ID       int     `json:"id" db:"id"`
	UserID   int     `json:"user_id" db:"user_id"`
	Total    float64 `json:"total" db:"total"`
	Status   string  `json:"status" db:"status"`
}
```

### Step B: Domain/Application service contract
`order/port.go`

```go
package order

import "ecommerce/domain"

type Service interface {
	Create(order domain.Order) (*domain.Order, error)
	Get(id int) (*domain.Order, error)
	List() ([]*domain.Order, error)
}

type OrderRepo interface {
	Create(order domain.Order) (*domain.Order, error)
	Get(id int) (*domain.Order, error)
	List() ([]*domain.Order, error)
}
```

### Step C: Service implementation
`order/service.go`  
এখানে business rule রাখবে (status transition, amount check, ownership check ইত্যাদি)।

### Step D: Repository implementation
`repo/order.go`  
এখানে SQL query থাকবে। `repo/product.go` / `repo/user.go` pattern follow করো।

### Step E: HTTP handler
`rest/handlers/order/`

1. `handler.go` (dependency receive)
2. `port.go` (service interface contract)
3. `create_order.go`, `get_order.go`, `get_orders.go`, ইত্যাদি
4. `route.go` (`mux.Handle(...)`)

### Step F: Route registration
`rest/server.go`-এ নতুন handler inject + `RegisterRoutes(...)` call।

### Step G: Dependency wiring
`cmd/serve.go`-তে:

1. `orderRepo := repo.NewOrderRepo(dbCon)`
2. `orderService := order.NewService(orderRepo)`
3. `orderHandler := orderHandler.NewHandler(..., orderService)`
4. `rest.NewServer(...)`-এ pass করো।

### Step H: Migration
`migrations/000003-create-orders.up.sql` এবং `.down.sql` যোগ করো।

---

## 6) Request flow (end-to-end)

নতুন endpoint hit হলে flow হবে:

**HTTP Request -> Handler -> Service -> Repository -> DB -> Repository -> Service -> Handler -> HTTP Response**

এই flow clean থাকলে:

- handler-এ business logic ঢুকবে না
- service-এ SQL ঢুকবে না
- repo-তে HTTP concern ঢুকবে না

এটাই DDD/Clean separation-এর বড় লাভ।

---

## 7) কোথা থেকে start করবে? (প্র্যাকটিক্যাল answer)

তুমি যদি আজ নতুন domain API শুরু করো, এই 3 টা কাজ আগে করো:

1. `domain/<new_domain>.go` লিখো (entity + rule fields)
2. `<new_domain>/port.go` + `<new_domain>/service.go` লিখো
3. `repo/<new_domain>.go` বানিয়ে DB integration করো

তারপর handler + route + serve wiring করো।

---

## 8) Common mistakes (এড়িয়ে চলবে)

1. Handler-এ business logic লেখা
2. Service-এ raw HTTP request object পাঠানো
3. Repo-তে domain rule validate করা
4. এক domain-এর logic অন্য domain-এ copy-paste করা
5. DB schema দেখে domain model force করা (উল্টোটা হওয়া উচিত)

---

## 9) Minimal checklist (প্রতি নতুন domain-এর জন্য)

- [ ] Domain entity তৈরি
- [ ] Service + Repo interface define
- [ ] Service logic implement
- [ ] Repo SQL implement
- [ ] Handler endpoints implement
- [ ] Route register
- [ ] `cmd/serve.go`-তে wiring
- [ ] Migration add
- [ ] API test (happy path + error path)

---

## 10) Final note

তোমার প্রজেক্টে DDD fully strict tactical style না হলেও, তুমি ইতিমধ্যে **domain + service + repo + handler separation** follow করছো।  
এখন consistency ধরে রেখে প্রতিটি নতুন domain একই sequence-এ build করলে architecture clean থাকবে, maintenance সহজ হবে, এবং নতুন feature যোগ করা অনেক দ্রুত হবে।

