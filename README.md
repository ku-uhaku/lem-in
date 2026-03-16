# Lem-in

A Go implementation of the Lem-in project, where ants need to find their way from a start room to an end room through a network of tunnels as efficiently as possible.

## What it does

This program reads a map file describing rooms, tunnels, and ants, then calculates the optimal paths for the ants to travel from start to end, and simulates their movement turn by turn.

## How to run

1. Make sure you have Go installed (version 1.26.1 or later)
2. Clone or download this project
3. Run with: `go run main.go <map_file>`

For example:
```
go run main.go example.map
```

## Input format

The input file should contain:
- Number of ants (first line)
- Room definitions: `name x y`
- Links between rooms: `room1-room2`
- Special lines: `##start` and `##end` before their respective rooms

Example input:
```
3
##start
start 0 0
room1 1 1
room2 2 2
##end
end 3 3
start-room1
room1-room2
room2-end
```

## Output

The program first prints the input map, then the ant movements like:
```
L1-room1 L2-room1
L1-room2 L3-room1
L2-room2 L3-room2
L3-end
```

Each line represents one turn, showing which ants move to which rooms.

## Project structure

- `main.go`: Entry point
- `kuuLemin/`: Main logic
- `kuuParser/`: Input parsing and validation
- `kuuFindPath/`: Path finding algorithms
- `kuuSimulate/`: Movement simulation
- `kuuType/`: Data structures

Enjoy watching the ants march!</content>
<parameter name="filePath">c:\Users\kuuhaku\Documents\mylem\README.md