<script>
    import { GetIds } from "../wailsjs/go/main/App.js";
    import { GetCategories } from "../wailsjs/go/main/App.js";
    import { GetKeywords } from "../wailsjs/go/main/App.js";
    import { GetUrls } from "../wailsjs/go/main/App.js";
    import { GetBookTitles } from "../wailsjs/go/main/App.js";
    import { GetIsbns } from "../wailsjs/go/main/App.js";
    import { QueryAll } from "../wailsjs/go/main/App.js";
    import { QueryById } from "../wailsjs/go/main/App.js";
    import { QueryByCategory } from "../wailsjs/go/main/App.js";
    import { QueryByKeywords } from "../wailsjs/go/main/App.js";
    import { QueryByDate } from "../wailsjs/go/main/App.js";
    import { QueryByUrl } from "../wailsjs/go/main/App.js";
    import { QueryByBookTitle } from "../wailsjs/go/main/App.js";
    import { QueryByIsbn } from "../wailsjs/go/main/App.js";
    import { Update } from "../wailsjs/go/main/App.js";
    import { Delete } from "../wailsjs/go/main/App.js";

    let ref = {};
    let refs = [];
    let idx = 0;
    let ids = [];
    let categories = [];
    let keywords = [];
    let urls = [];
    let bookTitles = [];
    let isbns = [];
    let selectedId;
    let selectedCategory;
    let selectedKeywords;
    let selectedStartDate;
    let selectedEndDate;
    let selectedUrl;
    let selectedBookTitle;
    let selectedIsbn;
    let selectedColumn;
    let selectedQueryType;
    let newValue;
    let updated = true;
    let deleted = false;
    let showRef = false;
    let showRefs = false;
    let showAllQuery = false;
    let showIdQuery = false;
    let showCategoryQuery = false;
    let showKeywordsQuery = false;
    let showDateQuery = false;
    let showUrlQuery = false;
    let showBookTitleQuery = false;
    let showIsbnQuery = false;
    let showQueryType = true;

    function reset() {
        selectedColumn = "";
        showRef = false;
        showRefs = false;
        showAllQuery = false;
        showIdQuery = false;
        showCategoryQuery = false;
        showKeywordsQuery = false;
        showDateQuery = false;
        showUrlQuery = false;
        showBookTitleQuery = false;
        showIsbnQuery = false;
        showQueryType = false;
    }

    function newQueryType() {
        reset();
        selectedColumn = "category";
        showQueryType = true;
    }

    function chooseQueryType() {
        if (selectedQueryType === "byAll") {
            showAllQuery = true;
        } else if (selectedQueryType === "byId") {
            getIds();
            showIdQuery = true;
        } else if (selectedQueryType === "byCategory") {
            getCategories();
            showCategoryQuery = true;
        } else if (selectedQueryType === "byKeywords") {
            getKeywords();
            showKeywordsQuery = true;
        } else if (selectedQueryType === "byDate") {
            showDateQuery = true;
        } else if (selectedQueryType === "byUrl") {
            getUrls();
            showUrlQuery = true;
        } else if (selectedQueryType === "byBookTitle") {
            getBookTitles();
            showBookTitleQuery = true;
        } else if (selectedQueryType === "byIsbn") {
            getIsbns();
            showIsbnQuery = true;
        }
        showQueryType = false;
    }

    function getIds() {
        GetIds().then((result) => (ids = result));
        selectedId = "";
    }

    function getCategories() {
        GetCategories().then((result) => (categories = result));
        selectedCategory = "";
    }

    function getKeywords() {
        GetKeywords().then((result) => (keywords = result));
        selectedKeywords = "";
    }

    function getUrls() {
        GetUrls().then((result) => (urls = result));
        selectedUrl = "";
    }

    function getBookTitles() {
        GetBookTitles().then((result) => (bookTitles = result));
        selectedBookTitle = "";
    }

    function getIsbns() {
        GetIsbns().then((result) => (isbns = result));
        selectedIsbn = "";
    }

    function queryAll() {
        idx = 0;
        QueryAll().then((result) => (refs = result));
        showRefs = true;
        updated = true;
        deleted = false;
    }

    function queryById() {
        QueryById(selectedId).then((result) => (ref = result));
        showRef = true;
        updated = true;
        deleted = false;
    }

    function queryByCategory() {
        idx = 0;
        QueryByCategory(selectedCategory).then((result) => (refs = result));
        showRefs = true;
        updated = true;
        deleted = false;
    }

    function queryByKeywords() {
        idx = 0;
        QueryByKeywords(selectedKeywords).then((result) => (refs = result));
        showRefs = true;
        updated = true;
        deleted = false;
    }

    function queryByDate() {
        idx = 0;
        QueryByDate(selectedStartDate, selectedEndDate).then(
            (result) => (refs = result)
        );
        showRefs = true;
        updated = true;
        deleted = false;
    }

    function queryByUrl() {
        idx = 0;
        QueryByUrl(selectedUrl).then((result) => (refs = result));
        showRefs = true;
        updated = true;
        deleted = false;
    }

    function queryByBookTitle() {
        idx = 0;
        QueryByBookTitle(selectedBookTitle).then((result) => (refs = result));
        showRefs = true;
        updated = true;
        deleted = false;
    }

    function queryByIsbn() {
        idx = 0;
        QueryByIsbn(selectedIsbn).then((result) => (refs = result));
        showRefs = true;
        updated = true;
        deleted = false;
    }

    function update() {
        if (selectedQueryType === "byId") {
            switch (selectedColumn) {
                case "category":
                    newValue = ref.Category;
                    break;
                case "keywords":
                    newValue = ref.Keywords;
                    break;
                case "url":
                    newValue = ref.Url;
                    break;
                case "book_title":
                    newValue = ref.Book_title;
                    break;
                case "isbn":
                    newValue = ref.Isbn;
                    break;
                case "excerpt":
                    newValue = ref.Excerpt;
                    break;
            }
        } else {
            switch (selectedColumn) {
                case "category":
                    newValue = refs[idx].Category;
                    break;
                case "keywords":
                    newValue = refs[idx].Keywords;
                    break;
                case "url":
                    newValue = refs[idx].Url;
                    break;
                case "book_title":
                    newValue = refs[idx].Book_title;
                    break;
                case "isbn":
                    newValue = refs[idx].Isbn;
                    break;
                case "excerpt":
                    newValue = refs[idx].Excerpt;
                    break;
            }
        }

        if (selectedQueryType === "byId") {
            Update(ref.Id.toString(), selectedColumn, newValue).then(
                (result) => (updated = result)
            );
        } else {
            Update(refs[idx].Id.toString(), selectedColumn, newValue).then(
                (result) => (updated = result)
            );
        }
    }

    function deleteRef() {
        if (selectedQueryType === "byId") {
            Delete(ref.Id.toString()).then((result) => (deleted = result));
        } else {
            Delete(refs[idx].Id.toString()).then(
                (result) => (deleted = result)
            );
        }
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

{#if showQueryType}
    <h3>Query, Update & Delete</h3>
    &nbsp;&nbsp;&nbsp; Choose a query type:&nbsp;
    <select bind:value={selectedQueryType}>
        <option value={"byAll"}>
            {"byAll"}
        </option>
        <option value={"byId"}>
            {"byId"}
        </option>
        <option value={"byCategory"}>
            {"byCategory"}
        </option>
        <option value={"byKeywords"}>
            {"byKeywords"}
        </option>
        <option value={"byDate"}>
            {"byDate"}
        </option>
        <option value={"byUrl"}>
            {"byUrl"}
        </option>
        <option value={"byBookTitle"}>
            {"byBookTitle"}
        </option>
        <option value={"byIsbn"}>
            {"byIsbn"}
        </option>
    </select>
    &nbsp;&nbsp;&nbsp;
    <button class="btn" on:click={chooseQueryType}>Show Query</button>
{/if}

{#if showAllQuery}
    <h3>Query for All References</h3>
    <button class="btn" on:click={queryAll}>Query for all references</button>
    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
    <button class="btn" on:click={newQueryType}>New Query Type</button>
{/if}

{#if showIdQuery}
    <h3>Query by Id</h3>
    &nbsp;&nbsp;&nbsp; Choose a Id:&nbsp;
    <select bind:value={selectedId} on:change={queryById}>
        {#each ids as id}
            <option value={id}>
                {id}
            </option>
        {/each}
    </select>
    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
    <button class="btn" on:click={newQueryType}>New Query Type</button>
    <br /><br />
{/if}

{#if showCategoryQuery}
    <h3>Query By Category</h3>
    &nbsp;&nbsp;&nbsp; Click on down arrow to select a category:&nbsp;
    <select bind:value={selectedCategory} on:change={queryByCategory}>
        {#each categories as category}
            <option value={category}>
                {category}
            </option>
        {/each}
    </select>
    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
    <button class="btn" on:click={newQueryType}>New Query Type</button>
{/if}

{#if showKeywordsQuery}
    <h3>Query By Keywords</h3>
    &nbsp;&nbsp;&nbsp; Click on down arrow to select a keyword:&nbsp;
    <select bind:value={selectedKeywords} on:change={queryByKeywords}>
        {#each keywords as keyword}
            <option value={keyword}>
                {keyword}
            </option>
        {/each}
    </select>
    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
    <button class="btn" on:click={newQueryType}>New Query Type</button>
{/if}

{#if showDateQuery}
    <h3>Query By Date</h3>

    <label for="start-date">Start date (eg 2022-08-13):</label>
    <input
        autocomplete="off"
        bind:value={selectedStartDate}
        on:change={queryByDate}
        id="start-date"
        type="text"
    />
    &nbsp;&nbsp;
    <label for="end-date">End date (eg 2022-09-27):</label>
    <input
        autocomplete="off"
        bind:value={selectedEndDate}
        id="end-date"
        type="text"
    />
    &nbsp;
    <button class="btn" on:click={queryByDate}>Query By Date</button>
    <br /><br />
    <button class="btn" on:click={newQueryType}>New Query Type</button>
{/if}

{#if showUrlQuery}
    <h3>Query By Url</h3>
    &nbsp;&nbsp;&nbsp; Click on down arrow to select a url:&nbsp;
    <select bind:value={selectedUrl} on:change={queryByUrl}>
        {#each urls as url}
            <option value={url}>
                {url}
            </option>
        {/each}
    </select>
    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
    <button class="btn" on:click={newQueryType}>New Query Type</button>
{/if}

{#if showBookTitleQuery}
    <h3>Query By Book Title</h3>
    &nbsp;&nbsp;&nbsp; Click on down arrow to select a book title:&nbsp;
    <select bind:value={selectedBookTitle} on:change={queryByBookTitle}>
        {#each bookTitles as bookTitle}
            <option value={bookTitle}>
                {bookTitle}
            </option>
        {/each}
    </select>
    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
    <button class="btn" on:click={newQueryType}>New Query Type</button>
{/if}

{#if showIsbnQuery}
    <h3>Query By ISBN</h3>
    &nbsp;&nbsp;&nbsp; Click on down arrow to select a ISBN:&nbsp;
    <select bind:value={selectedIsbn} on:change={queryByIsbn}>
        {#each isbns as isbn}
            <option value={isbn}>
                {isbn}
            </option>
        {/each}
    </select>
    &nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;
    <button class="btn" on:click={newQueryType}>New Query Type</button>
{/if}

{#if showRef}
    <div class="output">
        ID:
        <input
            disabled
            autocomplete="off"
            bind:value={ref.Id}
            class="input"
            type="text"
        />
        &nbsp;&nbsp; Date Created:
        <input
            disabled
            autocomplete="off"
            bind:value={ref.Date_created}
            class="input"
            type="text"
        />
        <br /><br />
        Category:
        {#if selectedColumn === "category"}
            <input
                autocomplete="off"
                bind:value={ref.Category}
                class="input"
                type="text"
            />
        {:else}
            <input
                disabled
                autocomplete="off"
                bind:value={ref.Category}
                class="input"
                type="text"
            />
        {/if}
        &nbsp;&nbsp; Keywords (separted by commas):
        {#if selectedColumn === "keywords"}
            <input
                autocomplete="off"
                bind:value={ref.Keywords}
                class="input"
                type="text"
            />
        {:else}
            <input
                disabled
                autocomplete="off"
                bind:value={ref.Keywords}
                class="input"
                type="text"
            />
        {/if}
        <br /><br />
        Url:
        {#if selectedColumn === "url"}
            <input
                autocomplete="off"
                bind:value={ref.Url}
                class="input"
                type="text"
            />
        {:else}
            <input
                disabled
                autocomplete="off"
                bind:value={ref.Url}
                class="input"
                type="text"
            />
        {/if}
        &nbsp;&nbsp; Book title:
        {#if selectedColumn === "book_title"}
            <input
                autocomplete="off"
                bind:value={ref.Book_title}
                class="input"
                type="text"
            />
        {:else}
            <input
                disabled
                autocomplete="off"
                bind:value={ref.Book_title}
                class="input"
                type="text"
            />
        {/if}
        &nbsp;&nbsp; ISBN:
        {#if selectedColumn === "isbn"}
            <input
                autocomplete="off"
                bind:value={ref.Isbn}
                class="input"
                type="text"
            />
        {:else}
            <input
                disabled
                autocomplete="off"
                bind:value={ref.Isbn}
                class="input"
                type="text"
            />
        {/if}
        <br /><br />
        Excerpt:
        {#if selectedColumn === "excerpt"}
            <textarea bind:value={ref.Excerpt} />
        {:else}
            <textarea disabled bind:value={ref.Excerpt} />
        {/if}
    </div>
{/if}

{#if showRefs}
    <br /><br />
    {#if updated && !deleted}
        <button class="btn" on:click={previous}>Previous</button>
    {:else}
        <button disabled class="btn" on:click={previous}>Previous</button>
    {/if}
    &nbsp;&nbsp;
    {#if updated && !deleted}
        <button class="btn" on:click={next}>Next</button>
    {:else}
        <button disabled class="btn" on:click={next}>Next</button>
    {/if}
    <br /><br />
    {#each refs as ref, index}
        {#if index === idx}
            <div class="output">
                ID:
                <input
                    disabled
                    bind:value={ref.Id}
                    class="disabled"
                    type="text"
                />
                &nbsp;&nbsp; Date Created:
                <input
                    disabled
                    bind:value={ref.Date_created}
                    class="disabled"
                    type="text"
                />
                <br /><br />
                Category:
                {#if selectedColumn === "category"}
                    <input
                        autocomplete="off"
                        bind:value={ref.Category}
                        class="edit"
                        type="text"
                    />
                {:else}
                    <input
                        disabled
                        bind:value={ref.Category}
                        class="disabled"
                        type="text"
                    />
                {/if}
                &nbsp;&nbsp; Keywords (separted by commas):
                {#if selectedColumn === "keywords"}
                    <input
                        autocomplete="off"
                        bind:value={ref.Keywords}
                        class="edit"
                        type="text"
                    />
                {:else}
                    <input
                        disabled
                        bind:value={ref.Keywords}
                        class="disabled"
                        type="text"
                    />
                {/if}
                <br /><br />
                Url:
                {#if selectedColumn === "url"}
                    <input
                        autocomplete="off"
                        bind:value={ref.Url}
                        class="edit"
                        type="text"
                    />
                {:else}
                    <input
                        disabled
                        bind:value={ref.Url}
                        class="disabled"
                        type="text"
                    />
                {/if}
                &nbsp;&nbsp; Book title:
                {#if selectedColumn === "book_title"}
                    <input
                        autocomplete="off"
                        bind:value={ref.Book_title}
                        class="edit"
                        type="text"
                    />
                {:else}
                    <input
                        disabled
                        bind:value={ref.Book_title}
                        class="disabled"
                        type="text"
                    />
                {/if}
                &nbsp;&nbsp; ISBN:
                {#if selectedColumn === "isbn"}
                    <input
                        autocomplete="off"
                        bind:value={ref.Isbn}
                        class="edit"
                        type="text"
                    />
                {:else}
                    <input
                        disabled
                        bind:value={ref.Isbn}
                        class="disabled"
                        type="text"
                    />
                {/if}
                <br /><br />
                Excerpt:
                {#if selectedColumn === "excerpt"}
                    <textarea bind:value={ref.Excerpt} class="edit" />
                {:else}
                    <textarea
                        disabled
                        bind:value={ref.Excerpt}
                        class="disabled"
                    />
                {/if}
            </div>
        {/if}
    {/each}
{/if}

{#if showRef || showRefs}
    {#if updated}
        &nbsp;&nbsp;&nbsp; Choose a column to update:&nbsp;
        <select bind:value={selectedColumn}>
            <option value={"category"}>
                {"category"}
            </option>
            <option value={"keywords"}>
                {"keywords"}
            </option>
            <option value={"url"}>
                {"url"}
            </option>
            <option value={"book_title"}>
                {"book_title"}
            </option>
            <option value={"isbn"}>
                {"isbn"}
            </option>
            <option value={"excerpt"}>
                {"excerpt"}
            </option>
        </select>
    {:else}
        &nbsp;&nbsp;&nbsp; Choose a column to update:&nbsp;
        <select disabled bind:value={selectedColumn}>
            <option value={"category"}>
                {"category"}
            </option>
            <option value={"keywords"}>
                {"keywords"}
            </option>
            <option value={"url"}>
                {"url"}
            </option>
            <option value={"book_title"}>
                {"book_title"}
            </option>
            <option value={"isbn"}>
                {"isbn"}
            </option>
            <option value={"excerpt"}>
                {"excerpt"}
            </option>
        </select>
    {/if}
    {#if !deleted}
        &nbsp;&nbsp;&nbsp;&nbsp;
        <button class="btn" on:click={update}>Update the reference</button>
    {:else}
        <button disabled class="btn" on:click={update}
            >Update the reference</button
        >
    {/if}
    &nbsp;&nbsp;&nbsp;&nbsp;
    <button class="btn" on:click={deleteRef}>Delete this reference</button>
    <br />
    {#if !updated}
        <span class="updated-msg"
            >The highlighed value might be out of date.</span
        >
    {/if}
    <br />
    {#if deleted}
        <span class="deleted-msg">This record has been deleted.</span>
    {/if}
{/if}

<style>
    select:focus {
        border-width: 3px;
        border-color: blue;
    }

    .btn:focus {
        border-width: 3px;
        border-color: blue;
    }

    textarea {
        width: 97%;
        height: 200px;
    }

    .edit {
        background-color: aquamarine;
    }

    .updated-msg {
        color: red;
        font-size: large;
    }

    .deleted-msg {
        color: red;
        font-size: large;
    }
</style>
