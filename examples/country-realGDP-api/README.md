# README

This example allows the user to select a country and see the real gross domestic product data for that country (if reported).  Optionally, the user can create a line chart that is stored as a PNG image file. This line chart uses the last ten values reported (typically quarters).  There is also an option to open the .png file in a photo viewer (or browser). For this option to work, the viewer must be in the user’s PATH or the full path must be specified. This option has not been tested on macOS (but in theory could work).

This example uses [this web site](https://restcountries.com/) to get the country codes and uses [this web site](https://www.econdb.com/) to get the data. It uses [this web site](https://github.com/wcharczuk/go-chart) to produce the charts.

See [this web site](https://github.com/public-apis/public-apis#books) for an extensive list of web sites that offer public APIs.




