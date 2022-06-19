# wails-v2-examples

These examples use Wails (version 2) with the Svelte template. If you are not yet familiar with Wails, then please see https://wails.io/ and https://github.com/wailsapp/wails.

# Prerequisites

See https://wails.io/docs/gettingstarted/installation for instructions on installing wails, paying attention to the Supported Platforms, Dependencies, and Platform Specific Dependencies sections.

Then run go install github.com/wailsapp/wails/v2/cmd/wails@latest.

It is recommended to then run wails doctor to determine whether or not your system is ready.

# Running the examples

Typically you would simply run wails dev from the directory of the example that you want to run (e.g. examples/pizza-order).

However, these examples were built using Linux and if you attempt to run them on Windows you are likely to get an error. In this case, instead of simply running wails dev, do the following:

while in the frontend directory execute the following commands:

	npm install

	npm run build
	
and then while in the project directory execute the following commands:

	go build -tags dev -gcflags "all=-N -l"

	wails dev
	
These extra steps might not be needed in future releases of Wails. They might not be needed using MacOS. The version of Wails used to build these examples is v2.0.0-beta.37.

Please note that the nation-db example has special requirements in order for the example to run as intended.  Please see the README.md file in the examples/nation-db for further information.  The mimic-nation-db example does not have these requirements.


For a more advanced example, you might want to check out RiftShare (https://riftshare.app/ and https://github.com/achhabra2/riftshare/tree/v0.1.9).

