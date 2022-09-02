<script>
    import { CreateRef } from "../wailsjs/go/main/App.js";

    let ref = {};

    function createRef() {
        CreateRef(ref);
        ref = {};
    }
</script>

<h3>Create References</h3>

<div>
    <label for="category">Category:</label>
    <input
        autocomplete="off"
        bind:value={ref.category}
        class="input"
        id="category"
        type="text"
    />
    &nbsp;
    <label for="keywords">Keywords (separated by commas):</label>
    <input
        autocomplete="off"
        bind:value={ref.keywords}
        class="input"
        id="keywords"
        type="text"
    />
</div>
<br /><br />

<div>
    <label for="url">Url:</label>
    <input
        autocomplete="off"
        bind:value={ref.url}
        class="input"
        id="url"
        type="text"
    />
    &nbsp;
    <label for="book_title">Book Title:</label>
    <input
        autocomplete="off"
        bind:value={ref.book_title}
        class="input"
        id="book_title"
        type="text"
    />
    &nbsp;
    <label for="isbn">ISBN:</label>
    <input
        autocomplete="off"
        bind:value={ref.isbn}
        class="input"
        id="isbn"
        type="text"
    />
</div>

<p>Excerpt:</p>
<textarea bind:value={ref.excerpt} />
<br /><br />
<div><button class="btn" on:click={createRef}>Submit</button></div>

<style>
    .btn:focus {
        border-width: 3px;
        border-color: blue;
    }
    
    textarea {
        width: 97%;
        height: 400px;
    }
</style>
