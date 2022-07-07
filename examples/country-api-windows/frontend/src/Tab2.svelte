<script>
    import { GetPopulations } from "../wailsjs/go/main/App.js";
    let populationsSorted = [];
    let highToLow = true;

    function getPopulations() {
        GetPopulations(highToLow).then(
            (result) => (populationsSorted = result)
        );
    }

    function reverseSortOrder() {
        highToLow = !highToLow;
        getPopulations();
    }

    getPopulations();

</script>

<div>
    <button class="btn" on:click={reverseSortOrder}>Reverse sort order</button>
    {#if highToLow}
        <h3>Population by Country -- High to Low</h3>
    {:else}
        <h3>Population by Country -- Low to High</h3>
    {/if}
    <ul>
        {#each populationsSorted as ps}
            <li>{ps}</li>
        {/each}
    </ul>
</div>

<style>
    ul {
        list-style-type: none;
    }
</style>
