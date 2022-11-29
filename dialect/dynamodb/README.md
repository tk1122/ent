# Ent framework usage in general
```
[Framework users code]

calls Open -> init *Client -> {
	dynamodb.NewClient()

	init UserClient (+ CarClient,...)
}

CreateUsers(client) -> {
	Client.User.Create().Save()

	return User (Entity)
}

--------------------------
[Generated code]

Client.User.Create() -> returns UserCreate

[Optional]
UserCreate.setID(u) -> { uc.mutation.SetID(u) }
[Optional]

UserCreate.Save() -> {
	UserCreate.dynomodbSave()
}

UserCreate.dynomodbSave() -> {
	userEntity, spec = UserCreate.createSpec()
	dynamodbgraph.CreateNode(spec)

	return userEntity
}

UserCreate.createSpec() -> returns User(Entity) + dynamodbgraph.CreateSpec

-------------------------
[Database dialect]

dynamodbgraph.CreateNode() -> {
	This one is already in database dialect (adapter) area -> Done!
}
```

# Packages

### dynamodb/dynamodbgraph
- Graph abstraction of entities and relationships


### dynamodb/schema
- Schema creation with schemas defined in ent/schema package

### dynamodb
- Driver to interact with AWS DynamoDB Golang SDK

# Supported features
- Schema
  - Fields with different types:
    - Go primitive types
      - [x] Map
      - [x] Struct
      - [x] Array
    - Edges: 
      - [x] O2O 
      - [x] O2M
      - [x] M2O
      - [x] M2M
- [x] Mixin
- [x] Indexes
- [x] Annotations
- Code generation
  - CRUD API
    - [x] Get many
    - [x] Get one
    - [x] Create
    - [x] Batch create
    - [x] Update one
    - [x] Update many
    - [x] Delete one
    - [x] Delete many
- [x] Graph Traversal
- [x] Predicates
- [x] Hooks
# Unsupported features
- [ ] Ordering
- [ ] Aggregation
- [ ] Query chaining 
- [ ] Edge related predicates
- [ ] Limit-Offset style pagination
- [ ] User level transaction APIs
- [ ] Schema migration
- [ ] Cascading deletion








