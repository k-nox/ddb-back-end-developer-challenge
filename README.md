# DDB Back End Developer Challenge

### Overview
This task focuses on creating an API for managing a player character's Hit Points (HP) within our game. The API will enable clients to perform various operations related to HP, including dealing damage of different types, considering character resistances and immunities, healing, and adding temporary Hit Points. The task requires building a service that interacts with HP data provided in the `briv.json` file and persists throughout the application's lifetime.

### Task Requirements

#### API Operations
1. **Deal Damage**
    - Implement the ability for clients to deal damage of different types (e.g., bludgeoning, fire) to a player character.
    - Ensure that the API calculates damage while considering character resistances and immunities.

    > Suppose a player character is hit by an attack that deals Piercing damage, and the attacker rolls a 14 on the damage's Hit Die (with a Piercing damage type). `[Character Hit Points - damage: 25 - 14 = 11]`

2. **Heal**
    - Enable clients to heal a player character, increasing their HP.

3. **Add Temporary Hit Points**
    - Implement the functionality to add temporary Hit Points to a player character.
    - Ensure that temporary Hit Points follow the rules: they are not additive, always taking the higher value, and cannot be healed.

    > Imagine a player character named "Eldric" currently has 11 Hit Points (HP) and no temporary Hit Points. He finds a magical item that grants him an additional 10 HP during the next fight. When the attacker rolls a 19, Eldric will lose all 10 temporary Hit Points and 9 from his player HP.

#### Implementation Details
- Build the API using your preferred technology stack.
- Ensure that character information, including HP, is initialized during the start of the application. Developers do not need to calculate HP; it is provided in the `briv.json` file.
- Retrieve character information, including HP, from the `briv.json` file.


#### Data Storage
- You have the flexibility to choose the data storage method for character information.

### Instructions to Run Locally
1. Clone the repository or obtain the project files.
2. Install any required dependencies using your preferred package manager.
3. Configure the API with necessary settings (e.g., database connection if applicable).
4. Build and run the API service locally.
5. Utilize the provided `briv.json` file as a sample character data, including HP, for testing the API.

### Additional Notes
- Temporary Hit Points take precedence over the regular HP pool and cannot be healed.
- Characters with resistance take half damage, while characters with immunity take no damage from a damage type.
- Use character filename as identifier

#### Possible Damage Types in D&D
Here is a list of possible damage types that can occur in Dungeons & Dragons (D&D). These damage types should be considered when dealing damage or implementing character resistances and immunities:
- Bludgeoning
- Piercing
- Slashing
- Fire
- Cold
- Acid
- Thunder
- Lightning
- Poison
- Radiant
- Necrotic
- Psychic
- Force

If you have any questions or require clarification, please reach out to your Wizards of the Coast contact, and we will provide prompt assistance.

Good luck with the implementation!

## Implementation Notes

### To run the server
1. Go 1.21.1 or greater must be installed
   - on macOS: `brew install go`
2. SQLite3 must be installed
   - on macOS: `brew install sqlite3`
3. `CGO_ENABLED=1` must be set in the environment
4. A gcc compiler must be installed
5. To run the server:
   - `make run`
   - this will install all dependencies, compile the binary, and start the server 
6. Go to http://localhost:8080/ to get the GraphQL playground - see below for example queries

### To run all tests
`make test` will install all test dependencies and run all tests

### Example Queries

The GraphQL playground should have autocomplete to help with discovery (try `<ctrl>+<space>`).
You can also view the GraphQL schema directly at `graph/schema/schema.graphqls`.

#### Get Character By Name
```graphql
query {
    characterByName(name: "Briv") {
       id,
       name,
       currentHitPoints,
       maxHitPoints,
       temporaryHitPoints,
       level,
       stats {
          charisma,
          constitution,
          wisdom,
          dexterity,
          intelligence,
          strength
       }
       defenses {
          damageType
          defenseType
       },
    }
}
```

returns by default:
```json
{
   "data": {
      "characterByName": {
         "id": "1",
         "name": "Briv",
         "currentHitPoints": 25,
         "maxHitPoints": 25,
         "temporaryHitPoints": null,
         "level": 5,
         "stats": {
            "charisma": 8,
            "constitution": 14,
            "wisdom": 10,
            "dexterity": 12,
            "intelligence": 13,
            "strength": 15
         },
         "defenses": [
            {
               "damageType": "FIRE",
               "defenseType": "IMMUNITY"
            },
            {
               "damageType": "SLASHING",
               "defenseType": "RESISTANCE"
            }
         ]
      }
   }
}
```

Note: All of the queries and mutations return a character object, and you can request any combination of the fields on that object.
For simplicity, the following examples will only include a subset of fields.

#### Get Character By ID
```graphql
query {
  character(id: 1) {
    name
  }
}
```

returns:
```json
{
  "data": {
    "character": {
      "name": "Briv"
    }
  }
}
```

#### Damage Character
```graphql
mutation {
   damageCharacter(input: {characterId: 1, damageType: BLUDGEONING, roll: 12}) {
      name
      currentHitPoints
      maxHitPoints
      temporaryHitPoints
   }
}
```

returns:
```json
{
  "data": {
    "damageCharacter": {
      "name": "Briv",
      "currentHitPoints": 13,
      "maxHitPoints": 25,
      "temporaryHitPoints": null
    }
  }
}
```

#### Heal Character
```graphql
mutation {
  healCharacter(input: {characterId: 1, roll: 5}) {
    name,
    currentHitPoints,
    maxHitPoints,
    temporaryHitPoints
  }
}
```

returns:
```json
{
  "data": {
    "healCharacter": {
      "name": "Briv",
      "currentHitPoints": 18,
      "maxHitPoints": 25,
      "temporaryHitPoints": null
    }
  }
}
```

#### Add Temporary Hit Points
```graphql
mutation {
  addTemporaryHitPoints(input: {characterId: 1, roll: 6}) {
    name
    currentHitPoints
    maxHitPoints
    temporaryHitPoints
  }
}
```

returns:
```json
{
  "data": {
    "addTemporaryHitPoints": {
      "name": "Briv",
      "currentHitPoints": 18,
      "maxHitPoints": 25,
      "temporaryHitPoints": 6
    }
  }
}
```

