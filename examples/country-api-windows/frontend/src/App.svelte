<script>
  import Tab1 from "./Tab1.svelte";
  import Tab2 from "./Tab2.svelte";
  import Tab3 from "./Tab3.svelte";
  import Tab4 from "./Tab4.svelte";
  import Tabs from "./Tabs.svelte";

  let items = [
    { label: "Country", value: 1, component: Tab1 },
    { label: "Population", value: 2, component: Tab2 },
    { label: "Area", value: 3, component: Tab3 },
    { label: "Population Density", value: 4, component: Tab4 },
  ];
</script>

<Tabs {items} />
