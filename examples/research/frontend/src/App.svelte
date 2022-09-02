<script>
  import CreateTab from "./CreateTab.svelte";
  import QueryTab from "./QueryTab.svelte";
  import CustomQueryTab from "./CustomQueryTab.svelte";
  import CustomUpdateTab from "./CustomUpdateTab.svelte";
  import CustomDeleteTab from "./CustomDeleteTab.svelte";
  import BackupTab from "./BackupTab.svelte";
  import Tabs from "./Tabs.svelte";

  let items = [
    { label: "Create References", value: 1, component: CreateTab },
    { label: "Query, Update & Delete", value: 2, component: QueryTab },
    { label: "Custom Select", value: 3, component: CustomQueryTab },
    { label: "Custom Update", value: 4, component: CustomUpdateTab },
    { label: "Custom Delete", value: 5, component: CustomDeleteTab },
    { label: "Backup Database", value: 6, component: BackupTab },
  ];
</script>

<Tabs {items} />
