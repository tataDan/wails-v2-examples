<script>
    import { GetCountryNames } from "../wailsjs/go/main/App.js";
    import { GetCountryCode } from "../wailsjs/go/main/App.js";
    import { GetRGDP } from "../wailsjs/go/main/App";
    import { OneCountryLineChart } from "../wailsjs/go/main/App";
    import { OpenLineChart } from "../wailsjs/go/main/App";

    let countryNames = [];
    let selectedCountry = "";
    let countryCode;
    let rgdp = [];
    let showData = false;
    let dataFound = false;
    let viewer = "";
    let pathToLineChartFile = "";
    let showViewers = false;
    let infoLink = "";

    function init() {
        getCountryNames();
    }

    init();

    function getCountryNames() {
        GetCountryNames().then((result) => (countryNames = result));
    }

    function getCountryCode() {
        GetCountryCode(selectedCountry).then(
            (result) => (countryCode = result)
        );
        rgdp.length = 0;
        dataFound = false;
        showViewers = false;
    }

    function getRGDB() {
        showData = false;
        showViewers = false;
        GetRGDP(countryCode).then((result) => (rgdp = result));
        if (countryCode === "GB") {
            infoLink = "https://www.econdb.com/api/series/RGDP" + "UK" + "/";
        } else {
            infoLink =
                "https://www.econdb.com/api/series/RGDP" + countryCode + "/";
        }
        showData = true;
    }

    function oneCountryLineChart() {
        showData = false;
        OneCountryLineChart(selectedCountry, countryCode).then(
            (result) => (pathToLineChartFile = result)
        );
        showData = true;
        showViewers = true;
    }

    function openLineChart() {
        OpenLineChart(pathToLineChartFile, viewer);
        viewer = "";
    }

    // @ts-ignore
    window.runtime.EventsOn("data-found", function () {
        dataFound = true;
    });
</script>

<main>
    <h3>Country Real Gross Domestic Product</h3>

    &nbsp;&nbsp;&nbsp; Select a country
    <select bind:value={selectedCountry} on:change={getCountryCode}>
        {#each countryNames as cn}
            <option value={cn}>
                {cn}
            </option>
        {/each}
    </select>

    <button class="btn" on:click={getRGDB}
        >Get Real Domestic Product Data</button
    >
    <br />

    {#if dataFound}
        <br />
        Place this URL into your browser to see monetary unit used: {infoLink}
        <br /><br />
        <button class="btn" on:click={oneCountryLineChart}
            >Create line chart for this country</button
        >
        using last 10 values reported (typically quarters)
        <br />
    {/if}

    {#if showViewers}
        <br />
        Select a image viewer from this list
        <select bind:value={viewer} on:change={openLineChart}>
            <option value="shotwell">Shotwell</option>
            <option value="eog">Gnome Image Viewer</option>
            <option value="xnview">XnViewMP on Linux</option>
            <option value="firefox">Firefox</option>
            <option value="C:\Program Files\IrfanView\i_view64"
                >IrfanView on Windows</option
            >
            <option
                value="C:\Program Files (x86)\Microsoft\Edge\Application\msedge.exe"
                >Edge</option
            >
            <option value="xnviewmp">XnViewMP on Windows</option>
            <!-- <option value="phiewer">Phiewer on macOS</option> NOT TESTED -->
        </select>
        or enter the name of your viewer
        <input type="text" bind:value={viewer} on:change={openLineChart} />
    {/if}

    {#if showData}
        <h3>{selectedCountry} ({countryCode})</h3>
        <table>
            <tr>
                <th>Value</th>
                <th>Date</th>
                <th>Status</th>
            </tr>
            {#each rgdp as data}
                <tr>
                    <td>{data.Value}</td>
                    <td>{data.Date}</td>
                    <td>{data.Status}</td>
                </tr>
            {/each}
        </table>
    {/if}
</main>

<style>
    table {
        width: 100%;
        background-color: lightblue;
        color: blue;
    }

    .btn:focus {
        border-width: 3px;
    }
</style>
