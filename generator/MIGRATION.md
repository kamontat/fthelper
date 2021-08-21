# FTGenerator

## v4.x.x -> v5.0.0 (22 Aug 2021)

- Remove support `fs` key and use inside generator directly
  - Example on v4 we need to define fs object in `fs` fields and link in generator via object name
  - But now you can add directly on generator
  - This change is because most of fs data is not reuseable, so separating is useless
- Remove support `fs.variables` and use to `variables` instead
