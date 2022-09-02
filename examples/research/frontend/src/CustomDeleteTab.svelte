<script>
    import { DeleteCustom } from "../wailsjs/go/main/App.js";
    import { OpenCustomStmtFile } from "../wailsjs/go/main/App.js";

    let sqlStmt;
    let sqlStmts = [];
    let rows_affected = 0;

    function deleteCustom() {
        DeleteCustom(sqlStmt).then((result) => (rows_affected = result));
    }

    function openCustomStmtFile() {
        OpenCustomStmtFile().then((result) => (sqlStmts = result));
    }
</script>

<h3>Custom Delete Statments</h3>

<div class="warning">
    WARNING: THIS FEATURE IS FOR THOSE WHO HAVE THE KNOWLEDGE TO WRITE AND
    UNDERSTAND SQL DELETE STATMENTS. IMPROPER USE CAN RESULT IN THE UNINTENDED
    LOSS OF DATA.
</div>

<br />
<div>
    Enter a SQL Delete statement:
    <input
        class="statement"
        autocomplete="off"
        bind:value={sqlStmt}
        type="text"
    />
    &nbsp;&nbsp;
    <button class="btn" on:click={deleteCustom}>Submit</button>
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

Rows deleted: {rows_affected}

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
