<script>
    import { LoadProgramData } from "../wailsjs/go/main/App.js";
    import { GetProgramTitles } from "../wailsjs/go/main/App.js";
    import { GetProgramCode } from "../wailsjs/go/main/App.js";
    import { SearchProgramTitles } from "../wailsjs/go/main/App.js";
    import { LoadStateData } from "../wailsjs/go/main/App.js";
    import { GetStateNames } from "../wailsjs/go/main/App.js";
    import { GetStateCode } from "../wailsjs/go/main/App.js";
    import { GetData } from "../wailsjs/go/main/App.js";

    let programTitles = [];
    let selectedTitle = "";
    let matchingTitles = [];
    let searchTitle = "";
    let stateNames = [];
    let selectedStateName = "";
    let city = "";
    let programCode = "";
    let stateCode = "";
    let studentSizeMax = "";
    let inStateTuitionMax = "";
    let data = [];
    let showData = false;
    let showMatchingTitles = false;

    function init() {
        loadProgramData();
        loadStateData();
    }

    init();

    async function loadProgramData() {
        await LoadProgramData();
        GetProgramTitles().then((result) => (programTitles = result));
    }

    function getProgramCode() {
        GetProgramCode(selectedTitle).then((result) => (programCode = result));
    }

    function searchProgramTitles() {
        showMatchingTitles = true;
        showData = false;
        SearchProgramTitles(searchTitle).then(
            (result) => (matchingTitles = result)
        );
    }

    async function loadStateData() {
        await LoadStateData();
        GetStateNames().then((result) => (stateNames = result));
    }

    function getStateCode() {
        GetStateCode(selectedStateName).then((result) => (stateCode = result));
    }

    async function getData() {
        showData = false;
        showMatchingTitles = false;

        await GetData(city, stateCode, programCode).then(
            (result) => (data = result)
        );
        showData = true;
    }
</script>

<h3>Colleges API</h3>

<div class="input">
    <div class="child-state">
        <div class="child" id="city">
            City:
            <input type="text" bind:value={city} />
        </div>
        <div class="child" id="state">
            State:
            <select bind:value={selectedStateName} on:change={getStateCode}>
                {#each stateNames as name}
                    <option value={name}>
                        {name}
                    </option>
                {/each}
            </select>
            {stateCode}
        </div>
    </div>

    <div class="programs">
        Programs:
        <select bind:value={selectedTitle} on:change={getProgramCode}>
            {#each programTitles as title}
                <option value={title}>
                    {title}
                </option>
            {/each}
        </select>
        {programCode}
    </div>

    <div>
        <button class="button" on:click={getData}>Get College Data</button>
        &nbsp;&nbsp; Text to search for in program titles:
        <input type="text" bind:value={searchTitle} />
        <button class="button" on:click={searchProgramTitles}
            >Search Program Titles</button
        >
    </div>

    <div class="filters">
        Filters: &nbsp;&nbsp; Undergraduate student size max:
        <input type="text" bind:value={studentSizeMax} /> &nbsp; In-state
        tuition max:
        <input type="text" bind:value={inStateTuitionMax} />
    </div>
</div>

<div class="result" id="result">
    {#if showData}
        {#if data.length === 0}
            <h3>No data found!</h3>
        {:else}
            {#each data as d}
                {#if (d.StudentSize <= studentSizeMax || studentSizeMax.length === 0) && (d.TuitionInState <= inStateTuitionMax || inStateTuitionMax.length === 0)}
                    {d.Name}<br />
                    {d.City},&nbsp;{d.State}<br />
                    Number of undergraduate students: {d.StudentSize}
                    {#if d.StudentSize === 0}(or unknown){/if}
                    <br />
                    Number of graduate students: {d.GradStudents}
                    {#if d.GradStudents === 0}(or unknown){/if}
                    <br />
                    In-state tuition: {d.TuitionInState}
                    {#if d.TuitionInState === 0} (or unknown){/if}
                    <br />
                    Out-of-state tuition: {d.TuitionOutOfState}
                    {#if d.TuitionOutOfState === 0} (or unknown) {/if}
                    <br />
                    School url: {d.SchoolUrl}<br />
                    Price calculator url: {d.PriceCalculatoUrl}<br />
                    <br />
                {/if}
            {/each}
        {/if}
    {/if}

    {#if showMatchingTitles}
        <h3>Matching Titles</h3>
        {#each matchingTitles as mt}
            {mt}<br />
        {/each}
    {/if}
</div>

<style>
    input {
        height: 1.5rem;
        font-size: small;
    }

    select {
        height: 2rem;
        font-size: small;
    }

    .input {
        text-align: left;
        padding-left: 2rem;
    }

    .child {
        display: inline-block;
    }

    #city {
        padding-right: 1rem;
    }

    .programs {
        margin: 0.5rem auto;
    }

    .filters {
        margin: 0.5rem auto;
        padding: 8px;
        border: 2px solid black;
        border-width: 2px;
    }

    .result {
        margin: 1.5rem auto;
    }

    .button {
        border-width: 3px;
    }
</style>
