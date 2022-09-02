<script>
    import { QueryCustom } from "../wailsjs/go/main/App.js";
    import { OpenCustomStmtFile } from "../wailsjs/go/main/App.js";

    let refs = [];
    let idx = 0;
    let sqlStmt = "";
    let sqlStmts = [];
    let showData = false;
    let showId = false;
    let showCategory = false;
    let showKeywords = false;
    let showDateCreated = false;
    let showUrl = false;
    let showBookTitle = false;
    let showIsbn = false;
    let showExcerpt = false;

    let showStmts = false;

    function changeShows(value) {
        showId = value;
        showCategory = value;
        showKeywords = value;
        showDateCreated = value;
        showUrl = value;
        showBookTitle = value;
        showIsbn = value;
        showExcerpt = value;
    }

    function queryCustom() {
        changeShows(false);
        const parts = sqlStmt.split(" ");
        for (let i in parts) {
            if (parts[i].toUpperCase() === "SELECT") {
                continue;
            }
            if (parts[i] === "*") {
                changeShows(true);
                break;
            }
            if (parts[i] === "id" || parts[i] === "id,") {
                showId = true;
            }
            if (parts[i] === "category" || parts[i] === "category,") {
                showCategory = true;
            }
            if (parts[i] === "keywords" || parts[i] === "keywords,") {
                showKeywords = true;
            }
            if (parts[i] === "date_created" || parts[i] === "date_created,") {
                showDateCreated = true;
            }
            if (parts[i] === "url" || parts[i] === "url,") {
                showUrl = true;
            }
            if (parts[i] === "book_title" || parts[i] === "book_title,") {
                showBookTitle = true;
            }
            if (parts[i] === "isbn" || parts[i] === "isbn,") {
                showIsbn = true;
            }
            if (parts[i] === "excerpt" || parts[i] === "excerpt,") {
                showExcerpt = true;
            }
            if (parts[i].toUpperCase() === "FROM") {
                break;
            }
        }
        idx = 0;
        QueryCustom(sqlStmt).then((result) => (refs = result));
        showData = true;
    }

    function openCustomStmtFile() {
        changeShows(false);
        OpenCustomStmtFile().then((result) => (sqlStmts = result));
        showStmts = true;
    }

    function next() {
        if (idx === refs.length - 1) {
            idx = 0;
        } else {
            idx++;
        }
    }

    function previous() {
        if (idx === 0) {
            idx = refs.length - 1;
        } else {
            idx--;
        }
    }
</script>

<h3>Custom Query Statments</h3>

<div class="notice">
    NOTICE: THIS FEATURE IS FOR THOSE WHO HAVE THE KNOWLEDGE TO WRITE AND
    UNDERSTAND SQL SELECT STATMENTS.
</div>
<br />
<div>
    Enter a SQL Select statement:
    <input
        class="statement"
        autocomplete="off"
        bind:value={sqlStmt}
        type="text"
    />
    &nbsp;&nbsp;
    <button class="btn" on:click={queryCustom}>Submit</button>
    <br /><br />
    OR &nbsp;&nbsp; &nbsp;&nbsp;
    <button class="btn" on:click={openCustomStmtFile}
        >Open file with custom SQL statements</button
    >
    &nbsp;&nbsp;&nbsp; Choose a SQL statement:&nbsp;
    <select bind:value={sqlStmt}>
        {#each sqlStmts as stmt}
            <option value={stmt}>
                {stmt}
            </option>
        {/each}
    </select>
</div>

<br />
<button class="btn" on:click={previous}>Previous</button>
<button class="btn" on:click={next}>Next</button>
<br /><br />

<div class="output">
    {#if showData}
        {#each refs as ref, index}
            {#if index === idx}
                <div class="output">
                    {#if showId}
                        ID:
                        <input disabled bind:value={ref.Id} type="text" />
                    {/if}
                    {#if showDateCreated}
                        &nbsp;&nbsp;Date Created:
                        <input
                            disabled
                            bind:value={ref.Date_created}
                            type="text"
                        />
                        <br /><br />
                    {/if}
                    {#if showCategory}
                        Category:
                        <input disabled bind:value={ref.Category} type="text" />
                    {/if}
                    {#if showKeywords}
                        &nbsp;&nbsp; Keywords:
                        <input disabled bind:value={ref.Keywords} type="text" />
                        <br /><br />
                    {/if}
                    {#if showUrl}
                        Url:
                        <input disabled bind:value={ref.Url} type="text" />
                    {/if}
                    {#if showBookTitle}
                        &nbsp;&nbsp; Book title:
                        <input
                            disabled
                            bind:value={ref.Book_title}
                            type="text"
                        />
                    {/if}
                    {#if showIsbn}
                        &nbsp;&nbsp; ISBN:
                        <input disabled bind:value={ref.Isbn} type="text" />
                    {/if}
                    {#if showExcerpt}
                        <br /><br />
                        Excerpt:
                        <textarea disabled bind:value={ref.Excerpt} />
                    {/if}
                </div>
            {/if}
        {/each}
    {/if}
</div>

<style>
   select:focus {
        border-width: 3px;
        border-color: blue;
    }

    .btn:focus {
        border-width: 3px;
        border-color: blue;
    }
    
    .notice {
        border-style: solid;
        border-color: blue;
        text-align: justify;
        text-justify: auto;
        padding: 10px;
    }

    .statement {
        width: 500px;
    }

    textarea {
        width: 97%;
        height: 200px;
    }
</style>
