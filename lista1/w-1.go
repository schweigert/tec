// https://turingmachinesimulator.com/

//LOAD AN EXAMPLE TO TRY
//then load an input and click play

//Syntax:

//-------CONFIGURATION
name: w-1
init: q0
accept: qacp

//-------DELTA FUNCTION:
// [current_state],[read_symbol]
// [new_state],[write_symbol],[>|<|-]

q0, 0
q1, 0, <

q0, 1
q1, 1, <

q1, _
q2, #, >

q2, 1
q3, #, <

q2, 0
q4, #, <

q2, #
q2, #, >

q2, _
qf, _, <

q3, #
q3, #, <

q3, 0
q3, 0, <

q3, 1
q3, 1, <

q3, _
q5, 1, >

q4, #
q4, #, <

q4, 0
q4, 0, <

q4, 1
q4, 1, <

q4, _
q5, 0, >

q5, 1
q5, 1, >

q5, 0
q5, 0, >

q5, #
q2, #, >

qf, #
qf, _, <

qf, 1
qacp, 1, >

qf, 0
qacp, 0, >

// < = left
// > = right
// - = hold
// use underscore for blank cells

//States and symbols are case-sensitive

//Load your code and click COMPILE.
//or load an example (top-right).

