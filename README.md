# boschtail

## Backend

> IN PROGRESS
generated with [ent](https://entgo.io/)

## Data model

> IN PROGRESS
better write proper datamodel, e.g. in [mermaid](https://mermaid-js.github.io/mermaid/#/classDiagram)

```mermaid
classDiagram
    Mixer --o "1" MixerConfig : hasA
    MixerConfig --o "1..*" Ingredient : configuredWith
    Order --> Mixer : refersTo
    Order --> Recipe : refersTo
    Recipe --o "1..*" RecipeIngredient : hasMany
    RecipeIngredient --o "1" Ingredient

    Mixer : +int id
    MixerConfig : +int id
    Order : +int id
    Recipe : +int id
    Recipe : +int name
    RecipeIngredient: +int id
    RecipeIngredient: +int amount
    Ingredient: +int id
    Ingredient: +String name
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

