<script>
    import { GetDensities } from "../wailsjs/go/main/App.js";
    let densitiesSorted = [];
    let highToLow = true;

    function getDensities() {
        GetDensities(highToLow).then((result) => (densitiesSorted = result));
    }

    function reverseSortOrder() {
        highToLow = !highToLow;
        getDensities();
    }

    getDensities();
</script>

<div>
    <button class="btn" on:click={reverseSortOrder}>Reverse sort order</button>
    {#if highToLow}
        <h3>Population Density by Country -- High to Low</h3>
    {:else}
        <h3>Population Density by Country -- Low to High</h3>
    {/if}
    (People per square kilometer)
    <ul>
        {#each densitiesSorted as ds}
            <li>{ds}</li>
        {/each}
    </ul>
</div>

<style>
    ul {
        list-style-type: none;
    }
</style>
