<script>
    import { BackupDb } from "../wailsjs/go/main/App.js";

    function backupDb() {
        BackupDb();
    }
</script>

<h3>Backup Database</h3>

<div><button class="btn" on:click={backupDb}>Backup Database</button></div>

<style>
    .btn:focus {
        border-width: 3px;
        border-color: blue;
    }
</style>