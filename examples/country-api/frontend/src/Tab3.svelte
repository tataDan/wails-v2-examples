<script>
    import { GetAreas } from "../wailsjs/go/main/App.js";
    let areasSorted = [];
    let highToLow = true;

    function getAreas() {
        GetAreas(highToLow).then((result) => (areasSorted = result));
    }

    function reverseSortOrder() {
        highToLow = !highToLow;
        getAreas();
    }

    getAreas();
</script>

<div>
    <button class="btn" on:click={reverseSortOrder}>Reverse sort order</button>
    {#if highToLow}
        <h3>Area by Country -- High to Low</h3>
    {:else}
        <h3>Area by Country -- Low to High</h3>
    {/if}
    <ul>
        {#each areasSorted as as}
            <li>{as}</li>
        {/each}
    </ul>
</div>

<style>
    ul {
        list-style-type: none;
    }
</style>
