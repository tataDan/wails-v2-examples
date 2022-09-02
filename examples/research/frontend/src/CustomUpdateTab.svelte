<script>
    import { UpdateCustom } from "../wailsjs/go/main/App.js";
    import { OpenCustomStmtFile } from "../wailsjs/go/main/App.js";

    let sqlStmt;
    let sqlStmts = [];
    let rows_affected = 0;

    function updateCustom() {
        UpdateCustom(sqlStmt).then((result) => (rows_affected = result));
    }

    function openCustomStmtFile() {
        OpenCustomStmtFile().then((result) => (sqlStmts = result));
    }
</script>

<h3>Custom Update Statments</h3>

<div class="warning">
    WARNING: THIS FEATURE IS FOR THOSE WHO HAVE THE KNOWLEDGE TO WRITE AND
    UNDERSTAND SQL UPDATE STATMENTS. IMPROPER USE CAN RESULT IN THE UNINTENDED
    MODIFICATION OF DATA.
</div>

<br />
<div>
    Enter a SQL Update statement:
    <input
        class="statement"
        autocomplete="off"
        bind:value={sqlStmt}
        type="text"
    />
    &nbsp;&nbsp;
    <button class="btn" on:click={updateCustom}>Submit</button>
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

Rows updated: {rows_affected}

<style>
  select:focus {
        border-width: 3px;
        border-color: blue;
    }

    .btn:focus {
        border-width: 3px;
        border-color: blue;
    }
    
    .warning {
        border-style: solid;
        border-color: blue;
        text-align: justify;
        text-justify: auto;
        padding: 10px;
    }

    .statement {
        width: 500px;
    }
</style>
