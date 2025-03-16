# Kotbot

Discord bot to manage inactive and unverified members of Warrior Cats RP server

## Commands

- `/get_inactive_humans` Gives a list of people with a `Human` role that have been inactive for a week.
- `/set_human_role` Sets the `Human` role.
- `/set_admin_role` Sets the `Admin` role that grants access to other commands. This command can only be used by the server owner.
- `/set_unverified_role` Sets the `Unverified` role. People with this role will be automatically kicked after a week.

## Hosting

Kotbot is provided as a docker image `koobson/kotbot`.  
It requires the following envs to be set:

- `BOT_TOKEN` - Discord bot token
- `DB_PATH` - Directory path for sqlite database
- `LOG_PATH` - Directory path for bot logs
