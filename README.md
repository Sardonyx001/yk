# yk: yaru-koto/やること

## what?

a dead simple tui app to manage tasks.

## yet another todo app?

coudln't find any simple app that has hotkeys to streamline usage so i wrote one.

## feature list

### data models

- task lists (`tasklist`) contain many tasks
  - seperate tasks into different contexts (work, personal, etc)
- tasks (`task`)
  - contains title: `str` and description: `str` (usually title should fit within one line while description can be way longer)
  - by default hide most of the description and only show it when selecting its respective task
-
