# boschtail

## Backend

> IN PROGRESS
generated with [ent](https://entgo.io/)

## Data model

> IN PROGRESS
better write proper datamodel, e.g. in [mermaid](https://mermaid-js.github.io/mermaid/#/classDiagram)

```mermaid
classDiagram
classA --|> classB : Inheritance
classC --* classD : Composition
classE --o classF : Aggregation
classG --> classH : Association
classI -- classJ : Link(Solid)
classK ..> classL : Dependency
classM ..|> classN : Realization
classO .. classP : Link(Dashed)`
```

mixer
- id
- mixer_config_id

mixer_config
- list of ingredient (ordered)

order
- id
- recipe_id
- mixer_id (calculated)

recipe
- id
- list of recipe_ingredient

recipe_ingredient
- id
- ingredient_id
- amount

ingredient
- id
- name

