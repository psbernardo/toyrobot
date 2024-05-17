# toyrobot

# Toy Robot Challenge

Welcome to the Toy Robot Challenge! Please complete this coding exercise and send through your solution for review. We are looking to understand how you solve problems, what approaches you believe are good practices and what you may decide to trade off on and why. Think of the code like a communication exercise where you are communicating the solution clearly to a group of peers, not just fulfilling the functional requirements of the task.

Please code your solution to the level as if it was “production ready”. You will be given detailed feedback on your solution and if you’re successful and proceed to the next stage, your code will be discussed and extended upon in the interview.

Most candidates require a few hours spread over a week to complete the exercise. If you need more time, please get in touch with us to discuss. If you’re short on time, it’s better to prioritise non-functional elements over functional requirements. If you take this approach, please document it so that the reviewer has this context when they review your code.

You can use any programming language that you'd like. It's probably best to use the language you're most familiar with so that you can show your best code.

The review team will be looking for things like:

- Striking a good balance of simiplicity vs complexity
- Appropriate use of design patterns
- High quality, reliable code that is easy to extend
- A great user experience

## Problem Description

- The application is a simulation of a toy robot moving on a square tabletop, of dimensions 5 units x 5 units
- There are no other obstructions on the table surface
- The robot is free to roam around the surface of the table, but must be prevented from falling to destruction
- Any movement that would result in the robot falling from the table must be prevented, however further valid movement commands must still be allowed

Create an application that can read in commands of the following form -

    PLACE X,Y,F
    MOVE
    LEFT
    RIGHT
    REPORT

PLACE will put the toy robot on the table in position X,Y and facing NORTH, SOUTH, EAST or WEST.

The origin (0,0) can be considered to be the SOUTH WEST most corner.

The first valid command to the robot is a PLACE command, after that, any sequence of commands may be issued, in any order, including another PLACE command. The application should discard all commands in the sequence until a valid PLACE command has been executed.

MOVE will move the toy robot one unit forward in the direction it is currently facing.

LEFT and RIGHT will rotate the robot 90 degrees in the specified direction without changing the position of the robot.

REPORT will announce the X,Y and F of the robot. This can be in any form, but standard output is sufficient.

A robot that is not on the table can choose the ignore the MOVE, LEFT, RIGHT and REPORT commands.

Input can be from a file, or from standard input, as the developer chooses.

Provide test data to exercise the application.

## Constraints

The toy robot must not fall off the table during movement. This also includes the initial placement of the toy robot. Any move that would cause the robot to fall must be ignored.

## Example Input/Output

    a) PLACE 0,0,NORTH
       MOVE
       REPORT
       Output: 0,1,NORTH

    b) PLACE 0,0,NORTH
       LEFT
       REPORT
       Output: 0,0,WEST

    c) PLACE 1,2,EAST
       MOVE
       MOVE
       LEFT
       MOVE
       REPORT
       Output: 3,3,NORTH
