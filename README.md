# Project 0: Resturant Selection Tool

The Resturant Selection Tool is a CLI program that help a group of people select a place to eat based on selection criteria.

## CLI Usage

./food [--type] [--price] [--distance]

--tpye : This option takes a string as an argument for the type of food you are interested in. (i.g. "mexican" "pizz")
--price : This option takes a float that represents the most you are willing to spend on a meal.
--distance : This option takes a float and represents the max distance you are willing to travel. (defaults to 28 miles)

## User Stories

- [x] User can get a list for potential eating locations
- [x] User can get a filtered list of locations based on price, food genre, and distance from Liv+ apartments based on CLI flags
- [X] User can get a list that is up to date with current Arlington restaurants.
- [X] User can see if a restaurant from list output is currently open at time of running program.

### To-Do List

- [x] Add Basic Functionality to Program
- [x] Add CLI flags to program arguments
- [x] Add Data Persistance via JSON file for restaurant list
- [x] Refactor Code into more standard format
- [x] Add Proper Error Handling

### Code Debt Refactoring Periodic Routine Branch

1) Bug Fixes
2) Seperate Logic into Functions
3) Remove Hardcoding
4) Avoid Casting Values
5) Promote Modularity with Packages
6) Documentation
7) Test/Benchmark/Example Additions

## Engineering Design Process

<https://www.sciencebuddies.org/Files/5082/8/2013-updated_engineering-method-steps_v6b_noheader.png>
