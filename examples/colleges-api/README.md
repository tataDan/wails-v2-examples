# README

## About

This application uses the collegescorecard api.

Documentation can be found at [this web site](https://collegescorecard.ed.gov/data/documentation/).

This application can be used to help find colleges using the following criteria:

    city,

    state,

    program (field of study) (ex. "Chemistry),

    undergraduate student size maximum limit,

    in-state tuition maximum limit,

    or any combination of the above.


## Building

To build this application you must first obtain an api key from [this web site](https://api.data.gov/signup/).
Then you will need to modify the line in app.go that reads:

`apiKey := "&api_key=XXXXXXXXXXXXXXXXXXXXXX"` by replacing all of the Xs with your own api key.

If the `wails dev` command does not work, please try `wails dev --forcebuild`.

To build the application, run `wails build`.


## Distributing

The `name-abbr.csv` file and the `program-code.csv` file must both be in the same directory as the executable file.

In some circumstances it might be useful to create a batch file that changes directories into that directory and then executes the application.

