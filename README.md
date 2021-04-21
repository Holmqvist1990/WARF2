# WARF (WIP)

A base building game, inspired by Roguelikes and the original Dungeon Keeper series.

Built with [Golang](https://golang.org/) and the [Ebiten](https://ebiten.org/) framework.

![image](./readme_screenshot.png)

## What's with all the comments?

The linter I used when first writing this was _very adament_ in enforcing a specific style
of commenting public functions, leading to great declarative comments like:
```
// Returns10 returns 10.
func Returns10() int { return 10 }
```

I have since switched to another linter ([golangci-lint](https://golangci-lint.run/)) just to avoid this.

## Implemented

✅ Basic graphics (world, dwarves, items).

✅ Cellular automata, flood fills, pathfinding.

✅ Mouse and keyboard input.

✅ Collision detection.

✅ Saving/loading

✅ Wall system

✅ Activity system.

## WIP

🔹 Library system.

🔹 UI system.

## TODO

❌ UI - Graphics, components, menus.

❌ Activities - Eating, resting, reading.

❌ Sound system.