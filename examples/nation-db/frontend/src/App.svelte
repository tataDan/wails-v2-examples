<script>
  import { Connect } from "../wailsjs/go/main/App.js";
  import { Query } from "../wailsjs/go/main/App.js";

  let username;
  let password;
  let is_connected = false;
  let country_name;
  let countries = [];
  let query_type = "Like";

  function connect() {
    Connect(username, password).then((result) => (is_connected = result));
  }

  function query() {
    Query(query_type, country_name).then((result) => (countries = result));
  }
</script>

<main>
  <h3>NATION DATABASE</h3>
  <div class="input-box" id="input">
    <label for="username">username:</label>
    {#if is_connected}
      <input
        autocomplete="off"
        bind:value={username}
        class="input"
        id="username"
        type="text"
        disabled
      />
    {:else}
      <input
        autocomplete="off"
        bind:value={username}
        class="input"
        id="username"
        type="text"
      />
    {/if}
    &nbsp;&nbsp;
    <label for="password">password:</label>
    <input
      autocomplete="off"
      bind:value={password}
      class="input"
      id="password"
      type="password"
    />
    {#if is_connected}
      <button disabled class="btn" on:click={connect}>Connect</button>
    {:else}
      <button class="btn" on:click={connect}>Connect</button>
    {/if}
  </div>
  <div class="input-box">
    <br />
    <input
      type="radio"
      class="radio"
      id="like"
      name="query_type"
      value="Like"
      bind:group={query_type}
    />
    <label for="like">Match on starting characters</label>
    &nbsp;
    <input
      type="radio"
      class="radio"
      id="equals"
      name="query_type"
      value="Equals"
      bind:group={query_type}
    />
    <label for="equals">Exact match</label>
    &nbsp;&nbsp;&nbsp;
    <label for="country_name">Country:</label>
    <input
      autocomplete="off"
      bind:value={country_name}
      class="input"
      id="country_name"
      type="text"
    />
    &nbsp;
    {#if is_connected}
      <button class="btn" on:click={query}>Query</button>
    {:else}
      <button disabled class="btn" on:click={query}>Query</button>
    {/if}
    <br /><br />
  </div>
  <div>
    <table>
      <tr>
        <th>Country</th>
        <th>Area</th>
        <th>National Holiday</th>
      </tr>
      {#each countries as data}
        <tr>
          <td>{data.Name}</td>
          <td>{data.Area}</td>
          <td>{data.NationalDay}</td>
        </tr>
      {/each}
    </table>
  </div>
</main>

<style>
  .btn {
    margin-left: 10px;
    margin-right: 10px;
  }

  .btn:focus {
    border-width: 3px;
  }

  .radio:focus {
    outline: 1px solid white;
  }

  .input-box .input {
    border: none;
    border-radius: 3px;
    outline: none;
    height: 30px;
    line-height: 30px;
    padding: 0 10px;
    background-color: rgba(240, 240, 240, 1);
    -webkit-font-smoothing: antialiased;
  }

  .input-box .input:hover {
    border: none;
    background-color: rgba(255, 255, 255, 1);
  }

  .input-box .input:focus {
    border: none;
    background-color: rgba(255, 255, 255, 1);
  }

  table {
    width: 100%;
    background-color: lightblue;
    color: blue;
  }

  #country_name {
    width: 200px;
  }
</style>
