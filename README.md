# DMN decision execution on the Cosmos blockchain

## Phase 1 - DMN model executed as a query

### Prerequisites

- In this PoC, the [dsntk](https://github.com/dsntk) engine is used as a DMN runtime.
- The example is based on the work [DMN Decision Execution on the Ethereum Blockchain](https://link.springer.com/chapter/10.1007/978-3-319-91563-0_20) by Stephan Haarmann.
- [Ignite CLI](https://docs.ignite.com/) is used to create and run the Cosmos Blockchain.

### Install DSNTK

```shell
$ cargo install dsntk
$ dsntk
dsntk 0.0.3
DSNTK | Decision Toolkit
dsntk: missing subcommand
Try 'dsntk --help' for more information.
```

### Create decision table for SLA

The decision table for calculating SLA is presented below. The source is saved in file [sla.dtb](sla.dtb). 
This table is identical like the one presented in Haarmann's work.

```text
 ┌───────┐
 │  SLA  │
 ├───┬───┴─────────────┬───────────────╥─────┐
 │ U │ YearsAsCustomer │ NumberOfUnits ║ SLA │
 │   ├─────────────────┼───────────────╫─────┤
 │   │    [0..100]     │ [0..1000000]  ║ 1,2 │
 ╞═══╪═════════════════╪═══════════════╬═════╡
 │ 1 │       <2        │    <1000      ║  1  │
 ├───┼─────────────────┼───────────────╫─────┤
 │ 2 │       <2        │   >=1000      ║  2  │
 ├───┼─────────────────┼───────────────╫─────┤
 │ 3 │      >=2        │     <500      ║  1  │
 ├───┼─────────────────┼───────────────╫─────┤
 │ 4 │      >=2        │    >=500      ║  2  │
 └───┴─────────────────┴───────────────╨─────┘
```

To evaluate this decision table, run:

```shell
$ dsntk edt sla.input sla.dtb
2
```
The [sla.input](sla.input) file contains input data presented to decision table during evaluation.

To test this decision table, run:

```shell
$ dsntk tdt sla.test sla.dtb
test 1 ... ok
test 2 ... ok
test 3 ... ok
test 4 ... ok
test 5 ... ok
test 6 ... ok
test 7 ... ok
test 8 ... ok
test 9 ... ok
test 10 ... ok
test 11 ... ok

test result: ok. 11 passed; 0 failed.
```

## Phase 2 - ?
(tbd)

## Phase 3 - ?
(tbd)

## Phase 4 - ?
(tbd)

## Phase 5 - ?
(tbd)
